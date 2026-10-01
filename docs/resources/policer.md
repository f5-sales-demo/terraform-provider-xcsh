---
page_title: "xcsh_policer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer landing."
---

# xcsh_policer landing

<a id="canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-602df954a8991d30840f6d436718d80cf6489791652584be178b1c1fad589e7a"></a>

## xcsh_policer — xcsh_policer / 19287e17873a / 2

Breadcrumbs:

- xcsh_policer

Manages new policer with traffic rate limits in F5 Distributed Cloud.

<a id="canonical-3cee3036db8e78b55ba4ee0b5340b4e1eb540ebdd504b58927aca3108f6b1f11"></a>

## Prerequisites — xcsh_policer / 19287e17873a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-63f35803fa20639c5cbc79701e4a6509cdc6a572eeee9233c93c8f3f159ca51c"></a>

## Minimal configuration — xcsh_policer / 19287e17873a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Policer configuration
resource "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"

  burst_size                 = 1
  committed_information_rate = 1
}
```

<a id="canonical-2c55362258263ee343448a3a2a3a6b9302fac15c4fa3339fc6f4e34128092aba"></a>

## Root configuration — xcsh_policer / 19287e17873a / 5

Required root properties: `burst_size`, `committed_information_rate`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-5c29cd5bf5a2a438fa81ce22661cd5f3cd61af252469737fbe7dd6f6d6683910"></a>

## Next pages — xcsh_policer / 19287e17873a / 6

- [Property reference](../guides/resources--policer--reference--group-001.md#canonical-25f1d1c8b4cdf547a8ecbbbcab78c6c88568d856b48014e86944888cf54e0170)
- [Examples](../guides/resources--policer--examples--group-001.md#canonical-75881073197b2939e6ca9b8bbdd4c65f83484474a98cf95df09244fd909361b9)
- [Import](../guides/resources--policer--lifecycle--group-001.md#canonical-63994a1d2fbd16b827033eed38c0efb05075a336cf88cd2a3226a074bad51844)
- [Timeouts](../guides/resources--policer--lifecycle--group-001.md#canonical-cab4f55c9ba5e8d9bd5a2356b537cd183cec28f7956e0794c1bb29c76e3fa5c9)
