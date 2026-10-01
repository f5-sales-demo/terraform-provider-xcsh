---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-22569355dc3c3cdffa40fb79465e6d8c6315de34f343f52db94e7bb4881e6a71"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router — aws.not_managed.node_list.interface_list.ipv6_auto_config.router / efd992cecc16 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-b6076e8da6fb4c9200f58f5b3bf9fa362f0e3ccdf5bf1708579ef20dbd23630f"></a>

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

<a id="canonical-5a7e7673a7e907b5eb6d5f5ffc1e1901bd4605fabf4ebf514867a067250abe00"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router / efd992cecc16 / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16): complete subsection reference.

<a id="canonical-5a973d42a24eef275fbc2dc8ea82a7180d229886a2f727636cb93f24a3c63840"></a>

<a id="canonical-2288a7f74de95b651814b06506f1e5ab6ab2cfcedbf69870f5ee657eb6cfa659"></a>

## network_prefix property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router / efd992cecc16 / 4

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

- [stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0): complete subsection reference.

<a id="canonical-097ee9e013b372720fd1409eb5544af508e9efe02192f40804c8a632a495c8ba"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router / efd992cecc16 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0cc4bcb85872e8d831f1cf553dd50d9a37499b9cc57c58180bd385014058d2b"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 5871b3b778ec / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-a8964787567c9694dd7de710ed044df6d9cd250ac9a2172772ee06b7384478f2"></a>

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

<a id="canonical-f1fadcc489fdf0a710f31c13129cf18dbe41a03a2dab843cba47e6a4091dfa76"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 5871b3b778ec / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-6a71898e337fadb4b48011f972a9aedcbc116f838b28e42da87a13e927cdcef1): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bb191101bf11131b2e1897c09be43f825b1cc5a16663e4bf45cf5b1973574aa7): complete subsection reference.

<a id="canonical-1f00728f388d5ac3c4d02885802515d38f24e1fd725df45a3ff06ae2244eabcd"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 5871b3b778ec / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-6a71898e337fadb4b48011f972a9aedcbc116f838b28e42da87a13e927cdcef1)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bb191101bf11131b2e1897c09be43f825b1cc5a16663e4bf45cf5b1973574aa7)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6a71898e337fadb4b48011f972a9aedcbc116f838b28e42da87a13e927cdcef1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f036331760d04edc9e0a87034abe6bf19be25292fc7921df5ccec811e3a13244"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 48f5d9f40a00 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-a2ec9737748cd18520b1bd4ed00698cd9212ee9dc1f3100e13bdd5862c7cf4f0"></a>

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

<a id="canonical-45e29459875479e55d1d280dd2d789b601d3717a023a4d17c84cf8b868b12274"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 48f5d9f40a00 / 3

<a id="canonical-12ea3c9a51d06b1c6feaad331613a0208b226c3aaf504800dcc99d75e37bf80d"></a>

<a id="canonical-6b22cee8902a9ebd5aaa11ea50159346bbd527ef5ef5ade80f50cf7fbf6f0450"></a>

## dns_list property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 48f5d9f40a00 / 4

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

<a id="canonical-e866a1484bdc563e45062150cf2625361883e7ea2fdd5c5e01486eb165f50118"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 48f5d9f40a00 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bb191101bf11131b2e1897c09be43f825b1cc5a16663e4bf45cf5b1973574aa7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-785d8d30f9ec0347bea4c37ab36d58ec787c9c895f8cb2b4d0623dd16c1b3070"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 6a0ad40891de / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-a6235097c1cc0b9dba03a7d1876b6ebdb63946084df7dd7ea615d6895f411973"></a>

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

<a id="canonical-90d842bf1eb1584039214366c3d28b818977a52599fed64e55a5012fc8bf6387"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 6a0ad40891de / 3

<a id="canonical-2574bb8f8cf8ab7d96eb93851e39e1bb614ccf5d07623080793e54bc0abb4e60"></a>

<a id="canonical-86784451df88956b63596e2154b122b2fc0b7b0a87dd6af568454d2dc2ab358d"></a>

## configured_address property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 6a0ad40891de / 4

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

- [first_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-615647db3041834a2887746e63a5b7b021b4923e4c277fde0828ef19726f8dab): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-754640a87bc14039b8f3b522fe83164c5e99f6158690e2c05fa5468e26563d09): complete subsection reference.

<a id="canonical-41affe6cbde83f89e2f3592e985e93ee45dc150f713b0134a8d9cce91a18d267"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 6a0ad40891de / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-615647db3041834a2887746e63a5b7b021b4923e4c277fde0828ef19726f8dab)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-754640a87bc14039b8f3b522fe83164c5e99f6158690e2c05fa5468e26563d09)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-615647db3041834a2887746e63a5b7b021b4923e4c277fde0828ef19726f8dab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-284ff24e913319d549294f1220c3343f910edaf38f8175122f56d4b1599f35f9"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / cddc508729b8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bb191101bf11131b2e1897c09be43f825b1cc5a16663e4bf45cf5b1973574aa7)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-412a9b1e047ca8bd77de69bed9884754d0f99f3398c0c6c30eb21f01b1a5980a"></a>

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

<a id="canonical-561f50cc84ecbea3fc2d95456be772348f37a30b48e95cd345464f3597a851c7"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / cddc508729b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d225d868c334b75538affb47966b4c5d4ec3f6d3daf9e0c2e9710d6f9fe7b1b4"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / cddc508729b8 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bb191101bf11131b2e1897c09be43f825b1cc5a16663e4bf45cf5b1973574aa7)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-754640a87bc14039b8f3b522fe83164c5e99f6158690e2c05fa5468e26563d09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae2d0b4295a32b65e016ac978e647d5b2c80c0685de9aab71c7925a72c983e05"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 671bf1ee50e3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-c6a43f054330e60a3afc78ddb0506018c0dcaf4a1bfad30808f073c93e82fb16)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bb191101bf11131b2e1897c09be43f825b1cc5a16663e4bf45cf5b1973574aa7)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-8d01ea17ae0c2ac433f4b053ac160632015cfb2f1e2d10100efc2624eee57e71"></a>

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

<a id="canonical-301ba6d1566a4a70c7ac01e1065e2e692b9f3beb3fc97e4908f5ed05c854b0b5"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 671bf1ee50e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04a2f352bf5e531731a3a89edb91d066ee0701c5821d2415940ef5981c366389"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 671bf1ee50e3 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bb191101bf11131b2e1897c09be43f825b1cc5a16663e4bf45cf5b1973574aa7)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b7c803e12c5247f73893b5e91db22e8068cf97e0e9376d8f0ae38e515bdc08f"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 36084ae1427b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-6d724ecc71b1c08d1fccb4ec4c1e6e9b5136165c926229351887bbab438dc674"></a>

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

<a id="canonical-5fab5bf0925350ea36160c5cd7925249ccfadae47c944da70a065f9072845173"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 36084ae1427b / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-004.md#canonical-70bdf3eba54cdcfb4b636017afaf9331e85067d2ce0a0c6bf0e5835c853623e7): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-004.md#canonical-fc9247ebb87ac8f42a56b9419e9d99a7b84b4ff1e17e7f832d032b8515f0b1c4): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-796692084501f3508cf624db28b78c6df90c6e915292603b099bc741f0b996c9): complete subsection reference.

<a id="canonical-c168472a623ed70aedf8debc4c2c22834616df63f647c2399eaafdc7812d956e"></a>

<a id="canonical-7d0688391ea7502d3b1d0937740c76f8fd9c1ad3158fcfbf006a900c8b83bfe6"></a>

## fixed_ip_map property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 36084ae1427b / 4

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-004.md#canonical-cfe2973e923143d7c953b420525f2a0c571372c8ff3574b95b5d0379dfaea0a4): complete subsection reference.

