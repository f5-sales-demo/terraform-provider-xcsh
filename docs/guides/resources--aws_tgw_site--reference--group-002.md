---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-6c62965235801366db44f1e8c72dd5638acedb237cfbb8a589513856021fa915"></a>

## aws_parameters.new_vpc — aws_parameters.new_vpc / 148823637b8e / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.new_vpc

<a id="canonical-74ee03345ff222625074b7fec80e3ec8e5cc0917698c667e6651e10d1ff0131b"></a>

Type: `"object"`. single nested block, Optional.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4"),
  validators.ConflictingObjectAttributes("autogenerate",
    "name_tag")}
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
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name_tag\"]"
}
```

Terraform syntax:

```terraform
new_vpc {
  # Configure direct properties listed below.
}
```

<a id="canonical-f57e885e45514b1a4262efbdb3d1b6e1b9c47edee502e56c0a5c360b266dbab3"></a>

## Direct properties — aws_parameters.new_vpc / 148823637b8e / 3

- [autogenerate](resources--aws_tgw_site--reference--group-002.md#canonical-62992f939ce14111cc35e7c3235acda84d3b0d598a86ade2d5b94fe17ea1c63b): complete subsection reference.

<a id="canonical-66e25dd9807a1eef7f8e9be8438e6ad1eb50400ed454ce8abd5e1093cb481fc8"></a>

<a id="canonical-a830c1968719606146b4c15bbe181d78737968286d31b683c462b30c1b7438fb"></a>

## name_tag property — aws_parameters.new_vpc / 148823637b8e / 4

Type: `"string"`. Optional.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

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

<a id="canonical-229233bf45e0c11fed1d4fd8848a3773a11b524811271a09362024178aa4c2ac"></a>

<a id="canonical-113321f349977222375aa24c7d026782e7c78a09864f5b8bfe7fb49a16bc1364"></a>

## primary_ipv4 property — aws_parameters.new_vpc / 148823637b8e / 5

Type: `"string"`. Optional.

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

Upstream description:

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

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
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  }
}
```

<a id="canonical-315f86aa354a1e5949a074b23e27280902315c3251892473e01d92e3aca6d466"></a>

## Next pages — aws_parameters.new_vpc / 148823637b8e / 6

- [aws_parameters.new_vpc.autogenerate](resources--aws_tgw_site--reference--group-002.md#canonical-62992f939ce14111cc35e7c3235acda84d3b0d598a86ade2d5b94fe17ea1c63b)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-62992f939ce14111cc35e7c3235acda84d3b0d598a86ade2d5b94fe17ea1c63b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2a93bd3170eddc1b77c6edb05f0ced4c76b66f272c003e0ab3e0d2b050b9718"></a>

## aws_parameters.new_vpc.autogenerate — aws_parameters.new_vpc.autogenerate / 26e714d10c6a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-001.md#canonical-2823b6f937a1792c17d1a89fa7317dd7838c44ea0e8ac666805abf0d98ea773a)
- aws_parameters.new_vpc.autogenerate

<a id="canonical-4cee8ea06a44d4415191ee85a94ed81ddcd51e9e999c65b5ed80cc2377a13747"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for autogenerate.

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
autogenerate = {}
```

<a id="canonical-f83dcbdda9f93d721ee46f592780b781f73742ccf815f2f8b8ac20039536ec1c"></a>

## Direct properties — aws_parameters.new_vpc.autogenerate / 26e714d10c6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62b57b5efcaf7de6c568c738adb2513ca1d467408c8db67220d6a2bff8b07040"></a>

## Next pages — aws_parameters.new_vpc.autogenerate / 26e714d10c6a / 4

- [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-001.md#canonical-2823b6f937a1792c17d1a89fa7317dd7838c44ea0e8ac666805abf0d98ea773a)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-a989844489c66d4e7faae9b8073925c332423fe87f9ed9161f123294f9f2b651"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73ee13e992949bdad0a72b3159db9d46c049f3c2540130ed8a2f145b7775ce8c"></a>

## aws_parameters.no_worker_nodes — aws_parameters.no_worker_nodes / 285c84b5508b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.no_worker_nodes

<a id="canonical-bc4bd9d19ebed9455f15464366fb78b8104af9a49d586a5aae92740fa9ef8cfa"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no worker nodes.

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
no_worker_nodes = {}
```

<a id="canonical-d3013179e6b621405f46775e2af0169a184974e5b0e668c2ce9db642b212eaa0"></a>

## Direct properties — aws_parameters.no_worker_nodes / 285c84b5508b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d826e61ad3d975909ba55b08a0f81f7179deab040bb9dbf537c965409d84d13"></a>

## Next pages — aws_parameters.no_worker_nodes / 285c84b5508b / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-d99f8f19c6a7e99ee542ccfb2979873c26ae54bd8d9b7fa065b15906cbe89931"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e12f0930912c574605c7c91ec7b103d9f6ceaab96012f8da3e1440698be692a4"></a>

## aws_parameters.reserved_tgw_cidr — aws_parameters.reserved_tgw_cidr / fc49c8707470 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.reserved_tgw_cidr

<a id="canonical-217df55ab72f020f9460ba67a6113a63d170f28bfba9445e25e61a582b449b8a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved tgw cidr.

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
reserved_tgw_cidr = {}
```

<a id="canonical-94d6a32975b577361a6e1a28b62710fb085193a762ca180151d0781db5f9fe38"></a>

## Direct properties — aws_parameters.reserved_tgw_cidr / fc49c8707470 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0d7b4f3b0fb23bc5e79258ba5dfff86733997d4d4e97ce0813b20b14c6f0b328"></a>

## Next pages — aws_parameters.reserved_tgw_cidr / fc49c8707470 / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-dda6f57721acd4eeedf2a9b8e0df09e5fe75a64f91d5fb5515a8f7664c2a2e98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68506e39d1cd335f03689f324d34efd157eaf52c932d45257548589cb71bb542"></a>

## aws_parameters.tgw_cidr — aws_parameters.tgw_cidr / eb77418e5da4 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.tgw_cidr

<a id="canonical-4071d78aefea4a0553ebeb7b73e00ac04a701f5a01673bef668dc78a1cc2bd14"></a>

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
tgw_cidr {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed7d42360473c8380f99d2e8cd0d4ca652e59d4909fc6b3e2e76b987ece665e3"></a>

## Direct properties — aws_parameters.tgw_cidr / eb77418e5da4 / 3

<a id="canonical-df5d082102250f01d4b558812d0b3af26b59bf9ca09b76fa72e29468c81b3e8b"></a>

<a id="canonical-b57ecb5968723f14cb9ee8eea262ed5cf7cb8eb198ff4f2b9ffdac4cc59a042b"></a>

## ipv4 property — aws_parameters.tgw_cidr / eb77418e5da4 / 4

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

<a id="canonical-ff607060529d005846a68660c1a42163b08232209435ccb39a10c714027e9499"></a>

## Next pages — aws_parameters.tgw_cidr / eb77418e5da4 / 5

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-f38ef4556d9454ee1fa3198b1de8b3eace6ebf064864c8599236cc65e500aa8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f995c6dd0a2905b6654de90e37a7cf40f57581e95d2602e24b304dc712e14b74"></a>

## block_all_services — block_all_services / 31dc8defcac1 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- block_all_services

<a id="canonical-92687f15442e205ba25c1be6414ce764510d97ca41eef734e413e6532ef940c5"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

- [block_all_services](resources--aws_tgw_site--reference--group-002.md#canonical-92687f15442e205ba25c1be6414ce764510d97ca41eef734e413e6532ef940c5)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-f2bc47ab189f0fdf332ba3670d67a08ace7945d34b26e4a91ec052c1166f8a63)
- [default_blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-807b0250ceb5f671f89c9359c068c15526f4a67bc2148a562a2ef648ce982475)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

<a id="canonical-b960bd6e9a047c8fa10fae1a0a0f2b452ec7fbfb9005f4c554716e6a69774a2b"></a>

## Direct properties — block_all_services / 31dc8defcac1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-818fdeaeac25f038baf5262ad480d8a2965ce20c0f4ea7f574e8e7a4a3db261a"></a>

## Next pages — block_all_services / 31dc8defcac1 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92b21ac29f010d4e62cb5a1977aff97e5f9526d0cd1ddae3484ba21447001d27"></a>

## blocked_services — blocked_services / ceffe77db969 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- blocked_services

<a id="canonical-f2bc47ab189f0fdf332ba3670d67a08ace7945d34b26e4a91ec052c1166f8a63"></a>

Type: `"object"`. single nested block, Optional.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-10016aab3cfd70b0c426b4d0bdbf7b4a4fc90cf4bbd73d3384822ebf300cc1ed"></a>

## Direct properties — blocked_services / ceffe77db969 / 3

- [blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34): complete subsection reference.

<a id="canonical-f189c046c5aba28d171e08d7d5ab0edd43ad7a70c7231f0a6d3624cbf0ef4306"></a>

## Next pages — blocked_services / ceffe77db969 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3b6878a88df6dee582240bb31a9f69a8335557612e978308c6fb0364c2301f8"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / a24cc2f57852 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51)
- blocked_services.blocked_service

<a id="canonical-eeec99388c63976fa00e0b755a89d013ae9ba1f8206cfe7f40697a82f450c53a"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-f343277c4b4e98e61b55ef792fdb46ecd80b092cc41eab347f852fb5f455cd45"></a>

## Direct properties — blocked_services.blocked_service / a24cc2f57852 / 3

- [dns](resources--aws_tgw_site--reference--group-002.md#canonical-c8e37ab5416508f913c91d935414ed2b750fafa1755384d929f4ce51f38f6555): complete subsection reference.

<a id="canonical-75f4b773c776279b418ffcfc1fd9f25a2c0c6cb3690d6612e04d7eb350a7b66b"></a>

<a id="canonical-f1b0a2328bfa36fb8f3f0aab8b45fc53eda4911fa1959c3e47e46932e524e643"></a>

## network_type property — blocked_services.blocked_service / a24cc2f57852 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ssh](resources--aws_tgw_site--reference--group-002.md#canonical-67bb1c860f509e630190435d279d83e388e770b52c6e94691f7a36be598c105c): complete subsection reference.

- [web_user_interface](resources--aws_tgw_site--reference--group-002.md#canonical-3d04196616feb7ac625ae23e40d0c9eb99677862485fcf1061ce61eafa94fa34): complete subsection reference.

<a id="canonical-e40e87377f014233fe939c645104e17fd200fc31b474bb75cd3a6e32f8504f50"></a>

## Next pages — blocked_services.blocked_service / a24cc2f57852 / 5

- [blocked_services.blocked_service.dns](resources--aws_tgw_site--reference--group-002.md#canonical-c8e37ab5416508f913c91d935414ed2b750fafa1755384d929f4ce51f38f6555)
- [blocked_services.blocked_service.ssh](resources--aws_tgw_site--reference--group-002.md#canonical-67bb1c860f509e630190435d279d83e388e770b52c6e94691f7a36be598c105c)
- [blocked_services.blocked_service.web_user_interface](resources--aws_tgw_site--reference--group-002.md#canonical-3d04196616feb7ac625ae23e40d0c9eb99677862485fcf1061ce61eafa94fa34)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-c8e37ab5416508f913c91d935414ed2b750fafa1755384d929f4ce51f38f6555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-692a5c3274b79995f2825a697a979eaaed4dcd7b2519f5c6cb4965a344a265d8"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / dcafb1f322f9 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34)
- blocked_services.blocked_service.dns

<a id="canonical-621b89128455bb18a4d94ea95deb72d8284ab2c596b7064d38605f1a14b07115"></a>

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
dns = {}
```

<a id="canonical-632ad4cc828f99748977ecd8208720371d2ac75276b8498dcb41f1dad2d8456a"></a>

## Direct properties — blocked_services.blocked_service.dns / dcafb1f322f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a65f9af59fbe5b4655441abdc7585521fec2cc862cefb69e57a8d9d671b0079"></a>

## Next pages — blocked_services.blocked_service.dns / dcafb1f322f9 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-67bb1c860f509e630190435d279d83e388e770b52c6e94691f7a36be598c105c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6761868f17efbc8aa7240145d421ed6c922f9cb141b1d32ff541513a290f91d1"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / 609195132098 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34)
- blocked_services.blocked_service.ssh

<a id="canonical-b060766320b33fe05c0fb93bf35318a72e658d649bdd2f5f7a037fee8bfd4ff4"></a>

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
ssh = {}
```

<a id="canonical-1ca5f7eac48e266106ae6fc7c4a211da013a4ce2c093fc8e33fe21a45835da2c"></a>

## Direct properties — blocked_services.blocked_service.ssh / 609195132098 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-479119a33f501a598a3dd18096f5276e8877580faf8784154bdebc201d3148b3"></a>

## Next pages — blocked_services.blocked_service.ssh / 609195132098 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-3d04196616feb7ac625ae23e40d0c9eb99677862485fcf1061ce61eafa94fa34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9adbfa2185bd50e54554f5b23b8095517e2e764f76b4c4645cccd9a0feb02674"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / d9398cf6a200 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-10005e84005b24614f640470ac6ee73bcf6de88469aa6e490f29296ca08974cf"></a>

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
web_user_interface = {}
```

<a id="canonical-cc4acd590b1dc9f0916bd1a80a242b342b1d6155ffc2b5ed7e701e98b6a2651b"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / d9398cf6a200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f458001bee7f67839928bad80e86b3a03a19ca523b15711f0b13ec9fec28c712"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / d9398cf6a200 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-fce4962825a12d03b6ee24c0665e00e554cd682443e47896e2312d5ea4dbed34)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-fc40e9958bb6e7178fac69151232bfbf4397fc7a8a01886162591f52e00b9404"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-669063e41f8eb888315a7d0d0a2ccc1aec490152d6f5f85f7247633ae2b713a7"></a>

## coordinates — coordinates / 3f4fe0a4cae8 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- coordinates

<a id="canonical-7a716c2ca189d903d4f26a83e0777f518d0b015913fca00111bbb64fdf4b93b8"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-45ef3787f9aac2e9c80886a96b162e32cd273bdab81faa0129bb950290b3aa15"></a>

## Direct properties — coordinates / 3f4fe0a4cae8 / 3

<a id="canonical-a3babf83d9f7722da27b24ebd29ae0cf5e9e480d5cb976135435538c7d37d632"></a>

<a id="canonical-ba4e2e8d0af4bc263f8caf8ac52ffb5a36c18749d71f8efa4a3524b128792cd1"></a>

## latitude property — coordinates / 3f4fe0a4cae8 / 4

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-63af6b5c77fc7f6763822f985c0bd5130c3f8c5b75fa46cc1e113b126f3ccd54"></a>

<a id="canonical-3f645443e7f6f21f6f04bba7086b74a6e7ff67e75b7405ab1bff0cfa98bdaad4"></a>

## longitude property — coordinates / 3f4fe0a4cae8 / 5

Type: `"number"`. Optional.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-b7eeb04117b6750c89d64eb64b69cf5319945462981ea8ea69858d4ba1eac3ba"></a>

## Next pages — coordinates / 3f4fe0a4cae8 / 6

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-347174fd79b6c86f16e86abc884c54253528abca8e5888200327edfd66b4729d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-023578d58d2ae05b91a8415b4bdc24ed4d9f9f1c3705e7f968a2b66bf916bfd9"></a>

## custom_dns — custom_dns / d8e84e5de87a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- custom_dns

<a id="canonical-f199fba880d5237e82ab901b5b39ce494c885d04ea7329b6f19bfa0d7c029dd4"></a>

Type: `"object"`. single nested block, Optional.

Custom DNS is the configured for specify CE site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-a7e047ee44572c911d9ae511e5a66e1707fba52a380f130cf559aa88106a93c8"></a>

## Direct properties — custom_dns / d8e84e5de87a / 3

<a id="canonical-e23470976891892ea64d4567a52a0b6b66851a97f1ef5d9acdd7fa773acd83f9"></a>

<a id="canonical-4589d030bd0f253be5551539ae630ab0d02ee592e9b290a5acc0102a35e97bb0"></a>

## inside_nameserver property — custom_dns / d8e84e5de87a / 4

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

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

<a id="canonical-4240b44ac7d801c6a7541e85ed48008a6225acf6986633aa10b789a99fa41cce"></a>

<a id="canonical-e2566060c0b65774eec760e41446799931f2c202f41fd36e48fae9ec2582154a"></a>

## outside_nameserver property — custom_dns / d8e84e5de87a / 5

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

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

<a id="canonical-48ad0faef192bc1e639dbf432e4285d5a7d0a2213c95fa8642d13fb5a0f6074c"></a>

## Next pages — custom_dns / d8e84e5de87a / 6

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-2e91e3b2b5e727c6e40a4216181f74d84039efa17ee193877ca6cab3695ef316"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d91539cfa4ff2ac1f0a15a077480feeaffcafd474581e0c4462920bd57dc0d07"></a>

## default_blocked_services — default_blocked_services / 7e8334b4a936 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- default_blocked_services

<a id="canonical-807b0250ceb5f671f89c9359c068c15526f4a67bc2148a562a2ef648ce982475"></a>

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
default_blocked_services = {}
```

<a id="canonical-d3fa7db99cdaf581e9e3b6c7ac19655b8ff72a8c19e3d30d6293b68c38c5e458"></a>

## Direct properties — default_blocked_services / 7e8334b4a936 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-658473ab576705b4754276951647e84639d9e8844e4363349ed4f5cd5625ab46"></a>

## Next pages — default_blocked_services / 7e8334b4a936 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-302989a4f6d557908c55a666fb3705955d8482c6561e1cab36d132b1cc5a1779"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae369cb1ef0b5fde5f340b62bc4dacf5d7cd7f4113f41ec1d75b94146505c449"></a>

## direct_connect_disabled — direct_connect_disabled / 2b10efe64fb6 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- direct_connect_disabled

<a id="canonical-dac64ee326fc8f93eab5e4cf182e098e5f4fa96a4a784d4e5780728193048c15"></a>

Type: `["object", {}]`. Optional.

\[OneOf: direct\_connect\_disabled, direct\_connect\_enabled, private\_connectivity\] Enable this
option

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

- [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-dac64ee326fc8f93eab5e4cf182e098e5f4fa96a4a784d4e5780728193048c15)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-119113580de10ee0036ce9fc950181b4c727b9a197f59a0f781654dbcc0f0c0b)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0e44d8f367a3dd341058d0f9af047fd63771d872e0cfe26499986d5462c83478)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
direct_connect_disabled = {}
```

<a id="canonical-b3d993a52efdbdb7a57ead317e54ffb45ddf3393003a7f070f7be820db163dec"></a>

## Direct properties — direct_connect_disabled / 2b10efe64fb6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7d84446e9f03db2d5fb69d1d68eb619288d2b458b5d9ca436e03c1b0ca642c8d"></a>

## Next pages — direct_connect_disabled / 2b10efe64fb6 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80cc274209eda2c45c7437e62c3d3297b5665907f736c801c56a5bccf39aa4d1"></a>

## direct_connect_enabled — direct_connect_enabled / d55725ab6013 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- direct_connect_enabled

<a id="canonical-119113580de10ee0036ce9fc950181b4c727b9a197f59a0f781654dbcc0f0c0b"></a>

Type: `"object"`. single nested block, Optional.

Direct Connect Configuration. Direct Connect Configuration.

Upstream description:

Direct Connect Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_asn",
    "custom_asn"),
  validators.ConflictingObjectAttributes("hosted_vifs",
    "standard_vifs")}
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
  "x-ves-oneof-field-vif_choice": "[\"hosted_vifs\",\"standard_vifs\"]"
}
```

Terraform syntax:

```terraform
direct_connect_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-b4e1d8ac83844fb8033b2882fbdfba703363d6a8618cedd71555b368fc138b13"></a>

