---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-099fb802627c66b17753d176576fcefe6bb60f647a0cc6be4187a544d5079a7d"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-00316da637695aa2487a72b19f06c9e4c092332f66b454c3cde4792a7dfc1fb7"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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

<a id="canonical-5812aabbee42d56c63cbeb7e97ed5b04310b6374ff694bb6aa15e692abe04b23"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 3

<a id="canonical-06f021f94077a54fe4c739077997d4d8b5c137bd113bc4512f4a3b6ebb534021"></a>

<a id="canonical-9dfeb4f64ecd78fc25f822cf67c4610aa71ea7e2cbf63318e4fefc5685285bb2"></a>

## kind property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 4

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

<a id="canonical-f43b35d0f27088d7c85d92c3a28ba1db57e4031f28a30ffa4ccb70739b61ef0c"></a>

<a id="canonical-0ed92f3b965633f9c50ff9904119afdfcfe79d741e363570a559fb87b7809327"></a>

## name property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 5

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

<a id="canonical-c5c7bd02676e151947a8af7bd0c91cbbc43a6eba61a937a8aba84bec057b0930"></a>

<a id="canonical-ebd0f67bf7b1adb1bcddc67f2bc55aecf5576f429013558eb486489dafff9f93"></a>

## namespace property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 6

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

<a id="canonical-300c4a8ba89310499c47b207107d8222e96a0ff8e79d4b0eeef88170439acf65"></a>

<a id="canonical-1e77be875284cc7cfaf2ca7b06acaf177aa7f7ee9f1b5851832128cb438a4982"></a>

## tenant property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 7

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

<a id="canonical-da7cf44da6d66276584e54478e20ba9685df2b6c22036d222a2e70eb25163502"></a>

<a id="canonical-ec0cfbd88b63a3b22f73ace29aee0f3c556a675a700debb18ad4f88223e6c23f"></a>

## uid property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 8

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

<a id="canonical-59d041eb4335540efbc6ffea24ac794d2a2bbad27c9af8cd3a79af4b4d3904dc"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / be19546f7722 / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a77f6b94b10cfe6c965cd837a831e4b27f067804f8b3b9a39a4d66f6d4a3e05e"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 7ba7e8b3b727 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-a8000feb975ccbbb05b72271ab5711179bcc0d24b6d04b2792c800832f9a8569"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-187f4cdb33bbe829d8348663f0b8090f993429edd34efb4f705f1f7cef91bc63"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 7ba7e8b3b727 / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-76e93b9fd13530b3d6eb419fab5ea48716b0e6aba92c28249c4c64cb2858b6dc): complete subsection reference.

- [ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-9f428dbe32b8bd6e33a6f56588d0ec98cf6097410f8c9b72835772e842e8832c): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-37d564b1f1ab7a43092e8a9a82a98aa24ce9a3451484b783c5f4ef12e1856e06): complete subsection reference.

<a id="canonical-3c581011219f3de82d0f38c4a14893ed13f212d494606afa839f49b99791e62e"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 7ba7e8b3b727 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-76e93b9fd13530b3d6eb419fab5ea48716b0e6aba92c28249c4c64cb2858b6dc)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-9f428dbe32b8bd6e33a6f56588d0ec98cf6097410f8c9b72835772e842e8832c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-37d564b1f1ab7a43092e8a9a82a98aa24ce9a3451484b783c5f4ef12e1856e06)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-76e93b9fd13530b3d6eb419fab5ea48716b0e6aba92c28249c4c64cb2858b6dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4bd89c19a567a89211ece009d27894c53aa850ee35291937353a3ac6a091045"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 9ae6290ab4a1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-4338ce2424d635b304af2804ae5d4cd3f6e1e25044d3427032534956e9636853"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e3b1dad2ee7059d9f524200d7ffcc0f321d935f354ebe3f6721b0036934be2c1"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 9ae6290ab4a1 / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-03e840c02bcc0f52a5f4ee58d070f370e4b00aab46ae5455f8b2a98b0641df06): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-cde511a6a84c480da8291221a3b38c5f975254bed6acaa225255649b11f3381e): complete subsection reference.

<a id="canonical-b69fc423ac0bc94080ee7a6e4fc921f25fe5eedb1913fedc9d76f5c15a0f325a"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 9ae6290ab4a1 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-03e840c02bcc0f52a5f4ee58d070f370e4b00aab46ae5455f8b2a98b0641df06)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-cde511a6a84c480da8291221a3b38c5f975254bed6acaa225255649b11f3381e)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-03e840c02bcc0f52a5f4ee58d070f370e4b00aab46ae5455f8b2a98b0641df06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b237cc92f5e57b158c4fda071431c17ae9958fc869f1bab27aee5792e719d2bb"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c97c20901aa3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-76e93b9fd13530b3d6eb419fab5ea48716b0e6aba92c28249c4c64cb2858b6dc)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-2e9b4967893129f41acf1ad447ac6c1b7b9a4668141bf58cb3102332657e4e32"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4fc6c8a23f05f44894d0f9acad0815fe045ec5b4acf6abefbdbbd36affe1ed76"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c97c20901aa3 / 3

<a id="canonical-b70096f63ce4685096e9a8ebaaa187c6448406b4994c7da447da44eb6bb3e04f"></a>

<a id="canonical-c772e7fda2f55cb9313a246204931c6f78f8877cd3e5957190113f523ff29e0d"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c97c20901aa3 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-cb282513439e300096a21c7efff234e38febdca8d67c0a0ca5098f47ca98b52b"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c97c20901aa3 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-76e93b9fd13530b3d6eb419fab5ea48716b0e6aba92c28249c4c64cb2858b6dc)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-cde511a6a84c480da8291221a3b38c5f975254bed6acaa225255649b11f3381e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e81da596fd10b53528816739035eea7152c606d6d47f5dd5ceacbd26dfc84481"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 15cf333ac462 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-76e93b9fd13530b3d6eb419fab5ea48716b0e6aba92c28249c4c64cb2858b6dc)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-e298eb13f193f3a4f23a369acbe4ed1d71d58debd259174bf6877415ac0759fe"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cc119e210a2fbe65d452784f48de8dfe12ff2ff83baac06d02ce8d37569f14f3"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 15cf333ac462 / 3

<a id="canonical-f9eae94d4a79d877ce47c2549d9300fa628d3bd50f0deedb519493f6420b8e33"></a>

<a id="canonical-83550fd6ecb5442cbdbb990f984ba2c1ba05452ee10678ebfe4ba00ff0c5367c"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 15cf333ac462 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-0d7fe2761dd9e0fd7b16644030536ae1d09d82d2c9ce87e6ea2530e62284334b"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 15cf333ac462 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-005.md#canonical-76e93b9fd13530b3d6eb419fab5ea48716b0e6aba92c28249c4c64cb2858b6dc)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-9f428dbe32b8bd6e33a6f56588d0ec98cf6097410f8c9b72835772e842e8832c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef37c0f4e24bf4bba7626a9459f1f38b15a72278ea4326c9d85863fde8d77481"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 4cc9ee379e0a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-9f42ab9fdb962a4e7bb4b6b7fb0b79c4c1457c51fa933ee9c6319c9a2404f94b"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ae8cbb80b51dc73f2aa9f1178a6342401691bc3df219d41e4c784e56f4032036"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 4cc9ee379e0a / 3

<a id="canonical-28b6f8c505d51ae51c94fc5124dcc422a38855ae240bdce8f5a02bf7184f21a1"></a>

<a id="canonical-76aa754f87a56c338490d528ddb6ae821917d6adf6c83a7034577094ff55383a"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 4cc9ee379e0a / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-f81c01df50db5c323540f61ab6c0f83902f28d87a91ac2ff2cfacc3575e2dece"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 4cc9ee379e0a / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-37d564b1f1ab7a43092e8a9a82a98aa24ce9a3451484b783c5f4ef12e1856e06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f77d351279facad2c82c05f830e4173458615f83b6c7567478f9d0f55a822664"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d5675e8b8c6a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-004.md#canonical-c9243f158c58281f5e253b5f1ac55f1e7422e63e27cab008d1a1d63a39e9b8c9)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-fc6c0df965982a99efc7eaef6bffd787fb3827deae551c818f45380a18b42d9e"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2de01ead6bfb50f522b0eea75ed62a20c99d302f3ad80c81cbbde1061f12c6ab"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d5675e8b8c6a / 3

<a id="canonical-8db301cb3a997aa7af71f7bac6f1bde952b31f988d1a98c120c8710ee297954f"></a>

<a id="canonical-9b4dd5619e057c1775b73a8fe6a5d9f89af356cf76c20cb0f3db6a264c050010"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d5675e8b8c6a / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-81481f227583d4a34f304f01d136ffbfc5b1dc71efc8bc62782ec6546e281d61"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / d5675e8b8c6a / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-005.md#canonical-f8fc1a84beccabb67d2f2b45de1c8652cd289bd2acbdbe103fdb83c6db74b79b)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-63f5c0f674002c7865e6c701692cbf192783d14f999679ef0eab9ed82799de00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70b1f6de681b7d7afccd965f54f8077365aed76e9b905f5f73169ffb600c2b1a"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 506d98af1120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-369e9d9cddf536794ea09dc9084fb7c269319bb2080c1ac9abd8470471da2fa6"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-12732bd12944b498ad2df2b87dc2dc7e08d46fa969ff8e25bb55aef6f261594b"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 506d98af1120 / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-93eb3a166e34367c00b6f0285e610801bec8bf76770ed41b6e592164d71ccaaf): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-d3017e313c49d393bba41c7ddb968c24a24c1a888f0ef88c2d6f0aff7a16bd0e): complete subsection reference.

