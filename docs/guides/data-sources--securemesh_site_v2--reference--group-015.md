---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-934f87bd37aba305f15751ea7cb860bf69e115e65006403b335cd093637be241"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option — openshift_virtualization.not_managed.node_list.interface_list.network_option / 57e9dd31ed65 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.network_option

<a id="canonical-79f0ef2c0f72a6e9844992a94ae586766ed8b9ca00ff19961fdaf90ad871cb0b"></a>

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

<a id="canonical-ab92baf236d2da8766657833b1439ee1971181ed6c3684707f61111eef46c95a"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.network_option / 57e9dd31ed65 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-015.md#canonical-bc7dd18f7aa3a6051ce1223fc891ff23fb4c7330e05402bdfe51b15365f2c5a4): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-015.md#canonical-69d52593ad61fa6d23c617b6e02ae352b567d52dfec8fbf6d7a6dc4eccb34c44): complete subsection reference.

<a id="canonical-1c1dce227fd440808df287395df8553dd9b30b8931a707f2c8a7cf173443e743"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.network_option / 57e9dd31ed65 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-015.md#canonical-bc7dd18f7aa3a6051ce1223fc891ff23fb4c7330e05402bdfe51b15365f2c5a4)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-015.md#canonical-69d52593ad61fa6d23c617b6e02ae352b567d52dfec8fbf6d7a6dc4eccb34c44)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bc7dd18f7aa3a6051ce1223fc891ff23fb4c7330e05402bdfe51b15365f2c5a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a2b6d2fb31ad3583369f26f796e9298cbc447b7265a4b856a93941f96b54c2c"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 1f3955b9d098 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-014.md#canonical-47c99d15b875d98500b2f63a9d523068bac29926489a39cb517af71aa3ebae6f)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-af6ec2e2c362b4e3ec10e41ba4f2c7f4096196811bcd0d010a7e2ef3a37cc8d1"></a>

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

<a id="canonical-1ae6db5a213588a6f54297f1f54b6732ab022377adb67d0a2d43f301529db531"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 1f3955b9d098 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9dc1b35987071798fdbb8c21a11904a7b5ecaa7f7ca8000d0aa5e1109151a99d"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 1f3955b9d098 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-014.md#canonical-47c99d15b875d98500b2f63a9d523068bac29926489a39cb517af71aa3ebae6f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-69d52593ad61fa6d23c617b6e02ae352b567d52dfec8fbf6d7a6dc4eccb34c44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36e2ad0da2bdae1affce62f1a94f548041b78a8f6e815758157d9cc629699a76"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 2df2267998ea / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-014.md#canonical-47c99d15b875d98500b2f63a9d523068bac29926489a39cb517af71aa3ebae6f)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-2534eaec3978351aa42c3fdbe52b25142f9bba7735b999a4279260a5c9d46a8c"></a>

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

<a id="canonical-dcca44a34f103b8b85559c70fa41dc24725809699f6e6cc964b7e0e4c2776a92"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 2df2267998ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b97c1e92d4f9bf7466bf78ebb6b14d51bbe40999c8786931db3dd5a4074c53e"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 2df2267998ea / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-014.md#canonical-47c99d15b875d98500b2f63a9d523068bac29926489a39cb517af71aa3ebae6f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-88a3a60a0a8ca4d117f745ec48f4e540a39e72200b272789e6d10df971f85571"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-008b5bfa5e50738d7774de43a526158eeedaf3947b80e130ba43faa1004caa7c"></a>

## openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address — openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address / 382ecb5daeff / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-ce8c4c90423651d99ffeed0a69ad07115ad01abcf60a15c728f843d2599b41cb"></a>

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

<a id="canonical-5e583a08579774096842bb71c699197cff29c364ab2ec550d143f7051201b378"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address / 382ecb5daeff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c373dfd65c3cf7e134286287497e90d12e2bb6ac8d49ccbc6b36985b9aa10947"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address / 382ecb5daeff / 4

- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7ca03389da4ea6ece4145779fddc4eb61d96d661bc86ce131570495507ce3df8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f69237c53022e859520d037ba58ab5b868363bbabbf0ace110779f6fb124b81"></a>

## openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address — openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address / 0c8a0990e609 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-6efa3a1f4d8e66cf55760955ac64917d68adf9983c13266d7c950b00282ca1fe"></a>

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

<a id="canonical-b0ed4801112effedb54e7791af28a87c184b1c0aa0631674095b6285cd13d631"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address / 0c8a0990e609 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8bb2c5df57ab9eef8b3276746f3b9ab278580910deca334218f08ffe3bf51ccc"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address / 0c8a0990e609 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-fda83053d91d0e8c89cbfc0417ef049564a7d4a250883e4c209ad2bc48e3ce38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-159b4b2ee9c96305fdec029a46bf0f4b211ffdc38176996c4fe38ce08d78e467"></a>

## openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 2dab27e5a728 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3ff43154af60530c7871433667bc1f3aaa4de39232e970667af9c15ad6b0a363"></a>

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

<a id="canonical-4efd275d3bfddeb2261ecbdfcadb5c4fbd1331cdbd94de4acf91ea9acf139dfe"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 2dab27e5a728 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb120297f63a5065965348f2b881400109d8bfa1b86e34bb90e57a454f9790b9"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 2dab27e5a728 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9fbe9f4f49f93ddea4a6be1b652851aac0e48f421c97a772ba8e845ee2775517"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b643234880c68c0d105024bb5887ec25c8d807cff070afa783c8dd67ae0bcca"></a>

## openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 47449f6d8c1d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-a6536655fdcba3bb225c524fcafe2e3e6b5a0b32c61cf5ef7b30b23693542c9e"></a>

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

<a id="canonical-edef4b923c6aeef6450e8a89c8dfe6a6dee6f001ca40c232025916a8d5d9111c"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 47449f6d8c1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4179938c0f60a20c05409fd12c1865d3ced0fa857ae595cac91f4b040d6e693"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 47449f6d8c1d / 4

- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-df6dc7bcc94b5998d52ede4049e241bb4e9413fd839e7a821e834be475da4db3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7035276d0f06f14c8d8afc8c60e3c0e47f02cff867810fa4dccd37bf44328e6"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ip — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 785b3d963217 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.static_ip

<a id="canonical-406cee3079c67261baafc6b10fda5113fe2d73d0fb6684f1e34d2facdc3750ea"></a>

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

<a id="canonical-707926beee4a6279ffd57cc6f0d0d667a4eb4465d9130dc01632749f78fe7c26"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 785b3d963217 / 3

<a id="canonical-7bc029dad98c1f30315533101a20b33e8a2413f40a28ca653b56bfe88a9a91aa"></a>

<a id="canonical-a7318531febdea236e7091e19bf5917ce0c2dabf14717de6c79ed52c17cee7e5"></a>

## default_gw property — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 785b3d963217 / 4

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

<a id="canonical-718242e4c4314f0f47ea38d250047afb60a74ac415c6f7ba1b7bb199803c1b73"></a>

<a id="canonical-8c527328ad293689fc6e733677f8cd2242badd82a49e03454746e4268e585014"></a>

## dns_server property — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 785b3d963217 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-ec20d31e7e79140e60c74e734f06148aeb66b54ccce550507a1da66b887a99f7"></a>

<a id="canonical-0374af9de39d9b002990d825f6a5740519274ab2abd53ef8109a133544bf14cc"></a>

