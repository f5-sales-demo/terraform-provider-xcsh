---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-c6e5794e1a8abf4859bb8f9c96c1a53002ccdda93e3668f416d03c9b8966dbff"></a>

## interface_ip_map property — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 9e81a5cc7935 / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-7ef4c444bda3e50d028bb2dc9955c2d3c262517d06b1b4f32a624d9de4d5144c"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 9e81a5cc7935 / 5

- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7c8ec9b093d5cbfd537c76ece2c04b0ffe953bf69b5fc0512b76e9e94a59c5fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6a637582482266a0c4ca396cd87910252cf42e8920fc909694c29b17a224c7e"></a>

## openstack.not_managed.node_list.interface_list.ethernet_interface — openstack.not_managed.node_list.interface_list.ethernet_interface / 5f2d0925e628 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-cf5d77c12ce5cb08a2b72ce514727d6f33af7302a8b23f1054dfe2d36cb78a56"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2cea54b3947f25af7295007b4ca61e687d270b198a6943ef8fd9cd318efc7b89"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ethernet_interface / 5f2d0925e628 / 3

<a id="canonical-733a8ffc329b1d54e25105f59816f89c31cffe351f16e3cf008942ccc6012041"></a>

<a id="canonical-82a50100993e75bb3c08736d8d52d66b655623e24e1f69c042091f87b38cc362"></a>

## device property — openstack.not_managed.node_list.interface_list.ethernet_interface / 5f2d0925e628 / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-802b95f3e28f25dbc7e3942d4ecd1c5de1ce4af2014de1d47d7df09a4902b8b4"></a>

<a id="canonical-90727106e5298708affc833981b80ab5e1fe29a28cc69effd3614a0fe64083f1"></a>

## mac property — openstack.not_managed.node_list.interface_list.ethernet_interface / 5f2d0925e628 / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

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

<a id="canonical-6b627cedc5efe2c57e322e66fa48a32125af4b7f0526b48ba69672ddfd5368d4"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ethernet_interface / 5f2d0925e628 / 6

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ac899b67984d43d0f94d01672890a0417ee3994341dc724b7ec7302a9025a5f"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config — openstack.not_managed.node_list.interface_list.ipv6_auto_config / 01c3e85bdecd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-c31e2905790afad1beb46291a31d19cea73c4ad7a24b604f0c08b35c40f3e5e1"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-647a8ae26952367a8e2b0756c9d84762deb1441d515292fb2311a1df976da405"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config / 01c3e85bdecd / 3

