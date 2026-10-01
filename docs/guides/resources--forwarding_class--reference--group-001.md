---
page_title: "xcsh_forwarding_class reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class reference."
---

# xcsh_forwarding_class reference

<a id="canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd5651f66ad7de34039fde92d855957d770365b2941f811762feb77c78daf051"></a>

## Property reference — Property reference / 3c536e356c29 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- Property reference

<a id="canonical-24f18b1570ed5c801d05bf033a46738a74f1fc99bbad40ce3fcdf6ad57170cc2"></a>

## Direct properties — Property reference / 3c536e356c29 / 3

<a id="canonical-7d0a1c992fb22a424c1681167d77f6c7dceabc5206af55795049b6b493a04558"></a>

<a id="canonical-b29ed6551fc255195707ebc51ae253ebbd1a511d5cf3739a8589f63f00184f33"></a>

## annotations property — Property reference / 3c536e356c29 / 4

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

<a id="canonical-fa45ea64f9c5f1182dfd77e6379102fc1578800e19c908ca942985f6f1146ada"></a>

<a id="canonical-9eb24d975256375c71eab3e5daed22c3107207174abeb8ed77313029d8112074"></a>

## description property — Property reference / 3c536e356c29 / 5

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

<a id="canonical-986618c868df0ea30853a0b4e7109c4f5b73a9d899b870c4973fef1e7e549450"></a>

<a id="canonical-8a9c3b16a581a824f105c192cd154aba2ba5b1105586044e94952f09449126e8"></a>

## disable property — Property reference / 3c536e356c29 / 6

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

