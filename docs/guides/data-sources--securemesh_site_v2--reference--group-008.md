---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2658ddb5ad28b63865e37bb82dd3dabb24d0cd562cf4ad0e5ef91bb23211f0b3"></a>

## dns_server property — eks_k8s.not_managed.node_list.interface_list.static_ip / 1bdb00514b7d / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-958bea74bfc4526083ba38a52c5b32f43a4f779b8273baebcd7d0f35b5db011b"></a>

<a id="canonical-d2f372909a093858e3bc51f9e939abe4e21247682b5330cc09a5ddf87661aa6a"></a>

## ip_address property — eks_k8s.not_managed.node_list.interface_list.static_ip / 1bdb00514b7d / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-9eec688a0b67556979f6f1cb07c450041337d65c21721b949a7d49c5fed77104"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ip / 1bdb00514b7d / 7

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6fbeb2b610cece275c75e13c9808134564c88d7778450503142d5d3f5e1bf285"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e11f95216c1639fb87f99119ab7161246d677ac33e5dcc9da9e6c113ab296d12"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address / d596ed713033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-e75ce74bbb253b8668c2e4439f552a1dfa65e7cb3624c2bbddcadd7c15fad8a3"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-d60dc3894757585c1f36b35243b3ad17849caf63bac4b6a6092e52401f03e42a"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address / d596ed713033 / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-008.md#canonical-118c954749f3cefb869ced97dc4f95cab2c782dd1baa491732a696f21ab833c8): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-008.md#canonical-8d58fa7ff92241d73675ae46e1bf9f3ad0a5e089f19ca5b537bebbe40f219c4f): complete subsection reference.

<a id="canonical-f1fc5da818960f55ac0267f9af8e5fcf18f2f5da4278b7b8ab7267a0d47ba8ef"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address / d596ed713033 / 4

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-008.md#canonical-118c954749f3cefb869ced97dc4f95cab2c782dd1baa491732a696f21ab833c8)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-008.md#canonical-8d58fa7ff92241d73675ae46e1bf9f3ad0a5e089f19ca5b537bebbe40f219c4f)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-118c954749f3cefb869ced97dc4f95cab2c782dd1baa491732a696f21ab833c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81f0ca83e8ce25226ecf1cbba1584e859ecf9fc803cf7cf8fce98903951e5949"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 10f97e74144d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-6fbeb2b610cece275c75e13c9808134564c88d7778450503142d5d3f5e1bf285)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-8c63137fbe9294d29e29f5dab8aad62af14c21831034da3f615c7301fc440c92"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-4a5baed11571439d29336e91e834d57e1492b7d2b9d80ffd355f56ac38a1e562"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 10f97e74144d / 3

<a id="canonical-dabb51434372b084f2df9c943cce94f0a0bc5c739cbec8bdf09f6009cf68ce48"></a>

<a id="canonical-fc73279f800546f45e84947a95b3c4e96c42be1b670f1f2458ba6d66ecd21a73"></a>

## interface_ip_map property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 10f97e74144d / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-f0ff1e0426ab2e48e5ac8fbb7b79116a709d350b3a2baf02419d593c51131974"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 10f97e74144d / 5

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-6fbeb2b610cece275c75e13c9808134564c88d7778450503142d5d3f5e1bf285)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8d58fa7ff92241d73675ae46e1bf9f3ad0a5e089f19ca5b537bebbe40f219c4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3989d0f2ad5e2a98daf15c3a23e7296245f694565aa8da39b214c93ab4cc1388"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / fa5adf91fa8f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-6fbeb2b610cece275c75e13c9808134564c88d7778450503142d5d3f5e1bf285)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-507cbaa913938a3ac9b016779b72019b74aba01aa0d63d3381db6e209ab3423d"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-e8203834cdb094e2d0e20d694f900e8c8b3b38cb76f967dc4dafad733353294c"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / fa5adf91fa8f / 3

<a id="canonical-91667a9dd496b7136f01f0459936ed10c6eb0ce91a1e894f1d14eedab44bddd8"></a>

<a id="canonical-e7ee4b5596ac1e860a124d89601c6cc9a874bd7e16ece190f2b748c54fdc3770"></a>

## default_gw property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / fa5adf91fa8f / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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

<a id="canonical-21362a73ed3384f25f11f18c7705f786cc163f07a84d5eeaef00514ecbdd696d"></a>

<a id="canonical-befd4a88e06a3026bd9ff892a539c319d6141fb31c2a8a4e45c8e578202830d2"></a>

## dns_server property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / fa5adf91fa8f / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-7531ea8d47a0015997c7bc0d76d9b63ae46f8eecaca0335c36580e3651b8b1c4"></a>

<a id="canonical-76cdbbedfa763cbda1666967c89e1472547229c293b70f534c36e0c7da241f11"></a>

## ip_address property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / fa5adf91fa8f / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-dfabd6518bf20161b94d60a1ef1709f0a6bff5674eb7a4e233a23a6fc7814e89"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / fa5adf91fa8f / 7

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-6fbeb2b610cece275c75e13c9808134564c88d7778450503142d5d3f5e1bf285)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d8ebbe80051ed8c3947c3a6d56a9cead9a4403c159fd18dde4e24fcd07a9a7b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa9f38ea809bf034d4d5b180b9c383c0e58bba3736050a99966a97e35a0a624b"></a>

## eks_k8s.not_managed.node_list.interface_list.vlan_interface — eks_k8s.not_managed.node_list.interface_list.vlan_interface / 09c0db9146a1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-4d6f079e931dbc379b9b61ced4c539da168d2d859a22345fe5fe22a3b47532f9)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-c26bfdf0cae1e75785957356e7945a0f9c30f1b927d6a46ab4d1645081ae3044)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- eks_k8s.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-cbc754bfbf7590b33c4d9acfe1982053e9517eef24495aac108a5bac555e41b7"></a>

Type: `"single"`. Computed.

Configuration parameter for vlan interface.

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

<a id="canonical-652ace8b335e765ed2d7d347671d2cf97df6562170f0f68d98ecfb3dd8350975"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.vlan_interface / 09c0db9146a1 / 3

<a id="canonical-189b7f7792434b5008946313f411d51a0bf2a1450ac15f3e0590ff5865dfc186"></a>

<a id="canonical-9cb09ec1c07e67479932ed9e0a030fd3af0d5b6fd79f20319b85a58b1ec674df"></a>

## device property — eks_k8s.not_managed.node_list.interface_list.vlan_interface / 09c0db9146a1 / 4

Type: `"string"`. Computed.

Select a parent interface from the dropdown.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-6b368883de29f00a53a5dcd31fb5bed29bd8dc4d874534486426aa111af1dd74"></a>

<a id="canonical-cc9a9603d1e7c90e0fcf005ea250de2453ecadf9737dfcdaf17f1f64c45d990c"></a>

## vlan_id property — eks_k8s.not_managed.node_list.interface_list.vlan_interface / 09c0db9146a1 / 5

Type: `"number"`. Computed.

Configure the VLAN tag for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-8c8a0632d05a2f4372747ff94ad3a51623f879cdef48f3804562d6e708acc3ba"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.vlan_interface / 09c0db9146a1 / 6

- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-60a457b3484e1651c301435375819ecd03e5bca84eb9347b2072ae08e179d557)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d4b6f8536582fa3d6ba3ebdfac41e7a754efe97f90c964c4ae33de0f7deaa51a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-419d787c3b9d5659a0afe06842f1abc0b3ed88c88b217c86b1201f42e6b41a54"></a>

## enable_advanced_delivery — enable_advanced_delivery / 4cb1366286d4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- enable_advanced_delivery

<a id="canonical-3241c2eb56b4585e5d89eb5b3af5db61afece4b1a74a088981bf7eebdcfd1868"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable advanced delivery.

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