## Direct properties — direct_connect_enabled / d55725ab6013 / 3

- [auto_asn](resources--aws_tgw_site--reference--group-002.md#canonical-669adc10c260ad3795c0da1dc665be790a3ab27000d6aaf11d32e5b2a852de4b): complete subsection reference.

<a id="canonical-b79e1ca8e5852a5aecf0406582cbc9ce647d8913b7f3e60fb5fc5912e5cdcf20"></a>

<a id="canonical-6bc80af67b0b7cfdff1d1c06010c05efaa03bd983659038322d998fd0f21b256"></a>

## custom_asn property — direct_connect_enabled / d55725ab6013 / 4

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Upstream description:

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0): complete subsection reference.

- [standard_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-23eb357da373bfd76be24f46c8fb7547fb447c51cf9a982a87548b3b93b5e880): complete subsection reference.

<a id="canonical-ac3ed748e6a30ccece4267edd8dc8f2a8a582811e12440faed8c7d236a48976c"></a>

## Next pages — direct_connect_enabled / d55725ab6013 / 5

- [direct_connect_enabled.auto_asn](resources--aws_tgw_site--reference--group-002.md#canonical-669adc10c260ad3795c0da1dc665be790a3ab27000d6aaf11d32e5b2a852de4b)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- [direct_connect_enabled.standard_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-23eb357da373bfd76be24f46c8fb7547fb447c51cf9a982a87548b3b93b5e880)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-669adc10c260ad3795c0da1dc665be790a3ab27000d6aaf11d32e5b2a852de4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd923f1e4138ae93e61185150cf6a4110cfc8f8b0db2cdab741a9ed8e72d0b5a"></a>

## direct_connect_enabled.auto_asn — direct_connect_enabled.auto_asn / 1f23aa5d297c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- direct_connect_enabled.auto_asn

<a id="canonical-d9ba56a7a1e52dfbf7b9a60123b96709fe536d4377064a2d9fda86ede6beb79c"></a>

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

<a id="canonical-51003bca385a578cfa7ec045c934f69a936d108e805c1be80c48b21b4dd7c2bd"></a>

## Direct properties — direct_connect_enabled.auto_asn / 1f23aa5d297c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d5c907143f73b646b0223328cfb5b765c6d8443db1944a30bf099190b8c921f"></a>

## Next pages — direct_connect_enabled.auto_asn / 1f23aa5d297c / 4

- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a8ee4602711566f536ae3ffb2aed141f4d3deeeda3d2955c44c963c2ce6b059"></a>

## direct_connect_enabled.hosted_vifs — direct_connect_enabled.hosted_vifs / 0da14edde432 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- direct_connect_enabled.hosted_vifs

<a id="canonical-119e032f4274ef3ffa7b6371aa7347e5c96c1ccf2b3b09ff0d4fd9720f578fb8"></a>

Type: `"object"`. single nested block, Optional.

AWS Direct Connect Hosted VIF Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_registration_over_direct_connect",
    "site_registration_over_internet")}
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
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

