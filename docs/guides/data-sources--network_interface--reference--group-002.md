---
page_title: "xcsh_network_interface reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface reference."
---

# xcsh_network_interface reference

<a id="canonical-8eb38d4feb4f316e45b3a232e7a4f121492116336ed8495d52ca86c1b08b11be"></a>

## Next pages — ethernet_interface.static_ip.node_static_ip / 380b2fecce0b / 7

- [ethernet_interface.static_ip](data-sources--network_interface--reference--group-001.md#canonical-f6a7a273d85d7cfa1c8614e9eecb3589eb3848255af051b6a00adcd274fd76ed)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-9a76ad147e80b2c51535b56cd403c93b9e85bbc9aa57014fc100c9bb299ff1f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41a3d8cbd1a15d4589459f67d25d0c8ca933058ebb7c1a06b275727621269d7e"></a>

## ethernet_interface.static_ipv6_address — ethernet_interface.static_ipv6_address / ffc2a01bc05b / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.static_ipv6_address

<a id="canonical-86190bd308c892284953123125e23154cf85a8118f24d845f5aec209bb3290dd"></a>

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

<a id="canonical-f8b4f8ec643e30cfe092b7342e65a383394f5ba2d29db9a95756409027ce601c"></a>

## Direct properties — ethernet_interface.static_ipv6_address / ffc2a01bc05b / 3

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-80fce281fadfaae3477d0e38cfd22793a35fe66bba2033965d600d289e6161f6): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-b28bb335f24385838fbace525157105468f2ca00ed9210593b4f6f7515218b46): complete subsection reference.

<a id="canonical-7212381e685825599e4d5c5ea6751ef30111fc7ef4744132b61efe89a892e41c"></a>

## Next pages — ethernet_interface.static_ipv6_address / ffc2a01bc05b / 4

- [ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-80fce281fadfaae3477d0e38cfd22793a35fe66bba2033965d600d289e6161f6)
- [ethernet_interface.static_ipv6_address.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-b28bb335f24385838fbace525157105468f2ca00ed9210593b4f6f7515218b46)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-80fce281fadfaae3477d0e38cfd22793a35fe66bba2033965d600d289e6161f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d8a39536a696e0c504a4be43cdef20ac58ac9c293831cb4d0ece6406ef9c492"></a>

## ethernet_interface.static_ipv6_address.cluster_static_ip — ethernet_interface.static_ipv6_address.cluster_static_ip / fd827ea9dac7 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-9a76ad147e80b2c51535b56cd403c93b9e85bbc9aa57014fc100c9bb299ff1f4)
- ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-9d10e91331d1012afdd8e980d8eeb35801a48aa81a393d21d7f2aefaa314245d"></a>

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

<a id="canonical-cc84b5bdba41630999b8b5a2477bf9c71cb4f0f338152602fdf3556dac8a19e2"></a>

## Direct properties — ethernet_interface.static_ipv6_address.cluster_static_ip / fd827ea9dac7 / 3

<a id="canonical-a5308dbc5f6a7372182239dfeb7b4f362b6a5e1edc28cd2dfc4d765c23afae68"></a>

<a id="canonical-72c57ccf9767f9634c54a98e2307c3804cb4483145f6af5895817c5a83168b5f"></a>

## interface_ip_map property — ethernet_interface.static_ipv6_address.cluster_static_ip / fd827ea9dac7 / 4

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

<a id="canonical-c0165b0fd34d190271eacbc92db0b621dfe23c0fcd50398fd1fd79f1d5471ae8"></a>

## Next pages — ethernet_interface.static_ipv6_address.cluster_static_ip / fd827ea9dac7 / 5

- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-9a76ad147e80b2c51535b56cd403c93b9e85bbc9aa57014fc100c9bb299ff1f4)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-b28bb335f24385838fbace525157105468f2ca00ed9210593b4f6f7515218b46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a45466a292e0a650712467ef880cd757fd3006f3f01b9cef6063d061a95fcc4"></a>

