---
page_title: "xcsh_policy_based_routing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing reference."
---

# xcsh_policy_based_routing reference

<a id="canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca300d78fb05a07400e32054171b585c01a115296f7d02cf0a00f262c63e517a"></a>

## Property reference — Property reference / 49c60d212bb9 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- Property reference

<a id="canonical-5108613071e4458d4857497a02f1ebf8a8fec2e00500f320eb899ea9adbfed38"></a>

## Direct properties — Property reference / 49c60d212bb9 / 3

<a id="canonical-24531696bdc7c839e9707d48f8de53db4ee2eb929f8458a1030ebe23b77cf030"></a>

<a id="canonical-92d524e043b081ea649bc4e904ffd60ea621d2495ec0e64d6f289f53d0207376"></a>

## annotations property — Property reference / 49c60d212bb9 / 4

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

<a id="canonical-51dc18726448fbd00f84eb57909cfacaf8e74a8906b4090d8ce6eac5917f1df5"></a>

<a id="canonical-ceb3b70d8664876863d43901d008a1c85fdf6606c79e68d770574d0e41641390"></a>

## description property — Property reference / 49c60d212bb9 / 5

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

<a id="canonical-13460366836cb23651ba98f3d920173e08dee854208cfd94a95f3df15e041e8f"></a>

<a id="canonical-3b0c32616bae1b25046c1bdafb8b5873ca7ace8ffffe70c0700f10d89b8e025a"></a>

## disable property — Property reference / 49c60d212bb9 / 6

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

- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328): complete subsection reference.

- [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-19eaad8e0e20a025d92dc4e8e8ad8b28101e842052d23859924f7873998d59f0): complete subsection reference.

<a id="canonical-bba2e59ca6f40a52f459f08aa58ae7151cfc52129a1ba037cf4eeb1cf431f7d3"></a>

<a id="canonical-32c6cab10027f1b5b505d4697007bd55791e1bbd2c4ee376f4352dc08f1c7995"></a>

## id property — Property reference / 49c60d212bb9 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-668614624d37cbb7a7e831a16caa8dccaed1980b1e267cd1e6876647dbcc0fb5"></a>

<a id="canonical-fbd21a933bfe83bd6d1d5741ae2fb83287984ab596467c28c7775f3a52e3a189"></a>

## labels property — Property reference / 49c60d212bb9 / 8

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

<a id="canonical-e6a2b8f72caec702c7f3b8b7010e07c69cea601c4a45a9526dd57589d41d07ef"></a>

<a id="canonical-f8db7b20d33b69e6fcb24f933b1f53896f5db5a604e227daaacaa5acd594923e"></a>

## name property — Property reference / 49c60d212bb9 / 9

Type: `"string"`. Required.

Name of the Policy Based Routing. Must be unique within the namespace.

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

<a id="canonical-373a6b7f5558ff6012d5ac18d2d6fb1c7d3aceeeaa432c38bd1cc4d0923c2648"></a>

<a id="canonical-7ffa3f495e5ffa0e110fadb91a846a17859152a919d3626440f7c673624695c3"></a>

## namespace property — Property reference / 49c60d212bb9 / 10

Type: `"string"`. Required.

Namespace where the Policy Based Routing is created.

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

- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5): complete subsection reference.

