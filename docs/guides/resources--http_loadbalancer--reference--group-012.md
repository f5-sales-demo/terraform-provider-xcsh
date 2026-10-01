---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2a8d0eac85265ff917e11175cb1f4aec79663bc80dbf9d46bf726024acb3a16a"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services — bot_defense.policy.protected_app_endpoints.flow_label.financial_services / 7ab5dbd14dbb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-0bde38b6fd8705f10865d9c6576690a81a15b739f542d7912b671f30f3fafa81"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Financial Services Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("apply",
    "money_transfer")}
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
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

Terraform syntax:

```terraform
financial_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-4249e6f53579e45ec3c47401420c78b1475bb62e7e6fc2a4b15d79e2e90c7a95"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.financial_services / 7ab5dbd14dbb / 3

- [apply](resources--http_loadbalancer--reference--group-012.md#canonical-ad896a62083378f2b88a6dcdb845ac4b689e34a81654b93103156e0c8b5249f9): complete subsection reference.

- [money_transfer](resources--http_loadbalancer--reference--group-012.md#canonical-81715fccbf765c298e70e781d98561d791b13a8f495013874fed1261a8b659c1): complete subsection reference.

<a id="canonical-3864ebb47113a45cae279fbf216531b5dc7c9e3714bcb48d79c7fbac785eb48b"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.financial_services / 7ab5dbd14dbb / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply](resources--http_loadbalancer--reference--group-012.md#canonical-ad896a62083378f2b88a6dcdb845ac4b689e34a81654b93103156e0c8b5249f9)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer](resources--http_loadbalancer--reference--group-012.md#canonical-81715fccbf765c298e70e781d98561d791b13a8f495013874fed1261a8b659c1)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ad896a62083378f2b88a6dcdb845ac4b689e34a81654b93103156e0c8b5249f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2453012014222505681e0a6f5cbe23b6e8d4a7eee5bae156da0528ade236c5f6"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply / fbe95ebd07ce / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-011.md#canonical-b7f2a614aa37817fa4cfda14645805e918d31e14f88c9070fb7fe9bf0e244a56)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-332e126367f4877c99b9cac70b4a79ba6eef8e11439a5496bfbc8233945702c8"></a>

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
apply = {}
```

<a id="canonical-6caf4358b0547e0cfdd842bc0fc9174dd633dfc56eaf181218fe0326a316dc54"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply / fbe95ebd07ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-180a6f131e02d0becc84dbb60548dcd3f1d15ff650865634a3a60a8b561b0a6b"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply / fbe95ebd07ce / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-011.md#canonical-b7f2a614aa37817fa4cfda14645805e918d31e14f88c9070fb7fe9bf0e244a56)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-81715fccbf765c298e70e781d98561d791b13a8f495013874fed1261a8b659c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0693e0e24d051b2ba75e893f0ab14135acd70dd58b8be8be3621f35a3fcdf847"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_t / 85ed830a1244 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-011.md#canonical-b7f2a614aa37817fa4cfda14645805e918d31e14f88c9070fb7fe9bf0e244a56)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-0000ff2a9f981a5fd8b320981603daf7281fcb933a364c98aeabc07d49b0f7a6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for money transfer.

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
money_transfer = {}
```

<a id="canonical-2cd72675bc09df3ec5a82b3d4bd34297037cd97716d2ef5df9b35a8011306945"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_t / 85ed830a1244 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07c11b4d02fb3fd87347733eaa6b10907ca18997b496beefb89176560e6d9e76"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_t / 85ed830a1244 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-011.md#canonical-b7f2a614aa37817fa4cfda14645805e918d31e14f88c9070fb7fe9bf0e244a56)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9d2291128668009bdee438ef7bdab823cf9e4a995b191a3ca18e9bd4e945aa5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4075bb88fd8220915c2a92ce22869261c660021739a2f0491c1850d2b89baa2a"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight — bot_defense.policy.protected_app_endpoints.flow_label.flight / 65a74e19768e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-16a9b723300e09bc2a35ab3133433c8dced1875ea83df05882d6e74ef36aa393"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

Terraform syntax:

```terraform
flight {
  # Configure direct properties listed below.
}
```

<a id="canonical-4cfa3498eb60d004c9106c2a2a11f00efabf2ef7aa561167e710c250dca9a2ae"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.flight / 65a74e19768e / 3

- [checkin](resources--http_loadbalancer--reference--group-012.md#canonical-247af7fd4927a43cb10fb4275946977ff757ffb45f592b80cba966c060a657e0): complete subsection reference.

<a id="canonical-bb560c92c39b18c3c92ba898008ae337202c1a5b33ce585b8d0b4748c1fe9906"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.flight / 65a74e19768e / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin](resources--http_loadbalancer--reference--group-012.md#canonical-247af7fd4927a43cb10fb4275946977ff757ffb45f592b80cba966c060a657e0)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-247af7fd4927a43cb10fb4275946977ff757ffb45f592b80cba966c060a657e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0127d728fae1870e3208a302a8055f70b1e4f9a71e83692e65919cac1137e401"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin — bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin / 93796e247f89 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-012.md#canonical-9d2291128668009bdee438ef7bdab823cf9e4a995b191a3ca18e9bd4e945aa5e)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-0a9653edfc84fbacb2fc684b745d965bdae9271664f5dbfb2116992a39b600e8"></a>

Type: `"object"`. single nested block, Optional.

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
checkin {}
```

<a id="canonical-f71eeeaf57ad589184279671ec58998311413dcfa58f254429bd440472ec2399"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin / 93796e247f89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-598a92464946494da068b49fb87a3af8fe01dc5b168f0b0b3b4eb9b77c521fc8"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin / 93796e247f89 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-012.md#canonical-9d2291128668009bdee438ef7bdab823cf9e4a995b191a3ca18e9bd4e945aa5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-439de2f247145caf96bc0b732e271e7c9f14f7172bbe1ab437d275cedec7cf75"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management — bot_defense.policy.protected_app_endpoints.flow_label.profile_management / 3389d7f95839 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-5bf50f69ef2b8ceac408b1373fc8f942af28b70c1da9f8f0fd9bb467ea685e71"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "update"),
  validators.ConflictingObjectAttributes("create",
    "view"),
  validators.ConflictingObjectAttributes("update",
    "view")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-3a8301274ec5383b454db7100467a91b37446a454f3a8bdd6c98621828a33d6a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management / 3389d7f95839 / 3