## ip_address property — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 785b3d963217 / 6

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

<a id="canonical-5e6cbcaf4f5b46688fe26b256e8540ee6f2c130e7e554399f9e3e3901a481a82"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 785b3d963217 / 7

- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-4d1c78c33d508d314bf1ec3d9d29ef41ec28eb0cbcca29861cb0730a70848542"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acb558242ba400e6675118542f2fd7d9a635fdf26408dd24e5d697766dc5a098"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 61b6bd9ebbd0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-0e4189c478ddb2cead4c74e363d635d09ff7e3c7e30e12632a665cc1f5065ab2"></a>

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

<a id="canonical-01524c622edd13f9956f9b93a506c316119ec8fd9bdbd8d06b361e4b7e55ae9b"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 61b6bd9ebbd0 / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-015.md#canonical-560d37e179ba326376ee2529854fe26a6e86076a1d53853dd24e7de2ddfd6a9d): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-015.md#canonical-6e5b930ece938323735aa1ed21d2e6222f3a5cd1ed6edc0289dab31cd10344ed): complete subsection reference.

<a id="canonical-d263e0ecb9d2aeb1260036e3f692668253aeb514f6dd1629e7262b7d59941b87"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 61b6bd9ebbd0 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-015.md#canonical-560d37e179ba326376ee2529854fe26a6e86076a1d53853dd24e7de2ddfd6a9d)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-015.md#canonical-6e5b930ece938323735aa1ed21d2e6222f3a5cd1ed6edc0289dab31cd10344ed)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-560d37e179ba326376ee2529854fe26a6e86076a1d53853dd24e7de2ddfd6a9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-993837792a46e7cb8b917195633d198641df9e7014878df6af6a0f90942a696f"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 6e2a8521ecdd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-4d1c78c33d508d314bf1ec3d9d29ef41ec28eb0cbcca29861cb0730a70848542)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-69b6e97cf0cc6e8fc6007eaba72215620b6d3b2195b2ff32956d8e202d6836bc"></a>

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

<a id="canonical-ff3fa0653c947b7c09e38b71c2deb8e4edbc442c70cc0c5b904d2447c03a2658"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 6e2a8521ecdd / 3

<a id="canonical-bd3917bf0fcd2c825399e110b363fb230f8c6f75a9ac4f20f4fe89a4ab2b2202"></a>

<a id="canonical-129dec3aef471190204cb1bc4c1b81ca90b64b96011bc69ec00d052aa5d8bf8c"></a>

## interface_ip_map property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 6e2a8521ecdd / 4

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

<a id="canonical-611f70be3efbba927c2b65cac12c582ab1ef3050007e7372555ce63902127e2c"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 6e2a8521ecdd / 5

- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-4d1c78c33d508d314bf1ec3d9d29ef41ec28eb0cbcca29861cb0730a70848542)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6e5b930ece938323735aa1ed21d2e6222f3a5cd1ed6edc0289dab31cd10344ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c2427ff6d670d0bd5f29116df7785a1fb1f21f816dc26a63f29c79ce590d087"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 0476eaa66ce1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-4d1c78c33d508d314bf1ec3d9d29ef41ec28eb0cbcca29861cb0730a70848542)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-fa290f1dd507ec62730fafa3ebab82d1c4d27c344b0fba051418d5e2e9fe4f1e"></a>

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

<a id="canonical-83f71d19fc9ca6bb62a9a7808bd78bed2b14ed0a332cd0279b4e2015eaec4033"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 0476eaa66ce1 / 3

<a id="canonical-f36a007123134852be341266337ab00052c079134fda96b4fc1f9dc9b5b75b86"></a>

<a id="canonical-ffee20a7c9caa79c1199811f297de986601cebcc417e197086b489fe1360a290"></a>

## default_gw property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 0476eaa66ce1 / 4

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

<a id="canonical-e1e4c9cd112abee9e783d93d68f197cad334a4024184bb86bd09e9538e04c468"></a>

<a id="canonical-ab3848a883cc4622f7e01923b04bb402e1ce93d9cf0710d1a52bfdcae200676c"></a>

## dns_server property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 0476eaa66ce1 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-952ecee52490925c5da8b7d31fd205d8d56b6c6a7dc1d80501d9b18cedcb5905"></a>

<a id="canonical-1c1602a4a2a42222d9818c954aa12d68a3cbf825fc178869bad0f35bbd505705"></a>

## ip_address property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 0476eaa66ce1 / 6

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

<a id="canonical-de78ae1ad90f6d7393d2002dd269b80c5684f1068f453d58d66718e4b69e4ae0"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 0476eaa66ce1 / 7

- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-4d1c78c33d508d314bf1ec3d9d29ef41ec28eb0cbcca29861cb0730a70848542)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-606378163f4ce95b35b7b97301be4f657a83820d985feca82ea8e93acacd566e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a166ef0369e9e2ca0731411e46dc5e657f6ddda8ceb27345178c63e23912effb"></a>