- [host](resources--securemesh_site_v2--reference--group-016.md#canonical-d1c9b527b2915b8f8f7536a0b7e59d817009ed89a783ac05a2d9c09235a1d114): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181): complete subsection reference.

<a id="canonical-2dd0ae046713dc0ab9608ecd8a10a0f1821c48d3408c0e0dd64955adb6e6f073"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config / 01c3e85bdecd / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-016.md#canonical-d1c9b527b2915b8f8f7536a0b7e59d817009ed89a783ac05a2d9c09235a1d114)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d1c9b527b2915b8f8f7536a0b7e59d817009ed89a783ac05a2d9c09235a1d114"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bd1c808b8eef71d1f66a10086935c25cb8b3df16be9cec0b0e042d31002560d"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.host — openstack.not_managed.node_list.interface_list.ipv6_auto_config.host / 6017fa066357 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-e4309f0337475d71138bdf91cc3969e9bec12ba706e821eb088160e7f3c47558"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
host = {}
```

<a id="canonical-9d09eb6878599a10eb0cf99ec9f8a78c90d3228f66a7f6b1020ccae80190e814"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.host / 6017fa066357 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a251749ee6735c0422ec9e88c18ce1cfd14ba6dd4682e8438c7cb438108e0d3"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.host / 6017fa066357 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb458919f40e7454ff481e7298f619ab442044f801b36bc8ceff559988588394"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / ca3d05157795 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-d30ef7e22c5a6c5d60f589644cc1911077988c5e417b8968e0cf1a35dd427f91"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-421153878adaeb87d6665b12561b294d27b2022cdb557a8ff8adf4b5cf696ccc"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / ca3d05157795 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee): complete subsection reference.

<a id="canonical-90e4a27faea6c7ee73617412346f2e3e51c422f0b04c5c2564a74d74275a2f53"></a>

<a id="canonical-c8c97446c3c690290ac979419e9e3f4690617f2ba796aeeed2a6fdbbb5fb5668"></a>

## network_prefix property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / ca3d05157795 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

- [stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1): complete subsection reference.

<a id="canonical-f50e005fb5cd22e05ab2229af93a39124fa31efdb6a8cb42d44a3d46113f42f0"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router / ca3d05157795 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95387253ec4ba1c343c8e1df2346bd14bc91b933eae06461550935563fa99e47"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f4bf263f4324 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-7ed60ccc595b383910a984ecc5fce01a3e3db5cd08bb135607b274bb34120664"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-6b0712a7ad534cc2ad05028cab1cb6b4fd856032bd1fbda0f34c264a39e89d77"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f4bf263f4324 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-016.md#canonical-3fada755790aeb178c56f45d47231d878617eaf40cb9179f7cd004e0f6594efc): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-046e2209ee25b522c6a2fcd471e11c2058e1303f5b5a433da6893c7746d9dc5e): complete subsection reference.

<a id="canonical-0e889b237e0b0fecf013f7708e08b623c7421e3485e1179a3489d1c7ffc468d2"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f4bf263f4324 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-016.md#canonical-3fada755790aeb178c56f45d47231d878617eaf40cb9179f7cd004e0f6594efc)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-046e2209ee25b522c6a2fcd471e11c2058e1303f5b5a433da6893c7746d9dc5e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3fada755790aeb178c56f45d47231d878617eaf40cb9179f7cd004e0f6594efc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9e9ba4336aebf7b261f910b943f03620df5ab492587637862271767feb07321"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f356a9c63712 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-0b63a27ab1c185f2e66c2ea255b045c5ff97d82618ad20e7ee16e9fa5cf8ba84"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-78c66492fbf5bd1bb3f716e652f15e72e6ccff846192eeb5815d02c8c0e06833"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f356a9c63712 / 3

<a id="canonical-83faf12b73fd7b95dd8dfcb840e9242e0ede0c771a32d7f0eb7d1f3ee25e67a3"></a>

<a id="canonical-bcc9c8c59d2bc3ce4de0bbe1b486eaffbcf1cb5d5a3cb11f0786de73db22d80d"></a>

## dns_list property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f356a9c63712 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-f584f1df4c57bb74d01c021bb19542ef714ce414ac50ee7ff17475efdf774cd6"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f356a9c63712 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-046e2209ee25b522c6a2fcd471e11c2058e1303f5b5a433da6893c7746d9dc5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d595adb583d476eece914fb595c2e8524b75f03703ac4c29a77582d2350ad639"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f7256c6aff91 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-e77a47b1cc1fc533fbc0976eb8c39061c0f82031815a0437f340615dafe75f05"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-614b762b17bb1266ac9c2c76271ef48cc1a43d3dca29a10fb79a90a74e4244c0"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f7256c6aff91 / 3

<a id="canonical-73158a6c96296d927cc97e3add4d44236507073d09ac3c0693a2f77df4f80aed"></a>

<a id="canonical-8846a3a66c74e84d1fc983d21ff50edbdfd65c7f74343a99330962dbb412ec4e"></a>

## configured_address property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f7256c6aff91 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

- [first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-d846ca6b11ac1774900a9e274d47ea34b7f9d15afe56fef00943d6e91a8cba81): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-d42e3cb78ca43260a2fc06b2f6030acf5a3a894a53a2153f2b267e2b2a17de08): complete subsection reference.

<a id="canonical-fc74073cc67a8451e0c55afbe6c6b3f896d229ea03cf940476e934769c43b651"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f7256c6aff91 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-016.md#canonical-d846ca6b11ac1774900a9e274d47ea34b7f9d15afe56fef00943d6e91a8cba81)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-016.md#canonical-d42e3cb78ca43260a2fc06b2f6030acf5a3a894a53a2153f2b267e2b2a17de08)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d846ca6b11ac1774900a9e274d47ea34b7f9d15afe56fef00943d6e91a8cba81"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58450cff9f45ed77896380d58dbae1c33a330f25ce8740df7ceb2ef5fcfcc013"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 640dbe56ba08 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-046e2209ee25b522c6a2fcd471e11c2058e1303f5b5a433da6893c7746d9dc5e)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-80f12d32782f2db2029be0f002f9ec66c972b2c239434eedeab867a44f2d33fc"></a>

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
first_address = {}
```

<a id="canonical-10af624f4807b661ea6fb17218e506f01dbee7cd188b903dddc3cacfde1a84b3"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 640dbe56ba08 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bef5c2bf1ca5d60b84f0a797295690410ec35a05b9ec186d42df21cde6d7a7e8"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / 640dbe56ba08 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-046e2209ee25b522c6a2fcd471e11c2058e1303f5b5a433da6893c7746d9dc5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d42e3cb78ca43260a2fc06b2f6030acf5a3a894a53a2153f2b267e2b2a17de08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17441a9bd7269a8fa8fc9754d1d51c470c0d4a65f1f52295d06a53244fdd3c51"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f71242ded8b2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-016.md#canonical-467088359d4331da6f62a6fae3077e8a6fd55856524423501296f15def0cdaee)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-046e2209ee25b522c6a2fcd471e11c2058e1303f5b5a433da6893c7746d9dc5e)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-bdf813c426c3294090614b0ca83dbd1c4717b04355874250fb87e15a4e091ee0"></a>

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
last_address = {}
```

<a id="canonical-bd31c6e9b2914dd5660bbfe9050df704b9e922fe3e3be4a50c71804769b6258e"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f71242ded8b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e3399b486649f8cf862410598c0358174e3ea90e77c618e5158e89daa04b052"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_confi / f71242ded8b2 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-016.md#canonical-046e2209ee25b522c6a2fcd471e11c2058e1303f5b5a433da6893c7746d9dc5e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5711e61d387d1aa0c84d699d989c664e023eba405d576dff365c8d9ecc7d79fc"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 4f044634d134 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-08fa0811445d5453935e1c1bf409f0cb86e72ff3425ab38d6258b907426c59c5"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-015e52b2ecddef823b2af8e5836c527055dfc17d4a1c7537c66f0116624d2ba9"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 4f044634d134 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-11b0d114228bd130575819115b8d98bde1945ebb94e04994d9d469c5b571e04c): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-ebaa3643ca4aae44978bcdaa0ab3b95a030dfb406a64823bf6002ea892665713): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-b47fee157ef797e3d9de6c8f5615f0ab2c6e4101fb397f2538baddd290c658a3): complete subsection reference.

<a id="canonical-68f4e82cdc20a4cf1b6bae0d3e03592634fbd257110e41db1ff25a7ec1dc43aa"></a>

<a id="canonical-e9eaf4c037346309ba6d4d106ef0a2cf3bcd83b9444a2fe5ee9aed2aedd642ed"></a>

## fixed_ip_map property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 4f044634d134 / 4

Type: `["map", "string"]`. Optional.

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-6f8ff032fb8db73c72a2e05c0b20154a1836fa2b37b2560cfec55c755a996d47): complete subsection reference.

<a id="canonical-cdea958dd5485a0bb8ae41dfd095ec5345d66f5041495acd9ebcac9693e749e9"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 4f044634d134 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-016.md#canonical-11b0d114228bd130575819115b8d98bde1945ebb94e04994d9d469c5b571e04c)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-016.md#canonical-ebaa3643ca4aae44978bcdaa0ab3b95a030dfb406a64823bf6002ea892665713)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-b47fee157ef797e3d9de6c8f5615f0ab2c6e4101fb397f2538baddd290c658a3)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-016.md#canonical-6f8ff032fb8db73c72a2e05c0b20154a1836fa2b37b2560cfec55c755a996d47)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-11b0d114228bd130575819115b8d98bde1945ebb94e04994d9d469c5b571e04c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1b601a31f65bf3761d021183c22908bee7a43b020a1e25eff182eb007a6d2b5"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 76189726d71b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-c4b222f2d3ebe3496e112533b6b2cef3683e9300537c6a6bb351e85a81a8fbb9"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

<a id="canonical-3da1dffefa5c81498fa65d6a1ac2d685e1f18fdd9385bf58a260409bde826897"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 76189726d71b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba762b1825ff0427866a1a60aa3c3c7c3a36b4f2943d41debae1af1523b0366c"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 76189726d71b / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ebaa3643ca4aae44978bcdaa0ab3b95a030dfb406a64823bf6002ea892665713"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f55ad9b5acc9159048eb69bd552f030c590735c77f36c57265a38a1caa81514"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / d37bab3c87b7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-dc4fdc1ec1541a4f60468a504cba1e27ff30c15461f768ca77a0de5b06130f0a"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

<a id="canonical-83b752419b596c411b5f7cd945e890d1f278788fa62eda2c5321635e99609dab"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / d37bab3c87b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc9c8c2e71de146aaa684cc7271c5adcbe868157b0ef60172be4c5476d61f5d0"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / d37bab3c87b7 / 4

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b47fee157ef797e3d9de6c8f5615f0ab2c6e4101fb397f2538baddd290c658a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2813005ca8c1982f309cfb4d0571517546bc520a5fe2ede547d50c939e506031"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 6d1708156182 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-381cd0ec51890e503e4b3ac3ae4a28326b72b53e649b57a4e4be48aef31e3941"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7bbda7353b1e5f1371b0ec30544dd969a4e55b3e33165fafe584203d642ad5a"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 6d1708156182 / 3

<a id="canonical-1c69ff641ce339db9dc4f384305f73c3d39aa713342a51022bcb374061328ba7"></a>

<a id="canonical-52acb2a30488bfc3ec178003e8e1557392edd23205e1852526e9e4b3a3578d47"></a>

## network_prefix property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 6d1708156182 / 4

Type: `"string"`. Optional.

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

<a id="canonical-aac9ff18095ab986056de88a7b8307ade2416ef9f628c7d0782a3a892da61ab9"></a>

<a id="canonical-b1f07001b12e4370fe27e3931ff2a84c2456de0188243db96e2ef74fb83a0e8f"></a>

## pool_settings property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 6d1708156182 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-016.md#canonical-5889c28169464fcd97c33d37287612e07f9bbbfb93f537a6d6cfcc332f15da98): complete subsection reference.

<a id="canonical-990fafc6771297738440b57c35b7c3d0b9ac22026c91aa186c7aa7b754148349"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / 6d1708156182 / 6

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-016.md#canonical-5889c28169464fcd97c33d37287612e07f9bbbfb93f537a6d6cfcc332f15da98)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5889c28169464fcd97c33d37287612e07f9bbbfb93f537a6d6cfcc332f15da98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac676a198d96043f9a10527fdac6182320fee1c9671038ee4a7afb5538d98083"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / c29b328b0a40 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-b47fee157ef797e3d9de6c8f5615f0ab2c6e4101fb397f2538baddd290c658a3)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-7726a292d0f052146b2e07c8a227977ebcdbacfe184016a98eb5f206ba305428"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2d3193b337f0de5b6c608c7f045941e23a481c4a0b261f7a82583772174561d0"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / c29b328b0a40 / 3

<a id="canonical-7af795ff9030ab0de61fcb0b059e16741e577767618d3afe0ab9e7dac723e28a"></a>

<a id="canonical-c204f0e7e3dc71129cf474099ce4701abc93dabe1f69f448e51c1758e558fae7"></a>

## end_ip property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / c29b328b0a40 / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-3325f023cd54b7cbec5107e61451fccd80bb17d0311fe5d1e41686dd592acd77"></a>

<a id="canonical-6f4afcb16a4607a9610f32e394217eba470b27fac4459cc20cb2216a1e6f97f6"></a>

## start_ip property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / c29b328b0a40 / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-b791b1fb55d806c37bfbd5a13d6b9a9fab62fa21594a650084feadeeea6b6d43"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / c29b328b0a40 / 6

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-016.md#canonical-b47fee157ef797e3d9de6c8f5615f0ab2c6e4101fb397f2538baddd290c658a3)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6f8ff032fb8db73c72a2e05c0b20154a1836fa2b37b2560cfec55c755a996d47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16de41ce447470e21745ac27a1163979e80397677fad028fd29e815eb71c46b8"></a>

## openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ec53b50d7231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-016.md#canonical-deb6ff9e926a7a6cf71c2d6263fe437709e1974ccbd1e67ba7892cccfdfe6181)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-df6086bc2a15cd00e6c4e29199589b0031b35f71df55190f13d10c504a9386bc"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ee9bf6952a554abac9b595bf8126ee12ae03b1d7d937e46742ee4958bd2126c"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ec53b50d7231 / 3

<a id="canonical-e67f315b2af9011a829124b7e8b51cd5283064f6e365858ff4fcc6d900f6a610"></a>

<a id="canonical-9356043670c9bc2e96466f07b7f0ec99e1f7b71d3057298c6f06c4c4996af7c2"></a>

## interface_ip_map property — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ec53b50d7231 / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-61640600da6ff50eb7b9d208dab412a5a153ac06bfbd9ec9b2c9fb9060f6d539"></a>

## Next pages — openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful. / ec53b50d7231 / 5

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-016.md#canonical-c0c79b04d5f9d86794b87fe19e9987fdb2343fd8f9030cb507b5f973cd5b85a1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d15ef01be5bcf8cdec599778c9387338a52d7c01027f83f0d0bcb23889c73822"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a92cccb23e7fd0015d1c31a6721c8f832252f8b634f5905c77da04bef612a98"></a>

## openstack.not_managed.node_list.interface_list.monitor — openstack.not_managed.node_list.interface_list.monitor / d337d360bc8b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.monitor

<a id="canonical-c223443a0f6ddcccb9db2050233b992ab041403ce26343c346fd1eda5fe84428"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor = {}
```

<a id="canonical-a010414b7dc7fdc7efc0a10962653ce1425435abdf16b847b2651f28dd817d91"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.monitor / d337d360bc8b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b4df04615f4fda88f169a58cb1e120a84769452c6c421e4508b7f4f5d919d9d"></a>

## Next pages — openstack.not_managed.node_list.interface_list.monitor / d337d360bc8b / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ea6fac01147da19415ddac1904373e25dc5aa43c5c7ef087ac3b54a525b801a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92cc0b99d7e7577ea5a8508a498fc80994551e09d9b52f43b0b42d1fb4585237"></a>

## openstack.not_managed.node_list.interface_list.monitor_disabled — openstack.not_managed.node_list.interface_list.monitor_disabled / 7fdcd9e75e8e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-12c42bd7f7651623c282e6f75507082a640604385da11d847782c902cd17b95f"></a>

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
monitor_disabled = {}
```

<a id="canonical-960eacda70c909b066a99cb13ce5b8ed3bc6cc72bc574f0324a958c4a23a2fe1"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.monitor_disabled / 7fdcd9e75e8e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c370bd68f42f451154bd6dd4ced768d05e14c3a17bee1d2f8a0b35173ba728cb"></a>

## Next pages — openstack.not_managed.node_list.interface_list.monitor_disabled / 7fdcd9e75e8e / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-cb95685c6151aa5ab0e934beb4bd98961832c7964b6c1fe11e5f1ae15bf27284"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5db5c2b1ef13340079c574bd97520d7f90a07a168b6cd79e85e025e66b3af4f"></a>

## openstack.not_managed.node_list.interface_list.network_option — openstack.not_managed.node_list.interface_list.network_option / f78a1dd34425 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.network_option

<a id="canonical-692b5c39728c747e83f1807ca449b3708ff8991f76b1f90f77af13578ba34855"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-a13bee1a89cad4fe5ce8a83642b0c61893d7c2a797b45035caa748262f2a5875"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.network_option / f78a1dd34425 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-016.md#canonical-0ecc571929ae6a2ece0560eae4dac0481b70db68315ca2a5158c91d1991fe172): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-016.md#canonical-59a60a45129e1625d7f6ff8599c15fc474f4e999d03ac63991a7fb32e05bc4c6): complete subsection reference.

<a id="canonical-34f522ee5dbb35ef62acfeac0a15526c73a839aac564cb9038684292b143cfd8"></a>

## Next pages — openstack.not_managed.node_list.interface_list.network_option / f78a1dd34425 / 4

- [openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-016.md#canonical-0ecc571929ae6a2ece0560eae4dac0481b70db68315ca2a5158c91d1991fe172)
- [openstack.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-016.md#canonical-59a60a45129e1625d7f6ff8599c15fc474f4e999d03ac63991a7fb32e05bc4c6)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0ecc571929ae6a2ece0560eae4dac0481b70db68315ca2a5158c91d1991fe172"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bacd70d64ff9b99dd00c3147269a5833edb4b4ed7fcf47e608cac25690fdc1ce"></a>

## openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network — openstack.not_managed.node_list.interface_list.network_option.site_local_inside_ / ef70cea6d581 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-cb95685c6151aa5ab0e934beb4bd98961832c7964b6c1fe11e5f1ae15bf27284)
- openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-350b699ad5e2d504eeb61a9c53c6e99487983399f8b24fc197c4be97e35c97d6"></a>

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

<a id="canonical-70eeda14f6055b6d311139642a30c0002d22e570842c96b2fdd56fb851907b6c"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.network_option.site_local_inside_ / ef70cea6d581 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc046bddf6742a3787e3418b3789643a868c328b8aac7f76c10d61251cf17f80"></a>

## Next pages — openstack.not_managed.node_list.interface_list.network_option.site_local_inside_ / ef70cea6d581 / 4

- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-cb95685c6151aa5ab0e934beb4bd98961832c7964b6c1fe11e5f1ae15bf27284)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-59a60a45129e1625d7f6ff8599c15fc474f4e999d03ac63991a7fb32e05bc4c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f965169b78f999479686aec947fd7f8c50ff881efb50e9593ffc64ddaf8f06d"></a>

## openstack.not_managed.node_list.interface_list.network_option.site_local_network — openstack.not_managed.node_list.interface_list.network_option.site_local_network / 6ccdaab5e70a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-cb95685c6151aa5ab0e934beb4bd98961832c7964b6c1fe11e5f1ae15bf27284)
- openstack.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-8a99b4635926d77a1a90cd6272021ff1fc2301aa23939193e3980efbe12cefc6"></a>

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

<a id="canonical-460b8e8b8823d972409c90c54e394f9b0923b01c6de7845dae9dda2b8f69f153"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.network_option.site_local_network / 6ccdaab5e70a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4bf010348aa0c331777f966e2e1497590f240b785f6af2f7c2c58653de0a8ba1"></a>

## Next pages — openstack.not_managed.node_list.interface_list.network_option.site_local_network / 6ccdaab5e70a / 4

- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-cb95685c6151aa5ab0e934beb4bd98961832c7964b6c1fe11e5f1ae15bf27284)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bed5566f4b80eb81da7cc71c5bdd2eec20a5a942ec7f68451a7a038c079eeb6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d3f5513ae66725d2fe356aae79b49fafe5f44d1873c4ad00859733cc9167101"></a>

## openstack.not_managed.node_list.interface_list.no_ipv4_address — openstack.not_managed.node_list.interface_list.no_ipv4_address / a92eeaeeff8c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-f1927622d732ec5125d11364912af940b65ef1c495ac693d8e3873a00390683c"></a>

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
no_ipv4_address = {}
```

<a id="canonical-f3e74491a6fc2b7e6f976662ca93f74a4e5139cc110fc48d44f93dea02994329"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.no_ipv4_address / a92eeaeeff8c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66f5f8fe1500376a38e15b24c5ec5b92e11428b02282ffdd4d4af85f3d22312d"></a>

## Next pages — openstack.not_managed.node_list.interface_list.no_ipv4_address / a92eeaeeff8c / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-94de12dbb480802e75b81333aa41337b25ddd2ccb03d31b86f2b3cfd0582523c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4e03b258848bef28a5bec061fc23261853120fba53bf90281fefdd9840f8f0b"></a>

## openstack.not_managed.node_list.interface_list.no_ipv6_address — openstack.not_managed.node_list.interface_list.no_ipv6_address / de7bc638d0d1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-eaf8f6a1e5c8cf2eff6c3993490d953bf24785b7489ea201091d372064f05935"></a>

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
no_ipv6_address = {}
```

<a id="canonical-081d3f2692d03056fad61498b6ee4705a5918fe37c8de7ee8c7d331674cafd14"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.no_ipv6_address / de7bc638d0d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe3323070080aecb1c522a02b74d1a3bf5f671ac66530925ba5006fa30707876"></a>

## Next pages — openstack.not_managed.node_list.interface_list.no_ipv6_address / de7bc638d0d1 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-241b19aa15996ff476d77b5d3b40e12c0c3521d8756ea8fc130e5432a2a2541f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-944cbf9df179bf41bc62a4cdceb0092df5614c40aec330ce486e33a67da2ef0a"></a>

## openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / e1173a7f9986 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-6951e62ec571d9e160e8be929a0187effa403ef16ea857e33ca65170ce9e8b3d"></a>

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
site_to_site_connectivity_interface_disabled = {}
```

<a id="canonical-435c16d3e9e3f02cf348e054cb860e1b9c329fcbad90cb2b1dc18ca485d5d944"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / e1173a7f9986 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bb5dba4d4c13a3f04ddabdff1fb6f854df6c3453fbef6556ec5346f8709b6df"></a>

## Next pages — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / e1173a7f9986 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8df2288afd00d05060dd889af762fd6ffaf46a730b62a5b95ce642efc5de3e45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a9c06acc02f50ccdcc63ed24874cbc1d5c115f9d9354e1bf51387ab80a90abe"></a>

## openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 9913c4dc29e9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-1cb5184b039d2232a8c0b22507ee1d1877bbe2aff471c05552b17f0ecda79437"></a>

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
site_to_site_connectivity_interface_enabled = {}
```

<a id="canonical-ebda385ae32e53e2a66a3aae35275a7b67891d49dd2da0e165447a1ec82a7f74"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 9913c4dc29e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-741266574db864b64a48489d9e461abc141fb16c6558006779d340047bc8597e"></a>

## Next pages — openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interfa / 9913c4dc29e9 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-96b72e06db4f8bdc4a8937b219e59de5d5162120aad6011fc939fafd3bf1c0b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df7560889c4ad4ccc1e59b4e8fd7318a68102d27e6fc1e8a4c51f50126190fe9"></a>

## openstack.not_managed.node_list.interface_list.static_ip — openstack.not_managed.node_list.interface_list.static_ip / 12a1c5762a75 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.static_ip

<a id="canonical-4f2a04365d9a845b752ece2a95d4f519154e20395c93006bd4563fc1be960998"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c0b8e2af00556e650d52992b3a98a2f14c98143f311aa3799c587c8c4f2099a"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ip / 12a1c5762a75 / 3

<a id="canonical-4be00e14502f61e040ede89365fb230d04beb63df7dfe43f97aac52598668155"></a>

<a id="canonical-842c3b49a54d659215823cddc3042423c000e191326aebeca99dd97a8745d143"></a>

## default_gw property — openstack.not_managed.node_list.interface_list.static_ip / 12a1c5762a75 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-b88b6f71c66da82c8d8483284849ed383e71eea1b130cad20df8c78c61c5da5d"></a>

<a id="canonical-a9d3557a907d46edf9c9435409a6f8e3134df876b4fbc69afe597fd7ad5fa1f9"></a>

## dns_server property — openstack.not_managed.node_list.interface_list.static_ip / 12a1c5762a75 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-e7ce535b4f990a2e56ab92c4f8bf3987c90bbe087e386264eb647be4271d445d"></a>

<a id="canonical-840a2588219a0847e159fd06ebb5f4aa9483cf2455153dcf70ba57a305ed341b"></a>

## ip_address property — openstack.not_managed.node_list.interface_list.static_ip / 12a1c5762a75 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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

<a id="canonical-e4040115337d987e73e36c10ab9909ef4e6237603bae81651c8008a40dcfb6dc"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ip / 12a1c5762a75 / 7

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a49e266a35a0b9739890bfa8a72a4e857cc5e14830ea12771c42acf75ea18783"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f314f4b5990bbc12e970ef11d0849b71fd39ffe72ecfceb4bcabf03258e756e3"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address — openstack.not_managed.node_list.interface_list.static_ipv6_address / 7de5e75ab54f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-fea3b6ca4d80a3146dceab080492b27b0efdeda2b5ae4ac53364dcc72569a8f0"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-e5bb5bbe5439e919e37704046053e2bf36832913e0e86abaa2d425a41300101e"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ipv6_address / 7de5e75ab54f / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-7a9f67ae631ea2dc8d6ad3e3d80c0e40676c90965f1c6f5284011afda0dbdc80): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-cb21c4768cf0778008308c12dd976063aaeb2735aabd9a06c3f1b9c00d3d2f2d): complete subsection reference.

<a id="canonical-a163e6b3ed86ba6ac825c84ba3e73e3baa3aeb1be497539b3a899cab67e995af"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ipv6_address / 7de5e75ab54f / 4

- [openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-7a9f67ae631ea2dc8d6ad3e3d80c0e40676c90965f1c6f5284011afda0dbdc80)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-cb21c4768cf0778008308c12dd976063aaeb2735aabd9a06c3f1b9c00d3d2f2d)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7a9f67ae631ea2dc8d6ad3e3d80c0e40676c90965f1c6f5284011afda0dbdc80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dfab61b7dd2cac918add2b0a29e8ae12d96a8fad20538ba75c8b9dfb31be1d2"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 532b803f3cbc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-a49e266a35a0b9739890bfa8a72a4e857cc5e14830ea12771c42acf75ea18783)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-d1d980aa14a9427434f2daa6dc52d211d88195e13f4a23278256eeacd2cc931d"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-88d47667149e1fdaaf6692ddd5256450ff49f0567d1741856502e905c1678371"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 532b803f3cbc / 3

<a id="canonical-f3c70346c51bc2276eb5f079687cbe368f3df1e0752402a55ccc15033ea65cf7"></a>

<a id="canonical-9275e66c868a754a134f93e59ea3b36d49f7dc263ba9f5fedc8d4acd72a62421"></a>

## interface_ip_map property — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 532b803f3cbc / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-0318906af5b4aab24dadffd6a9ac215f6c51c04be76d447c8c68ac56d1fdc293"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_stati / 532b803f3cbc / 5

- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-a49e266a35a0b9739890bfa8a72a4e857cc5e14830ea12771c42acf75ea18783)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-cb21c4768cf0778008308c12dd976063aaeb2735aabd9a06c3f1b9c00d3d2f2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdd29c8b0cb2904f157a0d900200a1024a6f89381908780048c281c886d69de8"></a>

## openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 2e685fe99754 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-a49e266a35a0b9739890bfa8a72a4e857cc5e14830ea12771c42acf75ea18783)
- openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-4c68c0f8fa43c10093eab57e7aba4a7d5309f7bf9c6320372f6284eea777070f"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-9600dab02d8e925e7742109c7aa4fcfc20f541a6a6b4d7c22b62bc80e0f59105"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 2e685fe99754 / 3

<a id="canonical-b7ced441eb554d1b4fd2b071022654705e614340a0ba3c0e5b4c26f8efe8b79a"></a>

<a id="canonical-cf325fe3e1835787867d77091f90a9a518acb282314675a0525f4dad5ac5e42e"></a>

## default_gw property — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 2e685fe99754 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-46508d7c199c317be44e89ad8fed5a98bfa401efbeebe1989e1b6b2f19324bc0"></a>

<a id="canonical-589ced88d83839ed28dd17ddc682a838542ab9667ecbb3c0a7ae9e2a0b1e7b53"></a>

## dns_server property — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 2e685fe99754 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-36b55047412a5dc4b44a1156b8963820612e437a0b5f9f3ced4dd1f29368d7f3"></a>

<a id="canonical-1af5669175b5087a5580fbf3b2703ac0b7539f25fb01d02bc7790031a3eb60f3"></a>

## ip_address property — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 2e685fe99754 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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

<a id="canonical-c72a249a63ee5f3ef8b1931a47fdf87dfb07bae1ae4ae1485e9d1d172af6162d"></a>

## Next pages — openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_i / 2e685fe99754 / 7

- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-a49e266a35a0b9739890bfa8a72a4e857cc5e14830ea12771c42acf75ea18783)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-08f097c51aaeacfbbee9ac8d133b0e2d7a629f88028a6cd15205e2fa8536115c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b20b5b77daf27869c38afd4fd9ae15445986136a50599f120d5db8784d92d074"></a>

## openstack.not_managed.node_list.interface_list.vlan_interface — openstack.not_managed.node_list.interface_list.vlan_interface / 008e2998ee58 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-9073fc1df0def01e293fd37f403badd01d78a0eb22cb51dad2bdb836b0ddd291"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-d2a136993c6db358b585094ad937e955270c8e86a09c60b1b87c2ff832828325"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.vlan_interface / 008e2998ee58 / 3

<a id="canonical-de8bef5a8089d9d3b0d31b9329c3e93c5abb5b1d9d328639027021efa6652bb6"></a>

<a id="canonical-f90ac739749404687ba6d3e3124ce2d7aeefcd1205c3ed4e14826ca51d49d227"></a>

## device property — openstack.not_managed.node_list.interface_list.vlan_interface / 008e2998ee58 / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-3f1465724bd6c5a9e079ff9735d969731ad6299c4bff3d80d2e65ee2d1184880"></a>

<a id="canonical-dbbe2f8e4324b4b22a4b36bc35ce61e7dfaed831273ef9868c572a05326f381e"></a>

## vlan_id property — openstack.not_managed.node_list.interface_list.vlan_interface / 008e2998ee58 / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

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

<a id="canonical-f02f2af50782c2f891c013e539d3242e10b5aca712c92e15140a71697b4bce12"></a>

## Next pages — openstack.not_managed.node_list.interface_list.vlan_interface / 008e2998ee58 / 6

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a22cb7ef0a5cfcfa4491159d5c61c08336bb2976b7b29d3626d569ab3cf8be47"></a>

## performance_enhancement_mode — performance_enhancement_mode / 550be1253446 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- performance_enhancement_mode

<a id="canonical-eade2ed42865adf59611ea79b3932091e5a15e1e7497541b9a9fd43b09039aa5"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-e86d1b52934f44ee7c2bd9c577976f44bb3d0670c58bf9dde7b82e3efc94f420"></a>

## Direct properties — performance_enhancement_mode / 550be1253446 / 3

- [perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-f6daeb8201c07ad705459cbb1bd0c0f6208dbc055ea9263cecb52d1e09504ab8): complete subsection reference.

- [perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3360556318dd25e26eee3b68b394be125c4a3d27e02124815935c74e22aee8f6): complete subsection reference.

<a id="canonical-ac9fb478d1e37d8ad6b67389b725c959285cfc1f089d07f04c814f8a786efeff"></a>

## Next pages — performance_enhancement_mode / 550be1253446 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-f6daeb8201c07ad705459cbb1bd0c0f6208dbc055ea9263cecb52d1e09504ab8)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3360556318dd25e26eee3b68b394be125c4a3d27e02124815935c74e22aee8f6)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f6daeb8201c07ad705459cbb1bd0c0f6208dbc055ea9263cecb52d1e09504ab8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d81d0717379c3a0bfcde086d1152001cc685f053317592f1fd0498d65c1c39a"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — performance_enhancement_mode.perf_mode_l3_enhanced / 8d60488067cf / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-98bca32084f7bce76bfc1a02cf3cebe5f191fd7e66390c12ec148a1883dc7521"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-eac8b8734885dde6191c014069d790dbaf2ae1f0cd803b5fbb3ec356e6a920cb"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced / 8d60488067cf / 3

- [jumbo](resources--securemesh_site_v2--reference--group-016.md#canonical-5624c2a22e41fcbad202b266b976f190e80ce3dee2ffd0b0d862a7a793801fa2): complete subsection reference.

- [no_jumbo](resources--securemesh_site_v2--reference--group-016.md#canonical-9ff4df2b6fb42579311cd63c0d89070631ebc91b1e36bae6a5986c96d5ae2df5): complete subsection reference.

<a id="canonical-faeac67e3701bc39615d957da122a01eb49c110ca6584869ee64bd8a5af7f3a1"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced / 8d60488067cf / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--securemesh_site_v2--reference--group-016.md#canonical-5624c2a22e41fcbad202b266b976f190e80ce3dee2ffd0b0d862a7a793801fa2)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--securemesh_site_v2--reference--group-016.md#canonical-9ff4df2b6fb42579311cd63c0d89070631ebc91b1e36bae6a5986c96d5ae2df5)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5624c2a22e41fcbad202b266b976f190e80ce3dee2ffd0b0d862a7a793801fa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aad96a1258ad08d2674d3dbca4fdfe5aa7f326632cc59ac633c5be630de58c9"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 8c8cec6ee8d6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-f6daeb8201c07ad705459cbb1bd0c0f6208dbc055ea9263cecb52d1e09504ab8)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-a473e2da05d4f9b8d9b97ed7afb09723d4b8ae6f3bedf4a1c2454f154b4fb1a0"></a>

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
jumbo = {}
```

<a id="canonical-b19a07f76f3aa841957717a3b5289f56a218ad859abb91cb869745c99fe59662"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 8c8cec6ee8d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c02d13f57aadb00bb68f91da9a4c755a87416ed4cfcb7924529817c5a9407904"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 8c8cec6ee8d6 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-f6daeb8201c07ad705459cbb1bd0c0f6208dbc055ea9263cecb52d1e09504ab8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9ff4df2b6fb42579311cd63c0d89070631ebc91b1e36bae6a5986c96d5ae2df5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8446b3c2eb5b0e4ef5d2003620ab04bd87a2a397f0b7066701e4da7ba44494e6"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 2f3bf664e390 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-f6daeb8201c07ad705459cbb1bd0c0f6208dbc055ea9263cecb52d1e09504ab8)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-4f49dcf9800109b4b3843deaa6c499d5067f1036964e39bf160dec31539aee26"></a>

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
no_jumbo = {}
```

<a id="canonical-6999df043854f891dea00c0d2d5373aa3f6b061373e7e5ffdb0d19f5c44e11cb"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 2f3bf664e390 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d925cdc76d919ee89ebce4d1876c71e42377f9cfd74ee7165f74d1ffdf0b76d"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 2f3bf664e390 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-f6daeb8201c07ad705459cbb1bd0c0f6208dbc055ea9263cecb52d1e09504ab8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3360556318dd25e26eee3b68b394be125c4a3d27e02124815935c74e22aee8f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b1e80f9681f5f9ee5c89f9588726799d1e4727773dcb37890964dd11c4bf5a4"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — performance_enhancement_mode.perf_mode_l7_enhanced / d33c09dcc695 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-468ac3a1bdd30c1ef95711f05555dc42dee738034e6b49a8a9bfc025a4fbfa4d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-b62a0a51ff3d0c0a6bfb6bf98ebc3fce6495c132d71a494e55995ccaa05e8031"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced / d33c09dcc695 / 3

- [jumbo_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-8eeefcd20d93f04b66a15c43c37c119872cfa65dccd65678b70264b46749605a): complete subsection reference.

- [jumbo_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0fa9731e8dc225aa0be07a6240b9f4d164da01f51d6a3f48f02af109212530df): complete subsection reference.

<a id="canonical-f320ba25bcdccaceb34e73a75e44ccb57d10cc348f50d1dbf279da5d10976143"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced / d33c09dcc695 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-8eeefcd20d93f04b66a15c43c37c119872cfa65dccd65678b70264b46749605a)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0fa9731e8dc225aa0be07a6240b9f4d164da01f51d6a3f48f02af109212530df)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8eeefcd20d93f04b66a15c43c37c119872cfa65dccd65678b70264b46749605a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0747367ea94e12a320a4e640d895cb1b4162c1f4c7d8928132dae0459dcd175c"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / fc83ba50f7e6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3360556318dd25e26eee3b68b394be125c4a3d27e02124815935c74e22aee8f6)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-d59271411abe9b326b8ed6b29c163b78d6d98754087c67b5983abfed9502080e"></a>

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
jumbo_disabled = {}
```

