---
page_title: "xcsh_ike1 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 landing."
---

# xcsh_ike1 landing

<a id="canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae506d7b3e53444043a5a079a64e08e3e72916a0a684072a850cdcca97483b41"></a>

## xcsh_ike1 — xcsh_ike1 / 0e219fbee081 / 2

Breadcrumbs:

- xcsh_ike1

Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.

<a id="canonical-115959fefa010f4ece827961d082ef0064240239b85615557b2da87200492366"></a>

## Prerequisites — xcsh_ike1 / 0e219fbee081 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-9c0e77c1d40c622fb44fe5bf4e358209536321ae93ac6d7260d11f016e72381c"></a>

## Minimal configuration — xcsh_ike1 / 0e219fbee081 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```

<a id="canonical-0e6baeea1445dbd92e5315913f435ccd5db16834ae712fef88adfd8d3dfba380"></a>

## Root configuration — xcsh_ike1 / 0e219fbee081 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-cff6290d335cfdd3b0b889d699d3f189fd83113d115f74745e85877bc93f467b"></a>

## Next pages — xcsh_ike1 / 0e219fbee081 / 6

- [Property reference](../guides/resources--ike1--reference--group-001.md#canonical-2debb276abd42f87e00a9690a0634f2eecd9a9c655336e52bd08eac86a5ede94)
- [Examples](../guides/resources--ike1--examples--group-001.md#canonical-a90ace16fce8dfc643a25b1907e73a1d38df1be791702aa2eaa0f7006aee33a1)
- [Import](../guides/resources--ike1--lifecycle--group-001.md#canonical-02964e4b1a0397370a82257e7a3a2ec431e800a2155f87d0503c5047bc793c99)
- [Timeouts](../guides/resources--ike1--lifecycle--group-001.md#canonical-045d2ec2f4d9e57680d007aa0d85b3b2e1146cafdb5eec4c7a9e33c188e74bdc)
