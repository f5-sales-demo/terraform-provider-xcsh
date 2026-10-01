---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-a35ff5782da43460191945315bcf8a4530d9b1a9a1de4e32cc8a51c4049c2607"></a>

## network_prefix property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 47378163445b / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0): complete subsection reference.

<a id="canonical-895a5ab586a7321231fce22ac928176f8690b391d03b261d99f3e0c41e60971e"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 47378163445b / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cee59fb39361f75f8c1442238d61589e67fac6e0ee85d2ef0b68f745f7b4da86"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8aa8d860ed04 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-74d20a562b1f52ec9ab36c827c9627afdb9e428a10d45d070f1fcb9b309ea606"></a>

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

<a id="canonical-202fd129c66843344116a3185eee80b2c3bba35ff002d889a6d2708861b3cfe7"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8aa8d860ed04 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0393c0eae97fbf94a3b29af2158d5e799d622f190f29cb838a867c0003c8a5ff): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5b6bf58e1b31fe8de14fdd34fd33e114d149327eb791d8e2340a3862c9766c): complete subsection reference.

<a id="canonical-39e4fbc97b41257e1d76fe3337a46fd2fc339c4a4a060fd6e36ffe1f67d34d05"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8aa8d860ed04 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0393c0eae97fbf94a3b29af2158d5e799d622f190f29cb838a867c0003c8a5ff)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5b6bf58e1b31fe8de14fdd34fd33e114d149327eb791d8e2340a3862c9766c)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0393c0eae97fbf94a3b29af2158d5e799d622f190f29cb838a867c0003c8a5ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6e457f1495dd8c4b0324279d609997186b2981405bf0254216bc3a78a58986b"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 9f1c7aca8678 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-81cf6c6e522ded5add4b51022f1150706fe7f24c435b9b4d229af1557b467f5e"></a>

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

<a id="canonical-4459a1d9361535f8336760933a61f10acc8b68267ece1f53bb42b384d0c50454"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 9f1c7aca8678 / 3

<a id="canonical-3df5b33fc0983c05db7ad1122b38863ce01c8377d4960c59ecaf4c57e488953b"></a>

<a id="canonical-cab5c864faf51f5585860e241e5ce1c58f16fe554fb3c534c3773486f1dfa265"></a>

## dns_list property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 9f1c7aca8678 / 4

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

<a id="canonical-cb4a4df7f3f0ec64fd8f8270abff990c4dbaf3ed9a90972473c367fefa1a55e8"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 9f1c7aca8678 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-cb5b6bf58e1b31fe8de14fdd34fd33e114d149327eb791d8e2340a3862c9766c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b908257c57a826edac5e3204bc5e645659ee98887b88ba56045b686a8426395"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 3e7a45dcd2ad / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-8b9c03740560506e93d429a43c8e639c310d874dc87f8a4c4a19d83e98acce0f"></a>

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

<a id="canonical-22aecf690cff3b767a9576cd868acd73c1fcee288775f7401c04c02facc8999b"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 3e7a45dcd2ad / 3

<a id="canonical-8a4b96282e2870784e0d60e5b3fa0c0735a37da04e66aed1a5d8b6772b9f0657"></a>

<a id="canonical-390f1dd182807d8b29b3fd32d0a2154917989714214e4625c98176d8af80d889"></a>

## configured_address property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 3e7a45dcd2ad / 4

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

- [first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-211598e41c4d1b5695e94257bcaac18d8f3b0021a808066b4ff543a8432d2097): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-ebb794fab24f977dad3a695e697b8f8c8014cd38975bfb0d58c2a10f976c9b8f): complete subsection reference.

<a id="canonical-25bbb62bf593833be65e4ed2f85980e1e968867b566eec01492a2c42c2f57b3b"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 3e7a45dcd2ad / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-211598e41c4d1b5695e94257bcaac18d8f3b0021a808066b4ff543a8432d2097)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-ebb794fab24f977dad3a695e697b8f8c8014cd38975bfb0d58c2a10f976c9b8f)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-211598e41c4d1b5695e94257bcaac18d8f3b0021a808066b4ff543a8432d2097"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e43f4ecc56d3acb0b0b3692f682d540146e24f53a5a5da8f7bcd54e75c46dc0f"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / e669bf4df870 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5b6bf58e1b31fe8de14fdd34fd33e114d149327eb791d8e2340a3862c9766c)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1d032b1d80ac68c5c301c5ed5fb9e42e63cda2a1298cda0a227da6798778a756"></a>

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

<a id="canonical-ea63db20f4982ce5f869dc90b2fef534ce07a83e73cdba3c66a5755c64162927"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / e669bf4df870 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d246e9c1969321d55eab76ae435fb5f756b3ee7893ac56eb842917af511428b8"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / e669bf4df870 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5b6bf58e1b31fe8de14fdd34fd33e114d149327eb791d8e2340a3862c9766c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ebb794fab24f977dad3a695e697b8f8c8014cd38975bfb0d58c2a10f976c9b8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4599e4cf43ece8d9a40391ee147a9c94010ba5126c66197387874ad9ed8c0cc"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 719084223bdd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-bebd9842466e927bdb6ed5f11500b98094b0c600a38a8fdd7c242ca0c7d85218)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5b6bf58e1b31fe8de14fdd34fd33e114d149327eb791d8e2340a3862c9766c)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-95792a6e75e7e32fe56b1d4cea977395469a92e6227b86486c9ad32f092bb00d"></a>

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

<a id="canonical-4db9692b06d53d5ffbc92bb19281f90412ecc5abfbf917aab2df7eb023eb4dde"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 719084223bdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-018fa9719e617c108602cfecba6c38e476d76c216f43fea01a67da25b5160c8a"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 719084223bdd / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5b6bf58e1b31fe8de14fdd34fd33e114d149327eb791d8e2340a3862c9766c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6e1f9258c6ae98b41304e15cacb1f41bb390f5e84524688215de972f67f473f"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 6422dd0e8563 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-630e3b338608cab076a7f83d9832dcdfd5a6a529a9704c56c05d8cea4f7681fb"></a>

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

<a id="canonical-e4d842c2898e9d2225db8f5f964a8158c303f4c15940be0c41409821f8fe0768"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 6422dd0e8563 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-032821caf19a7c47d47a9df095cc26118e5643aa8d256c7dc4040bfad89b1007): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-796e76bab0ef4f05b52bb1ee6580993d70b126a7eed4665ee0b0f4b21166216a): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5c2bed1ae8a6853a7c3b32306ab1780df79a21d1cdcac67e1c0e9e1e79a8f3): complete subsection reference.

<a id="canonical-5b3cdbb0c84f961a47a507b2107bf5f39902a780f01831a4ccd1b5beee95aa66"></a>

<a id="canonical-1dc8ae9cdb915032fc25068f1759af49e0dd29f720e28b82e5e6f1ff67ad9e17"></a>

## fixed_ip_map property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 6422dd0e8563 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-0ba1178676289ec91a9a22c34ca2f3a7a9bbbd146f6b05df530efe5ccf66216d): complete subsection reference.

