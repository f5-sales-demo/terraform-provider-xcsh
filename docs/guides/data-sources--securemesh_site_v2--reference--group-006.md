---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-d286603aee703e3fbeff056f2a9ebb0c1fa96f5da4b20293cdb8fabe726c01a2"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 4da6107682b9 / 5

- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-005.md#canonical-ce0ffda866ac9c62e438ff3dea61025b0ef740f2ed6977cd47624f1491396a57)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-61f72b922da72a0f24eeb13ac009cb5e5e3bde8fc3f171c6c64216358e135b48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3fad61fde5a5aa2caf10f1f5078940cd14ce6a07acb3f7de650975e2f10d770"></a>

## baremetal.not_managed.node_list.interface_list.ethernet_interface — baremetal.not_managed.node_list.interface_list.ethernet_interface / 9007f70f2658 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-2b8d29b33a31795f3c59cf773ec7a9090f58d835b09229e06d8a7be51461eea5"></a>

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

<a id="canonical-634d82aee75a0d3804cc1d1cb6392127cd0516ee357f9bab0c790432c97fc1f1"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ethernet_interface / 9007f70f2658 / 3

<a id="canonical-7059a64ca5036c9be2b1b7cd774fcd70587d09053b76bf6105eddfc1b762a224"></a>

<a id="canonical-900de7d62ebcc4ef3f17068301a32a7331032295852101cd4a623fee539a0e81"></a>

## device property — baremetal.not_managed.node_list.interface_list.ethernet_interface / 9007f70f2658 / 4

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

<a id="canonical-a6f6dc9820d47462ec67bc1c8354cea6221cb3ea5e9696dff249e468ba68f07f"></a>

<a id="canonical-8a15f41c51e6a5601a35f07d983a53401d7eb32d129c135cdca9336454c67a36"></a>

## mac property — baremetal.not_managed.node_list.interface_list.ethernet_interface / 9007f70f2658 / 5

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

<a id="canonical-e6b9a6c706f3e867860d15f235acbe76397e00f96351215103cbd857342d2f09"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ethernet_interface / 9007f70f2658 / 6

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-484241cfd6a03770abc54217678525891b3f3d42d2239d4572237a4a57ca68d6"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config — baremetal.not_managed.node_list.interface_list.ipv6_auto_config / 4f04db452c7d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-b8dab328c71e66a0c65125dca070c98eab42a2e07b7b077d99be90408a7ddc19"></a>

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

<a id="canonical-1a30ce0b29557ee0dcb2e4721be3932624c51b5890f0c5625111fdcba36ae885"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config / 4f04db452c7d / 3

- [host](data-sources--securemesh_site_v2--reference--group-006.md#canonical-eeb4a2ef4df6d95e6ce11579ba6651bde83c86558ac127a3b73dce7a2e6e2240): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499): complete subsection reference.

<a id="canonical-fcb422b1a83cb2223a94d76a114857945747e633ed39a15daaefd4a3b984bbd9"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config / 4f04db452c7d / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-006.md#canonical-eeb4a2ef4df6d95e6ce11579ba6651bde83c86558ac127a3b73dce7a2e6e2240)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-eeb4a2ef4df6d95e6ce11579ba6651bde83c86558ac127a3b73dce7a2e6e2240"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef79aec5904645c2e2997df16e329314f827f12c78125189ce0d85beb9fb00bb"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host / 48d76ea85cfe / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-398bb7b3c39a24b992ad6f1f8ea636505b54ae3469a906c990bcea65bdaaa621"></a>

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

<a id="canonical-fdbcf8ad50f0ecad2acf7aaa2d96f4cd142d183c04892cb9e629dda8fe361804"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host / 48d76ea85cfe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f86ef6d6f7df960091eb4f2bc992bc3a785b61446f1a6ed77f45326c715e19d"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.host / 48d76ea85cfe / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eae41b01386edc640f176601514cb9f30dff97e1b4c9157fae4cd12571dd5d2a"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router / c2df51c8fac0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-8e6eacf522b1e08fb05dc98088b60e685fb3206ee09e4b261d4b5d75afe9e69c"></a>

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

<a id="canonical-b0c024495376b089b7cf57ac9d5f7053d300e6946500a73077e840969de9c8b7"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router / c2df51c8fac0 / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c): complete subsection reference.

<a id="canonical-b28285f78e81eecaa121524d4347550a815c7f97ee1da32b069f407330967d41"></a>

<a id="canonical-bac9402dc66834f4dabf749f575c3ec80e5fefd331392f34169d6c8c71db23b1"></a>

## network_prefix property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router / c2df51c8fac0 / 4

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

- [stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3): complete subsection reference.

<a id="canonical-58b549f889dcb4663f76bd775012c6a59440271cae5d351491c70e7847661883"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router / c2df51c8fac0 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dd97b7983680ab262dc023cf73a6027582048387a845373338e3ea7ee466612"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 460f8700c855 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-dff8fe39043cca119a0277452af464dbb84a6e524ecbecce919b47e093a7fab7"></a>

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

<a id="canonical-76fc9d7ffc17d387184d752dbecdc05ea9a6097526afa5dfb904784599ded607"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 460f8700c855 / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1215fe43fd9327a01b8430e7a4669b9cfb1da5eb1df5f819d4b9fe173cc778d1): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-59fafb94f9e19bf88e0debdba225298a42316bda263940478d20d2ff901e199d): complete subsection reference.

<a id="canonical-3d0f2e4de979f6b315d28f8fd649593fccbe21f7e781b0734428d630e7f6411e"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 460f8700c855 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1215fe43fd9327a01b8430e7a4669b9cfb1da5eb1df5f819d4b9fe173cc778d1)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-59fafb94f9e19bf88e0debdba225298a42316bda263940478d20d2ff901e199d)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1215fe43fd9327a01b8430e7a4669b9cfb1da5eb1df5f819d4b9fe173cc778d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba25b24a469d6b12e71af30acfe69ad520a438d08de135f2c7eeecfb284305b2"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 3bfddc60ed92 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-b532c8d004dd315075196c59ac4ee08599d6057ee78bbb1b7d65bfe8ef82799f"></a>

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

<a id="canonical-87fb3e906891b2f7949a899bf3537658a4f1eeb40515b94c3da361f926b8de36"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 3bfddc60ed92 / 3

<a id="canonical-941197b21102e60f78bf51b802f4a23f758cfbe2675b5ded9b3223cad31f7d0f"></a>

<a id="canonical-f35cf3414f5727ed645becfd926530550b24029676823f9d4ab0e2a8a2a43ef6"></a>

## dns_list property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 3bfddc60ed92 / 4

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

<a id="canonical-d2619a45a935264ebae42266d4e6cc136db75feaca853c9dd7083783a0cc43ba"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 3bfddc60ed92 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-59fafb94f9e19bf88e0debdba225298a42316bda263940478d20d2ff901e199d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb7de8d77ca49e92aa7ccde05c26f654511352126f268f55b4c8a4e16ba48b38"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 8dc05ee24769 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-97ae0060e19fbdc39f836fb9e49a8a3be26cf32be85d127f22b9843782873da7"></a>

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

<a id="canonical-fda4664cccab301382744baf53d77cc1e3efd68ae9806bda6c68dae6c0418a0b"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 8dc05ee24769 / 3

<a id="canonical-5df37682366b3f26b5bb30ded83c6b11a1f3f1243c028f35f3bbd4e8d37c2dbe"></a>

<a id="canonical-d8ed5fb05b2bc5a38aafc7a2647c4bb8acf9a340f18cefb181df01d5e1cc1451"></a>

## configured_address property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 8dc05ee24769 / 4

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

- [first_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-156619564229d3cfc8261e2b5be18721db2378db845c555e027646ed91f2b8bd): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-348672da6b5984e669ae1c1050020160b3c89fc8ba5e22e1ab0ae5f8a370b0d0): complete subsection reference.

<a id="canonical-6638a3b2cf6ad7f8fd536a5a3538c9a20134bed34c564608b73a1ddbdc5049d3"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 8dc05ee24769 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-156619564229d3cfc8261e2b5be18721db2378db845c555e027646ed91f2b8bd)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-348672da6b5984e669ae1c1050020160b3c89fc8ba5e22e1ab0ae5f8a370b0d0)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-156619564229d3cfc8261e2b5be18721db2378db845c555e027646ed91f2b8bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2f485ae64774509325f885132ceda30f2c4a32a274aeca25b12fb7e8f277718"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 2c3e57c2b484 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-59fafb94f9e19bf88e0debdba225298a42316bda263940478d20d2ff901e199d)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-abb3916b7f58865034c2f87b0257c073a01c0330f4393a28c10756b19a07dbc9"></a>

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

<a id="canonical-3aef8e543027d2a3ef40f9816cbeaa2e47f21cb19c5ed6db5be4adea564c6d26"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 2c3e57c2b484 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a8a9b834afab1b791879973ed875b932b257bce8a08bf9ce729efe30bd197b9"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 2c3e57c2b484 / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-59fafb94f9e19bf88e0debdba225298a42316bda263940478d20d2ff901e199d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-348672da6b5984e669ae1c1050020160b3c89fc8ba5e22e1ab0ae5f8a370b0d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad8ad77734c7cea9df174d08e4ec80b6ae3547675d55447bd2c1ec078aa6e6c4"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 4f9d8de0219e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9ee2aadff93c546bf2e57862c58a2098b8a3bdff6567cb728f3d8e6fe87f157c)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-59fafb94f9e19bf88e0debdba225298a42316bda263940478d20d2ff901e199d)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-be5ff78a008a3d959a9de0bb4f407185088bf305e8071722ad2ff2daeecaa9f7"></a>

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