Terraform syntax:

```terraform
hosted_vifs {
  # Configure direct properties listed below.
}
```

<a id="canonical-330efc3307195b0ff14d62d7c3ac09ec8452d67defce7a4be4132d22fdfac826"></a>

## Direct properties — direct_connect_enabled.hosted_vifs / 0da14edde432 / 3

- [site_registration_over_direct_connect](resources--aws_tgw_site--reference--group-002.md#canonical-e5c09e1e1686e17dc5dad570ea49e43747ad9ff7130d9499a894ff60acc07120): complete subsection reference.

- [site_registration_over_internet](resources--aws_tgw_site--reference--group-002.md#canonical-afccdcfba036cca5586ca49f0cc7e9edf7ca8e05919044ef76bf6972fe62073c): complete subsection reference.

- [vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-05a7600b3f29b618ac78acfe9fb40657211ffaf74729618efc8a36ed0b32f833): complete subsection reference.

<a id="canonical-12aed17d54f4e62cc60ec1c4c23ba5328e7d224e682ac6bb2f67482c3f97bb2a"></a>

## Next pages — direct_connect_enabled.hosted_vifs / 0da14edde432 / 4

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_tgw_site--reference--group-002.md#canonical-e5c09e1e1686e17dc5dad570ea49e43747ad9ff7130d9499a894ff60acc07120)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_tgw_site--reference--group-002.md#canonical-afccdcfba036cca5586ca49f0cc7e9edf7ca8e05919044ef76bf6972fe62073c)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-05a7600b3f29b618ac78acfe9fb40657211ffaf74729618efc8a36ed0b32f833)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-e5c09e1e1686e17dc5dad570ea49e43747ad9ff7130d9499a894ff60acc07120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a1bb056e70d81269459a56d1149cd5c7f86092c3bb5d0919a307977a6dfe8af"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / fd6d8d623b40 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-726f2079f0b9942e52ac688f3aa00abd3b4560d8e407f9a2f440ba66d52d0356"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
site_registration_over_direct_connect {
  # Configure direct properties listed below.
}
```

<a id="canonical-5997fce4b9defe9f2180591066a1e69dd46be0ee665462bd331028add1e57a70"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / fd6d8d623b40 / 3

<a id="canonical-885eca1e38927dfb8b37d71883508731f8b4b365c3ea67d07b7b78d190bb4f79"></a>

<a id="canonical-43fad8756dd1ee7467bcabbf1d0c276a666b384528075d720e8d436f1f7c48a0"></a>

## cloudlink_network_name property — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / fd6d8d623b40 / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0227e7bbae58dcd3df3bea6e83cda1da65ae0b1ffa0136304cc342f947bfea2e"></a>

## Next pages — direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect / fd6d8d623b40 / 5

- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-afccdcfba036cca5586ca49f0cc7e9edf7ca8e05919044ef76bf6972fe62073c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b379ab05ea143b4719f517b77a1aec25a0e3cd1eb4057c8c5789fde473e80bdc"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_internet — direct_connect_enabled.hosted_vifs.site_registration_over_internet / 1cab7ad15d5b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-688d9d2bcef3f2ebd594b44c48b7dd51cf399c0ac25ac61885c5d74e72952024"></a>

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
site_registration_over_internet = {}
```

<a id="canonical-7ab180f67e76b938c6079e625243f933bfcec57b523d75daa52fe7702f334fa4"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.site_registration_over_internet / 1cab7ad15d5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-708e72dfc9a17ce31e39227def32531ba490fde9df96047f012b9aacb1719e1e"></a>

## Next pages — direct_connect_enabled.hosted_vifs.site_registration_over_internet / 1cab7ad15d5b / 4

- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-05a7600b3f29b618ac78acfe9fb40657211ffaf74729618efc8a36ed0b32f833"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db42dd4b023d6a7f3adb77ad38ae18b0ae1b598a17d65478924240fb118921c7"></a>

## direct_connect_enabled.hosted_vifs.vif_list — direct_connect_enabled.hosted_vifs.vif_list / 9bd4bc64cc79 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-52f854fec34897fd8ba709c668724e0fe511a4eb7c9a6258046c16a177895da8"></a>

Type: `"object"`. list nested block, Optional.

List of Hosted VIF Config. List of Hosted VIF Config.

Upstream description:

List of Hosted VIF Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("vif_id"),
  validators.ConflictingListObjectAttributes("other_region",
    "same_as_site_region")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
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
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
vif_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-7373cf0a67a48f3ff03a0e8d908ea9c8608a82016f5019ede1a971e300d1eec3"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.vif_list / 9bd4bc64cc79 / 3

<a id="canonical-6cfd0737db92fc3f387d5c2b64751a5ca5d23fec84090d6216926acf7a7fc547"></a>

<a id="canonical-4bbf2eeea0c1a3969d1174001caf542f6d3f8cb17d1db4af4bffd81d53cafdc6"></a>

## other_region property — direct_connect_enabled.hosted_vifs.vif_list / 9bd4bc64cc79 / 4

Type: `"string"`. Optional.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Upstream description:

Exclusive with \[same\_as\_site\_region\] Other Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
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
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](resources--aws_tgw_site--reference--group-002.md#canonical-538e18e9ae819866c1f82244c55013ebcfd87c2ec45118b69bf74cffa39eb7bd): complete subsection reference.

<a id="canonical-f57e856261732350d7293e4785c6f32d771c00b8f22a7f2b44f4467b1bdc9355"></a>

<a id="canonical-83820a29395510872c16b48f5e4f5bf4a9b1562fbfa43f9ad7dc18bf1b265c26"></a>

## vif_id property — direct_connect_enabled.hosted_vifs.vif_list / 9bd4bc64cc79 / 5

Type: `"string"`. Optional.

AWS Direct Connect VIF ID that needs to be connected to the site.

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
    "pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-6e9de8fa50768fce5ad63a9cb8092e11b34166dd48917c188cd35b8d0fe15184"></a>

## Next pages — direct_connect_enabled.hosted_vifs.vif_list / 9bd4bc64cc79 / 6

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_tgw_site--reference--group-002.md#canonical-538e18e9ae819866c1f82244c55013ebcfd87c2ec45118b69bf74cffa39eb7bd)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-538e18e9ae819866c1f82244c55013ebcfd87c2ec45118b69bf74cffa39eb7bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4cecc478e894234adf7d22e5a3522fd45504afefdb5766c6b97973cd0c009a9"></a>

## direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / b8e58aab7d90 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-be328f2805d8ef14ce2184c7845bee7bd7a56307466a16c28c2f2c5d110f16b0)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-05a7600b3f29b618ac78acfe9fb40657211ffaf74729618efc8a36ed0b32f833)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-232ba9595420745655e74675f55e3921efe87825970de3519b8e20e156b580e0"></a>

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
same_as_site_region = {}
```

<a id="canonical-0600012647bc9980a98ef820961f9f7e8c269cd25856b9a051e5c27fa2757646"></a>

## Direct properties — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / b8e58aab7d90 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a7f122d94db4f6727dbe8761ac8709c0675acb94d7e3779795ff8a7c06d9e7b"></a>

## Next pages — direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region / b8e58aab7d90 / 4

- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-05a7600b3f29b618ac78acfe9fb40657211ffaf74729618efc8a36ed0b32f833)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-23eb357da373bfd76be24f46c8fb7547fb447c51cf9a982a87548b3b93b5e880"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4534f506dc566daaebda78e72a05349e23f59456f4ffdf9e46f59dc11d451cc8"></a>

## direct_connect_enabled.standard_vifs — direct_connect_enabled.standard_vifs / 66e704a479de / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- direct_connect_enabled.standard_vifs

<a id="canonical-e2b6081b27a7cc7457005544c6e2de52bfc22b959f0b5a15cdd94bb594596f62"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for standard vifs.

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
standard_vifs = {}
```

<a id="canonical-4db5d9dc0fc18bad3a3bf0980f53def12b40de096f4303cb921f1dcd54bd7cc0"></a>

## Direct properties — direct_connect_enabled.standard_vifs / 66e704a479de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab54f82c7ef2e33b4f032d9127d4cf43a528cafffd8689cf28e8eec87c5c6eb8"></a>

## Next pages — direct_connect_enabled.standard_vifs / 66e704a479de / 4

- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72531356c80ab8c9bf9f0511d5712820768969949dfc8c3899697cabccc48a3f"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 5db80ad50dcd / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- kubernetes_upgrade_drain

<a id="canonical-4f80ba7a872482f05d380accb2aefbcf42330e69bf9fb45e73b042727556fd53"></a>

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

<a id="canonical-0610a6fe3f752ea11d3edabdcf87f3b8027d5ff212b344c4fcc9769811b7cddb"></a>

## Direct properties — kubernetes_upgrade_drain / 5db80ad50dcd / 3

- [disable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-38431793ebdc45dcb063fb47b5ef29797ea2a0f86a65abecb581f1a53bbc0b87): complete subsection reference.

- [enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-937263fb6cf7595520cb450eb79ba9056e59b7bd5a4f2a7700edaa92ddcc9775): complete subsection reference.

<a id="canonical-71737f750e0b6af58336193f655ae0666a38388cfc679ccebb9ccc9a5aa7402b"></a>

## Next pages — kubernetes_upgrade_drain / 5db80ad50dcd / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-38431793ebdc45dcb063fb47b5ef29797ea2a0f86a65abecb581f1a53bbc0b87)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-937263fb6cf7595520cb450eb79ba9056e59b7bd5a4f2a7700edaa92ddcc9775)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-38431793ebdc45dcb063fb47b5ef29797ea2a0f86a65abecb581f1a53bbc0b87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f3833eeae8646b3c02263764fdb72b4dc12a11f4350b5fdf5b12e5bf3da66aa"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / ebe349472b8a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-b9d4c33bf3714ac57754bf2cb644a4613545d4588fa7958c1d39d93e008335c3"></a>

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

<a id="canonical-e924e9100d5d6e477e84ce547eb271e9b4c03eba3906c5fe7073417bf4daa814"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / ebe349472b8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60153f3a6af758a1bec009feeff43e0ddb549aadd13e2e630495ffd522005d6f"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / ebe349472b8a / 4

- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-937263fb6cf7595520cb450eb79ba9056e59b7bd5a4f2a7700edaa92ddcc9775"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12ed9d4105d8cdf34e9fa95dd06ac117481f750dc92cdb2743a0d0e0819caf68"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / 8e8576887fa8 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-60b956b5d441b3cf8ac98642dde4c9624498ea0736a2546b2915bb750f76dcaa"></a>

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

<a id="canonical-f93d11521b15fcdae21967f68600b2316cd61df8ba26121b2f58155e9071cb92"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / 8e8576887fa8 / 3

- [disable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-0103462d482f5d93510b15af908a09879787e75af35455bef455f38e1fb10bb2): complete subsection reference.

<a id="canonical-1039210c2552d74cf4378dceeec97871f5c0cef09f1ad76326b7c42d0f75f49a"></a>

<a id="canonical-6fbbf15b57dc5ae29fdd97d7133ec30b9fdcccaac6673eafe21cadf8db0931b5"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / 8e8576887fa8 / 4

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

<a id="canonical-9e20e81d1ad9b0819fa3bf3be45d193e4c77a2a490cc310ce13b46c8f49567af"></a>

<a id="canonical-0697f3e5e508ca8044b7d04713fbea1808ac0785956315e5fceb73974d0ad3a2"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / 8e8576887fa8 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-befd8b74dca385a5ec123562ef377086bed50e5d1cd82e0b37dc6c49ba63ff27"></a>

<a id="canonical-82394eb0021019e378ed83450845e931fafa17db161cbc76da6c96adeee88869"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / 8e8576887fa8 / 6

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

- [enable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-a5397b2b5d8ee8de2b3e2365fe53a3e1c16b76e2d1f1217a46294a0702864760): complete subsection reference.

<a id="canonical-ca0e2356af6a60a54228ce609b4e01fa7edf1808ad3f7f3b8121a1b74805bc8f"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / 8e8576887fa8 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-0103462d482f5d93510b15af908a09879787e75af35455bef455f38e1fb10bb2)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-a5397b2b5d8ee8de2b3e2365fe53a3e1c16b76e2d1f1217a46294a0702864760)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-0103462d482f5d93510b15af908a09879787e75af35455bef455f38e1fb10bb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93ab872357e7e1c14e9ab767b817e2b0c50a8f85f85785afc01ba88d4f6c52d8"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 66ad18e2cd7c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-937263fb6cf7595520cb450eb79ba9056e59b7bd5a4f2a7700edaa92ddcc9775)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-5b5f04216e48cfec015b94fcbe7e8b1dd3a5751fefab338d7ad1debb6483dc63"></a>

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