<a id="canonical-b8382cc855d10b17836c3072e312ec0442431fc158ed42203897a4a8c04f8c28"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 6422dd0e8563 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-032821caf19a7c47d47a9df095cc26118e5643aa8d256c7dc4040bfad89b1007)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-796e76bab0ef4f05b52bb1ee6580993d70b126a7eed4665ee0b0f4b21166216a)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5c2bed1ae8a6853a7c3b32306ab1780df79a21d1cdcac67e1c0e9e1e79a8f3)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-0ba1178676289ec91a9a22c34ca2f3a7a9bbbd146f6b05df530efe5ccf66216d)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-032821caf19a7c47d47a9df095cc26118e5643aa8d256c7dc4040bfad89b1007"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-983470b4440fe31a88a31d1b6ae41e4d63f58dfb1896374491ee71b5571df2d6"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 221a49e6b06b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-11f30f95e0d7422e25e23a06a0d35922331f1061247f2505e4289413eac5d491"></a>

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

<a id="canonical-8baf4e756ae52d20370c6c895aa0b5efb2f31fe9acd46c4e871e5c29bb9b7b29"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 221a49e6b06b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-225ed50b6e842f746a2a95925f6895fbf42649587f9f608d7b861d8062f39463"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 221a49e6b06b / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-796e76bab0ef4f05b52bb1ee6580993d70b126a7eed4665ee0b0f4b21166216a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81c481e2ec8164bb13c831d950421417c6c2ea70010ff22e048dd9f4d9434e2f"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 7fbcdd4bd34f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-fd714839da344e4b5c104656c2a4a0b2b599868ae06c3affb0ce1fa24816039e"></a>

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

<a id="canonical-6ffe899c904a62004b06ef0eb1258ab041baee1d7acb216d9a8d02ce186eb5bd"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 7fbcdd4bd34f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59b61e4c192627c60cc2207af443f877e49e7ee210a290d885f105f3c6d14161"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 7fbcdd4bd34f / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-cb5c2bed1ae8a6853a7c3b32306ab1780df79a21d1cdcac67e1c0e9e1e79a8f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6086c9ebb15d03f93b46edb2ac947a90a3912161e0000e8b7d836985472693fb"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / b7ba34b990d2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-16ffa9b72ff93806cdbb9e4351bea8623027f9de94a45897932bcd96295b84af"></a>

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

<a id="canonical-2cb82450bd709196744ae499c6bfdb360fa948b29b2b765fb402ed8df2c1ba5c"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / b7ba34b990d2 / 3

<a id="canonical-122dfd925a313f5d25c1d1d7ca25fcd6a22fee2bd23b9c6fa996944e94c94b27"></a>

<a id="canonical-09e9c55f48e99a0f5c061337e8cee1d1743cc54ea5c2da89ac2c63ce4a2c1239"></a>

## network_prefix property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / b7ba34b990d2 / 4

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

<a id="canonical-8489f5fc58a76a3b4d3cde08f9d689915c45d8ff1a749a7276eb291bf03d99a2"></a>

<a id="canonical-1c11e8920e17b8533312b837e8655984a6b035cfbe5baac08d3b76742e8f9ba6"></a>

## pool_settings property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / b7ba34b990d2 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-015.md#canonical-b7f5ee755e069a1e3daa14e3c724c0b6ac3f97e9934e152bea75353cdca17a99): complete subsection reference.

<a id="canonical-52dd8a960797c3b2ef038c84cc3f724153a5be6f56be0c2d43b7c2734d93ecdc"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / b7ba34b990d2 / 6

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-015.md#canonical-b7f5ee755e069a1e3daa14e3c724c0b6ac3f97e9934e152bea75353cdca17a99)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b7f5ee755e069a1e3daa14e3c724c0b6ac3f97e9934e152bea75353cdca17a99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42ffeb58f7c6acd0860388215bcdfc60419862637d84af1719d32b475300bb54"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8ae3cfc33d8f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5c2bed1ae8a6853a7c3b32306ab1780df79a21d1cdcac67e1c0e9e1e79a8f3)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-2a311358a3b9713f0ded987b4f96c8aecd0fb9ab75a135815bd50e4db4316650"></a>

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

<a id="canonical-8dad6cdeb2cc7735223f094668fa33fc957c9c77a0f6e4ed8ec1e649bd06658e"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8ae3cfc33d8f / 3

<a id="canonical-26c2f095c273fdef91b7f9058e29933e475c89863c9fb9677696ee402fe94b80"></a>

<a id="canonical-22937ec2cfe2b5c90b677fa02e0b88684f0549cca195504d6e6c7a2f8193dd10"></a>

## end_ip property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8ae3cfc33d8f / 4

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

<a id="canonical-63e7ad3fd2ba3955fb4db65ebc8823226224abb142e53f46d07d05059c9ea52e"></a>

<a id="canonical-2fc0e24a15d1eb3424538d831836b5212a2d55224c9b6140ab3bc69fd32c8dc0"></a>

## start_ip property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8ae3cfc33d8f / 5

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

<a id="canonical-1315cf4db7fec0a8ec3bbe0500ab2895cf16cce06fe118eb1916c471794d7ab8"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 8ae3cfc33d8f / 6

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-cb5c2bed1ae8a6853a7c3b32306ab1780df79a21d1cdcac67e1c0e9e1e79a8f3)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0ba1178676289ec91a9a22c34ca2f3a7a9bbbd146f6b05df530efe5ccf66216d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f46421e2ad36e6ac38fb1ca02e8a8f0d1921a1664ffdb1b2dd0cf96075bcf47d"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 210e5b60e50a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-014.md#canonical-49cc9f2f74aaf8863321d26f1747f495bc83f5ecfa84aa30d4ebc136bbb822e0)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-014.md#canonical-8894804611653daf23fd17c901b0f78ace9f87f9d6848a18d13c5565e21184c4)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-7840876dcd31a0dbda5b8d805a200bc8c9ecb03cbc9f44bc3028bb235a41e13f"></a>

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

<a id="canonical-13feae964572cccd80a271c76f75b7c024f9c2695ee4f76a4013656eafb9da7f"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 210e5b60e50a / 3

<a id="canonical-8a2572d66848350303a1f8bcdfd9140df4a0b272e16039a632626f0472e3f8a8"></a>

<a id="canonical-1c3b92f2edd31ac9b11dc7ca9490024416853185d8c456f50038efce90de3955"></a>

## interface_ip_map property — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 210e5b60e50a / 4

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

<a id="canonical-8984818889ba4215f4d01808ba44c4bc3a9308cddf6d6bb4f263131a2a03a2d4"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.r / 210e5b60e50a / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0404780948111542039b6ad6b210a02c1a2adc751f12994945d33637c43d55e0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-672a7e92f37c16a66cd340f4d8dd81ba4f57d428c99521d62ff5d32f7f8f674c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2f6069a271480b48311fcdf47282c74ddd1d9c1f3f0eb1d2a843a0f866e87e7"></a>

## openshift_virtualization.not_managed.node_list.interface_list.monitor — openshift_virtualization.not_managed.node_list.interface_list.monitor / 9314751f5dd3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.monitor

<a id="canonical-cd55ad25f2d0b224b5dc92fb2257aa9143fd34ea6b031eda20837d4ea8859a9c"></a>

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