- [create](resources--http_loadbalancer--reference--group-012.md#canonical-fb544d73511fd27a442306d11cac35975a3cc31a9af498f315803cb17869d6dd): complete subsection reference.

- [update](resources--http_loadbalancer--reference--group-012.md#canonical-a2b6215d1d6c22e6d9da69963da216bc0623474e36986ec3a59098c0dc9acb21): complete subsection reference.

- [view](resources--http_loadbalancer--reference--group-012.md#canonical-852781cb7ca1bc91306cbdd3230f4757af143c5e4c4695a21a0493e88d5d0fe3): complete subsection reference.

<a id="canonical-765d452dece4e3d06e4e881f61dadef18c837bb1f05cb897c9647197ab9cf41d"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management / 3389d7f95839 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create](resources--http_loadbalancer--reference--group-012.md#canonical-fb544d73511fd27a442306d11cac35975a3cc31a9af498f315803cb17869d6dd)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update](resources--http_loadbalancer--reference--group-012.md#canonical-a2b6215d1d6c22e6d9da69963da216bc0623474e36986ec3a59098c0dc9acb21)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view](resources--http_loadbalancer--reference--group-012.md#canonical-852781cb7ca1bc91306cbdd3230f4757af143c5e4c4695a21a0493e88d5d0fe3)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fb544d73511fd27a442306d11cac35975a3cc31a9af498f315803cb17869d6dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cda629dfa770e19ab54e35a9329404e5fd459de5ea3080f5745f1755d4daeef6"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create / be1b8bdb9f37 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-d67a6e1b1e76fedcc0fe945fac843c0523c0b06c10be39eca98bd8ae6b25f13b"></a>

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
create = {}
```

<a id="canonical-84e8bda1b1172e90120ed6ad91a2b9577cb63ab7be61d8c15f4c2f1ca0171564"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create / be1b8bdb9f37 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7219accf6211f4dc4c3bd21ab528c9d0687a67d8b740ea1f2dd1b27f49486700"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create / be1b8bdb9f37 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a2b6215d1d6c22e6d9da69963da216bc0623474e36986ec3a59098c0dc9acb21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aee0d8a40bf90702af85d59c2021fcce40d2c3aaa0d0a2029999274c60c13f47"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update / 03d5a7ce8a75 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-a53c4cb6ad12cf0c461c35326e25179d99c2425e803671abcaea80f4c16230d4"></a>

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
update = {}
```

<a id="canonical-24d945f68384504d2b262b0a5d9a472aa80657ac3ec845a443acbf6256cc3549"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update / 03d5a7ce8a75 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4dc1614ae0a9f6bf9f9ae43f80d6c1dd71781b5de57b8e8acdb041f01ba9fb2"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update / 03d5a7ce8a75 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-852781cb7ca1bc91306cbdd3230f4757af143c5e4c4695a21a0493e88d5d0fe3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d71a04fbfc281e9fdf56991203e0ce8be3146b7fb8414f2aa08711296a44802"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view / eba422968c6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-d18ee6b58554f174474012752d74e643e9e0780cbb69ed722a52f37cd1b2155e"></a>

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
view = {}
```

<a id="canonical-cfefea4b3376472becb7c0e0e3748f11e96186ef6875ed06daf353bd69188ce6"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view / eba422968c6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-37bb8bf3cad587e5ba448c204e8e1824bb24ab2d65bcc9411cdac13dd14f623f"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view / eba422968c6c / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-3f7e8ce91a5a7a42ffcb08c818ed2ec58c3a5db7cd6d319ad9b2d00570648189)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6da33ef17d782f765baa3e1d1210ff8acf36f27c2e10207244f110a21f2335d"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search — bot_defense.policy.protected_app_endpoints.flow_label.search / b26ea9c8391c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-2cddbefb2885a68b3131afe5080da68b242ed97eda8b397e1e5f49214614690a"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("flight_search",
    "product_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "room_search"),
  validators.ConflictingObjectAttributes("product_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("product_search",
    "room_search"),
  validators.ConflictingObjectAttributes("reservation_search",
    "room_search")}
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
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

Terraform syntax:

```terraform
search {
  # Configure direct properties listed below.
}
```

<a id="canonical-a7b91844620cf093fc31b93bdf82fc84fa95bb30fc3c9ad1aed7ae8f59e4733a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search / b26ea9c8391c / 3

- [flight_search](resources--http_loadbalancer--reference--group-012.md#canonical-c36610a9b7a6d4e5edbdbf5150ef989d2eefe1c91eef53498511c6bbfa1bce8e): complete subsection reference.

- [product_search](resources--http_loadbalancer--reference--group-012.md#canonical-da4336e7fe2aba9e76a7cb4215a3f518598915c1871952dfbbd21a53cd15a352): complete subsection reference.

- [reservation_search](resources--http_loadbalancer--reference--group-012.md#canonical-56f26a1b469c0732712589e119b2f7e51ac132edf3345ba682600dc77454ad66): complete subsection reference.

- [room_search](resources--http_loadbalancer--reference--group-012.md#canonical-08963f219d148d080a877e20688f835cd72cb3fe7ddfb97c1548f8fc89f80fdf): complete subsection reference.

<a id="canonical-c18dd4e7d6b34250e5514ea07d7d2a7da80e557ecd32ffdd938b094cc8ca22b1"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search / b26ea9c8391c / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search](resources--http_loadbalancer--reference--group-012.md#canonical-c36610a9b7a6d4e5edbdbf5150ef989d2eefe1c91eef53498511c6bbfa1bce8e)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.product_search](resources--http_loadbalancer--reference--group-012.md#canonical-da4336e7fe2aba9e76a7cb4215a3f518598915c1871952dfbbd21a53cd15a352)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search](resources--http_loadbalancer--reference--group-012.md#canonical-56f26a1b469c0732712589e119b2f7e51ac132edf3345ba682600dc77454ad66)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.room_search](resources--http_loadbalancer--reference--group-012.md#canonical-08963f219d148d080a877e20688f835cd72cb3fe7ddfb97c1548f8fc89f80fdf)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c36610a9b7a6d4e5edbdbf5150ef989d2eefe1c91eef53498511c6bbfa1bce8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-452185cea14045e7ba232aa6086277e75bcd4d5bde1b0cbd3dafe0786c460f92"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search — bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search / 6e8b0bce79ea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-c816fd3c3bd4cbe7080d71d0fc08dd930ecf646346097e3c5315ba720efea410"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for flight search.

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
flight_search = {}
```

<a id="canonical-4e5f73ab5f1b6c7f5c4e32649a4e46b15a7387bc18581d995bfdba74262f6403"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search / 6e8b0bce79ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-105de4452123a1fba1b2d584ae0339401b9c89d9a10e884ff1c83cea13891da0"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search / 6e8b0bce79ea / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-da4336e7fe2aba9e76a7cb4215a3f518598915c1871952dfbbd21a53cd15a352"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e16914ae0ecdcbfb22efedbf3251f3b11c94042beb5f342545d9762386c43e50"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.product_search — bot_defense.policy.protected_app_endpoints.flow_label.search.product_search / 510115ff9c99 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-4d974c0740e3332da2ebd8e64925d33422f6be0cc3e0aff6d562e13027518afc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for product search.

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
product_search = {}
```

<a id="canonical-d7b35d3a88640672d5caa66188c944a1a59a00b775fb9fafb6024c18db4110b5"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.product_search / 510115ff9c99 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f54179ba889545fde8dc515af5b5756ff6d3205d8acfa3c2830a6192da09ec9"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.product_search / 510115ff9c99 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-56f26a1b469c0732712589e119b2f7e51ac132edf3345ba682600dc77454ad66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40f663523ba43c200d15054198446d7c37f31a1a058dcd4240b1dcf9c579e4de"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search — bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search / b175f6b8a747 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-445e3057d68c23166a216fae03bb2afae9f16a4967c80dfde2f0e05428293b26"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reservation search.

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
reservation_search = {}
```

<a id="canonical-e3f5fbc2cbdfc0c968ca543060fcd1789862a84d0c2f0f0c4d49514e822a6033"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search / b175f6b8a747 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-abda76c1b9d49f5db3dab47c4d1a09f7ed45e0b66ad085e27a7a11a12315f32b"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search / b175f6b8a747 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-08963f219d148d080a877e20688f835cd72cb3fe7ddfb97c1548f8fc89f80fdf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e46bb8db63c0372bd577008380e9220118dceb52431798c6ea7a983b8f46f7f"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.room_search — bot_defense.policy.protected_app_endpoints.flow_label.search.room_search / a0f5d83291ac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-5396b4b182b1b5d9b9744c646bbf0ca75c53977960afdccfa7fe6f9caee86411"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for room search.

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
room_search = {}
```

<a id="canonical-99a082b7177f2a4f37446706a85d8474cac3694275847359142f81ff5ec7af55"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.search.room_search / a0f5d83291ac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14740611631db151d61f6c48866d2a8d1fd626bb25ab1cd2320d412a1046f3e3"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.search.room_search / a0f5d83291ac / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-20f780669afe0a12646e944d9b2bddcdd5cd2a6c25288a422dcd320041ed2259)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e06de73b890627b6af6e6c7f0b07fff196d537225ec6dacbd67c477bca775ae"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards / 687b9238e4e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-b6f0a8af0d96d602c4707436bd99eb1b907bdac4bd9d5845648a557c5de23183"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "gift_card_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_purchase_gift_card",
    "shop_update_quantity")}
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
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

Terraform syntax:

```terraform
shopping_gift_cards {
  # Configure direct properties listed below.
}
```

<a id="canonical-2ea4b32e0187e8a9b6c8cf7668732263c4b59650677d27502260d6ea68dd5d68"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards / 687b9238e4e4 / 3

- [gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-1cc07d28089222312ca8cb79b597fb3fef8413879239024a38543ea5eb6da593): complete subsection reference.

- [gift_card_validation](resources--http_loadbalancer--reference--group-012.md#canonical-0264cb4d5486a6268fdeb6f1a3e73b5c633444743b018d3600529345196edeeb): complete subsection reference.

- [shop_add_to_cart](resources--http_loadbalancer--reference--group-012.md#canonical-9373003516bd8427c9a5d161fed8badc826071e7c994c5092ed393bce9e0cc63): complete subsection reference.

- [shop_checkout](resources--http_loadbalancer--reference--group-012.md#canonical-d39d09bbb1c1b39a39adf9afbaf18ea438f326786d91495c5c2f868e6b238138): complete subsection reference.

- [shop_choose_seat](resources--http_loadbalancer--reference--group-012.md#canonical-8803d586d51b6ff6d1c5841cb4cf35e38afcb36d33c77505b354f90fa277bbf0): complete subsection reference.

- [shop_enter_drawing_submission](resources--http_loadbalancer--reference--group-012.md#canonical-f8e7f944667289751c8cfd0027d5f813e298f29ac858cad687dbf9859fcd8671): complete subsection reference.

- [shop_make_payment](resources--http_loadbalancer--reference--group-012.md#canonical-bd149b9e9de8fdb2968a03777b77edaa35cc2089fda45e77f4eb6fc0c579a0c6): complete subsection reference.

- [shop_order](resources--http_loadbalancer--reference--group-012.md#canonical-4176c57c3b19ea3887ed4efc0114f14bb58477352bfdfbe2a00c00e44055cbea): complete subsection reference.

- [shop_price_inquiry](resources--http_loadbalancer--reference--group-012.md#canonical-7afe4a217aad2b2e250b59d0c632d9c99647fc181c2a70671fcb03963b3789d1): complete subsection reference.

- [shop_promo_code_validation](resources--http_loadbalancer--reference--group-012.md#canonical-7d5968b4f523ce74f598a4a7c0af1cfade5212d24e23b9c92adf906470aa3ea5): complete subsection reference.

- [shop_purchase_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-4e5e30d53fd3b1e0195c2c6b78c041637116f7b25bcbf30b45155e018aa4195b): complete subsection reference.

- [shop_update_quantity](resources--http_loadbalancer--reference--group-012.md#canonical-dca9b83bf456151b022690a3efb5b7e86afe8a379dead28c0768e2f39ad64dad): complete subsection reference.

<a id="canonical-aeb60a7da822d882fccfbbe1bf7b7dd1c69167bc87aef7a04bcc63b4d648ac35"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards / 687b9238e4e4 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-1cc07d28089222312ca8cb79b597fb3fef8413879239024a38543ea5eb6da593)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--http_loadbalancer--reference--group-012.md#canonical-0264cb4d5486a6268fdeb6f1a3e73b5c633444743b018d3600529345196edeeb)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--http_loadbalancer--reference--group-012.md#canonical-9373003516bd8427c9a5d161fed8badc826071e7c994c5092ed393bce9e0cc63)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--http_loadbalancer--reference--group-012.md#canonical-d39d09bbb1c1b39a39adf9afbaf18ea438f326786d91495c5c2f868e6b238138)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--http_loadbalancer--reference--group-012.md#canonical-8803d586d51b6ff6d1c5841cb4cf35e38afcb36d33c77505b354f90fa277bbf0)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--http_loadbalancer--reference--group-012.md#canonical-f8e7f944667289751c8cfd0027d5f813e298f29ac858cad687dbf9859fcd8671)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--http_loadbalancer--reference--group-012.md#canonical-bd149b9e9de8fdb2968a03777b77edaa35cc2089fda45e77f4eb6fc0c579a0c6)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order](resources--http_loadbalancer--reference--group-012.md#canonical-4176c57c3b19ea3887ed4efc0114f14bb58477352bfdfbe2a00c00e44055cbea)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--http_loadbalancer--reference--group-012.md#canonical-7afe4a217aad2b2e250b59d0c632d9c99647fc181c2a70671fcb03963b3789d1)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--http_loadbalancer--reference--group-012.md#canonical-7d5968b4f523ce74f598a4a7c0af1cfade5212d24e23b9c92adf906470aa3ea5)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-4e5e30d53fd3b1e0195c2c6b78c041637116f7b25bcbf30b45155e018aa4195b)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--http_loadbalancer--reference--group-012.md#canonical-dca9b83bf456151b022690a3efb5b7e86afe8a379dead28c0768e2f39ad64dad)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1cc07d28089222312ca8cb79b597fb3fef8413879239024a38543ea5eb6da593"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-750173697d25d7c26ae921cb2d661eda82917f0eb64e49f75c1090593ea05856"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_c / 054a15057e99 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-2384e2b9516861f25b2e10adbfd38ecd5d945fc563aabd880425ecdf0b84239a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card make purchase with gift card.

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
gift_card_make_purchase_with_gift_card = {}
```

<a id="canonical-ffe4207601b8da431da2257c09ad96dad61278ae31a9b77649b75f03d047a0d4"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_c / 054a15057e99 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1129d9f4f781e93bda7d6fb8e7f9a55eb545651e692d96e68b4f7b02f1bdffd7"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_c / 054a15057e99 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0264cb4d5486a6268fdeb6f1a3e73b5c633444743b018d3600529345196edeeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-128cb1763f6b3e0c4c70b2b0e2809fecd8946bf34bff69122f9505339b35a3e0"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_c / 3fa158922f14 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-b2c11dd729b04d65fcee8e49e5df1bd6ee901dbbd863821d00342636daa63ff5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card validation.

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
gift_card_validation = {}
```

<a id="canonical-b4649d68e106910e327b54d66f4685cbae211a088fae0676aecac1f61f9e6cf2"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_c / 3fa158922f14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6dadaf3c2423d76dcf5fd6db596eb7ecc07f3a93fdc859844d84e1027c906bb4"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_c / 3fa158922f14 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9373003516bd8427c9a5d161fed8badc826071e7c994c5092ed393bce9e0cc63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-119c7cf30bbfe3afdce95c14d3c0b428877395c839d2d78f2facd6b3f14da940"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_a / 370e65a47f5c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-769d1f9aea93d786af0e332189038167b910ad931838184f63926e08c282fa95"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop add to cart.

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
shop_add_to_cart = {}
```

<a id="canonical-292b4419ad5ac3707f0adcee8c8c51f8043eac5fd9f46c6d2fd7d50366d8dff0"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_a / 370e65a47f5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fb734064205881aedc7496a5fcbc4f4108ac4c1026d2a7848dd2b42d43b0d3c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_a / 370e65a47f5c / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d39d09bbb1c1b39a39adf9afbaf18ea438f326786d91495c5c2f868e6b238138"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f0f9618cda6e34a3a266f6ef66b79793e8e68c196aa459b8d43a9706a963352"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_c / f986fe768d66 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-3d00acbb3328d7d42c6c1cd71812cbcdd5abc745de839bca1ab1249e851ffde1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop checkout.

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
shop_checkout = {}
```

<a id="canonical-57a2a8c68ed46e77c23c17fb03a04ea441ac887cc68df0a92130b21a85256400"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_c / f986fe768d66 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-22e3426d19a0b9cf42d79baa2b5201dbbead6249baac1240b3c9ff3fb7938f61"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_c / f986fe768d66 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8803d586d51b6ff6d1c5841cb4cf35e38afcb36d33c77505b354f90fa277bbf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30952ea75df279de5e73df6e038094e83d38dc39e9829453763322aa1b92748a"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_c / 8998b44efb9b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-32f4f866837a44aa4aaa281d1d77f005063402597df5a9e32f0376569a8aa1b1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop choose seat.

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
shop_choose_seat = {}
```

<a id="canonical-ee9b741d3cf9fd4e30edbc09dd69213d11ee019063dac4f4713ed5c6c72908a9"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_c / 8998b44efb9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f67e175a2fe8bc934378f23b37b9bca1a8d7eca4e6c6ffea8aec0c17e05cf72b"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_c / 8998b44efb9b / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f8e7f944667289751c8cfd0027d5f813e298f29ac858cad687dbf9859fcd8671"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83913535593cee9b3c0efa8d4622953b48437e4dd02fcd73f1fefea2907dbc44"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_e / 0886da30ea5c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-bc81b7f5716788c6be7716991dc4bb521975bdfaf4fd07381df763f1bfc8b098"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop enter drawing submission.

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
shop_enter_drawing_submission = {}
```

<a id="canonical-8af39b58bbbe63d213b3addf5ce8ffd30eac1e7f2c8e523b9e61a0c01454049f"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_e / 0886da30ea5c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f1404bd74f3eab52e83d128266cb49e0c3aa8d5a679da5b5b248569413b1b3c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_e / 0886da30ea5c / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bd149b9e9de8fdb2968a03777b77edaa35cc2089fda45e77f4eb6fc0c579a0c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20df795884aebae59a6c7394e021919db936a90c5ed84d48b603047c4b0f691e"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_m / 48d1e531c38b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-a007f13cec666285d5d0f97ad6cbf18b807e63c14128f0087ea10ab2db69bd39"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop make payment.

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
shop_make_payment = {}
```

<a id="canonical-40a335f3b5e40764879c88b057eba8e4a911300fdd395fb42cc5dc8d335a3e54"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_m / 48d1e531c38b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f225ca754a1d048d7a66fe3a74e97671c48130e0b08f8b8176404c5168e87297"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_m / 48d1e531c38b / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4176c57c3b19ea3887ed4efc0114f14bb58477352bfdfbe2a00c00e44055cbea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cc4465a34d51cc0cb45093a25b3ab7dd2f257dc60acabd960c2fa056faae9e5"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_o / fc0bbeae8217 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-e8ca032e8967eef7cb8ff1aec25228c22d3a23246f4ab7357a455a315247450b"></a>

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
shop_order = {}
```

<a id="canonical-538d5ab36bdffa97252dc56d0e3cc0237d763807245b720900c3885c3f0546a2"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_o / fc0bbeae8217 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3151cd29baa82806b65753ce7ec1bf1c06657f7f6648477e95b0c4a3789ccacf"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_o / fc0bbeae8217 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7afe4a217aad2b2e250b59d0c632d9c99647fc181c2a70671fcb03963b3789d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f63c085c143aa149126bb192823baa2a00aa8664a8350d71796607dd6256e71"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 2e307bd4a65e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-ca16a54bebd77484231f06550d47d0b7bad3a2184d80cd925ed15c7be7bab851"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop price inquiry.

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
shop_price_inquiry = {}
```

<a id="canonical-b8ea912139d7985f26919602d0a16f0647ea19357ae95a6dff43c5683a9896c5"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 2e307bd4a65e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d004eb907235e99479d501dccf6449adb07d34e5ee167c2874dc6526182d48ea"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 2e307bd4a65e / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7d5968b4f523ce74f598a4a7c0af1cfade5212d24e23b9c92adf906470aa3ea5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbd00cdce1063024c3d1fe596d9052883168db873d8cdf907efa3a708d537b6c"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 639360f590ba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-7763154209ed9ab8d1b43da5e3136f31ae0ac6f4d87c1ed7983fb515cf13a369"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop promo code validation.

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
shop_promo_code_validation = {}
```

<a id="canonical-f06220a3cdbcf8c8293a3b402d4e7f443ecbdf9e575917f74bd604a13dc384c6"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 639360f590ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3e882bc9675619c5c2ca39d21e0a26d8b05b1aa965045c6815ed592119af1346"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 639360f590ba / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4e5e30d53fd3b1e0195c2c6b78c041637116f7b25bcbf30b45155e018aa4195b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f78aa94a505ae86d4f70d1a7292d5413402bafe0c1087e366a3131e97241e311"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 1eb9bd39cbf2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-9337ec8ec7f966471548d7b5c6547f5f2b0ebca03e234edd2ec722e0fbd364eb"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop purchase gift card.

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
shop_purchase_gift_card = {}
```

<a id="canonical-311a49b013e10bfce7bc16f70ebc20ccf4371ce0979ed9f6492bed391fe46fd5"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 1eb9bd39cbf2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19c837e7934f2563fe7f329915193a0ea73d1aeea981d04b9de09ea3187aebbc"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_p / 1eb9bd39cbf2 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dca9b83bf456151b022690a3efb5b7e86afe8a379dead28c0768e2f39ad64dad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8712b9cafa0a58a60b48fdd7f9a5aa8c43ade9149e123ced225a45f0df3ab096"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_u / d55f51194e42 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-3490cec2e5c3ffe9565e0862ea9926288cfaf32b972a3b0e8d8480398135ab82)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-27b0ca3e3b85144e0a0dd1f8e62f2d2347a4752fd3091a4ddbe63f979c342603"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop update quantity.

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
shop_update_quantity = {}
```

<a id="canonical-9512c9397df5c4e013e4c4d65d14634d5c72dd07a347164b6a3aea350836ef5c"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_u / d55f51194e42 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-efc2a108fefc8808527271bbf14c1cfe474c2b04dd8fe4f3c4bc4d58991dd337"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_u / d55f51194e42 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-e44dccc6c0c5614f6c31032d1204ce89a5f066aa9c0c0a540d64dbfef2633d28)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3acaad4c3ff013a5a3a10f054c0d3df0ca9e39002990b3d392f6d0ddd19e2c2b"></a>

## bot_defense.policy.protected_app_endpoints.headers — bot_defense.policy.protected_app_endpoints.headers / 243d46cac907 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.headers

<a id="canonical-7eebcb97742518bbae4687c3811bba04478c76f1a5b00cf7af09c11910b3f1e0"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-dcf4bc63921d8d00fb6e3058c7fbfeb0f52280aedbde918a8ab25fae72da9431"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.headers / 243d46cac907 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-ae8041f7ae784bdaaaf14485b14e64a822e6803970fbf1a438b97bf1e0b2a47b): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-012.md#canonical-16b7328396ac8cba435619f657dbc47c390b35edd95af85570f24003c2703219): complete subsection reference.

<a id="canonical-672b1ea524a14410f0c38043e3387eb7aabf91275b5c9644471ea1fc08ef4529"></a>

<a id="canonical-53eb8d9b4ae03685275190315f71d526da5e8be1abd7612536f9f84ebed7b27e"></a>

## invert_matcher property — bot_defense.policy.protected_app_endpoints.headers / 243d46cac907 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-012.md#canonical-366f0ed23d076f54879192325ba7f07cc09f61b8edd44221c3d7c596dec32755): complete subsection reference.

<a id="canonical-7bc92bd18ccf4dae440a99d6a99be27c39b695d9cbd58588930d1e24450e672e"></a>

<a id="canonical-5e61a87d84ef149ce38de2564de94a1da543ba9c45416b7c5b8dc1732eef4e89"></a>

## name property — bot_defense.policy.protected_app_endpoints.headers / 243d46cac907 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2c07424518b94e0c20d5e9a210cf6d1d425804fe2ad407a12dc9ee389d80275d"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.headers / 243d46cac907 / 6

- [bot_defense.policy.protected_app_endpoints.headers.check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-ae8041f7ae784bdaaaf14485b14e64a822e6803970fbf1a438b97bf1e0b2a47b)
- [bot_defense.policy.protected_app_endpoints.headers.check_present](resources--http_loadbalancer--reference--group-012.md#canonical-16b7328396ac8cba435619f657dbc47c390b35edd95af85570f24003c2703219)
- [bot_defense.policy.protected_app_endpoints.headers.item](resources--http_loadbalancer--reference--group-012.md#canonical-366f0ed23d076f54879192325ba7f07cc09f61b8edd44221c3d7c596dec32755)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ae8041f7ae784bdaaaf14485b14e64a822e6803970fbf1a438b97bf1e0b2a47b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000f684d123526bff0b3fd98a0b7b2090491ea5a03e3546e30fe58c7fc978c7"></a>

## bot_defense.policy.protected_app_endpoints.headers.check_not_present — bot_defense.policy.protected_app_endpoints.headers.check_not_present / a5ca75ffdd02 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3)
- bot_defense.policy.protected_app_endpoints.headers.check_not_present

<a id="canonical-d16ad524c5216fc70abde715773c0b597815e2256228c5dfd82e197da9459448"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-72a367da9fbdff5357c8c00accb5d3031fbb0b68eb0454b16fe96e33f437dd09"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.headers.check_not_present / a5ca75ffdd02 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c8f78b6a6c462db43c019c58f43b8e575d8a7e955b7c9b146fc76193fc206daf"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.headers.check_not_present / a5ca75ffdd02 / 4

- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-16b7328396ac8cba435619f657dbc47c390b35edd95af85570f24003c2703219"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15072a8fc2b65678da44df94d53d3f43e2c070e3a79abb7ec1bb9ab96e2bf74a"></a>

## bot_defense.policy.protected_app_endpoints.headers.check_present — bot_defense.policy.protected_app_endpoints.headers.check_present / 7104a335a2bb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3)
- bot_defense.policy.protected_app_endpoints.headers.check_present

<a id="canonical-4dc175940d8cfffec0d245fe405660559f218054306d544794a5285cef0c91a5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-02e62c4f9cfe61e31fecfd644d5b7200da7eba90c3437563c19ae67ce7af0056"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.headers.check_present / 7104a335a2bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39cd06b5c3e9129c46b24f2a6c20f0fd9f0d3eb0586c7037bfdf82697fdde505"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.headers.check_present / 7104a335a2bb / 4

- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-366f0ed23d076f54879192325ba7f07cc09f61b8edd44221c3d7c596dec32755"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-584dc08183a784f65078bf1242284492c47b7edbcc310fd37607c98557651975"></a>

## bot_defense.policy.protected_app_endpoints.headers.item — bot_defense.policy.protected_app_endpoints.headers.item / aac7e5990403 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3)
- bot_defense.policy.protected_app_endpoints.headers.item

<a id="canonical-ab94472b8b70124b1d22dcb1c16ea14f6ccd6d17691ec3d5402847a2238938eb"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-db8bab126f6b67babfc461f8a26a924b79736bb1aba0bec3788e51f787866c98"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.headers.item / aac7e5990403 / 3

<a id="canonical-e785ef750c858303654a67860068d4c52239eb6c6cf952ce6eef90de53bddf07"></a>

<a id="canonical-b7c0653b668df86c4a2f2f1c3e3c46e33a971f35c08faa28346d704a18f63942"></a>

## exact_values property — bot_defense.policy.protected_app_endpoints.headers.item / aac7e5990403 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ca6cdd27ba04f7878b6efdb98b9792d3c4ef2e9a27b734a7d690c0519f574fc8"></a>

<a id="canonical-71aec6770d9169bcd8d558677607e83cd96f3296ef36a9b5781e0e80b2101f28"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.headers.item / aac7e5990403 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6579d84b06fa0e38c40a59bb1af4836a46aeee41885f749a9c3bfcba05e86b09"></a>

<a id="canonical-214bcbd7e51c690f7d71f93982194accd41ff94671317a57ed43276e9e13fd9b"></a>

## transformers property — bot_defense.policy.protected_app_endpoints.headers.item / aac7e5990403 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-edeb21f5bda350517bb3cd3c5981faf3b8346fdecb19a8db8942477e512bb6c5"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.headers.item / aac7e5990403 / 7

- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-c50cfd111a202642cd8a6ca40991957a620e38ffdcbc622e946d7c9252a332a3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8b9bb68551efc78974162af61d69c50411f1fd66c3bbe0f9ba14caa05e79e09a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e8f591640666d49c5a5359f1e7f06c0b7df1a7ffe6a07837a331fc7157f51ed"></a>

## bot_defense.policy.protected_app_endpoints.metadata — bot_defense.policy.protected_app_endpoints.metadata / f589bd329d12 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.metadata

<a id="canonical-d73656aff8c426a48e06ff7696678c7a72bc396b771db6dc1a3448f806665483"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1775590fdb976550134488bc91eb474a5bffc00a0c028ede42f69abdb1a71113"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.metadata / f589bd329d12 / 3

<a id="canonical-1c0169cc9159ae607f26eb5c8deca29b8d8fd79ce862f37ec3f8d185dafc8dec"></a>

<a id="canonical-0e4a4a0085a5bbad0456b32aa7714558f70b17fa1b7550c15d3b050e39ce222e"></a>

## description_spec property — bot_defense.policy.protected_app_endpoints.metadata / f589bd329d12 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-46ce474fce3e2ba0fb11c15bf1017d2d020a670dc3fc79e69aa309ee1139255b"></a>

<a id="canonical-6dcca56f39424c9ddee712773f9bd50a5214473ddc791f3f3ed475c7e2bc6c38"></a>

## name property — bot_defense.policy.protected_app_endpoints.metadata / f589bd329d12 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-469e434659f0399259da6de9555938a14f71dad97b0fd9effd99069696fd2e43"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.metadata / f589bd329d12 / 6

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e0f93575fd4f9e09f02b9583a890ec54d9190308d4e103d4a770ed9bb6b8afca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19e6a9a8696295005e72151ff04dada48d8a835ebde90c3c518462b6d07b1cc9"></a>

## bot_defense.policy.protected_app_endpoints.mitigate_good_bots — bot_defense.policy.protected_app_endpoints.mitigate_good_bots / ee982481fe07 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.mitigate_good_bots

<a id="canonical-c053c828dbd4cd75a4459e1cb33e1c5bd20eeb4b097176cb0cc74083f7f9710b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for mitigate good bots.

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
mitigate_good_bots = {}
```

<a id="canonical-ee9399addaf881194bd9971d2dc43c235ee48f776cc5129b8d2360ccc35be8f8"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigate_good_bots / ee982481fe07 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0f06e16ba47bb9ac0861fbb0ef9cf3f91cc4bb12fe428585d1ff8e421599468"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigate_good_bots / ee982481fe07 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d93747ba5338a9c388b613bbbb5f8e070f122671985219f6afc588cd10a2623"></a>

## bot_defense.policy.protected_app_endpoints.mitigation — bot_defense.policy.protected_app_endpoints.mitigation / d8ad7f60e6b9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="canonical-9cbbce0d96e3442c8250b117c6ad66822e798075b9db37b51f144b89ef620e56"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot Defense behavior for a matching request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "flag"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("flag",
    "redirect")}
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
  "x-ves-oneof-field-action_type": "[\"block\",\"flag\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-39990db5254a75f1fa45e5cf41a5528ac1a7e4f2259dcad0e19d1d4e9bdd7a83"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigation / d8ad7f60e6b9 / 3