<a id="canonical-b02c2d4917045f1af848297e7e44705eb598c951a9bd8a1feaf9e89f9c3173e5"></a>

## Direct properties — enable_advanced_delivery / 4cb1366286d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d02f37881e7a9525ed23597316f1c7d24434fae314301005c4819b1a45440ed5"></a>

## Next pages — enable_advanced_delivery / 4cb1366286d4 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-110bfe58069a4c03d94d1db4fd495fb1a79d2b474c0ddaad5f2973457b7b5b61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b00e140a0387e8c10fd43bf62fe746bf2c23037196bb8029858626d0a9b536c"></a>

## enable_ha — enable_ha / f4df1a6c1a7f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- enable_ha

<a id="canonical-e5bd040dd34926397e376b81d1f9a808c2b3f5e754a2e4f1c63b03e6112c931d"></a>

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

<a id="canonical-e813f64c4ea09dc3ace079935e1fe8248315abf6b67b1456b0de63a6095f3b18"></a>

## Direct properties — enable_ha / f4df1a6c1a7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e5c9e895fc8c517165a47c218973e99e9e068da4c13073e5a1226290a3bc705"></a>

## Next pages — enable_ha / f4df1a6c1a7f / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bb396062a93950d63ade3ee0d4aafdb1b0d1538c84b2018f6d5de4923a32ff21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04fb0e89af5522c106ccb999f02611aa395167a98be2e06932011b5846f95686"></a>

## enable_log_anonymization — enable_log_anonymization / b1c807abdcad / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- enable_log_anonymization

<a id="canonical-fdca99a3fb88947ec15b6353aaa2ca758cf0a86ed5ce874c9b185ed4d902c4c6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable log anonymization.

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

<a id="canonical-389db1cd2086c9c3b402c3d294e24f9ee5e49e015dee080da5ce3c5725563809"></a>

## Direct properties — enable_log_anonymization / b1c807abdcad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be3575e6c57034425a829ac57f5ffa57adcd27aa3f401ce068fef76207de887d"></a>

## Next pages — enable_log_anonymization / b1c807abdcad / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5091645bfaa723812d5a8f2b667f8230bc7d521bc763a6858a94a5360e286f55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-877548994dcec44b1fdea9e34b2a19ed87195b1148c3f16f04808fe5688e9981"></a>

## enable_management_network — enable_management_network / 8cf83ac0c42d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- enable_management_network

<a id="canonical-45a1bcf6f61e659e90fe138f01a89ee4c354be889668e831e114ef8b64bd8c26"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable management network.

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

<a id="canonical-94a9c3f9c29db375913cd4c61474c9da1ff348841cef48489c1f387197cb4b6f"></a>

## Direct properties — enable_management_network / 8cf83ac0c42d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b95def3039f6d6e89aa0fa0dde3aa071cbdc435d99143b02151167716e5054b1"></a>

## Next pages — enable_management_network / 8cf83ac0c42d / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e2126e998b0f941149b8a149f21f7314dedb2d1712b4b282bc4e91af2320a305"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e970f7b1f8d5e74764d989eff0e616fc34c477dd47832f0e8a54c929f04cf79"></a>

## enable_url_categorization — enable_url_categorization / 4f852fec25c0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- enable_url_categorization

<a id="canonical-afc9dd89d440e4dc3318e2671b1cb46b11997dc760cc640528b5bbcb13f04975"></a>

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

<a id="canonical-0b6b01e843c1c6d22f9a7c69bbfbfbdc2893cc80230a31c602a382f1ee2c36d1"></a>

## Direct properties — enable_url_categorization / 4f852fec25c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8519703c2bbb7618c424c8279729a9bf8131f8f6e80d7e8a2d785833bc1ea57"></a>

## Next pages — enable_url_categorization / 4f852fec25c0 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c8f1da81054550d98ed5a5bd5cd414f2b85d4f9358b22d90ccb97209414f500"></a>

## equinix — equinix / d705c846389f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- equinix

<a id="canonical-a184215caa40d31a21b5c7091c11cb77aaee509297b0f0c20156f6e22044ee92"></a>

Type: `"single"`. Computed.

Equinix Provider Type. Equinix Provider Type.

Upstream description:

Equinix Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

<a id="canonical-5e65f9478f5ec9623d0ee83f1872716d6e56c25eaf876a8e744904b87a0cb90d"></a>

## Direct properties — equinix / d705c846389f / 3

- [not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd): complete subsection reference.

<a id="canonical-93644fd12d04f3bba8c104784eafb399d99435bdbb47668a433eaec83895be23"></a>

## Next pages — equinix / d705c846389f / 4

- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-050853605c2f489d0f4c952ab67daad5b95a5c72352fcfbb7558cc44287c7fa6"></a>

## equinix.not_managed — equinix.not_managed / 60efd080771d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- equinix.not_managed

<a id="canonical-aaa00013fdea83aa2466683f065d93452d2a6c9687c9d72d61952d29b97447df"></a>

Type: `"single"`. Computed.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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

<a id="canonical-5934cd024123f37dc6a7886923d3772486f5124d0564ab127e1b8a1d4a6e8194"></a>

## Direct properties — equinix.not_managed / 60efd080771d / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a): complete subsection reference.

<a id="canonical-50dd2746c8739ffa546c6255e639348dc4fef34daaed21eb4d607b70f0fb931b"></a>

## Next pages — equinix.not_managed / 60efd080771d / 4

- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abf7fa61e11caf8449992063640bbe7c10c3237a2588974d34ce30882ac01354"></a>

## equinix.not_managed.node_list — equinix.not_managed.node_list / 398a89d62573 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- equinix.not_managed.node_list

<a id="canonical-0f4d316e16f95d6bcb90a8d01432ebf00309597b4a2d4cffae9ea46b7b322d0e"></a>

Type: `"list"`. Computed.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-eb6931e59da42cfe2ae493d593742bcb05b2586c4969064585a8c3cc96708408"></a>

## Direct properties — equinix.not_managed.node_list / 398a89d62573 / 3

<a id="canonical-2a797c36a9b261dfda842a943d5bd86e5682df00b575d52bc822fba4027c355c"></a>

<a id="canonical-86e5dfd70d3a817d3f257088749fdc3b02b7b76183dcee0385e54e390ce0791a"></a>

## hostname property — equinix.not_managed.node_list / 398a89d62573 / 4

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067): complete subsection reference.

<a id="canonical-a0510ee14364c8780f0c607786cf62c21a51b3ef9e9a4368360c4d1c5dcc9e6c"></a>

<a id="canonical-67035ab6492b0afdcd3d80515457e26bf55d2284b823dafb5f33d8cda7fcebb2"></a>

## public_ip property — equinix.not_managed.node_list / 398a89d62573 / 5

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-6cb1447bf41104e750521b136ba2aafe09145dc57e4fc7c53f128759316a09b9"></a>

<a id="canonical-4fa7fd287b2e29fcacc5f5ed0e60b64ac0291ae0579531e6020f863a668c4780"></a>

## type property — equinix.not_managed.node_list / 398a89d62573 / 6

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-6d7dc35a8d3e3522d9421bcd989b8c785173237bcf0b791d8a07454e8d742b3a"></a>

## Next pages — equinix.not_managed.node_list / 398a89d62573 / 7

- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1de1e2b6ad014f95b4754a152115d38818d4f4e7c161e9a0123109b2f6f35714"></a>

## equinix.not_managed.node_list.interface_list — equinix.not_managed.node_list.interface_list / 25459b542a48 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- equinix.not_managed.node_list.interface_list

<a id="canonical-4f47f155d6d5903a76508065ee9bb577ee07941eeb78b68e2e074a043ebe553f"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f7d4a4b23a30ba0a5a07cf81c15bb05ac0845eaa66c0b39984fc618ec3c2936b"></a>

