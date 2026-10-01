---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-21bb313c6ee9290905a326e7b0aafe36fec0bac71d7d6099c60276c5a8d17d1c"></a>

## nutanix.not_managed.node_list.interface_list.static_ipv6_address — nutanix.not_managed.node_list.interface_list.static_ipv6_address / 9e21582a26d8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-afaa4ccccb7523424e4df02b4f1513c21cf529a0d8c2065d006bb7faff828456)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-27e834315878a609b33a46aa2d1dcf3dec6ae5dcbaac86662f0fc4a8b8fb61a4)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-9a64d29ca0f3dc0fea1da6aab9c87cdad5378ba770f09e4b8381dd5ccecd8527)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-8797a45af75a8995088809c546f2e21ec6a62806d33b715ae6456cb9838145ae)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-cad61617e97fbd1ae5a752ea0f3350a4de83d1574c9158a48f7ef8314d56af7f"></a>

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

<a id="canonical-e21e2db4eb4983f091019571654382a84348c2d0c92e0274f9f13416eb8b5a51"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.static_ipv6_address / 9e21582a26d8 / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-05da872b55ed51d41b0d8f0a9b720671dbf12addae9440c2338617f685eecf1a): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d4584ba49815db31a8c98c65ae0d05c551c3eb0e5a7e062c83498c9e95abaa8c): complete subsection reference.

<a id="canonical-17ed399fe6383008c10370ba8a943dd6edcfcdf683f3b0bd7879f0b4f0182cc4"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.static_ipv6_address / 9e21582a26d8 / 4

- [nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-05da872b55ed51d41b0d8f0a9b720671dbf12addae9440c2338617f685eecf1a)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d4584ba49815db31a8c98c65ae0d05c551c3eb0e5a7e062c83498c9e95abaa8c)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-8797a45af75a8995088809c546f2e21ec6a62806d33b715ae6456cb9838145ae)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-05da872b55ed51d41b0d8f0a9b720671dbf12addae9440c2338617f685eecf1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dec2a8e07d1ed05c084bef78e8390dff3154227703613859a96026cafd2bcc59"></a>

## nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 08bdad6959c0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-afaa4ccccb7523424e4df02b4f1513c21cf529a0d8c2065d006bb7faff828456)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-27e834315878a609b33a46aa2d1dcf3dec6ae5dcbaac86662f0fc4a8b8fb61a4)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-9a64d29ca0f3dc0fea1da6aab9c87cdad5378ba770f09e4b8381dd5ccecd8527)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-8797a45af75a8995088809c546f2e21ec6a62806d33b715ae6456cb9838145ae)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0350e9b79448f3304dfaa9688f059533f0b7947f21a22034e7e219cfdeb9552c)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-428a15ff83e609917ea728b2d6e28f89203f6bb7032889f8b0ec03879ece68d5"></a>

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

<a id="canonical-33d1eade1256b270756e9f21d0b205f045fc5efa6c3adb299fd1e782e582bd83"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 08bdad6959c0 / 3

<a id="canonical-dd09e2d45fd3dc703cf947e97e4df12ab8c4ac2a9e1adbd9135c469b407c26d4"></a>

<a id="canonical-64cee2a37b2c1f2886b384aaa3987147b2f5ef4a971a9fb20a30f7bdbd5514e3"></a>

## interface_ip_map property — nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 08bdad6959c0 / 4

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

<a id="canonical-89849e20af6d3620000c4664642f01da64c7581e02fbc9166868de68ecc323c0"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 08bdad6959c0 / 5

- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0350e9b79448f3304dfaa9688f059533f0b7947f21a22034e7e219cfdeb9552c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d4584ba49815db31a8c98c65ae0d05c551c3eb0e5a7e062c83498c9e95abaa8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00842e5f4150b852485ec191848946e77daaf95c195a2ad13810338119539d86"></a>

## nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 2c5e7ddf85ff / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-afaa4ccccb7523424e4df02b4f1513c21cf529a0d8c2065d006bb7faff828456)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-27e834315878a609b33a46aa2d1dcf3dec6ae5dcbaac86662f0fc4a8b8fb61a4)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-9a64d29ca0f3dc0fea1da6aab9c87cdad5378ba770f09e4b8381dd5ccecd8527)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-8797a45af75a8995088809c546f2e21ec6a62806d33b715ae6456cb9838145ae)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0350e9b79448f3304dfaa9688f059533f0b7947f21a22034e7e219cfdeb9552c)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-9bffc3fef36e52080032c27a0a0e5e57dee25af861cc0d7fe9893be93c981c17"></a>

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

<a id="canonical-b5a1dec81897b6022b051fdd8c923a98c4de965072d460f1f4bf4a5b93aeb8be"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 2c5e7ddf85ff / 3

<a id="canonical-1713823ce3a4a6a968ce62112901795fe79eb9eea281066264fc9c094789ccab"></a>

<a id="canonical-4182992b8298d0d95fa3325989e4024626dc4631b433f25eae4b3bc1f397ec43"></a>

## default_gw property — nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 2c5e7ddf85ff / 4

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

<a id="canonical-bf6378737fe9265ecabe61557e1c94bef6565c575bf54f92638f93fa25a6fe0a"></a>

<a id="canonical-47ba3f349f9b23eab2897034043f263a34374b78f80198ec9c3d1eb4978b36a5"></a>

## dns_server property — nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 2c5e7ddf85ff / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-6e806c944320fc2297f8694ec26f8dd286713a063de1102f0624227ba9d1fe50"></a>

<a id="canonical-fba095bafcac4adbafc4868e6f654f7b8f587c6009bac60bde357cc4e6d0bba4"></a>

## ip_address property — nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 2c5e7ddf85ff / 6

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

<a id="canonical-20d256f4f0fdfe0106bccd356110a825329df1a08ea89bb6f6aab2128bd9d8c2"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 2c5e7ddf85ff / 7

- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0350e9b79448f3304dfaa9688f059533f0b7947f21a22034e7e219cfdeb9552c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a64ac0e5e1c196d590761c3a5c0d10a75c39c8c3eff14f66ca47d36fbaec8df7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e2a3f2921adcb8bca825a7f865c418c1a09f985b28de61c6e508ef395015abc"></a>

## nutanix.not_managed.node_list.interface_list.vlan_interface — nutanix.not_managed.node_list.interface_list.vlan_interface / e999206d0f21 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-afaa4ccccb7523424e4df02b4f1513c21cf529a0d8c2065d006bb7faff828456)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-27e834315878a609b33a46aa2d1dcf3dec6ae5dcbaac86662f0fc4a8b8fb61a4)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-9a64d29ca0f3dc0fea1da6aab9c87cdad5378ba770f09e4b8381dd5ccecd8527)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-8797a45af75a8995088809c546f2e21ec6a62806d33b715ae6456cb9838145ae)
- nutanix.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-cf8529495f74e41524dc6304618ff16580a98cda045e4f7fab2e6bdbe74fbf53"></a>

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

<a id="canonical-22aedf0c541f79de6ddc46ff95054df9a656acde7ab7cc5cf802190c6487b0ea"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.vlan_interface / e999206d0f21 / 3

<a id="canonical-5746153c9c6f66b2ae7b0e5808eca5e33d8632bfb84b682f6e27c25f6f8f853b"></a>

<a id="canonical-2bdf7021aacd45cc93a2880c142f3907d42c0052df5a6709aaa92c6ab2350af7"></a>

## device property — nutanix.not_managed.node_list.interface_list.vlan_interface / e999206d0f21 / 4

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

<a id="canonical-8cbac982fcc2ccaef7f3edbf87dfded8659759637df4e147b69a29b550c096ca"></a>

<a id="canonical-75c54b0973a38f9bc4b6560a4f42c2310914c03eb9e94030c5ec60b9400243fc"></a>

## vlan_id property — nutanix.not_managed.node_list.interface_list.vlan_interface / e999206d0f21 / 5

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