- [dscp](resources--forwarding_class--reference--group-001.md#canonical-652a5476f9a2f31d345fd194c20a332b21b20119c14550b0ff9f70dacf68ea81): complete subsection reference.

- [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-f25fbab11f19cb9529d9189beb151a1b3309f2a20a11a4e92e878459c1742f38): complete subsection reference.

<a id="canonical-f8249cbe7fbcc1e6b6b452df1379c4a7d2beaa3df15fe97317c6918713713486"></a>

<a id="canonical-0870a231f076940f4bf3d9aa3197a376f95c9836cf206643d761e242d43f031a"></a>

## id property — Property reference / 3c536e356c29 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-985568ac06d6464eadc4428fce8e829f50484aacca6f901958ad7e6ba6082b7e"></a>

<a id="canonical-1fff4451354523424f9de2fc3e06cfa0325b7c0e28af1f108020099c5b1af6e4"></a>

## interface_group property — Property reference / 3c536e356c29 / 8

Type: `"string"`. Optional, Computed.

\[Enum: ANY\_AVAILABLE\_INTERFACE|INTERFACE\_GROUP1|INTERFACE\_GROUP2|INTERFACE\_GROUP3\] Interface
group, group membership by adding group label to interface Choose any of the available interfaces
Choose all interfaces with label group1 Choose all interfaces with label group2 Choose all
interfaces with label group3. Possible values are \`ANY\_AVAILABLE\_INTERFACE\`,
\`INTERFACE\_GROUP1\`, \`INTERFACE\_GROUP2\`, \`INTERFACE\_GROUP3\`. Defaults to
\`ANY\_AVAILABLE\_INTERFACE\`.

Upstream description:

Interface group, group membership by adding group label to interface

Choose any of the available interfaces Choose all interfaces with label group1 Choose all interfaces
with label group2 Choose all interfaces with label group3.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY_AVAILABLE_INTERFACE",
    "INTERFACE_GROUP1",
    "INTERFACE_GROUP2",
    "INTERFACE_GROUP3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY_AVAILABLE_INTERFACE",
  "enum": [
    "ANY_AVAILABLE_INTERFACE",
    "INTERFACE_GROUP1",
    "INTERFACE_GROUP2",
    "INTERFACE_GROUP3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-52087ae9e3bb5f06b197d7427ba983e1cdc621522d7f694735baf846395c121a"></a>

<a id="canonical-e669b352c0afa5cb45a0608baa34052f0d4b317f3fd79a3e6d96cef487fe96a0"></a>

## labels property — Property reference / 3c536e356c29 / 9

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

<a id="canonical-7a791aa96db4e0443a237fceed262c8e9901d027c7cabb85a8454ec4c5ddd431"></a>

<a id="canonical-9e37d9f21cd1de5ae2aa2d797d10aadf5a3ee7d44a495d18fe667118e145ff88"></a>

## name property — Property reference / 3c536e356c29 / 10

Type: `"string"`. Required.

Name of the Forwarding Class. Must be unique within the namespace.

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

<a id="canonical-e67854746d1fcdd86f9fe94d4f3bc1d216605a27b4fe1997611e2fcd7755e599"></a>

<a id="canonical-1c0117c0aea5136377a419274633070ea3cd1c3702e41d174a810c6c1865b19a"></a>

## namespace property — Property reference / 3c536e356c29 / 11

Type: `"string"`. Required.

Namespace where the Forwarding Class is created.

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

- [no_marking](resources--forwarding_class--reference--group-001.md#canonical-494dca555b3fadc38e06383c01fc218bfb27e9a5d7e3f06c515df83e74869dcf): complete subsection reference.

- [no_policer](resources--forwarding_class--reference--group-001.md#canonical-529620cf730077a8211f8c7b8a7a617ed5765227a32260eef5101e6e9ec5cede): complete subsection reference.

- [policer](resources--forwarding_class--reference--group-001.md#canonical-5499ccb2d8d3bd8429ec86b17e4ddd45ae6f8b6df1d09c3892e1d912d6e7d441): complete subsection reference.

<a id="canonical-7a55a353d3eef8c3d6b3740d103e9214a32006203e8399d6fb289e04c028b9ab"></a>

<a id="canonical-cc71f4a51b7d22e92dde8059ad9a7bc4568faa50ea1bb4267c82db9ebbc3501a"></a>

## queue_id_to_use property — Property reference / 3c536e356c29 / 12

Type: `"string"`. Optional, Computed.

\[Enum:
DSCP\_BEST\_EFFORT|DSCP\_CLASS1|DSCP\_CLASS2|DSCP\_CLASS3|DSCP\_CLASS4|DSCP\_EXPRESS\_FORWARDING|DSCP\_CONTROL\_L3|DSCP\_CONTROL\_L2\]
DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1
service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used
for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic
like.. Possible values are \`DSCP\_BEST\_EFFORT\`, \`DSCP\_CLASS1\`, \`DSCP\_CLASS2\`,
\`DSCP\_CLASS3\`, \`DSCP\_CLASS4\`, \`DSCP\_EXPRESS\_FORWARDING\`, \`DSCP\_CONTROL\_L3\`,
\`DSCP\_CONTROL\_L2\`. Defaults to \`DSCP\_BEST\_EFFORT\`.

Upstream description:

DSCP Precedence Level Values

Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP
Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is
used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not
recommended.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DSCP_BEST_EFFORT",
  "enum": [
    "DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--forwarding_class--reference--group-001.md#canonical-6e760ce30d8e419f86dc9614acc3c3489bdde9e9d18a34a3b1b3eef8ef984079): complete subsection reference.

<a id="canonical-58f4331eaf851461db9b763c8f56420c1a5cb2b1a75d7ef0c2c34a7abc6ac7ab"></a>

<a id="canonical-20b66e7096439935d49df309e3be610950341d25d33de8a46d05ebbeed5ed052"></a>

## tos_value property — Property reference / 3c536e356c29 / 13

Type: `"number"`. Optional, Computed.

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Upstream description:

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-240da4b042a1d28c069ed9a7efd711156c218e83266f67b0cd9cee1d63a4fb08"></a>

## All schema paths — Property reference / 3c536e356c29 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--forwarding_class--reference--group-001.md#canonical-7d0a1c992fb22a424c1681167d77f6c7dceabc5206af55795049b6b493a04558) |
| `description` | [description](resources--forwarding_class--reference--group-001.md#canonical-fa45ea64f9c5f1182dfd77e6379102fc1578800e19c908ca942985f6f1146ada) |
| `disable` | [disable](resources--forwarding_class--reference--group-001.md#canonical-986618c868df0ea30853a0b4e7109c4f5b73a9d899b870c4973fef1e7e549450) |
| `dscp` | [dscp](resources--forwarding_class--reference--group-001.md#canonical-d81552137e7e94ad5aee6c0c3edf49025ac26b528e80fc7c85798d5b989a8769) |
| `dscp.drop_precedence` | [dscp.drop_precedence](resources--forwarding_class--reference--group-001.md#canonical-61946da10543936e1ba7d78a9ea9c6bb39fa7261c6f0adf9385d6908c7dce387) |
| `dscp.dscp_class` | [dscp.dscp_class](resources--forwarding_class--reference--group-001.md#canonical-d89277b8a884586fd682b413a3bcf65c143b22ba60938e5f29b1ce41a75fd495) |
| `dscp_based_queue` | [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-83cce2716e945e99e781a002f0ae0943a30e1af64e11cb04b91dd12bb2a28aa6) |
| `id` | [id](resources--forwarding_class--reference--group-001.md#canonical-f8249cbe7fbcc1e6b6b452df1379c4a7d2beaa3df15fe97317c6918713713486) |
| `interface_group` | [interface_group](resources--forwarding_class--reference--group-001.md#canonical-985568ac06d6464eadc4428fce8e829f50484aacca6f901958ad7e6ba6082b7e) |
| `labels` | [labels](resources--forwarding_class--reference--group-001.md#canonical-52087ae9e3bb5f06b197d7427ba983e1cdc621522d7f694735baf846395c121a) |
| `name` | [name](resources--forwarding_class--reference--group-001.md#canonical-7a791aa96db4e0443a237fceed262c8e9901d027c7cabb85a8454ec4c5ddd431) |
| `namespace` | [namespace](resources--forwarding_class--reference--group-001.md#canonical-e67854746d1fcdd86f9fe94d4f3bc1d216605a27b4fe1997611e2fcd7755e599) |
| `no_marking` | [no_marking](resources--forwarding_class--reference--group-001.md#canonical-4926e89c3dcde7d1238158b39fc352bed7e39fa14768cd33080565f2c5724f34) |
| `no_policer` | [no_policer](resources--forwarding_class--reference--group-001.md#canonical-ca1b608a5eb7a07f12344dd08788984038cfb3ccb174bb5946d5e6cb4397dc6d) |
| `policer` | [policer](resources--forwarding_class--reference--group-001.md#canonical-bbf421bc83051a6f00aa6785cd162dbac807a30ae5198ae71d562f56525fa9f1) |
| `policer.name` | [policer.name](resources--forwarding_class--reference--group-001.md#canonical-75444f001165a602a0558bf3cb1a46ecf9b16768a71971374c0e22679849379f) |
| `policer.namespace` | [policer.namespace](resources--forwarding_class--reference--group-001.md#canonical-838804402c24d6a84ed448dd7f6539206ead1cc80089e98b6d35054618004e0e) |
| `policer.tenant` | [policer.tenant](resources--forwarding_class--reference--group-001.md#canonical-8cb01d0010f0938f5c2f1ae8b1833301ce42595a3b8d4666cd9fc868de760a03) |
| `queue_id_to_use` | [queue_id_to_use](resources--forwarding_class--reference--group-001.md#canonical-7a55a353d3eef8c3d6b3740d103e9214a32006203e8399d6fb289e04c028b9ab) |
| `timeouts` | [timeouts](resources--forwarding_class--reference--group-001.md#canonical-cfdbe5e2607ee83ba24e853c640948f292b2893cb5f60d32edbe0714b7e93c81) |
| `timeouts.create` | [timeouts.create](resources--forwarding_class--reference--group-001.md#canonical-61005361e3a251e857496c37289ceece5a5f80dec5b6bc4faa3c83cb5178789c) |
| `timeouts.delete` | [timeouts.delete](resources--forwarding_class--reference--group-001.md#canonical-c6fbb85c819f438917f9024af4dba55a38153dfb2e10f922b9787ae479df2fd5) |
| `timeouts.read` | [timeouts.read](resources--forwarding_class--reference--group-001.md#canonical-9beb1695968e8ca62da1f8541b30bcc082b6c1cb913342440151a79cdc487683) |
| `timeouts.update` | [timeouts.update](resources--forwarding_class--reference--group-001.md#canonical-6c147653219b4897d53fe4df5ac6a62bc7b69b12841ce0a43005055789fc6b70) |
| `tos_value` | [tos_value](resources--forwarding_class--reference--group-001.md#canonical-58f4331eaf851461db9b763c8f56420c1a5cb2b1a75d7ef0c2c34a7abc6ac7ab) |

<a id="canonical-a8e72209294f743f0651d0f666de240a36ae9a30a044b5940832c38e86e55807"></a>

## Next pages — Property reference / 3c536e356c29 / 15

- [dscp](resources--forwarding_class--reference--group-001.md#canonical-652a5476f9a2f31d345fd194c20a332b21b20119c14550b0ff9f70dacf68ea81)
- [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-f25fbab11f19cb9529d9189beb151a1b3309f2a20a11a4e92e878459c1742f38)
- [no_marking](resources--forwarding_class--reference--group-001.md#canonical-494dca555b3fadc38e06383c01fc218bfb27e9a5d7e3f06c515df83e74869dcf)
- [no_policer](resources--forwarding_class--reference--group-001.md#canonical-529620cf730077a8211f8c7b8a7a617ed5765227a32260eef5101e6e9ec5cede)
- [policer](resources--forwarding_class--reference--group-001.md#canonical-5499ccb2d8d3bd8429ec86b17e4ddd45ae6f8b6df1d09c3892e1d912d6e7d441)
- [timeouts](resources--forwarding_class--reference--group-001.md#canonical-6e760ce30d8e419f86dc9614acc3c3489bdde9e9d18a34a3b1b3eef8ef984079)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)

<a id="canonical-652a5476f9a2f31d345fd194c20a332b21b20119c14550b0ff9f70dacf68ea81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2afd09a5738190dcfa74788d72bf3a1bc5ee283ad4833ca0fe7ecd085234585e"></a>

## dscp — dscp / 7f6aa8e06e7c / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- dscp

<a id="canonical-d81552137e7e94ad5aee6c0c3edf49025ac26b528e80fc7c85798d5b989a8769"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dscp, no\_marking, tos\_value; Default: no\_marking\] DSCP Marking setting. DSCP marking
setting as per RFC 2475.

Upstream description:

DSCP marking setting as per RFC 2475.

Receipt-pinned upstream constraints:

```json
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

- [dscp](resources--forwarding_class--reference--group-001.md#canonical-d81552137e7e94ad5aee6c0c3edf49025ac26b528e80fc7c85798d5b989a8769)
- [no_marking](resources--forwarding_class--reference--group-001.md#canonical-4926e89c3dcde7d1238158b39fc352bed7e39fa14768cd33080565f2c5724f34)
- [tos_value](resources--forwarding_class--reference--group-001.md#canonical-58f4331eaf851461db9b763c8f56420c1a5cb2b1a75d7ef0c2c34a7abc6ac7ab)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dscp {
  # Configure direct properties listed below.
}
```

<a id="canonical-b65a04e9b8e620960bf9f3a7e0635dadd870c16883605c8cbf70de17e04808f5"></a>

## Direct properties — dscp / 7f6aa8e06e7c / 3

<a id="canonical-61946da10543936e1ba7d78a9ea9c6bb39fa7261c6f0adf9385d6908c7dce387"></a>

<a id="canonical-adfc1c9dddb569feb9af296462235616ff3965c109b8f7b47b496ff0d83a67f6"></a>

## drop_precedence property — dscp / 7f6aa8e06e7c / 4

Type: `"string"`. Optional.

\[Enum: DSCP\_AF\_LOW|DSCP\_AF\_MEDIUM|DSCP\_AF\_HIGH|DSCP\_AF\_POLICER\] DSCP Assured forwarding
drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop
precedence value is taken from output of policer. Possible values are \`DSCP\_AF\_LOW\`,
\`DSCP\_AF\_MEDIUM\`, \`DSCP\_AF\_HIGH\`, \`DSCP\_AF\_POLICER\`.

Upstream description:

DSCP Assured forwarding drop precedence

DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop precedence
value is taken from output of policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_AF_LOW",
    "DSCP_AF_MEDIUM",
    "DSCP_AF_HIGH",
    "DSCP_AF_POLICER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "DSCP_AF_LOW",
    "DSCP_AF_MEDIUM",
    "DSCP_AF_HIGH",
    "DSCP_AF_POLICER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d89277b8a884586fd682b413a3bcf65c143b22ba60938e5f29b1ce41a75fd495"></a>

<a id="canonical-ec975ffd2e936a1e61dc933e52ad0ef3e9b630ea34f46edab11ee7d0d88377c6"></a>

## dscp_class property — dscp / 7f6aa8e06e7c / 5

Type: `"string"`. Optional.

\[Enum:
DSCP\_BEST\_EFFORT|DSCP\_CLASS1|DSCP\_CLASS2|DSCP\_CLASS3|DSCP\_CLASS4|DSCP\_EXPRESS\_FORWARDING|DSCP\_CONTROL\_L3|DSCP\_CONTROL\_L2\]
DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1
service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used
for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic
like.. Possible values are \`DSCP\_BEST\_EFFORT\`, \`DSCP\_CLASS1\`, \`DSCP\_CLASS2\`,
\`DSCP\_CLASS3\`, \`DSCP\_CLASS4\`, \`DSCP\_EXPRESS\_FORWARDING\`, \`DSCP\_CONTROL\_L3\`,
\`DSCP\_CONTROL\_L2\`. Defaults to \`DSCP\_BEST\_EFFORT\`.

Upstream description:

DSCP Precedence Level Values

Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP
Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is
used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not
recommended.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DSCP_BEST_EFFORT",
  "enum": [
    "DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cbd9dd46c4abfeb7cee4fec810ebb808f1cb0ae4e9d74fcada6b9345646f9374"></a>

## Next pages — dscp / 7f6aa8e06e7c / 6

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)

<a id="canonical-f25fbab11f19cb9529d9189beb151a1b3309f2a20a11a4e92e878459c1742f38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-effd3e1c32d081b714a376a7d5962098b7384a20f63c45c394c25f51dace56f2"></a>

## dscp_based_queue — dscp_based_queue / ffca7f742aec / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- dscp_based_queue

<a id="canonical-83cce2716e945e99e781a002f0ae0943a30e1af64e11cb04b91dd12bb2a28aa6"></a>

Type: `["object", {}]`. Optional.

\[OneOf: dscp\_based\_queue, queue\_id\_to\_use\] Configuration parameter for dscp based queue.

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

- [dscp_based_queue](resources--forwarding_class--reference--group-001.md#canonical-83cce2716e945e99e781a002f0ae0943a30e1af64e11cb04b91dd12bb2a28aa6)
- [queue_id_to_use](resources--forwarding_class--reference--group-001.md#canonical-7a55a353d3eef8c3d6b3740d103e9214a32006203e8399d6fb289e04c028b9ab)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dscp_based_queue = {}
```

<a id="canonical-dc218c92d70d5e3fd9a2734c6163da89473ccbc6a5df250e03ad267290e4e5a6"></a>

## Direct properties — dscp_based_queue / ffca7f742aec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e66a5b047e6b8803110c9c21b3e2008d653ed8b9b04c2bf52d3d950c4b463c21"></a>

## Next pages — dscp_based_queue / ffca7f742aec / 4

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)

<a id="canonical-494dca555b3fadc38e06383c01fc218bfb27e9a5d7e3f06c515df83e74869dcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb73b6145924a32bc707a8128897b2620b17f53a932edcd0f8477bfd1a8ef52b"></a>

## no_marking — no_marking / 2f551a220682 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- no_marking

<a id="canonical-4926e89c3dcde7d1238158b39fc352bed7e39fa14768cd33080565f2c5724f34"></a>

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
no_marking = {}
```

<a id="canonical-80e7048943e6c155a07a2341f3901722ba62ee173c09e2777aa9b8d70bb01cb6"></a>

## Direct properties — no_marking / 2f551a220682 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7f25e277221e5617983564180925a88165848646ef12786daa0981e3386b015"></a>

## Next pages — no_marking / 2f551a220682 / 4

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)

<a id="canonical-529620cf730077a8211f8c7b8a7a617ed5765227a32260eef5101e6e9ec5cede"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9368714d50084f968f76065288d794602ae25c496be120d2c2c5af23c07313a"></a>

## no_policer — no_policer / 1712da3cf046 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- no_policer

<a id="canonical-ca1b608a5eb7a07f12344dd08788984038cfb3ccb174bb5946d5e6cb4397dc6d"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_policer, policer; Default: no\_policer\] Enable this option

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

- [no_policer](resources--forwarding_class--reference--group-001.md#canonical-ca1b608a5eb7a07f12344dd08788984038cfb3ccb174bb5946d5e6cb4397dc6d)
- [policer](resources--forwarding_class--reference--group-001.md#canonical-bbf421bc83051a6f00aa6785cd162dbac807a30ae5198ae71d562f56525fa9f1)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_policer = {}
```

<a id="canonical-94094966be859328fc7717d150d8ef786127a649764cc450673e7f53394f78aa"></a>

## Direct properties — no_policer / 1712da3cf046 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-81542dd796fd4cc367b8f51bb94aff85b9c2b4bf2bdac35b5193fe236f5f3b72"></a>

## Next pages — no_policer / 1712da3cf046 / 4

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)

<a id="canonical-5499ccb2d8d3bd8429ec86b17e4ddd45ae6f8b6df1d09c3892e1d912d6e7d441"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56f121c0363c2aa87d9c9a17821087c5c12120a121ce5e870455b1dca25d5b87"></a>

## policer — policer / 534ff26ea21d / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- policer

<a id="canonical-bbf421bc83051a6f00aa6785cd162dbac807a30ae5198ae71d562f56525fa9f1"></a>

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
policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-d20c175241a71cc935c349260fe63dc2e5346f6c2e4ec3914d760fd0920b9d93"></a>

## Direct properties — policer / 534ff26ea21d / 3

<a id="canonical-75444f001165a602a0558bf3cb1a46ecf9b16768a71971374c0e22679849379f"></a>

<a id="canonical-451cf71bf5fef5878bd093c85098d8ee1bd4e3f19d1498b9c01fdce7462492df"></a>

## name property — policer / 534ff26ea21d / 4

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

<a id="canonical-838804402c24d6a84ed448dd7f6539206ead1cc80089e98b6d35054618004e0e"></a>

<a id="canonical-febe76cc254034a61c1315f4c3e292eaa6183254b3dfb2bed2909f13ab485d19"></a>

## namespace property — policer / 534ff26ea21d / 5

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

<a id="canonical-8cb01d0010f0938f5c2f1ae8b1833301ce42595a3b8d4666cd9fc868de760a03"></a>

<a id="canonical-9ce2a5e30252082eac76052bb24c3a27fbde14c1af1af08172bff25f1f1d7f94"></a>

## tenant property — policer / 534ff26ea21d / 6

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

<a id="canonical-7c58a80788f12518237293c1933b17115f549ca40ae5f9238e2f860a2592d05a"></a>

## Next pages — policer / 534ff26ea21d / 7

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)

