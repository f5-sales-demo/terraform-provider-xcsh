---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-d52e452de5e1131b9883bf5aa7bb341b7a848c9769290a9951fc943c91899aed"></a>

## configured_address property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 92eb471ee230 / 4

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

- [first_address](resources--securemesh_site_v2--reference--group-011.md#canonical-1be2cea4f277343c1702884b696ea2fb50fb61eacc99a74a24c9b08977df1975): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-011.md#canonical-5d311f00080b64200cf093332f957b14a8e8c136aa3a82656b9781a3148decf8): complete subsection reference.

<a id="canonical-9cee98dcba9b7da4689d90a97e6f5e80f6f1330eb0c4e281d0194d8830dd74c3"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 92eb471ee230 / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-011.md#canonical-1be2cea4f277343c1702884b696ea2fb50fb61eacc99a74a24c9b08977df1975)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-011.md#canonical-5d311f00080b64200cf093332f957b14a8e8c136aa3a82656b9781a3148decf8)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1be2cea4f277343c1702884b696ea2fb50fb61eacc99a74a24c9b08977df1975"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21fcaa4f6eb6e27eb4fb8c17758544537ac0ffca68c6128f4bba648f1d96b539"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 0501fa841033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-5915cec528cdf8614eb1a1532ec9a6b289784ec0229a0f0dc88cc20c3831517e)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-b368441ba7b2b5b580fd93628980527ff9cd763b5f4c446ba0a138315b7d450e"></a>

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

<a id="canonical-f544ce707336133da7d094cbe4224d9bf81e4d8d9b9e4dd15b125b065efc7273"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 0501fa841033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e28f10100c5746821822c2a39e0bae323c5e95f335d900e15c90d8dd0956f214"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 0501fa841033 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-5915cec528cdf8614eb1a1532ec9a6b289784ec0229a0f0dc88cc20c3831517e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5d311f00080b64200cf093332f957b14a8e8c136aa3a82656b9781a3148decf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dbd67f44344dd884b901627a577dbfdfcf6347f140eaf41ea7c7da212e7df49"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / b4f61d3415e4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-5915cec528cdf8614eb1a1532ec9a6b289784ec0229a0f0dc88cc20c3831517e)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-572437b2c7aa93d5f0360b1b545581b221d02f2decb518c60c66384d1b7ade6d"></a>

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

<a id="canonical-11386653946229a18b7f5833112f768d1b1a35f18f29b8fbb059b14cf797ff99"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / b4f61d3415e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c4851e7f644fe9fc3f588028bebd08c6ed63c2a4adfd6e4d1929d42921da979"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / b4f61d3415e4 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-5915cec528cdf8614eb1a1532ec9a6b289784ec0229a0f0dc88cc20c3831517e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab9921eaf9db0e13978d47121fe62065fd412f7e001b6b740a2f9241723fad97"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 8423b77f1313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0b68674bb8cceacf03ac6f67442194c6f7efd3cc7065fe914b24f008f677a71b"></a>

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

<a id="canonical-d9bccb7755cc72cc28d8bf1a84a76ec6758656d95fc3336411bcd35b3a57471c"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 8423b77f1313 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-011.md#canonical-7645e43d6b6f33ecbb2de72455716cfda727c2f18b2c89548c2ce7b71bf9ae6f): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-011.md#canonical-0b11fff3f2d1c68a2059339e4d6cdc722eb8145397513ce16558474a5c3197a2): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-011.md#canonical-6dc3421a3525c824e5dbf3d58fcaf263aac401f61a426c9572010ac04b45714a): complete subsection reference.

<a id="canonical-0bdac60db6683b4e6c861d39e8c47391f4fb0d039b4e0a10aa8845d48e00c1bb"></a>

<a id="canonical-43b8e7509411ee00c45102a1795e2d4c1521926cdcde52250a548266370cf3dc"></a>

## fixed_ip_map property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 8423b77f1313 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-011.md#canonical-e9d399adf333d4b3e1e9e77e5fc7b48ceae8632b92bfaa6d97713732cd02dfc9): complete subsection reference.

<a id="canonical-92ba5a366171b3a97162332d3a83c06a51b72da0b1520c2b396f01d6a555a5f3"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 8423b77f1313 / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-011.md#canonical-7645e43d6b6f33ecbb2de72455716cfda727c2f18b2c89548c2ce7b71bf9ae6f)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-011.md#canonical-0b11fff3f2d1c68a2059339e4d6cdc722eb8145397513ce16558474a5c3197a2)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-011.md#canonical-6dc3421a3525c824e5dbf3d58fcaf263aac401f61a426c9572010ac04b45714a)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-011.md#canonical-e9d399adf333d4b3e1e9e77e5fc7b48ceae8632b92bfaa6d97713732cd02dfc9)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7645e43d6b6f33ecbb2de72455716cfda727c2f18b2c89548c2ce7b71bf9ae6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18cde4c3d7f55960b174193e115eb613724168a6b44541b9a969ddcbaf08de27"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 7e55bcc9c0b9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-ad2272de6cd2fd5ce55eb81d1ff1e136f79c166f653fae1ecd02fcb551d61d4f"></a>

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

<a id="canonical-27a602bc7471eabb1a4a92e717c6d21b68f4a71a0dd8bbf37ad5cc5bce951816"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 7e55bcc9c0b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6f7da0dd9546ee94bfcd8480e3cb19b2fc458d71b348459d7fb009daadbb13b"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / 7e55bcc9c0b9 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0b11fff3f2d1c68a2059339e4d6cdc722eb8145397513ce16558474a5c3197a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20a475dc3c2a63e08076a1f24858ba51c4e90485de49488bfae9dd1e3366012b"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / db546c70c8f1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-2f99d8a54a7fddfe2ff5cdc1162c8470c1ec3f0adf590182e6561e6c494d4339"></a>

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

<a id="canonical-2b2f644d8d48bc4b673e9bbcd12db571ec103602deb707a3d9f3293dd54da903"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / db546c70c8f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-accc4ffab2947dd046628835eae06ac46c2c47be80fe0572b10ae40b6b511430"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / db546c70c8f1 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6dc3421a3525c824e5dbf3d58fcaf263aac401f61a426c9572010ac04b45714a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5a0cd979050d34271a7b01040a8a60010a491f5a5e75dd82b5cff09976eac7a"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / cf1b6e0e2915 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-31945a4abb52a88bf58875878eeb25b8f80a10515bc71a03e849cc3677e53076"></a>

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

<a id="canonical-2e30787a725324d86011397d2be3f0346a5c77dc927e872d22ba2462f68f25e9"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / cf1b6e0e2915 / 3

<a id="canonical-ff50aa80d19d36f6aecdf859768d2804d1eb4a00dc6b362419a7d03fcc56db02"></a>

<a id="canonical-fd75834313dbeadba20b2f0e7ff295cc3a9652b84f05f8e29c5d2629d221e695"></a>

## network_prefix property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / cf1b6e0e2915 / 4

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

<a id="canonical-d3190032e1405d18d69a71544cbb57a034b8d7258ae22c2ba33157621dba37eb"></a>

<a id="canonical-64061bbcfcc6dc08f1d41d9947c7e32762e58fae2e723fbf904f000990ef40ef"></a>

## pool_settings property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / cf1b6e0e2915 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-011.md#canonical-f8a3c9075bb081bb60b783e8c827b00ef5de90aa9055b73a31312ddbf4528052): complete subsection reference.

<a id="canonical-80c971b8ef61700a937ad7cf282ae2d0e638c2697d6447387b3ab473df6498ff"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / cf1b6e0e2915 / 6

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-011.md#canonical-f8a3c9075bb081bb60b783e8c827b00ef5de90aa9055b73a31312ddbf4528052)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f8a3c9075bb081bb60b783e8c827b00ef5de90aa9055b73a31312ddbf4528052"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bba19564b172c0faa7dec700577c4a5bf39676b28015863d6c12121ca11c7421"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / af125e39ec4c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-011.md#canonical-6dc3421a3525c824e5dbf3d58fcaf263aac401f61a426c9572010ac04b45714a)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-e5e9c2434fdccb8037be5147b107b4c5aac5f46b144f5d2028951e4eb5d7ea3a"></a>

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

<a id="canonical-b5547f124fa502226d3beb310974575cdf1020a2f941829451b74ae2b06afc1d"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / af125e39ec4c / 3