<a id="canonical-2ceddd93544a7a60fc18015082d205133631b9ad9512b2642b8569b1ed9eecde"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.vlan_interface / e999206d0f21 / 6

- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-8797a45af75a8995088809c546f2e21ec6a62806d33b715ae6456cb9838145ae)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-438e7c7fc97ea4a3f3620c1e386bee189394d9140cc99dc3f3ea6f70e0e5a770"></a>

## oci — oci / 66ca6090db8d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- oci

<a id="canonical-c33baac5cdc4a5d7aa81f3522452ce4bde259307dde50bb889fdd551672c8334"></a>

Type: `"single"`. Computed.

OCI Provider Type. OCI Provider Type.

Upstream description:

OCI Provider Type.

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

<a id="canonical-0575bd0d7df0c657cb44d4d8021a18a30b219a2daf1fbb6368dda4b0e250fa34"></a>

## Direct properties — oci / 66ca6090db8d / 3

- [not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3): complete subsection reference.

<a id="canonical-c19e97881204638a68d09aa592d3710449e80f369ba79aaab535a0846456ed3f"></a>

## Next pages — oci / 66ca6090db8d / 4

- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7cec7246e6c006cd003850a42029adcae217fde131daeb7faae47046d3e8106"></a>

## oci.not_managed — oci.not_managed / 4acf0ba4a692 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- oci.not_managed

<a id="canonical-8c91850f40f8df05f5a0cc8a51d6577427ea8570df81e5e518ef310c075b2cc4"></a>

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

<a id="canonical-d408c98db47027566bbf14d910a091b8b132f49864f835cc568c6ef9badfe92e"></a>

## Direct properties — oci.not_managed / 4acf0ba4a692 / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01): complete subsection reference.

<a id="canonical-75b3fb9cfcef33e384f9434197a86d85fa5e0adf5df55e9a9c68b4453aafbe6a"></a>

## Next pages — oci.not_managed / 4acf0ba4a692 / 4

- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69096881512706cae6d5e40e4e46c5b758d182ee0b08c6d8b595ca96af4c338f"></a>

## oci.not_managed.node_list — oci.not_managed.node_list / e6c378b0c69e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- oci.not_managed.node_list

<a id="canonical-d41c4b1387de73f7967df5c3fa9dacadeb16013f03959d430a0c03a371fa7af2"></a>

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

<a id="canonical-3e39dbd2945a4d1a234514d1297d395ee4cfac3c2d87041fa91c41ba250137b5"></a>

## Direct properties — oci.not_managed.node_list / e6c378b0c69e / 3

<a id="canonical-9d8e57c5e284019a8cf8ac9a1ecaa1aec14108f5aa56713d848753fa640d5843"></a>

<a id="canonical-3a1831d9ed946fdfdebaabe01a6378c52707edb3cc318e8fbef58d73f8a2122f"></a>

## hostname property — oci.not_managed.node_list / e6c378b0c69e / 4

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

- [interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509): complete subsection reference.

<a id="canonical-20d20e3a903b699bd607ebb965438026caf2bf9a4fd30ffeaf5466cb4e162c9d"></a>

<a id="canonical-bc108ddf1b7117a7d4508969cdaa346ccbb311de28e6147bd65a2712a54d4b8b"></a>

## public_ip property — oci.not_managed.node_list / e6c378b0c69e / 5

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

<a id="canonical-46194f1cfc5b70790924f8cdf45b980ee4201f62e5928a072c013d3f37b02751"></a>

<a id="canonical-3be0cee91f70844b29db70ab96904fd8678757a656abc15e9717c42c50e02a0f"></a>

## type property — oci.not_managed.node_list / e6c378b0c69e / 6

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

<a id="canonical-fd269150d191ec645805e6e950caf912c746216b0bcfd972d2b58433f28f2511"></a>

## Next pages — oci.not_managed.node_list / e6c378b0c69e / 7

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5588c79dc9364f6f87ed4e1af4e2b018802e3bf01eb4f5d324a614d4fdd78b5"></a>

## oci.not_managed.node_list.interface_list — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- oci.not_managed.node_list.interface_list

<a id="canonical-33e26207f2ae52cc322e14e7c4b060171665c473071871a69f137096c99af69b"></a>

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

<a id="canonical-39f98a5f6b53a133aaa232b1920b6aa1eba2ed19dd2e26db1ea7229e7775af5b"></a>

## Direct properties — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2d791435c5deee82a37a9922c541830e3e32bcc43268af37c873b26c56a629f): complete subsection reference.

<a id="canonical-a53c208df79a89d6b5f0897c354bff2a8da4aa6a49261052c89c34651694b04b"></a>

<a id="canonical-8036e7919c0537d9b76b2c875013f2ac86539a4824660905ffab05d3a61e236d"></a>

## description_spec property — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-013.md#canonical-44b96806087b03f14492e6984792856ac8735c68e306b551da582db5e23b89bc): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2f588615b1dd8dcf61fd2f325831d1b8bac4435449a2ab759e8f0d642bec3b3): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe): complete subsection reference.

<a id="canonical-95792a3b4fdd2caa11fe68fc91092dbc963a5c1758670aa0ecaf0d3501a1f3cf"></a>

<a id="canonical-d5e14a9acb58e0944b36509d1076c048ac8e4a1ca9a7c24d6888694b54bd2ce9"></a>

## is_management property — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-00b991256d713c51aa416c24f27401fd3419e2fdc6fe1ac9b1005c8097951430"></a>

<a id="canonical-5586517c94b51cff902335cd0eb2341a728f1552b980003e1057320504d25826"></a>

## is_primary property — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-25ba038c8b66588be229f46296ea72ded226adb619c9b35a88d66eb02ee9cddd"></a>

<a id="canonical-2bb04c190f98042ea9813c297342afac8da12b41966322cdf1009b8efd8c5b6c"></a>

## labels property — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 7

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

- [monitor](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1d1b57547d6e3ba4d63d54151119c8eb3a3e2571cf20e75de813c271b439d350): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-92f17f37efcefacb5efb668878862626f6eaa89b66c1ca8590c03d3b48ab80fc): complete subsection reference.

<a id="canonical-2722847996696cbd2a828a88579a53c9ea03e9848a3233e21c8845feca2a1a01"></a>

<a id="canonical-afe382d2b5c41ba0640336379b2f002f43dcffd4dd6707b5844b3469f76381be"></a>

## mtu property — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 8

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

<a id="canonical-76d765a8eec2bf338802a3e8e333c37b15013229d42cdc9c4d2b43a1c4859b90"></a>

<a id="canonical-0443d01d6ec5ada718e88691860afd0c0789b311648c6a933ee49f1a90ff1dcc"></a>

## name property — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 9

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

- [network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-fd8cbc7831843ea8c86ad76975fcd65e449c1d6f3bda43f1d6668db7221acca3): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-212eec242d9540f5cfdef2f51ff5baa4ceef1cc9c72d5a39d07394736640e1ee): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-c09b92db0b51c5c722dded8fed9292fc0b35ffc7326a53b26d034e5aba97edb6): complete subsection reference.

<a id="canonical-cb18978c3358336e94046dd99178ddd55ef3cdc5cd48ee84775597f2a435312e"></a>

<a id="canonical-8469aa016d82e6dcced3cb623a6a5a4810d0e7058d1432c1cfee65355128a1af"></a>

## priority property — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 10

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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-ac4aa57df9087406fb7de380e727a76cabcc9e95613d056d5d650319fc10f76b): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f0ae47ba6fc1974934f4b0c3dad373ce4d41b3ae44946de81e4de15e507bebb2): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2b9d8ee83708fbb8c38c964b12bf3d5f4612d3995a47df661de88eb09793f44a): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-014.md#canonical-727d54bfa26ce247ee7dd7339ae5ab083b2edb054157e9a9f4d8922a7bdd8a6f): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0c281545db0bd4002c1fa685b2fde62057b0387efa7fe77f00d6b297c53cfc96): complete subsection reference.