<a id="canonical-327f5ea03f4748c197fed79c2869353024e987dbabc831a94c030ab69a106893"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 66ad18e2cd7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-948cd7dbc41dffcd6b1bf1793838071e7a748a61887fdf9d4982971535ecc806"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 66ad18e2cd7c / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-937263fb6cf7595520cb450eb79ba9056e59b7bd5a4f2a7700edaa92ddcc9775)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-a5397b2b5d8ee8de2b3e2365fe53a3e1c16b76e2d1f1217a46294a0702864760"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab62813734943f81fea2a190d234dfaafc5cea86f048fa9ee0f5d90fe1bcde50"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 05f4d79cbd32 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-937263fb6cf7595520cb450eb79ba9056e59b7bd5a4f2a7700edaa92ddcc9775)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-b1c187a1ac16b14917bfdf349d961da97426299c8d01a1910cee75eb08603366"></a>

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

<a id="canonical-561d31340e038baa705e509a5609b55f8d86dff0ca1d39702a96f85547abb4b8"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 05f4d79cbd32 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d16b6ec257eefd37bdf773aff773ead597941145a08340b930402dac7ddeb1fe"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 05f4d79cbd32 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-937263fb6cf7595520cb450eb79ba9056e59b7bd5a4f2a7700edaa92ddcc9775)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-5328bd4d13cd4b9bc6bf897e865e5f8b01e086064bde136bab9076d75b001faa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e63726de56a6c74abf5f7ffcb450db78d21b48a899e30d41c1331c17172563f"></a>

## log_receiver — log_receiver / 2d31ecbf11fa / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- log_receiver

<a id="canonical-e0f27dc08c3abcaea1c9442696f74690ea29b1c9c46d4e98f604c2ee93040813"></a>

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

- [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-e0f27dc08c3abcaea1c9442696f74690ea29b1c9c46d4e98f604c2ee93040813)
- [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-8244e24808d2c8df987ef85abc86cd1b6c7934aee1339ba9123c1a5c51d4ae48)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e96a828a12bd386a33dfbee7c8efb326f265465dc4a3e73536e4c1164a6aaa4"></a>

## Direct properties — log_receiver / 2d31ecbf11fa / 3

<a id="canonical-008b8073fb391c29d62520ebcb5371d09a72beaf5ef6ec3ba57619aeb1eac08a"></a>

<a id="canonical-e30f415c7811adbb68c13ae6c44b396418b51ebb64ef6aa2d00098771eebc1ed"></a>

## name property — log_receiver / 2d31ecbf11fa / 4

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

<a id="canonical-d482a55f2a1976d838e20a3c66a4022299c79fe97b805230732dcdafe3dadb00"></a>

<a id="canonical-44109e62af97960d53db5d0c8c8bc31d591fe69ef0509ccec4ad03ae3382a7fb"></a>

## namespace property — log_receiver / 2d31ecbf11fa / 5

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

<a id="canonical-7e1a5abd62b02327667ebe227464a05d3c9e2981377e98e6df4180ef8e300c9b"></a>

<a id="canonical-25754f47a628179a64d2a66534806752bd4637fc9c597e4e01a67b66199874f5"></a>

## tenant property — log_receiver / 2d31ecbf11fa / 6

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

<a id="canonical-95e816747752b4d820cbac6b2557e087a06e6a4bbfcc9303bf512f2e622b6540"></a>

## Next pages — log_receiver / 2d31ecbf11fa / 7

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-fdb168a79ccc613abd6af4e13fca87d0cf0e20f031d3af2ea06cada88b87c750"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-142e88fc4eceaac3d6394bde56672645fd1f8c1b5f5b2f4cda57635877f61d99"></a>

## logs_streaming_disabled — logs_streaming_disabled / 5727e51d8e52 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- logs_streaming_disabled

<a id="canonical-8244e24808d2c8df987ef85abc86cd1b6c7934aee1339ba9123c1a5c51d4ae48"></a>

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

<a id="canonical-e0fc4068c489c656a594e0395d04c657523d3ea0eceaaed80689ce87660873a9"></a>

## Direct properties — logs_streaming_disabled / 5727e51d8e52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd5f0f9c9297757fd4d761a5526c4d655ab495fd0290ff0e97c280a274a2ff4e"></a>

## Next pages — logs_streaming_disabled / 5727e51d8e52 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-c1d488ffe9bc7e32f5e0c9a74b5c69d528239d7d4ecd68080450066df847f5fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-418c2d0ae5f27bd088da12487d5af5ce0fba21d29452e0a16d31dd8c1aae24a5"></a>

## offline_survivability_mode — offline_survivability_mode / 5ba36d43e73b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- offline_survivability_mode

<a id="canonical-f0eac2e3c39277cb8020e491b2312dd937083c1b4c31e4ab2fd4185e3a320250"></a>

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

<a id="canonical-e324373a78fcc255185ddad0df87e9a27bf923901df93a5622b0a9d84c0d82a1"></a>

## Direct properties — offline_survivability_mode / 5ba36d43e73b / 3

- [enable_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-7c82c0d55ce8f8b92f2ba833983503768ff2d610bfc235b64af0aa09f0e8177f): complete subsection reference.

- [no_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-65b43f5aec52d2970a06053a344c9ec5007f42d7de03d0a33e6c5da55b49aea7): complete subsection reference.

<a id="canonical-b6a9461052e0091cf24b0c8aace79715f1c24d882a471c26e81ff56dd23799fc"></a>

## Next pages — offline_survivability_mode / 5ba36d43e73b / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-7c82c0d55ce8f8b92f2ba833983503768ff2d610bfc235b64af0aa09f0e8177f)
- [offline_survivability_mode.no_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-65b43f5aec52d2970a06053a344c9ec5007f42d7de03d0a33e6c5da55b49aea7)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-7c82c0d55ce8f8b92f2ba833983503768ff2d610bfc235b64af0aa09f0e8177f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d75ac940f3f61d6c2f9481ea7a7a48858a29bcaa3136d042de6b080c08581955"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / cf4ce82bb365 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-c1d488ffe9bc7e32f5e0c9a74b5c69d528239d7d4ecd68080450066df847f5fb)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-d2d553899274a4da3b1dcbc1a2eeb3eccf2d887bc1a7cdad98183b5652ec5716"></a>

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

<a id="canonical-328cbd7c2b78f04bc17c0d4f5f2a3c8268d40d082af2c783ec339c1ab67bcf38"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / cf4ce82bb365 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed13d3c59ee9d66a24be57f77399a0124bea03c3e033bd2ab5623fd60910cd76"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / cf4ce82bb365 / 4

- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-c1d488ffe9bc7e32f5e0c9a74b5c69d528239d7d4ecd68080450066df847f5fb)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-65b43f5aec52d2970a06053a344c9ec5007f42d7de03d0a33e6c5da55b49aea7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28a60aa2addca364a41707b89f8b9d384660f23436fb4f4d4743c927f8c3db9f"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / ef6ac540cae9 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-c1d488ffe9bc7e32f5e0c9a74b5c69d528239d7d4ecd68080450066df847f5fb)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-64fd5c393616b0dd40f9f354898b531a46770d4119891be460c66c002a5149b7"></a>

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

<a id="canonical-e7b3c9105dad68bd597a024a98112114bbff545f2e9400590327d338d202a8b7"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / ef6ac540cae9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bfd2783e743cca565cdb587680d328d2300556f0fcb373c5b797471e930a1598"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / ef6ac540cae9 / 4

- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-c1d488ffe9bc7e32f5e0c9a74b5c69d528239d7d4ecd68080450066df847f5fb)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b36c285d2a1299188f2809f555233d568667868bcfa2072e2d4e916468e5ee8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e9e424df0c56f2672bb71466f04d3a294ee03f564a99fa99b7e77d3002aca11"></a>

## os — os / ce51314dd96c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- os

<a id="canonical-8abc68ef5d6144c9ad6b9acf5479950114bd51be5f75ad84155e094839732c08"></a>

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

<a id="canonical-d0a9fa8949895f2d39745809ac125d856c974b349a39bb659a37840c2457ead4"></a>

## Direct properties — os / ce51314dd96c / 3

