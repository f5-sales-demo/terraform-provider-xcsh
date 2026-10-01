---
page_title: "xcsh_ike_phase1_profile examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile examples."
---

# xcsh_ike_phase1_profile examples

<a id="canonical-7b1a7d855a2f6a6f5e667b836423a2ae8f595c683b5178e68e6a6a1a2559eefc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-336f82fb8d579d68bdf2600a8b6154b26cbf68b5e563ae53d3824e7d27a1d774"></a>

## Examples — Examples / bde35010a0d2 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- Examples

<a id="canonical-ac14eabfd4d26f3dd3eb4e2832ece4cb0e622f9f4e1b8c220a70e9cf8042d0d4"></a>

## Complete configurations — Examples / bde35010a0d2 / 3

- [Resource](resources--ike_phase1_profile--examples--group-001.md#canonical-b9494dc75ff44002ae22600643a49cbfc1f2db1e165539c91720bd4d4352642e): valid configuration.

<a id="canonical-a798acf48e4d4ed08a8e7aae87e4e6fc5c1358af15311d37eb958443b00c0c3c"></a>

## Next pages — Examples / bde35010a0d2 / 4

- [Resource](resources--ike_phase1_profile--examples--group-001.md#canonical-b9494dc75ff44002ae22600643a49cbfc1f2db1e165539c91720bd4d4352642e)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-b9494dc75ff44002ae22600643a49cbfc1f2db1e165539c91720bd4d4352642e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-458e3b5fa97e5e8a83e0f41546f687f6f6f31d6ddb9408326d2dacbef97afc62"></a>

## Resource — Resource / 658a1eb4b277 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Examples](resources--ike_phase1_profile--examples--group-001.md#canonical-7b1a7d855a2f6a6f5e667b836423a2ae8f595c683b5178e68e6a6a1a2559eefc)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike_phase1_profile/resource.tf`; digest `sha256:8b0906176c52164cc70b7c3ef3103bf969192548192507a8a9df9fa2b687acf6`.

```terraform
# IKEPhase1Profile Resource Example
# Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase1Profile configuration
resource "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  dh_group             = ["example-value"]
  encryption_algos     = ["example-value"]
  prf                  = ["example-value"]
}
```

<a id="canonical-44794a75049431dced184fe90aae36369e1b3110d514b329eb016c5f2cce704a"></a>

## Next pages — Resource / 658a1eb4b277 / 3

- [Examples](resources--ike_phase1_profile--examples--group-001.md#canonical-7b1a7d855a2f6a6f5e667b836423a2ae8f595c683b5178e68e6a6a1a2559eefc)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