<a id="canonical-19521ec5030fa9949969d3ae17a8cb6843776e4820361ffeadee0c203d9ae0a4"></a>

## Next pages — oci.not_managed.node_list.interface_list / 9b1600cf25bf / 11

- [oci.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2d791435c5deee82a37a9922c541830e3e32bcc43268af37c873b26c56a629f)
- [oci.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-013.md#canonical-44b96806087b03f14492e6984792856ac8735c68e306b551da582db5e23b89bc)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [oci.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2f588615b1dd8dcf61fd2f325831d1b8bac4435449a2ab759e8f0d642bec3b3)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1d1b57547d6e3ba4d63d54151119c8eb3a3e2571cf20e75de813c271b439d350)
- [oci.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-92f17f37efcefacb5efb668878862626f6eaa89b66c1ca8590c03d3b48ab80fc)
- [oci.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-fd8cbc7831843ea8c86ad76975fcd65e449c1d6f3bda43f1d6668db7221acca3)
- [oci.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-212eec242d9540f5cfdef2f51ff5baa4ceef1cc9c72d5a39d07394736640e1ee)
- [oci.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-c09b92db0b51c5c722dded8fed9292fc0b35ffc7326a53b26d034e5aba97edb6)
- [oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-ac4aa57df9087406fb7de380e727a76cabcc9e95613d056d5d650319fc10f76b)
- [oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f0ae47ba6fc1974934f4b0c3dad373ce4d41b3ae44946de81e4de15e507bebb2)
- [oci.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-014.md#canonical-2b9d8ee83708fbb8c38c964b12bf3d5f4612d3995a47df661de88eb09793f44a)
- [oci.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-014.md#canonical-727d54bfa26ce247ee7dd7339ae5ab083b2edb054157e9a9f4d8922a7bdd8a6f)
- [oci.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-014.md#canonical-0c281545db0bd4002c1fa685b2fde62057b0387efa7fe77f00d6b297c53cfc96)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d2d791435c5deee82a37a9922c541830e3e32bcc43268af37c873b26c56a629f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71231e9e3c1a1e94944213b5994a2872ccfa376ac7f4e74c2f5c83f13e0a7680"></a>

## oci.not_managed.node_list.interface_list.bond_interface — oci.not_managed.node_list.interface_list.bond_interface / cd687acc4fd3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.bond_interface

<a id="canonical-7cb382e5b7059eb0777f259d9231a4bbef8e04ec9ace6f90f4c2b790adfe2766"></a>

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

<a id="canonical-b58dcd6f6625cb1e28592423c3f9ee00e75ff8c46560ba4f6d5915b086c0adec"></a>

## Direct properties — oci.not_managed.node_list.interface_list.bond_interface / cd687acc4fd3 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-013.md#canonical-17bbaa580c3882048de8a57335a815e3fc28e78a8465733af0a21a5d864a6b5b): complete subsection reference.

<a id="canonical-50a594bf0b486c2c1992eaaba769d92eb23019f64a73c093b31088d9f8e32e61"></a>

<a id="canonical-b12d7754c635ab46d38c8134b4f80b33945c5a26e9958ba11947500d1619044b"></a>

## devices property — oci.not_managed.node_list.interface_list.bond_interface / cd687acc4fd3 / 4

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

- [lacp](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4dd1c31c2910df0de439eba9e56fe9aecc538da494d26fe05ae29084a12d3bd): complete subsection reference.

<a id="canonical-943a03abcc31af17e970ad4a45c899c8dcbc13a3237a8ff0b18bd5aa73285d3e"></a>

<a id="canonical-bfc558995cbaa9c6146ac8602db8ab254158fbb4209a202c0b513fa85b284f78"></a>

## link_polling_interval property — oci.not_managed.node_list.interface_list.bond_interface / cd687acc4fd3 / 5

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

<a id="canonical-18e1f105e37476f2e0c535aa27bbf9b7c432e8356237869096ef0d89e9161930"></a>

<a id="canonical-ca5b9614c666edd04affba1abb02e610e7359ed1d622e4734801ef53981b955a"></a>

## link_up_delay property — oci.not_managed.node_list.interface_list.bond_interface / cd687acc4fd3 / 6

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

<a id="canonical-0d9ed64c40fc81b8d7f278ab6663012e56ac4ab1e0eccf03a217cb18a34121cc"></a>

<a id="canonical-5ef0af8097581feb1a256e1d27079f57d248fc8b30fe4bfe8328ec8b050c91bb"></a>

## name property — oci.not_managed.node_list.interface_list.bond_interface / cd687acc4fd3 / 7

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

<a id="canonical-d40580d63ee769e6081066b953f0c79922df03be38745b80c9dae526002cdeb5"></a>

## Next pages — oci.not_managed.node_list.interface_list.bond_interface / cd687acc4fd3 / 8

- [oci.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-013.md#canonical-17bbaa580c3882048de8a57335a815e3fc28e78a8465733af0a21a5d864a6b5b)
- [oci.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4dd1c31c2910df0de439eba9e56fe9aecc538da494d26fe05ae29084a12d3bd)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-17bbaa580c3882048de8a57335a815e3fc28e78a8465733af0a21a5d864a6b5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fcb41f494d556313b01d52588acfb4d4de7208ed6c84d50d3f040a955e5fbf8"></a>

## oci.not_managed.node_list.interface_list.bond_interface.active_backup — oci.not_managed.node_list.interface_list.bond_interface.active_backup / 4d05004c6eab / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2d791435c5deee82a37a9922c541830e3e32bcc43268af37c873b26c56a629f)
- oci.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-6559f439fcd49d305bf23195942f1e773d3885ea929f206a2ddbfd9d248ccd4c"></a>

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

<a id="canonical-d707803e4e529ac87a160ecabaedadb2dd636892a4245c5640e224e2f4b7260f"></a>

## Direct properties — oci.not_managed.node_list.interface_list.bond_interface.active_backup / 4d05004c6eab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3a1ed72f4e2a51f70e3e38dc07d8878f790aeab8870368aef2faf6497487893"></a>

## Next pages — oci.not_managed.node_list.interface_list.bond_interface.active_backup / 4d05004c6eab / 4

- [oci.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2d791435c5deee82a37a9922c541830e3e32bcc43268af37c873b26c56a629f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a4dd1c31c2910df0de439eba9e56fe9aecc538da494d26fe05ae29084a12d3bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-506bd260e8a788411623e1d3ec8b8a3ab9e4ec7c8c4cde13ff8dae15d6cc6f1e"></a>

## oci.not_managed.node_list.interface_list.bond_interface.lacp — oci.not_managed.node_list.interface_list.bond_interface.lacp / 1782c12ca75c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2d791435c5deee82a37a9922c541830e3e32bcc43268af37c873b26c56a629f)
- oci.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-57d24fce8d5afc4a773fecdb686980fdbb9b219051971dd1a736e6223d16a465"></a>

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

<a id="canonical-1479b03b62b7aec73843b19258d60b3e22066856a1df1249a6c600227119044f"></a>

## Direct properties — oci.not_managed.node_list.interface_list.bond_interface.lacp / 1782c12ca75c / 3

<a id="canonical-ef184e11048ceb8cc808e93325c1f64e5994079fbd004be204053a0669d5d4e8"></a>

<a id="canonical-d2f05982b52c455342434c953136e4bf55c46e8f20752951b4215c7b920df70b"></a>

## rate property — oci.not_managed.node_list.interface_list.bond_interface.lacp / 1782c12ca75c / 4

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

<a id="canonical-23ba9d599e983560f4b22767f1cb73f33e1672e2f57c3d4591a9002e65bc0c7d"></a>

## Next pages — oci.not_managed.node_list.interface_list.bond_interface.lacp / 1782c12ca75c / 5

- [oci.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d2d791435c5deee82a37a9922c541830e3e32bcc43268af37c873b26c56a629f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-44b96806087b03f14492e6984792856ac8735c68e306b551da582db5e23b89bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54c52a6986b60232781579099123a8c5c2aabbaca166ecf3b03e29a862ce61b1"></a>

## oci.not_managed.node_list.interface_list.dhcp_client — oci.not_managed.node_list.interface_list.dhcp_client / aa0dab00ad83 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-41b7c1cbc80359babaff8eee9a366a902a7d3d8d12eb6674c1f957183625be84"></a>

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

<a id="canonical-c7dd92433ba0a18f212019f77782c4ec48806e92a09ad77ecebd2746f962dac9"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_client / aa0dab00ad83 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01b8f5957cb9fafcbcacc76c0824982c930da91ead1f1c5acbcfea8ef390470c"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_client / aa0dab00ad83 / 4

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66c120905898f409de5a3f4bffd6ced542c2601a80e196d7996212b2ef3f132f"></a>

## oci.not_managed.node_list.interface_list.dhcp_server — oci.not_managed.node_list.interface_list.dhcp_server / 46d694bc9bd8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-30aaa6e8b35e100e1919b066f5f71cf7a2dca9ff25fc55ce939f9914f83630d1"></a>

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

<a id="canonical-80a9d6954e0dbeacaca3af726a2dfb0a798403da23d3193accb4ee8a0c4d8ea0"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server / 46d694bc9bd8 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-013.md#canonical-426514efccd20f2be78d6f132ae4d75c88f5910ac8b7c19679433f38cdd26202): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-013.md#canonical-9895ff76b5eb7d1393f2fcecd37a76c0858b37cc208cea5cd47c10cbef810e28): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8): complete subsection reference.

<a id="canonical-a14ceed3fb0297732b67b172b804412b5c2afded6347b8ab6964509ac4c88501"></a>

<a id="canonical-34e3738ab9b1f0224926fe46e2f0e329b544ce480e5451aa7393ccc003efb249"></a>

## dhcp_option82_tag property — oci.not_managed.node_list.interface_list.dhcp_server / 46d694bc9bd8 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-25c5881abf56034cc5bafd38901d9be37b483aa9118bf5bb8d21da14587911b3"></a>

<a id="canonical-e1653c0e9c6a692aee93cb6b7acd4bddb3fc7aaf9c813298bc5b3face40f8267"></a>

## fixed_ip_map property — oci.not_managed.node_list.interface_list.dhcp_server / 46d694bc9bd8 / 5

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-013.md#canonical-5726bc73544e0ed627a446bb9d7ddb3209bd939cc4b4dae46535845c07dd85fe): complete subsection reference.

<a id="canonical-a40d5d3bb519d7bd57b6bd6364d02b4ff9a443833fbd84d5e05fc48b0afd67bb"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server / 46d694bc9bd8 / 6

- [oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-013.md#canonical-426514efccd20f2be78d6f132ae4d75c88f5910ac8b7c19679433f38cdd26202)
- [oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-013.md#canonical-9895ff76b5eb7d1393f2fcecd37a76c0858b37cc208cea5cd47c10cbef810e28)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- [oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-013.md#canonical-5726bc73544e0ed627a446bb9d7ddb3209bd939cc4b4dae46535845c07dd85fe)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-426514efccd20f2be78d6f132ae4d75c88f5910ac8b7c19679433f38cdd26202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0d04ebf8047cd8d36bcef2473497a522cfbec4661ddaf4c515bc61db4fca387"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 97bed27b4ba3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-d1caf486e5c486d1abffe75b99e9fca3e11134727800ab3bf19262975721c6a5"></a>

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

<a id="canonical-de38129a583eccfb51a6ea5bdf9e4aef5dbaab68d1de561a752ef22ee4cd36d0"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 97bed27b4ba3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f8bd8cc88cfd03bbb336a7b8311636d12dca304ed4f1765e34547ce79074c507"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 97bed27b4ba3 / 4

- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9895ff76b5eb7d1393f2fcecd37a76c0858b37cc208cea5cd47c10cbef810e28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-387d32577426e0dd7157024a6d0a9175bff54c18283f31d93519a75861426d14"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / c24ed8cf2a62 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-6abb61648263c3f9b815603613274092d161e1c5d10bae4c34445db6df72b3ae"></a>

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

<a id="canonical-904bab9574690df47764d7f9e90718b9070b11e97347b37fde21c1e6841abd44"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / c24ed8cf2a62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6430c3b3f0757f386e220ba446bda429dc998707844aa62660fc70a6e1ee9622"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / c24ed8cf2a62 / 4

- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6b0c0474b477aa33b5d1c32922afefee6047f65e5479ffeff8857876b75ffd9"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 13850cef264b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-5cc8c2288faf2b29db29c59be19e47462183cfbf228ce873ed96671d114919f6"></a>

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

<a id="canonical-0415e01e535bd1f59f029976bba8970e81792a8bfddd376c75d3b4be4ccfb738"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 13850cef264b / 3

<a id="canonical-f01c40bb9226e190a86e6d0f57a41925468ec575243fe131d89b56255055b404"></a>

<a id="canonical-15f73687d5e5195e568bfc3eb757ae59702828f95959520c5d6f354fcdcd7d92"></a>

## dgw_address property — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 13850cef264b / 4

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

<a id="canonical-de0a2e5063770746a1850ca10e4153fbf5bb52dd8040ae345ecf543a44ddd883"></a>

<a id="canonical-7b0341ab0aa3a236fa8f98427d43562c6f411f047bb6cf2469fa30139bc26508"></a>

## dns_address property — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 13850cef264b / 5

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

- [first_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-39029a3cfbf6e681627241e986c7e5bd645fb2e71c7afc6bf2f188da5334b96b): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-7ad5a58095eacd976e43545b76a64017f71123eea314a5be3323a604daf2e27a): complete subsection reference.

<a id="canonical-fd06df2a0bbdc2738d0369594bb976e018a7624eb7a19a89567303f3cfb8db24"></a>

<a id="canonical-da9ac9cda8f2f169bf904f217a53824348a87de94f0a8870bf134f7b51f712a6"></a>

## network_prefix property — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 13850cef264b / 6

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

<a id="canonical-156eb99e72d0d5943a60a0846953fd128e4646bcdf1323f32e2e9730e5afdb35"></a>

<a id="canonical-fd15151ed784d2cd1624a8977080a233fa8cb0aa60fc3a4b9d3cbfad6e1fc5d1"></a>

## pool_settings property — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 13850cef264b / 7

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

- [pools](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d6e6145c73bd2e55d973a07c94e6b006bf0c0375b385bc7e3b25dc7069115635): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-013.md#canonical-8c443c179590bb1aa6026adf94347fd4e0ab7967c46bfb59a030423c7c82a7c0): complete subsection reference.

<a id="canonical-41e557784d8b89404c7b4893c5793a2a64557778f3e2dea787264e92696d457d"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 13850cef264b / 8

- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-39029a3cfbf6e681627241e986c7e5bd645fb2e71c7afc6bf2f188da5334b96b)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-7ad5a58095eacd976e43545b76a64017f71123eea314a5be3323a604daf2e27a)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-013.md#canonical-d6e6145c73bd2e55d973a07c94e6b006bf0c0375b385bc7e3b25dc7069115635)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-013.md#canonical-8c443c179590bb1aa6026adf94347fd4e0ab7967c46bfb59a030423c7c82a7c0)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-39029a3cfbf6e681627241e986c7e5bd645fb2e71c7afc6bf2f188da5334b96b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73bd46751631b1e6db7f31c9c656aa7116890f5506dae8d9c51776d0a70c0c2d"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 15047608f70b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-d30f6d13ae943fb6d63b760a095baf538d2317c05ce02101e347bd7721d6f29d"></a>

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

<a id="canonical-5be95ec24d379a6d0c21c63660d864710f31526bdc154d3c062ed16d3f94dc32"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 15047608f70b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7cc4a51a2b2ee44ec7b02bd5207257ce5e71b1c99625bb747f5d65907fe0bd4"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 15047608f70b / 4

- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7ad5a58095eacd976e43545b76a64017f71123eea314a5be3323a604daf2e27a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6869161220d10667cc38e83c4cacd97b963d7bcd8c4e2344423185f410780021"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / ca85809486e1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-c74a5f6fe4e15e40e8be93b277204180a98e15c58ab20313acd8306ba73cef79"></a>

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

<a id="canonical-0f5ca4f7a60ffea162af0703e403dafbab402a854016c9d8094dc5be97029d67"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / ca85809486e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-48f3a27356d533cd316ae75778fd75f7a587937348fde7ba8431508d7d609575"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / ca85809486e1 / 4

- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d6e6145c73bd2e55d973a07c94e6b006bf0c0375b385bc7e3b25dc7069115635"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58b9d020eb13a6c34c24c2ddc3c91aa5454e5c996b463b0fc34ba58c1278355a"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / fb3b851d0a85 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-00c259723119acb5e623d3e0d2dc7e8ecc3e1de71990f0d8b9dfa8b1cb30effa"></a>

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

<a id="canonical-3808710ccd72dedcbeea58ee29f0df77b40fcf9ef0c27706094e3431f67aac8c"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / fb3b851d0a85 / 3

<a id="canonical-cac01955a08f057597890c6905d98560981214bed98aa948dfa91d60f9c7cced"></a>

<a id="canonical-edd093298e78e08432a685a625054440767dc732c0a6e42b9e1042f1c586dfd8"></a>

## end_ip property — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / fb3b851d0a85 / 4

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

<a id="canonical-da3b0d70a3f38544eeb35cbea90e5d5d7b9d5532a465942e108df860c81610fa"></a>

<a id="canonical-f0d3ee18a1e970b363db34faef243dca2136070ee92c6d61833c7dd6236b4ad8"></a>

## exclude property — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / fb3b851d0a85 / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-55a3cf2e5a84bf98c84c4cc3fddfdbbd28aeb16baeed713195880c3b3ce73b8c"></a>

<a id="canonical-644f52afd213c4b7e07e058a43b0c6234656af72ccf0b40edd72884939fdde23"></a>

## start_ip property — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / fb3b851d0a85 / 6

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

<a id="canonical-d544c9196d5aab890b788aec4a3b93c580913cea1743c60ca2f8c9903d1a51cc"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / fb3b851d0a85 / 7

- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8c443c179590bb1aa6026adf94347fd4e0ab7967c46bfb59a030423c7c82a7c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2b56300341ddcda2b5d44ee9a2b90f055045d08a14cfceb1b7131610baaac0c"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / cdd4caa12ee5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0ba09a1ddecbe0f92a3be84248653b9281aecb5939a199b26c0a73645cd9f7de"></a>

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

<a id="canonical-5e5b851b3697222c790a042022b4554ad2be935a2dc69415fb52022788b1de50"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / cdd4caa12ee5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f177629117f14553bc56958cda26b799959f05fc6970bff2c7650287c4c506e0"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / cdd4caa12ee5 / 4

- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1fafec438ce72ec5c30404648eaa83dc01eda05eb2e116e70fbed3fd433d63f8)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5726bc73544e0ed627a446bb9d7ddb3209bd939cc4b4dae46535845c07dd85fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6445b6a9140a4c517c6df0150adb7c516f1514959b2155b382e5d8e3f7537bc"></a>

## oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 3411e44619ee / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-18a99a8a13baeaa5508545facadfd223c9b0ac4c477f6caaed1a6cdcee2f84ce"></a>

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

<a id="canonical-2522c32ec5544678f4b08d8aa984c748f923744607ecd51625fdeb8631d73d76"></a>

## Direct properties — oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 3411e44619ee / 3

<a id="canonical-67a5995474c276f39b7ce32045d5e1af1dc10f3c2706d0ac0ef4a800c4578613"></a>

<a id="canonical-e031d501dcd0c625d47e58c5a1a26ebbcf3f2777bbe1c148fccf9f11313d64d6"></a>

## interface_ip_map property — oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 3411e44619ee / 4

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

<a id="canonical-3aaa1dc7fd02e75993d3d857c27a797d7b7f5e3f14c34a204fb5b5b6c79d983c"></a>

## Next pages — oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 3411e44619ee / 5

- [oci.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-013.md#canonical-b208ef773490db40892db54b66633481a8ca9f19e39421ed26ec2250c4416240)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d2f588615b1dd8dcf61fd2f325831d1b8bac4435449a2ab759e8f0d642bec3b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6acc55273884256c47a10a02df29f6d4d37ce959a93399e39027d68ce799f9c"></a>

## oci.not_managed.node_list.interface_list.ethernet_interface — oci.not_managed.node_list.interface_list.ethernet_interface / a7e5488770fc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-08af573b9b7e8ae0fe0d210a41c5b7c8ee39d4e06d97e88832ad2e8fd6060cb0"></a>

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

<a id="canonical-d3b521014cd2ccb33e20be4401702fc9ff8d05f1e7c4b3a8a09c6a97cc8bff68"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ethernet_interface / a7e5488770fc / 3

<a id="canonical-5ef376184adb9971effc4ec2f3a4660d8bd94adbfb8763376a76b64258cd6660"></a>

<a id="canonical-ea77a35e81b367383ede6f005b7011270294f6c14c6e3db6f31b9c22fe65398b"></a>

## device property — oci.not_managed.node_list.interface_list.ethernet_interface / a7e5488770fc / 4

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

<a id="canonical-d916a4f564b1a9f64f423f3e13766ff0403d56ba009966985f6e385213ccb706"></a>

<a id="canonical-d7b2c42c4239a86b3e6395dde65dcfb44d81a3b7ab2a713ab38ec9a33cf0a68c"></a>

## mac property — oci.not_managed.node_list.interface_list.ethernet_interface / a7e5488770fc / 5

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

<a id="canonical-841ca0ec443abb34251c69e62c5e6bcf587ae68360a72d595b1f153c18c5318f"></a>

## Next pages — oci.not_managed.node_list.interface_list.ethernet_interface / a7e5488770fc / 6

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62c31df610fbf89bc528586e2eb30af48d9d68bd8da8a6a1bc3d31c4db44cd28"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config — oci.not_managed.node_list.interface_list.ipv6_auto_config / d0944989b9f8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-f5fb2eeb067f0678533ae783d0bb76918a451181faaff9b79e3f3f6ebde52b52"></a>

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

<a id="canonical-486e34364fcf716af5d0fe05a90d15c6384f792a62d8adc3e2c3d3cf3cf1417d"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config / d0944989b9f8 / 3

- [host](data-sources--securemesh_site_v2--reference--group-013.md#canonical-48329133e5ad9033d6421cee9617fd6aa06b3e38c2b5bd64266a5f84796f4a59): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0): complete subsection reference.

<a id="canonical-f1fe723e70bbbd50ca15c592686885812fc8681d9b8c8ef094e008fa43487879"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config / d0944989b9f8 / 4

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-013.md#canonical-48329133e5ad9033d6421cee9617fd6aa06b3e38c2b5bd64266a5f84796f4a59)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-48329133e5ad9033d6421cee9617fd6aa06b3e38c2b5bd64266a5f84796f4a59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1403d6f7f5269e0e7bbe18a1422d328d4751fb4135dbbe62d188dac5403af595"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.host — oci.not_managed.node_list.interface_list.ipv6_auto_config.host / 519f78483a09 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-bdb0a147981f186611c7223585b97a1a2b91eb3ffbc5170aef1445e16b176962"></a>

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

<a id="canonical-611407fff32f4f080243dc5992b328084ad90ff2cc5ec64f92e96064a04e85d5"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.host / 519f78483a09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b409cd2bd072d9f398978ce0e718d2df999089ea11ffc52a7adc9f5fd441914"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.host / 519f78483a09 / 4

- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70655625cf526966c8f0934cc4093bbd6aa5408ea35df807df9b02239e84bfd0"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router — oci.not_managed.node_list.interface_list.ipv6_auto_config.router / f3799915d269 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-392e4d451bac91e4dd1cc4405668df77da60ce6060f8c1bb24703e3010220209"></a>

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

<a id="canonical-9f7371ae4d8763dae8e1294f61a7f5c648cf474d0b8f3a6987ee1791377f74a5"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router / f3799915d269 / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630): complete subsection reference.

<a id="canonical-7a988ba25997fed70ef4f6dc835db6d022ba3b3206b68f68d2f9a8efbd698f59"></a>

<a id="canonical-40af07e10b3b2d736567461be80e588e50b3e6cd1dc5636d6d8bc80bc5b3745d"></a>

## network_prefix property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router / f3799915d269 / 4

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

- [stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579): complete subsection reference.

<a id="canonical-110b6cfaf542bf12fee2f84af01ae6646844c76b5a523b4e6ca5da86fbed9400"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router / f3799915d269 / 5

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ba22d8bbc3dfc75a497c921e8b2558737d795d93e9c9f98971ca213f1ad20ec"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 277c935368e4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-5f177d049f74b0f88e493b071ff3dc1b5ecc48d08d999526e0ba8cc0219a0e55"></a>

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

<a id="canonical-6b00540a38f76fc501b60efc79f887ba2d8ad06a1e90d4022f6163f13f10421b"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 277c935368e4 / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f7ed17551d94787dbbe28bd11eadbcc8637c5aa9e3ccfed66ad4b5d7ccdac718): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-38b9f2266514d46a007c885dd0099b319130b7c3652abba3af7eca99f66383a4): complete subsection reference.

<a id="canonical-4fcd4863013adfb2aad663b7e2eefcc287122b35862dee9b17b5a85502a3aae0"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 277c935368e4 / 4

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f7ed17551d94787dbbe28bd11eadbcc8637c5aa9e3ccfed66ad4b5d7ccdac718)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-38b9f2266514d46a007c885dd0099b319130b7c3652abba3af7eca99f66383a4)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f7ed17551d94787dbbe28bd11eadbcc8637c5aa9e3ccfed66ad4b5d7ccdac718"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-838f8b31cc30d139ad5ac09656672dee48963a48a6ec3d7d48ee5d5729a157ce"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / da460b140267 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-b411149f9653f3c5bccc43ee0cce9e88b1f089913d3a3962188485e7e9b2c997"></a>

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

<a id="canonical-1f2f8d3ee88a2db5e3b57780ca77edf99838daf63501045842673bf7cf941c99"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / da460b140267 / 3

<a id="canonical-3e6031b4947395de24a37f7a5e77284f21f081fd04d09d41aba052ee1769d02e"></a>

<a id="canonical-8dd42d595ac3c68690ecdd8bc4e24e56c96b9d5a70d6247a40da18d3540a3eca"></a>

## dns_list property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / da460b140267 / 4

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

<a id="canonical-dcc8be081a10d7ec0e048666e96be986bac5946392e4b041306f14508e0e9bd6"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / da460b140267 / 5

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-38b9f2266514d46a007c885dd0099b319130b7c3652abba3af7eca99f66383a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92b076ea2f3a611087e57ff20c54a023370f4cdb2fcc8b3e996e78d5405f4b0d"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 42a2784d1a29 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-445fe78c1a296009a656bc99d132b15a4bd9544029504eee6187f213f1c6b972"></a>

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

<a id="canonical-896b78f9b22932dad08d40173c820c357a8a29829dc1876b484a100c0c246941"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 42a2784d1a29 / 3

<a id="canonical-ff64bfa0839537f3ddaa288fde241f5db43736ad29f65760c3b62cdb047002f9"></a>

<a id="canonical-4ac10636cb9eeb914771fbfe0cfe02b858c4da288f1eec96f280107dbabde9f0"></a>

## configured_address property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 42a2784d1a29 / 4

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

- [first_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a3569861f635190e2515d1ed766813e25c031e384c9a244bf21288a62d92a95e): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-30ce27332467382c27f428258bd90b0c641dfc596a7c4172da5b13e96896733b): complete subsection reference.

<a id="canonical-6d98a31c42d7e8ff1199fc95ed6a45f1df9f8c15b307ba30f6903246691c3857"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 42a2784d1a29 / 5

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a3569861f635190e2515d1ed766813e25c031e384c9a244bf21288a62d92a95e)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-30ce27332467382c27f428258bd90b0c641dfc596a7c4172da5b13e96896733b)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a3569861f635190e2515d1ed766813e25c031e384c9a244bf21288a62d92a95e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99f30cefa180e4775d29c7637e7405b959a3496983bc9cdccc865e2bcfc5422a"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / fd57d2189636 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-38b9f2266514d46a007c885dd0099b319130b7c3652abba3af7eca99f66383a4)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2dccc1e9bcc9e7ec0ea3d2a226f1a27a7b0e296a5a9cfe5a74105d29da974ad8"></a>

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

