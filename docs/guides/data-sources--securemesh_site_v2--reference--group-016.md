---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3d2e0eaea1fce645ac4a3e856770dd912c085cb2703e66c145c896f5bfac7d24"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 9bdbd5732eb3 / 2

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
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-862c2fc729684068dd0fd04592919754b22dbdbbf4567720fc9991b967d4c178"></a>

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

<a id="canonical-d3f822883155bfe1b0f256eca9f45cadf2018c48b6daef8d6678f19e879b6734"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 9bdbd5732eb3 / 3

<a id="canonical-038136f3bfcf455a71bf90a6770304a38b77a847782b857c3cfc73b4e7baaf51"></a>

<a id="canonical-9f678c63aaf1882140c0327cd73b1e8540992ccfb6d64b694f11a87518db3035"></a>

## interface_ip_map property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 9bdbd5732eb3 / 4

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

<a id="canonical-4a6dfcf8220f4c05dbf99febc50a726369a3f48a77fb573a6d8c923827c0162e"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 9bdbd5732eb3 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-015.md#canonical-c743721e5d6f938f0ee2dce51eee4fc69d8ae819ab6b9f4063f0782e160ddd58)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-386e874527e47c2802096ad770112980b78ad02b225dec7b0cda2a20d3a80dda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d5ed25c79a34eb4da3be90f7e795a0a5d98be735e9d429fbc7c9e6a282df9fd"></a>

## openstack.not_managed.node_list.interface_list.monitor — openstack.not_managed.node_list.interface_list.monitor / a9c02947c5b9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.monitor

<a id="canonical-44758c9f4070a79617aeac2d32d49ef908842049c752703430b31b7dad7581db"></a>

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

<a id="canonical-b845915558c79bc954884a36bd8ef59ddac26f4c1eebd810b97166d3a1312ec7"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.monitor / a9c02947c5b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a6f1a0ae91fc797e606271229fbd6f0a4c1cb12595b4859f82b0073569c4a21"></a>

## Next pages — openstack.not_managed.node_list.interface_list.monitor / a9c02947c5b9 / 4

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-15446f4bf9a8f0b761a8db27550a0cbfc95b86a2745ec306a37724cd01e97853"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f15171cce00343d52598afe9208e1d196cf64d2611aa8cadbc43c93dcef4e9b2"></a>

## openstack.not_managed.node_list.interface_list.monitor_disabled — openstack.not_managed.node_list.interface_list.monitor_disabled / f8f09b948c11 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-5d41dd66afaff78081900a182f8b2c8913baa751d01a9b55f1335afd6d922de7"></a>

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

<a id="canonical-10741f2b042a4196898fe9c2d18c132fba2d3cd459d602109b2deb7bc7941f65"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.monitor_disabled / f8f09b948c11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6520e9d72eb259c9cb442d83b71d718e9aab2f20ba9d60f0d5e4e493aeb88a77"></a>

## Next pages — openstack.not_managed.node_list.interface_list.monitor_disabled / f8f09b948c11 / 4

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-44ebaca02f99000c50cc000d39ea92ddefe89d49f285f60303bf84c681273a29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7468f3e2fd1c63d42f591e670f1fcad0b32857d6cf3dca5db5b2d9c801fca8d"></a>

## openstack.not_managed.node_list.interface_list.network_option — openstack.not_managed.node_list.interface_list.network_option / f29386cd207b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.network_option

<a id="canonical-343829727bfac7a2ff74f689f5eeaaf853de64054da0aaa1141ed52a98850f4f"></a>

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

<a id="canonical-6db4eab72e52df962b4feabd0a29042a938a597de88dc089e5e97f733427f02a"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.network_option / f29386cd207b / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-016.md#canonical-73deb11246c80ae397610b7844b72ded3f55e46e26d3f397a195fb18799125ee): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1f0850f67d85968d64a8bf387a36cedbf44a47e1abd80dbc49b70d72ca579b17): complete subsection reference.

<a id="canonical-f1b8c85fe2d1cfe7955e9dcbe9f73eb0ba90330990d2efacfab7f7718290b322"></a>

## Next pages — openstack.not_managed.node_list.interface_list.network_option / f29386cd207b / 4

- [openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-016.md#canonical-73deb11246c80ae397610b7844b72ded3f55e46e26d3f397a195fb18799125ee)
- [openstack.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1f0850f67d85968d64a8bf387a36cedbf44a47e1abd80dbc49b70d72ca579b17)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-73deb11246c80ae397610b7844b72ded3f55e46e26d3f397a195fb18799125ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fa186c1a78271d6b24148fe4bd24376ac916b56efd8ce08f640127a48259263"></a>

## openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network — openstack.not_managed.node_list.interface_list.network_option.site_local_inside_ / c114897a3be3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-016.md#canonical-44ebaca02f99000c50cc000d39ea92ddefe89d49f285f60303bf84c681273a29)
- openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-3ff7ccd93b51bb5c50882a783d7774e29e75162a070262450502d556bb292e43"></a>

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

<a id="canonical-f4b3b2b1b943023a50f8e8c612be140bc600fb52849a717454b196568425c38c"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.network_option.site_local_inside_ / c114897a3be3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-012e06cb0f2776d0bd109bc225f2d8653601f6a8612d2a9d0805886fc1a2bf06"></a>

## Next pages — openstack.not_managed.node_list.interface_list.network_option.site_local_inside_ / c114897a3be3 / 4

- [openstack.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-016.md#canonical-44ebaca02f99000c50cc000d39ea92ddefe89d49f285f60303bf84c681273a29)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1f0850f67d85968d64a8bf387a36cedbf44a47e1abd80dbc49b70d72ca579b17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef6f511edc7e5285ce1f681368b6db3e474a14c402567f880564ceef68dc427c"></a>

## openstack.not_managed.node_list.interface_list.network_option.site_local_network — openstack.not_managed.node_list.interface_list.network_option.site_local_network / 0930b3116a8a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-016.md#canonical-44ebaca02f99000c50cc000d39ea92ddefe89d49f285f60303bf84c681273a29)
- openstack.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-c35392add5df25ecd06b5d1c094281ae3ff0af3e22963e2b7402359990c0de82"></a>

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

<a id="canonical-ea54b1b1eb453e1c6ad4e229dc775f65f25a0a451b57d20f1d3f4352e486d546"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.network_option.site_local_network / 0930b3116a8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e10ce0bf00d67f7e1454d6f4aff192dc7c24a84036d6a0ea18ce1ac122bb1351"></a>

## Next pages — openstack.not_managed.node_list.interface_list.network_option.site_local_network / 0930b3116a8a / 4

- [openstack.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-016.md#canonical-44ebaca02f99000c50cc000d39ea92ddefe89d49f285f60303bf84c681273a29)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0866cd68f4bd969b0b75029db83b5302a8ee320fd354864309154b13aa773cd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9a14dd89eae14e6d2fb93606bec839b55afb89d66b9a620a802c8298819d63a"></a>

## openstack.not_managed.node_list.interface_list.no_ipv4_address — openstack.not_managed.node_list.interface_list.no_ipv4_address / 6b31fa29d5f0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-483c72faf7b092cd73fd67e219ff3fc9d6c13cb426a5edd3458658c0fbe63c7d"></a>

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

<a id="canonical-53102104056de02fccbfa966022d0225bc7c109ed6a9fe9f4978644b5a0d4c2f"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.no_ipv4_address / 6b31fa29d5f0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9afecd1038fa17aebd6dfe2ee206e8b218d69ea458e8f7d913c1dd844887d0a0"></a>

## Next pages — openstack.not_managed.node_list.interface_list.no_ipv4_address / 6b31fa29d5f0 / 4

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0b98ca386f00bc96785d9fd36c3c2b9c428adc2c5b58c2480ed44d513db3cd03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97d2f71eb31d89c186f54f345cbf4e1484b531f787bc79c3fbde385edf3d39bc"></a>

## openstack.not_managed.node_list.interface_list.no_ipv6_address — openstack.not_managed.node_list.interface_list.no_ipv6_address / 692e2c73d42d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-34da3013f4780dc26c8f003ddac759c103f5f8307df3f5e2486ab1d9c3a952d7"></a>

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

<a id="canonical-06b33439d8325eab4baa19a57df0c1a298cccf1d4838d56da9ef088773d07a23"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.no_ipv6_address / 692e2c73d42d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ecb573b1d0f304db8bd269f0926341375c16592278cfb60d7ad97f3aa7f9b61"></a>

## Next pages — openstack.not_managed.node_list.interface_list.no_ipv6_address / 692e2c73d42d / 4

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-25b54c9542051a86a50f7cfb7b0d9568f5bedcd70f75c9ac82a62cc75f9a15e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-907745e9383223b0f881194b80909ec66a5351dcc9b87aefd1c338ffbccd1e53"></a>

## openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / dc6f51e2c131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-245b6df7033e14f5f55416a1f761a4ba8bd8c48745d4b39ad0bb444154abe3e4"></a>

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

<a id="canonical-441e1bf86f31a092344a2f15ce6e50eb04499e9bb0adc4afca5d6ef4abab3271"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / dc6f51e2c131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-67b45538c6bb567d8f1c78b033941a12f6a06d1480c8487a11d7e044ed65ad30"></a>

## Next pages — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / dc6f51e2c131 / 4

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d09f4f34403175abd7a38f76de3bf1d0e6589ea8f0093836b7c999e9cb663a62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bdd8a48319fdeac4169562772a4cb7a98f02f337714b1300b5c7957ef9c95bc"></a>

## openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 1d0927e8883c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-2d636c407b4ee8c38b063399f419d7abbe6efc9e6b75be45c15c5b8570d9986d"></a>

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

<a id="canonical-dd540700fa5b6543df46cccaf7bb192f1124d1a0a4b0ee4be164e264f5217469"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 1d0927e8883c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa5cf3e8d6d7dc26ef8faa853a3c0894ed7cf20d48acfe1f953061cbf4681bf1"></a>

## Next pages — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 1d0927e8883c / 4

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-141aeb5128e7ba8cc807867c87ac30da8e6d5c195093c29871f4aefc307dc9d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9155e7a2819fea754581ee397f2c03e1a9a3511193f1e54241a0017f2fa4623"></a>

## openstack.not_managed.node_list.interface_list.static_ip — openstack.not_managed.node_list.interface_list.static_ip / 761d7ce8a3de / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.static_ip

<a id="canonical-9260265984b3923f8a34fff6bbd23ccc3ce86c65cdb19e95d66e374e78f309ff"></a>

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

<a id="canonical-ac63b5ce3a042a24ef596de57101fefc31e79fa72630285fc5b94fe75c1205d2"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ip / 761d7ce8a3de / 3

<a id="canonical-3f6a88d1ad85e0d35d32a8c20ed39c9efd75e141f706414f233bbd1e625721e7"></a>

<a id="canonical-683765ac9557c0178cb90556c28edd3359a685906e09548a3f58d237a07f3356"></a>

## default_gw property — openstack.not_managed.node_list.interface_list.static_ip / 761d7ce8a3de / 4

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

<a id="canonical-97773ff7fc2cf4ea445fa53f3e6cd8c4de5a00cc1b1700d26765579498418105"></a>

<a id="canonical-bb27d1425f93cb308ce8461d76ea052fee1c94d54d7c0bfc80dd56b1f5e55dfd"></a>

## dns_server property — openstack.not_managed.node_list.interface_list.static_ip / 761d7ce8a3de / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-9a45141e79ca4b27415574447bc7700db217f98be860c2100156f0e30caf2f3c"></a>

<a id="canonical-01d74f8bc9874d7353763a799ac8547f6093633795476a616287e16988cf276b"></a>

## ip_address property — openstack.not_managed.node_list.interface_list.static_ip / 761d7ce8a3de / 6

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

<a id="canonical-66482a497ed2325e7567469bb3042557b66793109b28c968f1b2d6ae9130f359"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ip / 761d7ce8a3de / 7

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5ffceaf6f8a3df2f6a23389ba1cfef198694dfbd6ca38f93d1ac1923e4e4ee8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ec843d4134177460694e363a5ed2311bb711b07e3a4c85e3bcba7a62b8ccf9f"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address — openstack.not_managed.node_list.interface_list.static_ipv6_address / 2ec84ae370ac / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-4cd3d5fea63a5b1fa08c6e6a4e24b0fd8b1b53cb805a85fe917c725db37918f0"></a>

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

<a id="canonical-d7a58562687e769f4f192d3d17ee4c1900be1be5a7b2437094ace6256bd17700"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ipv6_address / 2ec84ae370ac / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c1c361abd2f3437c1bb5a7a66b8e06d279c135af15a45d1ff39722b3d99f4f07): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d6e18ba50bfec3fc6f68f7dd69a5398c411fd5fad218557b1a35788d29177d74): complete subsection reference.

<a id="canonical-97865cb3b24138b939fe26ead18542ba92b90d0efaf22620893c3515ae93ebea"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ipv6_address / 2ec84ae370ac / 4

- [openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c1c361abd2f3437c1bb5a7a66b8e06d279c135af15a45d1ff39722b3d99f4f07)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d6e18ba50bfec3fc6f68f7dd69a5398c411fd5fad218557b1a35788d29177d74)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c1c361abd2f3437c1bb5a7a66b8e06d279c135af15a45d1ff39722b3d99f4f07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-799e193422e2a87bd6a3bb7369631ada7c507f8f4cb856cb23e29eb41f2e3b13"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / a5dc9d44e949 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-5ffceaf6f8a3df2f6a23389ba1cfef198694dfbd6ca38f93d1ac1923e4e4ee8d)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2bb460bcab8ec752892f14f45c0a63100e8af7c99d73aa2e8934507c380eeb5c"></a>

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