- [default_os_version](resources--aws_tgw_site--reference--group-002.md#canonical-bd8f4966f2de9069876f093d333b5e8bd292ca7bc0cde926d2b1aa6c78ed4efc): complete subsection reference.

<a id="canonical-2c35bebd6b9e05ec8641085d1a61d6448432355f729aa5b15d0ca4f463d35166"></a>

<a id="canonical-5cfe1752e21749e67e03597b37ff2311e658144a42e98fc44deb99cd5cb19075"></a>

## operating_system_version property — os / ce51314dd96c / 4

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

<a id="canonical-a9ce9e6b1beea70e6bc9785c80283ffd09dfea855b327b27e808d9047cb791ef"></a>

## Next pages — os / ce51314dd96c / 5

- [os.default_os_version](resources--aws_tgw_site--reference--group-002.md#canonical-bd8f4966f2de9069876f093d333b5e8bd292ca7bc0cde926d2b1aa6c78ed4efc)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-bd8f4966f2de9069876f093d333b5e8bd292ca7bc0cde926d2b1aa6c78ed4efc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0901500006cb6facb97767553aad3a1a79332d33378dae442a85ff6a3b7258a"></a>

## os.default_os_version — os.default_os_version / fa0d55075d3c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [os](resources--aws_tgw_site--reference--group-002.md#canonical-b36c285d2a1299188f2809f555233d568667868bcfa2072e2d4e916468e5ee8c)
- os.default_os_version

<a id="canonical-54e2bdf5129eea60196e53e828c20526ef7afb1c5fb2630d65c4abef08b85990"></a>

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

<a id="canonical-28a6209ccf5f22d6a8378b26e2b965477b70fde6fbefbe98bc6c118aa17094d9"></a>

## Direct properties — os.default_os_version / fa0d55075d3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d96c279e933651256192b82d122b5fd97a23fde03d93d9d08cbc18854a2897e4"></a>

## Next pages — os.default_os_version / fa0d55075d3c / 4

- [os](resources--aws_tgw_site--reference--group-002.md#canonical-b36c285d2a1299188f2809f555233d568667868bcfa2072e2d4e916468e5ee8c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16e6a11d6b4455f7dd132e270af9166a123ada1cbcf58efc5965cbc6b427f029"></a>

## performance_enhancement_mode — performance_enhancement_mode / f12211f0d7da / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- performance_enhancement_mode

<a id="canonical-0f751d74fe8b13aa6d549d56934e712be9ebad56ad44b3020cbdc6480769845a"></a>

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

<a id="canonical-eb1e57b40feb8287be3b6a0b897d233448acaa236fbe0f00c4bac0326d76cc79"></a>

## Direct properties — performance_enhancement_mode / f12211f0d7da / 3

- [perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-c37992b5ab9bfb032fbaa559dd3e468a59c59c87f129ce6e77012956fa02a71b): complete subsection reference.

- [perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-e30d154621425de49081572d1d721dda3713d8138c4b40d3074364cf9c81ff7c): complete subsection reference.

<a id="canonical-ddb56b69a7780f6d47f5d78f99b5f84bc8dc17c6949648d5e4bffe6bd9d261a7"></a>

## Next pages — performance_enhancement_mode / f12211f0d7da / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-c37992b5ab9bfb032fbaa559dd3e468a59c59c87f129ce6e77012956fa02a71b)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-e30d154621425de49081572d1d721dda3713d8138c4b40d3074364cf9c81ff7c)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-c37992b5ab9bfb032fbaa559dd3e468a59c59c87f129ce6e77012956fa02a71b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cced079cf9e7ac34d7e9a8ae3a9cba1da8f80e6bd3580133999a15c20d6b3f12"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — performance_enhancement_mode.perf_mode_l3_enhanced / 61276e149c36 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-cd76f68ac940084410d8945b328530d8aa2b304511795f379220c6ace1e9bfbc"></a>

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

<a id="canonical-5931cacb3240e24109877a229bd2ab5e4d0ad65b9addb2350cbcbd7b7608b4f9"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced / 61276e149c36 / 3

- [jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-f90a7377fe64f7fe81faf12350c9d61e470a67d16bef155f5bab43c1bdf78af7): complete subsection reference.

- [no_jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-44f86bb5e01859e310634e4232ac6c983c9c11e6c194608dac880ddb2ce53df5): complete subsection reference.

<a id="canonical-5e90b282230210b723220738be0a0407fca1d70f61ed44275200d6f9e47fef06"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced / 61276e149c36 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-f90a7377fe64f7fe81faf12350c9d61e470a67d16bef155f5bab43c1bdf78af7)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-44f86bb5e01859e310634e4232ac6c983c9c11e6c194608dac880ddb2ce53df5)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-f90a7377fe64f7fe81faf12350c9d61e470a67d16bef155f5bab43c1bdf78af7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4561b95574c8b7a03a5c216958e26d8ce8e8a84a27430f3c1a45e32cd64c180e"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 5f4ec965d6ee / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-c37992b5ab9bfb032fbaa559dd3e468a59c59c87f129ce6e77012956fa02a71b)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-5f8ae67e816e9ae0208416825b268df9d22689b8af3ae9e335dcfa46296a4f55"></a>

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

<a id="canonical-e4e6e67daed628200e434532881e14a077bcb58609ec18c57f85c7bfb5d179fc"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 5f4ec965d6ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5396765817fa5e7689c0e65749bfda35e6fbeac8559370644b3c0b74836b3fd"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 5f4ec965d6ee / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-c37992b5ab9bfb032fbaa559dd3e468a59c59c87f129ce6e77012956fa02a71b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-44f86bb5e01859e310634e4232ac6c983c9c11e6c194608dac880ddb2ce53df5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c35c6a316e0c510829f3cea1cdcf26556f44fbdaee6e082cfdb3d93880a5bd33"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / ef9374bbd493 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-c37992b5ab9bfb032fbaa559dd3e468a59c59c87f129ce6e77012956fa02a71b)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-9245cb5bc3f298a3903e40687a5e39a382fa3df257e779db2e2bdb14e05800d5"></a>

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

<a id="canonical-b5ea223843d369c8c60b0e984d95813839bb602d1a5ba1d30fc88472befa5fcb"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / ef9374bbd493 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e854fa616d7722d57432c9f1854d97d24954221b4d3dd1dd079234a2e4c3b34a"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / ef9374bbd493 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-c37992b5ab9bfb032fbaa559dd3e468a59c59c87f129ce6e77012956fa02a71b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-e30d154621425de49081572d1d721dda3713d8138c4b40d3074364cf9c81ff7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a67ca819df97b8b3cd8b9a8446a807111a4882e9246285df77a113051a2971a"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — performance_enhancement_mode.perf_mode_l7_enhanced / 9dabdae56522 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-c6579b5591458cc7d60a8bc113b794d17148e03359b09c33197471e4e22249c5"></a>

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

<a id="canonical-5ed32f20c348d82d1c7e1971f4826ba5015a1ff448694bc240bc091fa66fd252"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced / 9dabdae56522 / 3

- [jumbo_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-b6c5aaccb70b39058993fcfc36ce6a25470a868ae10ae677ec9f524a04b367af): complete subsection reference.

- [jumbo_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ae0d2456004811fe86b5a1e250f502f040019ae8f3b6d16916f706cf0685f250): complete subsection reference.

<a id="canonical-87bb1b9408ef71784af03b1742ceaaf08b78f4296653f12c7242708b55aea159"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced / 9dabdae56522 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-b6c5aaccb70b39058993fcfc36ce6a25470a868ae10ae677ec9f524a04b367af)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ae0d2456004811fe86b5a1e250f502f040019ae8f3b6d16916f706cf0685f250)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b6c5aaccb70b39058993fcfc36ce6a25470a868ae10ae677ec9f524a04b367af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcd78ce5a542f07cf294a7aefccac983ac87f44aab1653135b17ae58b5711ab5"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / bf837576abdd / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-e30d154621425de49081572d1d721dda3713d8138c4b40d3074364cf9c81ff7c)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-ee5fe3a4cf21407ed8cfe6ea778739757bbbce12a47f02a5b1777f968df41641"></a>

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

<a id="canonical-ff7f0638f32216c3b19d9d31ec0cb442e8a8bdb416ba264b457a23ef7287fbde"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / bf837576abdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ea97487cc6bfe181f127e6c9ba556c13b52cc9ebf7a7c984f6b837a2420dedc"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / bf837576abdd / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-e30d154621425de49081572d1d721dda3713d8138c4b40d3074364cf9c81ff7c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-ae0d2456004811fe86b5a1e250f502f040019ae8f3b6d16916f706cf0685f250"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-478e352a197970558e369d71877c27596c0de1bbb3914566c0887894e21bd5dd"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / bf0195b82a52 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-e30d154621425de49081572d1d721dda3713d8138c4b40d3074364cf9c81ff7c)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-727be0eead130993e8ad1f234c4ee4dfb8a4effa172c7713259d1a67c9d61496"></a>

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

<a id="canonical-793d69042c93a4c994f485c78231df3dc1c98ccb815539dcde4d6d19ddfd1fba"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / bf0195b82a52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-543626dd4dddc38a0e00e6a2970a98e0a17d4c09fe7321eca669d8ce509e8dc5"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / bf0195b82a52 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-e30d154621425de49081572d1d721dda3713d8138c4b40d3074364cf9c81ff7c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b74d73ddebb8745d8fc947cf124904439aff7d1315004746f2b88d8c004c0ae6"></a>

## private_connectivity — private_connectivity / 56719cba145f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- private_connectivity

<a id="canonical-0e44d8f367a3dd341058d0f9af047fd63771d872e0cfe26499986d5462c83478"></a>

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

<a id="canonical-1965beeca4ddd9c565ca834fdfc60b758a719d693bc6457f7b8547f046797b4b"></a>

## Direct properties — private_connectivity / 56719cba145f / 3