## ethernet_interface.static_ipv6_address.node_static_ip — ethernet_interface.static_ipv6_address.node_static_ip / 6b37be6e8d19 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-9a76ad147e80b2c51535b56cd403c93b9e85bbc9aa57014fc100c9bb299ff1f4)
- ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-4e094e28389305c09b9ee16dd232f1089081aff3263b393d662c4ba4159518c8"></a>

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

<a id="canonical-a187d66d93191c442b7bf22b659f86d45b75bfe4272a348a5fbaf5456eafb59d"></a>

## Direct properties — ethernet_interface.static_ipv6_address.node_static_ip / 6b37be6e8d19 / 3

<a id="canonical-809520a2997efbe19d87d992581ed570bee25918b8715d4492ddb8a9e20a5ef2"></a>

<a id="canonical-4619f4397606f00953143a89cc53a0cb38f12f8719dd0356a4fe7ffea8d1c21a"></a>

## default_gw property — ethernet_interface.static_ipv6_address.node_static_ip / 6b37be6e8d19 / 4

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

<a id="canonical-83ae3ccc28db1dc1749925a132188110d0b69120b67cef72df705e68e7a26c5f"></a>

<a id="canonical-29ef91a9dbc4d7e2a942ca099e93347806fecab01f12b7067a48439c297f1247"></a>

## dns_server property — ethernet_interface.static_ipv6_address.node_static_ip / 6b37be6e8d19 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-231c2ad72141a37118c291802237a8d4328303ad83f01537d7faab39b9a481e1"></a>

<a id="canonical-35f2600af1d44e68a31e2a7e964a5e6059c496bb949dad439a8a12e12bf8004b"></a>

## ip_address property — ethernet_interface.static_ipv6_address.node_static_ip / 6b37be6e8d19 / 6

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

<a id="canonical-90a707970e49346c76ca04d061c86c3a377d99ca68df7829b9b18bf023c07d43"></a>

## Next pages — ethernet_interface.static_ipv6_address.node_static_ip / 6b37be6e8d19 / 7

