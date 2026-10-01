---
page_title: "xcsh_tunnel reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel reference."
---

# xcsh_tunnel reference

<a id="canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f40277aa9cdfdbf2be0c871665553446b46fba4be0d8e9df54fff8e2c2385c62"></a>

## Property reference — Property reference / 823f0b57e8ca / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- Property reference

<a id="canonical-9c24d1e254bc6d9cad222783b000ac513e4a53d12ec695e930c4f24cecbb5fc8"></a>

## Direct properties — Property reference / 823f0b57e8ca / 3

<a id="canonical-aebd0e8f2620c40ed15d5c018282251c5ad1c2408a5af701bfee8db0fa04e232"></a>

<a id="canonical-46ae4202ec3454d143d055758e28f8e288141f7942e28b8cfbad36bfebad71b2"></a>

## annotations property — Property reference / 823f0b57e8ca / 4

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

<a id="canonical-0d9a6bda77c5c5aded0afe0bbda1fad250fd90a1aad5722bf938baf6e0690534"></a>

<a id="canonical-fbd2c612b39863a38723e6c18b45f9d1c30b1176c794e2f9ba682609602252af"></a>

## description property — Property reference / 823f0b57e8ca / 5

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

<a id="canonical-8cd4772cb64a45e1c0fc2c36a56e301bd9efc27830a5ac2e4a80760c89179b0b"></a>

<a id="canonical-96a5ab2a440a6ea7a79d447f88df3c96a93462e171ef55da46f0fd01e9ba01ee"></a>

## disable property — Property reference / 823f0b57e8ca / 6

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

<a id="canonical-fb072a01a92f787b8307487ec730a34c2e91c0a06b33be666eef2bd6801f55dc"></a>

<a id="canonical-0318041ff88f9de5f2667a4b794429f949edc1c08711273cc219094a7743c344"></a>

## id property — Property reference / 823f0b57e8ca / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2e0d670fe9380c5b5935081003779185f0b971ec1ca1f759cd99b62a94c4d4c7"></a>

<a id="canonical-8e2dbbf20282df1165625e11fa51243112041651ca97d06ba432f3737eaa4520"></a>

## labels property — Property reference / 823f0b57e8ca / 8

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

- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7): complete subsection reference.

<a id="canonical-d0c7f7bac0d400e7d7a39030241498f6d6f9c04520e715689b1d51a354b26547"></a>

<a id="canonical-8ebee4850443eb5872e128c91c670fe1fe333ca1e51a386931e8401f916855f8"></a>

## name property — Property reference / 823f0b57e8ca / 9

Type: `"string"`. Required.

Name of the Tunnel. Must be unique within the namespace.

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

<a id="canonical-79b412bcd7192407234d21972ddad2eba837c8b82dbb8d1f4417e30d14695887"></a>

<a id="canonical-09e1441d850fdcc56cd38b6dce197ba33151c4c224c08df044225f077eeacdf5"></a>

## namespace property — Property reference / 823f0b57e8ca / 10

Type: `"string"`. Required.

Namespace where the Tunnel is created.

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

- [params](resources--tunnel--reference--group-001.md#canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627): complete subsection reference.

- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4): complete subsection reference.