<a id="canonical-269a69e2c461da0cbcae0195cbfd21fc33ea93263b620992a8f0411648eeab89"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / fd57d2189636 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d1559169a75e78961b55b8eca1e984c9850dffbe33f55903424a2bd63eb1bce"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / fd57d2189636 / 4

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-38b9f2266514d46a007c885dd0099b319130b7c3652abba3af7eca99f66383a4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-30ce27332467382c27f428258bd90b0c641dfc596a7c4172da5b13e96896733b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2d2f58fbebd63ea52bb8cf7c4cab01230f24eaa02fbf9a2d221fa7c079b9f49"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / d2351b1fbab0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3055b1bd98221aa2bb15a92b6709c893c4c7f5e7d85314c539187eb37c1eb630)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-38b9f2266514d46a007c885dd0099b319130b7c3652abba3af7eca99f66383a4)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-366074f3bf40f21438d46ff2cca39e23f42a0ec007063aeb2ce48e75714b03cf"></a>

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

<a id="canonical-378c670eee68059332373fd0b5f6246e191bdcd0a1ca2eefa72055c603fa14a7"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / d2351b1fbab0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c1784995736c1cda92ccc8e672091983b3ff46035cdb38360a1c6079c3a2891"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / d2351b1fbab0 / 4

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-013.md#canonical-38b9f2266514d46a007c885dd0099b319130b7c3652abba3af7eca99f66383a4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9a6ca78a5550111700977ea6908261719458e3868a62671d592f3c163d991a3"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ca60222cf75 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-d190972e0847e908c07ff911a91ac7332729f83cfa2cb20a33d78521a84ef385"></a>

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