- [ethernet_interface.static_ipv6_address](data-sources--network_interface--reference--group-002.md#canonical-9a76ad147e80b2c51535b56cd403c93b9e85bbc9aa57014fc100c9bb299ff1f4)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-3f3e41399e956ac76c0e4c72fc7f2a6417f8c1cc7de2f968c11e0059a3d8411a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4166ff528938b00b7330e0c2322f7e79e790babf65299c8117ec7e02f1032f3e"></a>

## ethernet_interface.storage_network — ethernet_interface.storage_network / 05a3dc8a4ed9 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.storage_network

<a id="canonical-aaa72a8a6c3e64516aefd2801c0e422a127e9315fdc0a5fc047c86fbf46edef4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for storage network.

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

<a id="canonical-0e5ac3da72f070af1fe2463117daf72f404e7d460f1fa6ad7242bb9b3dfe6f43"></a>

## Direct properties — ethernet_interface.storage_network / 05a3dc8a4ed9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26ebae5bfc1d01a7eb42e130e821053a7aa7fd1ac1baf75fa6a3e2df33bc3596"></a>

## Next pages — ethernet_interface.storage_network / 05a3dc8a4ed9 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-b3402af0a5d18205df42bc79bbeb04d66cf0db5fec4c9f1f4bea5d4601a46b76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5de7c560e263f382e9f004d65d54d27322505eaa0cc690036598e9f17e4ee3ae"></a>

## ethernet_interface.untagged — ethernet_interface.untagged / 526b6ca98673 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- ethernet_interface.untagged

<a id="canonical-fe8ae3fbcba88305ec8356fdc919a3be1d1283017e35bbc7d8ff2c8c5151c112"></a>

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

<a id="canonical-45279d7aa3ba10f6e996b1644027fb0369102887a838c77c511cbc935cdb5298"></a>

## Direct properties — ethernet_interface.untagged / 526b6ca98673 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b4eb2eb4ef1ca598a05fa0fa1e58b7cbdf80ee14b2d9a8a0a58e385a082bc3e"></a>

## Next pages — ethernet_interface.untagged / 526b6ca98673 / 4

- [ethernet_interface](data-sources--network_interface--reference--group-001.md#canonical-6da5e3e26326c09c2b17a5025193f1ac4df58638e4431ff4c3b6cfd13ebb42d6)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbede366373d537c961196620bdc695594e5a4a4167dd4e7fa65578943ad9ede"></a>

## layer2_interface — layer2_interface / 2c5754332ca9 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- layer2_interface

<a id="canonical-e25c4b23e2e1ec1252a64394f70c4aac2e7034a01010527a43b3b66f0f420ceb"></a>

Type: `"single"`. Computed.

Configuration parameter for layer2 interface.

Upstream description:

Layer2 Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-layer2_interface_choice": "[\"l2sriov_interface\",\"l2vlan_interface\",\"l2vlan_slo_interface\"]"
}
```

<a id="canonical-cb8c9c4b627319f5b9daa4596d58e2738e79ef90eb20701c1817e3fe4ec14f43"></a>

## Direct properties — layer2_interface / 2c5754332ca9 / 3

- [l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-b9bccf49a78eb986b0afad4bdc7ffcea25cb9e58aee6d828b214cd30e9fccc49): complete subsection reference.

- [l2vlan_interface](data-sources--network_interface--reference--group-002.md#canonical-3cf182136ac064cc6ebbd841b31abd806bc93c00695f769b47132c1e125f036e): complete subsection reference.

- [l2vlan_slo_interface](data-sources--network_interface--reference--group-002.md#canonical-7420e477cd8597c56c13c9f8aa7dbde82a60cf7169a0289170a2ade7d1f98f8a): complete subsection reference.

<a id="canonical-e8f46f57044aef015bf3b23ac4484a8ae06ab6283c2769896ea7c89c8fdf273b"></a>

## Next pages — layer2_interface / 2c5754332ca9 / 4

- [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-b9bccf49a78eb986b0afad4bdc7ffcea25cb9e58aee6d828b214cd30e9fccc49)
- [layer2_interface.l2vlan_interface](data-sources--network_interface--reference--group-002.md#canonical-3cf182136ac064cc6ebbd841b31abd806bc93c00695f769b47132c1e125f036e)
- [layer2_interface.l2vlan_slo_interface](data-sources--network_interface--reference--group-002.md#canonical-7420e477cd8597c56c13c9f8aa7dbde82a60cf7169a0289170a2ade7d1f98f8a)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-b9bccf49a78eb986b0afad4bdc7ffcea25cb9e58aee6d828b214cd30e9fccc49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e6858fb066122576a7c4b09ae53ae6e032c723be36b18b32d662b79c3e3a69f"></a>

## layer2_interface.l2sriov_interface — layer2_interface.l2sriov_interface / 7fa72b774849 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- layer2_interface.l2sriov_interface

<a id="canonical-de8091034db81fd6f72be817732751028b91672188945e42bf9d4f12a32d5886"></a>

Type: `"single"`. Computed.

Configuration parameter for l2sriov interface.

Upstream description:

Layer2 SR-IOV Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

<a id="canonical-646ee83c4e131dbece8ce84e8452cdd9f01081cdc0ee24d081f75dac02dcd552"></a>

## Direct properties — layer2_interface.l2sriov_interface / 7fa72b774849 / 3

<a id="canonical-563f58745d2099c61019eb401c116dfe0d497b6df9172f5cfcb1a5faf8eb3123"></a>

<a id="canonical-4fa8a264d37da7f5ccc71b7bbc59b149e4294e586b79b010f11e24774233cecb"></a>

## device property — layer2_interface.l2sriov_interface / 7fa72b774849 / 4

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

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

- [untagged](data-sources--network_interface--reference--group-002.md#canonical-932355f9cdc9584a69aa916b95ab543d0649196266c67001f1ed3d4424e54ba7): complete subsection reference.

<a id="canonical-7cff783164227e29be43971bd10b5bc399a9f605a966bc527b117df227fd5b90"></a>

<a id="canonical-835069997179d25e3436d1712eb28046cc0068d50246a2b769d7b8e131bbc485"></a>

## vlan_id property — layer2_interface.l2sriov_interface / 7fa72b774849 / 5

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged interface.

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
    "create": false,
    "minimum_config": false,
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

<a id="canonical-10e7aa1ce43e6d3b066b6634aefcbcae26f29a592e5bf1aa3fb6f5c8f20253a0"></a>

## Next pages — layer2_interface.l2sriov_interface / 7fa72b774849 / 6

- [layer2_interface.l2sriov_interface.untagged](data-sources--network_interface--reference--group-002.md#canonical-932355f9cdc9584a69aa916b95ab543d0649196266c67001f1ed3d4424e54ba7)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-932355f9cdc9584a69aa916b95ab543d0649196266c67001f1ed3d4424e54ba7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45c2f65fc6ef70531777929db44cdf7a04f9dbb0da102d276a9c6d7a140bc16a"></a>

## layer2_interface.l2sriov_interface.untagged — layer2_interface.l2sriov_interface.untagged / aa1e159c95e5 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-b9bccf49a78eb986b0afad4bdc7ffcea25cb9e58aee6d828b214cd30e9fccc49)
- layer2_interface.l2sriov_interface.untagged

<a id="canonical-85b06d296fc171d55fb3d610af76da9119fa7696e7398c4c526b35b20206d3ea"></a>

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

<a id="canonical-ff63bf2484902e9cc3c3093e6b72b5a7edb4e7a9364d0f7786bd8f4761d358b6"></a>

## Direct properties — layer2_interface.l2sriov_interface.untagged / aa1e159c95e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-942f443a9c24232257aa8e5d2a36f92b3932d216d6fed2e5dd45f4bb6f88e221"></a>

## Next pages — layer2_interface.l2sriov_interface.untagged / aa1e159c95e5 / 4

- [layer2_interface.l2sriov_interface](data-sources--network_interface--reference--group-002.md#canonical-b9bccf49a78eb986b0afad4bdc7ffcea25cb9e58aee6d828b214cd30e9fccc49)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-3cf182136ac064cc6ebbd841b31abd806bc93c00695f769b47132c1e125f036e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2fdaf27bfbfe800efb80b53c197c623f06c1568da8992b6167b8bb8f0053523"></a>

## layer2_interface.l2vlan_interface — layer2_interface.l2vlan_interface / 041a9787ed00 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- layer2_interface.l2vlan_interface

<a id="canonical-2e684f5ff27b6ff85205c9274e35ea3fc52289b5fd480adf4968919f7e3cd424"></a>

Type: `"single"`. Computed.

Configuration parameter for l2vlan interface.

Upstream description:

Layer2 VLAN Interface Configuration.

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

<a id="canonical-1981a7fbd5c51149eb5f1c650a87e2e8feba05f53b1b149013ac8cd247483309"></a>

## Direct properties — layer2_interface.l2vlan_interface / 041a9787ed00 / 3

<a id="canonical-172f0c411c33895503d64a66e56d1132a4684337e357d3fac2821f6a6c6fb940"></a>

<a id="canonical-d82d98e281118b7913ab0cd92f0bd5ac688e89c97e145bd67491d142ed8e9cde"></a>

## device property — layer2_interface.l2vlan_interface / 041a9787ed00 / 4

Type: `"string"`. Computed.

Ethernet Device. Physical ethernet interface.

Upstream description:

Physical ethernet interface.

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

<a id="canonical-0a29bba28098c7b2206fff78e66f9e082c978fd2b99393d67e990b0ad16eb13d"></a>

<a id="canonical-4acef8a1717b720abf829bef31ed641d7a1998c2c1267b8c08aa529fb000db15"></a>

## vlan_id property — layer2_interface.l2vlan_interface / 041a9787ed00 / 5

Type: `"number"`. Computed.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-d2cc5c3b7fe273e8dea7b78044e8933c1163bc0b2dbfd539086e8713850d5cf4"></a>

## Next pages — layer2_interface.l2vlan_interface / 041a9787ed00 / 6

- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-7420e477cd8597c56c13c9f8aa7dbde82a60cf7169a0289170a2ade7d1f98f8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ab63c01664fe2c293824b7451f3eef076c173198969a5f385fd04d561edf50d"></a>

## layer2_interface.l2vlan_slo_interface — layer2_interface.l2vlan_slo_interface / b95dfaf3b934 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- layer2_interface.l2vlan_slo_interface

<a id="canonical-a76b9c4cb3f24fecfc400619b2fa619033228f6f8ed2047ac0e387fd3350225e"></a>

Type: `"single"`. Computed.

Layer2 Site Local Outside VLAN Interface Configuration.

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

<a id="canonical-ea2720cfbe8e08e0537bbe3c65f82fe624625d79c533dd153b7d1c58dd48f741"></a>

## Direct properties — layer2_interface.l2vlan_slo_interface / b95dfaf3b934 / 3

<a id="canonical-975c0ff9b818e5b3dfb430cc3daeb9944a1aa8f0a86e2c212ae2a80a3ec85d55"></a>

<a id="canonical-b9f196a569b8f1f82e3dd7ac267880e65ec443961132df039bb3fe5667c14b88"></a>

## vlan_id property — layer2_interface.l2vlan_slo_interface / b95dfaf3b934 / 4

Type: `"number"`. Computed.

VLAN ID. VLAN ID

Upstream description:

VLAN ID

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-2ff982e3fbf7319fff6b5aadea737eb54c55de159517ca526f104c6d30520e56"></a>

## Next pages — layer2_interface.l2vlan_slo_interface / b95dfaf3b934 / 5

- [layer2_interface](data-sources--network_interface--reference--group-002.md#canonical-23fb46783886bf9cc3aab8b71636b36b489f59cff700e25eee93f42b2c95b70c)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1dcbd9ad550b47d24989242f5b6079c1933800f3d0737c283cd193c3f56b9e7"></a>

## tunnel_interface — tunnel_interface / 9d5f1ebf9092 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- tunnel_interface

<a id="canonical-2269d485ee9d6abbab5bf7aa51105026e13fedb756d7e5e7f41e68c320cf0cef"></a>

Type: `"single"`. Computed.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

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
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

<a id="canonical-828e2a8b5894a04e97cf566db332730dd89d46ab2f97e4e4d848f5bd9b89bd1d"></a>

## Direct properties — tunnel_interface / 9d5f1ebf9092 / 3

<a id="canonical-b16af4d8d658af6a2d6ce94ed4d95227a38d37ef20053ddb91cc794880697d2e"></a>

<a id="canonical-abd8069543d7d2c532a76ac5f9d384c62b9cd34a6117311d9b554bf54f1a53d7"></a>

## mtu property — tunnel_interface / 9d5f1ebf9092 / 4

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-affaa248268d9aa661fe2672faf0f9d1ab03a30897a5e02bb9c08cb6767f77e1"></a>

<a id="canonical-e8065f0c9e6dc8f88d8ed7e8cf96bc953d4a670a5a6b399976e9925f82fe3c5d"></a>

## node property — tunnel_interface / 9d5f1ebf9092 / 5

Type: `"string"`. Computed.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-9a463d81bb6733b17adbaca9fad7dbbf2ce260e621600eb998569bfd370eafef"></a>

<a id="canonical-4bb1c7d7779128dcf9fd08192a411ef3f9f93be14c5f8eb4a71f1409166988d6"></a>

## priority property — tunnel_interface / 9d5f1ebf9092 / 6

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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

- [site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-e0f9d70ad880b1ff82963e5c7ad0ebf343cddc9811ce1a27aa054348a9cde3d0): complete subsection reference.

- [site_local_network](data-sources--network_interface--reference--group-002.md#canonical-a9aaaa371339244aee6ee1cc4c8a2e6470fd704a00c8486474c9f1a74242acdc): complete subsection reference.

- [static_ip](data-sources--network_interface--reference--group-002.md#canonical-d998f40f1e817bd55c5f36113ddced31b6809f8412fe5ed13f6d55cdb38d294e): complete subsection reference.

- [tunnel](data-sources--network_interface--reference--group-002.md#canonical-12c7607c94c58ba6117f1102dbdc4d43295086477c6c24a64a68207b6253cd2d): complete subsection reference.

<a id="canonical-8c9ae70b538230ba7ba050b9977e420220d3eb6826952cfb7c3b33b7e52d7aa2"></a>

## Next pages — tunnel_interface / 9d5f1ebf9092 / 7

- [tunnel_interface.site_local_inside_network](data-sources--network_interface--reference--group-002.md#canonical-e0f9d70ad880b1ff82963e5c7ad0ebf343cddc9811ce1a27aa054348a9cde3d0)
- [tunnel_interface.site_local_network](data-sources--network_interface--reference--group-002.md#canonical-a9aaaa371339244aee6ee1cc4c8a2e6470fd704a00c8486474c9f1a74242acdc)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-d998f40f1e817bd55c5f36113ddced31b6809f8412fe5ed13f6d55cdb38d294e)
- [tunnel_interface.tunnel](data-sources--network_interface--reference--group-002.md#canonical-12c7607c94c58ba6117f1102dbdc4d43295086477c6c24a64a68207b6253cd2d)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-e0f9d70ad880b1ff82963e5c7ad0ebf343cddc9811ce1a27aa054348a9cde3d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c6d706bd0b2960468d9c74867dc8531d60f816399fb582c2aea9d4f1c6f63dd"></a>

## tunnel_interface.site_local_inside_network — tunnel_interface.site_local_inside_network / 0baaefb1371a / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- tunnel_interface.site_local_inside_network

<a id="canonical-ddfce1eb4985dad3d94241abc3dbd44edb756ed2ea830471afab9bce2b48151b"></a>

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

<a id="canonical-9c5ef03d9c1109e579461a38c71e112b1a575c6b29d526067e5732d3de15431b"></a>

## Direct properties — tunnel_interface.site_local_inside_network / 0baaefb1371a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c202314496fd69ee5a61a7f295fe1866381c1e3c279c16edf2a0d56884e5808"></a>

## Next pages — tunnel_interface.site_local_inside_network / 0baaefb1371a / 4

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-a9aaaa371339244aee6ee1cc4c8a2e6470fd704a00c8486474c9f1a74242acdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8de1d2542577de66d1490a972bc343077b04b7edb456a94e8a60dc26e8e27d2"></a>

## tunnel_interface.site_local_network — tunnel_interface.site_local_network / 98d1805a2ee3 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- tunnel_interface.site_local_network

<a id="canonical-40da0852fce20e65a6795890ccf3851f55b9ea511f4eebe016f215074d076519"></a>

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

<a id="canonical-277d984423a16f5190df4f5910a94fd4014e1ed47755dba7c67f9a85b13dcc02"></a>

## Direct properties — tunnel_interface.site_local_network / 98d1805a2ee3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b04bb14214c6a011a7819b2c6e46c4d0bf881b66552c8f473733355d2c66c702"></a>

## Next pages — tunnel_interface.site_local_network / 98d1805a2ee3 / 4

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-d998f40f1e817bd55c5f36113ddced31b6809f8412fe5ed13f6d55cdb38d294e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a69669d3e6b6b64e91fa343e7040788920b40d409de7684b898420ea4a9a6f4"></a>

## tunnel_interface.static_ip — tunnel_interface.static_ip / 76eaf389a3a6 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- tunnel_interface.static_ip

<a id="canonical-d759e4dee3392c091389e994fa5b71251ac2d86bf489e5f5b1b782eebe1bfddc"></a>

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

<a id="canonical-2641d7b119979766656d8def7a2a85589d8a654c70e9b26cd8efb62a39d4d0e6"></a>

## Direct properties — tunnel_interface.static_ip / 76eaf389a3a6 / 3

- [cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-00c22acaf24362b322232a94f82acb3de41e34c96de6b57643be8504fe362d01): complete subsection reference.

- [node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-57b51e56f301bd1d80b0361f77124fc69daaf028c0b0a6f44e20f5882b8d9e8d): complete subsection reference.

<a id="canonical-ccfc399594ea79a17bd5e42a3e2e362dbb3df9b7cc82d615b81fda10ad5ca97a"></a>

## Next pages — tunnel_interface.static_ip / 76eaf389a3a6 / 4

- [tunnel_interface.static_ip.cluster_static_ip](data-sources--network_interface--reference--group-002.md#canonical-00c22acaf24362b322232a94f82acb3de41e34c96de6b57643be8504fe362d01)
- [tunnel_interface.static_ip.node_static_ip](data-sources--network_interface--reference--group-002.md#canonical-57b51e56f301bd1d80b0361f77124fc69daaf028c0b0a6f44e20f5882b8d9e8d)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-00c22acaf24362b322232a94f82acb3de41e34c96de6b57643be8504fe362d01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b8058ac51f156b735f29db506d9dc751cb212cfcb6d49f5e5dad0847d8c9d1e"></a>

## tunnel_interface.static_ip.cluster_static_ip — tunnel_interface.static_ip.cluster_static_ip / 59d99160f569 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-d998f40f1e817bd55c5f36113ddced31b6809f8412fe5ed13f6d55cdb38d294e)
- tunnel_interface.static_ip.cluster_static_ip

<a id="canonical-3951090c5704fe4b24b3845979b1af9bc587e36046f2e9660a54a3d7c9858fa4"></a>

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

<a id="canonical-5bb498fd0694d36e3dafe0afb38863c3139b1ddad86d3b851f87888f14771992"></a>

## Direct properties — tunnel_interface.static_ip.cluster_static_ip / 59d99160f569 / 3

<a id="canonical-d17252b54bc679bb9c9fd0710f654221d0f17d369b1972f405a9d0bd5875239f"></a>

<a id="canonical-a7d0940a1404bd0c9f0f23b7b7f5880c817e8124ae0ca53eb35ba253538e108e"></a>

## interface_ip_map property — tunnel_interface.static_ip.cluster_static_ip / 59d99160f569 / 4

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

<a id="canonical-e364d39299a7905fab25e22bfe6fc00f3a9c5bac1a2f3a1778d069a72980bc96"></a>

## Next pages — tunnel_interface.static_ip.cluster_static_ip / 59d99160f569 / 5

- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-d998f40f1e817bd55c5f36113ddced31b6809f8412fe5ed13f6d55cdb38d294e)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-57b51e56f301bd1d80b0361f77124fc69daaf028c0b0a6f44e20f5882b8d9e8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7eedfb386b6ab9e35c18923aac544028c7ac7a5ec2e7200ad893375d67a9a260"></a>

## tunnel_interface.static_ip.node_static_ip — tunnel_interface.static_ip.node_static_ip / 20d7c7cdd475 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-d998f40f1e817bd55c5f36113ddced31b6809f8412fe5ed13f6d55cdb38d294e)
- tunnel_interface.static_ip.node_static_ip

<a id="canonical-75fba8069b09750e46990355a84447034150dad44ad0c2c2d3aac0d9a244bc1f"></a>

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

<a id="canonical-629abbd533cd2cf223a2a27676724d873d39283d188a96645e7450af7dbea533"></a>

## Direct properties — tunnel_interface.static_ip.node_static_ip / 20d7c7cdd475 / 3

<a id="canonical-0c78a0677c44b526d506638e3630ef7ceeeea99eda1bdcd6ddc6572209f368c9"></a>

<a id="canonical-b88511a4c404c2f954ee70a54f4394fdfadd70ed24705b239101074e1f4eecff"></a>

## default_gw property — tunnel_interface.static_ip.node_static_ip / 20d7c7cdd475 / 4

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

<a id="canonical-b365993c73b1fcce6c37f66fcba25f9334059b0ef3324ef83306a3198dde137b"></a>

<a id="canonical-08c8b196a1df870329393913118cb57f0cbb2662eb5a8a80259da2da685ce179"></a>

## dns_server property — tunnel_interface.static_ip.node_static_ip / 20d7c7cdd475 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-6b22ad6f23f80e4a5a4fb31288968511e7034e5ed3d7940915a1037cd63c8418"></a>

<a id="canonical-11c378ed6c83f45b4c6d5d550ea08ec92434c8b32b8a313223126deff507dc29"></a>

## ip_address property — tunnel_interface.static_ip.node_static_ip / 20d7c7cdd475 / 6

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

<a id="canonical-ec85e184c6634a3df73a64f69926b0b5f303a12f7d9bf4a0a98be4095b4ae49a"></a>

## Next pages — tunnel_interface.static_ip.node_static_ip / 20d7c7cdd475 / 7

- [tunnel_interface.static_ip](data-sources--network_interface--reference--group-002.md#canonical-d998f40f1e817bd55c5f36113ddced31b6809f8412fe5ed13f6d55cdb38d294e)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-12c7607c94c58ba6117f1102dbdc4d43295086477c6c24a64a68207b6253cd2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daf4df08f1529a303721001d27f18539de1f68de50af10655aa04198040b4997"></a>

## tunnel_interface.tunnel — tunnel_interface.tunnel / 277ecdc63a18 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Property reference](data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- tunnel_interface.tunnel

<a id="canonical-37d22a39d0a701213a15d6b37e005801363a06fb154abbc1c5d88afa0a13f6af"></a>

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

<a id="canonical-944e104667868ba8a1ffd404fde37113e8bfc7b289111e60b2a05d8c5b901ccc"></a>

## Direct properties — tunnel_interface.tunnel / 277ecdc63a18 / 3

<a id="canonical-ad6cb919c8091cd74ee36c3b3661456de78020eb9a8a38dc8fa8cd87843d5055"></a>

<a id="canonical-f0290eb64e4be745f6987bc2a099de74e7c176529d71fc8b08370b070d6420ef"></a>

## name property — tunnel_interface.tunnel / 277ecdc63a18 / 4

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

<a id="canonical-9ca364562be93c9d662267a70d9bdb137089324fd7b7590d98fa74430fcf48fc"></a>

<a id="canonical-6223da18632ae03481e3336526fcb8fe88f477d4e3b750247eef70f8869f3dcd"></a>

## namespace property — tunnel_interface.tunnel / 277ecdc63a18 / 5

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

<a id="canonical-c5a4034dbd7582f802d4e530cb65507fd8d1368bd2745ef20e3a393799187e7d"></a>

<a id="canonical-518acd374aedf2b1bf6a34e32a46507b6181e9b90681e7aca508c9ec4fe2499a"></a>

## tenant property — tunnel_interface.tunnel / 277ecdc63a18 / 6

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

<a id="canonical-8cc408bfc9c9116688f3f96a9bb9d9a5ab5f290375b138c1159d146ceb83f4e2"></a>

## Next pages — tunnel_interface.tunnel / 277ecdc63a18 / 7

- [tunnel_interface](data-sources--network_interface--reference--group-002.md#canonical-4c4a8d0d91a091335e2c8a7c47c83f1fbdc5b32befca2a5bf820c8b278cccbb7)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
