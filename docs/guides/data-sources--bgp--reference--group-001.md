---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1ffd5c38fadb9d3bcc8bd6830fbb8ded8f146ab4140baa4adda4f42cb6e9c1a"></a>

## Property reference — Property reference / 5c6a32aa2706 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- Property reference

<a id="canonical-7798a14f3ed87eff0159f18535ede5d7bc8501d960d6173eca0f96f4bd2bb6f0"></a>

## Direct properties — Property reference / 5c6a32aa2706 / 3

<a id="canonical-3f2f311261a8302323e0f81431ddfab8ee75c99b9d10065525a31f47c2e8b416"></a>

<a id="canonical-f321366ed995c1f8bdb5649685f9509404d949d8925b41a283775d777a2b94ea"></a>

## annotations property — Property reference / 5c6a32aa2706 / 4

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

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-7f432c686742ce5c6b276a53d92c6c021bab0da00649acb0c5a00d3656909885): complete subsection reference.

<a id="canonical-13aaae4992149e9d5230172fff4e973855ee8c7a1a78aeb74bce9fe6e1a5ba84"></a>

<a id="canonical-145aa5c4815221914471da0fca36722b852c80bd6ffe4ba54dd7a77c3976dff9"></a>

## description property — Property reference / 5c6a32aa2706 / 5

Type: `"string"`. Computed.

Description of the BGP.

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

<a id="canonical-1908edf9c8ab4def9b323a84e65535cbca334dc5b2fdbf4b99548420eb3fc656"></a>

<a id="canonical-8036c09617abf25244f9a0945ccf59f347f5717bc1b8052edc2f791db66ac195"></a>

## id property — Property reference / 5c6a32aa2706 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5cca288635763d294eac149ff4153e7683b0edfddd8c8403d08ce9ddb0cd915b"></a>

<a id="canonical-842274512723eed73e56ed2585a6d932d107683992d97a17659e822a80b92801"></a>

## labels property — Property reference / 5c6a32aa2706 / 7

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

<a id="canonical-53f6cd9899e8ab4fca907d0d960e6f8ccb00b88a8fe887464042ecbafc400c25"></a>

<a id="canonical-158372295e22f2f4f20a234f434c1898b29579666e50e70940ed1b340b6f19ed"></a>

## name property — Property reference / 5c6a32aa2706 / 8

Type: `"string"`. Required.

Name of the BGP.

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

<a id="canonical-5e842217b1bab46984447ffe241513b6dd5a83cb7c6c56721cd2fed0eeca7a39"></a>

<a id="canonical-d6abc4d66e480115e982df18954867ddf1eb2055486881756651c9dc03ec3e40"></a>

## namespace property — Property reference / 5c6a32aa2706 / 9

Type: `"string"`. Required.

Namespace where the BGP exists.

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

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3): complete subsection reference.

- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098): complete subsection reference.

<a id="canonical-62a235cecba9f61e559f2d96ee42c28b8aefe04affa0aeaab2fce9c29997c476"></a>