<a id="canonical-0e7ebe4a55b1fe8606ea1c2c0cac179efe05eab48cdeab967170b92b3653aea8"></a>

<a id="canonical-b68510990b7635dcfa8fd8da8adadef36838e8454b39b3a66188729df07dec96"></a>

## end_ip property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / af125e39ec4c / 4

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

<a id="canonical-b9f6e4400793e4313eb82ad3a4fd21dee5f26a952676fac9cbd05315f839a069"></a>

<a id="canonical-a3b00d8b5a8c56db364d8560d8a4db71b8d793aa2489cf2cf4c90b03a87b5715"></a>

## start_ip property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / af125e39ec4c / 5

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

<a id="canonical-2be7245c6b1570a1feab2d62668b5a49b8b3c2b0bc265f2b6291874cc883d03c"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / af125e39ec4c / 6

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-011.md#canonical-6dc3421a3525c824e5dbf3d58fcaf263aac401f61a426c9572010ac04b45714a)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e9d399adf333d4b3e1e9e77e5fc7b48ceae8632b92bfaa6d97713732cd02dfc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89d79ecaa96c67e6ddfb40a7800fafa2ee3a9d27ba1ecc0766ceb35398627105"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / ccf38bc7c63d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-d118d97255149c49326946beea1b63c23e13e683e9205c07fa56275cf5eeba36"></a>

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

<a id="canonical-449ac23b995ab3065c0d8af237160e9df2f394c40340368d295005603019ea38"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / ccf38bc7c63d / 3

<a id="canonical-ed064195bd0edfed08a400f58a19cde3c824e3d45046ef6b78a974182ae10baa"></a>

<a id="canonical-658f5aa7785a16b08cbbc4f03084097ecfdc650a93351b5ae2aa72bf7603cb24"></a>

## interface_ip_map property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / ccf38bc7c63d / 4

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

<a id="canonical-f1d89c15dff7b6d921945c703c3a8fe0142756591ce9955b0ad088b224565c68"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / ccf38bc7c63d / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-86d48cf74cb6ec20a931bbfb5aa21353bea131bd30bbe85fb0b609a188ff8267"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4481ef8b4b8e2a7ddeb2f4c3f66bd7c6d28262cb222cd830c252d62775ac9180"></a>

## kvm.not_managed.node_list.interface_list.monitor — kvm.not_managed.node_list.interface_list.monitor / c96eba6dd2af / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.monitor

<a id="canonical-adf214894567fc8cd65115542d4cbd8262ce7b86895e295170e6e9f6ab11ec38"></a>

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

<a id="canonical-f2d34d6119c6172f2b08e902fb5d40259ebc9f76a48262b8056f0c8999d85d7d"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.monitor / c96eba6dd2af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71bd269690b27abe53bf4a9616bbb4951425173815fcbc5ae317cdf6ba58e82f"></a>

## Next pages — kvm.not_managed.node_list.interface_list.monitor / c96eba6dd2af / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5089a3c8821ddebedbd32922dfc7b09737445d72d94eafa220b5d5fd99af019a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e4c5bbfc6885945cf5d419075c7047373d810c6b48cdfb80fde16f7823f342b"></a>

## kvm.not_managed.node_list.interface_list.monitor_disabled — kvm.not_managed.node_list.interface_list.monitor_disabled / 13f09562c712 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-62379f88b6fd6b09d91fdc3b4d9f80686edfb3c0d645c8f83ac781e5c61141e8"></a>

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

<a id="canonical-943162e2e065a82d02196264fef5cff029bf06fb72c537b2eaf1301f6d385835"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.monitor_disabled / 13f09562c712 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-453210ce8f4d36aec689e51a2eb3adca454c65c6a01902ef9c99f467b22ea033"></a>

## Next pages — kvm.not_managed.node_list.interface_list.monitor_disabled / 13f09562c712 / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-95195c28512bf633641d2a35dcb1fc4bac024249c31ef92cebebc0959b739017"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-723fec3cac56e3528c6aac8db286c08ec1568444abf63019315bc0d711cf7b3c"></a>

## kvm.not_managed.node_list.interface_list.network_option — kvm.not_managed.node_list.interface_list.network_option / 35cb5d526999 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.network_option

<a id="canonical-d7aca628ac1ad66cde2f7e10bfd06944aadd2dc7e355a214cac35f7a1f9b7526"></a>

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

<a id="canonical-ce4be73af765502a0a0edfc2ae2658534c5a060fc603da4bd09f5ac16bfb42b4"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.network_option / 35cb5d526999 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-011.md#canonical-9279ade966fcf8f8625bcba9f520ec8f28dcca287c973c2047cc24a029ccfb9d): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-011.md#canonical-f230ccfa6e810b512bdf34761489587d1036071c776ee3bcf217454dd363656d): complete subsection reference.

<a id="canonical-2b50598874ad81bcba4a12da979c6631664ebc4bd839f146c72feda82c77b5f5"></a>

## Next pages — kvm.not_managed.node_list.interface_list.network_option / 35cb5d526999 / 4

- [kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-011.md#canonical-9279ade966fcf8f8625bcba9f520ec8f28dcca287c973c2047cc24a029ccfb9d)
- [kvm.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-011.md#canonical-f230ccfa6e810b512bdf34761489587d1036071c776ee3bcf217454dd363656d)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9279ade966fcf8f8625bcba9f520ec8f28dcca287c973c2047cc24a029ccfb9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c54b4eee68acf5d665879e7855dea7b11b77a7a77138c0d8d02d2b66d4edec23"></a>

## kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network — kvm.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 0e65c00a8972 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-95195c28512bf633641d2a35dcb1fc4bac024249c31ef92cebebc0959b739017)
- kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-4a6e155268e30592b0b80f061775a56594efc87d868b72bd9906225cbd22e165"></a>

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

<a id="canonical-9d7d44bc480a9dbc4ed30bfa321541c24aa173be0590ae46be6ddc080b33a8d8"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 0e65c00a8972 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1fd0a411efe9911ca1559c3390f94f803678265fbc8cced9ad092bd9ff79deb"></a>

## Next pages — kvm.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 0e65c00a8972 / 4

- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-95195c28512bf633641d2a35dcb1fc4bac024249c31ef92cebebc0959b739017)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f230ccfa6e810b512bdf34761489587d1036071c776ee3bcf217454dd363656d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b89152d535cf70d2e392006bc8aa30e29e4e2057321a973bb32663eb9315a245"></a>

## kvm.not_managed.node_list.interface_list.network_option.site_local_network — kvm.not_managed.node_list.interface_list.network_option.site_local_network / ea62af2ff105 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-95195c28512bf633641d2a35dcb1fc4bac024249c31ef92cebebc0959b739017)
- kvm.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-3789a8257a73124efac71d39b98d151f8308a3b05b7274c2c217d130e8b42b2f"></a>

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

<a id="canonical-de2da460c8c019cf6840ccaa54f936303c3e108a9c832fb7955dbbb3574de0a9"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.network_option.site_local_network / ea62af2ff105 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4b08edf417ceb735d4b5fabe443d75dacdbc0a6109bb01970e8606c428e2c07"></a>

## Next pages — kvm.not_managed.node_list.interface_list.network_option.site_local_network / ea62af2ff105 / 4

- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-95195c28512bf633641d2a35dcb1fc4bac024249c31ef92cebebc0959b739017)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-db75d8f8359ee4b155a49ace03af36dbe9bed4e92bc51592bbd0065a32feaee1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-444c594b8491aeb752021947f8f39448c936c51338ef064d6946a8b8ec6a2ac3"></a>

## kvm.not_managed.node_list.interface_list.no_ipv4_address — kvm.not_managed.node_list.interface_list.no_ipv4_address / ddd126b001a1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-ea9140e929a3943ea47fdc0c945e7442f91458977192641db564df0e150cc0ad"></a>

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

<a id="canonical-8c953738fed57a8e4ee7c6bb6e58f463b9c27d42edb2963508047cbf7052ecd1"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.no_ipv4_address / ddd126b001a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15f3b0341fcb804d375f25a8e731d77423b08f134f793bcb95982e690404210e"></a>