## openshift_virtualization.not_managed.node_list.interface_list.vlan_interface — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4d47c0c568bb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-c02cb902cb11cbb24d08944bbb190d4b4f598ac07935807d790ba553815e160d)
- [openshift_virtualization.not_managed](data-sources--securemesh_site_v2--reference--group-014.md#canonical-82213915add19ef5f78381ab977094dd10aa3c7151a9811ed5585af902d0c464)
- [openshift_virtualization.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-f2ebd09282e3c39e55d7748fcf2514ec9a72346979415fc562e4597e09589ec7)
- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- openshift_virtualization.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0d118a71af13263e927d5856512d79dfcb7c6c8d34821132e3f49180da388c94"></a>

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

<a id="canonical-ed4c62beaf0f3ed95193ee5820067911110ef314c1b76c15bb3a721ed2cc6a69"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4d47c0c568bb / 3

<a id="canonical-995eedc704a165659fbafa178e96d2ea34d40587d59cd77319bfa477d43990c4"></a>

<a id="canonical-1e340b62bd3b23add326783bab867d8bf3c8518eab7ec47cd6b18762ef9828d6"></a>

## device property — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4d47c0c568bb / 4

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

<a id="canonical-0c7c8b737114f2b0fab72987f61d16447ba5acac6fa0f45fe19017a63799231b"></a>

<a id="canonical-b04dc5c135753eea6956b871fb646b5da550e728e8e4a49b1301461b0a99c603"></a>

## vlan_id property — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4d47c0c568bb / 5

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

<a id="canonical-bbd7580c34f4a9e615609bf99b881dd6299109fa6b9f1f48e1d3ac069ea870b7"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4d47c0c568bb / 6

- [openshift_virtualization.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-014.md#canonical-7e607b936a94dbd36c13d0aaec3d79e410aac596d1c2a2e191ca869747e2b966)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e31ca504e3bfe1f8cfbcc1e8b3be1b63d7af95172d2286c18c9703d1536d55f1"></a>

## openstack — openstack / 0ddf54a3c37e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- openstack

<a id="canonical-00e035357827e2311c4bfd8c670f4749ae111479fe40c80d5a605af70dfb1207"></a>

Type: `"single"`. Computed.

Openstack Provider Type. Openstack Provider Type.

Upstream description:

Openstack Provider Type.

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

<a id="canonical-debd799649b0a4d1f770f21672c62e2b6e231fe6679031e4e482b27fb4cc1ccc"></a>

## Direct properties — openstack / 0ddf54a3c37e / 3

- [not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f): complete subsection reference.

<a id="canonical-8661f0f7f1001c9830fde13431e629e439a45c31609ca434ec573125b31828c7"></a>

## Next pages — openstack / 0ddf54a3c37e / 4

- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdbfa55f8ea79b4fa5d6bffce20a4ed15c1d49584e037076d10f26ceaf19f17a"></a>

## openstack.not_managed — openstack.not_managed / ddb8fe855cab / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- openstack.not_managed

<a id="canonical-c2ce62b0eddd51734a8a136b3fc0e9c30893422f98f360082e6c113789a1bb23"></a>

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

<a id="canonical-a457a98d297de1698d8bedcd8df222bcbef3a543ef8a8eeed969ac891ea626c6"></a>

## Direct properties — openstack.not_managed / ddb8fe855cab / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743): complete subsection reference.

<a id="canonical-93f8236cbf7a8b6afc6364b6fbc7361950a6ff425ae11a9ea01deb5d128f4741"></a>

## Next pages — openstack.not_managed / ddb8fe855cab / 4

- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-843083c97cb2a31a354ded6b99e9bb901ce1ebb68199f4300a83aec0264a1429"></a>

## openstack.not_managed.node_list — openstack.not_managed.node_list / 694b5e347529 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- openstack.not_managed.node_list

<a id="canonical-1deb9a25c48ee6e0a5780638e38f68f33d7f84425d2642f46e2e8f1477379452"></a>

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

<a id="canonical-b03e71c08f8bf4115ddcbc41ca93ea3c98c34e9058fe8eb1c0fd9a38d4838d65"></a>

## Direct properties — openstack.not_managed.node_list / 694b5e347529 / 3

<a id="canonical-2c3ac7789d846ffbd9105a55e77b6af0b7cf7cb760890667a94d7c6f71129431"></a>

<a id="canonical-5782d94cc17056d7c972518dc1c32e93ed774d8ae0c29167d7a3982558103264"></a>

## hostname property — openstack.not_managed.node_list / 694b5e347529 / 4

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

- [interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112): complete subsection reference.

<a id="canonical-77a071a79513d69f5ae27906bc1066639fb9cfe76b64262f6ee5025f560dc5cb"></a>

<a id="canonical-4744b870d9caee9aa74db0b1053e3f6c62aee8f4b82f4c2bbc206721fcb68ec9"></a>

## public_ip property — openstack.not_managed.node_list / 694b5e347529 / 5

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

<a id="canonical-5586b05cd028040f74f0c2368a24ed721d3a0dfc3f1d00d51b1ad73ad74912fd"></a>

<a id="canonical-6d811e37c3e6b373423d4f2918a8469e024b275286b0afe3d10d56c7cee73ffd"></a>

## type property — openstack.not_managed.node_list / 694b5e347529 / 6

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

<a id="canonical-539404df78ff4a5c714aa758985572bf8f258d16f4bd23735e179fdf87a4ab94"></a>

## Next pages — openstack.not_managed.node_list / 694b5e347529 / 7

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5b2a8cc133ae833b1e25313c9c14cfa7b48e71e397677b48445f330ecb6209a"></a>

## openstack.not_managed.node_list.interface_list — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- openstack.not_managed.node_list.interface_list

<a id="canonical-3e41a860e0983624d61b60e4a64098fb962a9201ef98ec5de73f9760aa915827"></a>

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

<a id="canonical-77b896f278de944035a34aaa282f57916436696b35ca9cb7cacabc7f6bf21383"></a>

## Direct properties — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-f3efee42b72dd4a30c89f7d11ddc9d63f5d6b5e5699cb1bd903dfa4427651936): complete subsection reference.

<a id="canonical-2b42d10c5fba4ce1e8b03ef0399e72a752dd9dc0357d4d723d53d776c6ca62d0"></a>

<a id="canonical-be565e31ac0d4482a7df104d33f8e40e24ccbd7a22e51820f3746d1b5383fb27"></a>

## description_spec property — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0f1861d56bfbdc86de3b56f37a82695696544888ac0e954f383c74eed763cce7): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-8ffd7e6c1a0cfc43727e98c2d6b21ad42b86bd44f85fa88ab4b6c0d536a2024b): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71): complete subsection reference.

<a id="canonical-6ba9ed6b3f4abecccf8aa7a33106f0311a946245561cefd392e05764ff3984ff"></a>

<a id="canonical-06cd5cbaa77d756f3e1ee21d61d61f1bad14a7d525d6e33c107138ebc59e2dca"></a>

## is_management property — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-18d905d502296c3d07d6cd5269262274e6cea495e8c5aa50f3cb2f4c55ffb989"></a>

<a id="canonical-8399a99479ebd468e6895e2871d7fa4b6865fcf0be9adc7b82901b4fbed9337d"></a>

## is_primary property — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-d810f60cc0ce79da512d857c3010b6462bae82532aa3849710899e880829baf2"></a>

<a id="canonical-3d3f2281ac99f3da2bae8e7f9db785362ef934aa6e5587a72fe90a56673dd9be"></a>

## labels property — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 7

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

- [monitor](data-sources--securemesh_site_v2--reference--group-016.md#canonical-386e874527e47c2802096ad770112980b78ad02b225dec7b0cda2a20d3a80dda): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-15446f4bf9a8f0b761a8db27550a0cbfc95b86a2745ec306a37724cd01e97853): complete subsection reference.

<a id="canonical-ef7baa16d5b30149c889e1787577ccb905124a57b75e907e2ab058df1eb5a590"></a>

<a id="canonical-7b702e7f709f23114ba2cc80d00493219fb884c0931a795b212bd7d00c36770c"></a>

## mtu property — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 8

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

<a id="canonical-e5c12d46a319520c40e48b3419072b250482dd5248a57b87e8c1c058e9569e0e"></a>

<a id="canonical-a27423d4a27b4823d44d7f6786a8d30b368a20b381c7c33e2eb6718ae012143d"></a>

## name property — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 9

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

- [network_option](data-sources--securemesh_site_v2--reference--group-016.md#canonical-44ebaca02f99000c50cc000d39ea92ddefe89d49f285f60303bf84c681273a29): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0866cd68f4bd969b0b75029db83b5302a8ee320fd354864309154b13aa773cd2): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0b98ca386f00bc96785d9fd36c3c2b9c428adc2c5b58c2480ed44d513db3cd03): complete subsection reference.

<a id="canonical-3a127d5eb60b9f6948b00ee5750065d9aa779fd688ca076fe8f828fb3cdf8175"></a>

<a id="canonical-528b735712ccf607e0c8805d9a89a7e83a22ca1c0fdf92b8b8b42577c2dba6e3"></a>

## priority property — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 10

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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-25b54c9542051a86a50f7cfb7b0d9568f5bedcd70f75c9ac82a62cc75f9a15e6): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d09f4f34403175abd7a38f76de3bf1d0e6589ea8f0093836b7c999e9cb663a62): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-141aeb5128e7ba8cc807867c87ac30da8e6d5c195093c29871f4aefc307dc9d5): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-5ffceaf6f8a3df2f6a23389ba1cfef198694dfbd6ca38f93d1ac1923e4e4ee8d): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-abffec653af592fd5b551db70284d6adc5260214891b7b006288c421cef714b4): complete subsection reference.

