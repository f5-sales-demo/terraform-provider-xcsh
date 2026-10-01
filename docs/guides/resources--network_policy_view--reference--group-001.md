---
page_title: "xcsh_network_policy_view reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view reference."
---

# xcsh_network_policy_view reference

<a id="canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df381ba3498429ea846944a66e0ce97f036b815b36a4e22d41562c46ca06583a"></a>

## Property reference — Property reference / 4d85bc970357 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- Property reference

<a id="canonical-f28caf819506a041f693cea0811e2bb42a97aa2fdedebe72e5b6f8594b81a5a4"></a>

## Direct properties — Property reference / 4d85bc970357 / 3

<a id="canonical-37e6dcb9e02e206c5ed40de7e1fe6b17b1c029a2980ac8394a463cd3fc4c286c"></a>

<a id="canonical-757cc2c0e8746597d30903044c42b092ce9bba940c16edeb0b82ed8f1a46a3ab"></a>

## annotations property — Property reference / 4d85bc970357 / 4

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

<a id="canonical-bc36b27ded6cb36cfa9a00f61f225847e9313765c2122f8585c0bd59e7d0bc62"></a>

<a id="canonical-b742791b260335692989851d7102e6e6db4755a07b10cd1eb31874d0365644ff"></a>

## description property — Property reference / 4d85bc970357 / 5

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

<a id="canonical-e225f7a7c2163ac6d0dffc31ce98df3967548878073ce33e3023dd3a1934b2c6"></a>

<a id="canonical-a12a88fdb8931783fe654850b88e2440bb34b5a7f1968cf017ef30907694c303"></a>

## disable property — Property reference / 4d85bc970357 / 6

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

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27): complete subsection reference.

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b): complete subsection reference.

<a id="canonical-cd305920f8f8ca5ad9ef1458e38344971536c9fbb00db3af4ea7a20e8adbbee6"></a>

<a id="canonical-6b654c25db19861e5be278c3e124cf543dd5b59e5dc5144de7df4cff2fedade9"></a>

## id property — Property reference / 4d85bc970357 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007): complete subsection reference.

<a id="canonical-130ad6ff9b65bbda857a7d69f1aea237ff3af39fe52cfdf21f2039faa10004e2"></a>

<a id="canonical-0002bf89a95a34a4d272472c4f1a6fcfcad3fed31d72f86e3f15932abcfc8c84"></a>

## labels property — Property reference / 4d85bc970357 / 8

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

<a id="canonical-0352238be4368b05e3c6a5e6dbe83780de70057a0ecd7108579896c505725d42"></a>

<a id="canonical-a064c4e97cace4394f23462b4e6d1eec9da183d91c7993802f9715860d9addde"></a>

## name property — Property reference / 4d85bc970357 / 9

Type: `"string"`. Required.

Name of the Network Policy View. Must be unique within the namespace.

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

<a id="canonical-9330a74063b3a46d7edbc099b1b5900a7857bdbb3dfb2b7af05ba24dfca7c3cb"></a>

<a id="canonical-1eb23db9a75525ae7ae347e752412036bd96306ea53f318399c08152a5dae9ce"></a>

## namespace property — Property reference / 4d85bc970357 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Network Policy View. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

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

