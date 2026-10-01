---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-384cb2763791a7df8fa0232fc2d5853fc6a823ad4330f50a192ecec302ea0fd5"></a>

## tenant property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 0b3e73b07767 / 7

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

<a id="canonical-3ff2647bb173d09f9492defce9be6906172852f235a56a60c0ce20c9245b459d"></a>

<a id="canonical-17343d553f6267e35db5e7ae7e696d20f9d169bf2fae6ca370566309567ba161"></a>

## uid property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 0b3e73b07767 / 8

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

<a id="canonical-3ad409ceb1e5f4f02de6b3826adc8650b73f4dda1cda95640efd82befb90c5a4"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 0b3e73b07767 / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-957217b81a62ceb3d6cbe42b389e80863eb7f3df9c5090e907490822fb6d7e1e"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 2830870f997a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-4c69b0ac08c70f2d80be9ab54f43c51203e5ec08755a1ac32e2d929c3233b53d"></a>

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

<a id="canonical-8c1b1a1d703162a1fef7656c745c29b2ec3b0bb72bef9c3433a39eee51117a41"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 2830870f997a / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-77f569d933791cef204e9a23236bb75b36954b16a2e1449a3e683bbfd9d64fb4): complete subsection reference.

- [ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-db56320f2cc105f2d65706c7628318b00884fd748eb136b30fdf8e14be0f3de7): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-8087fb6c2bf1eace9a3a9d44e09736491f666cf93ecc282053f1168258666d80): complete subsection reference.

<a id="canonical-ceef77aa08eef766a21cf7c77486d624b1b9dab37a790d26428a532a83b4f196"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 2830870f997a / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-77f569d933791cef204e9a23236bb75b36954b16a2e1449a3e683bbfd9d64fb4)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-db56320f2cc105f2d65706c7628318b00884fd748eb136b30fdf8e14be0f3de7)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-8087fb6c2bf1eace9a3a9d44e09736491f666cf93ecc282053f1168258666d80)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-77f569d933791cef204e9a23236bb75b36954b16a2e1449a3e683bbfd9d64fb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8644da47fd986fcb6e6705984b90ef112025ce103a2ed6e61600434b6846cc11"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 612e91bab0a0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-f9d679011387c01be5436cf9ae06a7bf1e8aeaa5a9cf7896d99bc8a21ee00b98"></a>

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

<a id="canonical-04a43f64a1760271cc526cf2ee673574dd7727a5c82d0c39494cca171b792853"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 612e91bab0a0 / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-749c6a448a481e9187b7a5df2f7e5fec3591955ef853a0d1c859553e942e9eb9): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9c2591f892f709461bb5f9b3cf458a152e6ee7f4a5a8b0d1872ef0faca9f995): complete subsection reference.

<a id="canonical-7260b458e5ccdd0b3c04944429b4b118d4ee5943774cc7bbb94649e70e05f223"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 612e91bab0a0 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-749c6a448a481e9187b7a5df2f7e5fec3591955ef853a0d1c859553e942e9eb9)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9c2591f892f709461bb5f9b3cf458a152e6ee7f4a5a8b0d1872ef0faca9f995)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-749c6a448a481e9187b7a5df2f7e5fec3591955ef853a0d1c859553e942e9eb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eea424243643b6582f21980a156562ce26093da49b2da192bd9fd85ee8b3fc31"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 7c1a065d30a6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-77f569d933791cef204e9a23236bb75b36954b16a2e1449a3e683bbfd9d64fb4)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-53972ec3fd71bba30564ab090ab4270ade79152d15f4015a0e485122a4811c88"></a>

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

<a id="canonical-aaa441627a547d5869ebe84bf8b78149b1ab51159b759722fe6477f2785ba23c"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 7c1a065d30a6 / 3

<a id="canonical-f2a836b087af1d8b7fb5f258d2576cb1bed230e92c1ba012896923eb552c0e04"></a>

<a id="canonical-527bd0bc766e4c8c100c589ed3b9cc2d5cd1acb4e0656f027fd57fb1e3f28066"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 7c1a065d30a6 / 4

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

<a id="canonical-a20c8a5a19239616b23ce38a3a367a3f9048c035f21d238a02fb9e8f94a65ca5"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 7c1a065d30a6 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-77f569d933791cef204e9a23236bb75b36954b16a2e1449a3e683bbfd9d64fb4)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a9c2591f892f709461bb5f9b3cf458a152e6ee7f4a5a8b0d1872ef0faca9f995"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d17f8361d4a22077b257e7abd2e5f9704b096b7f1be2d0f5fe8e46c6ba082c13"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5bec14595b51 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-77f569d933791cef204e9a23236bb75b36954b16a2e1449a3e683bbfd9d64fb4)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-c84b34e371509a8cdf4a9fa1ca5de31c6795bcb3dbfb37c8a6943249bc53bdf6"></a>

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

<a id="canonical-b8a0b39c4ab57fc63b41132dc819e2a79a30214deecd758b85efd81e6d9eda40"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5bec14595b51 / 3

<a id="canonical-95ab0d7d7dab2fc6764d287a3bd89c0d5a2d3cb7627daa13dd413bd0e555c542"></a>

<a id="canonical-13e3e9ca30bda533fd698b7cb1b85357352f45d604133e745a16a510c8543a8c"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5bec14595b51 / 4

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

<a id="canonical-ab972b0f7aac5a729620123f846adb1b46875b026e1b4eb965e993a0c679b82a"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5bec14595b51 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-77f569d933791cef204e9a23236bb75b36954b16a2e1449a3e683bbfd9d64fb4)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-db56320f2cc105f2d65706c7628318b00884fd748eb136b30fdf8e14be0f3de7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-443a0ccd0a4feba64daf23e5b5015d3e0c692df022b3c995226f041f053cc1e1"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 87c97b5a0d47 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-00a0c535613d380dd86d38cc6c00e4ce64f220476c7380d9608fa10151921097"></a>

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

<a id="canonical-66169719a5821bc69becfe3e7bd55837358f89562883aa793cd272fd5086d377"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 87c97b5a0d47 / 3

<a id="canonical-72b9d6a6054aa4720022dc8eed0c0284e7a3212cf44cddc5aba65fd4bfcee3a7"></a>

<a id="canonical-549267b05ff8b521aefa61adbf7212f652af33f7998730fe361015be57312aab"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 87c97b5a0d47 / 4

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

<a id="canonical-c679edc5fb1f6664f3258a45a8e4e589e6c02dd3ccd0a2a7e0a6ace7b6c43423"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 87c97b5a0d47 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8087fb6c2bf1eace9a3a9d44e09736491f666cf93ecc282053f1168258666d80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dbf7af3a6777fc86bc57e70b7bc5997335223208ab6e24f64a1aa993e6cdf36"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / d7ab0aee62d0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-008.md#canonical-f5ad141b164198015aea65efe1f6e349f3d4a86213cd9fa57c060c2368a6e3f8)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-fc887e6abbfd59af684fd9ad4f6276221edc5e136cb70fc24954c1c7996d43cf"></a>

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

<a id="canonical-588fdf5d6f760ac45a319537defddd07025e012f24c9807deef0ce0a90684f10"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / d7ab0aee62d0 / 3

<a id="canonical-d43c2494e4bc6401cf82e6e66741798fa8a46968a2c3983410467e16c851da2f"></a>

<a id="canonical-3a0f5b498fbe2d51cd9578583b8a7face770a12768c85ef2ce640e07a409e22f"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / d7ab0aee62d0 / 4

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

<a id="canonical-cc9256765390a42688275979a48cdf3cf04f7688b7c0d95646fedd43e6c69618"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / d7ab0aee62d0 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-ba31b2ef817407bc1ea1a936ead6a883a217f78459db506495d0eaaf259df022)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-972c85e268533866b0cfe516b10018adfb68384a292003266b47f8728225a99d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e6d6ea7f9f0047623ae5145f0df8372205e4d2a5de2a131c82d505157500001"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 4b5dd68cfb21 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-e3c1f5ac100b2cfd20f1ce7d694633a84dd4491135bee2c3365de606be3371b6"></a>

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

<a id="canonical-3176c44b886cc3e40afff9cd4ede00fb18c333e5c23bdec18a3abe4733632d2f"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 4b5dd68cfb21 / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-a26c577caaba6e38436fe1203f03d31c177756ec723c1f177fe6e0966d702d09): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-b26bdf3fb15d950e1e0bb7336319aaf934b5afa002d637e8f7b969261908738b): complete subsection reference.

<a id="canonical-ff17fab1c8ddeedb25f0a7a0c2cb77fc871c02f5c77fb5ff0cd80a2cca623ff5"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 4b5dd68cfb21 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-a26c577caaba6e38436fe1203f03d31c177756ec723c1f177fe6e0966d702d09)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-b26bdf3fb15d950e1e0bb7336319aaf934b5afa002d637e8f7b969261908738b)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a26c577caaba6e38436fe1203f03d31c177756ec723c1f177fe6e0966d702d09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfdb22fce9db9803a1333cc31e50c9bbfb66b451caf3daefafab2c44afce75a0"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 9cb5bb8f6e2e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-972c85e268533866b0cfe516b10018adfb68384a292003266b47f8728225a99d)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-55b0a22ef579e6eabacdfa6209cf8ef2bb3eee044cee4ccff06b7ffb8c5a115e"></a>

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

<a id="canonical-461b8928a6463e5eac5d0c78d8e44e7e1770282bd9294ddb1c14ccc817225682"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 9cb5bb8f6e2e / 3

<a id="canonical-20fb051665aa39723926eb89aea1e464544ec70d18354c6ef8c245978df94eca"></a>

<a id="canonical-4cc0e4d28c5deadd5228e21919ab2001cbccf951dd1e351e24491f2426f1e052"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 9cb5bb8f6e2e / 4

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

