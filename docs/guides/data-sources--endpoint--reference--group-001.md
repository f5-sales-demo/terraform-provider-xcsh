---
page_title: "xcsh_endpoint reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint reference."
---

# xcsh_endpoint reference

<a id="canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2db4c501b53bb882bddfd19c7f769ec6bae68e51ac5b948f2f68cc2224ebdfe4"></a>

## Property reference — Property reference / 807fc46966a7 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- Property reference

<a id="canonical-ec72da0b352ab211945d632c6f0217d44937b48f347439cbac679f7123b77fb1"></a>

## Direct properties — Property reference / 807fc46966a7 / 3

<a id="canonical-f9e54189d4245957845e8cf35e94e1b8a9c3cbf5a32da47a23f83bb52a5eca0b"></a>

<a id="canonical-116d93d9e16836bb94ba1c1e8294dfb3e56924aa23a79628e83dbbeec04151c9"></a>

## annotations property — Property reference / 807fc46966a7 / 4

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

<a id="canonical-b17bc015a85bb5cc52f1a124d9e3fa1880ce6dc8e0d608ddf8565c665c0a77f9"></a>

<a id="canonical-f98d8a84f15172d5a1725106f0ecbd5c14e12b6edfb281cc5dee5d0830e9a80d"></a>

## description property — Property reference / 807fc46966a7 / 5

Type: `"string"`. Computed.

Description of the Endpoint.

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

<a id="canonical-8e08558436703d78acf72242d62f2a37a4bfd0501a72b933d11241d262151b67"></a>

<a id="canonical-2726f99d070046cbbb0d4d3cc88685311ac582e9fba85a42d00f36e6fdca851e"></a>

## dns_name property — Property reference / 807fc46966a7 / 6

Type: `"string"`. Computed.

\[OneOf: dns\_name, dns\_name\_advanced, ip, service\_info\] Exclusive with \[dns\_name\_advanced IP
service\_info\] Endpoint's IP address is discovered using DNS name resolution. The name given here
is fully qualified domain name.

Upstream description:

Exclusive with \[dns\_name\_advanced IP service\_info\] Endpoint's IP address is discovered using
DNS name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

OneOf alternatives in this subsection:

