---
page_title: "xcsh_nat_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy reference."
---

# xcsh_nat_policy reference

<a id="canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50f28adfe080d7bacee6023cd37ad1d610eab872412d18b1f4e1c42148095ddc"></a>

## Property reference — Property reference / d6ddf92779d9 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- Property reference

<a id="canonical-a8cc0083338ec6f9692277e9b66e98595e2b90ef7e82de752fb41b40258da317"></a>

## Direct properties — Property reference / d6ddf92779d9 / 3

<a id="canonical-3e6ae656d8e0125a0a7b4c96204cde95876797a5cb5b5461dabbc249cefd50ee"></a>

<a id="canonical-57baeb64259c166ca7c81d67d39fff18a71b19e74394c18d8547ec98febd6e33"></a>

## annotations property — Property reference / d6ddf92779d9 / 4

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

<a id="canonical-44a3e688396da8758c99d2ad6343138dda56ab2a087e0be422a94766b9c433a9"></a>

<a id="canonical-47903d069b6fc31814c9dca552ebf509dc8c49c58fdc6b9a54a2c53d42333f7d"></a>

## description property — Property reference / d6ddf92779d9 / 5

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

<a id="canonical-1da1f722eab7d40150f7b6a52bc61f9d63ddeb512fb86e2cc9304af0d68fe232"></a>

<a id="canonical-695952e08594f3492ebee5d12a287937dc8f78a4be5e46ca7226eded31624cb1"></a>

## disable property — Property reference / d6ddf92779d9 / 6

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

<a id="canonical-78230befe56e2dbdde352bca8f63b2f4ef918d33a1d3ec2d3b60d74e7e71cd9e"></a>

<a id="canonical-6e7e2ac6e96f0c08f5600589051874e9fae1340412f3fe20a970bb1a74d53d6f"></a>

## id property — Property reference / d6ddf92779d9 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-5db641a05e727f73a9545ae6a7e15415200ac8146edd8424f1537098a3b42f60"></a>

<a id="canonical-0101f2ad0c406881ba94add6b412ef789e4935207300388ab15ed2eeb0aaa034"></a>

## labels property — Property reference / d6ddf92779d9 / 8

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

<a id="canonical-6c9b47074082ddd0bc2c2ecc1df62a5cf3ecbfaa40bb5916c1dd76923474f194"></a>

<a id="canonical-cd719cd7ceb24872441d86391528dee405cc9d0d8c5c4e1aaf9cfbb7bcce8f3c"></a>

## name property — Property reference / d6ddf92779d9 / 9

Type: `"string"`. Required.

Name of the NAT Policy. Must be unique within the namespace.

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

<a id="canonical-a0ae5952bb64d36c1528dd2a5c3665ce9fda490a0a7bbf78158b9af1c7e8271a"></a>

<a id="canonical-abc4d7603237b5045f3de82b3c047cf1a9bd9966fc670edf0e7e581ad46b55ae"></a>

## namespace property — Property reference / d6ddf92779d9 / 10

Type: `"string"`. Required.

Namespace where the NAT Policy is created.

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

- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e): complete subsection reference.