<a id="canonical-70bc0b4e9bf9b34904996e315d7dae923ec327d2f67830838c98003581954149"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / a5dc9d44e949 / 3

<a id="canonical-b6610b377dc943e178b59dcad65bd5ab13e85b2955d114cca4ae6c1ecd88e34b"></a>

<a id="canonical-ba29338341eb298b0d39aa1c662d22dfe427bb9903df65f1db3a57ab925459d9"></a>

## interface_ip_map property — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / a5dc9d44e949 / 4

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

<a id="canonical-07e912cb28fa3414166a22c2335df89ac983e10365a7e27443fa8ecf40d82209"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / a5dc9d44e949 / 5

- [openstack.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-5ffceaf6f8a3df2f6a23389ba1cfef198694dfbd6ca38f93d1ac1923e4e4ee8d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d6e18ba50bfec3fc6f68f7dd69a5398c411fd5fad218557b1a35788d29177d74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a4eab3480b48fe77bcad37d10bf88f2b867ec63be7983b9a824cc24483fccf8"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 8a505d379f38 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-5ffceaf6f8a3df2f6a23389ba1cfef198694dfbd6ca38f93d1ac1923e4e4ee8d)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-d785e8afc4ea58981d0a20fb7a395e55b92190df3cd1a0bcf33a5211e1fa5c05"></a>

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

<a id="canonical-972f54bb83d8469c40c2a2c8c55ece1ac3e006f947365e52dec454453f6945f8"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 8a505d379f38 / 3

<a id="canonical-3c6399a93344673d0d3dca84b57bdde63739597e889a6ff91bc60c23b67d8d4b"></a>

<a id="canonical-28ec52654604d4af5be5ef3e12a759383516bf09196f85d8a3aac2ddbd4f82d1"></a>

## default_gw property — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 8a505d379f38 / 4

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

<a id="canonical-9076ee808fd633603ebb3722a4bb3fd22be9a8db8858a7ee235a182f606ac590"></a>

<a id="canonical-95cef7f2c8bb823cd6339f017bc68c50eaed629772c6a5a2185bf49f1151fede"></a>

## dns_server property — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 8a505d379f38 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-68cf063a90bfd538bdbc614000dfa9c8ba0b3fa4f159ccbf4dd432f825528b5b"></a>

<a id="canonical-00629936d35a5946570085908fc1191d55e00c4fa294dcdbf79fd7c5757f077c"></a>

## ip_address property — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 8a505d379f38 / 6

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

<a id="canonical-fa8777dc87309811bf8fa998668f9c66e1cd70fb1dbb80796be599163c846889"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 8a505d379f38 / 7

- [openstack.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-016.md#canonical-5ffceaf6f8a3df2f6a23389ba1cfef198694dfbd6ca38f93d1ac1923e4e4ee8d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-abffec653af592fd5b551db70284d6adc5260214891b7b006288c421cef714b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b29a0c5425cce2dd1c22201e9cbd9d916e2324102e78f71929b22ec61574b69c"></a>

## openstack.not_managed.node_list.interface_list.vlan_interface — openstack.not_managed.node_list.interface_list.vlan_interface / 2b4ad0d9f1c1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-ad8266eb8554942be0d758b2f7fc2e0cd13f5fd5eed7cc00d746a5c75d939063)
- [openstack.not_managed](data-sources--securemesh_site_v2--reference--group-015.md#canonical-19189d35cb70ae238492c722efa1be92c6dd906f53fc1848ac6fa15fef2a640f)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-1ef5f58736c43ee8225459e8ba831cf175ebf610f3ab92ca5c8a993570b31743)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- openstack.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-cb5b449c2af24d16da8b6f1c915863b5b55b341a2d6db80ca61076dddf871a32"></a>

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

<a id="canonical-41ebea56a8f5b597bc1a21c675ec37cb7db523cd3ff360d8dc30efd2a9d5972e"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.vlan_interface / 2b4ad0d9f1c1 / 3

<a id="canonical-d16774d0401b19c830617e2f9a194b6eaca51e04b5f6be6dd4b5a5dc941fb999"></a>

<a id="canonical-8747059260eccf88ec40efac6c93b387be39200f027f486522d278c2c8f3aa82"></a>

## device property — openstack.not_managed.node_list.interface_list.vlan_interface / 2b4ad0d9f1c1 / 4

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

<a id="canonical-89b46c2cfb0cfc459bb0c657ebf74f79872faef5e0d5d7d34cfb1c993e0a2085"></a>

<a id="canonical-2ca534a9f1969c26c1d2d5886e79053100bca84750c895afc6e350adab87eb92"></a>

## vlan_id property — openstack.not_managed.node_list.interface_list.vlan_interface / 2b4ad0d9f1c1 / 5

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

<a id="canonical-d1bb8030f1426ae231b7c570bb756f8671f2b2f279f9ea92c2da0f8b7464c00a"></a>

## Next pages — openstack.not_managed.node_list.interface_list.vlan_interface / 2b4ad0d9f1c1 / 6

- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-015.md#canonical-3b5c5e34fedd02a6af3f2acd553715b796336f2af3cbac4d2f38e6fe38278112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6342aabbe8b26eaf9a04b078913976f00bbfd9decef15f7ebaaff9c1bf41ec6"></a>

## performance_enhancement_mode — performance_enhancement_mode / e82c4babaedf / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- performance_enhancement_mode

<a id="canonical-5b1e1711ad7ee84040830235332631cd248ef99d6507cf99cb31b14497a2f4fc"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

<a id="canonical-dd6661d35154839dfc5fd4e819d7538fd2d4fd14f90ff3a12606b93765d89c2b"></a>

## Direct properties — performance_enhancement_mode / e82c4babaedf / 3

- [perf_mode_l3_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-cd9b763fc7ccd3c86c098a47ea4e3c2ef3ab42aff81bbd9be71c55714b1d81b9): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2952e4660e133e0b14e3bb53d417411fbdcb30b46d45b67a7c0a50574d849346): complete subsection reference.

<a id="canonical-acf457281bdc30d00a2f771f0ef27babd06836e407087b8800cbbd1c6677a74b"></a>

## Next pages — performance_enhancement_mode / e82c4babaedf / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-cd9b763fc7ccd3c86c098a47ea4e3c2ef3ab42aff81bbd9be71c55714b1d81b9)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2952e4660e133e0b14e3bb53d417411fbdcb30b46d45b67a7c0a50574d849346)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-cd9b763fc7ccd3c86c098a47ea4e3c2ef3ab42aff81bbd9be71c55714b1d81b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00046aab38f45e18796dee4dc492e4f21540d8c404eb080437d7c0ac6d36dc8e"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — performance_enhancement_mode.perf_mode_l3_enhanced / 0a29478a3cab / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-39a124c72d8a473a3480802c0ec5eb04d12b8ee87295f783b91fc660ccbb36d2"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

<a id="canonical-408eb26a851ac71cad66dfd8db97bcd8c843b2bf8082cc33e13be6acb14d0822"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced / 0a29478a3cab / 3

- [jumbo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-9be598bdbed703ea40a4687371546f2016c74d6ea3c454dfe6e6f175324ccbf8): complete subsection reference.

- [no_jumbo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-ac14c959af6a6af27017bb0919c10681e3e382409c59a6be3644e08bee187d5f): complete subsection reference.

<a id="canonical-8eea15cdae56d60b3c266d1ad79b9e2506ed3157bf35664e3c04b2cd712b4123"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced / 0a29478a3cab / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-9be598bdbed703ea40a4687371546f2016c74d6ea3c454dfe6e6f175324ccbf8)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-ac14c959af6a6af27017bb0919c10681e3e382409c59a6be3644e08bee187d5f)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9be598bdbed703ea40a4687371546f2016c74d6ea3c454dfe6e6f175324ccbf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0466dd096da591c50cbbc2c05bb5617fffc1aeb67790d8cf3729c2272835b5f8"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / c3b2754758b5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-cd9b763fc7ccd3c86c098a47ea4e3c2ef3ab42aff81bbd9be71c55714b1d81b9)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-ccf3b124f2ca83d4b03513b2231e3fd8f4b30346491f053f11289ab3f217b443"></a>

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

