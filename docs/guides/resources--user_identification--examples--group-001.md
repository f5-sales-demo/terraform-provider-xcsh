---
page_title: "xcsh_user_identification examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification examples."
---

# xcsh_user_identification examples

<a id="canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f340070f752133495ccbecb37d3a5b822dadd59c925b3be62786ae84a988ff4e"></a>

## Examples — Examples / 3973e55315f7 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- Examples

<a id="canonical-dd9758e37976fc1999d09de0ef4106d77506546a4176669f173006a5c9f93840"></a>

## Complete configurations — Examples / 3973e55315f7 / 3

- [All attributes](resources--user_identification--examples--group-001.md#canonical-de997fe2d0b68d410d9594e75d8aca6fc02ae504d5af9c58775fd960954f2851): valid configuration.

- [Cookie](resources--user_identification--examples--group-001.md#canonical-275f91e7e468c28f35f77de040bda17387669b6c2f3b0b6aea21a5f9bcb47b4f): valid configuration.

- [Http header](resources--user_identification--examples--group-001.md#canonical-53306daea91f0910660ae0d696f80bfaaaff7174711013cd9f01008dd637322e): valid configuration.

- [Resource](resources--user_identification--examples--group-001.md#canonical-0930fd0bfb65de18bdae2e19550dbadba7ae70b9190e8ce7d00701019a02bc72): valid configuration.

- [Tls fingerprint](resources--user_identification--examples--group-001.md#canonical-3da1b439ae09801f002b3b5ff2e6f62bbd96d4ad6efd27ccf0c3a0659dc12144): valid configuration.

- [With annotations](resources--user_identification--examples--group-001.md#canonical-f542cee90540e784394c41758440ecd71571a46ed9333845ab7b652f254e62bb): valid configuration.

- [With description](resources--user_identification--examples--group-001.md#canonical-ece20695ce72ef74977251dd4f1dad2de5de03e7502e9a9dcc4a6968071e329b): valid configuration.

- [With labels](resources--user_identification--examples--group-001.md#canonical-d06f119b32bc8a07bcfd7a59a877af8a304de5cd3243e5371418a7fe6905dead): valid configuration.

- [With rules](resources--user_identification--examples--group-001.md#canonical-1677b6f0fa9e374e802866260e51b1067cee7a684fdd0b5aedaa66cd3dafdf4a): valid configuration.

<a id="canonical-3e274041e8f054bc3a2e72acf4bedfe30b7166c93c2ee77a91ee33d815fac46b"></a>

## Next pages — Examples / 3973e55315f7 / 4

- [All attributes](resources--user_identification--examples--group-001.md#canonical-de997fe2d0b68d410d9594e75d8aca6fc02ae504d5af9c58775fd960954f2851)
- [Cookie](resources--user_identification--examples--group-001.md#canonical-275f91e7e468c28f35f77de040bda17387669b6c2f3b0b6aea21a5f9bcb47b4f)
- [Http header](resources--user_identification--examples--group-001.md#canonical-53306daea91f0910660ae0d696f80bfaaaff7174711013cd9f01008dd637322e)
- [Resource](resources--user_identification--examples--group-001.md#canonical-0930fd0bfb65de18bdae2e19550dbadba7ae70b9190e8ce7d00701019a02bc72)
- [Tls fingerprint](resources--user_identification--examples--group-001.md#canonical-3da1b439ae09801f002b3b5ff2e6f62bbd96d4ad6efd27ccf0c3a0659dc12144)
- [With annotations](resources--user_identification--examples--group-001.md#canonical-f542cee90540e784394c41758440ecd71571a46ed9333845ab7b652f254e62bb)
- [With description](resources--user_identification--examples--group-001.md#canonical-ece20695ce72ef74977251dd4f1dad2de5de03e7502e9a9dcc4a6968071e329b)
- [With labels](resources--user_identification--examples--group-001.md#canonical-d06f119b32bc8a07bcfd7a59a877af8a304de5cd3243e5371418a7fe6905dead)
- [With rules](resources--user_identification--examples--group-001.md#canonical-1677b6f0fa9e374e802866260e51b1067cee7a684fdd0b5aedaa66cd3dafdf4a)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-de997fe2d0b68d410d9594e75d8aca6fc02ae504d5af9c58775fd960954f2851"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72194468747b0235858088d5325dc0642a7d9122c660cd76f9699468dfb0a0b8"></a>

## All attributes — All attributes / 0b12d26a03be / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/all-attributes.tf`; digest `sha256:678a876b16600386f26af7457406ace3f301d5ca150aba0447d4b245d2b5a900`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test user identification with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "security"
  }

  annotations = {
    purpose = "testing"
  }

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-48fc8d743f596c0784b775b8d5ced9eed3eec931a2ba07f142dd423fa34c5a63"></a>

## Next pages — All attributes / 0b12d26a03be / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-275f91e7e468c28f35f77de040bda17387669b6c2f3b0b6aea21a5f9bcb47b4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bb2bcea9e0e617ad56e40c5df6c7c5f8b101af6f66cd879d6c278ad9b6d9103"></a>

## Cookie — Cookie / bc01cfc93a68 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- Cookie

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/cookie.tf`; digest `sha256:5603d5f0bc4ec0397625fee3d3eb29ea6bfe17b42e387a451f62ff082893af26`.

```terraform
# Cookie — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    cookie_name = "session_id"
  }
}
```

<a id="canonical-e7aceab7609c3e30bddb76ad5a558603ae5ea24828e04c9ddab22a933e49cc57"></a>

## Next pages — Cookie / bc01cfc93a68 / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-53306daea91f0910660ae0d696f80bfaaaff7174711013cd9f01008dd637322e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35c2097ebc1281b186974378597ade01846ae2314f4d2efad8a604920c4291e5"></a>

## Http header — Http header / 969892c19480 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- Http header

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/http-header.tf`; digest `sha256:b19736efd2e6603701aad68c2c3118fc31fc2264595ed7e6f49cbd6de0dc5c92`.

```terraform
# HttpHeader — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    http_header_name = "X-Forwarded-For"
  }
}
```

<a id="canonical-efb235dcc58456a30cc9312dfde75003e01a556b60450949a6540f58689a70cf"></a>

## Next pages — Http header / 969892c19480 / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-0930fd0bfb65de18bdae2e19550dbadba7ae70b9190e8ce7d00701019a02bc72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbcb51e83e42c3ee4106df0037feb5ecca70834378a11526497cce94f53fe383"></a>

## Resource — Resource / 1fa1490ea56a / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/resource.tf`; digest `sha256:a74347515a80c8e0aaaac40831664823cc6a3928018d3932b931b768606d6cb2`.

```terraform
# UserIdentification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UserIdentification configuration
resource "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}
```

<a id="canonical-699b31d5d64256bd01bf8b3192447e84e52c971b7ddc0c45d89c0f72d77be95a"></a>

## Next pages — Resource / 1fa1490ea56a / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-3da1b439ae09801f002b3b5ff2e6f62bbd96d4ad6efd27ccf0c3a0659dc12144"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f4045942ea776bc330d0811e57ace3b39bb2d5413e3ed1cd3f3a3924e9c2eec"></a>

## Tls fingerprint — Tls fingerprint / a7a5392bca84 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- Tls fingerprint

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/tls-fingerprint.tf`; digest `sha256:18aace91bd93f5201f1f6ed23d33d48d6c8f66a90c3a44614cea4e6f6f5a4a09`.

```terraform
# TlsFingerprint — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    tls_fingerprint = {}
  }
}
```

<a id="canonical-93b5b80e7f8e28d6c33ab11be1e3ae84ed3dd6230d1c20c35d5890647bdc317f"></a>

## Next pages — Tls fingerprint / a7a5392bca84 / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-f542cee90540e784394c41758440ecd71571a46ed9333845ab7b652f254e62bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-829acae7e3d552f435bf4bb522423dfe7a389d9815838d207534ba416a4a9744"></a>

## With annotations — With annotations / f28cfb744709 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-annotations.tf`; digest `sha256:903200cccd3aa8c5abe2eebfb30b2e2efe55bae4a37636c3dae4de46ed18a25f`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  annotations = {
    example-key = "example-value"
  }

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-7af0cf935c0c5ae99d26439dad87bfe57ff56a6bee2dfb8f7c718f5fec913ff8"></a>

## Next pages — With annotations / f28cfb744709 / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-ece20695ce72ef74977251dd4f1dad2de5de03e7502e9a9dcc4a6968071e329b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6877e01a299be7f53e4ca9db628dcc9022c23d8f2264fa88bcfac687207077b"></a>

## With description — With description / 9b9b4d48e047 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-description.tf`; digest `sha256:28e44c81af88136d36f13495510529b754168c2724c3ce7cec3eeb9297eaa127`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name        = "example"
  namespace   = "system"
  description = "example-value"

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-89cd4836e30b7f93b81f001bea9b00afa20c672560bb90fe942b3910b3e3cb5c"></a>

## Next pages — With description / 9b9b4d48e047 / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-d06f119b32bc8a07bcfd7a59a877af8a304de5cd3243e5371418a7fe6905dead"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4da4e4728a99eb14341a292d722599923b46bf7afb492cd398c532354213272d"></a>

## With labels — With labels / 9be8810d2d33 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-labels.tf`; digest `sha256:ba484df9375d0da693ea56c4aaf4288ad5030a88885849721fac1dec608ba5a3`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    example-key = "example-value"
  }

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-b2086908e2ae3b04b260b9865d956df94bb0b31601853f8626d65cc2e896c96c"></a>

## Next pages — With labels / 9be8810d2d33 / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-1677b6f0fa9e374e802866260e51b1067cee7a684fdd0b5aedaa66cd3dafdf4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13954c576c229d46cbe6c240ecc4a4b644641fd10d8c811f9862e7d3954026b8"></a>

## With rules — With rules / edcb489e74bc / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- With rules

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-rules.tf`; digest `sha256:b71ec64258f8256d380d65c27c4c945068a8b584275196cad8f9d2b7fcbb430f`.

```terraform
# WithRules — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name        = "example"
  namespace   = "system"
  description = "User identification with identification rules"

  rules {
    client_ip = {}
  }
}
```

<a id="canonical-a4f146db9aaa02eed9479c9dc5b583f572a71149feb7873763e573b949c439cb"></a>

## Next pages — With rules / edcb489e74bc / 3

- [Examples](resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