## Direct properties — equinix.not_managed.node_list.interface_list / 25459b542a48 / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-77d0a0b944b47299d6edf703f27669702724d47a7084efbef8247598f984db0c): complete subsection reference.

<a id="canonical-aadc54ffcb533ba635ed4bb9ae7fa776ce2c366077c45f070981091c3752ef51"></a>

<a id="canonical-93d4e53247a5c84aece1a8d19f20c98ba3a54b2c8cfa84639896e3f5d02a5f84"></a>

## description_spec property — equinix.not_managed.node_list.interface_list / 25459b542a48 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-008.md#canonical-b94f8d8e758c37f0fb9372cf8686f1c51824f1ad1056e2be00219f92e57227c8): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0c293626bc6be7d1fb03a3eba3dc6eef5a78971c8910eaf376cab8bd0d8bf514): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97): complete subsection reference.

<a id="canonical-f2129626e7f9ee25e90f9a290f43da29a52053e0fd5b6bb348ac61af6d70aa6c"></a>

<a id="canonical-a90806d26bd753754e025c6ebf485804fa07fd337172c1a5cf49160fa26e0999"></a>

## is_management property — equinix.not_managed.node_list.interface_list / 25459b542a48 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-09059467d11937d3bfa82dc0de8094e8c85f19024f1299922b0e042734702a7f"></a>

<a id="canonical-48e22b43fe1dd71cad3c27f4b0a989203f322b5bb123c4ea90d4b0c9524f79e3"></a>

## is_primary property — equinix.not_managed.node_list.interface_list / 25459b542a48 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-57798bd18f4bb8ccaacb21425c664fae682479ed5bb94d154103be89b0953191"></a>

<a id="canonical-640e6aced4bd0ca5b4e0c6e41fcf9aed97f1d60f958fa68d2238661c3a7a68be"></a>

## labels property — equinix.not_managed.node_list.interface_list / 25459b542a48 / 7

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5427fc0c20d6c5d9e94a2491696e0cfd31d3f15be75d030a01b0ed122c3e3e44): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7bed8611b98a0bad14f1472c9173640f72d47cecab5a2dfe1f6797652ada860b): complete subsection reference.

<a id="canonical-6b2aad5eeb2f58d55a0442cda58cc3c98679d02e2861e185044a5a6968d859de"></a>

<a id="canonical-4987060ee284824cddc5cc56abbb84ff45f68f5713e9de7095db127d0b9fa7d3"></a>

## mtu property — equinix.not_managed.node_list.interface_list / 25459b542a48 / 8

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-95db4397b7eaeda491add51e412fec27701e73ecca4c079f225d0320c0c152c7"></a>

<a id="canonical-3559be9f5be5e60ee8ba0dd682e19b5165149cb3a0451d395a89061752dc7346"></a>

## name property — equinix.not_managed.node_list.interface_list / 25459b542a48 / 9

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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
    "format": "dns-label",
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bb89ee17b354d7dcad9838f440e927d5a3ab48b5ad2d1067cf53bf9742d7229f): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-860d5b183be5c9e63857d838c17b844bceb29a4fd89d0d9509a54cea5346539b): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-83d73ae31e6a85c95bcc36161493a4a369ffa503787c2f6a60a4cc52c1187862): complete subsection reference.

<a id="canonical-41b1bd4b3e4020922f21df68d49353b83ecea44bb5983868c1dcc48f40b3490f"></a>

<a id="canonical-876caf068249bb48732fe9e2d35abcc3ffc8bc9ed6f74194248da15f77890058"></a>

## priority property — equinix.not_managed.node_list.interface_list / 25459b542a48 / 10

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-009.md#canonical-db5a4abc8f6da1038499037309d7f6a0fadc16b4344fc8fb275cdd7080eed76f): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-009.md#canonical-75d4d5f16a4786ae49388a71c10f12f9a45c802ab3266548baa8cca5c098136c): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0c9026688941bda6cb8558db5d6536dfd4e6928a3ac60bffe0ed872e7d953ad2): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-7f3582762ff671e9a24d16933fc7ab051945cca7705bf36338876c24fb07bf44): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-36975a8f40e8f2b0a80a1b330179e1b41fdc237e695f3d6bd3b130f400a81d4b): complete subsection reference.

<a id="canonical-9a96fb48ae7466ea67b36ea17ea49222d0b02a824c818eea65169ca08dfedb26"></a>

## Next pages — equinix.not_managed.node_list.interface_list / 25459b542a48 / 11