- [dns_name](data-sources--endpoint--reference--group-001.md#canonical-8e08558436703d78acf72242d62f2a37a4bfd0501a72b933d11241d262151b67)
- [dns_name_advanced](data-sources--endpoint--reference--group-001.md#canonical-3e72e61bb1f9ac688a1136213f4959e81ee359df3b066388b3add97ecaaad3cc)
- [ip](data-sources--endpoint--reference--group-001.md#canonical-cdf7e63d868e32c8f4b3725a15899e522d150471eaa74bd1acd3e93ea6a5a4d0)
- [service_info](data-sources--endpoint--reference--group-001.md#canonical-c02372e5969b85da12dc81968bbf646ff52a37a564932121566f9c49c3e87a40)

Select alternatives according to the provider validators above.

- [dns_name_advanced](data-sources--endpoint--reference--group-001.md#canonical-866d7aafef63c9c48034041f17869ba6f31a1209e7cbf4d6411baed7d02d1efb): complete subsection reference.

<a id="canonical-61086d1660d926c1990bd09d76a9eec35f23cc741cba69ee6d0112bdb2ace25d"></a>

<a id="canonical-d974f396a7faadd0c587fdc17442bb3358a58a44718cab5addcb27bfc6fd00ac"></a>

## health_check_port property — Property reference / 807fc46966a7 / 7

Type: `"number"`. Computed.

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

Upstream description:

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-db3d8c6b7ec6eee45f31fc590d79e0e135c75c734e3805f713b2ddff3893cde4"></a>

<a id="canonical-0cd5fdb7d085000f13320cfd3d5f283014a488ea50d83e1b5db83f05f5ac65e9"></a>

## id property — Property reference / 807fc46966a7 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-cdf7e63d868e32c8f4b3725a15899e522d150471eaa74bd1acd3e93ea6a5a4d0"></a>

<a id="canonical-97d19db071478131e51246bd00c71a9f6239fd787bdd18ad1425a73e0fb51ce7"></a>

## ip property — Property reference / 807fc46966a7 / 9

Type: `"string"`. Computed.

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

Upstream description:

Exclusive with \[dns\_name dns\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-efc7f8e187e796b274d6e00544c54980d7dfe701ae0b70a186c7df712516df4a"></a>

<a id="canonical-05b8fd3ac34d8fdcbf7b0e234e6042b84e29bf7ac988d488b8a528f4eb1ebfdd"></a>

## labels property — Property reference / 807fc46966a7 / 10

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

<a id="canonical-aa5c4d2a7d2bbc0f737317249c5b7e0a460407e8db39218ef953386166a6ef52"></a>

<a id="canonical-5d521aa42021887fc41b015039a92b6cee13f6d8b4cbf617b395bc36f6515e73"></a>

## name property — Property reference / 807fc46966a7 / 11

Type: `"string"`. Required.

Name of the Endpoint.

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

<a id="canonical-4321ece0953d526d956b7cbe85bf54767f3e6e2ee646963d4b810220b6c8b256"></a>

<a id="canonical-b6bf1fdd954e10de71f7b6c653fe01bcccba67c4e5abcb3cfc64006f982d1540"></a>

## namespace property — Property reference / 807fc46966a7 / 12

Type: `"string"`. Required.

Namespace where the Endpoint exists.

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

<a id="canonical-6713d4962c8b17b9cf646db29f657d3460b34efdde5baa322b23452b941276ef"></a>

<a id="canonical-57eb8d5cd695f8ed89eaa28e09282e69c00c6c7de581826c7e5d158075a513dd"></a>

## port property — Property reference / 807fc46966a7 / 13

Type: `"number"`. Computed.

Endpoint service is available on this port.

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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-07839a563ac30dcf273e4deafba495e1c144e232b708cd40e69ba51374cc19a2"></a>

<a id="canonical-2cf873eea6dcb0e16fb15314183e948d72b607fd1c0afb4d1bf9da40314ae5d5"></a>

## protocol property — Property reference / 807fc46966a7 / 14

Type: `"string"`. Computed.

\[Enum: TCP|UDP\] Protocol. Endpoint protocol. Default is TCP. Both TCP and UDP protocols are
supported. Possible values are \`TCP\`, \`UDP\`.

Upstream description:

Endpoint protocol. Default is TCP. Both TCP and UDP protocols are supported.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "TCP",
    "UDP"
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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [service_info](data-sources--endpoint--reference--group-001.md#canonical-b8e4e23d9f0eafe8ccfb8c65dc8f1bbaac349fdff3d1cfe6594f898ad027aeea): complete subsection reference.

- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-4be5cf2bc67dd275b7864b93a157bd4c14409d76dff07e211d56947f765a5a71): complete subsection reference.

- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014): complete subsection reference.

<a id="canonical-70e0331bce1ab0227fb8d08b58ba25fd4e9b05ef20c9db1b882ef95d31818851"></a>

## All schema paths — Property reference / 807fc46966a7 / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--endpoint--reference--group-001.md#canonical-f9e54189d4245957845e8cf35e94e1b8a9c3cbf5a32da47a23f83bb52a5eca0b) |
| `description` | [description](data-sources--endpoint--reference--group-001.md#canonical-b17bc015a85bb5cc52f1a124d9e3fa1880ce6dc8e0d608ddf8565c665c0a77f9) |
| `dns_name` | [dns_name](data-sources--endpoint--reference--group-001.md#canonical-8e08558436703d78acf72242d62f2a37a4bfd0501a72b933d11241d262151b67) |
| `dns_name_advanced` | [dns_name_advanced](data-sources--endpoint--reference--group-001.md#canonical-3e72e61bb1f9ac688a1136213f4959e81ee359df3b066388b3add97ecaaad3cc) |
| `dns_name_advanced.name` | [dns_name_advanced.name](data-sources--endpoint--reference--group-001.md#canonical-c8a9f12e475b79a226adeef3dacf672b0682176bb3f52f8af2bb0bc6fcf138af) |
| `dns_name_advanced.refresh_interval` | [dns_name_advanced.refresh_interval](data-sources--endpoint--reference--group-001.md#canonical-f0bc1011b96e1bf9fb62f93a88e815ba29d5c4597a17ce6cc43b5e18bebe0cd8) |
| `health_check_port` | [health_check_port](data-sources--endpoint--reference--group-001.md#canonical-61086d1660d926c1990bd09d76a9eec35f23cc741cba69ee6d0112bdb2ace25d) |
| `id` | [id](data-sources--endpoint--reference--group-001.md#canonical-db3d8c6b7ec6eee45f31fc590d79e0e135c75c734e3805f713b2ddff3893cde4) |
| `ip` | [ip](data-sources--endpoint--reference--group-001.md#canonical-cdf7e63d868e32c8f4b3725a15899e522d150471eaa74bd1acd3e93ea6a5a4d0) |
| `labels` | [labels](data-sources--endpoint--reference--group-001.md#canonical-efc7f8e187e796b274d6e00544c54980d7dfe701ae0b70a186c7df712516df4a) |
| `name` | [name](data-sources--endpoint--reference--group-001.md#canonical-aa5c4d2a7d2bbc0f737317249c5b7e0a460407e8db39218ef953386166a6ef52) |
| `namespace` | [namespace](data-sources--endpoint--reference--group-001.md#canonical-4321ece0953d526d956b7cbe85bf54767f3e6e2ee646963d4b810220b6c8b256) |
| `port` | [port](data-sources--endpoint--reference--group-001.md#canonical-6713d4962c8b17b9cf646db29f657d3460b34efdde5baa322b23452b941276ef) |
| `protocol` | [protocol](data-sources--endpoint--reference--group-001.md#canonical-07839a563ac30dcf273e4deafba495e1c144e232b708cd40e69ba51374cc19a2) |
| `service_info` | [service_info](data-sources--endpoint--reference--group-001.md#canonical-c02372e5969b85da12dc81968bbf646ff52a37a564932121566f9c49c3e87a40) |
| `service_info.discovery_type` | [service_info.discovery_type](data-sources--endpoint--reference--group-001.md#canonical-b260b33d8500896774a87b0096308117f18a6b1ca7602b0aeacc335f36bccd07) |
| `service_info.service_name` | [service_info.service_name](data-sources--endpoint--reference--group-001.md#canonical-44f22fa788adbca06d21f649fbf71acf93227eef23ea48de941f8608d8af73db) |
| `service_info.service_selector` | [service_info.service_selector](data-sources--endpoint--reference--group-001.md#canonical-44c434b708e8469917f5d69e72f2440a797425a4dd2f9f5553a511f6e4b37ff1) |
| `service_info.service_selector.expressions` | [service_info.service_selector.expressions](data-sources--endpoint--reference--group-001.md#canonical-9fa8033e44422f6880f06dd9257597001fecad690e4731dd81bb92b66290381a) |
| `snat_pool` | [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-90e634a417918754044d65c585a67493d25225ed3881c2f18df7dbcf5afa55f7) |
| `snat_pool.no_snat_pool` | [snat_pool.no_snat_pool](data-sources--endpoint--reference--group-001.md#canonical-cd0c9c68a3725b567d905bc2a9b5a2b6fd3e1a5bfcf500087c60ad2bc8e74133) |
| `snat_pool.snat_pool` | [snat_pool.snat_pool](data-sources--endpoint--reference--group-001.md#canonical-2c78dbeaacbf22b2c15bd6f3256a6ff4921539f4747819d03876c7a45d6efec7) |
| `snat_pool.snat_pool.prefixes` | [snat_pool.snat_pool.prefixes](data-sources--endpoint--reference--group-001.md#canonical-2699267a1c8795449679afd44a0dacd8d740eb9070fce7fa3ecf3488c3cb1681) |
| `where` | [where](data-sources--endpoint--reference--group-001.md#canonical-99fc03884a757fa3b6b85cc37651ed1086281d197fa1cb13cb30176454b472e5) |
| `where.site` | [where.site](data-sources--endpoint--reference--group-001.md#canonical-381f6bf54914b9e1394dfb646e86f677a034fee7f02a9c81049ad5b2480191aa) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-ff8a6abaf150893b5faea2e2fce78e56945d3dcf058f7b291ce38733cb7daf62) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-0a646008e514e2d3b8c5436caed7714054834e45e80df92b02b19ecb72042187) |
| `where.site.network_type` | [where.site.network_type](data-sources--endpoint--reference--group-001.md#canonical-518e1d6770e3a80b13010db61c4f71066ef48a9d28b26775096aa79c62180870) |
| `where.site.ref` | [where.site.ref](data-sources--endpoint--reference--group-001.md#canonical-01cdfbf04b5592b86a7e25e67c82dfda947889ccf91dff31e17bab7b94316639) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--endpoint--reference--group-001.md#canonical-6e88afc516f1118408c77f321f40469b6ad23d4a444a678d414c30bf633451d3) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--endpoint--reference--group-001.md#canonical-8cc8890b9ebd3b6f993a95a41b1e16564bcf0cd06c79d5c1dcad19187b16d2ac) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--endpoint--reference--group-001.md#canonical-d64b315c091a04d0f4146c1e1e855c255193bf769525d3ef7f041fef7fb2205f) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--endpoint--reference--group-001.md#canonical-0de7cf4ece88a8f4f8e8921aa5044fb606f22c5b661c70e9d29c3c4e2c162268) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--endpoint--reference--group-001.md#canonical-897656ace28d1b256893f6716e9d9dde8d47cdfdd2de9e28dd48ac07a36f27dd) |
| `where.virtual_network` | [where.virtual_network](data-sources--endpoint--reference--group-001.md#canonical-6f5d25bfe3837f83f5044b37de9196a47837e8c46559d5755775f715c89bc604) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--endpoint--reference--group-001.md#canonical-2328fb42942f7041468e0b5296cdc7a399855cad1df570663d69f46be084234b) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--endpoint--reference--group-001.md#canonical-b46b5a569ff021f0f1ee3e3fe9e0a050d2d0bdc7a507e15975784a7d91d46127) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--endpoint--reference--group-001.md#canonical-c1f07c114525dfe03c144f28e7b849db0ab32d65cf05a377a42ea5a0ca57d87c) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--endpoint--reference--group-001.md#canonical-ff9c4807d7775e3c3d6bbe921d236555f3b0a8e97feb4d59aca717354ad63abe) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--endpoint--reference--group-001.md#canonical-bf83401eb9c11eb9d9bd3cb8cf9bbf84aea51c05b31fbb5d41d79a8421f3c272) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--endpoint--reference--group-001.md#canonical-8f2ec99cd5e827f7aab78c920755e871476cf88008b086e03c00feac9fee1833) |
| `where.virtual_site` | [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-db39a57bf0248142247a2e5e8a79bd3903273da13741a7c637f9bbc9846d71f3) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-687e8dc08e468b12ea52dac0562164fe95ac5b0fff2cd15fee9e6f30997f9e28) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-2832066795cbe2820bd6e327a476960554d2a55098c126a716716aa8c8a53c87) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--endpoint--reference--group-001.md#canonical-d03cebb0aa62878e5ed8c5ceef4dd622a17de1dae48985b6ce7c5169317a05ee) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--endpoint--reference--group-001.md#canonical-1e9d574ed98768fad5f60dcba5c8be65a57ab230779bf07c9c2002fcf9b76ab8) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--endpoint--reference--group-001.md#canonical-a4e116335245776bc7297bed8d15a2aabddcb791cbf23027b9a2b3427234a5ec) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--endpoint--reference--group-001.md#canonical-7da1b95f7046dc2cd3eebcf9c803cf8ae37e371ab8468a732bd56e3a307dde19) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--endpoint--reference--group-001.md#canonical-c9b0b62100110e8bf7c8addb74ca8adaa8349f7761c7ff324a383c86d00e9692) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--endpoint--reference--group-001.md#canonical-a07011bcb3736b4bdc455a574810d7c1c89f270bcce05b4e8ca2a23dae610786) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--endpoint--reference--group-001.md#canonical-543e590e715d7a9f817b8e518c729a4bb300e990d8e76070c6db5133cb7bf584) |

<a id="canonical-088a21c4df0491d09f0dd1fd377e0a0c0660a684e54cd8b755b9191ec17a6187"></a>

## Next pages — Property reference / 807fc46966a7 / 16

- [dns_name_advanced](data-sources--endpoint--reference--group-001.md#canonical-866d7aafef63c9c48034041f17869ba6f31a1209e7cbf4d6411baed7d02d1efb)
- [service_info](data-sources--endpoint--reference--group-001.md#canonical-b8e4e23d9f0eafe8ccfb8c65dc8f1bbaac349fdff3d1cfe6594f898ad027aeea)
- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-4be5cf2bc67dd275b7864b93a157bd4c14409d76dff07e211d56947f765a5a71)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-866d7aafef63c9c48034041f17869ba6f31a1209e7cbf4d6411baed7d02d1efb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb071bc6a575f4d6fa3d39346547f99e3cf5eaaa7db7381d1a2a6259d861fddd"></a>

## dns_name_advanced — dns_name_advanced / f7524373900e / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- dns_name_advanced

<a id="canonical-3e72e61bb1f9ac688a1136213f4959e81ee359df3b066388b3add97ecaaad3cc"></a>

Type: `"single"`. Computed.

Specifies name and TTL used for DNS resolution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ttl_choice": "[\"refresh_interval\"]"
}
```

<a id="canonical-6079a6f82c1e98d2012a091085e1e6e11e71b3f50c9346ac7c1bf3ad3ac56f85"></a>

## Direct properties — dns_name_advanced / f7524373900e / 3

<a id="canonical-c8a9f12e475b79a226adeef3dacf672b0682176bb3f52f8af2bb0bc6fcf138af"></a>

<a id="canonical-09a376d6ccf827b8727039c52ddade28b4f163cc4683a7d251325008993ea073"></a>

## name property — dns_name_advanced / f7524373900e / 4

Type: `"string"`. Computed.

Endpoint's IP address is discovered using DNS name resolution. The name given here is fully
qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-f0bc1011b96e1bf9fb62f93a88e815ba29d5c4597a17ce6cc43b5e18bebe0cd8"></a>

<a id="canonical-9a9bd6dbb94c42d57515b4af03f518ff4944a3f13ac6e46e78f4f45b02c5e98e"></a>

## refresh_interval property — dns_name_advanced / f7524373900e / 5

Type: `"number"`. Computed.

Exclusive with \[\] Interval for DNS refresh in seconds.

Upstream description:

Exclusive with \[\] Interval for DNS refresh in seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  }
}
```

<a id="canonical-7131e21acc06f83e2ed8317abc9f62374691a83c0b5c354858a402870e73b237"></a>

## Next pages — dns_name_advanced / f7524373900e / 6

- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-b8e4e23d9f0eafe8ccfb8c65dc8f1bbaac349fdff3d1cfe6594f898ad027aeea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34bdd044e670e639d5f2ceba209974259f9e10dd69dcc98a4dbbef029200bb00"></a>

## service_info — service_info / 114410a83443 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- service_info

<a id="canonical-c02372e5969b85da12dc81968bbf646ff52a37a564932121566f9c49c3e87a40"></a>

Type: `"single"`. Computed.

Specifies whether endpoint service is discovered by name or labels.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_info": "[\"service_name\",\"service_selector\"]"
}
```

<a id="canonical-cb3645cc7a2526a594a6763d0ca791213a7a1a3da33a1c37c983958cd295bd73"></a>

## Direct properties — service_info / 114410a83443 / 3

<a id="canonical-b260b33d8500896774a87b0096308117f18a6b1ca7602b0aeacc335f36bccd07"></a>

<a id="canonical-2852247b9c6a644c716e678a7abc574ffa77d4ac4b9ba2c0d70b3fad4c96dd52"></a>

## discovery_type property — service_info / 114410a83443 / 4

Type: `"string"`. Computed.

\[Enum: INVALID\_DISCOVERY|K8S|CONSUL|CLASSIC\_BIGIP|THIRD\_PARTY|NGINX\_ONE\] Specifies the type of
discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service
Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.
Possible values are \`INVALID\_DISCOVERY\`, \`K8S\`, \`CONSUL\`, \`CLASSIC\_BIGIP\`,
\`THIRD\_PARTY\`, \`NGINX\_ONE\`. Defaults to \`INVALID\_DISCOVERY\`.

Upstream description:

Specifies the type of discovery

Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service Discover
from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALID_DISCOVERY",
  "enum": [
    "INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-44f22fa788adbca06d21f649fbf71acf93227eef23ea48de941f8608d8af73db"></a>

<a id="canonical-c3f26fc3cb22e753bc445222c4b9401d5bb6788e2bf9f86635941a533b1ef193"></a>

## service_name property — service_info / 114410a83443 / 5

Type: `"string"`. Computed.

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the..

Upstream description:

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the cluster identifier matches. If cluster identifier is not specified then discovery will be
done in all discovery objects of the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [service_selector](data-sources--endpoint--reference--group-001.md#canonical-4370456be517af62124ffe45cd68e5d399d1dfa528d35454c5af12a22e925969): complete subsection reference.

<a id="canonical-008a8fd2b6d2266d36f02d234a0b8c768c618a7f51367c4bcc8e53fc32dc5dda"></a>

## Next pages — service_info / 114410a83443 / 6

- [service_info.service_selector](data-sources--endpoint--reference--group-001.md#canonical-4370456be517af62124ffe45cd68e5d399d1dfa528d35454c5af12a22e925969)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-4370456be517af62124ffe45cd68e5d399d1dfa528d35454c5af12a22e925969"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f598df04c426a55e99fb0aa045b45c29041b76abee91331cbac22d5093d7a70"></a>

## service_info.service_selector — service_info.service_selector / 7791c7a0b437 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [service_info](data-sources--endpoint--reference--group-001.md#canonical-b8e4e23d9f0eafe8ccfb8c65dc8f1bbaac349fdff3d1cfe6594f898ad027aeea)
- service_info.service_selector

<a id="canonical-44c434b708e8469917f5d69e72f2440a797425a4dd2f9f5553a511f6e4b37ff1"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f9bc3a275e8f1c085a4fe249dcaa1a74bb048e89c8b82a2b9454997d8699c2dd"></a>

## Direct properties — service_info.service_selector / 7791c7a0b437 / 3

<a id="canonical-9fa8033e44422f6880f06dd9257597001fecad690e4731dd81bb92b66290381a"></a>

<a id="canonical-be00df3a768afc198cf62bf0ae8b1b26a19573b25a2d86620fd3cc6175e9a198"></a>

## expressions property — service_info.service_selector / 7791c7a0b437 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-197c0752991ed5105412ab7c57166983f5ceefcfbd261f44b20b6d65999777c5"></a>

## Next pages — service_info.service_selector / 7791c7a0b437 / 5

- [service_info](data-sources--endpoint--reference--group-001.md#canonical-b8e4e23d9f0eafe8ccfb8c65dc8f1bbaac349fdff3d1cfe6594f898ad027aeea)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-4be5cf2bc67dd275b7864b93a157bd4c14409d76dff07e211d56947f765a5a71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2b1a15a32ede3e6194f826fdd611b881e5b05f74069d0dc46f66f7a10f6c119"></a>

## snat_pool — snat_pool / 1efbb2ec35d5 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- snat_pool

<a id="canonical-90e634a417918754044d65c585a67493d25225ed3881c2f18df7dbcf5afa55f7"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-8a4d626da7dedb3060dbb9340c55aa28f8f063aae36deac361d750b1def207b1"></a>

## Direct properties — snat_pool / 1efbb2ec35d5 / 3

- [no_snat_pool](data-sources--endpoint--reference--group-001.md#canonical-f2356d9e4fdb9d03a460ada26f9793874bb5e638e9900fe247f6721a22a69bc9): complete subsection reference.

- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-0f31da38b81362d41550345ac655adf47d476df4bf04f8874cfad021780cdf9c): complete subsection reference.

<a id="canonical-542cc757aecbf4ba3f9afb946c27bdc671a945ba449852dc5de8340144f5a670"></a>

## Next pages — snat_pool / 1efbb2ec35d5 / 4

- [snat_pool.no_snat_pool](data-sources--endpoint--reference--group-001.md#canonical-f2356d9e4fdb9d03a460ada26f9793874bb5e638e9900fe247f6721a22a69bc9)
- [snat_pool.snat_pool](data-sources--endpoint--reference--group-001.md#canonical-0f31da38b81362d41550345ac655adf47d476df4bf04f8874cfad021780cdf9c)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-f2356d9e4fdb9d03a460ada26f9793874bb5e638e9900fe247f6721a22a69bc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efaf8e55b46f3d4c86ae2cd8d2d6590b2fe7c8b074a83c3736ef711d86b8c04d"></a>

## snat_pool.no_snat_pool — snat_pool.no_snat_pool / 8b8ba713b3ab / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-4be5cf2bc67dd275b7864b93a157bd4c14409d76dff07e211d56947f765a5a71)
- snat_pool.no_snat_pool

<a id="canonical-cd0c9c68a3725b567d905bc2a9b5a2b6fd3e1a5bfcf500087c60ad2bc8e74133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-af5f8b7c3f38b06e3ad6271d5c6af09d8d300bb07f9ed694269c7e60aba5b21e"></a>

## Direct properties — snat_pool.no_snat_pool / 8b8ba713b3ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e65fb163a4b7ce5756d3edf97a7f73c7c5bb1bca354c243b71bc030aa07f570"></a>

## Next pages — snat_pool.no_snat_pool / 8b8ba713b3ab / 4

- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-4be5cf2bc67dd275b7864b93a157bd4c14409d76dff07e211d56947f765a5a71)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-0f31da38b81362d41550345ac655adf47d476df4bf04f8874cfad021780cdf9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd65795ff0f7efcdf79d04b354e677ef80cb5340a9ecc6549c461da86f39ec75"></a>

## snat_pool.snat_pool — snat_pool.snat_pool / 20b46ec342f7 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-4be5cf2bc67dd275b7864b93a157bd4c14409d76dff07e211d56947f765a5a71)
- snat_pool.snat_pool

<a id="canonical-2c78dbeaacbf22b2c15bd6f3256a6ff4921539f4747819d03876c7a45d6efec7"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-13f5928aed48fe689645584fe736f5b4526c685ceac00339fe56c3c60d412df6"></a>

## Direct properties — snat_pool.snat_pool / 20b46ec342f7 / 3

<a id="canonical-2699267a1c8795449679afd44a0dacd8d740eb9070fce7fa3ecf3488c3cb1681"></a>

<a id="canonical-13135abd47f07fbf1d512afe8f91293c34bfdc92504a3287edaf1cfde066a5c4"></a>

## prefixes property — snat_pool.snat_pool / 20b46ec342f7 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0e7de5feca2b162f9bb735aeceb5a2de5a0e3bef99a6b21c607732d1f215d808"></a>

## Next pages — snat_pool.snat_pool / 20b46ec342f7 / 5

- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-4be5cf2bc67dd275b7864b93a157bd4c14409d76dff07e211d56947f765a5a71)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f30ede0d163ee47570e3f30b06997de0a9d9d48c41b3c39faa8ab49e6120f2e"></a>

## where — where / 4c5563b6f46f / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- where

<a id="canonical-99fc03884a757fa3b6b85cc37651ed1086281d197fa1cb13cb30176454b472e5"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-2125cc6f198d339cb1f15f8dd1086e9d194874ed06faf6005b2255c714785b45"></a>

## Direct properties — where / 4c5563b6f46f / 3

- [site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0): complete subsection reference.

- [virtual_network](data-sources--endpoint--reference--group-001.md#canonical-206ac2602617f3ab6d6869462c565d6ced235d3bd671c6d089bc3433d8c9e2a9): complete subsection reference.

- [virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0): complete subsection reference.

<a id="canonical-16361461ddce9138c1cc7c152e76ba302e9982161c9e9d75e636af14573712eb"></a>

## Next pages — where / 4c5563b6f46f / 4

- [where.site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0)
- [where.virtual_network](data-sources--endpoint--reference--group-001.md#canonical-206ac2602617f3ab6d6869462c565d6ced235d3bd671c6d089bc3433d8c9e2a9)
- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d62f725bb768af6dca347fd5e350363a400ecc2726e4a20c3ef20a3d3b6e5941"></a>

## where.site — where.site / 1dd483ae6a5b / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- where.site

<a id="canonical-381f6bf54914b9e1394dfb646e86f677a034fee7f02a9c81049ad5b2480191aa"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-207ca922e25f2a3ea16e326299be2ca536b39585f7767cd868fc23d37b02198c"></a>

## Direct properties — where.site / 1dd483ae6a5b / 3

- [disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-d26b9b5c74cec61c63e1b35c2343f6625a4c7261044d1882ea7525b1cc476e9d): complete subsection reference.

- [enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-edab4084323e562950f5cc6e70b56570eb71cbd779c918f622e5ef0c8f1db33d): complete subsection reference.

<a id="canonical-518e1d6770e3a80b13010db61c4f71066ef48a9d28b26775096aa79c62180870"></a>

<a id="canonical-bb37b5cf29af119b31aafe586b80898526b3e98fd2b176db3b9e43ebaf84ee6e"></a>

## network_type property — where.site / 1dd483ae6a5b / 4

Type: `"string"`. Computed.

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

- [ref](data-sources--endpoint--reference--group-001.md#canonical-85619c0b8f85e8cf545101b0416290795dc042ee8b6064f71a1cedb9e061f569): complete subsection reference.

<a id="canonical-21f33371e48c0e0019a0212ca5be612f094a72fbfc8d52aa05b347c18b9768fb"></a>

## Next pages — where.site / 1dd483ae6a5b / 5

- [where.site.disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-d26b9b5c74cec61c63e1b35c2343f6625a4c7261044d1882ea7525b1cc476e9d)
- [where.site.enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-edab4084323e562950f5cc6e70b56570eb71cbd779c918f622e5ef0c8f1db33d)
- [where.site.ref](data-sources--endpoint--reference--group-001.md#canonical-85619c0b8f85e8cf545101b0416290795dc042ee8b6064f71a1cedb9e061f569)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-d26b9b5c74cec61c63e1b35c2343f6625a4c7261044d1882ea7525b1cc476e9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-631c5093a0c4c3b5f6c39990746e27c30eff95172868ec5fb68eadcb3fc30a31"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / ab65bd731713 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [where.site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0)
- where.site.disable_internet_vip

<a id="canonical-ff8a6abaf150893b5faea2e2fce78e56945d3dcf058f7b291ce38733cb7daf62"></a>

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

<a id="canonical-f62a0428754de79539ea338c597b5ce9f8522e7b4ea9774b9c9c1fc969bad38d"></a>

## Direct properties — where.site.disable_internet_vip / ab65bd731713 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf2cf27202bef568106ece07a29a09e4ea1d33ee220e263016f17554f48eabd8"></a>

## Next pages — where.site.disable_internet_vip / ab65bd731713 / 4

- [where.site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-edab4084323e562950f5cc6e70b56570eb71cbd779c918f622e5ef0c8f1db33d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea6be84f93761f742739eecc52f8bac2f7d4eca2976a9f3bc16bf2131b5d1093"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / 7eae38ab9f02 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [where.site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0)
- where.site.enable_internet_vip

<a id="canonical-0a646008e514e2d3b8c5436caed7714054834e45e80df92b02b19ecb72042187"></a>

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

<a id="canonical-2adf30015138834e840a480c1eed4935ee401baf6a8f97fefcc3cf8b829f3450"></a>

## Direct properties — where.site.enable_internet_vip / 7eae38ab9f02 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd41628957c888165654fea25b693bde44d0b35b5b728646c0eaf0c7a1f534db"></a>

## Next pages — where.site.enable_internet_vip / 7eae38ab9f02 / 4

- [where.site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-85619c0b8f85e8cf545101b0416290795dc042ee8b6064f71a1cedb9e061f569"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a18895f74603d9adc9e0e35d1bc54a13ddda9529308695507aa334e58e32d7e6"></a>

## where.site.ref — where.site.ref / 3f38089fc9df / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [where.site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0)
- where.site.ref

<a id="canonical-01cdfbf04b5592b86a7e25e67c82dfda947889ccf91dff31e17bab7b94316639"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-c26f6bccd81aa02a95ab07ca79cd6651e015ff802d663f8ac492be519be2fbb8"></a>

## Direct properties — where.site.ref / 3f38089fc9df / 3

<a id="canonical-6e88afc516f1118408c77f321f40469b6ad23d4a444a678d414c30bf633451d3"></a>

<a id="canonical-39db5c372dfb218ff0e28896998f9f06c1c80a0beee5c0515dcd0a929407a1f9"></a>

## kind property — where.site.ref / 3f38089fc9df / 4

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

<a id="canonical-8cc8890b9ebd3b6f993a95a41b1e16564bcf0cd06c79d5c1dcad19187b16d2ac"></a>

<a id="canonical-2cdd01ce0dfb2757ad0049ffdaf89d9043963ee89f2df50c1fa296f378e2e598"></a>

## name property — where.site.ref / 3f38089fc9df / 5

Type: `"string"`. Computed.

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

<a id="canonical-d64b315c091a04d0f4146c1e1e855c255193bf769525d3ef7f041fef7fb2205f"></a>

<a id="canonical-a3fd21b19375fb6fd63550c11db7b7193aa0c99c359afa31a6273538b06f1521"></a>

## namespace property — where.site.ref / 3f38089fc9df / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0de7cf4ece88a8f4f8e8921aa5044fb606f22c5b661c70e9d29c3c4e2c162268"></a>

<a id="canonical-0263447aadebb0c6bca2c7c46ec0923236a60d02fda0dea7d4ed7425994a74a2"></a>

## tenant property — where.site.ref / 3f38089fc9df / 7

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

<a id="canonical-897656ace28d1b256893f6716e9d9dde8d47cdfdd2de9e28dd48ac07a36f27dd"></a>

<a id="canonical-76f872c2ee8e93a86151eefc75eef6dbe42f2e302e984bb1dfa23619ac4ee227"></a>

## uid property — where.site.ref / 3f38089fc9df / 8

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

<a id="canonical-cde8c6e95610296d2b0286e573879e0f56379b731e8ef79837e2a6b07bc243ec"></a>

## Next pages — where.site.ref / 3f38089fc9df / 9

- [where.site](data-sources--endpoint--reference--group-001.md#canonical-872d3c582f00644bf2ed8434d9c13d528267170691fe83a25b95529a95a213b0)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-206ac2602617f3ab6d6869462c565d6ced235d3bd671c6d089bc3433d8c9e2a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7b09ce061740788650d7d20a37f9f31c8deecfa34ad36972216c97ea4e02c32"></a>

## where.virtual_network — where.virtual_network / 91c42e924fc5 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- where.virtual_network

<a id="canonical-6f5d25bfe3837f83f5044b37de9196a47837e8c46559d5755775f715c89bc604"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e240dda2f4ca43ff06ec6821e7302d163671dabb3ea3136a18d85b0344ca35b9"></a>

## Direct properties — where.virtual_network / 91c42e924fc5 / 3

- [ref](data-sources--endpoint--reference--group-001.md#canonical-73e6adffa0f56b1ce663264623f52b6f672d8f4481830d97d96085c879774fd4): complete subsection reference.

<a id="canonical-0d947752bfd52295df87b1b00b4bc504da8704378a951fb40380e6250d1af972"></a>

## Next pages — where.virtual_network / 91c42e924fc5 / 4

- [where.virtual_network.ref](data-sources--endpoint--reference--group-001.md#canonical-73e6adffa0f56b1ce663264623f52b6f672d8f4481830d97d96085c879774fd4)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-73e6adffa0f56b1ce663264623f52b6f672d8f4481830d97d96085c879774fd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-776d7ac811683b7c4e4295407567896822b2e392ade5f9d74ef4c83e2bbcb8ac"></a>

## where.virtual_network.ref — where.virtual_network.ref / 1e14ebf6f93e / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [where.virtual_network](data-sources--endpoint--reference--group-001.md#canonical-206ac2602617f3ab6d6869462c565d6ced235d3bd671c6d089bc3433d8c9e2a9)
- where.virtual_network.ref

<a id="canonical-2328fb42942f7041468e0b5296cdc7a399855cad1df570663d69f46be084234b"></a>

Type: `"list"`. Computed.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-97ceca9af3a3cbbe9397c49b47c4f448ba73c42984f3d462d11bb5b74643a336"></a>

## Direct properties — where.virtual_network.ref / 1e14ebf6f93e / 3

<a id="canonical-b46b5a569ff021f0f1ee3e3fe9e0a050d2d0bdc7a507e15975784a7d91d46127"></a>

<a id="canonical-f9ea1f5aa32db6aae4ac705aa432f37b8994c73958400067e604bfc1a9ba261e"></a>

## kind property — where.virtual_network.ref / 1e14ebf6f93e / 4

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

<a id="canonical-c1f07c114525dfe03c144f28e7b849db0ab32d65cf05a377a42ea5a0ca57d87c"></a>

<a id="canonical-7b6b6d00b5a1b8c710bc111d2037952de358a9ebe401b45e162e4209ceefdc69"></a>

## name property — where.virtual_network.ref / 1e14ebf6f93e / 5

Type: `"string"`. Computed.

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

<a id="canonical-ff9c4807d7775e3c3d6bbe921d236555f3b0a8e97feb4d59aca717354ad63abe"></a>

<a id="canonical-1f3cea063faca702a911f46bea17da47c441688ed98fd2d46573b62fa44fc8a5"></a>

## namespace property — where.virtual_network.ref / 1e14ebf6f93e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-bf83401eb9c11eb9d9bd3cb8cf9bbf84aea51c05b31fbb5d41d79a8421f3c272"></a>

<a id="canonical-c3c8d7316fde82338308df3a3b29742cdf54e635147d3140b0a714e6ba685a07"></a>

## tenant property — where.virtual_network.ref / 1e14ebf6f93e / 7

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

<a id="canonical-8f2ec99cd5e827f7aab78c920755e871476cf88008b086e03c00feac9fee1833"></a>

<a id="canonical-75d2419cf5f595a3c30d5c1fd81d7dfb3c4ac140ff2a606e1628d4968e924814"></a>

## uid property — where.virtual_network.ref / 1e14ebf6f93e / 8

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

<a id="canonical-67d7c6deb159b112cc610f674ac0d4ae6ecb03f05375f79e29525a5d72e4db87"></a>

## Next pages — where.virtual_network.ref / 1e14ebf6f93e / 9

- [where.virtual_network](data-sources--endpoint--reference--group-001.md#canonical-206ac2602617f3ab6d6869462c565d6ced235d3bd671c6d089bc3433d8c9e2a9)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ff6082868220f2f52c983ac20296eb68d3bc944a04052acff1fe99b9647d351"></a>

## where.virtual_site — where.virtual_site / 5df3d30d1cef / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- where.virtual_site

<a id="canonical-db39a57bf0248142247a2e5e8a79bd3903273da13741a7c637f9bbc9846d71f3"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-3377b80fdced68d283d309d39ab696c326f596f8a1f01a8b143ab45220012ea7"></a>

## Direct properties — where.virtual_site / 5df3d30d1cef / 3

- [disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-ab332a1b0c8a9273cbf869df5990a60e839c95d7d4a539df222ae81435b64737): complete subsection reference.

- [enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-55e09bcaf9862e54b0da2287e4bf3f6101f04f7e2d7950d094cb084e11c1d62d): complete subsection reference.

<a id="canonical-d03cebb0aa62878e5ed8c5ceef4dd622a17de1dae48985b6ce7c5169317a05ee"></a>

<a id="canonical-cce2e97954005b9234c067250d78dbcc0791304a5a3e1607cac7fd27334611fb"></a>

## network_type property — where.virtual_site / 5df3d30d1cef / 4

Type: `"string"`. Computed.

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

- [ref](data-sources--endpoint--reference--group-001.md#canonical-8a6a1aa4a777a921c940fc8a60ea8d3b7ba4837c083545f1d6ac2483b953b4b1): complete subsection reference.

<a id="canonical-32eda7ae9170220a7734ab79c6657f4e0f467f8bb1c4602bf87b9552c1e26043"></a>

## Next pages — where.virtual_site / 5df3d30d1cef / 5

- [where.virtual_site.disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-ab332a1b0c8a9273cbf869df5990a60e839c95d7d4a539df222ae81435b64737)
- [where.virtual_site.enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-55e09bcaf9862e54b0da2287e4bf3f6101f04f7e2d7950d094cb084e11c1d62d)
- [where.virtual_site.ref](data-sources--endpoint--reference--group-001.md#canonical-8a6a1aa4a777a921c940fc8a60ea8d3b7ba4837c083545f1d6ac2483b953b4b1)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-ab332a1b0c8a9273cbf869df5990a60e839c95d7d4a539df222ae81435b64737"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c62cc3e90946aa8ba87548d9c9049c0c83914a65a7d454abf38eb8d7dbd2a05"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / 1a2d370bc24e / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0)
- where.virtual_site.disable_internet_vip

<a id="canonical-687e8dc08e468b12ea52dac0562164fe95ac5b0fff2cd15fee9e6f30997f9e28"></a>

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

<a id="canonical-3ea2a63010f436d6f2d315a426c96e0d518dded8da4e964af389cb430df6f7f5"></a>

## Direct properties — where.virtual_site.disable_internet_vip / 1a2d370bc24e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2324f44738a35d035184783e7877ebd585f8355defe67c89b092fb2b7b7c2c4b"></a>

## Next pages — where.virtual_site.disable_internet_vip / 1a2d370bc24e / 4

- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-55e09bcaf9862e54b0da2287e4bf3f6101f04f7e2d7950d094cb084e11c1d62d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-772b24adc01b6b9be70b9c7425d1f8fd8651f94afa57ab9fbdf574d948081973"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / e360f1e16712 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0)
- where.virtual_site.enable_internet_vip

<a id="canonical-2832066795cbe2820bd6e327a476960554d2a55098c126a716716aa8c8a53c87"></a>

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

<a id="canonical-9861a8a643010a50b3087921385db5651d1e5f0192b830dbb9d6e5475f75c05f"></a>

## Direct properties — where.virtual_site.enable_internet_vip / e360f1e16712 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11ba9387f786dcb8bf94ff10543edbf63d7b0617b30764d762a15b18cef501d5"></a>

## Next pages — where.virtual_site.enable_internet_vip / e360f1e16712 / 4

- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-8a6a1aa4a777a921c940fc8a60ea8d3b7ba4837c083545f1d6ac2483b953b4b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26027a28515f65c5664c476430686a71ddf194939464105bca8157eba3b26b8d"></a>

## where.virtual_site.ref — where.virtual_site.ref / c14a25a5d9f7 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-35bfc21ee54b7115fd6653a3c3c2e2fbf642183aadae40619ca774e9a22bb62b)
- [where](data-sources--endpoint--reference--group-001.md#canonical-29b30e4c1b9284f87fed14b05747e716d3b9733ce297534810544b567c997014)
- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0)
- where.virtual_site.ref

<a id="canonical-1e9d574ed98768fad5f60dcba5c8be65a57ab230779bf07c9c2002fcf9b76ab8"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-4a369a156c7e4aa7ed37190a7867db20d6847098c2365eaad9b6eab2e21ae185"></a>

## Direct properties — where.virtual_site.ref / c14a25a5d9f7 / 3

<a id="canonical-a4e116335245776bc7297bed8d15a2aabddcb791cbf23027b9a2b3427234a5ec"></a>

<a id="canonical-3f2405eba27d81da6c01c0fbfc00927fad7021c6fc75f4ddcf16920f1c65b1e0"></a>

## kind property — where.virtual_site.ref / c14a25a5d9f7 / 4

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

<a id="canonical-7da1b95f7046dc2cd3eebcf9c803cf8ae37e371ab8468a732bd56e3a307dde19"></a>

<a id="canonical-d188e1eda3550a7677ff851ccce263484240579262d87914283501880ab43f8a"></a>

## name property — where.virtual_site.ref / c14a25a5d9f7 / 5

Type: `"string"`. Computed.

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

<a id="canonical-c9b0b62100110e8bf7c8addb74ca8adaa8349f7761c7ff324a383c86d00e9692"></a>

<a id="canonical-b5fc4e97b341234cbecbef53c5265d2d5aae5ee1ef0fb38312bc7f270795ab2a"></a>

## namespace property — where.virtual_site.ref / c14a25a5d9f7 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-a07011bcb3736b4bdc455a574810d7c1c89f270bcce05b4e8ca2a23dae610786"></a>

<a id="canonical-ce76d41211fa57a2c9c43c7c026c89e1f1a6d663899cd836a9ac3c7226429d1c"></a>

## tenant property — where.virtual_site.ref / c14a25a5d9f7 / 7

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

<a id="canonical-543e590e715d7a9f817b8e518c729a4bb300e990d8e76070c6db5133cb7bf584"></a>

<a id="canonical-ad132833c582dd278ac14662f7c76f601d8f9af5bb96dd11d70be38ffff67303"></a>

## uid property — where.virtual_site.ref / c14a25a5d9f7 / 8

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

<a id="canonical-78be3eb4e16b5b672b691bd2d1c600c46ff2782163b71a9ca680aaf69b44266e"></a>

## Next pages — where.virtual_site.ref / c14a25a5d9f7 / 9

- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-fa39216e10593b0028a63f6db89e56eb5299068a5ec925cbcd3c21e0424f84a0)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