<a id="canonical-7a076cf2f8be98e785514b654f6dbc0f1239eb1e9573d56529a29c2bcd450059"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ca60222cf75 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-013.md#canonical-afee880aad27eadd6d641602921722358762c3b86af4033b3f454fc6e80a1691): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-013.md#canonical-01ff3bce9ed47fd96d260226e8cfecd608ebbc74cd8930ff9d647fface0ec3cb): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-70fbce96c11b64e90294578ce3297566e12a1ef8ffd4735de0daa209cecc590c): complete subsection reference.

<a id="canonical-2d9ed62773c42958e42647d2e5d13ddf1458da8f7522eb512d359f22bc5f4157"></a>

<a id="canonical-3ab31326016940dac14aae345be670c56f6b1d9b0b6f6fa7bf737f3e7c451941"></a>

## fixed_ip_map property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ca60222cf75 / 4

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-013.md#canonical-463faaa6adcea98d34b5abd92d3900b8b1149b68d87b4e478daf9beaa3fb05c1): complete subsection reference.

<a id="canonical-ae894969713a44d04cad20e2925e3cdac113d3b1252605c32f8c91d1273fb931"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ca60222cf75 / 5

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-013.md#canonical-afee880aad27eadd6d641602921722358762c3b86af4033b3f454fc6e80a1691)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-013.md#canonical-01ff3bce9ed47fd96d260226e8cfecd608ebbc74cd8930ff9d647fface0ec3cb)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-70fbce96c11b64e90294578ce3297566e12a1ef8ffd4735de0daa209cecc590c)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-013.md#canonical-463faaa6adcea98d34b5abd92d3900b8b1149b68d87b4e478daf9beaa3fb05c1)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-afee880aad27eadd6d641602921722358762c3b86af4033b3f454fc6e80a1691"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5315f371492891e7953a96cd15c94b1804c665789c78707fea8e3ba7ab5afee3"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 3c7fa5677f7d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-9a717ae643337c65e01c5760119375efd8d089a3f483bcc7e5b7f2d0eade28cb"></a>

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