- [site](resources--nat_policy--reference--group-001.md#canonical-c1aea168038d7fa6fe8f370374af11bfab68a850b95080a918405e78886cc2ec): complete subsection reference.

- [timeouts](resources--nat_policy--reference--group-001.md#canonical-3ed37e126f8c1e301894d45191d3eae175054a1ae68b7362934f480e7b2f1aee): complete subsection reference.

<a id="canonical-3b86538160b45af43fbc280fe13864ac4744b78e9d03340866821df248489c79"></a>

## All schema paths — Property reference / d6ddf92779d9 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nat_policy--reference--group-001.md#canonical-3e6ae656d8e0125a0a7b4c96204cde95876797a5cb5b5461dabbc249cefd50ee) |
| `description` | [description](resources--nat_policy--reference--group-001.md#canonical-44a3e688396da8758c99d2ad6343138dda56ab2a087e0be422a94766b9c433a9) |
| `disable` | [disable](resources--nat_policy--reference--group-001.md#canonical-1da1f722eab7d40150f7b6a52bc61f9d63ddeb512fb86e2cc9304af0d68fe232) |
| `id` | [id](resources--nat_policy--reference--group-001.md#canonical-78230befe56e2dbdde352bca8f63b2f4ef918d33a1d3ec2d3b60d74e7e71cd9e) |
| `labels` | [labels](resources--nat_policy--reference--group-001.md#canonical-5db641a05e727f73a9545ae6a7e15415200ac8146edd8424f1537098a3b42f60) |
| `name` | [name](resources--nat_policy--reference--group-001.md#canonical-6c9b47074082ddd0bc2c2ecc1df62a5cf3ecbfaa40bb5916c1dd76923474f194) |
| `namespace` | [namespace](resources--nat_policy--reference--group-001.md#canonical-a0ae5952bb64d36c1528dd2a5c3665ce9fda490a0a7bbf78158b9af1c7e8271a) |
| `rules` | [rules](resources--nat_policy--reference--group-001.md#canonical-67b0e61c7484e2aca4248c598ac743ff369284a302067ffc0e36f8b8f5a6f6b9) |
| `rules.action` | [rules.action](resources--nat_policy--reference--group-001.md#canonical-de39c0db4b49f194f09d5f614ac6ed575df1b330b500a7e0cb803570a22c3d1e) |
| `rules.action.dynamic` | [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-8af90d665370e19e4a8081c7d5ab8cd6620b0e2a5996d612a10e95d279b655af) |
| `rules.action.dynamic.elastic_ips` | [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-5c158c93d81045088733ef79d3fce5434bff4d9bdc2bc402dea5b96117ad0b08) |
| `rules.action.dynamic.elastic_ips.refs` | [rules.action.dynamic.elastic_ips.refs](resources--nat_policy--reference--group-001.md#canonical-ec09fd6d84490617578ed31440a349e376dbbc95bacd5973aae6fcc6251414a8) |
| `rules.action.dynamic.elastic_ips.refs.kind` | [rules.action.dynamic.elastic_ips.refs.kind](resources--nat_policy--reference--group-001.md#canonical-579ac0fc0dc3d7c6e03c81babc6c58d259a204889e4e6b4ac91fbb526573bdce) |
| `rules.action.dynamic.elastic_ips.refs.name` | [rules.action.dynamic.elastic_ips.refs.name](resources--nat_policy--reference--group-001.md#canonical-672951287156a507a3a25cea199bd8d0c85a4f3caa4c2d192c5d020bb90ae2a4) |
| `rules.action.dynamic.elastic_ips.refs.namespace` | [rules.action.dynamic.elastic_ips.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-6bb4fd04113d09034dd7f39bef30e9f3cfb38f2d3abe7fee775dd883d088413c) |
| `rules.action.dynamic.elastic_ips.refs.tenant` | [rules.action.dynamic.elastic_ips.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-49de7c06c328c1b8ed5482d707fba552299ccdec05c3e4f73853c799ff3b4768) |
| `rules.action.dynamic.elastic_ips.refs.uid` | [rules.action.dynamic.elastic_ips.refs.uid](resources--nat_policy--reference--group-001.md#canonical-3a12d5d32bc5eec2c8af5db2e4ecc848338bf134b23a521b8c3758c9a39eb2d8) |
| `rules.action.dynamic.pools` | [rules.action.dynamic.pools](resources--nat_policy--reference--group-001.md#canonical-bf68e6f59e6e285b022d991d8529ff64270488d32f798448406a4e83671964fe) |
| `rules.action.dynamic.pools.prefixes` | [rules.action.dynamic.pools.prefixes](resources--nat_policy--reference--group-001.md#canonical-dd30acacb901b59b025d14fa937550585c863bf2a9b9b3cb0d7792a5a7533860) |
| `rules.action.virtual_cidr` | [rules.action.virtual_cidr](resources--nat_policy--reference--group-001.md#canonical-63819da91329ce431e21d8b017d4819aec72542db333fc908b0b683389ad78a6) |
| `rules.cloud_connect` | [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-55d8d7c0079a69aff492b83d3724f4df590c7bc0f0f03a26b7a905d9ccd34f32) |
| `rules.cloud_connect.refs` | [rules.cloud_connect.refs](resources--nat_policy--reference--group-001.md#canonical-7505bfc803ec07588b5bc5890b2fbf9068696e916e754ca6529229ef1200a504) |
| `rules.cloud_connect.refs.kind` | [rules.cloud_connect.refs.kind](resources--nat_policy--reference--group-001.md#canonical-b44827260e0760976aba78eab6ca001769b7f6f13c0f3f5edb0154c168cc5cb6) |
| `rules.cloud_connect.refs.name` | [rules.cloud_connect.refs.name](resources--nat_policy--reference--group-001.md#canonical-0206dac9358378e5cff38bb3407a30d3e52d86fca0cbcfb2c788b4a506080f65) |
| `rules.cloud_connect.refs.namespace` | [rules.cloud_connect.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-097af11a2c534b37280b3dee5e376d635c58489ffff26ff838000b44ba847e61) |
| `rules.cloud_connect.refs.tenant` | [rules.cloud_connect.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-915f4bfb6f228b47347cb8aaa60eb393013791bba853268395c60b74fdad7730) |
| `rules.cloud_connect.refs.uid` | [rules.cloud_connect.refs.uid](resources--nat_policy--reference--group-001.md#canonical-94e10212c1024f37b2096f204c6998231484ef6fb569fba1754cd57184c53c04) |
| `rules.criteria` | [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-7dd48863468c787aa4452446dfa54400c7d1f848cc02182c2af7d04d08608d0b) |
| `rules.criteria.any` | [rules.criteria.any](resources--nat_policy--reference--group-001.md#canonical-1f7ac2a14cb88f41e55978be025d55724a3e6c3f23366eb5f7e56b0d8d2d7db5) |
| `rules.criteria.destination_cidr` | [rules.criteria.destination_cidr](resources--nat_policy--reference--group-001.md#canonical-cbe6715367bf33a4afc26d66f85cd58b8ffe7d4683466421d58c9b880e561662) |
| `rules.criteria.icmp` | [rules.criteria.icmp](resources--nat_policy--reference--group-001.md#canonical-e9a833ad5eb92acb80588cf64f271dd607f449a5df61346ef2a051b3b2dac47a) |
| `rules.criteria.site_local_inside_network` | [rules.criteria.site_local_inside_network](resources--nat_policy--reference--group-001.md#canonical-12599a852d474c0ee740a5a6d1f6f55aa14a65a21f62886d9d2422e98b947322) |
| `rules.criteria.site_local_network` | [rules.criteria.site_local_network](resources--nat_policy--reference--group-001.md#canonical-e31e2c6f06af7fc193c9dabef9b8c0e909bd2f587d46bdbc532b365fc8ff583f) |
| `rules.criteria.source_cidr` | [rules.criteria.source_cidr](resources--nat_policy--reference--group-001.md#canonical-ad8cb6dd7c3312510d621d12b9ea617c5ac29b75121eeed669020cc081aeab67) |
| `rules.criteria.tcp` | [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-24a5bf055d8b6110dccae6efef824388acc298bd61e290e67470adadced827e8) |
| `rules.criteria.tcp.destination_port` | [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-4ffbd20d22809f74863e0c13732eb5e07db8427daa6fc25a755fc18dbbd549c3) |
| `rules.criteria.tcp.destination_port.no_port_match` | [rules.criteria.tcp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-d51ba79caa86b577c0b9e39c1bcbafb4020ad26619ba0b4bcaaedefee5ee8435) |
| `rules.criteria.tcp.destination_port.port` | [rules.criteria.tcp.destination_port.port](resources--nat_policy--reference--group-001.md#canonical-d29fbcd3fb4b28716a6b07650385a5c16a304979c2813b4b0a1711b591de21bd) |
| `rules.criteria.tcp.destination_port.port_ranges` | [rules.criteria.tcp.destination_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-2957e7c9a14fc8984059c551f399cff22a2ef1f44a73494168bd2e0590a4edf7) |
| `rules.criteria.tcp.source_port` | [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-d2e0950161e1e63058d13a7f6df36207da2aae283c0f88d343bff16bc1e87785) |
| `rules.criteria.tcp.source_port.no_port_match` | [rules.criteria.tcp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-fab26929fc18b6adb0cee84a8b2405aede61e0fd9075503f4acdc511f3bee12c) |
| `rules.criteria.tcp.source_port.port` | [rules.criteria.tcp.source_port.port](resources--nat_policy--reference--group-001.md#canonical-1fc713c34b290a7207e864eb1756535ba761f14e583e3e88afebae0fc4478bed) |
| `rules.criteria.tcp.source_port.port_ranges` | [rules.criteria.tcp.source_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-0e4557d77924def2847802d78845720842aab283839a0df8e3e976abef0bd16a) |
| `rules.criteria.udp` | [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2b4960baf77e3f53a6ec09723fdacd831427aa76dd2af28e3d8c5b6f24781dee) |
| `rules.criteria.udp.destination_port` | [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-a795621116f8a2e03ace2d01af4d78512ccee5a97223c832ce36e8cbabe673b9) |
| `rules.criteria.udp.destination_port.no_port_match` | [rules.criteria.udp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-4ff9d52e91de016e0bc6ba1db04648d73adae9472d0d2ec86b93ac99101f1882) |
| `rules.criteria.udp.destination_port.port` | [rules.criteria.udp.destination_port.port](resources--nat_policy--reference--group-001.md#canonical-3e05f276fd4b764af633c5f301cfd8ec8178a4702d4116e07a9e138a5db00595) |
| `rules.criteria.udp.destination_port.port_ranges` | [rules.criteria.udp.destination_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-1586c28a18145a9dae55532b2f0c085a62c1dd0543592423b5436a94e9385bc5) |
| `rules.criteria.udp.source_port` | [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-b4977be6d64cdf8f95bb49bd5a5f48ebadb9960c382f9824dff857dfceeefe81) |
| `rules.criteria.udp.source_port.no_port_match` | [rules.criteria.udp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-e412ab6f608456fa69e649d824c11470c4c0d16217e74102b25781efb9845955) |
| `rules.criteria.udp.source_port.port` | [rules.criteria.udp.source_port.port](resources--nat_policy--reference--group-001.md#canonical-ed864ca1826eb6e38747e55e906734e0daea224ab5ef624332617d5d9e32df93) |
| `rules.criteria.udp.source_port.port_ranges` | [rules.criteria.udp.source_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-6d8d96502edfec0b1f6d86a3d63e62e4bca3f0490dc9d8fb9ffcc2c30d203335) |
| `rules.disable_spec` | [rules.disable_spec](resources--nat_policy--reference--group-001.md#canonical-6c85d7ac365c0b6539968523103d68b87d0f47da4879d2bac1b947c8c69ca892) |
| `rules.enable` | [rules.enable](resources--nat_policy--reference--group-001.md#canonical-95ebc55c6aaa91690375673ec7d62fac8a9fac239246b0dc404e2ac9c6c58cb1) |
| `rules.name` | [rules.name](resources--nat_policy--reference--group-001.md#canonical-862720555132537dba9c181ee7c26ebced230f601e44c1979be9ef2514e31a5f) |
| `rules.node_interface` | [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-38fc84da28372aa048be9c3f9cb3f90d481ba83b41143f80a5bd92519276409b) |
| `rules.node_interface.list` | [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-000626cb7493cab71930acb06a06c5c3b3f460c3fb9f0ab2764f819f615676c4) |
| `rules.node_interface.list.interface` | [rules.node_interface.list.interface](resources--nat_policy--reference--group-001.md#canonical-b69d82cb89ef323b338bc6d4cb0674d4dc8b94a69f846685f72ccd71fe0f89b6) |
| `rules.node_interface.list.interface.kind` | [rules.node_interface.list.interface.kind](resources--nat_policy--reference--group-001.md#canonical-aa44cd1026f67d8c8fe362af873ec75642044ea03bd4c2148e3c65803713cc56) |
| `rules.node_interface.list.interface.name` | [rules.node_interface.list.interface.name](resources--nat_policy--reference--group-001.md#canonical-3f1f998291f1b08e5fce5a2c9ebc43cb5f2c5375b2e7c1adf6c9421c8c400be9) |
| `rules.node_interface.list.interface.namespace` | [rules.node_interface.list.interface.namespace](resources--nat_policy--reference--group-001.md#canonical-d87bdca2131abf13b75075645d9f2968528b32f4acb02f62c79c02c5e038e48f) |
| `rules.node_interface.list.interface.tenant` | [rules.node_interface.list.interface.tenant](resources--nat_policy--reference--group-001.md#canonical-bd1faef0aa8db403283d5a298aac88ed70dcf791b0928189ec4aac00accbaa3d) |
| `rules.node_interface.list.interface.uid` | [rules.node_interface.list.interface.uid](resources--nat_policy--reference--group-001.md#canonical-36a7503d202c8b3753952919f378f3bbb85d066bb950d2d632a5b65b926360aa) |
| `rules.node_interface.list.node` | [rules.node_interface.list.node](resources--nat_policy--reference--group-001.md#canonical-badaa17ff37aed10801adb03ba182938a074ce468388dd7b157be0def0daf509) |
| `rules.segment` | [rules.segment](resources--nat_policy--reference--group-001.md#canonical-420fbae347cff4e30d1043f0a651189970e1a867edbe6f5a6645116228ebda85) |
| `rules.segment.refs` | [rules.segment.refs](resources--nat_policy--reference--group-001.md#canonical-3e7df79a11835050cacfa90d8bae3c956bb54b6f5a4e8293569d4ef3c2bbc627) |
| `rules.segment.refs.kind` | [rules.segment.refs.kind](resources--nat_policy--reference--group-001.md#canonical-73fd8c736cb4c7597b81b0af58ef0be5f98ffbbc46a6ae04904e238b8312b899) |
| `rules.segment.refs.name` | [rules.segment.refs.name](resources--nat_policy--reference--group-001.md#canonical-62c0c92aead1783e4d31b92b101153ef1c0f4a1f7abf74076506a85c64a6ac71) |
| `rules.segment.refs.namespace` | [rules.segment.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-f68e11041fc927edc6d3616fc155b31615cd1cbfabd134739c98d5532de0b97d) |
| `rules.segment.refs.tenant` | [rules.segment.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-c57bc6329229b91ecc887f4222618c183ddd088d97ac80fdaf6932cac3cd786f) |
| `rules.segment.refs.uid` | [rules.segment.refs.uid](resources--nat_policy--reference--group-001.md#canonical-8207de645a714450d7182f79f25917df24cfb6a03ccf899fbb983fce44c4cfd7) |
| `rules.virtual_network` | [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-421b524e49ff918eac8ade720a2512a4ee4528d9384ac48256d1b77e4fbc9dad) |
| `rules.virtual_network.refs` | [rules.virtual_network.refs](resources--nat_policy--reference--group-001.md#canonical-aa1793e4587086d8c25912cb1cdccaad76893637300906117e79ef188e29617c) |
| `rules.virtual_network.refs.kind` | [rules.virtual_network.refs.kind](resources--nat_policy--reference--group-001.md#canonical-5779b6653355a3885de3e901734c392585fa4cf65d278846067b90426fbae516) |
| `rules.virtual_network.refs.name` | [rules.virtual_network.refs.name](resources--nat_policy--reference--group-001.md#canonical-60d99a09e1d74c96b9b8073c6f4630a8f7bf72c09d56947b66fec0fbf648053a) |
| `rules.virtual_network.refs.namespace` | [rules.virtual_network.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-7dc9b077b3b4e5c75ad55d7b89dba23cafbe8723a932d1b0a9e1fe121ac055cf) |
| `rules.virtual_network.refs.tenant` | [rules.virtual_network.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-b51fdf3aadc36275d5902f7836719c19d95270797d39a3cc30ffcaee7202236c) |
| `rules.virtual_network.refs.uid` | [rules.virtual_network.refs.uid](resources--nat_policy--reference--group-001.md#canonical-0294a10e5de3ad08d91b28babe468f24229ec50c243486effb102e54ceedb136) |
| `site` | [site](resources--nat_policy--reference--group-001.md#canonical-8c4f0620116908cdcf7b9e80c2f041ae90a21e7caf6926958e2c0357aa72bd7d) |
| `site.refs` | [site.refs](resources--nat_policy--reference--group-001.md#canonical-a819e1b4a5787310556ee0ef25c488024b0bef665ad5ce704c227f2fde9ed533) |
| `site.refs.kind` | [site.refs.kind](resources--nat_policy--reference--group-001.md#canonical-355875b6f404341a01d789d54b7f182370a7cbdb482f08ba731c1d180144550a) |
| `site.refs.name` | [site.refs.name](resources--nat_policy--reference--group-001.md#canonical-783da7404918f7eacddfc07424d658a8e8b6c06c470f0755369269f063511bca) |
| `site.refs.namespace` | [site.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-57376f43fff4cffcba45ddaa0bc2d795f6640994076298bf1d84fe8113807366) |
| `site.refs.tenant` | [site.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-a8d6a4eab9db74b977d2b194119897e76d0963d605ebc27b31e8e27cb18b1f1b) |
| `site.refs.uid` | [site.refs.uid](resources--nat_policy--reference--group-001.md#canonical-d26fd94fb6fff9f202ce38845a42206ecc8c74f5183ac54585b1038018742c5e) |
| `timeouts` | [timeouts](resources--nat_policy--reference--group-001.md#canonical-5b02e38455b556b08cd7a472b909dca2be252b51e935d25b0aa939b4125e0812) |
| `timeouts.create` | [timeouts.create](resources--nat_policy--reference--group-001.md#canonical-45e888cc573ea621ee4b252a424d33f66db3823e823000826cb1afd351db819a) |
| `timeouts.delete` | [timeouts.delete](resources--nat_policy--reference--group-001.md#canonical-74ef950a7c48cc18ab85f95a71e886b88fab27dfaff68cef2c69056efc11419e) |
| `timeouts.read` | [timeouts.read](resources--nat_policy--reference--group-001.md#canonical-142eaa4f0281250bcc375b4363422a48e6c8d247b40fa99a75f35c821f5e758b) |
| `timeouts.update` | [timeouts.update](resources--nat_policy--reference--group-001.md#canonical-26b5bda5697c8cfbd69adbaf3ae38307be5a32ac542d19288953b025e7d08ce6) |

<a id="canonical-7fb176abec19e1cd8361b7efdc8d652a91c4d994e279948feaa5a5a4d99bfbf5"></a>

## Next pages — Property reference / d6ddf92779d9 / 12

- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [site](resources--nat_policy--reference--group-001.md#canonical-c1aea168038d7fa6fe8f370374af11bfab68a850b95080a918405e78886cc2ec)
- [timeouts](resources--nat_policy--reference--group-001.md#canonical-3ed37e126f8c1e301894d45191d3eae175054a1ae68b7362934f480e7b2f1aee)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56e832ef02915b29a453d398cd06e6f6e5c9b93f9ec0567f6f3c1b834f77c599"></a>

## rules — rules / ee0a5b9e2220 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- rules

<a id="canonical-67b0e61c7484e2aca4248c598ac743ff369284a302067ffc0e36f8b8f5a6f6b9"></a>

Type: `"object"`. list nested block, Optional.

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "node_interface"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "segment"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("disable_spec",
    "enable"),
  validators.ConflictingListObjectAttributes("node_interface",
    "segment"),
  validators.ConflictingListObjectAttributes("node_interface",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("segment",
    "virtual_network")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-54f28f020248b0059747fc3c566c81d0a01a75cefc0668b1d8cf288e7c98eb39"></a>

## Direct properties — rules / ee0a5b9e2220 / 3

- [action](resources--nat_policy--reference--group-001.md#canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66): complete subsection reference.

- [cloud_connect](resources--nat_policy--reference--group-001.md#canonical-4342fc9a3d91bb76d1a863120502c05f2c0913d3fcbf502d9f2184ff8c2714e5): complete subsection reference.

- [criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a): complete subsection reference.

- [disable_spec](resources--nat_policy--reference--group-001.md#canonical-bf54ee0daa4be07c6891585b45687b03ed6fc1e9d3784be804a7c14e20716e50): complete subsection reference.

- [enable](resources--nat_policy--reference--group-001.md#canonical-f180a8ae475eb7bf551c7e13e1734b9a6f2fce42081c03831911b5d5a32f2376): complete subsection reference.

<a id="canonical-862720555132537dba9c181ee7c26ebced230f601e44c1979be9ef2514e31a5f"></a>

<a id="canonical-024703a0a451c537c2f210e007cfb630886da79866d93cc56b60c78e1bc4ad81"></a>

## name property — rules / ee0a5b9e2220 / 4

Type: `"string"`. Optional.

Name. Name of the Rule.

Upstream description:

Name of the Rule.

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

- [node_interface](resources--nat_policy--reference--group-001.md#canonical-0368e87c2a3dd0f0db2452d2d3d60fab5e73b8386b388ae31ce6ed3019df8c28): complete subsection reference.

- [segment](resources--nat_policy--reference--group-001.md#canonical-80f720e5966bcdaa2dae3cdbbdfb1e5b4b4dfd78584937c7cd88946eb93daeac): complete subsection reference.

- [virtual_network](resources--nat_policy--reference--group-001.md#canonical-63554196d15ca33c0fd89f0529700793810d33309c458fbae48de4234f3a28a1): complete subsection reference.

<a id="canonical-e5a8b2f1d86394834cb9a86da35824bcee4ae8b64bc9067d633f7bf62dbb3725"></a>

## Next pages — rules / ee0a5b9e2220 / 5

- [rules.action](resources--nat_policy--reference--group-001.md#canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66)
- [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-4342fc9a3d91bb76d1a863120502c05f2c0913d3fcbf502d9f2184ff8c2714e5)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.disable_spec](resources--nat_policy--reference--group-001.md#canonical-bf54ee0daa4be07c6891585b45687b03ed6fc1e9d3784be804a7c14e20716e50)
- [rules.enable](resources--nat_policy--reference--group-001.md#canonical-f180a8ae475eb7bf551c7e13e1734b9a6f2fce42081c03831911b5d5a32f2376)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0368e87c2a3dd0f0db2452d2d3d60fab5e73b8386b388ae31ce6ed3019df8c28)
- [rules.segment](resources--nat_policy--reference--group-001.md#canonical-80f720e5966bcdaa2dae3cdbbdfb1e5b4b4dfd78584937c7cd88946eb93daeac)
- [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-63554196d15ca33c0fd89f0529700793810d33309c458fbae48de4234f3a28a1)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2143720be10e5112156fc2d36448855ed44af5f01a998bab90bc8bd462b9b9b6"></a>

## rules.action — rules.action / f7b961fce263 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.action

<a id="canonical-de39c0db4b49f194f09d5f614ac6ed575df1b330b500a7e0cb803570a22c3d1e"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dynamic",
    "virtual_cidr")}
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
  "x-ves-oneof-field-source_nat_choice": "[\"dynamic\",\"virtual_cidr\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0ac8352832c4916803f6bb453db42775f23456f53e58c61a02d9198623fb44e"></a>

## Direct properties — rules.action / f7b961fce263 / 3

- [dynamic](resources--nat_policy--reference--group-001.md#canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606): complete subsection reference.

<a id="canonical-63819da91329ce431e21d8b017d4819aec72542db333fc908b0b683389ad78a6"></a>

<a id="canonical-e00358253f7551738dd585ae18c2f13ea3e91407edc441597246bb3adcd50d11"></a>

## virtual_cidr property — rules.action / f7b961fce263 / 4

Type: `"string"`. Optional.

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR..

Upstream description:

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR
192.0.2.0/24, the virtual CIDR has 100.100.100.0/24.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-3ba2c41e69a54678ce58d854b6504d1f83843e829d7982b77b374bd7d03427b5"></a>

## Next pages — rules.action / f7b961fce263 / 5

- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab12aa65300b7b70106a69e1c77cbd6b28d0cf78ded6a26468e5776edc00fc73"></a>

## rules.action.dynamic — rules.action.dynamic / 16268b600882 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66)
- rules.action.dynamic

<a id="canonical-8af90d665370e19e4a8081c7d5ab8cd6620b0e2a5996d612a10e95d279b655af"></a>

Type: `"object"`. single nested block, Optional.

Dynamic Pool. Dynamic Pool Configuration.

Upstream description:

Dynamic Pool Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("elastic_ips",
    "pools")}
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
  "x-ves-oneof-field-pool_choice": "[\"elastic_ips\",\"pools\"]"
}
```

Terraform syntax:

```terraform
dynamic {
  # Configure direct properties listed below.
}
```

<a id="canonical-a5713367ab2a7d4b1a02d1269e68b1dde38a325e303e54aa6b858607863ca504"></a>

## Direct properties — rules.action.dynamic / 16268b600882 / 3

- [elastic_ips](resources--nat_policy--reference--group-001.md#canonical-78f86b1b602fc376c4979b1336a1c8fb492d4bd07e180d3c3ae320eb8e7fac83): complete subsection reference.

- [pools](resources--nat_policy--reference--group-001.md#canonical-d1f08b997061ce8f5880e46c0d16a54ed0e4edc6f059e4603f2549d6b06c9e12): complete subsection reference.

<a id="canonical-167980f4bf106acb61b73c2755d6aa7dc2fbb649da9b4cd600f4832ea1b4fd21"></a>

## Next pages — rules.action.dynamic / 16268b600882 / 4

- [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-78f86b1b602fc376c4979b1336a1c8fb492d4bd07e180d3c3ae320eb8e7fac83)
- [rules.action.dynamic.pools](resources--nat_policy--reference--group-001.md#canonical-d1f08b997061ce8f5880e46c0d16a54ed0e4edc6f059e4603f2549d6b06c9e12)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-78f86b1b602fc376c4979b1336a1c8fb492d4bd07e180d3c3ae320eb8e7fac83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35a6ed6e65e2f4a0299dc1e5074996f8f7f27dfa576fccd71f3145310a379a91"></a>

## rules.action.dynamic.elastic_ips — rules.action.dynamic.elastic_ips / d582ca747972 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606)
- rules.action.dynamic.elastic_ips

<a id="canonical-5c158c93d81045088733ef79d3fce5434bff4d9bdc2bc402dea5b96117ad0b08"></a>

Type: `"object"`. single nested block, Optional.

List of references to Cloud Elastic IP Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
elastic_ips {
  # Configure direct properties listed below.
}
```

<a id="canonical-a319f0e050738895d8905c84187aa7534633864daadc82521a850bcc987a7f8f"></a>

## Direct properties — rules.action.dynamic.elastic_ips / d582ca747972 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-c9ceb5c9e8224434093da60aa1141f1178a401280f796835cd89f1614f66574b): complete subsection reference.

<a id="canonical-4df9a3b826c01f47232f7373c734efac458bcd079db35c5435ae4afabf4e6f83"></a>

## Next pages — rules.action.dynamic.elastic_ips / d582ca747972 / 4

- [rules.action.dynamic.elastic_ips.refs](resources--nat_policy--reference--group-001.md#canonical-c9ceb5c9e8224434093da60aa1141f1178a401280f796835cd89f1614f66574b)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-c9ceb5c9e8224434093da60aa1141f1178a401280f796835cd89f1614f66574b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9f6fcc337d0474c4c669e68b472a96862246b61f8385ffb59a71a95a1890b50"></a>

## rules.action.dynamic.elastic_ips.refs — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606)
- [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-78f86b1b602fc376c4979b1336a1c8fb492d4bd07e180d3c3ae320eb8e7fac83)
- rules.action.dynamic.elastic_ips.refs

<a id="canonical-ec09fd6d84490617578ed31440a349e376dbbc95bacd5973aae6fcc6251414a8"></a>

Type: `"object"`. list nested block, Optional.

Reference to one or more cloud elastic IP objects.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-249f256b4494acf06ea101e32e0ec94b80abf7a407583caf99a9238be07dad75"></a>

## Direct properties — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 3

<a id="canonical-579ac0fc0dc3d7c6e03c81babc6c58d259a204889e4e6b4ac91fbb526573bdce"></a>

<a id="canonical-41bae1ab3e0a41e42429b26a998bfe2062afb9c0e3ab74d450ef7f64bd2071a0"></a>

## kind property — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 4

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

<a id="canonical-672951287156a507a3a25cea199bd8d0c85a4f3caa4c2d192c5d020bb90ae2a4"></a>

<a id="canonical-5bdf7d6ddfe7bdd83c260b4a6a03a5f40ba3754e88bd50caa33f7ff805abebff"></a>

## name property — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 5

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

<a id="canonical-6bb4fd04113d09034dd7f39bef30e9f3cfb38f2d3abe7fee775dd883d088413c"></a>

<a id="canonical-a3d7d75304a389d5750c8c1143b46219762a1b5b49da56c759f749acc32aae74"></a>

## namespace property — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 6

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

<a id="canonical-49de7c06c328c1b8ed5482d707fba552299ccdec05c3e4f73853c799ff3b4768"></a>

<a id="canonical-d65ee32dc907ce769d297ef1ad27ee4ac7a96eabd7c25428749d39cc81b017d2"></a>

## tenant property — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 7

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

<a id="canonical-3a12d5d32bc5eec2c8af5db2e4ecc848338bf134b23a521b8c3758c9a39eb2d8"></a>

<a id="canonical-b808f018ce631b374257e7b65dffe695e7ffa6a74f565d6ad61f69691f704913"></a>

## uid property — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 8

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

<a id="canonical-6ff23e9465a795b021293d6bf8fa0d70e6da9ef36fa9293dbe019c33f141879f"></a>

## Next pages — rules.action.dynamic.elastic_ips.refs / 769aef62e70b / 9

- [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-78f86b1b602fc376c4979b1336a1c8fb492d4bd07e180d3c3ae320eb8e7fac83)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-d1f08b997061ce8f5880e46c0d16a54ed0e4edc6f059e4603f2549d6b06c9e12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b8beb4072e6c11410c0db02674ce03f2d30fbf203a3ccc3e11d52ea0db49090"></a>

## rules.action.dynamic.pools — rules.action.dynamic.pools / 3da4f972fe33 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-5aa4931b9dbe9d9e2c0472d01ce78e85a83bccd70b08b1d3dc52294e4a037f66)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606)
- rules.action.dynamic.pools

<a id="canonical-bf68e6f59e6e285b022d991d8529ff64270488d32f798448406a4e83671964fe"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-e57930a90c0c822b2973fdab20a841fef8030e50e05902f6de58d1541cd77e82"></a>

## Direct properties — rules.action.dynamic.pools / 3da4f972fe33 / 3

<a id="canonical-dd30acacb901b59b025d14fa937550585c863bf2a9b9b3cb0d7792a5a7533860"></a>

<a id="canonical-a3ebcc61e59389f365e571845a19a0763d0baad0c675f8ba3ff5d4192185ff1d"></a>

## prefixes property — rules.action.dynamic.pools / 3da4f972fe33 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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

<a id="canonical-172d30587a2d0f8d15e1137a80b3d77e74b221d381b0e7594e0d27a975b207f9"></a>

## Next pages — rules.action.dynamic.pools / 3da4f972fe33 / 5

- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-299fe04765b48e22946fd905f5bf0ee63bfa16ffcb4f9c30470e9057504e9606)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-4342fc9a3d91bb76d1a863120502c05f2c0913d3fcbf502d9f2184ff8c2714e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32a74fe67698084efa4286c6c8009eacc2052e7cfc273f0f4a07231249606231"></a>

## rules.cloud_connect — rules.cloud_connect / 97060401874f / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.cloud_connect

<a id="canonical-55d8d7c0079a69aff492b83d3724f4df590c7bc0f0f03a26b7a905d9ccd34f32"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cloud connect.

Upstream description:

Reference to Cloud connect Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
cloud_connect {
  # Configure direct properties listed below.
}
```

<a id="canonical-8b8842b506e3ce59a0e327abac71d2a30b7453a32f77e773bfed1f3083e6ff21"></a>

## Direct properties — rules.cloud_connect / 97060401874f / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-31f8613a2ee41f0b210679f08957842133cdfb8aa1a724ae9bf8482a9cc359cb): complete subsection reference.

<a id="canonical-231ebaa0cb2232ce5f89250d77008fc08033eff76e46c2daf52ab067bb0e7b76"></a>

## Next pages — rules.cloud_connect / 97060401874f / 4

- [rules.cloud_connect.refs](resources--nat_policy--reference--group-001.md#canonical-31f8613a2ee41f0b210679f08957842133cdfb8aa1a724ae9bf8482a9cc359cb)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-31f8613a2ee41f0b210679f08957842133cdfb8aa1a724ae9bf8482a9cc359cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2d21e57a0eeea0dd7e7838528d844faf5d7ec05788cba7f0d1952e735d482b3"></a>

## rules.cloud_connect.refs — rules.cloud_connect.refs / 450fe7776430 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-4342fc9a3d91bb76d1a863120502c05f2c0913d3fcbf502d9f2184ff8c2714e5)
- rules.cloud_connect.refs

<a id="canonical-7505bfc803ec07588b5bc5890b2fbf9068696e916e754ca6529229ef1200a504"></a>

Type: `"object"`. list nested block, Optional.

Cloud Connect. Reference to Cloud Connect Object.

Upstream description:

Reference to Cloud Connect Object.

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

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-a73a72f5693c2e406a31aeb4fbc5011d2a6445ab825db5e06b45c1f7a1e8d0f6"></a>

## Direct properties — rules.cloud_connect.refs / 450fe7776430 / 3

<a id="canonical-b44827260e0760976aba78eab6ca001769b7f6f13c0f3f5edb0154c168cc5cb6"></a>

<a id="canonical-609153c5db2a924ddaa4ed34a515cad85f03e45e74c92c3fcec40736e45f2d88"></a>

## kind property — rules.cloud_connect.refs / 450fe7776430 / 4

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

<a id="canonical-0206dac9358378e5cff38bb3407a30d3e52d86fca0cbcfb2c788b4a506080f65"></a>

<a id="canonical-7aeffae5a192d141a1d474abb4d9f2fc1d860c7c092f3e25ea00b538eee91ddf"></a>

## name property — rules.cloud_connect.refs / 450fe7776430 / 5

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

<a id="canonical-097af11a2c534b37280b3dee5e376d635c58489ffff26ff838000b44ba847e61"></a>

<a id="canonical-6989d0aaad92a2399ea8c0ffb1609488c04f5fa0ec60235035cb86bf19186221"></a>

## namespace property — rules.cloud_connect.refs / 450fe7776430 / 6

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

<a id="canonical-915f4bfb6f228b47347cb8aaa60eb393013791bba853268395c60b74fdad7730"></a>

<a id="canonical-ddb159144492a7fa4800eff47f951eaa90804c5314f243179effa149609d02b1"></a>

## tenant property — rules.cloud_connect.refs / 450fe7776430 / 7

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

<a id="canonical-94e10212c1024f37b2096f204c6998231484ef6fb569fba1754cd57184c53c04"></a>

<a id="canonical-c32b7b9ef19704ce0546523771288a6f86dd9eef62b6422f800c5cf6f6285724"></a>

## uid property — rules.cloud_connect.refs / 450fe7776430 / 8

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

<a id="canonical-8e76614d9ec3ffca7570905ab2f8755a584370694d8f601c78219a683b879d76"></a>

## Next pages — rules.cloud_connect.refs / 450fe7776430 / 9

- [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-4342fc9a3d91bb76d1a863120502c05f2c0913d3fcbf502d9f2184ff8c2714e5)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31623547db3ccff817fb9ef0bd06b609a4022c49f05970614bf2a265ef90bdd8"></a>

## rules.criteria — rules.criteria / 43276f01b2bd / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.criteria

<a id="canonical-7dd48863468c787aa4452446dfa54400c7d1f848cc02182c2af7d04d08608d0b"></a>

Type: `"object"`. single nested block, Optional.

Match criteria of the packet to apply the NAT Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "icmp"),
  validators.ConflictingObjectAttributes("any",
    "tcp"),
  validators.ConflictingObjectAttributes("any",
    "udp"),
  validators.ConflictingObjectAttributes("icmp",
    "tcp"),
  validators.ConflictingObjectAttributes("icmp",
    "udp"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
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
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-protocol_choice": "[\"any\",\"icmp\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
criteria {
  # Configure direct properties listed below.
}
```

<a id="canonical-133e90f85ee9a7a01eaaffb230bf9c8da92437b979b3b373cd1c1fc414693aef"></a>

## Direct properties — rules.criteria / 43276f01b2bd / 3

- [any](resources--nat_policy--reference--group-001.md#canonical-fa715eada591669e7920ee25cc4273e944889daca2078b3c29c457dbf0f13c7e): complete subsection reference.

<a id="canonical-cbe6715367bf33a4afc26d66f85cd58b8ffe7d4683466421d58c9b880e561662"></a>

<a id="canonical-788f09330e362ca4dd05db8063c5fada28422e3162454f029182670346a86a25"></a>

## destination_cidr property — rules.criteria / 43276f01b2bd / 4

Type: `["list", "string"]`. Optional.

Destination IP. Destination IP of the packet to match.

Upstream description:

Destination IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [icmp](resources--nat_policy--reference--group-001.md#canonical-3574f0c4bb9cd0c0d53000435854fb14fd42fa0ae7f47b31d5b611d1ed6ba050): complete subsection reference.

- [site_local_inside_network](resources--nat_policy--reference--group-001.md#canonical-15f4fc5d998cbee70d75373b8fe9fbec15ffa8d78c28237055087ecd21ade379): complete subsection reference.

- [site_local_network](resources--nat_policy--reference--group-001.md#canonical-c109141953874a243723271643383e2cec1099ebf1ac6b5bb6161b4fd244b8d3): complete subsection reference.

<a id="canonical-ad8cb6dd7c3312510d621d12b9ea617c5ac29b75121eeed669020cc081aeab67"></a>

<a id="canonical-c8c659eda056972e564e73f09bd9b16d3a4d79999651e57d05adeaad1c93dae9"></a>

## source_cidr property — rules.criteria / 43276f01b2bd / 5

Type: `["list", "string"]`. Optional.

Source IP. Source IP of the packet to match.

Upstream description:

Source IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f): complete subsection reference.

- [udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27): complete subsection reference.

<a id="canonical-3bd636d43e42df4fc0529b916e7aa29d88dfb2ab0d257213724ba44f499c2349"></a>

## Next pages — rules.criteria / 43276f01b2bd / 6

- [rules.criteria.any](resources--nat_policy--reference--group-001.md#canonical-fa715eada591669e7920ee25cc4273e944889daca2078b3c29c457dbf0f13c7e)
- [rules.criteria.icmp](resources--nat_policy--reference--group-001.md#canonical-3574f0c4bb9cd0c0d53000435854fb14fd42fa0ae7f47b31d5b611d1ed6ba050)
- [rules.criteria.site_local_inside_network](resources--nat_policy--reference--group-001.md#canonical-15f4fc5d998cbee70d75373b8fe9fbec15ffa8d78c28237055087ecd21ade379)
- [rules.criteria.site_local_network](resources--nat_policy--reference--group-001.md#canonical-c109141953874a243723271643383e2cec1099ebf1ac6b5bb6161b4fd244b8d3)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-fa715eada591669e7920ee25cc4273e944889daca2078b3c29c457dbf0f13c7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1fdae57b4320024eb11c5daede0268fe2696ee6363892fc5f0ebd7d9e78f1b1"></a>

## rules.criteria.any — rules.criteria.any / 9fe5c8656bbe / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- rules.criteria.any

<a id="canonical-1f7ac2a14cb88f41e55978be025d55724a3e6c3f23366eb5f7e56b0d8d2d7db5"></a>

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
any = {}
```

<a id="canonical-1e8fcfc06ba747dcb19c23fc073483019bffcd046d6a226d41968bd6c1054386"></a>

## Direct properties — rules.criteria.any / 9fe5c8656bbe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5947ac8fc1485397d4b8db662ee419fb0476ed6bb931068f5578dcf5cea6423"></a>

## Next pages — rules.criteria.any / 9fe5c8656bbe / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-3574f0c4bb9cd0c0d53000435854fb14fd42fa0ae7f47b31d5b611d1ed6ba050"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21520f6f5fa1b14d3b966140b21ea672dab17487e83c94fe3398fe55692b9d26"></a>

## rules.criteria.icmp — rules.criteria.icmp / fa201d5d558e / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- rules.criteria.icmp

<a id="canonical-e9a833ad5eb92acb80588cf64f271dd607f449a5df61346ef2a051b3b2dac47a"></a>

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
icmp = {}
```

<a id="canonical-35b61c9a08f71056fdbfd3f1116c77a67c8eb0cccef0269c6f677d897324138b"></a>

## Direct properties — rules.criteria.icmp / fa201d5d558e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-218d8d71ebb8a30e270d17b666a0ee7033af5b8056d6115399516a0372e322ac"></a>

## Next pages — rules.criteria.icmp / fa201d5d558e / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-15f4fc5d998cbee70d75373b8fe9fbec15ffa8d78c28237055087ecd21ade379"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fbae4e66dcf1377b7d354d15c6c78c4602dbcba5d5d8dc134c08bdc8c4b10e8"></a>

## rules.criteria.site_local_inside_network — rules.criteria.site_local_inside_network / 3be3ae4040df / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- rules.criteria.site_local_inside_network

<a id="canonical-12599a852d474c0ee740a5a6d1f6f55aa14a65a21f62886d9d2422e98b947322"></a>

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
site_local_inside_network = {}
```

<a id="canonical-e58d61bbb1b1b172ad64304759af3dc40a1c11a5d935beba0c0d4d17a3564105"></a>

## Direct properties — rules.criteria.site_local_inside_network / 3be3ae4040df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e60c69e22b21f34fd1c364eb0ea96a514b9d0f2778900c2dd4de84fe265df9e7"></a>

## Next pages — rules.criteria.site_local_inside_network / 3be3ae4040df / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-c109141953874a243723271643383e2cec1099ebf1ac6b5bb6161b4fd244b8d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb293aec24cb99035856db74030d22ffc0f7965302b69301b4cdc4dfa692eec6"></a>

## rules.criteria.site_local_network — rules.criteria.site_local_network / 0f34cb9c814a / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- rules.criteria.site_local_network

<a id="canonical-e31e2c6f06af7fc193c9dabef9b8c0e909bd2f587d46bdbc532b365fc8ff583f"></a>

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
site_local_network = {}
```

<a id="canonical-229f646ae7883e85254919f2aa00dd8c213538c67f5f6d44dd590ed16830f436"></a>

## Direct properties — rules.criteria.site_local_network / 0f34cb9c814a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4aee3ab99680995f1f1138b6ad97d548de9be4c0cdf025b5c79134ccb738422"></a>

## Next pages — rules.criteria.site_local_network / 0f34cb9c814a / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c09e7fbdb782e567e73796a890c24549772926791ed464373d8d5c3a2a24f0dd"></a>

## rules.criteria.tcp — rules.criteria.tcp / a1625c00d973 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- rules.criteria.tcp

<a id="canonical-24a5bf055d8b6110dccae6efef824388acc298bd61e290e67470adadced827e8"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

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
tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-da0acfd9cb04520d5b66680b6dced81046a950a051ce3e377e2be710a7123b25"></a>

## Direct properties — rules.criteria.tcp / a1625c00d973 / 3

- [destination_port](resources--nat_policy--reference--group-001.md#canonical-5a25dd4b5c7dce57330dbee382ec664969befff58f1585eb73d7caac40c0c816): complete subsection reference.

- [source_port](resources--nat_policy--reference--group-001.md#canonical-06ff2ee02212d8ad11b2f8751f60d3539227c6e75945773306fb486dfecca78c): complete subsection reference.

<a id="canonical-75a2137c169423df26f1c8b7e2411ca5fb6798a9b8c7ccd2a9bb62bb0d6bc5a2"></a>

## Next pages — rules.criteria.tcp / a1625c00d973 / 4

- [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-5a25dd4b5c7dce57330dbee382ec664969befff58f1585eb73d7caac40c0c816)
- [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-06ff2ee02212d8ad11b2f8751f60d3539227c6e75945773306fb486dfecca78c)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-5a25dd4b5c7dce57330dbee382ec664969befff58f1585eb73d7caac40c0c816"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b547c78d93c938dd2f57adffd6c3cddb05b4dbc80ef919d2bdf58d502b65fdc"></a>

## rules.criteria.tcp.destination_port — rules.criteria.tcp.destination_port / a66dd0b59ac8 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f)
- rules.criteria.tcp.destination_port

<a id="canonical-4ffbd20d22809f74863e0c13732eb5e07db8427daa6fc25a755fc18dbbd549c3"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
destination_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-79eb8658cec74b1b7110c1e49df5a0b8ed43b2f2a0b619c25eff20828cf7f6ff"></a>

## Direct properties — rules.criteria.tcp.destination_port / a66dd0b59ac8 / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-0b4e06af8c1629d43dd13fc619c5a7df775260cd8318955235fe664fa02a9232): complete subsection reference.

<a id="canonical-d29fbcd3fb4b28716a6b07650385a5c16a304979c2813b4b0a1711b591de21bd"></a>

<a id="canonical-319323a9d2c25cff9ea48fcf57dd65f60efbaf7310a49808579125500e444bb4"></a>

## port property — rules.criteria.tcp.destination_port / a66dd0b59ac8 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-2957e7c9a14fc8984059c551f399cff22a2ef1f44a73494168bd2e0590a4edf7"></a>

<a id="canonical-4ba81f9e5cfe226a0042bb97c14d22ba492c4fe85e5b09fabad4f75477820a7f"></a>

## port_ranges property — rules.criteria.tcp.destination_port / a66dd0b59ac8 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-fdf714e97c0494acb9faa5633514d6c60c6a6cee8b9f02645a6e77c1a1e3b13f"></a>

## Next pages — rules.criteria.tcp.destination_port / a66dd0b59ac8 / 6

- [rules.criteria.tcp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-0b4e06af8c1629d43dd13fc619c5a7df775260cd8318955235fe664fa02a9232)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-0b4e06af8c1629d43dd13fc619c5a7df775260cd8318955235fe664fa02a9232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa0f9c67a5886e82024ee4aeb09f387732dc5e098ed5d8e6608bf95b0a0b28ac"></a>

## rules.criteria.tcp.destination_port.no_port_match — rules.criteria.tcp.destination_port.no_port_match / 20623425927b / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f)
- [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-5a25dd4b5c7dce57330dbee382ec664969befff58f1585eb73d7caac40c0c816)
- rules.criteria.tcp.destination_port.no_port_match

<a id="canonical-d51ba79caa86b577c0b9e39c1bcbafb4020ad26619ba0b4bcaaedefee5ee8435"></a>

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
no_port_match = {}
```

<a id="canonical-1103f9687d03aefb4d9af0fde08069950c3a526400b401c87f7d6d489d8824e9"></a>

## Direct properties — rules.criteria.tcp.destination_port.no_port_match / 20623425927b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48b4e77ef0fcc4d64686716db6ebc19199f2acbfc2fa9c7f8e21010e730d03c8"></a>

## Next pages — rules.criteria.tcp.destination_port.no_port_match / 20623425927b / 4

- [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-5a25dd4b5c7dce57330dbee382ec664969befff58f1585eb73d7caac40c0c816)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-06ff2ee02212d8ad11b2f8751f60d3539227c6e75945773306fb486dfecca78c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dab1fe2bc9618d0fc80c713f5708bd2a8c62d44c38be09f2b83bb01c347b49dc"></a>

## rules.criteria.tcp.source_port — rules.criteria.tcp.source_port / f59994ee883f / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f)
- rules.criteria.tcp.source_port

<a id="canonical-d2e0950161e1e63058d13a7f6df36207da2aae283c0f88d343bff16bc1e87785"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c5811a6a2ddce2a1f525d99986e5cb80222e4654c6699f638b52251d7783bce"></a>

## Direct properties — rules.criteria.tcp.source_port / f59994ee883f / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-1d03192a12e371be2e6e5909d5e81e78bb5d57f4a157ca8fbacd45082537e175): complete subsection reference.

<a id="canonical-1fc713c34b290a7207e864eb1756535ba761f14e583e3e88afebae0fc4478bed"></a>

<a id="canonical-49ab39fea7482207c7e6d8d5064f42baef31afc694cd538939403dcad6de3c27"></a>

## port property — rules.criteria.tcp.source_port / f59994ee883f / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-0e4557d77924def2847802d78845720842aab283839a0df8e3e976abef0bd16a"></a>

<a id="canonical-6bb07eb09a1996e20001a2a851685af963799798189429e146438d9bf99a6887"></a>

## port_ranges property — rules.criteria.tcp.source_port / f59994ee883f / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-ecbd80d8c98052d93200e40fe563513c6019925cc9e6311e7535f894586464bd"></a>

## Next pages — rules.criteria.tcp.source_port / f59994ee883f / 6

- [rules.criteria.tcp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-1d03192a12e371be2e6e5909d5e81e78bb5d57f4a157ca8fbacd45082537e175)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-1d03192a12e371be2e6e5909d5e81e78bb5d57f4a157ca8fbacd45082537e175"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-608b175519ad5abe84799aaac0a01e272f4abf79de1936cf30e8fb2261b14cc1"></a>

## rules.criteria.tcp.source_port.no_port_match — rules.criteria.tcp.source_port.no_port_match / eb9ef517ed5c / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-b1bdfbed6fd6c557705ae8b824b538aaccbac32100bb791845f110c252ff5d0f)
- [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-06ff2ee02212d8ad11b2f8751f60d3539227c6e75945773306fb486dfecca78c)
- rules.criteria.tcp.source_port.no_port_match

<a id="canonical-fab26929fc18b6adb0cee84a8b2405aede61e0fd9075503f4acdc511f3bee12c"></a>

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
no_port_match = {}
```

<a id="canonical-e8a38f9124583bc972567828eda4c8fae41ca732cdf3ce3fa1eddd410930fa72"></a>

## Direct properties — rules.criteria.tcp.source_port.no_port_match / eb9ef517ed5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1578fcef2e5f798d6a392555aa5f805d9bcd826af2899def908a1b7f6a2f446d"></a>

## Next pages — rules.criteria.tcp.source_port.no_port_match / eb9ef517ed5c / 4

- [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-06ff2ee02212d8ad11b2f8751f60d3539227c6e75945773306fb486dfecca78c)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-471ac449814817eed381370f322bfcf010fc0c4f6b70b03ab1d5b1497527bb1e"></a>

## rules.criteria.udp — rules.criteria.udp / 4d28848e60e0 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- rules.criteria.udp

<a id="canonical-2b4960baf77e3f53a6ec09723fdacd831427aa76dd2af28e3d8c5b6f24781dee"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

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
udp {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8c79612e000c168b22da44606199fe1fec125eef17142bca0e3ecc9fb0db1d4"></a>

## Direct properties — rules.criteria.udp / 4d28848e60e0 / 3

- [destination_port](resources--nat_policy--reference--group-001.md#canonical-dc70f589e549301d26f51441e60c48b80c042414a015bab9226e99621e765e89): complete subsection reference.

- [source_port](resources--nat_policy--reference--group-001.md#canonical-1a1a3b5c22fd1cd608c3bd3b4b4fd90b1a63dbd7b6ad5b68a2b159b9e1e6266d): complete subsection reference.

<a id="canonical-7e3b0a6f0a7f474751eee3a77d8389d593eeef95307da91997a2fd9c4e599edb"></a>

## Next pages — rules.criteria.udp / 4d28848e60e0 / 4

- [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-dc70f589e549301d26f51441e60c48b80c042414a015bab9226e99621e765e89)
- [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-1a1a3b5c22fd1cd608c3bd3b4b4fd90b1a63dbd7b6ad5b68a2b159b9e1e6266d)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-dc70f589e549301d26f51441e60c48b80c042414a015bab9226e99621e765e89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d78de91105199962b232b5a34202305fd9ac6973864e972bdb8fde03ef7d2187"></a>

## rules.criteria.udp.destination_port — rules.criteria.udp.destination_port / 6d08857fb310 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27)
- rules.criteria.udp.destination_port

<a id="canonical-a795621116f8a2e03ace2d01af4d78512ccee5a97223c832ce36e8cbabe673b9"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
destination_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-dca2eae7217d7d21ae8a5529ac35b2a686c168943b6c2aee4085fce80c6bec15"></a>

## Direct properties — rules.criteria.udp.destination_port / 6d08857fb310 / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-4b8b18318627efd495a99031665c36da7782558ae62b93ddf87275edbf248c67): complete subsection reference.

<a id="canonical-3e05f276fd4b764af633c5f301cfd8ec8178a4702d4116e07a9e138a5db00595"></a>

<a id="canonical-c751e6ff1f242e110d1c1fc02c4616e38c7ab1cb894af78270079406bb199be3"></a>

## port property — rules.criteria.udp.destination_port / 6d08857fb310 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-1586c28a18145a9dae55532b2f0c085a62c1dd0543592423b5436a94e9385bc5"></a>

<a id="canonical-7066a55d6b31b3486845297b7fa1518bf191837c8f92ce31802c9bda31803565"></a>

## port_ranges property — rules.criteria.udp.destination_port / 6d08857fb310 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-5933202611eab1619013abef19ae205490b4f7f49995285f893c4fcefdf77c6f"></a>

## Next pages — rules.criteria.udp.destination_port / 6d08857fb310 / 6

- [rules.criteria.udp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-4b8b18318627efd495a99031665c36da7782558ae62b93ddf87275edbf248c67)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-4b8b18318627efd495a99031665c36da7782558ae62b93ddf87275edbf248c67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0d11cabbec92cb6899d2a94b0ecb449b3e7ef32456dba60f80cff2923dc5af9"></a>

## rules.criteria.udp.destination_port.no_port_match — rules.criteria.udp.destination_port.no_port_match / 211d74417663 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27)
- [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-dc70f589e549301d26f51441e60c48b80c042414a015bab9226e99621e765e89)
- rules.criteria.udp.destination_port.no_port_match

<a id="canonical-4ff9d52e91de016e0bc6ba1db04648d73adae9472d0d2ec86b93ac99101f1882"></a>

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
no_port_match = {}
```

<a id="canonical-07f2e5e56a39d2c648768d57fb578cfe0b23dd568cd929c8ebd60b5d2778ddad"></a>

## Direct properties — rules.criteria.udp.destination_port.no_port_match / 211d74417663 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-715831218ffac41c8463e3860f42b61d7419a6d0068221803682cbd9d7dda266"></a>

## Next pages — rules.criteria.udp.destination_port.no_port_match / 211d74417663 / 4

- [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-dc70f589e549301d26f51441e60c48b80c042414a015bab9226e99621e765e89)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-1a1a3b5c22fd1cd608c3bd3b4b4fd90b1a63dbd7b6ad5b68a2b159b9e1e6266d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b5f125ca8e791e9ea350bbf490afc7619457d0269bb15ce1ef1e10e958da4b2"></a>

## rules.criteria.udp.source_port — rules.criteria.udp.source_port / 05242e7898b0 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27)
- rules.criteria.udp.source_port

<a id="canonical-b4977be6d64cdf8f95bb49bd5a5f48ebadb9960c382f9824dff857dfceeefe81"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-b209de851ae01258472216c4192c33a087c8e98fad3ecc8ce85c4326646350dc"></a>

## Direct properties — rules.criteria.udp.source_port / 05242e7898b0 / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-c068f0d10feb7aa191a2750eec6f4a57f743eac6d07320d3748bc3c5751d763e): complete subsection reference.

<a id="canonical-ed864ca1826eb6e38747e55e906734e0daea224ab5ef624332617d5d9e32df93"></a>

<a id="canonical-52b0cac5418542fcf14ad0ed8efb9ee7d46e41d26388aa94cb32303eee1f015f"></a>

## port property — rules.criteria.udp.source_port / 05242e7898b0 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-6d8d96502edfec0b1f6d86a3d63e62e4bca3f0490dc9d8fb9ffcc2c30d203335"></a>

<a id="canonical-f0830904d6850a0b274895ca8212aae041722d583f58105c7ff2a4516fb60817"></a>

## port_ranges property — rules.criteria.udp.source_port / 05242e7898b0 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-93a277c383cb06cb0da5339d387935a4d3166dcc4d9a084bfde8a77967dd0eab"></a>

## Next pages — rules.criteria.udp.source_port / 05242e7898b0 / 6

- [rules.criteria.udp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-c068f0d10feb7aa191a2750eec6f4a57f743eac6d07320d3748bc3c5751d763e)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-c068f0d10feb7aa191a2750eec6f4a57f743eac6d07320d3748bc3c5751d763e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15dbfd35c09bae41127939adf264cbd8f4c07a6fbce4794b0db6b1cb4936899f"></a>

## rules.criteria.udp.source_port.no_port_match — rules.criteria.udp.source_port.no_port_match / 609d887ec41e / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-2589e4111188ea2a05156570e8206a1422195c18fd0a3d3d2ba5f2c5a797ee0a)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-a5f4fbbf90064ccd9d0c209631845ee51e6996c1dfc9ac5dd7b8fbb148912a27)
- [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-1a1a3b5c22fd1cd608c3bd3b4b4fd90b1a63dbd7b6ad5b68a2b159b9e1e6266d)
- rules.criteria.udp.source_port.no_port_match

<a id="canonical-e412ab6f608456fa69e649d824c11470c4c0d16217e74102b25781efb9845955"></a>

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
no_port_match = {}
```

<a id="canonical-20b6e6173bf851999aea9c04622f4efdd94f90bf8e628541442d49e2f63d6a9d"></a>

## Direct properties — rules.criteria.udp.source_port.no_port_match / 609d887ec41e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f11e0b7af823d23f0a163254ae36050141344dc53d5a2a5537e1fdc186d7182"></a>

## Next pages — rules.criteria.udp.source_port.no_port_match / 609d887ec41e / 4

- [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-1a1a3b5c22fd1cd608c3bd3b4b4fd90b1a63dbd7b6ad5b68a2b159b9e1e6266d)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-bf54ee0daa4be07c6891585b45687b03ed6fc1e9d3784be804a7c14e20716e50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b74501fa745ab52c4918facdff375f503f88539245f41a51e5a68d320e88e70d"></a>

## rules.disable_spec — rules.disable_spec / f20cec165bc2 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.disable_spec

<a id="canonical-6c85d7ac365c0b6539968523103d68b87d0f47da4879d2bac1b947c8c69ca892"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-ad5b9f058b828d53bbd4f0198b139d0be0ef901e22b1c9b4c8fb84009a1f3fe9"></a>

## Direct properties — rules.disable_spec / f20cec165bc2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a7aca5132ec5e1d21f8258b76ae08debdb9962abed729ca95f789a8f3b5edc3"></a>

## Next pages — rules.disable_spec / f20cec165bc2 / 4

- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-f180a8ae475eb7bf551c7e13e1734b9a6f2fce42081c03831911b5d5a32f2376"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b1d0032e096e6f39fc1089e7144ec5332818eaa1dc06476b5d891f0e25871aa"></a>

## rules.enable — rules.enable / ba265abf68ce / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.enable

<a id="canonical-95ebc55c6aaa91690375673ec7d62fac8a9fac239246b0dc404e2ac9c6c58cb1"></a>

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
enable = {}
```

<a id="canonical-9424c68a3aa8ad295365120611aa6c9516cc9173c58a277c985a4e14d2110b01"></a>

## Direct properties — rules.enable / ba265abf68ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a6242422b05d37750bf6b95cab2bbc45ff03b49d3bcd0f98e99576275252387"></a>

## Next pages — rules.enable / ba265abf68ce / 4

- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-0368e87c2a3dd0f0db2452d2d3d60fab5e73b8386b388ae31ce6ed3019df8c28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b8aed164956632cd8dca3301c5b79b2f5cc90b6719748bffeeac830bd869ada"></a>

## rules.node_interface — rules.node_interface / 419201ed94b5 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.node_interface

<a id="canonical-38fc84da28372aa048be9c3f9cb3f90d481ba83b41143f80a5bd92519276409b"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-211a97a3142a928cb11d8f508321a3e7d23a48610ce7172974dcad673daab955"></a>

## Direct properties — rules.node_interface / 419201ed94b5 / 3

- [list](resources--nat_policy--reference--group-001.md#canonical-9a2aec8dbf3913892d85f347b954cb5c91cdd51c830bd2582ddb2f7fd40aab0e): complete subsection reference.

<a id="canonical-e9f26eefdae32822440ad4621ea4ff124bc4b2eb10f1087a1d838ab562d3fca4"></a>

## Next pages — rules.node_interface / 419201ed94b5 / 4

- [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-9a2aec8dbf3913892d85f347b954cb5c91cdd51c830bd2582ddb2f7fd40aab0e)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-9a2aec8dbf3913892d85f347b954cb5c91cdd51c830bd2582ddb2f7fd40aab0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffe9fe0da2ec513fed3ca8b1b753234afb10b5ad5b6e077f23a2037e546287d7"></a>

## rules.node_interface.list — rules.node_interface.list / 94dee5b9703e / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0368e87c2a3dd0f0db2452d2d3d60fab5e73b8386b388ae31ce6ed3019df8c28)
- rules.node_interface.list

<a id="canonical-000626cb7493cab71930acb06a06c5c3b3f460c3fb9f0ab2764f819f615676c4"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-382f68f60c2685f67b6ae75e7aacd5ff8be9bbabd22a09b4117ea9fb3eb2edae"></a>

## Direct properties — rules.node_interface.list / 94dee5b9703e / 3

- [interface](resources--nat_policy--reference--group-001.md#canonical-2685132bcc9ea0a24c338eafe4d25aa3de764e2dc0aa03ab1779fc538516d5be): complete subsection reference.

<a id="canonical-badaa17ff37aed10801adb03ba182938a074ce468388dd7b157be0def0daf509"></a>

<a id="canonical-04b16c8c90231491b93e83918daffd4576e18d258832498fbed50cb668d9f97a"></a>

## node property — rules.node_interface.list / 94dee5b9703e / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-5cdfd77f73d0287a8f42f21d05dadef057aa45ff201365e19073b753672fe62b"></a>

## Next pages — rules.node_interface.list / 94dee5b9703e / 5

- [rules.node_interface.list.interface](resources--nat_policy--reference--group-001.md#canonical-2685132bcc9ea0a24c338eafe4d25aa3de764e2dc0aa03ab1779fc538516d5be)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0368e87c2a3dd0f0db2452d2d3d60fab5e73b8386b388ae31ce6ed3019df8c28)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-2685132bcc9ea0a24c338eafe4d25aa3de764e2dc0aa03ab1779fc538516d5be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-371420efbda86c3a97fe0594a9a9e8cc8fc1c694f6140a3e91df78e54831d2c8"></a>

## rules.node_interface.list.interface — rules.node_interface.list.interface / 8f9fb4233362 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0368e87c2a3dd0f0db2452d2d3d60fab5e73b8386b388ae31ce6ed3019df8c28)
- [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-9a2aec8dbf3913892d85f347b954cb5c91cdd51c830bd2582ddb2f7fd40aab0e)
- rules.node_interface.list.interface

<a id="canonical-b69d82cb89ef323b338bc6d4cb0674d4dc8b94a69f846685f72ccd71fe0f89b6"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-59d366372888168c9fc9c6e300b072b26b4d9cbf774d51d70599b13c993615cf"></a>

## Direct properties — rules.node_interface.list.interface / 8f9fb4233362 / 3

<a id="canonical-aa44cd1026f67d8c8fe362af873ec75642044ea03bd4c2148e3c65803713cc56"></a>

<a id="canonical-54748604ac0d4d50cc9f42eb08d9ac37fa313658ff33a7fb04e18fb833f42f43"></a>

## kind property — rules.node_interface.list.interface / 8f9fb4233362 / 4

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

<a id="canonical-3f1f998291f1b08e5fce5a2c9ebc43cb5f2c5375b2e7c1adf6c9421c8c400be9"></a>

<a id="canonical-902748b44bc5af5cb27208d13c70c14307f6185932944eb799fe0a6b82c2410d"></a>

## name property — rules.node_interface.list.interface / 8f9fb4233362 / 5

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

<a id="canonical-d87bdca2131abf13b75075645d9f2968528b32f4acb02f62c79c02c5e038e48f"></a>

<a id="canonical-f59d990caacb3516d65cfcefcfd0bf06c8b8d859d915d5f52f5e2df84070ccde"></a>

## namespace property — rules.node_interface.list.interface / 8f9fb4233362 / 6

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

<a id="canonical-bd1faef0aa8db403283d5a298aac88ed70dcf791b0928189ec4aac00accbaa3d"></a>

<a id="canonical-c5a9cb3206db904570f0958ba9c0efdea89701363c9fef1cadee6567c23aca17"></a>

## tenant property — rules.node_interface.list.interface / 8f9fb4233362 / 7

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

<a id="canonical-36a7503d202c8b3753952919f378f3bbb85d066bb950d2d632a5b65b926360aa"></a>

<a id="canonical-4999f18120d3bbeebc72b5cd2b6db7a0b38471bd164aa7f0a96f57d88919e314"></a>

## uid property — rules.node_interface.list.interface / 8f9fb4233362 / 8

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

<a id="canonical-303f1cbd84068b432f0cec423c13962325d7232f661172579aaf56a019d0b9ac"></a>

## Next pages — rules.node_interface.list.interface / 8f9fb4233362 / 9

- [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-9a2aec8dbf3913892d85f347b954cb5c91cdd51c830bd2582ddb2f7fd40aab0e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-80f720e5966bcdaa2dae3cdbbdfb1e5b4b4dfd78584937c7cd88946eb93daeac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83069c560a4aa1d2332180974c50958983903a9681a40d5fced8fde958fc10fb"></a>

## rules.segment — rules.segment / 3a7207afe2ac / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.segment

<a id="canonical-420fbae347cff4e30d1043f0a651189970e1a867edbe6f5a6645116228ebda85"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-319847dc07b3b8e8346f587dc2a26cf26f8480f85b2bcd2ae99661120de7c0eb"></a>

## Direct properties — rules.segment / 3a7207afe2ac / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-534f8b1cba0945b5ff1af10098719b286493a0253ae651a3ebbcd90b2426004e): complete subsection reference.

<a id="canonical-2366034b55e378d658429788c91083c5d914f20bb4d22800ece06699c49aa103"></a>

## Next pages — rules.segment / 3a7207afe2ac / 4

- [rules.segment.refs](resources--nat_policy--reference--group-001.md#canonical-534f8b1cba0945b5ff1af10098719b286493a0253ae651a3ebbcd90b2426004e)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-534f8b1cba0945b5ff1af10098719b286493a0253ae651a3ebbcd90b2426004e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2b260fdaa0965c390e31578c6d8ff83d1726ae0b0cea901931cf12996e43b87"></a>

## rules.segment.refs — rules.segment.refs / 914baf7da0db / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.segment](resources--nat_policy--reference--group-001.md#canonical-80f720e5966bcdaa2dae3cdbbdfb1e5b4b4dfd78584937c7cd88946eb93daeac)
- rules.segment.refs

<a id="canonical-3e7df79a11835050cacfa90d8bae3c956bb54b6f5a4e8293569d4ef3c2bbc627"></a>

Type: `"object"`. list nested block, Optional.

Segment. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

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

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-d525a2504dee92c4d7638762a5c5e77f93f2873861c55a2ff4f3bc02d01015db"></a>

## Direct properties — rules.segment.refs / 914baf7da0db / 3

<a id="canonical-73fd8c736cb4c7597b81b0af58ef0be5f98ffbbc46a6ae04904e238b8312b899"></a>

<a id="canonical-1f9d33d1b1db46d3a66eea26bec3e1b858d1304decf4831360ce0f591ea61fb0"></a>

## kind property — rules.segment.refs / 914baf7da0db / 4

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

<a id="canonical-62c0c92aead1783e4d31b92b101153ef1c0f4a1f7abf74076506a85c64a6ac71"></a>

<a id="canonical-d55b419467fc15a59699aec7ca7a3a5b6ab8f60642347c3d4d1fb8ea262c4421"></a>

## name property — rules.segment.refs / 914baf7da0db / 5

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

<a id="canonical-f68e11041fc927edc6d3616fc155b31615cd1cbfabd134739c98d5532de0b97d"></a>

<a id="canonical-68bc3585c9239e0cc4caa299865d465f94cc57a890cb1ab4004e9b89743b9a44"></a>

## namespace property — rules.segment.refs / 914baf7da0db / 6

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

<a id="canonical-c57bc6329229b91ecc887f4222618c183ddd088d97ac80fdaf6932cac3cd786f"></a>

<a id="canonical-4858a02bb6ae53b33301d9cf5e8dfd1e5b91fa370cf61f789a3b75e4bc48f04d"></a>

## tenant property — rules.segment.refs / 914baf7da0db / 7

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

<a id="canonical-8207de645a714450d7182f79f25917df24cfb6a03ccf899fbb983fce44c4cfd7"></a>

<a id="canonical-ed3f70cfe0e47bb0c4a0098d7d963e35974ddc28a3d017f799087fb1ae55df14"></a>

## uid property — rules.segment.refs / 914baf7da0db / 8

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

<a id="canonical-d9dafb2fd66fb6a145a3252ea2cd4eef7c24307de50f05f94ffb6ff3560dcc22"></a>

## Next pages — rules.segment.refs / 914baf7da0db / 9

- [rules.segment](resources--nat_policy--reference--group-001.md#canonical-80f720e5966bcdaa2dae3cdbbdfb1e5b4b4dfd78584937c7cd88946eb93daeac)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-63554196d15ca33c0fd89f0529700793810d33309c458fbae48de4234f3a28a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38f1ffc196982efdb41038ed85980e8b43e2d1a9593c33e56203273bcb841a9c"></a>

## rules.virtual_network — rules.virtual_network / 579cb25980e2 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- rules.virtual_network

<a id="canonical-421b524e49ff918eac8ade720a2512a4ee4528d9384ac48256d1b77e4fbc9dad"></a>

Type: `"object"`. single nested block, Optional.

Carries the reference to virtual network.

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-83c736a75251703b958353739610b4f9dabd30720581cc1082a157f2fd99d618"></a>

## Direct properties — rules.virtual_network / 579cb25980e2 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-2cc1450ac8c7c22a6ccdedcab998ce37d236e023cb92bb352a33a367e77d6850): complete subsection reference.

<a id="canonical-9249cf4514965e3593c1a7914e27ec08efee0ea673dd5c93f4098f81df8fa344"></a>

## Next pages — rules.virtual_network / 579cb25980e2 / 4

- [rules.virtual_network.refs](resources--nat_policy--reference--group-001.md#canonical-2cc1450ac8c7c22a6ccdedcab998ce37d236e023cb92bb352a33a367e77d6850)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-2cc1450ac8c7c22a6ccdedcab998ce37d236e023cb92bb352a33a367e77d6850"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e236945421971d30c9432904edbd2bd32bc7b9ba24b16d961cd11532000bfa84"></a>

## rules.virtual_network.refs — rules.virtual_network.refs / 692ffabbe961 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [rules](resources--nat_policy--reference--group-001.md#canonical-86039b7815d03053539f52e1ad5d657e7b0dc599880562435fc68b67d5e8231e)
- [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-63554196d15ca33c0fd89f0529700793810d33309c458fbae48de4234f3a28a1)
- rules.virtual_network.refs

<a id="canonical-aa1793e4587086d8c25912cb1cdccaad76893637300906117e79ef188e29617c"></a>

Type: `"object"`. list nested block, Optional.

Virtual Network Reference. Reference to virtual network.

Upstream description:

Reference to virtual network.

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
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-2b6d0076593f55984dc2c635f35e1dd125bdd6bae23dee92481490d7839ad2ac"></a>

## Direct properties — rules.virtual_network.refs / 692ffabbe961 / 3

<a id="canonical-5779b6653355a3885de3e901734c392585fa4cf65d278846067b90426fbae516"></a>

<a id="canonical-3af865f9ba64e5ced27e10c5c4f5e735b1cd59dddc95406ebddc3215f89b7bb8"></a>

## kind property — rules.virtual_network.refs / 692ffabbe961 / 4

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

<a id="canonical-60d99a09e1d74c96b9b8073c6f4630a8f7bf72c09d56947b66fec0fbf648053a"></a>

<a id="canonical-257f7598d8497c6b1c687615cb08a3477798dc71791941ff9b64cf3e33ff1a66"></a>

## name property — rules.virtual_network.refs / 692ffabbe961 / 5

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

<a id="canonical-7dc9b077b3b4e5c75ad55d7b89dba23cafbe8723a932d1b0a9e1fe121ac055cf"></a>

<a id="canonical-d94a49d727929845b74abee704c290a7236fb1f44b211bb847c0db15bcf4d468"></a>

## namespace property — rules.virtual_network.refs / 692ffabbe961 / 6

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

<a id="canonical-b51fdf3aadc36275d5902f7836719c19d95270797d39a3cc30ffcaee7202236c"></a>

<a id="canonical-4091475617f627c6e904b4ca994824f12f0880f1d23a204ed186ea55768ce9ff"></a>

## tenant property — rules.virtual_network.refs / 692ffabbe961 / 7

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

<a id="canonical-0294a10e5de3ad08d91b28babe468f24229ec50c243486effb102e54ceedb136"></a>

<a id="canonical-4fa36b9cc276ee0f2395a113261e0ead72889a3b83bdf31b863400aa33e4b207"></a>

## uid property — rules.virtual_network.refs / 692ffabbe961 / 8

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

<a id="canonical-a7b8ca1b0544b7ad43ae8d0b25f71f56f9decbb4babcf75a2e3509aeb109744b"></a>

## Next pages — rules.virtual_network.refs / 692ffabbe961 / 9

- [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-63554196d15ca33c0fd89f0529700793810d33309c458fbae48de4234f3a28a1)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-c1aea168038d7fa6fe8f370374af11bfab68a850b95080a918405e78886cc2ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c09e4ea909e91cbad7ef183870757eb30d4cc955617ebe34da2c69e3fd48d69"></a>

## site — site / 8b37156bec95 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- site

<a id="canonical-8c4f0620116908cdcf7b9e80c2f041ae90a21e7caf6926958e2c0357aa72bd7d"></a>

Type: `"object"`. single nested block, Optional.

Site Reference Type. Reference to Site Object.

Upstream description:

Reference to Site Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-eab0aff81a4afe3a0f5c39a589d76a9dc8fe30fd75afce3e4d4102714ea56578"></a>

## Direct properties — site / 8b37156bec95 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-e72a34598a86fd4025e0e17f5d46ada34f8521cef6dd8c2a84a68b2fcb424e13): complete subsection reference.

<a id="canonical-7434e93a289ac83c8ea3175fa5034117ec568403b73661bac592810cf0b5375a"></a>

## Next pages — site / 8b37156bec95 / 4

- [site.refs](resources--nat_policy--reference--group-001.md#canonical-e72a34598a86fd4025e0e17f5d46ada34f8521cef6dd8c2a84a68b2fcb424e13)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-e72a34598a86fd4025e0e17f5d46ada34f8521cef6dd8c2a84a68b2fcb424e13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51b1aa59b345ac239e30d37519320461dcd13dc01163f99a541190b5816c48cf"></a>

## site.refs — site.refs / 51c59923584f / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [site](resources--nat_policy--reference--group-001.md#canonical-c1aea168038d7fa6fe8f370374af11bfab68a850b95080a918405e78886cc2ec)
- site.refs

<a id="canonical-a819e1b4a5787310556ee0ef25c488024b0bef665ad5ce704c227f2fde9ed533"></a>

Type: `"object"`. list nested block, Optional.

Site. Reference to Site Object.

Upstream description:

Reference to Site Object.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-89aae0eb5d4c2ecd5efcd10dd4dfd26e130a85854a8134ec657ae6701b654d85"></a>

## Direct properties — site.refs / 51c59923584f / 3

<a id="canonical-355875b6f404341a01d789d54b7f182370a7cbdb482f08ba731c1d180144550a"></a>

<a id="canonical-e93533a02c6de4a47c63fc72196c20b8377b8c17bc5e161e34b79396bd0314dd"></a>

## kind property — site.refs / 51c59923584f / 4

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

<a id="canonical-783da7404918f7eacddfc07424d658a8e8b6c06c470f0755369269f063511bca"></a>

<a id="canonical-29759b7b39b52d0388810d9b43942e0d2692d6e9410b2010e180622205376b65"></a>

## name property — site.refs / 51c59923584f / 5

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

<a id="canonical-57376f43fff4cffcba45ddaa0bc2d795f6640994076298bf1d84fe8113807366"></a>

<a id="canonical-71b084a35d4eede8244758cf429848c6715cd4ff1c5eff1e792764426ea5493d"></a>

## namespace property — site.refs / 51c59923584f / 6

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

<a id="canonical-a8d6a4eab9db74b977d2b194119897e76d0963d605ebc27b31e8e27cb18b1f1b"></a>

<a id="canonical-f9cf88d811d9dda9ba0ba864123b7768ddf6f3720c844e59c5e2ba78a5704d20"></a>

## tenant property — site.refs / 51c59923584f / 7

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

<a id="canonical-d26fd94fb6fff9f202ce38845a42206ecc8c74f5183ac54585b1038018742c5e"></a>

<a id="canonical-46f43b431b90c83b70a125d43731e46f9ddb7f2daa12c130d85da05c3a57272e"></a>

## uid property — site.refs / 51c59923584f / 8

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

<a id="canonical-0ddbf07537147ca6764471d3cb033eb3a213befd02db13c07286248ac862e2e8"></a>

## Next pages — site.refs / 51c59923584f / 9

- [site](resources--nat_policy--reference--group-001.md#canonical-c1aea168038d7fa6fe8f370374af11bfab68a850b95080a918405e78886cc2ec)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-3ed37e126f8c1e301894d45191d3eae175054a1ae68b7362934f480e7b2f1aee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fedb344a881a023608ef061676b0f333b85873362b7ba2b7854a0a4a03467e16"></a>

## timeouts — timeouts / 1b7756c61fd9 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- timeouts

<a id="canonical-5b02e38455b556b08cd7a472b909dca2be252b51e935d25b0aa939b4125e0812"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e722351586615a3c2b7fe7407077f6601a09d5e24d275c5b51a57588d95c541b"></a>

## Direct properties — timeouts / 1b7756c61fd9 / 3

<a id="canonical-45e888cc573ea621ee4b252a424d33f66db3823e823000826cb1afd351db819a"></a>

<a id="canonical-ab8cf97882105289e2c04005c4107c7ea68b0f86aad18fc4311c8e9c401f801c"></a>

## create property — timeouts / 1b7756c61fd9 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-74ef950a7c48cc18ab85f95a71e886b88fab27dfaff68cef2c69056efc11419e"></a>

<a id="canonical-b84624a19878f6793a1c1cbe00acad669b8c0e305b5555d7c0dcf8a3abd4ab61"></a>

## delete property — timeouts / 1b7756c61fd9 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-142eaa4f0281250bcc375b4363422a48e6c8d247b40fa99a75f35c821f5e758b"></a>

<a id="canonical-b6bf19ba5bf9ef499bb960636a457ef545efd829b98f13f88ef3bbd495ace410"></a>

## read property — timeouts / 1b7756c61fd9 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-26b5bda5697c8cfbd69adbaf3ae38307be5a32ac542d19288953b025e7d08ce6"></a>

<a id="canonical-8f8b6bd527b0a1d5bf7857f8b4cb955d5f03f03ddc49898897a465efe840e46f"></a>

## update property — timeouts / 1b7756c61fd9 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-306c5cc5289acb4d1acd0a9f8a07a1765c91cbe8a797ba64776a11c7dcfdb70a"></a>

## Next pages — timeouts / 1b7756c61fd9 / 8

- [Property reference](resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