<a id="canonical-ac57bde415d9dea72fc91b89027d5a39c470d9499ec2f91caa530e3424858eb8"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / fc83ba50f7e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72f323edb54079fae7f267fe170df6c94685595ce2d30ec0a8b06500eabd8e41"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / fc83ba50f7e6 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3360556318dd25e26eee3b68b394be125c4a3d27e02124815935c74e22aee8f6)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0fa9731e8dc225aa0be07a6240b9f4d164da01f51d6a3f48f02af109212530df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe71352e0a122db431cdbd2d80d5a62a341c2ee5af80172826af190a27945a7f"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / dd4b500c99bd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-9c1dd3edeeae4d7159220726451a7cd25819ff4322f4e558b219bfdd30d81450)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3360556318dd25e26eee3b68b394be125c4a3d27e02124815935c74e22aee8f6)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-42a192fdd473bf22d530d7154841e1583ffc3e4d9856baeb3259e83dac1771e3"></a>

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
jumbo_enabled = {}
```

<a id="canonical-06c210fa343aee97006119ec4b130c71a8ea04dc9f8f19ec6242680804d0ea7a"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / dd4b500c99bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31c581473778b5ab83d5a1758188c516ef35cfb43b35c44492ddef7b3789f06b"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / dd4b500c99bd / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site_v2--reference--group-016.md#canonical-3360556318dd25e26eee3b68b394be125c4a3d27e02124815935c74e22aee8f6)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5efc800fd9dc85dc5717b76d73721de2fd8a153cee308446056e512d71e786fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b02df9a90ce40b85c72b8de270260ab5b1c7a8c1c50e24468ba6724beb30b5d"></a>

## private_adn — private_adn / 7d41c75bb81a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- private_adn

<a id="canonical-c1405fb3ad8eb0256c830bf7862ab3fd42061bad7616d6f588d5181304384e28"></a>

Type: `"object"`. single nested block, Optional.

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Upstream description:

X-required Establish private connectivity with the F5 Distributed Cloud Global Network using a
Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud
support.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("private_adn")}
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
private_adn {
  # Configure direct properties listed below.
}
```