## Next pages — kvm.not_managed.node_list.interface_list.no_ipv4_address / ddd126b001a1 / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-35113afa738d3ff3faf3314fd5632e2c501dc921050553ef56458426a2fb7c91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23518af31096dde51fe518244b63d56f49ee9dfe49b76cfe13dca7ee051993e7"></a>

## kvm.not_managed.node_list.interface_list.no_ipv6_address — kvm.not_managed.node_list.interface_list.no_ipv6_address / 48fd89d6bb20 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-1826a5f53c7844403699cd4038ff6883e21a00c94c11241299894c6e31e67f64"></a>

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

<a id="canonical-1d357d63889a1a62efa5341c21fd5793d3e829fa0652ca48cb203ad23f3fb7c4"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.no_ipv6_address / 48fd89d6bb20 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b44a1a564b1fa4a6ed40549d7911296c06285810e8e42e814ba6ccd4360209b"></a>

## Next pages — kvm.not_managed.node_list.interface_list.no_ipv6_address / 48fd89d6bb20 / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fed971e3f1840028a982dba6a3186b22f92e720e2f8464631ef029a5c7557a2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-027fb0a385c6d8d496e623b6766fa567a3c8299ea96a768ce7ec1535acd18fd2"></a>

## kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / 73642fd5da65 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-09ccf2564986f449212711677f625f2771d191af84ca8fdcfdee291a6bcfe388"></a>

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

<a id="canonical-f48548314b5a346b01f17c4a52f013ae2b22e853bb36d732ddab32d80a1943de"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / 73642fd5da65 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-567783e05b0fac474253c251f68c7c20582419ef8141e2686f943e35a8926cbb"></a>

## Next pages — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / 73642fd5da65 / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c739f48589c4e0507b72d4974b60e5a2462c1bfee830f5dae826c634cd74a755"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89e47d39f1f47346b6b9a354821854f5ce0bf595de9202ee0243efb56ffaf311"></a>

## kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / 478beb73d2d9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-eb35fc40231262f737fba381840498e06d22f9093c42b313318d1d8ecd461d08"></a>

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

<a id="canonical-195a3e5c0565e63971ac4e11ef23d16c3d4548fb0bbe7abfd30f71d875c9d45f"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / 478beb73d2d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ba593b4dec947b60771ae1566193032607bfa1c80c3917ab04b688d7b94514f"></a>

## Next pages — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / 478beb73d2d9 / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1d904574b9d982a55620099dec99e0eb20a899e6103c21e86a73a891796659cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c08c4c3fff47b9fa584a411a2d71b9388040658505e6363a4b1c0ad465f9977"></a>

## kvm.not_managed.node_list.interface_list.static_ip — kvm.not_managed.node_list.interface_list.static_ip / 7cbb26904739 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.static_ip

<a id="canonical-711a3326a7947ab4a867cfe04077ce9a0591e2ab5773851c1f7a96f50e2a197f"></a>

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

<a id="canonical-58f9323749623fab04973eae75b84322bf0c1be4bce98a80af7d7a083bd9df03"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ip / 7cbb26904739 / 3

<a id="canonical-b0d0fca98e3c4bdb070234c9e54ff173701a72ea47a62d50ec5711eb33db2c6f"></a>

<a id="canonical-7e1e4b62d5cc761e23d5bd0eb35ce9864d1552d1978363ef56ac098ddd22b984"></a>

## default_gw property — kvm.not_managed.node_list.interface_list.static_ip / 7cbb26904739 / 4

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

<a id="canonical-d72399050245e6ebc48381e876cc6289e2e48c92f58f70e5837e870951c4566c"></a>

<a id="canonical-7cdc674783c8f71675beb914b6315fbb55ab0cc6c5ca1cc641b4dfd79d9e8183"></a>

## dns_server property — kvm.not_managed.node_list.interface_list.static_ip / 7cbb26904739 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-ecdb16a94b9c806a93a961cedc6ee23fc78979a7e5c7eaf293c2de813374738c"></a>

<a id="canonical-cf18e467dd7758e7af8f34fd69e25c30d5f6e7ec1f31efcfb2707cc71152046b"></a>

## ip_address property — kvm.not_managed.node_list.interface_list.static_ip / 7cbb26904739 / 6

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

<a id="canonical-52ddd589a0573d01c5a3d45832f46ebb02d1266484839ecdb7b3ce9575abda48"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ip / 7cbb26904739 / 7

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bfe28f2aebf796d26c8e6026a08039c4b9c5bbe882afd60bce2e1bc7b72e47f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3addb59e44535b76d6a8fc55df30b335180edfb7e22e0df4442f691b6bc6ffc4"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address — kvm.not_managed.node_list.interface_list.static_ipv6_address / 9dbb4426dce4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-9310e96d2191598872ef5edb06fbe3229c4ce2e617c24ce66f0a908d5eae88b6"></a>

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

<a id="canonical-4d85ee519fe08eaf64d7c1f6449586e571d856b94dfcd6c765bbfc067c3a2ef1"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ipv6_address / 9dbb4426dce4 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-db38b5280a9331d156aae0ee2986e7f5ddaa15f9c6d3281fc38d9ada633203de): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-eefe79d442d93434e49c91964b41f709afff985ef316def7a7fc911e47d9f637): complete subsection reference.

<a id="canonical-fda3ab929a877dcc44b4c109a8d54e7e5c2defc65187ef0c925b62ee8a298c90"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ipv6_address / 9dbb4426dce4 / 4

- [kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-db38b5280a9331d156aae0ee2986e7f5ddaa15f9c6d3281fc38d9ada633203de)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-eefe79d442d93434e49c91964b41f709afff985ef316def7a7fc911e47d9f637)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-db38b5280a9331d156aae0ee2986e7f5ddaa15f9c6d3281fc38d9ada633203de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87f7400e31ced924e65194fa4ec3a5148d4f560962f0d636c27774e747e347f7"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 9c163318aa18 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-bfe28f2aebf796d26c8e6026a08039c4b9c5bbe882afd60bce2e1bc7b72e47f9)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-bb02df49db548f464cb7b2d19f543ac636e88ab89c133dcc0b512c5298586eb0"></a>

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

<a id="canonical-33e27f6e65c273bf387efbdd8d0daa80d32ba3ae847023c38aea2a64c69b4149"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 9c163318aa18 / 3

<a id="canonical-dc29e3aab1adc9687c405a37210543f69a9a3183d2e2c127b6a6a3184f2e2189"></a>

<a id="canonical-84cd56588a7df0d8344626e238b7a5f7075fd308206c3aba369832f620d1b5a5"></a>

## interface_ip_map property — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 9c163318aa18 / 4

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

<a id="canonical-e780b620e8035e6196bdd5d1fc75d25ca15bad8c9062ce4aefa72940e56f7629"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 9c163318aa18 / 5

- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-bfe28f2aebf796d26c8e6026a08039c4b9c5bbe882afd60bce2e1bc7b72e47f9)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-eefe79d442d93434e49c91964b41f709afff985ef316def7a7fc911e47d9f637"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af26ab011ec8855ea553bfa530442f882be6a69c44d149e9482b8a7be07848d1"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / de0c3dec5c92 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-bfe28f2aebf796d26c8e6026a08039c4b9c5bbe882afd60bce2e1bc7b72e47f9)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-514aeb9cdaa95ba92a844efae569525d32d963b151cfcd8f9d6e826ad52282dc"></a>

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

<a id="canonical-726e20f6638b1c388f1be08a798e21b22de777d9d2260a47fc89b64dc144ff03"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / de0c3dec5c92 / 3

<a id="canonical-be7dd2bee7b3dd92117884ec9144781cae01e5df724dec445c9e10346c9998ae"></a>

<a id="canonical-cfcd1743e50a66f185c5b2f930a04f4935816a4db79ec8e99338cd818bd30877"></a>

## default_gw property — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / de0c3dec5c92 / 4

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

<a id="canonical-fb77283951b679e33e2bc4f6fc28c572ce7fd9928ea1e2bb4a32c131f1919657"></a>

<a id="canonical-1de0bcc79c746760c7a6802b42ff7207369865f9396202170747db55f2187b0e"></a>

## dns_server property — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / de0c3dec5c92 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-88f5ab4bcc764db7f9773dcbf2f19e89fc2f393d38c46f70429656349123b12b"></a>

<a id="canonical-af972c27c518dd724c439c9f0ee53eec25713fffcd37c95d8adc5cc6d77f6ce9"></a>

