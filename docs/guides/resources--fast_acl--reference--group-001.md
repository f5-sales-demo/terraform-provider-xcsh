---
page_title: "xcsh_fast_acl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl reference."
---

# xcsh_fast_acl reference

<a id="canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e74d7de103fbb5e2a8f09dab9db4abbea8cf48df0193995894103d6a01aed8e4"></a>

## Property reference — Property reference / 8541b123a441 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- Property reference

<a id="canonical-40013fcd42834f00cb811e49b6bff71860e1d917fd1f4ce4517451793d17b095"></a>

## Direct properties — Property reference / 8541b123a441 / 3

<a id="canonical-0d6b6887879e0043054e4f50b8a106d10e234dc8dd4af2a9e5f46cfda3f1d507"></a>

<a id="canonical-9c37175346cf023ef5976e5c23d2a1d5bfe02ae8c7f38a8980c7b342b8d66cf4"></a>

## annotations property — Property reference / 8541b123a441 / 4

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

<a id="canonical-b585cfe846cb2882ea2d709e7e90a74ede52b2c11ca3451aaeef367f81b63606"></a>

<a id="canonical-16e43af4f4c1285108fe03c1e34d6896581f3868fae2501f0a5387dd6daa0f8c"></a>

## description property — Property reference / 8541b123a441 / 5

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

<a id="canonical-deb8fb8fe588c1767df3a4fb767a087adc2632fb49a59c1f0cf180cafdffc2ee"></a>

<a id="canonical-131b06fd0a1a5a0827c0a39e25fdab026907ee61931ad2b66df43dc0361c41d1"></a>

## disable property — Property reference / 8541b123a441 / 6

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

<a id="canonical-e25887d2961dfe4743a74f001e5cd4dc60f3babd7c379b286a3b0a1d9ec1256f"></a>

<a id="canonical-974fe677f91e9ac7ff36bd9e548c63c134b1e3330af6579adf69e0fab28f9200"></a>

## id property — Property reference / 8541b123a441 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5d25a8c662de8e79cfbbb8763596f4df8798d0324c1f80cbffd845876a03c510"></a>

<a id="canonical-3d29f9f60c042e07a6069e3d2925abc753f7dfaa6a7661f004c2c035f16043ca"></a>

## labels property — Property reference / 8541b123a441 / 8

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

<a id="canonical-344a3b2bb804b40ebd706c19f0c8b8d100421128cd4224ef0ccda8dda9471173"></a>

<a id="canonical-62a41e5c85bc373ba73b82b8c411b8ad49052e669a202256c44f773bd8675eab"></a>

## name property — Property reference / 8541b123a441 / 9

Type: `"string"`. Required.

Name of the Fast ACL. Must be unique within the namespace.

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

<a id="canonical-e9936864df6bab46f8b6cdca14c467134a8e24c2014c32a8b160d3e5e28b3ba9"></a>

<a id="canonical-cc3e2f7db55387d417e5d11da4f8510108bff7b15cfc036a5a72e2a8260bee35"></a>

## namespace property — Property reference / 8541b123a441 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Fast ACL. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [protocol_policer](resources--fast_acl--reference--group-001.md#canonical-e78ede9db2605bca51fc2e340411f74b74449f2f0b6689f14e24882a40b223e5): complete subsection reference.

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9): complete subsection reference.

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa): complete subsection reference.