- [equinix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-77d0a0b944b47299d6edf703f27669702724d47a7084efbef8247598f984db0c)
- [equinix.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-008.md#canonical-b94f8d8e758c37f0fb9372cf8686f1c51824f1ad1056e2be00219f92e57227c8)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [equinix.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0c293626bc6be7d1fb03a3eba3dc6eef5a78971c8910eaf376cab8bd0d8bf514)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5427fc0c20d6c5d9e94a2491696e0cfd31d3f15be75d030a01b0ed122c3e3e44)
- [equinix.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7bed8611b98a0bad14f1472c9173640f72d47cecab5a2dfe1f6797652ada860b)
- [equinix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bb89ee17b354d7dcad9838f440e927d5a3ab48b5ad2d1067cf53bf9742d7229f)
- [equinix.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-860d5b183be5c9e63857d838c17b844bceb29a4fd89d0d9509a54cea5346539b)
- [equinix.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-83d73ae31e6a85c95bcc36161493a4a369ffa503787c2f6a60a4cc52c1187862)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-009.md#canonical-db5a4abc8f6da1038499037309d7f6a0fadc16b4344fc8fb275cdd7080eed76f)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-009.md#canonical-75d4d5f16a4786ae49388a71c10f12f9a45c802ab3266548baa8cca5c098136c)
- [equinix.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0c9026688941bda6cb8558db5d6536dfd4e6928a3ac60bffe0ed872e7d953ad2)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-009.md#canonical-7f3582762ff671e9a24d16933fc7ab051945cca7705bf36338876c24fb07bf44)
- [equinix.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-36975a8f40e8f2b0a80a1b330179e1b41fdc237e695f3d6bd3b130f400a81d4b)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-77d0a0b944b47299d6edf703f27669702724d47a7084efbef8247598f984db0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b26bf6b9d2137a0f1f77fb0785c4860c4fab8bc1c0d1e7e2804a5e6c351dc6a9"></a>

## equinix.not_managed.node_list.interface_list.bond_interface — equinix.not_managed.node_list.interface_list.bond_interface / 0689b5e72f95 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-d1794d0fb55487ef7f4cb9ebf15f1b81c72d8479a32bd9943dabc63a2c898f18"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

<a id="canonical-dd146ccdeb8b55012249473b9750b72726effc88e0914e1b0866811b37acfc7b"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.bond_interface / 0689b5e72f95 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-008.md#canonical-00209bf694c2bde1063892fb75ab684cda154342df580df37cab7194eca65118): complete subsection reference.

<a id="canonical-b7f3274f6fa99c6fff075ac96a44f94e8719726910abf9f3a1ae76d4cb1ed857"></a>

<a id="canonical-871955d290656b7d6cd91c6055760ac2b3b0c82c24af4cc9632957a92c5baf80"></a>

## devices property — equinix.not_managed.node_list.interface_list.bond_interface / 0689b5e72f95 / 4

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](data-sources--securemesh_site_v2--reference--group-008.md#canonical-c64212f501eda5708db03dd60e77c40c87f855781a9e2c97660b685a61c981bd): complete subsection reference.

<a id="canonical-fdd7a200170b10f82493ae5ec248b3435ec8526ced96f5d6e0db645a3a7c1df5"></a>

<a id="canonical-cdaf5df15b03e9fd12e9f97f45fc794e7fb0038c28e9d3da1577989c5a49a505"></a>

## link_polling_interval property — equinix.not_managed.node_list.interface_list.bond_interface / 0689b5e72f95 / 5

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-1dbb983f1748a99f18da4e76d3310cc667a5fa016d959a99b6a5d46f40ff1d71"></a>

<a id="canonical-835e0b1557cb26a18ec20d634e54c4ba4176bf5667a8bc09d446d49d58aebba1"></a>

## link_up_delay property — equinix.not_managed.node_list.interface_list.bond_interface / 0689b5e72f95 / 6

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-5239e7e6c0f7ff655a13164ecb06d94dd88a1454d704849a0a599c03d39da985"></a>

<a id="canonical-45cb3d93b2ec6f83fc0237e2c8ecdc52a94e3d7f22274f13f1caf5ec68b0fdb8"></a>

## name property — equinix.not_managed.node_list.interface_list.bond_interface / 0689b5e72f95 / 7

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-5c6da4fbeb08bf8c83f4eece52c8e899053ba57a1efa60f4faf9bf04233b1edb"></a>

## Next pages — equinix.not_managed.node_list.interface_list.bond_interface / 0689b5e72f95 / 8

- [equinix.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-008.md#canonical-00209bf694c2bde1063892fb75ab684cda154342df580df37cab7194eca65118)
- [equinix.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-008.md#canonical-c64212f501eda5708db03dd60e77c40c87f855781a9e2c97660b685a61c981bd)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-00209bf694c2bde1063892fb75ab684cda154342df580df37cab7194eca65118"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-201b704a79a8ff17ce41f2aa4baa1c9948fe2940ca9e05dfdd5cb36a92b771ea"></a>

## equinix.not_managed.node_list.interface_list.bond_interface.active_backup — equinix.not_managed.node_list.interface_list.bond_interface.active_backup / 45e04efd1259 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-77d0a0b944b47299d6edf703f27669702724d47a7084efbef8247598f984db0c)
- equinix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-75016e5b06a6424923bc3608f193dfebedfbc87dd954075f4d02a6a212573943"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for active backup.

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

<a id="canonical-b4d60aea1bc4e6931e7be45fee98cf679118b780680744fcafd8a01912baf063"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.bond_interface.active_backup / 45e04efd1259 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-daafd6131c6e3239f0e5d2990787ff9bfbb7f496fbf9c25b0629bc48d4cdd29e"></a>

## Next pages — equinix.not_managed.node_list.interface_list.bond_interface.active_backup / 45e04efd1259 / 4

- [equinix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-77d0a0b944b47299d6edf703f27669702724d47a7084efbef8247598f984db0c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c64212f501eda5708db03dd60e77c40c87f855781a9e2c97660b685a61c981bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39bd6033a26b2e06fc9a278bfaf8804571289db39baf6d6cffae2795dd36bf45"></a>

## equinix.not_managed.node_list.interface_list.bond_interface.lacp — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 6029ac5fc790 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-77d0a0b944b47299d6edf703f27669702724d47a7084efbef8247598f984db0c)
- equinix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-36a629cda5b73f598c91c84221185412e2ba4bd791c3566cb00e182377ef7d5e"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

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

<a id="canonical-d4200e07a41abf727722389fb7fd162f8cfab9438319e96d2af9d08cded35f01"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 6029ac5fc790 / 3

<a id="canonical-0055954dadca04b863316e73d5444e7058220a36e3f038910e110fd19105c5d5"></a>

<a id="canonical-148cfc1a4d2eef1b707fefc1be8b13fcbd07bafcaf324e269568a6075f87a42a"></a>

## rate property — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 6029ac5fc790 / 4

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-7518f94d1e6f1baa106b114de3341d7b5121f00e51a97707b024f10c5c445446"></a>

## Next pages — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 6029ac5fc790 / 5

- [equinix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-008.md#canonical-77d0a0b944b47299d6edf703f27669702724d47a7084efbef8247598f984db0c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b94f8d8e758c37f0fb9372cf8686f1c51824f1ad1056e2be00219f92e57227c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16bd12705a34bd82b97abf21930e716494084a2e6d46d15b8b013edd3171a525"></a>

## equinix.not_managed.node_list.interface_list.dhcp_client — equinix.not_managed.node_list.interface_list.dhcp_client / ed87c46399aa / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-886db3a3243746127d153806565d1a5d76d4cab859815b23c85f2158a1044922"></a>

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

<a id="canonical-e5a6c82cceef0c69dc2ea0f7c34cc1e41d1dfc2943a3d6d576181982452a90bb"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_client / ed87c46399aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65f1b853b2cade2ae3fdc0f23d8267e5affd68a1db8c15bfa09fbb64434c2a88"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_client / ed87c46399aa / 4

- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e65327a64645d9ab16ffc3d1d6e9d0d4418fc35ae4faa4a0309b790c237b218d"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server — equinix.not_managed.node_list.interface_list.dhcp_server / 91da066b7c3f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-fb77c4e5768466de8d7546db490916204b1b953a45838b3965ed55b68eea1187"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-a275f538f58f94b79b84650eaf38c3d31464a6d0f5bc40974422fc72093175c5"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server / 91da066b7c3f / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7d92bdb6eeea3a4bc5adb417c924f83354d6a5177c25f060c70c4ebdd6d77c0c): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bba79a95c85a83546a447054768ba9cc1251f9567ac723de277d1f28811d6382): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a): complete subsection reference.

<a id="canonical-ded1b83acbf2e70533c9ea0b5ca30fddc9e6126ab000f48899106fc7840e2524"></a>

<a id="canonical-9a8fefa963593b7678b33f1d443001ab69a67103d1b369c94c392bb3512ce8c3"></a>

## dhcp_option82_tag property — equinix.not_managed.node_list.interface_list.dhcp_server / 91da066b7c3f / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-6bdefef1f44424a5e38939fe6e7c5aae13c69b756cf3ff8418a92088ad5804d2"></a>

<a id="canonical-08612ff1164a0caebcde8bddb9fdfacb2e69e872f97e9dc399a9f194cfeca028"></a>

## fixed_ip_map property — equinix.not_managed.node_list.interface_list.dhcp_server / 91da066b7c3f / 5

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9e98f26f234fa60d3d2a086f985967363a45a4f4758fc49bb94356b05473a821): complete subsection reference.

<a id="canonical-1e750e88eec01e94ec0df547acbfc8801d15df44269f2cd1449c833ee7d7c117"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server / 91da066b7c3f / 6

- [equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7d92bdb6eeea3a4bc5adb417c924f83354d6a5177c25f060c70c4ebdd6d77c0c)
- [equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bba79a95c85a83546a447054768ba9cc1251f9567ac723de277d1f28811d6382)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- [equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9e98f26f234fa60d3d2a086f985967363a45a4f4758fc49bb94356b05473a821)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7d92bdb6eeea3a4bc5adb417c924f83354d6a5177c25f060c70c4ebdd6d77c0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e2e0ba757c021120a670865b993fdf7780c82ac26ee87906ad218f3a346cdfb"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 99480a70cefe / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-d072dfda0021d1f9f974854dc13eaebe8d0e87f1b3e3e4976c53f90296261091"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-1f4af53a7b8a54f27b055c30ccfa5e76a5f6f3f2d0cec88a93d5714288dd3d54"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 99480a70cefe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a029acc27a0b161cffe7db2e4ee41c6cac3f017ac51cfa04be471fe998eaace1"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 99480a70cefe / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bba79a95c85a83546a447054768ba9cc1251f9567ac723de277d1f28811d6382"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9b4f8f98e1504028e6fe79163b809a2ba759a0598e7bd37402657530b00497b"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / a16065ea3742 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-485635c5062ffedf05bbd0828f807142cfe00af44bb40fd9f3bd1d10f3b0408f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-8b2d26d4aec8909ac611923e064e9cba5dd62971b6de159ed9d757b4b2cb4ea9"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / a16065ea3742 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7bffd9ea35764b4656c469cb922b834f43a27ffa8cf700975753f53fee7b199a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / a16065ea3742 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-267c7dc4f548697419cd2e8330188fdceb80891d1ffd3aae7675fe3ba53a3d10"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / fea589ab7c36 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-cd211f96df67abbbaf91299ca55d982b119de9a72ffda43ba1be3306f4781107"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-855c9a0f843daf19a15833125f5d715058513a34f51d62c3ffd6f3ce875b4bad"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / fea589ab7c36 / 3

<a id="canonical-6b1b93233e38942ddf2f5f5f319a4aef94fdf4bd49412bfdca66145400741dcb"></a>

<a id="canonical-e9ed28f61e2825247aeb78e8ff1addaac2bfa2c18eba9ab36923c76458e809cf"></a>

## dgw_address property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / fea589ab7c36 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-b08148c578decbd7ea93e0713d1362e52a6b0dc70698438f1a003ae9f7cb4fc1"></a>

<a id="canonical-02060db4141d041b5325753b275bcafa6346c4eadf6a51371196173c0602a5c6"></a>

## dns_address property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / fea589ab7c36 / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-249d4fea972fc2ee1c55b4279778017fde0ec653bc8d37fe99b14e97ec2be4bb): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f917d43392fc89fd9e5c29980e437feee8670d3352faa681c7795fab256fce4b): complete subsection reference.

<a id="canonical-c70442a2706288b945d5e2fe51933d476631e995fe4f7643482081714ded2b79"></a>

<a id="canonical-d8b975c33afc3956142ee08ed5fd50ba11e92f20026997e1fa306feffc3bfe3b"></a>

## network_prefix property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / fea589ab7c36 / 6

Type: `"string"`. Computed.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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

<a id="canonical-4dc2bd6fc3917b862ba69bfec3cfbb0ee23df45b60caf1b2d7906a8a7ad01c49"></a>

<a id="canonical-2f9676a0e718cde5bbd233345a28304ca48a71a050e0d5024a3cd4111acd575b"></a>

## pool_settings property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / fea589ab7c36 / 7

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--securemesh_site_v2--reference--group-008.md#canonical-b02279a417d124f27a98de2f3e063664791973b69c821aca89b934491ab3858f): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-008.md#canonical-092a5b2bf300b6a49411e0c0fdaa8c2882a9db8035f6c8f22e52329e704ae4f5): complete subsection reference.

<a id="canonical-9e6121ffebe2abb8ea7db7bf3b82d909330ac9e841ea8a07c5e8fda82335a8b2"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / fea589ab7c36 / 8

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-249d4fea972fc2ee1c55b4279778017fde0ec653bc8d37fe99b14e97ec2be4bb)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f917d43392fc89fd9e5c29980e437feee8670d3352faa681c7795fab256fce4b)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-008.md#canonical-b02279a417d124f27a98de2f3e063664791973b69c821aca89b934491ab3858f)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-008.md#canonical-092a5b2bf300b6a49411e0c0fdaa8c2882a9db8035f6c8f22e52329e704ae4f5)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-249d4fea972fc2ee1c55b4279778017fde0ec653bc8d37fe99b14e97ec2be4bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faf69ad4bbf50e79c6a54e91649056c5eb46693644e2e45897baf8807ef20e5d"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / d262caef9cd1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-63e2af82f5ecf3c3b4af344e3fd564b00956d000e75766bc81672ad307e2e238"></a>

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

<a id="canonical-423cdadbe2398cd479bcb8069ee9b113ca627e2f9a4d7aeff5bc16082d02ef4c"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / d262caef9cd1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-329638e086fc9dc00a152e1a8e409b912f56f703751b87e3e687b27716f9811a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / d262caef9cd1 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f917d43392fc89fd9e5c29980e437feee8670d3352faa681c7795fab256fce4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a05235286f95d909a65c03485d408dc0fc45b1ca8047c8bf4fae8fc82d87230d"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / 89697aa21dd9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-b94b90e1bbbf836f2d9431bde76697fa3c07c03f00532a27ca0dedc1b524a964"></a>

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

<a id="canonical-6acff8c532f334db6c06c1b564d7b1dad9ce774286be698bafeeb2505778d4bc"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / 89697aa21dd9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87b16d9fce0f04a6137282c4ae981de25e4d2e8bf47052413e9411290de0dbb3"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / 89697aa21dd9 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b02279a417d124f27a98de2f3e063664791973b69c821aca89b934491ab3858f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9bb6c6094e00444b34867abc7f317465f725f76304477617c1580506205919e"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 2cdfc39e710c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-a5f49244b0a495192b5841121887a931d66be3bdd18512d7b6278ce257b6127f"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

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
    "minItems": 1,
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

<a id="canonical-ee084567ede6522388655c118f72c7fb3686c664595d6a2bdc1428c9487c5553"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 2cdfc39e710c / 3

<a id="canonical-e0880c439bda72af249d5edd483010f2216b69b0ace31143af785b6e445c98fe"></a>

<a id="canonical-d7616c16a929a0d2b23522b8a49784748cc1becfb86f35542640f09eb14913a4"></a>

## end_ip property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 2cdfc39e710c / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-01a9e62e1df0b3a76e5f2e6bee5f7483e94391738a95e9d4f232ad4148e77b59"></a>

<a id="canonical-442e1c0c7880e48e3478d59d2276efeba72c06c124aa3fd0c0c6c2869dc65316"></a>

## exclude property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 2cdfc39e710c / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-78f13ce0c631b827523898a7f2e37e1da655ba325e714150fbee74f1f29cc8b7"></a>

<a id="canonical-38a8208108db5f44bc6bcdcffb649e3f6bde7c7a5d85a2b862aba0b920f1dc58"></a>

## start_ip property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 2cdfc39e710c / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-925726d9c7cdf7174adebbe72ba68ab95b97f6075f049ee14b13da316655ac3e"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 2cdfc39e710c / 7

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-092a5b2bf300b6a49411e0c0fdaa8c2882a9db8035f6c8f22e52329e704ae4f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae0fea90524d3be726c9482cc878544a26d93c6b7f79aa3aa775258377c6e639"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 298a3d163697 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-78384c36e9d14f6c7b0a86670cd21568568bef599c88f1f09401751f94dbc103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

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

<a id="canonical-362f2c7de7d9b0e1ca5088401fb9cdc0b8c7d25cc3772d837887a326dfbc3162"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 298a3d163697 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff8aded3219459bb0d8ff5509cc8fbfb31bf52c83d87d36ae9fc61879660c50e"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 298a3d163697 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0466272bc4b5df5967eaa25a7e41e317e843eab9fb5d1900a141257da62a477a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9e98f26f234fa60d3d2a086f985967363a45a4f4758fc49bb94356b05473a821"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4d3e0b80b20c85fdc25fd60234017167e59d10f12cc70ff336f7ee20207feba"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 31cadaf0612a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-88a33306374c93c0e8341323e79a8036ca0996506f66ac67051d369b6911b619"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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

<a id="canonical-d5b99729fe8edd8b9fdad793b76389e6a773e5fcd4cda732c889c7c861b5cc35"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 31cadaf0612a / 3

<a id="canonical-463453492591922241d5b5ca500cc9c78a66342f98ef8003eeda06401829fad0"></a>

<a id="canonical-816d0501a21c8f40f095d8fd2c2156423332d51a060722c4cf8728b67973c1d4"></a>

## interface_ip_map property — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 31cadaf0612a / 4

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-c547f2a130fac01fff0454cfe0c7211eb35d70b4ce64e1a916af07b97342b61e"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 31cadaf0612a / 5

- [equinix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-008.md#canonical-966a1b191071edbb138c7c62fb2dda6b5369cdfd424a4fc14e606855773548f2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0c293626bc6be7d1fb03a3eba3dc6eef5a78971c8910eaf376cab8bd0d8bf514"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e35807b4580fc694d249f76a54e6afac2ff8fab6bae122fcc9b7e47349640ba5"></a>

## equinix.not_managed.node_list.interface_list.ethernet_interface — equinix.not_managed.node_list.interface_list.ethernet_interface / 9c24147e6461 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-a878a3d4a074f24e9ce0f8a29084ce76fadc46eb91e9cd686d6032e683aab1da"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

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

<a id="canonical-bbcc16e24b770d66e64526c00049313cd4ec82ee8b63f0a903d8887624db5cae"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ethernet_interface / 9c24147e6461 / 3

<a id="canonical-4ded3f400ae1713697d3edb753869ad1fd9d049e0dfbe95b72b44f59f21c168b"></a>

<a id="canonical-d913fffc9e34139664064c8a6387adf5de29559e2696329b737cb877b78b114b"></a>

## device property — equinix.not_managed.node_list.interface_list.ethernet_interface / 9c24147e6461 / 4

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3ae8769039fd591e5c212a8c65447bae67dc985ac119eee053ee7ae679ba2f93"></a>

<a id="canonical-ba33bd1e7ce888c99c18436b552ffe1cfea9d3610d34ed731704a1832b35d315"></a>

## mac property — equinix.not_managed.node_list.interface_list.ethernet_interface / 9c24147e6461 / 5

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-b31339c7f10b47bbc1a9c12b573b84d99af4543eeaafb886c2feacecc54e0f8a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ethernet_interface / 9c24147e6461 / 6

- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-712d955e20b7159d305a1b46446ff8729096871cbdea33d2c947edbe32988d15"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config — equinix.not_managed.node_list.interface_list.ipv6_auto_config / 55614d603f8a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3377cd0c9aa67d4c5fb06e9b5af67970aff0dfd12e325c9c09a8db539da1fe43"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

<a id="canonical-7cb932f818daeee4fe26432a30f7702b41b9c3ac5b27ab660d536370450d6f9e"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config / 55614d603f8a / 3

- [host](data-sources--securemesh_site_v2--reference--group-008.md#canonical-838dabe0802ff7455dab24067d00a03b0a8e2903e9eeadfbcc42462df9ef98f9): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016): complete subsection reference.

<a id="canonical-f64edacb4e7b956b725feffc6b811cf441b25feb0ce33406f769fd49fa611331"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config / 55614d603f8a / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-008.md#canonical-838dabe0802ff7455dab24067d00a03b0a8e2903e9eeadfbcc42462df9ef98f9)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-838dabe0802ff7455dab24067d00a03b0a8e2903e9eeadfbcc42462df9ef98f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68ebefe71d8cf47c19d334579fedce008529303e2bbfd65e6540188925e801ed"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.host — equinix.not_managed.node_list.interface_list.ipv6_auto_config.host / 7e0d2efa2abb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-54e025ea672fcdf5447e0deef9bf725847dbaa8f277904812f3ab824a6a484ea"></a>

Type: `["object", {}]`. Computed.

Hostname or IP address of the target server.

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

<a id="canonical-00352036541c02d3c001dfa4c2e4b4ba29bb87b2331c2f69d981783b910d6b4f"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.host / 7e0d2efa2abb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d7ef3227ec154130244fd2655e5f2dbf2f25da14d86d7cfa760776f85c35b63a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.host / 7e0d2efa2abb / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ba2b33f1bcbc4d7f7df9bed5a91433d175f6acc9f8b749e5d5f8154a2d07a85"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 68ad1df02b0e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-0b8eee36515740ea3bebe754e7b5d1554da4963098d16bbec477464f1f820cf9"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

<a id="canonical-04aa93b2fb70975c5bebe48d4154d5cddb890019a64cafb1735db684e30d8528"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 68ad1df02b0e / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce): complete subsection reference.

<a id="canonical-189340d71e1e48102a871836f4e96db1f13c6242c4022207cbd1e607c50415f3"></a>

<a id="canonical-263a1cb6d8088b608e9cbf029f7f38a616abdd6e3e4d029fe3c1108268c63daf"></a>

## network_prefix property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 68ad1df02b0e / 4

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc): complete subsection reference.

<a id="canonical-1804eedd474b97cb44677c8fc15db4bd1b9bd78f0929e00db8631113b2581354"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 68ad1df02b0e / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ba6bbbc63f6fd50cea73e54f15092907877facb4e12f05a58166486221520f1"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 496d56e96026 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-17437c182aff7996d3fd7bff6339321bd0c87b5576e59a6471d7976602ef4294"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

<a id="canonical-f8fb4ec27978dd7eed28fd84228b4f13f85c6458aa04f2c2f2d429c91f3fcc16"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 496d56e96026 / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-718df1b7ce3d43152da89f6f73ad14a19a42ec4ca24770c17a3df951098b3021): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-a50320bd4533a7e259e8393c8b1349b93cca26ca43b184facf31cc1d1ec33783): complete subsection reference.