## ip_address property — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / de0c3dec5c92 / 6

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

<a id="canonical-ce116744e9b939159ed4e8f6db13aa54220ea23ae406fed8143a3be7a6345280"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / de0c3dec5c92 / 7

- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-bfe28f2aebf796d26c8e6026a08039c4b9c5bbe882afd60bce2e1bc7b72e47f9)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6336643535d3cc49ea25eda2a16890f63b134e496b9b9d9c3aff069677cb615b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6db28795ed1afd5659beaf8dce177ba5d0e01cf002164ff024ce9d114cd93711"></a>

## kvm.not_managed.node_list.interface_list.vlan_interface — kvm.not_managed.node_list.interface_list.vlan_interface / 211a07322feb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-7e569110a6ef707934d7b69e47d16ea0640693444ec4edb74f104868717fc674"></a>

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

<a id="canonical-5ab6de6ca1b1b1989a7ef61062498978cf502fac0dfab9eea96a3ea96b20883d"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.vlan_interface / 211a07322feb / 3

<a id="canonical-6c4a118005d0b7035fcdb775049486447d4695710a65eb4a5782e82c43212c53"></a>

<a id="canonical-414fed2bc9ddbd4b961d47f23ce2c16935fec26c0c8829989a7c66026250fa3e"></a>

## device property — kvm.not_managed.node_list.interface_list.vlan_interface / 211a07322feb / 4

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

<a id="canonical-71ad38eb1f3b7519419629296c9e83438529273524d11b453e5985dfc4a71b2e"></a>

<a id="canonical-4607c179b766ac353be8b8d202fa3abb44fb9b7349dac627ff0327de12412903"></a>

## vlan_id property — kvm.not_managed.node_list.interface_list.vlan_interface / 211a07322feb / 5

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

<a id="canonical-2224ff00435ebdb98124a1b5048c5e5635f464bba8b5e868d42a507414ec6b86"></a>

## Next pages — kvm.not_managed.node_list.interface_list.vlan_interface / 211a07322feb / 6

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-872e4a68860accaf3e925d83a4ce2b9691840fd9d1ff4d1b969cea185980d877"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15cc827a606ddbec8cf86f4e461f49e74bd1678e6f0d841a0d350a0d3801e102"></a>

## load_balancing — load_balancing / 3f0b24230583 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- load_balancing

<a id="canonical-505c01b8b4809ee18d3fc000ede5397f3f911128b29c9deb489ceae51990739e"></a>

Type: `"object"`. single nested block, Optional.

Section contains settings on the site that relate to Load Balancing functionality.

Upstream description:

This section contains settings on the site that relate to Load Balancing functionality.

Receipt-pinned upstream constraints:

```json
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
load_balancing {
  # Configure direct properties listed below.
}
```

<a id="canonical-c9371faf61d55d0726ef17ee285c157fe843a86610cc8d0ad8446bd591937514"></a>

## Direct properties — load_balancing / 3f0b24230583 / 3

<a id="canonical-cd0da131be21127d2105d05ac066a86503e1956c984c5eed6176c18dd52e145e"></a>

<a id="canonical-30c3bb20357588ba8067c1fa9ed7d2bdf76d4eb57f9d654f0d44bdd28aed9efd"></a>

## vip_vrrp_mode property — load_balancing / 3f0b24230583 / 4

Type: `"string"`. Optional.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ed422b88c98aed5f4c351257de9c794b999f976f48c7124284514dfcc107def2"></a>

## Next pages — load_balancing / 3f0b24230583 / 5

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d192856d770531b87023906b67a57ed7c92b9da5cf8059b714da834d922785d6"></a>

## local_vrf — local_vrf / 4b626444b940 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- local_vrf

<a id="canonical-8613f83f90e4a29d2a831182bc0eace694d2bb8591066f20b4663b1a8c1aaf86"></a>

Type: `"object"`. single nested block, Optional.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF.

Upstream description:

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to
connect LAN side workloads to this site. SLI local VRF is optional.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_config",
    "slo_config"),
  validators.ConflictingObjectAttributes("default_sli_config",
    "sli_config")}
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
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

Terraform syntax:

```terraform
local_vrf {
  # Configure direct properties listed below.
}
```

<a id="canonical-81544974a4bf3dc538f3ffafe37631ee1645211e6ca643cc4dc735c63be14e09"></a>

## Direct properties — local_vrf / 4b626444b940 / 3

- [default_config](resources--securemesh_site_v2--reference--group-011.md#canonical-89fefbe48d9a7607e9a3e2a9967ec7b8cd0723b14e3979e0870e7ffe54f3b3cb): complete subsection reference.

- [default_sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-e07efb0b1fe8c78b10536faa3cb920248eaf7b72b0b96000f58073072b63a8c4): complete subsection reference.

- [sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0): complete subsection reference.

- [slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd): complete subsection reference.

<a id="canonical-d17a18838f1690908158256181ad0f0c272717dae2fb8422e6162c0392f37707"></a>

## Next pages — local_vrf / 4b626444b940 / 4

- [local_vrf.default_config](resources--securemesh_site_v2--reference--group-011.md#canonical-89fefbe48d9a7607e9a3e2a9967ec7b8cd0723b14e3979e0870e7ffe54f3b3cb)
- [local_vrf.default_sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-e07efb0b1fe8c78b10536faa3cb920248eaf7b72b0b96000f58073072b63a8c4)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-89fefbe48d9a7607e9a3e2a9967ec7b8cd0723b14e3979e0870e7ffe54f3b3cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da330ba00529762907f65a3e282c2f2b5b18690b873e01f90e106326a592c01c"></a>

## local_vrf.default_config — local_vrf.default_config / 4cd4b5d9b3d5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- local_vrf.default_config

<a id="canonical-70be569ab0c0c13ff7d3b9a6fe83cf9ec6ae5389871321190666bb2cbbbce79a"></a>

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
default_config = {}
```

<a id="canonical-b4ffee4bb464a45256b5a23ea5b8736905dbadfad3904a68637ff1ae846f7f87"></a>

## Direct properties — local_vrf.default_config / 4cd4b5d9b3d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e07a8997c0c5d343193e48dceb527c447f5d276cb46ebfba2396a5d2aad5f921"></a>

## Next pages — local_vrf.default_config / 4cd4b5d9b3d5 / 4

- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e07efb0b1fe8c78b10536faa3cb920248eaf7b72b0b96000f58073072b63a8c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-476bcc570e1f26404955c070f6bb254516392e3591ec53a8caf3d0947d88abfe"></a>

## local_vrf.default_sli_config — local_vrf.default_sli_config / 66a7a419180b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- local_vrf.default_sli_config

<a id="canonical-fffc04a9530d459ef0a89df6b3ec12760bda8fc4021b41b16af9f74809854361"></a>

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
default_sli_config = {}
```

<a id="canonical-c1d320a54f0fc6da652d37754c1ff826f32bc56206e1a844a9192e5878fd3495"></a>

## Direct properties — local_vrf.default_sli_config / 66a7a419180b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-839ec26a8eb81a992c72b355fd073f76ca39e4f7de9251f5fc9ce846c263b20b"></a>

## Next pages — local_vrf.default_sli_config / 66a7a419180b / 4

- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-030cec6a62601c78a53cbf71dca0e2f2ed3a016911971b25a26d7155c89635bd"></a>

## local_vrf.sli_config — local_vrf.sli_config / 0f931a9e876d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- local_vrf.sli_config

<a id="canonical-84161f3821d612756c4a1d9e1d074c692d40521102312d6d2eb2da539cd04247"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

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
sli_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-026392436649dfec3e40eb678a8e2ae2b78c6ed309a555789134b1a3ecbd5a32"></a>

## Direct properties — local_vrf.sli_config / 0f931a9e876d / 3

<a id="canonical-7f238525b3343953cc8290e41c544bb69f9e4b4394e7274da69cba7a08cf31a1"></a>

<a id="canonical-1fd6f06323693eda04fc256247bd6195306f69c55da10c40f505dd90e71fd3c9"></a>

## labels property — local_vrf.sli_config / 0f931a9e876d / 4

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-ef633145ab9fc500c50c67fcf746e3120341bae38f74d33012959e6f00381a60"></a>

<a id="canonical-2c98fe2d3a7cc4a4106a398fc2b85b845f740550c201f42de462d6a26ccc78c5"></a>

## nameserver property — local_vrf.sli_config / 0f931a9e876d / 5

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

- [no_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c02d134dbbc09185721bd351ba6641cef762a387296e0613a951733f6ae14d04): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-321950873b30bf9ab50b63472e3d94311e28b8bbfcd1a708e90ba2c5608ecaa5): complete subsection reference.

<a id="canonical-7a55f99e79fb967966392e378265bc8ea9368bfb5ca2b1b8c412de457657de70"></a>

<a id="canonical-7c0419dc382cec77ec061395546347ca028c9fcb9995caf8c18d6b5816c09e4e"></a>

## secondary_nameserver property — local_vrf.sli_config / 0f931a9e876d / 6

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

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040): complete subsection reference.