<a id="canonical-08fda99c5c6f4663b47f29f1026086a7272de98de6c40a7fe288ce963da9d76c"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 3c7fa5677f7d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6540fa7b410701a6753b63e7f4da26ac264d0340fb4a506880e965030a1ab1eb"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 3c7fa5677f7d / 4

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-01ff3bce9ed47fd96d260226e8cfecd608ebbc74cd8930ff9d647fface0ec3cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43195b0e9af14383950ca35af3cc539a39fc52914d63ae42f8c892901d26f67e"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 0a4f35142357 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-6915fb75799f4d6311fa1f84fe757afcf89795640fee99d1cfa1b8a188c42af5"></a>

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

<a id="canonical-72ce29f5ad696ee719752e66798dd302836b7aeac0a220c6f84d401b2dea39a4"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 0a4f35142357 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a983add2e2263ca1c5517236911140168bdcec16c64cbaeb6b96a18ea901a6f"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 0a4f35142357 / 4

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-70fbce96c11b64e90294578ce3297566e12a1ef8ffd4735de0daa209cecc590c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-652f8523d9e4d07d6cc0e39c602b8a0ac6c7647e0f78bc850d66644fcf71b1a9"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 255ae23790e8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-47bb2accd1617747bc06f85a5ebd05a99f0f855a0fe21f158194b90a6e5a839f"></a>

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