<a id="canonical-a31b8a43fd0121cc18376a7c31297152369b1aa341fe8ce5d4142ba3c1a39f37"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 36084ae1427b / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site_v2--reference--group-004.md#canonical-70bdf3eba54cdcfb4b636017afaf9331e85067d2ce0a0c6bf0e5835c853623e7)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site_v2--reference--group-004.md#canonical-fc9247ebb87ac8f42a56b9419e9d99a7b84b4ff1e17e7f832d032b8515f0b1c4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-796692084501f3508cf624db28b78c6df90c6e915292603b099bc741f0b996c9)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site_v2--reference--group-004.md#canonical-cfe2973e923143d7c953b420525f2a0c571372c8ff3574b95b5d0379dfaea0a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-70bdf3eba54cdcfb4b636017afaf9331e85067d2ce0a0c6bf0e5835c853623e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-242c8e9a5439f42b2eadf4a96c927d9505bd81db080aa4278728fe636a2e4c93"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / c3b00a2a6697 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-64952cfc2320161a5c8a3055993eebda1e7da9fe99585058294df989b3ec7792"></a>

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

<a id="canonical-f2b85eb13227a6b9c12ab96ead0fd4288bb6e013bd1f1a9725f6e2e20b77e422"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / c3b00a2a6697 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e70a36c25ce43b10e2f6af13326984cec6dc0f824d2adac88f52897a1084e37"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / c3b00a2a6697 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-fc9247ebb87ac8f42a56b9419e9d99a7b84b4ff1e17e7f832d032b8515f0b1c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43c6dd88616ccdecb2b1286e935df35a9e76fbebde8304fdbdd0b00ab898d49c"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / e687f07f1d3e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1777e945c9b25f5a8b477ab1ae7ccb36f46fbad90739a6230223875c1f4ee8d6"></a>

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

<a id="canonical-5c7145b796381c7fe9f45f800df69f5ab179578c47709e2b10b04344b2767262"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / e687f07f1d3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-545a2c27cdcc816cf164cae2724efc3f7f1e01101feacc4d016a2d8498aee246"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / e687f07f1d3e / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-796692084501f3508cf624db28b78c6df90c6e915292603b099bc741f0b996c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1c747202bafd33c9a283ecbfccd571677a72fc6a866a419f7d22ada66488c9c"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 3c5af94ade8f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-9d1a489855294f1968cfe01910b59277bb60190238b8d3cdb862bcb5f06275b5"></a>

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

<a id="canonical-6f9bc78e0c9c5a0afd2f0ef42e079d563dcfc4c8e2f00846aceb23e5205120e6"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 3c5af94ade8f / 3

<a id="canonical-06ccd04bed379e5837c7c06bb191b5d798ef9c57ae4c5f8ee35a1ad0f06fc522"></a>

<a id="canonical-30fd279194045847bd47f0e24a0147e9a1a7cd50e3c5d463f5fd9c75bcbf4d6a"></a>

## network_prefix property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 3c5af94ade8f / 4

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

<a id="canonical-768439079809c393f6c3c46c2cfef2a95aececf06bd9037b6da3ff4a868dd221"></a>

<a id="canonical-82e7ca15c36a7a9360c09734b8f6db53c53102e7e06149eeecbe4109fb57f4ac"></a>

## pool_settings property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 3c5af94ade8f / 5

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