<a id="canonical-215d15aa3363491b7f61ffb69fe919cacc48307c87ff916154bf979499daf222"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.monitor / 9314751f5dd3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e143aa405ad0872dbf1461847b4a6bf2a2d05c67798408f81eb494edb759ee1"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.monitor / 9314751f5dd3 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fb972d80ee2d6e5acbe4614554c6cea0bd442d373014163ea99743e41de72bc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb4be7f322fd5db0a620c432d6fd9168c460567e5971b9534c10e10038fa6d69"></a>

## openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled — openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled / 289d806f262f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-ff5c49380ff301d64e550aace0e5435203ffb8dd1941c8b5ff5ff664c0cf0ba4"></a>

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

<a id="canonical-b842cf53809d5625afe20ab678721f2789caae6204a937875daf3b705c77dc08"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled / 289d806f262f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65f8b5ebaee5c126de30029b254d8ba64139f1112692c8a7c74ebdafb33165d4"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled / 289d806f262f / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-01b51a492fb5a7069a34e2fa64bcb6b54185330e3b6c62e3cee7e513c86f1e65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54d113fa4dea9c961214ef3bf00c3fb10adc9927c0bb0c6a313814fd2c4ee0cd"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option — openshift_virtualization.not_managed.node_list.interface_list.network_option / 258a994544b4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.network_option

<a id="canonical-7138995267f6568373e27df7c8331a20e1c3bd3ebbbbc01a1c9b9aa0f39ce236"></a>

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

<a id="canonical-f0bea0f9cc07aa9d65bc03870150530881e1cedabd3ecbcecdd95da1485b4d0b"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.network_option / 258a994544b4 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-015.md#canonical-63d70f87a0cfe3246bd79ec99cc95968271c6f26985a64296cb5d86838141653): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-015.md#canonical-236e34f75a53f6db4bb6a283843d61ec3120a2deb4867070f12c0237e17d3d75): complete subsection reference.

<a id="canonical-ce3a358b3d722055aa3def96500d89428427c5fc17eca4b7344444a47615922a"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.network_option / 258a994544b4 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-015.md#canonical-63d70f87a0cfe3246bd79ec99cc95968271c6f26985a64296cb5d86838141653)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-015.md#canonical-236e34f75a53f6db4bb6a283843d61ec3120a2deb4867070f12c0237e17d3d75)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-63d70f87a0cfe3246bd79ec99cc95968271c6f26985a64296cb5d86838141653"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301376e6e564cba31a7b5979485e508a447e864ab4e982e72e4231d49cf4d89"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 3680eede8afc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-01b51a492fb5a7069a34e2fa64bcb6b54185330e3b6c62e3cee7e513c86f1e65)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-2d4994dc9afa82d025910c0f4868cda325ad34072bd1a21c5632ff4f8feaf2a8"></a>

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

<a id="canonical-818c3a5b03ad66eb0b82930c301f1899b14c3484575034333dc9431e5e7ada4f"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 3680eede8afc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1af9e1196d4899fff44b3d60d7fe675550171a036e58dbf2bdf74fb4a22df2ab"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / 3680eede8afc / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-01b51a492fb5a7069a34e2fa64bcb6b54185330e3b6c62e3cee7e513c86f1e65)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-236e34f75a53f6db4bb6a283843d61ec3120a2deb4867070f12c0237e17d3d75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da216907b94b34e221636fdf6d30022c2d41d99ea56461d1d36484400564a14b"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / a4372ead5179 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-01b51a492fb5a7069a34e2fa64bcb6b54185330e3b6c62e3cee7e513c86f1e65)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-20783e2d173fc99165a5f803dc3f6003d212240de13960ac043cbba7981bf2c5"></a>

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

<a id="canonical-2be4452fc93bd0329741d8b2bb3fcd5ac520c3b3b9e365c321aeecd491e2b4bd"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / a4372ead5179 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3a4dd705b13b40b6ead0e7ad98f1a579e9f70b3521f5ca46db588c1409c6450"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.network_option.sit / a4372ead5179 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-01b51a492fb5a7069a34e2fa64bcb6b54185330e3b6c62e3cee7e513c86f1e65)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-45167ee4da1034b03cf7148effc6bcd561996424fe56def4e221edb9ca4826e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a31878e4ee552c8119c6ce4235d28ffd5da2b1d01a1f7032f0bbec3ceafaccfd"></a>

## openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address — openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address / 9f0203187851 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-0df7dafcf48ccaffdf2a096d449116815e5471ba3b20206de01acba2bf4a4cc6"></a>

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

<a id="canonical-2caf5506d120d805923c9a47260bba776e1feb0b0e2f12be73188e5a25f27d96"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address / 9f0203187851 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff8436bb47b06c39385ebad4e1b08b6b9c4f432d53ad05ba39135080589ea00f"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address / 9f0203187851 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c0abef0a9f9bd5fef6e51a70e6053165fdde3aa3c706e0bdf1a33f98cb9016d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5e4fe80d2a568a0ae538ce50abb35442a4cdd261d253398da304d98742cbf1e"></a>

## openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address — openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address / 697f7d0eae3f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-7a9b0323e177372ffa4f93ebd2950ee24fde43c271fb9e23fbf904d390050338"></a>

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

<a id="canonical-ba0f466593ac61853aa1f333a6288829178c2859838dadc113471073b4e54d9c"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address / 697f7d0eae3f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c2b4ba7acf69c2dfd02ed7445c5f034ccfa2a094a019f69b3dfc2b9126e934ac"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address / 697f7d0eae3f / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-498d2b612dfc75447733e80ccaf6302ce8a04dd4adaec8e85b541b7b4f6f4b2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131fef8c1ab5f5123fea11c2ff72f4c58604b54936ecd5854fef61c0cd96f4f"></a>

## openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 9bea3a9dcb10 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-bf6a3809ea5768f6e4a82d9881dd25ab29756d54ed72816893f4a7508245584d"></a>

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

<a id="canonical-05abad0fc3596b2458f98e99e659e34e3068f6fb01b9c3ec3eb5a6d16359a63f"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 9bea3a9dcb10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-402d7341426dfc594c39ffe47ed6d6308362448b9854d1f5066aaa6734a19647"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / 9bea3a9dcb10 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3c8db1c9a8f94d775e79b970f2c9c53d980949bf11a225913f4c918ccaa49ad7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43fa2a2fc8cc83796c20acdcbc7739a0f437745fa1a0c077bd419c95f485640b"></a>

## openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / b7564935aa04 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-c0825558b4b4d5b3c19cf3012967966f6b270063c1ab28a61dbb09ab08289034"></a>

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

<a id="canonical-825b3184e5a13d4fb7e4badb3292cfa2abcc235f40124077dc471a4fbc94b6a9"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / b7564935aa04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d60addbffd3330b2c9a35f78b835144707b8a8b6680b728827a222a15c4c1df6"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.site_to_site_conne / b7564935aa04 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-85a82372e0c9afcedbfa07fe678056e494085eadf9926ab8ed1d4f27dd597256"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b93b5affbd232df77f8aad01b13d01afd3e4f15eac46cef609045c900fa79549"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ip — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 103eb5a6c1de / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.static_ip

<a id="canonical-eb1224795b5429eb724a372cedf9af8906e13c9a6587609fc0a9fb849253c0e2"></a>

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