<a id="canonical-6af9e69e3304af92c997d7d1c0b6f50ba7759de639c4cd68438867754d55c36d"></a>

## Direct properties — private_adn / 7d41c75bb81a / 3

<a id="canonical-fe1f7b664ab6e8639991c41d2f78ea1f7f24f43f86fb779efef4111bffc53b11"></a>

<a id="canonical-07595109ce18a54992b2e38405751d58a0739cd3568332038e034ba46d49b3cc"></a>

## private_adn property — private_adn / 7d41c75bb81a / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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

<a id="canonical-d48b3e600ced681f90cd61f1e93006d7a0db9f785214de9b8d3f6d3569062241"></a>

## Next pages — private_adn / 7d41c75bb81a / 5

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f20247f2bf62bff265884e5308bcb676fb69590f189a42d83f4a0123f809aa1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1510eabf4cec5a0788a36d01a051d821b95efa61f9106d296da259d262a02e8a"></a>

## re_select — re_select / a7e6b40c9564 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- re_select

<a id="canonical-d151b39e2822edde73f2705861dcdc72599dd5f44228c5e58a3c6ad3e2103be9"></a>

Type: `"object"`. single nested block, Optional.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("geo_proximity",
    "specific_geography"),
  validators.ConflictingObjectAttributes("geo_proximity",
    "specific_re"),
  validators.ConflictingObjectAttributes("specific_geography",
    "specific_re")}
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
  "x-ves-oneof-field-re_selection_choice": "[\"geo_proximity\", \"specific_geography\", \"specific_re\"]"
}
```

Terraform syntax:

```terraform
re_select {
  # Configure direct properties listed below.
}
```

<a id="canonical-677bde04b2d7268a8c960427a830bbe5a688d5429551b48c5a2a555a6f55bb08"></a>

## Direct properties — re_select / a7e6b40c9564 / 3

- [geo_proximity](resources--securemesh_site_v2--reference--group-016.md#canonical-c9e5252954e0f4b8b5ce8b686060205c792be3a10fb923f0eb1edf73c1b3aad0): complete subsection reference.

<a id="canonical-54240da1f45be8a4a372861788e6eacab85fdd9934e9ccad1afdd1eb36acedfc"></a>

<a id="canonical-bd4ec152a2f7d50d3d684bf474312c4e6b8965001187d435361a33eb03af8384"></a>

## specific_geography property — re_select / a7e6b40c9564 / 4

Type: `"string"`. Optional.

Geographic selection for the site's Regional Edge connections.

- [specific_re](resources--securemesh_site_v2--reference--group-016.md#canonical-65d836bc7d898a3c72016e012e2d570e29055f0528278b24cf3116f2f865effd): complete subsection reference.

<a id="canonical-e24a3d258753f4fab8a2e983d18368e21c3793389d184fc9e6e9ec3f9d6b44a8"></a>

## Next pages — re_select / a7e6b40c9564 / 5

- [re_select.geo_proximity](resources--securemesh_site_v2--reference--group-016.md#canonical-c9e5252954e0f4b8b5ce8b686060205c792be3a10fb923f0eb1edf73c1b3aad0)
- [re_select.specific_re](resources--securemesh_site_v2--reference--group-016.md#canonical-65d836bc7d898a3c72016e012e2d570e29055f0528278b24cf3116f2f865effd)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c9e5252954e0f4b8b5ce8b686060205c792be3a10fb923f0eb1edf73c1b3aad0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0748c9410171e18bcb8d498346be0112886b73e3065771106247f9972e4e55b3"></a>

## re_select.geo_proximity — re_select.geo_proximity / 9ec812fb3929 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-f20247f2bf62bff265884e5308bcb676fb69590f189a42d83f4a0123f809aa1e)
- re_select.geo_proximity

<a id="canonical-d10073289ddeef46e021a52c075182c3d9c5222c1c1bac3e7ee6ab4c14d7e412"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
geo_proximity = {}
```