<a id="canonical-b02d4559684c64399aa1cf690390da2a1529f85bf6c57fce7dc0f3630c9d034d"></a>

## Next pages — openstack.not_managed.node_list.interface_list / ed4d1dad27d0 / 11

- [openstack.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-f3efee42b72dd4a30c89f7d11ddc9d63f5d6b5e5699cb1bd903dfa4427651936)
- [openstack.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-015.md#canonical-0f1861d56bfbdc86de3b56f37a82695696544888ac0e954f383c74eed763cce7)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [openstack.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-8ffd7e6c1a0cfc43727e98c2d6b21ad42b86bd44f85fa88ab4b6c0d536a2024b)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-016.md#canonical-386e874527e47c2802096ad770112980b78ad02b225dec7b0cda2a20d3a80dda)
- [openstack.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-15446f4bf9a8f0b761a8db27550a0cbfc95b86a2745ec306a37724cd01e97853)
- [openstack.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-016.md#canonical-44ebaca02f99000c50cc000d39ea92ddefe89d49f285f60303bf84c681273a29)
- [openstack.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0866cd68f4bd969b0b75029db83b5302a8ee320fd354864309154b13aa773cd2)
- [openstack.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0b98ca386f00bc96785d9fd36c3c2b9c428adc2c5b58c2480ed44d513db3cd03)
- [openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-25b54c9542051a86a50f7cfb7b0d9568f5bedcd70f75c9ac82a62cc75f9a15e6)
- [openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d09f4f34403175abd7a38f76de3bf1d0e6589ea8f0093836b7c999e9cb663a62)
- [openstack.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-141aeb5128e7ba8cc807867c87ac30da8e6d5c195093c29871f4aefc307dc9d5)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-5ffceaf6f8a3df2f6a23389ba1cfef198694dfbd6ca38f93d1ac1923e4e4ee8d)
- [openstack.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-abffec653af592fd5b551db70284d6adc5260214891b7b006288c421cef714b4)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f3efee42b72dd4a30c89f7d11ddc9d63f5d6b5e5699cb1bd903dfa4427651936"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eee8c2fcbb7e62bf86056082fc6543bd5525768ccf526023b4f89751ebd2abc"></a>

## openstack.not_managed.node_list.interface_list.bond_interface — openstack.not_managed.node_list.interface_list.bond_interface / e891e9bda3b0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.bond_interface

<a id="canonical-dc1a520d02637eb99cb3f1fd662df6ac3ce8d8dd02fcf52b87377c5788180e36"></a>

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

<a id="canonical-574c4c7ca350c6cae86235d12f3675f82a3a8836a26bbd0e76360396d7cf6d30"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.bond_interface / e891e9bda3b0 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e36acae058b226866621c2e90f9155b0d8c5a8c6ce073d5b54034ac88d0031ac): complete subsection reference.

<a id="canonical-6d4a5e50d1111a5d139ba96f48f900fbc09def3f6e0b8028d80fecd03e6c4f4b"></a>

<a id="canonical-b80c9e79b1698f3b038d8ed9af54d1bc2d97abd61a163ff2a600058071bbc64d"></a>

## devices property — openstack.not_managed.node_list.interface_list.bond_interface / e891e9bda3b0 / 4

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

- [lacp](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2e622e787f480c5a5185a7be81ef32fb999c207d70b4fac1cbda290a52df1d93): complete subsection reference.

<a id="canonical-b8ce9c4f0c86e6f2b59fada73e03e9881702a015ca36da86cfbdae200a484597"></a>

<a id="canonical-d18733bd8c67f267b2ba6f8d01e976d020d373a7720ba04095486298556fb1f2"></a>

## link_polling_interval property — openstack.not_managed.node_list.interface_list.bond_interface / e891e9bda3b0 / 5

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

<a id="canonical-58267a8089bb38b4c80a6fea8ae9a3798df4016f74af4abe3b72f21336057676"></a>

<a id="canonical-3d98c2a117504d39491cead9026b11938a763394df5e402d599a456ab6a4087a"></a>

## link_up_delay property — openstack.not_managed.node_list.interface_list.bond_interface / e891e9bda3b0 / 6

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

<a id="canonical-73941a52ceb2543adfdc9c16eccb7603863e305a9e52a14d2cd6f71777c5bff5"></a>

<a id="canonical-98f927ce424191f1607d7a539f261d7340e21c3f241e845d80e82fa1ccfd02a8"></a>

## name property — openstack.not_managed.node_list.interface_list.bond_interface / e891e9bda3b0 / 7

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

<a id="canonical-c1d65f1127ce949a111798d5977364dac693bed4dfa8e1fc07a1460ca53eed3f"></a>

## Next pages — openstack.not_managed.node_list.interface_list.bond_interface / e891e9bda3b0 / 8

- [openstack.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e36acae058b226866621c2e90f9155b0d8c5a8c6ce073d5b54034ac88d0031ac)
- [openstack.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2e622e787f480c5a5185a7be81ef32fb999c207d70b4fac1cbda290a52df1d93)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e36acae058b226866621c2e90f9155b0d8c5a8c6ce073d5b54034ac88d0031ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-359f07f4a0de45e408ee274e199008e22a3b69ef377e36ac72f2cdebf3dd579d"></a>

## openstack.not_managed.node_list.interface_list.bond_interface.active_backup — openstack.not_managed.node_list.interface_list.bond_interface.active_backup / 5c3016e76858 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-f3efee42b72dd4a30c89f7d11ddc9d63f5d6b5e5699cb1bd903dfa4427651936)
- openstack.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-39f523503ccd4cb48ba20ebd131e827d1af1563f8633396cc2c4d00fed1a2a3f"></a>

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

<a id="canonical-c6274c8debd983065440252824dc2597f111e9ebe228218c5e865e1271277e67"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.bond_interface.active_backup / 5c3016e76858 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17b1694a51cfa0bc2404d6ab4d99f1039ddf2ce9177d04afa483ffdee5af1a2e"></a>

## Next pages — openstack.not_managed.node_list.interface_list.bond_interface.active_backup / 5c3016e76858 / 4

- [openstack.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-f3efee42b72dd4a30c89f7d11ddc9d63f5d6b5e5699cb1bd903dfa4427651936)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2e622e787f480c5a5185a7be81ef32fb999c207d70b4fac1cbda290a52df1d93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31348f45d04cf4fb9d123f908c63532e448c6900ff2b8815a12e82e48a25e30d"></a>

## openstack.not_managed.node_list.interface_list.bond_interface.lacp — openstack.not_managed.node_list.interface_list.bond_interface.lacp / e7c2a5501588 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-f3efee42b72dd4a30c89f7d11ddc9d63f5d6b5e5699cb1bd903dfa4427651936)
- openstack.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-9eeec609e2799f38dd698a4c408665889a95d78742ef5d8be5c22016c9ce834c"></a>

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

<a id="canonical-e54e74dc2eea7a4dea0b10d1134348634a169a7ce48564e2280a704a8392f593"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.bond_interface.lacp / e7c2a5501588 / 3

<a id="canonical-fc838cac7bb4c53f43b3922a8e2ee6aa2ccad8cbaaccfe40ad71ac88a59c4118"></a>

<a id="canonical-45de0fd6f49046ad37f21dc3201f9dd99df385efd50c323231d8abce42f93bfb"></a>