<a id="canonical-e6ae08e9a6db5d639ee55750cc5bd57f6686ac9fcff9e2905b11ab9d90ff7f5c"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 103eb5a6c1de / 3

<a id="canonical-775a9cf2aaec6611c492844f49729f4c3379f5f451b06b1bbfc89e9a70d1350a"></a>

<a id="canonical-4eee350c41dbe654910beed599814a9e29b8444c472fe27336ba5a039a74c495"></a>

## default_gw property — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 103eb5a6c1de / 4

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

<a id="canonical-4c74beb5914b9006cbe110356e19db479273f500384b43c5eb1682b5188a8af3"></a>

<a id="canonical-01a2713bebd30d7a5aef3e024a87dd591181dead9e76c4e6bcd1ac123772d2a7"></a>

## dns_server property — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 103eb5a6c1de / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0c360df6c6be90c9ea68c97619e6db0f4e8b9dc711833f7b437712c54605c551"></a>

<a id="canonical-11f5a276e27ae80a4dc8744e69ab82ad263c7ca5bbd68ffbee4c26607b6c70c1"></a>

## ip_address property — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 103eb5a6c1de / 6

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

<a id="canonical-7862037117fd479c28dc6babfe36a6f54043fad28b2e417b69ddcadce3eebbb2"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ip / 103eb5a6c1de / 7

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-19bbcacc4f9ea6747c8e777f3c64339d92e87b6ec2eebc7a461819e1c43c9fe3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-858fc2b5f7367e7a59797076558d7e0dd6aef85bdd14b282de5b7cf8ae9f3649"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 29ebfcef3536 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-c01681b91840adf5bbc34d8de50f1162bab8e4659c06fbb7fd5cd790cb4b5f27"></a>

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

<a id="canonical-2617f6b0d7daa02dacb3cc124b27bd8be22d78d23908214198c569d88231134d"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 29ebfcef3536 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-015.md#canonical-e8153d8f1e63bc3eab452c98073c02988595639016e36c01e468a5b6c269e6c7): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-015.md#canonical-a8dbfc152b24faf02a815b05ccf6b2a8138b76a363f8cebd5315a050efe1cf04): complete subsection reference.

<a id="canonical-0fd57c6de4eda1ea7ef2f85af9f4f7789a803ded80a542f28a1f10a6ead525b1"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 29ebfcef3536 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-015.md#canonical-e8153d8f1e63bc3eab452c98073c02988595639016e36c01e468a5b6c269e6c7)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-015.md#canonical-a8dbfc152b24faf02a815b05ccf6b2a8138b76a363f8cebd5315a050efe1cf04)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e8153d8f1e63bc3eab452c98073c02988595639016e36c01e468a5b6c269e6c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b16a4265a0d416fc39ee8cad3128fd6684a77c7237cc299a34a59eda5b230688"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 84f1610d7fe1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-19bbcacc4f9ea6747c8e777f3c64339d92e87b6ec2eebc7a461819e1c43c9fe3)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-91f6e9a7aed22a9f11158e64b61695fe4de06ee20535dbe5b067fc3850be5e67"></a>

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

<a id="canonical-2b9dd9dc4d0c04f370511c05572c06ab51dc70d7c24087c01062e6b90a2ce78b"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 84f1610d7fe1 / 3

<a id="canonical-0b765d909729b88fe3b7d3b1ac8639a523936c64539ea87e1646aed9bafaa90e"></a>

<a id="canonical-ff566e8a7aeed65a93c0fad8016e77b8e6bdfd04a1fcc0fc251339686c6d4551"></a>

## interface_ip_map property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 84f1610d7fe1 / 4

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

<a id="canonical-e761b7122caa8dc7c0bea2e6eec55bd32906c6d7ac9e63757db5dbcfab3339a0"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 84f1610d7fe1 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-19bbcacc4f9ea6747c8e777f3c64339d92e87b6ec2eebc7a461819e1c43c9fe3)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a8dbfc152b24faf02a815b05ccf6b2a8138b76a363f8cebd5315a050efe1cf04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c2736ccf4185f07fa346714c66f6d8e5f03460fe2e5c485eaee57e33b7937c9"></a>

## openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 4f406ac7ffba / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-19bbcacc4f9ea6747c8e777f3c64339d92e87b6ec2eebc7a461819e1c43c9fe3)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-861d97e48c882e2a8a71cd6bd5a459719fe9d6b2018db9e5df5198ccb0892703"></a>

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

<a id="canonical-9637825703f5251abf9686f37868ef35faece7501a8d219e25a106b054f5911c"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 4f406ac7ffba / 3

<a id="canonical-c719618d418ef4cdbcb4e86bbd9e4d8a889cdccfdc1a775b94958742b0c5a968"></a>

<a id="canonical-bf4243a17871f61960f43b1905a1ac4822d01f5c949425ec01da1b2833a255f3"></a>

## default_gw property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 4f406ac7ffba / 4

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

<a id="canonical-e245a5b83b0a03c514f03700107127442c1f431ff28f42806a2499291cbe42e0"></a>

<a id="canonical-77124fe3c6ae2a1ac72dad06e2f28edc34eb18f8e11a09e161d9b2b1555f7805"></a>

## dns_server property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 4f406ac7ffba / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-78787cd5b1c290c87cf0e15a7e45747dae0e09e68dbc9f9137316dc6f22d4d3f"></a>

<a id="canonical-dfb862cca6cbaf9254402f22460cb813ce28d311cefa35f9dc7c4949265c184f"></a>

## ip_address property — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 4f406ac7ffba / 6

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

<a id="canonical-267ca7d2e5c06faa10a855d2a6ac5c50c9aa2c62943ef1cfc23026389801a46e"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_addres / 4f406ac7ffba / 7

- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-015.md#canonical-19bbcacc4f9ea6747c8e777f3c64339d92e87b6ec2eebc7a461819e1c43c9fe3)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0c32f51ce747fafdd194a4634adfa694ab263f4c8c901b70bfb970d15359a4e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d364475e27f9b628ef84f1ed74fdc185bcc2870e447ef049c2832f5eb101a4a"></a>

## openshift_virtualization.not_managed.node_list.interface_list.vlan_interface — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4fe04fe1700c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-19cb9bfc065a0ea3715f0020d6a81383bd59b2f3e036c05ba47591547370d6a7)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-014.md#canonical-f66c486b3ab7db5741cfc3513f045502878c62e6f2da82a42f77e6b4b434fc47)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-014.md#canonical-1ec1d74e8971e5bd360c5be5919264cc65aac0a7ced9eafd314b9a0b3fa6b867)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- openshift_virtualization.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-7122326f085aaf7a6784bfc69bee817a49919e683926c1f16fcfc15f0229c0dd"></a>

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

<a id="canonical-5f64cc6e8ab17b32603be7e904f4391b5efbd29a8831d5dea42284156ea88e09"></a>

## Direct properties — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4fe04fe1700c / 3

<a id="canonical-0dfa96cb0b5a491d8228e6e4fe29b7fa87317fbd4856ae74e4ef5d1a50549258"></a>

<a id="canonical-9d5174beec287b49c93ce2e0a6791fab6a959974bfe31cf104487162ede7a93c"></a>

## device property — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4fe04fe1700c / 4

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