<a id="canonical-69e295ae977bc60ca6675b7c3009a3f06a5c5cb100ea456bc571efa0bc7c68b5"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 506d98af1120 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-005.md#canonical-93eb3a166e34367c00b6f0285e610801bec8bf76770ed41b6e592164d71ccaaf)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-005.md#canonical-d3017e313c49d393bba41c7ddb968c24a24c1a888f0ef88c2d6f0aff7a16bd0e)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-93eb3a166e34367c00b6f0285e610801bec8bf76770ed41b6e592164d71ccaaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c71534ac24f3eedde8e28eee653519e5f04763d5ad1d895e92794937b262dcea"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6557bf92e63c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-63f5c0f674002c7865e6c701692cbf192783d14f999679ef0eab9ed82799de00)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-5211161a33938209bc58b1623030e38d9eab027ef5e89a3359ed73cda01054c7"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2716143244d0e654d1ad2d980c8463b35fa8cb9ebadb0053536c44ae9a926cf6"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6557bf92e63c / 3

<a id="canonical-964151b2a2e81182a5c159eddc1d64bc6cc02957fb157b5904dc899ad1c5c220"></a>

<a id="canonical-f4170690748b7f980ff9fa6b9ce3ac33ef5f224b03a582ca9265d23158b3103d"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6557bf92e63c / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-46667366813075b46b084e52d7ba1e1860aa2ed7b7074833c08e22989c390498"></a>

<a id="canonical-ea4683559d5ce4736810d54d36b407532e80c2a61cae9c6468145ade1efbd6e4"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6557bf92e63c / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-1eb4cb9a0f842904dc857978e014f4c36c9a8c9db9a4a601d77d78ef1ded2b73"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6557bf92e63c / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-63f5c0f674002c7865e6c701692cbf192783d14f999679ef0eab9ed82799de00)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d3017e313c49d393bba41c7ddb968c24a24c1a888f0ef88c2d6f0aff7a16bd0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08699eed745371d7e4767ba9af65a36a4f271315351810b2653408f0875f4dcf"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 98db1bfc8736 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.outside_static_routes](data-sources--azure_vnet_site--reference--group-004.md#canonical-fdc114aeafc8f82ca9271b0c67cc328e8f38e8fd7646204058e84c77c3e8032c)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-004.md#canonical-554a2f6827068e15f956d36caa05ebb27d8e8735ba95defd5eb401ad32d009f5)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-004.md#canonical-7e029296414c8fbfff2a83ea8411e7dcf4e129ea5b2b49097234c9821f1ca130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-63f5c0f674002c7865e6c701692cbf192783d14f999679ef0eab9ed82799de00)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-82234e289d9bcb04d2204ef957e17eaa0186ee6c96c433c3755c6939642c3f22"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-37501fd4b34d31124264ecae70d32df2e4280718060473555b2b4994d63895d7"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 98db1bfc8736 / 3

<a id="canonical-1fb3d5159b4898625d06a2a7b8529552af7714061ca241261b22fd6d3dbac1b2"></a>

<a id="canonical-1c6bde82211932464f8ef0bd46e4b299d37111f55c60165c4df7906355f04445"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 98db1bfc8736 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-e70410e3800e28b70d4ffee550bfe3760de7906aa151fbd65ea7343d33cc59ea"></a>

<a id="canonical-bf7452441746dc433b7fbeb361c3ef6c1f9ece7e844b814ef6d1bef48152d1f6"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 98db1bfc8736 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-cf7839dc50d212c2cc8c7dcbb6b957639c989034458e9c8a8c49115640e3e790"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 98db1bfc8736 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-005.md#canonical-63f5c0f674002c7865e6c701692cbf192783d14f999679ef0eab9ed82799de00)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d77a6afb01a85ac4b8bbfa8a0f32cafb841fe9eecefa07b9da8c2879eb5c53f"></a>

## ingress_egress_gw.performance_enhancement_mode — ingress_egress_gw.performance_enhancement_mode / 9c10ef98eb1c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-4e4d39a53cede36e9c700262ca9a5b8b5434c28576cc3f6937d563d59f9b2c03"></a>

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

<a id="canonical-6bd0add6fe997796a588c88acd26dfe972ecceaedc607919dd2a16edaf152380"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode / 9c10ef98eb1c / 3

- [perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-b32ac1bb82fe36f98081a64060d0cd73635600f4ead64d5aef407a43094a012e): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-fa09517c3200a7c933e3199148de3ef016442b7f71026d091ee2ab4f685e6dfa): complete subsection reference.

<a id="canonical-478d509b14629bfbbacf88099b60c4c3d7e7ad7dea30744ac392a844e7dc176a"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode / 9c10ef98eb1c / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-b32ac1bb82fe36f98081a64060d0cd73635600f4ead64d5aef407a43094a012e)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-fa09517c3200a7c933e3199148de3ef016442b7f71026d091ee2ab4f685e6dfa)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b32ac1bb82fe36f98081a64060d0cd73635600f4ead64d5aef407a43094a012e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2655a08133d727e877ce47c4f19bd810946d00ca1b6fc409b968cfde2483e674"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 774e188e0f79 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-883fa83559d9f259299f7505e482a5891546bf84825bc515d3a497c0ad343338"></a>

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

<a id="canonical-374f27cb7dc4b1e79f04dda816a6a3714e6e084414523102a46bcecb668d531e"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 774e188e0f79 / 3

- [jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-ed6fbc8f39cde4130a1899e276e5edf016042a1b392dc0019a8410b87190ea07): complete subsection reference.

- [no_jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-f626416691fc344750957755ccb371f5f1255a22975fdb477d75ec019fe5f8f2): complete subsection reference.

<a id="canonical-c060aac1f5faf34c8ec2cbfe51ab19ac846f4149b4e92e8e5d96395cb5369b47"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 774e188e0f79 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-ed6fbc8f39cde4130a1899e276e5edf016042a1b392dc0019a8410b87190ea07)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--azure_vnet_site--reference--group-005.md#canonical-f626416691fc344750957755ccb371f5f1255a22975fdb477d75ec019fe5f8f2)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ed6fbc8f39cde4130a1899e276e5edf016042a1b392dc0019a8410b87190ea07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcfebfccf718be38cc7d58daf7349dd8f361421bd8b274e51ccb855528365786"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 48b760a7c5b3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-b32ac1bb82fe36f98081a64060d0cd73635600f4ead64d5aef407a43094a012e)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-fe5bd91209de55ac3bb5c8e8d411d0f28ac67f630964d9baf251e984cc30121b"></a>

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

<a id="canonical-40f9afae0fd8cd6991e96395be136b415bf30e9fef5e637c75a3969ca88fe6d3"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 48b760a7c5b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15bc6e6f58ad13f4f3fa757f552224b9e208cb880c2212ac1518701745c87deb"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 48b760a7c5b3 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-b32ac1bb82fe36f98081a64060d0cd73635600f4ead64d5aef407a43094a012e)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f626416691fc344750957755ccb371f5f1255a22975fdb477d75ec019fe5f8f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4c8dc4a7e124c85afa48d667a9ea1065029b97aad3e72a7aa6f26f31344d7fa"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / b42924f66ff1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-b32ac1bb82fe36f98081a64060d0cd73635600f4ead64d5aef407a43094a012e)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-1fdc957ba33410c615f255722b099630814c9434a52be50933e37e95eeb4a23e"></a>

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

<a id="canonical-69f0421b8a7673992f1d871cbca0fa4f669b40d9331529fcb619e963d42c2ed3"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / b42924f66ff1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9b40a2345d20a2967000882be1375bff5dab828e591f13934b193f26a4ee550d"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / b42924f66ff1 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-b32ac1bb82fe36f98081a64060d0cd73635600f4ead64d5aef407a43094a012e)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-fa09517c3200a7c933e3199148de3ef016442b7f71026d091ee2ab4f685e6dfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef2ef9324adde9046434337d1d93bafaa5c3c3e9d8f0cd77ba7ec90f94eb1ff0"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 561eb5190c8a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-6ea5d8ff59c3258a564ea0bb5ffa832c40d6691ff0b9412499b3ce7286cdf57e"></a>

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

<a id="canonical-a55675cc8239a4df3579023a572b8c6ba648ce1264299464b23710ff7607ce0d"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 561eb5190c8a / 3

- [jumbo_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-330cc5d8649fa9f581bd9733774e29132ba94a6ac5c876df16a9bca4dddb2f7a): complete subsection reference.

- [jumbo_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-7695696a2c33150a6a4db01a6426877915136f7a8ffbb9b47c666f4dddff6457): complete subsection reference.

<a id="canonical-bcb64cfbbd330cf4c36fefdee29f2da2bd1de903580944beeec04883fc2a0f52"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 561eb5190c8a / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-330cc5d8649fa9f581bd9733774e29132ba94a6ac5c876df16a9bca4dddb2f7a)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-7695696a2c33150a6a4db01a6426877915136f7a8ffbb9b47c666f4dddff6457)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-330cc5d8649fa9f581bd9733774e29132ba94a6ac5c876df16a9bca4dddb2f7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c9bfeacd89d6dac34fa8c279c486ed5147e2f92a8ac0618578e7e6320983949"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 90083d358249 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-fa09517c3200a7c933e3199148de3ef016442b7f71026d091ee2ab4f685e6dfa)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-894501e7a029bb2916cf05de23f4dd2b62ce16ee9f5344eaf2f0e8b4ce35fdb3"></a>

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

