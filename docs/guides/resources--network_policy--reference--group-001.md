---
page_title: "xcsh_network_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy reference."
---

# xcsh_network_policy reference

<a id="canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17a5217ce880a4ec6c82615275a0f0236b38834e14530bffca052f332700554a"></a>

## Property reference — Property reference / 978e3bdb6a7e / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- Property reference

<a id="canonical-1def042b5180f53eed40a11a7a76fed1c0ba0ea28cfe9de451415ec0027c30b9"></a>

## Direct properties — Property reference / 978e3bdb6a7e / 3

<a id="canonical-48f9f10c8407d5ba047afd42bc3491ada662000da62e6129053cb52338145aec"></a>

<a id="canonical-1b4bba81d05a062accf6f2fe978ad6de7077c8ad19f4b74f26feffdd32c9f994"></a>

## annotations property — Property reference / 978e3bdb6a7e / 4

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

<a id="canonical-794071bbc8461baa856bb2167a6c45c5a658b986f621c5d20cbf8f1cc7c3f26b"></a>

<a id="canonical-f23b3b3ae656ff345800b033f73b5712e043984909e3f706ad1b6122b6f09d5d"></a>

## description property — Property reference / 978e3bdb6a7e / 5

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

<a id="canonical-8d37b8dc0a5038fe140e6f120abec8c503690c6d94b70dbc482e459d1d7f5950"></a>

<a id="canonical-8937bf7697ee82b8b88b9898c26ce8bd65cbf9e2aef651159c8e3943c6d36971"></a>

## disable property — Property reference / 978e3bdb6a7e / 6

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

- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7): complete subsection reference.

<a id="canonical-5e3687873c714cfe3c67abe1198810e53c4d1b27a0c5ebaa9d2fd62f2e7176c1"></a>

<a id="canonical-0a7ca252dcef0382bab0cc1073c10497f7ba7357258c68d1712024751f90080b"></a>

## id property — Property reference / 978e3bdb6a7e / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-577b976cd4be9c45b20ec46fe67800f3b21f19cf68d9acc271b9505ae8fee28b"></a>

<a id="canonical-01c881e4d763941e7beeea014aac2646574d3121c7a2da9378788d9e7cf7f749"></a>

## labels property — Property reference / 978e3bdb6a7e / 8

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

<a id="canonical-49db2e0ce287acbc6dd141c7cae3671c31e41c9147902b438b4a393c8013901e"></a>

<a id="canonical-e044cbcf5b8e24dc69251a48272ea823d9e5b52395e5230d6ff75e46c41ce3d0"></a>

## name property — Property reference / 978e3bdb6a7e / 9

Type: `"string"`. Required.

Name of the Network Policy. Must be unique within the namespace.

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

<a id="canonical-a8e456b4d09a3f4393d2b7096871cd81782a53806a031662e96648bb62a85fac"></a>

<a id="canonical-af4652a14e55f7dcef0befa7b23e923cb51102cdf4d738f317d7f15c3a0e6f09"></a>

## namespace property — Property reference / 978e3bdb6a7e / 10

Type: `"string"`. Required.

Namespace where the Network Policy is created.

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

- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1): complete subsection reference.