<a id="canonical-6e760ce30d8e419f86dc9614acc3c3489bdde9e9d18a34a3b1b3eef8ef984079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-532b211f7bf5bd4c149194f34da4fd99059e3783fd329c233010510fc2c42775"></a>

## timeouts — timeouts / e2e842fdbd0f / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- timeouts

<a id="canonical-cfdbe5e2607ee83ba24e853c640948f292b2893cb5f60d32edbe0714b7e93c81"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-329fb609267a3b6fe730ecfdb4afdef56b0a85d328bdd43209aa0a4bebc87ee6"></a>

## Direct properties — timeouts / e2e842fdbd0f / 3

<a id="canonical-61005361e3a251e857496c37289ceece5a5f80dec5b6bc4faa3c83cb5178789c"></a>

<a id="canonical-4e3c7cd16fe6f3ecffcc2891695fc11c99267a5d59d15023d5ffe221299d18ca"></a>

## create property — timeouts / e2e842fdbd0f / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c6fbb85c819f438917f9024af4dba55a38153dfb2e10f922b9787ae479df2fd5"></a>

<a id="canonical-c5d491c2ae265bcc61e427bdbfadfadc4e922d7842ea9a8064479d522fe2216b"></a>

## delete property — timeouts / e2e842fdbd0f / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-9beb1695968e8ca62da1f8541b30bcc082b6c1cb913342440151a79cdc487683"></a>

<a id="canonical-184d6914d7fe1d7fb67638e04b3795f41540526d295341f74c7a99140fd9ade2"></a>

## read property — timeouts / e2e842fdbd0f / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-6c147653219b4897d53fe4df5ac6a62bc7b69b12841ce0a43005055789fc6b70"></a>

<a id="canonical-a6928467c77e7c41fa935db659a1a6391b431f95fbc17e2a32dd93594f7ed043"></a>

## update property — timeouts / e2e842fdbd0f / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9dc784bad75b4bbd4b59220856f26796968fab43bcff043a56d940aad60b8eab"></a>

## Next pages — timeouts / e2e842fdbd0f / 8

- [Property reference](resources--forwarding_class--reference--group-001.md#canonical-1255622e5d747b356be1777947f299b3ca6f69fbfbdab173a2cc2b147d98a70c)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