- [timeouts](resources--network_policy_view--reference--group-001.md#canonical-8d062b5a6c721be625f15b9aaf1708cf895efbca37bfbcf04384388757be2460): complete subsection reference.

<a id="canonical-e6adfa84171a3130c599e7adecf5eb97459787135eb2194784304f7a0993385b"></a>

## All schema paths — Property reference / 4d85bc970357 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_policy_view--reference--group-001.md#canonical-37e6dcb9e02e206c5ed40de7e1fe6b17b1c029a2980ac8394a463cd3fc4c286c) |
| `description` | [description](resources--network_policy_view--reference--group-001.md#canonical-bc36b27ded6cb36cfa9a00f61f225847e9313765c2122f8585c0bd59e7d0bc62) |
| `disable` | [disable](resources--network_policy_view--reference--group-001.md#canonical-e225f7a7c2163ac6d0dffc31ce98df3967548878073ce33e3023dd3a1934b2c6) |
| `egress_rules` | [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1910b2d9ab0f5bfe20edde4559b050f8af2e0f46e281197e331b4f9a3ab2e57a) |
| `egress_rules.action` | [egress_rules.action](resources--network_policy_view--reference--group-001.md#canonical-e1bcc8ec680c6c02146f3b150d57f9070029d4774f3d6b7d1647e6ee6e8a023f) |
| `egress_rules.adv_action` | [egress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-db053e8c6b54fd20220e4d18137c3fc51b4d239de3bb31b58f7b4a8c6392a9eb) |
| `egress_rules.adv_action.action` | [egress_rules.adv_action.action](resources--network_policy_view--reference--group-001.md#canonical-73a6a0e25b181f68876f90388dfb996352ce4bc59f00540a03301587d81a83eb) |
| `egress_rules.all_tcp_traffic` | [egress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-cd4249e887e7b26c09c8096c752dd43ac9e99bda20d57550a5c061d1c155d705) |
| `egress_rules.all_traffic` | [egress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-ed8ca84bd74cbf1fb061b64547e88a65046e14f8719833189bd6dee9e33b7ec7) |
| `egress_rules.all_udp_traffic` | [egress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-58d3b7ff1793e80fdbc3050934861b15f7a1fb3337c6c15a2754f72b2b160c11) |
| `egress_rules.any` | [egress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-a281ecdeb7d3d7912430084881fd26d3283e143e7d599d8690636d3767dce044) |
| `egress_rules.applications` | [egress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-94ee97aff33ecc9e6217688e8e5c2e370a1fa2303b429819792d30752015df6a) |
| `egress_rules.applications.applications` | [egress_rules.applications.applications](resources--network_policy_view--reference--group-001.md#canonical-e0ecc95ae2720bdd68a5f270509c9fa80814dec6e52e933a5b71fce0140e0310) |
| `egress_rules.inside_endpoints` | [egress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-a5930803adf3abab2e3dfd38405f5fe47a8be9e0d0f82922d98749376f817fd7) |
| `egress_rules.ip_prefix_set` | [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-ba347dd1d85782c39d533bd418833d3280d8fd7adb2bd540a8ade8ec05db5ed1) |
| `egress_rules.ip_prefix_set.ref` | [egress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-63e5e78f449f8f0b34635e0e5aff0b3a05bd07bd001248b224ca287de3b46297) |
| `egress_rules.ip_prefix_set.ref.kind` | [egress_rules.ip_prefix_set.ref.kind](resources--network_policy_view--reference--group-001.md#canonical-eafdf82f4ba482d81ac6de8758badade9393f01859e452779416a552f31bde46) |
| `egress_rules.ip_prefix_set.ref.name` | [egress_rules.ip_prefix_set.ref.name](resources--network_policy_view--reference--group-001.md#canonical-126ba41616d33c4c57885eee623ea55df501116dfbc01a909bb3f957bd2dfc60) |
| `egress_rules.ip_prefix_set.ref.namespace` | [egress_rules.ip_prefix_set.ref.namespace](resources--network_policy_view--reference--group-001.md#canonical-9f73a42e876f3afdeb8e3eb9c415c2e5a8c1057de782aa1c02624f65d3423174) |
| `egress_rules.ip_prefix_set.ref.tenant` | [egress_rules.ip_prefix_set.ref.tenant](resources--network_policy_view--reference--group-001.md#canonical-d9c29b613e867576302026554c05826208aef074fc16495904f6e7729c19e5f2) |
| `egress_rules.ip_prefix_set.ref.uid` | [egress_rules.ip_prefix_set.ref.uid](resources--network_policy_view--reference--group-001.md#canonical-a49acfb0e9f9e51a6c4d9bf28fffa9ffe3d82818740067354f52676a33a5a79b) |
| `egress_rules.label_matcher` | [egress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-8d9fb7454ea665515bd4551c15753ac7af38ba2766d4c0692416a07d8a8a1840) |
| `egress_rules.label_matcher.keys` | [egress_rules.label_matcher.keys](resources--network_policy_view--reference--group-001.md#canonical-596a7eaf5daeda8f810d5a5f312dde3010f016271694fd31b288bdda7efe1e6d) |
| `egress_rules.label_selector` | [egress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-f7e3d694d0c129a71794f4ab0efce0dfcaf14e063b6aebdee9cebcf9c620054d) |
| `egress_rules.label_selector.expressions` | [egress_rules.label_selector.expressions](resources--network_policy_view--reference--group-001.md#canonical-cb428613113756554976173bee20c2fc778dd90f22556ea15a1debb05a5811d4) |
| `egress_rules.metadata` | [egress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-5ffe43f827debd8ad07d882075853c2b51427d139aee573cdfb5fe571926e635) |
| `egress_rules.metadata.description_spec` | [egress_rules.metadata.description_spec](resources--network_policy_view--reference--group-001.md#canonical-902f137a7f9cff7f2b4af5318f7dfc4a7b119570e3a0e6df108118465492413f) |
| `egress_rules.metadata.name` | [egress_rules.metadata.name](resources--network_policy_view--reference--group-001.md#canonical-42d9d79e492dd7480856a8ec70e391ad6250ea566ad4bf426d22fd491ffe0441) |
| `egress_rules.outside_endpoints` | [egress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-1d8f2ea915da060aa15519312a1c39fd76948d2a14e2fd4a6580e8c0d593701c) |
| `egress_rules.prefix_list` | [egress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-d6f06794015c43bd7e1140bc00bb872ccde5583bbf65bf641a2f9e9f26066b3d) |
| `egress_rules.prefix_list.prefixes` | [egress_rules.prefix_list.prefixes](resources--network_policy_view--reference--group-001.md#canonical-1ac6e443c61f6c855ba1f3f1b40f32aa92e614f2c41e9ab4c9f5960857bbdf04) |
| `egress_rules.protocol_port_range` | [egress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-a4db90f9f47af2666006173af29c52df06aab4b9f981df51b6dfb6fecf92fcb6) |
| `egress_rules.protocol_port_range.port_ranges` | [egress_rules.protocol_port_range.port_ranges](resources--network_policy_view--reference--group-001.md#canonical-a8e9277f5e90a9e49de65a9f275fadfeaf46029271e5b20451102636c5e3afdf) |
| `egress_rules.protocol_port_range.protocol` | [egress_rules.protocol_port_range.protocol](resources--network_policy_view--reference--group-001.md#canonical-7f3cc03d479bdeaaec9c930c9eb06952d8a9461580e83766cf0fa99b62d7dfa6) |
| `endpoint` | [endpoint](resources--network_policy_view--reference--group-001.md#canonical-014d0d5b4a08804e205b3c2411491a04ca41f51a182760b5a7322ffdc711c010) |
| `endpoint.any` | [endpoint.any](resources--network_policy_view--reference--group-001.md#canonical-0524977894d36d8c062b34e6c435725e391636f173053e3dba1208348080e5c6) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-b817690e8e159602ccd2b16d891cd460c76454124e2bf463fe64faf01e682565) |
| `endpoint.label_selector` | [endpoint.label_selector](resources--network_policy_view--reference--group-001.md#canonical-7db05c617ae49f9404dbe582de12d123cd28ce20d8961b79f912f7ae9bf64390) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](resources--network_policy_view--reference--group-001.md#canonical-13641802134796117f0a480c3a17ca171cf5c9630268c2bf0b2fee47109d371a) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-c7ffa3f7eb582897cc5988824f722ba3363aec7a5a70d0707b1f95a8facfaf3a) |
| `endpoint.prefix_list` | [endpoint.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-be47c7a58aae4b12e1af924732177b6e9dc8d80b385350f09c9febce62f24e1c) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](resources--network_policy_view--reference--group-001.md#canonical-f8e37f8a3177c4a395f8c536f780640be8ab6f47f07f58ce463d85522305b265) |
| `id` | [id](resources--network_policy_view--reference--group-001.md#canonical-cd305920f8f8ca5ad9ef1458e38344971536c9fbb00db3af4ea7a20e8adbbee6) |
| `ingress_rules` | [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-c54f2c2a3f6e2bf5f9af039a96fa4453f45dd742831fb0ddbe81ee8d37464dd8) |
| `ingress_rules.action` | [ingress_rules.action](resources--network_policy_view--reference--group-001.md#canonical-b0bd2122d408db20ff4c56b5d305bdd84be1d94eb987f74b467f68738e1ca933) |
| `ingress_rules.adv_action` | [ingress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-a4b2a8600cd7ef5d54249816835243ff1fa2148860ecc5d0565bf00297a12e30) |
| `ingress_rules.adv_action.action` | [ingress_rules.adv_action.action](resources--network_policy_view--reference--group-001.md#canonical-92e6d1e1293a765e3916b0099415343cd14b7d8cd3ee220156720d5584b43b5f) |
| `ingress_rules.all_tcp_traffic` | [ingress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-5a172ff953db1ecf333d4e66d17f07393832c3a3ee46d07679a693a4e9616e1e) |
| `ingress_rules.all_traffic` | [ingress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-16c4f4c3639d1989a15e44b591f50907084fc8c0911089c11078545e1ffcb2fd) |
| `ingress_rules.all_udp_traffic` | [ingress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-00cb29625c8211969ab21cc5b9c13342ba03da0673f65c551bca31d37abe8a35) |
| `ingress_rules.any` | [ingress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-e68825c0b22208f8a31329c3e386bfff01a2125a8090b006c2d656f24bc5be24) |
| `ingress_rules.applications` | [ingress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-85adc2c23a613c1543bfea4c53359394f8fd4a7777c74b844574cbba003667b8) |
| `ingress_rules.applications.applications` | [ingress_rules.applications.applications](resources--network_policy_view--reference--group-001.md#canonical-1a0912e2c0e3029c7b01d2fa6aef5ba5fef2965d0589d9047f8335077f931070) |
| `ingress_rules.inside_endpoints` | [ingress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0fa6b1de9cd8da6a36e60842a7c267961500fcd84989fd5546816a24af1a9a63) |
| `ingress_rules.ip_prefix_set` | [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-3f4dbc82f485afb1e6df415af99cdf2d833cc312878bcf4de2b865e621cfde7e) |
| `ingress_rules.ip_prefix_set.ref` | [ingress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-e04d2e19ceb1bce852ac59a664ba0e30be74d35af0e55921a9f58ccaed57e8b2) |
| `ingress_rules.ip_prefix_set.ref.kind` | [ingress_rules.ip_prefix_set.ref.kind](resources--network_policy_view--reference--group-001.md#canonical-76a2d3cd62fceb7420caeac109115b098ff68e4ebe515898901db4cf66f2ae3e) |
| `ingress_rules.ip_prefix_set.ref.name` | [ingress_rules.ip_prefix_set.ref.name](resources--network_policy_view--reference--group-001.md#canonical-d6fc4c50e0fc61af2838f9a9a6f3492f234dd24c1d120e7fa49536ab454b06d0) |
| `ingress_rules.ip_prefix_set.ref.namespace` | [ingress_rules.ip_prefix_set.ref.namespace](resources--network_policy_view--reference--group-001.md#canonical-7e480e959c13ac31da91bfacb61dcce62f9267f015a76d65afdffa69403c5375) |
| `ingress_rules.ip_prefix_set.ref.tenant` | [ingress_rules.ip_prefix_set.ref.tenant](resources--network_policy_view--reference--group-001.md#canonical-b5a8d6300112726ba1afd6d55576fa31d0493f2dd8b5c7f977e94be7c7255495) |
| `ingress_rules.ip_prefix_set.ref.uid` | [ingress_rules.ip_prefix_set.ref.uid](resources--network_policy_view--reference--group-001.md#canonical-c8c47a90c590cf993cf3b01f97a90775d0e3e1dde1ed5c39e02558a259e6f7ee) |
| `ingress_rules.label_matcher` | [ingress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-d0ae1859256be6924ba797f292046f861bca9dd32cf76ed53dabdc61bb94beaf) |
| `ingress_rules.label_matcher.keys` | [ingress_rules.label_matcher.keys](resources--network_policy_view--reference--group-001.md#canonical-e9dfb84e607552236a51fbfdd3ad7bbbdaac83addc506214edb61458319255da) |
| `ingress_rules.label_selector` | [ingress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-f4a2c1ce4d9aed378e9db5b724dc4fb8f329e91993c5dc5a5ed292ed4159fa6d) |
| `ingress_rules.label_selector.expressions` | [ingress_rules.label_selector.expressions](resources--network_policy_view--reference--group-001.md#canonical-7d0197f5bc5b0ad567b2d0f966bf53341d4bdf6f150804bae770b692cfc2b50d) |
| `ingress_rules.metadata` | [ingress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-de84a1449a98dff10a71b369be888a340ac35df695d6643e6888dca1333cbc8b) |
| `ingress_rules.metadata.description_spec` | [ingress_rules.metadata.description_spec](resources--network_policy_view--reference--group-001.md#canonical-4c433b435236c452e24e87c5c1427bb27f769dce942bdd4904b1f9ad45558ecd) |
| `ingress_rules.metadata.name` | [ingress_rules.metadata.name](resources--network_policy_view--reference--group-001.md#canonical-16b19a5757ce14dc42da6fea3a510376651e4600f946b00043c6aee8348f895c) |
| `ingress_rules.outside_endpoints` | [ingress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-62aacd058418d151aee2ef8d179ef2f7b482ecd98ed25339125a39b9e562ea51) |
| `ingress_rules.prefix_list` | [ingress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-cbe40be3c94b4ba3e0013654402f46150440b69be19034d242e96ba52e990580) |
| `ingress_rules.prefix_list.prefixes` | [ingress_rules.prefix_list.prefixes](resources--network_policy_view--reference--group-001.md#canonical-95eef45b83aa6f4aef7ba5d6418848675fe2e25031b658c8a5d567cd1f893f4c) |
| `ingress_rules.protocol_port_range` | [ingress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-090464e0c2ac81f4f72a9b79f1d9b530c5f8db58e141061901c1fee2a1f89c47) |
| `ingress_rules.protocol_port_range.port_ranges` | [ingress_rules.protocol_port_range.port_ranges](resources--network_policy_view--reference--group-001.md#canonical-6a917827b78e9e2adac62a098913a27e4ee93eb8e28a3b88e3e3ee737c2377f5) |
| `ingress_rules.protocol_port_range.protocol` | [ingress_rules.protocol_port_range.protocol](resources--network_policy_view--reference--group-001.md#canonical-efb21fd48919f4192a0dc7eb6e5ca298e225d1e15e1cee6b910d8e4493356c0a) |
| `labels` | [labels](resources--network_policy_view--reference--group-001.md#canonical-130ad6ff9b65bbda857a7d69f1aea237ff3af39fe52cfdf21f2039faa10004e2) |
| `name` | [name](resources--network_policy_view--reference--group-001.md#canonical-0352238be4368b05e3c6a5e6dbe83780de70057a0ecd7108579896c505725d42) |
| `namespace` | [namespace](resources--network_policy_view--reference--group-001.md#canonical-9330a74063b3a46d7edbc099b1b5900a7857bdbb3dfb2b7af05ba24dfca7c3cb) |
| `timeouts` | [timeouts](resources--network_policy_view--reference--group-001.md#canonical-1a2fb1c095e234e02346376aee238b38499eb349299303b942eedcd1e0f258b0) |
| `timeouts.create` | [timeouts.create](resources--network_policy_view--reference--group-001.md#canonical-687e7f6942d71d56f9a62e01e56a713f49bc74a20ac462f22f9fb07174a9b895) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy_view--reference--group-001.md#canonical-0c81873655005f48fd5e147067bf096efb8c256d21fe4d1727ede347c1189b4b) |
| `timeouts.read` | [timeouts.read](resources--network_policy_view--reference--group-001.md#canonical-7d271c3673314ea7cc2bae41f05503d6fffb4066db3e800f58bedf0d85c1b675) |
| `timeouts.update` | [timeouts.update](resources--network_policy_view--reference--group-001.md#canonical-4293e1f8b49bb067304037908796da248dbf096e382060dc0e3eba07fccc6c05) |

<a id="canonical-3889712ef00d9a8751485e830aa6031a4810714cae6a05ba97926de949658f21"></a>

## Next pages — Property reference / 4d85bc970357 / 12

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [timeouts](resources--network_policy_view--reference--group-001.md#canonical-8d062b5a6c721be625f15b9aaf1708cf895efbca37bfbcf04384388757be2460)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e67804469e0cec3fadd68b71e427ce995574f875ebea62107bd2f01a223b08"></a>

## egress_rules — egress_rules / bd7a0f4b38cc / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- egress_rules

<a id="canonical-1910b2d9ab0f5bfe20edde4559b050f8af2e0f46e281197e331b4f9a3ab2e57a"></a>

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

<a id="canonical-3a98ed8bfecf7ea40eae276b999c005f87918e089ae3a52db0e7b136c719a30a"></a>

## Direct properties — egress_rules / bd7a0f4b38cc / 3

<a id="canonical-e1bcc8ec680c6c02146f3b150d57f9070029d4774f3d6b7d1647e6ee6e8a023f"></a>

<a id="canonical-81a4523799273f4396b9be7a7c0915c65fff893b3fed0643cff1c79b083bcb5d"></a>

## action property — egress_rules / bd7a0f4b38cc / 4

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

- [adv_action](resources--network_policy_view--reference--group-001.md#canonical-ae140796428285d260056247cf8f6fa436d255ac12d5f91d1fa15dcb3266eac7): complete subsection reference.

- [all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-a5eff9f21af2de46831bd7cdcef16f73963c0f77523c357f81856580db96f636): complete subsection reference.

- [all_traffic](resources--network_policy_view--reference--group-001.md#canonical-d4fd210a5a608864b877bff3f10b955d4ec279847d529d9cc2246d7bff145073): complete subsection reference.

- [all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-ab8cd755e9aa25e7f50ed17152c8880c1cf4502ef9793bc7e855f3bc481afbbf): complete subsection reference.

- [any](resources--network_policy_view--reference--group-001.md#canonical-384862fb9becf41c238511245bc59bb05045a0e43ff3d5e8742405b63c74180a): complete subsection reference.

- [applications](resources--network_policy_view--reference--group-001.md#canonical-ad61357e4dfd396696ce7625f139410cc38a94f06d57c46fd1ca47f3f7084867): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-a47ddc237554c4ff91a77225e2df7aef4f3d49cbe3e682bdc4dd1c1a940e2e01): complete subsection reference.

- [ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-9f67338b7227367b4ef7a29c44d7e7147c770bf0d791a798cc32bc20d2b8ce5a): complete subsection reference.

- [label_matcher](resources--network_policy_view--reference--group-001.md#canonical-a842b4a0ee50f5357600748ba9ccb4b6df83e0794373cb39f25d169214ea5069): complete subsection reference.

- [label_selector](resources--network_policy_view--reference--group-001.md#canonical-390d29a0f6a3d9c5ff758528f503705645996bc2cbf55d39abf33e9c640a28c1): complete subsection reference.

- [metadata](resources--network_policy_view--reference--group-001.md#canonical-6a4e5617d3c0c14b2ffdc9f53f80be79487d7cf51561eb32c9e2428e5f09999a): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-1f366c1d47211c18977886f1ccc711fc9ffff5f14ad25cb82f307fa830793b24): complete subsection reference.

- [prefix_list](resources--network_policy_view--reference--group-001.md#canonical-b0fc80ede6a2f9b9a4d75382e3c53785f5e6be2f9dfd964554add40360eff56e): complete subsection reference.

- [protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-7f1ab1e086aaff7b91deaaf64b7c23b5ace4e89d0769b4fd7d9732d014d8742e): complete subsection reference.

<a id="canonical-bb5099ac6fb7923287ffa837e0443abd80571927fa67e6bc04016c8621a04eeb"></a>

## Next pages — egress_rules / bd7a0f4b38cc / 5

- [egress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-ae140796428285d260056247cf8f6fa436d255ac12d5f91d1fa15dcb3266eac7)
- [egress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-a5eff9f21af2de46831bd7cdcef16f73963c0f77523c357f81856580db96f636)
- [egress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-d4fd210a5a608864b877bff3f10b955d4ec279847d529d9cc2246d7bff145073)
- [egress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-ab8cd755e9aa25e7f50ed17152c8880c1cf4502ef9793bc7e855f3bc481afbbf)
- [egress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-384862fb9becf41c238511245bc59bb05045a0e43ff3d5e8742405b63c74180a)
- [egress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-ad61357e4dfd396696ce7625f139410cc38a94f06d57c46fd1ca47f3f7084867)
- [egress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-a47ddc237554c4ff91a77225e2df7aef4f3d49cbe3e682bdc4dd1c1a940e2e01)
- [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-9f67338b7227367b4ef7a29c44d7e7147c770bf0d791a798cc32bc20d2b8ce5a)
- [egress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-a842b4a0ee50f5357600748ba9ccb4b6df83e0794373cb39f25d169214ea5069)
- [egress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-390d29a0f6a3d9c5ff758528f503705645996bc2cbf55d39abf33e9c640a28c1)
- [egress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-6a4e5617d3c0c14b2ffdc9f53f80be79487d7cf51561eb32c9e2428e5f09999a)
- [egress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-1f366c1d47211c18977886f1ccc711fc9ffff5f14ad25cb82f307fa830793b24)
- [egress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-b0fc80ede6a2f9b9a4d75382e3c53785f5e6be2f9dfd964554add40360eff56e)
- [egress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-7f1ab1e086aaff7b91deaaf64b7c23b5ace4e89d0769b4fd7d9732d014d8742e)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-ae140796428285d260056247cf8f6fa436d255ac12d5f91d1fa15dcb3266eac7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95f027213f7ba5e55490a63537a6f4f0697e0d2343a6a00c3b6d92d929a31d59"></a>

## egress_rules.adv_action — egress_rules.adv_action / 4aa6e5ed8d52 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.adv_action

<a id="canonical-db053e8c6b54fd20220e4d18137c3fc51b4d239de3bb31b58f7b4a8c6392a9eb"></a>

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

<a id="canonical-e6f4dc9eaeb3bc6a7ec5c2af62fa005bd9637898c15806242ad8837f1594eb41"></a>

## Direct properties — egress_rules.adv_action / 4aa6e5ed8d52 / 3

<a id="canonical-73a6a0e25b181f68876f90388dfb996352ce4bc59f00540a03301587d81a83eb"></a>

<a id="canonical-ee55678948d302ec3d32ec8ef3f1f7790cdb9baee2674c156edd3da4b513b792"></a>

## action property — egress_rules.adv_action / 4aa6e5ed8d52 / 4

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

<a id="canonical-16145f44ad91bf1fe219bba402858ba1e332d95d15050bd5a88320d1bb7c9c96"></a>

## Next pages — egress_rules.adv_action / 4aa6e5ed8d52 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-a5eff9f21af2de46831bd7cdcef16f73963c0f77523c357f81856580db96f636"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8941e4696f036027afc31bb2a87376d16cfea311e80cbf78623228d8e0c78757"></a>

## egress_rules.all_tcp_traffic — egress_rules.all_tcp_traffic / 051feb26feb1 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.all_tcp_traffic

<a id="canonical-cd4249e887e7b26c09c8096c752dd43ac9e99bda20d57550a5c061d1c155d705"></a>

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

<a id="canonical-8108a4bf98ebb894d2d8d14e30ea458a2d8fcc514259391d5109e82a894b558e"></a>

## Direct properties — egress_rules.all_tcp_traffic / 051feb26feb1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf0b0f257fb2920a911831f8cd7b925d53f2aac9a4ba3cdc9444f76ee2172672"></a>

## Next pages — egress_rules.all_tcp_traffic / 051feb26feb1 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-d4fd210a5a608864b877bff3f10b955d4ec279847d529d9cc2246d7bff145073"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5254c7ed45d49494a95f92d0cc0cad5a2fba8b86508c376951d8576e54907a22"></a>

## egress_rules.all_traffic — egress_rules.all_traffic / b7225b95be32 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.all_traffic

<a id="canonical-ed8ca84bd74cbf1fb061b64547e88a65046e14f8719833189bd6dee9e33b7ec7"></a>

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

<a id="canonical-c59fc9c89a7e0d93604642564af81ede628a31b114ebef8459ba3a6868293885"></a>

## Direct properties — egress_rules.all_traffic / b7225b95be32 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f41debbdd378b23d7bd070069266d4f67444323a52103dd69d630883ff0a49cb"></a>

## Next pages — egress_rules.all_traffic / b7225b95be32 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-ab8cd755e9aa25e7f50ed17152c8880c1cf4502ef9793bc7e855f3bc481afbbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ae24150d1387d92378adacaceecb1e93e4ecfd0b1a1303ac467e913eb8bb566"></a>

## egress_rules.all_udp_traffic — egress_rules.all_udp_traffic / 1ec360155874 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.all_udp_traffic

<a id="canonical-58d3b7ff1793e80fdbc3050934861b15f7a1fb3337c6c15a2754f72b2b160c11"></a>

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

<a id="canonical-a0d3e64e379ac994fbbd63e7efa2b762be5cd962c2790401516ed1a16b945be1"></a>

## Direct properties — egress_rules.all_udp_traffic / 1ec360155874 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58673435fd1d18aa453e230410765ae3c96f68b0b26998aad2bb4282525ae84f"></a>

## Next pages — egress_rules.all_udp_traffic / 1ec360155874 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-384862fb9becf41c238511245bc59bb05045a0e43ff3d5e8742405b63c74180a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88a9b9490d0fcfbeeace9fdc765e526111125c1dbeac334e84693029bfad0543"></a>

## egress_rules.any — egress_rules.any / fae9fb89bada / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.any

<a id="canonical-a281ecdeb7d3d7912430084881fd26d3283e143e7d599d8690636d3767dce044"></a>

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

<a id="canonical-1ce117892ceb2d5bd1e580baff8dba80dabf729c3fa565ab1724450a0502970e"></a>

## Direct properties — egress_rules.any / fae9fb89bada / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b104ae1793df552a7b9dfa9fd321e12935730f2b2901a58deadb24f702ac4a84"></a>

## Next pages — egress_rules.any / fae9fb89bada / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-ad61357e4dfd396696ce7625f139410cc38a94f06d57c46fd1ca47f3f7084867"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f820be32fefec793bca0345598168c90920c242cefb8aab66ce3a23d609ce404"></a>

## egress_rules.applications — egress_rules.applications / e67897456116 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.applications

<a id="canonical-94ee97aff33ecc9e6217688e8e5c2e370a1fa2303b429819792d30752015df6a"></a>

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

<a id="canonical-cdfac5b3c7406723f4442d4bef735a3133b5ef4647239b3bbc34bf77aa80eec8"></a>

## Direct properties — egress_rules.applications / e67897456116 / 3

<a id="canonical-e0ecc95ae2720bdd68a5f270509c9fa80814dec6e52e933a5b71fce0140e0310"></a>

<a id="canonical-cf6ea05c4a2e95497e606f85d7423d52077c2d079b479910b5cccb96d542e266"></a>

## applications property — egress_rules.applications / e67897456116 / 4

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

<a id="canonical-9ed103174d93f336fd53c8486cfae244d690e5b48431cff9d393088751ee9364"></a>

## Next pages — egress_rules.applications / e67897456116 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-a47ddc237554c4ff91a77225e2df7aef4f3d49cbe3e682bdc4dd1c1a940e2e01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a59c300d6eda9c51e27efac35d00376e732a0fc229a15b35f80a94cc4554e5e"></a>

## egress_rules.inside_endpoints — egress_rules.inside_endpoints / 56e5a7ab9e17 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.inside_endpoints

<a id="canonical-a5930803adf3abab2e3dfd38405f5fe47a8be9e0d0f82922d98749376f817fd7"></a>

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

<a id="canonical-bb6e9f59c7e948e8c2d1da042dc10ec13a4ab17352472bec5e5014445b302bae"></a>

## Direct properties — egress_rules.inside_endpoints / 56e5a7ab9e17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c409fc20a2238796252477096453200a5838af447b6ad0d207cc9e06dcbcf25"></a>

## Next pages — egress_rules.inside_endpoints / 56e5a7ab9e17 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-9f67338b7227367b4ef7a29c44d7e7147c770bf0d791a798cc32bc20d2b8ce5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16991fa81c29438691ba8ae12bb70843d0511ca7a08b0ec618b5e81395a8a1d9"></a>

## egress_rules.ip_prefix_set — egress_rules.ip_prefix_set / 58192b56324d / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.ip_prefix_set

<a id="canonical-ba347dd1d85782c39d533bd418833d3280d8fd7adb2bd540a8ade8ec05db5ed1"></a>

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

<a id="canonical-50d1ac4cc102f9f46c2ce236b027bad51a21c35efdd27e1746f028adabf7afbf"></a>

## Direct properties — egress_rules.ip_prefix_set / 58192b56324d / 3

- [ref](resources--network_policy_view--reference--group-001.md#canonical-13943caa2453509b23bee74509a9ccc8528d0d922b6664ac11b57f8a268f49ce): complete subsection reference.

<a id="canonical-f8fe48927da6cdb3747177127f72bb51d71b498cb7b4974f9f83283faccdc6f1"></a>

## Next pages — egress_rules.ip_prefix_set / 58192b56324d / 4

- [egress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-13943caa2453509b23bee74509a9ccc8528d0d922b6664ac11b57f8a268f49ce)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-13943caa2453509b23bee74509a9ccc8528d0d922b6664ac11b57f8a268f49ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75471be68387c2dee166bbe1c74667dc2c933fa17784da0164f332791c1e1eb4"></a>

## egress_rules.ip_prefix_set.ref — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-9f67338b7227367b4ef7a29c44d7e7147c770bf0d791a798cc32bc20d2b8ce5a)
- egress_rules.ip_prefix_set.ref

<a id="canonical-63e5e78f449f8f0b34635e0e5aff0b3a05bd07bd001248b224ca287de3b46297"></a>

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

<a id="canonical-dedd4c1a494a6fdc6ebdb3de3648835305150e9251ff959c99584dc675c0543f"></a>

## Direct properties — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 3

<a id="canonical-eafdf82f4ba482d81ac6de8758badade9393f01859e452779416a552f31bde46"></a>

<a id="canonical-ecda1b00bc60992e04d933dc4b4eadece66fcb1dcaf8b0778ec0b9fa038615f8"></a>

## kind property — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 4

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

<a id="canonical-126ba41616d33c4c57885eee623ea55df501116dfbc01a909bb3f957bd2dfc60"></a>

<a id="canonical-870566642f7dd5a58b323c807a59131f3d0f7cc045d8d8e88fa30d883b224401"></a>

## name property — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 5

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

<a id="canonical-9f73a42e876f3afdeb8e3eb9c415c2e5a8c1057de782aa1c02624f65d3423174"></a>

<a id="canonical-e156cee863435d4acec467af010a8f076aad39eee1d7940391b2473ec49928f6"></a>

## namespace property — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 6

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

<a id="canonical-d9c29b613e867576302026554c05826208aef074fc16495904f6e7729c19e5f2"></a>

<a id="canonical-781f825ada777b2c58451c08ceab34cedae8ec314a8d1b381de8234e171c35d6"></a>

## tenant property — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 7

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

<a id="canonical-a49acfb0e9f9e51a6c4d9bf28fffa9ffe3d82818740067354f52676a33a5a79b"></a>

<a id="canonical-b1ed05e5e39a07aea2cdfb864d8f39d2aec99c2db9ed23acd52bf6865756c9aa"></a>

## uid property — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 8

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

<a id="canonical-37fbcc85270c524755c4142087641968929cfb2f673b93b9f9d52107a70d3345"></a>

## Next pages — egress_rules.ip_prefix_set.ref / 01a11b5c9e84 / 9

- [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-9f67338b7227367b4ef7a29c44d7e7147c770bf0d791a798cc32bc20d2b8ce5a)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-a842b4a0ee50f5357600748ba9ccb4b6df83e0794373cb39f25d169214ea5069"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d1c9ebb8444f547c3911dedc60c400950b4705a68707561db2b8d93ea73f85e"></a>

## egress_rules.label_matcher — egress_rules.label_matcher / 093cc47a5c84 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.label_matcher

<a id="canonical-8d9fb7454ea665515bd4551c15753ac7af38ba2766d4c0692416a07d8a8a1840"></a>

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

<a id="canonical-c2c0cf855103cfe69cfc771722dd4241717c7f4b93b4371219777a6e40c893a1"></a>

## Direct properties — egress_rules.label_matcher / 093cc47a5c84 / 3

<a id="canonical-596a7eaf5daeda8f810d5a5f312dde3010f016271694fd31b288bdda7efe1e6d"></a>

<a id="canonical-4d0979170fe29d069630c9ed97a08704f13d58610b1b9e28549dcc98a9cf556b"></a>

## keys property — egress_rules.label_matcher / 093cc47a5c84 / 4

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

<a id="canonical-b0307a12d51cf81211b97113f5725414dba76a1b660611ba69e13c788cb7a818"></a>

## Next pages — egress_rules.label_matcher / 093cc47a5c84 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-390d29a0f6a3d9c5ff758528f503705645996bc2cbf55d39abf33e9c640a28c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9258a57cdabe71a3bf6ea50b15fcb46ba586c7a3996defc9542edd332bde6162"></a>

## egress_rules.label_selector — egress_rules.label_selector / a298b7f78477 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.label_selector

<a id="canonical-f7e3d694d0c129a71794f4ab0efce0dfcaf14e063b6aebdee9cebcf9c620054d"></a>

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

<a id="canonical-102898ad69b7137b559786162203b8612978a614cf15de1445a0706d903f26e9"></a>

## Direct properties — egress_rules.label_selector / a298b7f78477 / 3

<a id="canonical-cb428613113756554976173bee20c2fc778dd90f22556ea15a1debb05a5811d4"></a>

<a id="canonical-c938a6a3a44da3303c2d1dece08d4fc63c84af11e143f88ae472f028d5eef29a"></a>

## expressions property — egress_rules.label_selector / a298b7f78477 / 4

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

<a id="canonical-fc2da5f64724b758a512935107985730b04c7841cb203c3f5516e9a768e8bbb1"></a>

## Next pages — egress_rules.label_selector / a298b7f78477 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-6a4e5617d3c0c14b2ffdc9f53f80be79487d7cf51561eb32c9e2428e5f09999a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4b41ef6b92c88babb60c0cf5e3fe4bc8d97e4772d9f807cb9b5997f1e9ef84e"></a>

## egress_rules.metadata — egress_rules.metadata / bc02744b3baa / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.metadata

<a id="canonical-5ffe43f827debd8ad07d882075853c2b51427d139aee573cdfb5fe571926e635"></a>

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

<a id="canonical-d556c2c27f77ff27e4ad9adb746ed11fe2ffca7ae09df87585a4afa53aa62bf7"></a>

## Direct properties — egress_rules.metadata / bc02744b3baa / 3

<a id="canonical-902f137a7f9cff7f2b4af5318f7dfc4a7b119570e3a0e6df108118465492413f"></a>

<a id="canonical-efa6920531c0d4956206503fbbfeb4556eb221b78ce9098551661656525ccc79"></a>

## description_spec property — egress_rules.metadata / bc02744b3baa / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-42d9d79e492dd7480856a8ec70e391ad6250ea566ad4bf426d22fd491ffe0441"></a>

<a id="canonical-20fc7d041d5b8f677e382c419ff700f4b8ab0c5206b39d0980c434398c0ef2ae"></a>

## name property — egress_rules.metadata / bc02744b3baa / 5

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

<a id="canonical-a202c5deaa486222d0bae3549c89273aeb095c683420290f7b89aa2f24166b99"></a>

## Next pages — egress_rules.metadata / bc02744b3baa / 6

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-1f366c1d47211c18977886f1ccc711fc9ffff5f14ad25cb82f307fa830793b24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e54cc2b2d0682102aa83c8fa5732dc6ee0a03ee38ea49cec5201746935003f03"></a>

## egress_rules.outside_endpoints — egress_rules.outside_endpoints / fe1d8c034f93 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.outside_endpoints

<a id="canonical-1d8f2ea915da060aa15519312a1c39fd76948d2a14e2fd4a6580e8c0d593701c"></a>

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

<a id="canonical-902f11866082d5b47614b4613617c28c4759844223416e5c6f23a39bec9dfa4e"></a>

## Direct properties — egress_rules.outside_endpoints / fe1d8c034f93 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4d496e1dd5dc6d39822215c7db96488cfd1679d2cb5ef06a76c953755867cf7"></a>

## Next pages — egress_rules.outside_endpoints / fe1d8c034f93 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-b0fc80ede6a2f9b9a4d75382e3c53785f5e6be2f9dfd964554add40360eff56e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc23331db0e952b861340d56dd4cd930d334115d51988f9b2ce95938d637cdfd"></a>

## egress_rules.prefix_list — egress_rules.prefix_list / 02200f0b3187 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.prefix_list

<a id="canonical-d6f06794015c43bd7e1140bc00bb872ccde5583bbf65bf641a2f9e9f26066b3d"></a>

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

<a id="canonical-2446480a3549dc85d4f15ce9190659637b05e02e48d47f486661248aab03e5cf"></a>

## Direct properties — egress_rules.prefix_list / 02200f0b3187 / 3

<a id="canonical-1ac6e443c61f6c855ba1f3f1b40f32aa92e614f2c41e9ab4c9f5960857bbdf04"></a>

<a id="canonical-cde12c9ed366c448fc52b4af88d2dd8fe235b930355820322f763e6e1719daf2"></a>

## prefixes property — egress_rules.prefix_list / 02200f0b3187 / 4

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

<a id="canonical-9c9c23a157c1aa5492fbab04b946074e09e1c94ae8d10dd04bd6fe2c68454a55"></a>

## Next pages — egress_rules.prefix_list / 02200f0b3187 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-7f1ab1e086aaff7b91deaaf64b7c23b5ace4e89d0769b4fd7d9732d014d8742e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8da1379f2c9e95ff25df5595a3a32628e606714d39e1525cbb792ab578a388f"></a>

## egress_rules.protocol_port_range — egress_rules.protocol_port_range / 6db6bcaa269b / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- egress_rules.protocol_port_range

<a id="canonical-a4db90f9f47af2666006173af29c52df06aab4b9f981df51b6dfb6fecf92fcb6"></a>

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

<a id="canonical-d884833d2f6b0035496ccc8101bbeb6ae96696cb76b253212619b08a4b2723e7"></a>

## Direct properties — egress_rules.protocol_port_range / 6db6bcaa269b / 3

<a id="canonical-a8e9277f5e90a9e49de65a9f275fadfeaf46029271e5b20451102636c5e3afdf"></a>

<a id="canonical-42dc1c55b61482479913770abefd10a00ce6032249d2f0c9a07d7d06d78d958d"></a>

## port_ranges property — egress_rules.protocol_port_range / 6db6bcaa269b / 4

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

<a id="canonical-7f3cc03d479bdeaaec9c930c9eb06952d8a9461580e83766cf0fa99b62d7dfa6"></a>

<a id="canonical-214b398bf8a2f54d30dbc54f1fe2b7ac11a38cc3c5e5bee616cb3dc913c489fe"></a>

## protocol property — egress_rules.protocol_port_range / 6db6bcaa269b / 5

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

<a id="canonical-39d4883ee9063ed646726f29be48c85954adfd6213e5d48059cab378b9a3357e"></a>

## Next pages — egress_rules.protocol_port_range / 6db6bcaa269b / 6

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-5a1ceb3a762c34b6ca8f172017029990eae81c9396d2483d4e013c94704d2b27)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d4d7a3efe4a0984f9049901a8797f35bfe9a4ef5ec523e9f8f49518578538fd"></a>

## endpoint — endpoint / 60c578fe9785 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- endpoint

<a id="canonical-014d0d5b4a08804e205b3c2411491a04ca41f51a182760b5a7322ffdc711c010"></a>

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

<a id="canonical-216db60c62356409bbc8f803a7f981c1066bcb505f70fb4b8a6b25f9f06c8b69"></a>

## Direct properties — endpoint / 60c578fe9785 / 3

- [any](resources--network_policy_view--reference--group-001.md#canonical-abf47ab3829bb355c106803f4f8aff18c5811a6a83fa424c257a0f187f50af71): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2636dc131d4cfbaddb5550ebd4e3fd243bce702f7817af4a3cfdcb80448ba1a4): complete subsection reference.

- [label_selector](resources--network_policy_view--reference--group-001.md#canonical-4b48c6eada23e19aefa47804e2b383c5c1bc6f2b84cdaa27c86f5c2825c33e06): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-9600a61820455a05c73c55f03e23ece9cc1eea585cebc37ae40cb0ff01328af9): complete subsection reference.

- [prefix_list](resources--network_policy_view--reference--group-001.md#canonical-b762b3cf7c6fda24f482af7659acaf8b28f36005170a81df52ed856b9bf38fa9): complete subsection reference.

<a id="canonical-1f64377d1bc60ac7e872fa238365aa9d82f85b056ce33587a792d369c22c4803"></a>

## Next pages — endpoint / 60c578fe9785 / 4

- [endpoint.any](resources--network_policy_view--reference--group-001.md#canonical-abf47ab3829bb355c106803f4f8aff18c5811a6a83fa424c257a0f187f50af71)
- [endpoint.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2636dc131d4cfbaddb5550ebd4e3fd243bce702f7817af4a3cfdcb80448ba1a4)
- [endpoint.label_selector](resources--network_policy_view--reference--group-001.md#canonical-4b48c6eada23e19aefa47804e2b383c5c1bc6f2b84cdaa27c86f5c2825c33e06)
- [endpoint.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-9600a61820455a05c73c55f03e23ece9cc1eea585cebc37ae40cb0ff01328af9)
- [endpoint.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-b762b3cf7c6fda24f482af7659acaf8b28f36005170a81df52ed856b9bf38fa9)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-abf47ab3829bb355c106803f4f8aff18c5811a6a83fa424c257a0f187f50af71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcc99ffc1780d947d11281b519a16fe54ee5e58e365b4dc4462456d65f0939c9"></a>

## endpoint.any — endpoint.any / 23e8bd1d06e7 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- endpoint.any

<a id="canonical-0524977894d36d8c062b34e6c435725e391636f173053e3dba1208348080e5c6"></a>

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

<a id="canonical-2f1dbf577df59fa5c11c2cb6ca4b3762458ff0b17469966528d444585526a302"></a>

## Direct properties — endpoint.any / 23e8bd1d06e7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-93ef6428762736fde2435e57790023bcb183ec87e2ef4397f53124993328f7f7"></a>

## Next pages — endpoint.any / 23e8bd1d06e7 / 4

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-2636dc131d4cfbaddb5550ebd4e3fd243bce702f7817af4a3cfdcb80448ba1a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3508155e9018a73c85d9b3c87e6c532318f0d4def36adff8b79ea55ab8a62c2d"></a>

## endpoint.inside_endpoints — endpoint.inside_endpoints / 43e0a273c467 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- endpoint.inside_endpoints

<a id="canonical-b817690e8e159602ccd2b16d891cd460c76454124e2bf463fe64faf01e682565"></a>

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

<a id="canonical-05fab056252edff36c8b683e0199bd469f25f18b8e53c1e7606c324fcddb353e"></a>

## Direct properties — endpoint.inside_endpoints / 43e0a273c467 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6e493293553bb0ac86cdf819d64d2df9001cf3b9f673635e58df71cae241d41c"></a>

## Next pages — endpoint.inside_endpoints / 43e0a273c467 / 4

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-4b48c6eada23e19aefa47804e2b383c5c1bc6f2b84cdaa27c86f5c2825c33e06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd23ab49eae188e082ada21853640887c262a265645b41cd77039171e0309957"></a>

## endpoint.label_selector — endpoint.label_selector / a5e6c5f56ad0 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- endpoint.label_selector

<a id="canonical-7db05c617ae49f9404dbe582de12d123cd28ce20d8961b79f912f7ae9bf64390"></a>

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

<a id="canonical-b5b6c90b565d5bb31022337c8a2180487f8d17539c0b63ab6b5d9d7fd673e3ad"></a>

## Direct properties — endpoint.label_selector / a5e6c5f56ad0 / 3

<a id="canonical-13641802134796117f0a480c3a17ca171cf5c9630268c2bf0b2fee47109d371a"></a>

<a id="canonical-c2e04be981f825eb5bcfc8eee897381bdcb3739335ebb8470a31492ae189ad92"></a>

## expressions property — endpoint.label_selector / a5e6c5f56ad0 / 4

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

<a id="canonical-d8f3ee0ea6b1f317b9bd36ab2a5cf7935d7151f3fe20991e92b683573664b218"></a>

## Next pages — endpoint.label_selector / a5e6c5f56ad0 / 5

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-9600a61820455a05c73c55f03e23ece9cc1eea585cebc37ae40cb0ff01328af9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4019eec001b6ac74e8b84a9283ef8b75502e9eece8683474f436bbee00133b4d"></a>

## endpoint.outside_endpoints — endpoint.outside_endpoints / 693e4f43ced5 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- endpoint.outside_endpoints

<a id="canonical-c7ffa3f7eb582897cc5988824f722ba3363aec7a5a70d0707b1f95a8facfaf3a"></a>

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

<a id="canonical-0a0020904de20bb9627231971dc9c10edadceda74eeb96e78dbacf3d7b2d4f7e"></a>

## Direct properties — endpoint.outside_endpoints / 693e4f43ced5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8d23a78ff4cdc929697dbbd631f6ada9a20f425729a9749115070a1eb515ae72"></a>

## Next pages — endpoint.outside_endpoints / 693e4f43ced5 / 4

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-b762b3cf7c6fda24f482af7659acaf8b28f36005170a81df52ed856b9bf38fa9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8467ea401fb597495a9bbb3ceb17f5e538f767a34ff7c73b4b55b94ac1ed50d"></a>

## endpoint.prefix_list — endpoint.prefix_list / d5a50c95369c / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- endpoint.prefix_list

<a id="canonical-be47c7a58aae4b12e1af924732177b6e9dc8d80b385350f09c9febce62f24e1c"></a>

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

<a id="canonical-3461a27ca77cae92333ca5e4f2469eb5b9ec89b27bbc55855bf17bee08a82d9b"></a>

## Direct properties — endpoint.prefix_list / d5a50c95369c / 3

<a id="canonical-f8e37f8a3177c4a395f8c536f780640be8ab6f47f07f58ce463d85522305b265"></a>

<a id="canonical-c1fdd989d05a5304c7a5d45ba5beebaf5c71bbb1ee8f56e14e97b29dc513060b"></a>

## prefixes property — endpoint.prefix_list / d5a50c95369c / 4

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

<a id="canonical-11ca30d68b615a12ea9aaf0a27c6db39cdf872b6b6a3f0958ba0a6ac42d9dca0"></a>

## Next pages — endpoint.prefix_list / d5a50c95369c / 5

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-f0d84bc57e6e63dafcce08aa169712a01993b75050ebd089ee84378ae1d3719b)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25df5cd6d8b3549eab651149bfbf9a6ffb887dc59fefba667174b00b8c1f27ac"></a>

## ingress_rules — ingress_rules / d9c903587328 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- ingress_rules

<a id="canonical-c54f2c2a3f6e2bf5f9af039a96fa4453f45dd742831fb0ddbe81ee8d37464dd8"></a>

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

<a id="canonical-90cbe6cad128fc7621bbb85cd6e0d5d52fb0212a3c89e57a64e54a8080ca4397"></a>

## Direct properties — ingress_rules / d9c903587328 / 3

<a id="canonical-b0bd2122d408db20ff4c56b5d305bdd84be1d94eb987f74b467f68738e1ca933"></a>

<a id="canonical-b106ea12180b6f2648480c4e692b054473b713521ea495b92ca6871944a635ff"></a>

## action property — ingress_rules / d9c903587328 / 4

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

- [adv_action](resources--network_policy_view--reference--group-001.md#canonical-a984aa7a63f56c7aacad6986adbcf7dcf6db15fbc6d57c3a602bc2c5f84d20c6): complete subsection reference.

- [all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-bf5d24da1201eef4aa1270919985c3ae3eecb1ff1fc32c9da28b5aeeab2c5a2a): complete subsection reference.

- [all_traffic](resources--network_policy_view--reference--group-001.md#canonical-812619bc5169858ed48b57a05a86feac80426c9be07779459c9454609ea897f0): complete subsection reference.

- [all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-f886bc7bf72c670d246916009a33d640f647895cbd911d01afc3737083f3bc0f): complete subsection reference.

- [any](resources--network_policy_view--reference--group-001.md#canonical-142e96805e45f8a1cb63b4a6c8b87c2380a629b575ad34ac99f51c837fbe3ae5): complete subsection reference.

- [applications](resources--network_policy_view--reference--group-001.md#canonical-f94171facdd0e5dad81100de3a0ad18b9d48ad86aa1d96c62afb94799dea489d): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-316c10be5ea20c04293dc59f38bc0620f6a0c4f06b1217db1eddbbab16811e12): complete subsection reference.

- [ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-65927adbf2e77f9f4eef4735c4ee07f940c579e3f170ccdf06ceec71929ae163): complete subsection reference.

- [label_matcher](resources--network_policy_view--reference--group-001.md#canonical-d983abb1b6c5acd907cb53bcf94872c34e816354588062932021c9b62611fe33): complete subsection reference.

- [label_selector](resources--network_policy_view--reference--group-001.md#canonical-6b9323cf5323dc0b389a83829599af4ba4c2ebe13fe47b32636e8b45383a8642): complete subsection reference.

- [metadata](resources--network_policy_view--reference--group-001.md#canonical-fe28e30826d5f24960af0b5255db68f39d96ec07047d36bf03db7ca755eba970): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-29082cd786ff5813b1f5ef213058a2c22c9bbf6e6fd942a45e4d3da9895ebdd7): complete subsection reference.

- [prefix_list](resources--network_policy_view--reference--group-001.md#canonical-fdb3573834991aa6c0f449f9766ee10cc92a947c6eed689cb0c07d172f97e512): complete subsection reference.

- [protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-cee1b5b88a9b09fd35c338df45f73b756e0540c027ada98bf768ef9e64ddea1f): complete subsection reference.

<a id="canonical-e3e556ea0dd3097776beff4fdf834d2c31e06bbb1f22200ed5ecf4a9517250f8"></a>

## Next pages — ingress_rules / d9c903587328 / 5

- [ingress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-a984aa7a63f56c7aacad6986adbcf7dcf6db15fbc6d57c3a602bc2c5f84d20c6)
- [ingress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-bf5d24da1201eef4aa1270919985c3ae3eecb1ff1fc32c9da28b5aeeab2c5a2a)
- [ingress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-812619bc5169858ed48b57a05a86feac80426c9be07779459c9454609ea897f0)
- [ingress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-f886bc7bf72c670d246916009a33d640f647895cbd911d01afc3737083f3bc0f)
- [ingress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-142e96805e45f8a1cb63b4a6c8b87c2380a629b575ad34ac99f51c837fbe3ae5)
- [ingress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-f94171facdd0e5dad81100de3a0ad18b9d48ad86aa1d96c62afb94799dea489d)
- [ingress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-316c10be5ea20c04293dc59f38bc0620f6a0c4f06b1217db1eddbbab16811e12)
- [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-65927adbf2e77f9f4eef4735c4ee07f940c579e3f170ccdf06ceec71929ae163)
- [ingress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-d983abb1b6c5acd907cb53bcf94872c34e816354588062932021c9b62611fe33)
- [ingress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-6b9323cf5323dc0b389a83829599af4ba4c2ebe13fe47b32636e8b45383a8642)
- [ingress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-fe28e30826d5f24960af0b5255db68f39d96ec07047d36bf03db7ca755eba970)
- [ingress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-29082cd786ff5813b1f5ef213058a2c22c9bbf6e6fd942a45e4d3da9895ebdd7)
- [ingress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-fdb3573834991aa6c0f449f9766ee10cc92a947c6eed689cb0c07d172f97e512)
- [ingress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-cee1b5b88a9b09fd35c338df45f73b756e0540c027ada98bf768ef9e64ddea1f)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-a984aa7a63f56c7aacad6986adbcf7dcf6db15fbc6d57c3a602bc2c5f84d20c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef39ef38b50a2e01f911a2fd458a984f157cc4d9a4617a3ae6b4ebc71a2cbbe7"></a>

## ingress_rules.adv_action — ingress_rules.adv_action / 2f33d381ab3e / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.adv_action

<a id="canonical-a4b2a8600cd7ef5d54249816835243ff1fa2148860ecc5d0565bf00297a12e30"></a>

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

<a id="canonical-ec798b2e6898dc435013c3291137a2d585a0c3d440b6faa072cfddd919872208"></a>

## Direct properties — ingress_rules.adv_action / 2f33d381ab3e / 3

<a id="canonical-92e6d1e1293a765e3916b0099415343cd14b7d8cd3ee220156720d5584b43b5f"></a>

<a id="canonical-77d068befc8ace020b30881d897604ca6beb83ea36cc63c9acd2a4ee864ad243"></a>

## action property — ingress_rules.adv_action / 2f33d381ab3e / 4

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

<a id="canonical-9496914ecef9ff5af95008e333a7ff01d110e702b793e36f64abf192f48e1781"></a>

## Next pages — ingress_rules.adv_action / 2f33d381ab3e / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-bf5d24da1201eef4aa1270919985c3ae3eecb1ff1fc32c9da28b5aeeab2c5a2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efd882468e902be50fe39bd7e211cb038b91a67f9962a8c668ca5cb4e0af3253"></a>

## ingress_rules.all_tcp_traffic — ingress_rules.all_tcp_traffic / 020b09329cea / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.all_tcp_traffic

<a id="canonical-5a172ff953db1ecf333d4e66d17f07393832c3a3ee46d07679a693a4e9616e1e"></a>

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

<a id="canonical-22abd26e3124a636c10feeffd0dd2e91a98c835f838440f6598ffae95da86241"></a>

## Direct properties — ingress_rules.all_tcp_traffic / 020b09329cea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f014428600067ef7e960957d33cab6fbb10a9523f33d499ab864bfe615b9af30"></a>

## Next pages — ingress_rules.all_tcp_traffic / 020b09329cea / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-812619bc5169858ed48b57a05a86feac80426c9be07779459c9454609ea897f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-513fbdfe307c93a57da15d7edaa96f7909b17a66c91bfab94ef97c255c0cb67e"></a>

## ingress_rules.all_traffic — ingress_rules.all_traffic / 386a5327c036 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.all_traffic

<a id="canonical-16c4f4c3639d1989a15e44b591f50907084fc8c0911089c11078545e1ffcb2fd"></a>

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

<a id="canonical-ab41e4243f46a430d023adc91c7a08517b7fc750fac4082442b817e9ce9173e7"></a>

## Direct properties — ingress_rules.all_traffic / 386a5327c036 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b95de71829b43144fc688937bf57aec05f5d9988cfcda82e546245b6e81e3ae"></a>

## Next pages — ingress_rules.all_traffic / 386a5327c036 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-f886bc7bf72c670d246916009a33d640f647895cbd911d01afc3737083f3bc0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17e4df65c357d7e02f64742e696af4042e2f036caf82e3309b7a0ddd4c4f5b4d"></a>

## ingress_rules.all_udp_traffic — ingress_rules.all_udp_traffic / 00f11936a725 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.all_udp_traffic

<a id="canonical-00cb29625c8211969ab21cc5b9c13342ba03da0673f65c551bca31d37abe8a35"></a>

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

<a id="canonical-49be0cec7b10a545eb6b5638a5c60597610d171954b800439e726f3677e962ba"></a>

## Direct properties — ingress_rules.all_udp_traffic / 00f11936a725 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-70013485b89080a6d1213b3c7fa9de3aa0c88bf5afba5d5e6a4d3b248da2dbac"></a>

## Next pages — ingress_rules.all_udp_traffic / 00f11936a725 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-142e96805e45f8a1cb63b4a6c8b87c2380a629b575ad34ac99f51c837fbe3ae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4695e2226efe9c65b53987d287818cec4af7a5b046574b104fe8fe6da486f93f"></a>

## ingress_rules.any — ingress_rules.any / 9e9ebc2f30ea / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.any

<a id="canonical-e68825c0b22208f8a31329c3e386bfff01a2125a8090b006c2d656f24bc5be24"></a>

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

<a id="canonical-1ceff832c398ae325beda79a3764728ae700b326fc61e423bedeca1bc77b1b07"></a>

## Direct properties — ingress_rules.any / 9e9ebc2f30ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d97e8798b584ec086e967380eff25af0f59db44498cb8b94c39e24afdee5ff63"></a>

## Next pages — ingress_rules.any / 9e9ebc2f30ea / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-f94171facdd0e5dad81100de3a0ad18b9d48ad86aa1d96c62afb94799dea489d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71b9f9aeda1853ac4452e7d33d76624aacfaa393490086ef93563bb6a6a1ac88"></a>

## ingress_rules.applications — ingress_rules.applications / fcce4627e6b7 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.applications

<a id="canonical-85adc2c23a613c1543bfea4c53359394f8fd4a7777c74b844574cbba003667b8"></a>

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

<a id="canonical-0f8cfff757044ca68e22c5a9399eccc88245f3aac2344cfaf7a5418d92c13836"></a>

## Direct properties — ingress_rules.applications / fcce4627e6b7 / 3

<a id="canonical-1a0912e2c0e3029c7b01d2fa6aef5ba5fef2965d0589d9047f8335077f931070"></a>

<a id="canonical-1d8fc53d19d4eabb90b4979170c0ae8461f54842d74295aaf5f83dbe95441d1d"></a>

## applications property — ingress_rules.applications / fcce4627e6b7 / 4

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

<a id="canonical-a277a393c756d339e13e21fa968f6a5a827cae4f4433ea9ef42b99d468eba6e3"></a>

## Next pages — ingress_rules.applications / fcce4627e6b7 / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-316c10be5ea20c04293dc59f38bc0620f6a0c4f06b1217db1eddbbab16811e12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-507a48fada54ffd48b89889f15b32d5e7c8b844732aec33ba4d0388eb3ed5fe3"></a>

## ingress_rules.inside_endpoints — ingress_rules.inside_endpoints / bc5519ec2650 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.inside_endpoints

<a id="canonical-0fa6b1de9cd8da6a36e60842a7c267961500fcd84989fd5546816a24af1a9a63"></a>

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

<a id="canonical-e28835dd33b43a9ff4a984eba3fde03efc42824637b388f3e8ec8961f1b81d66"></a>

## Direct properties — ingress_rules.inside_endpoints / bc5519ec2650 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e993f1c2e6920545c6cfa355a6998f21951df9b97ea77f71b603ce22145faeb"></a>

## Next pages — ingress_rules.inside_endpoints / bc5519ec2650 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-65927adbf2e77f9f4eef4735c4ee07f940c579e3f170ccdf06ceec71929ae163"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e40e902eb944a45ac632d1adbe4cb4d582a5a19b97308079ce65c72237957f56"></a>

## ingress_rules.ip_prefix_set — ingress_rules.ip_prefix_set / 8b7df67ba3a6 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.ip_prefix_set

<a id="canonical-3f4dbc82f485afb1e6df415af99cdf2d833cc312878bcf4de2b865e621cfde7e"></a>

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

<a id="canonical-cc449e457e18574d7845637e76ca3b0a473b73f30142d395888e78ed9c7ebd80"></a>

## Direct properties — ingress_rules.ip_prefix_set / 8b7df67ba3a6 / 3

- [ref](resources--network_policy_view--reference--group-001.md#canonical-f4aa5afa02d33242ac21f1f3c8bb3d41239217b12764a64940ca1f897ff01ed0): complete subsection reference.

<a id="canonical-d00fdd4c0d463a255852af33dae8c87c95d4538f01301dda3f8e374272f7d03c"></a>

## Next pages — ingress_rules.ip_prefix_set / 8b7df67ba3a6 / 4

- [ingress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-f4aa5afa02d33242ac21f1f3c8bb3d41239217b12764a64940ca1f897ff01ed0)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-f4aa5afa02d33242ac21f1f3c8bb3d41239217b12764a64940ca1f897ff01ed0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8d7a325d678539290af37074e773222abcb35bfcf5a79d7528c3cbdafb784d7"></a>

## ingress_rules.ip_prefix_set.ref — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-65927adbf2e77f9f4eef4735c4ee07f940c579e3f170ccdf06ceec71929ae163)
- ingress_rules.ip_prefix_set.ref

<a id="canonical-e04d2e19ceb1bce852ac59a664ba0e30be74d35af0e55921a9f58ccaed57e8b2"></a>

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

<a id="canonical-6b4126b09bbf0c94499c21ea4d1d5c3efdae3bdd7a4bd6c5f642006b0ebb42d8"></a>

## Direct properties — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 3

<a id="canonical-76a2d3cd62fceb7420caeac109115b098ff68e4ebe515898901db4cf66f2ae3e"></a>

<a id="canonical-60df303f43e4670c2660cb6075071644ceb08ab134b3182b621ba0bed070e5f7"></a>

## kind property — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 4

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

<a id="canonical-d6fc4c50e0fc61af2838f9a9a6f3492f234dd24c1d120e7fa49536ab454b06d0"></a>

<a id="canonical-2c4ad29c3a6f7007d8361eb881222a44f1a61912ce7f6bf4678d0c351f4eee9d"></a>

## name property — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 5

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

<a id="canonical-7e480e959c13ac31da91bfacb61dcce62f9267f015a76d65afdffa69403c5375"></a>

<a id="canonical-6a46d7b41074d23193c8d722dab0f5d192829e2fe2dc3039bc008afeb1fd38fd"></a>

## namespace property — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 6

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

<a id="canonical-b5a8d6300112726ba1afd6d55576fa31d0493f2dd8b5c7f977e94be7c7255495"></a>

<a id="canonical-e058dbc5b6143d12c03c3976f1ca3514491996e3bd98dd9cb0dfa46913a55215"></a>

## tenant property — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 7

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

<a id="canonical-c8c47a90c590cf993cf3b01f97a90775d0e3e1dde1ed5c39e02558a259e6f7ee"></a>

<a id="canonical-f4c4cdb33c2a8ebe22adcd71c47ef8d6626ed0c78be0108c370c672d30edc567"></a>

## uid property — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 8

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

<a id="canonical-4bdafc134f8f8ba42fca6bb9d0a90d49d2ee0bba0e887f8ff1eaece5b58d0948"></a>

## Next pages — ingress_rules.ip_prefix_set.ref / 7f5e1817b113 / 9

- [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-65927adbf2e77f9f4eef4735c4ee07f940c579e3f170ccdf06ceec71929ae163)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-d983abb1b6c5acd907cb53bcf94872c34e816354588062932021c9b62611fe33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d65fbbc994419816d978887727714dd71ac0b1d4ebc9a879951cb5b5cf793762"></a>

## ingress_rules.label_matcher — ingress_rules.label_matcher / 472fcbe6899b / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.label_matcher

<a id="canonical-d0ae1859256be6924ba797f292046f861bca9dd32cf76ed53dabdc61bb94beaf"></a>

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

<a id="canonical-8e2f19442ea09f8c0cacff2dee6cb20746a86000c393b4587d88802c56eff957"></a>

## Direct properties — ingress_rules.label_matcher / 472fcbe6899b / 3

<a id="canonical-e9dfb84e607552236a51fbfdd3ad7bbbdaac83addc506214edb61458319255da"></a>

<a id="canonical-46b2c078abc96e4421f5c8d280bc18a29d15a40dc155b9cbdbda8ed3f0594fa4"></a>

## keys property — ingress_rules.label_matcher / 472fcbe6899b / 4

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

<a id="canonical-a9935cdcd527269e0f7926a4e281e27c495516eaa4043bbc31b9e4ee0d62a66c"></a>

## Next pages — ingress_rules.label_matcher / 472fcbe6899b / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-6b9323cf5323dc0b389a83829599af4ba4c2ebe13fe47b32636e8b45383a8642"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fe7dd6c81feba14755a749ec8596386e3c2d9d494294a71ebd42c9e1bacc606"></a>

## ingress_rules.label_selector — ingress_rules.label_selector / b98d78c61fab / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.label_selector

<a id="canonical-f4a2c1ce4d9aed378e9db5b724dc4fb8f329e91993c5dc5a5ed292ed4159fa6d"></a>

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

<a id="canonical-08655386b0856e5ad630913ab711ab8998d8cf6b9c92b9a779e8f4e6b5a69ec5"></a>

## Direct properties — ingress_rules.label_selector / b98d78c61fab / 3

<a id="canonical-7d0197f5bc5b0ad567b2d0f966bf53341d4bdf6f150804bae770b692cfc2b50d"></a>

<a id="canonical-31c30c3e8f34f1e2fb946ec175b0f4b6f24c1fdf11d717a1742dc90aa0b85ae7"></a>

## expressions property — ingress_rules.label_selector / b98d78c61fab / 4

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

<a id="canonical-51bc3f421b05e9c8cb300f606b646bf24d06dcc46bb7bbf73bcc53b9f9a2de8e"></a>

## Next pages — ingress_rules.label_selector / b98d78c61fab / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-fe28e30826d5f24960af0b5255db68f39d96ec07047d36bf03db7ca755eba970"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67ac88fae6960e2c03ddb882f9fae3a5a25faf720299cb76483ce4aa348e0f33"></a>

## ingress_rules.metadata — ingress_rules.metadata / 74ca7a2efcef / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.metadata

<a id="canonical-de84a1449a98dff10a71b369be888a340ac35df695d6643e6888dca1333cbc8b"></a>

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

<a id="canonical-0c1bc17d1c90308d2e55f06e4324ce2c0ecb9806af70cf4d162109d87d93dff4"></a>

## Direct properties — ingress_rules.metadata / 74ca7a2efcef / 3

<a id="canonical-4c433b435236c452e24e87c5c1427bb27f769dce942bdd4904b1f9ad45558ecd"></a>

<a id="canonical-ca62a1d2bc3b76ad39c74cf39f46d22ee720844688a753469d87ec8f2b00d10c"></a>

## description_spec property — ingress_rules.metadata / 74ca7a2efcef / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-16b19a5757ce14dc42da6fea3a510376651e4600f946b00043c6aee8348f895c"></a>

<a id="canonical-f59633c222e0666d0aa091d2c4fbba3ad553c726816c664f6e8f18c5d66a78a7"></a>

## name property — ingress_rules.metadata / 74ca7a2efcef / 5

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

<a id="canonical-585f61024490e29a81813d53c63f1a45ac85c6516b50bd521e0fa5e947be23d5"></a>

## Next pages — ingress_rules.metadata / 74ca7a2efcef / 6

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-29082cd786ff5813b1f5ef213058a2c22c9bbf6e6fd942a45e4d3da9895ebdd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2d2e060bc1b155a25d24414de705fc5291e5e2185fea697766bf875859980df"></a>

## ingress_rules.outside_endpoints — ingress_rules.outside_endpoints / ec379f3192c0 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.outside_endpoints

<a id="canonical-62aacd058418d151aee2ef8d179ef2f7b482ecd98ed25339125a39b9e562ea51"></a>

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

<a id="canonical-42cbd32522f40e7a1eb0c89e70ec238806891f5b236f9f6b285a90d0aa3641fb"></a>

## Direct properties — ingress_rules.outside_endpoints / ec379f3192c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62a5832342411c04fedeee2ad7ee5371890b009b7c8c360ff98b4fccc96c060d"></a>

## Next pages — ingress_rules.outside_endpoints / ec379f3192c0 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-fdb3573834991aa6c0f449f9766ee10cc92a947c6eed689cb0c07d172f97e512"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d8c74787fe30d7368b5b3829f6dff4e85155405a52f60f9058381a29c5eff52"></a>

## ingress_rules.prefix_list — ingress_rules.prefix_list / 87c33d136f41 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.prefix_list

<a id="canonical-cbe40be3c94b4ba3e0013654402f46150440b69be19034d242e96ba52e990580"></a>

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

<a id="canonical-1db6692300ae4c334405fcda8414eb529aba767bf880a474005c82a57abf032f"></a>

## Direct properties — ingress_rules.prefix_list / 87c33d136f41 / 3

<a id="canonical-95eef45b83aa6f4aef7ba5d6418848675fe2e25031b658c8a5d567cd1f893f4c"></a>

<a id="canonical-b7e338e855cc41f217f442886c2b8ee37ab9c5aabcce8ac316d26027ac3027c4"></a>

## prefixes property — ingress_rules.prefix_list / 87c33d136f41 / 4

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

<a id="canonical-35908b2131061d5aca7cfbb70958dbd9f85c62d1d53a2bc47dae0bcf21d88cc7"></a>

## Next pages — ingress_rules.prefix_list / 87c33d136f41 / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-cee1b5b88a9b09fd35c338df45f73b756e0540c027ada98bf768ef9e64ddea1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7b945dad02661d897d3d96e6fd289624f5fc4730e5d316c5275b85f36db6f9b"></a>

## ingress_rules.protocol_port_range — ingress_rules.protocol_port_range / fefc0e9b940a / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- ingress_rules.protocol_port_range

<a id="canonical-090464e0c2ac81f4f72a9b79f1d9b530c5f8db58e141061901c1fee2a1f89c47"></a>

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

<a id="canonical-28d6768943ea3b9bf3ec6dbcbb3f8fe2a67b05300880d36f30aad130d65bdfee"></a>

## Direct properties — ingress_rules.protocol_port_range / fefc0e9b940a / 3

<a id="canonical-6a917827b78e9e2adac62a098913a27e4ee93eb8e28a3b88e3e3ee737c2377f5"></a>

<a id="canonical-dcc72e2a3177de8be1e61d1cafcc3da9b7bb70224e426cb93a60e287fe34c3dc"></a>

## port_ranges property — ingress_rules.protocol_port_range / fefc0e9b940a / 4

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

<a id="canonical-efb21fd48919f4192a0dc7eb6e5ca298e225d1e15e1cee6b910d8e4493356c0a"></a>

<a id="canonical-4f5b9ad361401893e550d5a6bf76fe65c30f893f52899e8ee18f7f5cc55a044e"></a>

## protocol property — ingress_rules.protocol_port_range / fefc0e9b940a / 5

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

<a id="canonical-dc06cc7e5118af3a84441b7dfa76e41efb5ea519860acd332e69ece2a29aff3b"></a>

## Next pages — ingress_rules.protocol_port_range / fefc0e9b940a / 6

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3d2b16c04ecba450c7473356a03700d95e359d976bfffb8907f6e13e34d0d007)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-8d062b5a6c721be625f15b9aaf1708cf895efbca37bfbcf04384388757be2460"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eed655fa20ea7d439cdc5b88ccf707d9d86685fb464cead3cd16f38c185745b5"></a>

## timeouts — timeouts / 9b249b727eda / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- timeouts

<a id="canonical-1a2fb1c095e234e02346376aee238b38499eb349299303b942eedcd1e0f258b0"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a58e11e20ae070e238f7815ae15c28e266d81202ad02896f840a1a0c99bf56b5"></a>

## Direct properties — timeouts / 9b249b727eda / 3

<a id="canonical-687e7f6942d71d56f9a62e01e56a713f49bc74a20ac462f22f9fb07174a9b895"></a>

<a id="canonical-e4be8642218c9e1f01f5a624dedfdad7024336ed51e25639c0169f6724ea3432"></a>

## create property — timeouts / 9b249b727eda / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0c81873655005f48fd5e147067bf096efb8c256d21fe4d1727ede347c1189b4b"></a>

<a id="canonical-0cf5ef3e386822033c30359223390deb885c295121b91f88b086a20a70ee4844"></a>

## delete property — timeouts / 9b249b727eda / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-7d271c3673314ea7cc2bae41f05503d6fffb4066db3e800f58bedf0d85c1b675"></a>

<a id="canonical-26aff38e798c526080b25012dfd6b95d0f1e2e7eabd18f95c5b72fde3f22f548"></a>

## read property — timeouts / 9b249b727eda / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-4293e1f8b49bb067304037908796da248dbf096e382060dc0e3eba07fccc6c05"></a>

<a id="canonical-6e92f9d625bda166e1cbfd203059302966fdcc16721f0bb83e724f4249ba6f84"></a>

## update property — timeouts / 9b249b727eda / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-816bc1954d477340eeb4b013ca9afaa3282ef6b53e1f89f97bbcfcf0b6037b77"></a>

## Next pages — timeouts / 9b249b727eda / 8

- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