<a id="canonical-d522885fb328ae963d385051d71b49d139e7fd86fef50dfaae370ff3341ae321"></a>

<a id="canonical-a59e5d1ba08b18ebb0851910e9bd971384af7112e8ddbe151bf6167fc6130143"></a>

## vip property — local_vrf.sli_config / 0f931a9e876d / 7

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-ed5b8f89462384de099f8ad4fe2a01192896ea155f88e5ffbabacf3cda5d1d7e"></a>

## Next pages — local_vrf.sli_config / 0f931a9e876d / 8

- [local_vrf.sli_config.no_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c02d134dbbc09185721bd351ba6641cef762a387296e0613a951733f6ae14d04)
- [local_vrf.sli_config.no_v6_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-321950873b30bf9ab50b63472e3d94311e28b8bbfcd1a708e90ba2c5608ecaa5)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c02d134dbbc09185721bd351ba6641cef762a387296e0613a951733f6ae14d04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54e3c2959ee981f65f13c30cdf378ecbfe7153e23c5cec08a68617569a0a307c"></a>

## local_vrf.sli_config.no_static_routes — local_vrf.sli_config.no_static_routes / 8c7240e7c4a1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- local_vrf.sli_config.no_static_routes

<a id="canonical-ebc48b7e16c85c5a90cf35b282bdf8f6642edd81babb0728b9f21c89cbd14179"></a>

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

<a id="canonical-6a80366c3b5486cc7dd226c34d6f275ac5d143e80c062006c9c1962ace529ac4"></a>

## Direct properties — local_vrf.sli_config.no_static_routes / 8c7240e7c4a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-caca0ad6d0b81d624efb490208b43c37cfe387dcbe2d78d79c0472b3709b8867"></a>

## Next pages — local_vrf.sli_config.no_static_routes / 8c7240e7c4a1 / 4

- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-321950873b30bf9ab50b63472e3d94311e28b8bbfcd1a708e90ba2c5608ecaa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d94e93cfb9b136b7c7b1f2b1813554a017bd3a1304cdafa9c17c9a09a03f586f"></a>

## local_vrf.sli_config.no_v6_static_routes — local_vrf.sli_config.no_v6_static_routes / 1b74dfc0db21 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- local_vrf.sli_config.no_v6_static_routes

<a id="canonical-5f8f5ad352d793c0fbba3c51772c374509eaf11f7b420553fe816eb04884bbe8"></a>

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

<a id="canonical-efad816e2774894adeea0de9757c2a33948e16294767d2ba4d57b75558322d2f"></a>

## Direct properties — local_vrf.sli_config.no_v6_static_routes / 1b74dfc0db21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d40eafb8da73eb065541cf10d665cf0a2ecc0b7fbcf8b84e1e1820dcaf442a96"></a>

## Next pages — local_vrf.sli_config.no_v6_static_routes / 1b74dfc0db21 / 4

- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b31e36bfdb52cb75ca92adbf4ed72e6dc54af7a145d3af97c607e8e683e7597"></a>

## local_vrf.sli_config.static_routes — local_vrf.sli_config.static_routes / 76effc58081b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- local_vrf.sli_config.static_routes

<a id="canonical-f2c5742576169fb03ebf0b9287ae5b84706bc8d6146bc94e3b61c507333223e7"></a>

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

<a id="canonical-88c25b2660972244699cbf0be1582ea3083d73466f251c8ed90b05906ac1c698"></a>

## Direct properties — local_vrf.sli_config.static_routes / 76effc58081b / 3

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf): complete subsection reference.

<a id="canonical-ae375bc3352c39a187c980f508dc7dc0cd8201c1bb99c6d0694d2be14e3290b5"></a>

## Next pages — local_vrf.sli_config.static_routes / 76effc58081b / 4

- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04f444a32a6a8941980b9a0ba2eacff91de69054b4e5903e9bec65ca3b149d90"></a>

## local_vrf.sli_config.static_routes.static_routes — local_vrf.sli_config.static_routes.static_routes / b2d02bb438fb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167)
- local_vrf.sli_config.static_routes.static_routes

<a id="canonical-edd3cba7b8bde30cfbf47ad31208d38457e71d4f6b931d157f4557e67ce33d21"></a>

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

<a id="canonical-1c35bfb622b72e7638b79c45496743516c974759f57647355e3a2a605ae60125"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes / b2d02bb438fb / 3

<a id="canonical-30bd4f1d39b1cdd8580b29c99c648a65ceb5ef007dca1e8e74574b58a1ee799f"></a>

<a id="canonical-c791288e107f4e46130749d4bc3ffdb64ce82594d9c87788ddb943fddbe589e5"></a>

## attrs property — local_vrf.sli_config.static_routes.static_routes / b2d02bb438fb / 4

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

- [default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-6aa71e5860d1b7d527ef1e9c835978921918c877b41b88dc4d8efde11dfa01db): complete subsection reference.

<a id="canonical-130bec9a271fabfff18e68fca1a5125c34c910e234b883f454432af8fe5b7519"></a>

<a id="canonical-2e6ca59f90b7683ac64b04da33fedec45689f0e78d8818ffa021fe5e663dfd1a"></a>

## ip_address property — local_vrf.sli_config.static_routes.static_routes / b2d02bb438fb / 5

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

<a id="canonical-1dd4e0e89a84de867a6317df4d07a4ed1d0c972da14b46145b3ab168c9233731"></a>

<a id="canonical-305b8439076be6c21c9dfb850d70119dca3b121432e4487e06be02ea8aea0a11"></a>

## ip_prefixes property — local_vrf.sli_config.static_routes.static_routes / b2d02bb438fb / 6

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

- [node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-edb12d01934ec8c15f05a831a251bdd6e002e25a918ef289c9c08a3c2e73e005): complete subsection reference.

<a id="canonical-4e29597b904d00040b8d8982637f7f5fa2a29ec6ab2894a8effb0fb579dabfbe"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes / b2d02bb438fb / 7

- [local_vrf.sli_config.static_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-6aa71e5860d1b7d527ef1e9c835978921918c877b41b88dc4d8efde11dfa01db)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-edb12d01934ec8c15f05a831a251bdd6e002e25a918ef289c9c08a3c2e73e005)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6aa71e5860d1b7d527ef1e9c835978921918c877b41b88dc4d8efde11dfa01db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3f6de234994d967e4101e135af5c975881a112a1ebf2fc6fe464c13cdcd76e1"></a>

## local_vrf.sli_config.static_routes.static_routes.default_gateway — local_vrf.sli_config.static_routes.static_routes.default_gateway / 2434dc6c6272 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf)
- local_vrf.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-0c04b1bf9bceedc8b1c5ac0cdfc77361f1cb43bd720fa740fb5ac97acc821319"></a>

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

<a id="canonical-25c067836e29d5ead21a04d5de82eb15a176bd39cd3526e6f21673ce52d9365a"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.default_gateway / 2434dc6c6272 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3f27020c1373234174fe6df6397e54172eb2529a8a4029eac76a8a118fc57af"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.default_gateway / 2434dc6c6272 / 4

- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-edb12d01934ec8c15f05a831a251bdd6e002e25a918ef289c9c08a3c2e73e005"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09a1124dcdd58d2dcde649518e3dfd2663805fae3e7b35cedf452f6259d5f60c"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface — local_vrf.sli_config.static_routes.static_routes.node_interface / 6df1b8854ad0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf)
- local_vrf.sli_config.static_routes.static_routes.node_interface

<a id="canonical-ae3cdea2e6659142582575cecceab7539d5f6894fd40ac32b36b39f6648eaeb0"></a>

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

