---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-631e11939984e009d07a0b9c6dbbbc82713dc6d1becdeaf5c61940862d2092ad"></a>

## Property reference — Property reference / 691c9a728462 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- Property reference

<a id="canonical-b76056caceba14f10b03fbd9c026a11da8bec22d6e657d0e74a5a2377dad1924"></a>

## Direct properties — Property reference / 691c9a728462 / 3

<a id="canonical-0ac9c41df7cf2190f7c828ef44ea4ee0d4ab2e32cabdbfff38c7fa68fae4fdf4"></a>

<a id="canonical-92f8bf9d0e171e83e2a557391fe1d426f5a301e435e014a2d30cf18929fa2f87"></a>

## annotations property — Property reference / 691c9a728462 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48): complete subsection reference.

- [block_all_services](resources--aws_tgw_site--reference--group-002.md#canonical-f38ef4556d9454ee1fa3198b1de8b3eace6ebf064864c8599236cc65e500aa8f): complete subsection reference.

- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51): complete subsection reference.

- [coordinates](resources--aws_tgw_site--reference--group-002.md#canonical-fc40e9958bb6e7178fac69151232bfbf4397fc7a8a01886162591f52e00b9404): complete subsection reference.

- [custom_dns](resources--aws_tgw_site--reference--group-002.md#canonical-347174fd79b6c86f16e86abc884c54253528abca8e5888200327edfd66b4729d): complete subsection reference.

- [default_blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-2e91e3b2b5e727c6e40a4216181f74d84039efa17ee193877ca6cab3695ef316): complete subsection reference.

<a id="canonical-ddcd1139717f2d4600cd6e94ee4146d61921e8b44445d36fb40c274f886688c4"></a>

<a id="canonical-fc781a926d363e16cdce2dc054dff2ca0a2444ae474f9060a48d8bbdeb736d2e"></a>

## description property — Property reference / 691c9a728462 / 5

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-302989a4f6d557908c55a666fb3705955d8482c6561e1cab36d132b1cc5a1779): complete subsection reference.

- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db): complete subsection reference.

<a id="canonical-e95ada1ca86c3c5b5efad03febfaefe6e094023ecc91cc0c37ee1e14d2af9c7f"></a>

<a id="canonical-c2a6a28a1b2686c92cdf56a2526262de6eb2b072ca54da2a2c5f9e3120e8d08b"></a>

## disable property — Property reference / 691c9a728462 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-587dbfff4f4b620623f9182fe5ebc4e9b6275d917c8de9f95cadd9191acd80d9"></a>

<a id="canonical-102007453e2a9c7c0d17e3abe32c6810736358b03ff2581b0dc7781af6c611a5"></a>

## id property — Property reference / 691c9a728462 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e): complete subsection reference.

<a id="canonical-ac216435801a1a3ce16bfcaa7ead85613af16f64d2db1abbc555172be3dc3c0b"></a>

<a id="canonical-f71120bec23d85401185e1ecb7c8c24dd525509a7f74bfe764969f3cfb3e7910"></a>

## labels property — Property reference / 691c9a728462 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-5328bd4d13cd4b9bc6bf897e865e5f8b01e086064bde136bab9076d75b001faa): complete subsection reference.

- [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-fdb168a79ccc613abd6af4e13fca87d0cf0e20f031d3af2ea06cada88b87c750): complete subsection reference.

<a id="canonical-9c278855cba94834343a5719376c67c0769d8fdb5476699b6f985f50e97ed0a0"></a>

<a id="canonical-825b7ef160c5cde296cbd014725e83df047fce835f1ffdfade180fe65bd81527"></a>

## name property — Property reference / 691c9a728462 / 9

Type: `"string"`. Required.

Name of the AWS TGW Site. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-10ed9937cf5b7e976965bde4988f01d268420b318daa4d7ad5067727ec17df31"></a>

<a id="canonical-cb9e33052d6b66a3cbf8526836befd5e2648f4b6f0a4a94b8d7fe0b774be8a64"></a>

## namespace property — Property reference / 691c9a728462 / 10

Type: `"string"`. Required.

Namespace where the AWS TGW Site is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-c1d488ffe9bc7e32f5e0c9a74b5c69d528239d7d4ecd68080450066df847f5fb): complete subsection reference.

- [os](resources--aws_tgw_site--reference--group-002.md#canonical-b36c285d2a1299188f2809f555233d568667868bcfa2072e2d4e916468e5ee8c): complete subsection reference.

- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97): complete subsection reference.

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65): complete subsection reference.

- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-7076ed8e78694917e5f4df2c7aa41c7f303f575fbcc702ac483a7905e883932d): complete subsection reference.

<a id="canonical-56c105f82e5b3804d787b967ec605a02d4b533cb7f65cd0d1415a422f034dbe5"></a>

<a id="canonical-f3a1625b49bb67c17a86f008585acf3fe7d451813537129711c52a50e5f46628"></a>

## tags property — Property reference / 691c9a728462 / 11

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c): complete subsection reference.

- [timeouts](resources--aws_tgw_site--reference--group-002.md#canonical-7553cb4ac80cbd43a7a6cc1761d46be7559d8d7376ce03508ffdd968b305b1cf): complete subsection reference.

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d): complete subsection reference.

- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-6a136b15777bb8a98743650367f2764c51a7b782bf04d63d34a2defafee0ddf3): complete subsection reference.

- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-6d8e9aff961991a95f631650b225f1a27401e6966c47ee24453e16093a40748e): complete subsection reference.

<a id="canonical-584c5e22626056f5cbb527a64c169ab5604e64c9aab3522d69f26b46f7d9f219"></a>

