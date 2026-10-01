---
page_title: "xcsh_ike_phase1_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile landing."
---

# xcsh_ike_phase1_profile landing

<a id="canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da1d883eecb55a05e8ba96d836b955916a773a8c86395c17ac787a63d021dcae"></a>

## xcsh_ike_phase1_profile — xcsh_ike_phase1_profile / 46d1fbbd6597 / 2

Breadcrumbs:

- xcsh_ike_phase1_profile

Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.
configuration.

<a id="canonical-36354fc201e03984de4e56cee1ad6b4d302d4d69bff7c067fd52c81b422226fd"></a>

## Prerequisites — xcsh_ike_phase1_profile / 46d1fbbd6597 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e5194402ef154bf34b12655ab748aeaaddd0492ec989e9bed6f6499b9c2e5817"></a>

## Minimal configuration — xcsh_ike_phase1_profile / 46d1fbbd6597 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-42723c123dac886b1f78fa8bcb9d6290edcd21027cce444e64cd9eec02fc78f5"></a>

## Root configuration — xcsh_ike_phase1_profile / 46d1fbbd6597 / 5

Required root properties: `authentication_algos`, `dh_group`, `encryption_algos`, `name`, `namespace`, `prf`. Full root flags and choices appear in the property reference.

<a id="canonical-d20603a606a035203563904b49fc8949db9404814c7696de06dc0b12c6918de7"></a>

## Next pages — xcsh_ike_phase1_profile / 46d1fbbd6597 / 6

- [Property reference](../guides/resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [Examples](../guides/resources--ike_phase1_profile--examples--group-001.md#canonical-7b1a7d855a2f6a6f5e667b836423a2ae8f595c683b5178e68e6a6a1a2559eefc)
- [Import](../guides/resources--ike_phase1_profile--lifecycle--group-001.md#canonical-3ad820614ef7017cb7d00d786fa5e3140f25d59808384cf18072af8f505b49a8)
- [Timeouts](../guides/resources--ike_phase1_profile--lifecycle--group-001.md#canonical-d652ab502f3be7de3f9e59a6bc69f7d98615a2568271d4e5d9209970ce274ea9)