- [pools](data-sources--securemesh_site_v2--reference--group-004.md#canonical-20d4e442735098f34785c4f98e4dd55c81c65a8bb77ccc3d48e60371fd263cf8): complete subsection reference.

<a id="canonical-3d09f8cbcbffc609c69b48652f4863be5ac5afcef04b391bce4ce7ab06cc8fda"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 3c5af94ade8f / 6

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-004.md#canonical-20d4e442735098f34785c4f98e4dd55c81c65a8bb77ccc3d48e60371fd263cf8)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-20d4e442735098f34785c4f98e4dd55c81c65a8bb77ccc3d48e60371fd263cf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e182cd386c5bf71cb0db1ebd280e2ef2fe22f8d070ba5aa8ffa287052bbdedc7"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / df7130b8dc1e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-796692084501f3508cf624db28b78c6df90c6e915292603b099bc741f0b996c9)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-041eb8e0bcd33047d8509c0cd3e3a56601ddabb0844911e201e6dfb33ae255a4"></a>

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

<a id="canonical-d70da400ebb05b2ab5a6a88f82da78926753c9572fb299cf9d6107b017ecffaf"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / df7130b8dc1e / 3

<a id="canonical-8a92fc7a5cea6bba082599f96aacd68c89c9e67a2ff1dce4769dd1c2f321cb34"></a>

<a id="canonical-46b07e42d655c727e08c77b920f68de1a43b32ff7099963fab42fc2f2bfd551e"></a>

## end_ip property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / df7130b8dc1e / 4

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

<a id="canonical-8dd16d1fb9f86102dc02eae8a3d394f68682a8517d40a98461a3ce1a94c5e5cd"></a>

<a id="canonical-18909e15a017097c43fb4df497724bc99be8cbbb284369a08fab648bd068fd83"></a>

## start_ip property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / df7130b8dc1e / 5

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

<a id="canonical-0d81f25d8657935d647355cb520e8e1e48a7c78035a49db23ccec643347364a0"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / df7130b8dc1e / 6

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-796692084501f3508cf624db28b78c6df90c6e915292603b099bc741f0b996c9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-cfe2973e923143d7c953b420525f2a0c571372c8ff3574b95b5d0379dfaea0a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17399a9013de07b3a75da5610f3345a3ea0cbda20f3aee4ad7459c14bc590c0c"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 24577646c550 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-003.md#canonical-52b288aba7eb0ccfb61426163e002d3f8dbab4323d3c13eb3e5a5b5fde407038)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-003.md#canonical-779b1d6c25c411e663977c5a3dbf0ae7a810e04a9c6678e648f630263ee5f5a4)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-6f52ca67ee957db948e3d4f30ef011af965787246e5f1b6a6fe7ae43756ac75a"></a>

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

<a id="canonical-a194c8df3608f92ae782fc6a69b1e9d85290a7f09716ff4813d43bd5159cfde7"></a>

## Direct properties — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 24577646c550 / 3

<a id="canonical-594783177f69e8deb4aa5e15baf9697fe2189b03cb535e77d559fd906e1f726e"></a>

<a id="canonical-a04df03b1b2ce01074b262f59e4cd2bc0e99cff910cdf600db4debc75c8fb2af"></a>

## interface_ip_map property — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 24577646c550 / 4

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

<a id="canonical-9061d5dfa6ef684dc6aaeb37dc293583e750f370b42233f0b608816188515ee8"></a>

## Next pages — aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 24577646c550 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-004.md#canonical-45f8cd28bd358420c676646d34af1ac40b83f713d532fa7b81a1cd1dcb43b5b0)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-435b0278a3eedb776722af86c8aed64951423575022b7a4085c0c59a15ba1d40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e571e5812f68008daa883b516c2882d3d2a56d80afca34c7b53ca991ca1a722e"></a>

## aws.not_managed.node_list.interface_list.monitor — aws.not_managed.node_list.interface_list.monitor / 9ac804cb7cd5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.monitor

<a id="canonical-2b6c97f898dbd92db480fe951b07cac41340c0f7f115ae6af0da4067758f397f"></a>

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

<a id="canonical-6312f98306baaa0cb308cd193aea2c05d71f1be4fc003709e12e86cddcb0599f"></a>

## Direct properties — aws.not_managed.node_list.interface_list.monitor / 9ac804cb7cd5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bfb9186af2b5a8fd90563fa1e8aefb229c9e4ffecdc0e93b0649f73b7d03da8"></a>

## Next pages — aws.not_managed.node_list.interface_list.monitor / 9ac804cb7cd5 / 4

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7f2013db48acc0388182dfdc0966b3b2ed2283fc44d91762c789800a3fe206db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f2e59829df110887c314a6d37a39b0133cde7509d3af4e6d262f1bc0f781d0a"></a>

## aws.not_managed.node_list.interface_list.monitor_disabled — aws.not_managed.node_list.interface_list.monitor_disabled / b0e6edf06dbc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-06836bca28e1fcb033d0c19d6ff30346138ed62276e12e5fd6f09cb84b914a7a"></a>

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

<a id="canonical-400d180d6267fcd1cb3c67cbbb8e2d6392221f4e02ae75ec0821244a9f2c7118"></a>

## Direct properties — aws.not_managed.node_list.interface_list.monitor_disabled / b0e6edf06dbc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10140c88f416008ac820295b4a0148a25ff6a48aae930b9a4deabd480b6d7668"></a>

## Next pages — aws.not_managed.node_list.interface_list.monitor_disabled / b0e6edf06dbc / 4

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-6fd4cb11a8978a039b74b2749c52ea22393431638d4f5f76696e302e7b839f4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbc199915d6ed1c417af665f02b870d59929546758326601e8fac14a870ede3c"></a>

## aws.not_managed.node_list.interface_list.network_option — aws.not_managed.node_list.interface_list.network_option / 476763f9c4e0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.network_option

<a id="canonical-6d02cd5ea1eefba434d16dc3dab0128375119498183ba1c6b2e79601e36caaaa"></a>

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

<a id="canonical-60bcdf93f05cd55c80a35f5447d8ccd11e254d57497007cfade55c25b0132fc0"></a>

## Direct properties — aws.not_managed.node_list.interface_list.network_option / 476763f9c4e0 / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f06413b91665ad59bd1658a97bdb5a27237db3eaecc9773925b58cb4f6251d20): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0f245f02e688690308d58c68005f307414f47c2f96d76f6a4917eeadb0fbdf85): complete subsection reference.

<a id="canonical-af01e52c6c1ecd3c33f7ac41014ba4066131853bfe227746f44fce864adf047b"></a>

## Next pages — aws.not_managed.node_list.interface_list.network_option / 476763f9c4e0 / 4

- [aws.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f06413b91665ad59bd1658a97bdb5a27237db3eaecc9773925b58cb4f6251d20)
- [aws.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0f245f02e688690308d58c68005f307414f47c2f96d76f6a4917eeadb0fbdf85)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f06413b91665ad59bd1658a97bdb5a27237db3eaecc9773925b58cb4f6251d20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3028d69cbeda9153c6b25eb0c6248f3a834a7a60d114003311a14bc3cc35967"></a>

## aws.not_managed.node_list.interface_list.network_option.site_local_inside_network — aws.not_managed.node_list.interface_list.network_option.site_local_inside_networ / b0550c15f937 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-6fd4cb11a8978a039b74b2749c52ea22393431638d4f5f76696e302e7b839f4c)
- aws.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-02c69edda252a60b40e591e2072d753e7734ffe3b90b52b5a0807c4f2643a99e"></a>

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

<a id="canonical-a37014d2dc5a14a30233dbeb1ed040d3ba8311ae4cacb80e358a04980e188aa2"></a>

## Direct properties — aws.not_managed.node_list.interface_list.network_option.site_local_inside_networ / b0550c15f937 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3377e484603ed9396e4d897a634a42f547626fba91cc86df9cada9d56db7ea14"></a>

## Next pages — aws.not_managed.node_list.interface_list.network_option.site_local_inside_networ / b0550c15f937 / 4

- [aws.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-6fd4cb11a8978a039b74b2749c52ea22393431638d4f5f76696e302e7b839f4c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0f245f02e688690308d58c68005f307414f47c2f96d76f6a4917eeadb0fbdf85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de6930f324918f1390a56ce73fc3c14472b656ab031322127dd494cd9b6822fb"></a>

## aws.not_managed.node_list.interface_list.network_option.site_local_network — aws.not_managed.node_list.interface_list.network_option.site_local_network / 94fc5b2ace47 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-6fd4cb11a8978a039b74b2749c52ea22393431638d4f5f76696e302e7b839f4c)
- aws.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-8dea779a365b3161aa505453e79ddc42362e9a5b147a0484fa5dad98a9819c4e"></a>

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

<a id="canonical-21b28a047503562b5082011b6472bc3a198f4e1465d3dfb1fc466d850c8fba1d"></a>

## Direct properties — aws.not_managed.node_list.interface_list.network_option.site_local_network / 94fc5b2ace47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8127424c69a6c0353be1b8e372668aedf100525c756084a21054dcbf629acc1"></a>

## Next pages — aws.not_managed.node_list.interface_list.network_option.site_local_network / 94fc5b2ace47 / 4

- [aws.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-004.md#canonical-6fd4cb11a8978a039b74b2749c52ea22393431638d4f5f76696e302e7b839f4c)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-61d8d25fe8aa1236e6f3acbe35f59f0f78b04f3270e3e9f635aea547094b57d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bfa7d9977cfef3a359ed8333ba9f29a47e403e7412f93eb40488fdb33759a67"></a>

## aws.not_managed.node_list.interface_list.no_ipv4_address — aws.not_managed.node_list.interface_list.no_ipv4_address / a84e7cd76fa5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-9239687eedbbc7e1b67ef200fa3eb362244565599dee8b791cb21ec145cb3d3d"></a>

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

<a id="canonical-b5f6a51048ea84b7bc1f854cd2bfe5d56eb2d5a2024e948b1d6a7413bf87056a"></a>

## Direct properties — aws.not_managed.node_list.interface_list.no_ipv4_address / a84e7cd76fa5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d35148b6b733a85edc7a07992cde16f757f3925e028352956389021b0c3c7869"></a>

## Next pages — aws.not_managed.node_list.interface_list.no_ipv4_address / a84e7cd76fa5 / 4

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f14ecd1798cb09ac80b4c511a01bd244ee5f3430b209847f0ba2883b6a0fbb91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98dc341326c2b49cd35cc366723827efeb5e81302aa2ec64e29868597d34b539"></a>

## aws.not_managed.node_list.interface_list.no_ipv6_address — aws.not_managed.node_list.interface_list.no_ipv6_address / 3285bf169ae6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-dd86fe2f3df75b1e2d8e44c5533f9dcad0fd2485b68daf9c071cd6adb86d3972"></a>

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

<a id="canonical-dac245de6c5429b6c72c7da4152ba8f5bd7feba2ecc8a38657f74a5a90d13e2b"></a>

## Direct properties — aws.not_managed.node_list.interface_list.no_ipv6_address / 3285bf169ae6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e4d34cbc416bd1d6daa7ada9a4faa30649777a6794e52a9274836a6b54e5de1"></a>

## Next pages — aws.not_managed.node_list.interface_list.no_ipv6_address / 3285bf169ae6 / 4

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a7423d9370303fe286670b603c3566e4ff7b5a86a100308d39f313a6d1c4f926"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4846abb9d2abdd936fa6125379f05d488f2cabca12676620ec2e88b2ebc3890"></a>

## aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / 880df037fa86 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-5d2af599d12de5950be480c932f8562e3c431545c8cf6f098cd6935c2eb7024f"></a>

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

<a id="canonical-9bea7ab85de4d675e70014c1941008e6ae3b3eee5c3e051b58a44ad33393bbbb"></a>

## Direct properties — aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / 880df037fa86 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6dca05d21ca978cad8c5732bccf001e396123a4924f822893bddcb2bfdf6d35"></a>

## Next pages — aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / 880df037fa86 / 4

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ecf2f498087a2ff2077b099ed40e98413617e53b7df7ea0af604272c1b5fe71d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a20abb3372549531e188a311ec856467e73187983b0f6bd6489828b929f5fad"></a>

## aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / fb5211554a95 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-e206b5603dd526609224d3025ea06f4ef8087f2ac24b027b6a4c0f72883be29c"></a>

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

<a id="canonical-a2c671b41ec459f82e8a35d9270162a5bc9806158c23777be1f73a0ba36025a8"></a>

## Direct properties — aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / fb5211554a95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86a1e6316ae7d75fac7087decdec2b3565ad37204bd101e3d8e3b9031f2ae924"></a>

## Next pages — aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / fb5211554a95 / 4

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a2986406fc71975325d46ec91b796b7ef95f62999f84f4e2a62c84aa33b9c73a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e7d2ccce48d341134258da35fc1185b1dd25de9af9360932fe6f34bd939e48c"></a>

## aws.not_managed.node_list.interface_list.static_ip — aws.not_managed.node_list.interface_list.static_ip / b03096eace24 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.static_ip

<a id="canonical-f08bb951fd491694a0b26af293d9d969b05c79516e89ed56f2597605f9149b5b"></a>

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

<a id="canonical-ee781092e27636918165d0ca1c5474e156f9b676c0730226679cedc809beaf29"></a>

## Direct properties — aws.not_managed.node_list.interface_list.static_ip / b03096eace24 / 3

<a id="canonical-6baba347eefec92e2283c66c4889e3fd0cc4860b566f888e6f738ca31353593f"></a>

<a id="canonical-33ea7729ef23da6c232575f3a8b8bf4b7fe2db0140468988dfb9913597f9c61b"></a>

## default_gw property — aws.not_managed.node_list.interface_list.static_ip / b03096eace24 / 4

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

<a id="canonical-3ab70a5567b21d22c92ce4222b5cc4fcce76a3a49df6eb0c278b21a1fcc44057"></a>

<a id="canonical-e94086e88563a5999a2808253e0b2dab718d12c008f3895d3768082a0fa19e8d"></a>

## dns_server property — aws.not_managed.node_list.interface_list.static_ip / b03096eace24 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-edeaf55bd3f73a1a1485023d276ff94f818f3eebf51597e4bda7bc82810e570a"></a>

<a id="canonical-3bbfe2f7c12d24abd59da7df11e72b35f701383192b8919c0e0cd404c6d39c35"></a>

## ip_address property — aws.not_managed.node_list.interface_list.static_ip / b03096eace24 / 6

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

<a id="canonical-a567d4d9eb62100499384d5201196bb59195d1df6a8e8c646231ebad58ffa1ed"></a>

## Next pages — aws.not_managed.node_list.interface_list.static_ip / b03096eace24 / 7

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f83b37045d398833196e06e4a8ac3898b15f6b9993c02b2132ce3d11064303ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05464ffbfa8d4792f41d86f7da918c82800ebca05b7ce492d968cfed803b97e3"></a>

## aws.not_managed.node_list.interface_list.static_ipv6_address — aws.not_managed.node_list.interface_list.static_ipv6_address / 89301ad87c51 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-1d69072a730404d8a7326d231f28e0461dc9af74cf83aec3e0c6fc7a84627d36"></a>

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

<a id="canonical-192a83d28c2c7aaccd1d32dde8cf7ef7189933818f7cb2cb7660b5ec696273b9"></a>

## Direct properties — aws.not_managed.node_list.interface_list.static_ipv6_address / 89301ad87c51 / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1ff6ee8a8fb986b11e68fdbe44e5b57d5daa865953d0289a7db378ae0d45178b): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-5c1caed8dbe81a635d2a8618e9fe75e6036b98c1dd730fd488d3407a35aeca7d): complete subsection reference.

<a id="canonical-4add2656f05060c91f854cd821e0f86989416408e0b207ec5b46681736e7e475"></a>

## Next pages — aws.not_managed.node_list.interface_list.static_ipv6_address / 89301ad87c51 / 4

- [aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-1ff6ee8a8fb986b11e68fdbe44e5b57d5daa865953d0289a7db378ae0d45178b)
- [aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-004.md#canonical-5c1caed8dbe81a635d2a8618e9fe75e6036b98c1dd730fd488d3407a35aeca7d)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-1ff6ee8a8fb986b11e68fdbe44e5b57d5daa865953d0289a7db378ae0d45178b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26ae8d0c103a86d99b673dac4c132f078ab2a2f7c6c6d32b58763785b163f1b1"></a>

## aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / af24ed82ca07 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f83b37045d398833196e06e4a8ac3898b15f6b9993c02b2132ce3d11064303ff)
- aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-4be94a5db48e3888ebc5f4c43ba83c24283fd0c03d933bd98c5e5daf8851ed31"></a>

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

<a id="canonical-fdc496a4f03088b7a159180355ab140a8f2dde70c946027d5db5a6e924bb34dc"></a>

## Direct properties — aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / af24ed82ca07 / 3

<a id="canonical-7244a7d60768e4acb3ccb7613c6b723af8d2ad7940f845b78c3a90227fefcac5"></a>

<a id="canonical-0467ad0de3970573a1cbb516744a3b075e04523d02a212681de880b8815e1660"></a>

## interface_ip_map property — aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / af24ed82ca07 / 4

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

<a id="canonical-6c08fa57f83c5ed42202e4c050dc9ed9276a9b6829c975ea730614080803b08a"></a>

## Next pages — aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / af24ed82ca07 / 5

- [aws.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f83b37045d398833196e06e4a8ac3898b15f6b9993c02b2132ce3d11064303ff)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5c1caed8dbe81a635d2a8618e9fe75e6036b98c1dd730fd488d3407a35aeca7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1132acd6a71066de9b27d8dca591be3a06c4b35f4961d93d6f77e72a348d6f8"></a>

## aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 21e2bba7c464 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f83b37045d398833196e06e4a8ac3898b15f6b9993c02b2132ce3d11064303ff)
- aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-f9cf64fe8953e39407421d9fe30b44fa578321c61d65c72742b962c9e8e33a74"></a>

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

<a id="canonical-d21aab8a098c9656e721d67e921f3965eccef93d799cca4531ef802fb25ab7fc"></a>

## Direct properties — aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 21e2bba7c464 / 3

<a id="canonical-695bbe309d5d6247195d4f7706b44064138b5cded7fbd6bdf04d2806c5ad821c"></a>

<a id="canonical-157f98b6fb064c5e87d99adeab0440c7b30258f6498f8a8e666b7f38a0c28a35"></a>

## default_gw property — aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 21e2bba7c464 / 4

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

<a id="canonical-7a4788a66230dac9187ca22d7adda8c7d5bf098d76b1f462f0dc2d37e393c68f"></a>

<a id="canonical-0ed9eb59bd40762155d999e422fb29523034902473bee138a4823c743a93483f"></a>

## dns_server property — aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 21e2bba7c464 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-b9581d754c8e955527cf90219efe29ea26ccd2e3d0d651211fa25ef59d269e20"></a>

<a id="canonical-ccecf8464b8d9e0bad561e9b6f84e755f946aad1a3650a0214cd1914cce994b3"></a>

## ip_address property — aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 21e2bba7c464 / 6

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

<a id="canonical-ab3bea8a5bf700d6de6a2b91ccd25f40dd2b1499775f9941b5a7db6381738c15"></a>

## Next pages — aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 21e2bba7c464 / 7

- [aws.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f83b37045d398833196e06e4a8ac3898b15f6b9993c02b2132ce3d11064303ff)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ae4737053069c6fef88cc461a1bd82687816da9d46ae8b4afc746d36098e1e8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-001b3cbcf0a3ae069a27296c350e8025ec4981c49101d750c72f530775f34855"></a>

## aws.not_managed.node_list.interface_list.vlan_interface — aws.not_managed.node_list.interface_list.vlan_interface / c58967e3cc39 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-62d5e21dcfa4f3b673848d97041cca00221aead04682eaa3e75a721f89c53efd)
- [aws.not_managed](data-sources--securemesh_site_v2--reference--group-003.md#canonical-9e877b450151e45f8031915280d2cc7e4438aa897be0937e31d80ebc197d4c74)
- [aws.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-3dd518bfbe188f49a40fd3d6109f7941986faaf44f07a34009db964c8fd72637)
- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- aws.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-f71134d3319f5de1637ae382300f674820f3c9522e25610500b8116d40cd0382"></a>

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

<a id="canonical-5eb23ac44d711997c416f4426bea8bd14f770d913df28e21b3e67142f1913c05"></a>

## Direct properties — aws.not_managed.node_list.interface_list.vlan_interface / c58967e3cc39 / 3

<a id="canonical-56e9cda959d3f232e4d5c90a426f7b5a84d478dde17c20f746791cf05969fb7e"></a>

<a id="canonical-a2c4ba776cdf57afed4b376ea85a8e1d87288c8a7a73c965050e05bdf64ea10d"></a>

## device property — aws.not_managed.node_list.interface_list.vlan_interface / c58967e3cc39 / 4

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

<a id="canonical-737846a7ecc39feb2df8a36db52e375f149f508bc4c00d195636e0fd6c170c29"></a>

<a id="canonical-225b6debfb7aa17ed56a240bd877d07ce027734b27ff6d54d70522f2f827c090"></a>

## vlan_id property — aws.not_managed.node_list.interface_list.vlan_interface / c58967e3cc39 / 5

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

<a id="canonical-d2e224ce2f04f93328936811e070a4cd0c85379389c7dbab12f3e81eef4a4f2a"></a>

## Next pages — aws.not_managed.node_list.interface_list.vlan_interface / c58967e3cc39 / 6

- [aws.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1986b0335bb03c8d16100141de8484858cb687890657c98ecb716b831b69358f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4354da929131770d377bd00f951f2c3d036e7cd9987733acabe201aaf5c618f2"></a>

## azure — azure / 623aeb1dec99 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- azure

<a id="canonical-ea94d149342ac9231b16e09e3cf9bf5ae99c7440ee415e56b87e1e9d99ac1f91"></a>

Type: `"single"`. Computed.

Azure Provider Type. Azure Provider Type.

Upstream description:

Azure Provider Type.

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

<a id="canonical-8b50f1492b71f4f1a8ada0a2f1761996deb5e0310def97002e9d10b81e799139"></a>

## Direct properties — azure / 623aeb1dec99 / 3

- [not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf): complete subsection reference.

<a id="canonical-eaf5e578d0cc9bee1bfca238872cf76362d76fbc217dcbe5e7fccaf42897fdb0"></a>

## Next pages — azure / 623aeb1dec99 / 4

- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f9969d92c867dba094073d99bf707509507402ba3e4ed9a830e084988f185f8"></a>

## azure.not_managed — azure.not_managed / c232abe26f20 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- azure.not_managed

<a id="canonical-4cfbf4b059ec1c4f0f70ea55a9dadcc767003039c95d4feef4f370f0bae81a82"></a>

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

<a id="canonical-726a94036d78817a360823edd54a1724630a1c15a35b276a94bb998f17dfda4d"></a>

## Direct properties — azure.not_managed / c232abe26f20 / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b): complete subsection reference.

<a id="canonical-7ba577e2cf364a84039c1a20a3b856e7187aa34cd072aa629f879926b22e75a5"></a>

## Next pages — azure.not_managed / c232abe26f20 / 4

- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df37f70b0d1c174312f12f367688c1c4b27f1ce6a08c4425e646ecafc0966b08"></a>

## azure.not_managed.node_list — azure.not_managed.node_list / 55f82bdc4500 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- azure.not_managed.node_list

<a id="canonical-9c1da7b54a83ceede44cbda16f90b817c0de25b62b546c96adbe982700345561"></a>

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

<a id="canonical-956dfaf3611a8ff1f5c26cb13310948626b1561844f32f475d8aba5fee727551"></a>

## Direct properties — azure.not_managed.node_list / 55f82bdc4500 / 3

<a id="canonical-1b81db9c2df533616ffd33930e66925a16d396a4af7921487fadeba127da53d0"></a>

<a id="canonical-ad89206910b6c9de14fd68b71c6a6aa73b12b0eb6ca5f424fdbadbb0c59aa840"></a>

## hostname property — azure.not_managed.node_list / 55f82bdc4500 / 4

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

- [interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4): complete subsection reference.

<a id="canonical-26d01f50e53e31ad53a7fe0b451a4de74918d8d5ea0c6da4f45b435ab23b2885"></a>

<a id="canonical-e911cfb02b7eda2f4b623f21de95145ceca9b6d7476fcccb8de5fe480b7e31c8"></a>

## public_ip property — azure.not_managed.node_list / 55f82bdc4500 / 5

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

<a id="canonical-afb64f82bbb3a492d7f19b6e5e79f769e3f54c9e96e9df4e2b639e12695dead1"></a>

<a id="canonical-74c46c5fc87d1da9685b9923235f6ca7fc81bef29b50f7f21ae20d4a79eb280a"></a>

## type property — azure.not_managed.node_list / 55f82bdc4500 / 6

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

<a id="canonical-e5bb559e4f1d49a66b799e539ad9aa0769b17ce94c8a49dbacde7d272edd8458"></a>

## Next pages — azure.not_managed.node_list / 55f82bdc4500 / 7

- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2b41ba46fc3a19f450149ce64c53d4f108c0af569936612cb2eb00f6323de53"></a>

## azure.not_managed.node_list.interface_list — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- azure.not_managed.node_list.interface_list

<a id="canonical-77a5c713028b49626308e52c9f923872f59a4a7c22df89127a4a8a6e2a3013b0"></a>

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

<a id="canonical-849656a1ae2d5bc71dbcf9220a404f12e6ad11f8f5bfac581862b5a600c79b9a"></a>

## Direct properties — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-762b7593925ee8be744996f68862e495cc84e1095de99535dcd073d1dc837120): complete subsection reference.

<a id="canonical-41fc2404a5dbf6e6c8a9795ceef93c76e973370f2d4060c7f55bfe15587c00fb"></a>

<a id="canonical-cadd2b624e62461e6fab62222bdf4e6b4e65379218b35dc9c15f641a6ba76e67"></a>

## description_spec property — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-004.md#canonical-ec2f1007869b26454342c797bebf002abcb0107f58a7ded003151f4483372c97): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-090f985a66d65a3f7bc47d22c062b2f3867a88900c4e70be4085804b9d8687e1): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f28881d424ce9e5b22535af335b73f96509f38266e92d88c0c6b16aa5d5f97f0): complete subsection reference.

<a id="canonical-61416a983cfb329c5ae2f6f9b5c8a214f0f102671c41aa3d855960bb33ef420c"></a>

<a id="canonical-37185cc8608edc3f87fb5afcba3d9f53fb5d5d2e641a725830974d8ba8d1fbdc"></a>

## is_management property — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-7677474e7d44d371131b65eac6bf908a1658ea8134088992fdff6ec01d3ac803"></a>

<a id="canonical-9d2c05fcb2496c9ddec8b203e9ac57ae43486997d5c9687e3f947b3e8aa5f76b"></a>

## is_primary property — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-e3e43710b39465e84cb55c855961dcf80046bff460eb899451605207052ad852"></a>

<a id="canonical-adba4c9fc3087f58dfc5c66d0cf9530c3132e7860e6bc58e3826802f816c44ea"></a>

## labels property — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 7

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

- [monitor](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5ddf703fe5bc40836cbb04234626a21f97024ca8b26fb47ea1c894bd0fa8313d): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-005.md#canonical-7043227ef39fe0fab7950c13d8e1eec407cb357a8cace441f83bc6b90fc2a7f8): complete subsection reference.

<a id="canonical-9a9a396559ed2986857cb02c2a3ba298a78966e2ddb6a2d20936b2d642e3b139"></a>

<a id="canonical-04e8b1c4688d52b2eb5f57f41f07272801bcb69a08df407fa53bb9cfc22728ed"></a>

## mtu property — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 8

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

<a id="canonical-22942ea9902a5cbdd7627689465deb56acd5e40d8bcbf9fa7b8ae601c6b506cd"></a>

<a id="canonical-f6873cdbe8b733adcc384e8956a8b7dcf5ee849cac9936d495260ba005510518"></a>

## name property — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 9

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

- [network_option](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5feadc0aab7a9599e78466cc47e4e2f6d7d508a415aad723a1f92b9b1f114587): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-73f0da8cba75b9e5252ab6fca046fa47316d79544fdbf28eca5dacdd10d69a56): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-6cabf27523313c6ef6b48eb3ade98a9646dfbcdc150879ff144617fea98af889): complete subsection reference.

<a id="canonical-662aa51b81991d33f14851a39b5fa745f9569a1c1b527a635f9a1b5de833f880"></a>

<a id="canonical-7e3f883a1d641ada6bcfc22f00112053cddc962dcf5afc86d388a11afa1a18ff"></a>

## priority property — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 10

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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-005.md#canonical-f6c130164366cef04494034ee60c5c5921e74e3f97f09f07979127aff4aac0a3): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-005.md#canonical-7c750903b25966510fdcde5251090b324d52c0a9d1f9d9f04c012f7ace19033b): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-005.md#canonical-96acb7525e7ac27fdf0a0517fcb4832b634a3ae91e4d0f1530adb221e3e72711): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-6562fb6314edbbd084c0c9ad8c060494260c971567b05fe585a34808e545fcf0): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-005.md#canonical-f2e8113760599f5c8a8e677df3947ae1e6d067f1359a39f5a252a070184b8583): complete subsection reference.

<a id="canonical-0ffc92e2225b8fdba9d88ca45f78671b06c769fe7267edcc7fd6457d0c842de3"></a>

## Next pages — azure.not_managed.node_list.interface_list / 938ad6d7faf3 / 11

- [azure.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-762b7593925ee8be744996f68862e495cc84e1095de99535dcd073d1dc837120)
- [azure.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-004.md#canonical-ec2f1007869b26454342c797bebf002abcb0107f58a7ded003151f4483372c97)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [azure.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-090f985a66d65a3f7bc47d22c062b2f3867a88900c4e70be4085804b9d8687e1)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-004.md#canonical-f28881d424ce9e5b22535af335b73f96509f38266e92d88c0c6b16aa5d5f97f0)
- [azure.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5ddf703fe5bc40836cbb04234626a21f97024ca8b26fb47ea1c894bd0fa8313d)
- [azure.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-005.md#canonical-7043227ef39fe0fab7950c13d8e1eec407cb357a8cace441f83bc6b90fc2a7f8)
- [azure.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-005.md#canonical-5feadc0aab7a9599e78466cc47e4e2f6d7d508a415aad723a1f92b9b1f114587)
- [azure.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-73f0da8cba75b9e5252ab6fca046fa47316d79544fdbf28eca5dacdd10d69a56)
- [azure.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-6cabf27523313c6ef6b48eb3ade98a9646dfbcdc150879ff144617fea98af889)
- [azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-005.md#canonical-f6c130164366cef04494034ee60c5c5921e74e3f97f09f07979127aff4aac0a3)
- [azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-005.md#canonical-7c750903b25966510fdcde5251090b324d52c0a9d1f9d9f04c012f7ace19033b)
- [azure.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-005.md#canonical-96acb7525e7ac27fdf0a0517fcb4832b634a3ae91e4d0f1530adb221e3e72711)
- [azure.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-005.md#canonical-6562fb6314edbbd084c0c9ad8c060494260c971567b05fe585a34808e545fcf0)
- [azure.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-005.md#canonical-f2e8113760599f5c8a8e677df3947ae1e6d067f1359a39f5a252a070184b8583)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-762b7593925ee8be744996f68862e495cc84e1095de99535dcd073d1dc837120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18dbf1dc93bcef253363fee28a335d5ef50e8b258252a69703eddeb222419e93"></a>

## azure.not_managed.node_list.interface_list.bond_interface — azure.not_managed.node_list.interface_list.bond_interface / 534af82ab5a8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- azure.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2a147396789d6f608ddca63b1ee2e5bcdcace13bfe3bbd0cce8decaf33faa020"></a>

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

<a id="canonical-7e3821c54b14df8873be2d74d8da7e2e0dc48a2b516514a131d54f4101e1d6a0"></a>

## Direct properties — azure.not_managed.node_list.interface_list.bond_interface / 534af82ab5a8 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2705aec5e7d10380fa9e206b2b6b751ab1e28f2e3ff450b052126613db5d683a): complete subsection reference.

<a id="canonical-45c08178f02658adfd3b0f4b452c3368bc05ccdc6f534dc5d76c2d944a775d37"></a>

<a id="canonical-ff3babd61567c8e3f8ffa8ecfcbbb7afc63355a7bd8277701fcce9ccb1507630"></a>

## devices property — azure.not_managed.node_list.interface_list.bond_interface / 534af82ab5a8 / 4

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

- [lacp](data-sources--securemesh_site_v2--reference--group-004.md#canonical-849d387d6700ca86cc20795bfb4bfa6d8d5866e9119ac686b01158d7a232f0d7): complete subsection reference.

<a id="canonical-014ba048fd2c30fff4e84108901ddae1a249cc1dfb39c1c2e1d992497d244f1c"></a>

<a id="canonical-3e5342356016d26c5774a8f2838461b64a573a438f14c3f2b0ee0b20ab070b33"></a>

## link_polling_interval property — azure.not_managed.node_list.interface_list.bond_interface / 534af82ab5a8 / 5

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

<a id="canonical-ca77c97c1932d0837a5681ba7b971454615db3a922d0eb81e81dea0a45044c7e"></a>

<a id="canonical-c3f75dd5b133a0f7eb045eaabd3d488b6bd34265fbb75a45895c6cbec0241e5a"></a>

## link_up_delay property — azure.not_managed.node_list.interface_list.bond_interface / 534af82ab5a8 / 6

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

<a id="canonical-33e02a9450ddd2673e816fde9b142bab490367eb3b50015691e31fae45449f2c"></a>

<a id="canonical-169fb11b363477cf0ed4562ce09d0aaaedafaa928968643a49b38c2076c7bb20"></a>

## name property — azure.not_managed.node_list.interface_list.bond_interface / 534af82ab5a8 / 7

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

<a id="canonical-8260b82f86643b7ea9059fa67301ce9835b62005c72dd37b701fca2cfab698e4"></a>

## Next pages — azure.not_managed.node_list.interface_list.bond_interface / 534af82ab5a8 / 8

- [azure.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2705aec5e7d10380fa9e206b2b6b751ab1e28f2e3ff450b052126613db5d683a)
- [azure.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-004.md#canonical-849d387d6700ca86cc20795bfb4bfa6d8d5866e9119ac686b01158d7a232f0d7)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2705aec5e7d10380fa9e206b2b6b751ab1e28f2e3ff450b052126613db5d683a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-491eb839dd3f526d73b07f70c981d5e960497f253bde465e68aadc0f44fcdc76"></a>

## azure.not_managed.node_list.interface_list.bond_interface.active_backup — azure.not_managed.node_list.interface_list.bond_interface.active_backup / c78c231d4528 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-762b7593925ee8be744996f68862e495cc84e1095de99535dcd073d1dc837120)
- azure.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-191d48e54f57520df8702eb6c4ac0b0f27c5b95759ac63e6b7676eae08d07e41"></a>

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

<a id="canonical-d31f9416c93352d930476aa52a0a60553a269f3463aa98556f575eb2ae5b6ca7"></a>

## Direct properties — azure.not_managed.node_list.interface_list.bond_interface.active_backup / c78c231d4528 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f10aa15ba0ee0909b77ff142c9e9692a8373c27b52d9f506f9a7759422597c6"></a>

## Next pages — azure.not_managed.node_list.interface_list.bond_interface.active_backup / c78c231d4528 / 4

- [azure.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-762b7593925ee8be744996f68862e495cc84e1095de99535dcd073d1dc837120)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-849d387d6700ca86cc20795bfb4bfa6d8d5866e9119ac686b01158d7a232f0d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-021aabeee68858f85d7406e6dcfa6f990777b639310a2e160e63060e5711f3f9"></a>

## azure.not_managed.node_list.interface_list.bond_interface.lacp — azure.not_managed.node_list.interface_list.bond_interface.lacp / 8a3125cf2a8c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-762b7593925ee8be744996f68862e495cc84e1095de99535dcd073d1dc837120)
- azure.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-521f66035d1fe8360b2684214cef32b688e8be67bfe98dd9eca59801a433c044"></a>

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

<a id="canonical-7f63af2289b6d5787b4810713ddd67ee6263abeb680c62446ba33870989b1101"></a>

## Direct properties — azure.not_managed.node_list.interface_list.bond_interface.lacp / 8a3125cf2a8c / 3

<a id="canonical-2046ad353cdae0fff7a358e4b905363a4893299022c67e07a8a526120b01ed6b"></a>

<a id="canonical-f1d86761fe1361c66a43280ef46440497c8e3da645ae1c19e56c882ec9121a27"></a>

## rate property — azure.not_managed.node_list.interface_list.bond_interface.lacp / 8a3125cf2a8c / 4

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

<a id="canonical-3ad390545eb16da8539b1ec7e1f2cafd4af7cd4bcad3db76b8534955798c4dcc"></a>

## Next pages — azure.not_managed.node_list.interface_list.bond_interface.lacp / 8a3125cf2a8c / 5

- [azure.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-004.md#canonical-762b7593925ee8be744996f68862e495cc84e1095de99535dcd073d1dc837120)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ec2f1007869b26454342c797bebf002abcb0107f58a7ded003151f4483372c97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f66e9bf4912d88cd37eb9a65d315b008e1f7c5d399a06d17a66e44d04b0e9386"></a>

## azure.not_managed.node_list.interface_list.dhcp_client — azure.not_managed.node_list.interface_list.dhcp_client / c3d8917d074f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- azure.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-857bcd700814400a98c8e3dd96f7b8c13c25a9d0f2ff472d9921ea73b061fd08"></a>

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

<a id="canonical-40576a1d5a8f5bfc548069614433fe8322e2c5b759b74d042ac1205f9798a6b0"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_client / c3d8917d074f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46909fa75129cbdb1e764666d701c28c6cc535e3dfbf1d7c81582fc96ce4662f"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_client / c3d8917d074f / 4

- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdb8d29701a107915e7be91dccb13f1478be632ed9f6bed942f9dd6b8ae5edf4"></a>

## azure.not_managed.node_list.interface_list.dhcp_server — azure.not_managed.node_list.interface_list.dhcp_server / 065473fc9824 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- azure.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-162cc1b43aee1924c434c419e08d14d8d8b440b0e8ffd0e2e36402cd16367f7e"></a>

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

<a id="canonical-feb5868b1386e0db264c72445bb838e953f702bcfb0a009af3f403cbb6fbae85"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server / 065473fc9824 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-004.md#canonical-edd8b7caf8b314140ad0d097f95ab8e41bbf5d7182e0016b2d6136a4f5131b27): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9279070dc41101cc03d9abbbbad361970df0fbaa7418a73ca5a5176c5432053e): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3): complete subsection reference.

<a id="canonical-0a0d7b05ea7e941253115b013e52fd126a0234d59cce8d42cf7623697a4c896e"></a>

<a id="canonical-5361a84b82fa698d0a27f29bc709b16a68b3b9d57505a62c758add66fab03eac"></a>

## dhcp_option82_tag property — azure.not_managed.node_list.interface_list.dhcp_server / 065473fc9824 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-df1f2c5ef48535430e4754e9447cff9883e3ae9222eba535a13081a4ba150c14"></a>

<a id="canonical-32fbc29a3f6f8225f21cd946a0dcec48641032ec125886fd07f4be713728b8b6"></a>

## fixed_ip_map property — azure.not_managed.node_list.interface_list.dhcp_server / 065473fc9824 / 5

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-004.md#canonical-278edb6bddc8da55cf5368a99fe34416b2ae1030d13ac81cdc586a86aa03a255): complete subsection reference.

<a id="canonical-4bc987a8f7b7cf1187da78fa6717dd874d233ddf28128874d2e7d03ce5e791e4"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server / 065473fc9824 / 6

- [azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-004.md#canonical-edd8b7caf8b314140ad0d097f95ab8e41bbf5d7182e0016b2d6136a4f5131b27)
- [azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9279070dc41101cc03d9abbbbad361970df0fbaa7418a73ca5a5176c5432053e)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- [azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-004.md#canonical-278edb6bddc8da55cf5368a99fe34416b2ae1030d13ac81cdc586a86aa03a255)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-edd8b7caf8b314140ad0d097f95ab8e41bbf5d7182e0016b2d6136a4f5131b27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2935074effb89af9d15765a356200a0a1670eb686f2d86c153630db86d0d9f29"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 3479563935ee / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-2c386d6e6bd4cecdc60f20538a9a9a0a023e030b294ecb5a289e8ee78f3d4ce5"></a>

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

<a id="canonical-a4ad80c45bda694d76099dbbdaf4f34adf9ef2c96a73c5cdee1c96e155b54657"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 3479563935ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc63a084414affa28ed581f2c6763e289759d9b82eabcb30b480488a97f2acfa"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 3479563935ee / 4

- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9279070dc41101cc03d9abbbbad361970df0fbaa7418a73ca5a5176c5432053e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c647067644fa7afd561235430687972b812ec1c5dcf10884f2bae4ccd533ecfa"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 208b51f79e21 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-dd0a3d74de53ea04670b417d1343cf34d593e37b93b253c8a760db87774509c6"></a>

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

<a id="canonical-68c2b0b89ffb7a9ee68f6ebdb251985c58cb8ef80a24d2c9c710f7bd8916546b"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 208b51f79e21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-497a24cc6f129f75a2a3fce01c839b55e77644cc4897b8c7f8c4dcb605074d11"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 208b51f79e21 / 4

- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baa005335359de4babe68f53b741ce99045f7a027b3dc7545d76ca1a13e0501f"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / c8691933e992 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-4c4003e72de7748f0543e012a1196f159b490a391d7d5445a809aa023c1e1d72"></a>

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

<a id="canonical-4dceea4470892657ad81aa56fee63ecfe7d95bb3e9d88a9d41fc5ffd43d9a1a4"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / c8691933e992 / 3

<a id="canonical-9538dcce8a0a297ed7a4543c9add9a24d87949061cb115501b2c451b252cc353"></a>

<a id="canonical-c159f51acc724804ec788373e44b1db97a984f0678e5da3242173abfab12cf74"></a>

## dgw_address property — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / c8691933e992 / 4

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

<a id="canonical-57c4391fd0137cd0944c8b96d923283b95087a5cd682c1a32046ff6b917b6bb7"></a>

<a id="canonical-bcee99ab2e16bddde3dd6e9451867e804c4c318e776421f1ce90d38bf938ec61"></a>

## dns_address property — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / c8691933e992 / 5

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

- [first_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bf398ba71d67f6930cd5f2134e8ed9ff54e6b54b9c51bfe31d61bfa1fce9190e): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-a33bab602483fbbd9ba1ce750187f04c7f8469831d51b6ce886ddfad8afa56c3): complete subsection reference.

<a id="canonical-ec6794675e5034c73e401019f288112cd73d7c9abf2a3b7ed4e734038ecb39c0"></a>

<a id="canonical-35ec017d9e43a7c7dfd4a85600aa12b4ddfe24dc259a30bb330fbb57e2cf4611"></a>

## network_prefix property — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / c8691933e992 / 6

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

<a id="canonical-4b176a63761a999ee33399835404ea8ddb43e2cc944d55e2508c50969e0c1ed1"></a>

<a id="canonical-008928988bd3264ba0dec43983f27fe61e1df96b9f5c13c8eb2475149e341dea"></a>

## pool_settings property — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / c8691933e992 / 7

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

- [pools](data-sources--securemesh_site_v2--reference--group-004.md#canonical-b652ba6948599502bf1091be7bb174f9685f1f1ac6b56e105ec016be4f815564): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-004.md#canonical-e1a41012aaac162747b0550232352f09e654a541e9cda44f232ae5a587306e05): complete subsection reference.

<a id="canonical-3b8cf4f06016d6c66c25c75ee5e38532496901aebcb766c83fcf4183d56444df"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / c8691933e992 / 8

- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-bf398ba71d67f6930cd5f2134e8ed9ff54e6b54b9c51bfe31d61bfa1fce9190e)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-004.md#canonical-a33bab602483fbbd9ba1ce750187f04c7f8469831d51b6ce886ddfad8afa56c3)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-004.md#canonical-b652ba6948599502bf1091be7bb174f9685f1f1ac6b56e105ec016be4f815564)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-004.md#canonical-e1a41012aaac162747b0550232352f09e654a541e9cda44f232ae5a587306e05)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bf398ba71d67f6930cd5f2134e8ed9ff54e6b54b9c51bfe31d61bfa1fce9190e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e30aa6c5f0a2459fb5264e20a441d0401aec3fcd72e40bbc8ad2e80df804662"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addre / a9b96971bdc6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-13b2a6b1428708bf42df91ea1462ce4cb0fecf953c5693967cdd2eb517add3bf"></a>

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

<a id="canonical-e11497fa451e7766813b5605b5791b60b1b12e7812a84329dcf0cb04152aaa14"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addre / a9b96971bdc6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d369a1afbd00371e384bb18ddb92e90f48b28cb26f7e890b19334a9c5946767"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addre / a9b96971bdc6 / 4

- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a33bab602483fbbd9ba1ce750187f04c7f8469831d51b6ce886ddfad8afa56c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7634adda24480fb4a053da5e3c48ea8c1cace582a00e7eeb3bd4fe1802f8abbc"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addres / f04e242efcec / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-a2d8feee8310512286cc8a5f31bf637f865240d6345eed1d247889533d307caf"></a>

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

<a id="canonical-96129a458091e0cf6cf178a86cf290a1831819647af0b859c1e19500079e54a7"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addres / f04e242efcec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6581160aac1370d6ed490178fd6993960c03058cb367d7ad0c0bb0d4aad346af"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addres / f04e242efcec / 4

- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b652ba6948599502bf1091be7bb174f9685f1f1ac6b56e105ec016be4f815564"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4909a7f2791eb95226adad571f36c271240a4c0950818703b9714e28a438889a"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / ff0e58f8689e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-e831ec958ef6a3a4d055c812b60d286415f05fa8e78ccb29cdac55b8288543da"></a>

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

<a id="canonical-59ebb5497974b5aba3430a41c159b51a54770f3acdcb331c93b7d79e678e72cc"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / ff0e58f8689e / 3

<a id="canonical-67fba329b57ccd2cc6355a5320ff973205e97bcaf145686d0761133ae81b6c67"></a>

<a id="canonical-ba64672cced124b7e58b6395b9a49f9ba74cc434d51f5377de09abc0668fc84d"></a>

## end_ip property — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / ff0e58f8689e / 4

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

<a id="canonical-ab438901de775a49346c742aebed72c3f25426d51c08839b7ebab0101316999c"></a>

<a id="canonical-ca89b5f8fe35476c82a8a204d3745bb4e7bbea75888d1942da3d932faaf2c006"></a>

## exclude property — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / ff0e58f8689e / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-d6728a0e166069eb755b0d59a9c294b1985047365cf3918bb9369a17009cf42d"></a>

<a id="canonical-9b2f562859c4a8f72f5d09a96de897c2491c221c88b6db16b1f735b75ee5116e"></a>

## start_ip property — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / ff0e58f8689e / 6

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

<a id="canonical-020e37084324e6d77e623cb71adb74ff36c87156a1dd9218f6dc22ae38ed1fa8"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / ff0e58f8689e / 7

- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-e1a41012aaac162747b0550232352f09e654a541e9cda44f232ae5a587306e05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-310214b5e039cfa170a499e3b5c160cd2a96cb501f5df65b96875e8f53335c04"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / e10092a324bd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-4e27ddc6b7b42b4f241ea400a6b6d4aecd01eeb0f7f066649932dce9a0dbeeeb"></a>

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

<a id="canonical-fbea57a4d69e9055bdb398acd3afa943fe759db153b348194774efa87c257138"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / e10092a324bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7cfbcde40389cd5430c6805ae6b2b0aafdd9c645c9d810744fe0360b824fa0fb"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / e10092a324bd / 4

- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-004.md#canonical-9cd18056dd68d1e8c32db8f14e1b07a2245e3e5861f34869844bd3d6de3403d3)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-278edb6bddc8da55cf5368a99fe34416b2ae1030d13ac81cdc586a86aa03a255"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d10a7c028db78975be91c80e7c505ae9f8b530dbf601249dca71ef22ed55d536"></a>

## azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 7243992843de / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1dc6568161674d01e71a9fb1c3a0f0ae1c872b694873705a59d4dbc5b22b2968"></a>

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

<a id="canonical-5472e7b326d4e744e046221b2b15c4588682812003e87b9998c4a87a1d0e32c1"></a>

## Direct properties — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 7243992843de / 3

<a id="canonical-a91efc3e9b810e2a068a867d2b7c24c9a180eb5a2f8c4c352265bfdf62f43d11"></a>

<a id="canonical-0487397dd451ec056883bbdb21ed5133d00d6ef95dd841af33384d7647be81a7"></a>

## interface_ip_map property — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 7243992843de / 4

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

<a id="canonical-19ffb9c11da9cc36c93874baf89c2881ccbf964dc664c90789a35c20e760e3cf"></a>

## Next pages — azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 7243992843de / 5

- [azure.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-004.md#canonical-2e92235e7a95e3021029385d2a6d579887bc2a74534a643c83e1aa422da184b5)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-090f985a66d65a3f7bc47d22c062b2f3867a88900c4e70be4085804b9d8687e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6094261aaa3fce8e48a349eacc201bd5f5c652ceaeff7fd3b80a7c167ab2aba"></a>

## azure.not_managed.node_list.interface_list.ethernet_interface — azure.not_managed.node_list.interface_list.ethernet_interface / 451de1e2a736 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- azure.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-177f1bc8ce8b40d1b47032ea197a47d8bcf8127bbe65d7dd6e8b184272835bf4"></a>

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

<a id="canonical-2e27fee3229948230480c90a7246d71650c1eda34c44bf3fb5e4b1bef1ddcbe9"></a>

## Direct properties — azure.not_managed.node_list.interface_list.ethernet_interface / 451de1e2a736 / 3

<a id="canonical-29ab2daaf2bf0ff850c43f692e6b9ceaeef792e3e7641fed08e5b3c9fb4ec3a3"></a>

<a id="canonical-3d806d9fda49ce7e62f64cad1a0f8afaa5e87139b8cf43a0e19b9101b4a89ff8"></a>

## device property — azure.not_managed.node_list.interface_list.ethernet_interface / 451de1e2a736 / 4

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

<a id="canonical-06e2579e960c3ea59217b723a264861e4af3bc67d568f74f7192ab5f3b9bca57"></a>

<a id="canonical-ebc346d966b89660545895a9bf54b94f4b1b6ec9eef9b240ebf036f4ab651eb1"></a>

## mac property — azure.not_managed.node_list.interface_list.ethernet_interface / 451de1e2a736 / 5

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

<a id="canonical-de33c16cfc6d2c76819183817df9cae08d57b3c95c5dfdfe7cdf5cf6c22e45a7"></a>

## Next pages — azure.not_managed.node_list.interface_list.ethernet_interface / 451de1e2a736 / 6

- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f28881d424ce9e5b22535af335b73f96509f38266e92d88c0c6b16aa5d5f97f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c55c9c2416dd253bac4a53c370432de7a23d7b18a9f415d9cd271c7ed6a1daf0"></a>

## azure.not_managed.node_list.interface_list.ipv6_auto_config — azure.not_managed.node_list.interface_list.ipv6_auto_config / 8626bdfbf1d0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0869f169673e4654df8f600e75a61b00bf722c049e0984768f1cd9a6da4c3cd9)
- [azure.not_managed](data-sources--securemesh_site_v2--reference--group-004.md#canonical-8c57a588fe2d9ae854064ebf0b1d9ffc288808c8afd766ee44a60fbe8ae800cf)
- [azure.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-90336d61d9a86dd447154148a4e83808c0e1c2c7155f7be267a675766dfa743b)
- [azure.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-004.md#canonical-d4d0b4b4f3747a478c3071f08b3d4bd724fe1c4b354d91dbcd79ed06529bb1d4)
- azure.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-dd8645020f054db18aeb01e34fa4773be1ee259c8f9ea804d6e4778434743843"></a>

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