<a id="canonical-a5ec9cefd8eed1fab58047446e2e2a7c2d388965cf5e98eff98ed8051656e10f"></a>

## Direct properties — re_select.geo_proximity / 9ec812fb3929 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53d40f3ce24053d58da9ad225475bae164c0fc612cb9362c7e372cb997dd0d0a"></a>

## Next pages — re_select.geo_proximity / 9ec812fb3929 / 4

- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-f20247f2bf62bff265884e5308bcb676fb69590f189a42d83f4a0123f809aa1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-65d836bc7d898a3c72016e012e2d570e29055f0528278b24cf3116f2f865effd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbc988d8e282351d64750467abd8eb2b28f7af1371404f02acae7eddfec586b4"></a>

## re_select.specific_re — re_select.specific_re / b9b9917e5aa0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-f20247f2bf62bff265884e5308bcb676fb69590f189a42d83f4a0123f809aa1e)
- re_select.specific_re

<a id="canonical-eb590cae2bf17720c887d3d2f1285e27afa97912e9e376cce84a1b880e9671d3"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
specific_re {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e0c4cd24c165f9f3e439e4180364c871b9547ab162d2715ca199c411160c361"></a>

## Direct properties — re_select.specific_re / b9b9917e5aa0 / 3

<a id="canonical-3d84c4c17c894f722a6d71736fafd58d815f96b14ee41a545440a3d74f6ff34a"></a>

<a id="canonical-3fb8e60031242e11fc1c2cd9799e717afa071098e1ec7ac398ae09d0e994d275"></a>

## backup_re property — re_select.specific_re / b9b9917e5aa0 / 4

Type: `"string"`. Optional.

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

<a id="canonical-730bb71bc6f7bff4a78f22deaf0260ecc188679a05d4181124559c09d3ca74bd"></a>

<a id="canonical-7ee34e9a1b3cbb8b76be1bf274c555cd1edfead0d5343b93450adb07b9ab0212"></a>

## primary_re property — re_select.specific_re / b9b9917e5aa0 / 5

Type: `"string"`. Optional.

Primary RE Geography. Select primary RE for this site.

Upstream description:

Select primary RE for this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

<a id="canonical-dddadb14b02c37084286c817c346040e2c3fae9abca06392cbda259a1e3ce90c"></a>

## Next pages — re_select.specific_re / b9b9917e5aa0 / 6

- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-f20247f2bf62bff265884e5308bcb676fb69590f189a42d83f4a0123f809aa1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e921978e1569fbf8a1537bc07a0003cb05c47438b26ec138c823548220b8dd86"></a>

## segment_vrf — segment_vrf / 86e267231479 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- segment_vrf

<a id="canonical-ec26fec7a6e7baf4eece1aea0591412c2fe61db1fe3707220a62bc272929f5ab"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
segment_vrf {
  # Configure direct properties listed below.
}
```

<a id="canonical-a66bac95b8c6771b6d6cf74f0b6db4e6791f791e3683d1c86ca8fe70bf692cd2"></a>

## Direct properties — segment_vrf / 86e267231479 / 3

- [segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109): complete subsection reference.

- [segment_network](resources--securemesh_site_v2--reference--group-017.md#canonical-d49d241ba738275d8158a3a45f825bae3e3f5fd7e47bcac432b0148e376762a3): complete subsection reference.

<a id="canonical-9540ee0fc51ea5d765fa87c752c637b5f3ae55a945cb4ad002d545329db06228"></a>

## Next pages — segment_vrf / 86e267231479 / 4

- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_network](resources--securemesh_site_v2--reference--group-017.md#canonical-d49d241ba738275d8158a3a45f825bae3e3f5fd7e47bcac432b0148e376762a3)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a1032b866a245835ad5945034f4b3e2eba67d86244ebc488597675e4aace88f"></a>

## segment_vrf.segment_config — segment_vrf.segment_config / b8aaa68ddbdc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- segment_vrf.segment_config

<a id="canonical-59e409ed9e18848905fc11cb19d09250f56fdf6d4cd351034c7c20feca16ff38"></a>

Type: `"object"`. single nested block, Optional.

Segment Network Configuration. Segment Network Configuration.

Upstream description:

Segment Network Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
segment_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-270bd8aa8733a3d3f1e5b00167b41982ec5cec08cc57f415a5cb69f34fe99cdb"></a>

## Direct properties — segment_vrf.segment_config / b8aaa68ddbdc / 3

<a id="canonical-40a037f6a2c2121acebd763246c1e7696fd1346cb4eea8106260897d311a2b77"></a>

<a id="canonical-905a45c11ac8f9e0146ea848102e8e017608e9936ec0a4611c41130045c1138d"></a>

## nameserver property — segment_vrf.segment_config / b8aaa68ddbdc / 4

Type: `"string"`. Optional.

Optional IPv4 DNS server to be used for name resolution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

- [no_static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-7c392a4c9071eff79f4fda91043c27bed074c62386c605a9da4ceae84075955a): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-5a55c71974eac5d8f84dd2a7bc5d894ba761153d14c626f8e46cba1ad918e1b3): complete subsection reference.

<a id="canonical-08670df3941d66d844a6a35767a09534b91e7e13a39c40e6ce65f7a4bd86d81c"></a>

<a id="canonical-55e7c3179b838aefe019d2258111c39cc9980b4f6731a36d8ddb57f8ffe1fa40"></a>

## secondary_nameserver property — segment_vrf.segment_config / b8aaa68ddbdc / 5

Type: `"string"`. Optional.

Optional Secondary IPv4 DNS server to be used for name resolution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

- [static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a): complete subsection reference.

<a id="canonical-3b7d3cfc6faf33538c59482aeb5870e3319a8c78f98d13bf2d98b083837a89f6"></a>

## Next pages — segment_vrf.segment_config / b8aaa68ddbdc / 6

- [segment_vrf.segment_config.no_static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-7c392a4c9071eff79f4fda91043c27bed074c62386c605a9da4ceae84075955a)
- [segment_vrf.segment_config.no_v6_static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-5a55c71974eac5d8f84dd2a7bc5d894ba761153d14c626f8e46cba1ad918e1b3)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7c392a4c9071eff79f4fda91043c27bed074c62386c605a9da4ceae84075955a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a475804641caf93ba2a5e186cc99e3fc20e75770a8a14874dae48b01e2c9d1a6"></a>

## segment_vrf.segment_config.no_static_routes — segment_vrf.segment_config.no_static_routes / fd3e78ae8f7e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- segment_vrf.segment_config.no_static_routes

<a id="canonical-b604dc2a47574a40852eb389342dd5e9fac683f18a917f934088bbde75dbf654"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_static_routes = {}
```

<a id="canonical-f496da0ceabdbc08468c6253fbc42d2a831a7e61a6688543037196f346c91b16"></a>

## Direct properties — segment_vrf.segment_config.no_static_routes / fd3e78ae8f7e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d7b261a94fe15f72ef51e482dd86a5f3ff8a497220e477bc61e1e8e613cdb81e"></a>

## Next pages — segment_vrf.segment_config.no_static_routes / fd3e78ae8f7e / 4

- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5a55c71974eac5d8f84dd2a7bc5d894ba761153d14c626f8e46cba1ad918e1b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9216a885d0af8930689ef79bb1013a4d3fc86a6bc2db12d457d9dc3d3f8dcab3"></a>

## segment_vrf.segment_config.no_v6_static_routes — segment_vrf.segment_config.no_v6_static_routes / f2aa3b544fc3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- segment_vrf.segment_config.no_v6_static_routes

<a id="canonical-d8909d6c566fa4726f8eabe55809bd26e3c0b3ffaaf94b888542f82fb96e2987"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_v6_static_routes = {}
```

<a id="canonical-28e9b9f2bb8f586cf04158f755d90a952a836d6fd9e23a8c6d5b3fc9fea49a0c"></a>

## Direct properties — segment_vrf.segment_config.no_v6_static_routes / f2aa3b544fc3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aeb6db59e06f409606b91cb5472a0514bd09deef1e26af46deda4a30fb6b6ea0"></a>

## Next pages — segment_vrf.segment_config.no_v6_static_routes / f2aa3b544fc3 / 4

- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88e05b7cb1b5eae4b693cc0861ffa52a9badb69cf7a6d286eb55cdcbd5e1cb96"></a>

## segment_vrf.segment_config.static_routes — segment_vrf.segment_config.static_routes / 156a777e867b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- segment_vrf.segment_config.static_routes

<a id="canonical-d278929996f122cdf39835634add363ba29cc03ce1adcb1fdce9823a5389303e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-83232d4fce7ada9287d1355978bb0256abc87aee1b65943e5ddf0db938d08ae7"></a>

## Direct properties — segment_vrf.segment_config.static_routes / 156a777e867b / 3

- [static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5): complete subsection reference.

<a id="canonical-93a8a14670b5cda41386e1e7eacfae475af968bd1601eb3cc97b63f03965fae6"></a>

## Next pages — segment_vrf.segment_config.static_routes / 156a777e867b / 4

- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-479d685ac82c71a9d7efd0558161ac104ba1cd83aec84819dbc6b619739d8747"></a>

## segment_vrf.segment_config.static_routes.static_routes — segment_vrf.segment_config.static_routes.static_routes / ba00c067ef66 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97)
- segment_vrf.segment_config.static_routes.static_routes

<a id="canonical-b0af8c4e66d7a7c3713d57fd051a53abd32adf8c21f6111d4df94b300d97bf57"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

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

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-c45779ec132c028763a8b1187f932f57973181fe58aca9f448e1641abf426ea7"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes / ba00c067ef66 / 3

<a id="canonical-c7558144bffcf08d0a2e9742e2d4b609bdb39950dbc2ecd6b999a2b077716e71"></a>

<a id="canonical-18b8bb55287b31fd4ccf758c7120463bc7a5537e58933e709d982cf5226002bd"></a>

## attrs property — segment_vrf.segment_config.static_routes.static_routes / ba00c067ef66 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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

- [default_gateway](resources--securemesh_site_v2--reference--group-016.md#canonical-e664e76cca1396c0367097029f6d223ef5bc8431de3592b892115d6c11031f5c): complete subsection reference.

<a id="canonical-9cc7eeb883aba3da24c6eaf119b6d4a21ee3d6e3e9d1e445dcb1ce92af10d4cc"></a>

<a id="canonical-de5bc553e3434bb1ab92bb89efc0b476084f03e3b713497b3cc86f0488545932"></a>

## ip_address property — segment_vrf.segment_config.static_routes.static_routes / ba00c067ef66 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-a7771b8d9f27792634075dfbe18c7f6e72e0219f03694771e19a3e1139eea24c"></a>

<a id="canonical-2e06fc5b262a6e2290bc6b1dd32fe500e27f9fb3481f4a208ac625fbf4c250d1"></a>

## ip_prefixes property — segment_vrf.segment_config.static_routes.static_routes / ba00c067ef66 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

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

- [node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-932d6357b451862caad39347509ec03fa09aa90e6605d978f83c6e6af77c55cb): complete subsection reference.

<a id="canonical-24fb5b8b6c565c39f84336b8d3f644a0173c7e809fd179a6cea1b44f6a2fd945"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes / ba00c067ef66 / 7

- [segment_vrf.segment_config.static_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-016.md#canonical-e664e76cca1396c0367097029f6d223ef5bc8431de3592b892115d6c11031f5c)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-932d6357b451862caad39347509ec03fa09aa90e6605d978f83c6e6af77c55cb)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e664e76cca1396c0367097029f6d223ef5bc8431de3592b892115d6c11031f5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63f909c05990a15acfff4330e095575ea328a9aeec7628d098889d2458253ad8"></a>

## segment_vrf.segment_config.static_routes.static_routes.default_gateway — segment_vrf.segment_config.static_routes.static_routes.default_gateway / 162b6cd46e0e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5)
- segment_vrf.segment_config.static_routes.static_routes.default_gateway

<a id="canonical-276a6909b55a9b91c6f574b2a6385f7846f196df4d6269527a3f52fb2dbd5322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_gateway = {}
```

<a id="canonical-97ec44edadb66e30809b14c7f14115e5e26f2dc7c1e87d6de498b2bb80111fa8"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.default_gateway / 162b6cd46e0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1e904430c9f8a346eaa742372d403e2de6ae645596d579be2f1b491314f4c9c"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.default_gateway / 162b6cd46e0e / 4

- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-932d6357b451862caad39347509ec03fa09aa90e6605d978f83c6e6af77c55cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-212b4596b994f84044c792daa711e9dfbec41ab588f9b4f41102d0647fe66de1"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface — segment_vrf.segment_config.static_routes.static_routes.node_interface / aea809399b18 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5)
- segment_vrf.segment_config.static_routes.static_routes.node_interface

<a id="canonical-788e92f00935d09249d430c8fbec59614966421f5d56179b0ef1b28e6ac71f74"></a>

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

<a id="canonical-3b21167b762290ead92d41628c1d8f6b4ecf6bd721415f44d3ea86597e323eaf"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.node_interface / aea809399b18 / 3

- [list](resources--securemesh_site_v2--reference--group-016.md#canonical-96ef3758683fc8edf381c75d24dd8d3e5bcf05a9fe4fc77df97bc4a501e39dd2): complete subsection reference.

<a id="canonical-f813f2b547b67444f25a7e432baece9e47451f90bcc51bcb9b88942be448b772"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.node_interface / aea809399b18 / 4

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-016.md#canonical-96ef3758683fc8edf381c75d24dd8d3e5bcf05a9fe4fc77df97bc4a501e39dd2)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-96ef3758683fc8edf381c75d24dd8d3e5bcf05a9fe4fc77df97bc4a501e39dd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a847c49f16da1bb9b1741a6558e61136753e23528a6c6e7c3bf6aeff2d1f83d8"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / a0b7807f3437 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-932d6357b451862caad39347509ec03fa09aa90e6605d978f83c6e6af77c55cb)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list

<a id="canonical-93cee766c333d594f022c8ab15f0f264d17728214b48e061cab89827e6fcd0a4"></a>

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

<a id="canonical-e467e53ce5a2ea064b33d1183cfde6550f61d89f744fbcbcb9f03c0393db4bbf"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / a0b7807f3437 / 3

- [interface](resources--securemesh_site_v2--reference--group-016.md#canonical-5bbaf3fc64f368cb1a5207ea362c6df2ff6fd7a98d7c7315d09458f7ded1e4ce): complete subsection reference.

<a id="canonical-d101883ce10b13bc73ec67c514e89fcc0374d01a23df53210a66eb3d9c2ead26"></a>

<a id="canonical-5b43bbced919f2435ef69ee3bc21b38b7e49a5ede0a50b2073d7566e01620ae6"></a>

## node property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / a0b7807f3437 / 4

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

<a id="canonical-47243293b8bbfd298c6c07197ddf83cb1b9d339f7f9950f98b5521fbd11e0cb0"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.node_interface.list / a0b7807f3437 / 5

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-016.md#canonical-5bbaf3fc64f368cb1a5207ea362c6df2ff6fd7a98d7c7315d09458f7ded1e4ce)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-932d6357b451862caad39347509ec03fa09aa90e6605d978f83c6e6af77c55cb)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5bbaf3fc64f368cb1a5207ea362c6df2ff6fd7a98d7c7315d09458f7ded1e4ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-993b2e4908069ac8c08cb7b06012d2b9f416403532c3ed59e1f9d4e637667142"></a>

## segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-d8c6246eac4d04f19974bbc2184b1069abb22738ea76ca4d68c1b20402e21d97)
- [segment_vrf.segment_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-bc4fa8cd134d409ee3553b9b3275f9623d87f5ea8c45c7cc738d44976db158b5)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-932d6357b451862caad39347509ec03fa09aa90e6605d978f83c6e6af77c55cb)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-016.md#canonical-96ef3758683fc8edf381c75d24dd8d3e5bcf05a9fe4fc77df97bc4a501e39dd2)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-ca9686618a1d17f16f87fdd81996c19b18a3f301beba430fc49bb5ccf2a38bc8"></a>

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

<a id="canonical-5d36d83ffe1eb9521400b2c4611d759b8103e532820d022fad856bf4bcb7ffa4"></a>

## Direct properties — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 3

<a id="canonical-cc34f15a3db5da4e68d4075ca786820c9df03526f4bbd84e1af3eaba4d1036ef"></a>

<a id="canonical-8ccfb554800ed7c438783bee74f8e5dc4440b314ccc9f70f4499061ab85f22fc"></a>

## kind property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 4

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

<a id="canonical-e5b8729d5c126bec95394898d6e29a48c9895c31738ad5c0b8f3eb2ed0027ca6"></a>

<a id="canonical-13f3f6e37dc0ff82d4cabc75ece4a444175e11f17ad2123eb91459cb890ae3af"></a>

## name property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 5

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

<a id="canonical-62bd7ae6658d6420548ae70c3b596b7cfa8104130821265807f58dc7819e5f4b"></a>

<a id="canonical-47cf9ca08029130a9612c7f0f580281b317cec5e9ade15e8b90b1f4906e49b4e"></a>

## namespace property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 6

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

<a id="canonical-70e76bd52c29d0e2b2741ca685274541c4e7f4bb734fafe73ab3ae4f540b03c4"></a>

<a id="canonical-6569bd8584a5246edec382b63cc325bb1527278e5fe8fdc7c79abdf8d7805bf7"></a>

## tenant property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 7

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

<a id="canonical-45b5fde696852a81f915d6ff2a585a47f7a7f122e7a73cbd939250283e511651"></a>

<a id="canonical-9e84451ba64f8c6dad23ec5fc6ea5cc05ce1cdc83c6bf00dc7bf89f5b1f0f711"></a>

## uid property — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 8

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

<a id="canonical-32e41d41395a6dbe513cdafc7936bb9b061310d241d71ac65be282ae2ba934f1"></a>

## Next pages — segment_vrf.segment_config.static_routes.static_routes.node_interface.list.inter / 12fc26fdf318 / 9

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-016.md#canonical-96ef3758683fc8edf381c75d24dd8d3e5bcf05a9fe4fc77df97bc4a501e39dd2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