<a id="canonical-83a5dd6efc5b8bf53496b3b8c168e96f8e43d23feffa4f0e3591aa27b0634e21"></a>

<a id="canonical-e4addc3cbe9659c1a75c4c39d2054256ba147278b0dfd70de5d7dc8c14e1d3d8"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 9cb5bb8f6e2e / 5

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

<a id="canonical-7a5b9075ec0294ba00473306e99deffb83c89fae8f251abd1489254f61371a46"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 9cb5bb8f6e2e / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-972c85e268533866b0cfe516b10018adfb68384a292003266b47f8728225a99d)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b26bdf3fb15d950e1e0bb7336319aaf934b5afa002d637e8f7b969261908738b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-280bb5cca363ba0c5d4aa69f58be7cbf0734ea19f6658cf513dfee10fb4d9d34"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6398d51a1ff4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-008.md#canonical-9a131ec7031d937d57edb9c4a8f150e1689a75f40450bf738aeaf45f79beb67c)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-008.md#canonical-c3540407fb016d6d0b0fc52b901254300d36c39036064f2ddab11512c2124a5c)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-008.md#canonical-7f1662b182afb00e460341659f9690b3b3cfb8ae950d9a317c20d74e8ac00394)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-972c85e268533866b0cfe516b10018adfb68384a292003266b47f8728225a99d)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-7a700ff4e68f59765b705469c54e1228ac7139fc4454dfafdb0946d6b80819a0"></a>

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

<a id="canonical-16fe56f6dffe1e0d346ec41a2d283d18dfacebd292e3d5d23a4407912769023d"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6398d51a1ff4 / 3

<a id="canonical-e20a3f0a7d7b00d1079841e3733179ac2f90e12a19be3862988f89ca4aff6506"></a>

<a id="canonical-e1a7c262206a27b5f8d1db34323d60fbd812969bb482b453bfa530b0998e4b3b"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6398d51a1ff4 / 4

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

<a id="canonical-a21ab72de3d2705c94ca98398ceabb3bdf17bf6ba128985de41312efc9995256"></a>

<a id="canonical-feb1812febae78f11ddafff7dc9aa9dcd5ccd487a6978cc279a5af37c0993e3b"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6398d51a1ff4 / 5

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

<a id="canonical-2905340c6f30428998cd526aad43cf1dfa4c6288f237eab8688d0cc7f854f9b0"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6398d51a1ff4 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-972c85e268533866b0cfe516b10018adfb68384a292003266b47f8728225a99d)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-93d9875c35a530f18e18adcdca587ab63e1a3e1061c59582ca846c0cac42f35c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ab6d285af15a6c28c82fec81bb8da2860e1c3870d953ef46e24f4e227c387a6"></a>

## voltstack_cluster.sm_connection_public_ip — voltstack_cluster.sm_connection_public_ip / 8419d78cc25b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-0b8dc4371c488738b832b5280b5b6535359fb6182559ec3e94285c16460ab05b"></a>

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

<a id="canonical-19da86bd27f5e7a076971287677b4673f8f23c80236b9ad75dfdf88ce9e63ad2"></a>

## Direct properties — voltstack_cluster.sm_connection_public_ip / 8419d78cc25b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72fd6dba2f72db72aa2ac15685ade961c9e22627bc2c21516b516c0419e73109"></a>

## Next pages — voltstack_cluster.sm_connection_public_ip / 8419d78cc25b / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-253d150fb8696b5e940b7a9a1195a6ad393025683fd9f6520351385f5e3ba5fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fa0b15dafd350b2e404711939043e60349bbd3fb93e790bc1ff303c391e1bd7"></a>

## voltstack_cluster.sm_connection_pvt_ip — voltstack_cluster.sm_connection_pvt_ip / f9b9abd4b6b1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-dcca1330c1f77ae7f0a3e6ffcff4125e794c9cf6f7be2b65774ec49dac2f8267"></a>

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

<a id="canonical-b9bc1c7eb2c9dad9f9d2cc19e140f2d1b6ce6cff0335081e7b5a90a25177e194"></a>

## Direct properties — voltstack_cluster.sm_connection_pvt_ip / f9b9abd4b6b1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a57492ce44cc063484536e558f8d48fe1be46df37532b9dc262be5580d32e04b"></a>

## Next pages — voltstack_cluster.sm_connection_pvt_ip / f9b9abd4b6b1 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a5b14835ee7069d97522813844899d5351b81b3c100905df47eafa30c65ad0ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51f826683f6c7c2264f3c26986be02424bac8043e35d6153f28fefbab82f4367"></a>

## voltstack_cluster.storage_class_list — voltstack_cluster.storage_class_list / 28b6048f68cf / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- voltstack_cluster.storage_class_list

<a id="canonical-d1ff538f905429f397aae3354ec3046a46390592bc198647e78aeda907621d61"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-389c66f20480fbe9a32c06d4355d2927f1c51801f30912664ed3058943a8222a"></a>

## Direct properties — voltstack_cluster.storage_class_list / 28b6048f68cf / 3