<a id="canonical-7fdf2a7a5bb0f2692068bea63cda8a3dda59fcc335db25f2dc11e372945aee26"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 496d56e96026 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-718df1b7ce3d43152da89f6f73ad14a19a42ec4ca24770c17a3df951098b3021)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-a50320bd4533a7e259e8393c8b1349b93cca26ca43b184facf31cc1d1ec33783)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-718df1b7ce3d43152da89f6f73ad14a19a42ec4ca24770c17a3df951098b3021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecc1ca5ac92af309b6b84ea88988b3cbfcf4a90b619793d4eaa0ae2afd7d1750"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / e7013aa6c663 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-4da1b3cc73a3f25a6eb5cdb78354ba68db5bc0c9b92d12cf6f3a7796049f07f0"></a>

Type: `"single"`. Computed.

IPV6DnsList.

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

<a id="canonical-825985333a4121583a4d8703012c07d8b26db51472073370f1bf085c795f524a"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / e7013aa6c663 / 3

<a id="canonical-933ff7582a5410e85a66c683c17e8c7365e775ab54bf9f06cff8c3febbf83fd0"></a>

<a id="canonical-270659bd5355af67e7a59544ff8d078c3332a91079f68d193876e5cfe04c4544"></a>

## dns_list property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / e7013aa6c663 / 4

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7128038ef23eab2d77f958d61e18fbca60ac818fda852bc03221939a83368721"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / e7013aa6c663 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a50320bd4533a7e259e8393c8b1349b93cca26ca43b184facf31cc1d1ec33783"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5b99887e3825e519db3db80e20d31724559860f811b115d2a6b5ec62cdfbd74"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / af57ea650307 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-ff64b11fa54d24cab4cc513e91fc996957b546ad105332acfb35acf77f83fce7"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