<a id="canonical-d06d913e1de1ccaf9e63c17b5437de0bd49c5701cf06e440acfef864edb72035"></a>

<a id="canonical-127109578e855b0b16d43fe6b9feeb7bacf3cab6ee3e9a7c313ef5a7c9527030"></a>

## vlan_id property — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4fe04fe1700c / 5

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

<a id="canonical-4090bf0790a726edd87e04c4f3ce8313c45b7f82841d3e991172741843bd7e51"></a>

## Next pages — openshift_virtualization.not_managed.node_list.interface_list.vlan_interface / 4fe04fe1700c / 6

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-014.md#canonical-b44906d1752aa0e1b026395f1b984df820aea3e299f2618e0734606791e22046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5cde5e9507720d6f32a98e1007e178183bde2671a78aa4873d8e46da67f6fa74"></a>

## openstack — openstack / 21fce2c1f27b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- openstack

<a id="canonical-a1e6380ca91feec0bad577049023940699d0974f5b1d226b74c62696f7d93401"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
openstack {
  # Configure direct properties listed below.
}
```

<a id="canonical-c762bf17bd0cc42a1558e3c0ee032274705b4155fa176f47e5387af8c4aa9b03"></a>

## Direct properties — openstack / 21fce2c1f27b / 3

- [not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5): complete subsection reference.

<a id="canonical-5fa378448e56dd1f172f06d4896e224a82052ecfce68260754357e995f34bf00"></a>

## Next pages — openstack / 21fce2c1f27b / 4

- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af7787472dd71d00f8775f23783b3a5d3b6df493c9cc1d913aa99d36211705ac"></a>

## openstack.not_managed — openstack.not_managed / f5b7ea2427f7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- openstack.not_managed

<a id="canonical-a3f6afb6a0f3567e9fb15ab4b5d917261079942c8583a3155671deaf902aef67"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f72ce8b70763df4baab0d4df3ce73f7ed67c45a442831e4adeacfa89dfa93d8"></a>

## Direct properties — openstack.not_managed / f5b7ea2427f7 / 3

- [node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e): complete subsection reference.

<a id="canonical-0a7b5d7418809fa92e3491a8b6b4fa4cc0e20ffd4527ef1d4d817999304e7b4d"></a>

## Next pages — openstack.not_managed / f5b7ea2427f7 / 4

- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d48061fc45bfd6143e797fdd9e9133028a0043f6f2b3bde350eafecf1a2124f5"></a>

## openstack.not_managed.node_list — openstack.not_managed.node_list / 31d36583e0bb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- openstack.not_managed.node_list

<a id="canonical-942a9a7611e374f06f5c7020c155b090c5428918e30faabb2c777ef666dd8134"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-988b5bdfdde22d3765a5b3cd5f9adfd27803029ede0225bfc133bac9212ec61e"></a>

## Direct properties — openstack.not_managed.node_list / 31d36583e0bb / 3

<a id="canonical-7af55c932958af04302ffd4ceed498ff24728caf786e8953f327e977df33dd98"></a>

<a id="canonical-1eeb1856c4197abf71f9325b46a74b036e1cb696ac5199a44a87ec8c2fd5d553"></a>

## hostname property — openstack.not_managed.node_list / 31d36583e0bb / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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

- [interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e): complete subsection reference.

<a id="canonical-54368faa466f6bb2b24d4e0fd2702a54854a7c8bbe7112a1db11b9d1fb8a7cca"></a>

<a id="canonical-de0dd52b07932af3ef1355c4228dccee2c1dc3101fabe9a16854f566daa19f37"></a>

## public_ip property — openstack.not_managed.node_list / 31d36583e0bb / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-9ba4763f09adca743d8f3b55b0487cebef390185d548ea305887806f202e99be"></a>

<a id="canonical-ee315d58582e6ac5012f0cb288cc2970d3a3b0a11c1f64d8e8ddaddc1575aa3d"></a>

## type property — openstack.not_managed.node_list / 31d36583e0bb / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

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

<a id="canonical-695262af785cce257205d1a13064c34f22a554a94cef7b10fb02e4ceafccab3c"></a>

## Next pages — openstack.not_managed.node_list / 31d36583e0bb / 7

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31d7d3bcaf0361c40ba82dab2d54838e61b3e8a1f9e2c3c9b2ae3b02214e6d91"></a>

## openstack.not_managed.node_list.interface_list — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- openstack.not_managed.node_list.interface_list

<a id="canonical-34ed996a93b4271aa8a2382f2504be153a019644f3893fd8cbbb496b86628a47"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-09155d3681d5ff9f9be53c649e81047314c1f40a3eca5570565857efea47dc83"></a>

## Direct properties — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-ea51ba9152e986321d0054d226472ecd0621926a32586c4da00cfb31f3a6745f): complete subsection reference.

<a id="canonical-975532a59f94ac5640ca2d8471a2225046131fdda64796af90af2c3b535b8e5b"></a>

<a id="canonical-62991409796ab7a1ba82fff8cd4566e526fe57d7856b6c6bf97ce4632c013136"></a>

## description_spec property — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-015.md#canonical-ec5f69b77cbe8990e3d6e21cc61ad821ecf5ef2778ef28da34716059c2911bfa): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-7c8ec9b093d5cbfd537c76ece2c04b0ffe953bf69b5fc0512b76e9e94a59c5fb): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797): complete subsection reference.

<a id="canonical-e42f0c49331ef9eaf908fbad78f95a052cd73319090aeab3a5d0fda2b0c7edc3"></a>

<a id="canonical-ba1fbd8b4fe9507b2a4b49daa67b85c60e222db411ae5a42311d586c4c963e2e"></a>

## is_management property — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-920c2599d92428a7fe07a3753b2b96162b9ac3b04bd45da11ab59148b90701e2"></a>

<a id="canonical-0c2cdb76686eed54660fb2e473cfa2b8fbf878997d48e18ecde90fede02ab5db"></a>

## is_primary property — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-56448c06132e3b24d6d56b9dabf66db77c78d272e59be44c01be741851e941ad"></a>

<a id="canonical-1154c752bf35758a91258115462eade47ae7aadfc38d39e1b8885b1b5e376c47"></a>

## labels property — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 7

Type: `["map", "string"]`. Optional.

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

- [monitor](resources--securemesh_site_v2--reference--group-016.md#canonical-d15ef01be5bcf8cdec599778c9387338a52d7c01027f83f0d0bcb23889c73822): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-ea6fac01147da19415ddac1904373e25dc5aa43c5c7ef087ac3b54a525b801a0): complete subsection reference.

<a id="canonical-84a3e59283fa75e2c27ed91414c9d4c20a1a6b39fb34af0f03d127cbdb40dde1"></a>

<a id="canonical-d9a9f72cd33db4148a75764d57c1889b1115878a824e068987434c9f4bef850d"></a>

## mtu property — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

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

<a id="canonical-3f4bd50bdd6d72b83853bd133f00e79220575d2b23eacc8d7b7dd8003a8df9a2"></a>

<a id="canonical-8f6ffdefa20ed0c2d9f081b8d45a81221e4475b374b3f5dbcefeb6d3e7043015"></a>

## name property — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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

- [network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-cb95685c6151aa5ab0e934beb4bd98961832c7964b6c1fe11e5f1ae15bf27284): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-016.md#canonical-bed5566f4b80eb81da7cc71c5bdd2eec20a5a942ec7f68451a7a038c079eeb6e): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-94de12dbb480802e75b81333aa41337b25ddd2ccb03d31b86f2b3cfd0582523c): complete subsection reference.

<a id="canonical-cf462ae054fe23d7401f98d983db6f060402f56c1469b1923ba3ad305447be47"></a>

<a id="canonical-f2f5f00a3cfef0839bc4d0ae3fc1a79da782e7da1c1795e05a7546d6040ca981"></a>

## priority property — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-241b19aa15996ff476d77b5d3b40e12c0c3521d8756ea8fc130e5432a2a2541f): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-8df2288afd00d05060dd889af762fd6ffaf46a730b62a5b95ce642efc5de3e45): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-96b72e06db4f8bdc4a8937b219e59de5d5162120aad6011fc939fafd3bf1c0b9): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-a49e266a35a0b9739890bfa8a72a4e857cc5e14830ea12771c42acf75ea18783): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-08f097c51aaeacfbbee9ac8d133b0e2d7a629f88028a6cd15205e2fa8536115c): complete subsection reference.

<a id="canonical-cc83c8791ca3b90b1569617fe48a186abad73c54ccecacb11638e151fec38f60"></a>

## Next pages — openstack.not_managed.node_list.interface_list / 24bb7221ab8e / 11

- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-ea51ba9152e986321d0054d226472ecd0621926a32586c4da00cfb31f3a6745f)
- [openstack.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-015.md#canonical-ec5f69b77cbe8990e3d6e21cc61ad821ecf5ef2778ef28da34716059c2911bfa)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [openstack.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-7c8ec9b093d5cbfd537c76ece2c04b0ffe953bf69b5fc0512b76e9e94a59c5fb)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2c96d3171863ef4091404aba5c55bf3a3f9b4dcf97be0be197de192ea9d6b797)
- [openstack.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-016.md#canonical-d15ef01be5bcf8cdec599778c9387338a52d7c01027f83f0d0bcb23889c73822)
- [openstack.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-ea6fac01147da19415ddac1904373e25dc5aa43c5c7ef087ac3b54a525b801a0)
- [openstack.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-016.md#canonical-cb95685c6151aa5ab0e934beb4bd98961832c7964b6c1fe11e5f1ae15bf27284)
- [openstack.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-016.md#canonical-bed5566f4b80eb81da7cc71c5bdd2eec20a5a942ec7f68451a7a038c079eeb6e)
- [openstack.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-94de12dbb480802e75b81333aa41337b25ddd2ccb03d31b86f2b3cfd0582523c)
- [openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-241b19aa15996ff476d77b5d3b40e12c0c3521d8756ea8fc130e5432a2a2541f)
- [openstack.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-8df2288afd00d05060dd889af762fd6ffaf46a730b62a5b95ce642efc5de3e45)
- [openstack.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-96b72e06db4f8bdc4a8937b219e59de5d5162120aad6011fc939fafd3bf1c0b9)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-a49e266a35a0b9739890bfa8a72a4e857cc5e14830ea12771c42acf75ea18783)
- [openstack.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-08f097c51aaeacfbbee9ac8d133b0e2d7a629f88028a6cd15205e2fa8536115c)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ea51ba9152e986321d0054d226472ecd0621926a32586c4da00cfb31f3a6745f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7ef785421f6cf6e2da7c5100227b3a5dab86967da7be2668148ce30227b22b6"></a>

## openstack.not_managed.node_list.interface_list.bond_interface — openstack.not_managed.node_list.interface_list.bond_interface / 49a8eed96ae5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.bond_interface

<a id="canonical-515b8fe567c41dc01088bdc5598977f8e3fc99836ed58389887eace55ec1f8a5"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c29c6a41981307cb2db74bb096d2982b14fef5a451a8058b8cc18e1b2f9e1d0"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.bond_interface / 49a8eed96ae5 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-015.md#canonical-37dd51367e0cc385fe7914fe0c6dfef41019e64519bcb41e52398d86490ac1d9): complete subsection reference.

<a id="canonical-6067ab820677e3ada378d2fb29efd13ad6a90cbd7d40d1f32a25a5a191061107"></a>

<a id="canonical-d6dc9a4d7546a877c6ab23d2f7743b8bc33e0039090db452d670a89302b4c7b3"></a>

## devices property — openstack.not_managed.node_list.interface_list.bond_interface / 49a8eed96ae5 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

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

- [lacp](resources--securemesh_site_v2--reference--group-015.md#canonical-d2ae82f915a3e3f466aa169832ccfdbe3226ed73e45020e61fe8d2830e6c6ac3): complete subsection reference.

<a id="canonical-9749463a052f4a34af59ac5c543c320fcaf7aea847e6233511f09b84b4a15e0b"></a>

<a id="canonical-10f6e8359e20f848899d394afa3c8ff419eeefdc6e3daed764e770558cdd88d4"></a>

## link_polling_interval property — openstack.not_managed.node_list.interface_list.bond_interface / 49a8eed96ae5 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

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

<a id="canonical-db335b36ab8f74f93467eb1997f741ea8caba03812405eea5e79c4b94ed7833c"></a>

<a id="canonical-5a60dffe35f70165dd1bb6f8caf6b916595274e5c9c1e49fe6b6d53d74c43882"></a>

## link_up_delay property — openstack.not_managed.node_list.interface_list.bond_interface / 49a8eed96ae5 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

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

<a id="canonical-4d3c099d034c51fae38bccd27a0c8eae1966a6545db44ed52fd1298fc4c0e9af"></a>

<a id="canonical-fb4cc0c998eb4ed8d24fa872b1a0360bf3e49cac8d7cd89a9eb5ef91dfc88c35"></a>

## name property — openstack.not_managed.node_list.interface_list.bond_interface / 49a8eed96ae5 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

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

<a id="canonical-270622fbf7eb54c5572949efe639783652514253bc8e81efaf85de4da55a9aea"></a>

## Next pages — openstack.not_managed.node_list.interface_list.bond_interface / 49a8eed96ae5 / 8

- [openstack.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-015.md#canonical-37dd51367e0cc385fe7914fe0c6dfef41019e64519bcb41e52398d86490ac1d9)
- [openstack.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-015.md#canonical-d2ae82f915a3e3f466aa169832ccfdbe3226ed73e45020e61fe8d2830e6c6ac3)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-37dd51367e0cc385fe7914fe0c6dfef41019e64519bcb41e52398d86490ac1d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6a036d45502bb8695a38e1fdb8b82791dc516d19064d9acba3a45f32cc5b460"></a>

## openstack.not_managed.node_list.interface_list.bond_interface.active_backup — openstack.not_managed.node_list.interface_list.bond_interface.active_backup / cb658ced81f1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-ea51ba9152e986321d0054d226472ecd0621926a32586c4da00cfb31f3a6745f)
- openstack.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-54629afff146cf12208a9136243658117a07149a559d31292514eaa4ad3feebe"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
active_backup = {}
```