<a id="canonical-fc47dc67d06eecddfa5abfff7c19cb59fb54cc2768ff69dc7a6c5328816ee5b5"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.node_interface / 6df1b8854ad0 / 3

- [list](resources--securemesh_site_v2--reference--group-011.md#canonical-122c89b1e41209bf38d371539259b43db1614bd6e417559a0e3f37db334fba5b): complete subsection reference.

<a id="canonical-a9e997abdc0d3dcf8be68b6f1d4ffe0623bbb2a28de0e360c82a8a741e5252c6"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.node_interface / 6df1b8854ad0 / 4

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-122c89b1e41209bf38d371539259b43db1614bd6e417559a0e3f37db334fba5b)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-122c89b1e41209bf38d371539259b43db1614bd6e417559a0e3f37db334fba5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd196dd140afb8749c33066dd965258907d77f19548af098f8094d0a276bf087"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list — local_vrf.sli_config.static_routes.static_routes.node_interface.list / 74d05e4cf529 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-edb12d01934ec8c15f05a831a251bdd6e002e25a918ef289c9c08a3c2e73e005)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-cff5d4a6c4db48a41140611539f670861b0a7c89265bb01bc83a89bc311b4ea6"></a>

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

<a id="canonical-317e60c3dcbcd82742774767067785c3a4b9b7da491cc641c550aa77df0335d9"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.node_interface.list / 74d05e4cf529 / 3

- [interface](resources--securemesh_site_v2--reference--group-011.md#canonical-5c52d61d105e469243f4e1fc9107da5559f04891f03a735dbe55c41a1e87d3d3): complete subsection reference.

<a id="canonical-3536e4f4a14dceac372bf2f8e673cf6a3108e451dc6ef33c235f8e4d6a850c35"></a>

<a id="canonical-59b0ae805349a76f098175d8fc4ec5c777a7bcb8d31a3121a4310b0839c25c3b"></a>

## node property — local_vrf.sli_config.static_routes.static_routes.node_interface.list / 74d05e4cf529 / 4

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

<a id="canonical-9e894b5c5972ef2242cb046cb61b2af04b7db3be39524886d2ef5e10f300a739"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.node_interface.list / 74d05e4cf529 / 5

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-011.md#canonical-5c52d61d105e469243f4e1fc9107da5559f04891f03a735dbe55c41a1e87d3d3)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-edb12d01934ec8c15f05a831a251bdd6e002e25a918ef289c9c08a3c2e73e005)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5c52d61d105e469243f4e1fc9107da5559f04891f03a735dbe55c41a1e87d3d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96c1ae578f1bc634ee736626a90e3d2f9eb8c6b9093d7a6f0c4c0b6b5265a0f4"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3bb1097de112861dabe87dc282059cc23346f3d3a62cae67c565e8a424d2f167)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-d6c83939ba6d7e854f7d6201adfebb4db3eacd7f648fb816489fb0299f2656cf)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-edb12d01934ec8c15f05a831a251bdd6e002e25a918ef289c9c08a3c2e73e005)
- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-122c89b1e41209bf38d371539259b43db1614bd6e417559a0e3f37db334fba5b)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-eaa6044f6563ad0b3acff6341210f5c7082c92d24c13132b6d0faf2a5e4b0c2e"></a>

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

<a id="canonical-626a9bb8781d4708a561799280606f0e5ebbcb24b8ad8ec9095c02d9cf16c916"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 3

<a id="canonical-b7da42628f06982f9cb6d5cccfd16adad96681d21f619593c6d4d4cc69183008"></a>

<a id="canonical-b361578df4eb543406a51dd20b91523821d87f7f39a9fd4b31f5fd47ecdd08a6"></a>

## kind property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 4

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

<a id="canonical-92b80662c25b817cf710f8165257f7261e9346e7aeee105d854e70cd3b1ef35b"></a>

<a id="canonical-2846cc8560aa4c9b5202da6a14fc3bc13b11acfa11cd28afe706ac98adf23854"></a>

## name property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 5

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

<a id="canonical-084ac36cdcb41bb93b2e156cd795f1497f17269b090f8da3a4f5215df3912fff"></a>

<a id="canonical-b0d8e2dfa7c12855e1b6354ea0b62f113f5b5fef31a2e3ed35c1d1206d700f3e"></a>

## namespace property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 6

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

<a id="canonical-723751297185e35be66391f65fba65e5bbdc5cb9ad7537f79cd6ae1a8c17715b"></a>

<a id="canonical-30ff69febdd80386dbc98cebe2e8bc393e4aa6ad1f4a9389110b37366048e748"></a>

## tenant property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 7

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

<a id="canonical-abdbdbb8f6404b5c50aea3ca9402696ce6b980403b9e45fda02d4647900f7a5e"></a>

<a id="canonical-dc4bfe49a6729aaeb3fb2553458aa1a7254be56e405c4b7432b1bd578405958f"></a>

## uid property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 8

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

<a id="canonical-87981af5cf9954c6cbd8237886e0d34a204ef41ed7753b02d32955bcf5aa550c"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a6ceab613b34 / 9

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-122c89b1e41209bf38d371539259b43db1614bd6e417559a0e3f37db334fba5b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be0a2670559fdd2bb090304db5a9f09c4c857c2d0714acc1e324d54c9702dfd1"></a>

## local_vrf.sli_config.static_v6_routes — local_vrf.sli_config.static_v6_routes / d39c874d5c4f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- local_vrf.sli_config.static_v6_routes

<a id="canonical-33042f05bcf98a08eb6116ea568720a8344653dbed73e45667a55180a5b3789f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-a3fb6f5e2572f420d39ceca562da09e7cafec41a0099beab4d569690b3c77dd9"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes / d39c874d5c4f / 3

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26): complete subsection reference.

<a id="canonical-95ee67ac96c8b0605a6774798f28d2ae6a608b24fe04cc066f4220a0d4f8cf0b"></a>

## Next pages — local_vrf.sli_config.static_v6_routes / d39c874d5c4f / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a0e1d5dd6a5fb75f9da0d9dd987d7f73518491e2110c5674ae580419f4632c5"></a>

## local_vrf.sli_config.static_v6_routes.static_routes — local_vrf.sli_config.static_v6_routes.static_routes / 8a8b74a296fc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040)
- local_vrf.sli_config.static_v6_routes.static_routes

<a id="canonical-3957cf26e253e3f074fc4289ae346b06656760ab0d3737e7ece988df20d91cfa"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-318199163083e4095368600f31c00daac567627d319bcd5b1bdbaa8fdb567612"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes / 8a8b74a296fc / 3

<a id="canonical-4b3b4e2d63b322af8f5c83a7bd7dab7158c54e5825d7b307acf70e99d29bf96e"></a>

<a id="canonical-6fdffadbb211651c0f221636d922fcca50e471cebb39cb7ff1a3ff26a43e45d1"></a>

## attrs property — local_vrf.sli_config.static_v6_routes.static_routes / 8a8b74a296fc / 4

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

- [default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-68b44048b8e1b2259ff8e3a27d9d1acc52cb171b96595d020830343567cb3c18): complete subsection reference.

<a id="canonical-c6ed5af90c4f726b58406cbacbd2f50d1f4cb3093e230cfc400479e7def551c9"></a>

<a id="canonical-541ebd6f62243f3e65b3fd0d71b4387dadbff6538d18fd611788b39110324433"></a>

## ip_address property — local_vrf.sli_config.static_v6_routes.static_routes / 8a8b74a296fc / 5

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

<a id="canonical-0fbe47f8747cc88128354ec0d6319af42710179908aca046c45bf8913b7082da"></a>

<a id="canonical-1f71f12d5e015a201e960ca51dc3a882f9f38298be33b6739a23c3fce78774ae"></a>

## ip_prefixes property — local_vrf.sli_config.static_v6_routes.static_routes / 8a8b74a296fc / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

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

- [node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-c9323f21a245f3ed8892a4083db6241e29acb492cac438347deb9258dda15d8b): complete subsection reference.

<a id="canonical-45ec8852d5070a3c91799588cf56cb9f0a2af1c2a8de3f557cfbe31d8c3b2728"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes / 8a8b74a296fc / 7