<a id="canonical-0d2ee128309446e0466f9fec1c6f165437f7cda7ac070d9454493f7ee345cf70"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 90083d358249 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f48dcedd293cce263b727ce138b789326091de51cf30f487fb1544d7c13018e"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 90083d358249 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-fa09517c3200a7c933e3199148de3ef016442b7f71026d091ee2ab4f685e6dfa)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-7695696a2c33150a6a4db01a6426877915136f7a8ffbb9b47c666f4dddff6457"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6f0b46dc10dbfafbd7ae3282e975c15814abc720ed5f1c9e3f8a75d2bf02c7c"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 097b027f0541 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-005.md#canonical-f86faec1f2ef8fcf03f0ce7d864ed9b319a8e6f8fd0e7029a3934574202e9a79)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-fa09517c3200a7c933e3199148de3ef016442b7f71026d091ee2ab4f685e6dfa)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-dc836f0c2ebba1180770f88a6214043d4e42d681aae3f05a6c4c2b67b28b1603"></a>

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

<a id="canonical-15b448fec0a5554e99830366de0a563067d9448265862f0863095b355096f548"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 097b027f0541 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-45af11d6579bb76da00b45db8ea8d4f2df36897a859ad820f362a3329ea3d2ea"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 097b027f0541 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--azure_vnet_site--reference--group-005.md#canonical-fa09517c3200a7c933e3199148de3ef016442b7f71026d091ee2ab4f685e6dfa)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-e16def723f4eaec10e2a191557f03b595da763ff6021c47d0874a2c7bf66a4b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-878258ea59efbea8500216ccce79a4a082209a64f8e2721002ab167c19871237"></a>

## ingress_egress_gw.sm_connection_public_ip — ingress_egress_gw.sm_connection_public_ip / fd6029ad9012 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-d2079f24fdf9ab4a3e9bda73b7ffb01fad6e820554c4ed2d09f62c34062773ca"></a>

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

<a id="canonical-9b6ba76cc00930a29a146ad76d8007229dcae8f56a9695ff27732348729bfe62"></a>

## Direct properties — ingress_egress_gw.sm_connection_public_ip / fd6029ad9012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2d36072b565306a56a30c43219d4142cd7f101e70e75868cbb5d661723a2378"></a>

## Next pages — ingress_egress_gw.sm_connection_public_ip / fd6029ad9012 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-175603d3759575b22e9eee3f0c857a3d20732b4949b69907bb58b4b791a5c7f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78a43fac930692fb302d3b9b0c25761452f6252b9a29c8185494a330c1f24fef"></a>

## ingress_egress_gw.sm_connection_pvt_ip — ingress_egress_gw.sm_connection_pvt_ip / df5716042c78 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-0a53a1eb5d973a8f4cc0ad19ca618df33663f55f51c9d1c2e68c169d40780de6"></a>

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

<a id="canonical-0e630d6b26a1b5ba3c4759e3fdfa489f0c46783ff34c42e93826f85da7e4a282"></a>

## Direct properties — ingress_egress_gw.sm_connection_pvt_ip / df5716042c78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9b64313ec20cdd427b1f4e655d595f1318af801447b0bcc6c611d8a0367655f"></a>

## Next pages — ingress_egress_gw.sm_connection_pvt_ip / df5716042c78 / 4