<a id="canonical-43c679f311b763cf2baaaac126d415ce7cb5fccdf08d81fe4a6dd16cebc553ea"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / af57ea650307 / 3

<a id="canonical-8f2ceeeb547e81b8d28fa95919f690097c9426c6d68039be78b8d6425fa357af"></a>

<a id="canonical-275bb47693d7b09ff4ba1266fb9dc56db5b2aff67d8e8c75c3293a22c77c42da"></a>

## configured_address property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / af57ea650307 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-02b11b6fdb032c8553b620090fa316b9850cfc5d90839c185269139e9febcca9): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5c91e9fa41d2e6e19e076ed2f15dbacce066c186ab6f1a30c87754eeb3d3917e): complete subsection reference.

<a id="canonical-70e4c2c1f80f1fc3b4c9db7482d4057d7edcad2650709bf3dfe76e8a0647698a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / af57ea650307 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-02b11b6fdb032c8553b620090fa316b9850cfc5d90839c185269139e9febcca9)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5c91e9fa41d2e6e19e076ed2f15dbacce066c186ab6f1a30c87754eeb3d3917e)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-02b11b6fdb032c8553b620090fa316b9850cfc5d90839c185269139e9febcca9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-334112e917bf8af640fd86b564a9317281da0ec4ad43effd38db3f4384503ff6"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 7725219ede47 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-a50320bd4533a7e259e8393c8b1349b93cca26ca43b184facf31cc1d1ec33783)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-ca150b7a0ef5bfc3eee6f17b7fc6af6f0cbfce9738fbfa49d582d0152d690889"></a>

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