- [timeouts](resources--tunnel--reference--group-001.md#canonical-2e5aa6734285ac67d12db6f01e1708a97fc9b672be50d2ae079480b57b170b1d): complete subsection reference.

<a id="canonical-0afdda96fe00e97b1a313ab5ed3a77692d0e6082a15af5a2791e37d127f8e0bb"></a>

<a id="canonical-44bcdcfcdaae93e55e612003a893dbeb33f7f19b80f2ea272775f3e5973703ba"></a>

## tunnel_type property — Property reference / 823f0b57e8ca / 11

Type: `"string"`. Optional, Computed.

\[Enum: IPSEC\_PSK|GRE\] Supported tunnel types are IPsec IPsec tunnel type with PSK GRE tunnel
type. Possible values are \`IPSEC\_PSK\`, \`GRE\`. Defaults to \`IPSEC\_PSK\`.

Upstream description:

Supported tunnel types are IPsec

IPsec tunnel type with PSK GRE tunnel type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IPSEC_PSK",
    "GRE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "IPSEC_PSK",
  "enum": [
    "IPSEC_PSK",
    "GRE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1e5e9a77b4b640152713bcaff17aee51cecc799b5c2497b16cc40d64c6450f6d"></a>

## All schema paths — Property reference / 823f0b57e8ca / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--tunnel--reference--group-001.md#canonical-aebd0e8f2620c40ed15d5c018282251c5ad1c2408a5af701bfee8db0fa04e232) |
| `description` | [description](resources--tunnel--reference--group-001.md#canonical-0d9a6bda77c5c5aded0afe0bbda1fad250fd90a1aad5722bf938baf6e0690534) |
| `disable` | [disable](resources--tunnel--reference--group-001.md#canonical-8cd4772cb64a45e1c0fc2c36a56e301bd9efc27830a5ac2e4a80760c89179b0b) |
| `id` | [id](resources--tunnel--reference--group-001.md#canonical-fb072a01a92f787b8307487ec730a34c2e91c0a06b33be666eef2bd6801f55dc) |
| `labels` | [labels](resources--tunnel--reference--group-001.md#canonical-2e0d670fe9380c5b5935081003779185f0b971ec1ca1f759cd99b62a94c4d4c7) |
| `local_ip` | [local_ip](resources--tunnel--reference--group-001.md#canonical-62a074d37a48dcc8ce517c514410d0a245df52366975f988a6d86555ae0aa916) |
| `local_ip.intf` | [local_ip.intf](resources--tunnel--reference--group-001.md#canonical-ae03f906866006f2e8eb705a8b759ba2fcae91a6e7a31b7979d82b040470b719) |
| `local_ip.intf.local_intf` | [local_ip.intf.local_intf](resources--tunnel--reference--group-001.md#canonical-3fcea8c4997a5bf6e4c2843a077168c0429ec93a48dcae89b473b8f5eadbadc5) |
| `local_ip.intf.local_intf.kind` | [local_ip.intf.local_intf.kind](resources--tunnel--reference--group-001.md#canonical-13608d55fe892ede250d6f08cb6801757078cca4fc264775fa5dce5a94df7a5c) |
| `local_ip.intf.local_intf.name` | [local_ip.intf.local_intf.name](resources--tunnel--reference--group-001.md#canonical-903ddcc0bba0353fd68b8b63e5c19016f69b48f162a08bab8c8a478c5165fbb5) |
| `local_ip.intf.local_intf.namespace` | [local_ip.intf.local_intf.namespace](resources--tunnel--reference--group-001.md#canonical-12b82349fcc6d037f02103c56383ccd41d7e419690823754cfdc8a7c75ac647e) |
| `local_ip.intf.local_intf.tenant` | [local_ip.intf.local_intf.tenant](resources--tunnel--reference--group-001.md#canonical-072a16c63e4258c3820939d1238f524d36afc9c299ae1f0ef612e518eae1eadd) |
| `local_ip.intf.local_intf.uid` | [local_ip.intf.local_intf.uid](resources--tunnel--reference--group-001.md#canonical-566694b71ef15fbcef1982b57ebc168bf03ca9a050f3a6340689d8750974d1bf) |
| `local_ip.ip_address` | [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-e5ebf7486d6f5ed7b5803525874bb6f784dbc23889d9368fd829d05f7c9fd0f4) |
| `local_ip.ip_address.auto` | [local_ip.ip_address.auto](resources--tunnel--reference--group-001.md#canonical-378e08de2555749234985460a7db39d98eede44f0f6ad7e799fabe011de6596e) |
| `local_ip.ip_address.ip_address` | [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-528479dfb24e21902e9585cd895def9e4cfd0708aa8b2c056dd6450c1fcd0f05) |
| `local_ip.ip_address.ip_address.dual_stack` | [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-7aa84106e55d709cec6ff37223c2cb312f67a3c02c79627f439305a357cf439d) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4` | [local_ip.ip_address.ip_address.dual_stack.ipv4](resources--tunnel--reference--group-001.md#canonical-fea2ba50a487006b8c4926b3c802e4ffe915b7c51b17fbbac77aae09427a189d) |
| `local_ip.ip_address.ip_address.dual_stack.ipv4.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-477df662073931fb2b664719021e52747ab92647c56005532944672dee097b89) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6` | [local_ip.ip_address.ip_address.dual_stack.ipv6](resources--tunnel--reference--group-001.md#canonical-c360b483cb32f2d5df30ccbaee89088c599b638b6f4304b0ce0cbdee0b3a9859) |
| `local_ip.ip_address.ip_address.dual_stack.ipv6.addr` | [local_ip.ip_address.ip_address.dual_stack.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-d6f384f1592736d0dca8302d2103c3d03969d4dbaa90208b0e4ba2ce2c272a38) |
| `local_ip.ip_address.ip_address.ipv4` | [local_ip.ip_address.ip_address.ipv4](resources--tunnel--reference--group-001.md#canonical-c4a5011e4b991e8fc98ef748cf42644855e47ff2b8ede8db9c439e57b166ce7f) |
| `local_ip.ip_address.ip_address.ipv4.addr` | [local_ip.ip_address.ip_address.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-9481a50bfe3bf62f1c481eda625ee648e5c193ee65e59edd183c70af090dcdeb) |
| `local_ip.ip_address.ip_address.ipv6` | [local_ip.ip_address.ip_address.ipv6](resources--tunnel--reference--group-001.md#canonical-e4ea5931e58b94844ad3f42c912aa0a8b3e35869c9de23c21d0f4fbb8b823c6c) |
| `local_ip.ip_address.ip_address.ipv6.addr` | [local_ip.ip_address.ip_address.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-04e39f4c15c4498e50ef7d3bcc7c332bb98fc05ce239c78b515de428b1dfffe1) |
| `local_ip.ip_address.virtual_network_type` | [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-53458fc58a693c794b471fd47fa2d79aca41a1f63434eb78143553cfd14bc917) |
| `local_ip.ip_address.virtual_network_type.public` | [local_ip.ip_address.virtual_network_type.public](resources--tunnel--reference--group-001.md#canonical-a38be27de2305ba5bc4362bb33e92499f5faba64e6cbdba6a9e2954b9b6c74be) |
| `local_ip.ip_address.virtual_network_type.site_local` | [local_ip.ip_address.virtual_network_type.site_local](resources--tunnel--reference--group-001.md#canonical-2674b42184092c60b6e9ec2aee99f548fe105d9c27ca1397150470cc1093e81b) |
| `local_ip.ip_address.virtual_network_type.site_local_inside` | [local_ip.ip_address.virtual_network_type.site_local_inside](resources--tunnel--reference--group-001.md#canonical-fc358aa868fc2ec41f339272dcc1c45885383e5d1d02e37f419558bd77f257bd) |
| `name` | [name](resources--tunnel--reference--group-001.md#canonical-d0c7f7bac0d400e7d7a39030241498f6d6f9c04520e715689b1d51a354b26547) |
| `namespace` | [namespace](resources--tunnel--reference--group-001.md#canonical-79b412bcd7192407234d21972ddad2eba837c8b82dbb8d1f4417e30d14695887) |
| `params` | [params](resources--tunnel--reference--group-001.md#canonical-70223d998da3ac34fa1833e8da7512e2a44e99fc567baa15fbb48e04ae330d63) |
| `params.ipsec` | [params.ipsec](resources--tunnel--reference--group-001.md#canonical-a29bc07c75dffffd7d136c3a7d7779cf0431a3414e29ba3015c26f64326e773e) |
| `params.ipsec.ipsec_psk` | [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-2552a19fa5613597f7ad3ff241bd8d09f7d37f7b576a7d6553757a299df28de4) |
| `params.ipsec.ipsec_psk.blindfold_secret_info` | [params.ipsec.ipsec_psk.blindfold_secret_info](resources--tunnel--reference--group-001.md#canonical-c235204f06ab815e9e95fc2cef8585af326f7458fe0b4de3b4cff15de7621ce1) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.decryption_provider](resources--tunnel--reference--group-001.md#canonical-9639e8e7a6990198a54242c4febdf83e6557e1e8329e2ed592c717d4b1e57902) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.location` | [params.ipsec.ipsec_psk.blindfold_secret_info.location](resources--tunnel--reference--group-001.md#canonical-23452429c51cdcc2de05e6a2a487179938f2226e3da8720b471a345fd11fe1c8) |
| `params.ipsec.ipsec_psk.blindfold_secret_info.store_provider` | [params.ipsec.ipsec_psk.blindfold_secret_info.store_provider](resources--tunnel--reference--group-001.md#canonical-c34c1c9c1611a5429d9b561e07e47057edde45ecc56e6243966a8c665f15e5c6) |
| `params.ipsec.ipsec_psk.clear_secret_info` | [params.ipsec.ipsec_psk.clear_secret_info](resources--tunnel--reference--group-001.md#canonical-dcf8687659f6fb928dd110fda75563a4df0d7d5d85f387101b9c3159719723e3) |
| `params.ipsec.ipsec_psk.clear_secret_info.provider_ref` | [params.ipsec.ipsec_psk.clear_secret_info.provider_ref](resources--tunnel--reference--group-001.md#canonical-53758a039d65ec2e7aadb609f1bdba5f39c1f6f01222690c23cbc7da65032362) |
| `params.ipsec.ipsec_psk.clear_secret_info.url` | [params.ipsec.ipsec_psk.clear_secret_info.url](resources--tunnel--reference--group-001.md#canonical-00f501286e69887180d0b092ab69150047804eba989be97c4ee2fa92097b0309) |
| `remote_ip` | [remote_ip](resources--tunnel--reference--group-001.md#canonical-52464c7cf311ab703b97340dcc0eda678a061b6b108b6cd299a834de2d4ca866) |
| `remote_ip.endpoints` | [remote_ip.endpoints](resources--tunnel--reference--group-001.md#canonical-57de4f6dc7314f2da1ec195b1d77cc5b11f03581d2d8f09f4babd9bb2b796a30) |
| `remote_ip.endpoints.endpoints` | [remote_ip.endpoints.endpoints](resources--tunnel--reference--group-001.md#canonical-28f1143e34229b5b8e1898bf0b6225ee158e12a63aeeb4aeef5f48a15f10b99a) |
| `remote_ip.ip` | [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-eaa9415b3d8bbe25ab92b960a2f5aa434bddc77a24885aba3aff8b013d0c5764) |
| `remote_ip.ip.dual_stack` | [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-f14dfe7fe3ccce509017dde60a5e14edcb6a7beed5d68699792b5756df8b5682) |
| `remote_ip.ip.dual_stack.ipv4` | [remote_ip.ip.dual_stack.ipv4](resources--tunnel--reference--group-001.md#canonical-297b42522fb9fe19340f5f1b685d508d0a1141f24ea2003c9d3be87aa76910d5) |
| `remote_ip.ip.dual_stack.ipv4.addr` | [remote_ip.ip.dual_stack.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-35cb669f029de2d39dcb9f29072dbb17dc71dbb05b10eb2cf4ecff555e7419de) |
| `remote_ip.ip.dual_stack.ipv6` | [remote_ip.ip.dual_stack.ipv6](resources--tunnel--reference--group-001.md#canonical-2ceeefbcf717de775dcea323bf6da17258e86c2d148033304c03768ea8e79c73) |
| `remote_ip.ip.dual_stack.ipv6.addr` | [remote_ip.ip.dual_stack.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-c4bfb16ab1607fbd1c37805718c2c5d3cceb2294fdb7f454d580aa9d15f063cd) |
| `remote_ip.ip.ipv4` | [remote_ip.ip.ipv4](resources--tunnel--reference--group-001.md#canonical-8aaf9a2b179bcf2c84fc0af75b76b3b6735b50146f432abe686c0d87bcf10c8e) |
| `remote_ip.ip.ipv4.addr` | [remote_ip.ip.ipv4.addr](resources--tunnel--reference--group-001.md#canonical-b2cddcf832c0c3cd4d1c682acdd590e95cf4e08883b5356cf92e6a9274ad2693) |
| `remote_ip.ip.ipv6` | [remote_ip.ip.ipv6](resources--tunnel--reference--group-001.md#canonical-767dacb7cb501eae25a81b02d18ed255d4225a3b274a9b12e6d6f3cf05f800b7) |
| `remote_ip.ip.ipv6.addr` | [remote_ip.ip.ipv6.addr](resources--tunnel--reference--group-001.md#canonical-487b762a630d1762b05f453ac9b221143995e19eea8a59d10b6b66f75aec673d) |
| `timeouts` | [timeouts](resources--tunnel--reference--group-001.md#canonical-a31566559244c74d2e99aa324fdb09cf4a20f2b5152df02aa7c23c4c21d99b80) |
| `timeouts.create` | [timeouts.create](resources--tunnel--reference--group-001.md#canonical-55382f6fc55221fd90c9887bfaff3fbc07f945e9b5a85bd0be25b04acad041a0) |
| `timeouts.delete` | [timeouts.delete](resources--tunnel--reference--group-001.md#canonical-c3c2d64e44fe4cc870b354df783ff8f0aad2052a8986e52b9ea1d74575ef7c27) |
| `timeouts.read` | [timeouts.read](resources--tunnel--reference--group-001.md#canonical-da07924ff985d9d277cc48be4ee095004c84918c8caa58cda15f686f759d3cf8) |
| `timeouts.update` | [timeouts.update](resources--tunnel--reference--group-001.md#canonical-4c66580047f03ec051450f16b351600d1a1deb6bf799ddeb06fe63e11ae62d53) |
| `tunnel_type` | [tunnel_type](resources--tunnel--reference--group-001.md#canonical-0afdda96fe00e97b1a313ab5ed3a77692d0e6082a15af5a2791e37d127f8e0bb) |

<a id="canonical-6c816b5daa9aa3fcd9f0594db4cf57a90c6ce44c283b721973223511b59be208"></a>

## Next pages — Property reference / 823f0b57e8ca / 13

- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [params](resources--tunnel--reference--group-001.md#canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [timeouts](resources--tunnel--reference--group-001.md#canonical-2e5aa6734285ac67d12db6f01e1708a97fc9b672be50d2ae079480b57b170b1d)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a6327eba447998948a6ded85f12e08d1427cda7cc6b6dce5c00df751f342342"></a>

## local_ip — local_ip / 76c8983e98e3 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- local_ip

<a id="canonical-62a074d37a48dcc8ce517c514410d0a245df52366975f988a6d86555ae0aa916"></a>

Type: `"object"`. single nested block, Optional.

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - 1. Local Interface - Network Interface from which IP address and network will be
selected 2. IP Address - IP address and network can be configured explicitly.

Upstream description:

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - &#8203;1. Local Interface - Network Interface from which IP address and network will
be selected &#8203;2. IP Address - IP address and network can be configured explicitly.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("intf",
    "ip_address")}
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
  "x-ves-oneof-field-type": "[\"intf\",\"ip_address\"]"
}
```

Terraform syntax:

```terraform
local_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-52c7275665c0720710cb7936eeeb3003792ff95d201521019c0e64d42e794426"></a>

## Direct properties — local_ip / 76c8983e98e3 / 3

- [intf](resources--tunnel--reference--group-001.md#canonical-abb32c111a659793b9202a4ff80effdbf7ad1fd58016caac305e60784f38d027): complete subsection reference.

- [ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4): complete subsection reference.

<a id="canonical-60ccadbcf5bd7e28a24ae5a2d86e1c63e5290ca90f93cdd75ddcd75703d7fec8"></a>

## Next pages — local_ip / 76c8983e98e3 / 4

- [local_ip.intf](resources--tunnel--reference--group-001.md#canonical-abb32c111a659793b9202a4ff80effdbf7ad1fd58016caac305e60784f38d027)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-abb32c111a659793b9202a4ff80effdbf7ad1fd58016caac305e60784f38d027"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d472516e28f3c29856775732ac850a8cd0c59d731d7e5496611ddcf67ab5b701"></a>

## local_ip.intf — local_ip.intf / 78ab24eaaf12 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- local_ip.intf

<a id="canonical-ae03f906866006f2e8eb705a8b759ba2fcae91a6e7a31b7979d82b040470b719"></a>

Type: `"object"`. single nested block, Optional.

Provides the local interface to pick up source IP and network for transporting encapsulated packet.

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
intf {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8a32c80313eaac75a4cc160aeabbb1a7378a94ea24deeaa6bca9abc84da55e1"></a>

## Direct properties — local_ip.intf / 78ab24eaaf12 / 3

- [local_intf](resources--tunnel--reference--group-001.md#canonical-0fbba6a84148990fd7da24e95eb64f2ccdc8a5b43abbcac4012946c6ee57673c): complete subsection reference.

<a id="canonical-5ca2369dd30a0e72580a067a9f4d4f226f14910fad5f94439f05012dde34d1e2"></a>

## Next pages — local_ip.intf / 78ab24eaaf12 / 4

- [local_ip.intf.local_intf](resources--tunnel--reference--group-001.md#canonical-0fbba6a84148990fd7da24e95eb64f2ccdc8a5b43abbcac4012946c6ee57673c)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-0fbba6a84148990fd7da24e95eb64f2ccdc8a5b43abbcac4012946c6ee57673c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22e14bf5f7272fef4dd5c842d448d7593bcef4a3d7d5c67273d47d62b4fcb4d2"></a>

## local_ip.intf.local_intf — local_ip.intf.local_intf / fbeffc659707 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.intf](resources--tunnel--reference--group-001.md#canonical-abb32c111a659793b9202a4ff80effdbf7ad1fd58016caac305e60784f38d027)
- local_ip.intf.local_intf

<a id="canonical-3fcea8c4997a5bf6e4c2843a077168c0429ec93a48dcae89b473b8f5eadbadc5"></a>

Type: `"object"`. list nested block, Optional.

Local interface to be used for filling in source information of IP and network for transport.

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
local_intf {
  # Configure direct properties listed below.
}
```

<a id="canonical-c6f97ea899960dfe4398cfd5c4a22bed728a77b2614e8bb338d11a86c440e46a"></a>

## Direct properties — local_ip.intf.local_intf / fbeffc659707 / 3

<a id="canonical-13608d55fe892ede250d6f08cb6801757078cca4fc264775fa5dce5a94df7a5c"></a>

<a id="canonical-28a25aee15d8fcff529dc2332851c4893b0ed5b5f3a418cb3585bc6ba4c221fe"></a>

## kind property — local_ip.intf.local_intf / fbeffc659707 / 4

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

<a id="canonical-903ddcc0bba0353fd68b8b63e5c19016f69b48f162a08bab8c8a478c5165fbb5"></a>

<a id="canonical-57cce3f7ff7cd32dab898868100581b8b4a70c440769179d3d492a3eaebabd6c"></a>

## name property — local_ip.intf.local_intf / fbeffc659707 / 5

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

<a id="canonical-12b82349fcc6d037f02103c56383ccd41d7e419690823754cfdc8a7c75ac647e"></a>

<a id="canonical-09ec831d3bbdc2cc07849140c05fcf11fdc69586d02098f16a50352656d360b4"></a>

## namespace property — local_ip.intf.local_intf / fbeffc659707 / 6

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

<a id="canonical-072a16c63e4258c3820939d1238f524d36afc9c299ae1f0ef612e518eae1eadd"></a>

<a id="canonical-ef2456f7fd4ed7f8a35db6b4a8f78c94b0ed33c2f8fe1f66ede98f59a2f480cf"></a>

## tenant property — local_ip.intf.local_intf / fbeffc659707 / 7

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

<a id="canonical-566694b71ef15fbcef1982b57ebc168bf03ca9a050f3a6340689d8750974d1bf"></a>

<a id="canonical-03de08ae5d9b9c61daebc7b8be3612d8a12600338779485ed8bb1a815234cbb2"></a>

## uid property — local_ip.intf.local_intf / fbeffc659707 / 8

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

<a id="canonical-f1a54ff849dd12679daf0ccd7c485b5751e7a3833e7a9aa9fa6706b50148d5ec"></a>

## Next pages — local_ip.intf.local_intf / fbeffc659707 / 9

- [local_ip.intf](resources--tunnel--reference--group-001.md#canonical-abb32c111a659793b9202a4ff80effdbf7ad1fd58016caac305e60784f38d027)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1290aff6d7d15254a372f84f34adefe9e7eb91f646e8c0351c78ae33b12ab7a"></a>

## local_ip.ip_address — local_ip.ip_address / d725ff031d3a / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- local_ip.ip_address

<a id="canonical-e5ebf7486d6f5ed7b5803525874bb6f784dbc23889d9368fd829d05f7c9fd0f4"></a>

Type: `"object"`. single nested block, Optional.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "ip_address")}
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
  "x-ves-oneof-field-type": "[\"auto\",\"ip_address\"]"
}
```

Terraform syntax:

```terraform
ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-47474a5036304e9f6c31bf4faa503bca275fed1564b87dbcb8f1a4e54daafd13"></a>

## Direct properties — local_ip.ip_address / d725ff031d3a / 3

- [auto](resources--tunnel--reference--group-001.md#canonical-b45c36129b38b58a4f25c7e4e23b39df38f25e8f1944da8708b70fdc4ee0fa59): complete subsection reference.

- [ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a): complete subsection reference.

- [virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93): complete subsection reference.

<a id="canonical-096bc6a3e9c137d69e13dc1c45ca3500232a01b468bc3fad84fd584ec75b72e0"></a>

## Next pages — local_ip.ip_address / d725ff031d3a / 4

- [local_ip.ip_address.auto](resources--tunnel--reference--group-001.md#canonical-b45c36129b38b58a4f25c7e4e23b39df38f25e8f1944da8708b70fdc4ee0fa59)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-b45c36129b38b58a4f25c7e4e23b39df38f25e8f1944da8708b70fdc4ee0fa59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff4f703199f1f081bc45337aed2a73918245241822e044c31e640e9fb1f56c28"></a>

## local_ip.ip_address.auto — local_ip.ip_address.auto / 60e418633db0 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- local_ip.ip_address.auto

<a id="canonical-378e08de2555749234985460a7db39d98eede44f0f6ad7e799fabe011de6596e"></a>

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
auto = {}
```

<a id="canonical-581eb2174be1df4c77bc09975822502b9f3f3f738c2a20bfe850129ceb5c8a40"></a>

## Direct properties — local_ip.ip_address.auto / 60e418633db0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac6fa6f975ab79b1b847fdc906c0a84bd8515e8436e025c2b398f531baebc533"></a>

## Next pages — local_ip.ip_address.auto / 60e418633db0 / 4

- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9646e52683731f92e60ef462c031bafc6db6c8e39bf58bf1ef9c2deee67a6ab9"></a>

## local_ip.ip_address.ip_address — local_ip.ip_address.ip_address / c24ca3c44d14 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- local_ip.ip_address.ip_address

<a id="canonical-528479dfb24e21902e9585cd895def9e4cfd0708aa8b2c056dd6450c1fcd0f05"></a>

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
ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-cae6126e6a93dd5b7bc26ac4c6b3898471d21a16514424f81716922d1f288115"></a>

## Direct properties — local_ip.ip_address.ip_address / c24ca3c44d14 / 3

- [dual_stack](resources--tunnel--reference--group-001.md#canonical-c79b259e4c7ed28cea185542563d65facceafd9076da194d48d25463792a35cc): complete subsection reference.

- [ipv4](resources--tunnel--reference--group-001.md#canonical-c2e921008069084947a0a167536a59ede9dcf7bce118c59a1600e77e8c49b06a): complete subsection reference.

- [ipv6](resources--tunnel--reference--group-001.md#canonical-a48d35fb9eddad8b65fe11a73ece1e18e32490ab1d120a45b84fc66d38213563): complete subsection reference.

<a id="canonical-6eacfce4c9312318f43a249745503f1e35dc52c0ef396a22195bc47678cf3350"></a>

## Next pages — local_ip.ip_address.ip_address / c24ca3c44d14 / 4

- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-c79b259e4c7ed28cea185542563d65facceafd9076da194d48d25463792a35cc)
- [local_ip.ip_address.ip_address.ipv4](resources--tunnel--reference--group-001.md#canonical-c2e921008069084947a0a167536a59ede9dcf7bce118c59a1600e77e8c49b06a)
- [local_ip.ip_address.ip_address.ipv6](resources--tunnel--reference--group-001.md#canonical-a48d35fb9eddad8b65fe11a73ece1e18e32490ab1d120a45b84fc66d38213563)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-c79b259e4c7ed28cea185542563d65facceafd9076da194d48d25463792a35cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9875ce06b8ee7195f00afbd8c60c39138ac90346eefedec74ade0cc20c6c4248"></a>

## local_ip.ip_address.ip_address.dual_stack — local_ip.ip_address.ip_address.dual_stack / e11104783192 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- local_ip.ip_address.ip_address.dual_stack

<a id="canonical-7aa84106e55d709cec6ff37223c2cb312f67a3c02c79627f439305a357cf439d"></a>

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

<a id="canonical-024ebc2ccd10c6c3f3473bb1743a647d52a73d5414f319580cf244f4ab157bf9"></a>

## Direct properties — local_ip.ip_address.ip_address.dual_stack / e11104783192 / 3

- [ipv4](resources--tunnel--reference--group-001.md#canonical-673531e4d902fe2249bb31b16f355b47da5479d59c990d82ef4d3885c11bc7b6): complete subsection reference.

- [ipv6](resources--tunnel--reference--group-001.md#canonical-3f03e363bd70713726c88418533e7cfb3f41bbc29311de91760ca53beaf3d2b3): complete subsection reference.

<a id="canonical-5d110c338537545df8c7b1ec413be4aa6610f5b39a86d74f6f4c9849c16cae2b"></a>

## Next pages — local_ip.ip_address.ip_address.dual_stack / e11104783192 / 4

- [local_ip.ip_address.ip_address.dual_stack.ipv4](resources--tunnel--reference--group-001.md#canonical-673531e4d902fe2249bb31b16f355b47da5479d59c990d82ef4d3885c11bc7b6)
- [local_ip.ip_address.ip_address.dual_stack.ipv6](resources--tunnel--reference--group-001.md#canonical-3f03e363bd70713726c88418533e7cfb3f41bbc29311de91760ca53beaf3d2b3)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-673531e4d902fe2249bb31b16f355b47da5479d59c990d82ef4d3885c11bc7b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39def8158a9dff3c277ca2d2bbd896140a053a217291e92ba5cd5d61ac06c0a4"></a>

## local_ip.ip_address.ip_address.dual_stack.ipv4 — local_ip.ip_address.ip_address.dual_stack.ipv4 / d5e403627f9d / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-c79b259e4c7ed28cea185542563d65facceafd9076da194d48d25463792a35cc)
- local_ip.ip_address.ip_address.dual_stack.ipv4

<a id="canonical-fea2ba50a487006b8c4926b3c802e4ffe915b7c51b17fbbac77aae09427a189d"></a>

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

<a id="canonical-a0ff2c87ec595ee0cb95784acd6fd3443ec1de682ada73c1cf859717fd4f06cb"></a>

## Direct properties — local_ip.ip_address.ip_address.dual_stack.ipv4 / d5e403627f9d / 3

<a id="canonical-477df662073931fb2b664719021e52747ab92647c56005532944672dee097b89"></a>

<a id="canonical-c26763714051e181be756cf77d3e4e82796052cae2570d74897779cdf86c1812"></a>

## addr property — local_ip.ip_address.ip_address.dual_stack.ipv4 / d5e403627f9d / 4

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

<a id="canonical-11f4152a9f78bea829ee27f80ad8834ddac918af6d1b53b0e558b9140c7b04ff"></a>

## Next pages — local_ip.ip_address.ip_address.dual_stack.ipv4 / d5e403627f9d / 5

- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-c79b259e4c7ed28cea185542563d65facceafd9076da194d48d25463792a35cc)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-3f03e363bd70713726c88418533e7cfb3f41bbc29311de91760ca53beaf3d2b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7291b4af6730a5bce13cb1281af4e01ebf2a48648c18cb1a0f14c27bec0a0eb4"></a>

## local_ip.ip_address.ip_address.dual_stack.ipv6 — local_ip.ip_address.ip_address.dual_stack.ipv6 / d91aa88a15c7 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-c79b259e4c7ed28cea185542563d65facceafd9076da194d48d25463792a35cc)
- local_ip.ip_address.ip_address.dual_stack.ipv6

<a id="canonical-c360b483cb32f2d5df30ccbaee89088c599b638b6f4304b0ce0cbdee0b3a9859"></a>

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

<a id="canonical-74c9a2d0719a30654bd3821e92dbb23f52a859367e41e17cd92b69ba0bce1a79"></a>

## Direct properties — local_ip.ip_address.ip_address.dual_stack.ipv6 / d91aa88a15c7 / 3

<a id="canonical-d6f384f1592736d0dca8302d2103c3d03969d4dbaa90208b0e4ba2ce2c272a38"></a>

<a id="canonical-8befce31b9eccd5c0bf23075a7b209528afc8b05b5b8ac864388e805f1f0158f"></a>

## addr property — local_ip.ip_address.ip_address.dual_stack.ipv6 / d91aa88a15c7 / 4

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

<a id="canonical-16996db656f8014052ab65806b4e3f7f0b0664f0944615cb13ec0425ded80fd7"></a>

## Next pages — local_ip.ip_address.ip_address.dual_stack.ipv6 / d91aa88a15c7 / 5

- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--reference--group-001.md#canonical-c79b259e4c7ed28cea185542563d65facceafd9076da194d48d25463792a35cc)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-c2e921008069084947a0a167536a59ede9dcf7bce118c59a1600e77e8c49b06a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f22ca4bb32c1351a9436383f76e5e8ddba69d4739ba75d68fefb769184eaa119"></a>

## local_ip.ip_address.ip_address.ipv4 — local_ip.ip_address.ip_address.ipv4 / cbb7659d6992 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- local_ip.ip_address.ip_address.ipv4

<a id="canonical-c4a5011e4b991e8fc98ef748cf42644855e47ff2b8ede8db9c439e57b166ce7f"></a>

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

<a id="canonical-8091021fe94e9b78a10204a6f0c3308f22a9939a4981e2ae49262145d78435d6"></a>

## Direct properties — local_ip.ip_address.ip_address.ipv4 / cbb7659d6992 / 3

<a id="canonical-9481a50bfe3bf62f1c481eda625ee648e5c193ee65e59edd183c70af090dcdeb"></a>

<a id="canonical-0eb0db5b694d3530b3fe9628c3a7b0d81d47ecf75c880d43f8e507b476697a60"></a>

## addr property — local_ip.ip_address.ip_address.ipv4 / cbb7659d6992 / 4

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

<a id="canonical-2875e7538c5f3e841f3309acd212770f7c22e50836eeefdce2189802487be21d"></a>

## Next pages — local_ip.ip_address.ip_address.ipv4 / cbb7659d6992 / 5

- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-a48d35fb9eddad8b65fe11a73ece1e18e32490ab1d120a45b84fc66d38213563"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ae90f1567d6fe6044341b08d1de6451197e57d6a50fc258513891b601e729df"></a>

## local_ip.ip_address.ip_address.ipv6 — local_ip.ip_address.ip_address.ipv6 / cf590e80b0b0 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- local_ip.ip_address.ip_address.ipv6

<a id="canonical-e4ea5931e58b94844ad3f42c912aa0a8b3e35869c9de23c21d0f4fbb8b823c6c"></a>

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

<a id="canonical-1426fcd0ac7ca480f93f7132ed57731cd537c19cf65fbbfd77b1216a468c1edf"></a>

## Direct properties — local_ip.ip_address.ip_address.ipv6 / cf590e80b0b0 / 3

<a id="canonical-04e39f4c15c4498e50ef7d3bcc7c332bb98fc05ce239c78b515de428b1dfffe1"></a>

<a id="canonical-f9422f8f5595a764d0cb9d4bc962eb8407311a28c43451ee30b66790766abde2"></a>

## addr property — local_ip.ip_address.ip_address.ipv6 / cf590e80b0b0 / 4

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

<a id="canonical-892e7267a562e4a6cd53286950ae11d7b7a3263e6fcfaa4a9b24efbb815958eb"></a>

## Next pages — local_ip.ip_address.ip_address.ipv6 / cf590e80b0b0 / 5

- [local_ip.ip_address.ip_address](resources--tunnel--reference--group-001.md#canonical-1be34312c00be427cb8e4f1c25a13ea39fb77277b88e1df22c6370682686bb0a)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ced324432706a62433b61258748c3844fab3067bf6528a15f7ecdf0041b6e41"></a>

## local_ip.ip_address.virtual_network_type — local_ip.ip_address.virtual_network_type / 59e6656b78f3 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- local_ip.ip_address.virtual_network_type

<a id="canonical-53458fc58a693c794b471fd47fa2d79aca41a1f63434eb78143553cfd14bc917"></a>

Type: `"object"`. single nested block, Optional.

Different types of virtual networks understood by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("public",
    "site_local"),
  validators.ConflictingObjectAttributes("public",
    "site_local_inside"),
  validators.ConflictingObjectAttributes("site_local",
    "site_local_inside")}
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
  "x-ves-oneof-field-vn_type_choice": "[\"public\",\"site_local\",\"site_local_inside\"]"
}
```

Terraform syntax:

```terraform
virtual_network_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-5cf196cb4d521c25e1324e97c23277a4b1b0f1934faf26f98b9b1edc5fc68d6d"></a>

## Direct properties — local_ip.ip_address.virtual_network_type / 59e6656b78f3 / 3

- [public](resources--tunnel--reference--group-001.md#canonical-d864b3d0eab0e62949d38ea43631b5314561671c8634ab8a1797b314e67bebef): complete subsection reference.

- [site_local](resources--tunnel--reference--group-001.md#canonical-f58256109882878c33eca392a433f802f0910db42b49f2dde852ac2c61890d14): complete subsection reference.

- [site_local_inside](resources--tunnel--reference--group-001.md#canonical-77f4cf69a8ffe4f3ef5f0e453beeff15deee067aa5105306fff6bac9463aa619): complete subsection reference.

<a id="canonical-1dd7a8a02b4de3ed9d864ccad4de68b77d4514c45d9b287661180d0cdcecc6cf"></a>

## Next pages — local_ip.ip_address.virtual_network_type / 59e6656b78f3 / 4

- [local_ip.ip_address.virtual_network_type.public](resources--tunnel--reference--group-001.md#canonical-d864b3d0eab0e62949d38ea43631b5314561671c8634ab8a1797b314e67bebef)
- [local_ip.ip_address.virtual_network_type.site_local](resources--tunnel--reference--group-001.md#canonical-f58256109882878c33eca392a433f802f0910db42b49f2dde852ac2c61890d14)
- [local_ip.ip_address.virtual_network_type.site_local_inside](resources--tunnel--reference--group-001.md#canonical-77f4cf69a8ffe4f3ef5f0e453beeff15deee067aa5105306fff6bac9463aa619)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-d864b3d0eab0e62949d38ea43631b5314561671c8634ab8a1797b314e67bebef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcb3fc1b756f82a9d12b9ed94d935f1f9672cae784b410f9123a9271c06663fc"></a>

## local_ip.ip_address.virtual_network_type.public — local_ip.ip_address.virtual_network_type.public / 0d9daf2a982d / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93)
- local_ip.ip_address.virtual_network_type.public

<a id="canonical-a38be27de2305ba5bc4362bb33e92499f5faba64e6cbdba6a9e2954b9b6c74be"></a>

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
public = {}
```

<a id="canonical-776bd381b546462ac7b7a5f0c4aafed3e7d9c98d53a028422538866afb6822bb"></a>

## Direct properties — local_ip.ip_address.virtual_network_type.public / 0d9daf2a982d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0a2b38526844b609f0803cf41125bf2c822163cb115f4a2baf226b19a4a1a08e"></a>

## Next pages — local_ip.ip_address.virtual_network_type.public / 0d9daf2a982d / 4

- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-f58256109882878c33eca392a433f802f0910db42b49f2dde852ac2c61890d14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f350e16202ae4033cadea6d724e7162f2700cd6ba09cdbb259ddcaa480bad44d"></a>

## local_ip.ip_address.virtual_network_type.site_local — local_ip.ip_address.virtual_network_type.site_local / e59eb1ef5203 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93)
- local_ip.ip_address.virtual_network_type.site_local

<a id="canonical-2674b42184092c60b6e9ec2aee99f548fe105d9c27ca1397150470cc1093e81b"></a>

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
site_local = {}
```

<a id="canonical-2e3b1f58468830d0c0662e420b521ca97095713b7fe29b7c20bedc423d8a1280"></a>

## Direct properties — local_ip.ip_address.virtual_network_type.site_local / e59eb1ef5203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87e54fe92d0559d36c301e7b4b6439b49d708fa76320b475f4d483e58d4554fc"></a>

## Next pages — local_ip.ip_address.virtual_network_type.site_local / e59eb1ef5203 / 4

- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-77f4cf69a8ffe4f3ef5f0e453beeff15deee067aa5105306fff6bac9463aa619"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43c5d6af51b7e4db7d1dfece08fadedf5a79b5deb12dda2d16137855d2a27419"></a>

## local_ip.ip_address.virtual_network_type.site_local_inside — local_ip.ip_address.virtual_network_type.site_local_inside / b2c75a65fd4b / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [local_ip](resources--tunnel--reference--group-001.md#canonical-f4098d35ef5afc7ae20f337ae052ab6d54942e628bb1f9a9c820c6e3f61e01c7)
- [local_ip.ip_address](resources--tunnel--reference--group-001.md#canonical-11a1d28a4093af920933ef6c6baa0023d66192697f528276f8cc0dd4ce7369b4)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93)
- local_ip.ip_address.virtual_network_type.site_local_inside

<a id="canonical-fc358aa868fc2ec41f339272dcc1c45885383e5d1d02e37f419558bd77f257bd"></a>

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
site_local_inside = {}
```

<a id="canonical-11748c49415603fb31b60311eae40ea1d61655771a19324d1db86408315ca12e"></a>

## Direct properties — local_ip.ip_address.virtual_network_type.site_local_inside / b2c75a65fd4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9a4bb4a9cdd19a53caf49e9650e0a376abfe6687eefe14125e77d41413fdf6b"></a>

## Next pages — local_ip.ip_address.virtual_network_type.site_local_inside / b2c75a65fd4b / 4

- [local_ip.ip_address.virtual_network_type](resources--tunnel--reference--group-001.md#canonical-1188b399f949ab32d2705baa63e6a989c7fb8497fefef525e7e5e421ff9f4a93)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f5b137c62b1086895a095033015bc46fe1f501ee62c28c17e1131ab423a1ef2"></a>

## params — params / 06a6de775890 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- params

<a id="canonical-70223d998da3ac34fa1833e8da7512e2a44e99fc567baa15fbb48e04ae330d63"></a>

Type: `"object"`. single nested block, Optional.

Tunnel configuration parameters for supported encapsulation 1. IPsec is supported with PSK for which
PSK can be configured.

Upstream description:

Tunnel configuration parameters for supported encapsulation &#8203;1. IPsec is supported with PSK
for which PSK can be configured.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"ipsec\"]"
}
```

Terraform syntax:

```terraform
params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2fe981589bd4190bfad9b72d699340638e73e0e09bff67ca0614914d04d2bf71"></a>

## Direct properties — params / 06a6de775890 / 3

- [ipsec](resources--tunnel--reference--group-001.md#canonical-7fa261ac550da567a1ccd4b3e61c190a3f9829dfcc0c642cc0d3fc51458dbb1d): complete subsection reference.

<a id="canonical-897e54246040d9717e7de0d08810c894419a5aa9c9c61314f2d985db5ff11f90"></a>

## Next pages — params / 06a6de775890 / 4

- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-7fa261ac550da567a1ccd4b3e61c190a3f9829dfcc0c642cc0d3fc51458dbb1d)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-7fa261ac550da567a1ccd4b3e61c190a3f9829dfcc0c642cc0d3fc51458dbb1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09d90d58ab797f269c6676079385214f028bf884b8174cbb097461bc305ca1b7"></a>

## params.ipsec — params.ipsec / b554f54b6d72 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [params](resources--tunnel--reference--group-001.md#canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627)
- params.ipsec

<a id="canonical-a29bc07c75dffffd7d136c3a7d7779cf0431a3414e29ba3015c26f64326e773e"></a>

Type: `"object"`. single nested block, Optional.

Configuration for IPsec encapsulation are: 1. PSK - pre shared key to be used by IKE.

Upstream description:

Configuration for IPsec encapsulation are: &#8203;1. PSK - pre shared key to be used by IKE.

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
ipsec {
  # Configure direct properties listed below.
}
```

<a id="canonical-e50c7c726b2979709a7479f91d03aef595c9a5965951981f98e3c69790ae02a8"></a>

## Direct properties — params.ipsec / b554f54b6d72 / 3

- [ipsec_psk](resources--tunnel--reference--group-001.md#canonical-1841915d2a9af6b30fd3b48c156a82d906ea3a1ee71b0e4230201e4580daced4): complete subsection reference.

<a id="canonical-34cd03d81794556362e438225eefb98c723b0a5d9e19d8a297dd7edb43e41bf7"></a>

## Next pages — params.ipsec / b554f54b6d72 / 4

- [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-1841915d2a9af6b30fd3b48c156a82d906ea3a1ee71b0e4230201e4580daced4)
- [params](resources--tunnel--reference--group-001.md#canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-1841915d2a9af6b30fd3b48c156a82d906ea3a1ee71b0e4230201e4580daced4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6281defb8b222b658c2c7f3552b465e1a4f6a94b502ff329d4e3ed797c0bf75"></a>

## params.ipsec.ipsec_psk — params.ipsec.ipsec_psk / 145407a349ae / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [params](resources--tunnel--reference--group-001.md#canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627)
- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-7fa261ac550da567a1ccd4b3e61c190a3f9829dfcc0c642cc0d3fc51458dbb1d)
- params.ipsec.ipsec_psk

<a id="canonical-2552a19fa5613597f7ad3ff241bd8d09f7d37f7b576a7d6553757a299df28de4"></a>

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
ipsec_psk {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef2ebb1041f23705eead1ee73cd80702cb673e80c6cc33bd8969ff58306ae25e"></a>

## Direct properties — params.ipsec.ipsec_psk / 145407a349ae / 3

- [blindfold_secret_info](resources--tunnel--reference--group-001.md#canonical-9595cbb5dd102c44425d0a0c8cee1fa9c56834ee50de854ee8c71cddb3b524ad): complete subsection reference.

- [clear_secret_info](resources--tunnel--reference--group-001.md#canonical-7d9f06a2d185c38a2a9a8c4d4c5f1a9271b831c5005c96e80773da39fddd0791): complete subsection reference.

<a id="canonical-3ba682b5fc9eeeca1de2e5650d46e5688352c697aeadad76c8dfd495d9e17f7f"></a>

## Next pages — params.ipsec.ipsec_psk / 145407a349ae / 4

- [params.ipsec.ipsec_psk.blindfold_secret_info](resources--tunnel--reference--group-001.md#canonical-9595cbb5dd102c44425d0a0c8cee1fa9c56834ee50de854ee8c71cddb3b524ad)
- [params.ipsec.ipsec_psk.clear_secret_info](resources--tunnel--reference--group-001.md#canonical-7d9f06a2d185c38a2a9a8c4d4c5f1a9271b831c5005c96e80773da39fddd0791)
- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-7fa261ac550da567a1ccd4b3e61c190a3f9829dfcc0c642cc0d3fc51458dbb1d)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-9595cbb5dd102c44425d0a0c8cee1fa9c56834ee50de854ee8c71cddb3b524ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c17fc54ccb5a4889c82eac74bbcea9184b292fc26b4ef4f7364c4fd5f03167b"></a>

## params.ipsec.ipsec_psk.blindfold_secret_info — params.ipsec.ipsec_psk.blindfold_secret_info / bb8b2cba0afd / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [params](resources--tunnel--reference--group-001.md#canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627)
- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-7fa261ac550da567a1ccd4b3e61c190a3f9829dfcc0c642cc0d3fc51458dbb1d)
- [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-1841915d2a9af6b30fd3b48c156a82d906ea3a1ee71b0e4230201e4580daced4)
- params.ipsec.ipsec_psk.blindfold_secret_info

<a id="canonical-c235204f06ab815e9e95fc2cef8585af326f7458fe0b4de3b4cff15de7621ce1"></a>

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

<a id="canonical-7f6ca362f7c7ae60ba39519e893eb24ec56d766a6a897b0431fbff554ecbdad3"></a>

## Direct properties — params.ipsec.ipsec_psk.blindfold_secret_info / bb8b2cba0afd / 3

<a id="canonical-9639e8e7a6990198a54242c4febdf83e6557e1e8329e2ed592c717d4b1e57902"></a>

<a id="canonical-abc1ff48720aa0c710e582b7e8181dd97dbae0effef874028c7066252f1b8a52"></a>

## decryption_provider property — params.ipsec.ipsec_psk.blindfold_secret_info / bb8b2cba0afd / 4

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

<a id="canonical-23452429c51cdcc2de05e6a2a487179938f2226e3da8720b471a345fd11fe1c8"></a>

<a id="canonical-938fa3bc998a7e48ff9bbc5afe46028ea490d12ac5b0f5ac82be9984ca18994e"></a>

## location property — params.ipsec.ipsec_psk.blindfold_secret_info / bb8b2cba0afd / 5

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

<a id="canonical-c34c1c9c1611a5429d9b561e07e47057edde45ecc56e6243966a8c665f15e5c6"></a>

<a id="canonical-4d318b71172e8304ea72426713ae1d64eafd3487cf9d1432c0e5c4db719d8442"></a>

## store_provider property — params.ipsec.ipsec_psk.blindfold_secret_info / bb8b2cba0afd / 6

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

<a id="canonical-0e07c282a8fe9f0abb98753c989566683abea4f220988319816311e7c63c31a3"></a>

## Next pages — params.ipsec.ipsec_psk.blindfold_secret_info / bb8b2cba0afd / 7

- [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-1841915d2a9af6b30fd3b48c156a82d906ea3a1ee71b0e4230201e4580daced4)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-7d9f06a2d185c38a2a9a8c4d4c5f1a9271b831c5005c96e80773da39fddd0791"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05ffd3a232f9446d03e0b3d50a1d03f24b323e601ac400c79f745c532044aaf1"></a>

## params.ipsec.ipsec_psk.clear_secret_info — params.ipsec.ipsec_psk.clear_secret_info / cb3ea2c3f7ae / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [params](resources--tunnel--reference--group-001.md#canonical-d6a699cf73a8bb2c2ae89a3bbedba6c3a07dbb91c960fe1ce0e8f7f49dd34627)
- [params.ipsec](resources--tunnel--reference--group-001.md#canonical-7fa261ac550da567a1ccd4b3e61c190a3f9829dfcc0c642cc0d3fc51458dbb1d)
- [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-1841915d2a9af6b30fd3b48c156a82d906ea3a1ee71b0e4230201e4580daced4)
- params.ipsec.ipsec_psk.clear_secret_info

<a id="canonical-dcf8687659f6fb928dd110fda75563a4df0d7d5d85f387101b9c3159719723e3"></a>

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

<a id="canonical-903d94035181aa4b273105b3723ab459f0513718eb00d55542b361d99304ad9e"></a>

## Direct properties — params.ipsec.ipsec_psk.clear_secret_info / cb3ea2c3f7ae / 3

<a id="canonical-53758a039d65ec2e7aadb609f1bdba5f39c1f6f01222690c23cbc7da65032362"></a>

<a id="canonical-2f9404de4569ce8bcb4eee95aa56d652b5ea357f4e7ea2f2ae857735e22a9b89"></a>

## provider_ref property — params.ipsec.ipsec_psk.clear_secret_info / cb3ea2c3f7ae / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-00f501286e69887180d0b092ab69150047804eba989be97c4ee2fa92097b0309"></a>

<a id="canonical-b5ca8245dc65b5b055d229b73bbe1fb13d53e1a1950141fe59f655cf3c866783"></a>

## url property — params.ipsec.ipsec_psk.clear_secret_info / cb3ea2c3f7ae / 5

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

<a id="canonical-3a2a13a97dcea83c3016908eab31672b36a076bbc0dbe25cfcf158c833f690c6"></a>

## Next pages — params.ipsec.ipsec_psk.clear_secret_info / cb3ea2c3f7ae / 6

- [params.ipsec.ipsec_psk](resources--tunnel--reference--group-001.md#canonical-1841915d2a9af6b30fd3b48c156a82d906ea3a1ee71b0e4230201e4580daced4)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbfcfad289e1878ea66acbd49bfb8a7b566e4b5b67fc8079d2885dee7d3c736f"></a>

## remote_ip — remote_ip / 61697df2896e / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- remote_ip

<a id="canonical-52464c7cf311ab703b97340dcc0eda678a061b6b108b6cd299a834de2d4ca866"></a>

Type: `"object"`. single nested block, Optional.

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - 1. IP
Address - Specifies the remote IP to which tunnel has to be connected 2. Remote endpoint - Is a map
of IP address on per ver node basis.

Upstream description:

Defines the OPTIONS to select remote IP address for tunnel object OPTIONS available are - &#8203;1.
IP Address - Specifies the remote IP to which tunnel has to be connected &#8203;2. Remote endpoint -
Is a map of IP address on per ver node basis.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("endpoints",
    "ip")}
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
  "x-ves-oneof-field-type": "[\"endpoints\",\"ip\"]"
}
```

Terraform syntax:

```terraform
remote_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f5db47c784efd1a4ab64916aabe579576301d2d3d730836bf9f11ec471af76b"></a>

## Direct properties — remote_ip / 61697df2896e / 3

- [endpoints](resources--tunnel--reference--group-001.md#canonical-dd44f08c40299298075fbfd0395908548e967f615502b1e8f0cb35d148ab6929): complete subsection reference.

- [ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094): complete subsection reference.

<a id="canonical-060d7b7431e3b362c47806fc3291ab88c976ae2bc8ac3e8fd5e7f23bcbf35770"></a>

## Next pages — remote_ip / 61697df2896e / 4

- [remote_ip.endpoints](resources--tunnel--reference--group-001.md#canonical-dd44f08c40299298075fbfd0395908548e967f615502b1e8f0cb35d148ab6929)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-dd44f08c40299298075fbfd0395908548e967f615502b1e8f0cb35d148ab6929"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d60aa5d20d025340e5852217c42f5d4d51bbd6e481e1dff770ea8da995a8b3d5"></a>

## remote_ip.endpoints — remote_ip.endpoints / cbfdc32548cc / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- remote_ip.endpoints

<a id="canonical-57de4f6dc7314f2da1ec195b1d77cc5b11f03581d2d8f09f4babd9bb2b796a30"></a>

Type: `"object"`. single nested block, Optional.

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

Upstream description:

Provides a map of ver node name to remote node attributes Ver node should use these attributes to
configure as remote tunnel.

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
endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-4d7e8c6e664c6aeec72a6920e65ee2c52b7a219bdf373cd39167b9d4ee38ef20"></a>

## Direct properties — remote_ip.endpoints / cbfdc32548cc / 3

- [endpoints](resources--tunnel--reference--group-001.md#canonical-a07522f558dff0680a46e32bf0420f6aadf9f4300519dd60d7967d37455ef3e1): complete subsection reference.

<a id="canonical-bd14952105a672015663a671cf51ef56dae0f2e253b8478314ada415d351177b"></a>

## Next pages — remote_ip.endpoints / cbfdc32548cc / 4

- [remote_ip.endpoints.endpoints](resources--tunnel--reference--group-001.md#canonical-a07522f558dff0680a46e32bf0420f6aadf9f4300519dd60d7967d37455ef3e1)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-a07522f558dff0680a46e32bf0420f6aadf9f4300519dd60d7967d37455ef3e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a487c88d485059c635afa67f69a7d6506ab847d2cd09fffd5fd8f9c58ebfaafb"></a>

## remote_ip.endpoints.endpoints — remote_ip.endpoints.endpoints / 0dfe76df944f / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [remote_ip.endpoints](resources--tunnel--reference--group-001.md#canonical-dd44f08c40299298075fbfd0395908548e967f615502b1e8f0cb35d148ab6929)
- remote_ip.endpoints.endpoints

<a id="canonical-28f1143e34229b5b8e1898bf0b6225ee158e12a63aeeb4aeef5f48a15f10b99a"></a>

Type: `"object"`. single nested block, Optional.

Map of remote attributes to which tunnel will be established on per site node basis Every node can
have a different attributes and IP address to connect to Key is ver node name and value is Remote
node attributes.

Upstream description:

Map of remote attributes to which tunnel will be established on per site node basis Every node can
have a different attributes and IP address to connect to Key is ver node name and value is Remote
node attributes.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

Terraform syntax:

```terraform
endpoints {}
```

<a id="canonical-284c65325e401ab0c541345fe1ba95eecad8aa5428ae5418a2d80c47006fc09f"></a>

## Direct properties — remote_ip.endpoints.endpoints / 0dfe76df944f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-359908bc033b26f023835e7ee18aa7a7aad3b1a7c155db4511ecfdffdaa35f08"></a>

## Next pages — remote_ip.endpoints.endpoints / 0dfe76df944f / 4

- [remote_ip.endpoints](resources--tunnel--reference--group-001.md#canonical-dd44f08c40299298075fbfd0395908548e967f615502b1e8f0cb35d148ab6929)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5f2393e34ade26a8e2ec704e0d26588e1a51cf181c87f3f99e7562a804e80c9"></a>

## remote_ip.ip — remote_ip.ip / e9a7fa41a262 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- remote_ip.ip

<a id="canonical-eaa9415b3d8bbe25ab92b960a2f5aa434bddc77a24885aba3aff8b013d0c5764"></a>

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
ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-fcffed250a21077c296d06796fe8b8c1f78837629b0612195b4940feb45bd905"></a>

## Direct properties — remote_ip.ip / e9a7fa41a262 / 3

- [dual_stack](resources--tunnel--reference--group-001.md#canonical-27b230039631de0eabbdef861271b7c3e744161d568b9687da76036b9c93beb3): complete subsection reference.

- [ipv4](resources--tunnel--reference--group-001.md#canonical-2ed413f8e9c8b1342c1313fc986abd81f3e3915575628eefc280bbfc68810bb9): complete subsection reference.

- [ipv6](resources--tunnel--reference--group-001.md#canonical-dff547dbaf5c910c10e2b1e2ee636ae4a38149384f2f633b121eda24457ba5e1): complete subsection reference.

<a id="canonical-31119f241ad3208d65add6a15ae4da45990e0ed2b2a53b48657111c2f7c8a049"></a>

## Next pages — remote_ip.ip / e9a7fa41a262 / 4

- [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-27b230039631de0eabbdef861271b7c3e744161d568b9687da76036b9c93beb3)
- [remote_ip.ip.ipv4](resources--tunnel--reference--group-001.md#canonical-2ed413f8e9c8b1342c1313fc986abd81f3e3915575628eefc280bbfc68810bb9)
- [remote_ip.ip.ipv6](resources--tunnel--reference--group-001.md#canonical-dff547dbaf5c910c10e2b1e2ee636ae4a38149384f2f633b121eda24457ba5e1)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-27b230039631de0eabbdef861271b7c3e744161d568b9687da76036b9c93beb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28208a6d3e007341b28ef25a06d4fc1b98d4ccc8fe53710073d035f60ddeb18e"></a>

## remote_ip.ip.dual_stack — remote_ip.ip.dual_stack / 1c2abeee6694 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- remote_ip.ip.dual_stack

<a id="canonical-f14dfe7fe3ccce509017dde60a5e14edcb6a7beed5d68699792b5756df8b5682"></a>

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

<a id="canonical-0368c7983c6f8cca0b44ad262f6e9178154c9adcbd6e30bcd6a6c64e55542fc7"></a>

## Direct properties — remote_ip.ip.dual_stack / 1c2abeee6694 / 3

- [ipv4](resources--tunnel--reference--group-001.md#canonical-9be490bfded8a8c87138fac878e824b67ec46624aaadcde6acd9ed71708e595a): complete subsection reference.

- [ipv6](resources--tunnel--reference--group-001.md#canonical-63048bda7b2eb0cd614f23f79575e0dd2ef425e389116b29e4716167dc27f6a4): complete subsection reference.

<a id="canonical-762aaf707aa5a6b7cda110a53eb687b23b17b42d5a34365b52fda687e6407e32"></a>

## Next pages — remote_ip.ip.dual_stack / 1c2abeee6694 / 4

- [remote_ip.ip.dual_stack.ipv4](resources--tunnel--reference--group-001.md#canonical-9be490bfded8a8c87138fac878e824b67ec46624aaadcde6acd9ed71708e595a)
- [remote_ip.ip.dual_stack.ipv6](resources--tunnel--reference--group-001.md#canonical-63048bda7b2eb0cd614f23f79575e0dd2ef425e389116b29e4716167dc27f6a4)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-9be490bfded8a8c87138fac878e824b67ec46624aaadcde6acd9ed71708e595a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9372d6e3e13a51c24d1fd8c0fcef1cbcf74fb8f7f1113da50bdfdfaffad9101"></a>

## remote_ip.ip.dual_stack.ipv4 — remote_ip.ip.dual_stack.ipv4 / 705b8ca6bf6e / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-27b230039631de0eabbdef861271b7c3e744161d568b9687da76036b9c93beb3)
- remote_ip.ip.dual_stack.ipv4

<a id="canonical-297b42522fb9fe19340f5f1b685d508d0a1141f24ea2003c9d3be87aa76910d5"></a>

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

<a id="canonical-9ae2c61f5ff3ff2a0e50be32db12a4f90f0fd8a88390e86fb25513ae2ceaa79c"></a>

## Direct properties — remote_ip.ip.dual_stack.ipv4 / 705b8ca6bf6e / 3

<a id="canonical-35cb669f029de2d39dcb9f29072dbb17dc71dbb05b10eb2cf4ecff555e7419de"></a>

<a id="canonical-fd3530181cfd312cdfc76d64122dc70e97276c8550bba3d5e0481fabd8268bac"></a>

## addr property — remote_ip.ip.dual_stack.ipv4 / 705b8ca6bf6e / 4

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

<a id="canonical-5b4edca5cb4e259cec796d9f7f63f34e4c7d207de5ace2bb359f3d1392e52c22"></a>

## Next pages — remote_ip.ip.dual_stack.ipv4 / 705b8ca6bf6e / 5

- [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-27b230039631de0eabbdef861271b7c3e744161d568b9687da76036b9c93beb3)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-63048bda7b2eb0cd614f23f79575e0dd2ef425e389116b29e4716167dc27f6a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c6c0139b76497c99f870c762677b8d28db366bf6ccdca0b468f903503b6780d"></a>

## remote_ip.ip.dual_stack.ipv6 — remote_ip.ip.dual_stack.ipv6 / 705c0e05d908 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-27b230039631de0eabbdef861271b7c3e744161d568b9687da76036b9c93beb3)
- remote_ip.ip.dual_stack.ipv6

<a id="canonical-2ceeefbcf717de775dcea323bf6da17258e86c2d148033304c03768ea8e79c73"></a>

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

<a id="canonical-e9163b4db876dec0828a6fd1af5f8e4953c530c953686e0fb93cfd5eb6048a1d"></a>

## Direct properties — remote_ip.ip.dual_stack.ipv6 / 705c0e05d908 / 3

<a id="canonical-c4bfb16ab1607fbd1c37805718c2c5d3cceb2294fdb7f454d580aa9d15f063cd"></a>

<a id="canonical-d70d4767a60905e9d68be4c553aeabf4fbab7585046c6183e3a9dd7ed494fde3"></a>

## addr property — remote_ip.ip.dual_stack.ipv6 / 705c0e05d908 / 4

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

<a id="canonical-d7ba88de66fe518d595cc293d77c74448b8c449a85e9a3265f3431c5e65720c1"></a>

## Next pages — remote_ip.ip.dual_stack.ipv6 / 705c0e05d908 / 5

- [remote_ip.ip.dual_stack](resources--tunnel--reference--group-001.md#canonical-27b230039631de0eabbdef861271b7c3e744161d568b9687da76036b9c93beb3)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-2ed413f8e9c8b1342c1313fc986abd81f3e3915575628eefc280bbfc68810bb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a48a398ed2fd0d4553ddf0040491f729b57a7c7ddaecd23a14ac8172e9a127a7"></a>

## remote_ip.ip.ipv4 — remote_ip.ip.ipv4 / 32b852b386e2 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- remote_ip.ip.ipv4

<a id="canonical-8aaf9a2b179bcf2c84fc0af75b76b3b6735b50146f432abe686c0d87bcf10c8e"></a>

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

<a id="canonical-a9c28155f19d2cb5bd6b2c37a80fc38567bac21dddbc6926cd4ce72301043ae5"></a>

## Direct properties — remote_ip.ip.ipv4 / 32b852b386e2 / 3

<a id="canonical-b2cddcf832c0c3cd4d1c682acdd590e95cf4e08883b5356cf92e6a9274ad2693"></a>

<a id="canonical-5650d580f4aaf68e9faa1e267b8335d5c776d7012fa3a3a4fa5091d3e4cfcc77"></a>

## addr property — remote_ip.ip.ipv4 / 32b852b386e2 / 4

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

<a id="canonical-0055eb8b238b75ac1c48fe600bd064b3e26cf74cf8c95c89dee7538220b0139d"></a>

## Next pages — remote_ip.ip.ipv4 / 32b852b386e2 / 5

- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-dff547dbaf5c910c10e2b1e2ee636ae4a38149384f2f633b121eda24457ba5e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b48fc27da2b7fece31686b8b36e3bfb530735a56c2f4adf710b826d1fad1d408"></a>

## remote_ip.ip.ipv6 — remote_ip.ip.ipv6 / 78c55d1f6667 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [remote_ip](resources--tunnel--reference--group-001.md#canonical-0876ded3af35fb4ae96b24be04c79a0a600d3de000998845fb3fbcb0672344b4)
- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- remote_ip.ip.ipv6

<a id="canonical-767dacb7cb501eae25a81b02d18ed255d4225a3b274a9b12e6d6f3cf05f800b7"></a>

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

<a id="canonical-6103a5886c107cc0e8e19a1db03ae9394888809d6735c0eb682d89c41c839368"></a>

## Direct properties — remote_ip.ip.ipv6 / 78c55d1f6667 / 3

<a id="canonical-487b762a630d1762b05f453ac9b221143995e19eea8a59d10b6b66f75aec673d"></a>

<a id="canonical-340b65200c57d2f42e1ce2600b7cecae413b80436976888b090e9ad025b4596d"></a>

## addr property — remote_ip.ip.ipv6 / 78c55d1f6667 / 4

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

<a id="canonical-adb8102cbe4edeb30df5fda4a529f558a9c7185b3c0f64e6cf2f162d8d03e731"></a>

## Next pages — remote_ip.ip.ipv6 / 78c55d1f6667 / 5

- [remote_ip.ip](resources--tunnel--reference--group-001.md#canonical-b9cbe59967b3b25c4c146b24b20d87980dcaa9ab1f2cc04e521282b11e6f1094)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-2e5aa6734285ac67d12db6f01e1708a97fc9b672be50d2ae079480b57b170b1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dba85e441c4484c95af375dafe61bf7ff1b90fa7ba0185649fc4067332dd17e0"></a>

## timeouts — timeouts / 942d58f60ca3 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- timeouts

<a id="canonical-a31566559244c74d2e99aa324fdb09cf4a20f2b5152df02aa7c23c4c21d99b80"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f80ddc4fb22be9a46d3d0e8c2921f238e9f3d49656fd47946c9669bb3596cd7c"></a>

## Direct properties — timeouts / 942d58f60ca3 / 3

<a id="canonical-55382f6fc55221fd90c9887bfaff3fbc07f945e9b5a85bd0be25b04acad041a0"></a>

<a id="canonical-5257a8a1b9ea276b4cfd325a37b09dc8eb91c8c56dc98497f1fa99e78e267588"></a>

## create property — timeouts / 942d58f60ca3 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-c3c2d64e44fe4cc870b354df783ff8f0aad2052a8986e52b9ea1d74575ef7c27"></a>

<a id="canonical-cd26214812620f74c9c4c5111149794fb03200e9bd991a476a9a3de7328a6457"></a>

## delete property — timeouts / 942d58f60ca3 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-da07924ff985d9d277cc48be4ee095004c84918c8caa58cda15f686f759d3cf8"></a>

<a id="canonical-855a7d407bf2e1ee320c84d0d0ef09602fb584288e227e7e959f7d171945a256"></a>

## read property — timeouts / 942d58f60ca3 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-4c66580047f03ec051450f16b351600d1a1deb6bf799ddeb06fe63e11ae62d53"></a>

<a id="canonical-87e3b5e26cbdb307a38f110c17f362e658812a453b3c9f7ba6133c1801884fde"></a>

## update property — timeouts / 942d58f60ca3 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-88d33d76d4151118089e4dbbe642f314fe1f810b081a6862fc64bd2ac781d02a"></a>

## Next pages — timeouts / 942d58f60ca3 / 8

- [Property reference](resources--tunnel--reference--group-001.md#canonical-2086df7a33f0e7654b87d8c52dcf8dadc84e583239c313681b7b8c68a2f77ce7)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
