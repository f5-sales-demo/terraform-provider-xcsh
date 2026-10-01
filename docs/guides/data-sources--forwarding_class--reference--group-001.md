---
page_title: "xcsh_forwarding_class reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class reference."
---

# xcsh_forwarding_class reference

<a id="canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83c87d13f527b75dc6d90005053a52c1c08fd31572ffd4ad37eebd2c247169cb"></a>

## Property reference — Property reference / 48b82565670f / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- Property reference

<a id="canonical-a14b645242702a5da3431febf51be964c8660727bdf54de3b023683505409703"></a>

## Direct properties — Property reference / 48b82565670f / 3

<a id="canonical-2a7968735e6ac8d0e1ef5e1b08ede152b90ae77b7264697b20e13bf44875775c"></a>

<a id="canonical-9e76c550a8e3cbf4b88ffe0109c07fdc1e87f59d045ae615a4f1a62d3a1df9bf"></a>

## annotations property — Property reference / 48b82565670f / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="canonical-4f69f0db39b165d39606b9706098e8ee37b31e30c151dba775cfa59625ccda6f"></a>

<a id="canonical-3489ce2f23960915a3c20fece3fae25a9a0746d56b8caa836bb2759ee83a9941"></a>

## description property — Property reference / 48b82565670f / 5

Type: `"string"`. Computed.

Description of the ForwardingClass.

Upstream description:

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

- [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-b5887e4f1044fed8d5f66820dfc5e38fdae20fd7a17d1251b0b18f5e5569b38a): complete subsection reference.

- [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-c0f512f23169993bc10a23d98ba0be21042cc383cc3c43e47b20f7e6f75f9771): complete subsection reference.

<a id="canonical-264d651a03ac3e8b3a4fa0d1cb74908ac94ac532cca3c157295ded1d4db736d8"></a>

<a id="canonical-343d9098993a3244463491c78240d2f7ab857c0674e3235f964d53896b957472"></a>

## id property — Property reference / 48b82565670f / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-cfe046601a9cc272a4010b6fd3366758829145734f573ed9e51decc43e94dede"></a>

<a id="canonical-767ba7ddf99f5bdf4c83c693ecf76e3e9e935d5aa75f6cff1cef74cc72da039d"></a>

## interface_group property — Property reference / 48b82565670f / 7

Type: `"string"`. Computed.

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

<a id="canonical-874785a0feae312be29daaac874eefed315ffb76a093e6b3c67acc8c6f339481"></a>

<a id="canonical-9370b26cdfdcdb1b157a3665064ed42944e87e7775f54da711fc23c7206cf1ee"></a>

## labels property — Property reference / 48b82565670f / 8

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-03cab2abf44a39c65f450e42eba8ef60e69975330c55bd0e130fa68299035aef"></a>

<a id="canonical-a25347e37b344c3ce2b7c8808095ae90e5cba3b14b41f36f2f1dfbde02a3b8e1"></a>

## name property — Property reference / 48b82565670f / 9

Type: `"string"`. Required.

Name of the ForwardingClass.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

<a id="canonical-a31853a9c74bc4083e6ebdab7a8fdfef2d2595e553731adabe4a31428d62d5ee"></a>

<a id="canonical-baa53a246dbaf2e4deccc7723afa2ce660b12f989af41a6e738bf69be4c340af"></a>

## namespace property — Property reference / 48b82565670f / 10

Type: `"string"`. Required.

Namespace where the ForwardingClass exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-8488a3293531055deb7ce7e847c402ac4bd1619d95cef2b7486f363204ab99be): complete subsection reference.

- [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-1e086fadeca401981e969f1d16cd6d259ea979cd2b24d89ee096642e238a0963): complete subsection reference.