<a id="canonical-193172762053cdbeebaf1cf289bd9954427c6e91c53ee52b0a7d64435b8b8ae3"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 4f9d8de0219e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd261bd0debc74ea7318425ca150e29ecdd6ed008eb2a8fb67a60b4bb458e126"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 4f9d8de0219e / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-59fafb94f9e19bf88e0debdba225298a42316bda263940478d20d2ff901e199d)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfa3b66cc428f98d43d71fb18f4d3362ba5891137a31fd1288fe16ebb3b195e7"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 02ca32e2bba1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-ee486adaaa197d16c924e05711d4855b14a46b7bc00630a4a2d21444c97245e4"></a>

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

<a id="canonical-f189c47aadf056a40e15cc57acef33114ac7d32135c3bad368e954c072a9d9b1"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 02ca32e2bba1 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-006.md#canonical-da4d1134abbba077951fca6b47e58ac5e3aa3e1ec462f629ae99dec59a3bd135): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-006.md#canonical-2fcf5e525627564942486d3d904ebc10bd72a1cac504e8257cda0b19ae0665a6): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-006.md#canonical-08318ddd593ce7c239c604813a69aa546bfa3a2f5fc7739314d8ffea51578b93): complete subsection reference.

<a id="canonical-1b849059a3616f58e271332e4756b2683df7fa8dee6e3c99407da8dd5f9e557d"></a>

<a id="canonical-9825edb8fe6b1f1bfa761030b34eada8867f31d65aa455776fe22d9afbda7da4"></a>

## fixed_ip_map property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 02ca32e2bba1 / 4

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9130352bdaf487819a07f64a09fdea7f7004c27111ec7baf2023e5b5a8a15235): complete subsection reference.

<a id="canonical-aaf6030fdaff385395763aa929682db71dcab7cd0fb93b24e53e7ec0be888517"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 02ca32e2bba1 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-006.md#canonical-da4d1134abbba077951fca6b47e58ac5e3aa3e1ec462f629ae99dec59a3bd135)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-006.md#canonical-2fcf5e525627564942486d3d904ebc10bd72a1cac504e8257cda0b19ae0665a6)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-006.md#canonical-08318ddd593ce7c239c604813a69aa546bfa3a2f5fc7739314d8ffea51578b93)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9130352bdaf487819a07f64a09fdea7f7004c27111ec7baf2023e5b5a8a15235)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-da4d1134abbba077951fca6b47e58ac5e3aa3e1ec462f629ae99dec59a3bd135"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-400868731261fb61e6a9b64fc94bc33515029ee6baf766f9e9b9e3cb1ee96498"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 01aa6e6151ae / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-da83cd6b43dfc4cc6b1e66b76ceda7955a348092d2b8c05a221f564e3061316c"></a>

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

<a id="canonical-4111209b4f4aa8c761ea76f3cb2603f26acfb91ee972c69da7d71f382cffac2c"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 01aa6e6151ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c61a9fc0f2844e1cc28ca73481d8ecd8fef32c2393e268393343d196c81b6b5c"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 01aa6e6151ae / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2fcf5e525627564942486d3d904ebc10bd72a1cac504e8257cda0b19ae0665a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7d154639f07e72258a166cee1dd9ce7b01757c872e19990500026ab5a89771e"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 7420ce1606fe / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-af88c96c5cb4d25b65db21cf132762f166bc5cbf2e7700b1b5267f695749e66b"></a>

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

<a id="canonical-b42b934078f0f8e95dbe124f267006c9069977966ae83c09b04aa6621bde4f21"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 7420ce1606fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db780fdd67b75959c510c7344981dd2a406cc0bfa0776f76c678a91b03ab7ec6"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 7420ce1606fe / 4

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-08318ddd593ce7c239c604813a69aa546bfa3a2f5fc7739314d8ffea51578b93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d177a73a255b38871bf36e8b0c77f986cbc6644f5a925e5ce4c75d3b77d2a68"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 169aa4110924 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-cae04ce268ad8cec32776723efeed6550785c4e7d1746f04a8db3724265eb701"></a>

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

<a id="canonical-52ac3f8b78731ea69955c3ce76fd8dea41b363140e054daadd66103a5c385485"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 169aa4110924 / 3

<a id="canonical-5d3a3496a1372662a8f835642dd60ed3aee076f18b293ab45f6ae7275eb62fa1"></a>

<a id="canonical-42331b6ef5e1cdc874e833f53367684be642617e5fac42808ab7924e1a5a3803"></a>

## network_prefix property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 169aa4110924 / 4

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

<a id="canonical-b381604cb1b296d05969059549b981ed79f9b7a8e1cb2dee2c8686254f056f90"></a>

<a id="canonical-ca81be59b22987b5bb7a0fc7dc6046ee548bef4f6dd398f271c95c4739398bbe"></a>

## pool_settings property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 169aa4110924 / 5

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

- [pools](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1aaf34ab55ed17bc63c2a27c5f81bcd7dd2c9a2fb08b01a7ef56d40ab7688c68): complete subsection reference.

<a id="canonical-519e8f9350668e9d45b4b144a5c7ec4ff4bbf5d845dd29b82a966037ac3fec90"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 169aa4110924 / 6

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1aaf34ab55ed17bc63c2a27c5f81bcd7dd2c9a2fb08b01a7ef56d40ab7688c68)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1aaf34ab55ed17bc63c2a27c5f81bcd7dd2c9a2fb08b01a7ef56d40ab7688c68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-273e2058f63cae28e9ce48b2b4cdd317280644e4650e53ed0fc867dd5d57b29b"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5e1ca70630dd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-006.md#canonical-08318ddd593ce7c239c604813a69aa546bfa3a2f5fc7739314d8ffea51578b93)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-af8c5a85d81c6cffacf7d1b7851ed5572ab0f4d9f836012c58f13c869dfe1bf8"></a>

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

<a id="canonical-aca6a2659a6c90d475ffc54e6352f6ebf0b6d4d40c226a0400ad23f86a26ab39"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5e1ca70630dd / 3

<a id="canonical-f5fae61c4f06a2ff5a271054a3437fae7a9d38c6698821e4e6abf5e324216e8f"></a>

<a id="canonical-843aa77dfadfce9baa4bbb5ef35cea5738ff23931aab9559864f6e98d396c34d"></a>

## end_ip property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5e1ca70630dd / 4

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

<a id="canonical-f0b50675d2141376a189181341940ba34bf997fb8c9c60d6560c032c6f08a5aa"></a>

<a id="canonical-02b8a875824fd01d43b2a88df262e51b88291fa84d3abd31b712379aedbb66c0"></a>

## start_ip property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5e1ca70630dd / 5

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

<a id="canonical-9bfb43e25aa01983e689e00afea89b7ad32e4361dba9a623a19f3830f19ac1be"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 5e1ca70630dd / 6

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-006.md#canonical-08318ddd593ce7c239c604813a69aa546bfa3a2f5fc7739314d8ffea51578b93)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9130352bdaf487819a07f64a09fdea7f7004c27111ec7baf2023e5b5a8a15235"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1e804fa786ac7cbaec4ba7dad227010b2b9e16ee699797889c5734a889271ed"></a>

## baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 3bf6c2c70a93 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ed788020287e503f847433e529355794964fa1b695034cab9587b347dd68a6df)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-006.md#canonical-a8139d6d663d3731241dba4dba4b55fd561846868e13edcf02d0c81ce18bf499)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-6822bd7398d791c9917c6ef6a03adcb1ad79b3282256d42cc2f13d96f89660f6"></a>

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

<a id="canonical-d694aa16125b7345859b6a0b2c0355df1192830332edfe7f4c36263451cbcdc5"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 3bf6c2c70a93 / 3

<a id="canonical-2c9403df19cad6ae671cb335f36a775986f6312a0389a20be7ec3995f036f43d"></a>

<a id="canonical-ec20e9c2c84981a4ce522b6cd9fc54c8a2fa340be263344d538f8b0f7b1287f6"></a>

## interface_ip_map property — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 3bf6c2c70a93 / 4

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

<a id="canonical-afb8af68c8273a047deec25df1c93a6842fc3f485c63dc477522808711313cd1"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 3bf6c2c70a93 / 5

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-006.md#canonical-051b05bba5b6935e13acd53307b424229943cb07b58ee92b44668f0d548773d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8c727400ee3cd0bb070901fb656fb33770a1ee41a152df82756701bdb27606d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90979e364924565c8c50c5e830fd5e6c8379ba6de5ebb736de28eb899e68bd2d"></a>

## baremetal.not_managed.node_list.interface_list.monitor — baremetal.not_managed.node_list.interface_list.monitor / bb9ee6510e2b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.monitor

<a id="canonical-0d0910980d2e2d9bbebe6f8775aebb7c89583a6476cbd02705d06c1fe6fa3956"></a>

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

<a id="canonical-ce190f123f1398314a6589f71fa66a24947ee16c254394a8b8087d25d0e0c51a"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.monitor / bb9ee6510e2b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d3e84c3bc9729457e3bdd6a4bd7b09ba253911fb6e032101fe0534849f87672"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.monitor / bb9ee6510e2b / 4

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f87395695458cdddf404f9df98039d2312d8ced90699ce1529bc79d14fb44b58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dfe5fdd3bece74c929c118dabc8eaee8fb0304e60f922e316b1d9899c8539e3"></a>

## baremetal.not_managed.node_list.interface_list.monitor_disabled — baremetal.not_managed.node_list.interface_list.monitor_disabled / e0255711aa4f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-c8323a34e9a3bb98e6d35c33eb18dcbfc4bf3354080cf4325bc74c37614ac551"></a>

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

<a id="canonical-f8322d7c13d62c9dd1d66b66d9a581728ad0093d499a0a27cf4800cb2a42cc44"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.monitor_disabled / e0255711aa4f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d86164e9adee25fdea620de8fc47bef3c2ceb11e6a220a013608b4da30193c5b"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.monitor_disabled / e0255711aa4f / 4

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c68ac1bc928aa1e500d6a25aa4dec2cf35e6890702eca0c90a99f65338db1513"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6633873816155234a075017f84f04e877a7ea77e7c6572bece3a72764fc1f510"></a>