<a id="canonical-675d69fb873d3d1a2fdaf307a78cfa3eb486c8735129dfce5f03602eba4362e9"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / c3b2754758b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01857a45b8f6b1952b057696c29914f38a1339660efc209d0ff56788a36b19f3"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / c3b2754758b5 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-cd9b763fc7ccd3c86c098a47ea4e3c2ef3ab42aff81bbd9be71c55714b1d81b9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ac14c959af6a6af27017bb0919c10681e3e382409c59a6be3644e08bee187d5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de252ea3b30a738c96f1ec82ded7b53c0295761af8ea21de34bb813e6a78e887"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 17931d0650d7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-cd9b763fc7ccd3c86c098a47ea4e3c2ef3ab42aff81bbd9be71c55714b1d81b9)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-a69257f50e7be4d752d4e6714468acff8c58f4c7b0c53cd62279fae885b52ebf"></a>

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

<a id="canonical-b456e52f82fecf21a3272b79fcff472eae0fe5db5c3b02cc9ad5ad20029a2c69"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 17931d0650d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7c3395ac5733502d1400be713b6d21636585c743bdbfc8d474f3617a5d7c1a2"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 17931d0650d7 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-cd9b763fc7ccd3c86c098a47ea4e3c2ef3ab42aff81bbd9be71c55714b1d81b9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2952e4660e133e0b14e3bb53d417411fbdcb30b46d45b67a7c0a50574d849346"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52d9b5fa5953cc170bbb979e9e97e3542503a9ca2a5edfdd4cbd748f161c3e5c"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — performance_enhancement_mode.perf_mode_l7_enhanced / 5df5e8592a61 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-874663575881877f8abd1f2ad99250bc17cdf8988ca368465be26fdf6272910d"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

<a id="canonical-ba7c1e884ec8b42432810fe68b9e17f17095e16450a0c52554ad2a7bb05a8718"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced / 5df5e8592a61 / 3

- [jumbo_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-54654bf3f04d441dee72d3e3e40129aeaf829c6fa1166be494931239c1ebed94): complete subsection reference.

- [jumbo_enabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-933212e128dc94920b4acc8bfc3db2d1519abb3991d5bae435d8258cdff57b02): complete subsection reference.

<a id="canonical-f9c3aaa8b965f8d783d91c6c269656c0fe8dff32690163ec920121366ade49e1"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced / 5df5e8592a61 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-54654bf3f04d441dee72d3e3e40129aeaf829c6fa1166be494931239c1ebed94)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--securemesh_site_v2--reference--group-016.md#canonical-933212e128dc94920b4acc8bfc3db2d1519abb3991d5bae435d8258cdff57b02)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-54654bf3f04d441dee72d3e3e40129aeaf829c6fa1166be494931239c1ebed94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a95b41e34764ce6c7e4d7de1432dc7e3d7c8b36825b88ecde7aedc3350b355fc"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 665d28a5f5a2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2952e4660e133e0b14e3bb53d417411fbdcb30b46d45b67a7c0a50574d849346)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-380d196bb294a1b65a874ff2f8e810764c571f71ea0c6d1c7398dca02022a0ea"></a>

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

<a id="canonical-17f2fd989131458f8f32fa0c2535ebe86ba6f586f05d2e8db39b490baf0d1dd0"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 665d28a5f5a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6852df84f49399e681ac56c392acae79ecb8a2c527d71e4382bf6ac3a23c2ad8"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 665d28a5f5a2 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2952e4660e133e0b14e3bb53d417411fbdcb30b46d45b67a7c0a50574d849346)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-933212e128dc94920b4acc8bfc3db2d1519abb3991d5bae435d8258cdff57b02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21b879b64198effef002c229519d478dc398c51106c8454a63d8a4fcea695596"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / e1ad699b5d4b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-64f2ae85fa48c44bbb83615acc4fbb8fadc3ac1244737deb3cd3e8c54019a396)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2952e4660e133e0b14e3bb53d417411fbdcb30b46d45b67a7c0a50574d849346)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-301714fa9edcf95291641003e2e3adcbaf9cd773f7f7561eb32e8d57f28fd82b"></a>

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

<a id="canonical-b3b3305ed9a36de3a0e32cac84de3129ebfd976f526b03d71ceecd64a1f613a6"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / e1ad699b5d4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d61ff19a510e6cded87531935d9d1d85c729915df6132562e2e9512d7d477238"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / e1ad699b5d4b / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2952e4660e133e0b14e3bb53d417411fbdcb30b46d45b67a7c0a50574d849346)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-62dc9acc1977be7eb68ba5ade693dc3233e3ac4159fd0957f9d16196dff3b2ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-376782a8ae6f96d46106cfc1e64efbbcea4ec0476dbc1e863010ea0937f619cf"></a>

## private_adn — private_adn / a0c2ac0c6a22 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- private_adn

<a id="canonical-f49ea9315cf8fc4f7c44660c5c7bfddb1d3b4c507a389051f246fd47fd5576da"></a>

Type: `"single"`. Computed.

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Upstream description:

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4b7bdd095512b0c6ce90f6eaad7b215ee3e5af611b7df9b76270515d729e0e2e"></a>

## Direct properties — private_adn / a0c2ac0c6a22 / 3

<a id="canonical-444a0b4ea6792fcd5e9a9e99436ee5363a2f65e02955c9b682964e02002bf2b3"></a>

<a id="canonical-adc20e2c74d0b3c5890c0c8daf708dcae6cb21e8a4eaa5fae933e666af195994"></a>

## private_adn property — private_adn / a0c2ac0c6a22 / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-daa881da65cc995d881df7c7844090ca6478590b2cd715871d10819e902bc898"></a>

## Next pages — private_adn / a0c2ac0c6a22 / 5

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-995c3fca9a40cf2905d5a96884867c0a1ee30217eab9d5abee4111572773623e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e640acf0d2f8b9d89db3fb5ebd62cd19c6bc391e4190ca5458839f6704c06da"></a>

## re_select — re_select / ca5457b877fe / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- re_select

<a id="canonical-b4eb9ad4ba9a079b3c97ea8252454ade6caea09200a24a7aa67adbe4a97cfa54"></a>