## All schema paths — Property reference / 5c6a32aa2706 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp--reference--group-001.md#canonical-3f2f311261a8302323e0f81431ddfab8ee75c99b9d10065525a31f47c2e8b416) |
| `bgp_parameters` | [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-7d4a4041effc6ba49f9fae301ac7ac663d6159958095a71d609c0b0ab64e162a) |
| `bgp_parameters.asn` | [bgp_parameters.asn](data-sources--bgp--reference--group-001.md#canonical-337c277127c4df9d4da7d94c47138a17cd66ad9c96fc1d6b5c88aba88910acb9) |
| `bgp_parameters.from_site` | [bgp_parameters.from_site](data-sources--bgp--reference--group-001.md#canonical-e8ada238e545d64365567e4904b89c62e1255c620fbe8e92ce104c94b44d0d96) |
| `bgp_parameters.ip_address` | [bgp_parameters.ip_address](data-sources--bgp--reference--group-001.md#canonical-6d4f722eda4f76ca42277caff44e10eb33f9f76545a122d358aeb6903a1ada94) |
| `bgp_parameters.local_address` | [bgp_parameters.local_address](data-sources--bgp--reference--group-001.md#canonical-c3d638b59ffd9676e1afd3b38606044823599f1cb069ebb60bbb9e7538d72a6b) |
| `description` | [description](data-sources--bgp--reference--group-001.md#canonical-13aaae4992149e9d5230172fff4e973855ee8c7a1a78aeb74bce9fe6e1a5ba84) |
| `id` | [id](data-sources--bgp--reference--group-001.md#canonical-1908edf9c8ab4def9b323a84e65535cbca334dc5b2fdbf4b99548420eb3fc656) |
| `labels` | [labels](data-sources--bgp--reference--group-001.md#canonical-5cca288635763d294eac149ff4153e7683b0edfddd8c8403d08ce9ddb0cd915b) |
| `name` | [name](data-sources--bgp--reference--group-001.md#canonical-53f6cd9899e8ab4fca907d0d960e6f8ccb00b88a8fe887464042ecbafc400c25) |
| `namespace` | [namespace](data-sources--bgp--reference--group-001.md#canonical-5e842217b1bab46984447ffe241513b6dd5a83cb7c6c56721cd2fed0eeca7a39) |
| `peers` | [peers](data-sources--bgp--reference--group-001.md#canonical-13e0c063a45f2dfa92863ce1e8697ae5787bcc92a6a44bf9ccb2b593488e022d) |
| `peers.bfd_disabled` | [peers.bfd_disabled](data-sources--bgp--reference--group-001.md#canonical-30a5425146c75639a31d4927cb128df9c9d6a20d54c062275dedd9c253f6719b) |
| `peers.bfd_enabled` | [peers.bfd_enabled](data-sources--bgp--reference--group-001.md#canonical-c00e4e7e959f127f22268399f29bd4040c8954d4f8f34e02eb8e79bd97d79052) |
| `peers.bfd_enabled.multiplier` | [peers.bfd_enabled.multiplier](data-sources--bgp--reference--group-001.md#canonical-38c55825c2de0b92da06018c0664545bae43b365eacb64e0a6334630cead59d8) |
| `peers.bfd_enabled.receive_interval_milliseconds` | [peers.bfd_enabled.receive_interval_milliseconds](data-sources--bgp--reference--group-001.md#canonical-ffc4a5156e11bb6de7195118eb0e171a93c9070a244281842d73fc1772e05513) |
| `peers.bfd_enabled.transmit_interval_milliseconds` | [peers.bfd_enabled.transmit_interval_milliseconds](data-sources--bgp--reference--group-001.md#canonical-d65b74e8103fe6794506a9c9f8188f19a3abb288815fe0a6e9aa173c93a176a0) |
| `peers.disable_spec` | [peers.disable_spec](data-sources--bgp--reference--group-001.md#canonical-232540b4a1df3b9d6aa0a2612d2106a440bf02312126de65da1616bcaa33ab18) |
| `peers.ebgp_multihop_disabled` | [peers.ebgp_multihop_disabled](data-sources--bgp--reference--group-001.md#canonical-87e8ed37a0aae1696126e5cda70c12f2fc1f19ba4315351f3110c84fa0968b89) |
| `peers.ebgp_multihop_enabled` | [peers.ebgp_multihop_enabled](data-sources--bgp--reference--group-001.md#canonical-bd3c01a3867b1fb4a6aaca5d12d711b938a6798fe2cf57d8fdd3c05bd51d247a) |
| `peers.external` | [peers.external](data-sources--bgp--reference--group-001.md#canonical-f8b4c6b6ec8f6990ff8de56af59d30498a8db1a345f680ed843682b978e04ac3) |
| `peers.external.address` | [peers.external.address](data-sources--bgp--reference--group-001.md#canonical-59a2cb0be716462d1213071633771f886cb164915e500e43d17b11b2515a7279) |
| `peers.external.address_ipv6` | [peers.external.address_ipv6](data-sources--bgp--reference--group-001.md#canonical-52a2dfa1f88cd3d34d87446c1d50d6718057e6820eb840e479ceb172b6bb814e) |
| `peers.external.asn` | [peers.external.asn](data-sources--bgp--reference--group-001.md#canonical-49aaf06ee83db0cc5aa6211f7b37b3cc626b1fb62a861d0675f0663c886cf73d) |
| `peers.external.default_gateway` | [peers.external.default_gateway](data-sources--bgp--reference--group-001.md#canonical-09675cc0c43f9623041317631684ca3754a092b3232b0ce55501c86d9d3eaa22) |
| `peers.external.default_gateway_v6` | [peers.external.default_gateway_v6](data-sources--bgp--reference--group-001.md#canonical-2c62b6e810d7fe285fe6c5ac6ce044e98fbfe6794d457f45abb08e40d5004100) |
| `peers.external.disable_spec` | [peers.external.disable_spec](data-sources--bgp--reference--group-001.md#canonical-dff2dde2d05a41e191feff25103a81ae801f7f8e20e2ec0a51b933a4b1363465) |
| `peers.external.disable_v6` | [peers.external.disable_v6](data-sources--bgp--reference--group-001.md#canonical-39ac65780d9f06f751ed4bcef354864b4e2f1d9da463c313e7c981c0aa1eaf76) |
| `peers.external.external_connector` | [peers.external.external_connector](data-sources--bgp--reference--group-001.md#canonical-9068999fb414ed5912cff88b76507c72ea0f7cf69d3c5264ab08b8fcacf8940a) |
| `peers.external.family_inet` | [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-808d2a579129014b2399fb6655a46dbc66c909ed635478181115e445b825c87b) |
| `peers.external.family_inet.disable_spec` | [peers.external.family_inet.disable_spec](data-sources--bgp--reference--group-001.md#canonical-ea7a6bca20bbd829e717326f060419fcf9e21fdda3bdb094b9bee4bc7efa4735) |
| `peers.external.family_inet.enable` | [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-45be3d5388741d132473b33de8f2009b699fb83b20b7178ba6a88e76f132f1d0) |
| `peers.external.family_inet.enable.aggregation` | [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-0e8c65874da323ff74a828ec9d31cd58e8c8a2fdd092305b1ad7c602d957de18) |
| `peers.external.family_inet.enable.aggregation.ip_prefix` | [peers.external.family_inet.enable.aggregation.ip_prefix](data-sources--bgp--reference--group-001.md#canonical-97471307a0b44e8f73d16b0f76a317b8a04d487e02d425ff0f393151d54b0a15) |
| `peers.external.family_inet.enable.aggregation.options` | [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-2365903b12359d093a5a5a4b93fdd9703c28f6f374baad804996818b4d182456) |
| `peers.external.family_inet.enable.aggregation.options.summary_only` | [peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--bgp--reference--group-001.md#canonical-3c34517c52c85250dfaa47f249f107ecf2a68a10f8021f4fac05657de7bd244f) |
| `peers.external.from_site` | [peers.external.from_site](data-sources--bgp--reference--group-001.md#canonical-82f71c45038132c92ddafd7623124ee65bb7808ebd06bf923122ff6cf1c8535e) |
| `peers.external.from_site_v6` | [peers.external.from_site_v6](data-sources--bgp--reference--group-001.md#canonical-9221174021c97a00e575e082cec9683934cacf3d60f4fbb7e8f9b8bd8a322e25) |
| `peers.external.interface` | [peers.external.interface](data-sources--bgp--reference--group-001.md#canonical-35b6ed865f2da1dcbb530274fd7099bc31fb43ed4f33f24ce276a42ff5432d0f) |
| `peers.external.interface.name` | [peers.external.interface.name](data-sources--bgp--reference--group-001.md#canonical-bde97a8cdb534da663e16732231c5a5f434322b583583740365d4ec7d8103174) |
| `peers.external.interface.namespace` | [peers.external.interface.namespace](data-sources--bgp--reference--group-001.md#canonical-e78944a729a04f0001f6f3e8f487135300e67b8badfcf166e4c598161f620421) |
| `peers.external.interface.tenant` | [peers.external.interface.tenant](data-sources--bgp--reference--group-001.md#canonical-84ca59d3c70a9201fe5b6d5c1cca09e5f7c04e06f9d891c3c8e32f16da638465) |
| `peers.external.interface_list` | [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-5d0d23e53fa4ea2b0b80eb5087c16c5c0dcc85dd84e9e7f0468332fddd486fb0) |
| `peers.external.interface_list.interfaces` | [peers.external.interface_list.interfaces](data-sources--bgp--reference--group-001.md#canonical-14a9845636f942336f06d9768501f3e8608e1b651f10751755902c07dfbf0af8) |
| `peers.external.interface_list.interfaces.name` | [peers.external.interface_list.interfaces.name](data-sources--bgp--reference--group-001.md#canonical-4436209ff78d072a3c84749fcdfabbe2de94852764f987072eeb6241bb256df3) |
| `peers.external.interface_list.interfaces.namespace` | [peers.external.interface_list.interfaces.namespace](data-sources--bgp--reference--group-001.md#canonical-bb5d238ea92107f41f10255c6ac2aca6e16e6431bef63017c7eae7b90873e5e6) |
| `peers.external.interface_list.interfaces.tenant` | [peers.external.interface_list.interfaces.tenant](data-sources--bgp--reference--group-001.md#canonical-a6d845dfff304850b7d2fe2a93027112458eab663f284492e9ad0f4fc510abaa) |
| `peers.external.md5_auth_key` | [peers.external.md5_auth_key](data-sources--bgp--reference--group-001.md#canonical-4aed42603c600cc1463ae045ca7d79cc12e6c1ad62ae5f48060b8919cc426bb3) |
| `peers.external.no_authentication` | [peers.external.no_authentication](data-sources--bgp--reference--group-001.md#canonical-c48f05cb6d10ef4ce7ccbd26b6a2adbe862c757045b5ec4025acdb725b34ce5a) |
| `peers.external.port` | [peers.external.port](data-sources--bgp--reference--group-001.md#canonical-773e89e0492c476052c171577ff7cbf02f87dd4f3a6221e02ea70f315e371eff) |
| `peers.external.subnet_begin_offset` | [peers.external.subnet_begin_offset](data-sources--bgp--reference--group-001.md#canonical-bd6101e90b2bf252742d7b939e50d58aca0a773c37ac08269c9a0dc67df47421) |
| `peers.external.subnet_begin_offset_v6` | [peers.external.subnet_begin_offset_v6](data-sources--bgp--reference--group-001.md#canonical-d9395bef39286bd0448758e41bd895853d6ea622b711833dbc0101792147e313) |
| `peers.external.subnet_end_offset` | [peers.external.subnet_end_offset](data-sources--bgp--reference--group-001.md#canonical-ded9c00013201cc4b688801a6477aae71772815533b342d524fc0040db4fb49d) |
| `peers.external.subnet_end_offset_v6` | [peers.external.subnet_end_offset_v6](data-sources--bgp--reference--group-001.md#canonical-92cd88aa516a0c3db977a4f035c26dd4f352f99cec79eba941e2e36d7c9c9075) |
| `peers.label` | [peers.label](data-sources--bgp--reference--group-001.md#canonical-1b67a4c82b231c3ce56de9e70aecb6682fb1ba23c12baf486453f6ea55d7d0ce) |
| `peers.metadata` | [peers.metadata](data-sources--bgp--reference--group-001.md#canonical-3b52e1a8bb251df25d409d07b0540b2e28fefd35a47cbd04ba048e4f80c81902) |
| `peers.metadata.description_spec` | [peers.metadata.description_spec](data-sources--bgp--reference--group-001.md#canonical-abd4b127d9124dc88ca38b856079c9ce3ade6fdb3869e6a37a4d84644d36d62e) |
| `peers.metadata.name` | [peers.metadata.name](data-sources--bgp--reference--group-001.md#canonical-b21f5ba85ef904b4fbaf85c080d0b0523071637f891f4bc71cd5f721734ba4e4) |
| `peers.passive_mode_disabled` | [peers.passive_mode_disabled](data-sources--bgp--reference--group-001.md#canonical-83cde329c949737ae1baca652e39f9388294e308c8899cb83e99b7a21707966a) |
| `peers.passive_mode_enabled` | [peers.passive_mode_enabled](data-sources--bgp--reference--group-001.md#canonical-58d278b7937230dd68f1f1d0230b48843249ddd995ec633be632a386645a9b54) |
| `peers.routing_policies` | [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-33e8b65f508c2881286868cfd7761e7a23792e6aee67e33bcda7a154b3407cd2) |
| `peers.routing_policies.route_policy` | [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-66a858042524b39734cfbb1d8a7d970994d3abf06f7b3e82e18b20d43d7dddb8) |
| `peers.routing_policies.route_policy.all_nodes` | [peers.routing_policies.route_policy.all_nodes](data-sources--bgp--reference--group-001.md#canonical-755a8832117bc12df223d705c096a118dca965c436edd721cfd1dfd908fb01dd) |
| `peers.routing_policies.route_policy.inbound` | [peers.routing_policies.route_policy.inbound](data-sources--bgp--reference--group-001.md#canonical-7b813f25671252a4d46e159e9f570abed70b4179991d152249a6f6b63dafd2ed) |
| `peers.routing_policies.route_policy.node_name` | [peers.routing_policies.route_policy.node_name](data-sources--bgp--reference--group-001.md#canonical-a11df55948844dc9cf79361349a08b24aecf954c7bc741661bb779fa81ef3935) |
| `peers.routing_policies.route_policy.node_name.node` | [peers.routing_policies.route_policy.node_name.node](data-sources--bgp--reference--group-001.md#canonical-93249c82950374e6aabf7dbf88abca01a727e58f0ae4f2a0621dc4dd8cff9c84) |
| `peers.routing_policies.route_policy.object_refs` | [peers.routing_policies.route_policy.object_refs](data-sources--bgp--reference--group-001.md#canonical-951b1df316b6c4102b3c196e0a6f588b3ef3aa5b8f75a5d80011765fea6f723e) |
| `peers.routing_policies.route_policy.object_refs.kind` | [peers.routing_policies.route_policy.object_refs.kind](data-sources--bgp--reference--group-001.md#canonical-22f8363d572d2d76139b365d46d37e224578b1a8980e46827a490675ca14a4d8) |
| `peers.routing_policies.route_policy.object_refs.name` | [peers.routing_policies.route_policy.object_refs.name](data-sources--bgp--reference--group-001.md#canonical-f1b1f577e14993e2e1bf9e3175d1858bdd2cf1fdbc2540ca91b167e811f6f394) |
| `peers.routing_policies.route_policy.object_refs.namespace` | [peers.routing_policies.route_policy.object_refs.namespace](data-sources--bgp--reference--group-001.md#canonical-2628e6cf6fd2fc31185a794ed8a743d8746868bf81d81462be648085cea8b00d) |
| `peers.routing_policies.route_policy.object_refs.tenant` | [peers.routing_policies.route_policy.object_refs.tenant](data-sources--bgp--reference--group-001.md#canonical-083dcbef2951b77165eab30f5d37b77de27776b444344530d05715575169dc94) |
| `peers.routing_policies.route_policy.object_refs.uid` | [peers.routing_policies.route_policy.object_refs.uid](data-sources--bgp--reference--group-001.md#canonical-6317dda7a094f8ad344e3e2a407c192fe86bef601a6f496ec87b9f0e6cf4cd92) |
| `peers.routing_policies.route_policy.outbound` | [peers.routing_policies.route_policy.outbound](data-sources--bgp--reference--group-001.md#canonical-ac765d76f43b33149faa6fcdfaa35b1e2f86793a67dcff4eb74210d2d95ff042) |
| `where` | [where](data-sources--bgp--reference--group-001.md#canonical-4e81baa6eeb88c0b26bcf084623f147bdef01637bbd3c8efc1a0458bcd50b3ca) |
| `where.site` | [where.site](data-sources--bgp--reference--group-001.md#canonical-1c40d4c4134206c17fa5e909475f2479952b34335dac1cdef07f164081190ad4) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-9572689a8469ed674c977a77482ed2bd3ba505f01041779a3cf9c6be9d350fea) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-6b6d64fe971304d6394fdfed4028bc0e0ef960f62117a9a263af1b996b4df0d1) |
| `where.site.network_type` | [where.site.network_type](data-sources--bgp--reference--group-001.md#canonical-0784c2c7c7638af373b751e726ade2791ca889e63b116bc712bdb18504dcbdf2) |
| `where.site.ref` | [where.site.ref](data-sources--bgp--reference--group-001.md#canonical-e7b0ebcf7d759e1f246545d85eeab38aaad301159e3dda756f2c99dec43aa370) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--bgp--reference--group-001.md#canonical-f4e4133a175d42a52f96d45887a55ea353ce086cc1b26c93e98d122d5a03e0f5) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--bgp--reference--group-001.md#canonical-65631bd93a688302e4c6ba40e408f8203fa19cd6cb6343f81592a0e8b15db01f) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--bgp--reference--group-001.md#canonical-54d08e1f4b621ac6240a2ccdcc38a8c995c4135a7db9f7ffafdfe5b7e0132e36) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--bgp--reference--group-001.md#canonical-d37937503b74b00b58a187f7193ba4b453177f398859e80bd67f7c1e951f3e08) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--bgp--reference--group-001.md#canonical-df0be0ae3a96c184df2df83b538eeabeeb9546ed90af3221e612a145be1eebbb) |
| `where.virtual_site` | [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-f84546fa45d82bbed6e9b76d31aff132327e692acf8cb9534c8e566d37d114f9) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-dadf99503d102c067a09197240f09b0968ea9e5bb68c5d4bc813c88d6ff0c841) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-a5dafce89d557981b643d2e5396232b66db5ec80f61e0a84f8cb54501821b266) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--bgp--reference--group-001.md#canonical-92768af09156e2ad40ace1e5435dbe815d54c5c079ff4f43c6a5a4fd9384c2c5) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--bgp--reference--group-001.md#canonical-14eea49c909a5aa0c7d046688186d3839bd3c892b0428fce86db9784cbf69d46) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--bgp--reference--group-001.md#canonical-d53303aabe90c27a49d068782db83205c635b739395539f2c061cd77eb00a795) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--bgp--reference--group-001.md#canonical-77072e8dfa40eef10d328a2c22fef2647b5b6ceb2d1de28ab24178b6b6038926) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--bgp--reference--group-001.md#canonical-78660e8d03b6d04ded3a2c6cfd87a56dea65f88f039d844e86e81b81e1f82986) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--bgp--reference--group-001.md#canonical-0fcbdb094432f433ceedeccff8bf303b211339e4fdf2c97ab6f87311064ff511) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--bgp--reference--group-001.md#canonical-09df79cc9194e00a714c9db505856cf503c67267354bf75907c49dfae4b55547) |

<a id="canonical-0dec3a150693e4fb77b08ec712712064045404f58099629828f162729391428d"></a>

## Next pages — Property reference / 5c6a32aa2706 / 11

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-7f432c686742ce5c6b276a53d92c6c021bab0da00649acb0c5a00d3656909885)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-7f432c686742ce5c6b276a53d92c6c021bab0da00649acb0c5a00d3656909885"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb3cfc3998f4f95ee192b5f89dea1b67ab9359494d0e0a63c5bec9853e832d22"></a>

## bgp_parameters — bgp_parameters / b80a57157cbb / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- bgp_parameters

<a id="canonical-7d4a4041effc6ba49f9fae301ac7ac663d6159958095a71d609c0b0ab64e162a"></a>

Type: `"single"`. Computed.

Configuration parameter for bgp parameters.

Upstream description:

BGP parameters for the local site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-router_id_choice": "[\"from_site\",\"ip_address\",\"local_address\"]"
}
```

<a id="canonical-2d0fce62776f8b171517104577937556832353e373622b5d6bd83a9c0a77df4a"></a>

## Direct properties — bgp_parameters / b80a57157cbb / 3

<a id="canonical-337c277127c4df9d4da7d94c47138a17cd66ad9c96fc1d6b5c88aba88910acb9"></a>

<a id="canonical-f69796a88012739b16055c0cd16a0fedd6d93ff8fc7e632822a13825dc868937"></a>

## asn property — bgp_parameters / b80a57157cbb / 4

Type: `"number"`. Computed.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [from_site](data-sources--bgp--reference--group-001.md#canonical-0ad063296715970f438302b0de8e53652e57261c59fa05f6c4363974a54935d1): complete subsection reference.

<a id="canonical-6d4f722eda4f76ca42277caff44e10eb33f9f76545a122d358aeb6903a1ada94"></a>

<a id="canonical-2c66c85b9b724101c77c820ecc6bafce65d2b5547a19fdcf7faa312a5c711a0c"></a>

## ip_address property — bgp_parameters / b80a57157cbb / 5

Type: `"string"`. Computed.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Upstream description:

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

- [local_address](data-sources--bgp--reference--group-001.md#canonical-d51790f759824fb5b933d0ab88894a1a5e14c1540530296d511b9c3b196a3092): complete subsection reference.

<a id="canonical-812971bcd35be8a62878a17946ff18fc7b4866966e4d2ae99b29e2d5ae3f53b7"></a>

## Next pages — bgp_parameters / b80a57157cbb / 6

- [bgp_parameters.from_site](data-sources--bgp--reference--group-001.md#canonical-0ad063296715970f438302b0de8e53652e57261c59fa05f6c4363974a54935d1)
- [bgp_parameters.local_address](data-sources--bgp--reference--group-001.md#canonical-d51790f759824fb5b933d0ab88894a1a5e14c1540530296d511b9c3b196a3092)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-0ad063296715970f438302b0de8e53652e57261c59fa05f6c4363974a54935d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22570764e3cae755d13b98ddc2116a84a7729ed61e5ec56b2b54636c4400a2b5"></a>

## bgp_parameters.from_site — bgp_parameters.from_site / 6c4e46e687c1 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-7f432c686742ce5c6b276a53d92c6c021bab0da00649acb0c5a00d3656909885)
- bgp_parameters.from_site

<a id="canonical-e8ada238e545d64365567e4904b89c62e1255c620fbe8e92ce104c94b44d0d96"></a>

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

<a id="canonical-4a2b6e3c88bb8d4232f8a6dff6e6e4e19ea34d1694a1facea6060e60477d41d3"></a>

## Direct properties — bgp_parameters.from_site / 6c4e46e687c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c476ac0320bde6f38f39f9ebcb4d00f7062572c84b9535fef425179ac83aa54"></a>

## Next pages — bgp_parameters.from_site / 6c4e46e687c1 / 4

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-7f432c686742ce5c6b276a53d92c6c021bab0da00649acb0c5a00d3656909885)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-d51790f759824fb5b933d0ab88894a1a5e14c1540530296d511b9c3b196a3092"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8b391a3c03fe1df3c1311f949b99b2145356efe7d48d8622e7b88e50c070de0"></a>

## bgp_parameters.local_address — bgp_parameters.local_address / e705a6e2452e / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-7f432c686742ce5c6b276a53d92c6c021bab0da00649acb0c5a00d3656909885)
- bgp_parameters.local_address

<a id="canonical-c3d638b59ffd9676e1afd3b38606044823599f1cb069ebb60bbb9e7538d72a6b"></a>

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

<a id="canonical-a60bb68c033feb3bcabce5bbcc64ad139a43ec77d5711d19fe8689b23b29fdcd"></a>

## Direct properties — bgp_parameters.local_address / e705a6e2452e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef1783bd9ec967bdc6122afabb3089579320cb69537ec3511dc2c6e8c59a251b"></a>

## Next pages — bgp_parameters.local_address / e705a6e2452e / 4

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-7f432c686742ce5c6b276a53d92c6c021bab0da00649acb0c5a00d3656909885)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1d137c1687d7765d9bcabb6c229318f67503e6880d2fb85d27388aaebe6e09c"></a>

## peers — peers / fdccd76f984e / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- peers

<a id="canonical-13e0c063a45f2dfa92863ce1e8697ae5787bcc92a6a44bf9ccb2b593488e022d"></a>

Type: `"list"`. Computed.

Peers. List of peers.

Upstream description:

List of peers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-d9430340bb6e1c73e2cac4a44fdc764896f1a623a811979c2681d78c1add98c2"></a>

## Direct properties — peers / fdccd76f984e / 3

- [bfd_disabled](data-sources--bgp--reference--group-001.md#canonical-3f49d82e2646cc12189ea96e6932697f85e46ba5083125d48dc13e42db471dbc): complete subsection reference.

- [bfd_enabled](data-sources--bgp--reference--group-001.md#canonical-11dd1be4e14a02612422fd7fc47e4a34e5d952357ba338ab177ba029f17ccc5c): complete subsection reference.

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-2a42cfe2cd8093390b342d72569cb60abd8d75a608432966bb7cb4f82c4a2003): complete subsection reference.

- [ebgp_multihop_disabled](data-sources--bgp--reference--group-001.md#canonical-0ff53ac6fa782e239d35b5d68d70acaeb5acbb88f1a12dcc290c01eb149a0408): complete subsection reference.

- [ebgp_multihop_enabled](data-sources--bgp--reference--group-001.md#canonical-8598b51c125a58d045b3f4adb32a1b04d0e34a4c03a21b4cb22cbdae5fe10ec7): complete subsection reference.

- [external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72): complete subsection reference.

<a id="canonical-1b67a4c82b231c3ce56de9e70aecb6682fb1ba23c12baf486453f6ea55d7d0ce"></a>

<a id="canonical-3e804ab54dd7d90b7e1adea51d724d7acfec2bbdc716ee3fb303168d812d460d"></a>

## label property — peers / fdccd76f984e / 4

Type: `"string"`. Computed.

Label. Specify whether this peer should be.

Upstream description:

Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--bgp--reference--group-001.md#canonical-648e6bbcb68514c2c93dda03d0ed1085bf0fbfa3be5c718ff28a97f76d2a85fd): complete subsection reference.

- [passive_mode_disabled](data-sources--bgp--reference--group-001.md#canonical-fca83207c2d818c7d2d8cef72b4eee6f63d1a5e6ea2800a0f0f0cc08058e6ea8): complete subsection reference.

- [passive_mode_enabled](data-sources--bgp--reference--group-001.md#canonical-d58955caae51861025d930ed8afb777f88635da0349df5520c2557d4548cd5c2): complete subsection reference.

- [routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9): complete subsection reference.

<a id="canonical-39c34966a94f4be57ab8a177633e3b6b35fd8c1c3eca1737f196acfc065ea729"></a>

## Next pages — peers / fdccd76f984e / 5

- [peers.bfd_disabled](data-sources--bgp--reference--group-001.md#canonical-3f49d82e2646cc12189ea96e6932697f85e46ba5083125d48dc13e42db471dbc)
- [peers.bfd_enabled](data-sources--bgp--reference--group-001.md#canonical-11dd1be4e14a02612422fd7fc47e4a34e5d952357ba338ab177ba029f17ccc5c)
- [peers.disable_spec](data-sources--bgp--reference--group-001.md#canonical-2a42cfe2cd8093390b342d72569cb60abd8d75a608432966bb7cb4f82c4a2003)
- [peers.ebgp_multihop_disabled](data-sources--bgp--reference--group-001.md#canonical-0ff53ac6fa782e239d35b5d68d70acaeb5acbb88f1a12dcc290c01eb149a0408)
- [peers.ebgp_multihop_enabled](data-sources--bgp--reference--group-001.md#canonical-8598b51c125a58d045b3f4adb32a1b04d0e34a4c03a21b4cb22cbdae5fe10ec7)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [peers.metadata](data-sources--bgp--reference--group-001.md#canonical-648e6bbcb68514c2c93dda03d0ed1085bf0fbfa3be5c718ff28a97f76d2a85fd)
- [peers.passive_mode_disabled](data-sources--bgp--reference--group-001.md#canonical-fca83207c2d818c7d2d8cef72b4eee6f63d1a5e6ea2800a0f0f0cc08058e6ea8)
- [peers.passive_mode_enabled](data-sources--bgp--reference--group-001.md#canonical-d58955caae51861025d930ed8afb777f88635da0349df5520c2557d4548cd5c2)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-3f49d82e2646cc12189ea96e6932697f85e46ba5083125d48dc13e42db471dbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac1dfc2443a5fa7b5944e9bbd9a8b617b9d8499ce46138bcf010e8fad7e5b683"></a>

## peers.bfd_disabled — peers.bfd_disabled / 99a8d89ddfed / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.bfd_disabled

<a id="canonical-30a5425146c75639a31d4927cb128df9c9d6a20d54c062275dedd9c253f6719b"></a>

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

<a id="canonical-16e8c45bea8a7b3c4e2ae56bd14837bd2baf16f27f77eda5e38c4ef85d66acb1"></a>

## Direct properties — peers.bfd_disabled / 99a8d89ddfed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-09b82155d304c2c961a0b97464da755a4016570264da5a4dec030b2bb385ea83"></a>

## Next pages — peers.bfd_disabled / 99a8d89ddfed / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-11dd1be4e14a02612422fd7fc47e4a34e5d952357ba338ab177ba029f17ccc5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e4353b98f0c8b21dbc06290e1feedcae15549bf1f3008c12553f100358a91fc"></a>

## peers.bfd_enabled — peers.bfd_enabled / 49acf3f76420 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.bfd_enabled

<a id="canonical-c00e4e7e959f127f22268399f29bd4040c8954d4f8f34e02eb8e79bd97d79052"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

Upstream description:

BFD parameters.

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

<a id="canonical-5a7061da34ecd76362f32e249928718c1ed7b1ed5a08dc3d8ee0473f8b14764a"></a>

## Direct properties — peers.bfd_enabled / 49acf3f76420 / 3

<a id="canonical-38c55825c2de0b92da06018c0664545bae43b365eacb64e0a6334630cead59d8"></a>

<a id="canonical-b0c57ca93df30c569a9ff6294c6fef8e989221ccf091473a968d9fad3d017892"></a>

## multiplier property — peers.bfd_enabled / 49acf3f76420 / 4

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Upstream description:

Specify Number of missed packets to bring session down"

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
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-ffc4a5156e11bb6de7195118eb0e171a93c9070a244281842d73fc1772e05513"></a>

<a id="canonical-b453b8eb590515582387dee69b3c7a815884015d9ccf5f0006c5259ce15aaf7f"></a>

## receive_interval_milliseconds property — peers.bfd_enabled / 49acf3f76420 / 5

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-d65b74e8103fe6794506a9c9f8188f19a3abb288815fe0a6e9aa173c93a176a0"></a>

<a id="canonical-48bd9a7b097302fdc74c4ba4f97b25674ea3b40f75a73c46e12ebea9e2cc2c3b"></a>

## transmit_interval_milliseconds property — peers.bfd_enabled / 49acf3f76420 / 6

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-38e498718d8532b747fa9d5d1925cce0a38c323dc2a1ef7bf4f0015aa56b49da"></a>

## Next pages — peers.bfd_enabled / 49acf3f76420 / 7

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-2a42cfe2cd8093390b342d72569cb60abd8d75a608432966bb7cb4f82c4a2003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2976dfc53e257266224ba410faf8d550c14c5236497c153bf63c6f13d3671c8f"></a>

## peers.disable_spec — peers.disable_spec / 7b8528ca9f12 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.disable_spec

<a id="canonical-232540b4a1df3b9d6aa0a2612d2106a440bf02312126de65da1616bcaa33ab18"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-3f9ca92022af4a590b53365c9d0b65eda31902f99f71456927f05cb3a122b6a6"></a>

## Direct properties — peers.disable_spec / 7b8528ca9f12 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76fdb6c1fb64b9d5023ea561779b1f5897ac028605e89c96e7574ba58779d70a"></a>

## Next pages — peers.disable_spec / 7b8528ca9f12 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-0ff53ac6fa782e239d35b5d68d70acaeb5acbb88f1a12dcc290c01eb149a0408"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-541a129d02255d1a9e52aa317229623c4d930df29ce588d71e08152c2308f0f9"></a>

## peers.ebgp_multihop_disabled — peers.ebgp_multihop_disabled / 857bcc226344 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.ebgp_multihop_disabled

<a id="canonical-87e8ed37a0aae1696126e5cda70c12f2fc1f19ba4315351f3110c84fa0968b89"></a>

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

<a id="canonical-02f2704eb43d94c3ff7f83a5b326f13a9c82449907f3e6403743e925338bb78c"></a>

## Direct properties — peers.ebgp_multihop_disabled / 857bcc226344 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-756db02f1bbf2507c5cc9372262dccee56f9b0823a96826e660361ace3da24c9"></a>

## Next pages — peers.ebgp_multihop_disabled / 857bcc226344 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-8598b51c125a58d045b3f4adb32a1b04d0e34a4c03a21b4cb22cbdae5fe10ec7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-251e3051f1c78bf6bdde0533d006ff8ec018aa1d89dd1afee410434b73cd320f"></a>

## peers.ebgp_multihop_enabled — peers.ebgp_multihop_enabled / d57e03c9edda / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.ebgp_multihop_enabled

<a id="canonical-bd3c01a3867b1fb4a6aaca5d12d711b938a6798fe2cf57d8fdd3c05bd51d247a"></a>

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

<a id="canonical-fdeb64124e9a429407dd2ec4b74462d6a86cf5419b8a28563feae3b3e5dae09e"></a>

## Direct properties — peers.ebgp_multihop_enabled / d57e03c9edda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b60481d99b11416e0c80f467b264c99025210ccd6ffd88e4e93375d203c407f"></a>

## Next pages — peers.ebgp_multihop_enabled / d57e03c9edda / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-334cb1ddbe39ba2b74a6ab89ca50bbd2df04f6cc04be885abf0d58a72a7c0b58"></a>

## peers.external — peers.external / 37af7b448bad / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.external

<a id="canonical-f8b4c6b6ec8f6990ff8de56af59d30498a8db1a345f680ed843682b978e04ac3"></a>

Type: `"single"`. Computed.

External BGP Peer. External BGP Peer parameters.

Upstream description:

External BGP Peer parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

<a id="canonical-60879d1ee5c46063b9395ac382e57c55ec64ac279949432a0a9b2e78d89a97bc"></a>

## Direct properties — peers.external / 37af7b448bad / 3

<a id="canonical-59a2cb0be716462d1213071633771f886cb164915e500e43d17b11b2515a7279"></a>

<a id="canonical-4ff310110f0c9360fb649f1f63d37b47fe337f386343b7f58d8970d7b4c8a732"></a>

## address property — peers.external / 37af7b448bad / 4

Type: `"string"`. Computed.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Upstream description:

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

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

<a id="canonical-52a2dfa1f88cd3d34d87446c1d50d6718057e6820eb840e479ceb172b6bb814e"></a>

<a id="canonical-998af784768cd4925517e76c1851942410d35a903535e5e26b46538d8c8d3202"></a>

## address_ipv6 property — peers.external / 37af7b448bad / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Upstream description:

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

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

<a id="canonical-49aaf06ee83db0cc5aa6211f7b37b3cc626b1fb62a861d0675f0663c886cf73d"></a>

<a id="canonical-457bbca19538a0be7b8d21df5e95faed5b73632e02673467689aab4b11ad668e"></a>

## asn property — peers.external / 37af7b448bad / 6

Type: `"number"`. Computed.

ASN. Autonomous System Number for BGP peer.

Upstream description:

Autonomous System Number for BGP peer.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](data-sources--bgp--reference--group-001.md#canonical-f254283fae7611a8c674da632d3fc6cf36cdac7e8c0e00cfd342bb71998c8148): complete subsection reference.

- [default_gateway_v6](data-sources--bgp--reference--group-001.md#canonical-82e2d64dd04265cc2c67f8877baefeca1d247881042c3636834a8be0b9427702): complete subsection reference.

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-8dcabb8dfe2bfc2ceefd66e5a1acecedfa0a6292202ad431718e652fb1bad4c2): complete subsection reference.

- [disable_v6](data-sources--bgp--reference--group-001.md#canonical-0950a4cc550953fd613d5b453af39147fe9154133e375fbe0c5575e46120a1fc): complete subsection reference.

- [external_connector](data-sources--bgp--reference--group-001.md#canonical-baade174681c6a2ed9d542debf6e1c7f1a862b60f3489a5ce0c3d3a7f1677c70): complete subsection reference.

- [family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc): complete subsection reference.

- [from_site](data-sources--bgp--reference--group-001.md#canonical-8084bb055dab83912525af41e5de6c46ec96541ca8aa7e14f50e6215a97e1229): complete subsection reference.

- [from_site_v6](data-sources--bgp--reference--group-001.md#canonical-bc95bb841a325e594c5e591b2fec1bea324d20feeb8421ed13183767c4238615): complete subsection reference.

- [interface](data-sources--bgp--reference--group-001.md#canonical-62c0a0def5c692d1b8e57c4a6b77bc7aa95fd67dd408ea296c8674357c5b424b): complete subsection reference.

- [interface_list](data-sources--bgp--reference--group-001.md#canonical-3e580e618670e437c32d05896906644dea9fccaf4ef47ede76ab9df2627e50ab): complete subsection reference.

<a id="canonical-4aed42603c600cc1463ae045ca7d79cc12e6c1ad62ae5f48060b8919cc426bb3"></a>

<a id="canonical-68952cf99f3b3e479f6e4ac15e4fd06ee7a093cf1270b6ca1cfe17d71bd9453d"></a>

## md5_auth_key property — peers.external / 37af7b448bad / 7

Type: `"string"`. Computed.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Upstream description:

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385)

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

- [no_authentication](data-sources--bgp--reference--group-001.md#canonical-58306e25aae65df59212fe6087d9d10179f1040cf052de05c42fa38b153ecd16): complete subsection reference.

<a id="canonical-773e89e0492c476052c171577ff7cbf02f87dd4f3a6221e02ea70f315e371eff"></a>

<a id="canonical-e2c87f30e230105f6f70819651246aecd196478b8733836864bee29c83bcefaa"></a>

## port property — peers.external / 37af7b448bad / 8

Type: `"number"`. Computed.

Peer Port. Peer TCP port number.

Upstream description:

Peer TCP port number.

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
    "create": true,
    "minimum_config": true,
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

<a id="canonical-bd6101e90b2bf252742d7b939e50d58aca0a773c37ac08269c9a0dc67df47421"></a>

<a id="canonical-0254b5cf602a441be15ffef311f8889dc14db17c086a01f67efada9f42d2976d"></a>

## subnet_begin_offset property — peers.external / 37af7b448bad / 9

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-d9395bef39286bd0448758e41bd895853d6ea622b711833dbc0101792147e313"></a>

<a id="canonical-e0b53d1ef95d4eafc9e6947350e33812260be2de9a6509dbc0ae4858a4d650ee"></a>

## subnet_begin_offset_v6 property — peers.external / 37af7b448bad / 10

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-ded9c00013201cc4b688801a6477aae71772815533b342d524fc0040db4fb49d"></a>

<a id="canonical-5da845a02827c2c01a11bfdeaade2852c950f1c3bea5ba84c46ac30f697f1b16"></a>

## subnet_end_offset property — peers.external / 37af7b448bad / 11

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-92cd88aa516a0c3db977a4f035c26dd4f352f99cec79eba941e2e36d7c9c9075"></a>

<a id="canonical-2213ddf2d4f26961be7c410f8901285609491e70e013be017c9e21551656630d"></a>

## subnet_end_offset_v6 property — peers.external / 37af7b448bad / 12

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Upstream description:

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-02411018c0282fc068827328544d8bc06dcfc8572771bf07d8968191a217734a"></a>

## Next pages — peers.external / 37af7b448bad / 13

- [peers.external.default_gateway](data-sources--bgp--reference--group-001.md#canonical-f254283fae7611a8c674da632d3fc6cf36cdac7e8c0e00cfd342bb71998c8148)
- [peers.external.default_gateway_v6](data-sources--bgp--reference--group-001.md#canonical-82e2d64dd04265cc2c67f8877baefeca1d247881042c3636834a8be0b9427702)
- [peers.external.disable_spec](data-sources--bgp--reference--group-001.md#canonical-8dcabb8dfe2bfc2ceefd66e5a1acecedfa0a6292202ad431718e652fb1bad4c2)
- [peers.external.disable_v6](data-sources--bgp--reference--group-001.md#canonical-0950a4cc550953fd613d5b453af39147fe9154133e375fbe0c5575e46120a1fc)
- [peers.external.external_connector](data-sources--bgp--reference--group-001.md#canonical-baade174681c6a2ed9d542debf6e1c7f1a862b60f3489a5ce0c3d3a7f1677c70)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- [peers.external.from_site](data-sources--bgp--reference--group-001.md#canonical-8084bb055dab83912525af41e5de6c46ec96541ca8aa7e14f50e6215a97e1229)
- [peers.external.from_site_v6](data-sources--bgp--reference--group-001.md#canonical-bc95bb841a325e594c5e591b2fec1bea324d20feeb8421ed13183767c4238615)
- [peers.external.interface](data-sources--bgp--reference--group-001.md#canonical-62c0a0def5c692d1b8e57c4a6b77bc7aa95fd67dd408ea296c8674357c5b424b)
- [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-3e580e618670e437c32d05896906644dea9fccaf4ef47ede76ab9df2627e50ab)
- [peers.external.no_authentication](data-sources--bgp--reference--group-001.md#canonical-58306e25aae65df59212fe6087d9d10179f1040cf052de05c42fa38b153ecd16)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-f254283fae7611a8c674da632d3fc6cf36cdac7e8c0e00cfd342bb71998c8148"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d61468ee70a708d5fc818448b6b195879aee4d3fcd7963edd6e05a7130b6a1b9"></a>

## peers.external.default_gateway — peers.external.default_gateway / ec8ab576d465 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.default_gateway

<a id="canonical-09675cc0c43f9623041317631684ca3754a092b3232b0ce55501c86d9d3eaa22"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-93675bbb60ab2443fd647fae40ccb7c0d0889f5339db038ad6105f4a77c5b288"></a>

## Direct properties — peers.external.default_gateway / ec8ab576d465 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4424e8a1c64feae76c38ce053a9787a7b1b4f9e3c244c22c8aeecbc5266ec3e0"></a>

## Next pages — peers.external.default_gateway / ec8ab576d465 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-82e2d64dd04265cc2c67f8877baefeca1d247881042c3636834a8be0b9427702"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f33d4b991155c0be0f5a53e9d0d5a6a1844eb3c371abde056ce41eddc60f5821"></a>

## peers.external.default_gateway_v6 — peers.external.default_gateway_v6 / e3dc58d8398e / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.default_gateway_v6

<a id="canonical-2c62b6e810d7fe285fe6c5ac6ce044e98fbfe6794d457f45abb08e40d5004100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway v6.

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

<a id="canonical-b0df2dd04e6c6a700aa079c70f85401d2fdf3121d43277f1142c808c474e7761"></a>

## Direct properties — peers.external.default_gateway_v6 / e3dc58d8398e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4c14b9ce904e20b29c74ac2bfd9bd06204f2a4e787416b3d06a2b541e2ca163a"></a>

## Next pages — peers.external.default_gateway_v6 / e3dc58d8398e / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-8dcabb8dfe2bfc2ceefd66e5a1acecedfa0a6292202ad431718e652fb1bad4c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4da4373fc4afdced413db8dc58769121d2568123e12e74c69d2df1b0a8136e19"></a>

## peers.external.disable_spec — peers.external.disable_spec / ee4d071a9460 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.disable_spec

<a id="canonical-dff2dde2d05a41e191feff25103a81ae801f7f8e20e2ec0a51b933a4b1363465"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-13cd6c230f3170527b1e668bf4dfb40a363a7c65990ba38a5fb1737f2f523f3f"></a>

## Direct properties — peers.external.disable_spec / ee4d071a9460 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3253e0c768f2e3d78a83fc7721efd100ebcedff7494796a1f8b69db887250d5"></a>

## Next pages — peers.external.disable_spec / ee4d071a9460 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-0950a4cc550953fd613d5b453af39147fe9154133e375fbe0c5575e46120a1fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-971db604a9276e3874935381834164e2bd351b57085943b12c4c3c8ebad32169"></a>

## peers.external.disable_v6 — peers.external.disable_v6 / e45daebc8648 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.disable_v6

<a id="canonical-39ac65780d9f06f751ed4bcef354864b4e2f1d9da463c313e7c981c0aa1eaf76"></a>

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

<a id="canonical-bcf8f88eb595a2bdadadc390073a6aaba4906ea7a0b4f32a28e52f36b323dbc0"></a>

## Direct properties — peers.external.disable_v6 / e45daebc8648 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f0e49f747861f71966725ecbacf5bb4a025d04292f812a9d66e25ff7db683b2"></a>

## Next pages — peers.external.disable_v6 / e45daebc8648 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-baade174681c6a2ed9d542debf6e1c7f1a862b60f3489a5ce0c3d3a7f1677c70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a3a2f2285d82872f027e858560d2687cf5322cd5e6e30fb5207db3b9da2460b"></a>

## peers.external.external_connector — peers.external.external_connector / a34fefab5454 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.external_connector

<a id="canonical-9068999fb414ed5912cff88b76507c72ea0f7cf69d3c5264ab08b8fcacf8940a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external connector.

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

<a id="canonical-94647ab2df66b94c4c7f8e1407089f0ba335646747382f009349f5c23f671d8e"></a>

## Direct properties — peers.external.external_connector / a34fefab5454 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-234d316c6140e28d5e1c725b19577625f1090a20f8bdf7c1f166864877b66567"></a>

## Next pages — peers.external.external_connector / a34fefab5454 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfc969b8a772514ed7f679f8f52bc677ffc8c119e00e09a8efaa55756434962d"></a>

## peers.external.family_inet — peers.external.family_inet / c20fff65c4db / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.family_inet

<a id="canonical-808d2a579129014b2399fb6655a46dbc66c909ed635478181115e445b825c87b"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-29687c6f9bc555fd5c253c4f10f694be618e482b72226c515a53d9e4de044faa"></a>

## Direct properties — peers.external.family_inet / c20fff65c4db / 3

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-534a3be14032e088c8c0bc3061c5b4d938db79896444b626504db67e7689ae72): complete subsection reference.

- [enable](data-sources--bgp--reference--group-001.md#canonical-62a69e7e461498ef3a16a9016501803f299ccaa837dd8015ad58db73ea9bb49e): complete subsection reference.

<a id="canonical-802ee57681561f6a684f4ac39e6b048c3753b764a214f7d054a085e33b440ce1"></a>

## Next pages — peers.external.family_inet / c20fff65c4db / 4

- [peers.external.family_inet.disable_spec](data-sources--bgp--reference--group-001.md#canonical-534a3be14032e088c8c0bc3061c5b4d938db79896444b626504db67e7689ae72)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-62a69e7e461498ef3a16a9016501803f299ccaa837dd8015ad58db73ea9bb49e)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-534a3be14032e088c8c0bc3061c5b4d938db79896444b626504db67e7689ae72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abf7b7779654e363193c317ad29358451583ce7426e75c46eaeea50fb7a64478"></a>

## peers.external.family_inet.disable_spec — peers.external.family_inet.disable_spec / b91f708438b8 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- peers.external.family_inet.disable_spec

<a id="canonical-ea7a6bca20bbd829e717326f060419fcf9e21fdda3bdb094b9bee4bc7efa4735"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-25cb9149f68daeb763d767f7dc7d09317a991a10b1ac6b722c0747d152db8f35"></a>

## Direct properties — peers.external.family_inet.disable_spec / b91f708438b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd58164d05e3d3c0ea5c1a14705af0f7564c88baa22a1578639d89776ffa6d26"></a>

## Next pages — peers.external.family_inet.disable_spec / b91f708438b8 / 4

- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-62a69e7e461498ef3a16a9016501803f299ccaa837dd8015ad58db73ea9bb49e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e7a13325966296b9c9dd1720d8517752f0b531fdd926e6446e6bbd7f21c032d"></a>

## peers.external.family_inet.enable — peers.external.family_inet.enable / fa5c010cc7b8 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- peers.external.family_inet.enable

<a id="canonical-45be3d5388741d132473b33de8f2009b699fb83b20b7178ba6a88e76f132f1d0"></a>

Type: `"single"`. Computed.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

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

<a id="canonical-357a5eefbdde75c94246416abb190a688aa8759734e27eabb6ed553f0d5a09db"></a>

## Direct properties — peers.external.family_inet.enable / fa5c010cc7b8 / 3

- [aggregation](data-sources--bgp--reference--group-001.md#canonical-32e508af2d9fad43cca639b2d0ede9a0984836759f90aadeb4f894d7a2ac8a72): complete subsection reference.

<a id="canonical-ec6e2e61cce271fca565be2de64b911a8a6f6182ef2b8626008fa3b6002ebc83"></a>

## Next pages — peers.external.family_inet.enable / fa5c010cc7b8 / 4

- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-32e508af2d9fad43cca639b2d0ede9a0984836759f90aadeb4f894d7a2ac8a72)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-32e508af2d9fad43cca639b2d0ede9a0984836759f90aadeb4f894d7a2ac8a72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-714d4dd826633bd833b21e2b734c27f97e7bc457b053cadeb2e5caae8be194a6"></a>

## peers.external.family_inet.enable.aggregation — peers.external.family_inet.enable.aggregation / 6e35fb6f5302 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-62a69e7e461498ef3a16a9016501803f299ccaa837dd8015ad58db73ea9bb49e)
- peers.external.family_inet.enable.aggregation

<a id="canonical-0e8c65874da323ff74a828ec9d31cd58e8c8a2fdd092305b1ad7c602d957de18"></a>

Type: `"list"`. Computed.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

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

<a id="canonical-2691ce1be8f78df654b8e46a459dd349e2f79be5bfc43e9d79015bc145420e46"></a>

## Direct properties — peers.external.family_inet.enable.aggregation / 6e35fb6f5302 / 3

<a id="canonical-97471307a0b44e8f73d16b0f76a317b8a04d487e02d425ff0f393151d54b0a15"></a>

<a id="canonical-00aefb30a152474d67c0d748c6f2dd873057ab5b5d73dc2a8060ac533ed07d7b"></a>

## ip_prefix property — peers.external.family_inet.enable.aggregation / 6e35fb6f5302 / 4

Type: `"string"`. Computed.

IP Prefix. Specify IPv4 subnet for aggregation.

Upstream description:

Specify IPv4 subnet for aggregation.

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

- [options](data-sources--bgp--reference--group-001.md#canonical-471fce84614038e688ad9df86f656689db3234c9256d78bbf896145f1fd5b57d): complete subsection reference.

<a id="canonical-058ad01f00de16224a07f9fcffb9b45640faa6d006557eb2675d926e9a938b91"></a>

## Next pages — peers.external.family_inet.enable.aggregation / 6e35fb6f5302 / 5

- [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-471fce84614038e688ad9df86f656689db3234c9256d78bbf896145f1fd5b57d)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-62a69e7e461498ef3a16a9016501803f299ccaa837dd8015ad58db73ea9bb49e)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-471fce84614038e688ad9df86f656689db3234c9256d78bbf896145f1fd5b57d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb8e1168ad31c5df445e12e5cc7d1f87a0cdefe87a7072ddf0da8f0ce3ecc094"></a>

## peers.external.family_inet.enable.aggregation.options — peers.external.family_inet.enable.aggregation.options / 1addc15b0b67 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-62a69e7e461498ef3a16a9016501803f299ccaa837dd8015ad58db73ea9bb49e)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-32e508af2d9fad43cca639b2d0ede9a0984836759f90aadeb4f894d7a2ac8a72)
- peers.external.family_inet.enable.aggregation.options

<a id="canonical-2365903b12359d093a5a5a4b93fdd9703c28f6f374baad804996818b4d182456"></a>

Type: `"list"`. Computed.

Aggregation OPTIONS. Configuration parameter for options

Upstream description:

Configuration parameter for options

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

<a id="canonical-dd1eefa1b510c5d2f922f4d271ad62be7bccc353bebce6c91a2851cf0a7e6364"></a>

## Direct properties — peers.external.family_inet.enable.aggregation.options / 1addc15b0b67 / 3

- [summary_only](data-sources--bgp--reference--group-001.md#canonical-967b88f493bf99246348def8203278c95de644fa7f0d303ec33c12147fa02b6e): complete subsection reference.

<a id="canonical-03c1e2e7abd6dfb1418a573f7836e442198c0a5cbb8fca69c65c31049bf04c4f"></a>

## Next pages — peers.external.family_inet.enable.aggregation.options / 1addc15b0b67 / 4

- [peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--bgp--reference--group-001.md#canonical-967b88f493bf99246348def8203278c95de644fa7f0d303ec33c12147fa02b6e)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-32e508af2d9fad43cca639b2d0ede9a0984836759f90aadeb4f894d7a2ac8a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-967b88f493bf99246348def8203278c95de644fa7f0d303ec33c12147fa02b6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aceca766103f5c71f0c11da440afb365d81145f667b56c4cc3280583affcb7a"></a>

## peers.external.family_inet.enable.aggregation.options.summary_only — peers.external.family_inet.enable.aggregation.options.summary_only / 82f307bfb7c1 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-ef3e3a638b29222c996bbb1227a72dfbb115ceb9f34a8b0a41b55f5421afcecc)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-62a69e7e461498ef3a16a9016501803f299ccaa837dd8015ad58db73ea9bb49e)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-32e508af2d9fad43cca639b2d0ede9a0984836759f90aadeb4f894d7a2ac8a72)
- [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-471fce84614038e688ad9df86f656689db3234c9256d78bbf896145f1fd5b57d)
- peers.external.family_inet.enable.aggregation.options.summary_only

<a id="canonical-3c34517c52c85250dfaa47f249f107ecf2a68a10f8021f4fac05657de7bd244f"></a>

Type: `"single"`. Computed.

Configuration parameter for summary only.

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

<a id="canonical-0a054c07e452732cc9f88f678ed867d5512fe14e4d1d69c9eea333fd84e9a76e"></a>

## Direct properties — peers.external.family_inet.enable.aggregation.options.summary_only / 82f307bfb7c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ec8cf65da46a642234629534e68b7fba78136ea3ea734d26534cc40b03de0f8"></a>

## Next pages — peers.external.family_inet.enable.aggregation.options.summary_only / 82f307bfb7c1 / 4

- [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-471fce84614038e688ad9df86f656689db3234c9256d78bbf896145f1fd5b57d)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-8084bb055dab83912525af41e5de6c46ec96541ca8aa7e14f50e6215a97e1229"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a054ef98928319c644af8f430373a37f6a4c36ee50df667dd5269205bfd15bdc"></a>

## peers.external.from_site — peers.external.from_site / 5277edf1995d / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.from_site

<a id="canonical-82f71c45038132c92ddafd7623124ee65bb7808ebd06bf923122ff6cf1c8535e"></a>

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

<a id="canonical-cb578036c7c6a8c42588b2d3e91dcf52fbdd70d0db036b5cde104f020171ccd1"></a>

## Direct properties — peers.external.from_site / 5277edf1995d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1724acc27a64973df80c805cd31eb4cdd4ceddb1b2ed4eaedf14467d2873a854"></a>

## Next pages — peers.external.from_site / 5277edf1995d / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-bc95bb841a325e594c5e591b2fec1bea324d20feeb8421ed13183767c4238615"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d4af687c6e5603fd7257b7263437d034e0632601b2fa70b37ff0b7750eb8433"></a>

## peers.external.from_site_v6 — peers.external.from_site_v6 / 173a53e3988e / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.from_site_v6

<a id="canonical-9221174021c97a00e575e082cec9683934cacf3d60f4fbb7e8f9b8bd8a322e25"></a>

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

<a id="canonical-4ac600e41e2483591efadbf19335005a7514aad73cc19e44f92ce671f4deaafe"></a>

## Direct properties — peers.external.from_site_v6 / 173a53e3988e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d0e04953d46a2d03249c02727e3e4b4fbead8ce8cc0eed37a30a086d148eb3ce"></a>

## Next pages — peers.external.from_site_v6 / 173a53e3988e / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-62c0a0def5c692d1b8e57c4a6b77bc7aa95fd67dd408ea296c8674357c5b424b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd359f8dd877aa2dcd0a81b6075dfe3c3c6829c3e7a56d93125d0d194bfb8acd"></a>

## peers.external.interface — peers.external.interface / 1d058f67cfa2 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.interface

<a id="canonical-35b6ed865f2da1dcbb530274fd7099bc31fb43ed4f33f24ce276a42ff5432d0f"></a>

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

<a id="canonical-3345e426fe08ad582d398af25d152ba5eb429d6218cb8c363d38d574f3918e93"></a>

## Direct properties — peers.external.interface / 1d058f67cfa2 / 3

<a id="canonical-bde97a8cdb534da663e16732231c5a5f434322b583583740365d4ec7d8103174"></a>

<a id="canonical-d8535b2f4642cfb03be9e08cf1cf9b823879ecac16b642fc23f92e554d83a89b"></a>

## name property — peers.external.interface / 1d058f67cfa2 / 4

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

<a id="canonical-e78944a729a04f0001f6f3e8f487135300e67b8badfcf166e4c598161f620421"></a>

<a id="canonical-8ae0a1cb2535a7b8b35d29fa4d0253a9abe17ab9df298acd8dd0bbac3b71f3b8"></a>

## namespace property — peers.external.interface / 1d058f67cfa2 / 5

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

<a id="canonical-84ca59d3c70a9201fe5b6d5c1cca09e5f7c04e06f9d891c3c8e32f16da638465"></a>

<a id="canonical-ffc3d4b5a5d5e74ee9558ee999d261aa2e6342e6f30bf35723bea45499851add"></a>

## tenant property — peers.external.interface / 1d058f67cfa2 / 6

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

<a id="canonical-12a5fe552b949e8a0efa51f56f533eb08d5aa260ba528110940bb3b321913920"></a>

## Next pages — peers.external.interface / 1d058f67cfa2 / 7

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-3e580e618670e437c32d05896906644dea9fccaf4ef47ede76ab9df2627e50ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2b4e775349f843b91fb33afe6b8679b8679cedad4faa0d4c34c6045022b0a71"></a>

## peers.external.interface_list — peers.external.interface_list / 324edc86fb30 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.interface_list

<a id="canonical-5d0d23e53fa4ea2b0b80eb5087c16c5c0dcc85dd84e9e7f0468332fddd486fb0"></a>

Type: `"single"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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

<a id="canonical-ab2a2b4dd7be0bd36286f36c0b78d30c962502aa9df2066fd6be20b05551a703"></a>

## Direct properties — peers.external.interface_list / 324edc86fb30 / 3

- [interfaces](data-sources--bgp--reference--group-001.md#canonical-7c82ea562ff441c31f20a90101080747da52f5961ce70e11e7437c97e30bc2bb): complete subsection reference.

<a id="canonical-fe31ae75f38fd273282d67a5478063273c0dc944c2cbf4432bedefe48bc4d4ba"></a>

## Next pages — peers.external.interface_list / 324edc86fb30 / 4

- [peers.external.interface_list.interfaces](data-sources--bgp--reference--group-001.md#canonical-7c82ea562ff441c31f20a90101080747da52f5961ce70e11e7437c97e30bc2bb)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-7c82ea562ff441c31f20a90101080747da52f5961ce70e11e7437c97e30bc2bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a4ba64b911496b389827d97b0de23f04d22145cdaaa4f580570fd613a95a1d7"></a>

## peers.external.interface_list.interfaces — peers.external.interface_list.interfaces / 6991309739b6 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-3e580e618670e437c32d05896906644dea9fccaf4ef47ede76ab9df2627e50ab)
- peers.external.interface_list.interfaces

<a id="canonical-14a9845636f942336f06d9768501f3e8608e1b651f10751755902c07dfbf0af8"></a>

Type: `"list"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ad98f8a0345a0e7935d1b2a942ef454eb1f9abb784821e34c983e672b4f88898"></a>

## Direct properties — peers.external.interface_list.interfaces / 6991309739b6 / 3

<a id="canonical-4436209ff78d072a3c84749fcdfabbe2de94852764f987072eeb6241bb256df3"></a>

<a id="canonical-42baef3899b7ea2930674c1914b410d2105573407f984c4f27604e9b4790ec14"></a>

## name property — peers.external.interface_list.interfaces / 6991309739b6 / 4

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

<a id="canonical-bb5d238ea92107f41f10255c6ac2aca6e16e6431bef63017c7eae7b90873e5e6"></a>

<a id="canonical-be5ba069b90524a25f8ada1123af73fa0ee9f97a5019084577d0657dd3437fc0"></a>

## namespace property — peers.external.interface_list.interfaces / 6991309739b6 / 5

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

<a id="canonical-a6d845dfff304850b7d2fe2a93027112458eab663f284492e9ad0f4fc510abaa"></a>

<a id="canonical-e1ba18543acd1eb82288d5959a77ce5299c657b048ba06d510302d1bcb62844d"></a>

## tenant property — peers.external.interface_list.interfaces / 6991309739b6 / 6

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

<a id="canonical-037d4a12a73bd506df40e1196b1921afa6c9a8133dd5b1d007099c8ab4568843"></a>

## Next pages — peers.external.interface_list.interfaces / 6991309739b6 / 7

- [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-3e580e618670e437c32d05896906644dea9fccaf4ef47ede76ab9df2627e50ab)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-58306e25aae65df59212fe6087d9d10179f1040cf052de05c42fa38b153ecd16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c12d03794dcfd194789905a50740194759f1a5b8a2a58c9dcd78af7315a47fdb"></a>

## peers.external.no_authentication — peers.external.no_authentication / 90049b8913ef / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- peers.external.no_authentication

<a id="canonical-c48f05cb6d10ef4ce7ccbd26b6a2adbe862c757045b5ec4025acdb725b34ce5a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-0bb6373ba67e8dc3f8f88d8c1d4af93667c4fd795c1f975299c90cee0df95fa2"></a>

## Direct properties — peers.external.no_authentication / 90049b8913ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a023c8526f3af0e00490d6dfee3af10d72b76df92dfe198f9356f37f8bae6d75"></a>

## Next pages — peers.external.no_authentication / 90049b8913ef / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-50d0782463214fda4978dfdc91aefbe75d52dcb3132931126731c058ade45a72)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-648e6bbcb68514c2c93dda03d0ed1085bf0fbfa3be5c718ff28a97f76d2a85fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f046bf2154f2384b10ed1e3e5e466c745c3bf75ea51cf1cce48f4b0c0736c36"></a>

## peers.metadata — peers.metadata / 5fcabad71d93 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.metadata

<a id="canonical-3b52e1a8bb251df25d409d07b0540b2e28fefd35a47cbd04ba048e4f80c81902"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-9ffb1c91aee1450881c144ea4093c8632475109efa0cd1465caabf7637d2ba7b"></a>

## Direct properties — peers.metadata / 5fcabad71d93 / 3

<a id="canonical-abd4b127d9124dc88ca38b856079c9ce3ade6fdb3869e6a37a4d84644d36d62e"></a>

<a id="canonical-cfbce6f4aa8da9b1df33eecd5938f05163b7bda0ca2c129c655dbd2eae372b68"></a>

## description_spec property — peers.metadata / 5fcabad71d93 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-b21f5ba85ef904b4fbaf85c080d0b0523071637f891f4bc71cd5f721734ba4e4"></a>

<a id="canonical-3a0df90ad85b31a64ffdd13d331337a2a48941608d533aaf98bd7400b57063a0"></a>

## name property — peers.metadata / 5fcabad71d93 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-3bbd5a61e50b96686239c6e401f5d9584f37489d4e733a061c03663d5573171a"></a>

## Next pages — peers.metadata / 5fcabad71d93 / 6

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-fca83207c2d818c7d2d8cef72b4eee6f63d1a5e6ea2800a0f0f0cc08058e6ea8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-134f2c22bd1774c33b4ab7f4550ff71bff7c7e80b581c1169f0ae782966146bf"></a>

## peers.passive_mode_disabled — peers.passive_mode_disabled / 4856b0609fd2 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.passive_mode_disabled

<a id="canonical-83cde329c949737ae1baca652e39f9388294e308c8899cb83e99b7a21707966a"></a>

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

<a id="canonical-76f85a90bc42f73fefdb28fbb6e91bff0f9e3abb92036a88800432d7bb76b134"></a>

## Direct properties — peers.passive_mode_disabled / 4856b0609fd2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9baf9db01a9bbbe96426d06058bca74c896a425ea17a68e6c7604eb02668b485"></a>

## Next pages — peers.passive_mode_disabled / 4856b0609fd2 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-d58955caae51861025d930ed8afb777f88635da0349df5520c2557d4548cd5c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65b00a8b7b899cd1a505d2ce2a6e54fdeb825dc7fa8019756b159ba409b7f88f"></a>

## peers.passive_mode_enabled — peers.passive_mode_enabled / 2af0db884fb2 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.passive_mode_enabled

<a id="canonical-58d278b7937230dd68f1f1d0230b48843249ddd995ec633be632a386645a9b54"></a>

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

<a id="canonical-47e1e3b29136f6ac58a3c112f521847d4ed5af822aa96b288f0543eaceb6fb09"></a>

## Direct properties — peers.passive_mode_enabled / 2af0db884fb2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4afa63d407356ef0a4a1bc89fc7a8c50059198fe33a9c512036f614250bda7b5"></a>

## Next pages — peers.passive_mode_enabled / 2af0db884fb2 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36b1b46f2d6a8f23cd00852ab4cc97c5531b80abc686365f9cadf7ddaa6ddffe"></a>

## peers.routing_policies — peers.routing_policies / e5d296bde3c2 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- peers.routing_policies

<a id="canonical-33e8b65f508c2881286868cfd7761e7a23792e6aee67e33bcda7a154b3407cd2"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

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

<a id="canonical-c70050295b64290ad9c4ce4c4edbfa759b732ba38b483f034a9afdfdaf4db97d"></a>

## Direct properties — peers.routing_policies / e5d296bde3c2 / 3

- [route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327): complete subsection reference.

<a id="canonical-6738e3df1357a4fb625280245fb35192f2ff19b2a57da8bd23e1d452a223e750"></a>

## Next pages — peers.routing_policies / e5d296bde3c2 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d1c5aef8b6ceec3be70d7c838067cf37e9d5e70614637bb963d541d5c423c8d"></a>

## peers.routing_policies.route_policy — peers.routing_policies.route_policy / e90b29873d6f / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- peers.routing_policies.route_policy

<a id="canonical-66a858042524b39734cfbb1d8a7d970994d3abf06f7b3e82e18b20d43d7dddb8"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

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

<a id="canonical-96fe147e5d9063d68ed946e0fe7cd4c2267b134e620b9757d1a8a3706ce241b1"></a>

## Direct properties — peers.routing_policies.route_policy / e90b29873d6f / 3

- [all_nodes](data-sources--bgp--reference--group-001.md#canonical-68cae395ed3cd99ab78d3512891513fa80be87f1eb161a93e5207bca2edc9c55): complete subsection reference.

- [inbound](data-sources--bgp--reference--group-001.md#canonical-6efd32095da9547b02622f54386d4b2332e7b818a65f8cb8e72bb95404a14a67): complete subsection reference.

- [node_name](data-sources--bgp--reference--group-001.md#canonical-dcd89b92529b4a74150fe94816ed3414094e9a38a2e611c77fc6002773d43259): complete subsection reference.

- [object_refs](data-sources--bgp--reference--group-001.md#canonical-423bb5d56539eff47654011e39d8706d9b02a1a9c2350d5e93d96f4e848771d0): complete subsection reference.

- [outbound](data-sources--bgp--reference--group-001.md#canonical-be97d6057327dabbde80e3712ac86433974dd51937b19fa94b7c805079a69f52): complete subsection reference.

<a id="canonical-f8d27d0e12b5c32a7b371d14144801ae35e77ec3b0c04e54a24966e409bd5a03"></a>

## Next pages — peers.routing_policies.route_policy / e90b29873d6f / 4

- [peers.routing_policies.route_policy.all_nodes](data-sources--bgp--reference--group-001.md#canonical-68cae395ed3cd99ab78d3512891513fa80be87f1eb161a93e5207bca2edc9c55)
- [peers.routing_policies.route_policy.inbound](data-sources--bgp--reference--group-001.md#canonical-6efd32095da9547b02622f54386d4b2332e7b818a65f8cb8e72bb95404a14a67)
- [peers.routing_policies.route_policy.node_name](data-sources--bgp--reference--group-001.md#canonical-dcd89b92529b4a74150fe94816ed3414094e9a38a2e611c77fc6002773d43259)
- [peers.routing_policies.route_policy.object_refs](data-sources--bgp--reference--group-001.md#canonical-423bb5d56539eff47654011e39d8706d9b02a1a9c2350d5e93d96f4e848771d0)
- [peers.routing_policies.route_policy.outbound](data-sources--bgp--reference--group-001.md#canonical-be97d6057327dabbde80e3712ac86433974dd51937b19fa94b7c805079a69f52)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-68cae395ed3cd99ab78d3512891513fa80be87f1eb161a93e5207bca2edc9c55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0d4ccd4e6d56c38710bec2cdc3d6ec0ff118d30def0c48ebe2ed17e18cb4dd5"></a>

## peers.routing_policies.route_policy.all_nodes — peers.routing_policies.route_policy.all_nodes / 3abe61467816 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- peers.routing_policies.route_policy.all_nodes

<a id="canonical-755a8832117bc12df223d705c096a118dca965c436edd721cfd1dfd908fb01dd"></a>

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

<a id="canonical-e2d3f4ffd9132672fb0b02b1ea46d5d4f4c946b5e5fc0b44d68e11edddff0532"></a>

## Direct properties — peers.routing_policies.route_policy.all_nodes / 3abe61467816 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7cb18549ee04acada93e351d00253b428f254b1169fe9f50926d7391b26c9308"></a>

## Next pages — peers.routing_policies.route_policy.all_nodes / 3abe61467816 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-6efd32095da9547b02622f54386d4b2332e7b818a65f8cb8e72bb95404a14a67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85c10b75e6bff5d4ac78857210e08ef85f3779c1b4c868d66c56017b00a7cea1"></a>

## peers.routing_policies.route_policy.inbound — peers.routing_policies.route_policy.inbound / 10282fd6d081 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- peers.routing_policies.route_policy.inbound

<a id="canonical-7b813f25671252a4d46e159e9f570abed70b4179991d152249a6f6b63dafd2ed"></a>

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

<a id="canonical-6f778221e709454c516342facd51c73320826abb09156ddaf77938da7d5db49e"></a>

## Direct properties — peers.routing_policies.route_policy.inbound / 10282fd6d081 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9d680dbefb9ab7c4e32b14e22dfce5d19b2b82c8aff59f430f36a3621d363b0"></a>

## Next pages — peers.routing_policies.route_policy.inbound / 10282fd6d081 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-dcd89b92529b4a74150fe94816ed3414094e9a38a2e611c77fc6002773d43259"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f12e344023539da858d21f542fa6c0c881b02f0615b5a40a832214f5f312af8"></a>

## peers.routing_policies.route_policy.node_name — peers.routing_policies.route_policy.node_name / 10a907e0b8e3 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- peers.routing_policies.route_policy.node_name

<a id="canonical-a11df55948844dc9cf79361349a08b24aecf954c7bc741661bb779fa81ef3935"></a>

Type: `"single"`. Computed.

List of nodes on which BGP routing policy has to be applied.

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

<a id="canonical-d2a0b1e6aa81f90bf34334d2f740cebad196c009424f2acc7cf66734c4e0f191"></a>

## Direct properties — peers.routing_policies.route_policy.node_name / 10a907e0b8e3 / 3

<a id="canonical-93249c82950374e6aabf7dbf88abca01a727e58f0ae4f2a0621dc4dd8cff9c84"></a>

<a id="canonical-cdb68e5a047421b4fb6d7c43a711880b7705dbc4d0ea6d04003d1b63a4a4ce35"></a>

## node property — peers.routing_policies.route_policy.node_name / 10a907e0b8e3 / 4

Type: `["list", "string"]`. Computed.

Select BGP Session on which policy will be applied.

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

<a id="canonical-c79221d056a7372cb2f8deab4f7f20c859ecb7b9552f8317f5c2106f71d1784f"></a>

## Next pages — peers.routing_policies.route_policy.node_name / 10a907e0b8e3 / 5

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-423bb5d56539eff47654011e39d8706d9b02a1a9c2350d5e93d96f4e848771d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b508be3a1cb819d7c8cf31ae52b91eb2572e4809f5678a38063789f5fe6840fc"></a>

## peers.routing_policies.route_policy.object_refs — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- peers.routing_policies.route_policy.object_refs

<a id="canonical-951b1df316b6c4102b3c196e0a6f588b3ef3aa5b8f75a5d80011765fea6f723e"></a>

Type: `"list"`. Computed.

BGP routing policy. Select route policy to apply.

Upstream description:

Select route policy to apply.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-f3f73d600212fa87f56d92540049081b0e4bd57f6c4166ef56a4458561ef13af"></a>

## Direct properties — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 3

<a id="canonical-22f8363d572d2d76139b365d46d37e224578b1a8980e46827a490675ca14a4d8"></a>

<a id="canonical-4e5c7e3176c48baa6f7c68561deaf585e549654c00e0e8f388a7a88fa171cfb2"></a>

## kind property — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 4

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

<a id="canonical-f1b1f577e14993e2e1bf9e3175d1858bdd2cf1fdbc2540ca91b167e811f6f394"></a>

<a id="canonical-c6892ea417666e484ab8f7c979ec8e08e29ca4e7b2fa697a1aa0d35f7a952c0a"></a>

## name property — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 5

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

<a id="canonical-2628e6cf6fd2fc31185a794ed8a743d8746868bf81d81462be648085cea8b00d"></a>

<a id="canonical-cec72d118ed0313835a6790651eab995a91e7d371a8b2d65a684c18da95e4f8e"></a>

## namespace property — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 6

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

<a id="canonical-083dcbef2951b77165eab30f5d37b77de27776b444344530d05715575169dc94"></a>

<a id="canonical-b23122db8b7f8a7e2cf3dce7bc9956d08147abb13252bed589408b1634a5b077"></a>

## tenant property — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 7

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

<a id="canonical-6317dda7a094f8ad344e3e2a407c192fe86bef601a6f496ec87b9f0e6cf4cd92"></a>

<a id="canonical-dbb9cbea76aa1658da487b654d429b1e8b31baf4ea661f111b83ff02b2cbfd1f"></a>

## uid property — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 8

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

<a id="canonical-960e8cf866cdff457c93f706e4e0ee457e80b2b0484f3a60c8a2b4c25726487b"></a>

## Next pages — peers.routing_policies.route_policy.object_refs / 605f0b90bda2 / 9

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-be97d6057327dabbde80e3712ac86433974dd51937b19fa94b7c805079a69f52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a79c4401459b2b0e03181cb06abda62c09258389c1a4e7ceaf486fac69f7234a"></a>

## peers.routing_policies.route_policy.outbound — peers.routing_policies.route_policy.outbound / fbe89ad1ed84 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [peers](data-sources--bgp--reference--group-001.md#canonical-911887ab20c58816497ade96b820a77668537eda8fc9df94a555490a61f249e3)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-4b1e0185bbee5ee5190ddef279b5b9b52372d6050c9ed2dcf7a0958451d91cf9)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- peers.routing_policies.route_policy.outbound

<a id="canonical-ac765d76f43b33149faa6fcdfaa35b1e2f86793a67dcff4eb74210d2d95ff042"></a>

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

<a id="canonical-99ca74dcaa4bbfdfe89d1a7067480798c7f91a036c0dfaf08552e58f5d18dbc7"></a>

## Direct properties — peers.routing_policies.route_policy.outbound / fbe89ad1ed84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c813142bfc7dabdee2c53628895dea5a3f7503bf0bad1d591fe9a991bab70ee5"></a>

## Next pages — peers.routing_policies.route_policy.outbound / fbe89ad1ed84 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1d6fa0d9abd244529cd1bb6c162389ba531d347bca1e747e4f1160877030a327)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43d95b9492fe02dc6fe6c8e65a84082f0573d59e5ea59d31e3da8dcef3d1e84e"></a>

## where — where / 955014b19118 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- where

<a id="canonical-4e81baa6eeb88c0b26bcf084623f147bdef01637bbd3c8efc1a0458bcd50b3ca"></a>

Type: `"single"`. Computed.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Upstream description:

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-93a026f7beed96bdad878791a58d3fe6412f2e55232ede318067c3354893b4c9"></a>

## Direct properties — where / 955014b19118 / 3

- [site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d): complete subsection reference.

- [virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94): complete subsection reference.

<a id="canonical-07c5b4c1e0353bdf33a2c8ca47649d6bc8c6bc9594bc3e0c7339b7d4bf21ede0"></a>

## Next pages — where / 955014b19118 / 4

- [where.site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2241a087ec93218dd1e2bfd9b796ca51b6a60192a94f96f02b7a2b1ecb8c0123"></a>

## where.site — where.site / 011c1c966b81 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- where.site

<a id="canonical-1c40d4c4134206c17fa5e909475f2479952b34335dac1cdef07f164081190ad4"></a>

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

<a id="canonical-088a6fc3ba46b904da47fb96ed8d8c6c526b725af9652acd7070b57302dc5727"></a>

## Direct properties — where.site / 011c1c966b81 / 3

- [disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-da2df0f2cb8a7f032d8019ebada6b211e63566f5b7632521acdf293c463883b8): complete subsection reference.

- [enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-558639c9f1515bf4a02bb86912598f7432c206f3640542ae4b082b13a0c79a1f): complete subsection reference.

<a id="canonical-0784c2c7c7638af373b751e726ade2791ca889e63b116bc712bdb18504dcbdf2"></a>

<a id="canonical-068947713bdd35acc5f85b142d39bf7a5207e2c9d63bf601ef1bb475e422cf94"></a>

## network_type property — where.site / 011c1c966b81 / 4

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

- [ref](data-sources--bgp--reference--group-001.md#canonical-8e7648c22e32822d35aa09ad96c8811d9e11e582f9fa69df773ad36718d7f051): complete subsection reference.

<a id="canonical-e868653775dd7f87042543d97db26d79efe25ae11089e8a27476bf093fb0fa60"></a>

## Next pages — where.site / 011c1c966b81 / 5

- [where.site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-da2df0f2cb8a7f032d8019ebada6b211e63566f5b7632521acdf293c463883b8)
- [where.site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-558639c9f1515bf4a02bb86912598f7432c206f3640542ae4b082b13a0c79a1f)
- [where.site.ref](data-sources--bgp--reference--group-001.md#canonical-8e7648c22e32822d35aa09ad96c8811d9e11e582f9fa69df773ad36718d7f051)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-da2df0f2cb8a7f032d8019ebada6b211e63566f5b7632521acdf293c463883b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b51e203d5acf3eb255a6c52bd8df389aed13b48dfc65a46ebf5b81693ec6e1f"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / b64ec25e7a3a / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [where.site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d)
- where.site.disable_internet_vip

<a id="canonical-9572689a8469ed674c977a77482ed2bd3ba505f01041779a3cf9c6be9d350fea"></a>

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

<a id="canonical-92675f153deb6b9a172b2525ce71bb9dc29257f4ebb0bece16108efc2fdfa635"></a>

## Direct properties — where.site.disable_internet_vip / b64ec25e7a3a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b27b21bc9ad4643046439186b49b5ca5dfbebe57c343cdfd56d01a4e0d0b52fe"></a>

## Next pages — where.site.disable_internet_vip / b64ec25e7a3a / 4

- [where.site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-558639c9f1515bf4a02bb86912598f7432c206f3640542ae4b082b13a0c79a1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a8bb2471f93d62e5b7afd3786c6e28795d540cd54de5d78152587637c5db484"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / 072e77757d4b / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [where.site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d)
- where.site.enable_internet_vip

<a id="canonical-6b6d64fe971304d6394fdfed4028bc0e0ef960f62117a9a263af1b996b4df0d1"></a>

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

<a id="canonical-6f793cce4ba8a3ac92f4890bf24d02942fcbb6832c50601f6b14bf0a3869f590"></a>

## Direct properties — where.site.enable_internet_vip / 072e77757d4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8fb2210f4bb2f74bd1429b8518855b07b8928d9dbbfed954057e95a4b709bdd8"></a>

## Next pages — where.site.enable_internet_vip / 072e77757d4b / 4

- [where.site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-8e7648c22e32822d35aa09ad96c8811d9e11e582f9fa69df773ad36718d7f051"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d9f480ffcc30202a085c98788840b32737db81ec857d32d88efe6a3d691eaf1"></a>

## where.site.ref — where.site.ref / 2dd1ec4bb261 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [where.site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d)
- where.site.ref

<a id="canonical-e7b0ebcf7d759e1f246545d85eeab38aaad301159e3dda756f2c99dec43aa370"></a>

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

<a id="canonical-2a76f482ecbe43f71948f79affd53b33ff7a6150c0b2b0bd5aefc5418e742700"></a>

## Direct properties — where.site.ref / 2dd1ec4bb261 / 3

<a id="canonical-f4e4133a175d42a52f96d45887a55ea353ce086cc1b26c93e98d122d5a03e0f5"></a>

<a id="canonical-549e43370f9ec94f2bd530ee0a1206bc2a0cd574dae746812f9bc1e8e1b3a013"></a>

## kind property — where.site.ref / 2dd1ec4bb261 / 4

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

<a id="canonical-65631bd93a688302e4c6ba40e408f8203fa19cd6cb6343f81592a0e8b15db01f"></a>

<a id="canonical-94ac7d211694745f085991e7d6373ed2630d9831d1bc70440aa6171a30a67d15"></a>

## name property — where.site.ref / 2dd1ec4bb261 / 5

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

<a id="canonical-54d08e1f4b621ac6240a2ccdcc38a8c995c4135a7db9f7ffafdfe5b7e0132e36"></a>

<a id="canonical-65d30e10c21fc396278fffe30251cf9a4bc555d597e6f20cbc34da3616055e68"></a>

## namespace property — where.site.ref / 2dd1ec4bb261 / 6

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

<a id="canonical-d37937503b74b00b58a187f7193ba4b453177f398859e80bd67f7c1e951f3e08"></a>

<a id="canonical-cc90f7777df7550123e0b7edc28679db7932427bc8f7d21a1d69cb3c5ccb26db"></a>

## tenant property — where.site.ref / 2dd1ec4bb261 / 7

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

<a id="canonical-df0be0ae3a96c184df2df83b538eeabeeb9546ed90af3221e612a145be1eebbb"></a>

<a id="canonical-4f25e52febaa79bef50a723fb1b92edb84ad4a1bb6ef07323627c28140dfb98c"></a>

## uid property — where.site.ref / 2dd1ec4bb261 / 8

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

<a id="canonical-3a93e865deae61d601a0bbc6a74b5a05be9dfcea7d3e5178955a40a16f5571fc"></a>

## Next pages — where.site.ref / 2dd1ec4bb261 / 9

- [where.site](data-sources--bgp--reference--group-001.md#canonical-3ae8fb9ab3e1616b1787b6b2c68e8ff82e6bca22ce66c12fe579bb81f091b16d)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a0620181fbc79c7ba2326b5df31c32572e7292703a8387f9d4ce7803b3a2dab"></a>

## where.virtual_site — where.virtual_site / 0905f90ade6c / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- where.virtual_site

<a id="canonical-f84546fa45d82bbed6e9b76d31aff132327e692acf8cb9534c8e566d37d114f9"></a>

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

<a id="canonical-3dd35bd436ec7a4d58d4c7417d962cd63b387e2d6158e1ec1dcfcac718ef452d"></a>

## Direct properties — where.virtual_site / 0905f90ade6c / 3

- [disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-03ebce743965e4f1c1c286481f0967c87b80f9a83bf4808ee2e16b9a0fff873b): complete subsection reference.

- [enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-4d298ed16cf5d36b04714612c237f8f8ad0461a645fa255e0c81fb5bdf7732e7): complete subsection reference.

<a id="canonical-92768af09156e2ad40ace1e5435dbe815d54c5c079ff4f43c6a5a4fd9384c2c5"></a>

<a id="canonical-7573ecc7187c5f381e7299bbe321be1196dc29b6b1bc2277a5e98aacad144a66"></a>

## network_type property — where.virtual_site / 0905f90ade6c / 4

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

- [ref](data-sources--bgp--reference--group-001.md#canonical-8d56408d4ac7db4bc19a94f9a018050aba1cecdd8fe5cf9253b7017e74de94c1): complete subsection reference.

<a id="canonical-9abac90565f1f6a3fb78e44a05eb1e08ef07f3d727a1bf664154787ee1149ad2"></a>

## Next pages — where.virtual_site / 0905f90ade6c / 5

- [where.virtual_site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-03ebce743965e4f1c1c286481f0967c87b80f9a83bf4808ee2e16b9a0fff873b)
- [where.virtual_site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-4d298ed16cf5d36b04714612c237f8f8ad0461a645fa255e0c81fb5bdf7732e7)
- [where.virtual_site.ref](data-sources--bgp--reference--group-001.md#canonical-8d56408d4ac7db4bc19a94f9a018050aba1cecdd8fe5cf9253b7017e74de94c1)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-03ebce743965e4f1c1c286481f0967c87b80f9a83bf4808ee2e16b9a0fff873b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fa90aab502e3f5f24bfc3c73cf74c7d46c0b4aa92e6f83bc6cb613d9608a306"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / 20d88907b17a / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94)
- where.virtual_site.disable_internet_vip

<a id="canonical-dadf99503d102c067a09197240f09b0968ea9e5bb68c5d4bc813c88d6ff0c841"></a>

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

<a id="canonical-2d76a4886bb14ccd81fb3bb2e3de7f343be86c08b629727748250704bc96a168"></a>

## Direct properties — where.virtual_site.disable_internet_vip / 20d88907b17a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15299d78dd20ba6292693ba39003fb03bbe5c8ebcd79766a5fbdc1bbb092aac5"></a>

## Next pages — where.virtual_site.disable_internet_vip / 20d88907b17a / 4

- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-4d298ed16cf5d36b04714612c237f8f8ad0461a645fa255e0c81fb5bdf7732e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-765f054fe57ef10883211aaa8ddec2733a8d1058462e95868dc6fddce4ec7e1b"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / 28d9f5ae5c94 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94)
- where.virtual_site.enable_internet_vip

<a id="canonical-a5dafce89d557981b643d2e5396232b66db5ec80f61e0a84f8cb54501821b266"></a>

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

<a id="canonical-055d8e01f0a4e3d0307c287b4436f21a878f8153da704e5e2bb9507286b89468"></a>

## Direct properties — where.virtual_site.enable_internet_vip / 28d9f5ae5c94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e75b4cd712e4f4bebd5f28f19c1b41aa993512caad39bbcb71d240cba5f6d53a"></a>

## Next pages — where.virtual_site.enable_internet_vip / 28d9f5ae5c94 / 4

- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-8d56408d4ac7db4bc19a94f9a018050aba1cecdd8fe5cf9253b7017e74de94c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf57fa3e1aa7e2217ef9a6d18fb022d5d9ce96f62da79c68a92e1ae62bab548a"></a>

## where.virtual_site.ref — where.virtual_site.ref / 1fc7abe3df01 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-9f0b0b9a966f6df2c2702c17a4a3fe01c00f56a4879ae13620d620958b4ad058)
- [where](data-sources--bgp--reference--group-001.md#canonical-30f9f9c8e6b341f1e60d20e8bd6d99c9202590bbfff6c358ce90d8fd4c026098)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94)
- where.virtual_site.ref

<a id="canonical-14eea49c909a5aa0c7d046688186d3839bd3c892b0428fce86db9784cbf69d46"></a>

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

<a id="canonical-f61a714d42eb0d61ffd4feb6af0c89967782e0ba6234fe091ed659c5b1436b8c"></a>

## Direct properties — where.virtual_site.ref / 1fc7abe3df01 / 3

<a id="canonical-d53303aabe90c27a49d068782db83205c635b739395539f2c061cd77eb00a795"></a>

<a id="canonical-fee00de17d9ff6e54abdfaa840ec6716e6b5423d6fb868fc9427b109d1f7c26b"></a>

## kind property — where.virtual_site.ref / 1fc7abe3df01 / 4

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

<a id="canonical-77072e8dfa40eef10d328a2c22fef2647b5b6ceb2d1de28ab24178b6b6038926"></a>

<a id="canonical-3924fde9cdad2b47c4fc71f6ac947301c8c89290ad7d18e738670f3340340f06"></a>

## name property — where.virtual_site.ref / 1fc7abe3df01 / 5

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

<a id="canonical-78660e8d03b6d04ded3a2c6cfd87a56dea65f88f039d844e86e81b81e1f82986"></a>

<a id="canonical-4c2fb79b083bf7ad3e2bf0951ad77431b9d9b0c2d534c1b21020b0199c7fc854"></a>

## namespace property — where.virtual_site.ref / 1fc7abe3df01 / 6

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

<a id="canonical-0fcbdb094432f433ceedeccff8bf303b211339e4fdf2c97ab6f87311064ff511"></a>

<a id="canonical-8eeba2af051c81c4df3ea85d122ce80c32f27cebed92bf75ea8c1166d8ea6bb0"></a>

## tenant property — where.virtual_site.ref / 1fc7abe3df01 / 7

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

<a id="canonical-09df79cc9194e00a714c9db505856cf503c67267354bf75907c49dfae4b55547"></a>

<a id="canonical-955169a5bdb518cc9bd156e3e127ca4e5e6cb534c1b0fe5e0c3fb152d5356299"></a>

## uid property — where.virtual_site.ref / 1fc7abe3df01 / 8

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

<a id="canonical-03dcfda3a06a7ff3c5556fc1a1384864793929090df4d60df09e0fae3cb9d978"></a>

## Next pages — where.virtual_site.ref / 1fc7abe3df01 / 9

- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-016d7bb89f593fb8d945a755ce651d08c2f5ec6cedcff6315840b3ef84f18a94)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