## baremetal.not_managed.node_list.interface_list.network_option — baremetal.not_managed.node_list.interface_list.network_option / d31e7230cf65 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.network_option

<a id="canonical-add21c1349f41a196638c51cd7a9da400e82919b37d6b8b9e948f5b75d94564b"></a>

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

<a id="canonical-65335898a057ed30d66c622561d7064eb9e12c4b10152e15906d173753e70f1d"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.network_option / d31e7230cf65 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-006.md#canonical-00a41ae5fc3c7a2eabd2294379d4a63cd58f4e79189b47e7941b4cdb0ca190b5): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-006.md#canonical-8b84e1706c1ce118a9ec0ef08846fe3f42fdd263aef1c319ad7edde991d421cd): complete subsection reference.

<a id="canonical-25958271fa42a7234a01678de2563e19153352fdb10c907c54ce387c4f6a6534"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.network_option / d31e7230cf65 / 4

- [baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-006.md#canonical-00a41ae5fc3c7a2eabd2294379d4a63cd58f4e79189b47e7941b4cdb0ca190b5)
- [baremetal.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-006.md#canonical-8b84e1706c1ce118a9ec0ef08846fe3f42fdd263aef1c319ad7edde991d421cd)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-00a41ae5fc3c7a2eabd2294379d4a63cd58f4e79189b47e7941b4cdb0ca190b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab7ce9680a4f1921ba7cf1508e246a2b458282a9426a043bb0032350ce841679"></a>

## baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_network — baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_ / 2b2f8f628741 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-006.md#canonical-c68ac1bc928aa1e500d6a25aa4dec2cf35e6890702eca0c90a99f65338db1513)
- baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-f15a5d817fcbf77d63e5489350aef8ffb3278a9e5ca12c53c705f443fbfef45c"></a>

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

<a id="canonical-751c2db9a81f3c5492c558efa68d3a77d5a73ba9e80df5c3f35c0ea582cf9380"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_ / 2b2f8f628741 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a902affce0f45edcc43de0126a84ace6e7c16cedfc3b581edd1351dc2b683629"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.network_option.site_local_inside_ / 2b2f8f628741 / 4

- [baremetal.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-006.md#canonical-c68ac1bc928aa1e500d6a25aa4dec2cf35e6890702eca0c90a99f65338db1513)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8b84e1706c1ce118a9ec0ef08846fe3f42fdd263aef1c319ad7edde991d421cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d850f91277d05b03a1abddd3215f6f82b6c8eec1e3dd7c0693accdbabff81ff5"></a>

## baremetal.not_managed.node_list.interface_list.network_option.site_local_network — baremetal.not_managed.node_list.interface_list.network_option.site_local_network / 66e036961cc8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-006.md#canonical-c68ac1bc928aa1e500d6a25aa4dec2cf35e6890702eca0c90a99f65338db1513)
- baremetal.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-6981f0a9bd921f9eb37f07e756a729e6f4f3271b46b191fb73312aa373711f3b"></a>

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

<a id="canonical-22c50ca25128d97b11324ad9719d08b3ed4d700ccc63f4b1c05090a58e4dfdd7"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.network_option.site_local_network / 66e036961cc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-50ac556ef83f0a8d7d7336a9364e34ab02c46700e18e65af407a20e9a845fdcd"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.network_option.site_local_network / 66e036961cc8 / 4

- [baremetal.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-006.md#canonical-c68ac1bc928aa1e500d6a25aa4dec2cf35e6890702eca0c90a99f65338db1513)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2f94ba4694c095f5bc21adda6bd3b2dfaadc2f110e54da31aa78cf0fa8fbe2dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28784e9f99e498d65c289bf2b46a9a04eaccd25c7af716c8029a015f184f8dd5"></a>

## baremetal.not_managed.node_list.interface_list.no_ipv4_address — baremetal.not_managed.node_list.interface_list.no_ipv4_address / 3548f85cb6d1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-c0c5c0bba88ec2df92f341474f4b6719a7b5d22c3104e1f7f9a94b7f823a1580"></a>

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

<a id="canonical-ffffe6fa2dd7074cdfd63d0840f36671ca899a9dc0f7af83193176fb899f4457"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.no_ipv4_address / 3548f85cb6d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f23623bbc02355f5a03bf7674ba3de820d705eccb9decbe950b52c106dd5b10d"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.no_ipv4_address / 3548f85cb6d1 / 4

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-baef686c60b77dda5a082e6d6203228acd1f48c6298b30c058457a8c6a70ee47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c39b56945f20e1a1df6d6f9d74496bc523ab4e2a348f99456814dcde5d5595f8"></a>

## baremetal.not_managed.node_list.interface_list.no_ipv6_address — baremetal.not_managed.node_list.interface_list.no_ipv6_address / 9c2a10a71336 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-e6c59993dd626d6b3ec971c72f66ee6714d2b30e08abe099f42de4ee0212b841"></a>

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

<a id="canonical-34744d9ad377e008a9b7e7bf5ac1b51a8afb43db3dc326898eb5b015e9be4450"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.no_ipv6_address / 9c2a10a71336 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea3dfe6b945d0e8e04d45844c25432c62418a344e4a6e7651acd33a0252f30d9"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.no_ipv6_address / 9c2a10a71336 / 4

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-3cc028162db95746a83494833969f3c58278f1956efbe83b26d77a3fac296a3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb6d68f0356b6487c643340065d4a31537202b6f2b26183b5111194c2b343ddf"></a>

## baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 25ef45e56300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-7c4f73a6fd0a333c36ab94ef113171464c35ab24bfffc8c8de585f0e5464f4f9"></a>

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

<a id="canonical-e9957cb6b0c7866823a8157fe159791c4f0dd46fb6e4c01ecce0905c19f79dd2"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 25ef45e56300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c6228ba33c49cd1553d928a41d76606c1eb6881dbeaca0bbe0d99c001349a4c5"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 25ef45e56300 / 4

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0453e7b7a86c465d1b6da7414a707be51f6cf4dcdf2c5b2ae73fb337df206d85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bf4c171b2368e91ae2aa5946ec8d42ecc3a1b03c1fb9222ed90a4ff129696f3"></a>

## baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 38559b2ee01f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-88a8d82f5483eebf81d43f2869b561c7b4bf01b9920ad41a8fd291195e954c60"></a>

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

<a id="canonical-5f506b3e93f3f496cefe51fa824e01e8966a4b736318fd6e0106388c41cbd1a4"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 38559b2ee01f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15c513f618e70c0d3a18ffcb36d8cc47b2d21373e46763d9b1875bf3af5d841a"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 38559b2ee01f / 4

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e84246afaea91dab2d2850f9e0733a7c1bb44d6ffe93649b8a60f5847e441a7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23c05e2fa093f8fd9ddc4bb159dcc86aa849c1dea54c9bdff2213f8d913d7766"></a>

## baremetal.not_managed.node_list.interface_list.static_ip — baremetal.not_managed.node_list.interface_list.static_ip / eb615455bb1a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.static_ip

<a id="canonical-58d1f78cc313b5d0b89f26fdc0bcfc0118aab90bd4a9e4cc7ad97efaddc154f1"></a>

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

<a id="canonical-cb4da31621992beb0fd37e690dd70d32febf9e7fb40b62becebce06a637c0168"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.static_ip / eb615455bb1a / 3

<a id="canonical-735f59bbf07ad66ac61c7a64282e94dc7b00e1df019fefaea6cc2aeb156970b2"></a>

<a id="canonical-cf60e99beabe310419eea1247aadc0c58a28e0f72392293b092b1cbfd23f591d"></a>

## default_gw property — baremetal.not_managed.node_list.interface_list.static_ip / eb615455bb1a / 4

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

<a id="canonical-b0d8191fb186ca3c979a7668d70c89279f447b9cd3c453d1326590cf612dc188"></a>

<a id="canonical-73ed181fc962c70a7ec722c54343a8d6aecd054adb57111cc73f0d2adffc4878"></a>

## dns_server property — baremetal.not_managed.node_list.interface_list.static_ip / eb615455bb1a / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-e793cba7f9f48031eae0c7b7342724cae528ab94efb7b7b53cacc616788fedb5"></a>

<a id="canonical-d56a0b7193f252f18e4addb4ed6cbcfd74f1fa70edc3dcab674f16f8d9d857e2"></a>

## ip_address property — baremetal.not_managed.node_list.interface_list.static_ip / eb615455bb1a / 6

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

<a id="canonical-9e024d5a7c991d9ba62d72c4a5aeadc5ec188e9d646d9dc4c64f5831f4b7e4bc"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.static_ip / eb615455bb1a / 7

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bfcc2c8c2f7e417eb1ffe055d119d7d9049185e0963836c6f75023b594574370"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-170f33d0571faa687e6b60f71986158864f1ed6f918c55ab8521ea12d65f01c1"></a>

## baremetal.not_managed.node_list.interface_list.static_ipv6_address — baremetal.not_managed.node_list.interface_list.static_ipv6_address / b29b4e4fece3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-bea743040a9397e6272aa149420a22b0edf89ba3e0b891e54e67b736d2052ee5"></a>

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

<a id="canonical-43d829b31a9973912f45e1ea343c47b2b977a7074dfdc1988ad796feadeff19c"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.static_ipv6_address / b29b4e4fece3 / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bf440f67e237b34fcda4a45b106f694bbf0cff14a21f04e8a490eab96f7aa3a2): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-006.md#canonical-6bdb21ff347643d45836fa6c076ea7ada7d057e13807f830a6d3fe45c66d5468): complete subsection reference.

<a id="canonical-8c61e2a8fb8e8ea5f688dcd7a1db3f2b77356329773765328eb99765eced3b1a"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.static_ipv6_address / b29b4e4fece3 / 4

- [baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bf440f67e237b34fcda4a45b106f694bbf0cff14a21f04e8a490eab96f7aa3a2)
- [baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-006.md#canonical-6bdb21ff347643d45836fa6c076ea7ada7d057e13807f830a6d3fe45c66d5468)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bf440f67e237b34fcda4a45b106f694bbf0cff14a21f04e8a490eab96f7aa3a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ced3426397074c0a4747f718f190caf5c26bb4a312ff39a46e2e42c801f9f05"></a>

## baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 62c62d909f8d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bfcc2c8c2f7e417eb1ffe055d119d7d9049185e0963836c6f75023b594574370)
- baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-ac96b9a3b27f097144dfc5d86e0a42ad648e3debd494658dadc93d86d887a7e5"></a>

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

<a id="canonical-50a3d40c375cfec5c2908e91310c6f564e840b0f3b214074bfef6a0f5e219df7"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 62c62d909f8d / 3

<a id="canonical-f303fbfcc592ae0288e47152fd4ae935028b8a2af1f6f4ffdfd93f96e8a0f64a"></a>

<a id="canonical-58871a552b619ed24c2b380102774e8cf53eb3258fec251862754dbacc6d9d07"></a>

## interface_ip_map property — baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 62c62d909f8d / 4

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

<a id="canonical-767e8fea168a3fa377bdca6e3ebc645a62313630c06172fa765a4e81fa1dc199"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 62c62d909f8d / 5

- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bfcc2c8c2f7e417eb1ffe055d119d7d9049185e0963836c6f75023b594574370)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6bdb21ff347643d45836fa6c076ea7ada7d057e13807f830a6d3fe45c66d5468"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbb2ed03f66d41d589c62b43866ab18399e09607ac386ed5222888f3d02a3b68"></a>

## baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / f1cc9abf148c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bfcc2c8c2f7e417eb1ffe055d119d7d9049185e0963836c6f75023b594574370)
- baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-5a90b74a28aebe3df0be299a5b2b99d146ca011ccbe014ed565fedd87d2f72ee"></a>

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

<a id="canonical-35109996ba21e6d650e08dcec96d13bacb8819262a612981dca8fa89de5c8dfc"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / f1cc9abf148c / 3

<a id="canonical-7b81fcb6e19b6270e8c07685cf0eabe76172447621831edff5fca1e2b6cf1f55"></a>

<a id="canonical-20146f8d919d2b57290936e09ba3b34d77e7d5b9436448c44b7ec1ab7388f331"></a>

## default_gw property — baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / f1cc9abf148c / 4

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

<a id="canonical-b293b582dd2145f5e51e8bc663a049b8a7085d85af4473eb4180101aeb2dc0a5"></a>

<a id="canonical-c45c9a48dcd023eb09f48572686f40a38b83794d3c9f599595b4a1bbb7178061"></a>

## dns_server property — baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / f1cc9abf148c / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-b957ba9f47e83451e031ed7c8418f81489ae1c1f05ca62872cd573f33ff9619d"></a>

<a id="canonical-4fa770f573317c87614bdd28ccde3f08a1dd66ec0b6380b032a29a9180da8698"></a>

## ip_address property — baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / f1cc9abf148c / 6

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

<a id="canonical-bf3aa7bfc0bb516b765c69970df38957ac4eb1d225dcdb622a373c8029aeeeb4"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / f1cc9abf148c / 7

- [baremetal.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bfcc2c8c2f7e417eb1ffe055d119d7d9049185e0963836c6f75023b594574370)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-10f842b0035011fd72f80128387f7ea1b74da2ae6f48284699b4a14e8e56f374"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d65d2eb5d753edbe16d384e5c3da5816a4e135b7f4658e56cceef612a823736"></a>

## baremetal.not_managed.node_list.interface_list.vlan_interface — baremetal.not_managed.node_list.interface_list.vlan_interface / d4fdb34e343c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5be45692217f4e1ad758a54b0a12fa9e69b659f93e60defa41c8dc59cbe511e3)
- [baremetal.not_managed](data-sources--securemesh_site_v2--reference--group-005.md#canonical-d67284f92ac52194dc409488dd405f464fac9c1ca326b699632aca2d6ae6b7df)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-e33244fca9238de8ff11d0b89a26f6f1a66d75fb099020bed03d18c3859d9558)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- baremetal.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-af736d11b56f754a1ab680cf7919925b7df353fdf31fd5566461217b42811a13"></a>

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

<a id="canonical-4acc408f459a09a8cb5a0d01238ffaab8eee521920e2efb1b88a81d39f44e559"></a>

## Direct properties — baremetal.not_managed.node_list.interface_list.vlan_interface / d4fdb34e343c / 3

<a id="canonical-b675ac6df058accb4002a3167e625df251e4f0d97738a58222e0eb4e05bf4ab3"></a>

<a id="canonical-a2737594c7b9c6c8d73f65dee06cc343f6caec6a3b4a43a4fd8a3bb9ef144a00"></a>

## device property — baremetal.not_managed.node_list.interface_list.vlan_interface / d4fdb34e343c / 4

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

<a id="canonical-74b1a48ad2cde681e092377efc328c3c61e9a9ed089b7ac53950a24ddd325678"></a>

<a id="canonical-36051b9167bf46c0216488b8e4912f0ada68ad42a4d071e82d53d0fff53591db"></a>

## vlan_id property — baremetal.not_managed.node_list.interface_list.vlan_interface / d4fdb34e343c / 5

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

<a id="canonical-db80b3e1f88d4abff3c6804a43035120833fe2e8d674c0b235c0bc4267c0a2dd"></a>

## Next pages — baremetal.not_managed.node_list.interface_list.vlan_interface / d4fdb34e343c / 6

- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-005.md#canonical-a302891d8aea496eed798ad4c194039578127c868646f9955842fb626cf835d2)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0a7cf87c580938dbb83972241837a48690174f115478c3f0483ef78da1c8bc51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96be6a5b0ba3b8df43591cc59235d31fdfd732833ed701beea2491bd7402ef4c"></a>

## block_all_services — block_all_services / cf61d7187335 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- block_all_services

<a id="canonical-0a1404077f065dba2060bf14ceb47eb56f5ef4c0eab6d26d64272d5737b7c26a"></a>

Type: `["object", {}]`. Computed.

\[OneOf: block\_all\_services, blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0a1404077f065dba2060bf14ceb47eb56f5ef4c0eab6d26d64272d5737b7c26a)
- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df4acaa208bae6deda87bed7db1e03a0a2ff5e41820c41e5a5a2886d1ba53891)

Select alternatives according to the provider validators above.

<a id="canonical-89ca04734971777ad27c90b4dc262f377a8ff8b863bcf96e70059960dffe513e"></a>

## Direct properties — block_all_services / cf61d7187335 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c5a4516dd6c3dd6e17f8345ea60474718a6f2338dfc0a40d91ead7d33103d36"></a>

## Next pages — block_all_services / cf61d7187335 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bdc07c6fdb88576d02aecaba3b780f257c3969a706dc14bee0ae1b77043879c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3fa8c4e4405ae3fce6d6299a412503946433871d20949f5e3b3c48fed6e643c"></a>

## blocked_services — blocked_services / e386850be2e7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- blocked_services

<a id="canonical-df4acaa208bae6deda87bed7db1e03a0a2ff5e41820c41e5a5a2886d1ba53891"></a>

Type: `"single"`. Computed.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b4396a976dfeca069a4bd2985fff58e5d4454a3aa9bb6060fb29e85079bf9570"></a>

## Direct properties — blocked_services / e386850be2e7 / 3

- [blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3): complete subsection reference.

<a id="canonical-7298f4fd5d846e46c2dbdf24c4d4bdea8cd3df28b2a95157e243808d46f8f2ab"></a>

## Next pages — blocked_services / e386850be2e7 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ab6ecafc22a1434396c9233782aaa34f9d69512f728b6247aae04885e5846b6"></a>

## blocked_services.blocked_service — blocked_services.blocked_service / 66101cb9f56b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bdc07c6fdb88576d02aecaba3b780f257c3969a706dc14bee0ae1b77043879c0)
- blocked_services.blocked_service

<a id="canonical-425fccec6c7fdd8a1a77c4acb3c1b2d6cb35517d04ee65c8c4a16ec897ecd75c"></a>

Type: `"list"`. Computed.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

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

<a id="canonical-74497c613346793ad7d78fbf7a4d02b3a9f6b797ab81fc6238988507abe730bf"></a>

## Direct properties — blocked_services.blocked_service / 66101cb9f56b / 3

- [dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-26d235f565c36d6cba80ebb8d0b00cd8807b580731f9e2db2f5e52446559907b): complete subsection reference.

<a id="canonical-e3f7ed8d81892d032c549f427af1e05f407b292e2411df1e8b623042bd3e7d20"></a>

<a id="canonical-00c11cd72f9219328d37175db3e3b0bf5ebeb75952db64af87ca2fdb119f03ed"></a>

## network_type property — blocked_services.blocked_service / 66101cb9f56b / 4

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

- [ssh](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bcd59fdd7f96f1e8c088bee68bc7d9224c070fbc816b945f10edce0ab0cf4572): complete subsection reference.

- [web_user_interface](data-sources--securemesh_site_v2--reference--group-006.md#canonical-6a230cbb6c54bc9d26dbb4bc11ddc6a4cb44c98a2d34f8282d20610e2879a92c): complete subsection reference.

<a id="canonical-aba7a469785db14c3cc5a51b786b296668758cff2cee6ba6274f075ed0fd5b52"></a>

## Next pages — blocked_services.blocked_service / 66101cb9f56b / 5

- [blocked_services.blocked_service.dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-26d235f565c36d6cba80ebb8d0b00cd8807b580731f9e2db2f5e52446559907b)
- [blocked_services.blocked_service.ssh](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bcd59fdd7f96f1e8c088bee68bc7d9224c070fbc816b945f10edce0ab0cf4572)
- [blocked_services.blocked_service.web_user_interface](data-sources--securemesh_site_v2--reference--group-006.md#canonical-6a230cbb6c54bc9d26dbb4bc11ddc6a4cb44c98a2d34f8282d20610e2879a92c)
- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bdc07c6fdb88576d02aecaba3b780f257c3969a706dc14bee0ae1b77043879c0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-26d235f565c36d6cba80ebb8d0b00cd8807b580731f9e2db2f5e52446559907b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38d4aa4e034e6f71d0b3f9b85978c379f70a11b8e882a4dd7f9b560ca33eeb4a"></a>

## blocked_services.blocked_service.dns — blocked_services.blocked_service.dns / 89e828bf63d0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bdc07c6fdb88576d02aecaba3b780f257c3969a706dc14bee0ae1b77043879c0)
- [blocked_services.blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3)
- blocked_services.blocked_service.dns

<a id="canonical-9610a9a91f5d051e2d7938d0e976b3f494495bb418eee1c613af5b9e74f6f265"></a>

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

<a id="canonical-73d5aa549fec19d807526bdc9d73b83e18e724c0ad98224d684841de8454a471"></a>

## Direct properties — blocked_services.blocked_service.dns / 89e828bf63d0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e2d2eb1a517b522b12330f37946c7d630fcbef24495573a196736ae10030a317"></a>

## Next pages — blocked_services.blocked_service.dns / 89e828bf63d0 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bcd59fdd7f96f1e8c088bee68bc7d9224c070fbc816b945f10edce0ab0cf4572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64b45e747bb40d497ed287dfb8b12fd7821040be23aba5a98ce089d3558bb574"></a>

## blocked_services.blocked_service.ssh — blocked_services.blocked_service.ssh / af0f9b21c86d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bdc07c6fdb88576d02aecaba3b780f257c3969a706dc14bee0ae1b77043879c0)
- [blocked_services.blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3)
- blocked_services.blocked_service.ssh

<a id="canonical-ea73dac6017e21fcd782bf276dc8c945b18b8ddf5d662a8bf6d651d9a62899a1"></a>

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

<a id="canonical-2ecc956b2ce9c17e628a262908aebf4dce30f2b3925930fcaa05e7dc058aaa29"></a>

## Direct properties — blocked_services.blocked_service.ssh / af0f9b21c86d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4643a3691665526542c9cbdf3840d1a19a59e503c84734cb39dcc1ec9c023bd6"></a>

## Next pages — blocked_services.blocked_service.ssh / af0f9b21c86d / 4

- [blocked_services.blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6a230cbb6c54bc9d26dbb4bc11ddc6a4cb44c98a2d34f8282d20610e2879a92c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-311f08482a8f26ce89e3e72add57c55c3f5eb84c382439ce6418ac453c8c4f7a"></a>

## blocked_services.blocked_service.web_user_interface — blocked_services.blocked_service.web_user_interface / e5499c9c5933 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-bdc07c6fdb88576d02aecaba3b780f257c3969a706dc14bee0ae1b77043879c0)
- [blocked_services.blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-4d84977c859d6f1f1e0e8db579a5ebc9f39d6825b4fb8b1194866c10dd5da4a6"></a>

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

<a id="canonical-9ccf5ab428d528097a7768bc8f07953cf66ec541bab7f2fe5bb641ad5fd2088b"></a>

## Direct properties — blocked_services.blocked_service.web_user_interface / e5499c9c5933 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd0d8d0710dffca4bf41030c161814e20e06d0e19503f367e1f2a34df1350001"></a>

## Next pages — blocked_services.blocked_service.web_user_interface / e5499c9c5933 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site_v2--reference--group-006.md#canonical-83d70de52694666dc3db422442ee53ba1187029db9204d21228896a8a6cf25a3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-995c42abbe7787ec8847d5b2557072711a57567ea01a23a7b1ee5a651bed489f"></a>

## custom_proxy — custom_proxy / 89544a0fbc79 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- custom_proxy

<a id="canonical-dfeea95b00213736dc35e5a4a72857df4916c93b0235e9cb57eef46ab6335f6a"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_proxy, f5\_proxy, private\_adn\] Configuration parameter for custom proxy.

Upstream description:

Custom Enterprise Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-use_for_re_tunnel_choice": "[\"disable_re_tunnel\",\"enable_re_tunnel\"]"
}
```

OneOf alternatives in this subsection:

- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-dfeea95b00213736dc35e5a4a72857df4916c93b0235e9cb57eef46ab6335f6a)
- [f5_proxy](data-sources--securemesh_site_v2--reference--group-009.md#canonical-d43bf6b388ecaecb16048fa9be3a53b79a65e1767ef5282f4cd7661313e03120)
- [private_adn](data-sources--securemesh_site_v2--reference--group-016.md#canonical-f49ea9315cf8fc4f7c44660c5c7bfddb1d3b4c507a389051f246fd47fd5576da)

Select alternatives according to the provider validators above.

<a id="canonical-ab2a047b2fac9a0accd76c8ee0375961a7a6a13ddd9920ee51d38981208579b9"></a>

## Direct properties — custom_proxy / 89544a0fbc79 / 3

- [disable_re_tunnel](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0b8fbe4b5e897e7220db74c72c0d5d64b1fc20c2ec87f1465440c7d1eb17e9ed): complete subsection reference.

- [enable_re_tunnel](data-sources--securemesh_site_v2--reference--group-006.md#canonical-6929c3a16b479853495dee289310356e065d3273c91a7d528fe66af4bbac3820): complete subsection reference.

- [password](data-sources--securemesh_site_v2--reference--group-006.md#canonical-68cc0ee6f6275e1e3503dd44f561b428c28718aaebca56ca52cb44494cd37e27): complete subsection reference.

<a id="canonical-740763ce95ab1a047dc984bf9db21254649654ab09916e9c05210de24ed1a6d0"></a>

<a id="canonical-e68445f7a41c7d0e36cdc3e22672b0de21375662dc5b3044636b82f347832d98"></a>

## proxy_ip_address property — custom_proxy / 89544a0fbc79 / 4

Type: `"string"`. Computed.

Specify the IPv4 Address of the internal Enterprise Proxy.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-a609cff7a01381b492c60c1c7fb85c9ae80e3b6b8f1fcfa1751c5a78702eb217"></a>

<a id="canonical-43774484a82fce55ef3a0e2e934fbd7721b01c28f2b3d6c20435893620e1a3a0"></a>

## proxy_port property — custom_proxy / 89544a0fbc79 / 5

Type: `"number"`. Computed.

Specify the Port of the internal Enterprise Proxy.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-dd902600083b7bc3ad894b1a270960095d9e9d55fbd5af8a19e0b24adf5de0d5"></a>

<a id="canonical-9b501129984919df36f374b4fb1289dfa4016c63a69f8b25575e39d12d22f728"></a>

## username property — custom_proxy / 89544a0fbc79 / 6

Type: `"string"`. Computed.

If the internal Enterprise Proxy is using basic authentication, specify the username. This is an
optional field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
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

<a id="canonical-dbf4e5285731b2f9d76fdc489920ee503369d943c32982430366233877f861cc"></a>

## Next pages — custom_proxy / 89544a0fbc79 / 7

- [custom_proxy.disable_re_tunnel](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0b8fbe4b5e897e7220db74c72c0d5d64b1fc20c2ec87f1465440c7d1eb17e9ed)
- [custom_proxy.enable_re_tunnel](data-sources--securemesh_site_v2--reference--group-006.md#canonical-6929c3a16b479853495dee289310356e065d3273c91a7d528fe66af4bbac3820)
- [custom_proxy.password](data-sources--securemesh_site_v2--reference--group-006.md#canonical-68cc0ee6f6275e1e3503dd44f561b428c28718aaebca56ca52cb44494cd37e27)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0b8fbe4b5e897e7220db74c72c0d5d64b1fc20c2ec87f1465440c7d1eb17e9ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18215966037afad603cd2214e219cb880957315ed999f38a1fe385add0c08886"></a>

## custom_proxy.disable_re_tunnel — custom_proxy.disable_re_tunnel / 7ad1506e486f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- custom_proxy.disable_re_tunnel

<a id="canonical-8ba442a6108a0a6ace59798ea9b4809bd1c807e68e7751db4beb94273558c3d8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable re tunnel.

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

<a id="canonical-440875772b8ff928afa694a1e8675c7e7ac3095b3c2b43b8fd6deca68a1a8d00"></a>

## Direct properties — custom_proxy.disable_re_tunnel / 7ad1506e486f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c53d4f4ee4db259b3e1604b8443aea630bdd2e5ff3ac5d52eb80c897e0e27d5d"></a>

## Next pages — custom_proxy.disable_re_tunnel / 7ad1506e486f / 4

- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6929c3a16b479853495dee289310356e065d3273c91a7d528fe66af4bbac3820"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04394bd991865f55cfc020039848ff23027f3715a5d98759b6415d5c4c04d274"></a>

## custom_proxy.enable_re_tunnel — custom_proxy.enable_re_tunnel / 1f14a068e6f3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- custom_proxy.enable_re_tunnel

<a id="canonical-1523fcbef75b8193fa66dce0bd5e9223fc2f93f7e3f4ff82d830f1ce01149a92"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable re tunnel.

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

<a id="canonical-2cff3df1eb03ab85b5c42ed54df3dba10ec2b56b75a4d389181b835295955909"></a>

## Direct properties — custom_proxy.enable_re_tunnel / 1f14a068e6f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc7dd77fb61d2b6a4bf0a23efdba8715821640abb064bbf6dd2bb25da15d3f7f"></a>

## Next pages — custom_proxy.enable_re_tunnel / 1f14a068e6f3 / 4

- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-68cc0ee6f6275e1e3503dd44f561b428c28718aaebca56ca52cb44494cd37e27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44cc72c915c5c8a0f73971cce46a4b2b29f9a8ed6e33e1ae4bf6d102ce3f6614"></a>

## custom_proxy.password — custom_proxy.password / d960ce131870 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- custom_proxy.password

<a id="canonical-a893d034ef6d4781c1bf6b84472353950a321f1be51844c5fd10fd9ebdbbdc7d"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-5f88d0398a0ed04e0f14dd168213b3069eb3a3fbe1c893ed1ed6e1f311ae2bef"></a>

## Direct properties — custom_proxy.password / d960ce131870 / 3

- [blindfold_secret_info](data-sources--securemesh_site_v2--reference--group-006.md#canonical-269390b76d742644f5ecf8789eb09bb507fff6e0d23d36615881809e184db10e): complete subsection reference.

- [clear_secret_info](data-sources--securemesh_site_v2--reference--group-006.md#canonical-09b0d9e7ef427172205fd31a67dc8b60ceb681592b133bd534a4a2b052d3fe4d): complete subsection reference.

<a id="canonical-1b50669e39c69ad5e362b32859cf71490336f4b752fd54c2aa80c69b7ea57e8f"></a>

## Next pages — custom_proxy.password / d960ce131870 / 4

- [custom_proxy.password.blindfold_secret_info](data-sources--securemesh_site_v2--reference--group-006.md#canonical-269390b76d742644f5ecf8789eb09bb507fff6e0d23d36615881809e184db10e)
- [custom_proxy.password.clear_secret_info](data-sources--securemesh_site_v2--reference--group-006.md#canonical-09b0d9e7ef427172205fd31a67dc8b60ceb681592b133bd534a4a2b052d3fe4d)
- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-269390b76d742644f5ecf8789eb09bb507fff6e0d23d36615881809e184db10e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9654b68a61f5245f5beee307bdd379f6562490c5f9a31802de17dd0a7d8c3132"></a>

## custom_proxy.password.blindfold_secret_info — custom_proxy.password.blindfold_secret_info / 6d2196b26ca3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- [custom_proxy.password](data-sources--securemesh_site_v2--reference--group-006.md#canonical-68cc0ee6f6275e1e3503dd44f561b428c28718aaebca56ca52cb44494cd37e27)
- custom_proxy.password.blindfold_secret_info

<a id="canonical-e8a9776ec95c77e97b9d7c115fff8492e8b11991e64e9103e99f4beb3de51113"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9d563e59b16738e7a75979d39f892c89e1ba21ace1fb5be47263f3fd5c84a491"></a>

## Direct properties — custom_proxy.password.blindfold_secret_info / 6d2196b26ca3 / 3

<a id="canonical-76fd89173f29962fc4b595abac3291197d21b070c2e8ca649a93facfe9ebfc13"></a>

<a id="canonical-65bdfd9d28ede23124c48a2e8dafac189a6c819381f55880ec7dd583fa092d3d"></a>

## decryption_provider property — custom_proxy.password.blindfold_secret_info / 6d2196b26ca3 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-a5a7d99e41375b344affb7381463b70f53489c7ae4c4124eff234093479dce88"></a>

<a id="canonical-5bc3aecba2c917e6c9ed58eae5649e1f15a015f1fe19cc65455e737b68407e22"></a>

## location property — custom_proxy.password.blindfold_secret_info / 6d2196b26ca3 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-08d55f956c8603b4e899947398b520d9adec0bfaef4e80073ebb9bd6ea0526b4"></a>

<a id="canonical-474aad02427272150387635f5fa61735b66c9df1aab175a6862360a1a1b24ffc"></a>

## store_provider property — custom_proxy.password.blindfold_secret_info / 6d2196b26ca3 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-02838f5fa7f913bd1660fa9cd9afca6a68228673ccdd2d7985bc18ecad75bc1f"></a>

## Next pages — custom_proxy.password.blindfold_secret_info / 6d2196b26ca3 / 7

- [custom_proxy.password](data-sources--securemesh_site_v2--reference--group-006.md#canonical-68cc0ee6f6275e1e3503dd44f561b428c28718aaebca56ca52cb44494cd37e27)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-09b0d9e7ef427172205fd31a67dc8b60ceb681592b133bd534a4a2b052d3fe4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae932c4357d206ad52bcf3bd4c215b231d3f736f23e586988490c59c37d6e5b1"></a>

## custom_proxy.password.clear_secret_info — custom_proxy.password.clear_secret_info / e0426765027c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-df3a000486b8f3fa28eaabd732100163b19fcea69dffea5684c9b1e55e56db10)
- [custom_proxy.password](data-sources--securemesh_site_v2--reference--group-006.md#canonical-68cc0ee6f6275e1e3503dd44f561b428c28718aaebca56ca52cb44494cd37e27)
- custom_proxy.password.clear_secret_info

<a id="canonical-fbd86153da8a80a332ef21de89a275f855114156d0b3ff5974b068fbfd685897"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3561ad3b7bda51b30b038983f3ae9319110d9fac1c20f02bbb98ec3086b18e24"></a>

## Direct properties — custom_proxy.password.clear_secret_info / e0426765027c / 3

<a id="canonical-7b1e63b59ee3f682c409936ca6c220ee80c74e9b75888f19f80c6362e9245b0b"></a>

<a id="canonical-c1da4b5013c05581f892ecf718482a55be1f733ea8f7b589e3b201a6a8ec4d8d"></a>

## provider_ref property — custom_proxy.password.clear_secret_info / e0426765027c / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2a66b9581604934699ad03f3f2e4acd12a3d7a42098457e2d51e010a1d1b1d9f"></a>

<a id="canonical-bc4b9efafd62f8f0ddb0fab50188aff362fc90ac2e6862574978275cc036fdfd"></a>

## url property — custom_proxy.password.clear_secret_info / e0426765027c / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-f67708958f49b6f60fc6975a4fd3e2d7a84911d38a6e94a0240674e9747b72aa"></a>

## Next pages — custom_proxy.password.clear_secret_info / e0426765027c / 6

- [custom_proxy.password](data-sources--securemesh_site_v2--reference--group-006.md#canonical-68cc0ee6f6275e1e3503dd44f561b428c28718aaebca56ca52cb44494cd37e27)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-4d3f4e8a317a63047b869979f7e14de5e4ebbc8faaf48da0a26ca6d4d07d1d08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3410986f308ac3dffdee860517af2285d9c84a54b7062c86e9bcf99d3c369c71"></a>

## custom_proxy_bypass — custom_proxy_bypass / 0dccb42d58a7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- custom_proxy_bypass

<a id="canonical-3579d01a26d8591ac2454680ed2261cc5cfe40e4d816e6ac127d64d68f7c92ca"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_proxy\_bypass, no\_proxy\_bypass; Default: no\_proxy\_bypass\] Configuration
parameter for custom proxy bypass.

Upstream description:

List of domains to bypass the proxy.

Receipt-pinned upstream constraints:

```json
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

- [custom_proxy_bypass](data-sources--securemesh_site_v2--reference--group-006.md#canonical-3579d01a26d8591ac2454680ed2261cc5cfe40e4d816e6ac127d64d68f7c92ca)
- [no_proxy_bypass](data-sources--securemesh_site_v2--reference--group-012.md#canonical-ef35083386368c146005a3b5f0d27ecf28c3613841b46dc03180dbb86d8f724a)

Select alternatives according to the provider validators above.

<a id="canonical-b0c1d45be1578ea488d06c5606c4306370cc100c409d9ba36963ffec4655fee6"></a>

## Direct properties — custom_proxy_bypass / 0dccb42d58a7 / 3

<a id="canonical-a82ab52deb76dd2eae2acb69015181c76f7aec425eb55bd40deaabf96550ee82"></a>

<a id="canonical-4a6896e3c648784ed5fa9de60af6f506c1513dbfd77b388b19b097d1c44887dd"></a>

## proxy_bypass property — custom_proxy_bypass / 0dccb42d58a7 / 4

Type: `["list", "string"]`. Computed.

Proxy Bypass. List of domains to bypass the proxy.

Upstream description:

List of domains to bypass the proxy.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-63a0c0af50e317612da35a16f92d01f787afbf60bf412e14737af5f1916a704a"></a>

## Next pages — custom_proxy_bypass / 0dccb42d58a7 / 5

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-3dfd8a56f9719e72c10f919b04b4ac0a5237eb8192c28cc4f9b3b4c05e8731c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3f8cef8f7219ce7338d2c161f1198d5697f5d404cc6e134b35e2a25baab71e3"></a>

## dc_cluster_group_sli — dc_cluster_group_sli / 2530b5529385 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- dc_cluster_group_sli

<a id="canonical-7c04b90b3afbb972a127d4a00a315c8b6bd96040dae9ca10112f18af343d033f"></a>

Type: `"single"`. Computed.

\[OneOf: dc\_cluster\_group\_sli, no\_s2s\_connectivity\_sli; Default: no\_s2s\_connectivity\_sli\]
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

OneOf alternatives in this subsection:

- [dc_cluster_group_sli](data-sources--securemesh_site_v2--reference--group-006.md#canonical-7c04b90b3afbb972a127d4a00a315c8b6bd96040dae9ca10112f18af343d033f)
- [no_s2s_connectivity_sli](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0b959449d629fd5647013247ddc8e9b6d671e75230972ed654c28f8ca2951e36)

Select alternatives according to the provider validators above.

<a id="canonical-c9bedefd023fa03e2f4f7214962fb4ad36ef511feb3a3b7412cae3fac70c9497"></a>

## Direct properties — dc_cluster_group_sli / 2530b5529385 / 3

<a id="canonical-df791bee0fe12cddf0be016636216dbbe1e658306016d76b0c1f41cd316d9880"></a>

<a id="canonical-90f94a51f5afee83048e61efca1d00af74f2ac02d75d1c3bcad649e6ee1413f3"></a>

## name property — dc_cluster_group_sli / 2530b5529385 / 4

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

<a id="canonical-736a9d9b1398a454227dd40ea19a12c3a28e16492b58fd75d836a6de40c66dd8"></a>

<a id="canonical-366865429755c86275bcfaec269693810d97681e482e27126f50f88436b9d785"></a>

## namespace property — dc_cluster_group_sli / 2530b5529385 / 5

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

<a id="canonical-aa71a4fab20e1636f0be091382639e1dfb108c3ce1514a59aea511f205454e77"></a>

<a id="canonical-743a350fbd9aeac4ab56031950e51c15a3ff7033adfcb0f453564e76fa298cd4"></a>

## tenant property — dc_cluster_group_sli / 2530b5529385 / 6

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

<a id="canonical-67a128f54b9ed58521a5e62af8c01ff40a25d70edcc80783aa845af9377ec9e9"></a>

## Next pages — dc_cluster_group_sli / 2530b5529385 / 7

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1cd8beb30390eb9dff0352e7eaba35691cbae29677af2490983826c0247f7293"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6397ba08c0bc6f5de16bd801e7e52109294b68d0525d4ef8123f1b23d55f14da"></a>

## dc_cluster_group_slo — dc_cluster_group_slo / 5b9a2367bb2e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- dc_cluster_group_slo

<a id="canonical-13a743e97c5ed8b2b198fbdc3cdce8b4127eb97010a8c0de17da498db7612b34"></a>

Type: `"single"`. Computed.

\[OneOf: dc\_cluster\_group\_slo, no\_s2s\_connectivity\_slo, site\_mesh\_group\_on\_slo; Default:
no\_s2s\_connectivity\_slo\] Type establishes a direct reference from one object(the referrer) to
another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [dc_cluster_group_slo](data-sources--securemesh_site_v2--reference--group-006.md#canonical-13a743e97c5ed8b2b198fbdc3cdce8b4127eb97010a8c0de17da498db7612b34)
- [no_s2s_connectivity_slo](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2c514fcd35556199f33b4e7e37e4bc9a58c985963ff92fe9dd767265d2cdddfa)
- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-016.md#canonical-e96691a958d856d3aeb554371c04a13896cf452a2de85b05476675125addf845)

Select alternatives according to the provider validators above.

<a id="canonical-75cfc65720da3b63ed48fe6ff7411e3e09907291e0abf5f0791f6924bbe59218"></a>

## Direct properties — dc_cluster_group_slo / 5b9a2367bb2e / 3

<a id="canonical-ebf8e66a7b726265b998c5bd354d6f3ad326c3d7095763edcbe06b169ce46489"></a>

<a id="canonical-f26892150dba73a5ca84b4d66131cfe8cd8d8e96283da87a0bd40a1e264b2b1c"></a>

## name property — dc_cluster_group_slo / 5b9a2367bb2e / 4

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

<a id="canonical-7cae828070715c08f884d2ad37c375732d0511ddb5f3a4bba4f299355fc434ca"></a>

<a id="canonical-362b37599ee5dd8b707b6fbe2c8378d22ceb271fc95cb8d893e1b9b4a56b47cd"></a>

## namespace property — dc_cluster_group_slo / 5b9a2367bb2e / 5

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

<a id="canonical-7338171aa87f175c46f744a9c2129869fe67681904b98a5fcb42c5c16df522d4"></a>

<a id="canonical-a6aa8e20803c596acde47557a33ba7a96793dbb7c8fc39ec1132d0a49ac4f7a4"></a>

## tenant property — dc_cluster_group_slo / 5b9a2367bb2e / 6

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

<a id="canonical-29993c442e9569ba9ad6dbdb3b299f28c73819222e02b24cf0b2af830665d3a9"></a>

## Next pages — dc_cluster_group_slo / 5b9a2367bb2e / 7

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-672ec91c4e0d8791a168094e3a55b37bcc49ad0407eb2b1d7e760924d8087c7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecbdf3bc3c8b323661cafac38868fd5924d740f47278dc85e43ab8b1bf34e5c1"></a>

## disable_advanced_delivery — disable_advanced_delivery / bca1b454616c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- disable_advanced_delivery

<a id="canonical-5ed14b0aade8e07a46f6a58deba4f519f80569f48103081d8280a39615a98d2e"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_advanced\_delivery, enable\_advanced\_delivery; Default:
disable\_advanced\_delivery\] Configuration parameter for disable advanced delivery.

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

OneOf alternatives in this subsection:

- [disable_advanced_delivery](data-sources--securemesh_site_v2--reference--group-006.md#canonical-5ed14b0aade8e07a46f6a58deba4f519f80569f48103081d8280a39615a98d2e)
- [enable_advanced_delivery](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3241c2eb56b4585e5d89eb5b3af5db61afece4b1a74a088981bf7eebdcfd1868)

Select alternatives according to the provider validators above.

<a id="canonical-d4efa329d948ca7c716e94fe5200f3f3ddf200d9b3c981b882b2d4fef146a4e9"></a>

## Direct properties — disable_advanced_delivery / bca1b454616c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-077605b828f1e3694cec9094552cced51dc162b4c23c738518597b61bfc1bb5a"></a>

## Next pages — disable_advanced_delivery / bca1b454616c / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-879e23ae61314061c8979fa698c980b6ebc7ce92552213fe62e4693c7cacf58a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-076392302394536a6f431cbffa5aeaaceea9cab9664e683cf6cead13507e2c60"></a>

## disable_ha — disable_ha / f0a9d820a52e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- disable_ha

<a id="canonical-1eb226b113a7cc88d8f8a555c52ec7e0aa4b17e4c6dbd632ce7771a732bc779d"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ha, enable\_ha; Default: disable\_ha\] Enable this option

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

OneOf alternatives in this subsection:

- [disable_ha](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1eb226b113a7cc88d8f8a555c52ec7e0aa4b17e4c6dbd632ce7771a732bc779d)
- [enable_ha](data-sources--securemesh_site_v2--reference--group-008.md#canonical-e5bd040dd34926397e376b81d1f9a808c2b3f5e754a2e4f1c63b03e6112c931d)

Select alternatives according to the provider validators above.

<a id="canonical-b099f4ce724bc65ed7a6e07ad88539852c9e287b639cf96d0e51cf1fbca6b8e2"></a>

## Direct properties — disable_ha / f0a9d820a52e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3d707574482f9b069bfc9017fb28a5d1faa50d8662e8b4e4d5f9232b990bd1b"></a>

## Next pages — disable_ha / f0a9d820a52e / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-806e429bdc10c108ea7fb0bc7c4a9bf6abcbfb6acc1bca8b6147ff9d43804711"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24aefeea4aae4a0ec667d2adb8bb789b9c9bc540255d23be90a640b438ae7f11"></a>

## disable_log_anonymization — disable_log_anonymization / 5781f2273658 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- disable_log_anonymization

<a id="canonical-49eb5376a4c642d6a861837486357a80f28b82d5087f55ecc5f795d087081b28"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

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

OneOf alternatives in this subsection:

- [disable_log_anonymization](data-sources--securemesh_site_v2--reference--group-006.md#canonical-49eb5376a4c642d6a861837486357a80f28b82d5087f55ecc5f795d087081b28)
- [enable_log_anonymization](data-sources--securemesh_site_v2--reference--group-008.md#canonical-fdca99a3fb88947ec15b6353aaa2ca758cf0a86ed5ce874c9b185ed4d902c4c6)

Select alternatives according to the provider validators above.

<a id="canonical-cb18d2e9f675508911557b53e555add864ed7656f2c46a8bd9c40810c905af33"></a>

## Direct properties — disable_log_anonymization / 5781f2273658 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7183da49052bdca1a686e014f942c43de093dda9a5828ed2c18944c6423f0c4c"></a>

## Next pages — disable_log_anonymization / 5781f2273658 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ec6e84c00f01885007916a5339b1bd2e894e8f79e2f84c9b5074b10559010d0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-013f19e048bbfc833e77f603eec37233ab095ba6377fc308125dcde652d80b93"></a>

## disable_management_network — disable_management_network / 442ebcb7ea46 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- disable_management_network

<a id="canonical-060f415f8416fe2d34c228132c6831911cb011cddfff633ad185c3c540c810fe"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_management\_network, enable\_management\_network; Default:
disable\_management\_network\] Configuration parameter for disable management network.

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

OneOf alternatives in this subsection:

- [disable_management_network](data-sources--securemesh_site_v2--reference--group-006.md#canonical-060f415f8416fe2d34c228132c6831911cb011cddfff633ad185c3c540c810fe)
- [enable_management_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-45a1bcf6f61e659e90fe138f01a89ee4c354be889668e831e114ef8b64bd8c26)

Select alternatives according to the provider validators above.

<a id="canonical-b7e5fb54ec08753afa9618084655b13758b8696e87861a8d547e5764b925ecbd"></a>

## Direct properties — disable_management_network / 442ebcb7ea46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6451e1c223f8c9d8523c55143f2902dd52c3026a600778e2812aa4c03f337656"></a>

## Next pages — disable_management_network / 442ebcb7ea46 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7390627391f5fd1c7117c27a9d1df6c5aa028d6e39dd04493b74000fd77f471f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19196424aec96e521f390ca79e55316e62802085057046591b26336a71b10823"></a>

## disable_url_categorization — disable_url_categorization / e3369390ef4e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- disable_url_categorization

<a id="canonical-ae8b265200cfc0e4d6778971fe3a70866a012106025c3ffd4b66a6f283fdc644"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_url\_categorization, enable\_url\_categorization; Default:
disable\_url\_categorization\] Enable this option

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

OneOf alternatives in this subsection:

- [disable_url_categorization](data-sources--securemesh_site_v2--reference--group-006.md#canonical-ae8b265200cfc0e4d6778971fe3a70866a012106025c3ffd4b66a6f283fdc644)
- [enable_url_categorization](data-sources--securemesh_site_v2--reference--group-008.md#canonical-afc9dd89d440e4dc3318e2671b1cb46b11997dc760cc640528b5bbcb13f04975)

Select alternatives according to the provider validators above.

<a id="canonical-489db10c8a22aec194702ab82f38c1056122f927984708d422e3b6216fcddcba"></a>

## Direct properties — disable_url_categorization / e3369390ef4e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af6e2a87d698c0e86f59113feb45d4e8ccc74abd8d7a3eb8750d8be3ac35a142"></a>

## Next pages — disable_url_categorization / e3369390ef4e / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f46540468d23116d84d9d4fa6eb00dcd2f255ff558cdf538fa9f77a12d14261"></a>

## dns_ntp_config — dns_ntp_config / a2e42aac1529 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- dns_ntp_config

<a id="canonical-02fb5dec366ffe16c080949a0ac9dcab6d37e348cc1a86f47a7ac37b4f84f48c"></a>

Type: `"single"`. Computed.

Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_server_choice": "[\"custom_dns\",\"f5_dns_default\"]",
  "x-ves-oneof-field-ntp_server_choice": "[\"custom_ntp\",\"f5_ntp_default\"]"
}
```

<a id="canonical-10f97852a56ef168f2bbb663d86e28878137c282fc563969c37251341793960b"></a>

## Direct properties — dns_ntp_config / a2e42aac1529 / 3

- [custom_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-d43941bfec97110e822fc2c46be630772a7a249491c68793e844ebe713cf9571): complete subsection reference.

- [custom_ntp](data-sources--securemesh_site_v2--reference--group-006.md#canonical-75a651441c367bc8c326788a7d763de51e2ea54e501420c130a8e4ac10ae35c7): complete subsection reference.

- [f5_dns_default](data-sources--securemesh_site_v2--reference--group-006.md#canonical-efa9ddfb43bcb0d503ba180ff57a5577075be291799f7abe70a4585883ffc60b): complete subsection reference.

- [f5_ntp_default](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9e61bc3e55a79a40504577060b62c22cb80f9e39026e00475f124d6d55ae1056): complete subsection reference.

<a id="canonical-0d02ba1564e436eb065badd58c52d218f0a1eab02052453886ee76f941b9c9f7"></a>

## Next pages — dns_ntp_config / a2e42aac1529 / 4

- [dns_ntp_config.custom_dns](data-sources--securemesh_site_v2--reference--group-006.md#canonical-d43941bfec97110e822fc2c46be630772a7a249491c68793e844ebe713cf9571)
- [dns_ntp_config.custom_ntp](data-sources--securemesh_site_v2--reference--group-006.md#canonical-75a651441c367bc8c326788a7d763de51e2ea54e501420c130a8e4ac10ae35c7)
- [dns_ntp_config.f5_dns_default](data-sources--securemesh_site_v2--reference--group-006.md#canonical-efa9ddfb43bcb0d503ba180ff57a5577075be291799f7abe70a4585883ffc60b)
- [dns_ntp_config.f5_ntp_default](data-sources--securemesh_site_v2--reference--group-006.md#canonical-9e61bc3e55a79a40504577060b62c22cb80f9e39026e00475f124d6d55ae1056)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d43941bfec97110e822fc2c46be630772a7a249491c68793e844ebe713cf9571"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77b7218188b0821f47bb11ab3e7d3c9dfc572543d2f84135d9616eaaf9c2d797"></a>

## dns_ntp_config.custom_dns — dns_ntp_config.custom_dns / f4d8ebf4fe98 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- dns_ntp_config.custom_dns

<a id="canonical-0d3da6fbe6f883b39c4d9f19f92e73bf195628f4107117e207c371c3a37bcb90"></a>

Type: `"single"`. Computed.

DNS Servers. DNS Servers.

Upstream description:

DNS Servers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-57019782d43bf169a182a8141eb446879d1d789a4cb79cc677cf95e1cda2e3fd"></a>

## Direct properties — dns_ntp_config.custom_dns / f4d8ebf4fe98 / 3

<a id="canonical-81a0d4b97c48e400ac77f15f4a13e42a3d3db77c4dbca30828cf190cfd0baaf7"></a>

<a id="canonical-b85fc03f96a56d058761e96edddb7cfae3f70d14bbd9d29959018eb183aa5bcd"></a>

## dns_servers property — dns_ntp_config.custom_dns / f4d8ebf4fe98 / 4

Type: `["list", "string"]`. Computed.

DNS Servers. DNS Servers.

Upstream description:

DNS Servers.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f6a497d54593fba566098f077ba0ae7887a52daeca838448c04b76d1e0371720"></a>

## Next pages — dns_ntp_config.custom_dns / f4d8ebf4fe98 / 5

- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-75a651441c367bc8c326788a7d763de51e2ea54e501420c130a8e4ac10ae35c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22c0a472cc598a327c0fe78f3a045c5138c4290972c89cf0c78e0ec9178038f3"></a>

## dns_ntp_config.custom_ntp — dns_ntp_config.custom_ntp / 6516b1b34d70 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- dns_ntp_config.custom_ntp

<a id="canonical-4bd0815d2e31f4f88e0293509a2c5e708f3aef35f3949f0e7fc1eb3db0e46bf3"></a>

Type: `"single"`. Computed.

NTP Servers. NTP Servers.

Upstream description:

NTP Servers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-52f87a517a3d4ff8bc59a2ca0554085a250594a0d012d0312437d013ed353634"></a>

## Direct properties — dns_ntp_config.custom_ntp / 6516b1b34d70 / 3

<a id="canonical-2002955c2121277503f67a1ea1a4a8106ef80b5752557f61f766f8b307e8bf16"></a>

<a id="canonical-d24128c212acee3bf564047b50791f5b99383e8bf5edc27163c74015d7a54a5d"></a>

## ntp_servers property — dns_ntp_config.custom_ntp / 6516b1b34d70 / 4

Type: `["list", "string"]`. Computed.

NTP Servers. NTP Servers.

Upstream description:

NTP Servers.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-54c9bf897aebcbad20657a87d812e468e476d6d96bca6149ebc8fd2419d0a5e2"></a>

## Next pages — dns_ntp_config.custom_ntp / 6516b1b34d70 / 5

- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-efa9ddfb43bcb0d503ba180ff57a5577075be291799f7abe70a4585883ffc60b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbb2f803577cc39fc46237b1e6f697843bda101069febfb1096b6c3a42976739"></a>

## dns_ntp_config.f5_dns_default — dns_ntp_config.f5_dns_default / 0531085c7c09 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- dns_ntp_config.f5_dns_default

<a id="canonical-90493b33117a1f90843910cea5e26bbd940679c28e529226e0d167b37f573096"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for f5 dns default.

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

<a id="canonical-e38bf5affed09f7fe72257e16350ac79a93424f07c0163ff69c5077f1c8c7ddb"></a>

## Direct properties — dns_ntp_config.f5_dns_default / 0531085c7c09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14e58a7eb57d568d7cec208ce4ed89392416478f8facca424880e3ccd9de297c"></a>

## Next pages — dns_ntp_config.f5_dns_default / 0531085c7c09 / 4

- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9e61bc3e55a79a40504577060b62c22cb80f9e39026e00475f124d6d55ae1056"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff0f8a9c1730bf02a39b0ec25d277313a9a2d497b026c40286d263413efcc646"></a>

## dns_ntp_config.f5_ntp_default — dns_ntp_config.f5_ntp_default / 5b40396039cd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- dns_ntp_config.f5_ntp_default

<a id="canonical-d0548c367633167d5b68013731a978759d699993c33efe6908c31fc412b43aad"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for f5 ntp default.

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

<a id="canonical-f0de299f5d2f152cefbba00da3eed5c9779d60e64d21a5d619512d5dd9a3ba96"></a>

## Direct properties — dns_ntp_config.f5_ntp_default / 5b40396039cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-68ff93606b59c222846b4b93996fbae354bed872bc963d3b687d8221b5dcbafc"></a>

## Next pages — dns_ntp_config.f5_ntp_default / 5b40396039cd / 4

- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-006.md#canonical-fc91c77d554a75aed311fab57b0cf31e82dbecdb9a2297a2cc41bcff7ab2c2a6)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9b9a08412aea228d084300841733c26648aeb3be1139135fddefcaaba51b1deb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd83ea3d2a8968618f137fe409000c48e26c3e8fd6a0c7a77c89bb66a4076df4"></a>

## eks_k8s — eks_k8s / 426f599ff496 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- eks_k8s

<a id="canonical-64f2e4271a2ba72a599ad4101168af34bdd9d87c274a59b89326ad85ad7e401a"></a>

Type: `"single"`. Computed.

Kubernetes Provider Type. Kubernetes Provider Type.

Upstream description:

Kubernetes Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-anti_affinity_choice": "[\"disable_anti_affinity\",\"enable_anti_affinity\"]"
}
```

<a id="canonical-81340866b2f3b8f806435b16013eb39544e2f36a6338afb179278c9f8074d9ca"></a>

## Direct properties — eks_k8s / 426f599ff496 / 3

<a id="canonical-fa6dd65cdc1ca5e9e86682dbbf2cd54bb4c573bd435b927f7a72851c576970c3"></a>