Type: `"single"`. Computed.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-re_selection_choice": "[\"geo_proximity\", \"specific_geography\", \"specific_re\"]"
}
```

<a id="canonical-36becde96a14083ba6b5789de586083170f65fd49400d10b13f7eb11b1d6d13d"></a>

## Direct properties — re_select / ca5457b877fe / 3

- [geo_proximity](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c1f6c02782467e29b1a87a5f5477a4a71410a54e87817429161c4e97618168ec): complete subsection reference.

<a id="canonical-d142855128ed32d6b0f9ebfc45a5d52abfd4bde33856fbf4fd096287e2db548f"></a>

<a id="canonical-b7a83045a6ab6dab403c6b1a295e6d3c2727dcd282391e6895c3b59fcbcad7d3"></a>

## specific_geography property — re_select / ca5457b877fe / 4

Type: `"string"`. Computed.

Geographic selection for the site's Regional Edge connections.

- [specific_re](data-sources--securemesh_site_v2--reference--group-016.md#canonical-aca051c847c33fe70b4817bd7847e1604f3856ea791958145a5fb6265d72f3c3): complete subsection reference.

<a id="canonical-b5c5634caa8a4c425414371984c83314ccfaafd971fa78fea1f2cb7ba28ef8fe"></a>

## Next pages — re_select / ca5457b877fe / 5

- [re_select.geo_proximity](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c1f6c02782467e29b1a87a5f5477a4a71410a54e87817429161c4e97618168ec)
- [re_select.specific_re](data-sources--securemesh_site_v2--reference--group-016.md#canonical-aca051c847c33fe70b4817bd7847e1604f3856ea791958145a5fb6265d72f3c3)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c1f6c02782467e29b1a87a5f5477a4a71410a54e87817429161c4e97618168ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26b95678c706eb4aefeb92cccebc2013cc9b0c31abb2eebf87d96ab2ff62e735"></a>

## re_select.geo_proximity — re_select.geo_proximity / cd776e5d1d6a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [re_select](data-sources--securemesh_site_v2--reference--group-016.md#canonical-995c3fca9a40cf2905d5a96884867c0a1ee30217eab9d5abee4111572773623e)
- re_select.geo_proximity

<a id="canonical-187a157e5309e97c4bd1eedec8d4c3f98850a75a52916f62f92fdde79df7e6bb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for geo proximity.

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

<a id="canonical-a9bed57dd887770a4bafa3803bf7901695a957b53ae9450c183d49136205fa31"></a>

## Direct properties — re_select.geo_proximity / cd776e5d1d6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb5de0bd0a7402b12acb81f9eaabe2dc0eaf75410ce8326b99d85abe70d2acf0"></a>

## Next pages — re_select.geo_proximity / cd776e5d1d6a / 4

- [re_select](data-sources--securemesh_site_v2--reference--group-016.md#canonical-995c3fca9a40cf2905d5a96884867c0a1ee30217eab9d5abee4111572773623e)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-aca051c847c33fe70b4817bd7847e1604f3856ea791958145a5fb6265d72f3c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61abbe4956438c3d29a5b0d40eca13360a9783709fd512b9ed13fb4103a030a7"></a>

## re_select.specific_re — re_select.specific_re / c3930d577ae2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [re_select](data-sources--securemesh_site_v2--reference--group-016.md#canonical-995c3fca9a40cf2905d5a96884867c0a1ee30217eab9d5abee4111572773623e)
- re_select.specific_re

<a id="canonical-2ed3dce054b65c1baba2215b008550db04f018683ac7745575c02dddec249e05"></a>

Type: `"single"`. Computed.

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs.
A site will always be connected to 2 REs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-56a472d442f4c04eb308432bc4f6a6cff6f8f34950e359194d8e5d68e3fa1633"></a>

## Direct properties — re_select.specific_re / c3930d577ae2 / 3

<a id="canonical-adfbf3c1d83339d5545f9ed03c19a32c7ed11cec5a32668d9314c5a393fbe0cc"></a>

<a id="canonical-79ece1b81062bd197a5cf39fc0732a08ac1a562e13bdd5a1aa3607d7f4459251"></a>

## backup_re property — re_select.specific_re / c3930d577ae2 / 4

Type: `"string"`. Computed.

Select backup RE for this site, cannot be the same as Primary RE.

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

<a id="canonical-f6d3dbe35396062c007efee6c047b4615bd42b2ad70183498df3cab2f65cd7f5"></a>

<a id="canonical-19b528bb6370275ac2b64594c3f4576a024061c7e029628332be99960f423ab2"></a>

## primary_re property — re_select.specific_re / c3930d577ae2 / 5

Type: `"string"`. Computed.

Primary RE Geography. Select primary RE for this site.

Upstream description:

Select primary RE for this site.

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

<a id="canonical-ce680c16a0250e4f072d355e667343bb0039c3d6a67cbc6a0d8140b130760ac3"></a>

## Next pages — re_select.specific_re / c3930d577ae2 / 6

- [re_select](data-sources--securemesh_site_v2--reference--group-016.md#canonical-995c3fca9a40cf2905d5a96884867c0a1ee30217eab9d5abee4111572773623e)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa0d06e5ac50932697af6add97111fcbd044d30015089fd10d19f561de4691c7"></a>

## segment_vrf — segment_vrf / a9c94f8717bd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- segment_vrf

<a id="canonical-183316700eb619bcd6c26409791c01fb99296075d318fb834d7b73f746b900de"></a>

Type: `"list"`. Computed.

The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name.
Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per
Site that can be configured here.

Upstream description:

The Segment VRF is valid across all Sites of a Tenant. These are identified with a Segment name.
Though these VRFs are across all Sites of a Tenant, there are some configurations that are valid per
Site that can be configured here.

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

<a id="canonical-18b649b7f4d4b9289680587236e1b5db7cd9a207e1b42151cca45160e27ac111"></a>

## Direct properties — segment_vrf / a9c94f8717bd / 3

- [segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38): complete subsection reference.

- [segment_network](data-sources--securemesh_site_v2--reference--group-016.md#canonical-095eb7a1ca67df8ef66d628863d062dcb3466c1b013aa3a74b1fed4aadb8f496): complete subsection reference.

<a id="canonical-fe3ecb7356c17c9b1f6804112e07eb1c501ac90dd787ca5543530c0b643fd6f6"></a>

## Next pages — segment_vrf / a9c94f8717bd / 4

- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_network](data-sources--securemesh_site_v2--reference--group-016.md#canonical-095eb7a1ca67df8ef66d628863d062dcb3466c1b013aa3a74b1fed4aadb8f496)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ef7ceb79962e4ab83d896d999641f978e14158f51bada00327b58d899d7274f"></a>

## segment_vrf.segment_config — segment_vrf.segment_config / 5a894a787c34 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- segment_vrf.segment_config

<a id="canonical-b44ead161e135fad45088fd2b0e660f8b6533509d383a586457f0da6df37654b"></a>

Type: `"single"`. Computed.

Segment Network Configuration. Segment Network Configuration.

Upstream description:

Segment Network Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-fda48cccb133d29c74124d1dc12cf9bce17ff2cd3db4c517a95b62edf867901a"></a>

## Direct properties — segment_vrf.segment_config / 5a894a787c34 / 3

<a id="canonical-49161689d5ad5433477c162b603b92704b71ea469339123efed19f6bb6170491"></a>

<a id="canonical-a9cab630f61b667f18ee494413ecd0e04fc30826c433f4a59b043f91da257d26"></a>

## nameserver property — segment_vrf.segment_config / 5a894a787c34 / 4

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution.

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

- [no_static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a799704bea8112c2c3bd6b8e598bf603d9d460c3501c415d40f24a503d8d33dc): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-772db225186d65b0e938ecb8cd9d4105c17ebe41d226bf1431228a2d8156d7de): complete subsection reference.

<a id="canonical-0a05281448cf8269d029b91a9792a7aa7f1abdee507312e07bfbc98218f554d9"></a>

<a id="canonical-101c79fbdce470bd4c486ff60d3560f5edd8c1f4cab4b58aa1d7f610f3e957df"></a>

## secondary_nameserver property — segment_vrf.segment_config / 5a894a787c34 / 5

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76): complete subsection reference.

<a id="canonical-23e6c8bb1ce9912212c78ad27e05e5b95deb05cc33cab0eaf69ca163aca78296"></a>

## Next pages — segment_vrf.segment_config / 5a894a787c34 / 6

- [segment_vrf.segment_config.no_static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a799704bea8112c2c3bd6b8e598bf603d9d460c3501c415d40f24a503d8d33dc)
- [segment_vrf.segment_config.no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-772db225186d65b0e938ecb8cd9d4105c17ebe41d226bf1431228a2d8156d7de)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a799704bea8112c2c3bd6b8e598bf603d9d460c3501c415d40f24a503d8d33dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-491b6c1af03d5edde89ed2c1de403a7f0e33430227c2eb5f87d92484d35de6cb"></a>

## segment_vrf.segment_config.no_static_routes — segment_vrf.segment_config.no_static_routes / 849d1731d4b4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- segment_vrf.segment_config.no_static_routes

<a id="canonical-8fef76dd78063b923f8eab1d530df93c50275f62cf0c7f38c8131577d411e850"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-bd4091dff2cacb4c0d9739a3ef82f12c9a979b557efbab4d4bf6da9055882b73"></a>

## Direct properties — segment_vrf.segment_config.no_static_routes / 849d1731d4b4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9df8d80f8ec21940f2d5b5592f9b6f7dc762fc4a3bb39fa992b1147e4c60f15"></a>

## Next pages — segment_vrf.segment_config.no_static_routes / 849d1731d4b4 / 4

- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-772db225186d65b0e938ecb8cd9d4105c17ebe41d226bf1431228a2d8156d7de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5be92eb5bbae0de6398aaa66bf4d5c61d85ec3a92fd67ed569e691b6b9984526"></a>

## segment_vrf.segment_config.no_v6_static_routes — segment_vrf.segment_config.no_v6_static_routes / 8c8f901c8e2a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- segment_vrf.segment_config.no_v6_static_routes

<a id="canonical-58b897c0b5ca02fb3a09be833bf0b2a0da74eb2c872d6808205c191ae846b790"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-f7fe97b4d8792d5b1d888cd6f6666a4981bfdce0c8129830e345da021f836624"></a>

## Direct properties — segment_vrf.segment_config.no_v6_static_routes / 8c8f901c8e2a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3f3cc547f6bdd8a988228a9a0f9f4a38ab420e30ed8d50817de463c09aa8542"></a>

## Next pages — segment_vrf.segment_config.no_v6_static_routes / 8c8f901c8e2a / 4

- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-507bad4c86748898924cd3be74d2f8cfc321c9d408dc93620a0b5708d76b8697"></a>

## segment_vrf.segment_config.static_routes — segment_vrf.segment_config.static_routes / ac2b3cd50525 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- segment_vrf.segment_config.static_routes

<a id="canonical-81f4c0bb61c57570046a0d35c180f957a62dbc8a9f0112ff75cbd9ea888e2eee"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-de4d9e7ecb8bedbf74bbbdccd8c40d57fdda89210d1e0ffa1c821f80ae038354"></a>

## Direct properties — segment_vrf.segment_config.static_routes / ac2b3cd50525 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204): complete subsection reference.

<a id="canonical-64fe6dfc1a2c18f872a670ffc6b8e06ea7f9c0a7459d31e3c277632090ae28ca"></a>

## Next pages — segment_vrf.segment_config.static_routes / ac2b3cd50525 / 4

- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68b0140bd7f51074111f526e871fa22c3dd14d51688685b27b3edae11ca7e1b3"></a>

## segment_vrf.segment_config.static_routes.static_routes — segment_vrf.segment_config.static_routes.static_routes / b53e8816f213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd)
- segment_vrf.segment_config.static_routes.static_routes

<a id="canonical-1866e1df6e96d4495e3a7199491f2e59ce555b188c02de9b9ed794f55dd55d48"></a>

Type: `"list"`. Computed.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

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

<a id="canonical-329b2d0af1897d4d83afe3df87f25549a1eaceab5f7d7ebb93bb6a2e5c9575bf"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes / b53e8816f213 / 3

<a id="canonical-ba269efab34f98c68e238abca97f474b84ce972cfa1b260c19660689a5deaf19"></a>

<a id="canonical-26a1b78c78557d89928c14e17b2f6dc01343bc715e5256311676a13ab9d5762e"></a>

## attrs property — segment_vrf.segment_config.static_routes.static_routes / b53e8816f213 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c35e2f6ca3844828aa2389383767e1547ae5a41f5b09b3768c078fd30890bfb5): complete subsection reference.

<a id="canonical-4cbae4b1529fbf2f2a499de06ecc1637d5fb3353c79e59f72a4d2722f640445c"></a>

<a id="canonical-d711fde530913b124f7fb169765f2ea15ccb55a378e0d86ceedf4897c11542db"></a>

## ip_address property — segment_vrf.segment_config.static_routes.static_routes / b53e8816f213 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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

<a id="canonical-7d2ad6e9100ca4c75f36933079bc5244dc589609bfe960b0ebb9cd705845d580"></a>

<a id="canonical-7b66dbc4b19e61c942cd82f7b0032b8c14da7f8ae7c33e195a31cddc08523a85"></a>

## ip_prefixes property — segment_vrf.segment_config.static_routes.static_routes / b53e8816f213 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d975dd8fd240600bc829a66fa8c4ae069d8893254b0ac9c66440e88ce6e11ad5): complete subsection reference.

<a id="canonical-2a112565e19eccabab73ff833c0d0e3495e747b991696ab52015d806e60338d4"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes / b53e8816f213 / 7

- [segment_vrf.segment_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c35e2f6ca3844828aa2389383767e1547ae5a41f5b09b3768c078fd30890bfb5)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d975dd8fd240600bc829a66fa8c4ae069d8893254b0ac9c66440e88ce6e11ad5)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c35e2f6ca3844828aa2389383767e1547ae5a41f5b09b3768c078fd30890bfb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b3d9983a9bac097822269935d5728367b510931e512101b191b424d1f01cfd9"></a>

## segment_vrf.segment_config.static_routes.static_routes.default_gateway — segment_vrf.segment_config.static_routes.static_routes.default_gateway / a2ef3bca8c3c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204)
- segment_vrf.segment_config.static_routes.static_routes.default_gateway

<a id="canonical-2f067157fded2bacd4e0936f838ad00ce5da508c1ddcfe80d9482a5cdfb869b8"></a>

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

<a id="canonical-523d571ee00b6133cd7e6aa97a9ea56e22cf371be13440a6c62d0b91df2cd365"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.default_gateway / a2ef3bca8c3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1baa0fcb0752614b2f487b93af19d0a71c1f0035577a607e1a64e208039e4155"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.default_gateway / a2ef3bca8c3c / 4

- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d975dd8fd240600bc829a66fa8c4ae069d8893254b0ac9c66440e88ce6e11ad5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c5f0b6ef90f97a99ed59ee51a8af89a77f50e7f01e13f74d100ccd520351436"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface — segment_vrf.segment_config.static_routes.static_routes.node_interface / b5990d68cc55 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204)
- segment_vrf.segment_config.static_routes.static_routes.node_interface

<a id="canonical-c6f38c89403e31b376e08c21186990698ac42830c8eb3270a7b90588d25c88aa"></a>

Type: `"single"`. Computed.

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

<a id="canonical-6d32b6ad9526dd2e5e81b320aa5c3961fb19429279ea1a4cb407b217ee8a3800"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.node_interface / b5990d68cc55 / 3

- [list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-978845ecce07d58cd8c3216ac8873f1d26d729ce1705142e3a902295fa0814d9): complete subsection reference.

<a id="canonical-38075bfaec4510e168d4f53ba9410dd2b7ac08b1cdf2b8e66664ce6f400b6e09"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.node_interface / b5990d68cc55 / 4

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-978845ecce07d58cd8c3216ac8873f1d26d729ce1705142e3a902295fa0814d9)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-978845ecce07d58cd8c3216ac8873f1d26d729ce1705142e3a902295fa0814d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9ef61606a2a1ba0ccb1a3410f1a4d21c6bc75214ff4039a54a0f15eb25aaa59"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / d4b308096487 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d975dd8fd240600bc829a66fa8c4ae069d8893254b0ac9c66440e88ce6e11ad5)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list

<a id="canonical-8f584cc8c0a646065a8815ef9b2386b072d313646de6e2c7e68104c3bea5c292"></a>

Type: `"list"`. Computed.

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

<a id="canonical-9ec10f94449980920bea4749427fae29c810a4fe646315a495b902c7b7f42d32"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / d4b308096487 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a79609189243bac3a1fdcde7796894e17d2ef527f879e27f57699b313127b837): complete subsection reference.

<a id="canonical-5e74a22a16b76ac16009a680cfc4302750d41609099d0d3027e5bec8fd95b17c"></a>

<a id="canonical-4a772bd4d685ae974a2d52f326cf49f2d11d07c8a1a18ef484368cd4bedc93f2"></a>

## node property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / d4b308096487 / 4

Type: `"string"`. Computed.

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

<a id="canonical-bc0ea4c6e19e5bb987a5b478486b2ff22177bdad1c86a43e56c20eafcd215aa3"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / d4b308096487 / 5

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a79609189243bac3a1fdcde7796894e17d2ef527f879e27f57699b313127b837)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d975dd8fd240600bc829a66fa8c4ae069d8893254b0ac9c66440e88ce6e11ad5)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a79609189243bac3a1fdcde7796894e17d2ef527f879e27f57699b313127b837"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cced25d1c6612a77794b994c22ee5b50677ba0fe3fea94249f0392281cecf26"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-c71e49bb9ca9bf5ad2a39bdee576fd1f7b06db393598bb197f6576fb63fe65fd)
- [segment_vrf.segment_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-345ab2ff3d3aaf060b6c3caad03004a9ffec7270e3d7bc59fb8e79be3ea64204)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-d975dd8fd240600bc829a66fa8c4ae069d8893254b0ac9c66440e88ce6e11ad5)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-978845ecce07d58cd8c3216ac8873f1d26d729ce1705142e3a902295fa0814d9)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-034226cbf5a353741de4c635497bba9975b934c38a65d870680f9cefc1d8ab94"></a>

Type: `"list"`. Computed.

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

<a id="canonical-237f91828b6a51be83e6f1f04559c0aba2c71ec6a9216cf1d9b5d9da75ba9641"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 3

<a id="canonical-81f891eb4cd53e129be0e366a9556df6b90b815792453ed72239965d76390421"></a>

<a id="canonical-63b9c347554b5bc7563ba56fd27fcb05e33d1c3b340f69eceac9db7a50233881"></a>

## kind property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 4

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

<a id="canonical-16edbf8110ea21e8b830ca7e3095e758c6cda2f236971d1c369ec7a50071ae78"></a>

<a id="canonical-c8a5085ccbf95162bc46b43ba27ea16897f56f8c54acefe4a22be57797e5b5e7"></a>

## name property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 5

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

<a id="canonical-2804c6859b38367c906fc554a8df03bbbc9561c80b91e90b385fd00538354e92"></a>

<a id="canonical-92ad5f47dc86dcba232beb096be13010c738043db9cda2906221643a3bcb1235"></a>

## namespace property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 6

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

<a id="canonical-c0f6462560c8ed83391c0d1c551a6859e2edc4f4756c02b356b3898885e5e4c6"></a>

<a id="canonical-afbc72e52b9aa672c6c02269f19ee20af47e4cc999039e022aaa2947a64ce8d4"></a>

## tenant property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 7

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

<a id="canonical-f11e9d926623f95d0306c911175a7608ac0ab41faf55b55ca712f682b7a5fa35"></a>

<a id="canonical-8eea704d7cdbc2e438f28957ac11d50ec1c39d1f65d0c9ad9f3ce20e439a2471"></a>

## uid property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 8

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

<a id="canonical-97def8f1f5fd67dfa139a4ea4152dddeea81fa386426f1da01ed9964f01ec454"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 5181679cea9c / 9

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-978845ecce07d58cd8c3216ac8873f1d26d729ce1705142e3a902295fa0814d9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2647b4a96bc8674d2d2aa9e4c1b71f4a535ea0da2f8195566bc71bfb6f625dbd"></a>

## segment_vrf.segment_config.static_v6_routes — segment_vrf.segment_config.static_v6_routes / 9b4db52cecf3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- segment_vrf.segment_config.static_v6_routes

<a id="canonical-b99161fb206404018cef84adc42e3ea0ef1591de8218149b5e0faf1dfaceabec"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e64cc34000e29f621a02a96bb3dfe62c01a8fd99830eccb3ee137d59b164d7ae"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes / 9b4db52cecf3 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8): complete subsection reference.

<a id="canonical-92fa4a2d06291845d787a514c7337e63af5139d2087fe61639e10a4d901398a8"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes / 9b4db52cecf3 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-331417f427c6b45af351a7f9fe778f3b93010f175522f9ff35bf51f3638491ad"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes — segment_vrf.segment_config.static_v6_routes.static_routes / c460f56815c1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76)
- segment_vrf.segment_config.static_v6_routes.static_routes

<a id="canonical-c8492219581f9da886febd88aecab8d2aad89533063b06ca2812dedf18b2f2c1"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-88fd2c424bc79a2bef9016250ef69eb127ceadbad910f863417638862952f029"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes / c460f56815c1 / 3

<a id="canonical-ecb48d3a902555dd3e5c5ee941eb1969a0a136bfb21e13b7616ad6d7a5198fac"></a>

<a id="canonical-aeae631fd2c2c1376038d616ff440415c0226ad942abb144fb87d87ed824042a"></a>

## attrs property — segment_vrf.segment_config.static_v6_routes.static_routes / c460f56815c1 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0955f14419c49e6b88d7c79d124d131213baf8c8cb21517d1bde00880b40decf): complete subsection reference.

<a id="canonical-00ee79a661900e7a73c215adf572305cd8b202cf474bc0d77753be7aa7cb927f"></a>

<a id="canonical-4eb0e56f20de8981de5ac8e233c81824d437fdac0d615597ec2e7ee65cac1e1c"></a>

## ip_address property — segment_vrf.segment_config.static_v6_routes.static_routes / c460f56815c1 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-e3d6bc27f6cf8a7b3215ccaa385bfde7bd5ee16a9533a2f863739b07b848da74"></a>

<a id="canonical-830a529a2c211a1b7226035391f94c2082aa8b977d628769156e4fd5b3349fce"></a>

## ip_prefixes property — segment_vrf.segment_config.static_v6_routes.static_routes / c460f56815c1 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-6da357482a9d5898e58563af3493e0f546f9a0c866cb0a0ef548aa8dbe15e893): complete subsection reference.

<a id="canonical-b271390ff82a7610d48ad310cd56de9948ec32ac9420bb7701da66f1fb6f396c"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes / c460f56815c1 / 7

- [segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0955f14419c49e6b88d7c79d124d131213baf8c8cb21517d1bde00880b40decf)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-6da357482a9d5898e58563af3493e0f546f9a0c866cb0a0ef548aa8dbe15e893)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0955f14419c49e6b88d7c79d124d131213baf8c8cb21517d1bde00880b40decf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd2d5b0efefff6ef5987c43b9768a864dbc8df54394f1d3a7fe427bf9b90749b"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway — segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway / 37f78b8b50fc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8)
- segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-43e43e007cc8900e22bc9a05b6bff6416320896306b9c056352179cd78891d5a"></a>

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

<a id="canonical-bf549efc12a56981d791d07be06213897efa59faabc5d75713ef818364e1403c"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway / 37f78b8b50fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd69774f35047bbe0e5b88ee91c9f4bbf84155af412b88c83a32697cfd528614"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway / 37f78b8b50fc / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6da357482a9d5898e58563af3493e0f546f9a0c866cb0a0ef548aa8dbe15e893"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9295bab299dc5618d52ec2ec76a19512a488a79b6884979c78bb4a7e1b646e71"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface / 6df3d20001e4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

<a id="canonical-d8d4e7414f526836da034bc8ee1b94e9ef4635bc61cd854b892402be9eca657c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3bc638dc5d6cb8580cb47f0eed04b65ebf2fa84e1fad25f02eebdba2399a5dfb"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface / 6df3d20001e4 / 3

- [list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-95143c62dfbd9931fabfece3abda5d3c20f232f4e94878a9f99862e546dbf589): complete subsection reference.

<a id="canonical-29c87e34949e6c76e1ba81c8af746ed536b4faaf276cb265aa74e092ca2cb753"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface / 6df3d20001e4 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-95143c62dfbd9931fabfece3abda5d3c20f232f4e94878a9f99862e546dbf589)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-95143c62dfbd9931fabfece3abda5d3c20f232f4e94878a9f99862e546dbf589"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-733ece96cb72422a410b5c337c4ea5ad4f839c72cb90218b40676b48f0b9ee52"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / d402de03acd3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-6da357482a9d5898e58563af3493e0f546f9a0c866cb0a0ef548aa8dbe15e893)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-ed8641bc8d5b021b955887f54773267f6e109e645df2c9c781731047f2ac226c"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2c9ba06cfd4b04d7bbbbf7eed62d7a3853900031042be8b4c42547a49f480684"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / d402de03acd3 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0bc360aef71611c2c6574992ac721950a798a836bc0895d17eaf0ee0eae860a5): complete subsection reference.

<a id="canonical-4d9e4748f4b7b1e841a4600ae2a9b6577a311f0d3e4d5e80b6a9bed03af0cc54"></a>

<a id="canonical-dc39e8e8b5c08c9d515427b560fda8f869f2c436e0d7e194452d3ec271b9d6f9"></a>

## node property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / d402de03acd3 / 4

Type: `"string"`. Computed.

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

<a id="canonical-793e2fc8cad11647d59acb59645823b076f2ef411914eeed8173f2b56df86ad5"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / d402de03acd3 / 5

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0bc360aef71611c2c6574992ac721950a798a836bc0895d17eaf0ee0eae860a5)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-6da357482a9d5898e58563af3493e0f546f9a0c866cb0a0ef548aa8dbe15e893)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0bc360aef71611c2c6574992ac721950a798a836bc0895d17eaf0ee0eae860a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7afa9d52bb4a1ec398e6976bb475c51c03558e215a419e0d13885a0e3edd381"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [segment_vrf.segment_config](data-sources--securemesh_site_v2--reference--group-016.md#canonical-06230dead5c76fb07643002b3e5eed309c79887dc11b691ee381c98d97a6bd38)
- [segment_vrf.segment_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-736247ce08c5f03bf5e682b2d7a89be4a7dc2f7250feffaa19fe83979ffbca76)
- [segment_vrf.segment_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-016.md#canonical-476fd61a0f59ce457fc157fd79a441c8402ce2c1bc384c03640c17fdb942f5a8)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-016.md#canonical-6da357482a9d5898e58563af3493e0f546f9a0c866cb0a0ef548aa8dbe15e893)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-95143c62dfbd9931fabfece3abda5d3c20f232f4e94878a9f99862e546dbf589)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-cac9d278a1a794605d9cf39a44846cc765f3d269d8d0108bd2c8db7f425b1344"></a>

Type: `"list"`. Computed.

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

<a id="canonical-6ba7b02ffc3f7a6d72e5e93c867f55e48862645fa0ac0aa6fe5a840f9be2741e"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 3

<a id="canonical-be13b06f0f9ace3606e0e76e7b8661b51e676c5de861dc255a1a6ff664abeff1"></a>

<a id="canonical-fe7e744347aec59c501ed3e7267f68cf8d2bae6eebc2abaff8b493c46512b1e3"></a>

## kind property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 4

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

<a id="canonical-2b6e98f260d317865a11e118c572f42c8b7ce318572c131e555d51fc705eebe8"></a>

<a id="canonical-ccb530d5c8dbfd60772da9eaf9dc49ad362ffd07837a7e55f6201d330216656b"></a>

## name property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 5

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

<a id="canonical-95aa7b4bfed61ac1c9effb4c8663a01c692bf0cdcd92a6a26549992905b1da08"></a>

<a id="canonical-dceb5ceb5b64dbdd54c92233b21537c7bb4f9762954eba9315a1e37828b5df1c"></a>

## namespace property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 6

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

<a id="canonical-13dfd263670c860a111d95533c47e74ac1fdb2c37c3b0726895a3c77f72b78a9"></a>

<a id="canonical-3a6478124ba140bbe9400eb07c37ce9d9d7c43e7709b13906b0d5bd67e9f6c7c"></a>

## tenant property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 7

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

<a id="canonical-f9a1f88bb763c86f0b7e8683457cde0a92ac2abf09e13706c2dda4509f8995ff"></a>

<a id="canonical-30f29ef42595dd7c8dbf38f72b406b8143a101fedde510799de1ec1a067c83f3"></a>

## uid property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 8

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

<a id="canonical-e51dadd3ec9ba41e89ea08488edcbee9833b49af37edb5ceb72df5711bf99d3d"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ff8a31c0a851 / 9

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-95143c62dfbd9931fabfece3abda5d3c20f232f4e94878a9f99862e546dbf589)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-095eb7a1ca67df8ef66d628863d062dcb3466c1b013aa3a74b1fed4aadb8f496"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5015238df474520951684c1d454ece3f0578af3d13e0dbabfa2f0e90982bb92"></a>

## segment_vrf.segment_network — segment_vrf.segment_network / a40be3ab7157 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- segment_vrf.segment_network

<a id="canonical-30e29aa34f2121d98962b61b5b1f82da9e44f4066a6a64371dbce0b9a7eaa352"></a>

Type: `"single"`. Computed.

Type establishes a 'direct reference' from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name for public API and Uid for private API This type of
reference is called direct because the relation is explicit and concrete (as opposed to selector..

Upstream description:

This type establishes a 'direct reference' from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name for public API and Uid for private API This
type of reference is called direct because the relation is explicit and concrete (as opposed to
selector reference which builds a group based on labels of selectee objects)

<a id="canonical-a20ca969a7e89c0290d2edc581ac5906f57b248b390ec5f295c35c0810ca4e63"></a>

## Direct properties — segment_vrf.segment_network / a40be3ab7157 / 3

<a id="canonical-177e8467877fb209ee751d69c8105bab1173fba5e36836fcc6e381719a50ff76"></a>

<a id="canonical-48e08ad08ae7f5645f7756c820a5e3750e80f40366b0504d5343566d4ca41c37"></a>

## kind property — segment_vrf.segment_network / a40be3ab7157 / 4

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

<a id="canonical-50ba6c96ac2596a6230544ef11d3ac15bebcce4d1e1e670c59e84b19103df426"></a>

<a id="canonical-1aabe61b3ebdca1b2df8eee777c5e0ec6eb57760a2d411dd8476e0785392ddba"></a>

## name property — segment_vrf.segment_network / a40be3ab7157 / 5

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

<a id="canonical-73d6360fad1a0354be27607862c4e1e6e6b2a5cffc9574221cc0f29a66980984"></a>

<a id="canonical-47516cec0b56ebfe31d4c1a9951d977ebe73278d0df8398fd9810b52a3e62cc4"></a>

## namespace property — segment_vrf.segment_network / a40be3ab7157 / 6

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

<a id="canonical-f368c2f6dc103848adbc5d1b4a21cedf4f1b68c09a450d96c0df94167ca416f3"></a>

<a id="canonical-26a3f94bb2dc95d49321db19db47cc5e922b794a347b1291644adba2a7c88a66"></a>

## tenant property — segment_vrf.segment_network / a40be3ab7157 / 7

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

<a id="canonical-77610fb5d8a44bec39f9799a66fa6561a9c6e7aa7904af9d72087b714401d092"></a>

<a id="canonical-f48065056f4b03b16338fc0e45f6981bed45347a2b6829ecaab095981e8dbe0e"></a>

## uid property — segment_vrf.segment_network / a40be3ab7157 / 8

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

<a id="canonical-4ee438efa69a81ef38fe16e177331bf60dcd0f41dbb61187935813194ea5e473"></a>

## Next pages — segment_vrf.segment_network / a40be3ab7157 / 9

- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2c730e24052cdf5bed8649e2387ffe17dc259ad692e7eedd4a9c2f8d3fd9cbf2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-003cce9d5bf1a9cf834f34aa208b338d566941fdb9c5efa052d7586bfa84906f"></a>

## site_mesh_group_on_slo — site_mesh_group_on_slo / efa0deb0f521 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- site_mesh_group_on_slo

<a id="canonical-e96691a958d856d3aeb554371c04a13896cf452a2de85b05476675125addf845"></a>

Type: `"single"`. Computed.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

<a id="canonical-b4a2144b1abc9c2631763d17b1cbce9ea6d5b9e8bab3df9f27c93483525797af"></a>

## Direct properties — site_mesh_group_on_slo / efa0deb0f521 / 3

- [no_site_mesh_group](data-sources--securemesh_site_v2--reference--group-016.md#canonical-7d2cb5b44f1dc755d8f531971effde7bb81778242e2c3c5b6277d1cfcb688a32): complete subsection reference.

- [site_mesh_group](data-sources--securemesh_site_v2--reference--group-016.md#canonical-346dfdf4d166e34872bf10997555df8d2898102b1c39c53cd9fa2070753d75d9): complete subsection reference.

- [sm_connection_public_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-099bc0c92d214bfb577f9308c27dd714ca6ea9f572afd9947275838f72120b2a): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-6a3b5828b5a6586a291568e6f06c51db5b6855a58e105d700ff952e7822f8440): complete subsection reference.

<a id="canonical-24c2cd7989b75fb940bfe24db1096f19d426bef3e4b7b365c58af293f594a0d5"></a>

## Next pages — site_mesh_group_on_slo / efa0deb0f521 / 4

- [site_mesh_group_on_slo.no_site_mesh_group](data-sources--securemesh_site_v2--reference--group-016.md#canonical-7d2cb5b44f1dc755d8f531971effde7bb81778242e2c3c5b6277d1cfcb688a32)
- [site_mesh_group_on_slo.site_mesh_group](data-sources--securemesh_site_v2--reference--group-016.md#canonical-346dfdf4d166e34872bf10997555df8d2898102b1c39c53cd9fa2070753d75d9)
- [site_mesh_group_on_slo.sm_connection_public_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-099bc0c92d214bfb577f9308c27dd714ca6ea9f572afd9947275838f72120b2a)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](data-sources--securemesh_site_v2--reference--group-016.md#canonical-6a3b5828b5a6586a291568e6f06c51db5b6855a58e105d700ff952e7822f8440)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7d2cb5b44f1dc755d8f531971effde7bb81778242e2c3c5b6277d1cfcb688a32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22f104ce3e915201fd42af3fa39dd3487ffc4da25b1e05798fcd7a7755280975"></a>

## site_mesh_group_on_slo.no_site_mesh_group — site_mesh_group_on_slo.no_site_mesh_group / 69ed79ee600e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- site_mesh_group_on_slo.no_site_mesh_group

<a id="canonical-fbf0bcb0b61c2d011ec7a406103f46a3bed289c432018f77900e5c51bdae0a03"></a>

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

<a id="canonical-0dab4b2cf9761902394d7e84bd34b1ddd0c982c98f9a789a07b2148ed13165c7"></a>

## Direct properties — site_mesh_group_on_slo.no_site_mesh_group / 69ed79ee600e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8195708c9299d2d603674816186b05cded1f3a5587aee21e6c207b1e4be55476"></a>

## Next pages — site_mesh_group_on_slo.no_site_mesh_group / 69ed79ee600e / 4

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-346dfdf4d166e34872bf10997555df8d2898102b1c39c53cd9fa2070753d75d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc1d231fb653e7e5ad4c741b2be7f0313b298a964454bb4dee53e2d5f13ca15f"></a>

## site_mesh_group_on_slo.site_mesh_group — site_mesh_group_on_slo.site_mesh_group / 959dfaceba8a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- site_mesh_group_on_slo.site_mesh_group

<a id="canonical-7949635ff24e0df0a4a0ca6e0f33461e7cf433b8e963f1dcf534b7cea82aba2f"></a>

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

<a id="canonical-f0efa2aeb375c24d1a632239c418f71f7717f7d099f333e8e75587dee51a9bc5"></a>

## Direct properties — site_mesh_group_on_slo.site_mesh_group / 959dfaceba8a / 3

<a id="canonical-250968f2f026e23ef437b002e16aea5e2fb79a140fb8ac0edce6da1f9f1d358a"></a>

<a id="canonical-637e143e8fb16f72089ff1a5d385dffbe9027bc855da7eff1e6c2402aaa6c33e"></a>

## name property — site_mesh_group_on_slo.site_mesh_group / 959dfaceba8a / 4

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

<a id="canonical-a592c55af348d3b6b22444547b4af49abe4a353f8839d5d71e01a463065e5ec8"></a>

<a id="canonical-63bb3e98810b767bb0782748f1ffe68e9fb364d4b81a113a3b3ec8dbffa4a82e"></a>

## namespace property — site_mesh_group_on_slo.site_mesh_group / 959dfaceba8a / 5

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

<a id="canonical-0caff155de7415001e8f7e681cc9f28d994b76781fe41c65f17521d080ebfe58"></a>

<a id="canonical-a68683c171284eef54ee54428433dc83aefac98fb01323df95bb733f62448473"></a>

## tenant property — site_mesh_group_on_slo.site_mesh_group / 959dfaceba8a / 6

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

<a id="canonical-aec300144e424caf5ac099c946d169981013d9c75a116852136b22465dd505dd"></a>

## Next pages — site_mesh_group_on_slo.site_mesh_group / 959dfaceba8a / 7

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-099bc0c92d214bfb577f9308c27dd714ca6ea9f572afd9947275838f72120b2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c26f270f87cbe168eafb6679abc7d0245e9c7443d94045eb5e7f709b4a9afbf"></a>

## site_mesh_group_on_slo.sm_connection_public_ip — site_mesh_group_on_slo.sm_connection_public_ip / 665ec1af34c2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- site_mesh_group_on_slo.sm_connection_public_ip

<a id="canonical-15e8eac254deedad64b1752d106f55f20e695327c88467fa9783913c325c5419"></a>

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

<a id="canonical-c9aaacb0fa6ef3af0beba56d302c434c3347bb45f38fc26a83765d5f99f2116d"></a>

## Direct properties — site_mesh_group_on_slo.sm_connection_public_ip / 665ec1af34c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0a45c030a0367b5b42e12a996496758964fd909d46c06c0e07ff569ba6d38a5"></a>

## Next pages — site_mesh_group_on_slo.sm_connection_public_ip / 665ec1af34c2 / 4

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6a3b5828b5a6586a291568e6f06c51db5b6855a58e105d700ff952e7822f8440"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9574f7283f698e54c3b814c655645ad3b431a2688b2b8483f2e6e040e8cb8a6c"></a>

## site_mesh_group_on_slo.sm_connection_pvt_ip — site_mesh_group_on_slo.sm_connection_pvt_ip / f31340f95b23 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- site_mesh_group_on_slo.sm_connection_pvt_ip

<a id="canonical-390c0f5620d9a5204f1188b40ecd8d3c251bff7ade0c4121fc6dddd88948c7bb"></a>

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

<a id="canonical-0c824960f406c69e8ad22a7744346a816a6fafcf12b6fcfbc14a85fad47d4e10"></a>

## Direct properties — site_mesh_group_on_slo.sm_connection_pvt_ip / f31340f95b23 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce81f6c156958710d59576d8416f1de29fcbbc5da7fcccb6b04731a525de1ae1"></a>

## Next pages — site_mesh_group_on_slo.sm_connection_pvt_ip / f31340f95b23 / 4

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1cc3e86d92efd4ec57d162359e35a8358882499a7f7471702a4568e46798b04d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45e24600398bc6c386ef1bb6cc2ec3feac490c39250a42ebeccb1d1c98384898"></a>

## upgrade_settings — upgrade_settings / 806ed5f429c8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- upgrade_settings

<a id="canonical-c00c4bd4dabfbca9fc2431aaa2ed0b6abd6b2f9124acf75386b1a34ac7047569"></a>

Type: `"single"`. Computed.

Configuration parameter for upgrade settings.

Upstream description:

Specify how a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0b62dc53d0b3bf6a5a1944f0a0c56d464d89de842dff84c632356bcd1184231d"></a>

## Direct properties — upgrade_settings / 806ed5f429c8 / 3

- [kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115): complete subsection reference.

<a id="canonical-4f94716b9569dbabc1fdca8b18ab17fd874637dd723c9a38dd4941ff814e3134"></a>

## Next pages — upgrade_settings / 806ed5f429c8 / 4

- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd00e7b22ead0f62f6cf79998dc376ffb97dfe09a05435e1a69bc9043c587ea9"></a>

## upgrade_settings.kubernetes_upgrade_drain — upgrade_settings.kubernetes_upgrade_drain / 1fd9ad21f0cb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd)
- upgrade_settings.kubernetes_upgrade_drain

<a id="canonical-dd60205ea2f00d21cec03ae037ba7e5284bf3e190a7ed68093f10d3b59990895"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

<a id="canonical-275fc88510c2271e563647c92ed96e23e8721b45764c37ede0eb5d4d2f1af0c4"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain / 1fd9ad21f0cb / 3

- [disable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-510b149a8dc7967f95154a92840cfece12cd07f9086fb3bc3f5388f405a5e952): complete subsection reference.

- [enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0992bae962531dd3013eea58df2b45456cc60aeac107d34ec98d919ea3d7bf5b): complete subsection reference.

<a id="canonical-e4dec142894b1a5fac921b1e4a884c7190e2986606b6a02e0e52d504c737b26b"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain / 1fd9ad21f0cb / 4

- [upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-510b149a8dc7967f95154a92840cfece12cd07f9086fb3bc3f5388f405a5e952)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0992bae962531dd3013eea58df2b45456cc60aeac107d34ec98d919ea3d7bf5b)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-510b149a8dc7967f95154a92840cfece12cd07f9086fb3bc3f5388f405a5e952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4df68a464d2cbd3a6f9e7620e4aaa0c3e5a5a41a73bcbe269fcd3baeee23ca9e"></a>

## upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain — upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain / 4540be96d2eb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115)
- upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-c4af24ed199ab57f0f82c5d7b60248cf47351b7530238ad529e20bd32f540d3f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable upgrade drain.

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

<a id="canonical-f06bc44ce4a3007aa8d3f9f1209b9f50b6b42b8feacd94aa742ce4d573206876"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain / 4540be96d2eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20ca21e21172aba1db9edc43b94d62d8ccd70c7588bf9a3e900ff4c22e923240"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain / 4540be96d2eb / 4

- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0992bae962531dd3013eea58df2b45456cc60aeac107d34ec98d919ea3d7bf5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-810bb95f844092f2a2ba17e33b59653b344843d878e5d403dd429aecd632e137"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 01da15c2df7d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-300f538bf2d87fd547c29ffd5e02d97c1dd33642f343bc8a298dd866745e9da6"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

<a id="canonical-cd1bfd0d3678c8c23ade3b8006be441ad5b8a834d6e7397201dd18de716bb661"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 01da15c2df7d / 3

- [disable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-402def5432b6963121e9777e0e55be6d2fc8bb35191a094021e31ba4a0c30daa): complete subsection reference.

<a id="canonical-35e920f9b19a480c249a5fea3b023ccb46e4946f3045e3dd501caef734b25357"></a>

<a id="canonical-a3a2c66272d853f794b3f44e25db085034ac4dac95e32136379e454e4f8f4825"></a>

## drain_max_unavailable_node_count property — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 01da15c2df7d / 4

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

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
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-ff10a6f2c7b5ef44fe3a0b91f8a8e614832c75aea923a914fc9ab777b19a1a74"></a>

<a id="canonical-f1ed3fd77ff55999481772da4629eb830a4c44ede454a1f60d2924f13cc5674c"></a>

## drain_max_unavailable_node_percentage property — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 01da15c2df7d / 5

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-73107082b91b666ba58f73369b21c1bc7024f49961d9e50a08fbf6e9596b9d71"></a>

<a id="canonical-7b4c39ae076ab109b78e031590d1ffd026c473b1f2e42b66e85448490464e438"></a>

## drain_node_timeout property — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 01da15c2df7d / 6

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-b0fffa0fa9382edd7070f00876eaf7947993e465ce83c5591114513f1e447ea4): complete subsection reference.

<a id="canonical-d1d466b7b469e57bbaba5982b516d6b474ce7ef3b01f0feee902713f010a9264"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 01da15c2df7d / 7

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-402def5432b6963121e9777e0e55be6d2fc8bb35191a094021e31ba4a0c30daa)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-b0fffa0fa9382edd7070f00876eaf7947993e465ce83c5591114513f1e447ea4)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-402def5432b6963121e9777e0e55be6d2fc8bb35191a094021e31ba4a0c30daa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d22f4a23cac9495b1e223cedf56f630eae50a25b4c4cded2539f54db1f95a9d"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgr / ce3e6787965d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0992bae962531dd3013eea58df2b45456cc60aeac107d34ec98d919ea3d7bf5b)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-7c8ca8c64fa0f6d5c0eda1cf326edce9801c668f2bd350db80ddf35ffbb4996d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable vega upgrade mode.

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

<a id="canonical-bdd65c1a5edbc2a3b6def7db1b0e754ba026365b83b0ff7f9101dcab68fba64c"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgr / ce3e6787965d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97ba34e0126714863770f6177e334bdaa53dcf7c469dd16f08c18bebb8ba73e4"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgr / ce3e6787965d / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0992bae962531dd3013eea58df2b45456cc60aeac107d34ec98d919ea3d7bf5b)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b0fffa0fa9382edd7070f00876eaf7947993e465ce83c5591114513f1e447ea4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6dece7a4145b38647f0274ced0222543b68e5fb547c6f1894cf9cd82cefdeaa"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgra / 171004b627cc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0c41f12660315276db3be349044c62d8a28c6e5717bd4e1297c55801131822dd)
- [upgrade_settings.kubernetes_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-75db66de9e68164bb4bac00dfd3f06d596fd936a99f0e36928cf3fcaf01ea115)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0992bae962531dd3013eea58df2b45456cc60aeac107d34ec98d919ea3d7bf5b)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-9916b15b726e89433e60ef26815ea738d4d3c10ede9a71ec7cf478f7d819fe5b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable vega upgrade mode.

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

<a id="canonical-28d9d73ac48d98f57637fb658bc57ab0ec22122ce39fdbe8f19d6a69edcfa07c"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgra / 171004b627cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65bcaa52e8a17e9f05fa8e4c48f0ada2f7ace9b7386528120bbb4acbf5c1be53"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgra / 171004b627cc / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0992bae962531dd3013eea58df2b45456cc60aeac107d34ec98d919ea3d7bf5b)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b0f00bb1aa2d2d981d53a70728c56b6a46ac96dcb599531fa47a2772354a7fe"></a>

## vmware — vmware / 5a68d32e3f9d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- vmware

<a id="canonical-400afa49d3081b8723b25c4a02e8658a13b95844a2fbac64934a2f0c064f913e"></a>

Type: `"single"`. Computed.

VMware Provider Type. VMware Provider Type.

Upstream description:

VMware Provider Type.

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

<a id="canonical-1f842f36d2da3bb632b9874ad0251283346a122b8a618edd69645e950d35d983"></a>

## Direct properties — vmware / 5a68d32e3f9d / 3

- [not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6): complete subsection reference.

<a id="canonical-e56dca07b1673270fbe84cbd121cf5b028af5717f493c95b8e4eae5f4e925568"></a>

## Next pages — vmware / 5a68d32e3f9d / 4

- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b865ff973357a8ed7bb8cc0ea0aa741491bc5bc98fde93f25ed504875298f24"></a>

## vmware.not_managed — vmware.not_managed / 9702a51d7a27 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- vmware.not_managed

<a id="canonical-67277ab4d9ee724bd0405317a36fe9d2c298dac13e3fbd3c8d2c0c25ef20e893"></a>

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

<a id="canonical-09e3493f453a7f0d19cf48e7035888d0fa3bf6b4d4f94774fffcb34123a50fe9"></a>

## Direct properties — vmware.not_managed / 9702a51d7a27 / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1): complete subsection reference.

<a id="canonical-ee3f711b1f954fb09944b50e8d06b011a839f6193ddf5f2cc46d0dc1c6479097"></a>

## Next pages — vmware.not_managed / 9702a51d7a27 / 4

- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1717e310237e21d4c015b1360d428773c840dde5cdbfad14838e045badb957d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6aaa48abef1d1b6e282d0a600d97c4305b1190708c4030e16201095ac1dde1a"></a>

## vmware.not_managed.node_list — vmware.not_managed.node_list / 32b26ba7282c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [vmware](data-sources--securemesh_site_v2--reference--group-016.md#canonical-a40005702c893ea5cf5b53bc3bb512942b2d9786f57b8397e8b8deaeba7a4fee)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- vmware.not_managed.node_list

<a id="canonical-fb967e4b301bdea2d77bc66b22b5e33bd48d90630c01a842ddf6d88a3a767cff"></a>

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

<a id="canonical-adfbeeea59b4e9033575d46c034d46a663fb4a9b924bb2e9b5cc78af3d8e0ea8"></a>

## Direct properties — vmware.not_managed.node_list / 32b26ba7282c / 3

<a id="canonical-7ce30792c385f9efcf82be92243f98327107b354e021129d5a80be1ca8627f05"></a>

<a id="canonical-ba8b7eb4bf8aa59dfd45554e1f2a965124cbb62c2faffaf92d4c29a619e35647"></a>

## hostname property — vmware.not_managed.node_list / 32b26ba7282c / 4

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

- [interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737): complete subsection reference.

<a id="canonical-d6d9c35dd7b6f9cfdc700c64c86e602582f6612389e254a4706e9791a81a4d34"></a>

<a id="canonical-f3e33026877e6211ac1d887bb406893dd3db84cb58fe093051ec14f1e32d3d8d"></a>

## public_ip property — vmware.not_managed.node_list / 32b26ba7282c / 5

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

<a id="canonical-fdb846b6331aef6a7907ba058947f4e11af6cc4a3a83c1b19f09a9f4b841ae82"></a>

<a id="canonical-c5cf8ac64f75041d6384c691acbb3556dc20c465500d362d0c679235f1323094"></a>

## type property — vmware.not_managed.node_list / 32b26ba7282c / 6

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

<a id="canonical-923e20c55a044bb1dad122da5f34238d050a67d78e27ae2ec7aaf3531f74e849"></a>

## Next pages — vmware.not_managed.node_list / 32b26ba7282c / 7

- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-016.md#canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-016.md#canonical-8e03e4e0991fd9ffcfd09d5dd56ddb7e84065017268d8f0f6af6b92ae3cdd1e6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-249f885a042353d5ba1541c02458289c62f7563e9b4ffc7da6b521ebff7ff737"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