- [block](resources--http_loadbalancer--reference--group-012.md#canonical-545d27365a39313d3787ca67300b95d2228c2e68a4c86c4fbfeb047d1ecc2d6c): complete subsection reference.

- [flag](resources--http_loadbalancer--reference--group-012.md#canonical-8f3db5ab3fe2f3cdae3e20dfd154644409c0eb2f715016c6e13333fcbb7890fc): complete subsection reference.

- [redirect](resources--http_loadbalancer--reference--group-012.md#canonical-e20a43c7a00610efe64ae3f79b398569502c2515f1f6b063ab4a92a72e1dba25): complete subsection reference.

<a id="canonical-78cf99cba502b90d5fd26520e26562da6efa5d7288c6ea30927f07236a23bd22"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation / d8ad7f60e6b9 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.block](resources--http_loadbalancer--reference--group-012.md#canonical-545d27365a39313d3787ca67300b95d2228c2e68a4c86c4fbfeb047d1ecc2d6c)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-012.md#canonical-8f3db5ab3fe2f3cdae3e20dfd154644409c0eb2f715016c6e13333fcbb7890fc)
- [bot_defense.policy.protected_app_endpoints.mitigation.redirect](resources--http_loadbalancer--reference--group-012.md#canonical-e20a43c7a00610efe64ae3f79b398569502c2515f1f6b063ab4a92a72e1dba25)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-545d27365a39313d3787ca67300b95d2228c2e68a4c86c4fbfeb047d1ecc2d6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f1595a249861486dfcbeb9db8ef1167d33585dc212777dc936e99ce0b2acb16"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.block — bot_defense.policy.protected_app_endpoints.mitigation.block / 77a7a6649d5d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- bot_defense.policy.protected_app_endpoints.mitigation.block

<a id="canonical-1d4236490a73f5e6167e6293ab60c6b39c6e017f7c8402153c667e14cdd02b3c"></a>

Type: `"object"`. single nested block, Optional.

Block request and respond with custom content.

Receipt-pinned upstream constraints:

```json
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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-e23768a5d80e0787a59fb151c22ba4fbb5f71ef592ce8141c89c6d66186597b3"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigation.block / 77a7a6649d5d / 3

<a id="canonical-83a6e845c947e0869dc6e8a18c43b60ec49fb950ad7f1b3b357b346daf25136b"></a>

<a id="canonical-af892f979daff81ddd74be2e6ba6bd6a660dd81cffdccb78e870eed386b5151f"></a>

## body property — bot_defense.policy.protected_app_endpoints.mitigation.block / 77a7a6649d5d / 4

Type: `"string"`. Optional.

Custom body message is of type uri\_ref. Currently supported URL schemes is string:///. For
string:/// scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom body message is of type uri\_ref. Currently supported URL schemes is string:///. For
string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as
base64 encoded plain text message e.g. "Your request was blocked" or it can be HTML paragraph or a
body string encoded as base64 string E.g. "&lt;p&gt; Your request was blocked &lt;/p&gt;". Base64
encoded string for this HTML is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0f384faf8acb73ba45cc0d953093b665fddcbc439ffbe72ac1cff8c7a7592d3e"></a>

<a id="canonical-6c334dff7d7c80853e43f82c51ee20a06397912cd478a31b7dc5082bc94a0c7d"></a>

## status property — bot_defense.policy.protected_app_endpoints.mitigation.block / 77a7a6649d5d / 5

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-88966d4a68ca31df6d03d42c5fc1d8476c7fa613dbd48c66441532ac9bc9e4e8"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation.block / 77a7a6649d5d / 6

- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8f3db5ab3fe2f3cdae3e20dfd154644409c0eb2f715016c6e13333fcbb7890fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-403955a0e50dbb7e05e5bd32ee115b785c0e41062daae320b7b2f344e2644bde"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag — bot_defense.policy.protected_app_endpoints.mitigation.flag / b67823365de0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="canonical-55863d462822163024f665ecfb9593e5840d7fd6bd53d8a76debedaaeb82e5c3"></a>

Type: `"object"`. single nested block, Optional.

Select Flag Bot Mitigation Action. Flag mitigation action.

Upstream description:

Flag mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_headers",
    "no_headers")}
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
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

Terraform syntax:

```terraform
flag {
  # Configure direct properties listed below.
}
```

<a id="canonical-dee304cf8a6e68614c7f4ba904d42ea077dbba88a216aa8ec61d2ff67f95ffe7"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigation.flag / b67823365de0 / 3

- [append_headers](resources--http_loadbalancer--reference--group-012.md#canonical-67da9a678d0aa8e28996be5160387dd4220e7e5180feee3b1eb435614a13c2d4): complete subsection reference.

- [no_headers](resources--http_loadbalancer--reference--group-012.md#canonical-25c500b582a71dee7961ace71078ef78b2c8a0b56beb9c6591d335ccfb9a24aa): complete subsection reference.

<a id="canonical-ac2cf8b9010426635b1648f592f36ae80387a0211cc661f4b41409fad6a48ec9"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation.flag / b67823365de0 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers](resources--http_loadbalancer--reference--group-012.md#canonical-67da9a678d0aa8e28996be5160387dd4220e7e5180feee3b1eb435614a13c2d4)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers](resources--http_loadbalancer--reference--group-012.md#canonical-25c500b582a71dee7961ace71078ef78b2c8a0b56beb9c6591d335ccfb9a24aa)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-67da9a678d0aa8e28996be5160387dd4220e7e5180feee3b1eb435614a13c2d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2562719af59be86980a319066b84dc21657137b6bdb131cdef50a9a6e0c59155"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers — bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers / a0b757ee747b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-012.md#canonical-8f3db5ab3fe2f3cdae3e20dfd154644409c0eb2f715016c6e13333fcbb7890fc)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers

<a id="canonical-789b96b33c2a236f48a5e18e45b6883ab65a444d0687b86678da8281c123a7c8"></a>

Type: `"object"`. single nested block, Optional.

Append flag mitigation headers to forwarded request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("auto_type_header_name",
    "inference_header_name")}
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
append_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-6993676e5ac38d69e6162d1274ee39f6541088d9de1af218656f1732af6aeb36"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers / a0b757ee747b / 3

<a id="canonical-f05a2255658ac54859456e8bdbe04d34c6cd6c49aa79e5a90dbf99c2fd51d6a7"></a>

<a id="canonical-52da9a657b9ee660d22df1e50508d3bd841c60dac3df56c959e10c3709aa2a7c"></a>

## auto_type_header_name property — bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers / a0b757ee747b / 4

Type: `"string"`. Optional.

Automation Type Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "byteLength": {
      "max": 256
    },
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-8d76fb6566d76b986e62f06fe06f6cc8e578512edad8115899c865eac6bff644"></a>

<a id="canonical-4cba04ad23f7ec522d73388e0c6f8dfbef64cecb7c92e46877162118e8e4841c"></a>

## inference_header_name property — bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers / a0b757ee747b / 5

Type: `"string"`. Optional.

Inference Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "byteLength": {
      "max": 256
    },
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-fb2ad45428d5ed19f74ef39a8ee5962d8f80cb358ad9fa9df55671c4f54be4be"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers / a0b757ee747b / 6

- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-012.md#canonical-8f3db5ab3fe2f3cdae3e20dfd154644409c0eb2f715016c6e13333fcbb7890fc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-25c500b582a71dee7961ace71078ef78b2c8a0b56beb9c6591d335ccfb9a24aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eddbc19d7e407caa083cfea511c5351413340b0b2486f3978f17c519793923ca"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers — bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers / 0325d5ec7c1e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-012.md#canonical-8f3db5ab3fe2f3cdae3e20dfd154644409c0eb2f715016c6e13333fcbb7890fc)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers

<a id="canonical-718433e1f1da048f054b7e6b1d863429d2d91aac56dc8d6ed05d7ad1cedb1ca8"></a>

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
no_headers = {}
```

<a id="canonical-6fa9fe21d3741b7a343ed0e9ea1a9ad9bd76fd5f7d50cd6bae54c5f4d7960649"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers / 0325d5ec7c1e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a311cf78dd33b91ce4ea4c9a43cf6739fc30d653baea8b82937c0a3c7c0cae4"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers / 0325d5ec7c1e / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-012.md#canonical-8f3db5ab3fe2f3cdae3e20dfd154644409c0eb2f715016c6e13333fcbb7890fc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e20a43c7a00610efe64ae3f79b398569502c2515f1f6b063ab4a92a72e1dba25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7d48b19935855962d541811babe506f70b52c59c177f1ee0f390e7057dba761"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.redirect — bot_defense.policy.protected_app_endpoints.mitigation.redirect / f5a60bc8c7d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- bot_defense.policy.protected_app_endpoints.mitigation.redirect

<a id="canonical-94f70ac0adb2be1c528fdc7c9f075a70173f2ae2a22e2daf426028feec5f11ca"></a>

Type: `"object"`. single nested block, Optional.

Redirect bot mitigation. Redirect request to a custom URI.

Upstream description:

Redirect request to a custom URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri")}
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
redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-607e06da9a579dd1bbf5442e1992d457d770e75505907fabd66b1f9e90d6163e"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigation.redirect / f5a60bc8c7d9 / 3

<a id="canonical-349b4e731a77e0dc62c664c4b50f2fbe08c93b89ddb78ef187472280b07a2a9f"></a>

<a id="canonical-8b825b1e50fcf0487a37303dea8b533d5bdd6d64256a8962d7d86f8f64eb3c1c"></a>

## uri property — bot_defense.policy.protected_app_endpoints.mitigation.redirect / f5a60bc8c7d9 / 4

Type: `"string"`. Optional.

URI location for redirect may be relative or absolute.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-315308f1a104fadcc11bc1a12fbeacc49e418a13c6e859d703efdcfb989a188f"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation.redirect / f5a60bc8c7d9 / 5

- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-ea61011b5b44901d7f05b8d65cfa06ec62f1c302f3b1e7c4e32d6f8191aceffd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-20d0718577d0874b23e50c5f919b2484f8b798aecd77321c1833315668f5d538"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbbe7d70dde24619450b3642e76c88df8d6b39a0883c9c86e157bca4a16e4d5d"></a>

## bot_defense.policy.protected_app_endpoints.mobile — bot_defense.policy.protected_app_endpoints.mobile / cccce7cf45be / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.mobile

<a id="canonical-f7ccef60eec09272e18a46ca9b6855a607532aeef21888fe3e60738a10554258"></a>

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
mobile = {}
```

<a id="canonical-89b945a5dd5c8ecee7ee96b19f3b9d944fc3303ce8991e6061189f703564fe8c"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mobile / cccce7cf45be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4247066f78c490d04f73e01ab0854aa6d2ff5750d2f827a64a8d39b2f6e433c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mobile / cccce7cf45be / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-de209870a7c605b348e30eacce83fff5f5442f9b4484b18a1aa1170b78643f4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07b521e2fe0ed7371a01abd41bcd19259b5c34ee505d4a68f1734ca7dab75c4e"></a>

## bot_defense.policy.protected_app_endpoints.path — bot_defense.policy.protected_app_endpoints.path / 24eea3e97448 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.path

<a id="canonical-924bd87ecdd28d7424d3db339de9dabc4e9fe8754d9c8d1c41b7d65d9f53bc83"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3d7b98eb50cde54479dfc752370b4ab1237ac109aa97e32b6716c25c4df59926"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.path / 24eea3e97448 / 3

<a id="canonical-2d63ab8f6af319d4b38fabac6cfb50213bd969eda800f28540c3cc94318ef8f2"></a>

<a id="canonical-194f31bafd88800ab8180a13817a52b7433116e004f88f4c4d09279e7b7004d5"></a>

## path property — bot_defense.policy.protected_app_endpoints.path / 24eea3e97448 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-24e2c72a9e1e32fc6e3ee9f808679b97bf1cfcd6b3114b5abe0f07a307af3304"></a>

<a id="canonical-654964e18764393fda2b7f16d8ceed7a08ec19683fcfa8a652549b0a551c6b65"></a>

## prefix property — bot_defense.policy.protected_app_endpoints.path / 24eea3e97448 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-80eee0d3736813919440519b1412989aa52c22e193fa73731da11149317e47de"></a>

<a id="canonical-39e6d55e1fa6cdc5df469c26933ad79ef4f56b183fd21a5907b4bb5f90110237"></a>

## regex property — bot_defense.policy.protected_app_endpoints.path / 24eea3e97448 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-9a4ed0b0e0f025a7b5b399baad3744b4a5f7aaf9dc5590c8039fe7166fd8134c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.path / 24eea3e97448 / 7

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c70318e008ef4bdefe3a4e9edcb1d785fc8518bc2e25e6751f2804d39c6f0e5"></a>

## bot_defense.policy.protected_app_endpoints.query_params — bot_defense.policy.protected_app_endpoints.query_params / 2cf8de3ba682 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.query_params

<a id="canonical-05671a083a5f9c99afa78bc50fa58a3ee0dd55b3d27bdb9420110f0b5b8634de"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-820dfce921ac67a0d5bb4787fc0a92db5f228e775bef4088a9d2e8f95fbfe649"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params / 2cf8de3ba682 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-91fbaa4c901aa1e157dde0c27d5eac6046fa18669c2494a24ee7edbac624d5ab): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-012.md#canonical-b3127519179a81db1c12dc571eeb1fef26320f7766c5f6c09a1edf9e088bcb01): complete subsection reference.

<a id="canonical-683db4d4f813337761e46dbad5f225b38ec52ba5bca9816c1952a6faa4d039da"></a>

<a id="canonical-81f2d486cd5fef2c84a0b415fe9abcb4f10619659530583eccea5b7b96fb03c3"></a>

## invert_matcher property — bot_defense.policy.protected_app_endpoints.query_params / 2cf8de3ba682 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-012.md#canonical-c75efc456cebabc28fe251e2f80c43e2c9ba0740f493e377595ec30fdc34bc0b): complete subsection reference.

<a id="canonical-9a0f1414c78a9e9fbdf8ed381da4d060bb47ab91d094d155b7fd6e027e27ead0"></a>

<a id="canonical-e34449bf50192f62ff3f3f342743567056096dfd10d8b56c4bdd026dc74aa032"></a>

## key property — bot_defense.policy.protected_app_endpoints.query_params / 2cf8de3ba682 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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
    "byteLength": {
      "max": 256
    },
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-c6ad9eaa4b382365e2f8e418c776876c68bae6c1193b27e3b633bff5d3501e4a"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params / 2cf8de3ba682 / 6

- [bot_defense.policy.protected_app_endpoints.query_params.check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-91fbaa4c901aa1e157dde0c27d5eac6046fa18669c2494a24ee7edbac624d5ab)
- [bot_defense.policy.protected_app_endpoints.query_params.check_present](resources--http_loadbalancer--reference--group-012.md#canonical-b3127519179a81db1c12dc571eeb1fef26320f7766c5f6c09a1edf9e088bcb01)
- [bot_defense.policy.protected_app_endpoints.query_params.item](resources--http_loadbalancer--reference--group-012.md#canonical-c75efc456cebabc28fe251e2f80c43e2c9ba0740f493e377595ec30fdc34bc0b)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-91fbaa4c901aa1e157dde0c27d5eac6046fa18669c2494a24ee7edbac624d5ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5db37aa69488f95f2ea4047ee530f03c47647524782c06b0cbe48a07975ff12"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_not_present — bot_defense.policy.protected_app_endpoints.query_params.check_not_present / e9690cc17666 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e)
- bot_defense.policy.protected_app_endpoints.query_params.check_not_present

<a id="canonical-5e448294591f1521bae8a78187171ef6c9e1d8d14c1034ffa50d1d527f7e8d45"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-7a28a3720e83b5484d1f3ea65d8d6aa44be12a403aaed8bd752c7c489a375f80"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params.check_not_present / e9690cc17666 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1085c2a8c49e5cc5ae8e39d7a2fe6f453408d459114a586311299f23b16975ed"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params.check_not_present / e9690cc17666 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b3127519179a81db1c12dc571eeb1fef26320f7766c5f6c09a1edf9e088bcb01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-787a610ebb617a43adbc70ee07d9f9ad21124b7897ef41f8415b0e941ec0da9a"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_present — bot_defense.policy.protected_app_endpoints.query_params.check_present / d76aff102f85 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e)
- bot_defense.policy.protected_app_endpoints.query_params.check_present

<a id="canonical-963f295da6957d47f846b3e24744732e35e3d7356cad00305b95e7be6c1e04f7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-c99feba9e6c6d8b2f144ee49a3b5ce1fabfff3ad911a8c6d60156675262f40cd"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params.check_present / d76aff102f85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a28916206468d466a552971875afa7728890c30b0a513cefcdcbf52723447be"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params.check_present / d76aff102f85 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c75efc456cebabc28fe251e2f80c43e2c9ba0740f493e377595ec30fdc34bc0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8690831c61cd070a4779e5f44d0428d48c139d3d0146d935a5fe62b49b172a6a"></a>

## bot_defense.policy.protected_app_endpoints.query_params.item — bot_defense.policy.protected_app_endpoints.query_params.item / 0e3f69493448 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e)
- bot_defense.policy.protected_app_endpoints.query_params.item

<a id="canonical-07c40e16dd8c99b3a93dec650ff7bba1913f44278f85c8f6d140cec2911e72eb"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-5b24794996f4c2a6628032c2ad390c96c21846953a146492d72d45db53f4060c"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params.item / 0e3f69493448 / 3

<a id="canonical-46627d3e68031e1313eb93cd3a9e57033349efe7248a69317219362d470e0091"></a>

<a id="canonical-94494e913d389188f611e8dd4b957c15445a0847b8b3eda8b13cb0e8a095051b"></a>

## exact_values property — bot_defense.policy.protected_app_endpoints.query_params.item / 0e3f69493448 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-109541e9c16cc42320ed45c628cb031e5e7e41ec817844c3d59276b17ce75b6d"></a>

<a id="canonical-71cf8f018d25a73164021a28243dd075410da010a08f19f9a553c0f56f4de327"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.query_params.item / 0e3f69493448 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1c6392a64ab3fb797c2a73142a8ad0553cb789f8b08da8acc976e377153e27e8"></a>

<a id="canonical-cff280177d2841363b760b2a67e77fe018ebb7c3660284d2dbb1cf3faecace77"></a>

## transformers property — bot_defense.policy.protected_app_endpoints.query_params.item / 0e3f69493448 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4a662c2a790b7b8a7254ba96c5e040b4540cc7ac570c6c03745c5812cb245d71"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params.item / 0e3f69493448 / 7

- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-012.md#canonical-47b8ba621183a1aa542dafd9635d726b2f9d6fe73de146e5d57618e8f7e16b5e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-63939a980f6231f79d98a18e6fe71dfcd8d673eef2c80a60cbc851752e35c7ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-929f9ae0e16a6f68e1076c773d8df5ad16b2bf83d6e21c878a08946babe2bed3"></a>

## bot_defense.policy.protected_app_endpoints.undefined_flow_label — bot_defense.policy.protected_app_endpoints.undefined_flow_label / b2a00a051418 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.undefined_flow_label

<a id="canonical-edff40ac95e2cca8c3b0d7d82fc7eb46236e585e01d8738b9b521ac5f3b01845"></a>

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
undefined_flow_label = {}
```

<a id="canonical-a6075c8bf4e8d0bc1ea807d0262314f9fcd0b3444246ddf5cc03812e3b65e7a5"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.undefined_flow_label / b2a00a051418 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f3017eabcb3d7fee2b259a2fd23b1ef01b0a4aa2d69aa72fd9c4d4f3ef7bcdf"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.undefined_flow_label / b2a00a051418 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1c21b5e02cda152a7fc62c5cf2972c46ada5da9df9ee1c3f1142fd678808b4f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b45682a9fe0773150ab9236f49615981995f236b9e894efeaed71fab8bc07e9"></a>

## bot_defense.policy.protected_app_endpoints.web — bot_defense.policy.protected_app_endpoints.web / c31b11970df2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.web

<a id="canonical-fe3d7fdd1210ddd483c63aae27f92bdf338794fc8783fd34ebc69ebd2459ae2e"></a>

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
web = {}
```

<a id="canonical-8b15d529291e849000b7d21e927a0e61e99af658ab5bbe33a12490a8f5b52732"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.web / c31b11970df2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a5693b21111bb93f913134af912a0e2955a6ea3e8bc40b3514013c95e6d60e3"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.web / c31b11970df2 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d1dd9a62713870cb68c6af01a169c34eb791ba8dcbd8fba2919ee95722f29d72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b060e65f72ccc1114eecd2bc9ee8485dc0718f79f7cfb54bf8e9e0f1a16c1c85"></a>

## bot_defense.policy.protected_app_endpoints.web_mobile — bot_defense.policy.protected_app_endpoints.web_mobile / 6dab5389e7a3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-1ceefcd7848fe72af6e57cd67c1eb345d75bbcf89c9565f6706954be2fbb26cc)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-c3491869a8602d9e75d608a96a08db159a549313d547c3e9a32375ad76dda047)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="canonical-b4f5ed6fd9ac30b14c126175f9dddf18bf372c483d4469dfbdac512d5dd22b62"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile traffic type. Web and Mobile traffic type.

Upstream description:

Web and Mobile traffic type.

Receipt-pinned upstream constraints:

```json
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
web_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-053bb54c6f5d1e37f69512286d484901a6cdc2f9e7323de32ca63895e445946b"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.web_mobile / 6dab5389e7a3 / 3

<a id="canonical-c24f8df7c0e4f4d6def42ecb1a1fe7d727165f6137ddd7c381fbf2d307e17528"></a>

<a id="canonical-40756b077273215ce79ca315566356a74bf15e90cba3fd3753a684e82b8f0470"></a>

## mobile_identifier property — bot_defense.policy.protected_app_endpoints.web_mobile / 6dab5389e7a3 / 4

Type: `"string"`. Optional.

\[Enum: HEADERS\] Mobile identifier type - HEADERS: Headers Headers. The only possible value is
\`HEADERS\`. Defaults to \`HEADERS\`.

Upstream description:

Mobile identifier type

&#8203;- HEADERS: Headers

Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HEADERS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HEADERS",
  "enum": [
    "HEADERS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-755a6cc4f92623f0ee304cf3419b1ba19cafb9dc3c21f3dc6a2e2897e2c5531c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.web_mobile / 6dab5389e7a3 / 5

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-70a6db04e5e846a171922dcaeb05302300103195acdeae11076bed313abfa5c9)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b64c05612002f941244db509c6a358307202dae3bfdc6df7c422740e6b6a541"></a>

## bot_defense_advanced_protection — bot_defense_advanced_protection / 0ef559d6ef97 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- bot_defense_advanced_protection

<a id="canonical-7d2fbebc18d0457e5b88baf2edf2236a96be148f95e0b8f95831cb57e63378ca"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Advanced Protection - replaces BotDefenseAdvancedType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("both_web_and_mobile",
    "mobile_only"),
  validators.ConflictingObjectAttributes("both_web_and_mobile",
    "web_only"),
  validators.ConflictingObjectAttributes("mobile_only",
    "web_only")}
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
  "x-ves-oneof-field-client_type_choice": "[\"both_web_and_mobile\",\"mobile_only\",\"web_only\"]"
}
```

Terraform syntax:

```terraform
bot_defense_advanced_protection {
  # Configure direct properties listed below.
}
```

<a id="canonical-213de63d5044c4b4c2d22be1f063032b7ce95d1f14fa7671e45aebaa7b36a754"></a>

## Direct properties — bot_defense_advanced_protection / 0ef559d6ef97 / 3

- [both_web_and_mobile](resources--http_loadbalancer--reference--group-012.md#canonical-25ec8053d1533fd105b80fbd19384e27907467ca7f8e9b2a3909e8a8af1f206f): complete subsection reference.

- [mobile_only](resources--http_loadbalancer--reference--group-013.md#canonical-77a5406e4055d50198fe268911f96875f3b0428b6a4d306440915cfb83362d8f): complete subsection reference.

- [web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626): complete subsection reference.

<a id="canonical-d6f88ab16f2ca48e0a1d4d263f1f7568f0cb93c15dea1731c64184ff02297645"></a>

## Next pages — bot_defense_advanced_protection / 0ef559d6ef97 / 4

- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-012.md#canonical-25ec8053d1533fd105b80fbd19384e27907467ca7f8e9b2a3909e8a8af1f206f)
- [bot_defense_advanced_protection.mobile_only](resources--http_loadbalancer--reference--group-013.md#canonical-77a5406e4055d50198fe268911f96875f3b0428b6a4d306440915cfb83362d8f)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-25ec8053d1533fd105b80fbd19384e27907467ca7f8e9b2a3909e8a8af1f206f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3ba156aaa7a7929598331f640ce4d4a1d48020601a7111f43185825f63875e7"></a>

## bot_defense_advanced_protection.both_web_and_mobile — bot_defense_advanced_protection.both_web_and_mobile / 64af82c3be1a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- bot_defense_advanced_protection.both_web_and_mobile

<a id="canonical-0554112acd001b053d115eed5edb87e1e030395c9f0d9292b73da7c84daa218b"></a>

Type: `"object"`. single nested block, Optional.

Both Web &amp; Mobile. Both Web and Mobile configuration.

Upstream description:

Both Web and Mobile configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
both_web_and_mobile {
  # Configure direct properties listed below.
}
```