<a id="canonical-2b271c854e17b379544c2e14a43eb0d0737ef217ab1d11ff9ec111782a242af2"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.bond_interface.active_backup / cb658ced81f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc6d6ecbfd23e5ac974176560b9b0c533d83f3a150a028c8f15d300362168c30"></a>

## Next pages — openstack.not_managed.node_list.interface_list.bond_interface.active_backup / cb658ced81f1 / 4

- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-ea51ba9152e986321d0054d226472ecd0621926a32586c4da00cfb31f3a6745f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d2ae82f915a3e3f466aa169832ccfdbe3226ed73e45020e61fe8d2830e6c6ac3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c81b62182f5f07d66754490858e17cd3a1ce75dca304fd087646386311879db2"></a>

## openstack.not_managed.node_list.interface_list.bond_interface.lacp — openstack.not_managed.node_list.interface_list.bond_interface.lacp / db7b4d4492c8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-ea51ba9152e986321d0054d226472ecd0621926a32586c4da00cfb31f3a6745f)
- openstack.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-475541a81eb7f8e5a7168433bd6b4aea0416923fbf559f0af59fac738f56b7d4"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2cdf9930e4f7b1d0c1473e617847b0258c5ac9e42c4c611c8ee1279ae13c75fe"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.bond_interface.lacp / db7b4d4492c8 / 3