- [cloud_link](resources--aws_tgw_site--reference--group-002.md#canonical-4faaef29101f15a5738b2f158e2dd1803c59ccbfa2548a3f3de5ac9e64a4271a): complete subsection reference.

- [inside](resources--aws_tgw_site--reference--group-002.md#canonical-8bdec5fbe6f14371b9bc7a6a334fb80ad84d8b6100f07f5c20892c3c88e0459b): complete subsection reference.

- [outside](resources--aws_tgw_site--reference--group-002.md#canonical-18c723784f22bccb792f92b551018c36a0089be693312db95a3b42beb61717cf): complete subsection reference.

<a id="canonical-1c1de232787ca04d7bf9a7094be7ad2030fc60b86e46b461dd5a0c596df22b00"></a>

## Next pages — private_connectivity / 56719cba145f / 4

- [private_connectivity.cloud_link](resources--aws_tgw_site--reference--group-002.md#canonical-4faaef29101f15a5738b2f158e2dd1803c59ccbfa2548a3f3de5ac9e64a4271a)
- [private_connectivity.inside](resources--aws_tgw_site--reference--group-002.md#canonical-8bdec5fbe6f14371b9bc7a6a334fb80ad84d8b6100f07f5c20892c3c88e0459b)
- [private_connectivity.outside](resources--aws_tgw_site--reference--group-002.md#canonical-18c723784f22bccb792f92b551018c36a0089be693312db95a3b42beb61717cf)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-4faaef29101f15a5738b2f158e2dd1803c59ccbfa2548a3f3de5ac9e64a4271a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1a40fc234cfd91286e336a0f4bc0a825166a78e0a9f31832ba3e0704de19eac"></a>

## private_connectivity.cloud_link — private_connectivity.cloud_link / 904ce617c33c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65)
- private_connectivity.cloud_link

<a id="canonical-d5e7365ab614ad33d0b97a329d5ff6de05b656b83fbe75c9249e5f6d68622b43"></a>

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

<a id="canonical-48fb25c0313ff75077736561ac72120fbf0465cadcc8181fc3de7a56126f6af5"></a>

## Direct properties — private_connectivity.cloud_link / 904ce617c33c / 3

<a id="canonical-2e8a982009f56602e4376f38a368756b78a25007cc565047caaf415bbee167fc"></a>

<a id="canonical-f0b8972bdd0ee1b8c5919fc0687e40f496c4e27ff7dc9c2525d7237b9808324e"></a>

## name property — private_connectivity.cloud_link / 904ce617c33c / 4

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

<a id="canonical-07b3da90560b7935051a2bb8181a94283367c15b4ee869e047bc7b95d4e2289f"></a>

<a id="canonical-943fe762a94aab82a4dbaaca2aa833844446adcec3036caed3a473ee9731cbca"></a>

## namespace property — private_connectivity.cloud_link / 904ce617c33c / 5

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

<a id="canonical-b0d4bfb91c2a372493933c046495aba56f1663046531b79adea1be5cf8f2f698"></a>

<a id="canonical-f0f61a8c4fe87a3a1cd5d1f2b312cc02d20231556981c06bed4d31b2d88bb030"></a>

## tenant property — private_connectivity.cloud_link / 904ce617c33c / 6

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

<a id="canonical-bff5cbc55156571c51823078ab73005aa32f690adc25be11b00ba4e3cede06c7"></a>

## Next pages — private_connectivity.cloud_link / 904ce617c33c / 7

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-8bdec5fbe6f14371b9bc7a6a334fb80ad84d8b6100f07f5c20892c3c88e0459b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2dbe3cb7e4496e55f8900cbfda1102523d3bda8860c3cb6f2ff872491bcaae68"></a>

## private_connectivity.inside — private_connectivity.inside / 9b349da6b4df / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65)
- private_connectivity.inside

<a id="canonical-faa8d5c50ac1ce021302bfce8278aefdf1bd5ea849892e84f98580ebf0569d70"></a>

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

<a id="canonical-bc45613a119b7806c758595e40d78906b2ca5bac5510876db85ba22896949960"></a>

## Direct properties — private_connectivity.inside / 9b349da6b4df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19d94534d90158a0940c68ccca824cb5bb4094dc8aea35d2a11780c288514ee8"></a>

## Next pages — private_connectivity.inside / 9b349da6b4df / 4

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-18c723784f22bccb792f92b551018c36a0089be693312db95a3b42beb61717cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d59b004465369859e1c178a5fb9ac62d04f19866557094172c3208237429cd1f"></a>

## private_connectivity.outside — private_connectivity.outside / f14fcb367cb2 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65)
- private_connectivity.outside

<a id="canonical-c4a2c241d523a5b606d2bb7d43a91a029b8a2876ec9ed3cdf91f1649afab8ead"></a>

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

<a id="canonical-350630a830539419ee17c72cc4b00b7d39dc25aae35371c29fa5560d0e8364d2"></a>

## Direct properties — private_connectivity.outside / f14fcb367cb2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb2e6de5925c9109c77da96c448b7aebbca7f6c265192da3eff542a3ab091946"></a>

## Next pages — private_connectivity.outside / f14fcb367cb2 / 4

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-7076ed8e78694917e5f4df2c7aa41c7f303f575fbcc702ac483a7905e883932d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-494c4d75e52a8327921925d6534d0a7ac88332b28ee5e9993be51a0cf536720c"></a>

## sw — sw / 508bda938aac / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- sw

<a id="canonical-688e5f8d462864c3a29ffd340314a04cbfc5ff2479892e76f2568ef428bc909e"></a>

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

<a id="canonical-29ed231841b14151e1444292b29f6bc1d3102190f2db59628cc195961df46d4f"></a>

## Direct properties — sw / 508bda938aac / 3

- [default_sw_version](resources--aws_tgw_site--reference--group-002.md#canonical-9a459debfa1aa620c5cc7037e8afe952eb03b18f798d4957528c636ee26c48bd): complete subsection reference.

<a id="canonical-79e6354f473ec8351bb7d621cf866612679e58b72602a1331b549cad572e9bb9"></a>

<a id="canonical-e0ccf3648d2cea65bbf9ba36705c62ff21b4fe5d773f9403ee130233bc9c5888"></a>

## volterra_software_version property — sw / 508bda938aac / 4

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

<a id="canonical-0c36b68b6e6fe8eae1e5d6f78de6f77cbeaa626bf9f1ace08bec4612437df834"></a>

## Next pages — sw / 508bda938aac / 5

- [sw.default_sw_version](resources--aws_tgw_site--reference--group-002.md#canonical-9a459debfa1aa620c5cc7037e8afe952eb03b18f798d4957528c636ee26c48bd)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-9a459debfa1aa620c5cc7037e8afe952eb03b18f798d4957528c636ee26c48bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ba97c4f63ef57349deb9454da5819b1c57d13fa6c822e13ac655cba68130d5e"></a>

## sw.default_sw_version — sw.default_sw_version / f4ca48495291 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-7076ed8e78694917e5f4df2c7aa41c7f303f575fbcc702ac483a7905e883932d)
- sw.default_sw_version

<a id="canonical-bde73d3e70e7ce34013056d02dcb38f064825997d17a0747978e2c9018304a6e"></a>

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

<a id="canonical-f89402a5feba77414c83d505f068156fad80b49337737b5664e67b6a86095ff9"></a>

## Direct properties — sw.default_sw_version / f4ca48495291 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3a681e9f2b81836bea387a2cf06cfec9ddbafc33529388330676a3854b194c5"></a>

## Next pages — sw.default_sw_version / f4ca48495291 / 4

- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-7076ed8e78694917e5f4df2c7aa41c7f303f575fbcc702ac483a7905e883932d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac67fc7d79e9489405d68e82f86954d3916f77a775f54ad960229872661be7cc"></a>

## tgw_security — tgw_security / 3fcadffd8893 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- tgw_security

<a id="canonical-76d50759ef620219023b57f8f9583cce4c1c3ed0cbe324c1912f1787d543ca47"></a>

Type: `"object"`. single nested block, Optional.

Security Configuration for transit gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "east_west_service_policy_allow_all"),
  validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "no_east_west_policy"),
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
  validators.ConflictingObjectAttributes("east_west_service_policy_allow_all",
    "no_east_west_policy"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy")}
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
  "x-ves-oneof-field-east_west_service_policy_choice": "[\"active_east_west_service_policies\",\"east_west_service_policy_allow_all\",\"no_east_west_policy\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]"
}
```

Terraform syntax:

```terraform
tgw_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c7635a3d33b86f8355df9fec825606967d18126c90c1c4fea6532377c922807"></a>

## Direct properties — tgw_security / 3fcadffd8893 / 3

- [active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-af4ee271e5518d775aab8426c5f40f72f50b97a5aa01c5e7391899a7f6f4dd0c): complete subsection reference.

- [active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-b44d6ed1ee3aa2dac77ea81f0578fba8a84dbad194beebedd32df7c9a267e9df): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-4d1f4211a668bb07878e424c5dc24cb0fb3124d3a33041905f57d861027bc775): complete subsection reference.

- [active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-f711d4613c1e3af2f43a1aeed8b9854b7bc296b0550673bfcfce72de17fda258): complete subsection reference.

- [east_west_service_policy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-66b73a2f845f368fc5e3dadd393f8a0c68fcdc9087f034c8de9e1332c78baf3e): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-61602b71196ea1672ae9091ebd8cdf74499bb69e93dab7ea60c7d7d3c0e0af21): complete subsection reference.

- [no_east_west_policy](resources--aws_tgw_site--reference--group-002.md#canonical-8b28679da2660b0558d699d5ae2f0748c9ecb55adf4146772965e524847ac111): complete subsection reference.

- [no_forward_proxy](resources--aws_tgw_site--reference--group-002.md#canonical-b92573a556c559329a287b35b4f32c5d66eec9960192a781aa2a22d2fd1c33f4): complete subsection reference.

- [no_network_policy](resources--aws_tgw_site--reference--group-002.md#canonical-63369bcb606956a588e21960767782e5c2c987b0e333d43bc9a375812c2e85cc): complete subsection reference.

<a id="canonical-29b4ed6d24153b1674eeb994ce912c34283b31b2be9788fd96c941cc0a2ec111"></a>

## Next pages — tgw_security / 3fcadffd8893 / 4

- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-af4ee271e5518d775aab8426c5f40f72f50b97a5aa01c5e7391899a7f6f4dd0c)
- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-b44d6ed1ee3aa2dac77ea81f0578fba8a84dbad194beebedd32df7c9a267e9df)
- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-4d1f4211a668bb07878e424c5dc24cb0fb3124d3a33041905f57d861027bc775)
- [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-f711d4613c1e3af2f43a1aeed8b9854b7bc296b0550673bfcfce72de17fda258)
- [tgw_security.east_west_service_policy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-66b73a2f845f368fc5e3dadd393f8a0c68fcdc9087f034c8de9e1332c78baf3e)
- [tgw_security.forward_proxy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-61602b71196ea1672ae9091ebd8cdf74499bb69e93dab7ea60c7d7d3c0e0af21)
- [tgw_security.no_east_west_policy](resources--aws_tgw_site--reference--group-002.md#canonical-8b28679da2660b0558d699d5ae2f0748c9ecb55adf4146772965e524847ac111)
- [tgw_security.no_forward_proxy](resources--aws_tgw_site--reference--group-002.md#canonical-b92573a556c559329a287b35b4f32c5d66eec9960192a781aa2a22d2fd1c33f4)
- [tgw_security.no_network_policy](resources--aws_tgw_site--reference--group-002.md#canonical-63369bcb606956a588e21960767782e5c2c987b0e333d43bc9a375812c2e85cc)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-af4ee271e5518d775aab8426c5f40f72f50b97a5aa01c5e7391899a7f6f4dd0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f2ff1bb06ac3057ce7f007f723741b45d70a147a73bc9e556f456ae3c47e3fd"></a>

## tgw_security.active_east_west_service_policies — tgw_security.active_east_west_service_policies / 766d5659051a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.active_east_west_service_policies

<a id="canonical-546c58bb71187c8cf2e0de09dcb75b70f03fa5325839abed1fb0ef75d35a44a9"></a>

Type: `"object"`. single nested block, Optional.

Active service policies for the east-west proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
active_east_west_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-b802a9d72ba0cf70ef990635fe140ce9cb0e1d1b5632711c7d6b0ea621438366"></a>

## Direct properties — tgw_security.active_east_west_service_policies / 766d5659051a / 3

- [service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-cd92af7879f5749d32aef26a56430c1347be9e9704a0227159259475af133652): complete subsection reference.

<a id="canonical-cd6a6029026ad267c79329857ad227684bbee754c8404a63e3c0650f5b426428"></a>

## Next pages — tgw_security.active_east_west_service_policies / 766d5659051a / 4

- [tgw_security.active_east_west_service_policies.service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-cd92af7879f5749d32aef26a56430c1347be9e9704a0227159259475af133652)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-cd92af7879f5749d32aef26a56430c1347be9e9704a0227159259475af133652"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a868d59391a0f36009b7795570b731a9d48d796738592ad34e3d2e8168f764f"></a>

## tgw_security.active_east_west_service_policies.service_policies — tgw_security.active_east_west_service_policies.service_policies / 953413c6dda4 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-af4ee271e5518d775aab8426c5f40f72f50b97a5aa01c5e7391899a7f6f4dd0c)
- tgw_security.active_east_west_service_policies.service_policies

<a id="canonical-0e1fef266b64f42d251c4c36e0d13548fcf2407bdc2b15459d59f9a383eed0bf"></a>

Type: `"object"`. list nested block, Optional.

List of references to service\_policy objects.

Upstream description:

A list of references to service\_policy objects.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-4deac80f70b35bb790ff09967df9e539270778be5b88d811393d473216a032a7"></a>

## Direct properties — tgw_security.active_east_west_service_policies.service_policies / 953413c6dda4 / 3

<a id="canonical-1e286906324dce3b16e33b7d0f981a900e9849f02bf96721198ced1bc2d39497"></a>

<a id="canonical-ecfdab0e2145ba037935f7dfd724f8446be0ab509d5a2e4ddbed2e1ebb060a8a"></a>

## name property — tgw_security.active_east_west_service_policies.service_policies / 953413c6dda4 / 4

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

<a id="canonical-bbcbe9abcc3c3b42d5567a35fb6ac48acba84e299e077e8915edfe7a5a73bfdd"></a>

<a id="canonical-17e210367c682c04093781ed34e79eb612f522525c54a45b295df36fbd306759"></a>

## namespace property — tgw_security.active_east_west_service_policies.service_policies / 953413c6dda4 / 5

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

<a id="canonical-63d2d93b6279117fd3b963635bb0dfe2b2948b87d5fbc3236d17a59a42a0e4b8"></a>

<a id="canonical-cbc7c81518408e5c231361198db95c3ee01fd1b055334ec2a6b7f7a0b45f5049"></a>

## tenant property — tgw_security.active_east_west_service_policies.service_policies / 953413c6dda4 / 6

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

<a id="canonical-2cb8db111382512cd57b66eb491f44cf2f8904c2837c223cf302274d91bd2b5b"></a>

## Next pages — tgw_security.active_east_west_service_policies.service_policies / 953413c6dda4 / 7

- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-af4ee271e5518d775aab8426c5f40f72f50b97a5aa01c5e7391899a7f6f4dd0c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b44d6ed1ee3aa2dac77ea81f0578fba8a84dbad194beebedd32df7c9a267e9df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63a5f2cc9d392b813af73655e1c91779a824cf19b74aa986a935e1e515f2e7de"></a>

## tgw_security.active_enhanced_firewall_policies — tgw_security.active_enhanced_firewall_policies / 1f377cda1e34 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.active_enhanced_firewall_policies

<a id="canonical-2b2541ae92baee7770f500699975dabe117ca66edff058e711daef0f8d25958f"></a>

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

<a id="canonical-47c3e7c533401321d6317cab05f685bd229c67c89857e9200fd764c277a96b3f"></a>

## Direct properties — tgw_security.active_enhanced_firewall_policies / 1f377cda1e34 / 3

- [enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-bdaa3e61a06bbe931bfab9fd94412edba893bb97aabb6a64926a564ed635b72e): complete subsection reference.

<a id="canonical-9b0c6da39b15a49b034ce5770e0e3ff2917c38f87be4c79de7bd5afed690abf8"></a>

## Next pages — tgw_security.active_enhanced_firewall_policies / 1f377cda1e34 / 4

- [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-bdaa3e61a06bbe931bfab9fd94412edba893bb97aabb6a64926a564ed635b72e)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-bdaa3e61a06bbe931bfab9fd94412edba893bb97aabb6a64926a564ed635b72e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90f0bce038218d040c770a24b2ed377064a6ca775c60fa237d1c1b4c96d91778"></a>

## tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies — tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies / 3a3fe6c73fa5 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-b44d6ed1ee3aa2dac77ea81f0578fba8a84dbad194beebedd32df7c9a267e9df)
- tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-1f367bba8542016b47dadd0240aba83003c76685b48709f22e3d2d4ab95c7a49"></a>

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

<a id="canonical-c5ed8cf4c9769502fd556aee17dad2be97b82df0be4e5f293d3635ea3c4998f7"></a>

## Direct properties — tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies / 3a3fe6c73fa5 / 3

<a id="canonical-02f7daa8c1a2549edbe8bc7937098b8f1d9b768f1a59d847fef008c353a1f0d4"></a>

<a id="canonical-7cc09af9a434144c2633eddaa28fc2ff366e8a45ececd923e3219dff0a085557"></a>

## name property — tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies / 3a3fe6c73fa5 / 4

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

<a id="canonical-78b76582ef8c063e185cce4c0f9d01e09fe2e79ef2dc08a88d7f794f3ee393b3"></a>

<a id="canonical-c2558b9b0503b7508f63c085df5ec347386a168ae44aa6ce2c79eb1ea00084bf"></a>

## namespace property — tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies / 3a3fe6c73fa5 / 5

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

<a id="canonical-fcca5b445a3e57dbb75f28ecd7d29ed88cf8309704e05c044074afdb9da7cdec"></a>

<a id="canonical-628e090866ad1c128956d3f379482266f6a416d61e97c647f2afa765fdf3e3e5"></a>

## tenant property — tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies / 3a3fe6c73fa5 / 6

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

<a id="canonical-d9d2983923967957664aa3083e27c351fc73eaf1a904be43707a761a66e82aa2"></a>

## Next pages — tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies / 3a3fe6c73fa5 / 7

- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-b44d6ed1ee3aa2dac77ea81f0578fba8a84dbad194beebedd32df7c9a267e9df)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-4d1f4211a668bb07878e424c5dc24cb0fb3124d3a33041905f57d861027bc775"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b87895c2c2a63e8c09f3754bac97348a37838397bf3e4ca40ade2cb7d0fdf1f"></a>

## tgw_security.active_forward_proxy_policies — tgw_security.active_forward_proxy_policies / 65e482618bdb / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.active_forward_proxy_policies

<a id="canonical-0f7d3b567a14a0bf518148dba55bef2c8f54e0f14e109024d84c816f76fde0e1"></a>

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

<a id="canonical-5e7be7599c38ba1f66f4464f58fb942a69c59879edfb3fd0b2f2e62b12531704"></a>

## Direct properties — tgw_security.active_forward_proxy_policies / 65e482618bdb / 3

- [forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-55514d8df59cf32525efcb37e583c8a816684a611b19ddd1d3c5585c5c836df2): complete subsection reference.

<a id="canonical-dda586064b19471148c3c624ad03649eb7b2fcc5551918a883e78b23c9bc2ca7"></a>

## Next pages — tgw_security.active_forward_proxy_policies / 65e482618bdb / 4

- [tgw_security.active_forward_proxy_policies.forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-55514d8df59cf32525efcb37e583c8a816684a611b19ddd1d3c5585c5c836df2)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-55514d8df59cf32525efcb37e583c8a816684a611b19ddd1d3c5585c5c836df2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89232d288faad6f3a8837f7c94653fa385b08445b1b163038401923105e7c632"></a>

## tgw_security.active_forward_proxy_policies.forward_proxy_policies — tgw_security.active_forward_proxy_policies.forward_proxy_policies / c0bd31d620bf / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-4d1f4211a668bb07878e424c5dc24cb0fb3124d3a33041905f57d861027bc775)
- tgw_security.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-181a240d439c37f404932563ba560731d1a451fe7c56f4a95442c67b0a4fd55a"></a>

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

<a id="canonical-03867131264fdb5b9b738dead05563f46849579d4a1ee5b782daeeedfcc1ec3a"></a>

## Direct properties — tgw_security.active_forward_proxy_policies.forward_proxy_policies / c0bd31d620bf / 3

<a id="canonical-f818aa3cb150cc044405927ba334aaa6ce46a7dc64234b262a320d2504c5de69"></a>

<a id="canonical-8bd52a6dc21e69364613e668b30b60c25e844d74362993934f57974bd09b8d0a"></a>

## name property — tgw_security.active_forward_proxy_policies.forward_proxy_policies / c0bd31d620bf / 4

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

<a id="canonical-81704ea759db82e2279ceeef6a5fac7f95200bbf37d67b7fcfa000a9c447655c"></a>

<a id="canonical-748e5a8cfcec99b0076fb1805a4d8620c2c09a825a2362411ed537ef6d511592"></a>

## namespace property — tgw_security.active_forward_proxy_policies.forward_proxy_policies / c0bd31d620bf / 5

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

<a id="canonical-bcb1109912269e993776ff0c81a55eeb8126a5883b9b8d8a2a6336ed1a146d51"></a>

<a id="canonical-88bdb3b92d725458371deaa7ad643f5ebdc512f3cc1bb17436df7bb6bdd6c0c2"></a>

## tenant property — tgw_security.active_forward_proxy_policies.forward_proxy_policies / c0bd31d620bf / 6

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

<a id="canonical-b493bb67ef887353185d592fea6778f376043a58c678c4abc503f235ed25c39e"></a>

## Next pages — tgw_security.active_forward_proxy_policies.forward_proxy_policies / c0bd31d620bf / 7

- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-4d1f4211a668bb07878e424c5dc24cb0fb3124d3a33041905f57d861027bc775)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-f711d4613c1e3af2f43a1aeed8b9854b7bc296b0550673bfcfce72de17fda258"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9881b2432a717335e52105d942f5d9365f2a63852a7b35bd498989cb7f60ad25"></a>

## tgw_security.active_network_policies — tgw_security.active_network_policies / 455ad904dff7 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.active_network_policies

<a id="canonical-47c1bb875fb6eb216f7fcc4babee0a94f7db2d8a8e2753e86755ff0c439d8d7c"></a>

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

<a id="canonical-06062c6d5311ea5d4330b010649082aabff101b047fd0e9c351143ccb9194d68"></a>

## Direct properties — tgw_security.active_network_policies / 455ad904dff7 / 3

- [network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-b393222c3c3b2f5ecfeaca1bfdba5187e185dc4eebfdcbf2e4f096a107557ec5): complete subsection reference.

<a id="canonical-df0f25b87e2af33191f764899be242d8dd21997bce8a618d205b27c0d87652b9"></a>

## Next pages — tgw_security.active_network_policies / 455ad904dff7 / 4

- [tgw_security.active_network_policies.network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-b393222c3c3b2f5ecfeaca1bfdba5187e185dc4eebfdcbf2e4f096a107557ec5)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b393222c3c3b2f5ecfeaca1bfdba5187e185dc4eebfdcbf2e4f096a107557ec5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ede462bc99dd04a52a264639f98c38641145067d11b0b1013d99e1a1ad44bded"></a>

## tgw_security.active_network_policies.network_policies — tgw_security.active_network_policies.network_policies / 894b273248e7 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-f711d4613c1e3af2f43a1aeed8b9854b7bc296b0550673bfcfce72de17fda258)
- tgw_security.active_network_policies.network_policies

<a id="canonical-fa35526f956abd6e5bf2b15c578f3193119a846574b4c38b0c9de42dccd381ef"></a>

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

<a id="canonical-eb25ab9fd795467efdcf30e6b2c142769be9c3be82ef9fb532da1aa8fc8361c5"></a>

## Direct properties — tgw_security.active_network_policies.network_policies / 894b273248e7 / 3

<a id="canonical-8110123b6d20e7f94aae25c16aa22373bd3a2c08ce37f8b2f75474a147c73700"></a>

<a id="canonical-c166e6b7920a916e8b60e96d27deeea752dcd6bde5b09be217651091db5ad13f"></a>

## name property — tgw_security.active_network_policies.network_policies / 894b273248e7 / 4

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

<a id="canonical-ed4c08fd4e2ed13f698752d72e4539bbacd5de3719abc9902aae8b7c598cb9c4"></a>

<a id="canonical-2392d91b5cdf08e871e8183eff1a3d583b5644b36b5be7105dca593b0e2b5605"></a>

## namespace property — tgw_security.active_network_policies.network_policies / 894b273248e7 / 5

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

<a id="canonical-0b27c45cd488e53df91cb0eac6e9470ce63d51250e3b39d6cbf09ada0f311d8b"></a>

<a id="canonical-7ccd2781478a55f2c5dd4f786edb88f2605fd4baf44258bceea2149b0c0b6626"></a>

## tenant property — tgw_security.active_network_policies.network_policies / 894b273248e7 / 6

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

<a id="canonical-cf44791da188258a92b6fd35b4180bd269c5ef5ec0da9b3a6c8128bfff513766"></a>

## Next pages — tgw_security.active_network_policies.network_policies / 894b273248e7 / 7

- [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-f711d4613c1e3af2f43a1aeed8b9854b7bc296b0550673bfcfce72de17fda258)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-66b73a2f845f368fc5e3dadd393f8a0c68fcdc9087f034c8de9e1332c78baf3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d31381476d85a9f5053132496f9af51f13c0dce1ac7a00e267a68c2f13438445"></a>

## tgw_security.east_west_service_policy_allow_all — tgw_security.east_west_service_policy_allow_all / bfd5aa42860a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.east_west_service_policy_allow_all

<a id="canonical-ac04e59c692ab1bd6136a083e7424b87d913a586a319177f4afe1cb301ac4d0b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for east west service policy allow all.

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
east_west_service_policy_allow_all = {}
```

<a id="canonical-d8f5a5d34dcf2d5e6cf70f27e7a869e989dd170ad4ace8fe47d01535dfb26795"></a>

## Direct properties — tgw_security.east_west_service_policy_allow_all / bfd5aa42860a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64314be16d86b12ae7f84db5af1ef782d5448659f01a5765316bf0a650505174"></a>

## Next pages — tgw_security.east_west_service_policy_allow_all / bfd5aa42860a / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-61602b71196ea1672ae9091ebd8cdf74499bb69e93dab7ea60c7d7d3c0e0af21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5baf0fca487d00e502f3cd18c5b37407d4b2ac324f024ce5ce44bf4d3f4c5d28"></a>

## tgw_security.forward_proxy_allow_all — tgw_security.forward_proxy_allow_all / b173bb4bc8ef / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.forward_proxy_allow_all

<a id="canonical-bf0dd5e4aa6889bcc4dc270207578b0c1486fcaa37be5c9c70777d813ed6c1b2"></a>

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

<a id="canonical-b97cd6092b76e690401d7f9f1294acca72a98f243c098bc857207d048d160850"></a>

## Direct properties — tgw_security.forward_proxy_allow_all / b173bb4bc8ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01b0c9ad8083eb3fe234cc46cf724f303c8c61fd4acc128db34e927f8ab1050c"></a>

## Next pages — tgw_security.forward_proxy_allow_all / b173bb4bc8ef / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-8b28679da2660b0558d699d5ae2f0748c9ecb55adf4146772965e524847ac111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f9d2d4b1215b3d19090890f2fab9ca69ae41e5f20c3552754e92d25d2ba4c20"></a>

## tgw_security.no_east_west_policy — tgw_security.no_east_west_policy / 1b84d21ee07a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.no_east_west_policy

<a id="canonical-3b2e7305066ac0cdc6dd6d0a2ffd0476908542f5c39c295dfa26a7641b4a07fa"></a>

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
no_east_west_policy = {}
```

<a id="canonical-c7dad0529fe8a27c459a2d6392001a6f18a0b5f8cee1849fa3311d07b616072f"></a>

## Direct properties — tgw_security.no_east_west_policy / 1b84d21ee07a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b65cbefe5975dec38ea1c361f6f1aad7f5a7f82b5869b7a0955917dabccd250"></a>

## Next pages — tgw_security.no_east_west_policy / 1b84d21ee07a / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b92573a556c559329a287b35b4f32c5d66eec9960192a781aa2a22d2fd1c33f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c88e70778fe1ab2509d4ebf4f50184b12080ea1dd1d74a3d940066062f2cf9bd"></a>

## tgw_security.no_forward_proxy — tgw_security.no_forward_proxy / e7d6abe6e95e / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.no_forward_proxy

<a id="canonical-410e6710d0c426ef3181f220c73abcd483f68f7329720910210602ebe8b21e72"></a>

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

<a id="canonical-6e51a4e5ba23ff210b7838ad6ec4bc61ee95090f6cee019400e879bddc0104a2"></a>

## Direct properties — tgw_security.no_forward_proxy / e7d6abe6e95e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-407b78d117892ffa9cedf3714b39b26b8ccde040599d566cd1d5e5ec8ed1ecfb"></a>

## Next pages — tgw_security.no_forward_proxy / e7d6abe6e95e / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-63369bcb606956a588e21960767782e5c2c987b0e333d43bc9a375812c2e85cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375578ad8b750e6b13744141fd60c87be389b353add59337642eb36d7e6497e8"></a>

## tgw_security.no_network_policy — tgw_security.no_network_policy / 041e0c32b935 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- tgw_security.no_network_policy

<a id="canonical-3ed71b92661bfefdd4075e100aac67ff5c8fd7a566ebe4c48374d12494572270"></a>

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

<a id="canonical-1bf7ce81edfab5b213bd4d7867542a46030dcab941f10cfa20f2d322c655dc77"></a>

## Direct properties — tgw_security.no_network_policy / 041e0c32b935 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7147d1219b3ffef52b7de1abd3bfb673ab28c26c16b2113306b1556ef2994ce"></a>

## Next pages — tgw_security.no_network_policy / 041e0c32b935 / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-7553cb4ac80cbd43a7a6cc1761d46be7559d8d7376ce03508ffdd968b305b1cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-001e5bffd2ab6b590cf3dd89415010f016e469be13a8b1328795dac6540123c5"></a>

## timeouts — timeouts / c085aa58edf3 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- timeouts

<a id="canonical-41f232eeea5a1df0a1a020c28d487fe0038c962e99dcf7fc2bec4cb43848e39d"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-99c6e0680fbbae3d288bd9a361273d097835c019b02d59486e0e98a85432ad1b"></a>

## Direct properties — timeouts / c085aa58edf3 / 3

<a id="canonical-aed28684150ff33be4d7796a79d3972a2d5895ae3c3bbafccd368c20c0686048"></a>

<a id="canonical-06198fe3c11b809e22c84264edca26cebab3454813e2e3a8637136a94390eb3a"></a>

## create property — timeouts / c085aa58edf3 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-cf79cf44a96147795771fda0526b1c0d95dd9935b8cec60de5b59be6c5e2b6f5"></a>

<a id="canonical-6e014b7f5d066aa5eeda4f89d91888036372ef922a71c2ba9e3591e4e1504a58"></a>

## delete property — timeouts / c085aa58edf3 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-36844258841e6e8daec6432c8c19deb7edd884c8dfa326170c0f1a2bf9502a08"></a>

<a id="canonical-22e75af0f3d9fd2e28efc3f732accc802cc74017812814844fbd71effa4aaee7"></a>

## read property — timeouts / c085aa58edf3 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-e5e6c84241cdff0d268ddf93b621331978ac3c7fe4da94a8555ca571fc64657b"></a>

<a id="canonical-0f31a7624ba65623c267619d69011250818f0410121fe4d799669116aa90851a"></a>

## update property — timeouts / c085aa58edf3 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-4773ded0b08b9f9014ecef8cdbfbde433fd7cc61363a73b7bb2bace17c4ba573"></a>

## Next pages — timeouts / c085aa58edf3 / 8

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4ad39f56e2ccab0f0d65fd424cb623c1d86728676adeee2208b4e0bb05ba7bb"></a>

## vn_config — vn_config / 9ebdc035d1a8 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- vn_config

<a id="canonical-2d9d64fb77b97c6f1294b2d91695434a49cbb6efba293b7a88503902791d127a"></a>

Type: `"object"`. single nested block, Optional.

Virtual Network Configuration. Virtual Network Configuration.

Upstream description:

Virtual Network Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
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
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
vn_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-186a2bab70c57aebfc9b087708fda79ef203c76a99cd00d04311348abbf3952d"></a>

## Direct properties — vn_config / 9ebdc035d1a8 / 3

- [allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27): complete subsection reference.

- [allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329): complete subsection reference.

- [dc_cluster_group_inside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-fb609dc791b3d1048ce594023281cdbcf0a207f1656de710ea962117ede022be): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-63828422b1f3043f756939e95d1857c4026f64a856a9f5b90e54537161daf8dc): complete subsection reference.

- [global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c): complete subsection reference.

- [inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c): complete subsection reference.

- [no_dc_cluster_group](resources--aws_tgw_site--reference--group-003.md#canonical-856b6b6cc7bb0db6a24803d3cca4cece48ab9d8f72abd0805e8c4dacf4124d3e): complete subsection reference.

- [no_global_network](resources--aws_tgw_site--reference--group-003.md#canonical-6894bde4d974e73e62ab179f77209c20d1a7090a847a0f0d7e08f439c2631f50): complete subsection reference.

- [no_inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-307b5c27a74a99b14ef331a84e64d2222da5fdb0113aedbeec7cf78fb17b8abe): complete subsection reference.

- [no_outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-19f38b9fb951ac9d47b805ec6640ed9487e8fc9463758825e27e609b3268288b): complete subsection reference.

- [outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f): complete subsection reference.

- [sm_connection_public_ip](resources--aws_tgw_site--reference--group-004.md#canonical-f320fa4d1906e5ee3d44cc33f3445eb2733ff209b91d622b368f1fdcc13dd93f): complete subsection reference.

- [sm_connection_pvt_ip](resources--aws_tgw_site--reference--group-004.md#canonical-0919a404c2ee92ed821488904172bb0250b874ac4d2a9ff45cbc802eb327c0c5): complete subsection reference.

<a id="canonical-99a2a9e3f36c8d008acc96b9c36558c0cab8132655541c476d497e7538bcb41f"></a>

## Next pages — vn_config / 9ebdc035d1a8 / 4

- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- [vn_config.dc_cluster_group_inside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-fb609dc791b3d1048ce594023281cdbcf0a207f1656de710ea962117ede022be)
- [vn_config.dc_cluster_group_outside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-63828422b1f3043f756939e95d1857c4026f64a856a9f5b90e54537161daf8dc)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.no_dc_cluster_group](resources--aws_tgw_site--reference--group-003.md#canonical-856b6b6cc7bb0db6a24803d3cca4cece48ab9d8f72abd0805e8c4dacf4124d3e)
- [vn_config.no_global_network](resources--aws_tgw_site--reference--group-003.md#canonical-6894bde4d974e73e62ab179f77209c20d1a7090a847a0f0d7e08f439c2631f50)
- [vn_config.no_inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-307b5c27a74a99b14ef331a84e64d2222da5fdb0113aedbeec7cf78fb17b8abe)
- [vn_config.no_outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-19f38b9fb951ac9d47b805ec6640ed9487e8fc9463758825e27e609b3268288b)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.sm_connection_public_ip](resources--aws_tgw_site--reference--group-004.md#canonical-f320fa4d1906e5ee3d44cc33f3445eb2733ff209b91d622b368f1fdcc13dd93f)
- [vn_config.sm_connection_pvt_ip](resources--aws_tgw_site--reference--group-004.md#canonical-0919a404c2ee92ed821488904172bb0250b874ac4d2a9ff45cbc802eb327c0c5)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
