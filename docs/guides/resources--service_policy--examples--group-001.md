---
page_title: "xcsh_service_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy examples."
---

# xcsh_service_policy examples

<a id="canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-200681942f9c83f8ea20cb4193ccec0ba1b9c5d7aebb05457ef226f5d5266b6b"></a>

## Examples — Examples / 05c362778369 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- Examples

<a id="canonical-d7d8e62c9a5c772addd704e7dd72eba3c8660e07cf3fa78c7df9e55da99e86d4"></a>

## Complete configurations — Examples / 05c362778369 / 3

- [Allow list](resources--service_policy--examples--group-001.md#canonical-2f4babdb35c48396dd09a128918eb6d96ea5a2513e9f7274da31e6ec0cdbd91e): valid configuration.

- [Deny all](resources--service_policy--examples--group-001.md#canonical-44a387b24e7ad4860368284e6c43bed6cbebb3a390710a1e82e21f4f25c7646a): valid configuration.

- [Deny list](resources--service_policy--examples--group-001.md#canonical-484897bb6d35a334a51db3b78efa1125a3450cc3c32fd7a000e2541892956ac3): valid configuration.

- [Resource](resources--service_policy--examples--group-001.md#canonical-f8966b7117a23c8fca0d1392345c1f0fbff1521dbf8a7f2311d9b16ceeca6748): valid configuration.

- [With labels](resources--service_policy--examples--group-001.md#canonical-50719778c203e1931c3a083759487a0a80004e270f4c4e173787b741ba775a17): valid configuration.

<a id="canonical-8135f2c963cc4c1a7ebdf43e916777b701a52ab71b1680943354fb4a2e7e0f8c"></a>

## Next pages — Examples / 05c362778369 / 4

- [Allow list](resources--service_policy--examples--group-001.md#canonical-2f4babdb35c48396dd09a128918eb6d96ea5a2513e9f7274da31e6ec0cdbd91e)
- [Deny all](resources--service_policy--examples--group-001.md#canonical-44a387b24e7ad4860368284e6c43bed6cbebb3a390710a1e82e21f4f25c7646a)
- [Deny list](resources--service_policy--examples--group-001.md#canonical-484897bb6d35a334a51db3b78efa1125a3450cc3c32fd7a000e2541892956ac3)
- [Resource](resources--service_policy--examples--group-001.md#canonical-f8966b7117a23c8fca0d1392345c1f0fbff1521dbf8a7f2311d9b16ceeca6748)
- [With labels](resources--service_policy--examples--group-001.md#canonical-50719778c203e1931c3a083759487a0a80004e270f4c4e173787b741ba775a17)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-2f4babdb35c48396dd09a128918eb6d96ea5a2513e9f7274da31e6ec0cdbd91e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5321799ab6fcf5aac06fb134cb56af24c2e4acf97c6ae5e2adb1ffbd2c05fa3"></a>

## Allow list — Allow list / 215f773954d6 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- Allow list

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/allow-list.tf`; digest `sha256:e3e5a70e5d1b6c7d8436b9826fb90da18093ff0f1b51c42b00f77c75bb40e353`.

```terraform
# AllowList — Acceptance-test-derived Configuration
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

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  # Allow list with IP prefix
  allow_list {
    prefix_list {
      prefixes = ["10.0.0.0/8", "192.168.0.0/16"]
    }
    default_action_deny = {}
  }

  # Apply to any server
  any_server = {}
}
```

<a id="canonical-9228e4c18617363f1f0338d9341bb1d8c65e677c2640338ba48f2090d7432bcf"></a>

## Next pages — Allow list / 215f773954d6 / 3

- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-44a387b24e7ad4860368284e6c43bed6cbebb3a390710a1e82e21f4f25c7646a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3fad76075d8cdebfc0757bd5d20abeed1d2ac5e496ba0204fc71428cf6a5d2f"></a>

## Deny all — Deny all / a444c9785e87 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- Deny all

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/deny-all.tf`; digest `sha256:9da9688512df86906769c34d1e33802dc2db45bff21587219ced7d18b6172282`.

```terraform
# DenyAll — Acceptance-test-derived Configuration
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

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  # Deny all requests
  deny_all_requests = {}

  # Apply to any server
  any_server = {}
}
```

<a id="canonical-68e1b304106d3ccf3f4c559cf3e2ade4c82e933b8b0d95466a1b62225a38eb02"></a>

## Next pages — Deny all / a444c9785e87 / 3

- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-484897bb6d35a334a51db3b78efa1125a3450cc3c32fd7a000e2541892956ac3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24eabfb2ef6277803e4ec3de06f867dc89c0c356e1b763cb4b2906dacbcc2880"></a>

## Deny list — Deny list / b86daf8d6360 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- Deny list

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/deny-list.tf`; digest `sha256:6000e015862d1bb25f8e6b5bc94664685fdab920f11f2608c4f47416c65ff393`.

```terraform
# DenyList — Acceptance-test-derived Configuration
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

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  deny_list {
    prefix_list {
      prefixes = ["172.16.0.0/12"]
    }
    default_action_allow = {}
  }

  any_server = {}
}
```

<a id="canonical-50262f95148010939969303a954fc71fa7d6efeaf6f986578a9f251bf943091e"></a>

## Next pages — Deny list / b86daf8d6360 / 3

- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-f8966b7117a23c8fca0d1392345c1f0fbff1521dbf8a7f2311d9b16ceeca6748"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e56f468fdd1804f29a9fe45cd654c0ef75adc469608a4bd0460ba8e926e6e653"></a>

## Resource — Resource / 269aea7c9ff2 / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/resource.tf`; digest `sha256:e0fbf1c5446df6211fe69296e8020455f7d5f52c63778a8e6218ac43e92a6906`.

```terraform
# ServicePolicy Resource Example
# Manages service_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicy configuration
resource "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}
```

<a id="canonical-f2abb035132ed12a55a0aee42c3cce3a3f831bc6b879f6d22022294482de2df9"></a>

## Next pages — Resource / 269aea7c9ff2 / 3

- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)

<a id="canonical-50719778c203e1931c3a083759487a0a80004e270f4c4e173787b741ba775a17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-016b1bca66cbb658a6657fe6fffefddcfa9b45678a39d5d2bc241af1f1fa0a6d"></a>

## With labels — With labels / 8d81e383444d / 2

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/with-labels.tf`; digest `sha256:27966c953ab74ec12398f4c688232398a984b67f1bec6a9179d66139d051b625`.

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

resource "xcsh_service_policy" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test service policy"

  labels = {
    environment = "test"
    team        = "security"
  }

  # Allow all requests
  allow_all_requests = {}

  # Apply to any server
  any_server = {}
}
```

<a id="canonical-bbc6e62c560f74a8ecc3bab06376f25c9ed4db6089e82ef6af98509cee38a9fa"></a>

## Next pages — With labels / 8d81e383444d / 3

- [Examples](resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- [xcsh_service_policy](../resources/service_policy.md#canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1)