<a id="canonical-57570d7d32240164f44452b0727811782f20520f8061e20324f356863c838bff"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 255ae23790e8 / 3

<a id="canonical-7ada42f2a14cdd9bc123949f6b7d8e499c02e651155303bc140435780e2e89e0"></a>

<a id="canonical-2aed51ba03ad47c55b0932902a9a8510a1f32e21f8197af7be304a30b51964bd"></a>

## network_prefix property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 255ae23790e8 / 4

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

<a id="canonical-8eb158d6ade704c37c8afce4383417a03c552f4ba467990c570fec8a60c66672"></a>

<a id="canonical-a7dc135f88deb6bf428c3ad2952bdc7aa1c0381f5bb5180d7853a77599dcacd5"></a>

## pool_settings property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 255ae23790e8 / 5

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

- [pools](data-sources--securemesh_site_v2--reference--group-013.md#canonical-869ac1be7d9f30014fb6c8d3d667ccb2430f75ec42510b66e908512da790199f): complete subsection reference.

<a id="canonical-30efc307165d324f9d379358dad8975e50d53727890775c24bebfcbb83c34535"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 255ae23790e8 / 6

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-013.md#canonical-869ac1be7d9f30014fb6c8d3d667ccb2430f75ec42510b66e908512da790199f)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-869ac1be7d9f30014fb6c8d3d667ccb2430f75ec42510b66e908512da790199f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-534499e3fab91688b5095c9d53994bf93aa715556828bb716d9ebfcf05b43bc3"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 904e872e2fea / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-70fbce96c11b64e90294578ce3297566e12a1ef8ffd4735de0daa209cecc590c)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-6d02e59ee58fd03947eabcf86824ccb8546813d6da59546ce3d980e1b2813345"></a>

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

<a id="canonical-fa1767bcad2d4c547fed29c99cebc807178bd34edaf2485d18ccebd35bd707df"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 904e872e2fea / 3

<a id="canonical-a406793151f5d8bd30450632b79cf15ce714d27ea246e4a188c289087050863b"></a>

<a id="canonical-9df868e1c88fc96553cf090c40b4ddd43fae0162a81ddfc74f5f32e15766bd1d"></a>

## end_ip property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 904e872e2fea / 4

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

<a id="canonical-d29d045afacb60123c6c7ccaf8a9bba73311ce13a2b010ef31ed24cf1e6c7e18"></a>

<a id="canonical-ad709ca3a9bde7703850b3b8a0604bacabe664627dd179e646f3f9fddfa72da2"></a>

## start_ip property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 904e872e2fea / 5

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

<a id="canonical-45c04826211a52a9f8d3a26bdecb8f35cee752bd815d306fe6d8311229accdce"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 904e872e2fea / 6

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-013.md#canonical-70fbce96c11b64e90294578ce3297566e12a1ef8ffd4735de0daa209cecc590c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-463faaa6adcea98d34b5abd92d3900b8b1149b68d87b4e478daf9beaa3fb05c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8d71b2cc995b0c9fc16913e64c51fcc582226484c3eeb1f14352a7ee4c4aaec"></a>

## oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 978e64bb967b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-013.md#canonical-f894e63e32d293025282e52726fbebcb396c5950e6a84f4159cc4f3f877c04fe)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-013.md#canonical-778d826c7703620e5df21154cfecae50e9c389e4f8c5b8ed24a22e1bb301d8a0)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-fb568b2e572c00ab3442942f36ed44e4f4cd260252bb9cd275b565bec8af70cb"></a>

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

<a id="canonical-95516f1c63eae92f9487de5e5a458df8e3d27a548458e492943b2b76fadc7a32"></a>

## Direct properties — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 978e64bb967b / 3

<a id="canonical-29f95a03db31fcdd8bcb2f2840b49fae85ff33adc315e337bedbc8b7eb9ebd72"></a>

<a id="canonical-b30c8f2aa08958b355d514575cfc083147dea89305e4aa2db86723fa2b089574"></a>

## interface_ip_map property — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 978e64bb967b / 4

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

<a id="canonical-6a6550ee42795c7513ac975fd59298f431ce782a897f1b2aa44724f8916e1d37"></a>

## Next pages — oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 978e64bb967b / 5

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-61116259c2571c8df58f71ab38fcaa38d883c7661cb744ca772039449ac31579)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1d1b57547d6e3ba4d63d54151119c8eb3a3e2571cf20e75de813c271b439d350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a953781b2d3b2dec65fed9c89181c246ed1a0f59216ef3ae73e68af38048938"></a>

## oci.not_managed.node_list.interface_list.monitor — oci.not_managed.node_list.interface_list.monitor / ff95a9f4df3d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.monitor

<a id="canonical-2301ffff0830247e32dc00c787f18a49e1967a81db1cce6ce4dcae59fe1afb7d"></a>

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

<a id="canonical-c91140fcd9bb9aad270861cf02c4c8af48260ada37382f6ad4bd383ff195c0dc"></a>

## Direct properties — oci.not_managed.node_list.interface_list.monitor / ff95a9f4df3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ecc6c5bce64477e8aeff23e47bb1e2b083847e7bd2d4571c582543d69244bad2"></a>

## Next pages — oci.not_managed.node_list.interface_list.monitor / ff95a9f4df3d / 4

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-92f17f37efcefacb5efb668878862626f6eaa89b66c1ca8590c03d3b48ab80fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e62a2099ac08d9cafdf8895989c4cf84a252a99e4abab9a7a9d32ed5b38d2431"></a>