<a id="canonical-a2ae6a0957f7164e89d3c8d6700ac8e9bf0f94d2fae6c03da79a26f308c9cfb4"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 7725219ede47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8ec39c05537fca47674546122ef28f69c7e39004d6ad8e4b1d599bb897f1f5bb"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 7725219ede47 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-a50320bd4533a7e259e8393c8b1349b93cca26ca43b184facf31cc1d1ec33783)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5c91e9fa41d2e6e19e076ed2f15dbacce066c186ab6f1a30c87754eeb3d3917e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3462669e93b20b535de2c9421bea452addf8238a94b489bf8f1dd38263850f50"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 21eb33359cb5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-083b2b1501df6e16b959ccd5cfb3afab7608e0f95cf94f254de0b641187730ce)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-a50320bd4533a7e259e8393c8b1349b93cca26ca43b184facf31cc1d1ec33783)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-0a2ea346fe2955d70d2382a7856b64dd05c709dae84769a138a61868add0bc85"></a>

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

<a id="canonical-5609fa43e269691a7ee68346fc5448707e7052b808eec4a92e794dee4b507642"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 21eb33359cb5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78d88671faa30188fa794640c696707c60ae61d01ec5fc22722c64d994b7aea6"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 21eb33359cb5 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-a50320bd4533a7e259e8393c8b1349b93cca26ca43b184facf31cc1d1ec33783)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f817956253e6e4d719c6c2ebdd67b3d8913126f4fc45bc845f042e9581b1dad5"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9b6391abd56d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-7ef3719523754be8a5aec932c64f048aaa84f15dd78afa02a83b2abb3e4ac0f1"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-e9f8c737ad2b95600eb83b7368505de98737e6951a6277d886e7acc8a7cd3d95"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9b6391abd56d / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-008.md#canonical-8cebc6ac6c851dcac9089d58e15ef1793cdb450239a9e0ffccec7d6ecec1287c): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-008.md#canonical-13ae7698bc4d541a79d9c7e38349b9f8a40d22c7bab736f98c621f79cfec8373): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-19eaf01caa35adc0d06b8149fdfabe074c8e8a1ded5d84f4bdc439890e53a22f): complete subsection reference.

<a id="canonical-b0bbd45ceae3db58f31febc23c6d484957ae66a22d7e43d8948574475b28e692"></a>

<a id="canonical-37d2630c3b4ce0b470fc983699abee2e4b3b5a595e27d728db4cd24d407adc68"></a>

## fixed_ip_map property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9b6391abd56d / 4

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-008.md#canonical-dfb128e85c46b0085ef7fe18518be57c2ed96b53d66e63416491bffccc6773f0): complete subsection reference.

<a id="canonical-cd433ec1ce8b5018431e8bf770dcc5d0a2a65813de46533cd1be84f124865f4f"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 9b6391abd56d / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-008.md#canonical-8cebc6ac6c851dcac9089d58e15ef1793cdb450239a9e0ffccec7d6ecec1287c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-008.md#canonical-13ae7698bc4d541a79d9c7e38349b9f8a40d22c7bab736f98c621f79cfec8373)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-19eaf01caa35adc0d06b8149fdfabe074c8e8a1ded5d84f4bdc439890e53a22f)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-008.md#canonical-dfb128e85c46b0085ef7fe18518be57c2ed96b53d66e63416491bffccc6773f0)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8cebc6ac6c851dcac9089d58e15ef1793cdb450239a9e0ffccec7d6ecec1287c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26747580802fd786b2fd4ea29d1da4d8b1ba677aaa12cb42f34f3e8335d1ee56"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 8e7cde87769a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-bc077d413891b3aa176d0e54d4603d7038c6ae1b7964be54fe86250286cd99a3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-969501e4dad0a30f517cee4b3317b19bd21da6db60c0c51ca82ad893b674844b"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 8e7cde87769a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-920bfa857b780f13738e15386ec3fb0b2c4b0d09d0c3b7288849f78a86052344"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 8e7cde87769a / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-13ae7698bc4d541a79d9c7e38349b9f8a40d22c7bab736f98c621f79cfec8373"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70c7db7d7275d1fea398e2494ac12e8bd00ac76de598c413e44d634ec1e5549d"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 6553c001e303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-77e70a555281046c42afccfe3f12413c3ea6f91724480adccadb812308098342"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-ba5fb3bdc3c11b40e568673c0d15804a30b1abf8ad1c26c3f7e631c40209282f"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 6553c001e303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-93ba3f6f791bf1512e6e1ed819a4f344db71469b7bec5143fe7824fa2a39038c"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 6553c001e303 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-19eaf01caa35adc0d06b8149fdfabe074c8e8a1ded5d84f4bdc439890e53a22f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03e23b808942fe0b11be4ea8783baf883ed7a520722e163d7dd6c79e8fe2b35f"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f267367b41dc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-ba3ac6b98d2e0794708577544b99c9026d1eb288da61ce1b9e76fdb19f1f7908"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-09721a051bc1ed1c6fe2439cd7967b02d5061f0b44681ba0aea02153438761f0"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f267367b41dc / 3

<a id="canonical-a4ce495576c5bbf6474d49cd3a85d70483c2c7cf4c653b06141fd11b78e00ae4"></a>

<a id="canonical-7ee4a0ecb7d00be992a700dc516789b2956eb9b28dff0289dc2abb304232010f"></a>

## network_prefix property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f267367b41dc / 4

Type: `"string"`. Computed.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-b9c0b7f417d871cd1686f5ad6e7e298013ed1d1c71f3bbc561e4fb0c6530f394"></a>

<a id="canonical-08f027f826c38d7b666b65adbec782a08b1df83837f2fb82e2ba4f06591cbb5e"></a>

## pool_settings property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f267367b41dc / 5

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--securemesh_site_v2--reference--group-008.md#canonical-ba0b614c2175be2cbc724f7728ea95b3ecabaa1732d747ec8d2b38ad1dcbfbe9): complete subsection reference.

<a id="canonical-a466daf4048edb2612f2be180b0738cecd3f7d3e00390c2bd7e5f4824bd3fd40"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f267367b41dc / 6

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-008.md#canonical-ba0b614c2175be2cbc724f7728ea95b3ecabaa1732d747ec8d2b38ad1dcbfbe9)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ba0b614c2175be2cbc724f7728ea95b3ecabaa1732d747ec8d2b38ad1dcbfbe9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef57ab39c011c15bf34dd4871834c35cc10447dbd7d0cb5de7170e3dc4c3d5de"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 393f6bda7da1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-19eaf01caa35adc0d06b8149fdfabe074c8e8a1ded5d84f4bdc439890e53a22f)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-a1dcd59014a436f9fbb6a2b687c7226d4bcd50587a282bada96fd09eac61135a"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

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
    "minItems": 1,
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

<a id="canonical-515a3bf4faf23cdffc1245efb3daa4ff2d8514a3355345847525d3b9e29d1fcd"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 393f6bda7da1 / 3

<a id="canonical-4904ec233ef8e7e2ce4ef2bd1ffb5d3f98b9e1b57912426b8244c57c376568b5"></a>

<a id="canonical-0df7dc0dbd51072183a519bf6df71d4767075fc8ab40cc944902583ec451162e"></a>