- [storage_classes](data-sources--azure_vnet_site--reference--group-009.md#canonical-8f9365e1d3688cc9724241b967cf4ec1aa963b9d1d6a5284338001080638539c): complete subsection reference.

<a id="canonical-7a700ecaef6fa463f2c46061254a4f501497523268f18ed1fbf65ce91eb4c8db"></a>

## Next pages — voltstack_cluster.storage_class_list / 28b6048f68cf / 4

- [voltstack_cluster.storage_class_list.storage_classes](data-sources--azure_vnet_site--reference--group-009.md#canonical-8f9365e1d3688cc9724241b967cf4ec1aa963b9d1d6a5284338001080638539c)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8f9365e1d3688cc9724241b967cf4ec1aa963b9d1d6a5284338001080638539c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3809771a5b22e98c57a04abb33df405a757a9548d60bae4d3b340b2a4d16fd19"></a>

## voltstack_cluster.storage_class_list.storage_classes — voltstack_cluster.storage_class_list.storage_classes / dfd240b1d979 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0be67d2d7b5bbccba8043c7cc827083a7accdc2484c337bdd8b596fe11dde700)
- [voltstack_cluster.storage_class_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-a5b14835ee7069d97522813844899d5351b81b3c100905df47eafa30c65ad0ab)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-8c4b3b755b345d321e2af4cb7bba56a5992d56198a813be1b7df5eab9f9849f4"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

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

<a id="canonical-3d297bf7e005264e4c966f6fb5a7fd9c4a69cf7810f44c378ce0482ee03f47fc"></a>

## Direct properties — voltstack_cluster.storage_class_list.storage_classes / dfd240b1d979 / 3

<a id="canonical-8dfdf05de0aa921576a0e58795195c18ed4a297e0a2df28e4b2aa923c309faf1"></a>

<a id="canonical-ea87031063984fa288fbf84d9cef087f485afbb9dc1f8e1d01b284a7e175eb1b"></a>

## default_storage_class property — voltstack_cluster.storage_class_list.storage_classes / dfd240b1d979 / 4

Type: `"bool"`. Computed.

Make this storage class default storage class for the K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ba6ef9f5662a1b1ffe314dc6b7800f9c2d75fbd34a2452a5350ee8825d83a763"></a>

<a id="canonical-2b633ce089e8cc510518598839a4a12d180dcd3fef3ff8ae3675091bbc7b25c6"></a>

## storage_class_name property — voltstack_cluster.storage_class_list.storage_classes / dfd240b1d979 / 5

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-bc65e46c8ba4cc62d686e5f35a1e105b3d95e9779b9d33c67fce448b707d8ccf"></a>

## Next pages — voltstack_cluster.storage_class_list.storage_classes / dfd240b1d979 / 6

- [voltstack_cluster.storage_class_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-a5b14835ee7069d97522813844899d5351b81b3c100905df47eafa30c65ad0ab)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f20cb3c07349bc7dc4c69961371ea673515e49ba9b9fd040d889dd818c30447"></a>

## voltstack_cluster_ar — voltstack_cluster_ar / 73a945f7e896 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- voltstack_cluster_ar

<a id="canonical-5e3dd1f0dced227f34115de8514d94ae2790083d0eaacfb278e8494db94eb612"></a>

Type: `"single"`. Computed.

App Stack Cluster of single interface Azure nodes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

<a id="canonical-4adb19d01a6da5a0046b6e9c49510ab79e80bc89463d061ce45a4d5d5619fcfa"></a>

## Direct properties — voltstack_cluster_ar / 73a945f7e896 / 3

- [accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-d5e1b7a2152e82a99f5d910cbc8e1ba56495dc04ce00940534e56d90ee2c5f86): complete subsection reference.

- [active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-dcd77dd26bc732a2354140283470665b9df2466c8be936f97d1eb1c7ed643d0e): complete subsection reference.

- [active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-a5b131a73e65adc0e811dbcef693db09c0d537dd1eb99224b25965b578654981): complete subsection reference.

- [active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-95daf8f5047512b12e269a1d0b68b26ad4c49ca102decb338e0d3ba96cc557de): complete subsection reference.

<a id="canonical-94ff9b052b3686e08e3b370a301df00d8a06e086ca38ff3286d21e3b17274f97"></a>

<a id="canonical-b78d98b1a9107a5ac4a4cf3295f975a4f084d77aab730044e690ad6ac72d6fea"></a>

## azure_certified_hw property — voltstack_cluster_ar / 73a945f7e896 / 4

Type: `"string"`. Computed.

\[Enum: azure-byol-voltstack-combo\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-voltstack-combo\`.

Upstream description:

Name for Azure certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-b0bcd9866e7fd6642d4650a33600c474e83fdfe443c4bd586398c35480c56c3d): complete subsection reference.

- [default_storage](data-sources--azure_vnet_site--reference--group-009.md#canonical-c9511d506bb8a71655e35c8822643d3140a7682648d6f0c34213c754fa3e3096): complete subsection reference.

- [forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-009.md#canonical-433a56aa39be244c9e31dd32ae0f689659c25dac91beed57b2d5ec1d7fbda829): complete subsection reference.

- [global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263): complete subsection reference.

- [k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-a4732fd25019dd2da4fcda8a6f9224587fadbd5acffcefaee82e7aee0ffb03a9): complete subsection reference.

- [no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-ae6e0cdf9444360bf277b441872027f8c4cd60bb5018512e8e4b4a3100dc095b): complete subsection reference.

- [no_forward_proxy](data-sources--azure_vnet_site--reference--group-009.md#canonical-689e6c85553974c4f315e416838f05b7f707b51c6e5a05bc827e6b8d94f189e0): complete subsection reference.

- [no_global_network](data-sources--azure_vnet_site--reference--group-009.md#canonical-d3bdf49ae2bb1b8bd5c7797401199b07dd6bbabda699cb6cbb267be2972a4278): complete subsection reference.

- [no_k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-309ff2d1789f13ed9b981ff6f7a85ed57648688e6f1893bcb868d58b179d7cf0): complete subsection reference.

- [no_network_policy](data-sources--azure_vnet_site--reference--group-009.md#canonical-ddff7218ccd5d6e45cd856f00a9ed6ee04e3a03fb7004e3ff2744fb1b1d09df0): complete subsection reference.

- [no_outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-11c9b02e198b98e08ec66a9051e2ccb9f556948c90dcc9bf622ffc00a0a001e5): complete subsection reference.

- [node](data-sources--azure_vnet_site--reference--group-009.md#canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd): complete subsection reference.

- [outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a): complete subsection reference.

- [sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-54766ac35fe5d2e703fc587b4b548e36ccf4d460bfe1a3dd89fff3b9d8dac5f1): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-11b6c7ec3f181c658986f7a28da94476eae4b676b2dff84b12bf72c28798cf96): complete subsection reference.

- [storage_class_list](data-sources--azure_vnet_site--reference--group-010.md#canonical-6481f0ed982a057ed19611192cec83707d0e581d79647d711c22c851f1aa307f): complete subsection reference.

<a id="canonical-8e72548c8efc09350138a174bbcef0492697f49e4413cad4da24561ba1fd0b41"></a>

## Next pages — voltstack_cluster_ar / 73a945f7e896 / 5

- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-d5e1b7a2152e82a99f5d910cbc8e1ba56495dc04ce00940534e56d90ee2c5f86)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-dcd77dd26bc732a2354140283470665b9df2466c8be936f97d1eb1c7ed643d0e)
- [voltstack_cluster_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-a5b131a73e65adc0e811dbcef693db09c0d537dd1eb99224b25965b578654981)
- [voltstack_cluster_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-95daf8f5047512b12e269a1d0b68b26ad4c49ca102decb338e0d3ba96cc557de)
- [voltstack_cluster_ar.dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-b0bcd9866e7fd6642d4650a33600c474e83fdfe443c4bd586398c35480c56c3d)
- [voltstack_cluster_ar.default_storage](data-sources--azure_vnet_site--reference--group-009.md#canonical-c9511d506bb8a71655e35c8822643d3140a7682648d6f0c34213c754fa3e3096)
- [voltstack_cluster_ar.forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-009.md#canonical-433a56aa39be244c9e31dd32ae0f689659c25dac91beed57b2d5ec1d7fbda829)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263)
- [voltstack_cluster_ar.k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-a4732fd25019dd2da4fcda8a6f9224587fadbd5acffcefaee82e7aee0ffb03a9)
- [voltstack_cluster_ar.no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-ae6e0cdf9444360bf277b441872027f8c4cd60bb5018512e8e4b4a3100dc095b)
- [voltstack_cluster_ar.no_forward_proxy](data-sources--azure_vnet_site--reference--group-009.md#canonical-689e6c85553974c4f315e416838f05b7f707b51c6e5a05bc827e6b8d94f189e0)
- [voltstack_cluster_ar.no_global_network](data-sources--azure_vnet_site--reference--group-009.md#canonical-d3bdf49ae2bb1b8bd5c7797401199b07dd6bbabda699cb6cbb267be2972a4278)
- [voltstack_cluster_ar.no_k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-309ff2d1789f13ed9b981ff6f7a85ed57648688e6f1893bcb868d58b179d7cf0)
- [voltstack_cluster_ar.no_network_policy](data-sources--azure_vnet_site--reference--group-009.md#canonical-ddff7218ccd5d6e45cd856f00a9ed6ee04e3a03fb7004e3ff2744fb1b1d09df0)
- [voltstack_cluster_ar.no_outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-11c9b02e198b98e08ec66a9051e2ccb9f556948c90dcc9bf622ffc00a0a001e5)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--reference--group-009.md#canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- [voltstack_cluster_ar.sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-54766ac35fe5d2e703fc587b4b548e36ccf4d460bfe1a3dd89fff3b9d8dac5f1)
- [voltstack_cluster_ar.sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-11b6c7ec3f181c658986f7a28da94476eae4b676b2dff84b12bf72c28798cf96)
- [voltstack_cluster_ar.storage_class_list](data-sources--azure_vnet_site--reference--group-010.md#canonical-6481f0ed982a057ed19611192cec83707d0e581d79647d711c22c851f1aa307f)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d5e1b7a2152e82a99f5d910cbc8e1ba56495dc04ce00940534e56d90ee2c5f86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-740f1dcc522096be74fe617a28e0a94f134a2fdbcb06489c59376364a513bcad"></a>

## voltstack_cluster_ar.accelerated_networking — voltstack_cluster_ar.accelerated_networking / 0291bfaa5e29 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.accelerated_networking

<a id="canonical-b5f24dc967f7fc2686c9da49274d78327f092049a824d6b0dddab75f3c8da7fb"></a>

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

<a id="canonical-966ebd8a30ba7f136a7786ace6b05026328d4594a86b6cea470291c1556770e2"></a>

## Direct properties — voltstack_cluster_ar.accelerated_networking / 0291bfaa5e29 / 3

- [disable_spec](data-sources--azure_vnet_site--reference--group-009.md#canonical-24779564f5e30b253dc6389af82eb177a4b5c6cd46c05b4b8c8dc4d919197d0e): complete subsection reference.

- [enable](data-sources--azure_vnet_site--reference--group-009.md#canonical-4c9fa1dc38d4af82f6ec1e5d5ffee01e6adfd8f574251a9886cae777179045c5): complete subsection reference.

<a id="canonical-0ca13a0fbd8113ec97019d937a333c33b9344286c5403c0c89ef0c34a2b2c5a1"></a>

## Next pages — voltstack_cluster_ar.accelerated_networking / 0291bfaa5e29 / 4

- [voltstack_cluster_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--reference--group-009.md#canonical-24779564f5e30b253dc6389af82eb177a4b5c6cd46c05b4b8c8dc4d919197d0e)
- [voltstack_cluster_ar.accelerated_networking.enable](data-sources--azure_vnet_site--reference--group-009.md#canonical-4c9fa1dc38d4af82f6ec1e5d5ffee01e6adfd8f574251a9886cae777179045c5)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-24779564f5e30b253dc6389af82eb177a4b5c6cd46c05b4b8c8dc4d919197d0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c864c258d026ba3ee27526f0c79d562c97b96f2aeeebd21a44216a6aee3b636"></a>

## voltstack_cluster_ar.accelerated_networking.disable_spec — voltstack_cluster_ar.accelerated_networking.disable_spec / e4e13b1cf743 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-d5e1b7a2152e82a99f5d910cbc8e1ba56495dc04ce00940534e56d90ee2c5f86)
- voltstack_cluster_ar.accelerated_networking.disable_spec

<a id="canonical-dbb767e7c285ba5c003f3da90ff57b400871b3a55282be7f902e96fcddc388ad"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-24be29a1ddeb55a965bda1b9889efbe7b7bad7ddb9faf71882cbcad24f635140"></a>

## Direct properties — voltstack_cluster_ar.accelerated_networking.disable_spec / e4e13b1cf743 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd40f6cf700b804f0a0a9681a1a4f5f22d17fa607294b7b5fcea2989d1f4a77d"></a>

## Next pages — voltstack_cluster_ar.accelerated_networking.disable_spec / e4e13b1cf743 / 4

- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-d5e1b7a2152e82a99f5d910cbc8e1ba56495dc04ce00940534e56d90ee2c5f86)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-4c9fa1dc38d4af82f6ec1e5d5ffee01e6adfd8f574251a9886cae777179045c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0025caf94ab3d570dc801ffa499d1ec76fb69cb40de4113e1ad4799e2b12a69"></a>

## voltstack_cluster_ar.accelerated_networking.enable — voltstack_cluster_ar.accelerated_networking.enable / 6bf630ef4b29 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-d5e1b7a2152e82a99f5d910cbc8e1ba56495dc04ce00940534e56d90ee2c5f86)
- voltstack_cluster_ar.accelerated_networking.enable

<a id="canonical-10fbd9fb164bf9c90d43fd9396e6a9bcdb90381a193ceceed03fd14538752a15"></a>

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

<a id="canonical-a31afbc53efca3f77c1a215bba23ca045789eea9fe6783dc8b3e9da3b1a5d598"></a>

## Direct properties — voltstack_cluster_ar.accelerated_networking.enable / 6bf630ef4b29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff836f773285f256cf04330843ad13d156c91e9ec945a10b335d77a924530dfe"></a>

## Next pages — voltstack_cluster_ar.accelerated_networking.enable / 6bf630ef4b29 / 4

- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-d5e1b7a2152e82a99f5d910cbc8e1ba56495dc04ce00940534e56d90ee2c5f86)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-dcd77dd26bc732a2354140283470665b9df2466c8be936f97d1eb1c7ed643d0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02de0df469d13db751d92258c116e2acbde56f8a78b227294ac2c9dcecc4dde1"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies — voltstack_cluster_ar.active_enhanced_firewall_policies / 3df84e3835bd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.active_enhanced_firewall_policies

<a id="canonical-25d5793403bcd23d9fc2ee5f1704153dd034cc2d705c9b67b9890bc42a313d56"></a>

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

<a id="canonical-1dcadc5cf4858c2505f985d96e96c6a44a093831e8c7d03a46b3c3fe6d6bbdcc"></a>

## Direct properties — voltstack_cluster_ar.active_enhanced_firewall_policies / 3df84e3835bd / 3

- [enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-ffc34bf7d556b0714c56361d65e9c6fa580168c833808b3197833151a88222b9): complete subsection reference.

<a id="canonical-c35008d2a21b47cb7c362e6a461a092750586fa025a6c5729dd11afd84ae6c43"></a>

## Next pages — voltstack_cluster_ar.active_enhanced_firewall_policies / 3df84e3835bd / 4

- [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-ffc34bf7d556b0714c56361d65e9c6fa580168c833808b3197833151a88222b9)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ffc34bf7d556b0714c56361d65e9c6fa580168c833808b3197833151a88222b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-529d121f88baa5d39ee0a13bde80c940f96876c8d7185990048833396dd65b5a"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / af5c8b75bc3c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-dcd77dd26bc732a2354140283470665b9df2466c8be936f97d1eb1c7ed643d0e)
- voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-cd0a7df9f2f3be8c04b0d66849e25d9942209b9fe6884db0cb55f095ee0fc7c7"></a>

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

<a id="canonical-b92e83e2d4b04af2f44ce48e6dac259de450b600f217f187cd15898cfa13921e"></a>

## Direct properties — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / af5c8b75bc3c / 3

<a id="canonical-64713c35ce90c23df76daf3bf515d3e44e47a2a0aece4e3cd392efa35d68ef96"></a>

<a id="canonical-ff996dfecbe3051ee31d6e523fdcb4bb36dcf1a4ad2babe9b4e5ba3bf35e51a6"></a>

## name property — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / af5c8b75bc3c / 4

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

<a id="canonical-5fdb1eb6d3e3a84fc4895895c65a52ab8da78c4089b91ff108ab8a0d4182d022"></a>

<a id="canonical-a90cf2ce71ce7115b0b2b7324d1d55ffcdd8c61c3024e5269ea099341c7baf00"></a>

## namespace property — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / af5c8b75bc3c / 5

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

<a id="canonical-4e713b5aa000a6d799a0b623f8686b0133623c2aecb085371ea9f266f9b8791d"></a>

<a id="canonical-5b2544c6a6e1b049723525239e3214d0e0c683bcf33bd31804c2a93afe66576d"></a>

## tenant property — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / af5c8b75bc3c / 6

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

<a id="canonical-6e0fc2d242f55233d6f65866f0be9e74d3d77ec07dfe036f18c28b7e4cb59e59"></a>

## Next pages — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / af5c8b75bc3c / 7

- [voltstack_cluster_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-dcd77dd26bc732a2354140283470665b9df2466c8be936f97d1eb1c7ed643d0e)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a5b131a73e65adc0e811dbcef693db09c0d537dd1eb99224b25965b578654981"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57a08c72bbf6802c8939bc9ec9ffa2db33faec72a94271bd97e2b995cdab75d3"></a>

## voltstack_cluster_ar.active_forward_proxy_policies — voltstack_cluster_ar.active_forward_proxy_policies / 774b5c39bff0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.active_forward_proxy_policies

<a id="canonical-9aaf47add637c09c492aeecc6aabf7ea53796926e18760b2b1115710cfd4141e"></a>

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

<a id="canonical-c41b8c3c9b9c3f4f3eafc3308166ff0fb2d334a2c7a2ea43f8d394edb9ddbdbe"></a>

## Direct properties — voltstack_cluster_ar.active_forward_proxy_policies / 774b5c39bff0 / 3

- [forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-224d23a80eb0daf510b3b8644667b784a258ccbb3022ab4bfeb5ceccbe6df7fb): complete subsection reference.

<a id="canonical-f02fdcd95adf209036fd0f6e85bd4b95dac49090ba4ea4efcb215446450641e7"></a>

## Next pages — voltstack_cluster_ar.active_forward_proxy_policies / 774b5c39bff0 / 4

- [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-224d23a80eb0daf510b3b8644667b784a258ccbb3022ab4bfeb5ceccbe6df7fb)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-224d23a80eb0daf510b3b8644667b784a258ccbb3022ab4bfeb5ceccbe6df7fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-135f8f125a5ff049a2d44efe84f690363213b0a11b5cb30aa60ca9865f24c3b0"></a>

## voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / 1ae8597e3e2c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-a5b131a73e65adc0e811dbcef693db09c0d537dd1eb99224b25965b578654981)
- voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-d631783de6677e2bed4022361f9700fc62fa778a6b097b752bc2956ccaaf2dc2"></a>

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

<a id="canonical-3eca153dc882aa9935333fc5b91f8c2ab9e2a9f46690e2fadc4644cf4dfc2b42"></a>

## Direct properties — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / 1ae8597e3e2c / 3

<a id="canonical-c558d55653c19e086c104c62d5a7afb7e61913813dcb989c1fa4fe993d1ca2be"></a>

<a id="canonical-d468d5fb3df9e256a68df1b37729068ca4671603fadf9d70c0c72085f477ffbe"></a>

## name property — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / 1ae8597e3e2c / 4

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

<a id="canonical-3feceecf28227853b658f5eb6210ec17bd318ac8ed9294de5ae8e13b3a046e86"></a>

<a id="canonical-ed90558a12a4eb5b3c6bc133150955eb996aaa59142204963b9ece2a13fe3333"></a>

## namespace property — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / 1ae8597e3e2c / 5

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

<a id="canonical-9520716d355cee6c8fcd148d3757bef20470eff2b2c603b100f65bf5607bdb28"></a>

<a id="canonical-94146cbc247d6b8269a1baecac0f714eeb8579ab11e2c0111e4f4fe4b16ef13f"></a>

## tenant property — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / 1ae8597e3e2c / 6

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

<a id="canonical-271e99a11abd207a53f30aed56fcd1c87a84b9ef43fba3d0f8bbcf95c781d548"></a>

## Next pages — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / 1ae8597e3e2c / 7

- [voltstack_cluster_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-a5b131a73e65adc0e811dbcef693db09c0d537dd1eb99224b25965b578654981)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-95daf8f5047512b12e269a1d0b68b26ad4c49ca102decb338e0d3ba96cc557de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f55bb6325f937f8c69aedbbb5dd005d6fa3ebbf9cd1a44b2fcb36e9dcc7b3b9c"></a>

## voltstack_cluster_ar.active_network_policies — voltstack_cluster_ar.active_network_policies / dc80972adec0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.active_network_policies

<a id="canonical-fffca38c938fe80840ca757723739d6a8fa447ba38db16c0a12fd4bab1431607"></a>

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

<a id="canonical-ff0ebab576049b49e38d797be1c3315b51eeff3d7911cf76b5d2e09a9992778f"></a>

## Direct properties — voltstack_cluster_ar.active_network_policies / dc80972adec0 / 3

- [network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-ddc02745974dcca64b96a1bb440343a21dbd247363fd9397e9e07857815d15b9): complete subsection reference.

<a id="canonical-7b46279ef59e97ef2538aa96349199cdf1567169aba32f3c659cee98142d87f5"></a>

## Next pages — voltstack_cluster_ar.active_network_policies / dc80972adec0 / 4

- [voltstack_cluster_ar.active_network_policies.network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-ddc02745974dcca64b96a1bb440343a21dbd247363fd9397e9e07857815d15b9)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ddc02745974dcca64b96a1bb440343a21dbd247363fd9397e9e07857815d15b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200c9c9a498567a81cf6c98d4104f565d7e6d1c6f0f759571dbead32d2716ad"></a>

## voltstack_cluster_ar.active_network_policies.network_policies — voltstack_cluster_ar.active_network_policies.network_policies / 4d262550b2bd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-95daf8f5047512b12e269a1d0b68b26ad4c49ca102decb338e0d3ba96cc557de)
- voltstack_cluster_ar.active_network_policies.network_policies

<a id="canonical-b8e72fefb7df99ec5bc20791881dc8ab3f0ca9363695c52502309f25da2c602b"></a>

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

<a id="canonical-89fc2c3219c577cf7e1a61fa6e61e21e3c655221b14a152ff52933054e0484d9"></a>

## Direct properties — voltstack_cluster_ar.active_network_policies.network_policies / 4d262550b2bd / 3

<a id="canonical-4495c15e42d11e078deaa1b1887a0c8efc59425e407cc96758684552e7e720a6"></a>

<a id="canonical-10995ff96ee6b6a1e958f11c58da36a2620b189468bc436aa336e15dfbe86db8"></a>

## name property — voltstack_cluster_ar.active_network_policies.network_policies / 4d262550b2bd / 4

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

<a id="canonical-ae8e1facf60adda7fe9941d4495ea173afb75c4950cdaa9dcb555e79f1ed4895"></a>

<a id="canonical-65ff16a91add69a226c8b26277824a615653c3b10188a563ec84ce8bdf980657"></a>

## namespace property — voltstack_cluster_ar.active_network_policies.network_policies / 4d262550b2bd / 5

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

<a id="canonical-ab0792afd5d16f489b6d2fde1fbc4cf838756fa410e124880c5f60dc9f631c17"></a>

<a id="canonical-efadf538b7a78dba5c155627b24f7e29f6047459b510033319db44351a20f691"></a>

## tenant property — voltstack_cluster_ar.active_network_policies.network_policies / 4d262550b2bd / 6

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

<a id="canonical-ac73e5498d8039aec85cae8d30e101d38c5f818286e155a7312c93aa81bc8cec"></a>

## Next pages — voltstack_cluster_ar.active_network_policies.network_policies / 4d262550b2bd / 7

- [voltstack_cluster_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-95daf8f5047512b12e269a1d0b68b26ad4c49ca102decb338e0d3ba96cc557de)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-b0bcd9866e7fd6642d4650a33600c474e83fdfe443c4bd586398c35480c56c3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e65d9c45d1602599a5aa23f3023c920167befe97cf8be0db33be2f052f0110f"></a>

## voltstack_cluster_ar.dc_cluster_group — voltstack_cluster_ar.dc_cluster_group / 6d93e7ad3bd7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.dc_cluster_group

<a id="canonical-97e37c274b0532ae68bf23c33ae5056027d5ae159105c7a9fd5a7e7a4b0f485e"></a>

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

<a id="canonical-7aaaa55021f787af60189c819e3e04220a2bde4e8577842f87db2413eceb89ed"></a>

## Direct properties — voltstack_cluster_ar.dc_cluster_group / 6d93e7ad3bd7 / 3

<a id="canonical-1e8f0ed2ccb8cd3aea9e21d28a9260eade4e9635a28bce01c6a0597afa1d6117"></a>

<a id="canonical-42771a90d843144eab5d6dfcd1014561fb2ceb850f7213479d7482be5d784239"></a>

## name property — voltstack_cluster_ar.dc_cluster_group / 6d93e7ad3bd7 / 4

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

<a id="canonical-b5cecc032da3b32900833931685a5a9809c2d12134c37f7eb50a616200d3622d"></a>

<a id="canonical-a1987a9057b21c1a10313084d2fa7ffade9a53e75ff78cfb3a4cf0352c56c387"></a>

## namespace property — voltstack_cluster_ar.dc_cluster_group / 6d93e7ad3bd7 / 5

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

<a id="canonical-e7b5f32bec53a0fbee2f57f2f2e19eab1a4954a5dfa0b95d5a362c0eca2e37cf"></a>

<a id="canonical-e96b13f31632cb16392702bf579c50e2d821b2ce9d47fa590773cd1252d651bd"></a>

## tenant property — voltstack_cluster_ar.dc_cluster_group / 6d93e7ad3bd7 / 6

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

<a id="canonical-61c558c24615ecb44162b1443eec4fe29bf4bed9fc91e83b240954a8733500ff"></a>

## Next pages — voltstack_cluster_ar.dc_cluster_group / 6d93e7ad3bd7 / 7

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-c9511d506bb8a71655e35c8822643d3140a7682648d6f0c34213c754fa3e3096"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e3acc595b4b5c39542f6f4b69d1fb8286b12b61ac592f00b5334fa56c51db7a"></a>

## voltstack_cluster_ar.default_storage — voltstack_cluster_ar.default_storage / fac95fa641ec / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.default_storage

<a id="canonical-4df089975e2d745516e2238886c7155d30aec683478dc8cb39386fe26b6b0b30"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default storage.

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

<a id="canonical-1e83cfa6d83e1bc5ec6c17050d63187cc98f7881e016280c5d413dd1988b6c67"></a>

## Direct properties — voltstack_cluster_ar.default_storage / fac95fa641ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4508e2b02e40a0f1fd915aadc9b8eec2ab98da0f19ee97d8166742c0dcd7e61f"></a>

## Next pages — voltstack_cluster_ar.default_storage / fac95fa641ec / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-433a56aa39be244c9e31dd32ae0f689659c25dac91beed57b2d5ec1d7fbda829"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c83e0eb2e12b8cdb73fc69bc01de28c28aaa6f8eee676fd59014be09574d2cf6"></a>

## voltstack_cluster_ar.forward_proxy_allow_all — voltstack_cluster_ar.forward_proxy_allow_all / e37adc8da527 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.forward_proxy_allow_all

<a id="canonical-e6ca858a88d87595bfa85259707f04b2b7000552f2e37bc39b59ea181bcd6cb5"></a>

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

<a id="canonical-ed3e4a61bb405cb24af7f12f47d3d5d45695d4ddd6dd6699c4d1b7aaf4640ca9"></a>

## Direct properties — voltstack_cluster_ar.forward_proxy_allow_all / e37adc8da527 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a89f708d40ace6214e96898f0d111da1ae6d5ed36db644ee8c90fc3571478e0d"></a>

## Next pages — voltstack_cluster_ar.forward_proxy_allow_all / e37adc8da527 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6103f6b89021fa4eb9bbc7e6448bea705c85bf4bfa991d19bfcb04c360c9849f"></a>

## voltstack_cluster_ar.global_network_list — voltstack_cluster_ar.global_network_list / 8c6dcc947a63 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.global_network_list

<a id="canonical-2496a4b6973e76499d43e0433f8fc74c75785d2be503ce5807ff20c0d2e519b4"></a>

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

<a id="canonical-346a0fd1756fbc5540ad47882ddeffbfacf3668271961e99041d4f19b0d7687a"></a>

## Direct properties — voltstack_cluster_ar.global_network_list / 8c6dcc947a63 / 3

- [global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a): complete subsection reference.

<a id="canonical-9a83068d0ccc1d2eb280b16e6d771eafaad18b668dc49b6b0b0ce9818d93577d"></a>

## Next pages — voltstack_cluster_ar.global_network_list / 8c6dcc947a63 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-644650f4dfc4ba3774a7402ded2f4b5160f4d1a1bbcca8337bdb1d8eb9f4f0f4"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections — voltstack_cluster_ar.global_network_list.global_network_connections / e102b794dae0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263)
- voltstack_cluster_ar.global_network_list.global_network_connections

<a id="canonical-69b8db4d0c42792f3a1e0dcab6106ef8149acf2f69433da184d4206272e17ce6"></a>

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

<a id="canonical-a499f11d13772953a0be626d67791275baebb1214bbdc51693f59896712b2254"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections / e102b794dae0 / 3

- [sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-736902c84c156a8d64dc244e54ad40dac83f2b5390c109bd6d28b234ee89b895): complete subsection reference.

- [slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0a408f73f47d29d853a104d8e3a3616794229bb1851067bbfbaffeb0748ad466): complete subsection reference.

<a id="canonical-d1f5c326a8caa698efaffa3012afa18ec8f6949526cc081f7f711e548ea80403"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections / e102b794dae0 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-736902c84c156a8d64dc244e54ad40dac83f2b5390c109bd6d28b234ee89b895)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0a408f73f47d29d853a104d8e3a3616794229bb1851067bbfbaffeb0748ad466)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-736902c84c156a8d64dc244e54ad40dac83f2b5390c109bd6d28b234ee89b895"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-901c848c94306dea61782a65b500836bf621226a6a1a5549598365364eee4329"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / c8055e2656b3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-eb051b56b010a00461ae6261a99a3111db18a9c7514e3157e83d70a655ab4099"></a>

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

<a id="canonical-11c41d12f4ee306c36689b3261d69936e19db57ad7d89a4a98e9ad8335dcd050"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / c8055e2656b3 / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-8ae0dbf5b7e250b17a0317061bca893ee7bf247c24d040e98392f8e091585879): complete subsection reference.

<a id="canonical-fc180b34a277da8950776a2a3fe7be18804217c7f6389dfbd5c570b3146a9a9c"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / c8055e2656b3 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-8ae0dbf5b7e250b17a0317061bca893ee7bf247c24d040e98392f8e091585879)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-8ae0dbf5b7e250b17a0317061bca893ee7bf247c24d040e98392f8e091585879"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80cf1b604cdc056637bcf22fde6179d14cfdcbf1dd696af113263bfe74813a7f"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 57afd1a52d75 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a)
- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-736902c84c156a8d64dc244e54ad40dac83f2b5390c109bd6d28b234ee89b895)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-8f4699aa22a013968779a1ed4f6597066c3fad1e1bfb9f9b4caaec266fb16da2"></a>

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

<a id="canonical-931e61363acee853624b26c6dcfc1c77c3b1985c232157f8a40ec5f5d58eef1a"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 57afd1a52d75 / 3

<a id="canonical-74629dd14f417b17eedcf4f2a37623e1477a65c550d37d6ef81e5567112965a1"></a>

<a id="canonical-4744739b0ae55f451d481ce6087d1f2a6eda999acaa8a086d72472e1b6ced134"></a>

## name property — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 57afd1a52d75 / 4

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

<a id="canonical-b5d2cfb0ac670e1881347b81b5b9b768762c6729319591293d5b329c72cfb975"></a>

<a id="canonical-812c844390f3fc7c272558f2899855d73c4ff4705ef8761d4c1730bfb5a5bb9d"></a>

## namespace property — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 57afd1a52d75 / 5

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

<a id="canonical-80c9668a2deaa98b56148922a018b3db2c10312616d2658dff5a9a5d5d83f453"></a>

<a id="canonical-e02ca86165a10472b9f76f3807529abfe69239edfdcbeb7c30dc2a550ba934db"></a>

## tenant property — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 57afd1a52d75 / 6

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

<a id="canonical-bce6dec4e65f08446c8014c2d293e8bfd2449405d729b66b6ffe2c76a3c4003a"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 57afd1a52d75 / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-736902c84c156a8d64dc244e54ad40dac83f2b5390c109bd6d28b234ee89b895)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-0a408f73f47d29d853a104d8e3a3616794229bb1851067bbfbaffeb0748ad466"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7660a1b034403387ed84373d341b4221e70cd38c01533445b5f29398b129afc4"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 74226c4850c2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-73401188bf1d23d00616fcd6711d9bddc2e3f50724e4d0a1839d79154045a68e"></a>

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

<a id="canonical-54cccfc7cf6f1a72ea40d8ab12f564bb4c36242d6799d5dd83c2c0100ce1ff02"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 74226c4850c2 / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-a858938a610efee356cccdd4ef39d695c65fa7c5f0bc367c7ff4ca9f28724503): complete subsection reference.

<a id="canonical-da21affda24d9404054a6e4359e873548dc4e4924c2aa2854082114178e1cf8e"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 74226c4850c2 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-a858938a610efee356cccdd4ef39d695c65fa7c5f0bc367c7ff4ca9f28724503)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a858938a610efee356cccdd4ef39d695c65fa7c5f0bc367c7ff4ca9f28724503"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19f6a82531bf8e3df4001d6f801e5e7fb3699ba5473d386b7075c7158e489451"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 9fffcac1fa44 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-e1ab7f3215f02d2d178821b7f51b9d801f47f9e4a964b92565ee98aaec7c7263)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-a9cf4778dd55caea0d3a2c037a42d3d231094d1ec06a7d49db0741d118d2fa9a)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0a408f73f47d29d853a104d8e3a3616794229bb1851067bbfbaffeb0748ad466)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-80ac428cd4e61d9422ad9596048d620a0e313d5078d0e516c40225cde82c8363"></a>

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

<a id="canonical-bbfdc03e712ca1e7d26c92b184ab0aa3069153c5a5150a44878264d439c73bc1"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 9fffcac1fa44 / 3

<a id="canonical-adc9a75d8a97e8a58cf8a017b3fe7263865b61386ea52613e4b0d2fd127cd42f"></a>

<a id="canonical-6af26de9e41fd4bf0ebc3f9822ea0e1607982a4b037c5f1a505ae46a32a85419"></a>

## name property — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 9fffcac1fa44 / 4

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

<a id="canonical-67bba025c56fb63c797619c68c126510e57a4b0546b6b17f50161bc45bd58879"></a>

<a id="canonical-df10dd2c88cc77040fa2d80428a6ea994f365df0273551f514d724b4b174f17b"></a>

## namespace property — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 9fffcac1fa44 / 5

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

<a id="canonical-937966aa2eae9e01a4b3a6320a0697934f7b79377d17f1a60d4db00da264469c"></a>

<a id="canonical-76bd4ec4a2a56a94cda325eec0a864b764e8ccaa9f5ad85003243782de424095"></a>

## tenant property — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 9fffcac1fa44 / 6

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

<a id="canonical-68d7ef0712e06384aa04a0d963b53f5531700f50fb9795e23f2fda6c68f61303"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 9fffcac1fa44 / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0a408f73f47d29d853a104d8e3a3616794229bb1851067bbfbaffeb0748ad466)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a4732fd25019dd2da4fcda8a6f9224587fadbd5acffcefaee82e7aee0ffb03a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf10903c44d992a54814eb56de20e445a9ae46d3343c27c9dca52f81992343cb"></a>

## voltstack_cluster_ar.k8s_cluster — voltstack_cluster_ar.k8s_cluster / 82c1d183388b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.k8s_cluster

<a id="canonical-fdfcbc1a276645765783ce0b6a6fa8158090516ddbbb2eac292dbe9b6cabb483"></a>

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

<a id="canonical-76f41dfbce723348cd547f9d8fb29428735e948e930188b3b6211b61071a82fe"></a>

## Direct properties — voltstack_cluster_ar.k8s_cluster / 82c1d183388b / 3

<a id="canonical-a905a4934722c48f7e8a59db4474fbc2d87d44655eab72e95ce77d2bd736959b"></a>

<a id="canonical-0c0c6a81c8edb69057fb1a82e9da693132364e1434ff73747f6db4a1009b3259"></a>

## name property — voltstack_cluster_ar.k8s_cluster / 82c1d183388b / 4

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

<a id="canonical-8598723c3caff2e6e47c8a976e1cae5c279921fb61477e28cf94be6fed77e8d7"></a>

<a id="canonical-952bd0e893f34bcaaa9825d90528258fc1a1639ed15bde3702d340eeccb45365"></a>

## namespace property — voltstack_cluster_ar.k8s_cluster / 82c1d183388b / 5

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

<a id="canonical-340e70855f58f0e67f7843bc84bca2224b7417e1d87b694ed112f0d1d7fd2052"></a>

<a id="canonical-9db6421fe2db767e683f4cc543147e8cad7dbb736d90d7048aef6c517101e762"></a>

## tenant property — voltstack_cluster_ar.k8s_cluster / 82c1d183388b / 6

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

<a id="canonical-811a497b01ff57f9ad73961fd7ed9e8be68fbc5728d94b71eeda8d0734712637"></a>

## Next pages — voltstack_cluster_ar.k8s_cluster / 82c1d183388b / 7

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ae6e0cdf9444360bf277b441872027f8c4cd60bb5018512e8e4b4a3100dc095b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d654f294858e713ebf3722b8a62f1758742ccd9f877bb989527e7d6e3f7011a"></a>

## voltstack_cluster_ar.no_dc_cluster_group — voltstack_cluster_ar.no_dc_cluster_group / ff0558288fe3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.no_dc_cluster_group

<a id="canonical-3da215176406e8cbe49799ade5bc0bd208ba9a0ec5fee0a781ed8e6c8bd594e7"></a>

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

<a id="canonical-63916c25bdf77402c27ad5822074bf78891ec70c9b6e03bcb4f921dd1fbff8e1"></a>

## Direct properties — voltstack_cluster_ar.no_dc_cluster_group / ff0558288fe3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd5e37ca7d8fa19262e968b8f53cc1d60daa8e05bc8fe4295c72643508cd88ef"></a>

## Next pages — voltstack_cluster_ar.no_dc_cluster_group / ff0558288fe3 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-689e6c85553974c4f315e416838f05b7f707b51c6e5a05bc827e6b8d94f189e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63c67bd48626e83b01c6e2e02e43157caf269cf41d920097b0e4b3e6d7f8f6d7"></a>

## voltstack_cluster_ar.no_forward_proxy — voltstack_cluster_ar.no_forward_proxy / faf55af50db0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.no_forward_proxy

<a id="canonical-ec27f364ce9d0109cbae4e8d7b89d554cab0ae115f8c4b5ace6aacdb3ec2bdd8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-1d074c31f6e2bece5badc015821b9c1e5904e463109a32c9a93155318c0a1d47"></a>

## Direct properties — voltstack_cluster_ar.no_forward_proxy / faf55af50db0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-942b6c7041be3c9bd520db75260b25390639c6dfb2f08cea21060f241d28cc2e"></a>

## Next pages — voltstack_cluster_ar.no_forward_proxy / faf55af50db0 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d3bdf49ae2bb1b8bd5c7797401199b07dd6bbabda699cb6cbb267be2972a4278"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dc3979c58e9325cded1282c676b271a1b74701f9e2719645e098b702f1d51fa"></a>

## voltstack_cluster_ar.no_global_network — voltstack_cluster_ar.no_global_network / d0378e6e1498 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.no_global_network

<a id="canonical-f18606e4677bb9411adff1b069e1adf4e3426d055e523d833f7161850b04fc5e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-b3c349defac557cc942c29fa5aeccd15f652776d71c2b4da9c465acda84061de"></a>

## Direct properties — voltstack_cluster_ar.no_global_network / d0378e6e1498 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1f57df468719bba1fb065be72a4a4fd460e7020eedc0d16d2ebff83d6034d50"></a>

## Next pages — voltstack_cluster_ar.no_global_network / d0378e6e1498 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-309ff2d1789f13ed9b981ff6f7a85ed57648688e6f1893bcb868d58b179d7cf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6663a0929883a9ce83b24ec68b021fffecd73d3964e9154bb1f04a2707df74b4"></a>

## voltstack_cluster_ar.no_k8s_cluster — voltstack_cluster_ar.no_k8s_cluster / 9d11b846f549 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.no_k8s_cluster

<a id="canonical-cc7a4640933c0d4b172a8d34b9affa3eb089434c87b4868e0188cc2105adfced"></a>

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

<a id="canonical-0ae4170ae791816d7957ba228989de140fce180767c1308f0cedf4b4dd86c1f1"></a>

## Direct properties — voltstack_cluster_ar.no_k8s_cluster / 9d11b846f549 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f0e1f926438f6efa086971eb635050a850f6b75597a6071ef3584d34da61deb"></a>

## Next pages — voltstack_cluster_ar.no_k8s_cluster / 9d11b846f549 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ddff7218ccd5d6e45cd856f00a9ed6ee04e3a03fb7004e3ff2744fb1b1d09df0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78cb932a16e9e87380a2fe8bbd45e4801656f658b5d8488d11c1dcabde321a14"></a>

## voltstack_cluster_ar.no_network_policy — voltstack_cluster_ar.no_network_policy / 7f257c8ca64d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.no_network_policy

<a id="canonical-d606647b0e9c72a9bc7d82ae77a362d5bc9c9be7b74e10e61e02078a0130d784"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-2f26b33c16b5570991bedf94b35a44f8148aa8b92cd3db14e7d7a6e6cde8c615"></a>

## Direct properties — voltstack_cluster_ar.no_network_policy / 7f257c8ca64d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac79e632b58a037d7e480ec4d550f74a75d62edeec46c0b409392a653c8c07d5"></a>

## Next pages — voltstack_cluster_ar.no_network_policy / 7f257c8ca64d / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-11c9b02e198b98e08ec66a9051e2ccb9f556948c90dcc9bf622ffc00a0a001e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e565475af7130cea1ddd4103778f4ae41fca3003d55bd6e2b0d902b5d8bb7aef"></a>

## voltstack_cluster_ar.no_outside_static_routes — voltstack_cluster_ar.no_outside_static_routes / d2ce4ce8fe94 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.no_outside_static_routes

<a id="canonical-e0c26ceecf96e64fed3112628545fcb534a47824970f5176cf3225e27450f311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no outside static routes.

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

<a id="canonical-9ca872ab5707e06a7a2242ca13c55f2f399db4357348266273d0d7d33b7a6126"></a>

## Direct properties — voltstack_cluster_ar.no_outside_static_routes / d2ce4ce8fe94 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a592d1e762cefc99cc41344970817b89ea90b8b4812e94d5bbd7c93b793a679"></a>

## Next pages — voltstack_cluster_ar.no_outside_static_routes / d2ce4ce8fe94 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f672f4a3aad6da14aaabb887bafbff9f6d73da12da53990a203e2d6d7a56ee16"></a>

## voltstack_cluster_ar.node — voltstack_cluster_ar.node / d9b4a4f9604e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.node

<a id="canonical-51638beaa5db21e3e4a56a621c75406460b789abdfee055c899e32cf53b47bc1"></a>

Type: `"single"`. Computed.

Parameters for creating Single interface Node for Alternate Region.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7d5a71c3d63056d5b7df02c506baa20536f6bc6adfb1efecdec193db4bbc94d3"></a>

## Direct properties — voltstack_cluster_ar.node / d9b4a4f9604e / 3

<a id="canonical-da0caa31f7e174c636ad09c0bf91fe30da5a9e604ca39bb8a45150e99fd28887"></a>

<a id="canonical-e4cdac22c7d8e27abf8fb76de47c49ee4749074fef434288406846fc5e3be59c"></a>

## fault_domain property — voltstack_cluster_ar.node / d9b4a4f9604e / 4

Type: `"number"`. Computed.

Namuber of fault domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
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
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca): complete subsection reference.

<a id="canonical-ac3762e9362e9c518c5c11da4c88ba5040dfab4b9930151e4c63fe92ab4f677d"></a>

<a id="canonical-e75eb883e84ace0145841fe9e8097e21f644e8a46de185d1af84e77ba528777e"></a>

## node_number property — voltstack_cluster_ar.node / d9b4a4f9604e / 5

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

<a id="canonical-0083680732bed1d3af0cf2bcf23c7e6e2d5ce9b6227dcde0f55ce84f62be652c"></a>

<a id="canonical-c2e09291c985ee0e4d1a26142c7265e84723d43af41aa3e5dc3bbec2adb8d521"></a>

## update_domain property — voltstack_cluster_ar.node / d9b4a4f9604e / 6

Type: `"number"`. Computed.

Namuber of update domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-36256f9a5073f380909cf2fe6c0c07ffdf84566bd3a2f6a2f2679268e0e9c6a9"></a>

## Next pages — voltstack_cluster_ar.node / d9b4a4f9604e / 7

- [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d351caf919a3c870af942865b75f8d2861f6d6910043289e13265f1cb1b5cc05"></a>

## voltstack_cluster_ar.node.local_subnet — voltstack_cluster_ar.node.local_subnet / a22a23cf0015 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--reference--group-009.md#canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd)
- voltstack_cluster_ar.node.local_subnet

<a id="canonical-29436f3d5e7ad27c7290272348f607441e351c0019dec77d2907c644d14b671f"></a>

Type: `"single"`. Computed.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-fc073c1e33e60779ac4045c10e8e1462948d396849376d404a7fe7a746c1ba83"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet / a22a23cf0015 / 3

- [subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-ec16187a2f0235bf77b17b1e2cdca44c4c0ce3d5ab93d76c52490bfeebdef6b8): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-009.md#canonical-353a3ca70725954a3a3dd5fa8c71dcfdac2031eeabf3e7f6d3157705d414fe26): complete subsection reference.

<a id="canonical-6a16da696210bca93754871723d42b811c0773a3d36b543b81bf2ebe1f66bf75"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet / a22a23cf0015 / 4

- [voltstack_cluster_ar.node.local_subnet.subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-ec16187a2f0235bf77b17b1e2cdca44c4c0ce3d5ab93d76c52490bfeebdef6b8)
- [voltstack_cluster_ar.node.local_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-009.md#canonical-353a3ca70725954a3a3dd5fa8c71dcfdac2031eeabf3e7f6d3157705d414fe26)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--reference--group-009.md#canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-ec16187a2f0235bf77b17b1e2cdca44c4c0ce3d5ab93d76c52490bfeebdef6b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-885dc281d67a209accb9df8bbf84aed962bde8aea23db0d61834483f7306b598"></a>

## voltstack_cluster_ar.node.local_subnet.subnet — voltstack_cluster_ar.node.local_subnet.subnet / 8e3eee25e40e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--reference--group-009.md#canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd)
- [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca)
- voltstack_cluster_ar.node.local_subnet.subnet

<a id="canonical-3ae49b7306f58a0b66153386deed5a8806bbaf00ad4485cd537ea92096276b3d"></a>

Type: `"single"`. Computed.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-814525d47653cdaecc3585d841b9f6bb1c2af929a95930c7ceecf815d9356def"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet.subnet / 8e3eee25e40e / 3

<a id="canonical-1f2276dae048b19bb301b1a80388231da4dc3c860019adc78dd19c1beb1963a1"></a>

<a id="canonical-298cc706145bbcd2788b5c8262d7f8fb9273c5b55986016a60f9646951fb7a1d"></a>

## subnet_name property — voltstack_cluster_ar.node.local_subnet.subnet / 8e3eee25e40e / 4

Type: `"string"`. Computed.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

Receipt-pinned upstream constraints:

```json
{
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-a9b63e29fbba45de6fc505472ecc3a6d0e26576c222e2ffd0adddd10854a712d"></a>

<a id="canonical-4f88069103dfba8fb01a3f3e2465d9520adfe96c5eece38848e66b4d02b9d850"></a>

## subnet_resource_grp property — voltstack_cluster_ar.node.local_subnet.subnet / 8e3eee25e40e / 5

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Receipt-pinned upstream constraints:

```json
{
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-a521f2c08e5ea309fa58d05751d8f5f74ad5b63905547ab01ac9581222b7bbd0): complete subsection reference.

<a id="canonical-179e0c40980adee3a706628776b37af83cafeb38f2c3793076b3b56a0b30e2c6"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet.subnet / 8e3eee25e40e / 6

- [voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-a521f2c08e5ea309fa58d05751d8f5f74ad5b63905547ab01ac9581222b7bbd0)
- [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-a521f2c08e5ea309fa58d05751d8f5f74ad5b63905547ab01ac9581222b7bbd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e9ee333994820d24794a6e0ff308d30268f3e41f003eaf6f8d4a2e13e6e6d50"></a>

## voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group — voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group / 0528f75b2b9b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--reference--group-009.md#canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd)
- [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca)
- [voltstack_cluster_ar.node.local_subnet.subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-ec16187a2f0235bf77b17b1e2cdca44c4c0ce3d5ab93d76c52490bfeebdef6b8)
- voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group

<a id="canonical-8e8320a9b98ad9a27f56dba13e30be75f9e2c64f44fe7949f6d1fd7718db5e41"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-4a2ebea20dbf1aee7e774e23c5ccc0f7f7413c0c99e9dbd1018beead11df6e99"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group / 0528f75b2b9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d736bc28c2e5f7ea0434d16bfb4d3458d5187cb2b4950427300fdb1bb523bfd"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group / 0528f75b2b9b / 4

- [voltstack_cluster_ar.node.local_subnet.subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-ec16187a2f0235bf77b17b1e2cdca44c4c0ce3d5ab93d76c52490bfeebdef6b8)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-353a3ca70725954a3a3dd5fa8c71dcfdac2031eeabf3e7f6d3157705d414fe26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c62bbb16ed1fe48d55f577cb13d230f9101794102c0a8277a4f7e29daee834cd"></a>

## voltstack_cluster_ar.node.local_subnet.subnet_param — voltstack_cluster_ar.node.local_subnet.subnet_param / fdc88a8906e0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--reference--group-009.md#canonical-818d2019cae2f0552b7a08c2d4d4c3c97f9b0b10738801796144dd1ec36fc3fd)
- [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca)
- voltstack_cluster_ar.node.local_subnet.subnet_param

<a id="canonical-0342b0b96d8b88cac38e4392dccace149a790fbf2427296fec874de1d6182a1f"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-629feb4a79e9cc2db571fc9537f82e2f3c824c9a303b17500598268881118b59"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet.subnet_param / fdc88a8906e0 / 3

<a id="canonical-e3818f0d3880733512fdde48f3733b659a51da627fbaaa0be8ca451cd65c7137"></a>

<a id="canonical-5f67d91165a4669b5e2452691e322a5b3dff2f5b00dc7cc6f880b61dd2eefb96"></a>

## ipv4 property — voltstack_cluster_ar.node.local_subnet.subnet_param / fdc88a8906e0 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-fec14f052421f2e503c9123bc17d426d05ff5642b03f4ef22746a31cc00fb92a"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet.subnet_param / fdc88a8906e0 / 5

- [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-f4d34399d8096a786079179ab70bb3a56754bb1f6e839eaf7b553d7376e496ca)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f060d03e2a467a64785a8e3f4c297c88ae5b35d5f3481452b33a133823b24f47"></a>

## voltstack_cluster_ar.outside_static_routes — voltstack_cluster_ar.outside_static_routes / 95a27588fb55 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- voltstack_cluster_ar.outside_static_routes

<a id="canonical-0170bddd287b8b3fc74b748dad58846fd3b5fd7caa037aab183f0fa8338adc20"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3d5c19b463d784f9f682c2334129b13ff2b1076976bb8ee76574cd2655d88de2"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes / 95a27588fb55 / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb): complete subsection reference.

<a id="canonical-b7bddc167ce6fc71ef48ff2baa0f7284ecee453f767d7cd9bcd29367291ab6a7"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes / 95a27588fb55 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d2e4924d0d3aa250302e8e57ae422819a88f3dba83c4f924dc64e55f00330ea"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list — voltstack_cluster_ar.outside_static_routes.static_route_list / 3c1643606edd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- voltstack_cluster_ar.outside_static_routes.static_route_list

<a id="canonical-5c8179336adc100d14d307fe1146e4993b49187eb7383f0895b7d57a60e6540b"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-ec3fd41fe78f60f64d3ff9387150c87a5672d617f7c2761bf93750f250a5407d"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list / 3c1643606edd / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5): complete subsection reference.

<a id="canonical-8bbd2155cb796293fc97d6263b108b61e38fca5dcbf343d9c53efc16163f9bdc"></a>

<a id="canonical-ec06d06ef8381c39ba15bfb19ed82658e313d635ef64b3d55e1956b0769156cb"></a>

## simple_static_route property — voltstack_cluster_ar.outside_static_routes.static_route_list / 3c1643606edd / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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

<a id="canonical-8d4f178ffd2848be53bd76df2b93e1d4125f558e050e730de0ab9bafd2ea822c"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list / 3c1643606edd / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5dc6315fee97326257c8e9b7494b3af686529605b5f911c378c4f2af84ce528"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 4218d2d10aba / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-da535556a0e4817d2d6b11a8968664cf3b2bed97206b3d95253daf797384ca9f"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ff6c6368083b933061f5c1639a0bc1c8f15245160d89c0bd58e45a2311d49b93"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 4218d2d10aba / 3

<a id="canonical-96bc4d1703ce15063b983612778ae53a6b0c58693cdf1287fa4d2e14d1a62764"></a>

<a id="canonical-30851bf6312eb56ab14e6164ded76a49915283db0e1d11b1dc72d8b28c66d7e5"></a>

## attrs property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 4218d2d10aba / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--azure_vnet_site--reference--group-009.md#canonical-2571c4358375c9551101979a1a8302f993a645c99c9c313b07c57d8647b9c678): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-d93d243c0bc10fd10bd47dd2c03ae61c441190b0df14d3fb08a024f18fa44736): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-010.md#canonical-3fe61c504a6d1ca7971273b45e33ace39c379dbec5ccfae5e27c97a2086f0186): complete subsection reference.

<a id="canonical-0649ebc1dc26c22b40eedc55450896cc2e4da3d1b4009cf48c51a26812be9ef1"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 4218d2d10aba / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-009.md#canonical-2571c4358375c9551101979a1a8302f993a645c99c9c313b07c57d8647b9c678)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-d93d243c0bc10fd10bd47dd2c03ae61c441190b0df14d3fb08a024f18fa44736)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-010.md#canonical-3fe61c504a6d1ca7971273b45e33ace39c379dbec5ccfae5e27c97a2086f0186)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-2571c4358375c9551101979a1a8302f993a645c99c9c313b07c57d8647b9c678"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21fa625581a24bab096aa1dc66942d6a1a05afc1a0f308eebeb8f08000746366"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 7d9c85d37769 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-d87ac3d6703b6d03cbbd679c3c72ac0cd4118b22a73e897000113eadf1f1225f"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-42104f141543f82ba2917fe25b7bc0cda161f46f9c232290549efe97052d417f"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 7d9c85d37769 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b388c334f78712ab5a83903b222cad8cbda2760a3d7c96235c06418e8d7f5bac"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 7d9c85d37769 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-d93d243c0bc10fd10bd47dd2c03ae61c441190b0df14d3fb08a024f18fa44736"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40b8b4d247e72b0f6f36628908175625edb6f98a5fc62a6de2cb413a9c88e74e"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / ee4630958d93 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-65c5af89c02279b45b69f1224f60d16cb90aa2c3ef42d3a3191c7296c514edc1"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f0d7dcca6b69e882cde8fe92dbde9dc8be9abb2b190528106631ac52fe5e33ee"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / ee4630958d93 / 3

- [interface](data-sources--azure_vnet_site--reference--group-009.md#canonical-edf35103c35b2411c9496004796ce550901c3aba0681b6810ec03c9e19000b8a): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-f9152fd1f094e1989d9b6be5129091432054b59461f17745ee798a490697be57): complete subsection reference.

<a id="canonical-515f4b394d4f94630358830f02d0b5970dc21e794942dd5a2ca9f22f1ad16eaf"></a>

<a id="canonical-2ed28d3f49fc6a5253189107256aa109a876163b59eda36a74ee091e524b3397"></a>

## type property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / ee4630958d93 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0ebf43195ccb4e37bd383ce59d3d4ce299d6677272433c81cdb30a13ca34c5fd"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / ee4630958d93 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-009.md#canonical-edf35103c35b2411c9496004796ce550901c3aba0681b6810ec03c9e19000b8a)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-f9152fd1f094e1989d9b6be5129091432054b59461f17745ee798a490697be57)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-edf35103c35b2411c9496004796ce550901c3aba0681b6810ec03c9e19000b8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74aac4d4c5f94bf12572a3170e88c503bf65b348dcf2c728d37157b5bad6c53e"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-d93d243c0bc10fd10bd47dd2c03ae61c441190b0df14d3fb08a024f18fa44736)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-ec9ee56626eb1f53e384dcbd687a36d457b7999592c9ea1cba509a5c9840b8f9"></a>

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

<a id="canonical-e5d81eb1902252c4ab7be3257b510797f03e751c506baae3ea34db665a65b1f6"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 3

<a id="canonical-d1c0641b182ca7d1eb56caeff1c3984b52234e9e0ce117047c887ffc87f44c7f"></a>

<a id="canonical-e2a4fe24101eb274bfce17269d10deeca76471535e4c7ae9ca06147869e37551"></a>

## kind property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 4

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

<a id="canonical-9a07c302a0c14bdc2480b0460a67fe2a2f53e06cabaaec43b94b37f5fc869cd9"></a>

<a id="canonical-bfc67bb0504f9a24e696a9c81937a6f72681377fd70dee8784f227eec33fc763"></a>

## name property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 5

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

<a id="canonical-e75e6b08df9c3a5a13084125a247d26dce57d1c7e08561c19909e8138ec441bb"></a>

<a id="canonical-b88803e6cf8c27d44cd8afa112c921e95d26c58d2d21d1bbb9e8a65bb33507ee"></a>

## namespace property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 6

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

<a id="canonical-91902fe35596fc4fad2ccf88c60fd8e4afce69d1c53cb73613a861678e06a607"></a>

<a id="canonical-15761cd32df2ba2a89a81443aa71cac65c53dc38a82cf58ae0f373ebb1822383"></a>

## tenant property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 7

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

<a id="canonical-d22a8ad9996003fbd89ac039446f95eaaa72eb237f31a1268502010a49c89f2c"></a>

<a id="canonical-c9fe501c44a09562ce8bb3c3f321e29d6430fb101733a89493e8140fa6c9de57"></a>

## uid property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 8

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

<a id="canonical-a2c9313552b3f1dc6cbf5c65e3becedba6627f26a092575232e83b02288c5cf1"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 46e4a4cd999a / 9

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-d93d243c0bc10fd10bd47dd2c03ae61c441190b0df14d3fb08a024f18fa44736)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f9152fd1f094e1989d9b6be5129091432054b59461f17745ee798a490697be57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6bb3db8b3f362f61fb296224e457df46fc57f50438f2a6d78fe8f94569ffd88"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / acee83ce8906 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-07f7cbdbcc30193eb1eca8a3acc6cf23b4fc6c1b14d5c89d19158bd01a5f0197)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-80bfa02ae81def835ca7c9ac3e1c5aa752cb2328b9c9aa9357576b035055470a)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-cff00fde2b4c16f2ea7f8200d3e9b6a05c8313690b32c0cf12e22b21c30cadbb)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-9126be8c7c2f7c567ec2e319adb6742c6d24821f429ded66084938d91b88d9d5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-d93d243c0bc10fd10bd47dd2c03ae61c441190b0df14d3fb08a024f18fa44736)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-4f4be4e8e2140dca17f61923173a988e00c8dc823c11e4319359bfb5c232443f"></a>

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

<a id="canonical-0cd3692a227ad184f24d784052bea20a0616e127595a3c83be1a4e69fa3fefb4"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / acee83ce8906 / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-f94bd9ef48934aad9e0f1d7159532da3903b1f5ac3cc888e7399ca317aced24f): complete subsection reference.

- [ipv4](data-sources--azure_vnet_site--reference--group-010.md#canonical-f04ed0ddfa83c3b87306443ffed9eb5451ffe13f307facae6b2bcdd1c91b0562): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-010.md#canonical-47c3472b10d1affe3307b4a231da2c815b78c997f8d643c272f925d5c9d161a1): complete subsection reference.

<a id="canonical-b1105b0f22a90a2c82caf5736c55437402f3c983498a8c7581abc04f254d95a1"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / acee83ce8906 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-f94bd9ef48934aad9e0f1d7159532da3903b1f5ac3cc888e7399ca317aced24f)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-010.md#canonical-f04ed0ddfa83c3b87306443ffed9eb5451ffe13f307facae6b2bcdd1c91b0562)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-010.md#canonical-47c3472b10d1affe3307b4a231da2c815b78c997f8d643c272f925d5c9d161a1)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-d93d243c0bc10fd10bd47dd2c03ae61c441190b0df14d3fb08a024f18fa44736)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-f94bd9ef48934aad9e0f1d7159532da3903b1f5ac3cc888e7399ca317aced24f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