## oci.not_managed.node_list.interface_list.monitor_disabled — oci.not_managed.node_list.interface_list.monitor_disabled / 57c3b5f6cc7c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-c393bdbad729d4a0a13e38e7e7a729dc423c14ec1e05f888286fa6cb568e8cdb"></a>

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

<a id="canonical-91592b03a65509f37634d545ffefb6bd6119b55e9fb0fa2c31feabdef91247e7"></a>

## Direct properties — oci.not_managed.node_list.interface_list.monitor_disabled / 57c3b5f6cc7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d936f6dc2a7708f0c5701ad432d289f4e587f2d016cef8f7e032a076c6fd917"></a>

## Next pages — oci.not_managed.node_list.interface_list.monitor_disabled / 57c3b5f6cc7c / 4

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-fd8cbc7831843ea8c86ad76975fcd65e449c1d6f3bda43f1d6668db7221acca3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-370cb6b7e3639023e96726de9f3f0b027a291f7d5fcf2108349528b805d69877"></a>

## oci.not_managed.node_list.interface_list.network_option — oci.not_managed.node_list.interface_list.network_option / 5584e976fd11 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.network_option

<a id="canonical-cf049d641ba9ad0d0a60e865280c0706442bcbdfc456dca22b02effc9469bd71"></a>

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

<a id="canonical-f10edd532f259c39afa9e2eb9a77032206e79339bd6128aee263db0cdd6c57b8"></a>

## Direct properties — oci.not_managed.node_list.interface_list.network_option / 5584e976fd11 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-013.md#canonical-7c161677ea0b2b0778c94fc4f94db03386fca9977564119665d1279cc18c1683): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-013.md#canonical-459a08d16e3c4c8d2819f65af97487641a9ef0015bb5098650414ef69ec99834): complete subsection reference.

<a id="canonical-d43aa681582047a50457fe9649c891cf3666f49e035918c6e405f48b3e6671b9"></a>

## Next pages — oci.not_managed.node_list.interface_list.network_option / 5584e976fd11 / 4

- [oci.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-013.md#canonical-7c161677ea0b2b0778c94fc4f94db03386fca9977564119665d1279cc18c1683)
- [oci.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-013.md#canonical-459a08d16e3c4c8d2819f65af97487641a9ef0015bb5098650414ef69ec99834)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7c161677ea0b2b0778c94fc4f94db03386fca9977564119665d1279cc18c1683"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff25bae4e2e712dba76fe340e1a5e8db54d379864055c02507ba44c34191c60a"></a>

## oci.not_managed.node_list.interface_list.network_option.site_local_inside_network — oci.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 533f1c35555c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-fd8cbc7831843ea8c86ad76975fcd65e449c1d6f3bda43f1d6668db7221acca3)
- oci.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-214f6c999c97066c4f7b20b24b590924714c76f4ced026fbbe4524a69a2b8de0"></a>

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

<a id="canonical-1544a897a9dd7ac73c2635cd7ea5b0d1dc5fb69e84682068ae540aed52a584fb"></a>

## Direct properties — oci.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 533f1c35555c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9815adfd378fb70043ea3fde873a427b03969805dd0af0d78ad2d9bc41aafa25"></a>

## Next pages — oci.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 533f1c35555c / 4

- [oci.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-fd8cbc7831843ea8c86ad76975fcd65e449c1d6f3bda43f1d6668db7221acca3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-459a08d16e3c4c8d2819f65af97487641a9ef0015bb5098650414ef69ec99834"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38794d1a6c23c13a8e16ea7dc874c99ab35ee273d0d55589520763af0225cb84"></a>

## oci.not_managed.node_list.interface_list.network_option.site_local_network — oci.not_managed.node_list.interface_list.network_option.site_local_network / 5244e311244d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [oci.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-fd8cbc7831843ea8c86ad76975fcd65e449c1d6f3bda43f1d6668db7221acca3)
- oci.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-3076454540dbb727853fb601f2971ab05b621b163dc0ac5deb7bf29e899a78d0"></a>

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

<a id="canonical-9d59559c6e7393010322da358f7a3d8bbe525e5edd258a27c00cea1306c09278"></a>

## Direct properties — oci.not_managed.node_list.interface_list.network_option.site_local_network / 5244e311244d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c016004423cae01502b9e8051a1dc1cfbe519f619d5ce89c07c8268a9846d2d1"></a>

## Next pages — oci.not_managed.node_list.interface_list.network_option.site_local_network / 5244e311244d / 4

- [oci.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-fd8cbc7831843ea8c86ad76975fcd65e449c1d6f3bda43f1d6668db7221acca3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-212eec242d9540f5cfdef2f51ff5baa4ceef1cc9c72d5a39d07394736640e1ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caa6b6a39bb0cb025d23f09be257f0029dffa9666649f68ede1f4b572a0aa3b7"></a>

## oci.not_managed.node_list.interface_list.no_ipv4_address — oci.not_managed.node_list.interface_list.no_ipv4_address / 468a930283be / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-f5a39d941770ffd381dccbf136e24154ea4558fd06b1c8e779eea7454125808e"></a>

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

<a id="canonical-25544be54bdd28ffae4ea52e09101d5f04d45e3f809b95b04f4fd54ef96ac36e"></a>

## Direct properties — oci.not_managed.node_list.interface_list.no_ipv4_address / 468a930283be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76cf542a6fee0627a852552ea4be5d8ca89e2258c2b2a9150953841ff1fcf173"></a>

## Next pages — oci.not_managed.node_list.interface_list.no_ipv4_address / 468a930283be / 4

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c09b92db0b51c5c722dded8fed9292fc0b35ffc7326a53b26d034e5aba97edb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a4ff1dad967929013a6fa250cbf3db1797f2e6e63c8b601edf19ebfdf57ced5"></a>

## oci.not_managed.node_list.interface_list.no_ipv6_address — oci.not_managed.node_list.interface_list.no_ipv6_address / 2d69e48c3e21 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-53f1b520fd8f33ad082e6f25e3a68d09a99a2dcf834b2957e78979fa8f6278d6"></a>

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

<a id="canonical-9eea7ae8d46568ed836c4491405da58e3ef2b72580b4de0d44c84e9b29752acf"></a>

## Direct properties — oci.not_managed.node_list.interface_list.no_ipv6_address / 2d69e48c3e21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-501ee7dc278c318aa9ba8f3a1848198c3dc7fcc8fa7c5664ebc7ab5ff4dece44"></a>

## Next pages — oci.not_managed.node_list.interface_list.no_ipv6_address / 2d69e48c3e21 / 4

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ac4aa57df9087406fb7de380e727a76cabcc9e95613d056d5d650319fc10f76b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8656e6208162424fd9428124c326afd04591ecbb2cbc31eedc924eaedaa4cb9"></a>

## oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / f80de12562ed / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-a4f291a55e24027cf99f6e56801d242cd7029a5ee81e30a38be30b38d9b986b5)
- [oci.not_managed](data-sources--securemesh_site_v2--reference--group-013.md#canonical-114586c14d1cf947c3e4a4b114cab214e7ebab329b62b9853aca03294e3730e3)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-479d1cd547f8ecbd6a616396083f37f5c123e09bead1e02a38c10491cad7ac01)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-8663e0f4dd0ca832a9deca63bda3a2f46ee27a36a4ac31965ae826574bec0a63"></a>

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

<a id="canonical-e766cb73ea4a917e5056b62093ab7de824b0259aff14224cd22e8cd1cd2003e8"></a>

## Direct properties — oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / f80de12562ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f01d165e1e01f8a5478b93b3c570d9546937d0b72e95105d2830107bc64cf20"></a>

## Next pages — oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / f80de12562ed / 4

- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-013.md#canonical-eb23bc88b903b96739bc82f191e51e188690f21f4984b99d444853fee9556509)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f0ae47ba6fc1974934f4b0c3dad373ce4d41b3ae44946de81e4de15e507bebb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