- [local_vrf.sli_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-68b44048b8e1b2259ff8e3a27d9d1acc52cb171b96595d020830343567cb3c18)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-c9323f21a245f3ed8892a4083db6241e29acb492cac438347deb9258dda15d8b)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-68b44048b8e1b2259ff8e3a27d9d1acc52cb171b96595d020830343567cb3c18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fddb9691928f5785cacca830b8eb5585aaf24e56d3bb97c0637a4cda818294ac"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.default_gateway — local_vrf.sli_config.static_v6_routes.static_routes.default_gateway / cfadf4c84bed / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26)
- local_vrf.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-d519402ad3fb655823a3ef63c15794f4b2bfe746695693c94ee845cf878dfbe3"></a>

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

<a id="canonical-f38d41f80197ba73ebeff69a53e375c5650b87f872c9e68982cd96efceb349d7"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.default_gateway / cfadf4c84bed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5f45d7a8c0bf6c33e6e2a5b35dd108f4fc1d4f8fb191521e38bc184579313b2"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.default_gateway / cfadf4c84bed / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c9323f21a245f3ed8892a4083db6241e29acb492cac438347deb9258dda15d8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aee3c954711a0faa8b92767bd422e7a9179514a3773c00bdc56689b4f0fe3e2f"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface — local_vrf.sli_config.static_v6_routes.static_routes.node_interface / 41f24734677a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-58f6b9a84229ba8ad754c509ae8c3045e4f534e04da717834d5089e21adfa886"></a>

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

<a id="canonical-df4510bbfc15c987b462b26a964376fd08306757598addd9241ef6713baef036"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.node_interface / 41f24734677a / 3

- [list](resources--securemesh_site_v2--reference--group-011.md#canonical-4ff7a0bf9d77c87bca38724fa7ad850d973ca8f58c5a7bea10d776b1e9b984f1): complete subsection reference.

<a id="canonical-fce156946bd0a0096ac2267d3dd86a0246690472a716fbda3561da182c6f41ba"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.node_interface / 41f24734677a / 4

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-4ff7a0bf9d77c87bca38724fa7ad850d973ca8f58c5a7bea10d776b1e9b984f1)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4ff7a0bf9d77c87bca38724fa7ad850d973ca8f58c5a7bea10d776b1e9b984f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c20fb9d25ba2bd9cab0f31110401a6db3323c6e053c409358fd0373c42e81c4d"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 8f508c14a16a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-c9323f21a245f3ed8892a4083db6241e29acb492cac438347deb9258dda15d8b)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-cd550efa43951364e265ae547e9f4bb12354ea04b8b7cf98b5ebc4759efac518"></a>

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

<a id="canonical-2591bf485f04e3709eace6db61c744a2cadf41a7401c99507715681c91b081f1"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 8f508c14a16a / 3

- [interface](resources--securemesh_site_v2--reference--group-011.md#canonical-46d2464b9b6457344fe780a4cbffb1c52dec95b5c7276464508d60fc2586951e): complete subsection reference.

<a id="canonical-d16e873206ae5e62a6ebd1ec67e060f75745365763cd3f74c81620dfa2cf9039"></a>

<a id="canonical-b3d4b689132dc2ff5337bb22f6f858ed0924c6707e2a1813a2cf181fdcc23619"></a>

## node property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 8f508c14a16a / 4

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

<a id="canonical-754beb479cb7820c7a5eb6854d8765936fd549e627e97617d3a0490ac6341a6a"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 8f508c14a16a / 5

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-011.md#canonical-46d2464b9b6457344fe780a4cbffb1c52dec95b5c7276464508d60fc2586951e)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-c9323f21a245f3ed8892a4083db6241e29acb492cac438347deb9258dda15d8b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-46d2464b9b6457344fe780a4cbffb1c52dec95b5c7276464508d60fc2586951e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b0e47fc58835c4ceefcfda07cadd75f4706bf81bc65df5a75edbc569c571795"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-433c4e620f2098503c844fbe51c8ea9bd071cd78fdff98a438ff39894f6480c0)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c1784c4628b75ec835e92d59193c1f6c9a95e6bf12f978218fc89bbc36d83040)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-94a907ae9111bee6f7f43f8336e7c496b61cfbfc2d9483d44d0a45d83753cd26)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-c9323f21a245f3ed8892a4083db6241e29acb492cac438347deb9258dda15d8b)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-4ff7a0bf9d77c87bca38724fa7ad850d973ca8f58c5a7bea10d776b1e9b984f1)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-948cd402714f9003875c56b3d14fa6c6b9aa98506bf268f63f4f76ce4f94ed0d"></a>

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

<a id="canonical-e55825335754dffd0d076158213ff8e623aef46f590b1d78d9e3a9a62d56ab0c"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 3

<a id="canonical-cb996f40fee9a2672762d7caf7336eeb7d24663486860d93ab300d32a9380c09"></a>

<a id="canonical-34cbd2343dc63ffe52042e5ba5b664b8be4273d91adc8ed7c4161709057413a2"></a>

## kind property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 4

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

<a id="canonical-d7986710568d9309d5b6ecdec763963f5b45c081b32b5ef8d35549b3e091a4f4"></a>

<a id="canonical-29115baa77bf065f5f85140380dcfe4d0fccd43e7b0242fc9942e10b4cc44482"></a>

## name property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 5

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

<a id="canonical-dcaa81fe2096f2135dcf888138cfd3a50101bce9bef7bc5df5ff1fb42f6193b4"></a>

<a id="canonical-95cf63e4fe84b909a95e237b945be9663ec7a61dcdfa12924538f5e8d5479c18"></a>

## namespace property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 6

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

<a id="canonical-c1de6b5408b50fa769c22e16fb74db62309535a92a9534179fafbf08ebd503e4"></a>

<a id="canonical-12e0e7197206d243ba8dc34a6b923577902d64ea36ca1a3e3cf6ba8edbeebb73"></a>

## tenant property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 7

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

<a id="canonical-eb70f714cb246f0ce6d16a8fd196beeb6b63c9a1e9923a94fde1473f0e65675c"></a>

<a id="canonical-7b84a0430c190bc91f252bec24360d8eef32830dc5b3195291111a8abd0d2a96"></a>

## uid property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 8

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

<a id="canonical-13dbf6e3239934f62f4a567843afef1fbac525dbfa7e7545ca15424c5cc0522c"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 580ded282b11 / 9

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-4ff7a0bf9d77c87bca38724fa7ad850d973ca8f58c5a7bea10d776b1e9b984f1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d852c744eb426670717e7a03d03142beaf8c83463a7b3fdc6edb4e9dc3303f2"></a>

## local_vrf.slo_config — local_vrf.slo_config / 6d2558f4c0d0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- local_vrf.slo_config

<a id="canonical-0b87eaf21b0f6d224553e8c08fe12ee146e904536af5dac2c0b29a8455e874fc"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

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
slo_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c2d996cf6638f30a4c69d0e2053ad53c74be398d55f2762f715accbe604248c"></a>

## Direct properties — local_vrf.slo_config / 6d2558f4c0d0 / 3

<a id="canonical-6f599f45eb404aa48592efe3fc1eceb2b35681522496fa865247e09ef5483d89"></a>

<a id="canonical-9c5d91463ec78369f89aec32563bfd7ded33988c5e65341221c3db87cba4c2ee"></a>

## labels property — local_vrf.slo_config / 6d2558f4c0d0 / 4

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-18d0883b1afe9136a99d99412f1fdc8fb1919157b152c6d28586ae0a10c2f5f5"></a>

<a id="canonical-cdb925d6c630ecd4d7bb5d63336764001297209afd7b47595704383a80e75fdb"></a>

## nameserver property — local_vrf.slo_config / 6d2558f4c0d0 / 5

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

- [no_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c4770eb843cbffb04895186b7cb39540f243d30f9d944e1117115e6ad4edaa31): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-06b2a99c228d6735043147502c84edba3e8ea01181d118e39838609ac13c46d6): complete subsection reference.

<a id="canonical-b6b04e8615b95bcbf90f9b1492dfbe3a86dcf20835d3aa6a11b14f773541a5a2"></a>

<a id="canonical-c0439674fd60ccdc43f2f4fcedd3ad749b157090ecc2cbce25bff355c00b41e7"></a>

## secondary_nameserver property — local_vrf.slo_config / 6d2558f4c0d0 / 6

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

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8): complete subsection reference.

<a id="canonical-e3937b614e41d29d3688618944bdafe0d345d9160bcac22de9468883a891c0f6"></a>