- [timeouts](resources--network_policy--reference--group-001.md#canonical-1b1756d0c2a28e1e88cc694220b1e946446c2046746e8e3e176999919bc2aeb8): complete subsection reference.

<a id="canonical-a7961723bd19c14226b5dd5e7ad25ef7599a2e6885f47d7493e367a5997501de"></a>

## All schema paths — Property reference / 978e3bdb6a7e / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_policy--reference--group-001.md#canonical-48f9f10c8407d5ba047afd42bc3491ada662000da62e6129053cb52338145aec) |
| `description` | [description](resources--network_policy--reference--group-001.md#canonical-794071bbc8461baa856bb2167a6c45c5a658b986f621c5d20cbf8f1cc7c3f26b) |
| `disable` | [disable](resources--network_policy--reference--group-001.md#canonical-8d37b8dc0a5038fe140e6f120abec8c503690c6d94b70dbc482e459d1d7f5950) |
| `endpoint` | [endpoint](resources--network_policy--reference--group-001.md#canonical-c2fc6108ee86710fd6def8ba004dfc4782745135076b09df0c69afc2bfe1af15) |
| `endpoint.any` | [endpoint.any](resources--network_policy--reference--group-001.md#canonical-0a646dc85d65a2585abf56d047a4197059cc8371439570e2159c7720405593fa) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-77ddb38620d3435586c816d9206aaef4eebebe526b1f50ad463d81b53194fa2d) |
| `endpoint.label_selector` | [endpoint.label_selector](resources--network_policy--reference--group-001.md#canonical-2eca9f0d8847acbf909483c9d4814a9eb28e5e8e0693ab41234e9a5f511afe24) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](resources--network_policy--reference--group-001.md#canonical-89a5ab018d175cbb3eeef404ceaae6c8a621c14199d93d57ad92331bb476b7d3) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-fab89b126d956733e793a3d8db7601c6dede2d91fd256ab9603b4051314d3422) |
| `endpoint.prefix_list` | [endpoint.prefix_list](resources--network_policy--reference--group-001.md#canonical-850371ad6872fc7efb09309a2ceb86c96a61c454a8ebd9043e9f0cce67920900) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](resources--network_policy--reference--group-001.md#canonical-506e21019678bcaab887450726235ee8c0a7c9e135e16b1a1256112577091290) |
| `id` | [id](resources--network_policy--reference--group-001.md#canonical-5e3687873c714cfe3c67abe1198810e53c4d1b27a0c5ebaa9d2fd62f2e7176c1) |
| `labels` | [labels](resources--network_policy--reference--group-001.md#canonical-577b976cd4be9c45b20ec46fe67800f3b21f19cf68d9acc271b9505ae8fee28b) |
| `name` | [name](resources--network_policy--reference--group-001.md#canonical-49db2e0ce287acbc6dd141c7cae3671c31e41c9147902b438b4a393c8013901e) |
| `namespace` | [namespace](resources--network_policy--reference--group-001.md#canonical-a8e456b4d09a3f4393d2b7096871cd81782a53806a031662e96648bb62a85fac) |
| `rules` | [rules](resources--network_policy--reference--group-001.md#canonical-d520dbf5b4ba6d6dc94bd1d45510aaa7dd92f5508bcd677f9c42a94c2128c0c6) |
| `rules.egress_rules` | [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-2a57fc3ae9548f663b38bee1c47e6be23133db006b8841eba91aa61031802717) |
| `rules.egress_rules.action` | [rules.egress_rules.action](resources--network_policy--reference--group-001.md#canonical-6e3bdf0fd3bc6465a764f93b3fa878215423dd4770e5c598c92decd2e8cbb61e) |
| `rules.egress_rules.adv_action` | [rules.egress_rules.adv_action](resources--network_policy--reference--group-001.md#canonical-e3389b187d85e017315c394910a88273d158796e35f313b8bebb6b820f7f6053) |
| `rules.egress_rules.adv_action.action` | [rules.egress_rules.adv_action.action](resources--network_policy--reference--group-001.md#canonical-4bfc55c4d398852b2119362cedd00033a0ab0b2188dd5151378b294e28ed5202) |
| `rules.egress_rules.all_tcp_traffic` | [rules.egress_rules.all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-f0b1f64da3bc3eaf0dca2fe7a2fe663861b306c120c2343caa1d30707815afc1) |
| `rules.egress_rules.all_traffic` | [rules.egress_rules.all_traffic](resources--network_policy--reference--group-001.md#canonical-6bd615307d83013b56f4e99ce3770e2b9ef322561d6de8f6321010b3a642f020) |
| `rules.egress_rules.all_udp_traffic` | [rules.egress_rules.all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-52f3856f0657ab4243feb821b83b307196bfbc891b92133be9e750869df39316) |
| `rules.egress_rules.any` | [rules.egress_rules.any](resources--network_policy--reference--group-001.md#canonical-efbf4f01d1c0d0c4e03a156c4454893ee7824158f828762025c0f86ee80d6340) |
| `rules.egress_rules.applications` | [rules.egress_rules.applications](resources--network_policy--reference--group-001.md#canonical-ad5939b1d1fdf8219bb09fcda44d522b9f5cc1df69a508989f2b4d965738443c) |
| `rules.egress_rules.applications.applications` | [rules.egress_rules.applications.applications](resources--network_policy--reference--group-001.md#canonical-9da6ddd9443600cd6f0ae8961064bcd2eea02ddc76e210d4bd10a634ffd20679) |
| `rules.egress_rules.inside_endpoints` | [rules.egress_rules.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-1d7713c3c70c5195c511be334fc3afd3d8b6f91cdc4e188cfa6ddafc21b28a1c) |
| `rules.egress_rules.ip_prefix_set` | [rules.egress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-f7786c7dffa2d6de13323762a28ecf3ede90fd787e6519f0595e0e2a37fd6baf) |
| `rules.egress_rules.ip_prefix_set.ref` | [rules.egress_rules.ip_prefix_set.ref](resources--network_policy--reference--group-001.md#canonical-6e3df8118cd4097f1b8e2cbf21bb7e4af04c8a64cdaa50ae9664c603614bb100) |
| `rules.egress_rules.ip_prefix_set.ref.kind` | [rules.egress_rules.ip_prefix_set.ref.kind](resources--network_policy--reference--group-001.md#canonical-bf142740783632ad5db733abb07b139936f250d0951b6a9e064d579159f7a80e) |
| `rules.egress_rules.ip_prefix_set.ref.name` | [rules.egress_rules.ip_prefix_set.ref.name](resources--network_policy--reference--group-001.md#canonical-d1f3b32917d70dc89b4ff0f7abb7d947e8104fd5489d28c2361e5c1ab11d911c) |
| `rules.egress_rules.ip_prefix_set.ref.namespace` | [rules.egress_rules.ip_prefix_set.ref.namespace](resources--network_policy--reference--group-001.md#canonical-4a87d013695bd0528c5ecc59a01e9889db4ed3a0b4fe1a4c03912aa2eed14ccf) |
| `rules.egress_rules.ip_prefix_set.ref.tenant` | [rules.egress_rules.ip_prefix_set.ref.tenant](resources--network_policy--reference--group-001.md#canonical-2814a87362ceecf039ad640bfc05ef3b1a31fa9a8670f42ef4ef4558d8a892e6) |
| `rules.egress_rules.ip_prefix_set.ref.uid` | [rules.egress_rules.ip_prefix_set.ref.uid](resources--network_policy--reference--group-001.md#canonical-999b4806205fcb3a619fc48277004107ab0421eed5ce8db58a6ff02c03087865) |
| `rules.egress_rules.label_matcher` | [rules.egress_rules.label_matcher](resources--network_policy--reference--group-001.md#canonical-182d39d0faa38e669e70068c56502d2d510a92ce67fcc1b75383fb83d53394cb) |
| `rules.egress_rules.label_matcher.keys` | [rules.egress_rules.label_matcher.keys](resources--network_policy--reference--group-001.md#canonical-76f7c4292b6ce1283853e8eb3f0019caa62068da63be7f8f5992cfd57d0bb8a3) |
| `rules.egress_rules.label_selector` | [rules.egress_rules.label_selector](resources--network_policy--reference--group-001.md#canonical-3e9d427b90035d4301bf1aa9403e552c3bc28d1b121aa5dac1e7ea9d3a82af46) |
| `rules.egress_rules.label_selector.expressions` | [rules.egress_rules.label_selector.expressions](resources--network_policy--reference--group-001.md#canonical-17578fa2a9fecf6c4b6f8ab206381ec2814050f113f596f8716065cfefc7ceb8) |
| `rules.egress_rules.metadata` | [rules.egress_rules.metadata](resources--network_policy--reference--group-001.md#canonical-6e344b05851ef80915ae6f1a83780728552f09ea0086c58c1b7a47bfef4724fb) |
| `rules.egress_rules.metadata.description_spec` | [rules.egress_rules.metadata.description_spec](resources--network_policy--reference--group-001.md#canonical-f1b5a4eabb10871178aacda086c503c79401a2ce6bd1297facf2e2929bed05a4) |
| `rules.egress_rules.metadata.name` | [rules.egress_rules.metadata.name](resources--network_policy--reference--group-001.md#canonical-d6b9aa170a359437916d945888686f980e441f2293ce05714ffe0a2e4ba4212d) |
| `rules.egress_rules.outside_endpoints` | [rules.egress_rules.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-c4a42c078cdb717d12077217f18390129dfec8eed6889569332fdff33074c8fa) |
| `rules.egress_rules.prefix_list` | [rules.egress_rules.prefix_list](resources--network_policy--reference--group-001.md#canonical-abaecdf1b073b00bf3e4eede16f888a988602e540b9e744079855e5f99c31d35) |
| `rules.egress_rules.prefix_list.prefixes` | [rules.egress_rules.prefix_list.prefixes](resources--network_policy--reference--group-001.md#canonical-2833e9610ebf65cfdc5d6957790ecca05b03c1d30eefc7eadd6c0dcddddfe7f0) |
| `rules.egress_rules.protocol_port_range` | [rules.egress_rules.protocol_port_range](resources--network_policy--reference--group-001.md#canonical-957c534e0e0e3c9139e07314b1d92a0a80d8ca949574f7b204062bbcc9992bef) |
| `rules.egress_rules.protocol_port_range.port_ranges` | [rules.egress_rules.protocol_port_range.port_ranges](resources--network_policy--reference--group-001.md#canonical-d80cebdbe9d29f622b67c66e2cb87af6314d4c25128bf0c87ca1c796768d233f) |
| `rules.egress_rules.protocol_port_range.protocol` | [rules.egress_rules.protocol_port_range.protocol](resources--network_policy--reference--group-001.md#canonical-554dcb051d1ff850e9b9866fe927eeac590207ff6b6ffb3baa872847101f08bc) |
| `rules.ingress_rules` | [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-d4fba6a7307e25a2a49b3f669ec8a0b9d5e0c946b4363491ed9da2c70e0f3f4b) |
| `rules.ingress_rules.action` | [rules.ingress_rules.action](resources--network_policy--reference--group-001.md#canonical-09e21ccbd5cbaeec7d4062e71d46e231b4f424fd623f8bac0e763931d41a3f8e) |
| `rules.ingress_rules.adv_action` | [rules.ingress_rules.adv_action](resources--network_policy--reference--group-001.md#canonical-681daf8927f4ca61e628ab815ae0db5c47cf3e3c461528a23451f35b1d8edb20) |
| `rules.ingress_rules.adv_action.action` | [rules.ingress_rules.adv_action.action](resources--network_policy--reference--group-001.md#canonical-f4966a4ef4e7c05e4407e55b7e5a1da91741cfadc29bf62366f8d7c14ac95104) |
| `rules.ingress_rules.all_tcp_traffic` | [rules.ingress_rules.all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-b24f161cb3f9e950a86b0d63aa807c6087b57d1be471bbc1edcda93259f8a125) |
| `rules.ingress_rules.all_traffic` | [rules.ingress_rules.all_traffic](resources--network_policy--reference--group-001.md#canonical-7654d6e53cd7c6a746531f99b4156f0e0440a9ad58fa7b60459e2aa3332479db) |
| `rules.ingress_rules.all_udp_traffic` | [rules.ingress_rules.all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-cd41ff73730c34f3469b268e87701e7fe431ce323ae931a3bf873444c804aa9b) |
| `rules.ingress_rules.any` | [rules.ingress_rules.any](resources--network_policy--reference--group-001.md#canonical-9a876fcdd011c13a1396f4fa89085507971bd113912c4578934e83d0d6e7bec0) |
| `rules.ingress_rules.applications` | [rules.ingress_rules.applications](resources--network_policy--reference--group-001.md#canonical-f79523d68e7e9676cd6711f4046de251b7b0ac2456c78c08d4a26c3202b8f384) |
| `rules.ingress_rules.applications.applications` | [rules.ingress_rules.applications.applications](resources--network_policy--reference--group-001.md#canonical-f039c72a36f9f3a0fead14834264b265e80bf71a90b0dfc44d1d0b1f0f30fdc0) |
| `rules.ingress_rules.inside_endpoints` | [rules.ingress_rules.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-23bde8769bff6921112d0e8406b39c8fb692d99b80ffc533e62a0ede90fbff8b) |
| `rules.ingress_rules.ip_prefix_set` | [rules.ingress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-aa5b637af7da336abba8b7e191afb53da274a861c996dad6143b26f50d2f5dec) |
| `rules.ingress_rules.ip_prefix_set.ref` | [rules.ingress_rules.ip_prefix_set.ref](resources--network_policy--reference--group-001.md#canonical-ea314b9e3aa12e86dc3aa3889f45c8acc014ca481ce4cdc07933b55c12963879) |
| `rules.ingress_rules.ip_prefix_set.ref.kind` | [rules.ingress_rules.ip_prefix_set.ref.kind](resources--network_policy--reference--group-001.md#canonical-1bbecedd0d5850c5e0c4ef8aaf02fc71c68976f1e25ed60884784454ab618e44) |
| `rules.ingress_rules.ip_prefix_set.ref.name` | [rules.ingress_rules.ip_prefix_set.ref.name](resources--network_policy--reference--group-001.md#canonical-f57d213fe8ef18a4bc371f2988b8d201f3158ff89d5c5d1f0b886baf9c8d405f) |
| `rules.ingress_rules.ip_prefix_set.ref.namespace` | [rules.ingress_rules.ip_prefix_set.ref.namespace](resources--network_policy--reference--group-001.md#canonical-1971d77f1e80eea5d8a7142e022aa6965ae23a93eb4f4d377d9a7753a738a5f1) |
| `rules.ingress_rules.ip_prefix_set.ref.tenant` | [rules.ingress_rules.ip_prefix_set.ref.tenant](resources--network_policy--reference--group-001.md#canonical-655bb79b72af35a834ecd88d868bb8585b0518e35ff24cb9ac3a0d837a8c8767) |
| `rules.ingress_rules.ip_prefix_set.ref.uid` | [rules.ingress_rules.ip_prefix_set.ref.uid](resources--network_policy--reference--group-001.md#canonical-0935a943df2d9bcee3e513b938983292fa7c6976571dc4c9792ec65d59c1b315) |
| `rules.ingress_rules.label_matcher` | [rules.ingress_rules.label_matcher](resources--network_policy--reference--group-001.md#canonical-7f71ce56846e007d1880d338a6fc14cbb3c551bd267a103f5cd9f244a2cacc46) |
| `rules.ingress_rules.label_matcher.keys` | [rules.ingress_rules.label_matcher.keys](resources--network_policy--reference--group-001.md#canonical-9723a1904814d4502d1e2202773b628826d29b51b8e172477b98e561205cac5f) |
| `rules.ingress_rules.label_selector` | [rules.ingress_rules.label_selector](resources--network_policy--reference--group-001.md#canonical-ea2a315a1b458f88aff51105e297b872db463c14f7aee0e06c765e4837f3c1df) |
| `rules.ingress_rules.label_selector.expressions` | [rules.ingress_rules.label_selector.expressions](resources--network_policy--reference--group-001.md#canonical-706965ec4afb407cbf7f029e72c98cf5b8a7d38834177e81247a0e7641726555) |
| `rules.ingress_rules.metadata` | [rules.ingress_rules.metadata](resources--network_policy--reference--group-001.md#canonical-33fb34d2c13ba66bb8c8bb043041fb9d48e9564ab0651fb626004a03d3dc737a) |
| `rules.ingress_rules.metadata.description_spec` | [rules.ingress_rules.metadata.description_spec](resources--network_policy--reference--group-001.md#canonical-f346c9bcbaa44f80ce88c9f10a2a9be6ed89964453fbdd8a4264f370af66d996) |
| `rules.ingress_rules.metadata.name` | [rules.ingress_rules.metadata.name](resources--network_policy--reference--group-001.md#canonical-7d59448d8a8cf850dbadea0c72085c5701bc81985d7f466aa2e49ee3761f8738) |
| `rules.ingress_rules.outside_endpoints` | [rules.ingress_rules.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-a6b219a5653ef90ae48934dabf4d3d58798cb235e98ecd02eb876a791b51b4fa) |
| `rules.ingress_rules.prefix_list` | [rules.ingress_rules.prefix_list](resources--network_policy--reference--group-001.md#canonical-862911aeb304d2ff8029f956cf09a9e6955f87ee8cfa71bdf1453815f88a4cbc) |
| `rules.ingress_rules.prefix_list.prefixes` | [rules.ingress_rules.prefix_list.prefixes](resources--network_policy--reference--group-001.md#canonical-167109372e9d78a8c982985b8e018dd5df273025c254cd33e6e6c5ba15fecc9f) |
| `rules.ingress_rules.protocol_port_range` | [rules.ingress_rules.protocol_port_range](resources--network_policy--reference--group-001.md#canonical-0868bd0390f1852d13c9d64fdca0958704fcac3a8155ed0da6a7c69f97f6627d) |
| `rules.ingress_rules.protocol_port_range.port_ranges` | [rules.ingress_rules.protocol_port_range.port_ranges](resources--network_policy--reference--group-001.md#canonical-7966f8cecba51f55ab03212c80ccc8a3055de13dc8a10e26b5fc64570bff38eb) |
| `rules.ingress_rules.protocol_port_range.protocol` | [rules.ingress_rules.protocol_port_range.protocol](resources--network_policy--reference--group-001.md#canonical-1522f5c7957b612fd96d86aedeb49a209e0593d2f1b140fb5cd992450742024c) |
| `timeouts` | [timeouts](resources--network_policy--reference--group-001.md#canonical-5d6714412be8be94de068037b30f1316147ea379858c8543c79cd963182d4865) |
| `timeouts.create` | [timeouts.create](resources--network_policy--reference--group-001.md#canonical-f1c0b2ddabb31777a523e0023df63262039e697aa5f8d1174358a9e3f35f50cc) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy--reference--group-001.md#canonical-3da976b1e3ee598e7fa45f590a98e60bd998f86649c20586353d02ae4b3cc1e4) |
| `timeouts.read` | [timeouts.read](resources--network_policy--reference--group-001.md#canonical-79527f77021a17323d4dc6968e26e495ff527f7b31367f222fe4f64dbc108301) |
| `timeouts.update` | [timeouts.update](resources--network_policy--reference--group-001.md#canonical-6893e67ea8b9bbb82e63d809d45c68ff890b9a4e93bc6af1cb23fbb4ca8b75e3) |

<a id="canonical-f258e154011043f114168df75789af612ab40114f29b09dbb173395a3c4a1b80"></a>

## Next pages — Property reference / 978e3bdb6a7e / 12

- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [timeouts](resources--network_policy--reference--group-001.md#canonical-1b1756d0c2a28e1e88cc694220b1e946446c2046746e8e3e176999919bc2aeb8)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-829856dc9281d491bbd2044d0e2201fb747c2a1755c7c4f0e45758bac30ab684"></a>

## endpoint — endpoint / 39de730699cd / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- endpoint

<a id="canonical-c2fc6108ee86710fd6def8ba004dfc4782745135076b09df0c69afc2bfe1af15"></a>

Type: `"object"`. single nested block, Optional.

Shape of the endpoint choices for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingObjectAttributes("outside_endpoints",
    "prefix_list")}
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
  "x-ves-oneof-field-endpoint_choice": "[\"any\",\"inside_endpoints\",\"label_selector\",\"outside_endpoints\",\"prefix_list\"]"
}
```

Terraform syntax:

```terraform
endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7b6b1ed3adfa69a4cdc38c21d3a1511b41b040329c9ebd52ea57ef3c2484e79"></a>

## Direct properties — endpoint / 39de730699cd / 3

- [any](resources--network_policy--reference--group-001.md#canonical-19c8b55f3c41d1d4e408c5515ffbc7b9b636b743fbba92dcdeab4bbde748cc17): complete subsection reference.

- [inside_endpoints](resources--network_policy--reference--group-001.md#canonical-2ea8fc7445149f441a60fff2f09c837eb2b409006d48a434eeafa5d1f7760c0d): complete subsection reference.

- [label_selector](resources--network_policy--reference--group-001.md#canonical-cbbbb4abb5eb2ca198e1380cfc90b5d06de9dd3ecb04bbfb913a183889e0b79a): complete subsection reference.

- [outside_endpoints](resources--network_policy--reference--group-001.md#canonical-b97f2b17244b5bffadb9e9dbf4c5eb1a2542e6c3c038024cf2139b62823d6681): complete subsection reference.

- [prefix_list](resources--network_policy--reference--group-001.md#canonical-1145fbedd0764e046354be29f9236b7a28d6c5535dfe5eda63e6955c4e59fbe0): complete subsection reference.

<a id="canonical-f50dddd849be7062a7178e4ff34b4bc4f04d55635d1fce08ec9bd2e5935587ea"></a>

## Next pages — endpoint / 39de730699cd / 4

- [endpoint.any](resources--network_policy--reference--group-001.md#canonical-19c8b55f3c41d1d4e408c5515ffbc7b9b636b743fbba92dcdeab4bbde748cc17)
- [endpoint.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-2ea8fc7445149f441a60fff2f09c837eb2b409006d48a434eeafa5d1f7760c0d)
- [endpoint.label_selector](resources--network_policy--reference--group-001.md#canonical-cbbbb4abb5eb2ca198e1380cfc90b5d06de9dd3ecb04bbfb913a183889e0b79a)
- [endpoint.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-b97f2b17244b5bffadb9e9dbf4c5eb1a2542e6c3c038024cf2139b62823d6681)
- [endpoint.prefix_list](resources--network_policy--reference--group-001.md#canonical-1145fbedd0764e046354be29f9236b7a28d6c5535dfe5eda63e6955c4e59fbe0)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-19c8b55f3c41d1d4e408c5515ffbc7b9b636b743fbba92dcdeab4bbde748cc17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75485309ca89fca56bf8aa0040dd1fdf198788134a52125d190a130bd9e543b8"></a>

## endpoint.any — endpoint.any / 2631c8ec160d / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- endpoint.any

<a id="canonical-0a646dc85d65a2585abf56d047a4197059cc8371439570e2159c7720405593fa"></a>

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

<a id="canonical-73a4dfbad83d24710ab6542bd4e19989380fda3ee76b46e10453259663f17001"></a>

## Direct properties — endpoint.any / 2631c8ec160d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-91991dc6b0f6dbca7ff430ad4f6c96d0cf679db3051308e21afe1111d095e98d"></a>

## Next pages — endpoint.any / 2631c8ec160d / 4

- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-2ea8fc7445149f441a60fff2f09c837eb2b409006d48a434eeafa5d1f7760c0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7e4777ab3d3bc6d02997e67b01715cc73dda5751e35987e1c7575c59617fdd8"></a>

## endpoint.inside_endpoints — endpoint.inside_endpoints / b024687c9c35 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- endpoint.inside_endpoints

<a id="canonical-77ddb38620d3435586c816d9206aaef4eebebe526b1f50ad463d81b53194fa2d"></a>

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
inside_endpoints = {}
```

<a id="canonical-c60534c73bced7242523531c223631ee1dbb038e1c0040817cb93d891b5b366d"></a>

## Direct properties — endpoint.inside_endpoints / b024687c9c35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d308a7622595a01a8c82b9692c76ef894386d110618e3d3cf9ccf17c831a377"></a>

## Next pages — endpoint.inside_endpoints / b024687c9c35 / 4

- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-cbbbb4abb5eb2ca198e1380cfc90b5d06de9dd3ecb04bbfb913a183889e0b79a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e588f6b36156e4ee0509a180d17ead3ff38a109d7469e007707369d1b6c23b9"></a>

## endpoint.label_selector — endpoint.label_selector / eee1c064fb31 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- endpoint.label_selector

<a id="canonical-2eca9f0d8847acbf909483c9d4814a9eb28e5e8e0693ab41234e9a5f511afe24"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c789ce1baf7d7544c97673d34285455bc5706b1c45b8cca538077935b4c08ec"></a>

## Direct properties — endpoint.label_selector / eee1c064fb31 / 3

<a id="canonical-89a5ab018d175cbb3eeef404ceaae6c8a621c14199d93d57ad92331bb476b7d3"></a>

<a id="canonical-0949fa116c125aa4fa227f56410e1b0c53680121764501cf7d081f4dfb8ed20f"></a>

## expressions property — endpoint.label_selector / eee1c064fb31 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-e10fb583300bd1cfad64fa7f0ea21c6eb10811d62f9331a6a3d697a0975c1c42"></a>

## Next pages — endpoint.label_selector / eee1c064fb31 / 5

- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-b97f2b17244b5bffadb9e9dbf4c5eb1a2542e6c3c038024cf2139b62823d6681"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fbb255ab9f51e164003f423b555020132a143e01c508f32f4f4ec255562ace2"></a>

## endpoint.outside_endpoints — endpoint.outside_endpoints / 7311990e523f / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- endpoint.outside_endpoints

<a id="canonical-fab89b126d956733e793a3d8db7601c6dede2d91fd256ab9603b4051314d3422"></a>

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
outside_endpoints = {}
```

<a id="canonical-b0216ecfce0ba59bee12c4557b01b97f29ce8cab5dd8b555ce3fe2bd91068f24"></a>

## Direct properties — endpoint.outside_endpoints / 7311990e523f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea36e3164a920831fc1681cbc6a9ca308ac85a4d0798f9d384c83618d2e6efe8"></a>

## Next pages — endpoint.outside_endpoints / 7311990e523f / 4

- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-1145fbedd0764e046354be29f9236b7a28d6c5535dfe5eda63e6955c4e59fbe0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14ea87c21492dffb93a15a4f951684341942ab756491c1c12fa28c810d907ba7"></a>

## endpoint.prefix_list — endpoint.prefix_list / ef03f54cd16f / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- endpoint.prefix_list

<a id="canonical-850371ad6872fc7efb09309a2ceb86c96a61c454a8ebd9043e9f0cce67920900"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f9df32a213399a642fbd47b5fbe5e5f9d5f9ef77136bbfe1194945fe7502815"></a>

## Direct properties — endpoint.prefix_list / ef03f54cd16f / 3

<a id="canonical-506e21019678bcaab887450726235ee8c0a7c9e135e16b1a1256112577091290"></a>

<a id="canonical-cf4973d8884c734b4150589cddc591e0f69343d9df23334a913fd5a140622a1f"></a>

## prefixes property — endpoint.prefix_list / ef03f54cd16f / 4

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

<a id="canonical-dbc3f2720b0810f4bdae6b0442ff3a32b5d07cd3339661cfa9c3640f775acc05"></a>

## Next pages — endpoint.prefix_list / ef03f54cd16f / 5

- [endpoint](resources--network_policy--reference--group-001.md#canonical-ac6dd0690321bedca454d30e18e90105daf46fd4674b79282bfe27af264eecc7)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-991aeb8ed347b7ba096ef6144f2f1abdc22b327316e2504b9a32684789344830"></a>

## rules — rules / 3d50ce972ab3 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- rules

<a id="canonical-d520dbf5b4ba6d6dc94bd1d45510aaa7dd92f5508bcd677f9c42a94c2128c0c6"></a>

Type: `"object"`. single nested block, Optional.

Rule Choice. Shape of Rule Choice.

Upstream description:

Shape of Rule Choice.

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
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-63f6d54aaa9c27902101cdd6963fc045619ffb08aec044a6d80890148b8d3b01"></a>

## Direct properties — rules / 3d50ce972ab3 / 3

- [egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3): complete subsection reference.

- [ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df): complete subsection reference.

<a id="canonical-39ba79dbe833fb2c9566df6528354839813461435d539d4142913fa830c1a891"></a>

## Next pages — rules / 3d50ce972ab3 / 4

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8bd956a991d096a39b5efe0b84d8303d2b5d9f915c437924e84629e64b69882"></a>

## rules.egress_rules — rules.egress_rules / 9a9e349dbbce / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- rules.egress_rules

<a id="canonical-2a57fc3ae9548f663b38bee1c47e6be23133db006b8841eba91aa61031802717"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections from policy endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
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
egress_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-ae87ceb0dae6638b788418c68cae1b5fe3fb2a3b7748249d9651cf2000f85c6a"></a>

## Direct properties — rules.egress_rules / 9a9e349dbbce / 3

<a id="canonical-6e3bdf0fd3bc6465a764f93b3fa878215423dd4770e5c598c92decd2e8cbb61e"></a>

<a id="canonical-d4a49a944d588444951e3639fed88398a8c42e68ac4583bb4985ec0f3754a858"></a>

## action property — rules.egress_rules / 9a9e349dbbce / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

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

- [adv_action](resources--network_policy--reference--group-001.md#canonical-d73c8de3d304d1fa7f2a7ec9f46503dd80bf1efc72c11e06a54044750544f0f7): complete subsection reference.

- [all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-38924ff3b0919bc60b8add6d372d9ee7dfce6a23cdd17da5df2fed2ad71524bb): complete subsection reference.

- [all_traffic](resources--network_policy--reference--group-001.md#canonical-cd57880b35e58e57215418a7f14f04a920606500fb319b315e7217fe5383722c): complete subsection reference.

- [all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-da8e62d03bccdcba54c8671d33f390473f994eef2f031175d38b0afc7474d9fc): complete subsection reference.

- [any](resources--network_policy--reference--group-001.md#canonical-f5555c24a6d5d7f0f35e5bef6378bc4e1240eaaff614df47cdea1b092603a176): complete subsection reference.

- [applications](resources--network_policy--reference--group-001.md#canonical-bae11aa47e185f01d2a43b06c9f2932bd30c1e33130ae7127548fa300608b724): complete subsection reference.

- [inside_endpoints](resources--network_policy--reference--group-001.md#canonical-de941d88c7f6fd5dadf7d2948c0647a8215e8bf7fec0e38fa5c2acd565869667): complete subsection reference.

- [ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-1e4c43354fbd01b2193e7d05b626ef6375586ba3f56829bb7f589f12ea935185): complete subsection reference.

- [label_matcher](resources--network_policy--reference--group-001.md#canonical-37cc1b7067ea7fe251cce72615bd16f7553d7cd785f5fedf122052d593e16373): complete subsection reference.

- [label_selector](resources--network_policy--reference--group-001.md#canonical-3782f29a45559a7cb6ba46f1df4e5a0f9d82bca638a68232a4dde6cc4cdd84a6): complete subsection reference.

- [metadata](resources--network_policy--reference--group-001.md#canonical-37d9ae63737878871475550eeee02ecb98e1e637f66cfa01998738eabbcc336d): complete subsection reference.

- [outside_endpoints](resources--network_policy--reference--group-001.md#canonical-94ec9fcc67f80bb987da23ce458b3f719062f4ea086859b8c14022f38eeab4ab): complete subsection reference.

- [prefix_list](resources--network_policy--reference--group-001.md#canonical-94bd8b75ab142c767843b3aaddd7798eee4e99d300753aaf7ec12cb111ef8fd0): complete subsection reference.

- [protocol_port_range](resources--network_policy--reference--group-001.md#canonical-869dddae57d8e65c24edad07a172858bc70b28eff189bcd4ab4a2dd5bbc3145a): complete subsection reference.

<a id="canonical-0f3c5416f7d60a204c054317a33052980865f41ba3a0614fcef588f2e61154ae"></a>

## Next pages — rules.egress_rules / 9a9e349dbbce / 5

- [rules.egress_rules.adv_action](resources--network_policy--reference--group-001.md#canonical-d73c8de3d304d1fa7f2a7ec9f46503dd80bf1efc72c11e06a54044750544f0f7)
- [rules.egress_rules.all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-38924ff3b0919bc60b8add6d372d9ee7dfce6a23cdd17da5df2fed2ad71524bb)
- [rules.egress_rules.all_traffic](resources--network_policy--reference--group-001.md#canonical-cd57880b35e58e57215418a7f14f04a920606500fb319b315e7217fe5383722c)
- [rules.egress_rules.all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-da8e62d03bccdcba54c8671d33f390473f994eef2f031175d38b0afc7474d9fc)
- [rules.egress_rules.any](resources--network_policy--reference--group-001.md#canonical-f5555c24a6d5d7f0f35e5bef6378bc4e1240eaaff614df47cdea1b092603a176)
- [rules.egress_rules.applications](resources--network_policy--reference--group-001.md#canonical-bae11aa47e185f01d2a43b06c9f2932bd30c1e33130ae7127548fa300608b724)
- [rules.egress_rules.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-de941d88c7f6fd5dadf7d2948c0647a8215e8bf7fec0e38fa5c2acd565869667)
- [rules.egress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-1e4c43354fbd01b2193e7d05b626ef6375586ba3f56829bb7f589f12ea935185)
- [rules.egress_rules.label_matcher](resources--network_policy--reference--group-001.md#canonical-37cc1b7067ea7fe251cce72615bd16f7553d7cd785f5fedf122052d593e16373)
- [rules.egress_rules.label_selector](resources--network_policy--reference--group-001.md#canonical-3782f29a45559a7cb6ba46f1df4e5a0f9d82bca638a68232a4dde6cc4cdd84a6)
- [rules.egress_rules.metadata](resources--network_policy--reference--group-001.md#canonical-37d9ae63737878871475550eeee02ecb98e1e637f66cfa01998738eabbcc336d)
- [rules.egress_rules.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-94ec9fcc67f80bb987da23ce458b3f719062f4ea086859b8c14022f38eeab4ab)
- [rules.egress_rules.prefix_list](resources--network_policy--reference--group-001.md#canonical-94bd8b75ab142c767843b3aaddd7798eee4e99d300753aaf7ec12cb111ef8fd0)
- [rules.egress_rules.protocol_port_range](resources--network_policy--reference--group-001.md#canonical-869dddae57d8e65c24edad07a172858bc70b28eff189bcd4ab4a2dd5bbc3145a)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-d73c8de3d304d1fa7f2a7ec9f46503dd80bf1efc72c11e06a54044750544f0f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52fe227b3842ad271834ae85358c09142718b0ad8348fcc5e94a7cbf13a2d737"></a>

## rules.egress_rules.adv_action — rules.egress_rules.adv_action / 80e8a83dbc35 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.adv_action

<a id="canonical-e3389b187d85e017315c394910a88273d158796e35f313b8bebb6b820f7f6053"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
adv_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-63c8c3f13812be38c1d6e7e5d0fab5801c10dc073b6f297ff59e80f5c11ec021"></a>

## Direct properties — rules.egress_rules.adv_action / 80e8a83dbc35 / 3

<a id="canonical-4bfc55c4d398852b2119362cedd00033a0ab0b2188dd5151378b294e28ed5202"></a>

<a id="canonical-ccc7b5b4b201969cafa3accd037119e2bd4d4f6a2395c4e152f191efa6c5e802"></a>

## action property — rules.egress_rules.adv_action / 80e8a83dbc35 / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4ef19e920e40b958ea546c68aafa702163b56c3b2cc13dff1f823d68494af938"></a>

## Next pages — rules.egress_rules.adv_action / 80e8a83dbc35 / 5

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-38924ff3b0919bc60b8add6d372d9ee7dfce6a23cdd17da5df2fed2ad71524bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11bcd645f6f846b54e2ff4a05e792ec7639c5742a9798cee25faffd5818458aa"></a>

## rules.egress_rules.all_tcp_traffic — rules.egress_rules.all_tcp_traffic / 3c3efe32267b / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.all_tcp_traffic

<a id="canonical-f0b1f64da3bc3eaf0dca2fe7a2fe663861b306c120c2343caa1d30707815afc1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

<a id="canonical-06a268e9101a060795e1ab435d4094650e8654214e4e77e52d6dcf63e3185f0f"></a>

## Direct properties — rules.egress_rules.all_tcp_traffic / 3c3efe32267b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0993464db4b467c1bfe698f839bc053c1523a37ae0fdf5c0aa5bbf9208d51a20"></a>

## Next pages — rules.egress_rules.all_tcp_traffic / 3c3efe32267b / 4

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-cd57880b35e58e57215418a7f14f04a920606500fb319b315e7217fe5383722c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54e943ffcb5c0ecd387c9251c62724d034b577dbc17f501cd04990879780125c"></a>

## rules.egress_rules.all_traffic — rules.egress_rules.all_traffic / 7bc9503d3332 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.all_traffic

<a id="canonical-6bd615307d83013b56f4e99ce3770e2b9ef322561d6de8f6321010b3a642f020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

<a id="canonical-c52ed5b9e506c1ce41a9b644bd01b43a26735f6accbfb99dace2da5cc790bba2"></a>

## Direct properties — rules.egress_rules.all_traffic / 7bc9503d3332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e356908759e0db1e7b18d1f486d0fda6fd710ec7b3e583de63140f8e337a801c"></a>

## Next pages — rules.egress_rules.all_traffic / 7bc9503d3332 / 4

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-da8e62d03bccdcba54c8671d33f390473f994eef2f031175d38b0afc7474d9fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b495ace521859a2dc8302e1019268d2bca3ad7ae52f2c3112775adcfd6d49c03"></a>

## rules.egress_rules.all_udp_traffic — rules.egress_rules.all_udp_traffic / 5132296a6177 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.all_udp_traffic

<a id="canonical-52f3856f0657ab4243feb821b83b307196bfbc891b92133be9e750869df39316"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

<a id="canonical-073481886a1a05980d1e2e66249cba474472c61434acd4d6565016507686b002"></a>

## Direct properties — rules.egress_rules.all_udp_traffic / 5132296a6177 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c1020674c0a19d837d2e7d420a5a2c394f69fcda4336289ebbdf125c6dd6bcaf"></a>

## Next pages — rules.egress_rules.all_udp_traffic / 5132296a6177 / 4

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-f5555c24a6d5d7f0f35e5bef6378bc4e1240eaaff614df47cdea1b092603a176"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3d7a2234b88778cd47a6f666b2de430591820150b19a264394006e527cc138a"></a>

## rules.egress_rules.any — rules.egress_rules.any / f47ff6a24db3 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.any

<a id="canonical-efbf4f01d1c0d0c4e03a156c4454893ee7824158f828762025c0f86ee80d6340"></a>

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

<a id="canonical-1b6d8253d925682e3e2ef6c18b9c4eb1053cb50496150160c3680b629442cca9"></a>

## Direct properties — rules.egress_rules.any / f47ff6a24db3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1a1e8be8b18fb115414befafa9a4ab3f58d1004e23dacb9cb6f0d2227ec885d"></a>

## Next pages — rules.egress_rules.any / f47ff6a24db3 / 4

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-bae11aa47e185f01d2a43b06c9f2932bd30c1e33130ae7127548fa300608b724"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a81f76db62f0c59a6afc2d420c34f4d9aec72074bbd4a6fef1e79de5e8e7c2b"></a>

## rules.egress_rules.applications — rules.egress_rules.applications / 2e2ab6e96e4d / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.applications

<a id="canonical-ad5939b1d1fdf8219bb09fcda44d522b9f5cc1df69a508989f2b4d965738443c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-c5edd485e8e0f56fe52b0767f616cd79fcee8bb13ebfc56f41e2229cc66d685a"></a>

## Direct properties — rules.egress_rules.applications / 2e2ab6e96e4d / 3

<a id="canonical-9da6ddd9443600cd6f0ae8961064bcd2eea02ddc76e210d4bd10a634ffd20679"></a>

<a id="canonical-cf74e5ea8cfee57d1855c73d4ac71efe8961e77a79a5f64aaa0728eb23ec42ba"></a>

## applications property — rules.egress_rules.applications / 2e2ab6e96e4d / 4

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-af539e6549c0f910a4780c055daf40282d316a0a352f3a6f1bea3eb06201ba1b"></a>

## Next pages — rules.egress_rules.applications / 2e2ab6e96e4d / 5

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-de941d88c7f6fd5dadf7d2948c0647a8215e8bf7fec0e38fa5c2acd565869667"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecc644c70871a1680899e7727044f35d13b010aa2f9819f653edfbb28e1c95ac"></a>

## rules.egress_rules.inside_endpoints — rules.egress_rules.inside_endpoints / 10f7e1d7a9f2 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.inside_endpoints

<a id="canonical-1d7713c3c70c5195c511be334fc3afd3d8b6f91cdc4e188cfa6ddafc21b28a1c"></a>

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
inside_endpoints = {}
```

<a id="canonical-683b73046e8e7f2a69df1005f72c723508b420385ecbc3162b9d0a18e034ae5e"></a>

## Direct properties — rules.egress_rules.inside_endpoints / 10f7e1d7a9f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-81c6dd47a02d1bdffe78cecec04730b62d73fc3e6534a35a7ac5a56555a02480"></a>

## Next pages — rules.egress_rules.inside_endpoints / 10f7e1d7a9f2 / 4

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-1e4c43354fbd01b2193e7d05b626ef6375586ba3f56829bb7f589f12ea935185"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fc37db8c1eb9843669e7703d172ea5acda080de0110766c0606b6892bb06d68"></a>

## rules.egress_rules.ip_prefix_set — rules.egress_rules.ip_prefix_set / 29ad39d24395 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.ip_prefix_set

<a id="canonical-f7786c7dffa2d6de13323762a28ecf3ede90fd787e6519f0595e0e2a37fd6baf"></a>

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

<a id="canonical-46db4389b160cd93ff8cd54f6b3414fbf377183a836da0460d0be1896c71b2dc"></a>

## Direct properties — rules.egress_rules.ip_prefix_set / 29ad39d24395 / 3

- [ref](resources--network_policy--reference--group-001.md#canonical-0694688ea5b147761ad0076b2187ad95a85cb57e88d1673f0579ef946059b6bc): complete subsection reference.

<a id="canonical-b71b9466a53981183a460a0cc2e097ed8ebdb519d1b1701a6f9b22d9f7a052f1"></a>

## Next pages — rules.egress_rules.ip_prefix_set / 29ad39d24395 / 4

- [rules.egress_rules.ip_prefix_set.ref](resources--network_policy--reference--group-001.md#canonical-0694688ea5b147761ad0076b2187ad95a85cb57e88d1673f0579ef946059b6bc)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-0694688ea5b147761ad0076b2187ad95a85cb57e88d1673f0579ef946059b6bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32f51f00b0142ee81566e949bbcc4d358433ffc436f160059b39663eeaf857ac"></a>

## rules.egress_rules.ip_prefix_set.ref — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [rules.egress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-1e4c43354fbd01b2193e7d05b626ef6375586ba3f56829bb7f589f12ea935185)
- rules.egress_rules.ip_prefix_set.ref

<a id="canonical-6e3df8118cd4097f1b8e2cbf21bb7e4af04c8a64cdaa50ae9664c603614bb100"></a>

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

<a id="canonical-835b171c9160e380f38f1becba23d7705d7203ad3bcc70ce38e57bb2fb5c53a6"></a>

## Direct properties — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 3

<a id="canonical-bf142740783632ad5db733abb07b139936f250d0951b6a9e064d579159f7a80e"></a>

<a id="canonical-0c6bcce73fa7e46435f995bc02139eaafe7927eabfe213b62b384ca53a8e3a8c"></a>

## kind property — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 4

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

<a id="canonical-d1f3b32917d70dc89b4ff0f7abb7d947e8104fd5489d28c2361e5c1ab11d911c"></a>

<a id="canonical-a59a4f99fc92a8c06af8d5ccc5817e6217d7a21a9afa4688a6af91ec70c789b1"></a>

## name property — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 5

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

<a id="canonical-4a87d013695bd0528c5ecc59a01e9889db4ed3a0b4fe1a4c03912aa2eed14ccf"></a>

<a id="canonical-c557a0d02fc64568c8a0754302f028d59525d5c285b5c4d69c07581a9d0f175d"></a>

## namespace property — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 6

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

<a id="canonical-2814a87362ceecf039ad640bfc05ef3b1a31fa9a8670f42ef4ef4558d8a892e6"></a>

<a id="canonical-4cd9ca4a39f82216689ce6e0666b537305e254f4533dc9237c4c9224e8a5c198"></a>

## tenant property — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 7

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

<a id="canonical-999b4806205fcb3a619fc48277004107ab0421eed5ce8db58a6ff02c03087865"></a>

<a id="canonical-5fa812d68eebe0c07f33baec7f2aee224803c239d612e4440b7c81bcd6dfda20"></a>

## uid property — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 8

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

<a id="canonical-5e8de0340c6e2053e42aec6016ac6e42cc44ba5ee457637eab112efaee51cf26"></a>

## Next pages — rules.egress_rules.ip_prefix_set.ref / 970cc0ee48cd / 9

- [rules.egress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-1e4c43354fbd01b2193e7d05b626ef6375586ba3f56829bb7f589f12ea935185)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-37cc1b7067ea7fe251cce72615bd16f7553d7cd785f5fedf122052d593e16373"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19302d16b98e7a80b80d0e3cdabf7b530708ee6c40095668b2f5d29335163079"></a>

## rules.egress_rules.label_matcher — rules.egress_rules.label_matcher / bf1ef8102e0e / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.label_matcher

<a id="canonical-182d39d0faa38e669e70068c56502d2d510a92ce67fcc1b75383fb83d53394cb"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-b64c47b997ae07abf87874d741e8a0477c9229b0f8c91e2a0430de2338abc973"></a>

## Direct properties — rules.egress_rules.label_matcher / bf1ef8102e0e / 3

<a id="canonical-76f7c4292b6ce1283853e8eb3f0019caa62068da63be7f8f5992cfd57d0bb8a3"></a>

<a id="canonical-734fb598e5f8760f834d3b9f089736cc2e46becaaf63f55c5c8b40a60ad6a3ca"></a>

## keys property — rules.egress_rules.label_matcher / bf1ef8102e0e / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fc9ba66465cd3674c9247916b6fad0664555318ff14afe63228862916f192806"></a>

## Next pages — rules.egress_rules.label_matcher / bf1ef8102e0e / 5

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-3782f29a45559a7cb6ba46f1df4e5a0f9d82bca638a68232a4dde6cc4cdd84a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c6a39e2c7ed331a46e8d84523bb8eda6b05c10e54cfe7b6f518f5ca9740f791"></a>

## rules.egress_rules.label_selector — rules.egress_rules.label_selector / c9f67ea1a4f8 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.label_selector

<a id="canonical-3e9d427b90035d4301bf1aa9403e552c3bc28d1b121aa5dac1e7ea9d3a82af46"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c8fab3f691ddede59b5adb67330c060753f3c0e2905f9f786b0c6cfe99af07e"></a>

## Direct properties — rules.egress_rules.label_selector / c9f67ea1a4f8 / 3

<a id="canonical-17578fa2a9fecf6c4b6f8ab206381ec2814050f113f596f8716065cfefc7ceb8"></a>

<a id="canonical-79e0fa96fdba1f021718308f5a8298f94c4ad47ef93b246051ca6c8143d16923"></a>

## expressions property — rules.egress_rules.label_selector / c9f67ea1a4f8 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-656fec254800afe57b736a16cd32a9638dc1bbca0fb2420873bafa2613b88414"></a>

## Next pages — rules.egress_rules.label_selector / c9f67ea1a4f8 / 5

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-37d9ae63737878871475550eeee02ecb98e1e637f66cfa01998738eabbcc336d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b897d81425456cf887edebd739d23de47a955b5ef32822783874845ab46382d0"></a>

## rules.egress_rules.metadata — rules.egress_rules.metadata / d4794a9211d0 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.metadata

<a id="canonical-6e344b05851ef80915ae6f1a83780728552f09ea0086c58c1b7a47bfef4724fb"></a>

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

<a id="canonical-7df1e02e73832059c161e8556bd0da8659a341f057b7d7f8c1a13be7eaec89b9"></a>

## Direct properties — rules.egress_rules.metadata / d4794a9211d0 / 3

<a id="canonical-f1b5a4eabb10871178aacda086c503c79401a2ce6bd1297facf2e2929bed05a4"></a>

<a id="canonical-f3a2c1beae1c4a17d727684f0674acdcef35c6f47d25bec1a0b550eeffe85fae"></a>

## description_spec property — rules.egress_rules.metadata / d4794a9211d0 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-d6b9aa170a359437916d945888686f980e441f2293ce05714ffe0a2e4ba4212d"></a>

<a id="canonical-b26f08f02f9c7e6f212034efbaeeac9062257c51be8fa53f0dda2633f80342fe"></a>

## name property — rules.egress_rules.metadata / d4794a9211d0 / 5

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

<a id="canonical-03fd0e1fa5ddd5b1d9605b79c797776281a1fc0f9fc4d01adf2e7f36d235999d"></a>

## Next pages — rules.egress_rules.metadata / d4794a9211d0 / 6

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-94ec9fcc67f80bb987da23ce458b3f719062f4ea086859b8c14022f38eeab4ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73771d36c7a7b9a7273998c23c37886522b2fbafd804a04843d12e0c0b5a4e21"></a>

## rules.egress_rules.outside_endpoints — rules.egress_rules.outside_endpoints / e68cc7fbd10b / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.outside_endpoints

<a id="canonical-c4a42c078cdb717d12077217f18390129dfec8eed6889569332fdff33074c8fa"></a>

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
outside_endpoints = {}
```

<a id="canonical-5b72fe68ac3027e2df26e6cc2af0cd574c24c93af7bf556b42a9fa50f5c44837"></a>

## Direct properties — rules.egress_rules.outside_endpoints / e68cc7fbd10b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-781eb6a2153a04fa1284ecd6bbed68fb3b034651748016a8b349985356a124fa"></a>

## Next pages — rules.egress_rules.outside_endpoints / e68cc7fbd10b / 4

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-94bd8b75ab142c767843b3aaddd7798eee4e99d300753aaf7ec12cb111ef8fd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b86cf4a877f703664395f79933b3dc1858aef0569fa9b70fc41081e7c28f9009"></a>

## rules.egress_rules.prefix_list — rules.egress_rules.prefix_list / dd29306fe5d7 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.prefix_list

<a id="canonical-abaecdf1b073b00bf3e4eede16f888a988602e540b9e744079855e5f99c31d35"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-141fb17f939dda29f329e6c36fdd1ec81c93ae49d3b6f1425ed4fa71828600fd"></a>

## Direct properties — rules.egress_rules.prefix_list / dd29306fe5d7 / 3

<a id="canonical-2833e9610ebf65cfdc5d6957790ecca05b03c1d30eefc7eadd6c0dcddddfe7f0"></a>

<a id="canonical-f07ee7a8b22bf80589c1ef6b284c3fc58e0478ca456d6130bd8d1fbbb9a0c7ff"></a>

## prefixes property — rules.egress_rules.prefix_list / dd29306fe5d7 / 4

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

<a id="canonical-2c97ffa60587e6f12970c7047b81a25fa8579dafd9a3985d85fd5d9f20532044"></a>

## Next pages — rules.egress_rules.prefix_list / dd29306fe5d7 / 5

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-869dddae57d8e65c24edad07a172858bc70b28eff189bcd4ab4a2dd5bbc3145a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ddc59011b7ccd759f63aa502c7a3e348c9498a97998dd2a3dfc8e124ff7168f"></a>

## rules.egress_rules.protocol_port_range — rules.egress_rules.protocol_port_range / 0309e2f15d24 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- rules.egress_rules.protocol_port_range

<a id="canonical-957c534e0e0e3c9139e07314b1d92a0a80d8ca949574f7b204062bbcc9992bef"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-e6b91405e637734064330d1581a32bfdd5072c7e9eb0fcca14159d5003759b92"></a>

## Direct properties — rules.egress_rules.protocol_port_range / 0309e2f15d24 / 3

<a id="canonical-d80cebdbe9d29f622b67c66e2cb87af6314d4c25128bf0c87ca1c796768d233f"></a>

<a id="canonical-5572d9ce484742bb638f973ca77d7c3f78a984df3aceff833a1f01e62f0c61af"></a>

## port_ranges property — rules.egress_rules.protocol_port_range / 0309e2f15d24 / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-554dcb051d1ff850e9b9866fe927eeac590207ff6b6ffb3baa872847101f08bc"></a>

<a id="canonical-f694b49aea56196516e08a488685def0170ae1b3d4d3d742a3de469b4a8dfc02"></a>

## protocol property — rules.egress_rules.protocol_port_range / 0309e2f15d24 / 5

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-4c83e9bdccf57d1b3fb2ae6654d8d1f5f7338f10508c3da41edeb5447dff3156"></a>

## Next pages — rules.egress_rules.protocol_port_range / 0309e2f15d24 / 6

- [rules.egress_rules](resources--network_policy--reference--group-001.md#canonical-78b68c8c4560b85cbfda2703cfad118de8def36542eace3ea93048167c5e23a3)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e9dfc1fc993689eb307a3cc9670ae425ec9bbf1ed2a01af4c168dc9da6ef2f7"></a>

## rules.ingress_rules — rules.ingress_rules / e257bd1fb372 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- rules.ingress_rules

<a id="canonical-d4fba6a7307e25a2a49b3f669ec8a0b9d5e0c946b4363491ed9da2c70e0f3f4b"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections to policy endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
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
ingress_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-88a125d9aeca51a7592df31fd6d5e60a4c04e0533269bae81db3eb7de8b74bf3"></a>

## Direct properties — rules.ingress_rules / e257bd1fb372 / 3

<a id="canonical-09e21ccbd5cbaeec7d4062e71d46e231b4f424fd623f8bac0e763931d41a3f8e"></a>

<a id="canonical-9694c395ba6a722ae6acd5467594670ab5c0421fa4e6517295ae32f32c337aa0"></a>

## action property — rules.ingress_rules / e257bd1fb372 / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

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

- [adv_action](resources--network_policy--reference--group-001.md#canonical-c57cfb2c0ba611fef9312bd9838b1d0ff7cefbed92a7d36172e9b7397ed01dda): complete subsection reference.

- [all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-1a0d94e15c548e2cae46343d8e9d3a38ac662dcbf9e1bee0a597ad1a938ba581): complete subsection reference.

- [all_traffic](resources--network_policy--reference--group-001.md#canonical-71ff44d6a51e541f994149956d9525e12e7a533f037bd24f28debd74a4882bf2): complete subsection reference.

- [all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-771d7c890111653ebf230eac5086656059adfd645445a817bf36c882b2e56ae5): complete subsection reference.

- [any](resources--network_policy--reference--group-001.md#canonical-49da52ba363e4b221c7bcfc1dbd44d3de57d76e5b4caea91430b5bda1c165c4f): complete subsection reference.

- [applications](resources--network_policy--reference--group-001.md#canonical-131c2d6155c68481810caad9a422a17611f284b691f805d21c90d567e6891a66): complete subsection reference.

- [inside_endpoints](resources--network_policy--reference--group-001.md#canonical-1e48e7209637ed33eeffe736f98f499f25f11e4549b64ab6d3c9128a284bfca6): complete subsection reference.

- [ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0bfcd79c501c159c0c8536a50a83c2f46848adb3b439f5c07ff8082d97a1e8ee): complete subsection reference.

- [label_matcher](resources--network_policy--reference--group-001.md#canonical-5a0fbc325705ba6dfa02bf9b07c2093a23f0b2c95fc4b6ba64a783c1eac563b8): complete subsection reference.

- [label_selector](resources--network_policy--reference--group-001.md#canonical-53d1f6d7285c2d2490572819fb3f67ac3f6894e2084720e9a5b988a57d319bd4): complete subsection reference.

- [metadata](resources--network_policy--reference--group-001.md#canonical-613683426ee52bae9ab346e5656172f8f066cbcdbb7067318da6ed14bc771b7d): complete subsection reference.

- [outside_endpoints](resources--network_policy--reference--group-001.md#canonical-1356fe2214963a163f8e1ada555b09e1ced208ebd307be84599542b83710bef1): complete subsection reference.

- [prefix_list](resources--network_policy--reference--group-001.md#canonical-739526e292c109a2736fe052c1baa3fb0f526b8263ad653b5d6cc4baad830f83): complete subsection reference.

- [protocol_port_range](resources--network_policy--reference--group-001.md#canonical-fcbfe10b56a329f2c6d5ac5aa6c6a174e1a7dd8a1f0b436f88403cba4c6fbbf8): complete subsection reference.

<a id="canonical-46f1d845a7ccafaf515ead4060a3bb1894073f1f94ca6f917ee2b0152203ad37"></a>

## Next pages — rules.ingress_rules / e257bd1fb372 / 5

- [rules.ingress_rules.adv_action](resources--network_policy--reference--group-001.md#canonical-c57cfb2c0ba611fef9312bd9838b1d0ff7cefbed92a7d36172e9b7397ed01dda)
- [rules.ingress_rules.all_tcp_traffic](resources--network_policy--reference--group-001.md#canonical-1a0d94e15c548e2cae46343d8e9d3a38ac662dcbf9e1bee0a597ad1a938ba581)
- [rules.ingress_rules.all_traffic](resources--network_policy--reference--group-001.md#canonical-71ff44d6a51e541f994149956d9525e12e7a533f037bd24f28debd74a4882bf2)
- [rules.ingress_rules.all_udp_traffic](resources--network_policy--reference--group-001.md#canonical-771d7c890111653ebf230eac5086656059adfd645445a817bf36c882b2e56ae5)
- [rules.ingress_rules.any](resources--network_policy--reference--group-001.md#canonical-49da52ba363e4b221c7bcfc1dbd44d3de57d76e5b4caea91430b5bda1c165c4f)
- [rules.ingress_rules.applications](resources--network_policy--reference--group-001.md#canonical-131c2d6155c68481810caad9a422a17611f284b691f805d21c90d567e6891a66)
- [rules.ingress_rules.inside_endpoints](resources--network_policy--reference--group-001.md#canonical-1e48e7209637ed33eeffe736f98f499f25f11e4549b64ab6d3c9128a284bfca6)
- [rules.ingress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0bfcd79c501c159c0c8536a50a83c2f46848adb3b439f5c07ff8082d97a1e8ee)
- [rules.ingress_rules.label_matcher](resources--network_policy--reference--group-001.md#canonical-5a0fbc325705ba6dfa02bf9b07c2093a23f0b2c95fc4b6ba64a783c1eac563b8)
- [rules.ingress_rules.label_selector](resources--network_policy--reference--group-001.md#canonical-53d1f6d7285c2d2490572819fb3f67ac3f6894e2084720e9a5b988a57d319bd4)
- [rules.ingress_rules.metadata](resources--network_policy--reference--group-001.md#canonical-613683426ee52bae9ab346e5656172f8f066cbcdbb7067318da6ed14bc771b7d)
- [rules.ingress_rules.outside_endpoints](resources--network_policy--reference--group-001.md#canonical-1356fe2214963a163f8e1ada555b09e1ced208ebd307be84599542b83710bef1)
- [rules.ingress_rules.prefix_list](resources--network_policy--reference--group-001.md#canonical-739526e292c109a2736fe052c1baa3fb0f526b8263ad653b5d6cc4baad830f83)
- [rules.ingress_rules.protocol_port_range](resources--network_policy--reference--group-001.md#canonical-fcbfe10b56a329f2c6d5ac5aa6c6a174e1a7dd8a1f0b436f88403cba4c6fbbf8)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-c57cfb2c0ba611fef9312bd9838b1d0ff7cefbed92a7d36172e9b7397ed01dda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de799765884d5cb6686ef76b0cd45db60a5e8278970185986b28b8706fc0d4e7"></a>

## rules.ingress_rules.adv_action — rules.ingress_rules.adv_action / db44d35dc14b / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.adv_action

<a id="canonical-681daf8927f4ca61e628ab815ae0db5c47cf3e3c461528a23451f35b1d8edb20"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
adv_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-b98d4947cb5563dee1920ccb3fdef4670b7fc543ebac7cabe378b96521fc32c9"></a>

## Direct properties — rules.ingress_rules.adv_action / db44d35dc14b / 3

<a id="canonical-f4966a4ef4e7c05e4407e55b7e5a1da91741cfadc29bf62366f8d7c14ac95104"></a>

<a id="canonical-0e51b82ea9f8c71080830cf3f3caa454a3573f1062a3fcd8ba5addfb51f42864"></a>

## action property — rules.ingress_rules.adv_action / db44d35dc14b / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3f3a800df313ae2efc115d8c38ba648b2e4d04149a3d68bad43443c572435630"></a>

## Next pages — rules.ingress_rules.adv_action / db44d35dc14b / 5

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-1a0d94e15c548e2cae46343d8e9d3a38ac662dcbf9e1bee0a597ad1a938ba581"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faa4ca5e42e4df1d2730d91ebf3435cd9b0f0c660c17c61308911ae0149dae8f"></a>

## rules.ingress_rules.all_tcp_traffic — rules.ingress_rules.all_tcp_traffic / dd2d7494528c / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.all_tcp_traffic

<a id="canonical-b24f161cb3f9e950a86b0d63aa807c6087b57d1be471bbc1edcda93259f8a125"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

<a id="canonical-73d4421ab4306e23a1c91f9e419cc8aec08dfcfacbcc9a1df02a8908bd507440"></a>

## Direct properties — rules.ingress_rules.all_tcp_traffic / dd2d7494528c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7967eb8bde8027a88112fe7a725462de3d06fe1c17e9fdb11d9b82d18572906d"></a>

## Next pages — rules.ingress_rules.all_tcp_traffic / dd2d7494528c / 4

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-71ff44d6a51e541f994149956d9525e12e7a533f037bd24f28debd74a4882bf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9360d7444d1a5610c2d058130dcdcaecf78ee189dd29a02700c96932111de049"></a>

## rules.ingress_rules.all_traffic — rules.ingress_rules.all_traffic / 42028ecf070d / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.all_traffic

<a id="canonical-7654d6e53cd7c6a746531f99b4156f0e0440a9ad58fa7b60459e2aa3332479db"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

<a id="canonical-3c64a586ac3fea0e19aa1ae9a5799a96c89a12c1fb8ccae5b1ea6356abc58d12"></a>

## Direct properties — rules.ingress_rules.all_traffic / 42028ecf070d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-186fc8cc3494098ac299f96b429552cfc7d8938aedda2452779c87b36ff9c695"></a>

## Next pages — rules.ingress_rules.all_traffic / 42028ecf070d / 4

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-771d7c890111653ebf230eac5086656059adfd645445a817bf36c882b2e56ae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6a9664e9e016447ccbc027b4e47e29b6b310e0771663260838231411e874f29"></a>

## rules.ingress_rules.all_udp_traffic — rules.ingress_rules.all_udp_traffic / cbf13f141853 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.all_udp_traffic

<a id="canonical-cd41ff73730c34f3469b268e87701e7fe431ce323ae931a3bf873444c804aa9b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

<a id="canonical-0f896e0e7d184fb7a3c0663f3ef7c360fdc800f4e6daf2facc5174c532d9e7cd"></a>

## Direct properties — rules.ingress_rules.all_udp_traffic / cbf13f141853 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-074a1a125168873240b0c2d7d9db7ea5173eeedbb21fc82b3fb115b5834a4c5b"></a>

## Next pages — rules.ingress_rules.all_udp_traffic / cbf13f141853 / 4

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-49da52ba363e4b221c7bcfc1dbd44d3de57d76e5b4caea91430b5bda1c165c4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02a9f60831369b9e745d30180bbac945bfcb5e20a863e69a363f2e96b4471524"></a>

## rules.ingress_rules.any — rules.ingress_rules.any / 181bd2a98d92 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.any

<a id="canonical-9a876fcdd011c13a1396f4fa89085507971bd113912c4578934e83d0d6e7bec0"></a>

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

<a id="canonical-8c4a76e07af307c8a54ccb1cadb06f3aa18860440f8ee7a7b3d0ee124b79e112"></a>

## Direct properties — rules.ingress_rules.any / 181bd2a98d92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1896f98f71a794df09627c2950aa8246d5d929e9b2d80f37d072be6e33cb739c"></a>

## Next pages — rules.ingress_rules.any / 181bd2a98d92 / 4

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-131c2d6155c68481810caad9a422a17611f284b691f805d21c90d567e6891a66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4e5891d4e8943dd474975d55272fa44836b2a9f3293122721bd59404067e27a"></a>

## rules.ingress_rules.applications — rules.ingress_rules.applications / 848ff2dd0a95 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.applications

<a id="canonical-f79523d68e7e9676cd6711f4046de251b7b0ac2456c78c08d4a26c3202b8f384"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-f52be6bfbc3cf5740563c0d44452aaa0115a53bac0049fcb2af1c502e597dbed"></a>

## Direct properties — rules.ingress_rules.applications / 848ff2dd0a95 / 3

<a id="canonical-f039c72a36f9f3a0fead14834264b265e80bf71a90b0dfc44d1d0b1f0f30fdc0"></a>

<a id="canonical-ef45975218521323dbbcab8281341ffff6224111e35c8cdba8c5736c9d089f64"></a>

## applications property — rules.ingress_rules.applications / 848ff2dd0a95 / 4

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-c801d9eb7ddece173ffcf26f2e8fb45167937cb2d82f737a262cf3da2e4da88a"></a>

## Next pages — rules.ingress_rules.applications / 848ff2dd0a95 / 5

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-1e48e7209637ed33eeffe736f98f499f25f11e4549b64ab6d3c9128a284bfca6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-131b7ae34ead3e57a5baeaa6bd89391d552879e444b6bdeca0960775e65b4433"></a>

## rules.ingress_rules.inside_endpoints — rules.ingress_rules.inside_endpoints / 2eafadcb83a9 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.inside_endpoints

<a id="canonical-23bde8769bff6921112d0e8406b39c8fb692d99b80ffc533e62a0ede90fbff8b"></a>

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
inside_endpoints = {}
```

<a id="canonical-f7e1685e14a0d71be26035d8e39644417bc9a1a645cd0beef00f3aafd69bb419"></a>

## Direct properties — rules.ingress_rules.inside_endpoints / 2eafadcb83a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64575f2c64fda5f75e523c9a9e943e855ef351be8cb99918e7f542183055e599"></a>

## Next pages — rules.ingress_rules.inside_endpoints / 2eafadcb83a9 / 4

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-0bfcd79c501c159c0c8536a50a83c2f46848adb3b439f5c07ff8082d97a1e8ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1842285f5467123053d3745355eddbcba6eba921595b8ce1cfa08e9f06ccabe8"></a>

## rules.ingress_rules.ip_prefix_set — rules.ingress_rules.ip_prefix_set / dd2a62f3a3d0 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.ip_prefix_set

<a id="canonical-aa5b637af7da336abba8b7e191afb53da274a861c996dad6143b26f50d2f5dec"></a>

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

<a id="canonical-eb92bfb4cd7fb09d633bf92f51a1d24353235e6ece05b55e812a1d8c6fbfba9e"></a>

## Direct properties — rules.ingress_rules.ip_prefix_set / dd2a62f3a3d0 / 3

- [ref](resources--network_policy--reference--group-001.md#canonical-fc57aa8512918b975da3aaa7e17537cab8c1cb6b0dd191482912dc44b49df5aa): complete subsection reference.

<a id="canonical-33057576daa3a271ceed0371ebe5625a4e591b946a969249da955b4322694a35"></a>

## Next pages — rules.ingress_rules.ip_prefix_set / dd2a62f3a3d0 / 4

- [rules.ingress_rules.ip_prefix_set.ref](resources--network_policy--reference--group-001.md#canonical-fc57aa8512918b975da3aaa7e17537cab8c1cb6b0dd191482912dc44b49df5aa)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-fc57aa8512918b975da3aaa7e17537cab8c1cb6b0dd191482912dc44b49df5aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fd426609d079f2007d2a2849f7f4d360063e4ab354da3b4eeb292946e0533d3"></a>

## rules.ingress_rules.ip_prefix_set.ref — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [rules.ingress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0bfcd79c501c159c0c8536a50a83c2f46848adb3b439f5c07ff8082d97a1e8ee)
- rules.ingress_rules.ip_prefix_set.ref

<a id="canonical-ea314b9e3aa12e86dc3aa3889f45c8acc014ca481ce4cdc07933b55c12963879"></a>

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

<a id="canonical-7a3fa9ee51df4803b8f25dfcead28046cb56a80d4df903ad7fc882ed72982a39"></a>

## Direct properties — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 3

<a id="canonical-1bbecedd0d5850c5e0c4ef8aaf02fc71c68976f1e25ed60884784454ab618e44"></a>

<a id="canonical-97ff2f0f4e64c57811777f497505aef278645bd56b3771dee910d5cda10e08f1"></a>

## kind property — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 4

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

<a id="canonical-f57d213fe8ef18a4bc371f2988b8d201f3158ff89d5c5d1f0b886baf9c8d405f"></a>

<a id="canonical-63095ed7f000cf77d04e4a881787e6ddbc23c22b1f38d70f177e37dd1559c39e"></a>

## name property — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 5

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

<a id="canonical-1971d77f1e80eea5d8a7142e022aa6965ae23a93eb4f4d377d9a7753a738a5f1"></a>

<a id="canonical-10b017994ac4e94f5b80f7b7f6ffaf094499f031895dea16e86940b1a5f4ffab"></a>

## namespace property — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 6

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

<a id="canonical-655bb79b72af35a834ecd88d868bb8585b0518e35ff24cb9ac3a0d837a8c8767"></a>

<a id="canonical-525df2b46dc3d4dea24231505797d88052de0571301a63f395f5437036b6ace9"></a>

## tenant property — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 7

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

<a id="canonical-0935a943df2d9bcee3e513b938983292fa7c6976571dc4c9792ec65d59c1b315"></a>

<a id="canonical-b397972301d72ecf08c22d9bfc8be32d1e85170be0fdd2ea3f763877511a9730"></a>

## uid property — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 8

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

<a id="canonical-b0afd4826bcc944a8c5785c345c32fdb6c468a47869e652da169bff98964d310"></a>

## Next pages — rules.ingress_rules.ip_prefix_set.ref / 002224810cf5 / 9

- [rules.ingress_rules.ip_prefix_set](resources--network_policy--reference--group-001.md#canonical-0bfcd79c501c159c0c8536a50a83c2f46848adb3b439f5c07ff8082d97a1e8ee)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-5a0fbc325705ba6dfa02bf9b07c2093a23f0b2c95fc4b6ba64a783c1eac563b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d33b0af66ba07e3c422535f4b9facebf8540f495154cac57a34a110db94d307"></a>

## rules.ingress_rules.label_matcher — rules.ingress_rules.label_matcher / f06b62913790 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.label_matcher

<a id="canonical-7f71ce56846e007d1880d338a6fc14cbb3c551bd267a103f5cd9f244a2cacc46"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-291d9cb451146744e627d29ddcb16734102c8d0db1941484ae970b7f0ff6a644"></a>

## Direct properties — rules.ingress_rules.label_matcher / f06b62913790 / 3

<a id="canonical-9723a1904814d4502d1e2202773b628826d29b51b8e172477b98e561205cac5f"></a>

<a id="canonical-a716f85a80d8b9352dc95d819f2565120b1cd3c68066c40239e08672ee038a7c"></a>

## keys property — rules.ingress_rules.label_matcher / f06b62913790 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4d592f639f5cc080f1fdbc6ff7885600bf4dbb35e888cf3eb22abb9916d91f96"></a>

## Next pages — rules.ingress_rules.label_matcher / f06b62913790 / 5

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-53d1f6d7285c2d2490572819fb3f67ac3f6894e2084720e9a5b988a57d319bd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e645e3901f2b3340c7b7b436e2d4fe77765f33f77c79d13c5fa7b382b5204e7f"></a>

## rules.ingress_rules.label_selector — rules.ingress_rules.label_selector / 3e19940002b9 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.label_selector

<a id="canonical-ea2a315a1b458f88aff51105e297b872db463c14f7aee0e06c765e4837f3c1df"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-cedbc279775c5fa1ee230fc6bfd67ed233c52a4e01186c17c9960810b5f736c0"></a>

## Direct properties — rules.ingress_rules.label_selector / 3e19940002b9 / 3

<a id="canonical-706965ec4afb407cbf7f029e72c98cf5b8a7d38834177e81247a0e7641726555"></a>

<a id="canonical-58c7b4e8fd5c3d06ef73a45877521d83a2ea8363bd7ce24555f8d0e313005145"></a>

## expressions property — rules.ingress_rules.label_selector / 3e19940002b9 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-10f85081c85c5a27ac99f701d23cc49454df6950838cf95dedb778a5d0f19514"></a>

## Next pages — rules.ingress_rules.label_selector / 3e19940002b9 / 5

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-613683426ee52bae9ab346e5656172f8f066cbcdbb7067318da6ed14bc771b7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9aa1337f7bd40a709df44119da6b633d73c943aaf36a970b21b9611045c7afb1"></a>

## rules.ingress_rules.metadata — rules.ingress_rules.metadata / e2079d3c354b / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.metadata

<a id="canonical-33fb34d2c13ba66bb8c8bb043041fb9d48e9564ab0651fb626004a03d3dc737a"></a>

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

<a id="canonical-162705fe9aa512fbb30d4b2aa43ebc2f65792a9f57fc058f22b74ee0baa39716"></a>

## Direct properties — rules.ingress_rules.metadata / e2079d3c354b / 3

<a id="canonical-f346c9bcbaa44f80ce88c9f10a2a9be6ed89964453fbdd8a4264f370af66d996"></a>

<a id="canonical-d538e79cdf32f5e0bd17d2f4d969af8bac33af958ebfdca412f25e14b8c4d10b"></a>

## description_spec property — rules.ingress_rules.metadata / e2079d3c354b / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-7d59448d8a8cf850dbadea0c72085c5701bc81985d7f466aa2e49ee3761f8738"></a>

<a id="canonical-8784106534be3cd32b26928ad86f417ad7b338aaf4eb07a6b107bc279ce7aa72"></a>

## name property — rules.ingress_rules.metadata / e2079d3c354b / 5

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

<a id="canonical-5e7cdb22076fd83c69d1462e9bd661242e4dfb1161c6f1bd43119af905f96397"></a>

## Next pages — rules.ingress_rules.metadata / e2079d3c354b / 6

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-1356fe2214963a163f8e1ada555b09e1ced208ebd307be84599542b83710bef1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72ce5eeb874720c1a0a1f2b84fcf2e60f0035d715f0ae2b022b57de843a6a020"></a>

## rules.ingress_rules.outside_endpoints — rules.ingress_rules.outside_endpoints / e8507a2024c3 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.outside_endpoints

<a id="canonical-a6b219a5653ef90ae48934dabf4d3d58798cb235e98ecd02eb876a791b51b4fa"></a>

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
outside_endpoints = {}
```

<a id="canonical-00af937ad30ff7fb7dcc0afc2a06782bc5804fad96e8747b2d09f382f577b5fd"></a>

## Direct properties — rules.ingress_rules.outside_endpoints / e8507a2024c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6718473ac785b427ba531c10f48041f3b437f11f613a78aaf8ccddce510ff578"></a>

## Next pages — rules.ingress_rules.outside_endpoints / e8507a2024c3 / 4

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-739526e292c109a2736fe052c1baa3fb0f526b8263ad653b5d6cc4baad830f83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-551caeaa6150eb65e18d88ae8c555a3bfb4c527356f670f03181a75106ea9320"></a>

## rules.ingress_rules.prefix_list — rules.ingress_rules.prefix_list / 70ee13578384 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.prefix_list

<a id="canonical-862911aeb304d2ff8029f956cf09a9e6955f87ee8cfa71bdf1453815f88a4cbc"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a11db6baebbe52bfaf9565b1aa132175ff5a8092c1bdcfa42338002ac5df465"></a>

## Direct properties — rules.ingress_rules.prefix_list / 70ee13578384 / 3

<a id="canonical-167109372e9d78a8c982985b8e018dd5df273025c254cd33e6e6c5ba15fecc9f"></a>

<a id="canonical-dfc7048f47fa8d2acfb75a03939f67225b28dfc1423d3ec70c06d156edbd3a82"></a>

## prefixes property — rules.ingress_rules.prefix_list / 70ee13578384 / 4

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

<a id="canonical-a17ddebd1903384c6fd659dcc70d87cab025f77227328b46e530c982dcc2baa0"></a>

## Next pages — rules.ingress_rules.prefix_list / 70ee13578384 / 5

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-fcbfe10b56a329f2c6d5ac5aa6c6a174e1a7dd8a1f0b436f88403cba4c6fbbf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daa0056f42feacac1c67cbdeec0466098fc04ebf2115b5da50c2f54ad42b192c"></a>

## rules.ingress_rules.protocol_port_range — rules.ingress_rules.protocol_port_range / ed59e120833b / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [rules](resources--network_policy--reference--group-001.md#canonical-945f39bc276ad2c5f65684f85b091f9c69e0ce8649381629aaa58fc2995cffb1)
- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- rules.ingress_rules.protocol_port_range

<a id="canonical-0868bd0390f1852d13c9d64fdca0958704fcac3a8155ed0da6a7c69f97f6627d"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-fb0201fd2365ed0a688fe69a0208a7a906c3e58a3ecf1caff4b95408f1c43027"></a>

## Direct properties — rules.ingress_rules.protocol_port_range / ed59e120833b / 3

<a id="canonical-7966f8cecba51f55ab03212c80ccc8a3055de13dc8a10e26b5fc64570bff38eb"></a>

<a id="canonical-db00872c75740b44eeeb21388dd31054fde11a2c47c66505799b84297f79b38e"></a>

## port_ranges property — rules.ingress_rules.protocol_port_range / ed59e120833b / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-1522f5c7957b612fd96d86aedeb49a209e0593d2f1b140fb5cd992450742024c"></a>

<a id="canonical-2a4ad84d0649ae9588bf8b3beddf9437bd55aa23c5ba6a491df094e59d4b69cc"></a>

## protocol property — rules.ingress_rules.protocol_port_range / ed59e120833b / 5

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-7561b32a3f8c4e55aadb144f02e3e01154e2d657f7440e72c7a093866dfeb97d"></a>

## Next pages — rules.ingress_rules.protocol_port_range / ed59e120833b / 6

- [rules.ingress_rules](resources--network_policy--reference--group-001.md#canonical-0c26ddd39799b322eb1614346bab94ce710061d6e66af086e3d42321939de5df)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-1b1756d0c2a28e1e88cc694220b1e946446c2046746e8e3e176999919bc2aeb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e80c0cd98e7a916a6da02c3b6f72ff5d6f717e6ad886a42539fe4d5fd5d4f54d"></a>

## timeouts — timeouts / 227fdda82355 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- timeouts

<a id="canonical-5d6714412be8be94de068037b30f1316147ea379858c8543c79cd963182d4865"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e43eacc5256747764d7577e211a3b60523a9e4c0e701cc0daaa5f51940cafa62"></a>

## Direct properties — timeouts / 227fdda82355 / 3

<a id="canonical-f1c0b2ddabb31777a523e0023df63262039e697aa5f8d1174358a9e3f35f50cc"></a>

<a id="canonical-492d0748cee31d87f106a698566c4cd5eeffd4c457e14980c3b2b061daedfbc4"></a>

## create property — timeouts / 227fdda82355 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3da976b1e3ee598e7fa45f590a98e60bd998f86649c20586353d02ae4b3cc1e4"></a>

<a id="canonical-d1a368ccd6bbb2d0cebdbcb4011d984f479f7d73735b2e02600253672f3021a8"></a>

## delete property — timeouts / 227fdda82355 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-79527f77021a17323d4dc6968e26e495ff527f7b31367f222fe4f64dbc108301"></a>

<a id="canonical-c870d021d15024fadc15a06cadc5ae0dbdfc360099d9f9029e3e3746b5f5eb82"></a>

## read property — timeouts / 227fdda82355 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-6893e67ea8b9bbb82e63d809d45c68ff890b9a4e93bc6af1cb23fbb4ca8b75e3"></a>

<a id="canonical-7f7371e9be04e5c22f7aa7f063a865436b12759a9ba327218045f2be80da6fbd"></a>

## update property — timeouts / 227fdda82355 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9734eeb8b793eeac07ebd2485527a545e6f5a9222992bb28432b1a0992ffefad"></a>

## Next pages — timeouts / 227fdda82355 / 8

- [Property reference](resources--network_policy--reference--group-001.md#canonical-5af3803b7443432c13c3a16926914991b434f7699a5f87eaec9a797ca463d013)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