## rate property — openstack.not_managed.node_list.interface_list.bond_interface.lacp / e7c2a5501588 / 4

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

<a id="canonical-43d9c3291f819af7691ebce518478fdb85e786ab02855ee6c5f49a86367369b2"></a>

## Next pages — openstack.not_managed.node_list.interface_list.bond_interface.lacp / e7c2a5501588 / 5

- [openstack.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-015.md#canonical-f3efee42b72dd4a30c89f7d11ddc9d63f5d6b5e5699cb1bd903dfa4427651936)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0f1861d56bfbdc86de3b56f37a82695696544888ac0e954f383c74eed763cce7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d64f68a34ebcf613a8f955c941b6009da954ce0dde4d2f9c39963f2fa3eb4fa4"></a>

## openstack.not_managed.node_list.interface_list.dhcp_client — openstack.not_managed.node_list.interface_list.dhcp_client / bcab9c64180f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-d3a4d0e591d9b034ca648b4acf329935dfd898f3b7866f87f306c6a3c72dbee3"></a>

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

<a id="canonical-1d695a34326df3ab16e71120efe056d7e5e62c74ec20119c3135bb19b0183c45"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_client / bcab9c64180f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c598d1a844102f362f093f954862d6773285b4ed22e6c535ceca1bb7482b1d1d"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_client / bcab9c64180f / 4

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f179fac376c48f06321265b95d691e201ca9efad7d1740828eb386efd28eb879"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server — openstack.not_managed.node_list.interface_list.dhcp_server / 50bbc5c451b6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-94f151e23bb68649e119965e82bb2621f0e00fb19b68be192ac6bcecaacbe862"></a>

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

<a id="canonical-ca9d8a27240f9e8e718fb49272900099ed61f03b483cb43736d3b4a0c675fc14"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server / 50bbc5c451b6 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-015.md#canonical-729c5a51b114935d81c63af6c2fab9070e2f1bff3341c3a85af2a71088a550d9): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-015.md#canonical-d9fe4346c0119b97b2a4d2159d6b4ec9d7df22f9d2477d37d575b2ad3c9c2a78): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb): complete subsection reference.

<a id="canonical-d41263df0a30d532cb995b0ae2da15ee0b3e38c254ad0c2010303d5808c00a46"></a>

<a id="canonical-29b4501ce4b09e4fc106b231bb8a41f3cd46260a5ddb20e02d8fdf23d7136dcc"></a>

## dhcp_option82_tag property — openstack.not_managed.node_list.interface_list.dhcp_server / 50bbc5c451b6 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-a75c9b13e4d34f42fef511ec08847810dacb556f3298414e6d9893384eca698b"></a>

<a id="canonical-fac0a8363638add9f2123a0276ee3e813433f792cc1c42dc955576d563c13930"></a>

## fixed_ip_map property — openstack.not_managed.node_list.interface_list.dhcp_server / 50bbc5c451b6 / 5

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-015.md#canonical-48914842ae990cc4ad956324cf2853a0c9c98137e4c2d0a77f7a4dbf9b6dabdf): complete subsection reference.

<a id="canonical-1f2c1751928b2f58bba1438888128c8f482206f9ef1cc9321703e9f303b2616b"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server / 50bbc5c451b6 / 6

- [openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-015.md#canonical-729c5a51b114935d81c63af6c2fab9070e2f1bff3341c3a85af2a71088a550d9)
- [openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-015.md#canonical-d9fe4346c0119b97b2a4d2159d6b4ec9d7df22f9d2477d37d575b2ad3c9c2a78)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- [openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-015.md#canonical-48914842ae990cc4ad956324cf2853a0c9c98137e4c2d0a77f7a4dbf9b6dabdf)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-729c5a51b114935d81c63af6c2fab9070e2f1bff3341c3a85af2a71088a550d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb8fcdf803376356618635a6ae59fb06d87ed849ee53246740b9dbee7cf931ff"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / c8098872a355 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-6d59e4f9d5631f31829d7eff0abc5097d66e6d29d0b5b31ece3c012984c2e70e"></a>

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

<a id="canonical-b8afd20a693c744783d9f3c76822ab2745350a903da43644da008fe98b385554"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / c8098872a355 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd5841e3e3ff4d8304acffedea380204d3a0fe420b025c1f9dda66f806a8e7d2"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / c8098872a355 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d9fe4346c0119b97b2a4d2159d6b4ec9d7df22f9d2477d37d575b2ad3c9c2a78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa091a0242fc5548558c2abadde74a5389f780981107826618272106337ad5d1"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 2863853eee8b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-d0405dcbb806cdb842b459dff6c67efd2fb9b853d64ff3a928f35e8971489c45"></a>

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

<a id="canonical-40f08e0ac3d54bed872423323aaabe6d8cfdaac28f226087959749d35cade9ec"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 2863853eee8b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-216b0f0b016723801094c90752ada9783374e005da09df16d2efa8fe99235251"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 2863853eee8b / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b54c16d4a215b3a8e5ad51fa10eab2256a3b65b39e7e97f29e309bd4a7788a1"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 2a188bac2b3a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-7d1a0999562d9cfb84e8e36447d16ed00909f25e01ba59cf3a3040dccf5ea438"></a>

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

<a id="canonical-650956c2107bd5e1a56f73e5b3c3ecb27d66348b23c0312eec07b68d2020b12d"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 2a188bac2b3a / 3

<a id="canonical-86dc83f7429f5caacfe2f33719472719135bcf25924f4323a8043ca877853463"></a>

<a id="canonical-810475d940d060c458e4adf496f49b5c3d6d74222a6e6b605d33a82e68298dcf"></a>

## dgw_address property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 2a188bac2b3a / 4

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

<a id="canonical-bdd7c2d81ded219dac131cbeffcd6fb92ee260513b6049d96e2e61a1dd317d59"></a>

<a id="canonical-cdcf842f41c6a07027924a16cb5e3fbad87b7b9e137c8bfd61f1a470f856f79a"></a>

## dns_address property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 2a188bac2b3a / 5

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

- [first_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e21a0d77636390b806e46fc4dd71784bdf2d29c70eb1d0056c583528450ccb6f): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-a564e0a752d1250657e86a9789f0c2a005911b6061b80bc133178de0446a1fe7): complete subsection reference.

<a id="canonical-fe87dab2be865f6aef7d4b2c50d4b400fcc0927e839ff295ef6372bf52b0b006"></a>

<a id="canonical-19a9b98e3da4d4ec680b661ae96fd17159dd044bef6bef0858632339de71f1d9"></a>

## network_prefix property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 2a188bac2b3a / 6

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

<a id="canonical-73846020ea00270c2a661210eefa478c783f354b444b55e14834a903c4686658"></a>

<a id="canonical-7ec18c75951bebb29e0160350d70eac18dcf3cc287abf388ed8f486e72c97ced"></a>

## pool_settings property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 2a188bac2b3a / 7

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

- [pools](data-sources--securemesh_site_v2--reference--group-015.md#canonical-4edd1600a1d1d56bab618e9741f3da9f42d2ec0339a5d01435c7b652a4cb0185): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-015.md#canonical-493ee4dda8b047e33d2e949ef1e4b7eed7c54c7e47184587b7291cd1aa941654): complete subsection reference.

<a id="canonical-94b0893aa41bea00cb06a661c90c96f7774ff85614f39629a257e38a76461e2f"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 2a188bac2b3a / 8

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e21a0d77636390b806e46fc4dd71784bdf2d29c70eb1d0056c583528450ccb6f)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-a564e0a752d1250657e86a9789f0c2a005911b6061b80bc133178de0446a1fe7)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-015.md#canonical-4edd1600a1d1d56bab618e9741f3da9f42d2ec0339a5d01435c7b652a4cb0185)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-015.md#canonical-493ee4dda8b047e33d2e949ef1e4b7eed7c54c7e47184587b7291cd1aa941654)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e21a0d77636390b806e46fc4dd71784bdf2d29c70eb1d0056c583528450ccb6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed198919106173f87ae8cf5d9ca5fce6e43111c958274de4aa57147b757c8a74"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / 457a2a09f540 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-66c323f2d7c8a08d25090315a1fb898d6ae0ecc9f7c03f7e95424c35a5bdb0ae"></a>

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

<a id="canonical-70deee4af61408b65767592b9a2a32afd990b8a300a5032870324e5d2092979f"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / 457a2a09f540 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c8fec301902d7bc3cc4a6bf8fcf3a06e53d1ef4063533647662ea335883cc0b"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / 457a2a09f540 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a564e0a752d1250657e86a9789f0c2a005911b6061b80bc133178de0446a1fe7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19b5cbffba01a9464435f011c836ef2052fd2094ccea4436d020dd01c9e21555"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_ad / cbd3a083e712 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-70cdd902fa8976b0eb735215ee422847f1e3a727e18a11f0109124feaf8f731d"></a>

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

<a id="canonical-e676ef0d89e96983b06504992413e668e886bc219468b418202d771b5e629c67"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_ad / cbd3a083e712 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c3144911623d16fe5d30d82c10cb8b75ead71536ad5057d1a13fd64a721e3ea"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_ad / cbd3a083e712 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-4edd1600a1d1d56bab618e9741f3da9f42d2ec0339a5d01435c7b652a4cb0185"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4a0090834bcda8e63ea20375815d875a7ded5292b0ab91776192dba1db1c223"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32c9da668a4e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-d6ea048eeca6982f31b4c06681814927520463d2a1f6b022c180991587944131"></a>

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

<a id="canonical-a4c17732f5eb7668e2fcdc468e4d83bac7021eb47f4bc47a1aabbee4b0b32402"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32c9da668a4e / 3

<a id="canonical-607c78c582083a9702503c50e89e5c6ee62f12502192bb0aba1364198e54bd67"></a>

<a id="canonical-e582b848342e46f95093ae523cfafd41abf8057dba5e29979b76db5ceec29878"></a>

## end_ip property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32c9da668a4e / 4

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

<a id="canonical-116a00a11c8867c36cc34ab4a018a87fe9cc6cc87ef270b1ae5bd58e5efdc014"></a>

<a id="canonical-4cef5a712eb8f392717e4932cfa4db83edbaeaf86a34268c1177915d76f6cce2"></a>

## exclude property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32c9da668a4e / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-c3cafa8de36e8319045e71123556d26b8d4fcd78a16b500e00551956540b2ff7"></a>

<a id="canonical-1d6bc43ef935766e6207e11cd6891722ef340e130300f0597b87530cd9053e56"></a>

## start_ip property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32c9da668a4e / 6

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

<a id="canonical-297baa9e0bb76409f21545586513d87ddce699c347d54d00eba272704672be5d"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 32c9da668a4e / 7

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-493ee4dda8b047e33d2e949ef1e4b7eed7c54c7e47184587b7291cd1aa941654"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2702bfdbeb6c9bd0ee7c83284d548a0773db369b60606eebc625c108937d1e2"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as / 5ef22682a537 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-16e0dab5ac1d704b2ba7fb039e88fc5bf60e1c4e7378a0b77733857496972dc7"></a>

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

<a id="canonical-36efde3979e7d05fb1bbe0f9449e5124adf83aca1507c9b26abc699bb227bdde"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as / 5ef22682a537 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b24f8f656f271098aa67f2f1e904485fffa95030a0060e58074113669e23464"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as / 5ef22682a537 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-e1123b4b5db32632e65ded0ef67530909b7b7cc252960d71b7bde5aa00acdffb)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-48914842ae990cc4ad956324cf2853a0c9c98137e4c2d0a77f7a4dbf9b6dabdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6512988dfef6057be7cef0161d9ce75173c11cdf2bd2c3ddfa40b1276b7f107"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 248948f659d8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-a2c5a705f8ea338b56e7bee9276baad393881268f93f259535bfda70d5b6ec28"></a>

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

<a id="canonical-e7b0e6ef812134be8d64da4d813ed158cd49f15d001ba089f03432168127ec24"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 248948f659d8 / 3

<a id="canonical-7e7a00e4c0a6095a35ab3916d6fc1f341018fdcbf8eb26cd8425a7ac91130ece"></a>

<a id="canonical-2b0b69e0ec38de7370d1a9be1181f560aa4dcf64add7d0cde681086b7258afd5"></a>

## interface_ip_map property — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 248948f659d8 / 4

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

<a id="canonical-256ee7d1ec754864bf3c62d3da6c7fb5503879ea4fe5da823c6cee8beeabbce0"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 248948f659d8 / 5

- [openstack.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2f1eb700bac383e361ec413cb77070ac6810c0b83da60306b3d429e3e2eb43b6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8ffd7e6c1a0cfc43727e98c2d6b21ad42b86bd44f85fa88ab4b6c0d536a2024b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d13fc1953ce6af0a55d5dd6f691794005070f610c1fd71247eb9267f7bdafeb3"></a>

## openstack.not_managed.node_list.interface_list.ethernet_interface — openstack.not_managed.node_list.interface_list.ethernet_interface / 97bed0a04eae / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-5a687f7c61acf8440d0557451f6f29d43d83d6b300e7dc21eb4b1766b3ee0fdd"></a>

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

<a id="canonical-0f0fa08aa2a378d639a282626c3a62a552b284946c219f5c2e4bbc01341c58ff"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ethernet_interface / 97bed0a04eae / 3

<a id="canonical-6ca2832bce2024f01adc4b03a57045208bdf70407f22a9c960d873e2b3b1a5eb"></a>

<a id="canonical-e52accc9da19a33250cadbecc8070f6d6ccf7620c4b5813474707b514eb3e27a"></a>

## device property — openstack.not_managed.node_list.interface_list.ethernet_interface / 97bed0a04eae / 4

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

<a id="canonical-638cc1d0586cb83f05601e32fa61e9571301d551588eb7d39604c4868b6ca0c0"></a>

<a id="canonical-b0027975d7c341969bc5cbcb12b5c175a504f0b657a294c83e87eba5f3d13bc4"></a>

## mac property — openstack.not_managed.node_list.interface_list.ethernet_interface / 97bed0a04eae / 5

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

<a id="canonical-c72954cf6176a517d5e2c3b2f66ba9e4c6048fa891912ed08c7d2814247d409a"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ethernet_interface / 97bed0a04eae / 6

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5b5d473613f0afff7d52a42c5763368992738ca6cfecb99a9d08e75392535dc"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config — openstack.not_managed.node_list.interface_list.ipv6_auto_config / 56fb642a24f6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-84b250b4e1abeea7fe0ad8291176707922dd61c17ece94bfc179ff29c31eecb4"></a>

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

<a id="canonical-0cfb289d7c4c37426992f9d15fb0bb0581b37094060a05df963868bcdeaabcd9"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config / 56fb642a24f6 / 3

- [host](data-sources--securemesh_site_v2--reference--group-015.md#canonical-bc310f34b5315963f476d8a8370ace1841ee412209b597a69cc2310ea22cf6d9): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9): complete subsection reference.

<a id="canonical-528889126148c721674a55d809a438903ba1af5731472216e83583dd65b664e6"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config / 56fb642a24f6 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-015.md#canonical-bc310f34b5315963f476d8a8370ace1841ee412209b597a69cc2310ea22cf6d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bc310f34b5315963f476d8a8370ace1841ee412209b597a69cc2310ea22cf6d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee0fe078126717b12157f82a703bbae34db9748cbf414c3ccc0ac8d88c3badda"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.host — openstack.not_managed.node_list.interface_list.ipv6_auto_config.host / 64f37b362ee2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-20791506107cc6d56e8e2597ee244991d10aa3b291b3e3e85924e09209a5cd18"></a>

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

<a id="canonical-3742ae796c55244ebc5f1f87690500e05e0b85fb634265b2b825cbcebea666bc"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.host / 64f37b362ee2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-81449114d4da692491753b6e9289dc965a85bd38bff635d8301ddd2a58ad8b54"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.host / 64f37b362ee2 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7207e9537a26d034160d1529f97c31809f16f3ea6d74e9a4db7cdefeae6a4d1e"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / 4267d5944c67 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-4c00e0d1ee52ea53d8a7567cc414aec91a08e00ee447bea9923ac62b83cd32e6"></a>

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

<a id="canonical-8636dff31144b41d49dc13c08c9b9c913b45afa2cf8a03a1a4cc4a950c3c7459"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / 4267d5944c67 / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19): complete subsection reference.

<a id="canonical-439a418d0bf93cad0b3d8c091c0c9dc335104e0fa69d752ddd31c17c0e99546b"></a>

<a id="canonical-131959933c77b6039189be8214cf9b1e00ac3173c1c42a02240e75c078b94388"></a>

## network_prefix property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / 4267d5944c67 / 4

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

- [stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58): complete subsection reference.

<a id="canonical-476190ce71900d3de3ae63cd1ad685e9a00b193cc4c0c436e98f9b31baad5a63"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / 4267d5944c67 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e91d2b6cd7125ac6eb648108883c3833e85e5fa2e30d1cbd68f440f3f5dfb652"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / e459a05d8d26 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-b6682b645925d16181aae57aba7e50aab7a50b364acb085328f1be2902a98e41"></a>

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

<a id="canonical-52daa78717b2a61bbfd022b57ee297307469c94a96a6e50a37854fdbff360107"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / e459a05d8d26 / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-35892c90633c0ee278d525ccfda33b9e200c3d9fb182f5c9c09dd644f62090b6): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-08e78d7998be05fe3f3aa8375fd78973d5bea62c77bd44ba9bf9c9a9e9723e20): complete subsection reference.