- [timeouts](resources--fast_acl--reference--group-001.md#canonical-aa8259fa24db8bdfd1ab7042e02d60c3ad6a53370cad9578f959eeec65409c35): complete subsection reference.

<a id="canonical-25edca4f42ee30093ac68d3ef62cf6c24ed2e06f6bb1882af68f59c5e66a59c8"></a>

## All schema paths — Property reference / 8541b123a441 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--fast_acl--reference--group-001.md#canonical-0d6b6887879e0043054e4f50b8a106d10e234dc8dd4af2a9e5f46cfda3f1d507) |
| `description` | [description](resources--fast_acl--reference--group-001.md#canonical-b585cfe846cb2882ea2d709e7e90a74ede52b2c11ca3451aaeef367f81b63606) |
| `disable` | [disable](resources--fast_acl--reference--group-001.md#canonical-deb8fb8fe588c1767df3a4fb767a087adc2632fb49a59c1f0cf180cafdffc2ee) |
| `id` | [id](resources--fast_acl--reference--group-001.md#canonical-e25887d2961dfe4743a74f001e5cd4dc60f3babd7c379b286a3b0a1d9ec1256f) |
| `labels` | [labels](resources--fast_acl--reference--group-001.md#canonical-5d25a8c662de8e79cfbbb8763596f4df8798d0324c1f80cbffd845876a03c510) |
| `name` | [name](resources--fast_acl--reference--group-001.md#canonical-344a3b2bb804b40ebd706c19f0c8b8d100421128cd4224ef0ccda8dda9471173) |
| `namespace` | [namespace](resources--fast_acl--reference--group-001.md#canonical-e9936864df6bab46f8b6cdca14c467134a8e24c2014c32a8b160d3e5e28b3ba9) |
| `protocol_policer` | [protocol_policer](resources--fast_acl--reference--group-001.md#canonical-d24f7c70adbcbe6117fef1bfb0053b9103ad7b3fb666c1f9034e900ce4e21b3d) |
| `protocol_policer.name` | [protocol_policer.name](resources--fast_acl--reference--group-001.md#canonical-6dc8297dd4acd5b483b611af1aa81a6ed2b576032555b3ef827925ae26dd07d0) |
| `protocol_policer.namespace` | [protocol_policer.namespace](resources--fast_acl--reference--group-001.md#canonical-46543feba1d63cb340ff0b8c262c968d7903024fcb5c333053d844f2ba11bf83) |
| `protocol_policer.tenant` | [protocol_policer.tenant](resources--fast_acl--reference--group-001.md#canonical-d4e992ce1bcfe2078799661895cc1b1bdf7915bf65a2a49a94f522eebba31de1) |
| `re_acl` | [re_acl](resources--fast_acl--reference--group-001.md#canonical-8df6ba5f3fab70b9dd077b0e4fbd6abd7baeefb6f780e4b9e04c2e09679ce9f5) |
| `re_acl.all_public_vips` | [re_acl.all_public_vips](resources--fast_acl--reference--group-001.md#canonical-5bb9cc2cea9bfcfef6c5d096f67cf19fc5abed4f23d2a6ef16136d74d9811a50) |
| `re_acl.default_tenant_vip` | [re_acl.default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-12459950887bfbd5b3d160da0d91e2ad0b9085906ede418575ae7b2a7559db21) |
| `re_acl.fast_acl_rules` | [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-54cfe77ee1ed0ebcf8f83137be0fcd43f7a6ea3a5236e7353eee53ba2c684ad7) |
| `re_acl.fast_acl_rules.action` | [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-d00e439f707c80fe0a67e613ef1ec4eea6e9598c5bc4989f627bae4cba380e20) |
| `re_acl.fast_acl_rules.action.policer_action` | [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-a9e6026849b44eea03ace8fa714ea1363b005189aaad83ddbbe4039d7d203a2e) |
| `re_acl.fast_acl_rules.action.policer_action.ref` | [re_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-f3c6ec3a2187ffeafe55b66b75850a947cd7f8406397a0c1105ca63e6f502d7b) |
| `re_acl.fast_acl_rules.action.policer_action.ref.kind` | [re_acl.fast_acl_rules.action.policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-8c4a9cd87cc86ed8076598ad2b8e3c2ccbf37500a239d7d28f595efa24aec7ff) |
| `re_acl.fast_acl_rules.action.policer_action.ref.name` | [re_acl.fast_acl_rules.action.policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-9d8762c5c85206195b7c35b6aa1eca898da83dd9680ffd0b39892a1e66d1e4b3) |
| `re_acl.fast_acl_rules.action.policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-263ee18d1ea2caafc4e4624dbdbf773f0d0311d93555f7ef9214dc02331d0420) |
| `re_acl.fast_acl_rules.action.policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-81bb96da644447bdd8173cbfc6c232273532ee67507b9f47c902a13de069421e) |
| `re_acl.fast_acl_rules.action.policer_action.ref.uid` | [re_acl.fast_acl_rules.action.policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-3c224a88e4c2e8acf184dddb403c656c858b08a8f5a0bb4de61d3d2aeb83b563) |
| `re_acl.fast_acl_rules.action.protocol_policer_action` | [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-9f9355baaa86856d2a56489d11d849ac2374bd6cd4e05e2143beba47016df444) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-77f29f8b79745f80dac649574ee8ed272239e1f684ed87cd8ebc78abb5dc79cf) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-30d1e5e1e6e592594485e226181f1179c013de767419495274c6598bd1767488) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-efc9650dcb7010ca3e2bc5e67b3045e02c30b8a5f0150ed2ab90866211c4203a) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-3a228ad4e0d58aa7172b85a41457d9c19935a4e5c76e6d97f9c6e337fce0cc65) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-c12c553f206b6e84918c5d99dc63a6ad24098f3031e30dc1c14b620875434e60) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-340ee92b4ae88d8297a1ad16f51ad6b2c38c6fc2972ccf68b18c5d20dfa24693) |
| `re_acl.fast_acl_rules.action.simple_action` | [re_acl.fast_acl_rules.action.simple_action](resources--fast_acl--reference--group-001.md#canonical-41612e6845e99e7495817189d3a623210ebbe6c3c6bb439a3f7c742becd1c250) |
| `re_acl.fast_acl_rules.ip_prefix_set` | [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-b19bb8e21fe8a8da481020943516175749a832a1af6a1cc4c97a03fe5c6762f5) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref` | [re_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-1d5d2e227725df101cdbda173379544d529ed3369e8a7116458f3de15e9c510a) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [re_acl.fast_acl_rules.ip_prefix_set.ref.kind](resources--fast_acl--reference--group-001.md#canonical-d5c029f3b1f4f2b9e492b86040a4f078e028b7a400c74993d078c1f79bbb97ca) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.name` | [re_acl.fast_acl_rules.ip_prefix_set.ref.name](resources--fast_acl--reference--group-001.md#canonical-041f9f81b05af4944c7136ab326edc5726a4eed566c22d7d97da859f0cdc610b) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [re_acl.fast_acl_rules.ip_prefix_set.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-58478d180d70cdd145a6a00ca2bed50a8472bb189d13dace26b7bf6ba39149ca) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [re_acl.fast_acl_rules.ip_prefix_set.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-67e1674e854697c5381acfe3b3705c819382e5513b84206a744729f25b2e3263) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [re_acl.fast_acl_rules.ip_prefix_set.ref.uid](resources--fast_acl--reference--group-001.md#canonical-38d15e93e9e0c9c6ffe8fd949b064830cd9da78c6ffde703996b0a8a31212e0d) |
| `re_acl.fast_acl_rules.metadata` | [re_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-8f35bd4ea899b69890edc0c8b38a2c156171e60b9da0dd78b62e84f6c4d95488) |
| `re_acl.fast_acl_rules.metadata.description_spec` | [re_acl.fast_acl_rules.metadata.description_spec](resources--fast_acl--reference--group-001.md#canonical-bad09e46f6246dab4ee41eb18fb449c0c368923037bc0aee0c44becf2f3d5fd6) |
| `re_acl.fast_acl_rules.metadata.name` | [re_acl.fast_acl_rules.metadata.name](resources--fast_acl--reference--group-001.md#canonical-42742cd33260628a397ae50d622209774942b829e2405eeeb357c3e82cf59243) |
| `re_acl.fast_acl_rules.port` | [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-cad142d5e07e0ea7aeadab213c088f2bcb7a6eb44ee54766a31f0195b7b54ec2) |
| `re_acl.fast_acl_rules.port.all` | [re_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-9db0b0e21ce812e0094034a461e918e29cb59ae6b435ec3adf538a926aa0bbc7) |
| `re_acl.fast_acl_rules.port.dns` | [re_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-8edc5aed6baf6fc9f43de101afc2c67ab77043f04626f57c1715edd799ae9f67) |
| `re_acl.fast_acl_rules.port.user_defined` | [re_acl.fast_acl_rules.port.user_defined](resources--fast_acl--reference--group-001.md#canonical-ee9b7023eb3abdf6268a56aafa2202047d05eb75c628fccf0994af7f190f251c) |
| `re_acl.fast_acl_rules.prefix` | [re_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-46fc68cb61f559d01d50b0aa10ad7a58c7a4bd330677da8b81c1ddb4d10cbaad) |
| `re_acl.fast_acl_rules.prefix.prefix` | [re_acl.fast_acl_rules.prefix.prefix](resources--fast_acl--reference--group-001.md#canonical-b374c069d5e2357e283eadb1f77ac45850f16f4d902571264098d75ee61e5d18) |
| `re_acl.selected_tenant_vip` | [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-0fb2648107694b5189d753a0e601eea218b355a5fc781af072203d466552e6a4) |
| `re_acl.selected_tenant_vip.default_tenant_vip` | [re_acl.selected_tenant_vip.default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-5dfeeace05f9ef61160401511bc678334b94961fb869118f9b57fca3a9da62ab) |
| `re_acl.selected_tenant_vip.public_ip_refs` | [re_acl.selected_tenant_vip.public_ip_refs](resources--fast_acl--reference--group-001.md#canonical-e8f453344a91f06470fdb070c5a07f68adc99cfd38e8d005a6e181f1fca123e5) |
| `re_acl.selected_tenant_vip.public_ip_refs.name` | [re_acl.selected_tenant_vip.public_ip_refs.name](resources--fast_acl--reference--group-001.md#canonical-a92c93023faec5350aa32ada07a3703a71fee8491dae4cb571046d23c747b07b) |
| `re_acl.selected_tenant_vip.public_ip_refs.namespace` | [re_acl.selected_tenant_vip.public_ip_refs.namespace](resources--fast_acl--reference--group-001.md#canonical-f1c4070390e92f31793b12badfc982c83c9c647a635610bef66ef13da5f23d94) |
| `re_acl.selected_tenant_vip.public_ip_refs.tenant` | [re_acl.selected_tenant_vip.public_ip_refs.tenant](resources--fast_acl--reference--group-001.md#canonical-8af88fe2801b6b929d5bab2ebad89841ccd665670edd411dc86fbb79062f6e14) |
| `site_acl` | [site_acl](resources--fast_acl--reference--group-001.md#canonical-2458b7ab90751514fe9c081d85b5bcd6d44f3a20996e72b1025b3069e78abda7) |
| `site_acl.all_services` | [site_acl.all_services](resources--fast_acl--reference--group-001.md#canonical-81894e5e3b8c65516b240e1a02de731caee6dea8092f7aae22322fafb1f2bdd7) |
| `site_acl.fast_acl_rules` | [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-6701fdefd8163ea2cc139eb893b6d7086a96b5f1a836dff8b66224cc11a796e7) |
| `site_acl.fast_acl_rules.action` | [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-f20826b7b14cd205e0e1e475f5e9c585b92b3600b79cda80ae73cc86e84be837) |
| `site_acl.fast_acl_rules.action.policer_action` | [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-d51f73141536fa80fc5e4b70fe91ef19edb6f287c81db4e377687b86b2367414) |
| `site_acl.fast_acl_rules.action.policer_action.ref` | [site_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-611ebc487646fe3e3d2e6595a5916b640ce9b5f5b5c55f4ff4e17a03f918cc08) |
| `site_acl.fast_acl_rules.action.policer_action.ref.kind` | [site_acl.fast_acl_rules.action.policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-2a129f1a82ab68c0ce2c8b1d6d49b3b77e054b8aefd57e37723c743abaad4020) |
| `site_acl.fast_acl_rules.action.policer_action.ref.name` | [site_acl.fast_acl_rules.action.policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-9550b04a4413916d3c9629713e461f2353bb613d01fbfd6c885875fe9702f5fa) |
| `site_acl.fast_acl_rules.action.policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-e4964cc3e45c45daa030db31ce3a738d4343f0df9a7b8e1143ae8634aee8e65d) |
| `site_acl.fast_acl_rules.action.policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-a47804a981b5ac689a1e835811a3472bdf7b6050d1aff4456b6d6265bc8871c6) |
| `site_acl.fast_acl_rules.action.policer_action.ref.uid` | [site_acl.fast_acl_rules.action.policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-d0b80d7c4ef3ce7980f86812ad3730510417d4c43e6deaab5bcdc41756c840bd) |
| `site_acl.fast_acl_rules.action.protocol_policer_action` | [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-37e844396760f8f7e0fb2ac1df3d27406a221f80d54182332755d29464d94ef6) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-3c06b1af8084a6a0ad704f2376e03568c71cd7d1e3a14f4f979be7bd66eda38b) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](resources--fast_acl--reference--group-001.md#canonical-3147bf5d8f591ad4688429864cd8c404efd4b75643059727c18199d2502d28c2) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.name](resources--fast_acl--reference--group-001.md#canonical-4095e15ade95db4ee358b0b621f48c760bf423b1dde4b0dd1664291ba8e8eae9) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-de51f300b0dad6ef55cc0809b13aa5682568a30a412958e440cac40b62aea9c3) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-5b69b43243a89dc34c8a567f6daba66b3c366596fd7de376088560804516671b) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](resources--fast_acl--reference--group-001.md#canonical-3ede1a9e6d6f34744e2a90bfef8b6942ea584c9269ee125f0e33a5182d3237cd) |
| `site_acl.fast_acl_rules.action.simple_action` | [site_acl.fast_acl_rules.action.simple_action](resources--fast_acl--reference--group-001.md#canonical-19caa509ef34a9e102b99d593ec3739f2bf0aace227b654cbb6f14ab1ce9997a) |
| `site_acl.fast_acl_rules.ip_prefix_set` | [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-faac95cce9b6dae5418e29f5b71caef9ec1be258ccac321c2e647fbbdd0155b0) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref` | [site_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-044d98634fa5e9aaf791e19a2156392e2bfe86de1b19ae74c3a31ed17a6455c6) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [site_acl.fast_acl_rules.ip_prefix_set.ref.kind](resources--fast_acl--reference--group-001.md#canonical-7bbe705f2e45f3b441fec4b090726948f976741177d5ba701f89d89288ee3185) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.name` | [site_acl.fast_acl_rules.ip_prefix_set.ref.name](resources--fast_acl--reference--group-001.md#canonical-7ec01ecd4a678ea76bdf2f2206464f0a1660b248c5d7fa70e32dc0bd23f8d3c1) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [site_acl.fast_acl_rules.ip_prefix_set.ref.namespace](resources--fast_acl--reference--group-001.md#canonical-a035ccb19ff1262c29ff00f6988357a298ebb18e478bb71997b8f22edd05e39b) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [site_acl.fast_acl_rules.ip_prefix_set.ref.tenant](resources--fast_acl--reference--group-001.md#canonical-6c936044c0394a4e533718a27a9fa4a76a72d3453fff65d3b90d15d6e7bf51cc) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [site_acl.fast_acl_rules.ip_prefix_set.ref.uid](resources--fast_acl--reference--group-001.md#canonical-80f5dd9d415e8476363f5a3dd82beaaa0b22cb1bedd6834549f2677b05edd7ec) |
| `site_acl.fast_acl_rules.metadata` | [site_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-b7eb1be7324e442e15802d6beab8522440e4cf718051468c030d0784f58979a3) |
| `site_acl.fast_acl_rules.metadata.description_spec` | [site_acl.fast_acl_rules.metadata.description_spec](resources--fast_acl--reference--group-001.md#canonical-840403d784eb25775000aaa727e01c9acceeb18afe710278d7d590d8ea80c842) |
| `site_acl.fast_acl_rules.metadata.name` | [site_acl.fast_acl_rules.metadata.name](resources--fast_acl--reference--group-001.md#canonical-a7178d9ec1b8957426d0a6f0bee5f478babd5d975829801efe4946308d338223) |
| `site_acl.fast_acl_rules.port` | [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-65933cae7afd531ad93bd80fc8003fdc5e9db17a802ad74365739f78b7a4a781) |
| `site_acl.fast_acl_rules.port.all` | [site_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-f81da70f841f6f35a72ff9301c82613fa5d6302ab498cad736de8223f221b0a8) |
| `site_acl.fast_acl_rules.port.dns` | [site_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-606f38503ea99fc47d72c86ca96fddb851109c7ac6be03639e57311e8604142e) |
| `site_acl.fast_acl_rules.port.user_defined` | [site_acl.fast_acl_rules.port.user_defined](resources--fast_acl--reference--group-001.md#canonical-4f1e676c44d66b52b406c5db4546b2876a2d6a6ea6b65b45a580160ce4d414d1) |
| `site_acl.fast_acl_rules.prefix` | [site_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-8b2dce192e2535c2eaa06854edc0795272dd0417bfc0c2de4274b5081b63e201) |
| `site_acl.fast_acl_rules.prefix.prefix` | [site_acl.fast_acl_rules.prefix.prefix](resources--fast_acl--reference--group-001.md#canonical-115b6474eaa43c1c720eba2d7010fa739741ed9409a33724d345ba315d7b4356) |
| `site_acl.inside_network` | [site_acl.inside_network](resources--fast_acl--reference--group-001.md#canonical-411db025c49565219504b2727942602f20358f666dae900d1967fe089fd88b79) |
| `site_acl.interface_services` | [site_acl.interface_services](resources--fast_acl--reference--group-001.md#canonical-fab7c5d4a30e915498f87762d9c93a8c231c1706fa60b68aa9a61698589c4def) |
| `site_acl.outside_network` | [site_acl.outside_network](resources--fast_acl--reference--group-001.md#canonical-41711b1ad57ddab3aef2f545683331114c8d0ef4182e28b3a1e8033ac2beb6d6) |
| `site_acl.vip_services` | [site_acl.vip_services](resources--fast_acl--reference--group-001.md#canonical-40aa9741213119e8a7ea059cac6013255579c8706ac9924d42a9b3ea9f1f07c9) |
| `timeouts` | [timeouts](resources--fast_acl--reference--group-001.md#canonical-974f8b7bb7dcb969d97122fe835658d9a89a4c9e90d58eb6ac49152862b4292d) |
| `timeouts.create` | [timeouts.create](resources--fast_acl--reference--group-001.md#canonical-db5dec63902c5d4209fa0e3c7757f090296d06a6abca9fe1d5a96ff442653529) |
| `timeouts.delete` | [timeouts.delete](resources--fast_acl--reference--group-001.md#canonical-495abfdc5c92ebefab18ce74c75538767d2ae7bbbd7028b15ccbf47dbe53572a) |
| `timeouts.read` | [timeouts.read](resources--fast_acl--reference--group-001.md#canonical-fb0f413f6a8bd516262f86f02c738724a78edb5e454d373879931e623fd18270) |
| `timeouts.update` | [timeouts.update](resources--fast_acl--reference--group-001.md#canonical-16401a93b3f55b48fc5319f93479011de8e8675cff813cef3a1f90659504fd6f) |

<a id="canonical-825bc72071d10dcb7df105e748b8e0d258571763dd77d6bdd4061ee5dd35acdc"></a>

## Next pages — Property reference / 8541b123a441 / 12

- [protocol_policer](resources--fast_acl--reference--group-001.md#canonical-e78ede9db2605bca51fc2e340411f74b74449f2f0b6689f14e24882a40b223e5)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [timeouts](resources--fast_acl--reference--group-001.md#canonical-aa8259fa24db8bdfd1ab7042e02d60c3ad6a53370cad9578f959eeec65409c35)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-e78ede9db2605bca51fc2e340411f74b74449f2f0b6689f14e24882a40b223e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-590df4d93ce08a869bc1d0642d030ca49c203dd4d1cc98636b1d634115df79d0"></a>

## protocol_policer — protocol_policer / 4a5629d940ac / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- protocol_policer

<a id="canonical-d24f7c70adbcbe6117fef1bfb0053b9103ad7b3fb666c1f9034e900ce4e21b3d"></a>

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
protocol_policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-2efcc7198eb812bbe774907ad395ec5486344cca5edd33d741ea5a241268c3aa"></a>

## Direct properties — protocol_policer / 4a5629d940ac / 3

<a id="canonical-6dc8297dd4acd5b483b611af1aa81a6ed2b576032555b3ef827925ae26dd07d0"></a>

<a id="canonical-a26aebcb97265247f72a1e336c492c0bd4c19f02ce4ac1b543504408cd13b495"></a>

## name property — protocol_policer / 4a5629d940ac / 4

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

<a id="canonical-46543feba1d63cb340ff0b8c262c968d7903024fcb5c333053d844f2ba11bf83"></a>

<a id="canonical-ccde82ee83e02bc4f5981779924a7598a3d73ae0dad0c02e79b50898b85ef45d"></a>

## namespace property — protocol_policer / 4a5629d940ac / 5

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

<a id="canonical-d4e992ce1bcfe2078799661895cc1b1bdf7915bf65a2a49a94f522eebba31de1"></a>

<a id="canonical-f2fa833b700c86e303045182c2493cd5e7df98813aff654aa394483ce1d6dc59"></a>

## tenant property — protocol_policer / 4a5629d940ac / 6

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

<a id="canonical-6b84147ad96fe30356ad53288872f353fbfee08d930d86696afa7c2df3163878"></a>

## Next pages — protocol_policer / 4a5629d940ac / 7

- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da8db40a0b47d989b38288dab4d07f78fc4b5782219cd5fe95f1907385d01b4e"></a>

## re_acl — re_acl / 6dbe1d8be99c / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- re_acl

<a id="canonical-8df6ba5f3fab70b9dd077b0e4fbd6abd7baeefb6f780e4b9e04c2e09679ce9f5"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Upstream description:

Fast ACL definition for RE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_public_vips",
    "default_tenant_vip"),
  validators.ConflictingObjectAttributes("all_public_vips",
    "selected_tenant_vip"),
  validators.ConflictingObjectAttributes("default_tenant_vip",
    "selected_tenant_vip")}
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
  "x-ves-oneof-field-vip_choice": "[\"all_public_vips\",\"default_tenant_vip\",\"selected_tenant_vip\"]"
}
```

OneOf alternatives in this subsection:

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-8df6ba5f3fab70b9dd077b0e4fbd6abd7baeefb6f780e4b9e04c2e09679ce9f5)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-2458b7ab90751514fe9c081d85b5bcd6d44f3a20996e72b1025b3069e78abda7)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
re_acl {
  # Configure direct properties listed below.
}
```

<a id="canonical-06d3d8d85bbd193974faf0b43d9c9a919d28c05d7c018d7fc90ea052b01dcd73"></a>

## Direct properties — re_acl / 6dbe1d8be99c / 3

- [all_public_vips](resources--fast_acl--reference--group-001.md#canonical-80da462f74b16ae64c8a3532cdf33b7df97fb3747a846d32cc2a2a2e77799a65): complete subsection reference.

- [default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-463c8a62aa5273cef5eee97d26cf0bb385eaf2fa2d71442b7f7f2c854828fe9a): complete subsection reference.

- [fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04): complete subsection reference.

- [selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-b75b4b8d1eb71cf96a1d309b6304288399b79b20a45906f905cb4002a2692482): complete subsection reference.

<a id="canonical-aa0e73767e03bbcf1d54bb9b45ec7e150376fbc02feb4d63538a5ef834ebfbc0"></a>

## Next pages — re_acl / 6dbe1d8be99c / 4

- [re_acl.all_public_vips](resources--fast_acl--reference--group-001.md#canonical-80da462f74b16ae64c8a3532cdf33b7df97fb3747a846d32cc2a2a2e77799a65)
- [re_acl.default_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-463c8a62aa5273cef5eee97d26cf0bb385eaf2fa2d71442b7f7f2c854828fe9a)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-b75b4b8d1eb71cf96a1d309b6304288399b79b20a45906f905cb4002a2692482)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-80da462f74b16ae64c8a3532cdf33b7df97fb3747a846d32cc2a2a2e77799a65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6fa7034b760dfbe6f13ecbf45d9c8fe0a944877b38821b4fb81392ca023ecc9"></a>

## re_acl.all_public_vips — re_acl.all_public_vips / 6c9f1fdcddca / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- re_acl.all_public_vips

<a id="canonical-5bb9cc2cea9bfcfef6c5d096f67cf19fc5abed4f23d2a6ef16136d74d9811a50"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
all_public_vips = {}
```

<a id="canonical-7fe7a13037c84f2435774609422ec3f143a0955aa66c8fc53bc5b12e31d1146e"></a>

## Direct properties — re_acl.all_public_vips / 6c9f1fdcddca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf6aa2b7cb7275b76477f1127c76b5c82b4cb65efd3fb60b16e713bb9fabc5b3"></a>

## Next pages — re_acl.all_public_vips / 6c9f1fdcddca / 4

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-463c8a62aa5273cef5eee97d26cf0bb385eaf2fa2d71442b7f7f2c854828fe9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27702330f2cff0f98a59053be1d5a58d4ff6dd7fc3358661ed2f368e36de41d0"></a>

## re_acl.default_tenant_vip — re_acl.default_tenant_vip / 3cc106192bbd / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- re_acl.default_tenant_vip

<a id="canonical-12459950887bfbd5b3d160da0d91e2ad0b9085906ede418575ae7b2a7559db21"></a>

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
default_tenant_vip = {}
```

<a id="canonical-783708c8d0cd97c140c621e382badcb6efe2775b0b8c8a187951d2678e3f6909"></a>

## Direct properties — re_acl.default_tenant_vip / 3cc106192bbd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b13b46236ee9cafb5e4384acbe64df7b03a903e85f260d57cc6b3ff25d4c2d8c"></a>

## Next pages — re_acl.default_tenant_vip / 3cc106192bbd / 4

- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4beb6512dbcdc90604b6261f890f6cf1940229c739b1a91d00bd3d44b1108ae5"></a>

## re_acl.fast_acl_rules — re_acl.fast_acl_rules / 6dce7960ffd6 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- re_acl.fast_acl_rules

<a id="canonical-54cfe77ee1ed0ebcf8f83137be0fcd43f7a6ea3a5236e7353eee53ba2c684ad7"></a>

Type: `"object"`. list nested block, Optional.

Rules. Fast ACL rules to match. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Fast ACL rules to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
fast_acl_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-57a629cf95a9296244a984289f85e7aa81cfaf6b9450f8f9f22e8f290c306b2c"></a>

## Direct properties — re_acl.fast_acl_rules / 6dce7960ffd6 / 3

- [action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650): complete subsection reference.

- [ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-75225aa31d9175f96b9986ac016f51e012c9a9fb1fb9511771a36d70e4f8f191): complete subsection reference.

- [metadata](resources--fast_acl--reference--group-001.md#canonical-434879650adbe8f6fa7a92ecedb05a7535395ca80082d2b1d04c404f81c22f80): complete subsection reference.

- [port](resources--fast_acl--reference--group-001.md#canonical-1068e02786137d0f452f14499b95a0f43ac011e1bb03abdb77aa7f5daeb621d8): complete subsection reference.

- [prefix](resources--fast_acl--reference--group-001.md#canonical-37e1a9c946d06657f2f9ee517dcc5f21883041b123d15893376f1222de9ce560): complete subsection reference.

<a id="canonical-585997b9e5daa29e3ea0ccc81b64ec76a171af27403dd1bf3efa60016167e5d5"></a>

## Next pages — re_acl.fast_acl_rules / 6dce7960ffd6 / 4

- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650)
- [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-75225aa31d9175f96b9986ac016f51e012c9a9fb1fb9511771a36d70e4f8f191)
- [re_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-434879650adbe8f6fa7a92ecedb05a7535395ca80082d2b1d04c404f81c22f80)
- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-1068e02786137d0f452f14499b95a0f43ac011e1bb03abdb77aa7f5daeb621d8)
- [re_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-37e1a9c946d06657f2f9ee517dcc5f21883041b123d15893376f1222de9ce560)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-593e82bd48de47b9db21cede11f6ff7e4e357fe6ab717699805596ad8e635c21"></a>

## re_acl.fast_acl_rules.action — re_acl.fast_acl_rules.action / a28fe58fb5c9 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- re_acl.fast_acl_rules.action

<a id="canonical-d00e439f707c80fe0a67e613ef1ec4eea6e9598c5bc4989f627bae4cba380e20"></a>

Type: `"object"`. single nested block, Optional.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("policer_action",
    "protocol_policer_action"),
  validators.ConflictingObjectAttributes("policer_action",
    "simple_action"),
  validators.ConflictingObjectAttributes("protocol_policer_action",
    "simple_action")}
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
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-8c9782b41ad45ab4be0353b92a86f86026a9faab178f1b07430598dc6251b41b"></a>

## Direct properties — re_acl.fast_acl_rules.action / a28fe58fb5c9 / 3

- [policer_action](resources--fast_acl--reference--group-001.md#canonical-bb08bf4d667a121c5fae2a52b1d269d68cb4b60a1e9fc5b701219a5d82d373ce): complete subsection reference.

- [protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-dba567b5b8e15018996ecf0f4f7f0f6c3dc5095ac755e1dbd52673da7c37a5b2): complete subsection reference.

<a id="canonical-41612e6845e99e7495817189d3a623210ebbe6c3c6bb439a3f7c742becd1c250"></a>

<a id="canonical-a568aaabb8024c0952cf32b668645cb11ed4d5f2f21e6486e439d05522bfe536"></a>

## simple_action property — re_acl.fast_acl_rules.action / a28fe58fb5c9 / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-37b56283a37fff1383114540ee48687b0b15cec4a9e79940b3d5c85698b8668f"></a>

## Next pages — re_acl.fast_acl_rules.action / a28fe58fb5c9 / 5

- [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-bb08bf4d667a121c5fae2a52b1d269d68cb4b60a1e9fc5b701219a5d82d373ce)
- [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-dba567b5b8e15018996ecf0f4f7f0f6c3dc5095ac755e1dbd52673da7c37a5b2)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-bb08bf4d667a121c5fae2a52b1d269d68cb4b60a1e9fc5b701219a5d82d373ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21f659ec333312a93c564ff93108c35cbefa08260d263599eef174b563b6e01d"></a>

## re_acl.fast_acl_rules.action.policer_action — re_acl.fast_acl_rules.action.policer_action / 35ed6e763798 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650)
- re_acl.fast_acl_rules.action.policer_action

<a id="canonical-a9e6026849b44eea03ace8fa714ea1363b005189aaad83ddbbe4039d7d203a2e"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-d6405eb6dcb72e14b259efaee90e07ec62e216578d28702180ce5db4072d73f6"></a>

## Direct properties — re_acl.fast_acl_rules.action.policer_action / 35ed6e763798 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-091e17273d1c83b1c8b94fb7091463f0a7c6f105e1dbf5388298bc710790f4fa): complete subsection reference.

<a id="canonical-31bf2d04ad5ced7bafacc1d460aa25eae12b844e35c4583e4d8ef4427b430caf"></a>

## Next pages — re_acl.fast_acl_rules.action.policer_action / 35ed6e763798 / 4

- [re_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-091e17273d1c83b1c8b94fb7091463f0a7c6f105e1dbf5388298bc710790f4fa)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-091e17273d1c83b1c8b94fb7091463f0a7c6f105e1dbf5388298bc710790f4fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f38891fbe2c5448f68573ff941b5f76cfc0611228c141068b8dbd3a7d74c85e"></a>

## re_acl.fast_acl_rules.action.policer_action.ref — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650)
- [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-bb08bf4d667a121c5fae2a52b1d269d68cb4b60a1e9fc5b701219a5d82d373ce)
- re_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-f3c6ec3a2187ffeafe55b66b75850a947cd7f8406397a0c1105ca63e6f502d7b"></a>

Type: `"object"`. list nested block, Optional.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-70ff7cfbe16d0c7cd1af04de5cec9f62da5dd5f7801c094e231c0c23d0b70081"></a>

## Direct properties — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 3

<a id="canonical-8c4a9cd87cc86ed8076598ad2b8e3c2ccbf37500a239d7d28f595efa24aec7ff"></a>

<a id="canonical-34ab8f6881aad1945a2e40c2a1ba73c408250093c35862a8e5fdd49dfeedb459"></a>

## kind property — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 4

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

<a id="canonical-9d8762c5c85206195b7c35b6aa1eca898da83dd9680ffd0b39892a1e66d1e4b3"></a>

<a id="canonical-d4850745246ab87ac0d50e139392feeadf5ceeb919688347826e343df32d6c16"></a>

## name property — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 5

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

<a id="canonical-263ee18d1ea2caafc4e4624dbdbf773f0d0311d93555f7ef9214dc02331d0420"></a>

<a id="canonical-9c5684b69f7c4deedbcdc6ea1fcf980c707ab7717c7a33ffb4c27e573ccf2680"></a>

## namespace property — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 6

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

<a id="canonical-81bb96da644447bdd8173cbfc6c232273532ee67507b9f47c902a13de069421e"></a>

<a id="canonical-94cbf221f51c3e735a39db5f96fa36037c9cfe6a294b5e2e38aca1af8a3416fd"></a>

## tenant property — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 7

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

<a id="canonical-3c224a88e4c2e8acf184dddb403c656c858b08a8f5a0bb4de61d3d2aeb83b563"></a>

<a id="canonical-5ce740072d05aa53faf1a3f8a79ceb80ec604e1d24a52cc52f150fcfe8f2ee67"></a>

## uid property — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 8

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

<a id="canonical-e03ba757e409f458914f91f60c59d7c2ddd5916c924a0572efbb1cf6a789cde7"></a>

## Next pages — re_acl.fast_acl_rules.action.policer_action.ref / fe1594e374c5 / 9

- [re_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-bb08bf4d667a121c5fae2a52b1d269d68cb4b60a1e9fc5b701219a5d82d373ce)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-dba567b5b8e15018996ecf0f4f7f0f6c3dc5095ac755e1dbd52673da7c37a5b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5796e9baaa1300df63f2f0b8fecadd50e0ed300c07d5a1cf3fb709fbbf680fc6"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action — re_acl.fast_acl_rules.action.protocol_policer_action / 212fe050f8ab / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650)
- re_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-9f9355baaa86856d2a56489d11d849ac2374bd6cd4e05e2143beba47016df444"></a>

Type: `"object"`. single nested block, Optional.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
protocol_policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-09363356a410564a18449224a192c464dccb5370c601301ac87d4dfe7b6f541a"></a>

## Direct properties — re_acl.fast_acl_rules.action.protocol_policer_action / 212fe050f8ab / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-2a58534980b099108b00f45cae940273fab1405d013b50da42d1a67ab0fc6a0e): complete subsection reference.

<a id="canonical-3de4326446b157d63684a98692d27a4cc4c83e12e0c0efe012969560f3508a9c"></a>

## Next pages — re_acl.fast_acl_rules.action.protocol_policer_action / 212fe050f8ab / 4

- [re_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-2a58534980b099108b00f45cae940273fab1405d013b50da42d1a67ab0fc6a0e)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-2a58534980b099108b00f45cae940273fab1405d013b50da42d1a67ab0fc6a0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-988cef25e817108c2164378390f4c92f3d1930f5da9e3c81867b7b2810ad9bac"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action.ref — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-daf90eb78d75e4d7f52d9230b45417f91817a391eced254c1517049435316650)
- [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-dba567b5b8e15018996ecf0f4f7f0f6c3dc5095ac755e1dbd52673da7c37a5b2)
- re_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-77f29f8b79745f80dac649574ee8ed272239e1f684ed87cd8ebc78abb5dc79cf"></a>

Type: `"object"`. list nested block, Optional.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-446aa5266892395435603d919cd3cfc393bfcba2482a06ce90b3f49473f7502e"></a>

## Direct properties — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 3

<a id="canonical-30d1e5e1e6e592594485e226181f1179c013de767419495274c6598bd1767488"></a>

<a id="canonical-1837a5eb13d139bf6977a731e377ed711331ee7d56f1e90f59c2c2f2511050ea"></a>

## kind property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 4

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

<a id="canonical-efc9650dcb7010ca3e2bc5e67b3045e02c30b8a5f0150ed2ab90866211c4203a"></a>

<a id="canonical-1b794083db0bdd6ed9d43fb5c922ed5a75c19a9921153c2510be65f551421642"></a>

## name property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 5

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

<a id="canonical-3a228ad4e0d58aa7172b85a41457d9c19935a4e5c76e6d97f9c6e337fce0cc65"></a>

<a id="canonical-b9c1a27d0140a03534a106e3833eb689e9320f267b880926f9097540c052ba36"></a>

## namespace property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 6

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

<a id="canonical-c12c553f206b6e84918c5d99dc63a6ad24098f3031e30dc1c14b620875434e60"></a>

<a id="canonical-95cf1213b979ac3382bdf4afd87cd4d4e3e88d3ec940acf7ed00f78ebcf7bd8c"></a>

## tenant property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 7

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

<a id="canonical-340ee92b4ae88d8297a1ad16f51ad6b2c38c6fc2972ccf68b18c5d20dfa24693"></a>

<a id="canonical-745d93765c8e5b38419698c3b56cafea42ac919502a3d96b65d3c326e29dc120"></a>

## uid property — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 8

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

<a id="canonical-30d83ed900a86afeebfba335174eff534d5a06bf4a6176c258cbe17df86b7e09"></a>

## Next pages — re_acl.fast_acl_rules.action.protocol_policer_action.ref / b8d7e661f3e3 / 9

- [re_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-dba567b5b8e15018996ecf0f4f7f0f6c3dc5095ac755e1dbd52673da7c37a5b2)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-75225aa31d9175f96b9986ac016f51e012c9a9fb1fb9511771a36d70e4f8f191"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9c9cad29eb622e233acfca093b81192d0de4eb4452570472a07e6c06f840f40"></a>

## re_acl.fast_acl_rules.ip_prefix_set — re_acl.fast_acl_rules.ip_prefix_set / e82408e0795d / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- re_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-b19bb8e21fe8a8da481020943516175749a832a1af6a1cc4c97a03fe5c6762f5"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef1db3ee48aea13b3d1f427e413355c47a871f4ee208cfa1b63c38df617dc609"></a>

## Direct properties — re_acl.fast_acl_rules.ip_prefix_set / e82408e0795d / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-36af56ad1b519f4412bc641235a1e2731799a7119352d7b0602523fa5d4a3dc1): complete subsection reference.

<a id="canonical-ea4aeec1a2f3019726b73606d7b4340dc2440400005cfbd0188b3cdf900d694e"></a>

## Next pages — re_acl.fast_acl_rules.ip_prefix_set / e82408e0795d / 4

- [re_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-36af56ad1b519f4412bc641235a1e2731799a7119352d7b0602523fa5d4a3dc1)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-36af56ad1b519f4412bc641235a1e2731799a7119352d7b0602523fa5d4a3dc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9618a58e15ca4027f60cf7ab72b0625335ca445364bfb82d4f57af606e122b7f"></a>

## re_acl.fast_acl_rules.ip_prefix_set.ref — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-75225aa31d9175f96b9986ac016f51e012c9a9fb1fb9511771a36d70e4f8f191)
- re_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-1d5d2e227725df101cdbda173379544d529ed3369e8a7116458f3de15e9c510a"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-495f18c7cde8f1b0a84f13086303643c0d54fe25572269d4fbe0336a0e3d1e6a"></a>

## Direct properties — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 3

<a id="canonical-d5c029f3b1f4f2b9e492b86040a4f078e028b7a400c74993d078c1f79bbb97ca"></a>

<a id="canonical-49160d007ae0a78c5d52226550fe684becd62f448930d06ccf8c57d014238ebf"></a>

## kind property — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 4

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

<a id="canonical-041f9f81b05af4944c7136ab326edc5726a4eed566c22d7d97da859f0cdc610b"></a>

<a id="canonical-c21e6df6055b88b94ec33920dd9c48c5e07b11bf17b3826851c0ee87fa9e9464"></a>

## name property — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 5

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

<a id="canonical-58478d180d70cdd145a6a00ca2bed50a8472bb189d13dace26b7bf6ba39149ca"></a>

<a id="canonical-3f4792c3f85a1ee3216e6096b2f113a7c74a57b53cd7ca99cf48372b25bcff4b"></a>

## namespace property — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 6

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

<a id="canonical-67e1674e854697c5381acfe3b3705c819382e5513b84206a744729f25b2e3263"></a>

<a id="canonical-9160c32685e9562945f9081ebc9c5da3abf2137bcc648e50826b41c3afdb58fd"></a>

## tenant property — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 7

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

<a id="canonical-38d15e93e9e0c9c6ffe8fd949b064830cd9da78c6ffde703996b0a8a31212e0d"></a>

<a id="canonical-ad816668409f8cc3592e3b645f8cf80fa11ce0e752f87f6f02b928c2935a6313"></a>

## uid property — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 8

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

<a id="canonical-f6ed063f9f17de565862ef35a89ac6d3a6df3277e8bcb556af1eed6e05d82a5d"></a>

## Next pages — re_acl.fast_acl_rules.ip_prefix_set.ref / 7a892e580787 / 9

- [re_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-75225aa31d9175f96b9986ac016f51e012c9a9fb1fb9511771a36d70e4f8f191)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-434879650adbe8f6fa7a92ecedb05a7535395ca80082d2b1d04c404f81c22f80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9583eea4ba3bad2fde909322f97aebd20ee216b7dbcd14a9dd9a3b61ae08e694"></a>

## re_acl.fast_acl_rules.metadata — re_acl.fast_acl_rules.metadata / 884fc978de5e / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- re_acl.fast_acl_rules.metadata

<a id="canonical-8f35bd4ea899b69890edc0c8b38a2c156171e60b9da0dd78b62e84f6c4d95488"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-87b3479286bc1a38230b0e7abd86881389dbac9e8db6a9ad72ec69e72c51858d"></a>

## Direct properties — re_acl.fast_acl_rules.metadata / 884fc978de5e / 3

<a id="canonical-bad09e46f6246dab4ee41eb18fb449c0c368923037bc0aee0c44becf2f3d5fd6"></a>

<a id="canonical-c2fc975a7c2f4623a8795ab58422ee0392f960c7a86f751fe62695581e9d3755"></a>

## description_spec property — re_acl.fast_acl_rules.metadata / 884fc978de5e / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-42742cd33260628a397ae50d622209774942b829e2405eeeb357c3e82cf59243"></a>

<a id="canonical-7b59f0c1ffb284965e645ec2ad85a3c2ccaccf45b43f3c95f89ed3ea668b8d6f"></a>

## name property — re_acl.fast_acl_rules.metadata / 884fc978de5e / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-db4bf212754ba11565ff940cbc89f795116e92217b6edad4d4b36efa2780de30"></a>

## Next pages — re_acl.fast_acl_rules.metadata / 884fc978de5e / 6

- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-1068e02786137d0f452f14499b95a0f43ac011e1bb03abdb77aa7f5daeb621d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54e0e8165200a47b8d8b421dbf7a1609aec29881ecd0009624fa1c29d45ca8bf"></a>

## re_acl.fast_acl_rules.port — re_acl.fast_acl_rules.port / e51934857bb9 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- re_acl.fast_acl_rules.port

<a id="canonical-cad142d5e07e0ea7aeadab213c088f2bcb7a6eb44ee54766a31f0195b7b54ec2"></a>

Type: `"object"`. list nested block, Optional.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all",
    "dns"),
  validators.ConflictingListObjectAttributes("all",
    "user_defined"),
  validators.ConflictingListObjectAttributes("dns",
    "user_defined")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb673113f833c5f919192e0e6dbd62a117965b409f4f6069f52751eb832fb812"></a>

## Direct properties — re_acl.fast_acl_rules.port / e51934857bb9 / 3

- [all](resources--fast_acl--reference--group-001.md#canonical-8abb650565567195bc1dfe57d533cebf185196490212f5bc26607a171b4a8329): complete subsection reference.

- [dns](resources--fast_acl--reference--group-001.md#canonical-5cb9fdc9d1a74215f191069bb1799e01d1041dfd81a4da4fe63504993aa54860): complete subsection reference.

<a id="canonical-ee9b7023eb3abdf6268a56aafa2202047d05eb75c628fccf0994af7f190f251c"></a>

<a id="canonical-3abf9a97d00811333df0c951a62f2b2cce270a1246a88763ac3069951c118cc8"></a>

## user_defined property — re_acl.fast_acl_rules.port / e51934857bb9 / 4

Type: `"number"`. Optional.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

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

<a id="canonical-11f27c5248d5c95dc9d2c50f2029a3256a0e3695ecbd0b27db19c30853fa9baa"></a>

## Next pages — re_acl.fast_acl_rules.port / e51934857bb9 / 5

- [re_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-8abb650565567195bc1dfe57d533cebf185196490212f5bc26607a171b4a8329)
- [re_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-5cb9fdc9d1a74215f191069bb1799e01d1041dfd81a4da4fe63504993aa54860)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-8abb650565567195bc1dfe57d533cebf185196490212f5bc26607a171b4a8329"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5020ab0c76cd9f33e39ea65287d09057203e04616c7488c1e2c97eb76bfc7cd2"></a>

## re_acl.fast_acl_rules.port.all — re_acl.fast_acl_rules.port.all / 9a2a8fb42b10 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-1068e02786137d0f452f14499b95a0f43ac011e1bb03abdb77aa7f5daeb621d8)
- re_acl.fast_acl_rules.port.all

<a id="canonical-9db0b0e21ce812e0094034a461e918e29cb59ae6b435ec3adf538a926aa0bbc7"></a>

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
all = {}
```

<a id="canonical-b82bc4fc547ab7875a5062d9b83cc3515c73b4020d455af260787f13740ecbcf"></a>

## Direct properties — re_acl.fast_acl_rules.port.all / 9a2a8fb42b10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69976c84878582e40d0c6bac6d361792be449c295cce06b385038f639e22be05"></a>

## Next pages — re_acl.fast_acl_rules.port.all / 9a2a8fb42b10 / 4

- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-1068e02786137d0f452f14499b95a0f43ac011e1bb03abdb77aa7f5daeb621d8)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-5cb9fdc9d1a74215f191069bb1799e01d1041dfd81a4da4fe63504993aa54860"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d04b1bee3533348343469a47d40ffc47e2585925bcd208ba14d927b50150fa97"></a>

## re_acl.fast_acl_rules.port.dns — re_acl.fast_acl_rules.port.dns / ec7e5b01007f / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-1068e02786137d0f452f14499b95a0f43ac011e1bb03abdb77aa7f5daeb621d8)
- re_acl.fast_acl_rules.port.dns

<a id="canonical-8edc5aed6baf6fc9f43de101afc2c67ab77043f04626f57c1715edd799ae9f67"></a>

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

<a id="canonical-d90a7cf8338dcc5f73732e3b65db188478a02b8ff1b4913148eeb30a41fd7862"></a>

## Direct properties — re_acl.fast_acl_rules.port.dns / ec7e5b01007f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65eef2de653b907788d32606286944b13ef6da073d3ebea5c65bbf6285191c40"></a>

## Next pages — re_acl.fast_acl_rules.port.dns / ec7e5b01007f / 4

- [re_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-1068e02786137d0f452f14499b95a0f43ac011e1bb03abdb77aa7f5daeb621d8)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-37e1a9c946d06657f2f9ee517dcc5f21883041b123d15893376f1222de9ce560"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-179ae852912a65f603dea72b7b0fffcbc5144c0a89fc1b895502125dd7dc75c9"></a>

## re_acl.fast_acl_rules.prefix — re_acl.fast_acl_rules.prefix / 6ff554bf56dd / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- re_acl.fast_acl_rules.prefix

<a id="canonical-46fc68cb61f559d01d50b0aa10ad7a58c7a4bd330677da8b81c1ddb4d10cbaad"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
prefix {
  # Configure direct properties listed below.
}
```

<a id="canonical-7c454e1062e97193b0e473b2568618223d2a3d3a2a8b0347f8b8b5fb85d9c52b"></a>

## Direct properties — re_acl.fast_acl_rules.prefix / 6ff554bf56dd / 3

<a id="canonical-b374c069d5e2357e283eadb1f77ac45850f16f4d902571264098d75ee61e5d18"></a>

<a id="canonical-527782bf689c0e0e885b3616e134ea1534ac5ca0ef8227b897ea4327864097c3"></a>

## prefix property — re_acl.fast_acl_rules.prefix / 6ff554bf56dd / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-6d35ac5eef12f01b45c9e3e25cbccb0bf83f255e484f84b51c58817f263f32df"></a>

## Next pages — re_acl.fast_acl_rules.prefix / 6ff554bf56dd / 5

- [re_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-5ccf8a6145cbc7cddfde99ca397242c35d88e545047bd41066d45e5289153e04)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-b75b4b8d1eb71cf96a1d309b6304288399b79b20a45906f905cb4002a2692482"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aacfcf3fa7045703544a860365e3278787c3d503fcc8a9ecbc8660fe78405dc4"></a>

## re_acl.selected_tenant_vip — re_acl.selected_tenant_vip / 4b48c4e2e6c7 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- re_acl.selected_tenant_vip

<a id="canonical-0fb2648107694b5189d753a0e601eea218b355a5fc781af072203d466552e6a4"></a>

Type: `"object"`. single nested block, Optional.

Specific Tenant VIP. Select various tenant public VIP(s)

Upstream description:

Select various tenant public VIP(s)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("public_ip_refs")}
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
selected_tenant_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-a185c4258cda9b3347f3bf358ba5ab3adf1ea701411773c05479fc301bd83360"></a>

## Direct properties — re_acl.selected_tenant_vip / 4b48c4e2e6c7 / 3

<a id="canonical-5dfeeace05f9ef61160401511bc678334b94961fb869118f9b57fca3a9da62ab"></a>

<a id="canonical-69d978c4a2d12f937aec15d2c7a86fbcc6cdb2f9738f026a3c86314278a5fb02"></a>

## default_tenant_vip property — re_acl.selected_tenant_vip / 4b48c4e2e6c7 / 4

Type: `"bool"`. Optional.

Include tenant VIP in list of specific VIP(s).

Upstream description:

Include tenant VIP in list of specific VIP(s)

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

- [public_ip_refs](resources--fast_acl--reference--group-001.md#canonical-90cea813b45693dee835c801497a4a01c7442b8afbcc136061532ffe36400143): complete subsection reference.

<a id="canonical-6810609757bb23ddeafe4921d50c2066863e1e265864e8fb7501fade61144ad1"></a>

## Next pages — re_acl.selected_tenant_vip / 4b48c4e2e6c7 / 5

- [re_acl.selected_tenant_vip.public_ip_refs](resources--fast_acl--reference--group-001.md#canonical-90cea813b45693dee835c801497a4a01c7442b8afbcc136061532ffe36400143)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-90cea813b45693dee835c801497a4a01c7442b8afbcc136061532ffe36400143"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-059d7cf5c1ca6e2fea6ff2b0263877ce7a318adb9307385475b32cb99903594f"></a>

## re_acl.selected_tenant_vip.public_ip_refs — re_acl.selected_tenant_vip.public_ip_refs / 3d7bfbdb80a4 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [re_acl](resources--fast_acl--reference--group-001.md#canonical-6d2aac9644c53fc6a77eefe3bd0ef7859882b00da52129c10715f4ffd23eb8a9)
- [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-b75b4b8d1eb71cf96a1d309b6304288399b79b20a45906f905cb4002a2692482)
- re_acl.selected_tenant_vip.public_ip_refs

<a id="canonical-e8f453344a91f06470fdb070c5a07f68adc99cfd38e8d005a6e181f1fca123e5"></a>

Type: `"object"`. list nested block, Optional.

Select Public VIP(s). Select additional public VIP(s)

Upstream description:

Select additional public VIP(s)

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
public_ip_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-9b425c0728400aeefa06a1578768c65fb9d4e43c7a6978b4beafd37f35ba4813"></a>

## Direct properties — re_acl.selected_tenant_vip.public_ip_refs / 3d7bfbdb80a4 / 3

<a id="canonical-a92c93023faec5350aa32ada07a3703a71fee8491dae4cb571046d23c747b07b"></a>

<a id="canonical-09f706e5dbfb1d7ead76c21fc240883b8b7ab7f47caa536393fb853f21d70529"></a>

## name property — re_acl.selected_tenant_vip.public_ip_refs / 3d7bfbdb80a4 / 4

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

<a id="canonical-f1c4070390e92f31793b12badfc982c83c9c647a635610bef66ef13da5f23d94"></a>

<a id="canonical-8e8bfa03c8c9fb9d6551418fa31a5b1f06634155dc2e2f9cf8023cfdd5153306"></a>

## namespace property — re_acl.selected_tenant_vip.public_ip_refs / 3d7bfbdb80a4 / 5

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

<a id="canonical-8af88fe2801b6b929d5bab2ebad89841ccd665670edd411dc86fbb79062f6e14"></a>

<a id="canonical-2971dc0d1ad3f6366715f187acf82c30d7fa74416f4221bc67a12bdc393944ff"></a>

## tenant property — re_acl.selected_tenant_vip.public_ip_refs / 3d7bfbdb80a4 / 6

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

<a id="canonical-cbd973c7489805868a32b59254b6f8c79391fbf68acd9c13d6fd8e790953b8d6"></a>

## Next pages — re_acl.selected_tenant_vip.public_ip_refs / 3d7bfbdb80a4 / 7

- [re_acl.selected_tenant_vip](resources--fast_acl--reference--group-001.md#canonical-b75b4b8d1eb71cf96a1d309b6304288399b79b20a45906f905cb4002a2692482)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e081f56610b3068964a1d4078c28166c270063036f6cde9d5fa1473e1fd8e414"></a>

## site_acl — site_acl / dcdeea874a06 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- site_acl

<a id="canonical-2458b7ab90751514fe9c081d85b5bcd6d44f3a20996e72b1025b3069e78abda7"></a>

Type: `"object"`. single nested block, Optional.

Fast ACL for Site. Fast ACL definition for Site.

Upstream description:

Fast ACL definition for Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_services",
    "interface_services"),
  validators.ConflictingObjectAttributes("all_services",
    "vip_services"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("interface_services",
    "vip_services")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]",
  "x-ves-oneof-field-vip_choice": "[\"all_services\",\"interface_services\",\"vip_services\"]"
}
```

Terraform syntax:

```terraform
site_acl {
  # Configure direct properties listed below.
}
```

<a id="canonical-44947f418b92aa04abe7b31129b00aa9db88ef12a277d6c41e8b0aa6610cac8f"></a>

## Direct properties — site_acl / dcdeea874a06 / 3

- [all_services](resources--fast_acl--reference--group-001.md#canonical-c2f7765358e28e9dff6f948b6cd96fd2a4f6756255960fbc63ad05bb0a6c685b): complete subsection reference.

- [fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9): complete subsection reference.

- [inside_network](resources--fast_acl--reference--group-001.md#canonical-455e5c0734aa4e53d8893ae4b23530caf37f03dec57285620bdac82f1c9a8926): complete subsection reference.

- [interface_services](resources--fast_acl--reference--group-001.md#canonical-7980ac9d18076f9d21ec150822f9b762557552b195b174490fe718216b8de7fa): complete subsection reference.

- [outside_network](resources--fast_acl--reference--group-001.md#canonical-10098691a31f88a74498999bb34c120c806e5c433fd33ac8eb85fc9707780bd4): complete subsection reference.

- [vip_services](resources--fast_acl--reference--group-001.md#canonical-59d17558c6ec64d80167a83cc56d3e3e9178c42f258a640cbd1c902dbf32a851): complete subsection reference.

<a id="canonical-8937cf57e41fbd177a8186c5f434d00fcdf657c3b8540e1d5ef787d78a1e483c"></a>

## Next pages — site_acl / dcdeea874a06 / 4

- [site_acl.all_services](resources--fast_acl--reference--group-001.md#canonical-c2f7765358e28e9dff6f948b6cd96fd2a4f6756255960fbc63ad05bb0a6c685b)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.inside_network](resources--fast_acl--reference--group-001.md#canonical-455e5c0734aa4e53d8893ae4b23530caf37f03dec57285620bdac82f1c9a8926)
- [site_acl.interface_services](resources--fast_acl--reference--group-001.md#canonical-7980ac9d18076f9d21ec150822f9b762557552b195b174490fe718216b8de7fa)
- [site_acl.outside_network](resources--fast_acl--reference--group-001.md#canonical-10098691a31f88a74498999bb34c120c806e5c433fd33ac8eb85fc9707780bd4)
- [site_acl.vip_services](resources--fast_acl--reference--group-001.md#canonical-59d17558c6ec64d80167a83cc56d3e3e9178c42f258a640cbd1c902dbf32a851)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-c2f7765358e28e9dff6f948b6cd96fd2a4f6756255960fbc63ad05bb0a6c685b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a42fbb27d29bafb8a25e5d6f609225c33b5a4766bc7cd0ec8c2dbbf08748811c"></a>

## site_acl.all_services — site_acl.all_services / 639406096d2b / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- site_acl.all_services

<a id="canonical-81894e5e3b8c65516b240e1a02de731caee6dea8092f7aae22322fafb1f2bdd7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all services.

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
all_services = {}
```

<a id="canonical-ab5fda0583cc7b5a0df34e2657344c90e620eeab739de28417ea282958ef4bfa"></a>

## Direct properties — site_acl.all_services / 639406096d2b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e430c1cabcc57de7c6e7834af41c3f4cc977b1e66c7cbafe7ec18683a3231bc"></a>

## Next pages — site_acl.all_services / 639406096d2b / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbabd576d38b2eb904653b3beeb46dcbc6698840df0aa3701d17c11074581475"></a>

## site_acl.fast_acl_rules — site_acl.fast_acl_rules / 35556e2da098 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- site_acl.fast_acl_rules

<a id="canonical-6701fdefd8163ea2cc139eb893b6d7086a96b5f1a836dff8b66224cc11a796e7"></a>

Type: `"object"`. list nested block, Optional.

Rules. Fast ACL rules to match.

Upstream description:

Fast ACL rules to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
fast_acl_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e57a31f8ff2f3d55a1aef4f42295122d5beea7debe753eb2e518c68906004dd"></a>

## Direct properties — site_acl.fast_acl_rules / 35556e2da098 / 3

- [action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a): complete subsection reference.

- [ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-3bdde70d9f86bc4ab339f80e97e279527b39c7a92c58e20b918bb1b3ad81f5c4): complete subsection reference.

- [metadata](resources--fast_acl--reference--group-001.md#canonical-91b94cbe58997c6d4e6136e25b06501c2a3d4ea0fc3551d9a445e3255d6b13ef): complete subsection reference.

- [port](resources--fast_acl--reference--group-001.md#canonical-95d80e88d1961822b37e11b50d5773c35cc39b5fb71cd6c3b0797c744f2ec242): complete subsection reference.

- [prefix](resources--fast_acl--reference--group-001.md#canonical-8a4c50d1b3c1f23a5d56b15a28c7570a7fd37b5a76b2d8df6a07c35e9bc423c9): complete subsection reference.

<a id="canonical-747642ec23eca8457ba6cb72c88f39840b84f75e7cfb52cf8d0414eadd1c2430"></a>

## Next pages — site_acl.fast_acl_rules / 35556e2da098 / 4

- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a)
- [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-3bdde70d9f86bc4ab339f80e97e279527b39c7a92c58e20b918bb1b3ad81f5c4)
- [site_acl.fast_acl_rules.metadata](resources--fast_acl--reference--group-001.md#canonical-91b94cbe58997c6d4e6136e25b06501c2a3d4ea0fc3551d9a445e3255d6b13ef)
- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-95d80e88d1961822b37e11b50d5773c35cc39b5fb71cd6c3b0797c744f2ec242)
- [site_acl.fast_acl_rules.prefix](resources--fast_acl--reference--group-001.md#canonical-8a4c50d1b3c1f23a5d56b15a28c7570a7fd37b5a76b2d8df6a07c35e9bc423c9)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f03bc4cba55bb98dc54be139b5dc6cf251b2d1fcac113abd5c7c8ccf3a4bc0ef"></a>

## site_acl.fast_acl_rules.action — site_acl.fast_acl_rules.action / 3498accdfa79 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- site_acl.fast_acl_rules.action

<a id="canonical-f20826b7b14cd205e0e1e475f5e9c585b92b3600b79cda80ae73cc86e84be837"></a>

Type: `"object"`. single nested block, Optional.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("policer_action",
    "protocol_policer_action"),
  validators.ConflictingObjectAttributes("policer_action",
    "simple_action"),
  validators.ConflictingObjectAttributes("protocol_policer_action",
    "simple_action")}
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
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3b769e2f8ea497a3acf3187d984176be50ffb300c390e715ef8ce9b937244345"></a>

## Direct properties — site_acl.fast_acl_rules.action / 3498accdfa79 / 3

- [policer_action](resources--fast_acl--reference--group-001.md#canonical-a98f96081a0cc62b521a0709fb8a65955dc9421b7ce6ded1b811a520a1c09cc1): complete subsection reference.

- [protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-eb83f3dcffca45eecaa610df6066008f3323863624ebb1505429a9d7b701c7c0): complete subsection reference.

<a id="canonical-19caa509ef34a9e102b99d593ec3739f2bf0aace227b654cbb6f14ab1ce9997a"></a>

<a id="canonical-fdd0cf3f3c11afb1269dddf059ff09bf1c911109718e5337b6e0c20f1f788fe3"></a>

## simple_action property — site_acl.fast_acl_rules.action / 3498accdfa79 / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-758b67f7aa14b4357b41b677c181c6f32b76b03b62731f100710557b8f60bc69"></a>

## Next pages — site_acl.fast_acl_rules.action / 3498accdfa79 / 5

- [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-a98f96081a0cc62b521a0709fb8a65955dc9421b7ce6ded1b811a520a1c09cc1)
- [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-eb83f3dcffca45eecaa610df6066008f3323863624ebb1505429a9d7b701c7c0)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-a98f96081a0cc62b521a0709fb8a65955dc9421b7ce6ded1b811a520a1c09cc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a77c302eaecadfc4040d3ec02b5eb52add23481a4aca1eb36aaf86fcb215220"></a>

## site_acl.fast_acl_rules.action.policer_action — site_acl.fast_acl_rules.action.policer_action / 8728aa521b25 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a)
- site_acl.fast_acl_rules.action.policer_action

<a id="canonical-d51f73141536fa80fc5e4b70fe91ef19edb6f287c81db4e377687b86b2367414"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-9345ab5d4526e86d216fc56cb34b48f0e9097e50fa298ad94a93858a2e618a21"></a>

## Direct properties — site_acl.fast_acl_rules.action.policer_action / 8728aa521b25 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-e1afcce271f9a7bbe64644eb4c73111b55ab0db918a0c4b8a03dbdebde9684d8): complete subsection reference.

<a id="canonical-919594f706193f7cde8fbefaf7885bc887b36ca3c7c32e1be0f797b08a7227de"></a>

## Next pages — site_acl.fast_acl_rules.action.policer_action / 8728aa521b25 / 4

- [site_acl.fast_acl_rules.action.policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-e1afcce271f9a7bbe64644eb4c73111b55ab0db918a0c4b8a03dbdebde9684d8)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-e1afcce271f9a7bbe64644eb4c73111b55ab0db918a0c4b8a03dbdebde9684d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9fb8dc8aeb968160ae26495197fa8dfa9118633bfb39d41462be6cf256c325c"></a>

## site_acl.fast_acl_rules.action.policer_action.ref — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a)
- [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-a98f96081a0cc62b521a0709fb8a65955dc9421b7ce6ded1b811a520a1c09cc1)
- site_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-611ebc487646fe3e3d2e6595a5916b640ce9b5f5b5c55f4ff4e17a03f918cc08"></a>

Type: `"object"`. list nested block, Optional.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-7664be2b96385cb08b460a96927517139f48c43e2351353126cf86e4b2bb23ca"></a>

## Direct properties — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 3

<a id="canonical-2a129f1a82ab68c0ce2c8b1d6d49b3b77e054b8aefd57e37723c743abaad4020"></a>

<a id="canonical-a6df97cf43eb80292e367578cb751dba7b6a795e1c72962e8ad54ae53a83d2c2"></a>

## kind property — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 4

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

<a id="canonical-9550b04a4413916d3c9629713e461f2353bb613d01fbfd6c885875fe9702f5fa"></a>

<a id="canonical-356f8a964bd2f8da568065e382ccd5de410666e734917f1044028825246e3ce0"></a>

## name property — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 5

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

<a id="canonical-e4964cc3e45c45daa030db31ce3a738d4343f0df9a7b8e1143ae8634aee8e65d"></a>

<a id="canonical-24bf2e320e60cccf971722d56ed302c5d266779542e561e8fb40544959f34a97"></a>

## namespace property — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 6

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

<a id="canonical-a47804a981b5ac689a1e835811a3472bdf7b6050d1aff4456b6d6265bc8871c6"></a>

<a id="canonical-b73574c744ae106b7a4f261bf3c6d82eedea754072594e2425461ed4480d62a2"></a>

## tenant property — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 7

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

<a id="canonical-d0b80d7c4ef3ce7980f86812ad3730510417d4c43e6deaab5bcdc41756c840bd"></a>

<a id="canonical-61bc0c9090c28e41dbc96792c1859a11b2c6bdfdb24bf989b098cb70491010fd"></a>

## uid property — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 8

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

<a id="canonical-ea6d99e2426a16322fce790e246df5cd3c5dcc4d2c1de67180dde76f5e0877b9"></a>

## Next pages — site_acl.fast_acl_rules.action.policer_action.ref / 295b66a98131 / 9

- [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--reference--group-001.md#canonical-a98f96081a0cc62b521a0709fb8a65955dc9421b7ce6ded1b811a520a1c09cc1)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-eb83f3dcffca45eecaa610df6066008f3323863624ebb1505429a9d7b701c7c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c210aa1226b0d7d46c811778f065bf4b6cf965837e272faf86fbb9c671f75c6"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action — site_acl.fast_acl_rules.action.protocol_policer_action / 733156b19aaa / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a)
- site_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-37e844396760f8f7e0fb2ac1df3d27406a221f80d54182332755d29464d94ef6"></a>

Type: `"object"`. single nested block, Optional.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
protocol_policer_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac78569ceb8fe9333ce7a55f76373fabb2f54559172acc9ee6ea71d165a2fd11"></a>

## Direct properties — site_acl.fast_acl_rules.action.protocol_policer_action / 733156b19aaa / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-4993b2418266bda750ce27857cfdf0c9ef39575c72f98589576c6fb4d81c006e): complete subsection reference.

<a id="canonical-e5a07b8d7507b6a9e7a38f6cf7121cf390cbc6ecaa0eaa6b05110fefadc264bf"></a>

## Next pages — site_acl.fast_acl_rules.action.protocol_policer_action / 733156b19aaa / 4

- [site_acl.fast_acl_rules.action.protocol_policer_action.ref](resources--fast_acl--reference--group-001.md#canonical-4993b2418266bda750ce27857cfdf0c9ef39575c72f98589576c6fb4d81c006e)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-4993b2418266bda750ce27857cfdf0c9ef39575c72f98589576c6fb4d81c006e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-def039b24f7fd495583b5e0ab98a74953445668124d51ee3cb4bf5b5643a48c2"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action.ref — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.fast_acl_rules.action](resources--fast_acl--reference--group-001.md#canonical-b92caa8920afe675875578d0ec7ab8f6de1e1d5fc60c10622d96e674eb2c536a)
- [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-eb83f3dcffca45eecaa610df6066008f3323863624ebb1505429a9d7b701c7c0)
- site_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-3c06b1af8084a6a0ad704f2376e03568c71cd7d1e3a14f4f979be7bd66eda38b"></a>

Type: `"object"`. list nested block, Optional.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-97559a543cb48babfc0136b9106eeeefb7bfc4d7e45446a43b244f2e1ad4d01d"></a>

## Direct properties — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 3

<a id="canonical-3147bf5d8f591ad4688429864cd8c404efd4b75643059727c18199d2502d28c2"></a>

<a id="canonical-54c761d197c3fa33edcaaf1c08d062840ab6195525f16c5566ac1921080c3b80"></a>

## kind property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 4

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

<a id="canonical-4095e15ade95db4ee358b0b621f48c760bf423b1dde4b0dd1664291ba8e8eae9"></a>

<a id="canonical-409554c99c7c9ac115b6dd2057cfca6b135a4a636c5ca9080c047ef0564a3874"></a>

## name property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 5

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

<a id="canonical-de51f300b0dad6ef55cc0809b13aa5682568a30a412958e440cac40b62aea9c3"></a>

<a id="canonical-f0878095347a02194513d3b3980c034aa0566c86c84edf1e222abd456aba6e8a"></a>

## namespace property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 6

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

<a id="canonical-5b69b43243a89dc34c8a567f6daba66b3c366596fd7de376088560804516671b"></a>

<a id="canonical-ad48b7c82b28df1a9e5d8ab29a57fc8c7bfac13a1ddd650e326be48cd7f2ffee"></a>

## tenant property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 7

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

<a id="canonical-3ede1a9e6d6f34744e2a90bfef8b6942ea584c9269ee125f0e33a5182d3237cd"></a>

<a id="canonical-6c7f0e2796ba53a5b85b32c3a4af2dd25e2532a4304f308b8f7b21abcd60b6c8"></a>

## uid property — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 8

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

<a id="canonical-5c1a48c9139af6829f98bcc999432413717da0f556e70c7937cb5294bc8af99d"></a>

## Next pages — site_acl.fast_acl_rules.action.protocol_policer_action.ref / fa97793688ff / 9

- [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--reference--group-001.md#canonical-eb83f3dcffca45eecaa610df6066008f3323863624ebb1505429a9d7b701c7c0)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-3bdde70d9f86bc4ab339f80e97e279527b39c7a92c58e20b918bb1b3ad81f5c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecd69cc69fa35d945d8bf699367c942df42c142cdc5e33f852067b6b0ee8423b"></a>

## site_acl.fast_acl_rules.ip_prefix_set — site_acl.fast_acl_rules.ip_prefix_set / 617e733f1668 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- site_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-faac95cce9b6dae5418e29f5b71caef9ec1be258ccac321c2e647fbbdd0155b0"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-2088012fbd7bc32f062e528f7d5c6b8fdd309701afc1309d772c80c08facb2f7"></a>

## Direct properties — site_acl.fast_acl_rules.ip_prefix_set / 617e733f1668 / 3

- [ref](resources--fast_acl--reference--group-001.md#canonical-b7d2df8811fc4fb082e70b0d89d9b4c2e38fdebd46865f029d9330bd6bfb0b6c): complete subsection reference.

<a id="canonical-ded3972c572353e8e43c537c8b785fa6adb8a0d72eeb94621a0f5605a96f8fdc"></a>

## Next pages — site_acl.fast_acl_rules.ip_prefix_set / 617e733f1668 / 4

- [site_acl.fast_acl_rules.ip_prefix_set.ref](resources--fast_acl--reference--group-001.md#canonical-b7d2df8811fc4fb082e70b0d89d9b4c2e38fdebd46865f029d9330bd6bfb0b6c)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-b7d2df8811fc4fb082e70b0d89d9b4c2e38fdebd46865f029d9330bd6bfb0b6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d8e1a0ae4ad4a0d8bd27155a474c2944b6c6757648ac9a8f2873ec35627bcfa"></a>

## site_acl.fast_acl_rules.ip_prefix_set.ref — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-3bdde70d9f86bc4ab339f80e97e279527b39c7a92c58e20b918bb1b3ad81f5c4)
- site_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-044d98634fa5e9aaf791e19a2156392e2bfe86de1b19ae74c3a31ed17a6455c6"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1a2fbbb19087e6b8f0ed602e63608508efc18719b3796677a93cae8accb7a38"></a>

## Direct properties — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 3

<a id="canonical-7bbe705f2e45f3b441fec4b090726948f976741177d5ba701f89d89288ee3185"></a>

<a id="canonical-d3b3a1d294db2a25f16d7c407ce4c57b3c6d84f46c99a146e727eb81a103050b"></a>

## kind property — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 4

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

<a id="canonical-7ec01ecd4a678ea76bdf2f2206464f0a1660b248c5d7fa70e32dc0bd23f8d3c1"></a>

<a id="canonical-ddb92fa651ba984de3d649642dc2dac1eaaa4d120cb0560d3a28a77ffec96ea7"></a>

## name property — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 5

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

<a id="canonical-a035ccb19ff1262c29ff00f6988357a298ebb18e478bb71997b8f22edd05e39b"></a>

<a id="canonical-906b958f9d61bb35bc331ec2fa8e9dae2dab17293d22421695178fc05cb769e2"></a>

## namespace property — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 6

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

<a id="canonical-6c936044c0394a4e533718a27a9fa4a76a72d3453fff65d3b90d15d6e7bf51cc"></a>

<a id="canonical-e9d8b2fb5fab67831108c1aeecd67139b4d632bb1b31b4d3aeda50e40a6fa3d1"></a>

## tenant property — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 7

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

<a id="canonical-80f5dd9d415e8476363f5a3dd82beaaa0b22cb1bedd6834549f2677b05edd7ec"></a>

<a id="canonical-12192170050ec79ab68f58bf3a82cdec7cb2664e31bc0242c9087ef77d5abd54"></a>

## uid property — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 8

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

<a id="canonical-9e7aaeba2ad4c8f535385837a542c70d671bc5ecc448943189a76ac6d9c59899"></a>

## Next pages — site_acl.fast_acl_rules.ip_prefix_set.ref / aec6f2d23683 / 9

- [site_acl.fast_acl_rules.ip_prefix_set](resources--fast_acl--reference--group-001.md#canonical-3bdde70d9f86bc4ab339f80e97e279527b39c7a92c58e20b918bb1b3ad81f5c4)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-91b94cbe58997c6d4e6136e25b06501c2a3d4ea0fc3551d9a445e3255d6b13ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5a7a44574ac9fe0b50fd54fedfffe7f43487d1d2cb52c018fe022a2d07fb9fc"></a>

## site_acl.fast_acl_rules.metadata — site_acl.fast_acl_rules.metadata / 5c251839ee42 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- site_acl.fast_acl_rules.metadata

<a id="canonical-b7eb1be7324e442e15802d6beab8522440e4cf718051468c030d0784f58979a3"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d83b0b1c6bb4aa901e0e4bd1be54525251fa3d3f5781fe86caa635f069818b3"></a>

## Direct properties — site_acl.fast_acl_rules.metadata / 5c251839ee42 / 3

<a id="canonical-840403d784eb25775000aaa727e01c9acceeb18afe710278d7d590d8ea80c842"></a>

<a id="canonical-f10259ff85f6ce6322e4c539a0a8a4ea5240b06e4b37799b693cec4550c3adb1"></a>

## description_spec property — site_acl.fast_acl_rules.metadata / 5c251839ee42 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-a7178d9ec1b8957426d0a6f0bee5f478babd5d975829801efe4946308d338223"></a>

<a id="canonical-a9a4d5d6aa597e1bbc3e548a30162ffb468227fd1bc2e587c5832e72a5f32406"></a>

## name property — site_acl.fast_acl_rules.metadata / 5c251839ee42 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-9a87783c75898b5bb3deb1aaf6dcba68d7c2532ad716116bb31c757a7b3f9f83"></a>

## Next pages — site_acl.fast_acl_rules.metadata / 5c251839ee42 / 6

- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-95d80e88d1961822b37e11b50d5773c35cc39b5fb71cd6c3b0797c744f2ec242"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-993feb3a2120af00639779702f94756bb8f431538e2095779f8774367351324e"></a>

## site_acl.fast_acl_rules.port — site_acl.fast_acl_rules.port / ec174e89da5c / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- site_acl.fast_acl_rules.port

<a id="canonical-65933cae7afd531ad93bd80fc8003fdc5e9db17a802ad74365739f78b7a4a781"></a>

Type: `"object"`. list nested block, Optional.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all",
    "dns"),
  validators.ConflictingListObjectAttributes("all",
    "user_defined"),
  validators.ConflictingListObjectAttributes("dns",
    "user_defined")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1bdb22e1dfbe46e6d208001f4598cd3ac9f21b174a42a88e9acecde8aab88cb1"></a>

## Direct properties — site_acl.fast_acl_rules.port / ec174e89da5c / 3

- [all](resources--fast_acl--reference--group-001.md#canonical-110f7ebc1605ac37f9cddb08ce9441c8874fc80135d8bea4b0047bb849f60d32): complete subsection reference.

- [dns](resources--fast_acl--reference--group-001.md#canonical-721aae1186531f3f5a96b3109ab92b395397cd77fb2a8ad516abacec650fec21): complete subsection reference.

<a id="canonical-4f1e676c44d66b52b406c5db4546b2876a2d6a6ea6b65b45a580160ce4d414d1"></a>

<a id="canonical-6e3248887a70dfeb48356fd1fb30e6943bb5b7b56fa98b5137641a37a29ecb2f"></a>

## user_defined property — site_acl.fast_acl_rules.port / ec174e89da5c / 4

Type: `"number"`. Optional.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

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

<a id="canonical-81e4207b794caf0b031e4bbeb27289d17ff0408e010f36eac5c70ca412fa48f0"></a>

## Next pages — site_acl.fast_acl_rules.port / ec174e89da5c / 5

- [site_acl.fast_acl_rules.port.all](resources--fast_acl--reference--group-001.md#canonical-110f7ebc1605ac37f9cddb08ce9441c8874fc80135d8bea4b0047bb849f60d32)
- [site_acl.fast_acl_rules.port.dns](resources--fast_acl--reference--group-001.md#canonical-721aae1186531f3f5a96b3109ab92b395397cd77fb2a8ad516abacec650fec21)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-110f7ebc1605ac37f9cddb08ce9441c8874fc80135d8bea4b0047bb849f60d32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57de99fe4aa86acdbb9bffd3a0c24cf610038bd16223b3bbd8619f518c68167b"></a>

## site_acl.fast_acl_rules.port.all — site_acl.fast_acl_rules.port.all / 3e755fad9c8f / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-95d80e88d1961822b37e11b50d5773c35cc39b5fb71cd6c3b0797c744f2ec242)
- site_acl.fast_acl_rules.port.all

<a id="canonical-f81da70f841f6f35a72ff9301c82613fa5d6302ab498cad736de8223f221b0a8"></a>

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
all = {}
```

<a id="canonical-84a2c8c776eaf7217358149369bece47842b1755b70f250c5b0e9ff3280b207d"></a>

## Direct properties — site_acl.fast_acl_rules.port.all / 3e755fad9c8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-69bad7aebb7a8640e5b15b1bfb2722eccfe8b3f8e9093bef0e1ed01dab023ca4"></a>

## Next pages — site_acl.fast_acl_rules.port.all / 3e755fad9c8f / 4

- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-95d80e88d1961822b37e11b50d5773c35cc39b5fb71cd6c3b0797c744f2ec242)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-721aae1186531f3f5a96b3109ab92b395397cd77fb2a8ad516abacec650fec21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf7dd90a94eb9ef3e2bc51a64a572c0ed7bbcdf0d6c9406bb46a5fd759a34286"></a>

## site_acl.fast_acl_rules.port.dns — site_acl.fast_acl_rules.port.dns / 658c336744a8 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-95d80e88d1961822b37e11b50d5773c35cc39b5fb71cd6c3b0797c744f2ec242)
- site_acl.fast_acl_rules.port.dns

<a id="canonical-606f38503ea99fc47d72c86ca96fddb851109c7ac6be03639e57311e8604142e"></a>

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

<a id="canonical-b8cf10e84f3e1b85e75f2a7e257dd871ca05ea95b8ad02379e5469bde49e71fe"></a>

## Direct properties — site_acl.fast_acl_rules.port.dns / 658c336744a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d6d2a9069b6dc41ba772ba4691b6b7d19135c7976f1aaf39ba17d5d1f842c6e"></a>

## Next pages — site_acl.fast_acl_rules.port.dns / 658c336744a8 / 4

- [site_acl.fast_acl_rules.port](resources--fast_acl--reference--group-001.md#canonical-95d80e88d1961822b37e11b50d5773c35cc39b5fb71cd6c3b0797c744f2ec242)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-8a4c50d1b3c1f23a5d56b15a28c7570a7fd37b5a76b2d8df6a07c35e9bc423c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f71c20b7ca5f8c0d4ee14464095cbad831bf547b9b57bf6d64ecefd1477be861"></a>

## site_acl.fast_acl_rules.prefix — site_acl.fast_acl_rules.prefix / 4971db9cdafe / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- site_acl.fast_acl_rules.prefix

<a id="canonical-8b2dce192e2535c2eaa06854edc0795272dd0417bfc0c2de4274b5081b63e201"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
prefix {
  # Configure direct properties listed below.
}
```

<a id="canonical-d2b6bf3ad18f88e3443c5d312113b481394e6fb6475fff0988febf766ee1fe0d"></a>

## Direct properties — site_acl.fast_acl_rules.prefix / 4971db9cdafe / 3

<a id="canonical-115b6474eaa43c1c720eba2d7010fa739741ed9409a33724d345ba315d7b4356"></a>

<a id="canonical-470f359e33aad0d175e646c0e53707681d25c8ba20988d270bf7feada66e1c2a"></a>

## prefix property — site_acl.fast_acl_rules.prefix / 4971db9cdafe / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-f295c3f733f7d092eb81692189ffa7874450728390039bb83363159d1522f632"></a>

## Next pages — site_acl.fast_acl_rules.prefix / 4971db9cdafe / 5

- [site_acl.fast_acl_rules](resources--fast_acl--reference--group-001.md#canonical-8fbbab894d096682c7bf391943f8d7d3dfd26e4e775b522634e436b6fc4ed7c9)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-455e5c0734aa4e53d8893ae4b23530caf37f03dec57285620bdac82f1c9a8926"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f77b11271a4a0e164be420832ac28f19c9b7c649e00dc88c743cdb6c7bf172b"></a>

## site_acl.inside_network — site_acl.inside_network / aa153e03dfa6 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- site_acl.inside_network

<a id="canonical-411db025c49565219504b2727942602f20358f666dae900d1967fe089fd88b79"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-63f6bf7d7c3d2b9a401e6e644af31741359da76ebb7ca4359dbb466ab3f23875"></a>

## Direct properties — site_acl.inside_network / aa153e03dfa6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7cf1ac59bc4e00478deb543ea1ff5c6fce11695c0385f54d2ccf7c1ea0c13c4"></a>

## Next pages — site_acl.inside_network / aa153e03dfa6 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-7980ac9d18076f9d21ec150822f9b762557552b195b174490fe718216b8de7fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c52232566e27483ea1b3236825028b17e1f1ccaefb9010d1a770fccb8407e70"></a>

## site_acl.interface_services — site_acl.interface_services / d90623c800fc / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- site_acl.interface_services

<a id="canonical-fab7c5d4a30e915498f87762d9c93a8c231c1706fa60b68aa9a61698589c4def"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for interface services.

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
interface_services = {}
```

<a id="canonical-b22556de1c174a525bafde5b574db46e7b01503ea61c53ce2498e7f2b397b34d"></a>

## Direct properties — site_acl.interface_services / d90623c800fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ecc46fcfd872336ecedc38f086e701b222586aa3e01c3f9a2f41d64fda82f7f7"></a>

## Next pages — site_acl.interface_services / d90623c800fc / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-10098691a31f88a74498999bb34c120c806e5c433fd33ac8eb85fc9707780bd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dac7fd74c80744f928deeceae470fb8036c89caf991b39f7c4ffc845b7424e4d"></a>

## site_acl.outside_network — site_acl.outside_network / 7dc8eb56c2b3 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- site_acl.outside_network

<a id="canonical-41711b1ad57ddab3aef2f545683331114c8d0ef4182e28b3a1e8033ac2beb6d6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-e8440d9d703b4ec04b8ea10a3dbf5ec5677fc30351f850abe2ea45ad88fbe99e"></a>

## Direct properties — site_acl.outside_network / 7dc8eb56c2b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30a7734e9eb7a7035caac328741886f35e9052965b9104d23e9251f0d490db72"></a>

## Next pages — site_acl.outside_network / 7dc8eb56c2b3 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-59d17558c6ec64d80167a83cc56d3e3e9178c42f258a640cbd1c902dbf32a851"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd4055b26af95f2b1bc31f8db13b1ad0dc3c5d0d108750decc12b4ec125b8b48"></a>

## site_acl.vip_services — site_acl.vip_services / 9c9438ed5650 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- site_acl.vip_services

<a id="canonical-40aa9741213119e8a7ea059cac6013255579c8706ac9924d42a9b3ea9f1f07c9"></a>

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
vip_services = {}
```

<a id="canonical-7dbcb42ddfadfe6655fa21a84e650b981b837c06c1a9bb3b1623e1bb16f6ef8d"></a>

## Direct properties — site_acl.vip_services / 9c9438ed5650 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0cd6d20f3c6c3bc4197cb44e3b8c85fa12aa14f628c849e7ed22b35c3ee77b3"></a>

## Next pages — site_acl.vip_services / 9c9438ed5650 / 4

- [site_acl](resources--fast_acl--reference--group-001.md#canonical-219fcf5656603ab98159941813d7e0ac211171be15e6a225e520df52128420fa)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-aa8259fa24db8bdfd1ab7042e02d60c3ad6a53370cad9578f959eeec65409c35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59fecc69318b71d7621ab06bb451018986e0f286c9e4ba160eb34ff55a7d3827"></a>

## timeouts — timeouts / e11cd67a6471 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- timeouts

<a id="canonical-974f8b7bb7dcb969d97122fe835658d9a89a4c9e90d58eb6ac49152862b4292d"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-c62fbf855b307cb0e87d2e4a1e0e70e034ff446942af04e81b13a5489ba0bcb4"></a>

## Direct properties — timeouts / e11cd67a6471 / 3

<a id="canonical-db5dec63902c5d4209fa0e3c7757f090296d06a6abca9fe1d5a96ff442653529"></a>

<a id="canonical-dd13d1d88ab5f52b1ceacb8b410d50a44920933161ff083119fae4266d82a54d"></a>

## create property — timeouts / e11cd67a6471 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-495abfdc5c92ebefab18ce74c75538767d2ae7bbbd7028b15ccbf47dbe53572a"></a>

<a id="canonical-30df69edf6177740587221893ffbc53406f52713c720b2bfcbc71cf67ecc4c90"></a>

## delete property — timeouts / e11cd67a6471 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-fb0f413f6a8bd516262f86f02c738724a78edb5e454d373879931e623fd18270"></a>

<a id="canonical-52102481939a0962b8777d5a9e52ecdb450409a499129ba27713704294b33f68"></a>

## read property — timeouts / e11cd67a6471 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-16401a93b3f55b48fc5319f93479011de8e8675cff813cef3a1f90659504fd6f"></a>

<a id="canonical-6f86a9400c6f64d11881a6a7995af5faa276ad8e13f9533c8f68fe8e1c6a7861"></a>

## update property — timeouts / e11cd67a6471 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-42b5d5813d921956e7e45fb2ccb0bbd3dfc9d79395b29983b9cc7607ec4c4971"></a>

## Next pages — timeouts / e11cd67a6471 / 8

- [Property reference](resources--fast_acl--reference--group-001.md#canonical-e17ee761bab5ee394fe0dbef08b5ba16b640f983545d13a5adb24c265eb3c9b6)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