- [policer](data-sources--forwarding_class--reference--group-001.md#canonical-712b7d1959dea32cb31c6a555103e15b43f4ff016c8c62d8354b66bca8364cc7): complete subsection reference.

<a id="canonical-7a627cf7af1980bc3c0067c1d1e6f07fe8c79796e36cfd65d8495261c5bac8aa"></a>

<a id="canonical-95cc6c131a7cb067b269757591f71686628e6f7640c5ca72fb22806f9af53f5c"></a>

## queue_id_to_use property — Property reference / 48b82565670f / 11

Type: `"string"`. Computed.

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

<a id="canonical-610e4ef53c88f156523970e90d14034b72a38ac6cbc05d75f9e813b05e6b1c26"></a>

<a id="canonical-43b237bbaaa28ab167a0ea311e11415e142fa6ea1c3ae0274c61cc7e27bf1a19"></a>

## tos_value property — Property reference / 48b82565670f / 12

Type: `"number"`. Computed.

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Upstream description:

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

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

<a id="canonical-eb1d0bfd8fd4fe784a830d8e8f1311f5ccf9e1d81d7b67296193d8a45bd687f1"></a>

## All schema paths — Property reference / 48b82565670f / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--forwarding_class--reference--group-001.md#canonical-2a7968735e6ac8d0e1ef5e1b08ede152b90ae77b7264697b20e13bf44875775c) |
| `description` | [description](data-sources--forwarding_class--reference--group-001.md#canonical-4f69f0db39b165d39606b9706098e8ee37b31e30c151dba775cfa59625ccda6f) |
| `dscp` | [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-1865fb53596c6eec8139d1979e8f5dfb0db7a820e2a9ca4c58d0fea8a0feab43) |
| `dscp.drop_precedence` | [dscp.drop_precedence](data-sources--forwarding_class--reference--group-001.md#canonical-bddedc97b67c71437e933d086729c141c65c5e9db4f26c44767b2d4befa2dc32) |
| `dscp.dscp_class` | [dscp.dscp_class](data-sources--forwarding_class--reference--group-001.md#canonical-4e81b717af2250c897965f8fbf276127570326988a01878ef6c6d93cbe35cf6d) |
| `dscp_based_queue` | [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-1278fdf2ce09c78740261b859af34ca815c8f5a88d27be312a752b3001136583) |
| `id` | [id](data-sources--forwarding_class--reference--group-001.md#canonical-264d651a03ac3e8b3a4fa0d1cb74908ac94ac532cca3c157295ded1d4db736d8) |
| `interface_group` | [interface_group](data-sources--forwarding_class--reference--group-001.md#canonical-cfe046601a9cc272a4010b6fd3366758829145734f573ed9e51decc43e94dede) |
| `labels` | [labels](data-sources--forwarding_class--reference--group-001.md#canonical-874785a0feae312be29daaac874eefed315ffb76a093e6b3c67acc8c6f339481) |
| `name` | [name](data-sources--forwarding_class--reference--group-001.md#canonical-03cab2abf44a39c65f450e42eba8ef60e69975330c55bd0e130fa68299035aef) |
| `namespace` | [namespace](data-sources--forwarding_class--reference--group-001.md#canonical-a31853a9c74bc4083e6ebdab7a8fdfef2d2595e553731adabe4a31428d62d5ee) |
| `no_marking` | [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-34d62f4cb92016e3412dd1383e1b8f7c52fe2bfdc84d38ac19a38b13ec3ec78f) |
| `no_policer` | [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-05453b340433107ccae17a1b474a8db81cb507fad22c70e5bc5d00d0649dd8b1) |
| `policer` | [policer](data-sources--forwarding_class--reference--group-001.md#canonical-9467423299469f0cdef085e3d7017c929c3516a6c78503888c12a7ad4430378b) |
| `policer.name` | [policer.name](data-sources--forwarding_class--reference--group-001.md#canonical-ac417417edd7530661d45b792365ede3246dbbd6d813cc361e06d4e3dce669e9) |
| `policer.namespace` | [policer.namespace](data-sources--forwarding_class--reference--group-001.md#canonical-9015a8f0de81b1fa5a7828da4e1c39a4e7872683fc098f17d565de0ba0ef76a3) |
| `policer.tenant` | [policer.tenant](data-sources--forwarding_class--reference--group-001.md#canonical-7996a1ed122235183398dbe62ddc5a2638434217f33a684e7c0030201d0edaa8) |
| `queue_id_to_use` | [queue_id_to_use](data-sources--forwarding_class--reference--group-001.md#canonical-7a627cf7af1980bc3c0067c1d1e6f07fe8c79796e36cfd65d8495261c5bac8aa) |
| `tos_value` | [tos_value](data-sources--forwarding_class--reference--group-001.md#canonical-610e4ef53c88f156523970e90d14034b72a38ac6cbc05d75f9e813b05e6b1c26) |

<a id="canonical-e68cb12e21ab70b82fd996ac370a4dd729217f0a7799497e672b48e794a8f90c"></a>

## Next pages — Property reference / 48b82565670f / 14

- [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-b5887e4f1044fed8d5f66820dfc5e38fdae20fd7a17d1251b0b18f5e5569b38a)
- [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-c0f512f23169993bc10a23d98ba0be21042cc383cc3c43e47b20f7e6f75f9771)
- [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-8488a3293531055deb7ce7e847c402ac4bd1619d95cef2b7486f363204ab99be)
- [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-1e086fadeca401981e969f1d16cd6d259ea979cd2b24d89ee096642e238a0963)
- [policer](data-sources--forwarding_class--reference--group-001.md#canonical-712b7d1959dea32cb31c6a555103e15b43f4ff016c8c62d8354b66bca8364cc7)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)

<a id="canonical-b5887e4f1044fed8d5f66820dfc5e38fdae20fd7a17d1251b0b18f5e5569b38a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aef2b41b359efae771be8cfd317c1adf499c85c0b6e767537a0fcb7f3e8e064"></a>

## dscp — dscp / a7f76c873eda / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- dscp

<a id="canonical-1865fb53596c6eec8139d1979e8f5dfb0db7a820e2a9ca4c58d0fea8a0feab43"></a>

Type: `"single"`. Computed.

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

- [dscp](data-sources--forwarding_class--reference--group-001.md#canonical-1865fb53596c6eec8139d1979e8f5dfb0db7a820e2a9ca4c58d0fea8a0feab43)
- [no_marking](data-sources--forwarding_class--reference--group-001.md#canonical-34d62f4cb92016e3412dd1383e1b8f7c52fe2bfdc84d38ac19a38b13ec3ec78f)
- [tos_value](data-sources--forwarding_class--reference--group-001.md#canonical-610e4ef53c88f156523970e90d14034b72a38ac6cbc05d75f9e813b05e6b1c26)

Select alternatives according to the provider validators above.

<a id="canonical-01d7f0e9c4d1bdf12555a5cbebc2883046f9170ced7d035ff1f56c6508d69871"></a>

## Direct properties — dscp / a7f76c873eda / 3

<a id="canonical-bddedc97b67c71437e933d086729c141c65c5e9db4f26c44767b2d4befa2dc32"></a>

<a id="canonical-f264243fabd913b43c1fc5e4ae76ce0a6e3176a6cd51f76db3bb747ec9e28dfc"></a>

## drop_precedence property — dscp / a7f76c873eda / 4

Type: `"string"`. Computed.

\[Enum: DSCP\_AF\_LOW|DSCP\_AF\_MEDIUM|DSCP\_AF\_HIGH|DSCP\_AF\_POLICER\] DSCP Assured forwarding
drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop
precedence value is taken from output of policer. Possible values are \`DSCP\_AF\_LOW\`,
\`DSCP\_AF\_MEDIUM\`, \`DSCP\_AF\_HIGH\`, \`DSCP\_AF\_POLICER\`.

Upstream description:

DSCP Assured forwarding drop precedence

DSCP Low drop precedence DSCP Low drop precedence DSCP Low drop precedence DSCP drop precedence
value is taken from output of policer.

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

<a id="canonical-4e81b717af2250c897965f8fbf276127570326988a01878ef6c6d93cbe35cf6d"></a>

<a id="canonical-fef115e1b2d5ce1e613356e108455f2e39b4c3858870a4f4d44eb5336d2dca90"></a>

## dscp_class property — dscp / a7f76c873eda / 5

Type: `"string"`. Computed.

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

<a id="canonical-8b3c3417940e94277d075b93411cdb9a01772ad97e1f245350fc89a293c4ef26"></a>

## Next pages — dscp / a7f76c873eda / 6

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)

<a id="canonical-c0f512f23169993bc10a23d98ba0be21042cc383cc3c43e47b20f7e6f75f9771"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc1b1ff0801eaed92a569d19cfd8a531e8ac7fe73fd2e15f73fcf1ebbfb6cce3"></a>

## dscp_based_queue — dscp_based_queue / c804e2b12363 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- dscp_based_queue

<a id="canonical-1278fdf2ce09c78740261b859af34ca815c8f5a88d27be312a752b3001136583"></a>

Type: `["object", {}]`. Computed.

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

- [dscp_based_queue](data-sources--forwarding_class--reference--group-001.md#canonical-1278fdf2ce09c78740261b859af34ca815c8f5a88d27be312a752b3001136583)
- [queue_id_to_use](data-sources--forwarding_class--reference--group-001.md#canonical-7a627cf7af1980bc3c0067c1d1e6f07fe8c79796e36cfd65d8495261c5bac8aa)

Select alternatives according to the provider validators above.

<a id="canonical-600bd62aaa715bc074dc71ba39142d5ca9ae6fe3f44fab544ccc29969e54d170"></a>

## Direct properties — dscp_based_queue / c804e2b12363 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2f4523653b0dc3052108fd7d15d743828ae9e9a6cf005b54d8b4b2d3c49f60c"></a>

## Next pages — dscp_based_queue / c804e2b12363 / 4

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)

<a id="canonical-8488a3293531055deb7ce7e847c402ac4bd1619d95cef2b7486f363204ab99be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b7fe5eda2f83c7b72e0663294d39795f352a9332e66c623002bf9f958f153e5"></a>

## no_marking — no_marking / 99979a94012e / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- no_marking

<a id="canonical-34d62f4cb92016e3412dd1383e1b8f7c52fe2bfdc84d38ac19a38b13ec3ec78f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1d13502268a8036687a67dae1b3af22c0d9e4a0e7cbbc9a11a55724177b43cff"></a>

## Direct properties — no_marking / 99979a94012e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e839522b99c0a014e6354e5121417ee6cb627faef64ec325c806400537f96e2e"></a>

## Next pages — no_marking / 99979a94012e / 4

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)

<a id="canonical-1e086fadeca401981e969f1d16cd6d259ea979cd2b24d89ee096642e238a0963"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-703a44a7fcf0d8845e5705e784c3488f733df267d9ea7267960122f171bd38d2"></a>

## no_policer — no_policer / 38b6bbdd1e11 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- no_policer

<a id="canonical-05453b340433107ccae17a1b474a8db81cb507fad22c70e5bc5d00d0649dd8b1"></a>

Type: `["object", {}]`. Computed.

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

- [no_policer](data-sources--forwarding_class--reference--group-001.md#canonical-05453b340433107ccae17a1b474a8db81cb507fad22c70e5bc5d00d0649dd8b1)
- [policer](data-sources--forwarding_class--reference--group-001.md#canonical-9467423299469f0cdef085e3d7017c929c3516a6c78503888c12a7ad4430378b)

Select alternatives according to the provider validators above.

<a id="canonical-91eff19a6bcf0f4b13a45b4830e985681b026dec7d98e383ab053a74cf92873a"></a>

## Direct properties — no_policer / 38b6bbdd1e11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7338956bee355eb6c32e2315ed7c8ac3110422c12010dbfccb9c9ce23186024a"></a>

## Next pages — no_policer / 38b6bbdd1e11 / 4

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)

<a id="canonical-712b7d1959dea32cb31c6a555103e15b43f4ff016c8c62d8354b66bca8364cc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16d88c6cf4541d9bbc03b414fc148c205eff44ad4b067854d0922f347e186a39"></a>

## policer — policer / c5b7185d0ea6 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- policer

<a id="canonical-9467423299469f0cdef085e3d7017c929c3516a6c78503888c12a7ad4430378b"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-33883b74f8c3f136084f206174f631d39456e018d7d10a452644fb27dd471b36"></a>

## Direct properties — policer / c5b7185d0ea6 / 3

<a id="canonical-ac417417edd7530661d45b792365ede3246dbbd6d813cc361e06d4e3dce669e9"></a>

<a id="canonical-a8400ccff2f2394134e2d1bc5b9f8dbae272e2875c55462045ac3dbfca05340a"></a>

## name property — policer / c5b7185d0ea6 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-9015a8f0de81b1fa5a7828da4e1c39a4e7872683fc098f17d565de0ba0ef76a3"></a>

<a id="canonical-66b100ffc2d102fdf7fa4eb32480f7ecf650ecc622a4e7f5faac9fd3813273c5"></a>

## namespace property — policer / c5b7185d0ea6 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-7996a1ed122235183398dbe62ddc5a2638434217f33a684e7c0030201d0edaa8"></a>

<a id="canonical-280b2eb4b7a1c83da0e100d255bbbc45aacb269e42e539f5255935aeb013a245"></a>

## tenant property — policer / c5b7185d0ea6 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-7a238a65f19b0e2a9d4a6e5fa3f0814cf051f420fa234735690adcf2026be200"></a>

## Next pages — policer / c5b7185d0ea6 / 7

- [Property reference](data-sources--forwarding_class--reference--group-001.md#canonical-6ec71ff62f075a2e37a43517a8a9b211feae94a629c2e1fc0af087e3ac82abe7)
- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-3b133379aa5b31b88a0783e6e5d2d0a53474466c614a38dd2cbdc055e1e46e23)