<a id="canonical-43f300b77f8a9f38b9037cbc6bdb59f9b6ddb299c20a54b26348e3e7e64a7e9b"></a>

## vip property — local_vrf.slo_config / 6d2558f4c0d0 / 7

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-d9d7f0a8a713fd4d05639025fc4eaec919a1aa58e229db10565f3d321c54cbb4"></a>

## Next pages — local_vrf.slo_config / 6d2558f4c0d0 / 8

- [local_vrf.slo_config.no_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-c4770eb843cbffb04895186b7cb39540f243d30f9d944e1117115e6ad4edaa31)
- [local_vrf.slo_config.no_v6_static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-06b2a99c228d6735043147502c84edba3e8ea01181d118e39838609ac13c46d6)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c4770eb843cbffb04895186b7cb39540f243d30f9d944e1117115e6ad4edaa31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34404567c08f451937e52039b17b56f6f337ef2d5ce5c128a2dfa3ea6d0ee29d"></a>

## local_vrf.slo_config.no_static_routes — local_vrf.slo_config.no_static_routes / b1ec55837933 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- local_vrf.slo_config.no_static_routes

<a id="canonical-ee6747c0b894074793590f17c4d30c51d658e679af116c92efc273774563e8cc"></a>

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

<a id="canonical-2dff2de00ea592e22f340687b89aace127ddcfd39d6a4f47007846d1095948c8"></a>

## Direct properties — local_vrf.slo_config.no_static_routes / b1ec55837933 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04813a4c90213d4fb92554c22675f62a32d625f87cd673a834be35ebfab48d8c"></a>

## Next pages — local_vrf.slo_config.no_static_routes / b1ec55837933 / 4

- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-06b2a99c228d6735043147502c84edba3e8ea01181d118e39838609ac13c46d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4da82680ddd1dbd2ac7aeff363d3d2d966897891acca04f66f47b55eaaeac4ef"></a>

## local_vrf.slo_config.no_v6_static_routes — local_vrf.slo_config.no_v6_static_routes / 47ef1be63dd9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- local_vrf.slo_config.no_v6_static_routes

<a id="canonical-92f2b3dba9ea7b1ac3cb0195ce8dea506a53021f301e5b34baafafee60449f4f"></a>

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

<a id="canonical-a64c9ef8059565dde781288d6eb950cf5447cd32988c87ac1dfb818d580087c6"></a>

## Direct properties — local_vrf.slo_config.no_v6_static_routes / 47ef1be63dd9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-804d63691243a4eaa838c65818c9a618c8debd2d029e37044b1e6451d05ac486"></a>

## Next pages — local_vrf.slo_config.no_v6_static_routes / 47ef1be63dd9 / 4

- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fe0969e312a93f094f0c778b48f6ba7dcbf71b0d52e89f43a14943e751400db"></a>

## local_vrf.slo_config.static_routes — local_vrf.slo_config.static_routes / 32752c43ac12 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- local_vrf.slo_config.static_routes

<a id="canonical-7998b7142601447a344525c622c5f816f73e068722eeefceadd216bd6737032a"></a>

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

<a id="canonical-2faa5f1a2acbff26007daff4a96adc44856cc8e84d7e00a3dbfe51905b7df27b"></a>

## Direct properties — local_vrf.slo_config.static_routes / 32752c43ac12 / 3

- [static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0): complete subsection reference.

<a id="canonical-9d8e7441392d45c64775263abab39ef815bdf3031724c3dd37600e9f34cef661"></a>

## Next pages — local_vrf.slo_config.static_routes / 32752c43ac12 / 4

- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81804a73bc2f77e252c07d814e6f776f45ea22b7c36160cfab68ff115fb42b12"></a>

## local_vrf.slo_config.static_routes.static_routes — local_vrf.slo_config.static_routes.static_routes / 0bcbbb0f31d9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564)
- local_vrf.slo_config.static_routes.static_routes

<a id="canonical-4e22e4abe913250a8d242d3e00086a9b92bd3b5150526802bb4d94aa954a0216"></a>

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

<a id="canonical-741b22daedbd37a65453da5c66ca21c549cdc05791153a6682bc9a74cfd42177"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes / 0bcbbb0f31d9 / 3

<a id="canonical-212af2375281e33a850ddc71d1f1e181a561deec3c5fc6d80a21cbcc13b590de"></a>

<a id="canonical-435e89534d88ac5d007aef184b0914dfd431224c76d71a3f69f4ea7adb099a2b"></a>

## attrs property — local_vrf.slo_config.static_routes.static_routes / 0bcbbb0f31d9 / 4

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

- [default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-44dcc95d8518ee0f8f05669fd60f993a56ed3d6937a8cc276eb4514d5e39097b): complete subsection reference.

<a id="canonical-21ee9572283e10b85248d040c28c1f2cf9827696a3ed9207686ef13888d0d8b2"></a>

<a id="canonical-c3e445c31a89528afb5aa065b98868d49fa33acb81a036dd92e1aece36047f6e"></a>

## ip_address property — local_vrf.slo_config.static_routes.static_routes / 0bcbbb0f31d9 / 5

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

<a id="canonical-da9d4035d024cea33ff6056fc2c223ab96e3cbe5c23eeb26f20aa4df4b1464a9"></a>

<a id="canonical-14fa05d418d75680e92686ba3ecdeec4c1228d33c159a107f739577d4f739047"></a>

## ip_prefixes property — local_vrf.slo_config.static_routes.static_routes / 0bcbbb0f31d9 / 6

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

- [node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-e5a2d87110f268db0b060ad49ac4400b9813118d308e9ae9c97941f9d717a41a): complete subsection reference.

<a id="canonical-6cfc4ca3209f22f51de8bd4a9e0fff6386d7adbcdd1c3d2be1ef0fe1a308fbcf"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes / 0bcbbb0f31d9 / 7

- [local_vrf.slo_config.static_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-011.md#canonical-44dcc95d8518ee0f8f05669fd60f993a56ed3d6937a8cc276eb4514d5e39097b)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-e5a2d87110f268db0b060ad49ac4400b9813118d308e9ae9c97941f9d717a41a)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-44dcc95d8518ee0f8f05669fd60f993a56ed3d6937a8cc276eb4514d5e39097b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4106d6d46450cff6356ac8ece9265759aaeecaae2c52d3e47746e81868da6dec"></a>

## local_vrf.slo_config.static_routes.static_routes.default_gateway — local_vrf.slo_config.static_routes.static_routes.default_gateway / 19ea36726ca2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0)
- local_vrf.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-93da652f6388cc2cc3807f6bcc8c6980732cbf10fb5c63c9029e5fa94ffe87fd"></a>

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

<a id="canonical-6b784fd7f01b6a9f67439d59c7bba4bedc7d79c079948dfbc4b5a29388dbd754"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.default_gateway / 19ea36726ca2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-abdede054a68b4ac1ad1e4ee6bf429a5cef7da849584320936c13eeaee6bfbbc"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.default_gateway / 19ea36726ca2 / 4

- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e5a2d87110f268db0b060ad49ac4400b9813118d308e9ae9c97941f9d717a41a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a838fcee98c784baf7b25b09770d3fb48cc947fcc6d41ade7cf300650865f8a0"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface — local_vrf.slo_config.static_routes.static_routes.node_interface / 0c66970d2dd4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0)
- local_vrf.slo_config.static_routes.static_routes.node_interface

<a id="canonical-1ffc64d7396a3be4c512a23718e6862654f699e4232eae9e653dbe4b093f6b93"></a>

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

<a id="canonical-81d68a6490f80bce68fd56ada45ed163cabf4a09ede348249b8afac143eafa8f"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.node_interface / 0c66970d2dd4 / 3

- [list](resources--securemesh_site_v2--reference--group-011.md#canonical-08b51f16fa711b354f0067ee85777bc4870af5335cf76f1e4aebea417e3d5317): complete subsection reference.

<a id="canonical-6f62eafea2e1401bb3612daef972e336fbb9937e1affcd00b6946122a902c9e7"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.node_interface / 0c66970d2dd4 / 4

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-08b51f16fa711b354f0067ee85777bc4870af5335cf76f1e4aebea417e3d5317)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-08b51f16fa711b354f0067ee85777bc4870af5335cf76f1e4aebea417e3d5317"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