- [ingress_egress_gw](data-sources--azure_vnet_site--reference--group-003.md#canonical-62bd839d85e58a58fd242f704ee2ebdffdbab208160d824c7aba98fb9a0caf16)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32dd20d6f6b2bfef79943625402be9bbc28e245c6389649cdaf1d7df19788bf2"></a>

## ingress_egress_gw_ar — ingress_egress_gw_ar / 5bbbeb9d27bb / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- ingress_egress_gw_ar

<a id="canonical-ae000b2f481efeb407e0e1bb074a410080d294c06b43aa2cd5931e3ea8221cde"></a>

Type: `"single"`. Computed.

Two interface Azure ingress/egress site on Alternate Region with no support for zones.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-hub_choice": "[\"hub\",\"not_hub\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

<a id="canonical-b3ad0af827d5f8cadcb34337782ba0937f5004603875d44e9a64c7f181f13363"></a>

## Direct properties — ingress_egress_gw_ar / 5bbbeb9d27bb / 3

- [accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-ec6e2e0d3c410dde27d221253263df0c0f31cae3a44e855a8e5bc75964d13d9c): complete subsection reference.

- [active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-8c975ff7e5d6e57534cc559384c22d77107e562dd5de9299f98f53f151460470): complete subsection reference.

- [active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-b1c872cfbcfb1135ba0bc06fa813ec0c68a90ef019b59ba44626c8a853c6fdf4): complete subsection reference.

- [active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-beaa5a73cabb86215df5883a42811b828ca6a68712b7910eb715fe075a8e3a64): complete subsection reference.

<a id="canonical-68366a70ca65c3a5b90dab1e9d8028061e2301be691262221d3ef5805a46bfcd"></a>

<a id="canonical-353a28df99e03c341aa27e07b79d7bf3f0edf42498a00e50c7e2e9b9a39e87c7"></a>

## azure_certified_hw property — ingress_egress_gw_ar / 5bbbeb9d27bb / 4

Type: `"string"`. Computed.

\[Enum: azure-byol-multi-nic-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-multi-nic-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group_inside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-6abdc93bf697a046924dd3adf2aec5bdfade91ec6b278f4ceccb07c0d55d7781): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-b8183756be5efce221052ffd7b2c46bb3992737ec370a7d763f98ae6ecc65cdb): complete subsection reference.

- [forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-005.md#canonical-8c6e032f6f419ec4d25e3364f8533783f4faabb58e3f538dca36d21b7458e3dc): complete subsection reference.

- [global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183): complete subsection reference.

- [hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6): complete subsection reference.

- [inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67): complete subsection reference.

- [no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-1adc55cf4e964eca81eda4ae9033b5497809c0dd2e417faee5856f72d7912936): complete subsection reference.

- [no_forward_proxy](data-sources--azure_vnet_site--reference--group-006.md#canonical-9104b6a30f4a654bf9bcdc4f16f1561bf3b0e8ee75a843ef3c1de495ac60ea7e): complete subsection reference.

- [no_global_network](data-sources--azure_vnet_site--reference--group-006.md#canonical-ae5945d1bb672f90650e08c6886fd4978545223f3e8e51d824d7579ae7d2f074): complete subsection reference.

- [no_inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-b1ffdac3082b8b06a008d52c4072ccac2a90903b13767ca88f126b4ca6138998): complete subsection reference.

- [no_network_policy](data-sources--azure_vnet_site--reference--group-006.md#canonical-fdda130c92b94013b60116a4bbba0410e68076b4a47c6dc701bcc29bbb1bb8ad): complete subsection reference.

- [no_outside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-6a379509da9c1ee5a61e9229fb3c0c0b629543b3a032d49ff588e2d82f923fb3): complete subsection reference.

- [node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c): complete subsection reference.

- [not_hub](data-sources--azure_vnet_site--reference--group-007.md#canonical-4487197c09ca09c9d4c1dcfc2e8a477907d99dc54bb4d5e0672be14f5d6780b2): complete subsection reference.

- [outside_static_routes](data-sources--azure_vnet_site--reference--group-007.md#canonical-f3d519299980415ba99fda6d5b233ee4f8c3084cc3ffa6e447f95638f40aa3a0): complete subsection reference.

- [performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-007.md#canonical-33b2f64cd384fa152f6eb45d27579b468bd36683fbfee5fe5edf07ade573b1af): complete subsection reference.

- [sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-fd9201b3774a879d4cedaa3202c74b5d1c3df7a10f97b388af22849f56aa7085): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-22acf03628be14315205a4f748a7677298fb78200c1a36cf89f924e8f876da25): complete subsection reference.

<a id="canonical-025a3fec8fc7febdf7f6849907b3dd5a7ef1b0f7d4f430228979f410b7e29d45"></a>

## Next pages — ingress_egress_gw_ar / 5bbbeb9d27bb / 5

- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-ec6e2e0d3c410dde27d221253263df0c0f31cae3a44e855a8e5bc75964d13d9c)
- [ingress_egress_gw_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-8c975ff7e5d6e57534cc559384c22d77107e562dd5de9299f98f53f151460470)
- [ingress_egress_gw_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-b1c872cfbcfb1135ba0bc06fa813ec0c68a90ef019b59ba44626c8a853c6fdf4)
- [ingress_egress_gw_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-beaa5a73cabb86215df5883a42811b828ca6a68712b7910eb715fe075a8e3a64)
- [ingress_egress_gw_ar.dc_cluster_group_inside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-6abdc93bf697a046924dd3adf2aec5bdfade91ec6b278f4ceccb07c0d55d7781)
- [ingress_egress_gw_ar.dc_cluster_group_outside_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-b8183756be5efce221052ffd7b2c46bb3992737ec370a7d763f98ae6ecc65cdb)
- [ingress_egress_gw_ar.forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-005.md#canonical-8c6e032f6f419ec4d25e3364f8533783f4faabb58e3f538dca36d21b7458e3dc)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-9f508b120e0aab82c790e7a53bad089d5b00aac91b98ed8e3471bba7b9197a67)
- [ingress_egress_gw_ar.no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-1adc55cf4e964eca81eda4ae9033b5497809c0dd2e417faee5856f72d7912936)
- [ingress_egress_gw_ar.no_forward_proxy](data-sources--azure_vnet_site--reference--group-006.md#canonical-9104b6a30f4a654bf9bcdc4f16f1561bf3b0e8ee75a843ef3c1de495ac60ea7e)
- [ingress_egress_gw_ar.no_global_network](data-sources--azure_vnet_site--reference--group-006.md#canonical-ae5945d1bb672f90650e08c6886fd4978545223f3e8e51d824d7579ae7d2f074)
- [ingress_egress_gw_ar.no_inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-b1ffdac3082b8b06a008d52c4072ccac2a90903b13767ca88f126b4ca6138998)
- [ingress_egress_gw_ar.no_network_policy](data-sources--azure_vnet_site--reference--group-006.md#canonical-fdda130c92b94013b60116a4bbba0410e68076b4a47c6dc701bcc29bbb1bb8ad)
- [ingress_egress_gw_ar.no_outside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-6a379509da9c1ee5a61e9229fb3c0c0b629543b3a032d49ff588e2d82f923fb3)
- [ingress_egress_gw_ar.node](data-sources--azure_vnet_site--reference--group-006.md#canonical-31fa9317ade280d019ac0c4610154c8d18e89ceb533de5fd1c1e050160203c9c)
- [ingress_egress_gw_ar.not_hub](data-sources--azure_vnet_site--reference--group-007.md#canonical-4487197c09ca09c9d4c1dcfc2e8a477907d99dc54bb4d5e0672be14f5d6780b2)
- [ingress_egress_gw_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-007.md#canonical-f3d519299980415ba99fda6d5b233ee4f8c3084cc3ffa6e447f95638f40aa3a0)
- [ingress_egress_gw_ar.performance_enhancement_mode](data-sources--azure_vnet_site--reference--group-007.md#canonical-33b2f64cd384fa152f6eb45d27579b468bd36683fbfee5fe5edf07ade573b1af)
- [ingress_egress_gw_ar.sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-fd9201b3774a879d4cedaa3202c74b5d1c3df7a10f97b388af22849f56aa7085)
- [ingress_egress_gw_ar.sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-007.md#canonical-22acf03628be14315205a4f748a7677298fb78200c1a36cf89f924e8f876da25)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ec6e2e0d3c410dde27d221253263df0c0f31cae3a44e855a8e5bc75964d13d9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6a9df661735ed950525c190c9d3e6a25a18e1568b09b4058ec685dd3b49c452"></a>

## ingress_egress_gw_ar.accelerated_networking — ingress_egress_gw_ar.accelerated_networking / 407507a1a007 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.accelerated_networking

<a id="canonical-e229931fa20bac5a7be5b5df2f8377f055b72471eeb19b73cdc2eec5a6143916"></a>

Type: `"single"`. Computed.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-2ffa2e4c96f6b7cfdcf710db2858626846c4f19e83e8326441ca2287d9be57ec"></a>

## Direct properties — ingress_egress_gw_ar.accelerated_networking / 407507a1a007 / 3

- [disable_spec](data-sources--azure_vnet_site--reference--group-005.md#canonical-b6b1f2e76044536fdeca235ef33fea52704caac32b9b91e4ec5152e7dab39287): complete subsection reference.

- [enable](data-sources--azure_vnet_site--reference--group-005.md#canonical-bf9df11959cc7b2804b7d9e2c38df5e8d1596af793fda47eb325080d8f002edb): complete subsection reference.

<a id="canonical-1ba7d0234811db0f3cd42beba4f9764c787e64ce3f9e552fc9c5012a2a36208e"></a>

## Next pages — ingress_egress_gw_ar.accelerated_networking / 407507a1a007 / 4

- [ingress_egress_gw_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--reference--group-005.md#canonical-b6b1f2e76044536fdeca235ef33fea52704caac32b9b91e4ec5152e7dab39287)
- [ingress_egress_gw_ar.accelerated_networking.enable](data-sources--azure_vnet_site--reference--group-005.md#canonical-bf9df11959cc7b2804b7d9e2c38df5e8d1596af793fda47eb325080d8f002edb)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b6b1f2e76044536fdeca235ef33fea52704caac32b9b91e4ec5152e7dab39287"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-671e5fa18b736d90888086b5bfed2867532619b7fd7706dc0a70a94049f2708f"></a>

## ingress_egress_gw_ar.accelerated_networking.disable_spec — ingress_egress_gw_ar.accelerated_networking.disable_spec / 904a4e511f68 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-ec6e2e0d3c410dde27d221253263df0c0f31cae3a44e855a8e5bc75964d13d9c)
- ingress_egress_gw_ar.accelerated_networking.disable_spec

<a id="canonical-37b394c01a3e40c93b0c67adb3eaf3833e2cac2696bd91396e25b5aa9e5fbc67"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-1227f2ef43c3e042befc97d999149c837fae5e5f63fb7b83a01ed340cd74a36e"></a>

## Direct properties — ingress_egress_gw_ar.accelerated_networking.disable_spec / 904a4e511f68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-095bb1172c7ccd1f220516316829b04e271c097e69c15bc995f59bf4f0bfe1fb"></a>

## Next pages — ingress_egress_gw_ar.accelerated_networking.disable_spec / 904a4e511f68 / 4

- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-ec6e2e0d3c410dde27d221253263df0c0f31cae3a44e855a8e5bc75964d13d9c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-bf9df11959cc7b2804b7d9e2c38df5e8d1596af793fda47eb325080d8f002edb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b09e1358dd5aa309cb59c94e8ce1883bcf6f511eae30085fd9049551b4cc9042"></a>

## ingress_egress_gw_ar.accelerated_networking.enable — ingress_egress_gw_ar.accelerated_networking.enable / b41507540e82 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-ec6e2e0d3c410dde27d221253263df0c0f31cae3a44e855a8e5bc75964d13d9c)
- ingress_egress_gw_ar.accelerated_networking.enable

<a id="canonical-200cec9ee075a711f508015e990ca754fafc0ca65aa6abfd962ae0536f766367"></a>

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

<a id="canonical-a05c084dbf5e62da4b9cecb61ddaa0ad2d6d00d817296fa8101d147841d5328a"></a>

## Direct properties — ingress_egress_gw_ar.accelerated_networking.enable / b41507540e82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0d55659256c1d1cefa6d6f3773e8b330831bb722a0618610939035022bc6a49"></a>

## Next pages — ingress_egress_gw_ar.accelerated_networking.enable / b41507540e82 / 4

- [ingress_egress_gw_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-005.md#canonical-ec6e2e0d3c410dde27d221253263df0c0f31cae3a44e855a8e5bc75964d13d9c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8c975ff7e5d6e57534cc559384c22d77107e562dd5de9299f98f53f151460470"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc0455f3565224b4fc3a940ea521e696ad337702d57782229063d8561cf5b18d"></a>

## ingress_egress_gw_ar.active_enhanced_firewall_policies — ingress_egress_gw_ar.active_enhanced_firewall_policies / fe4c91dfe2fa / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.active_enhanced_firewall_policies

<a id="canonical-c5bc38c95f6c5ab3de7e34914f4669227454db737ea1e8d19c5ca3e78f832926"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b5bad89cbd62467592bc170c5be2359b159c612b6fbf7b0046e6f27dd1210bc6"></a>

## Direct properties — ingress_egress_gw_ar.active_enhanced_firewall_policies / fe4c91dfe2fa / 3

- [enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-00c72e5a62f2d538ef4fff267ed562446e5164bc4aea85787e22dcbd89af3535): complete subsection reference.

<a id="canonical-91a6de149c8f87ad305568a249ddcbe01f120517bc747e5556eb36e5bb5731af"></a>

## Next pages — ingress_egress_gw_ar.active_enhanced_firewall_policies / fe4c91dfe2fa / 4

- [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-00c72e5a62f2d538ef4fff267ed562446e5164bc4aea85787e22dcbd89af3535)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-00c72e5a62f2d538ef4fff267ed562446e5164bc4aea85787e22dcbd89af3535"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8d76fa6bed20b47d7c905b053b01c09db850f4cad1fedb485ec748f1e2bc867"></a>

## ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / e3245e6a0338 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-8c975ff7e5d6e57534cc559384c22d77107e562dd5de9299f98f53f151460470)
- ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-63632cb37e47c1ed66cc0239a198cb0fa29524c13dd9a8e3d1017526087bf482"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-82f848da3d074d9dc0771687717a0adfd97f34e232ff2a67ec1072979b966cd7"></a>

## Direct properties — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / e3245e6a0338 / 3

<a id="canonical-2b0560729553c11bd3699fc7cc7ff907d46c1badb03b6153c14d0cfbadcd7df0"></a>

<a id="canonical-57507f88885cb194c3da28202ad60190fbfb2ed1b3fd653707fde7a2d6aa9a31"></a>

## name property — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / e3245e6a0338 / 4

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

<a id="canonical-a666d1da0cd56846be578701958432f60e2b8e80c09ad65b8c985179d8007dd9"></a>

<a id="canonical-d109ce7d448f04d742a4e63ee51c4350e174799bfe09c4cdc3a6e35ceae51f02"></a>

## namespace property — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / e3245e6a0338 / 5

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

<a id="canonical-177e5423b6d7dac40b84561a42c36a1fb598c0f0b6f72eae1fd6ce218b7056c9"></a>

<a id="canonical-6b2293ffe6210abd253b983cb11dd4e495b8febeccfbab93af3596f9ab9c569d"></a>

## tenant property — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / e3245e6a0338 / 6

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

<a id="canonical-99352ed619d11c20896b902bd0a74edddd7f34fc92a8989d947132e5a7e2c66b"></a>

## Next pages — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / e3245e6a0338 / 7

- [ingress_egress_gw_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-8c975ff7e5d6e57534cc559384c22d77107e562dd5de9299f98f53f151460470)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b1c872cfbcfb1135ba0bc06fa813ec0c68a90ef019b59ba44626c8a853c6fdf4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95880f154d87b2bdff47affb161f2bfcde11d53466623dfe35cf816a11b0112f"></a>

## ingress_egress_gw_ar.active_forward_proxy_policies — ingress_egress_gw_ar.active_forward_proxy_policies / 7238d8a9c0b8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.active_forward_proxy_policies

<a id="canonical-f1deca7a9d32f21cf65ceb84bf5433f443bd1e3b6755667d10089b3b27f12679"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9906f84be5b26e94ba0dffb2c06d831b1568873df460470f65168c69284cf82c"></a>

## Direct properties — ingress_egress_gw_ar.active_forward_proxy_policies / 7238d8a9c0b8 / 3

- [forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2290901781ad17d9dcc27d9b0a3c681ae416787df19f77e0bceae40ae0204f53): complete subsection reference.

<a id="canonical-6fca4390f010ddd78567df69b79d7f70617d4190fe504d4503979996861400a3"></a>

## Next pages — ingress_egress_gw_ar.active_forward_proxy_policies / 7238d8a9c0b8 / 4

- [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-2290901781ad17d9dcc27d9b0a3c681ae416787df19f77e0bceae40ae0204f53)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-2290901781ad17d9dcc27d9b0a3c681ae416787df19f77e0bceae40ae0204f53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c7198177cd2985954b50891dd30f4faf7dfc62817efc48135125c30f03cd273"></a>

## ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / e06f2b2f6846 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-b1c872cfbcfb1135ba0bc06fa813ec0c68a90ef019b59ba44626c8a853c6fdf4)
- ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-1cd6fd4eab13d197321b90fd61cd25b502eca4acfda1744b3e507421d14cd582"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-7fefa0b5c8d689dd4528dd418e69e7880c6bf65f905a43ba70140bc74ddbae2a"></a>

## Direct properties — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / e06f2b2f6846 / 3

<a id="canonical-1b1064c40934d9ba1554a0d2571fcba5901dedea644092ff74a75e5989488bc6"></a>

<a id="canonical-6e4335c3c307c492c4ffb9a9ecfe9b8325ceeeae160b8970537e8f4a0c334c55"></a>

## name property — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / e06f2b2f6846 / 4

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

<a id="canonical-d5c45c0da58839f54c2b13c2ef06c98b3e449217e687ff6b2516cfd08c176fc8"></a>

<a id="canonical-06ac2099a1f6b7ff0993f1e5f3b9d3b42bf841839aa252175ee1c692fee2cbd7"></a>

## namespace property — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / e06f2b2f6846 / 5

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

<a id="canonical-5d35fba77c00ca649fd5cb96577cf822131ffeec0dedc52e2009e73d7b05aff7"></a>

<a id="canonical-04921667d4714d7933d6a3e8b7311e3395a2ccc04d08b739ac45ec469a51a89e"></a>

## tenant property — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / e06f2b2f6846 / 6

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

<a id="canonical-44bf9019a59c499dfb8cd5065e674464544a8db63e418e9edc95a22ba03e4ba8"></a>

## Next pages — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / e06f2b2f6846 / 7

- [ingress_egress_gw_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-b1c872cfbcfb1135ba0bc06fa813ec0c68a90ef019b59ba44626c8a853c6fdf4)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-beaa5a73cabb86215df5883a42811b828ca6a68712b7910eb715fe075a8e3a64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-807c74ab5deb06761d03dc0b623a2ea4eff132792d0c3a361006595d1947507f"></a>

## ingress_egress_gw_ar.active_network_policies — ingress_egress_gw_ar.active_network_policies / a061bcb730fb / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.active_network_policies

<a id="canonical-807c85da68e521562586be7f73411ebe3fab6ef831f60f289d4e739b709aecbf"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0917034db4dd0fbb191a68742c78d3e60766dcbab33124e85008de710a8e85f1"></a>

## Direct properties — ingress_egress_gw_ar.active_network_policies / a061bcb730fb / 3

- [network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-84c45930daf0627c3202a7273b61531f5ad8ab1891557ae05496c5a0f01182dd): complete subsection reference.

<a id="canonical-af8eba915cd27399b94b18df6a77ca9e0379c7984ad6e34e0d5814e128a5886e"></a>

## Next pages — ingress_egress_gw_ar.active_network_policies / a061bcb730fb / 4

- [ingress_egress_gw_ar.active_network_policies.network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-84c45930daf0627c3202a7273b61531f5ad8ab1891557ae05496c5a0f01182dd)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-84c45930daf0627c3202a7273b61531f5ad8ab1891557ae05496c5a0f01182dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8444fd8d9fb7e163e3e2140af78f2ad9d75229829aaf1c8950fc783fcfa0f9c4"></a>

## ingress_egress_gw_ar.active_network_policies.network_policies — ingress_egress_gw_ar.active_network_policies.network_policies / 0da1396c25da / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-beaa5a73cabb86215df5883a42811b828ca6a68712b7910eb715fe075a8e3a64)
- ingress_egress_gw_ar.active_network_policies.network_policies

<a id="canonical-8a755abb6bcc2b5cbe64d71ec6d3a352cda9b70f36a35a0a2a489a72add9f528"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0c77a1096b9d17f5052e37a92559da849161cb37592cf058656f36f43d827701"></a>

## Direct properties — ingress_egress_gw_ar.active_network_policies.network_policies / 0da1396c25da / 3

<a id="canonical-886a69892e4aba1cb08b5081dd96b8331d0e4f0bb2b7072b8d2203091e3b87c4"></a>

<a id="canonical-f297fc677a813db4d8bbf3caef9883dd27d622bfac5d18a83614106a134d9e9b"></a>

## name property — ingress_egress_gw_ar.active_network_policies.network_policies / 0da1396c25da / 4

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

<a id="canonical-08652b3371c9b7d9993a75d433de0dadb32510bab51c3508f54ad2e6df891c46"></a>

<a id="canonical-efbae2deb73e7e444155d896104c241c6b66429d856b79769bfd608c963584fb"></a>

## namespace property — ingress_egress_gw_ar.active_network_policies.network_policies / 0da1396c25da / 5

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

<a id="canonical-c5b9c46f80b0c72c94d9ca54df37ad46f9ccf0b2b1ae74cf9fada0262f9ddf25"></a>

<a id="canonical-f389300324f6d525110e1e8781fcb404160b9a70a22c82b5cacc2415dee0102b"></a>

## tenant property — ingress_egress_gw_ar.active_network_policies.network_policies / 0da1396c25da / 6

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

<a id="canonical-e31c93f594ee442bad5d3fed8d111de9408052f79b70a02ccb6918038d53c37f"></a>

## Next pages — ingress_egress_gw_ar.active_network_policies.network_policies / 0da1396c25da / 7

- [ingress_egress_gw_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-005.md#canonical-beaa5a73cabb86215df5883a42811b828ca6a68712b7910eb715fe075a8e3a64)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-6abdc93bf697a046924dd3adf2aec5bdfade91ec6b278f4ceccb07c0d55d7781"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1ca3f31e14984416b275fc2fc81d1121bfb6ba73d1e944bc1d0da625a273840"></a>

## ingress_egress_gw_ar.dc_cluster_group_inside_vn — ingress_egress_gw_ar.dc_cluster_group_inside_vn / 9ed4ac800f0d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.dc_cluster_group_inside_vn

<a id="canonical-8999f2a4d6a1a504a293370edb00ea05dd53f612e04f56a99abd236182b096b0"></a>

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

<a id="canonical-32ee8c72ebe2047b876097f8918df27f7f2eef486db072dd3547257e8d4d4177"></a>

## Direct properties — ingress_egress_gw_ar.dc_cluster_group_inside_vn / 9ed4ac800f0d / 3

<a id="canonical-0c4daa0fac62a4e978a97ab562c65f163a1f7744057d4ea749b42b6284bca330"></a>

<a id="canonical-b0ddba6b276603971b5aece1c42f4b25e9673ce4db882b48aa547e1e14f40a5e"></a>

## name property — ingress_egress_gw_ar.dc_cluster_group_inside_vn / 9ed4ac800f0d / 4

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

<a id="canonical-bd574b0e2ea4ae948822029fcc2cc8e3652c6040a1cf4586666266a0e7288be5"></a>

<a id="canonical-915f95a92e66d70a6cbf6f35b3cf451f0bd64f81ce1ccaf6b8419d1075630122"></a>

## namespace property — ingress_egress_gw_ar.dc_cluster_group_inside_vn / 9ed4ac800f0d / 5

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

<a id="canonical-d898814ea209f3f8cebc9c953c5f6481853874f4500c37e0059e60fbf106091b"></a>

<a id="canonical-cf9250558cd2cdc59a8d1fd7cedd2b861fb5d87bff824a17e0e68f71f75c1a33"></a>

## tenant property — ingress_egress_gw_ar.dc_cluster_group_inside_vn / 9ed4ac800f0d / 6

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

<a id="canonical-a91422f68fa31dcebec8289c99e4067feaa3de621336c89ed69910d7419e777e"></a>

## Next pages — ingress_egress_gw_ar.dc_cluster_group_inside_vn / 9ed4ac800f0d / 7

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b8183756be5efce221052ffd7b2c46bb3992737ec370a7d763f98ae6ecc65cdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b049fed35501fba38e0e00e03e6844e61816ff304602d85e4b3034d12cfe63a5"></a>

## ingress_egress_gw_ar.dc_cluster_group_outside_vn — ingress_egress_gw_ar.dc_cluster_group_outside_vn / e96abd08d812 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.dc_cluster_group_outside_vn

<a id="canonical-bbd0555d7c7d491364ec45216070c312671363534b2df4acd8b64f2e0cdc1f0c"></a>

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

<a id="canonical-b8d88fd2a8af9b90d728b604efb991ec34c11dea88a8791e0b4d342e70b4554a"></a>

## Direct properties — ingress_egress_gw_ar.dc_cluster_group_outside_vn / e96abd08d812 / 3

<a id="canonical-66f50c1b87bbc712c3bca72d9a9f3055b944202a24e98350d5125ab18ebeda0b"></a>

<a id="canonical-6205958e0af1be4a77e7dc01bb1b5bec610af8c42a76a11cde531f9946e43aef"></a>

## name property — ingress_egress_gw_ar.dc_cluster_group_outside_vn / e96abd08d812 / 4

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

<a id="canonical-c1a0f4bea8ba64d2879ee8d78728b466d1f6ae965a7b4b7e5d72b48e6f6d0753"></a>

<a id="canonical-f647afbf71558b9d963ee76d333ac8f4ae9f9c7ed3e73b245b70c2c1387deb63"></a>

## namespace property — ingress_egress_gw_ar.dc_cluster_group_outside_vn / e96abd08d812 / 5

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

<a id="canonical-23d989893b1735d75d0e2a346100aa839f050d98827f86a3fbf1cd62c4c02e04"></a>

<a id="canonical-30e88fe3eb0f713ad470a2a020b37f7fe3080020af2be9109d32209264e754d5"></a>

## tenant property — ingress_egress_gw_ar.dc_cluster_group_outside_vn / e96abd08d812 / 6

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

<a id="canonical-916ee1f38d514243045a865ab9bfb681aa386e9dcd81098b94ac411c6bf5544c"></a>

## Next pages — ingress_egress_gw_ar.dc_cluster_group_outside_vn / e96abd08d812 / 7

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8c6e032f6f419ec4d25e3364f8533783f4faabb58e3f538dca36d21b7458e3dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5145221f5e21b2acd4c4976da1eaae3b7fe87fb1a77c0362764d84d12b55c09f"></a>

## ingress_egress_gw_ar.forward_proxy_allow_all — ingress_egress_gw_ar.forward_proxy_allow_all / 3ef5974b08c2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.forward_proxy_allow_all

<a id="canonical-bc693cde0363b0e4ac678018b332ee655b1daf18d60b07db28c410740cbf7658"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for forward proxy allow all.

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

<a id="canonical-6dcb53c007868066bdecb769967419ea7be3a505a3b2fd88a8e052d9a77da27f"></a>

## Direct properties — ingress_egress_gw_ar.forward_proxy_allow_all / 3ef5974b08c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a5c4c5171f94be3fe4194b887569e2cb61d6c68b1e10a5dcfa97489ceefa1d8"></a>

## Next pages — ingress_egress_gw_ar.forward_proxy_allow_all / 3ef5974b08c2 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5059037177b0c68e2467238bb26d8834e16a0483f9af1ad59c57f547a5fb9f70"></a>

## ingress_egress_gw_ar.global_network_list — ingress_egress_gw_ar.global_network_list / c4e9529352c9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.global_network_list

<a id="canonical-313a37a8932b26048a176ab6e5a237dd2358015cf8bcbd69561bf5e7d56f7709"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4d09b48ddf190ebea2e597d4eb917bb7b2b2989b967bbf07ab6fe0b902402680"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list / c4e9529352c9 / 3

- [global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3): complete subsection reference.

<a id="canonical-dc9e8bfff569b5eee292a0cd8e1ac190796f1ad50a84fb3ecf06e3b5332183b9"></a>

## Next pages — ingress_egress_gw_ar.global_network_list / c4e9529352c9 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0daad39be18ddb4cf9de3171317b190814b85ac2f5b4b0cc30bcfa883787ba19"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections — ingress_egress_gw_ar.global_network_list.global_network_connections / c88d8dba1531 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183)
- ingress_egress_gw_ar.global_network_list.global_network_connections

<a id="canonical-b2eb423d465e0268768950b7683d37433d8b538ff0cce32bf0576d06c7eeb1db"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-13b583d29405a2372410bf0e19c4d0feae2eeb6d8d5f44a31867c509f260222d"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections / c88d8dba1531 / 3

- [sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-65812226c52654d588de2f9d0e51bc20914eda0e4f05dcfd076b1a6f520a5494): complete subsection reference.

- [slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-8e7ae540d88e737dccabeb7ca18234978223ba59d7c399a485d84f319d1b314c): complete subsection reference.

<a id="canonical-ac34f1202ca594c2298be4bdf992b22ba57595644cd4b960d75d567237b7e09d"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections / c88d8dba1531 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-65812226c52654d588de2f9d0e51bc20914eda0e4f05dcfd076b1a6f520a5494)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-8e7ae540d88e737dccabeb7ca18234978223ba59d7c399a485d84f319d1b314c)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-65812226c52654d588de2f9d0e51bc20914eda0e4f05dcfd076b1a6f520a5494"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-735c68fe9705ec27ce44ced747eb33bf43535ecd6d2ea5a7f4f50128c4f789aa"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 9294792ff8fc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3)
- ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-bd85ed32e91cafa0d2c43a7e0ef5bbdb8980f6a602b8d9a100816039755c733d"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f482fcff04cac2bbe87a75b13003f016f5a56504b19d78423387283e9368f053"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 9294792ff8fc / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-f62dd13147db2a86641b14b9fdf65d236e28f98745b37957eac8e653bb01d485): complete subsection reference.

<a id="canonical-befd528329b8ed23933fe2092b43814caa8ff59bcd1a38d53928345d1143af70"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 9294792ff8fc / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-f62dd13147db2a86641b14b9fdf65d236e28f98745b37957eac8e653bb01d485)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f62dd13147db2a86641b14b9fdf65d236e28f98745b37957eac8e653bb01d485"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a6e1115b00d1348bac201f144c4cfd9f1a3b4d477b028fafb2e8e35605b1b06"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / e88480344033 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-65812226c52654d588de2f9d0e51bc20914eda0e4f05dcfd076b1a6f520a5494)
- ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-ed7512a4e961f488bbb4d2ea97c42b1b3c2b24e0fc3dd34dd01993902e2f19aa"></a>

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

<a id="canonical-747b41ce4d292a64d3b35ede31a7e0566d7145fcaed96d695de7039a09347130"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / e88480344033 / 3

<a id="canonical-2ab6d26e6509946ae98f326a022355d0ad479e633e26bb802a291dc6eb944152"></a>

<a id="canonical-040652b313a48397adf28f11840a3ea69a0cb955535daca41c6eed3c7a367d8d"></a>

## name property — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / e88480344033 / 4

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

<a id="canonical-295158a533fdd5e2e57fe0bdfd8e481bdd6c966841edac63055cf70adf214ca0"></a>

<a id="canonical-0089117f237fe5beae632e091c8131898a1e8bb85e480043a43ac6a275b67797"></a>

## namespace property — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / e88480344033 / 5

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

<a id="canonical-c6994e8b3a66051729f1313171903ede367d1bd19a7a03e43da36744eb91a4c3"></a>

<a id="canonical-adf435fcf374abda70d1928eb6faa094a54e373ca7dbb5a641d8bb3b2544f371"></a>

## tenant property — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / e88480344033 / 6

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

<a id="canonical-bd896db1da69dfb06374baedc48749b4ee0e913d20c5e7eea2be88c8fbc2d75a"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / e88480344033 / 7

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-65812226c52654d588de2f9d0e51bc20914eda0e4f05dcfd076b1a6f520a5494)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8e7ae540d88e737dccabeb7ca18234978223ba59d7c399a485d84f319d1b314c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bfff4fa6686bbbc66cc61d89d4fab2766d2ca1d43e8ea14ce459be0f3433b4d"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / c5d1fc179073 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3)
- ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1bf5bb7f26d472c14548bcd2cdee730fabd10fc16aad25e852f84005b61c0f97"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-386b5fc3ad1d72c9da9a713ac8305acce66ecbe1bce048794174bbf66196a434"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / c5d1fc179073 / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-d02232984db48d2e3514f3aae5c522269f78093f7923ba44a6c2f8122ebdcbbd): complete subsection reference.

<a id="canonical-c40d9eafeaf14b1f3e4b9d703387f21c3b5aee6ff66939d457f864b515241cfb"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / c5d1fc179073 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-005.md#canonical-d02232984db48d2e3514f3aae5c522269f78093f7923ba44a6c2f8122ebdcbbd)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d02232984db48d2e3514f3aae5c522269f78093f7923ba44a6c2f8122ebdcbbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b00bf2adba480af7fc744c0304a74bbd7eaa2e444bdb4a708c500d4803262198"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / d6518a7bcc06 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.global_network_list](data-sources--azure_vnet_site--reference--group-005.md#canonical-54c6fd28d68cea00e61a827218202127917cdf36162e3c1b818524aa05465183)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-f526cbe300e05adf5442220dd912adf909ec618b0695b29fcf8caa7b8630d7e3)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-8e7ae540d88e737dccabeb7ca18234978223ba59d7c399a485d84f319d1b314c)
- ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-8a07ce33aa55e050409c9d5cf1ce0fcc73a68aac1ca2c60fe4acd60f3406fc19"></a>

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

<a id="canonical-1998b6f80baf6b43c430b04daee623d0215865ec4e0bde5e0e95e3fa23fc315b"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / d6518a7bcc06 / 3

<a id="canonical-7ce24d36209866f1768dbae6b179a8bcf2aeccd0450a718fc896f346454eaab7"></a>

<a id="canonical-cc452b0937540ea074eeb4b9c66b3626f8c0b453dcd17255659e1f0bacba829b"></a>

## name property — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / d6518a7bcc06 / 4

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

<a id="canonical-10254a28d4b8da9737557ac6ea43b51ebccaaa57fe51e126ca79520fb13361ae"></a>

<a id="canonical-e73df3c72f37ee4aec1889a064dada2d8ba952467a82de3b2854324a8f87f4e0"></a>

## namespace property — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / d6518a7bcc06 / 5

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

<a id="canonical-3ddba56e09f741bfe527d94cabdc85d8e3f24355a8d4deb71514b92c8dd09de1"></a>

<a id="canonical-28973c50faab9c9c91c08f10311bd14e31b39213219eff67e04814cdd8845802"></a>

## tenant property — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / d6518a7bcc06 / 6

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

<a id="canonical-03ff1f5e481ef6eaea4e944474240155f2cf76b522823700417864e11810d7ab"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / d6518a7bcc06 / 7

- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-005.md#canonical-8e7ae540d88e737dccabeb7ca18234978223ba59d7c399a485d84f319d1b314c)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d33f1bf276fb7351882f676dfe894afc7c6e391a967d72c01ab541b84c20e7ac"></a>

## ingress_egress_gw_ar.hub — ingress_egress_gw_ar.hub / 4015425c8072 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- ingress_egress_gw_ar.hub

<a id="canonical-f108a786966d5352b3645f6c9845d9c15a656d2b7ee4e827466da8414ce8cd61"></a>

Type: `"single"`. Computed.

Hub VNet type. Hub VNet type.

Upstream description:

Hub VNet type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-express_route_choice": "[\"express_route_disabled\",\"express_route_enabled\"]"
}
```

<a id="canonical-9857da4b7ecf35f5ce27b0b5ef2a51594f67002875d369cc93cb64cca1ab87e1"></a>

## Direct properties — ingress_egress_gw_ar.hub / 4015425c8072 / 3

- [express_route_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-05827c68a63a7173a15192458c4862b441918f316acb6b3cbdc791ff8cd94a08): complete subsection reference.

- [express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675): complete subsection reference.

- [spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85): complete subsection reference.

<a id="canonical-b37a602378b49e5ac66e45612dc8b3cfa634851818e27886fdd4749d555a6957"></a>

## Next pages — ingress_egress_gw_ar.hub / 4015425c8072 / 4

- [ingress_egress_gw_ar.hub.express_route_disabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-05827c68a63a7173a15192458c4862b441918f316acb6b3cbdc791ff8cd94a08)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-7981ce90f70a82a13fe90fb257dea70672c33353817df24e4fef3e41afd38a85)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-05827c68a63a7173a15192458c4862b441918f316acb6b3cbdc791ff8cd94a08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d20b97a051ecf1d06e6f514f5b0a32c8dacfaea34109847ee46284d20df21380"></a>

## ingress_egress_gw_ar.hub.express_route_disabled — ingress_egress_gw_ar.hub.express_route_disabled / cb2c0d1c1eb4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- ingress_egress_gw_ar.hub.express_route_disabled

<a id="canonical-b3716a283349e41ac14a93fa154a71f0b2a5281d7c869047c4cb1e35b927530c"></a>

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

<a id="canonical-f76b99f74b1841459488e9e8bd0c2aeab2e57416d02e4f7d74bdfacf5f6bbab6"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_disabled / cb2c0d1c1eb4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eee058a767484a34ef3c351aba9ca6519e6dde1354f680e598032a6305d64294"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_disabled / cb2c0d1c1eb4 / 4

- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e17623e16df41a62adb945fda6281ca74b41a1774b076f300bf7d35996b7e3e"></a>

## ingress_egress_gw_ar.hub.express_route_enabled — ingress_egress_gw_ar.hub.express_route_enabled / eb9e69225eb5 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- ingress_egress_gw_ar.hub.express_route_enabled

<a id="canonical-f7c3f60bc97d4b1a4c3ee48c105567e0b95eea556695f119319fb8b36200db42"></a>

Type: `"single"`. Computed.

Express Route Configuration. Express Route Configuration.

Upstream description:

Express Route Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_express_route\",\"site_registration_over_internet\"]",
  "x-ves-oneof-field-sku_choice": "[\"sku_ergw1az\",\"sku_ergw2az\",\"sku_high_perf\",\"sku_standard\"]",
  "x-ves-oneof-field-spoke_vnet_routes": "[\"advertise_to_route_server\",\"do_not_advertise_to_route_server\"]"
}
```

<a id="canonical-e0dff8258a8214ca4028e8f34d286c553fc99b9401dea1ac56103d76dbada040"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled / eb9e69225eb5 / 3

- [advertise_to_route_server](data-sources--azure_vnet_site--reference--group-005.md#canonical-58145a4012c1d43bac59e3bd859b95cb43957c52ddd68e095e09bdf35d6cdc75): complete subsection reference.

- [auto_asn](data-sources--azure_vnet_site--reference--group-005.md#canonical-ee8d7e7df307d2b51c1b985a0c9250ef4d00be87b4fc2f6983e26c4df154adf8): complete subsection reference.

- [connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb): complete subsection reference.

<a id="canonical-a4e8055218eda104bdf47bff36279f7e3c155fde5de47912b5a54adc98153123"></a>

<a id="canonical-fd37dbc5d253cd46fa9d4af6493baab45e19c9d917ba340d9e91ad36891fc243"></a>

## custom_asn property — ingress_egress_gw_ar.hub.express_route_enabled / eb9e69225eb5 / 4

Type: `"number"`. Computed.

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Upstream description:

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

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
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  }
}
```

- [do_not_advertise_to_route_server](data-sources--azure_vnet_site--reference--group-006.md#canonical-8b3fd92aacfcb186818fbd2890ea4b7fa6fd03bd4af97c65aaa065c4dbbbacd5): complete subsection reference.

- [gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684): complete subsection reference.

- [route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862): complete subsection reference.

- [site_registration_over_express_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-299278ea065d5b3db3d839922a4b47db34f8cd9b753d64361f5a360120c8d2e9): complete subsection reference.

- [site_registration_over_internet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fa5a94ca7a39b2d36fc2759b911feb90e360b9105d49563df9d195b3c8b998fe): complete subsection reference.

- [sku_ergw1az](data-sources--azure_vnet_site--reference--group-006.md#canonical-c5430ea326d3a0e63f9d8011d6bc6a8aa46668371d4c21e9817c8d4947740b2e): complete subsection reference.

- [sku_ergw2az](data-sources--azure_vnet_site--reference--group-006.md#canonical-aad40f15c0a8bfb1b5fd48281b49ceec4dd12e554d021a431c09c5c570bb97fa): complete subsection reference.

- [sku_high_perf](data-sources--azure_vnet_site--reference--group-006.md#canonical-259b6cdc06ea3436858053d9c540e2da17f985c597c8a108a157046f71159958): complete subsection reference.

- [sku_standard](data-sources--azure_vnet_site--reference--group-006.md#canonical-647fd346ef093b5ddf530586a56202953953f81023603e85f6041ed7b47ab228): complete subsection reference.

<a id="canonical-49267e1138d761eef62e33fa73f28087f9ccb2f955c2bf54ee67bf09c7a15086"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled / eb9e69225eb5 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server](data-sources--azure_vnet_site--reference--group-005.md#canonical-58145a4012c1d43bac59e3bd859b95cb43957c52ddd68e095e09bdf35d6cdc75)
- [ingress_egress_gw_ar.hub.express_route_enabled.auto_asn](data-sources--azure_vnet_site--reference--group-005.md#canonical-ee8d7e7df307d2b51c1b985a0c9250ef4d00be87b4fc2f6983e26c4df154adf8)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- [ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server](data-sources--azure_vnet_site--reference--group-006.md#canonical-8b3fd92aacfcb186818fbd2890ea4b7fa6fd03bd4af97c65aaa065c4dbbbacd5)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2e6e73f8071e94139791f37a1ff2701840b8a1e164b8aef1929579cf0e701684)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-d5a902fd21a77e5ef701cc29f8a337368c9081de008a21d43003df2dd8e2c862)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-299278ea065d5b3db3d839922a4b47db34f8cd9b753d64361f5a360120c8d2e9)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet](data-sources--azure_vnet_site--reference--group-006.md#canonical-fa5a94ca7a39b2d36fc2759b911feb90e360b9105d49563df9d195b3c8b998fe)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az](data-sources--azure_vnet_site--reference--group-006.md#canonical-c5430ea326d3a0e63f9d8011d6bc6a8aa46668371d4c21e9817c8d4947740b2e)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az](data-sources--azure_vnet_site--reference--group-006.md#canonical-aad40f15c0a8bfb1b5fd48281b49ceec4dd12e554d021a431c09c5c570bb97fa)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf](data-sources--azure_vnet_site--reference--group-006.md#canonical-259b6cdc06ea3436858053d9c540e2da17f985c597c8a108a157046f71159958)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_standard](data-sources--azure_vnet_site--reference--group-006.md#canonical-647fd346ef093b5ddf530586a56202953953f81023603e85f6041ed7b47ab228)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-58145a4012c1d43bac59e3bd859b95cb43957c52ddd68e095e09bdf35d6cdc75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0fbb29ebfb2d28c375af75dae76d3983e4c79188f8ce6bb628005c902db6745"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server — ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server / 237c6db0d58f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server

<a id="canonical-abd9dad8493125f6e3364eb5be1d81121855326604994f588c771200feaa33e0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise to route server.

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

<a id="canonical-51a424cf51a5f26da943c6f44804e4a565ba95b788737c046d8fb3ec5efd7acd"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server / 237c6db0d58f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07cd26c9b3882e97c29ca1976bc49e1466a88648915db28173dfb237d089e534"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server / 237c6db0d58f / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ee8d7e7df307d2b51c1b985a0c9250ef4d00be87b4fc2f6983e26c4df154adf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aeb0c2964e10c8f181beada20b71a4c0eee365e952255c38f7f1ecdab0925a2"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.auto_asn — ingress_egress_gw_ar.hub.express_route_enabled.auto_asn / 173180990095 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.auto_asn

<a id="canonical-b1e5333030b4d2d5936403d6fecb344dc098f17639f27e769467781cb67ed26c"></a>

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

<a id="canonical-fa70e7ef55aa42561aa234367c515abe4e72c888e3c0e4f2d834248d2b7ccc17"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.auto_asn / 173180990095 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c10ae3c7caa735ceb07e5ab7ecc6cb995001f32cf9989bdc3dfb84b4bc745b1"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.auto_asn / 173180990095 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aa0294399665a967f35a96c11ac5202098c00c6ee1cb34fc616527ea4009d17"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections — ingress_egress_gw_ar.hub.express_route_enabled.connections / 6ee562502313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- ingress_egress_gw_ar.hub.express_route_enabled.connections

<a id="canonical-9a0755e369c5c33e549403a521553175635d35231e3b67f692368607ba8a818a"></a>

Type: `"list"`. Computed.

Add the ExpressRoute Circuit Connections to this site.

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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-d7ca79adefc124c128b432d4c3e337b540b82b6c688183955716f003ff5b8e1a"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections / 6ee562502313 / 3

<a id="canonical-7c51111ecd634c1bf63dbeb5f5988dd5e0d6e8a966165714834f03afe4ce0205"></a>

<a id="canonical-a6347b250ae59b293184a474b8ec504afe84bfbfe0f9459868a758d1d57d545c"></a>

## circuit_id property — ingress_egress_gw_ar.hub.express_route_enabled.connections / 6ee562502313 / 4

Type: `"string"`. Computed.

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

Upstream description:

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

- [metadata](data-sources--azure_vnet_site--reference--group-005.md#canonical-d1e4bd9e3a17f18aeb519070b42177ac3fc0356b71da3605b230b6f22946d86a): complete subsection reference.

- [other_subscription](data-sources--azure_vnet_site--reference--group-005.md#canonical-51b1ff4f98ef8a3b44521b61c6a0f2c6841ca187ea397e3a6a583e7fba8fe100): complete subsection reference.

<a id="canonical-d6b2bad2b546cd218e461ce8d13175bd669112bbf6bca6dadea235be0c3e1ba2"></a>

<a id="canonical-c539c8f971d312e68c7e940cede426369f409915d3538265b5847600feb688b2"></a>

## weight property — ingress_egress_gw_ar.hub.express_route_enabled.connections / 6ee562502313 / 5

Type: `"number"`. Computed.

The weight (or priority) for the routes received from this connection. The. Defaults to \`10\`.

Upstream description:

The weight (or priority) for the routes received from this connection. The default value is 10.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

<a id="canonical-c676b1c11313378145f233e8a55f5692f158c094aad7085e67589383ed1c4625"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections / 6ee562502313 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata](data-sources--azure_vnet_site--reference--group-005.md#canonical-d1e4bd9e3a17f18aeb519070b42177ac3fc0356b71da3605b230b6f22946d86a)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-005.md#canonical-51b1ff4f98ef8a3b44521b61c6a0f2c6841ca187ea397e3a6a583e7fba8fe100)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d1e4bd9e3a17f18aeb519070b42177ac3fc0356b71da3605b230b6f22946d86a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81c6590b1a57c1371fe23a7410195afd443ece1779e19013b3300f3bd2894133"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / e11b18eb1943 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata

<a id="canonical-bc31479a1f97e46c8c28370427e1f353d72ee1d5e546e28faff1f8b59de1ec7b"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-26895d789eba895175799b9f44244fcc3b8863e6dcd308c73adba2cb1e6805ef"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / e11b18eb1943 / 3

<a id="canonical-e1c52d0a2e74ec71f3c289f01449ed3b71d724e3fb2dbc1912f1c03eb11f753c"></a>

<a id="canonical-8f16b9c8efd6bd90069aaf84c9672ff357fbb8d6bbe93c1225787ba792cbf854"></a>

## description_spec property — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / e11b18eb1943 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-d24e5103bb6092034e71f1d28c8f3853387eeab12a7573349f72330316dc8fc5"></a>

<a id="canonical-e9b0d322e1f37bbfb65552566a779ac9e05572b0e12388d43e9ffd5e095112f8"></a>

## name property — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / e11b18eb1943 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0e34879218e0f6242c788168fb394ad311f2a6b711facd66d76fa99d44dc5e9c"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata / e11b18eb1943 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-51b1ff4f98ef8a3b44521b61c6a0f2c6841ca187ea397e3a6a583e7fba8fe100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8570f8463457fe2c1fa9d29639fa437ac5a31c3eb5dc978817f43a2cb537993f"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / 1aefc95c75f0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription

<a id="canonical-95be07977fb32b0aaeb5fe766a667d3b99e99557c004ec933ec9aa19adaa487c"></a>

Type: `"single"`. Computed.

Express Route Circuit Config From Other Subscription.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2ed388b6aebe984b29afefcef1a0c001ccd379b2e7fbddf173817901c92ea1fc"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / 1aefc95c75f0 / 3

- [authorized_key](data-sources--azure_vnet_site--reference--group-005.md#canonical-df2dbf4f80a90c6182064e66dceaaeb134c52ad6442e0065f745c89e7a250559): complete subsection reference.

<a id="canonical-2c5cf1763d4c7988643d45cb7a803d693c8b2149ff2edd4b3aba6fd741b96524"></a>

<a id="canonical-bc18154fe1a3e7127f8d6bbf3c5a151eeb3aa55d6f657aa7a33f14f3339b3d65"></a>

## circuit_id property — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / 1aefc95c75f0 / 4

Type: `"string"`. Computed.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-bc5ad5078c20aad4dd68e2d6d2d50b6db767f33d682abfbffc8376b31b09d82b"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription / 1aefc95c75f0 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-005.md#canonical-df2dbf4f80a90c6182064e66dceaaeb134c52ad6442e0065f745c89e7a250559)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-df2dbf4f80a90c6182064e66dceaaeb134c52ad6442e0065f745c89e7a250559"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e9ff4d7246642ebcb335ec27cc69e9ef1a6e1b22f8adfd7eba63f3ec06e8684"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 3905982e8e7d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-b163f0d783eafafeca1d0adc24b82ee9fc49f58ffe6625aff98d034258949ed5)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-49de7737a02c1c4c5e112eb01ede09af9877a10ecb21c13da95c573cae2af7c6)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-da86de1bea8911eb9d7f0c9f2cfff8345b7a717714f16e8c4bf74c9cac38f675)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-6ce8e5b6f4d03779cec0058e7cbac6dded02139d92f1cd3a16254724c7654fbb)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-005.md#canonical-51b1ff4f98ef8a3b44521b61c6a0f2c6841ca187ea397e3a6a583e7fba8fe100)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="canonical-f1abe8226560504fa5c803312be049e8794053a2da8a19344fb198ea1963be2d"></a>

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

<a id="canonical-febbee71862bc98e5bd6e883ca7fe539f439ba18945c0cd78eccf2c04a729b95"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 3905982e8e7d / 3

- [blindfold_secret_info](data-sources--azure_vnet_site--reference--group-005.md#canonical-b87776148b78b41662f99388c3748fe81b4d0e0fafa941b7291326f7600618c0): complete subsection reference.

- [clear_secret_info](data-sources--azure_vnet_site--reference--group-006.md#canonical-69777777604c41afdbd849d3c55ee081f2a6c4069540639264a034cb44904a06): complete subsection reference.

<a id="canonical-9ed16b686b4d0fcecfe45857731eb0b837c0e3537e04fc80616769755b0f9abf"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.au / 3905982e8e7d / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](data-sources--azure_vnet_site--reference--group-005.md#canonical-b87776148b78b41662f99388c3748fe81b4d0e0fafa941b7291326f7600618c0)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](data-sources--azure_vnet_site--reference--group-006.md#canonical-69777777604c41afdbd849d3c55ee081f2a6c4069540639264a034cb44904a06)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-005.md#canonical-51b1ff4f98ef8a3b44521b61c6a0f2c6841ca187ea397e3a6a583e7fba8fe100)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b87776148b78b41662f99388c3748fe81b4d0e0fafa941b7291326f7600618c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