<a id="canonical-9ad3b2a6c6c58694c1d680415ca03e939f9ca3feb28f2e494df29108b8090b46"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / e459a05d8d26 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-35892c90633c0ee278d525ccfda33b9e200c3d9fb182f5c9c09dd644f62090b6)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-08e78d7998be05fe3f3aa8375fd78973d5bea62c77bd44ba9bf9c9a9e9723e20)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-35892c90633c0ee278d525ccfda33b9e200c3d9fb182f5c9c09dd644f62090b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b56806644fdceb5eea75c037189beb391974787effdca460e447ae469932ed8"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / b60ed5a2ad2a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-42bdd51a2c5cdd51f6e999bfbed4fca9444dcc264d64fa2d87c0730f549ebb96"></a>

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

<a id="canonical-932031cf37593472b8f8ae2b17b321130e7b44acde83fc7ea620568882016697"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / b60ed5a2ad2a / 3

<a id="canonical-b8bfb29ceb6d5fc423fef8fa70025ee1e64c405439cc08a3f94ed50cbc8c9a68"></a>

<a id="canonical-877739261f3f748fd46ed013c23f7d2065aaaf699e476ea0f0b4f4d52d6c306f"></a>

## dns_list property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / b60ed5a2ad2a / 4

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

<a id="canonical-f3ec8587b85e85f5c7282002c5d644b25108f1e684bbfd3f7782228716ac884e"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / b60ed5a2ad2a / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-08e78d7998be05fe3f3aa8375fd78973d5bea62c77bd44ba9bf9c9a9e9723e20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0e389f4f1f51b03ddee8234e53d3e7011f2d6d8c8e68a3ae0bbd4335824b384"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / e47fd586f6b0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-ae43729e3c5b9ac7d9ede43657c1d4eb3423d3b01f67d61e3eb0d1d24282dbbd"></a>

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