- [timeouts](resources--policy_based_routing--reference--group-001.md#canonical-e6b0b7cdc591feff0e35c2994ef5aa4ca05a7f1daf4ce44cb331cdfbaf70222f): complete subsection reference.

<a id="canonical-c99ac8ee8a16b96f2f9151cd2560f84c04b521751f60e645d75f21fe3c500bbd"></a>

## All schema paths — Property reference / 49c60d212bb9 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--policy_based_routing--reference--group-001.md#canonical-24531696bdc7c839e9707d48f8de53db4ee2eb929f8458a1030ebe23b77cf030) |
| `description` | [description](resources--policy_based_routing--reference--group-001.md#canonical-51dc18726448fbd00f84eb57909cfacaf8e74a8906b4090d8ce6eac5917f1df5) |
| `disable` | [disable](resources--policy_based_routing--reference--group-001.md#canonical-13460366836cb23651ba98f3d920173e08dee854208cfd94a95f3df15e041e8f) |
| `forward_proxy_pbr` | [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-fdc913c5b28f6fc83e5367d877d36c0c4c95d0c025b9e9a813b01d99c7163571) |
| `forward_proxy_pbr.forward_proxy_pbr_rules` | [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-aceabb519bacda19bed38acb97b017ba0b2d315ca612a0a8a544f9485ad0b866) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](resources--policy_based_routing--reference--group-001.md#canonical-41f315bc8505e1781f936176c9bd605111a6755462fa431392097c3834bdb728) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.all_sources` | [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](resources--policy_based_routing--reference--group-001.md#canonical-bc272481d94b39087a67dab9a5a67dfcc2fd7f4d627a50c41155d99fd63d8172) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-3d8f74b4b43a580c51fef38d65a534783593998e65d4f0cb387cb20a5bbf4af0) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.name](resources--policy_based_routing--reference--group-001.md#canonical-18b695e6870ca4da4f8bcbc24e987f1ab8d8aa3ba1d77d7b6460d768e176a34c) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.namespace](resources--policy_based_routing--reference--group-001.md#canonical-c0b54358df05770993ac76fee97ed8dc0e44209595c8aff2f8cf973cf055d44f) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list.tenant](resources--policy_based_routing--reference--group-001.md#canonical-11dc29b9e5a111d0c6b810ec16bf8848ecd61321c9876adf721c0609ad368871) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-8362d01d4c6288c292909d718b62db009024dd63abfc100f11e88978748d949a) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](resources--policy_based_routing--reference--group-001.md#canonical-2e7a45512ec168be491e1523f44dec36b9f7415db8804869b74a7fbf7f5e2990) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](resources--policy_based_routing--reference--group-001.md#canonical-ed5ef77b24b9eb5163c0261b31ccc4c95e065d9121a9721a4facf40ea34ba597) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.exact_value](resources--policy_based_routing--reference--group-001.md#canonical-8427c2340d101d28a7fb13a723999bf44b40129dfc7aa46dac0375286b751a33) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_exact_value](resources--policy_based_routing--reference--group-001.md#canonical-c1d5d362e5f2016ce64cd8444dcd0637c99f09042ac0edf270cc1b511178d413) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_prefix_value](resources--policy_based_routing--reference--group-001.md#canonical-542f57834f3b5e96dea92098a4f09bdc4432ccc4e7c63885bda539ac947a6ce0) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.path_regex_value](resources--policy_based_routing--reference--group-001.md#canonical-0a3d19bd90924f96b8f8c3de3644bf655cd311f0e7e92d2fc649608bef9bf26d) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.regex_value](resources--policy_based_routing--reference--group-001.md#canonical-09d93ac86271ae99a224edae2c362a434f94e0d5a1e9f1eef33754458db94ed6) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.suffix_value](resources--policy_based_routing--reference--group-001.md#canonical-0cf1b12e84a63a9d4b849a2a64fa5aaf73cb7011e5284cfd7fedf814a91fbea5) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-48c3c1e7026f2132ec61850192813567d8fa7a713ca5787e0ec6ddaf5c681937) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.name](resources--policy_based_routing--reference--group-001.md#canonical-6340fea934ac80002805bcca306e99afd471be4b2d72890002719756c788d009) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.namespace](resources--policy_based_routing--reference--group-001.md#canonical-63707862ac197cd7c2c10962578dfa2cafb6fd4470a5224aeda8e9d387b96c6c) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant` | [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set.tenant](resources--policy_based_routing--reference--group-001.md#canonical-4c4e51215480441565f223f6ee3a2f5303670db4faeff775b3eb7ad092c44673) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](resources--policy_based_routing--reference--group-001.md#canonical-4de854d1c48c3e8c6a8b357c8b2ed85b15dfa78d34f842b9675335d70eee62a0) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions` | [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector.expressions](resources--policy_based_routing--reference--group-001.md#canonical-52a0d56df08b1c56d008476bb7aa8cb712bcf75096105dbaa71fff9cf91c37f0) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](resources--policy_based_routing--reference--group-001.md#canonical-3717acefc18fcda77d56bb13f1209378eeaae41a1cafb5b05bf42064f0611642) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.description_spec](resources--policy_based_routing--reference--group-001.md#canonical-800fd0350a3df47fa3ef5e9c5c371f922370242afd5648cdaf59806a466b19fa) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name` | [forward_proxy_pbr.forward_proxy_pbr_rules.metadata.name](resources--policy_based_routing--reference--group-001.md#canonical-8600cec6686bc8f08ea2a03a74057206b9d77339576bf550230b71f424da8b70) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-c63173681f275b877e540cc6589bf0973417e82ae0244cb3d073a97d1e24be50) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes` | [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list.prefixes](resources--policy_based_routing--reference--group-001.md#canonical-f12e56cf60a34b58945c9d4447889b84d9ffc6cb51aaf554465c54a3880b86b0) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-5989ada4e7f5877b274afc1af4175b383e254392f9610eee0c8c8d9f1048b81d) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-1871e5e8419007027dd5d262c7360239348bc4a059f0c20461efec494153ccdb) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.exact_value](resources--policy_based_routing--reference--group-001.md#canonical-c7b1ddffe93559e408ae197abf43df04e377f8db19acc143977749131aae4300) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.regex_value](resources--policy_based_routing--reference--group-001.md#canonical-39243fffdedfbb86fcaa43dad30acc627124f72adb9220b268dafb283f310418) |
| `forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value` | [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list.suffix_value](resources--policy_based_routing--reference--group-001.md#canonical-6a9f9b0a575b500283dfaf3ae2c38b5731139d018a9b731177da62f6dfdcb70d) |
| `forwarding_class_list` | [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-84db5b2fb7ca6792f867f1544155e55dfde34f0d0a3113c61b6c77d793c5da77) |
| `forwarding_class_list.name` | [forwarding_class_list.name](resources--policy_based_routing--reference--group-001.md#canonical-cb05d37a0a248cf6d6f1de9ab74a84116b91d078d336f1e406c257bc5aac830c) |
| `forwarding_class_list.namespace` | [forwarding_class_list.namespace](resources--policy_based_routing--reference--group-001.md#canonical-1ef17adc7499da58058667fd1e0636c106d2a45eb5df84521d7a5f55bfc2f3b3) |
| `forwarding_class_list.tenant` | [forwarding_class_list.tenant](resources--policy_based_routing--reference--group-001.md#canonical-d4bdbcea3b85e34b6eddae278d1e7e260b3f095b28317d50124c59c9044aae45) |
| `id` | [id](resources--policy_based_routing--reference--group-001.md#canonical-bba2e59ca6f40a52f459f08aa58ae7151cfc52129a1ba037cf4eeb1cf431f7d3) |
| `labels` | [labels](resources--policy_based_routing--reference--group-001.md#canonical-668614624d37cbb7a7e831a16caa8dccaed1980b1e267cd1e6876647dbcc0fb5) |
| `name` | [name](resources--policy_based_routing--reference--group-001.md#canonical-e6a2b8f72caec702c7f3b8b7010e07c69cea601c4a45a9526dd57589d41d07ef) |
| `namespace` | [namespace](resources--policy_based_routing--reference--group-001.md#canonical-373a6b7f5558ff6012d5ac18d2d6fb1c7d3aceeeaa432c38bd1cc4d0923c2648) |
| `network_pbr` | [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-61e77f1647a97b87b16e8cd5881378d133386bc11d93a00b3da7d5c381cdd067) |
| `network_pbr.any` | [network_pbr.any](resources--policy_based_routing--reference--group-001.md#canonical-c96ca785df867c597ccab56829c113d14d8552c6051b570ac5f5107dc4145d50) |
| `network_pbr.label_selector` | [network_pbr.label_selector](resources--policy_based_routing--reference--group-001.md#canonical-c6e8e1ef7822e121c94fa7fb503d902c55aa0019a3b6b8d2d9620768798d317b) |
| `network_pbr.label_selector.expressions` | [network_pbr.label_selector.expressions](resources--policy_based_routing--reference--group-001.md#canonical-d732f257f69ee8146f882031cd269f09170446010f908dc347382e40b52851e1) |
| `network_pbr.network_pbr_rules` | [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-89113ef586777010e760e956ec36b82930207a5bec9601a4d78bd8513cda45e7) |
| `network_pbr.network_pbr_rules.all_tcp_traffic` | [network_pbr.network_pbr_rules.all_tcp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-6fb96838392a6a9c7860db3c2d1f8e52f9f55abbdb08bbe094684bf4349778fa) |
| `network_pbr.network_pbr_rules.all_traffic` | [network_pbr.network_pbr_rules.all_traffic](resources--policy_based_routing--reference--group-001.md#canonical-a5cea3ffa9b81e59670201ce336eddb40c7b14a205f98e9e6cc48a15761dbf09) |
| `network_pbr.network_pbr_rules.all_udp_traffic` | [network_pbr.network_pbr_rules.all_udp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-b1e46191a6f4382de50ac72e5079e8ec1a97bbada1a9c59ec562bc9e7992146f) |
| `network_pbr.network_pbr_rules.any` | [network_pbr.network_pbr_rules.any](resources--policy_based_routing--reference--group-001.md#canonical-e425bbfedff900d302212916319c4e8790df122eb82237d42de8abdf75f579ee) |
| `network_pbr.network_pbr_rules.applications` | [network_pbr.network_pbr_rules.applications](resources--policy_based_routing--reference--group-001.md#canonical-0ec5c339bce94ae7ef3bbb1dd88269d2d8dbf02e5431107875b83225473d2cfa) |
| `network_pbr.network_pbr_rules.applications.applications` | [network_pbr.network_pbr_rules.applications.applications](resources--policy_based_routing--reference--group-001.md#canonical-b1ded9b50b4883ae2f508569d636fdc46dc59c806b0d14b1d4513ff0a14925b1) |
| `network_pbr.network_pbr_rules.dns_name` | [network_pbr.network_pbr_rules.dns_name](resources--policy_based_routing--reference--group-001.md#canonical-cad4270754c4faa6b9ddda799d516471c8458083195eff9b52e0624fa1ec833c) |
| `network_pbr.network_pbr_rules.forwarding_class_list` | [network_pbr.network_pbr_rules.forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-94f71d3c81a076d40939a4a40603cc9d56da224107d34b36ffb88b2c07157d23) |
| `network_pbr.network_pbr_rules.forwarding_class_list.name` | [network_pbr.network_pbr_rules.forwarding_class_list.name](resources--policy_based_routing--reference--group-001.md#canonical-f1b568d173cdecd76e35cf02b3a3699a4286a9a9a68220e21ae1b88a669847c5) |
| `network_pbr.network_pbr_rules.forwarding_class_list.namespace` | [network_pbr.network_pbr_rules.forwarding_class_list.namespace](resources--policy_based_routing--reference--group-001.md#canonical-ac7995c308b1e86dd8d92e615cecbd342aecf085a64c08517df8dc1e888211f8) |
| `network_pbr.network_pbr_rules.forwarding_class_list.tenant` | [network_pbr.network_pbr_rules.forwarding_class_list.tenant](resources--policy_based_routing--reference--group-001.md#canonical-d25cd9d21e28f7c36b183538e93f223829ab4467a79642223df719488db8092d) |
| `network_pbr.network_pbr_rules.ip_prefix_set` | [network_pbr.network_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-471538d8a75ca030df70e2640c22c62e9a78dcfa0675ee3c60a3df103cfb5c49) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref` | [network_pbr.network_pbr_rules.ip_prefix_set.ref](resources--policy_based_routing--reference--group-001.md#canonical-bbb01a4b896c925729dbfe1dc5f37d8fe92d905586bf339f49b4cb9dd9ca908e) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.kind` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.kind](resources--policy_based_routing--reference--group-001.md#canonical-3ebb1bea2ff97d13d64a98457694e561339c69c251258bb65ee6411dc83b0039) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.name` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.name](resources--policy_based_routing--reference--group-001.md#canonical-1eecc7dfef1443bf326bfb740dcf0e95acd9d81344af5c01499f80b58219cb66) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.namespace](resources--policy_based_routing--reference--group-001.md#canonical-f6562b7abd35ebbf32cc7a485ca8cf0d5a9e5675d141fbfd8fa493f771e63b06) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.tenant](resources--policy_based_routing--reference--group-001.md#canonical-626cb0aaa9984d87663dc05edb23e70f4bf276b5826baece433cf25680fa0daa) |
| `network_pbr.network_pbr_rules.ip_prefix_set.ref.uid` | [network_pbr.network_pbr_rules.ip_prefix_set.ref.uid](resources--policy_based_routing--reference--group-001.md#canonical-4a45da6501a72e3cf352c62153358df5a3ae478e625489ba76b61c16d66adf4c) |
| `network_pbr.network_pbr_rules.metadata` | [network_pbr.network_pbr_rules.metadata](resources--policy_based_routing--reference--group-001.md#canonical-591c39c3d0cebd72cf4f4609d815bc70d6dd4c50873d85e756937e1ebeb7e497) |
| `network_pbr.network_pbr_rules.metadata.description_spec` | [network_pbr.network_pbr_rules.metadata.description_spec](resources--policy_based_routing--reference--group-001.md#canonical-a26891fa5d379edd6ccabd64836f68edc32f919569cbd145b1b9cfb0eaebac2a) |
| `network_pbr.network_pbr_rules.metadata.name` | [network_pbr.network_pbr_rules.metadata.name](resources--policy_based_routing--reference--group-001.md#canonical-6292a390a10fdddf08dae9a537da47b2b1d282dc2cf01d233dcc4f89533f662d) |
| `network_pbr.network_pbr_rules.prefix_list` | [network_pbr.network_pbr_rules.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-982a4f5d95f6eeca4233018e616c5d0827e552a015af70a2e7f6cc8cc3afa032) |
| `network_pbr.network_pbr_rules.prefix_list.prefixes` | [network_pbr.network_pbr_rules.prefix_list.prefixes](resources--policy_based_routing--reference--group-001.md#canonical-c57cc1fe687fc4bb80a991f47f7513c810bc8e2c26ac866eb35de49eaf1f186c) |
| `network_pbr.network_pbr_rules.protocol_port_range` | [network_pbr.network_pbr_rules.protocol_port_range](resources--policy_based_routing--reference--group-001.md#canonical-75d2e1111e01a1a19e76efe480666f9ef0db6d8290b0f7b0e87e2af2ef339f98) |
| `network_pbr.network_pbr_rules.protocol_port_range.port_ranges` | [network_pbr.network_pbr_rules.protocol_port_range.port_ranges](resources--policy_based_routing--reference--group-001.md#canonical-234719337b1848cd5add3a0ca94cb8ea2679c008288465beaf8462aa3d97bc67) |
| `network_pbr.network_pbr_rules.protocol_port_range.protocol` | [network_pbr.network_pbr_rules.protocol_port_range.protocol](resources--policy_based_routing--reference--group-001.md#canonical-a62f5c6f5129c03792ed02fb9a379bc4978c1169a3850d2413484740f5345c9e) |
| `network_pbr.prefix_list` | [network_pbr.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-d2dea93b4e58534023680b4ca05c0011e80661409a584372b2210bdaba3f037d) |
| `network_pbr.prefix_list.prefixes` | [network_pbr.prefix_list.prefixes](resources--policy_based_routing--reference--group-001.md#canonical-ed26bffcc3d444ff499e909eeb701f25d5294b33229e4809d041bce9bb8ea878) |
| `timeouts` | [timeouts](resources--policy_based_routing--reference--group-001.md#canonical-bf50f8189e76663681d9b08046d04d8bf1380937f99a6dc0251c35a8f2999fa0) |
| `timeouts.create` | [timeouts.create](resources--policy_based_routing--reference--group-001.md#canonical-117951dac4a3f59ea9590b3ddb7869ea6bdbf70527286ed939ed27e4eeca692e) |
| `timeouts.delete` | [timeouts.delete](resources--policy_based_routing--reference--group-001.md#canonical-9c52b9d14fe622057ffa2bd26a1131b14df210e3d6f4b6a17616c006fed907a1) |
| `timeouts.read` | [timeouts.read](resources--policy_based_routing--reference--group-001.md#canonical-19442ca3398f627720942504c6c319bf34e4e3ad41da323da8553ef3492b553a) |
| `timeouts.update` | [timeouts.update](resources--policy_based_routing--reference--group-001.md#canonical-89448f7269c12549be8111939e879e9fa2e5e12d7f9e023f2241d3f0ed2b4429) |

<a id="canonical-3baa2deb683e6e26f5e65ce3b6de09ac5f56047fd93a1650d7227f28f15dfc56"></a>

## Next pages — Property reference / 49c60d212bb9 / 12

- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-19eaad8e0e20a025d92dc4e8e8ad8b28101e842052d23859924f7873998d59f0)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [timeouts](resources--policy_based_routing--reference--group-001.md#canonical-e6b0b7cdc591feff0e35c2994ef5aa4ca05a7f1daf4ce44cb331cdfbaf70222f)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f7608c1d62b5c55d88fef84760ea25b6d45741cbd60215eb24f196f2377a584"></a>

## forward_proxy_pbr — forward_proxy_pbr / 0fcf13bd4509 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- forward_proxy_pbr

<a id="canonical-fdc913c5b28f6fc83e5367d877d36c0c4c95d0c025b9e9a813b01d99c7163571"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: forward\_proxy\_pbr, network\_pbr\] Configuration parameter for forward proxy pbr.

Upstream description:

Network(L3/L4) routing policy rule.

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

- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-fdc913c5b28f6fc83e5367d877d36c0c4c95d0c025b9e9a813b01d99c7163571)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-61e77f1647a97b87b16e8cd5881378d133386bc11d93a00b3da7d5c381cdd067)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
forward_proxy_pbr {
  # Configure direct properties listed below.
}
```

<a id="canonical-29880ab11f8f8186a3c4f77c7ef356c664cdb43e382de51cf444c78ed5504c88"></a>

## Direct properties — forward_proxy_pbr / 0fcf13bd4509 / 3

- [forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976): complete subsection reference.

<a id="canonical-d4e3eb5d677573ddd73057091a796f9a0a917ddacb84ba1c83f2cdbe1de6b95d"></a>

## Next pages — forward_proxy_pbr / 0fcf13bd4509 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-581779fcd7057138110e99f900cdd3ef98ea1ac34461bcc2fc841ee97cf89c1d"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules — forward_proxy_pbr.forward_proxy_pbr_rules / 784111eb2fd2 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="canonical-aceabb519bacda19bed38acb97b017ba0b2d315ca612a0a8a544f9485ad0b866"></a>

Type: `"object"`. list nested block, Optional.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

Upstream description:

Network(L3/L4) routing policy rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("forwarding_class_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "http_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "tls_list"),
  validators.ConflictingListObjectAttributes("all_sources",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
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
forward_proxy_pbr_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-89c3c08200395a7155a1f3253a2a980f66783a15e06f73f8ff228b384c193149"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules / 784111eb2fd2 / 3

- [all_destinations](resources--policy_based_routing--reference--group-001.md#canonical-83b2f8de353ae96ecc27194f42fc63bf06a088b434e3d5f6bcf352cd23bb5c28): complete subsection reference.

- [all_sources](resources--policy_based_routing--reference--group-001.md#canonical-c807ce852225feba8ef81124a373fff88ed52e6c2085237004f3b801a5f36c54): complete subsection reference.

- [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-e20e578bb8a952910a122161215d8973c7b1a688c5af7eba637174a4e6746f24): complete subsection reference.

- [http_list](resources--policy_based_routing--reference--group-001.md#canonical-e4b8025dc951bf93ec31246b149fb6eb732478ec40ce0560b77ce9a0790677e9): complete subsection reference.

- [ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-5a11d72f003b591aceceafab969d4c26f04b6668031bdd409f04854ee9a822c8): complete subsection reference.

- [label_selector](resources--policy_based_routing--reference--group-001.md#canonical-bd903591a97928cd06bcff469c6a1b05d6821b727af64811d40e400600586910): complete subsection reference.

- [metadata](resources--policy_based_routing--reference--group-001.md#canonical-6fec9c9800d071a0285425954231027a4bf03934a96ef88eba5e6003d3f855fc): complete subsection reference.

- [prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-b69321dbe5fdbcb336a8e25d29b0d8c740a824613c2d991abca08321e364a889): complete subsection reference.

- [tls_list](resources--policy_based_routing--reference--group-001.md#canonical-b202b7be18666aadd2cc7e096f4e2b57a06de48c9447531e3d5512e4757ce9cc): complete subsection reference.

<a id="canonical-36a99885a2f1a9a9f060b40497048b7075823219aad3b1649866f07bd1940d76"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules / 784111eb2fd2 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](resources--policy_based_routing--reference--group-001.md#canonical-83b2f8de353ae96ecc27194f42fc63bf06a088b434e3d5f6bcf352cd23bb5c28)
- [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](resources--policy_based_routing--reference--group-001.md#canonical-c807ce852225feba8ef81124a373fff88ed52e6c2085237004f3b801a5f36c54)
- [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-e20e578bb8a952910a122161215d8973c7b1a688c5af7eba637174a4e6746f24)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-e4b8025dc951bf93ec31246b149fb6eb732478ec40ce0560b77ce9a0790677e9)
- [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-5a11d72f003b591aceceafab969d4c26f04b6668031bdd409f04854ee9a822c8)
- [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](resources--policy_based_routing--reference--group-001.md#canonical-bd903591a97928cd06bcff469c6a1b05d6821b727af64811d40e400600586910)
- [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](resources--policy_based_routing--reference--group-001.md#canonical-6fec9c9800d071a0285425954231027a4bf03934a96ef88eba5e6003d3f855fc)
- [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-b69321dbe5fdbcb336a8e25d29b0d8c740a824613c2d991abca08321e364a889)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-b202b7be18666aadd2cc7e096f4e2b57a06de48c9447531e3d5512e4757ce9cc)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-83b2f8de353ae96ecc27194f42fc63bf06a088b434e3d5f6bcf352cd23bb5c28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b3ca6781aa378242f056db39bc085622c8520fe8288c69036b7f16bb5a37e37"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations — forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations / 15e932ce6502 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations

<a id="canonical-41f315bc8505e1781f936176c9bd605111a6755462fa431392097c3834bdb728"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

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
all_destinations = {}
```

<a id="canonical-325ea597d78ecf35fbedb07300cbaef2c703073c993b4984b876833d1751089f"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations / 15e932ce6502 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-27486c095377b63bd732e13bef75e8a9d638ac7b0f7316f84539a355dba7be94"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations / 15e932ce6502 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-c807ce852225feba8ef81124a373fff88ed52e6c2085237004f3b801a5f36c54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5485b9561dba9b887ce092c48b29efd175d1fd344470c4925f832d69c10955fe"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.all_sources — forward_proxy_pbr.forward_proxy_pbr_rules.all_sources / e106f98191ca / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.all_sources

<a id="canonical-bc272481d94b39087a67dab9a5a67dfcc2fd7f4d627a50c41155d99fd63d8172"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all sources.

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
all_sources = {}
```

<a id="canonical-2888ff397b80e112435b084f1b7dc841d0ee3f6a8a7ac0ec5211231b5cc0471a"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.all_sources / e106f98191ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25c76c2eab8ec6691f0dc65733141c2c42d4b77ef6736d085f680e3c7219b054"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.all_sources / e106f98191ca / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-e20e578bb8a952910a122161215d8973c7b1a688c5af7eba637174a4e6746f24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a359bd268707a128a4d9b08186e9aaa427b05f1bc18c4de522bc46fb83a4d111"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / fcaac0b76b30 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list

<a id="canonical-3d8f74b4b43a580c51fef38d65a534783593998e65d4f0cb387cb20a5bbf4af0"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of forwarding Class to be used if no rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forwarding_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-15312d70cf526df561f4e9b429404fd2c28e628c5a4a2f3abf024ba272910432"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / fcaac0b76b30 / 3

<a id="canonical-18b695e6870ca4da4f8bcbc24e987f1ab8d8aa3ba1d77d7b6460d768e176a34c"></a>

<a id="canonical-9c361fb85a7b6e24e5ab916070aa0d60c435bd384751a63774581795341e0cd4"></a>

## name property — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / fcaac0b76b30 / 4

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

<a id="canonical-c0b54358df05770993ac76fee97ed8dc0e44209595c8aff2f8cf973cf055d44f"></a>

<a id="canonical-e642858669e52d14ed7335a6930f9447d4d64d004741d411f69d4654b41a0af7"></a>

## namespace property — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / fcaac0b76b30 / 5

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

<a id="canonical-11dc29b9e5a111d0c6b810ec16bf8848ecd61321c9876adf721c0609ad368871"></a>

<a id="canonical-cdbc02429266b07813d379bdc57afeea8c0e4f9e269a750f458909f7a0d06af1"></a>

## tenant property — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / fcaac0b76b30 / 6

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

<a id="canonical-f1bb8cd030c5ace1013eb02eeafda0505ff5f1202a122a14402bf6282d6dc722"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list / fcaac0b76b30 / 7

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-e4b8025dc951bf93ec31246b149fb6eb732478ec40ce0560b77ce9a0790677e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2da87b54f54e75a5a1b61d822c13c984c67731802db21cd8cf9b29bb584acd20"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.http_list — forward_proxy_pbr.forward_proxy_pbr_rules.http_list / 3cf7820f51c6 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list

<a id="canonical-8362d01d4c6288c292909d718b62db009024dd63abfc100f11e88978748d949a"></a>

Type: `"object"`. single nested block, Optional.

URLListType.

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
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d3b2c4f271840cfd1526e8457071eb4dd2d6ee3355cf13c73639813f510e813"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.http_list / 3cf7820f51c6 / 3

- [http_list](resources--policy_based_routing--reference--group-001.md#canonical-5b9c42c0fb1c0ca558421131712d4df34f354c27204635f302cc4f2dbf408b55): complete subsection reference.

<a id="canonical-4f66d0d74ab43cf86d739a7ea4a8a22591d9ac7f8518eb0e6b70c356281a4d41"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.http_list / 3cf7820f51c6 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](resources--policy_based_routing--reference--group-001.md#canonical-5b9c42c0fb1c0ca558421131712d4df34f354c27204635f302cc4f2dbf408b55)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-5b9c42c0fb1c0ca558421131712d4df34f354c27204635f302cc4f2dbf408b55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f002b1b62502083c8f80b1c35516895e471f133c128b798a9c50d87403e7728d"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-e4b8025dc951bf93ec31246b149fb6eb732478ec40ce0560b77ce9a0790677e9)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list

<a id="canonical-2e7a45512ec168be491e1523f44dec36b9f7415db8804869b74a7fbf7f5e2990"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-7c2af2e84224d5b8a6f2de33a03e10505bf8fc16740cee2478844d2abcc0a9d9"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 3

- [any_path](resources--policy_based_routing--reference--group-001.md#canonical-835f8351f4b7a17310d8bb44bdc7b59642cb2461f4bd19883fd7dfafef451f48): complete subsection reference.

<a id="canonical-8427c2340d101d28a7fb13a723999bf44b40129dfc7aa46dac0375286b751a33"></a>

<a id="canonical-54ffcfe4daac90536403652f9716b5c6c3f39177e6c7b8fd2bfe32f58cdd6b07"></a>

## exact_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c1d5d362e5f2016ce64cd8444dcd0637c99f09042ac0edf270cc1b511178d413"></a>

<a id="canonical-9673629fc14a1086e1cfd9f2f332eea650e97f07b32e80c5d85d6303d060d7c8"></a>

## path_exact_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 5

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-542f57834f3b5e96dea92098a4f09bdc4432ccc4e7c63885bda539ac947a6ce0"></a>

<a id="canonical-d343b63f1d838bdef9248bba61438d47de306b110009abd90cdc8344b2cd064e"></a>

## path_prefix_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 6

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0a3d19bd90924f96b8f8c3de3644bf655cd311f0e7e92d2fc649608bef9bf26d"></a>

<a id="canonical-e9a2a6c3b8c14606d86b24a54a1fbc727a5582670f10f59a78292e79661e5206"></a>

## path_regex_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-09d93ac86271ae99a224edae2c362a434f94e0d5a1e9f1eef33754458db94ed6"></a>

<a id="canonical-6f8eecfe2dd968af98e4be5e8193552fc43e5b1f8997fcfde553aaec2fb51e74"></a>

## regex_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 8

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0cf1b12e84a63a9d4b849a2a64fa5aaf73cb7011e5284cfd7fedf814a91fbea5"></a>

<a id="canonical-c861d76c1cd95540a7e9b13e347378b0f866b939557e74e0c00661c1b73dc874"></a>

## suffix_value property — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 9

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-07790fc7805e125295b14aa8c2da9598b7353350e833337f093ae74591ac0aa2"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list / bf9db8fb95af / 10

- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path](resources--policy_based_routing--reference--group-001.md#canonical-835f8351f4b7a17310d8bb44bdc7b59642cb2461f4bd19883fd7dfafef451f48)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-e4b8025dc951bf93ec31246b149fb6eb732478ec40ce0560b77ce9a0790677e9)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-835f8351f4b7a17310d8bb44bdc7b59642cb2461f4bd19883fd7dfafef451f48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ded0f18a190ac399fb68a17f4057ad9c661a2cbea5367adc0d59e15e285bd3c"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path / 2444c93c9536 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](resources--policy_based_routing--reference--group-001.md#canonical-e4b8025dc951bf93ec31246b149fb6eb732478ec40ce0560b77ce9a0790677e9)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](resources--policy_based_routing--reference--group-001.md#canonical-5b9c42c0fb1c0ca558421131712d4df34f354c27204635f302cc4f2dbf408b55)
- forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path

<a id="canonical-ed5ef77b24b9eb5163c0261b31ccc4c95e065d9121a9721a4facf40ea34ba597"></a>

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
any_path = {}
```

<a id="canonical-b16662902dde9e52f3b9538355713ffe163182da956f49b701248c203983d0e5"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path / 2444c93c9536 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c8d0080192cd98bc19de77dfa42e1c606f0ace4c98d6327c1c8ad929370827fc"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list.any_path / 2444c93c9536 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list.http_list](resources--policy_based_routing--reference--group-001.md#canonical-5b9c42c0fb1c0ca558421131712d4df34f354c27204635f302cc4f2dbf408b55)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-5a11d72f003b591aceceafab969d4c26f04b6668031bdd409f04854ee9a822c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d87e2dc04f58e1c3ea183ade3d127fdd78c7207c651cf2fb304a025274fc62e"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 5f8d381f6dcb / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set

<a id="canonical-48c3c1e7026f2132ec61850192813567d8fa7a713ca5787e0ec6ddaf5c681937"></a>

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-91fb43d04954ea2743af368d1a2d1bcfb36c6364e3faf93de6587acdcbf3e9a3"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 5f8d381f6dcb / 3

<a id="canonical-6340fea934ac80002805bcca306e99afd471be4b2d72890002719756c788d009"></a>

<a id="canonical-c1a776bc780cf142f89f80d0993311adb4389fabf9603ed8c84061c247b0ffc4"></a>

## name property — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 5f8d381f6dcb / 4

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

<a id="canonical-63707862ac197cd7c2c10962578dfa2cafb6fd4470a5224aeda8e9d387b96c6c"></a>

<a id="canonical-a6266eeb5a16c5162b614cbd0c801f48e06f065fb6573ef9ff68e29584ce0a36"></a>

## namespace property — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 5f8d381f6dcb / 5

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

<a id="canonical-4c4e51215480441565f223f6ee3a2f5303670db4faeff775b3eb7ad092c44673"></a>

<a id="canonical-e84d6859f5960dc600ce85f95840235bf3ca3d666f82e759326529c2faa17053"></a>

## tenant property — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 5f8d381f6dcb / 6

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

<a id="canonical-44808fed8c9b95f443e2d687d7e2ec5f09ca3a28b58b8031b899340082913082"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set / 5f8d381f6dcb / 7

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-bd903591a97928cd06bcff469c6a1b05d6821b727af64811d40e400600586910"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4d3d49792e311868dc77f8850782c9043c83f3a9207b21554b488ed492bf88f"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.label_selector — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 8b6a7bee8754 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.label_selector

<a id="canonical-4de854d1c48c3e8c6a8b357c8b2ed85b15dfa78d34f842b9675335d70eee62a0"></a>

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

<a id="canonical-f35153ac205c24e4d6b7da13c3fa46abfe564fb346d42eef3925b59442d561af"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 8b6a7bee8754 / 3

<a id="canonical-52a0d56df08b1c56d008476bb7aa8cb712bcf75096105dbaa71fff9cf91c37f0"></a>

<a id="canonical-f2cfa0b72d1ec48536c5b58aff1124f97485db2a3c83cfe90409b2b8f7d37d35"></a>

## expressions property — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 8b6a7bee8754 / 4

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

<a id="canonical-e63cefa23a77a9aa6e64700c04ed259fcdb9cc0f1a5a7bb4f34aee559905ae4f"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.label_selector / 8b6a7bee8754 / 5

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-6fec9c9800d071a0285425954231027a4bf03934a96ef88eba5e6003d3f855fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6dac0cb07c64495d3cecf46bc3814e4cdc9bb1a195f7a399a1ca19b4f9358112"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.metadata — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 75be47335152 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.metadata

<a id="canonical-3717acefc18fcda77d56bb13f1209378eeaae41a1cafb5b05bf42064f0611642"></a>

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

<a id="canonical-10df771b8a69d1ef008ff4c03da0dfe0bc8734ff50625d8ad2294acdba288572"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 75be47335152 / 3

<a id="canonical-800fd0350a3df47fa3ef5e9c5c371f922370242afd5648cdaf59806a466b19fa"></a>

<a id="canonical-c620265cd5628d5c1903289ecd4578b80d6cb0441673d8893b4a09c2f4955592"></a>

## description_spec property — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 75be47335152 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-8600cec6686bc8f08ea2a03a74057206b9d77339576bf550230b71f424da8b70"></a>

<a id="canonical-061e1fe47a4810801b1950b082319e129f2534a9f424bb5f873b62c8e598d301"></a>

## name property — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 75be47335152 / 5

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

<a id="canonical-886687f17550be40c401c9a6c80c94ca377e9cd613c3e2e26f1e1d11aafb9f39"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.metadata / 75be47335152 / 6

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-b69321dbe5fdbcb336a8e25d29b0d8c740a824613c2d991abca08321e364a889"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33d47ce653a0ae7fe35e40a012ad68f7c8e1ff15f3b3be61df1a2e3eef1fcfc1"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / 2491bc962e29 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list

<a id="canonical-c63173681f275b877e540cc6589bf0973417e82ae0244cb3d073a97d1e24be50"></a>

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

<a id="canonical-5049fbd1e27cf968edede35ed2f137eccbc4195ce829bf62d94ead6caef2cc24"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / 2491bc962e29 / 3

<a id="canonical-f12e56cf60a34b58945c9d4447889b84d9ffc6cb51aaf554465c54a3880b86b0"></a>

<a id="canonical-2e4c36dd4060e780067106b4282b71ea92f83393d4890b0e7dd8ae79b2412c42"></a>

## prefixes property — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / 2491bc962e29 / 4

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

<a id="canonical-78d447c15e53a25884d34f8f9e18314f0e89e011931d35c7305c8cb43559bb37"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list / 2491bc962e29 / 5

- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-b202b7be18666aadd2cc7e096f4e2b57a06de48c9447531e3d5512e4757ce9cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02aa1ca54e66c15428c5553f4b2bc9c39b3213832ca8dab997de2841592f7956"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.tls_list — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list / 2d1b6f675b33 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list

<a id="canonical-5989ada4e7f5877b274afc1af4175b383e254392f9610eee0c8c8d9f1048b81d"></a>

Type: `"object"`. single nested block, Optional.

DomainListType.

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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea5e3f50d783fa84d3c9af9c7f1aacb575d7b1013e9bd37c97dbc672357f97cc"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list / 2d1b6f675b33 / 3

- [tls_list](resources--policy_based_routing--reference--group-001.md#canonical-dcc8de322abe28c0d9c9c012f8a24817ec00d92dab0e747d3c4ef6f80c9eff7a): complete subsection reference.

<a id="canonical-2e685bda8a1a773f5acf3fff981dc517e72bfedd32f8f0f550dd4c8b0cd042bb"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list / 2d1b6f675b33 / 4

- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-dcc8de322abe28c0d9c9c012f8a24817ec00d92dab0e747d3c4ef6f80c9eff7a)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-dcc8de322abe28c0d9c9c012f8a24817ec00d92dab0e747d3c4ef6f80c9eff7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-321bf02469db638c809bd908d9bf54bf428b040c902dce6181e7d3283093fdfc"></a>

## forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / 2acb920c10b7 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [forward_proxy_pbr](resources--policy_based_routing--reference--group-001.md#canonical-a8fa6450016fd088e7c571b6318cd4047b936b67aa7dc9fa7b12f82fa0022328)
- [forward_proxy_pbr.forward_proxy_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-4b2bd07791c16e99795d713e6597bc3070a2f0207a57c0698b0266e66a07b976)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-b202b7be18666aadd2cc7e096f4e2b57a06de48c9447531e3d5512e4757ce9cc)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list

<a id="canonical-1871e5e8419007027dd5d262c7360239348bc4a059f0c20461efec494153ccdb"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-8fa304b4a67dd07f727a5dc9d4eec494188b8393a8b5dc9244a37ae31d8c1144"></a>

## Direct properties — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / 2acb920c10b7 / 3

<a id="canonical-c7b1ddffe93559e408ae197abf43df04e377f8db19acc143977749131aae4300"></a>

<a id="canonical-9c24b665ab500b3cdd749ce947556c7286954ffcf918d4e65f243488cdc72209"></a>

## exact_value property — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / 2acb920c10b7 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-39243fffdedfbb86fcaa43dad30acc627124f72adb9220b268dafb283f310418"></a>

<a id="canonical-8c93ee2cc1edca3db9e1cf7c357e4a928dbf85f17d7ca842e57ffd79a6312a3a"></a>

## regex_value property — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / 2acb920c10b7 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-6a9f9b0a575b500283dfaf3ae2c38b5731139d018a9b731177da62f6dfdcb70d"></a>

<a id="canonical-3c1e9660f1952cb0e6bf449f5d04718a0d86681d933c13f66faf83e564d756d0"></a>

## suffix_value property — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / 2acb920c10b7 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b3aed09940b82e8efbb3eba84046f206857dba19ffd135ed3d846f3aea3885db"></a>

## Next pages — forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list / 2acb920c10b7 / 7

- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](resources--policy_based_routing--reference--group-001.md#canonical-b202b7be18666aadd2cc7e096f4e2b57a06de48c9447531e3d5512e4757ce9cc)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-19eaad8e0e20a025d92dc4e8e8ad8b28101e842052d23859924f7873998d59f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b06a4e7d20b5ff04bd9ec4c4f29cb7f0c1283480205574e554fb47ad4b39ebc3"></a>

## forwarding_class_list — forwarding_class_list / 1608dc2c8676 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- forwarding_class_list

<a id="canonical-84db5b2fb7ca6792f867f1544155e55dfde34f0d0a3113c61b6c77d793c5da77"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of forwarding Class to be used if source application match and no rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forwarding_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-834200f43f551013b5bf1d4c3ab2a913cd95cf250400997413681a9cc316b838"></a>

## Direct properties — forwarding_class_list / 1608dc2c8676 / 3

<a id="canonical-cb05d37a0a248cf6d6f1de9ab74a84116b91d078d336f1e406c257bc5aac830c"></a>

<a id="canonical-c8c4e538e823a09a5fc16f36f57557dda5e96677d8ff694eeda81d72d8fb9747"></a>

## name property — forwarding_class_list / 1608dc2c8676 / 4

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

<a id="canonical-1ef17adc7499da58058667fd1e0636c106d2a45eb5df84521d7a5f55bfc2f3b3"></a>

<a id="canonical-6538a122c50e3b62b785a5706c81222e1e62ac4b92340834bf025879f0e2c39e"></a>

## namespace property — forwarding_class_list / 1608dc2c8676 / 5

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

<a id="canonical-d4bdbcea3b85e34b6eddae278d1e7e260b3f095b28317d50124c59c9044aae45"></a>

<a id="canonical-1ccd3780086c82a1eeacf4a6dc80817c65735ae44ae858d939fcc19471c5052b"></a>

## tenant property — forwarding_class_list / 1608dc2c8676 / 6

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

<a id="canonical-21290541eaca2e9957d708ff8097047f6143c15e5f2600d7990816f72b969b2d"></a>

## Next pages — forwarding_class_list / 1608dc2c8676 / 7

- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98f864828a829c2676a5c3091067603a7b3bf8755720eb7a0374132b8cdb1307"></a>

## network_pbr — network_pbr / defdabf6db92 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- network_pbr

<a id="canonical-61e77f1647a97b87b16e8cd5881378d133386bc11d93a00b3da7d5c381cdd067"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for network pbr.

Upstream description:

Network(L3/L4) routing policy rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
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
  "x-ves-oneof-field-source_choice": "[\"any\",\"label_selector\",\"prefix_list\"]"
}
```

Terraform syntax:

```terraform
network_pbr {
  # Configure direct properties listed below.
}
```

<a id="canonical-f24cca6119078b56591b42a4a3385fe4d1a20d2c1313a45c2037577d97e5dfa4"></a>

## Direct properties — network_pbr / defdabf6db92 / 3

- [any](resources--policy_based_routing--reference--group-001.md#canonical-147a498a3be14c859cc8c172f85117cabb456c9ac767b1d76bf4eb7c4a1799b8): complete subsection reference.

- [label_selector](resources--policy_based_routing--reference--group-001.md#canonical-e36c0b356ddfe4de2a43fcd862640e7338eb98c6bcf0b889e769112918842d37): complete subsection reference.

- [network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0): complete subsection reference.

- [prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-d19ae0c9bb95b5010653fabd6116db2d9235b7817e72fcb0ac948dcc3eb8d621): complete subsection reference.

<a id="canonical-a902c8fc1337076a22a707e432925b0bb87b5e7aa82bd8b13bdf109e258f01fa"></a>

## Next pages — network_pbr / defdabf6db92 / 4

- [network_pbr.any](resources--policy_based_routing--reference--group-001.md#canonical-147a498a3be14c859cc8c172f85117cabb456c9ac767b1d76bf4eb7c4a1799b8)
- [network_pbr.label_selector](resources--policy_based_routing--reference--group-001.md#canonical-e36c0b356ddfe4de2a43fcd862640e7338eb98c6bcf0b889e769112918842d37)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [network_pbr.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-d19ae0c9bb95b5010653fabd6116db2d9235b7817e72fcb0ac948dcc3eb8d621)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-147a498a3be14c859cc8c172f85117cabb456c9ac767b1d76bf4eb7c4a1799b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e5b34bd8fe50d428908bd3cc79a81f097cf5690eb2eb60a96b56c48b0c41b94"></a>

## network_pbr.any — network_pbr.any / c80896389743 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- network_pbr.any

<a id="canonical-c96ca785df867c597ccab56829c113d14d8552c6051b570ac5f5107dc4145d50"></a>

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

<a id="canonical-5eb56442a74aeaf60da79a4744e24e1cc63e8d124458719053ec50bfeba544e3"></a>

## Direct properties — network_pbr.any / c80896389743 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-56e58da80653d7673012571bcf861ab2babb643c0da61288caedf683ae420ac8"></a>

## Next pages — network_pbr.any / c80896389743 / 4

- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-e36c0b356ddfe4de2a43fcd862640e7338eb98c6bcf0b889e769112918842d37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05c3fdcc8cc2b37cbdc7aa8c330faba544b56951ccb14c3c11b00f8fd4f3b3c6"></a>

## network_pbr.label_selector — network_pbr.label_selector / 85f615206719 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- network_pbr.label_selector

<a id="canonical-c6e8e1ef7822e121c94fa7fb503d902c55aa0019a3b6b8d2d9620768798d317b"></a>

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

<a id="canonical-4ece9c70d808adab681173755b17c9919b553a993adbe4a11775ae1979b7bee7"></a>

## Direct properties — network_pbr.label_selector / 85f615206719 / 3

<a id="canonical-d732f257f69ee8146f882031cd269f09170446010f908dc347382e40b52851e1"></a>

<a id="canonical-c10e3703b54260864ead513f7dce77ba25b20a180c9c2366e679f4ee30803654"></a>

## expressions property — network_pbr.label_selector / 85f615206719 / 4

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

<a id="canonical-557fcc44146f50a5f05b526ff52d14d01ab5352746fecb90ebf0ac4515ffb106"></a>

## Next pages — network_pbr.label_selector / 85f615206719 / 5

- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ce3174e190c54b35f75df55f31c9d4d2329bdc7ad8eea57d6ab716b650c3e7a"></a>

## network_pbr.network_pbr_rules — network_pbr.network_pbr_rules / 83acd4f8a632 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- network_pbr.network_pbr_rules

<a id="canonical-89113ef586777010e760e956ec36b82930207a5bec9601a4d78bd8513cda45e7"></a>

Type: `"object"`. list nested block, Optional.

L3/L4 Destination Routing Rules. Network(L3/L4) routing policy rule.

Upstream description:

Network(L3/L4) routing policy rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("forwarding_class_list"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
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
    "dns_name"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("dns_name",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dns_name",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
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
network_pbr_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-8de1971430ce1d9438b7f6ae6dc204f0cc5c82959b049a21f3752699f60fc845"></a>

## Direct properties — network_pbr.network_pbr_rules / 83acd4f8a632 / 3

- [all_tcp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-59ace201b2c434d968390a5e0995751a0821933c1e426e5a5c344b30ef6fe0fd): complete subsection reference.

- [all_traffic](resources--policy_based_routing--reference--group-001.md#canonical-d81526fb2b5ed6a9ddfd26abe345b224415ae35c33b76085c6e75930dde0cfae): complete subsection reference.

- [all_udp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-327cebf933308484a21cc5dc090c715c15f4a5f35d61057a4caa26d745d4a27d): complete subsection reference.

- [any](resources--policy_based_routing--reference--group-001.md#canonical-0789ab10d71bb49fe3bfbf0e0f9221b0103b1f5a70587bbc5b5ff65ec4d18cd1): complete subsection reference.

- [applications](resources--policy_based_routing--reference--group-001.md#canonical-0eccdfbe476bb68c4ce8f8b4696e44ca09407edaef6d2abc28e5756517ba193c): complete subsection reference.

<a id="canonical-cad4270754c4faa6b9ddda799d516471c8458083195eff9b52e0624fa1ec833c"></a>

<a id="canonical-5f45aa64a9b4ad95f1f13c7550ebd3c88bab848be5f5d7170e0a5852f61afc5a"></a>

## dns_name property — network_pbr.network_pbr_rules / 83acd4f8a632 / 4

Type: `"string"`. Optional.

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Upstream description:

Exclusive with \[any ip\_prefix\_set prefix\_list\] Resolve hostname to GET the IP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-023fd91ef592f0861bdc27e2568f4bea117f14e9ae91afd0ccb0fcc6b4d9ef38): complete subsection reference.

- [ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-28c5523a277bf151cca104ebfff87e2324e74107899f4a97b578761a3bead976): complete subsection reference.

- [metadata](resources--policy_based_routing--reference--group-001.md#canonical-0c93a1918183d0c4e61929c375ddc25b7c500d730b883ef55c0edcb5ccc7308d): complete subsection reference.

- [prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-9e7a0506ca13e72959a49dca526cd737bf77c253b260941b330118d95c83ad7e): complete subsection reference.

- [protocol_port_range](resources--policy_based_routing--reference--group-001.md#canonical-58a7d1612de059d29372c9a107bb7bd930e27c8f95a55c36f6bb37d194f82b92): complete subsection reference.

<a id="canonical-05b0d08e88869a019f662a5e43047906f25ccb6d8b1f762773e2197e05709f8d"></a>

## Next pages — network_pbr.network_pbr_rules / 83acd4f8a632 / 5

- [network_pbr.network_pbr_rules.all_tcp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-59ace201b2c434d968390a5e0995751a0821933c1e426e5a5c344b30ef6fe0fd)
- [network_pbr.network_pbr_rules.all_traffic](resources--policy_based_routing--reference--group-001.md#canonical-d81526fb2b5ed6a9ddfd26abe345b224415ae35c33b76085c6e75930dde0cfae)
- [network_pbr.network_pbr_rules.all_udp_traffic](resources--policy_based_routing--reference--group-001.md#canonical-327cebf933308484a21cc5dc090c715c15f4a5f35d61057a4caa26d745d4a27d)
- [network_pbr.network_pbr_rules.any](resources--policy_based_routing--reference--group-001.md#canonical-0789ab10d71bb49fe3bfbf0e0f9221b0103b1f5a70587bbc5b5ff65ec4d18cd1)
- [network_pbr.network_pbr_rules.applications](resources--policy_based_routing--reference--group-001.md#canonical-0eccdfbe476bb68c4ce8f8b4696e44ca09407edaef6d2abc28e5756517ba193c)
- [network_pbr.network_pbr_rules.forwarding_class_list](resources--policy_based_routing--reference--group-001.md#canonical-023fd91ef592f0861bdc27e2568f4bea117f14e9ae91afd0ccb0fcc6b4d9ef38)
- [network_pbr.network_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-28c5523a277bf151cca104ebfff87e2324e74107899f4a97b578761a3bead976)
- [network_pbr.network_pbr_rules.metadata](resources--policy_based_routing--reference--group-001.md#canonical-0c93a1918183d0c4e61929c375ddc25b7c500d730b883ef55c0edcb5ccc7308d)
- [network_pbr.network_pbr_rules.prefix_list](resources--policy_based_routing--reference--group-001.md#canonical-9e7a0506ca13e72959a49dca526cd737bf77c253b260941b330118d95c83ad7e)
- [network_pbr.network_pbr_rules.protocol_port_range](resources--policy_based_routing--reference--group-001.md#canonical-58a7d1612de059d29372c9a107bb7bd930e27c8f95a55c36f6bb37d194f82b92)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-59ace201b2c434d968390a5e0995751a0821933c1e426e5a5c344b30ef6fe0fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b81c11f746109adf3875fc33e983aa43f0edfea07fe02c253875f3983facc31"></a>

## network_pbr.network_pbr_rules.all_tcp_traffic — network_pbr.network_pbr_rules.all_tcp_traffic / 89a32d912aff / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.all_tcp_traffic

<a id="canonical-6fb96838392a6a9c7860db3c2d1f8e52f9f55abbdb08bbe094684bf4349778fa"></a>

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

<a id="canonical-3bbe4ae3de042472358e41a50be899d6502cc4f4c92edfc694dfa054f7bfbf66"></a>

## Direct properties — network_pbr.network_pbr_rules.all_tcp_traffic / 89a32d912aff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8139a4b53c27b8a56705aee4a4a11a08d3a9ad02b0cecb198f386d272b59ee21"></a>

## Next pages — network_pbr.network_pbr_rules.all_tcp_traffic / 89a32d912aff / 4

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-d81526fb2b5ed6a9ddfd26abe345b224415ae35c33b76085c6e75930dde0cfae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e88ad6431d465541c8d970e4002466ccfc8d28cb0df343c471ceab663d765f7b"></a>

## network_pbr.network_pbr_rules.all_traffic — network_pbr.network_pbr_rules.all_traffic / 9486f73f6d0e / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.all_traffic

<a id="canonical-a5cea3ffa9b81e59670201ce336eddb40c7b14a205f98e9e6cc48a15761dbf09"></a>

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

<a id="canonical-5bec8c63c7a9724fa4961c3695b33abbcc6030ad4080482de4525a655bd9a24a"></a>

## Direct properties — network_pbr.network_pbr_rules.all_traffic / 9486f73f6d0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6160cbd5ddff4b80ebd3f9437516583b67465e168910feb3e8e388a913a3ae96"></a>

## Next pages — network_pbr.network_pbr_rules.all_traffic / 9486f73f6d0e / 4

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-327cebf933308484a21cc5dc090c715c15f4a5f35d61057a4caa26d745d4a27d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1edc1cb690125c89601f7052066f29fff598947e1344fdb7018cd927cc13558"></a>

## network_pbr.network_pbr_rules.all_udp_traffic — network_pbr.network_pbr_rules.all_udp_traffic / 3d33c7e2d572 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.all_udp_traffic

<a id="canonical-b1e46191a6f4382de50ac72e5079e8ec1a97bbada1a9c59ec562bc9e7992146f"></a>

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

<a id="canonical-e3cdaca186b1c273d4145814583ad256ca27041b9117e3df970fd7434d457ba3"></a>

## Direct properties — network_pbr.network_pbr_rules.all_udp_traffic / 3d33c7e2d572 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9b1a28e676c5f2358dbbf917995f2c7733846819fdf69ad20b14b5b3759d8e1b"></a>

## Next pages — network_pbr.network_pbr_rules.all_udp_traffic / 3d33c7e2d572 / 4

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-0789ab10d71bb49fe3bfbf0e0f9221b0103b1f5a70587bbc5b5ff65ec4d18cd1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-127cdc35cc142a0cae5955e65ae8d221e9ec33272b4c6243e4310a6a0642e66d"></a>

## network_pbr.network_pbr_rules.any — network_pbr.network_pbr_rules.any / a443cb9855c0 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.any

<a id="canonical-e425bbfedff900d302212916319c4e8790df122eb82237d42de8abdf75f579ee"></a>

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

<a id="canonical-733cc05f38ad575707fe285b500ce8554b50cc21079ad516b2df1f1244dfc747"></a>

## Direct properties — network_pbr.network_pbr_rules.any / a443cb9855c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-96828a21ac91ef67733e658eb2955a73c7ab9834939acfeebb5385b52a6b3009"></a>

## Next pages — network_pbr.network_pbr_rules.any / a443cb9855c0 / 4

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-0eccdfbe476bb68c4ce8f8b4696e44ca09407edaef6d2abc28e5756517ba193c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad3888d95849dddbecd26083bd8e972453678e57048ad11b1ea96d4a3e14952e"></a>

## network_pbr.network_pbr_rules.applications — network_pbr.network_pbr_rules.applications / be32c93b5e20 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.applications

<a id="canonical-0ec5c339bce94ae7ef3bbb1dd88269d2d8dbf02e5431107875b83225473d2cfa"></a>

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

<a id="canonical-ced05a89d5732ff1f2b7f9bd8c249c13f9a80b4f9014feed143254332ff97035"></a>

## Direct properties — network_pbr.network_pbr_rules.applications / be32c93b5e20 / 3

<a id="canonical-b1ded9b50b4883ae2f508569d636fdc46dc59c806b0d14b1d4513ff0a14925b1"></a>

<a id="canonical-e42e78f6016c8818279b9ddfd9cfb6a1fac9b544c914770e25a805c481214966"></a>

## applications property — network_pbr.network_pbr_rules.applications / be32c93b5e20 / 4

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

<a id="canonical-abb24de707d948f977a605bd782d2b989622b2f6932143d0074d492df3eda51a"></a>

## Next pages — network_pbr.network_pbr_rules.applications / be32c93b5e20 / 5

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-023fd91ef592f0861bdc27e2568f4bea117f14e9ae91afd0ccb0fcc6b4d9ef38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59c9feff2a061760b838462c87644916f807073393362e4d0efd02fd55fe35ec"></a>

## network_pbr.network_pbr_rules.forwarding_class_list — network_pbr.network_pbr_rules.forwarding_class_list / 74862f02911d / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.forwarding_class_list

<a id="canonical-94f71d3c81a076d40939a4a40603cc9d56da224107d34b36ffb88b2c07157d23"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of forwarding Class to be used if rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forwarding_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-7c4b2e758d46c3a3a21da92d427b56f31155e628caa1abf2af14aee7d23bab98"></a>

## Direct properties — network_pbr.network_pbr_rules.forwarding_class_list / 74862f02911d / 3

<a id="canonical-f1b568d173cdecd76e35cf02b3a3699a4286a9a9a68220e21ae1b88a669847c5"></a>

<a id="canonical-af09878891eace147edb85d327381797aa4ed28e8d26dc2aed20b77894cf12b1"></a>

## name property — network_pbr.network_pbr_rules.forwarding_class_list / 74862f02911d / 4

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

<a id="canonical-ac7995c308b1e86dd8d92e615cecbd342aecf085a64c08517df8dc1e888211f8"></a>

<a id="canonical-8c45d2e36f05255bcd73c068922495c6c3eb67ed57ac00d38dbb6b18eebb1787"></a>

## namespace property — network_pbr.network_pbr_rules.forwarding_class_list / 74862f02911d / 5

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

<a id="canonical-d25cd9d21e28f7c36b183538e93f223829ab4467a79642223df719488db8092d"></a>

<a id="canonical-2c08a2cde163417ad8edfb5bfa979a335b454ce53fb8cd14a75b6855ac294d09"></a>

## tenant property — network_pbr.network_pbr_rules.forwarding_class_list / 74862f02911d / 6

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

<a id="canonical-94961f55bc5e3d18b7126920b9b7bcac4e74e6d106d987e99376ba4572ef16a5"></a>

## Next pages — network_pbr.network_pbr_rules.forwarding_class_list / 74862f02911d / 7

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-28c5523a277bf151cca104ebfff87e2324e74107899f4a97b578761a3bead976"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b282af08b4700e5744ef80bd01239e09f3fd3b0a68138b9094617e01f356eab7"></a>

## network_pbr.network_pbr_rules.ip_prefix_set — network_pbr.network_pbr_rules.ip_prefix_set / a693b8dee61a / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.ip_prefix_set

<a id="canonical-471538d8a75ca030df70e2640c22c62e9a78dcfa0675ee3c60a3df103cfb5c49"></a>

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

<a id="canonical-0dccbbd7f11c43f7009d9639d32957ee9ca1e646680daad4d83fcefbd84af4d6"></a>

## Direct properties — network_pbr.network_pbr_rules.ip_prefix_set / a693b8dee61a / 3

- [ref](resources--policy_based_routing--reference--group-001.md#canonical-bc9fd7f2f8849b074307a5207b75431c758536163a421adb62704050726bf61a): complete subsection reference.

<a id="canonical-6de345c29fb8434aa60704d8ee02aed2f52c0fca5c4073c370801ed0d539caa7"></a>

## Next pages — network_pbr.network_pbr_rules.ip_prefix_set / a693b8dee61a / 4

- [network_pbr.network_pbr_rules.ip_prefix_set.ref](resources--policy_based_routing--reference--group-001.md#canonical-bc9fd7f2f8849b074307a5207b75431c758536163a421adb62704050726bf61a)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-bc9fd7f2f8849b074307a5207b75431c758536163a421adb62704050726bf61a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d00040d868a2646a7ba0f1d05b99ade67459e494ce1abcf8d95aacc839e5c7bf"></a>

## network_pbr.network_pbr_rules.ip_prefix_set.ref — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [network_pbr.network_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-28c5523a277bf151cca104ebfff87e2324e74107899f4a97b578761a3bead976)
- network_pbr.network_pbr_rules.ip_prefix_set.ref

<a id="canonical-bbb01a4b896c925729dbfe1dc5f37d8fe92d905586bf339f49b4cb9dd9ca908e"></a>

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

<a id="canonical-22f62b27d0ec8bfc8f21d415491f2d06855daaaa26e49f8aec519a94dcb7886f"></a>

## Direct properties — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 3

<a id="canonical-3ebb1bea2ff97d13d64a98457694e561339c69c251258bb65ee6411dc83b0039"></a>

<a id="canonical-a83fa32fd65b2177b59c1c0ed9e0299e9ea4e3c6ca17133cad5f4fd87c1848e3"></a>

## kind property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 4

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

<a id="canonical-1eecc7dfef1443bf326bfb740dcf0e95acd9d81344af5c01499f80b58219cb66"></a>

<a id="canonical-0b6f6c086c343240d83bbb7e8c194f444aa09bce2332787ad84d2bc486e58941"></a>

## name property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 5

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

<a id="canonical-f6562b7abd35ebbf32cc7a485ca8cf0d5a9e5675d141fbfd8fa493f771e63b06"></a>

<a id="canonical-781fe34e9b264a4fd682ae6669b367109c3828364a32bfcacf286fb5c9fb6cfd"></a>

## namespace property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 6

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

<a id="canonical-626cb0aaa9984d87663dc05edb23e70f4bf276b5826baece433cf25680fa0daa"></a>

<a id="canonical-262a47c3fd607c4a15c11952190b86c115f4fb06b23e12d501d42aa4aff282c7"></a>

## tenant property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 7

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

<a id="canonical-4a45da6501a72e3cf352c62153358df5a3ae478e625489ba76b61c16d66adf4c"></a>

<a id="canonical-3cb0059c2ee8e70f04b73576f55d040c093a15293bff35c6fdc8991fbd375ca5"></a>

## uid property — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 8

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

<a id="canonical-70c19c904b96a5e8b03aa44d08f2eb605b17565ff8ecad391157fac7d6e80c3f"></a>

## Next pages — network_pbr.network_pbr_rules.ip_prefix_set.ref / 5bdaf70a949e / 9

- [network_pbr.network_pbr_rules.ip_prefix_set](resources--policy_based_routing--reference--group-001.md#canonical-28c5523a277bf151cca104ebfff87e2324e74107899f4a97b578761a3bead976)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-0c93a1918183d0c4e61929c375ddc25b7c500d730b883ef55c0edcb5ccc7308d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-018c34ab8cc4533da36d2d4cc2c2b4add36932c1d3f9033820ca2ac48c749c59"></a>

## network_pbr.network_pbr_rules.metadata — network_pbr.network_pbr_rules.metadata / 942fbc42bf42 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.metadata

<a id="canonical-591c39c3d0cebd72cf4f4609d815bc70d6dd4c50873d85e756937e1ebeb7e497"></a>

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

<a id="canonical-355ceeb411c4c8396d34525738bbbb52cb3f61388873024b1e282f81932a1884"></a>

## Direct properties — network_pbr.network_pbr_rules.metadata / 942fbc42bf42 / 3

<a id="canonical-a26891fa5d379edd6ccabd64836f68edc32f919569cbd145b1b9cfb0eaebac2a"></a>

<a id="canonical-adab542c229fd6cf494c304cfd9f5f5aafc79cb153ea66ddc08110338d5ddea9"></a>

## description_spec property — network_pbr.network_pbr_rules.metadata / 942fbc42bf42 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-6292a390a10fdddf08dae9a537da47b2b1d282dc2cf01d233dcc4f89533f662d"></a>

<a id="canonical-777ef5ccc2731c520450c48f57b8a3e08f150ebc4b618eb39e9f8769b9c8e8cb"></a>

## name property — network_pbr.network_pbr_rules.metadata / 942fbc42bf42 / 5

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

<a id="canonical-3b58d698c81e375e39d5deb1e546745689332f12cf9a344b55aeddb773bdbf3d"></a>

## Next pages — network_pbr.network_pbr_rules.metadata / 942fbc42bf42 / 6

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-9e7a0506ca13e72959a49dca526cd737bf77c253b260941b330118d95c83ad7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-795374b089364465e9c5d887338dfba68676a356b95ff27f3e02f194c2b54ec4"></a>

## network_pbr.network_pbr_rules.prefix_list — network_pbr.network_pbr_rules.prefix_list / 01c2370afe87 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.prefix_list

<a id="canonical-982a4f5d95f6eeca4233018e616c5d0827e552a015af70a2e7f6cc8cc3afa032"></a>

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

<a id="canonical-90242ec7e1fde89bbd26b09ab994069fae621a7ef6753a5516619ac10c4c486d"></a>

## Direct properties — network_pbr.network_pbr_rules.prefix_list / 01c2370afe87 / 3

<a id="canonical-c57cc1fe687fc4bb80a991f47f7513c810bc8e2c26ac866eb35de49eaf1f186c"></a>

<a id="canonical-4be8648aff80a90f7c10898dfb8699487915265a50dbd10ad10575bed8095ecc"></a>

## prefixes property — network_pbr.network_pbr_rules.prefix_list / 01c2370afe87 / 4

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

<a id="canonical-e6232dac08c0ec8b7250005d52f94206701ade7e79783bb2effc1dfdfd0a00eb"></a>

## Next pages — network_pbr.network_pbr_rules.prefix_list / 01c2370afe87 / 5

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-58a7d1612de059d29372c9a107bb7bd930e27c8f95a55c36f6bb37d194f82b92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-954f550370fb4229cf1332a3b3f20243abc24d5f2bba8b73f1020a33b37781e4"></a>

## network_pbr.network_pbr_rules.protocol_port_range — network_pbr.network_pbr_rules.protocol_port_range / 511fc964c733 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- network_pbr.network_pbr_rules.protocol_port_range

<a id="canonical-75d2e1111e01a1a19e76efe480666f9ef0db6d8290b0f7b0e87e2af2ef339f98"></a>

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

<a id="canonical-1f60146c249b63ec9fc533c6b036846d4c961cdecceba07fb2b8225b47e38a12"></a>

## Direct properties — network_pbr.network_pbr_rules.protocol_port_range / 511fc964c733 / 3

<a id="canonical-234719337b1848cd5add3a0ca94cb8ea2679c008288465beaf8462aa3d97bc67"></a>

<a id="canonical-b6858f3218dde743e0c51a7d9280075c0cad0a920b3258fb9e8ecce673bfeb83"></a>

## port_ranges property — network_pbr.network_pbr_rules.protocol_port_range / 511fc964c733 / 4

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

<a id="canonical-a62f5c6f5129c03792ed02fb9a379bc4978c1169a3850d2413484740f5345c9e"></a>

<a id="canonical-c02901106eb424199db5badeff82fac92371720b378156c674a72b1eafc3e56a"></a>

## protocol property — network_pbr.network_pbr_rules.protocol_port_range / 511fc964c733 / 5

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

<a id="canonical-81483affde7fd37af4c1fe9631ecf1263eea68528f268df87f0063bc31549233"></a>

## Next pages — network_pbr.network_pbr_rules.protocol_port_range / 511fc964c733 / 6

- [network_pbr.network_pbr_rules](resources--policy_based_routing--reference--group-001.md#canonical-7fde0f0ef93a1cfc4209bcdad11a0fe64e45d314561d112e14a8fb3366c8e7c0)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-d19ae0c9bb95b5010653fabd6116db2d9235b7817e72fcb0ac948dcc3eb8d621"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0aaaaca817544ce3dc7c6c2a8097edd9e1413ed0317d4923a936e2fcfdd56ca8"></a>

## network_pbr.prefix_list — network_pbr.prefix_list / 9acbd88eeb1e / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- network_pbr.prefix_list

<a id="canonical-d2dea93b4e58534023680b4ca05c0011e80661409a584372b2210bdaba3f037d"></a>

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

<a id="canonical-b088fe335285bff39ccc270e4c7dd21f5ac5ead67a59e2701df26c0b881433ce"></a>

## Direct properties — network_pbr.prefix_list / 9acbd88eeb1e / 3

<a id="canonical-ed26bffcc3d444ff499e909eeb701f25d5294b33229e4809d041bce9bb8ea878"></a>

<a id="canonical-394bfc80c5739b93703735dabb96401f7c543ac613a71bc4f0ac81d27dd6b023"></a>

## prefixes property — network_pbr.prefix_list / 9acbd88eeb1e / 4

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

<a id="canonical-3cd8866c2725f42f51f7c8db7613d0679ad9080cbdbd8fa0c97082e0273c8654"></a>

## Next pages — network_pbr.prefix_list / 9acbd88eeb1e / 5

- [network_pbr](resources--policy_based_routing--reference--group-001.md#canonical-f140dbda49b0864a651a1339d2b7728e452eed0f17036dd312dae9d5009ff3d5)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-e6b0b7cdc591feff0e35c2994ef5aa4ca05a7f1daf4ce44cb331cdfbaf70222f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb72f23ecd506f0f242275934773e9e6258938bd276cdd9cff79ee6694bb2dda"></a>

## timeouts — timeouts / 030aa99a1a80 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- timeouts

<a id="canonical-bf50f8189e76663681d9b08046d04d8bf1380937f99a6dc0251c35a8f2999fa0"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8be3dabde7f32e82a14555b609d2a44608392c6611a69784d1be326f4a4b63e"></a>

## Direct properties — timeouts / 030aa99a1a80 / 3

<a id="canonical-117951dac4a3f59ea9590b3ddb7869ea6bdbf70527286ed939ed27e4eeca692e"></a>

<a id="canonical-dfb2f079cdffede709e0f4ce4cc0bf4180e9ba207b9aebb3d750dbda5feee6b4"></a>

## create property — timeouts / 030aa99a1a80 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9c52b9d14fe622057ffa2bd26a1131b14df210e3d6f4b6a17616c006fed907a1"></a>

<a id="canonical-2af5e7a296af837c1b3a3c42afe77f5bc4b0c15e8f1a34c632677934ac918f89"></a>

## delete property — timeouts / 030aa99a1a80 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-19442ca3398f627720942504c6c319bf34e4e3ad41da323da8553ef3492b553a"></a>

<a id="canonical-1e97a6b574b3e8f7a535a920c1e60c826b9864f751f5fcfb38e33106c4233b39"></a>

## read property — timeouts / 030aa99a1a80 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-89448f7269c12549be8111939e879e9fa2e5e12d7f9e023f2241d3f0ed2b4429"></a>

<a id="canonical-a61d0150855f5ad46fdab2c1ddc72cc8ebb9ca85eb72ba0f39978a66e978c8a3"></a>

## update property — timeouts / 030aa99a1a80 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2d1dba7c7caf728beffaa1b33d98f6fb9b46dc2128ac6d4c7356aad1c4c8fdcc"></a>

## Next pages — timeouts / 030aa99a1a80 / 8

- [Property reference](resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