<a id="canonical-51c0a1184ce0b1d81da89f12facc6a1d2fb537dc59dbe48ca2012811e44f014f"></a>

<a id="canonical-6a5b440b8f2a29d41c331b01e03d5f6c15ce7ce3f698e2875d5892b58b14ea89"></a>

## rate property — openstack.not_managed.node_list.interface_list.bond_interface.lacp / db7b4d4492c8 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

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

<a id="canonical-b4d99ff1e885fe36e43d032c64f4df02d4eac7faeb263ea1712d04bef2f8aa6e"></a>

## Next pages — openstack.not_managed.node_list.interface_list.bond_interface.lacp / db7b4d4492c8 / 5

- [openstack.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-ea51ba9152e986321d0054d226472ecd0621926a32586c4da00cfb31f3a6745f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ec5f69b77cbe8990e3d6e21cc61ad821ecf5ef2778ef28da34716059c2911bfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9361fab30717bc60fefbdb0beeb6923e8c7b851deb207b60ced19a9da613e5d0"></a>

## openstack.not_managed.node_list.interface_list.dhcp_client — openstack.not_managed.node_list.interface_list.dhcp_client / 856539ba8670 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-db22c7b23695f40cd2cc532ffc3262865ac9229d838c6b48a3c036105ec5affd"></a>

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
dhcp_client = {}
```

<a id="canonical-9f7bf3ce4dfd6321b7323b74db0c707b4192d94aac0a93b33ac057b352eda5f7"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_client / 856539ba8670 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e3d6eed3a49d11c17c724d1c5eb0acb35234aefa064cc0415e790477c95789e"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_client / 856539ba8670 / 4

- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a71d66a734c6232c3f2bcf28e4a933c88ff5303cf3c38a3ce3aaac800445b75"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server — openstack.not_managed.node_list.interface_list.dhcp_server / b4fca952b209 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- openstack.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-9c5cd6a896357097b6f9f35a665a03208df06deea585ffc67c17bd8d96f1e00c"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-e1b0530a9346f8b4d733e9059f75d8d16c9f64a43c5cc2d412ed086d6a311a32"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server / b4fca952b209 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-b2d54b3abdc1648b74fd7acbd15d09349923c9f98b0c55870afb75909238e579): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-f9be14feae0e3d45fb4af6c0df4aa89977fc665aff00c370c8147fb03ecd784f): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64): complete subsection reference.

<a id="canonical-6846988ff3198facaa8a3614d6071b75490e3aff6206ccbd4d44bc63857ce577"></a>

<a id="canonical-ff7c2f856d0f627ff93e74bbcd8879beb8093162e472f23aab368cf45c919485"></a>

## dhcp_option82_tag property — openstack.not_managed.node_list.interface_list.dhcp_server / b4fca952b209 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-3b3576cd563c3fff20c5f6e9656440aa96ec4b56ab236076c780008a3860fad5"></a>

<a id="canonical-597e0bb0d9aae742bc0895736dce5182215bd18b841323ba27096ea45cfdca2f"></a>

## fixed_ip_map property — openstack.not_managed.node_list.interface_list.dhcp_server / b4fca952b209 / 5

Type: `["map", "string"]`. Optional.

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-8d89c61cf75f2d73f73bca80ea6e66695a770ed715def7cbbbb862bb1fc9bc79): complete subsection reference.

<a id="canonical-0c9fa26f857034e54859f2aea056aba6f899ddb0fa30bbe244b1801be504847f"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server / b4fca952b209 / 6

- [openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-b2d54b3abdc1648b74fd7acbd15d09349923c9f98b0c55870afb75909238e579)
- [openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-f9be14feae0e3d45fb4af6c0df4aa89977fc665aff00c370c8147fb03ecd784f)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- [openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-8d89c61cf75f2d73f73bca80ea6e66695a770ed715def7cbbbb862bb1fc9bc79)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b2d54b3abdc1648b74fd7acbd15d09349923c9f98b0c55870afb75909238e579"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49c04d26d45183450b934321054b4feff75eb4a3b06fee3246cc0aef35858b76"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / ccdf43bcfbc9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-350f2c4d3808e1540a7c7855748d03fe0fdb19566b7bb6a0916ce5ca24474b34"></a>

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

<a id="canonical-08eaf6e7bcea28f61151e043256a8fe7770af606e1fface881c1ae1d4e20413c"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / ccdf43bcfbc9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58e04978c65085a5ab3d4a144f261513022a4939d52a8ec585df2791fe89bfb7"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / ccdf43bcfbc9 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f9be14feae0e3d45fb4af6c0df4aa89977fc665aff00c370c8147fb03ecd784f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01391bc9fe92ba01e115755f3eb2ee50e60bfb020f4ad3b0976afca9ccae37c8"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / e8cb3d575b10 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-eda5504a9b300a858f235cfdf00aaceca4ddf7861e9930f1441a3056070b910d"></a>

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

<a id="canonical-ce2f5efdddb946f544dba64e2559de56ca07e7c60686e639c246e253f9aeb818"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / e8cb3d575b10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-541c2a6fdf4b6f13997a2ed6eff59a06b4f28bc4de7b1abeef144da5de6efed2"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / e8cb3d575b10 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf2d7b7b353391672d3b28798d34171d60345feb40eda8a1a00df79256003a44"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9ba6fbc88f36 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-63afb381455ff0cd2a362d3b9effecece9aa6ac5ef3d14249381fbfa90c02872"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

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

<a id="canonical-be308dd2fbe36aca389e9eeabd01fb3245e2d67ec5aa008f5252b506b10f44df"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9ba6fbc88f36 / 3

<a id="canonical-c270681e832d25aef2fea7e189f897f5eccd51940dee7c56b0db27d5a2dbcf9a"></a>

<a id="canonical-1749a14ea5fb8a76bcdaedd4be373b6c8b3e5990da7857abe51c8e75e1744fdf"></a>

## dgw_address property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9ba6fbc88f36 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-354e83f279cf7e40ff5789b97f5521da9c336c5e47f56409b9d2f700603900b0"></a>

<a id="canonical-329303b0e1a3414f0b0c4cfdf148358f61932ebcadd6b34af5e928287c3008e4"></a>

## dns_address property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9ba6fbc88f36 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-8ce59ca45ed5e4527e84a9ec402eb0b16fe1c0631168199a860db5e06970f37f): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-586add1f8929a32a9e8cf1beff26f14dc6a546a68a5c61c8d9c9026e6532627c): complete subsection reference.

<a id="canonical-986142580d0c516938c9c1bf8482b106a7a20719cada67ad459a8add0eec8d35"></a>

<a id="canonical-9acc0459be49eb650d44f03b219aeb62a0c43e78b50384e218ae8de98f54863b"></a>

## network_prefix property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9ba6fbc88f36 / 6

Type: `"string"`. Optional.

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

<a id="canonical-53c7335f7bbbd2e58113ec9fa9ad6f54db54cd88dd28fcd9672494f45d5bbfb4"></a>

<a id="canonical-52314b0be72f98c4b634718a43d9e7fa77849b5ca4746d68a5946f461f3230a2"></a>

## pool_settings property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9ba6fbc88f36 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-015.md#canonical-bc26c8dae666ebb043debd4e847a79224eb5598b559c2a8ca02ee85ed62372f4): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-015.md#canonical-af3aa550884516fa6ff77eeb91f25be7d240b723a9c765b582a13b40d1232f8b): complete subsection reference.

<a id="canonical-c13d3fad09f9f5f94e316ed7d9c25734cb7227b97af1baa6e293afc780c3d8e7"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 9ba6fbc88f36 / 8

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-8ce59ca45ed5e4527e84a9ec402eb0b16fe1c0631168199a860db5e06970f37f)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-586add1f8929a32a9e8cf1beff26f14dc6a546a68a5c61c8d9c9026e6532627c)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-015.md#canonical-bc26c8dae666ebb043debd4e847a79224eb5598b559c2a8ca02ee85ed62372f4)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-015.md#canonical-af3aa550884516fa6ff77eeb91f25be7d240b723a9c765b582a13b40d1232f8b)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8ce59ca45ed5e4527e84a9ec402eb0b16fe1c0631168199a860db5e06970f37f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a213c6b061fbdc72a4198125d2f4006a58645dcce9d7dbc5dd4b9c937bb53a74"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / 36e4eb355c07 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-21a1751884217d7ccfebf8e40269a223173824ec2776790c2bd8f9ee9a48e8bd"></a>

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

<a id="canonical-f00dc9c01563ce9540ca7bf094dda8628a80a061f8f656dd09e760cbaf3a33ec"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / 36e4eb355c07 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe4e952ec32a30c2ca6c6b4bf84b9cab19eea7e8038c20e6e7b7725122d6ef28"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_a / 36e4eb355c07 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-586add1f8929a32a9e8cf1beff26f14dc6a546a68a5c61c8d9c9026e6532627c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26dd8dbee9867a2b6fc18d4aed7f15a9b1f451e276333339cc21dbe955b51a54"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_ad / f75d50057948 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-b3219f1b415398611d83aa42c9b2a3c04270dc37f3432217b7ae73ae1ad1dc4f"></a>

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

<a id="canonical-a0c8e0cbdd0db51e0dcac30ab5977da022402581418936f51db08fcf3154d822"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_ad / f75d50057948 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4b78a7df7da6fedb4e491893600154824cc283df8dbe81e9777dfee7390497c"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_ad / f75d50057948 / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bc26c8dae666ebb043debd4e847a79224eb5598b559c2a8ca02ee85ed62372f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc1957904088ecb9e93e35838eb228ad0bc85c5524f049c92570ade8fd4c1f16"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 956808bb12f1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-b6031a6ab6a6cf5b4af6f52974abb4b39f76171a39697de8bbf260d83c6c6684"></a>

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

<a id="canonical-506026bda547000105389c4adb807abcbcba3554d9d82adeaed3568c8b93ef1b"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 956808bb12f1 / 3

<a id="canonical-4b259a396e427060ecba2f1c7d83cfc031a3ad6d27d86976d72e6a39d83f776c"></a>

<a id="canonical-34e3b2b06584de54378310210641365c8d10682aeee35ef57fbc3a4ba7acaeca"></a>

## end_ip property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 956808bb12f1 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-e07b8074cacebcfa9e1828438be3c2750b41cf07feda39c11aaa38b2a8869995"></a>

<a id="canonical-0500c533b7a37fd44daa3278b3e3a2dc030abba05113cfa6029c61f6109f78ec"></a>

## exclude property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 956808bb12f1 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-62057b86d6c8c96adab0e8c748309a1dbe13c7824df3814f20c9f934bec5c49b"></a>

<a id="canonical-878bfd524216f9a030ea9249680b4c4cf42876fef9e9f0b4a94467dca9e76e64"></a>

## start_ip property — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 956808bb12f1 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-e8769adb45ac81b58480df5165bcb4f4e5f2db3693685896f88759820c2b34bc"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 956808bb12f1 / 7

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-af3aa550884516fa6ff77eeb91f25be7d240b723a9c765b582a13b40d1232f8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cfa6a40ada28404fe628bb14bd62d0c84be9ddab441bd8ace1eab6aa66df6d1"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as / 96a6305a32fb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-bd21ff093db10f41c3f5c7dfe7a8715e580f5d85c7fa1377dc2bff7a4ad82ea2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
same_as_dgw = {}
```

<a id="canonical-c5cc0ce4a05457003c22afad1a82993bda9d32e6ca4d4f8b0a477aca278605b9"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as / 96a6305a32fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-599fd69379778954582e807aa0e77ebfe2ae60fd2ddaa29997d801685034c756"></a>

## Next pages — openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as / 96a6305a32fb / 4

- [openstack.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-df7fa893633723c93c159ae7003c11297914d757eafd10b6f90abbf2d931de64)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8d89c61cf75f2d73f73bca80ea6e66695a770ed715def7cbbbb862bb1fc9bc79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47cd0c23be7f06163c52809f92587dac89ecdb9daeb1d72e56b17b1479de63c6"></a>

## openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 9e81a5cc7935 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-d3f295e7510b96417be62c8d0fe80f49efa6df17a64dace2f2a7c5081597b77a)
- [openstack.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-363e948e5067c61e9323e184d592949c384bb960f5d0007a3171c041a75c9fe5)
- [openstack.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5413a778a90b5c050a7b177a3b67d585269fe3828d02c6c1e1240c1f239c9b7e)
- [openstack.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-5d49e2133d9f358bf0ba7988228aad82a62822454c92a774f69488e6703ccd1e)
- [openstack.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-6d7b4951fec5db6a3779915c8806b3aabe001f7124a360fd98072b9a2df7022c)
- openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-99cb0eeecfd9cb4632c5c6cfc2de10f810d43e7528dd75dfc4b4a3631b0345d2"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-a76381371eb46b358008fdecb1d11be2ea092d21eb3b47df69d306c819d071a9"></a>

## Direct properties — openstack.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 9e81a5cc7935 / 3

<a id="canonical-3333fc73af0c217e596f56a63cc863e46d931a5b5a6ceea419c51e1a1bcd7d46"></a>