<a id="canonical-76198a455461ba252e89a83a43cb566a09cf8f1938a05374908e0f6919f693b0"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / e47fd586f6b0 / 3

<a id="canonical-e08ecfdf33bf61e1349ec347672ae46efe69ac017eb066a5ff5ec2d229207cd1"></a>

<a id="canonical-004364d040067b75ccca51f56baf27ddee7067bc95b17dc0b793a7c06ce3f024"></a>

## configured_address property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / e47fd586f6b0 / 4

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

- [first_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-5f07f127a48f23d2aef513cd9c3fbe318586ae56b316ec8533acc625952c120f): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-27d0b69b7808d4cad0693e88c6659420a559e40b7a40a8bf6b9e9d446ce6bf3f): complete subsection reference.

<a id="canonical-c0a5d2a3246167f3c031f41d082125034787913836b6de9a0d141372cf46d909"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / e47fd586f6b0 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-5f07f127a48f23d2aef513cd9c3fbe318586ae56b316ec8533acc625952c120f)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-015.md#canonical-27d0b69b7808d4cad0693e88c6659420a559e40b7a40a8bf6b9e9d446ce6bf3f)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5f07f127a48f23d2aef513cd9c3fbe318586ae56b316ec8533acc625952c120f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92b49da373f068d500bf7c0fcb4f0c041fd96a7498475242a48dec8fa57b2bc0"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f1d1f6dc3721 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-08e78d7998be05fe3f3aa8375fd78973d5bea62c77bd44ba9bf9c9a9e9723e20)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-5aad9879bef9f152f0c428e198ae3d46c298116b952e3bda6b02cb128b6eea0a"></a>

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

<a id="canonical-fc967753252ab8500825ebfef84bfbdbe2b0155f727ee8af8520328bca2c9b33"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f1d1f6dc3721 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f39a3fa5221bd2d0e7e83da201623302c5386a287d39f021eb4542a21f569a27"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f1d1f6dc3721 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-08e78d7998be05fe3f3aa8375fd78973d5bea62c77bd44ba9bf9c9a9e9723e20)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-27d0b69b7808d4cad0693e88c6659420a559e40b7a40a8bf6b9e9d446ce6bf3f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-009a8294996aa90d34bdd297fed7f84354c1c5f82c379a81bba738f066f512ab"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 6de9c7a733fa / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ec1f07baedda969a692b6d0c0f81f938cd94cd46e30f9c302280403e8856fb19)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-08e78d7998be05fe3f3aa8375fd78973d5bea62c77bd44ba9bf9c9a9e9723e20)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-6cec6e79bb07dc42f67a3263600e493523fcb92ec927eaf1213ae40a19c0ced6"></a>

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

