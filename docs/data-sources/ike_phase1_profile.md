---
page_title: "xcsh_ike_phase1_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile landing."
---

# xcsh_ike_phase1_profile landing

<a id="canonical-7e900fd939fc8632d9be51fd78b7c94f953410b370971c904dd66cb6676b2f0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a4c4640f90f836d5980c64655158989e22c77ddcc850f8861549ba8e5708c68"></a>

## xcsh_ike_phase1_profile — xcsh_ike_phase1_profile / 9fb8303ab62f / 2

Breadcrumbs:

- xcsh_ike_phase1_profile

Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.
configuration.

<a id="canonical-747f3a60ddf71d09c9ca38d37be1cc8c69a4f9bb54a23745a9da7bc8d771bac2"></a>

## Prerequisites — xcsh_ike_phase1_profile / 9fb8303ab62f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-952cc5053fbbbd6accf2e3ddb82fe96cb198dc8326975e31b6b00977cf028df6"></a>

## Minimal configuration — xcsh_ike_phase1_profile / 9fb8303ab62f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase1Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase1Profile by name
data "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"
}

output "ike_phase1_profile_id" {
  value = data.xcsh_ike_phase1_profile.example.id
}
```

<a id="canonical-eb49d68b028276a68ce5b992eb78eecd7d105bf8d543ac143ee53115ba0800d4"></a>

## Root configuration — xcsh_ike_phase1_profile / 9fb8303ab62f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d4d8ba5e068c73bc909ecb45b6a5b568d5dd82c5b2de2ae6878e9fa71a0eeddf"></a>

## Next pages — xcsh_ike_phase1_profile / 9fb8303ab62f / 6

- [Property reference](../guides/data-sources--ike_phase1_profile--reference--group-001.md#canonical-680a0a7af1996f5da96b1557527fd012e4f56985b5c17aba003ac0aa65e34fdb)
- [Examples](../guides/data-sources--ike_phase1_profile--examples--group-001.md#canonical-dd4f9b52237c7b7acd3c241cb898af511b1254c0fa3b3b89791832c09351fdea)