## end_ip property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 393f6bda7da1 / 4

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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

<a id="canonical-fe1300ecc5767bbadb1f85a863336dc2558d7b47a2afdacee7974d4651deb0da"></a>

<a id="canonical-dc309bfc5af56d6291fdad25eb1abe9a2f51741acc5a671015cc74a69a61052a"></a>

## start_ip property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 393f6bda7da1 / 5

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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

<a id="canonical-b0b6e8bbc76f4df210a90be15fa0613542227e685b8ca5cd60d320fcfd05786c"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / 393f6bda7da1 / 6

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-19eaf01caa35adc0d06b8149fdfabe074c8e8a1ded5d84f4bdc439890e53a22f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-dfb128e85c46b0085ef7fe18518be57c2ed96b53d66e63416491bffccc6773f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8390889830df4bfcecdf545baaf990a4bc0a24988f0fab3592194c3bb701ffcd"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / b86a4f5f8090 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-36b0761d317957cdd944d3e39d417f2945b2ef8b9cbc9cf957f91ff053d85a97)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-60addeb2190fadb87026c5aab1d72fd488562dd4a64b31f623b12d3eb3a91016)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-564367668901c89562b605423e10168f12aa852f655e9651d5e411b935cca38e"></a>

Type: `"single"`. Computed.

Map of Interface IPv6 assignments per node.

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

<a id="canonical-c3157a95047c73b7f23e103f58a0bcc34f1e22a3fe74d26387a4f925710ecb96"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / b86a4f5f8090 / 3

<a id="canonical-a56194729bdacc50d0b329072be41ddeb4750ae8b5602f039dc14cf5f2a393be"></a>

<a id="canonical-3365df64ac7b9e6d216c3545de2faddca2976e8d8c1dd0e18e1b73f469ec7680"></a>

## interface_ip_map property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / b86a4f5f8090 / 4

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-4594ddd9a37c8045d5a38fbe70c9c9ad2b000ba022c1395f8afa1c4ced54231b"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / b86a4f5f8090 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-f992fd7cad548b5c870c9f65dfea17fa9fba85b390f1ba41b54827a055e214cc)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5427fc0c20d6c5d9e94a2491696e0cfd31d3f15be75d030a01b0ed122c3e3e44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e00874ac220f7b274558a52e6d4b67c27551f495b0c3e4c995bfe814cfc8aa43"></a>

## equinix.not_managed.node_list.interface_list.monitor — equinix.not_managed.node_list.interface_list.monitor / 9c9649661b2c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.monitor

<a id="canonical-68244b0df33db574b58f1d386d84008bce7079bdfa704dca5cef40ae4ec35195"></a>

Type: `["object", {}]`. Computed.

Link Quality Monitoring configuration for a network interface.

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

<a id="canonical-2ecee6cdbbef0b22d9d6c1620526cdcdacf992e8f7f83cb319db5a6932788659"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.monitor / 9c9649661b2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53d8165917e3f147950300ae2dd160b18a718ccfa1e96b748b99a8fdaf09152d"></a>

## Next pages — equinix.not_managed.node_list.interface_list.monitor / 9c9649661b2c / 4

- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7bed8611b98a0bad14f1472c9173640f72d47cecab5a2dfe1f6797652ada860b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3be05bafa9b46e704a1b009a9781a86f76ed3790e5603aaaf26bd2d9dbeababf"></a>

## equinix.not_managed.node_list.interface_list.monitor_disabled — equinix.not_managed.node_list.interface_list.monitor_disabled / 4794a3200890 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-d4de9cc088b00dfe97c6ee8e8600bdb65f04d74b6b8936ed1234c0ad2c4a2547"></a>

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

<a id="canonical-9ea3c91a735fb7797e07c3e4bc293147ceef3dff1c0c12d9f643054157fb3dd0"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.monitor_disabled / 4794a3200890 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d023a99fc248b7b95896784b88a6f9bd69e6c2a1911a229815d16215117c13e7"></a>

## Next pages — equinix.not_managed.node_list.interface_list.monitor_disabled / 4794a3200890 / 4

- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bb89ee17b354d7dcad9838f440e927d5a3ab48b5ad2d1067cf53bf9742d7229f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58a6bb7eb3b8281c0ab5264c577b47a558e091cb45f0b911130c5ad96431f641"></a>

## equinix.not_managed.node_list.interface_list.network_option — equinix.not_managed.node_list.interface_list.network_option / 39ecc44952d8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- equinix.not_managed.node_list.interface_list.network_option

<a id="canonical-70ac362f1606c6e72f94246597bf39fd132282128a6906fa9ef23d7982b120e4"></a>

Type: `"single"`. Computed.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

<a id="canonical-991a1cea682bb0a9a828dde2362b176a8f831d1673c2214e51fc065ae4f01f9f"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.network_option / 39ecc44952d8 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9a6aad6969820285a30fc7069534f35baa1904c02f8899bd9710519b90c15538): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-71f67d3a6e79782b83caeb8b0e9b2ff34c67748c0511037cc1c5f687bf614396): complete subsection reference.

<a id="canonical-aa1dead34fd46b9507752f9d082197ba68019f83541010655880e0d580f18511"></a>

## Next pages — equinix.not_managed.node_list.interface_list.network_option / 39ecc44952d8 / 4

- [equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9a6aad6969820285a30fc7069534f35baa1904c02f8899bd9710519b90c15538)
- [equinix.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-71f67d3a6e79782b83caeb8b0e9b2ff34c67748c0511037cc1c5f687bf614396)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9a6aad6969820285a30fc7069534f35baa1904c02f8899bd9710519b90c15538"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d621c145df13fb33dcd1229d60bc19666430f35d61278e99d92e9137ec04a516"></a>

## equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network — equinix.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 6b907bcd0c4c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bb89ee17b354d7dcad9838f440e927d5a3ab48b5ad2d1067cf53bf9742d7229f)
- equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-141e38cbaf9f26f94da6303fd39cd40baf10e6a3faf3596fe2ef9cbac5e8c737"></a>

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

<a id="canonical-a007a72ee0d5d84fdc5a2195f084987069f4245bdcbbc73a239452bddda69008"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 6b907bcd0c4c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1790fc9fd1f5434a5bf903753733cc7857d46ec8284a430e0e900b7d0a081421"></a>

## Next pages — equinix.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 6b907bcd0c4c / 4

- [equinix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bb89ee17b354d7dcad9838f440e927d5a3ab48b5ad2d1067cf53bf9742d7229f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-71f67d3a6e79782b83caeb8b0e9b2ff34c67748c0511037cc1c5f687bf614396"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cb13ffc6cf29274def9efbda94e7df9f742f6d8ac0009df9c1be7456032f9bc"></a>

## equinix.not_managed.node_list.interface_list.network_option.site_local_network — equinix.not_managed.node_list.interface_list.network_option.site_local_network / bc7011c96c93 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0cd2f3f6376b9c4dc4b518d2dc357fb02984629062f267c297004d089415a383)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-9544967a50f5566a6932a9ca2ae7dc73a426c90c8f638f70f7928dbfd38837dd)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-7b644edd1301eaa381b8b780eb33a34e8a1012714d0cc450f43354e4004ccd0a)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-5b6302e2bc7554d6932c78047b194b6cae2e44283a922fe886e86651ec44a067)
- [equinix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-008.md#canonical-bb89ee17b354d7dcad9838f440e927d5a3ab48b5ad2d1067cf53bf9742d7229f)
- equinix.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-d8ca89281feb9797f737c5e009be804c3b6404bb3df3f026ead17dd6d6c85760"></a>

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

<a id="canonical-02e195632a84bf0b5e3d5c427cf7087e95552d5323e220ccfbfafb19be951069"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.network_option.site_local_network / bc7011c96c93 / 3

This is an empty object or choice marker. It has no direct properties.