<a id="canonical-5dff21267741fc0356b89e4fae68699121ad7a98149fa79b4e7567c06fdc55e7"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 6de9c7a733fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd134449b8da64f8c5fa779847e48c3b4c6451b3b40ff6646cef218e7ca0f3b5"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 6de9c7a733fa / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-015.md#canonical-08e78d7998be05fe3f3aa8375fd78973d5bea62c77bd44ba9bf9c9a9e9723e20)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ac62adb579a2d568133be342ddd8eab3695c773ccbec3bd60de489750a762c0"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f722a405953d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-b50a2829fb121f94a55d7b464bc9de0e1383ab2f9d5982a4e7d14b194f4eb661"></a>

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

<a id="canonical-a3bb8c7d0bf5fb4bc33b25d4f121567fae5e1041075de9cc3740af0f8dfbf148"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f722a405953d / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-015.md#canonical-24b1c3cb907b3603455e49f7507b53794d0478dc7e903f9b2993929073ac3239): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-015.md#canonical-17a4a268dbd84c56698836f530213974d0bfb8b5cb6a0f370bd91f96af02cf3b): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-512480b924ceef17c9418144af4d27c1a0f9849ad7208f092d7ed3f3c652912b): complete subsection reference.

<a id="canonical-b58b44251179fbd1bd50a5b7240a65b6e3ca5c8a79ebdc1f0fdb1c2c9ee0194b"></a>

<a id="canonical-79b0c22446b3a8fe138d5de756439d07caca793f05fb97905c5da65d6509479f"></a>

## fixed_ip_map property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f722a405953d / 4

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-015.md#canonical-477bbdd164022da21518ba2f26673562956e931de8654f9f48fc166fbc0b0677): complete subsection reference.

<a id="canonical-a8537359feec09c7335cdbbca4eea4d4b4eb8bcb3b478b0e1f184d8bf4f7d6bb"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / f722a405953d / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-015.md#canonical-24b1c3cb907b3603455e49f7507b53794d0478dc7e903f9b2993929073ac3239)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-015.md#canonical-17a4a268dbd84c56698836f530213974d0bfb8b5cb6a0f370bd91f96af02cf3b)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-512480b924ceef17c9418144af4d27c1a0f9849ad7208f092d7ed3f3c652912b)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-015.md#canonical-477bbdd164022da21518ba2f26673562956e931de8654f9f48fc166fbc0b0677)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-24b1c3cb907b3603455e49f7507b53794d0478dc7e903f9b2993929073ac3239"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-941646a833f72b27b92062b74498868d21c2ff0ab8339f39320431526dc00534"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 4e617ac0e60e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-6e238e61b673b90aef2c29da12d9f95010fd49bdeed063d81066f4495a182906"></a>

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

<a id="canonical-66e10119c337f8dc83f3d4dd783ae017d77f4c4d21130b7002238c3952aa7223"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 4e617ac0e60e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b26754f4b8310cff23fe205476ec97b88af3cfe26940499e0dc74e71e0fc0661"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 4e617ac0e60e / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-17a4a268dbd84c56698836f530213974d0bfb8b5cb6a0f370bd91f96af02cf3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff8ad800e03427cb9912b0a894d0ed5a5fd0149fb9ce5fe6a7bc5e80199cdd25"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5745371f8c17 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-b4f2adad04ca4cf65146987cc16699abf537108f5f4f0075c523adb2c5aae959"></a>

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

<a id="canonical-68a84337db4504531cda8509fcbb714287e013419a8453a71f6cc5e7dd5a1734"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5745371f8c17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42e4c671e195c28e3e8772aae40e0212deff632af534538d23a8b15d76bb9c28"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5745371f8c17 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-512480b924ceef17c9418144af4d27c1a0f9849ad7208f092d7ed3f3c652912b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45e3ad4f291d3bd91fbfa7949255829f2f72c1ca5b73806be9b87e76b345c88c"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / fe7b037db958 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-d27661f1a88b1e3af9df7186de0fda4df6864ce7973fd13a60278a1afbc00389"></a>

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

<a id="canonical-db47afabc9a4225ade1da3b6c169848deac1bf0a4fc7a9f25d7d4878caefe4ab"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / fe7b037db958 / 3

<a id="canonical-040579e1023505bd2d3549c6b638c63205f8d4ff83151df704738a9159f026a5"></a>

<a id="canonical-728d88208fe411d4a7ec10a28745091b2dbc588e39eb45260f8907a09efc012d"></a>

## network_prefix property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / fe7b037db958 / 4

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

<a id="canonical-7a36c227955b59bfdc1618a041e15fa7a5cf3b49e2efaea26764218f02c3fa2c"></a>

<a id="canonical-543a02f6e9272cdbec9a0f2c0c16b070a6f57ff4a6d29e57df23b652b1bdf04a"></a>

## pool_settings property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / fe7b037db958 / 5

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

- [pools](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1713b07191864b905623759b8a9744b23f80adc73fa5fdfc66045eb019d7aa67): complete subsection reference.

<a id="canonical-749e23c57ef9b113fd7d97fd8462489b53b70a80efaa14387237d01d70f71811"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / fe7b037db958 / 6

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1713b07191864b905623759b8a9744b23f80adc73fa5fdfc66045eb019d7aa67)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1713b07191864b905623759b8a9744b23f80adc73fa5fdfc66045eb019d7aa67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fae22ed48eab56605819e7a0919d2ef04624b46dc7fe57f99cfab33bcea7a31e"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ac030bdcb43f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7815396d285d15a14da901a8d68bbf960451e303fd25252fb51efb800a086b71)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-015.md#canonical-7207c3a8895948a6ee891ea9a628f1b3b055ae7b00dad30cba7a494bf1dd33d9)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-512480b924ceef17c9418144af4d27c1a0f9849ad7208f092d7ed3f3c652912b)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-c2efa54fe480323d59d80c9006430056a84bc4fe41ead91792fff1f2b6c1e2c3"></a>

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

<a id="canonical-f03876d4e26a3f88a71c1b8bdb7bc3db2b40629b8b1c1736e90d6747d112c93a"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ac030bdcb43f / 3

<a id="canonical-6d3cf10a82bb70d27768e0f6d7f0d5833d6d2cac61f4c7906329435c292af1cf"></a>

<a id="canonical-1a64b65a12c7c41b2719c7ab2599952e9c808c20bf394dfce7a87c74c20655f6"></a>

## end_ip property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ac030bdcb43f / 4

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

<a id="canonical-424ba933690961ad20978bacdc08c71aeaa0b55a86b1a64db31ff09236d06f33"></a>

<a id="canonical-dccfa2ce87cd6b28b83602549a9998f1c0e265fba6723307970e3f2d400c078f"></a>

## start_ip property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ac030bdcb43f / 5

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

<a id="canonical-8f3c09fe6787271f984d11561bed220e87f0607086727147e9a34fd83bb87f7a"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ac030bdcb43f / 6

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-015.md#canonical-512480b924ceef17c9418144af4d27c1a0f9849ad7208f092d7ed3f3c652912b)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-477bbdd164022da21518ba2f26673562956e931de8654f9f48fc166fbc0b0677"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