## All schema paths — Property reference / 691c9a728462 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--aws_tgw_site--reference--group-001.md#canonical-0ac9c41df7cf2190f7c828ef44ea4ee0d4ab2e32cabdbfff38c7fa68fae4fdf4) |
| `aws_parameters` | [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-b8df39ec01df9616f8675ba90bbd827321ccbfe415aa31e0ae4c30299796519a) |
| `aws_parameters.admin_password` | [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-47fe691e9721e0d7e2e3f4df2200aa1fed77dbf18c335b5d9aca87226ab68197) |
| `aws_parameters.admin_password.blindfold_secret_info` | [aws_parameters.admin_password.blindfold_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-4adb868bb34fb595fef71192fa480a4637398817e54deff1d7e6d9b581bc2897) |
| `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` | [aws_parameters.admin_password.blindfold_secret_info.decryption_provider](resources--aws_tgw_site--reference--group-001.md#canonical-3066e9510c16101dd4008390fe8f3db63f94a3146816cb8904ef8ec10c2f0a7f) |
| `aws_parameters.admin_password.blindfold_secret_info.location` | [aws_parameters.admin_password.blindfold_secret_info.location](resources--aws_tgw_site--reference--group-001.md#canonical-902388d1b73e7ed0e01782f3517ada8992d60fdb5b694e9edc8e664b5697d500) |
| `aws_parameters.admin_password.blindfold_secret_info.store_provider` | [aws_parameters.admin_password.blindfold_secret_info.store_provider](resources--aws_tgw_site--reference--group-001.md#canonical-e0437efedf39e3a131f29a7e3ffc6323c4e0edf8c0444e3593fe7a0ae58274d4) |
| `aws_parameters.admin_password.clear_secret_info` | [aws_parameters.admin_password.clear_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-95a9fa8f85097f67a04129fd26fcfe2eae0d7566f10f62624676956ea703e322) |
| `aws_parameters.admin_password.clear_secret_info.provider_ref` | [aws_parameters.admin_password.clear_secret_info.provider_ref](resources--aws_tgw_site--reference--group-001.md#canonical-b9828acaaa98b63162c5b92b97aed389d69bc5b8278fa2ed9dacaffc878573d1) |
| `aws_parameters.admin_password.clear_secret_info.url` | [aws_parameters.admin_password.clear_secret_info.url](resources--aws_tgw_site--reference--group-001.md#canonical-b19f97b8fbcfcc528c808b756f37285bb0dd2fce073aefecf9a19c00852ada3f) |
| `aws_parameters.aws_cred` | [aws_parameters.aws_cred](resources--aws_tgw_site--reference--group-001.md#canonical-afbd740e212a335c1aa9e082e7f7f41ccbda3c87ff5716cb0bae22ce0bcb3a8d) |
| `aws_parameters.aws_cred.name` | [aws_parameters.aws_cred.name](resources--aws_tgw_site--reference--group-001.md#canonical-aa086c855a2fb5c590686b1ba17600cd5ee27e6cf53216523c460e318ab399c6) |
| `aws_parameters.aws_cred.namespace` | [aws_parameters.aws_cred.namespace](resources--aws_tgw_site--reference--group-001.md#canonical-346f0bbc0bb98b81e6dd7da3fc41dcbcd3284a19b73e6a79c6656148a97e891d) |
| `aws_parameters.aws_cred.tenant` | [aws_parameters.aws_cred.tenant](resources--aws_tgw_site--reference--group-001.md#canonical-edfebec6a589264264f8d0777379fa7400efb8e92ac54c65bf3cbd237978cb91) |
| `aws_parameters.aws_region` | [aws_parameters.aws_region](resources--aws_tgw_site--reference--group-001.md#canonical-f66dbbe11a5283416386c195913e774b4f5fdc207091ccb5a4106944007e6559) |
| `aws_parameters.az_nodes` | [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-fb2bfb41dc0d82bb84cf60336d281ff240717dc003051f4e096c0ac5eb296ef5) |
| `aws_parameters.az_nodes.aws_az_name` | [aws_parameters.az_nodes.aws_az_name](resources--aws_tgw_site--reference--group-001.md#canonical-463365584c3c9af0b8b4b4039e915e9b4808c85fc0ca0541b32ab6a792e373ea) |
| `aws_parameters.az_nodes.inside_subnet` | [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-b910dc8e743677ff887c22e8152161c24f45563eec30bb4ca18fa8c69b090805) |
| `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.inside_subnet.existing_subnet_id](resources--aws_tgw_site--reference--group-001.md#canonical-4de66bbd0c8714a337e588f55443bcaf9dce84cc9c37a7ad01357e82b5044edd) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param` | [aws_parameters.az_nodes.inside_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-97dd02a6aa4fec37729fea8b6d3ae177bb50268a962b48dbe60822a8d4d1d50e) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-716c070e124589efc5da18a08de6c22654bcd7d69909be2115c4bfc9598e484c) |
| `aws_parameters.az_nodes.outside_subnet` | [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-04960f4424d2b0d4b470c4108cfc828a45c706130ef83c88103d1c68bbcfb980) |
| `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.outside_subnet.existing_subnet_id](resources--aws_tgw_site--reference--group-001.md#canonical-b0d593e80d170c975b17abea8a26a811d14963e57c9666a85120b5b8abca6d89) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param` | [aws_parameters.az_nodes.outside_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-178d71bcf91e185d748bfc5e5131913de3c1a022c3a15f0b73b8a2560ebf4c65) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-15d5fcdf0af88ae551c6ce4295c0a276806d8e83d22dd9b0cb775452d2d10094) |
| `aws_parameters.az_nodes.reserved_inside_subnet` | [aws_parameters.az_nodes.reserved_inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-a4f81741908c1e0df6e57f9979abcc23eb421dbee371e9483f326f3c1d6681ef) |
| `aws_parameters.az_nodes.workload_subnet` | [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-e9a5aabc1dea32ce8e69ddbcede71fd7434c2766e1d86e0ec19bde91e2a92040) |
| `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` | [aws_parameters.az_nodes.workload_subnet.existing_subnet_id](resources--aws_tgw_site--reference--group-001.md#canonical-47656bb0b68796fac88a63a7b8cdb441c2bf73d4c45d46e63a16fcfb8812e3f4) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param` | [aws_parameters.az_nodes.workload_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-0b16c5d9dbbc4cf45c24b402ac63acd86c958febbef20a9144278a36b53026e4) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4](resources--aws_tgw_site--reference--group-001.md#canonical-1933ddd419baa21c6ce98d3bcc4e6aaf8314dda0a1a6471c7c0ab42ecd8d55f2) |
| `aws_parameters.custom_security_group` | [aws_parameters.custom_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-d0968f23e68e7a95578266883f00caedbe487b227ad438793aa54b48c01a608f) |
| `aws_parameters.custom_security_group.inside_security_group_id` | [aws_parameters.custom_security_group.inside_security_group_id](resources--aws_tgw_site--reference--group-001.md#canonical-ceef3b608f66430b9eb21ee06ae873678a5b777ffb65bbd1012d8a338aa17878) |
| `aws_parameters.custom_security_group.outside_security_group_id` | [aws_parameters.custom_security_group.outside_security_group_id](resources--aws_tgw_site--reference--group-001.md#canonical-73c04a5f6c4034636ba865a9e297869ff571fee2c6b32c885bcc862d50766e54) |
| `aws_parameters.disable_encryption` | [aws_parameters.disable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-03805be8ca9cfad4b1b4847a9fe4bdb0dff83133d50b93768a5fb5088e83cd43) |
| `aws_parameters.disable_internet_vip` | [aws_parameters.disable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-ae4f0f3c1ae9d386a14f7801b9f668f892477920d5ef7ebc005ea7e8bd30ee86) |
| `aws_parameters.disk_size` | [aws_parameters.disk_size](resources--aws_tgw_site--reference--group-001.md#canonical-9ce8a7de9ae5acd74cfc7c6988f38bd40de7a32ed3ab1a13397a437af4b802d6) |
| `aws_parameters.enable_encryption` | [aws_parameters.enable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-0ec3db85e4f1fd47f8640657a8e2a0f2b5ce43c53ccb5b4fb7e01e0dac474b38) |
| `aws_parameters.enable_encryption.kms_key_id` | [aws_parameters.enable_encryption.kms_key_id](resources--aws_tgw_site--reference--group-001.md#canonical-74f7898dfe6160c450abda0d9e33cf3536bbbee4cdc91f9549cf9265dd670ebc) |
| `aws_parameters.enable_internet_vip` | [aws_parameters.enable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-8bba561b6295cefd101a9efb4d3de7b9e60b889ead0b4e79725bbf15dfa0a78f) |
| `aws_parameters.existing_tgw` | [aws_parameters.existing_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-cdf4b580fd1cb25dc424c53737109e1d9f0185b0cf336a041e012b0d6a1e4eb4) |
| `aws_parameters.existing_tgw.tgw_asn` | [aws_parameters.existing_tgw.tgw_asn](resources--aws_tgw_site--reference--group-001.md#canonical-e3b1bbc23c815bd9f414fa1fa6c65d747a8d6d4d5cfeb40edd1aa76f699085fd) |
| `aws_parameters.existing_tgw.tgw_id` | [aws_parameters.existing_tgw.tgw_id](resources--aws_tgw_site--reference--group-001.md#canonical-4eb547e195d687d888811ec0f055f936b62f394e98e8a085cb8f9313f1496cd3) |
| `aws_parameters.existing_tgw.volterra_site_asn` | [aws_parameters.existing_tgw.volterra_site_asn](resources--aws_tgw_site--reference--group-001.md#canonical-c5c1152ed7ffc46dec87e95e587a97d5c6134c45567bd61721fcd85e45325e9e) |
| `aws_parameters.f5xc_security_group` | [aws_parameters.f5xc_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-54dc7e3625aa67ce268c0817dc425f18de17a8bcf921471f3da7ba24341646b4) |
| `aws_parameters.instance_type` | [aws_parameters.instance_type](resources--aws_tgw_site--reference--group-001.md#canonical-c74273294dbeae216bffdcd30fe0266bf850af20d81a3111dcc083f1cc686016) |
| `aws_parameters.new_tgw` | [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-5a5f926fed2cc993355bc753558edc5ecf309df9817bc17f3df85df1086d9b5c) |
| `aws_parameters.new_tgw.system_generated` | [aws_parameters.new_tgw.system_generated](resources--aws_tgw_site--reference--group-001.md#canonical-156529bf2718cd0dccd2afbffec20cc05742611a626fd1d9799d5386fd0f592e) |
| `aws_parameters.new_tgw.user_assigned` | [aws_parameters.new_tgw.user_assigned](resources--aws_tgw_site--reference--group-001.md#canonical-3c055260d3e6ece4fd16bec67061ce12670b6722073190ae9818f7371d6994ca) |
| `aws_parameters.new_tgw.user_assigned.tgw_asn` | [aws_parameters.new_tgw.user_assigned.tgw_asn](resources--aws_tgw_site--reference--group-001.md#canonical-45f4ae92e7113fa99cb4ab87ca978f3f5f160c47eeb246e2959fb3d8253dd5f0) |
| `aws_parameters.new_tgw.user_assigned.volterra_site_asn` | [aws_parameters.new_tgw.user_assigned.volterra_site_asn](resources--aws_tgw_site--reference--group-001.md#canonical-6e4859bffd629ffb333ef65785181d82f301cae7c9337b87f886d3ccb3262f96) |
| `aws_parameters.new_vpc` | [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-002.md#canonical-74ee03345ff222625074b7fec80e3ec8e5cc0917698c667e6651e10d1ff0131b) |
| `aws_parameters.new_vpc.autogenerate` | [aws_parameters.new_vpc.autogenerate](resources--aws_tgw_site--reference--group-002.md#canonical-4cee8ea06a44d4415191ee85a94ed81ddcd51e9e999c65b5ed80cc2377a13747) |
| `aws_parameters.new_vpc.name_tag` | [aws_parameters.new_vpc.name_tag](resources--aws_tgw_site--reference--group-002.md#canonical-66e25dd9807a1eef7f8e9be8438e6ad1eb50400ed454ce8abd5e1093cb481fc8) |
| `aws_parameters.new_vpc.primary_ipv4` | [aws_parameters.new_vpc.primary_ipv4](resources--aws_tgw_site--reference--group-002.md#canonical-229233bf45e0c11fed1d4fd8848a3773a11b524811271a09362024178aa4c2ac) |
| `aws_parameters.no_worker_nodes` | [aws_parameters.no_worker_nodes](resources--aws_tgw_site--reference--group-002.md#canonical-bc4bd9d19ebed9455f15464366fb78b8104af9a49d586a5aae92740fa9ef8cfa) |
| `aws_parameters.nodes_per_az` | [aws_parameters.nodes_per_az](resources--aws_tgw_site--reference--group-001.md#canonical-78069ab3841acd2f64c991a506a11a9130680e2d8945617d108ac0b7530f3635) |
| `aws_parameters.reserved_tgw_cidr` | [aws_parameters.reserved_tgw_cidr](resources--aws_tgw_site--reference--group-002.md#canonical-217df55ab72f020f9460ba67a6113a63d170f28bfba9445e25e61a582b449b8a) |
| `aws_parameters.ssh_key` | [aws_parameters.ssh_key](resources--aws_tgw_site--reference--group-001.md#canonical-a14f180aecced62e682152714b37b4919851254fabba835c0d057376389459b9) |
| `aws_parameters.tgw_cidr` | [aws_parameters.tgw_cidr](resources--aws_tgw_site--reference--group-002.md#canonical-4071d78aefea4a0553ebeb7b73e00ac04a701f5a01673bef668dc78a1cc2bd14) |
| `aws_parameters.tgw_cidr.ipv4` | [aws_parameters.tgw_cidr.ipv4](resources--aws_tgw_site--reference--group-002.md#canonical-df5d082102250f01d4b558812d0b3af26b59bf9ca09b76fa72e29468c81b3e8b) |
| `aws_parameters.total_nodes` | [aws_parameters.total_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-3c0984a6024b5552778ebf98575be07d5c379ab431f747ca1b1e9b9db4979b65) |
| `aws_parameters.vpc_id` | [aws_parameters.vpc_id](resources--aws_tgw_site--reference--group-001.md#canonical-4912b3f08118b46648c3b7c746cfaf4f39e633d78350ad60aa0bf3666ab64267) |
| `block_all_services` | [block_all_services](resources--aws_tgw_site--reference--group-002.md#canonical-92687f15442e205ba25c1be6414ce764510d97ca41eef734e413e6532ef940c5) |
| `blocked_services` | [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-f2bc47ab189f0fdf332ba3670d67a08ace7945d34b26e4a91ec052c1166f8a63) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-eeec99388c63976fa00e0b755a89d013ae9ba1f8206cfe7f40697a82f450c53a) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--aws_tgw_site--reference--group-002.md#canonical-621b89128455bb18a4d94ea95deb72d8284ab2c596b7064d38605f1a14b07115) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--aws_tgw_site--reference--group-002.md#canonical-75f4b773c776279b418ffcfc1fd9f25a2c0c6cb3690d6612e04d7eb350a7b66b) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--aws_tgw_site--reference--group-002.md#canonical-b060766320b33fe05c0fb93bf35318a72e658d649bdd2f5f7a037fee8bfd4ff4) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--aws_tgw_site--reference--group-002.md#canonical-10005e84005b24614f640470ac6ee73bcf6de88469aa6e490f29296ca08974cf) |
| `coordinates` | [coordinates](resources--aws_tgw_site--reference--group-002.md#canonical-7a716c2ca189d903d4f26a83e0777f518d0b015913fca00111bbb64fdf4b93b8) |
| `coordinates.latitude` | [coordinates.latitude](resources--aws_tgw_site--reference--group-002.md#canonical-a3babf83d9f7722da27b24ebd29ae0cf5e9e480d5cb976135435538c7d37d632) |
| `coordinates.longitude` | [coordinates.longitude](resources--aws_tgw_site--reference--group-002.md#canonical-63af6b5c77fc7f6763822f985c0bd5130c3f8c5b75fa46cc1e113b126f3ccd54) |
| `custom_dns` | [custom_dns](resources--aws_tgw_site--reference--group-002.md#canonical-f199fba880d5237e82ab901b5b39ce494c885d04ea7329b6f19bfa0d7c029dd4) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--aws_tgw_site--reference--group-002.md#canonical-e23470976891892ea64d4567a52a0b6b66851a97f1ef5d9acdd7fa773acd83f9) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--aws_tgw_site--reference--group-002.md#canonical-4240b44ac7d801c6a7541e85ed48008a6225acf6986633aa10b789a99fa41cce) |
| `default_blocked_services` | [default_blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-807b0250ceb5f671f89c9359c068c15526f4a67bc2148a562a2ef648ce982475) |
| `description` | [description](resources--aws_tgw_site--reference--group-001.md#canonical-ddcd1139717f2d4600cd6e94ee4146d61921e8b44445d36fb40c274f886688c4) |
| `direct_connect_disabled` | [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-dac64ee326fc8f93eab5e4cf182e098e5f4fa96a4a784d4e5780728193048c15) |
| `direct_connect_enabled` | [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-119113580de10ee0036ce9fc950181b4c727b9a197f59a0f781654dbcc0f0c0b) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](resources--aws_tgw_site--reference--group-002.md#canonical-d9ba56a7a1e52dfbf7b9a60123b96709fe536d4377064a2d9fda86ede6beb79c) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](resources--aws_tgw_site--reference--group-002.md#canonical-b79e1ca8e5852a5aecf0406582cbc9ce647d8913b7f3e60fb5fc5912e5cdcf20) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-119e032f4274ef3ffa7b6371aa7347e5c96c1ccf2b3b09ff0d4fd9720f578fb8) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_tgw_site--reference--group-002.md#canonical-726f2079f0b9942e52ac688f3aa00abd3b4560d8e407f9a2f440ba66d52d0356) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](resources--aws_tgw_site--reference--group-002.md#canonical-885eca1e38927dfb8b37d71883508731f8b4b365c3ea67d07b7b78d190bb4f79) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_tgw_site--reference--group-002.md#canonical-688d9d2bcef3f2ebd594b44c48b7dd51cf399c0ac25ac61885c5d74e72952024) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-52f854fec34897fd8ba709c668724e0fe511a4eb7c9a6258046c16a177895da8) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](resources--aws_tgw_site--reference--group-002.md#canonical-6cfd0737db92fc3f387d5c2b64751a5ca5d23fec84090d6216926acf7a7fc547) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_tgw_site--reference--group-002.md#canonical-232ba9595420745655e74675f55e3921efe87825970de3519b8e20e156b580e0) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](resources--aws_tgw_site--reference--group-002.md#canonical-f57e856261732350d7293e4785c6f32d771c00b8f22a7f2b44f4467b1bdc9355) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-e2b6081b27a7cc7457005544c6e2de52bfc22b959f0b5a15cdd94bb594596f62) |
| `disable` | [disable](resources--aws_tgw_site--reference--group-001.md#canonical-e95ada1ca86c3c5b5efad03febfaefe6e094023ecc91cc0c37ee1e14d2af9c7f) |
| `id` | [id](resources--aws_tgw_site--reference--group-001.md#canonical-587dbfff4f4b620623f9182fe5ebc4e9b6275d917c8de9f95cadd9191acd80d9) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-4f80ba7a872482f05d380accb2aefbcf42330e69bf9fb45e73b042727556fd53) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-b9d4c33bf3714ac57754bf2cb644a4613545d4588fa7958c1d39d93e008335c3) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-60b956b5d441b3cf8ac98642dde4c9624498ea0736a2546b2915bb750f76dcaa) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-5b5f04216e48cfec015b94fcbe7e8b1dd3a5751fefab338d7ad1debb6483dc63) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--aws_tgw_site--reference--group-002.md#canonical-1039210c2552d74cf4378dceeec97871f5c0cef09f1ad76326b7c42d0f75f49a) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--aws_tgw_site--reference--group-002.md#canonical-9e20e81d1ad9b0819fa3bf3be45d193e4c77a2a490cc310ce13b46c8f49567af) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--aws_tgw_site--reference--group-002.md#canonical-befd8b74dca385a5ec123562ef377086bed50e5d1cd82e0b37dc6c49ba63ff27) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-b1c187a1ac16b14917bfdf349d961da97426299c8d01a1910cee75eb08603366) |
| `labels` | [labels](resources--aws_tgw_site--reference--group-001.md#canonical-ac216435801a1a3ce16bfcaa7ead85613af16f64d2db1abbc555172be3dc3c0b) |
| `log_receiver` | [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-e0f27dc08c3abcaea1c9442696f74690ea29b1c9c46d4e98f604c2ee93040813) |
| `log_receiver.name` | [log_receiver.name](resources--aws_tgw_site--reference--group-002.md#canonical-008b8073fb391c29d62520ebcb5371d09a72beaf5ef6ec3ba57619aeb1eac08a) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-d482a55f2a1976d838e20a3c66a4022299c79fe97b805230732dcdafe3dadb00) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-7e1a5abd62b02327667ebe227464a05d3c9e2981377e98e6df4180ef8e300c9b) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-8244e24808d2c8df987ef85abc86cd1b6c7934aee1339ba9123c1a5c51d4ae48) |
| `name` | [name](resources--aws_tgw_site--reference--group-001.md#canonical-9c278855cba94834343a5719376c67c0769d8fdb5476699b6f985f50e97ed0a0) |
| `namespace` | [namespace](resources--aws_tgw_site--reference--group-001.md#canonical-10ed9937cf5b7e976965bde4988f01d268420b318daa4d7ad5067727ec17df31) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-f0eac2e3c39277cb8020e491b2312dd937083c1b4c31e4ab2fd4185e3a320250) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-d2d553899274a4da3b1dcbc1a2eeb3eccf2d887bc1a7cdad98183b5652ec5716) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-64fd5c393616b0dd40f9f354898b531a46770d4119891be460c66c002a5149b7) |
| `os` | [os](resources--aws_tgw_site--reference--group-002.md#canonical-8abc68ef5d6144c9ad6b9acf5479950114bd51be5f75ad84155e094839732c08) |
| `os.default_os_version` | [os.default_os_version](resources--aws_tgw_site--reference--group-002.md#canonical-54e2bdf5129eea60196e53e828c20526ef7afb1c5fb2630d65c4abef08b85990) |
| `os.operating_system_version` | [os.operating_system_version](resources--aws_tgw_site--reference--group-002.md#canonical-2c35bebd6b9e05ec8641085d1a61d6448432355f729aa5b15d0ca4f463d35166) |
| `performance_enhancement_mode` | [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-0f751d74fe8b13aa6d549d56934e712be9ebad56ad44b3020cbdc6480769845a) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-cd76f68ac940084410d8945b328530d8aa2b304511795f379220c6ace1e9bfbc) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-5f8ae67e816e9ae0208416825b268df9d22689b8af3ae9e335dcfa46296a4f55) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-9245cb5bc3f298a3903e40687a5e39a382fa3df257e779db2e2bdb14e05800d5) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-c6579b5591458cc7d60a8bc113b794d17148e03359b09c33197471e4e22249c5) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-ee5fe3a4cf21407ed8cfe6ea778739757bbbce12a47f02a5b1777f968df41641) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-727be0eead130993e8ad1f234c4ee4dfb8a4effa172c7713259d1a67c9d61496) |
| `private_connectivity` | [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0e44d8f367a3dd341058d0f9af047fd63771d872e0cfe26499986d5462c83478) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--aws_tgw_site--reference--group-002.md#canonical-d5e7365ab614ad33d0b97a329d5ff6de05b656b83fbe75c9249e5f6d68622b43) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--aws_tgw_site--reference--group-002.md#canonical-2e8a982009f56602e4376f38a368756b78a25007cc565047caaf415bbee167fc) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-07b3da90560b7935051a2bb8181a94283367c15b4ee869e047bc7b95d4e2289f) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-b0d4bfb91c2a372493933c046495aba56f1663046531b79adea1be5cf8f2f698) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--aws_tgw_site--reference--group-002.md#canonical-faa8d5c50ac1ce021302bfce8278aefdf1bd5ea849892e84f98580ebf0569d70) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--aws_tgw_site--reference--group-002.md#canonical-c4a2c241d523a5b606d2bb7d43a91a029b8a2876ec9ed3cdf91f1649afab8ead) |
| `sw` | [sw](resources--aws_tgw_site--reference--group-002.md#canonical-688e5f8d462864c3a29ffd340314a04cbfc5ff2479892e76f2568ef428bc909e) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--aws_tgw_site--reference--group-002.md#canonical-bde73d3e70e7ce34013056d02dcb38f064825997d17a0747978e2c9018304a6e) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--aws_tgw_site--reference--group-002.md#canonical-79e6354f473ec8351bb7d621cf866612679e58b72602a1331b549cad572e9bb9) |
| `tags` | [tags](resources--aws_tgw_site--reference--group-001.md#canonical-56c105f82e5b3804d787b967ec605a02d4b533cb7f65cd0d1415a422f034dbe5) |
| `tgw_security` | [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-76d50759ef620219023b57f8f9583cce4c1c3ed0cbe324c1912f1787d543ca47) |
| `tgw_security.active_east_west_service_policies` | [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-546c58bb71187c8cf2e0de09dcb75b70f03fa5325839abed1fb0ef75d35a44a9) |
| `tgw_security.active_east_west_service_policies.service_policies` | [tgw_security.active_east_west_service_policies.service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-0e1fef266b64f42d251c4c36e0d13548fcf2407bdc2b15459d59f9a383eed0bf) |
| `tgw_security.active_east_west_service_policies.service_policies.name` | [tgw_security.active_east_west_service_policies.service_policies.name](resources--aws_tgw_site--reference--group-002.md#canonical-1e286906324dce3b16e33b7d0f981a900e9849f02bf96721198ced1bc2d39497) |
| `tgw_security.active_east_west_service_policies.service_policies.namespace` | [tgw_security.active_east_west_service_policies.service_policies.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-bbcbe9abcc3c3b42d5567a35fb6ac48acba84e299e077e8915edfe7a5a73bfdd) |
| `tgw_security.active_east_west_service_policies.service_policies.tenant` | [tgw_security.active_east_west_service_policies.service_policies.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-63d2d93b6279117fd3b963635bb0dfe2b2948b87d5fbc3236d17a59a42a0e4b8) |
| `tgw_security.active_enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2b2541ae92baee7770f500699975dabe117ca66edff058e711daef0f8d25958f) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1f367bba8542016b47dadd0240aba83003c76685b48709f22e3d2d4ab95c7a49) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_tgw_site--reference--group-002.md#canonical-02f7daa8c1a2549edbe8bc7937098b8f1d9b768f1a59d847fef008c353a1f0d4) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-78b76582ef8c063e185cce4c0f9d01e09fe2e79ef2dc08a88d7f794f3ee393b3) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-fcca5b445a3e57dbb75f28ecd7d29ed88cf8309704e05c044074afdb9da7cdec) |
| `tgw_security.active_forward_proxy_policies` | [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-0f7d3b567a14a0bf518148dba55bef2c8f54e0f14e109024d84c816f76fde0e1) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-181a240d439c37f404932563ba560731d1a451fe7c56f4a95442c67b0a4fd55a) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_tgw_site--reference--group-002.md#canonical-f818aa3cb150cc044405927ba334aaa6ce46a7dc64234b262a320d2504c5de69) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-81704ea759db82e2279ceeef6a5fac7f95200bbf37d67b7fcfa000a9c447655c) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-bcb1109912269e993776ff0c81a55eeb8126a5883b9b8d8a2a6336ed1a146d51) |
| `tgw_security.active_network_policies` | [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-47c1bb875fb6eb216f7fcc4babee0a94f7db2d8a8e2753e86755ff0c439d8d7c) |
| `tgw_security.active_network_policies.network_policies` | [tgw_security.active_network_policies.network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-fa35526f956abd6e5bf2b15c578f3193119a846574b4c38b0c9de42dccd381ef) |
| `tgw_security.active_network_policies.network_policies.name` | [tgw_security.active_network_policies.network_policies.name](resources--aws_tgw_site--reference--group-002.md#canonical-8110123b6d20e7f94aae25c16aa22373bd3a2c08ce37f8b2f75474a147c73700) |
| `tgw_security.active_network_policies.network_policies.namespace` | [tgw_security.active_network_policies.network_policies.namespace](resources--aws_tgw_site--reference--group-002.md#canonical-ed4c08fd4e2ed13f698752d72e4539bbacd5de3719abc9902aae8b7c598cb9c4) |
| `tgw_security.active_network_policies.network_policies.tenant` | [tgw_security.active_network_policies.network_policies.tenant](resources--aws_tgw_site--reference--group-002.md#canonical-0b27c45cd488e53df91cb0eac6e9470ce63d51250e3b39d6cbf09ada0f311d8b) |
| `tgw_security.east_west_service_policy_allow_all` | [tgw_security.east_west_service_policy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-ac04e59c692ab1bd6136a083e7424b87d913a586a319177f4afe1cb301ac4d0b) |
| `tgw_security.forward_proxy_allow_all` | [tgw_security.forward_proxy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-bf0dd5e4aa6889bcc4dc270207578b0c1486fcaa37be5c9c70777d813ed6c1b2) |
| `tgw_security.no_east_west_policy` | [tgw_security.no_east_west_policy](resources--aws_tgw_site--reference--group-002.md#canonical-3b2e7305066ac0cdc6dd6d0a2ffd0476908542f5c39c295dfa26a7641b4a07fa) |
| `tgw_security.no_forward_proxy` | [tgw_security.no_forward_proxy](resources--aws_tgw_site--reference--group-002.md#canonical-410e6710d0c426ef3181f220c73abcd483f68f7329720910210602ebe8b21e72) |
| `tgw_security.no_network_policy` | [tgw_security.no_network_policy](resources--aws_tgw_site--reference--group-002.md#canonical-3ed71b92661bfefdd4075e100aac67ff5c8fd7a566ebe4c48374d12494572270) |
| `timeouts` | [timeouts](resources--aws_tgw_site--reference--group-002.md#canonical-41f232eeea5a1df0a1a020c28d487fe0038c962e99dcf7fc2bec4cb43848e39d) |
| `timeouts.create` | [timeouts.create](resources--aws_tgw_site--reference--group-002.md#canonical-aed28684150ff33be4d7796a79d3972a2d5895ae3c3bbafccd368c20c0686048) |
| `timeouts.delete` | [timeouts.delete](resources--aws_tgw_site--reference--group-002.md#canonical-cf79cf44a96147795771fda0526b1c0d95dd9935b8cec60de5b59be6c5e2b6f5) |
| `timeouts.read` | [timeouts.read](resources--aws_tgw_site--reference--group-002.md#canonical-36844258841e6e8daec6432c8c19deb7edd884c8dfa326170c0f1a2bf9502a08) |
| `timeouts.update` | [timeouts.update](resources--aws_tgw_site--reference--group-002.md#canonical-e5e6c84241cdff0d268ddf93b621331978ac3c7fe4da94a8555ca571fc64657b) |
| `vn_config` | [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-2d9d64fb77b97c6f1294b2d91695434a49cbb6efba293b7a88503902791d127a) |
| `vn_config.allowed_vip_port` | [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-3a6f8b9cae264f949303180413bbf7a52575cdd9ebc8c79c1d9ddb01f9f0c927) |
| `vn_config.allowed_vip_port.custom_ports` | [vn_config.allowed_vip_port.custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-a4a61d6c084e6d4861b8dacf869d4a92901320edb1e37b1c0fd7a086bbd6a455) |
| `vn_config.allowed_vip_port.custom_ports.port_ranges` | [vn_config.allowed_vip_port.custom_ports.port_ranges](resources--aws_tgw_site--reference--group-003.md#canonical-0ad6517e5187bbdac18c59b1cb66d3d2431abc45f4f2fc3d136b9f105423be0a) |
| `vn_config.allowed_vip_port.disable_allowed_vip_port` | [vn_config.allowed_vip_port.disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-ca6cc2ac5593cd01cc77981d09bd43cd805c49dc8eccfcdfaedbbd21bd48f762) |
| `vn_config.allowed_vip_port.use_http_https_port` | [vn_config.allowed_vip_port.use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-ca2d97ad209edeb8d550b7d897195e30dd36efa8b9b275f425ab079b0f61ef4f) |
| `vn_config.allowed_vip_port.use_http_port` | [vn_config.allowed_vip_port.use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-613a685dac6043a8f4f6ab107b38b57e4348dee7b6258b628215cd6aba386d2b) |
| `vn_config.allowed_vip_port.use_https_port` | [vn_config.allowed_vip_port.use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-2ca78f0ff3515f2b3a1f5cd5ed2bcc2c4f4e4542bc11b04b893ba6e32f65faf2) |
| `vn_config.allowed_vip_port_sli` | [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-ad51201dfcdfc13578233e5cde7d14f603c24abb79768cfdf2812a4312c849b4) |
| `vn_config.allowed_vip_port_sli.custom_ports` | [vn_config.allowed_vip_port_sli.custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-4b1d99a9291c2d1ec427d94b06de02845bd1299ade1186d0fb26341a3096c14c) |
| `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` | [vn_config.allowed_vip_port_sli.custom_ports.port_ranges](resources--aws_tgw_site--reference--group-003.md#canonical-0dd27f44b80cbd3f01cd429789897d3931878363e88a8bd2b215ee8e969028c9) |
| `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` | [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-2a553c75361f0add0104d8602b968fadd30eb0ed204339022383847cfc156b8a) |
| `vn_config.allowed_vip_port_sli.use_http_https_port` | [vn_config.allowed_vip_port_sli.use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-5b4faa612541843bf9e4df1263f8669f12e2497b815fadfdb2ae0de6e5bad6d0) |
| `vn_config.allowed_vip_port_sli.use_http_port` | [vn_config.allowed_vip_port_sli.use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-8c4aedd707bb0c993bc113bf69942b21e18e89e9b060fa4b43b48ef87635e1ea) |
| `vn_config.allowed_vip_port_sli.use_https_port` | [vn_config.allowed_vip_port_sli.use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-cbb22e0d5930a60ccb875d5e6955225e0014a779e1c4b6d2467275b1321552a1) |
| `vn_config.dc_cluster_group_inside_vn` | [vn_config.dc_cluster_group_inside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-32b5d510ca4d7dd340a778452f046ebae147467f1eaeb52fc8b42251ae4f7953) |
| `vn_config.dc_cluster_group_inside_vn.name` | [vn_config.dc_cluster_group_inside_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-33a76fbf3b2176dff8855a7abe1503cd32af824e2d543d1ea1475b8f9c98936f) |
| `vn_config.dc_cluster_group_inside_vn.namespace` | [vn_config.dc_cluster_group_inside_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-5aea276ec576d450ec1352ee0393948cb1cf60a57696fff9dfc09a75c431b66a) |
| `vn_config.dc_cluster_group_inside_vn.tenant` | [vn_config.dc_cluster_group_inside_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-0b73b1aabbf68a4f05b02d1baa79805eee6a8bb4e4233cc446b5818b5a9493f1) |
| `vn_config.dc_cluster_group_outside_vn` | [vn_config.dc_cluster_group_outside_vn](resources--aws_tgw_site--reference--group-003.md#canonical-0b97b402fe8921d8338faa719c984a3c38e46efbd16c1e78662e20340f66ca50) |
| `vn_config.dc_cluster_group_outside_vn.name` | [vn_config.dc_cluster_group_outside_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-631a660b961c0a07c29866c641aaede62b75957b7c7cef385252ca4ce81e2633) |
| `vn_config.dc_cluster_group_outside_vn.namespace` | [vn_config.dc_cluster_group_outside_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-20813f8aacf05df081eb1b48f584619ec980826dd15b71a5b33109e28025d9cc) |
| `vn_config.dc_cluster_group_outside_vn.tenant` | [vn_config.dc_cluster_group_outside_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-daee89284ae1e9d08d0d825dbe7c7544944d5f375448001fc500bff13cb3306e) |
| `vn_config.global_network_list` | [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-9be859efac131b19296b21f17afe56fa2ee646558ed8e588e6da15241b9bf68c) |
| `vn_config.global_network_list.global_network_connections` | [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-37ea0243159277f22c59c0c97a84c938f112aa88f7c72487d264cbe2c5b886aa) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-45056b4c42ef5d1bed309f07f9249cffe13ae4396c1d24150d74c1f518241a4e) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-e52ca3b1482b677725071f34650551ffac08cda0acf90380db7acdc3f0ab12d0) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-230fefb0eabe1fd4d7b640e5ea954a6a1f365c2a1c66bc6392dc2dbc8faf4593) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-a1f80918bc704635371ead6c736b70c21018657a837c30ce15e479f02a85f60e) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-6660550b14502c1f18b78136ee16c5c7b4c8318bdbc926889d44bceb0195b282) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-4b220696ec5e61951cba0de0b449596e66c75705eb5910e5b6664dabfbb3b53f) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-0dbd6d4cc42a533d46891f925bb369b8e105d1a7bf9081c8083dc6fc99d7c4d2) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_tgw_site--reference--group-003.md#canonical-3d0d1c1c9e0557b70a1856e764b807e0cb10c23296fd74ffec02fcc685542b89) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-e735588c4b37f5d440c48cc3504c46265716e6f132fff7f1c0a71042fc367d59) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-34667d990cc7695a7d55bb87909d9c354cbc729d541eeca24bfd2a7b8071790b) |
| `vn_config.inside_static_routes` | [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-e9d4c945d12a5181370ebf86c118bdd8e4f7fbc7d13111b80e88b55fca9eee7d) |
| `vn_config.inside_static_routes.static_route_list` | [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-bf19c8f28b7dae1f59db06764aed1a0773128a91f6e449a46554812b06ca021e) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route` | [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-141186e7f666d11a3865c5cfc02f44fa04e72a2b46b8521306440bde02914f21) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.inside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_tgw_site--reference--group-003.md#canonical-572b299eac60cccc24853cdd3b6892a8ba2a515cbf6e1964b47a191eb5eccde0) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--reference--group-003.md#canonical-610d9be4e688527a47eabb299762d250e4838b41b5878d842a54fc8e866d1272) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1a5378682ba319746be5cfbd653c6626a4c5ba467a22fd725a77a5e19f5d8a28) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--reference--group-003.md#canonical-c442b7d2c1e6a739aa08398a51eb6169af8e23191f60666a95ee2c1bd7b6de45) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_tgw_site--reference--group-003.md#canonical-cc0170fdea027451204c999f65449366aa6a9720f33c0e1bb797b7a39c5e3f92) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_tgw_site--reference--group-003.md#canonical-409ced29e2a7d0ad5c3a0f8f060ffd3534e2ea34935dff79b70ab6ec89781ad9) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-202bc5bcabe4198cec8dada5e13538496d823d23f83dc8779ca00a4757724ba7) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-cd4b8565bce845d3facb25c085ebf37b5d44b326b01ba4b13e41d21fb6d730f1) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_tgw_site--reference--group-003.md#canonical-e9150d09a3fcc16036cf3a8a6e4758c245190348a8dc598372b839b16cb2708e) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-6fdde452222857aaf8b5591fb0a5bc6fe84a178ee96748079c9bcce287f1622d) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-b06fb45f3fbd999f7b657fb0201a9b299e8dc2545d5ba97f7aafff1e76c178e5) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-fa98f7e17063309afd487927be49b9a72b18596a8ef629f336f48cd261b21d97) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_tgw_site--reference--group-003.md#canonical-950a7b17480e5e5da9a6acfafc3bae9976975acb9b3f1f3b48ea46e56412af7d) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-6677f3ad5914da54fa17c8fc7c4f984e6d8ca31468552a5b067b343bebb319e9) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_tgw_site--reference--group-003.md#canonical-2347ace543598f690061418e3114ca656abd2835aef1d4141498e536fde05555) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-4badde14e648b4191aa9707927d54195111403f3b0d1fab305e45be7325e6353) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_tgw_site--reference--group-003.md#canonical-ca08830ceb6ffbed250e9d2778d171926c21078bd4a80af02ac6bb38a8c1e889) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-46d9ae94dde00ec53495c90c2fcfbee6776bfb73a02401627885e96232c1c44a) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_tgw_site--reference--group-003.md#canonical-ad798ae0ad4ba2228a678866b1cbf19b999cabe5ae3ea83c1176f50aff80169c) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_tgw_site--reference--group-003.md#canonical-59f88511c1940c8fa795f8e6203e361f9a13167488f096dbd72bf61244f2a5d3) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-003.md#canonical-f301e050e8b2a057dd0ce5e28e620dab933507eb388ce16a0b992c44a589b345) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-2c6b9bdcd7ed16bc708705eda94c8cb275e0f7484aa0f5f38f874e4499e0f09d) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_tgw_site--reference--group-003.md#canonical-fa671e65d6fe8bfbb3fcffd84eea55b564a30c6a0c62b9ea7b8f8c7104ec52b9) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_tgw_site--reference--group-003.md#canonical-7cf6bd6a144249fe1ee478d11e0a5fcd2284984b88679adcc552e91eae6ed3fc) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-d8c496090d0285aedc798d5a8ea8da20df42e86dedebaf554b72656909ad1043) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_tgw_site--reference--group-003.md#canonical-d426422583334276659c5d9f22a0cee49e81a401c096f546c4b5ce7998ae82b1) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_tgw_site--reference--group-003.md#canonical-7f1ef5e41197fc018f3679821dd6ccbd30d9ac9aaba962148e4ca2864503f934) |
| `vn_config.inside_static_routes.static_route_list.simple_static_route` | [vn_config.inside_static_routes.static_route_list.simple_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-f9c87c320f33f80c511b8cd297c5d1090e2bda0d656418e2e7d2efdedc556e6f) |
| `vn_config.no_dc_cluster_group` | [vn_config.no_dc_cluster_group](resources--aws_tgw_site--reference--group-003.md#canonical-9710f656b3aef731f795ba0202544316d932aa24090d85645f324adbc0e365e2) |
| `vn_config.no_global_network` | [vn_config.no_global_network](resources--aws_tgw_site--reference--group-003.md#canonical-a987642426f88fce3124633f56616d897e642eb771d8c9e239bbad1c0b7be533) |
| `vn_config.no_inside_static_routes` | [vn_config.no_inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-e375cdada23bca0526a524b50d8350af8e96e925797f60c37d5ba7ab959740bd) |
| `vn_config.no_outside_static_routes` | [vn_config.no_outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-c2df28f72af3d73a7a97eeaa0310c808eeda847d095e385bc98235a4d5a4fe6a) |
| `vn_config.outside_static_routes` | [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-ae75c78654be04a04bb78c2e12956605a360a6da1059caa087e36869721ee7bc) |
| `vn_config.outside_static_routes.static_route_list` | [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-2574ced6230122eec924c58971ef5a92ae3973866a065cf35d88a6d3c54dd9cb) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route` | [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-c99557f3aecc77b94fa554eff001598cee0aa1431943086473f1a9bcb2133f44) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_tgw_site--reference--group-003.md#canonical-0356635cc6e94e850a41a0613ff6d3bc7eeabfc9e29fd04c674cd88d44ea9bf7) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--reference--group-003.md#canonical-7eab1046c2093f6e8a2c4830fed40c59809444adebe8a97688035749f0431aa4) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-f929093866295e582cc6ddb53980e2feba810b8f03bc399dc0a66c5cbfe8459d) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--reference--group-003.md#canonical-19fb2219bd74230573241ab403a66ce7e539a420a60ec015a680c1b408609715) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_tgw_site--reference--group-003.md#canonical-e90bbdfe1bf893fd4d319b5ee61f13528a408b7fbe3c1021cb2c033690a0d98f) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_tgw_site--reference--group-003.md#canonical-0c1fa5b5b610f34dcf2f315a21051f9fa7d3615ee190d22a66c34c252f12851e) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_tgw_site--reference--group-003.md#canonical-3c2991b34c5f19b034e9ef8cf785b02a2803416fc496917f993f23bfec25face) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_tgw_site--reference--group-003.md#canonical-42847b715c59efd84e8f90e51be2555503d449880c4d2a2b0dbf1f3e6f1fb4d1) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_tgw_site--reference--group-003.md#canonical-3b86f7c8f475cf5d9d5eb5f315659db2a0bb1a16517787fc07d8f5f98b54b3ee) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-e5ba286e4dc29ebefe41341ab7ef302cf9733e34198d930a71827a8f180d1f4e) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-84cce9475c09d37e14b426fb590cda20187317b6689b3a125ff30973b6877454) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-1afa34227b64774358f4c6e24f04b388649bb7c04ddecfbce973f212b30fe780) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_tgw_site--reference--group-003.md#canonical-1e330e171479ef688aca97d6f3d87738be6f6c62329b59c0b30ec147d4be57d7) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-d327fbd44591de809be627924f7b66c17ef88823d75ae436a2824cb51c936e25) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_tgw_site--reference--group-003.md#canonical-c0aa458f3714b91a916a4bf352eeb257c4111282b25e1515dc216c6c88c7638c) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-0a4b8f0d71bc907a1f25c361776aa1c65f5e0074b8ca8c8bf67ceae32f6a91a6) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_tgw_site--reference--group-003.md#canonical-c63983ec42e1d46c32d1437a2a6d53f45c17f64fbd941e3b89aeb13e607f01d5) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-d4b199f18d955fd86b070279e4b5d070aa5d70fea944727fde8e175881b5e61a) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_tgw_site--reference--group-004.md#canonical-daedee21f2b82700145605c74a12ee3883e52dcb9d0cada80f2ea68f694092e1) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_tgw_site--reference--group-003.md#canonical-5a2fb32d22e4ff942f51852df057470036d84734236d8a953bcc9650979a3ea3) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-cffde52c11e4ec77a0ffea28fcebd409e94ba14e97284057a8f3134dee329ef1) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-3bb4c4b68dbb1b8428aa0f2e2dada765e72f4eab37dc7d0cb1796cc258cd67ca) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_tgw_site--reference--group-004.md#canonical-6e12564e2e26cfce4ac2889ffbb118c8c5551e2121de7cc07fd3f3eb3443c9a9) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_tgw_site--reference--group-004.md#canonical-d32b7ca202631164c4c8c108461db5c41bc3e675b00d2498f89849a007d3645b) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-4bbba6f398597b9013af3d3e77324a2b0eca83fcbc8cc97dff409822d66446ce) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_tgw_site--reference--group-004.md#canonical-1e6a24e125335c225e163e83a846e4cc14219ddfa4090af2634d82e4306c9ebb) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_tgw_site--reference--group-004.md#canonical-d5d7b44bf5c7f514bc820ecf98b92987ffc639042b047e6795380ec138dfb1e9) |
| `vn_config.outside_static_routes.static_route_list.simple_static_route` | [vn_config.outside_static_routes.static_route_list.simple_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-e8a1ff8dfedac86e4affc2ad0abc5bd31b3150e158ead546be5dbbcf20b54e89) |
| `vn_config.sm_connection_public_ip` | [vn_config.sm_connection_public_ip](resources--aws_tgw_site--reference--group-004.md#canonical-62934311464c78654206088afcad0e66574bc52752dfbd07e6237934a5b3b17a) |
| `vn_config.sm_connection_pvt_ip` | [vn_config.sm_connection_pvt_ip](resources--aws_tgw_site--reference--group-004.md#canonical-c1a93302970bffa7d1c78b6891b41c0cfc2bc5fe9c1d75f09afb4763c25ade09) |
| `vpc_attachments` | [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-bce0d35c2b76f1d20d33b4f67b533b620e7758adb0eae939a5d0a9d989834d5e) |
| `vpc_attachments.vpc_list` | [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-64e51b24a01c68a625a6ce81995a1274e5be7c48822cf59ddfe5e2282aa55e68) |
| `vpc_attachments.vpc_list.labels` | [vpc_attachments.vpc_list.labels](resources--aws_tgw_site--reference--group-004.md#canonical-e41605155bcea2cdd94df0e6f184d47db2e624175241c3f3941c25ee6f9f2822) |
| `vpc_attachments.vpc_list.vpc_id` | [vpc_attachments.vpc_list.vpc_id](resources--aws_tgw_site--reference--group-004.md#canonical-e0043c45eb68468b70a72db18dd59949567ad5a7309ad2bbf28db29944ac2fab) |
| `waf_signatures` | [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-ea9e313c1fd8aba779020efd5a0f862c453f8cbffa600e7e2d723416b917abed) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--aws_tgw_site--reference--group-004.md#canonical-37f040218e8a7bd1a144d6c14bba7ab887b37a8791ed9e1ecd6fb8198cb87ae7) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--aws_tgw_site--reference--group-004.md#canonical-06472057b4314a3c5613b1ee1254a6d14e3c15386a27d2f6f2ec773a252689a3) |

<a id="canonical-b81b31988c0c8217fb1379d45c73ca2e2eb4a87002302073411f266cc3f16cff"></a>

## Next pages — Property reference / 691c9a728462 / 13

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [block_all_services](resources--aws_tgw_site--reference--group-002.md#canonical-f38ef4556d9454ee1fa3198b1de8b3eace6ebf064864c8599236cc65e500aa8f)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-e8d8b9fb2aee79c8f7b4a5747a0783a249ff6ae310ab6774a3c05184d3e24c51)
- [coordinates](resources--aws_tgw_site--reference--group-002.md#canonical-fc40e9958bb6e7178fac69151232bfbf4397fc7a8a01886162591f52e00b9404)
- [custom_dns](resources--aws_tgw_site--reference--group-002.md#canonical-347174fd79b6c86f16e86abc884c54253528abca8e5888200327edfd66b4729d)
- [default_blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-2e91e3b2b5e727c6e40a4216181f74d84039efa17ee193877ca6cab3695ef316)
- [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-302989a4f6d557908c55a666fb3705955d8482c6561e1cab36d132b1cc5a1779)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-ac82732ae4c6ed1fc3fe57328e909a84eec3d10b2e429345d448109c2a5fb3db)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-5f0e8beb13147df524b7d641760ac19c7109b05bd03b9cd7f614d7b329893a3e)
- [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-5328bd4d13cd4b9bc6bf897e865e5f8b01e086064bde136bab9076d75b001faa)
- [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-fdb168a79ccc613abd6af4e13fca87d0cf0e20f031d3af2ea06cada88b87c750)
- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-c1d488ffe9bc7e32f5e0c9a74b5c69d528239d7d4ecd68080450066df847f5fb)
- [os](resources--aws_tgw_site--reference--group-002.md#canonical-b36c285d2a1299188f2809f555233d568667868bcfa2072e2d4e916468e5ee8c)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-4303abda634caa4786db98cdc8377f6b9ba8897cbe625e6d6e5da785aaec3f97)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-167e0e87b06b0a771760b8cbb360bf5f698651ea97f72ae4267354bc6877cb65)
- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-7076ed8e78694917e5f4df2c7aa41c7f303f575fbcc702ac483a7905e883932d)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-5ff04c1933b8606e7ed6ccd8a5bd654c5b878146645f88ddcebe941fe708934c)
- [timeouts](resources--aws_tgw_site--reference--group-002.md#canonical-7553cb4ac80cbd43a7a6cc1761d46be7559d8d7376ce03508ffdd968b305b1cf)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-6a136b15777bb8a98743650367f2764c51a7b782bf04d63d34a2defafee0ddf3)
- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-6d8e9aff961991a95f631650b225f1a27401e6966c47ee24453e16093a40748e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f53d240bf3eadb95fd589814ec66d2648958ebfa85f728b217ca93dee36b5c2"></a>

## aws_parameters — aws_parameters / da1487af8f02 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- aws_parameters

<a id="canonical-b8df39ec01df9616f8675ba90bbd827321ccbfe415aa31e0ae4c30299796519a"></a>

Type: `"object"`. single nested block, Optional.

Setup AWS services VPC, transit gateway and site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "az_nodes",
    "instance_type",
    "ssh_key"),
  validators.ConflictingObjectAttributes("custom_security_group",
    "f5xc_security_group"),
  validators.ConflictingObjectAttributes("disable_encryption",
    "enable_encryption"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip"),
  validators.ConflictingObjectAttributes("existing_tgw",
    "new_tgw"),
  validators.ConflictingObjectAttributes("new_vpc",
    "vpc_id"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "nodes_per_az"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "total_nodes"),
  validators.ConflictingObjectAttributes("nodes_per_az",
    "total_nodes"),
  validators.ConflictingObjectAttributes("reserved_tgw_cidr",
    "tgw_cidr")}
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
  "x-ves-oneof-field-deployment": "[\"aws_cred\"]",
  "x-ves-oneof-field-encryption_choice": "[\"disable_encryption\",\"enable_encryption\"]",
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]",
  "x-ves-oneof-field-security_group_choice": "[\"custom_security_group\",\"f5xc_security_group\"]",
  "x-ves-oneof-field-service_vpc_choice": "[\"new_vpc\",\"vpc_id\"]",
  "x-ves-oneof-field-tgw_choice": "[\"existing_tgw\",\"new_tgw\"]",
  "x-ves-oneof-field-tgw_cidr_choice": "[\"reserved_tgw_cidr\",\"tgw_cidr\"]",
  "x-ves-oneof-field-worker_nodes": "[\"no_worker_nodes\",\"nodes_per_az\",\"total_nodes\"]"
}
```

Terraform syntax:

```terraform
aws_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-56f045a18237b9cbd34811601d653641d89af212fd5278d831acd1d69413c388"></a>

## Direct properties — aws_parameters / da1487af8f02 / 3

- [admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-88f8143d60eb6d965531df056eb61c3c957a4e86ddbed6d734b8eb35169a4452): complete subsection reference.

- [aws_cred](resources--aws_tgw_site--reference--group-001.md#canonical-e9e60ee6cd97042783e932829df6acea0fc6b60f4182d9f8d517f9c7228511e0): complete subsection reference.

<a id="canonical-f66dbbe11a5283416386c195913e774b4f5fdc207091ccb5a4106944007e6559"></a>

<a id="canonical-31f6e3c500d7e17582e43557a49d38a7a7aef28dd69255b07acdaa6232ee3413"></a>

## aws_region property — aws_parameters / da1487af8f02 / 4

Type: `"string"`. Optional.

AWS Region of your services VPC, where F5XC site will be deployed.

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

- [az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b): complete subsection reference.

- [custom_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-a2dbf3068620cc2ba0398e09b0bfc71396676e57c14fdee9a4a4b6a35bfcb755): complete subsection reference.

- [disable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-6f4daa9c6f5acfe04beb470e293d168eda234f99982312b18c311291fe554551): complete subsection reference.

- [disable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-94343b10b6742d3b41a087823b366f63704533082ddd021cf173675cefd4954f): complete subsection reference.

<a id="canonical-9ce8a7de9ae5acd74cfc7c6988f38bd40de7a32ed3ab1a13397a437af4b802d6"></a>

<a id="canonical-f75b6d77e46493c36b167bfe889aa1a4e4ac4eae87422fdeb67fdc79ab4c5c95"></a>

## disk_size property — aws_parameters / da1487af8f02 / 5

Type: `"number"`. Optional.

Node disk size for all node in the F5XC site. Unit is GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-9fcf4edf24bb9ac5a87ed674acb3693271da466bc5df4c2dc505b7bf598f61c9): complete subsection reference.

- [enable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-00ffbb634b9197ae73007e874f86eb0a47a982e3437a4789288f9a7ed47bf3f7): complete subsection reference.

- [existing_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-13688fa212f3629952f7851b68ae9ce15d01ad9608fffecf507600292ce14c98): complete subsection reference.

- [f5xc_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-b34cebbd4a5b43e28221b06df679ba86048fdab745ea758667b7f8d8473e73e0): complete subsection reference.

<a id="canonical-c74273294dbeae216bffdcd30fe0266bf850af20d81a3111dcc083f1cc686016"></a>

<a id="canonical-fab11e844adf2a57a6ce48cfd360d8fbd57a0dd28ad19516bda426c2e54bbfd8"></a>

## instance_type property — aws_parameters / da1487af8f02 / 6

Type: `"string"`. Optional.

AWS Instance Type for Node. Instance size based on the performance.

Upstream description:

Instance size based on the performance.

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

- [new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-2112fc494e90b1dc8d39e4b1b6f2967d6fe4a8cc5da84ec25fff2ba07bb87f52): complete subsection reference.

- [new_vpc](resources--aws_tgw_site--reference--group-001.md#canonical-2823b6f937a1792c17d1a89fa7317dd7838c44ea0e8ac666805abf0d98ea773a): complete subsection reference.

- [no_worker_nodes](resources--aws_tgw_site--reference--group-002.md#canonical-a989844489c66d4e7faae9b8073925c332423fe87f9ed9161f123294f9f2b651): complete subsection reference.

<a id="canonical-78069ab3841acd2f64c991a506a11a9130680e2d8945617d108ac0b7530f3635"></a>

<a id="canonical-5588af80dbda673828081ad5e52412d7bd928b163628e0fd9f091c2587527d04"></a>

## nodes_per_az property — aws_parameters / da1487af8f02 / 7

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [reserved_tgw_cidr](resources--aws_tgw_site--reference--group-002.md#canonical-d99f8f19c6a7e99ee542ccfb2979873c26ae54bd8d9b7fa065b15906cbe89931): complete subsection reference.

<a id="canonical-a14f180aecced62e682152714b37b4919851254fabba835c0d057376389459b9"></a>

<a id="canonical-83a235841ab634e0200c932a4211ac51c754c6cb62c8e47368663ee9643122e3"></a>

## ssh_key property — aws_parameters / da1487af8f02 / 8

Type: `"string"`. Optional.

Public SSH key for accessing nodes of the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tgw_cidr](resources--aws_tgw_site--reference--group-002.md#canonical-dda6f57721acd4eeedf2a9b8e0df09e5fe75a64f91d5fb5515a8f7664c2a2e98): complete subsection reference.

<a id="canonical-3c0984a6024b5552778ebf98575be07d5c379ab431f747ca1b1e9b9db4979b65"></a>

<a id="canonical-8aa5a5eb26198bb0726b86d7cf2426695b4cf70f0b25d0955feaf32d027d3f4f"></a>

## total_nodes property — aws_parameters / da1487af8f02 / 9

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

<a id="canonical-4912b3f08118b46648c3b7c746cfaf4f39e633d78350ad60aa0bf3666ab64267"></a>

<a id="canonical-f857abc48af75154c57c4cea96742326b111dc854169a4ce48c60635d034d732"></a>

## vpc_id property — aws_parameters / da1487af8f02 / 10

Type: `"string"`. Optional.

Exclusive with \[new\_vpc\] Existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Existing VPC ID.

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
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-15f8ab6d97583488daf87aec89bbfa916ffd77548bb57ddcbfd651f184f21e49"></a>

## Next pages — aws_parameters / da1487af8f02 / 11

- [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-88f8143d60eb6d965531df056eb61c3c957a4e86ddbed6d734b8eb35169a4452)
- [aws_parameters.aws_cred](resources--aws_tgw_site--reference--group-001.md#canonical-e9e60ee6cd97042783e932829df6acea0fc6b60f4182d9f8d517f9c7228511e0)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [aws_parameters.custom_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-a2dbf3068620cc2ba0398e09b0bfc71396676e57c14fdee9a4a4b6a35bfcb755)
- [aws_parameters.disable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-6f4daa9c6f5acfe04beb470e293d168eda234f99982312b18c311291fe554551)
- [aws_parameters.disable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-94343b10b6742d3b41a087823b366f63704533082ddd021cf173675cefd4954f)
- [aws_parameters.enable_encryption](resources--aws_tgw_site--reference--group-001.md#canonical-9fcf4edf24bb9ac5a87ed674acb3693271da466bc5df4c2dc505b7bf598f61c9)
- [aws_parameters.enable_internet_vip](resources--aws_tgw_site--reference--group-001.md#canonical-00ffbb634b9197ae73007e874f86eb0a47a982e3437a4789288f9a7ed47bf3f7)
- [aws_parameters.existing_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-13688fa212f3629952f7851b68ae9ce15d01ad9608fffecf507600292ce14c98)
- [aws_parameters.f5xc_security_group](resources--aws_tgw_site--reference--group-001.md#canonical-b34cebbd4a5b43e28221b06df679ba86048fdab745ea758667b7f8d8473e73e0)
- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-2112fc494e90b1dc8d39e4b1b6f2967d6fe4a8cc5da84ec25fff2ba07bb87f52)
- [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-001.md#canonical-2823b6f937a1792c17d1a89fa7317dd7838c44ea0e8ac666805abf0d98ea773a)
- [aws_parameters.no_worker_nodes](resources--aws_tgw_site--reference--group-002.md#canonical-a989844489c66d4e7faae9b8073925c332423fe87f9ed9161f123294f9f2b651)
- [aws_parameters.reserved_tgw_cidr](resources--aws_tgw_site--reference--group-002.md#canonical-d99f8f19c6a7e99ee542ccfb2979873c26ae54bd8d9b7fa065b15906cbe89931)
- [aws_parameters.tgw_cidr](resources--aws_tgw_site--reference--group-002.md#canonical-dda6f57721acd4eeedf2a9b8e0df09e5fe75a64f91d5fb5515a8f7664c2a2e98)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-88f8143d60eb6d965531df056eb61c3c957a4e86ddbed6d734b8eb35169a4452"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c30e5dba1831dbd2826a136d3c508fea1eec9f9a6cb3aadec701848a838a8d55"></a>

## aws_parameters.admin_password — aws_parameters.admin_password / 998b16a12df7 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.admin_password

<a id="canonical-47fe691e9721e0d7e2e3f4df2200aa1fed77dbf18c335b5d9aca87226ab68197"></a>

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
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-b8d70a6b5faa28bc2e45e37617ef2cb3bf66a35b9e66b566b4a4854c0ad1c7a6"></a>

## Direct properties — aws_parameters.admin_password / 998b16a12df7 / 3

- [blindfold_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-8ae03ea9d521b70b12de43648ec5a70ff58b0dfca5379f3fb0159af39d9ae54e): complete subsection reference.

- [clear_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-cf1cc6f0432406c467b75996913f7366edaa8a69808753af1edc300950bf49b4): complete subsection reference.

<a id="canonical-61e0ecd7b1ed7cf5a8b30ef5dc43fad4ce2dbc6ca094edec6185af3707aac2b6"></a>

## Next pages — aws_parameters.admin_password / 998b16a12df7 / 4

- [aws_parameters.admin_password.blindfold_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-8ae03ea9d521b70b12de43648ec5a70ff58b0dfca5379f3fb0159af39d9ae54e)
- [aws_parameters.admin_password.clear_secret_info](resources--aws_tgw_site--reference--group-001.md#canonical-cf1cc6f0432406c467b75996913f7366edaa8a69808753af1edc300950bf49b4)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-8ae03ea9d521b70b12de43648ec5a70ff58b0dfca5379f3fb0159af39d9ae54e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffea1db1e2970f7e1fb0b52d7ce4cd7ed9966a6bdcfb74f8eee65f5dfde10d7b"></a>

## aws_parameters.admin_password.blindfold_secret_info — aws_parameters.admin_password.blindfold_secret_info / a2b745dc908f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-88f8143d60eb6d965531df056eb61c3c957a4e86ddbed6d734b8eb35169a4452)
- aws_parameters.admin_password.blindfold_secret_info

<a id="canonical-4adb868bb34fb595fef71192fa480a4637398817e54deff1d7e6d9b581bc2897"></a>

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

<a id="canonical-a3aebf79f5a7ff23331e5a5d70c824b1edd8bf18cafc0edcc16622dfbfa13924"></a>

## Direct properties — aws_parameters.admin_password.blindfold_secret_info / a2b745dc908f / 3

<a id="canonical-3066e9510c16101dd4008390fe8f3db63f94a3146816cb8904ef8ec10c2f0a7f"></a>

<a id="canonical-8c4793e2c5476d6da93330730383ff17aa9d3a35418aa27c64ebb66bc073a8c8"></a>

## decryption_provider property — aws_parameters.admin_password.blindfold_secret_info / a2b745dc908f / 4

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

<a id="canonical-902388d1b73e7ed0e01782f3517ada8992d60fdb5b694e9edc8e664b5697d500"></a>

<a id="canonical-535a2de7488d06d54244e07c176f3e91c4f02e18fed77ae90d2dd75702f9113d"></a>

## location property — aws_parameters.admin_password.blindfold_secret_info / a2b745dc908f / 5

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

<a id="canonical-e0437efedf39e3a131f29a7e3ffc6323c4e0edf8c0444e3593fe7a0ae58274d4"></a>

<a id="canonical-e72c125a4f1ad3b923becc49c180afab34f5e219a6d186b5f09847b229491397"></a>

## store_provider property — aws_parameters.admin_password.blindfold_secret_info / a2b745dc908f / 6

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

<a id="canonical-93d0de8deeca6ddfe2f2d752f41f5138da16aed700dec5545aa80e2786eeabeb"></a>

## Next pages — aws_parameters.admin_password.blindfold_secret_info / a2b745dc908f / 7

- [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-88f8143d60eb6d965531df056eb61c3c957a4e86ddbed6d734b8eb35169a4452)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-cf1cc6f0432406c467b75996913f7366edaa8a69808753af1edc300950bf49b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75515449fc2d22837ed9eb569309d331b94a35fae02b30ace1432d1312c2a87c"></a>

## aws_parameters.admin_password.clear_secret_info — aws_parameters.admin_password.clear_secret_info / 33fa9641f62b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-88f8143d60eb6d965531df056eb61c3c957a4e86ddbed6d734b8eb35169a4452)
- aws_parameters.admin_password.clear_secret_info

<a id="canonical-95a9fa8f85097f67a04129fd26fcfe2eae0d7566f10f62624676956ea703e322"></a>

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

<a id="canonical-efa9bebf033f108eb755e419cfed5eec8c8f4361994ff51ca16646865540ad9b"></a>

## Direct properties — aws_parameters.admin_password.clear_secret_info / 33fa9641f62b / 3

<a id="canonical-b9828acaaa98b63162c5b92b97aed389d69bc5b8278fa2ed9dacaffc878573d1"></a>

<a id="canonical-d193d4142e00b25cd51dcb04c0a5e9c9b04a3fcf138af15f10032cc37796415c"></a>

## provider_ref property — aws_parameters.admin_password.clear_secret_info / 33fa9641f62b / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b19f97b8fbcfcc528c808b756f37285bb0dd2fce073aefecf9a19c00852ada3f"></a>

<a id="canonical-eedb8ff94aab874b84fd548da6b8854254b408ae4311ce912361198b19887cd4"></a>

## url property — aws_parameters.admin_password.clear_secret_info / 33fa9641f62b / 5

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

<a id="canonical-e262001651924ff3742f1fd99df4c1d5b89f533f85a887288dc554f737f6cad5"></a>

## Next pages — aws_parameters.admin_password.clear_secret_info / 33fa9641f62b / 6

- [aws_parameters.admin_password](resources--aws_tgw_site--reference--group-001.md#canonical-88f8143d60eb6d965531df056eb61c3c957a4e86ddbed6d734b8eb35169a4452)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-e9e60ee6cd97042783e932829df6acea0fc6b60f4182d9f8d517f9c7228511e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27a003a18d26b4020c5d3032365881a1d813dd61994f7f7c4241be1a00aaf6f2"></a>

## aws_parameters.aws_cred — aws_parameters.aws_cred / 7e3dfbb0b040 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.aws_cred

<a id="canonical-afbd740e212a335c1aa9e082e7f7f41ccbda3c87ff5716cb0bae22ce0bcb3a8d"></a>

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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-78dad7b555dba61b9b6b25a7e5cc0010a07f5144ef7afd8a7a74d8e7c4791637"></a>

## Direct properties — aws_parameters.aws_cred / 7e3dfbb0b040 / 3

<a id="canonical-aa086c855a2fb5c590686b1ba17600cd5ee27e6cf53216523c460e318ab399c6"></a>

<a id="canonical-18ddafe7a79a37272366d4f526935c62ed310e5eb68aa42e4892976b5834df50"></a>

## name property — aws_parameters.aws_cred / 7e3dfbb0b040 / 4

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

<a id="canonical-346f0bbc0bb98b81e6dd7da3fc41dcbcd3284a19b73e6a79c6656148a97e891d"></a>

<a id="canonical-d6cf8008c44b52ec1f50fe2b2dc11af4958b717d55c827dfcd3a6ea38e271705"></a>

## namespace property — aws_parameters.aws_cred / 7e3dfbb0b040 / 5

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

<a id="canonical-edfebec6a589264264f8d0777379fa7400efb8e92ac54c65bf3cbd237978cb91"></a>

<a id="canonical-db1b0ba5af52b0dc4939010f0597ea20bbc624cf882fd474f15f5afe22c6e7a0"></a>

## tenant property — aws_parameters.aws_cred / 7e3dfbb0b040 / 6

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

<a id="canonical-b75c2b655ca91063c1a6a306399627feb8afc9412f6fec2c9d83f255705b76cd"></a>

## Next pages — aws_parameters.aws_cred / 7e3dfbb0b040 / 7

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e330f2e65af5eadc273602b5ece8a07252df372265581bdb69a023333259a0b"></a>

## aws_parameters.az_nodes — aws_parameters.az_nodes / e1ec34b58bbf / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.az_nodes

<a id="canonical-fb2bfb41dc0d82bb84cf60336d281ff240717dc003051f4e096c0ac5eb296ef5"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name"),
  validators.ConflictingListObjectAttributes("inside_subnet",
    "reserved_inside_subnet")}
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

<a id="canonical-c9a16c6007f9de4c6c1cab345f9982e78bd890ec373e8ea234c57ea3c4596d8f"></a>

## Direct properties — aws_parameters.az_nodes / e1ec34b58bbf / 3

<a id="canonical-463365584c3c9af0b8b4b4039e915e9b4808c85fc0ca0541b32ab6a792e373ea"></a>

<a id="canonical-4b9c469eaf320df7a84fef2744476641f609928f767d74aa66994bae9c4d3b21"></a>

## aws_az_name property — aws_parameters.az_nodes / e1ec34b58bbf / 4

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

- [inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-5999cb63066f4166ef223df059ecacd92d68d1c9c5fb3bf1939d19d196064e6d): complete subsection reference.

- [outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-17065352d24fe947e9b2a45a7a5e303178f7297ecbde3da1d267d5884f4ab8fa): complete subsection reference.

- [reserved_inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-2536a8e26a9c3af28ee612d67c0121f6c18e077ac6264c76a92a4fe63b5a3505): complete subsection reference.

- [workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-9e9660a632e8abfc5cbd116c62206b38ce0dfaeae340c4ca3691816ad1d88903): complete subsection reference.

<a id="canonical-cac299216713b5ba478bd0b1023e5103d046fe3b96b719ecb0eb415fb2c942b7"></a>

## Next pages — aws_parameters.az_nodes / e1ec34b58bbf / 5

- [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-5999cb63066f4166ef223df059ecacd92d68d1c9c5fb3bf1939d19d196064e6d)
- [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-17065352d24fe947e9b2a45a7a5e303178f7297ecbde3da1d267d5884f4ab8fa)
- [aws_parameters.az_nodes.reserved_inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-2536a8e26a9c3af28ee612d67c0121f6c18e077ac6264c76a92a4fe63b5a3505)
- [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-9e9660a632e8abfc5cbd116c62206b38ce0dfaeae340c4ca3691816ad1d88903)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-5999cb63066f4166ef223df059ecacd92d68d1c9c5fb3bf1939d19d196064e6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-921996683690cb7f752fc41a04a8a6b25aac2c76cf07b229d6dbe2ff80acb23b"></a>

## aws_parameters.az_nodes.inside_subnet — aws_parameters.az_nodes.inside_subnet / 0da927da1eba / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- aws_parameters.az_nodes.inside_subnet

<a id="canonical-b910dc8e743677ff887c22e8152161c24f45563eec30bb4ca18fa8c69b090805"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

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
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2ea9f989404652690173a8820380d5ac5334b73327fadad7d7ae2d8181db1d2"></a>

## Direct properties — aws_parameters.az_nodes.inside_subnet / 0da927da1eba / 3

<a id="canonical-4de66bbd0c8714a337e588f55443bcaf9dce84cc9c37a7ad01357e82b5044edd"></a>

<a id="canonical-874649d7579d0be6e8c2f394aab4ff0b1e7dcf47d34b24d8c2e9ee84758957e8"></a>

## existing_subnet_id property — aws_parameters.az_nodes.inside_subnet / 0da927da1eba / 4

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

- [subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-26038bb035a9bdde78a139dc5f6b0be0799552749daff7322378f5e2da9c479d): complete subsection reference.

<a id="canonical-502c029d2f0f44e5fca16fae7bf8f453be3f664542f5df57cede2f13e50a2c85"></a>

## Next pages — aws_parameters.az_nodes.inside_subnet / 0da927da1eba / 5

- [aws_parameters.az_nodes.inside_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-26038bb035a9bdde78a139dc5f6b0be0799552749daff7322378f5e2da9c479d)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-26038bb035a9bdde78a139dc5f6b0be0799552749daff7322378f5e2da9c479d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0951e947eb14c05f10b50718dc23782c5ea032b5b728b958c1042199cb6db6d"></a>

## aws_parameters.az_nodes.inside_subnet.subnet_param — aws_parameters.az_nodes.inside_subnet.subnet_param / 3f20c45830d9 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-5999cb63066f4166ef223df059ecacd92d68d1c9c5fb3bf1939d19d196064e6d)
- aws_parameters.az_nodes.inside_subnet.subnet_param

<a id="canonical-97dd02a6aa4fec37729fea8b6d3ae177bb50268a962b48dbe60822a8d4d1d50e"></a>

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

<a id="canonical-31083ef8f9916679063d8dca4b4feb68ea18e9eb2d78746002d54cbb8bd674cd"></a>

## Direct properties — aws_parameters.az_nodes.inside_subnet.subnet_param / 3f20c45830d9 / 3

<a id="canonical-716c070e124589efc5da18a08de6c22654bcd7d69909be2115c4bfc9598e484c"></a>

<a id="canonical-203e2183ac465dbeaa634516533e81958842fdf5e8ca81701693f5b10ff4b577"></a>

## ipv4 property — aws_parameters.az_nodes.inside_subnet.subnet_param / 3f20c45830d9 / 4

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

<a id="canonical-97bbdd0c13ea150d27afd07dbd8c53866904b5768ad20c3eb9398e717488fbaa"></a>

## Next pages — aws_parameters.az_nodes.inside_subnet.subnet_param / 3f20c45830d9 / 5

- [aws_parameters.az_nodes.inside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-5999cb63066f4166ef223df059ecacd92d68d1c9c5fb3bf1939d19d196064e6d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-17065352d24fe947e9b2a45a7a5e303178f7297ecbde3da1d267d5884f4ab8fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84076648983481e7cf83e53c7c758564ec79b3cd03126a7436e646f80744f518"></a>

## aws_parameters.az_nodes.outside_subnet — aws_parameters.az_nodes.outside_subnet / 5aa1af600721 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- aws_parameters.az_nodes.outside_subnet

<a id="canonical-04960f4424d2b0d4b470c4108cfc828a45c706130ef83c88103d1c68bbcfb980"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

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
outside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1c8fc5108b2b76585b8a13feb438765411e73c381e56629e082da25a610d6b1d"></a>

## Direct properties — aws_parameters.az_nodes.outside_subnet / 5aa1af600721 / 3

<a id="canonical-b0d593e80d170c975b17abea8a26a811d14963e57c9666a85120b5b8abca6d89"></a>

<a id="canonical-1ceeaa240c00ac0cb568cb38a4c0aeca5ec5e424bc62ffbe27eb5cefe253ff71"></a>

## existing_subnet_id property — aws_parameters.az_nodes.outside_subnet / 5aa1af600721 / 4

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

- [subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-8a32a0e11d89ffe0cfc83670add60dcacefb26e51857505ec2aef766f4ca6773): complete subsection reference.

<a id="canonical-a7b2a03c0089154c3c6ab3e74bbd2309836e9f85b6e839db39a034c194bd8fc8"></a>

## Next pages — aws_parameters.az_nodes.outside_subnet / 5aa1af600721 / 5

- [aws_parameters.az_nodes.outside_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-8a32a0e11d89ffe0cfc83670add60dcacefb26e51857505ec2aef766f4ca6773)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-8a32a0e11d89ffe0cfc83670add60dcacefb26e51857505ec2aef766f4ca6773"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14ba74ee5c43740a68c9b65f69356a0f257e51737ea2781ae4c58d2583180287"></a>

## aws_parameters.az_nodes.outside_subnet.subnet_param — aws_parameters.az_nodes.outside_subnet.subnet_param / 8367c25a963f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-17065352d24fe947e9b2a45a7a5e303178f7297ecbde3da1d267d5884f4ab8fa)
- aws_parameters.az_nodes.outside_subnet.subnet_param

<a id="canonical-178d71bcf91e185d748bfc5e5131913de3c1a022c3a15f0b73b8a2560ebf4c65"></a>

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

<a id="canonical-5104f5732f44d6fb27ddf7a7bf34f8659d07e7361488e5c8cdced7f2712dfa7c"></a>

## Direct properties — aws_parameters.az_nodes.outside_subnet.subnet_param / 8367c25a963f / 3

<a id="canonical-15d5fcdf0af88ae551c6ce4295c0a276806d8e83d22dd9b0cb775452d2d10094"></a>

<a id="canonical-bc35838cab7f162b797bbb35b490c8dd50adcb1659b273bbde4a9ffa9a2fff91"></a>

## ipv4 property — aws_parameters.az_nodes.outside_subnet.subnet_param / 8367c25a963f / 4

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

<a id="canonical-0f814fbd1319d2e320a5989b16483f843e6bc782c01c378e889b4e41384c313a"></a>

## Next pages — aws_parameters.az_nodes.outside_subnet.subnet_param / 8367c25a963f / 5

- [aws_parameters.az_nodes.outside_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-17065352d24fe947e9b2a45a7a5e303178f7297ecbde3da1d267d5884f4ab8fa)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-2536a8e26a9c3af28ee612d67c0121f6c18e077ac6264c76a92a4fe63b5a3505"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4f0020aa76f4d6f1b057efea3bbe2479c20ecf5015952f690d52c9525ee91b0"></a>

## aws_parameters.az_nodes.reserved_inside_subnet — aws_parameters.az_nodes.reserved_inside_subnet / 577bd992daeb / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- aws_parameters.az_nodes.reserved_inside_subnet

<a id="canonical-a4f81741908c1e0df6e57f9979abcc23eb421dbee371e9483f326f3c1d6681ef"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved inside subnet.

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
reserved_inside_subnet = {}
```

<a id="canonical-54bf346d91c0b8cca36b92d57a44d24ca241746e13d8ddbdb412545be632c19a"></a>

## Direct properties — aws_parameters.az_nodes.reserved_inside_subnet / 577bd992daeb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f09b60fc51ccdd7c8fd76bfb4bc53583563ef2607dc170b55a1da6d3202c1d5"></a>

## Next pages — aws_parameters.az_nodes.reserved_inside_subnet / 577bd992daeb / 4

- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-9e9660a632e8abfc5cbd116c62206b38ce0dfaeae340c4ca3691816ad1d88903"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7fc10866fc17ac96eb85ccfdce1a89f4067df222dc000228af7566cd0b5ed08"></a>

## aws_parameters.az_nodes.workload_subnet — aws_parameters.az_nodes.workload_subnet / ddc3f792a98b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- aws_parameters.az_nodes.workload_subnet

<a id="canonical-e9a5aabc1dea32ce8e69ddbcede71fd7434c2766e1d86e0ec19bde91e2a92040"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for workload subnet.

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
workload_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-e1283ceb59ef1e5e25babcb9a7c2ba41799dd1c577227688cd0146eeea2b598d"></a>

## Direct properties — aws_parameters.az_nodes.workload_subnet / ddc3f792a98b / 3

<a id="canonical-47656bb0b68796fac88a63a7b8cdb441c2bf73d4c45d46e63a16fcfb8812e3f4"></a>

<a id="canonical-16b999a025449bbb9a8a0c3496daa96495faac477a0e2aee73dd8c0a02ba3626"></a>

## existing_subnet_id property — aws_parameters.az_nodes.workload_subnet / ddc3f792a98b / 4

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

- [subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-25e8ae67a29cb47b720d0e651da85a2811680609f8d6de8e49f3037e1b353520): complete subsection reference.

<a id="canonical-05db9f64968faead3aa4bd7a5372c2873bf92d3c7c918812a21a4e3ab2cf5270"></a>

## Next pages — aws_parameters.az_nodes.workload_subnet / ddc3f792a98b / 5

- [aws_parameters.az_nodes.workload_subnet.subnet_param](resources--aws_tgw_site--reference--group-001.md#canonical-25e8ae67a29cb47b720d0e651da85a2811680609f8d6de8e49f3037e1b353520)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-25e8ae67a29cb47b720d0e651da85a2811680609f8d6de8e49f3037e1b353520"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b6549bcaef64aa87d2df0046426dafb12763e40163226a92d55724e871d1a79"></a>

## aws_parameters.az_nodes.workload_subnet.subnet_param — aws_parameters.az_nodes.workload_subnet.subnet_param / c0950076f940 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.az_nodes](resources--aws_tgw_site--reference--group-001.md#canonical-b9b100e7deba782559419b98d00a975f113e2952ad307da115a4d1ccae964b7b)
- [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-9e9660a632e8abfc5cbd116c62206b38ce0dfaeae340c4ca3691816ad1d88903)
- aws_parameters.az_nodes.workload_subnet.subnet_param

<a id="canonical-0b16c5d9dbbc4cf45c24b402ac63acd86c958febbef20a9144278a36b53026e4"></a>

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

<a id="canonical-73b3266190b7d29734eb51c455f72785b0bfc4fc6a4f09172d7ce6e2f3fa0571"></a>

## Direct properties — aws_parameters.az_nodes.workload_subnet.subnet_param / c0950076f940 / 3

<a id="canonical-1933ddd419baa21c6ce98d3bcc4e6aaf8314dda0a1a6471c7c0ab42ecd8d55f2"></a>

<a id="canonical-6fddda25dedd83106eed930b8179458e9caca442259702a1cc06bb092fe7aae4"></a>

## ipv4 property — aws_parameters.az_nodes.workload_subnet.subnet_param / c0950076f940 / 4

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

<a id="canonical-aecdbfb2122a1dd8831825363df170492a0dcb70a0badc167e9ff1fae8352544"></a>

## Next pages — aws_parameters.az_nodes.workload_subnet.subnet_param / c0950076f940 / 5

- [aws_parameters.az_nodes.workload_subnet](resources--aws_tgw_site--reference--group-001.md#canonical-9e9660a632e8abfc5cbd116c62206b38ce0dfaeae340c4ca3691816ad1d88903)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-a2dbf3068620cc2ba0398e09b0bfc71396676e57c14fdee9a4a4b6a35bfcb755"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfeff6c4f6facb4da35f767ae9e479de84e89190a1b0d0eefc497d58b1b9b5c6"></a>

## aws_parameters.custom_security_group — aws_parameters.custom_security_group / d8a1e0c73b1d / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.custom_security_group

<a id="canonical-d0968f23e68e7a95578266883f00caedbe487b227ad438793aa54b48c01a608f"></a>

Type: `"object"`. single nested block, Optional.

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
custom_security_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-9f7bee099e8d2b2b8699fe6e7ee8702244d1d23ac4886c3272494bed4f522629"></a>

## Direct properties — aws_parameters.custom_security_group / d8a1e0c73b1d / 3

<a id="canonical-ceef3b608f66430b9eb21ee06ae873678a5b777ffb65bbd1012d8a338aa17878"></a>

<a id="canonical-9f729db2d33ecae006ec937cf0a304c45e5653411e49b12d04cc18ef64f3a503"></a>

## inside_security_group_id property — aws_parameters.custom_security_group / d8a1e0c73b1d / 4

Type: `"string"`. Optional.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

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
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-73c04a5f6c4034636ba865a9e297869ff571fee2c6b32c885bcc862d50766e54"></a>

<a id="canonical-c0d1699e7456ebeefe66ed7c4b11dc66f2fdf0214b167e3890e389ea899fe84f"></a>

## outside_security_group_id property — aws_parameters.custom_security_group / d8a1e0c73b1d / 5

Type: `"string"`. Optional.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

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
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-cbce3b6ea1e344305fa492ec2a71640245ddb09c1026e09e72338639c25641b3"></a>

## Next pages — aws_parameters.custom_security_group / d8a1e0c73b1d / 6

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-6f4daa9c6f5acfe04beb470e293d168eda234f99982312b18c311291fe554551"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83f5ffa802846075799935bb26e3df19d493fc43fb7741f43bb3299baa3f5205"></a>

## aws_parameters.disable_encryption — aws_parameters.disable_encryption / d466f7221f93 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.disable_encryption

<a id="canonical-03805be8ca9cfad4b1b4847a9fe4bdb0dff83133d50b93768a5fb5088e83cd43"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable encryption.

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
disable_encryption = {}
```

<a id="canonical-87b7aa23ca6b104c62aaf1f4544addc11378dc9b5a2b1a1e98320ee4b738f8e1"></a>

## Direct properties — aws_parameters.disable_encryption / d466f7221f93 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a533f190f4f14ae6fcfbd2be61eba4eb411420bcd2c0ba55afaa3785765ad7f0"></a>

## Next pages — aws_parameters.disable_encryption / d466f7221f93 / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-94343b10b6742d3b41a087823b366f63704533082ddd021cf173675cefd4954f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9093d3bc0ee242e3bf9314b655be5b020293ebf8e51e39fe07735feebddc042e"></a>

## aws_parameters.disable_internet_vip — aws_parameters.disable_internet_vip / db50ffeea0bd / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.disable_internet_vip

<a id="canonical-ae4f0f3c1ae9d386a14f7801b9f668f892477920d5ef7ebc005ea7e8bd30ee86"></a>

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
disable_internet_vip = {}
```

<a id="canonical-14dcd2b8b33f1dfadcdec1076ff7c305ce174e3b7f32b9f442bdecce6cd656dd"></a>

## Direct properties — aws_parameters.disable_internet_vip / db50ffeea0bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8909fd4d1dbe4d7e569c08ef0d53490107bf83178a51707630a2d5e8cb5531df"></a>

## Next pages — aws_parameters.disable_internet_vip / db50ffeea0bd / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-9fcf4edf24bb9ac5a87ed674acb3693271da466bc5df4c2dc505b7bf598f61c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d47f41212a6dc61983676a215826c385f3b71209fc95b57585cc37937beada4f"></a>

## aws_parameters.enable_encryption — aws_parameters.enable_encryption / 78753a627025 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.enable_encryption

<a id="canonical-0ec3db85e4f1fd47f8640657a8e2a0f2b5ce43c53ccb5b4fb7e01e0dac474b38"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("kms_key_id")}
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
enable_encryption {
  # Configure direct properties listed below.
}
```

<a id="canonical-24f1458fe7043a42d2fb51a2f1d6a0f7a68d262efb3ab881f11aa424b9abf66f"></a>

## Direct properties — aws_parameters.enable_encryption / 78753a627025 / 3

<a id="canonical-74f7898dfe6160c450abda0d9e33cf3536bbbee4cdc91f9549cf9265dd670ebc"></a>

<a id="canonical-6ede9d4366b0efcbff2bf4fcf1e90060f6432f5cca3c58289c86458e38316c7a"></a>

## kms_key_id property — aws_parameters.enable_encryption / 78753a627025 / 4

Type: `"string"`. Optional.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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

<a id="canonical-19cca620caa73e09e62b63637051d47837c273e9a3291390a0f55ee72cca2343"></a>

## Next pages — aws_parameters.enable_encryption / 78753a627025 / 5

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-00ffbb634b9197ae73007e874f86eb0a47a982e3437a4789288f9a7ed47bf3f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fc8d23496f7ac63d9a9d74ba0b4674fe5a61c85f14e25cac273d60b9b546901"></a>

## aws_parameters.enable_internet_vip — aws_parameters.enable_internet_vip / 8d3cfba8d13b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.enable_internet_vip

<a id="canonical-8bba561b6295cefd101a9efb4d3de7b9e60b889ead0b4e79725bbf15dfa0a78f"></a>

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
enable_internet_vip = {}
```

<a id="canonical-cb0bb63b5734734404ecd113fe1d95c07f5ac1193ec72cfe9923dcc80eec7635"></a>

## Direct properties — aws_parameters.enable_internet_vip / 8d3cfba8d13b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb72f2aec47b0bdbdac0da8c513752fa3b2da8f3ff5047da8438be7fbf30811c"></a>

## Next pages — aws_parameters.enable_internet_vip / 8d3cfba8d13b / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-13688fa212f3629952f7851b68ae9ce15d01ad9608fffecf507600292ce14c98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f3bef92a7862ad747d78e65defed8c2bae5cc8fecebc151b272585932c927c9"></a>

## aws_parameters.existing_tgw — aws_parameters.existing_tgw / 20abfd274e87 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.existing_tgw

<a id="canonical-cdf4b580fd1cb25dc424c53737109e1d9f0185b0cf336a041e012b0d6a1e4eb4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing tgw.

Upstream description:

Information needed for existing TGW.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
existing_tgw {
  # Configure direct properties listed below.
}
```

<a id="canonical-a8f8875a94d9674011fd1aa3055bf3ee8217fb22a6b270f261e953393d11dc8d"></a>

## Direct properties — aws_parameters.existing_tgw / 20abfd274e87 / 3

<a id="canonical-e3b1bbc23c815bd9f414fa1fa6c65d747a8d6d4d5cfeb40edd1aa76f699085fd"></a>

<a id="canonical-79bbffd6c111183673f11ed0db7d168dc9d76c26fa043d6e2eb4891baf21f7e4"></a>

## tgw_asn property — aws_parameters.existing_tgw / 20abfd274e87 / 4

Type: `"number"`. Optional.

Enter TGW ASN. TGW ASN.

Upstream description:

TGW ASN.

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

<a id="canonical-4eb547e195d687d888811ec0f055f936b62f394e98e8a085cb8f9313f1496cd3"></a>

<a id="canonical-8f1fc9cff28499387bc330a92635b5d22221a9a115d7854b492053e20b185f7e"></a>

## tgw_id property — aws_parameters.existing_tgw / 20abfd274e87 / 5

Type: `"string"`. Optional.

Existing TGW ID. Existing TGW ID.

Upstream description:

Existing TGW ID.

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
    "pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-c5c1152ed7ffc46dec87e95e587a97d5c6134c45567bd61721fcd85e45325e9e"></a>

<a id="canonical-a0e17612daf1d426c5c5b0cd70f771a10ee143a0d9522343c084b74be085f801"></a>

## volterra_site_asn property — aws_parameters.existing_tgw / 20abfd274e87 / 6

Type: `"number"`. Optional.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

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

<a id="canonical-10887fa4891f3a69389ab8c0c442f24fe63d22f68e7815d4b58b5313fe7e24fa"></a>

## Next pages — aws_parameters.existing_tgw / 20abfd274e87 / 7

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b34cebbd4a5b43e28221b06df679ba86048fdab745ea758667b7f8d8473e73e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad7d8c5e8cd13e2c8133d6c66324cb04becebc3afdad8ad5e88cdc88ed617e87"></a>

## aws_parameters.f5xc_security_group — aws_parameters.f5xc_security_group / 0c8ba203a5cf / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.f5xc_security_group

<a id="canonical-54dc7e3625aa67ce268c0817dc425f18de17a8bcf921471f3da7ba24341646b4"></a>

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
f5xc_security_group = {}
```

<a id="canonical-e74c0935a8beba1b526b16fa457ef5f941ecb1ff94d6720f68933486351126ce"></a>

## Direct properties — aws_parameters.f5xc_security_group / 0c8ba203a5cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07459583612a92cf2c5aba72786aec7fafabfb523501d07db25cdb76dd42ad4e"></a>

## Next pages — aws_parameters.f5xc_security_group / 0c8ba203a5cf / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-2112fc494e90b1dc8d39e4b1b6f2967d6fe4a8cc5da84ec25fff2ba07bb87f52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ea61fb5112ff78f627702f75e2d6afe014bf603651dcd927c15dc6524102713"></a>

## aws_parameters.new_tgw — aws_parameters.new_tgw / 65c9a3b4becf / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- aws_parameters.new_tgw

<a id="canonical-5a5f926fed2cc993355bc753558edc5ecf309df9817bc17f3df85df1086d9b5c"></a>

Type: `"object"`. single nested block, Optional.

TGWParamsType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("system_generated",
    "user_assigned")}
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
  "x-ves-oneof-field-asn_choice": "[\"system_generated\",\"user_assigned\"]"
}
```

Terraform syntax:

```terraform
new_tgw {
  # Configure direct properties listed below.
}
```

<a id="canonical-e787e23b3c0ad4d78b9539b02bea679f6a50e7aae6855a58477ef131b9b5613d"></a>

## Direct properties — aws_parameters.new_tgw / 65c9a3b4becf / 3

- [system_generated](resources--aws_tgw_site--reference--group-001.md#canonical-13a572284b3b05205ae03250f1ad44c06e6843609b0f41f811a21db37ef3a5ca): complete subsection reference.

- [user_assigned](resources--aws_tgw_site--reference--group-001.md#canonical-46e61e66ff905e68701d95cb84a8a4d87d7dd02d84f8a17529630edc5ce9f23e): complete subsection reference.

<a id="canonical-32459c6b788912273c8a3ede314a2b328909686fdea76efe468b8327d8470bdf"></a>

## Next pages — aws_parameters.new_tgw / 65c9a3b4becf / 4

- [aws_parameters.new_tgw.system_generated](resources--aws_tgw_site--reference--group-001.md#canonical-13a572284b3b05205ae03250f1ad44c06e6843609b0f41f811a21db37ef3a5ca)
- [aws_parameters.new_tgw.user_assigned](resources--aws_tgw_site--reference--group-001.md#canonical-46e61e66ff905e68701d95cb84a8a4d87d7dd02d84f8a17529630edc5ce9f23e)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-13a572284b3b05205ae03250f1ad44c06e6843609b0f41f811a21db37ef3a5ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fcf0072b636002d6bf5f1eef7542c784535c7bdcbaa6838e4ffc76189f8ab6b"></a>

## aws_parameters.new_tgw.system_generated — aws_parameters.new_tgw.system_generated / a662fdfec0cc / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-2112fc494e90b1dc8d39e4b1b6f2967d6fe4a8cc5da84ec25fff2ba07bb87f52)
- aws_parameters.new_tgw.system_generated

<a id="canonical-156529bf2718cd0dccd2afbffec20cc05742611a626fd1d9799d5386fd0f592e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system generated.

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
system_generated = {}
```

<a id="canonical-d1272ecc62acb1a0a6c1a1db44b1a40bac9b5c344654f2e6cd40191bce229789"></a>

## Direct properties — aws_parameters.new_tgw.system_generated / a662fdfec0cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ec1164d4cf518d5bf345b00f063033a2ab79d106605c994e3cd07aa52525d19"></a>

## Next pages — aws_parameters.new_tgw.system_generated / a662fdfec0cc / 4

- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-2112fc494e90b1dc8d39e4b1b6f2967d6fe4a8cc5da84ec25fff2ba07bb87f52)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-46e61e66ff905e68701d95cb84a8a4d87d7dd02d84f8a17529630edc5ce9f23e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5ec29ef8968baf77f1fd4903bfce15ef91ede04f769fbaeac147d5b231d3361"></a>

## aws_parameters.new_tgw.user_assigned — aws_parameters.new_tgw.user_assigned / 949228739ee8 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-2597973ff44d50cebddd1b59403ed1be7dd86c9583f43455ed40766869b89f48)
- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-2112fc494e90b1dc8d39e4b1b6f2967d6fe4a8cc5da84ec25fff2ba07bb87f52)
- aws_parameters.new_tgw.user_assigned

<a id="canonical-3c055260d3e6ece4fd16bec67061ce12670b6722073190ae9818f7371d6994ca"></a>

Type: `"object"`. single nested block, Optional.

Information needed when ASNs are assigned by the user.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
user_assigned {
  # Configure direct properties listed below.
}
```

<a id="canonical-f22e5363624ec0a0cd5b3a9af747d9a33f2022022d4783fdf74b1df2d0aa237c"></a>

## Direct properties — aws_parameters.new_tgw.user_assigned / 949228739ee8 / 3

<a id="canonical-45f4ae92e7113fa99cb4ab87ca978f3f5f160c47eeb246e2959fb3d8253dd5f0"></a>

<a id="canonical-b633437c333fc783ce69adca0a4161f107f6636c6777248d004e6f929af16414"></a>

## tgw_asn property — aws_parameters.new_tgw.user_assigned / 949228739ee8 / 4

Type: `"number"`. Optional.

TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(64513, 65534),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65534,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 64513
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  }
}
```

<a id="canonical-6e4859bffd629ffb333ef65785181d82f301cae7c9337b87f886d3ccb3262f96"></a>

<a id="canonical-e38e0d34acdafc5e888cf917a7209fe4917b84679b6384c76b8f16d8815a9451"></a>

## volterra_site_asn property — aws_parameters.new_tgw.user_assigned / 949228739ee8 / 5

Type: `"number"`. Optional.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

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

<a id="canonical-2c243f2a86c03b777935da96e5757c90c208b40ae659c4c5531522a6c2f95511"></a>

## Next pages — aws_parameters.new_tgw.user_assigned / 949228739ee8 / 6

- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-2112fc494e90b1dc8d39e4b1b6f2967d6fe4a8cc5da84ec25fff2ba07bb87f52)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-2823b6f937a1792c17d1a89fa7317dd7838c44ea0e8ac666805abf0d98ea773a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
