---
page_title: "xcsh_policer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer examples."
---

# xcsh_policer examples

<a id="canonical-75881073197b2939e6ca9b8bbdd4c65f83484474a98cf95df09244fd909361b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1517b9753c995795d4e4bbc627034cb3ebe82264f1c856d64b49676902964a8b"></a>

## Examples — Examples / 5e521336799c / 2

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)
- Examples

<a id="canonical-375c2f4a1c79f2b0ae6d5bea36cafc6d1ae1637032a901cebe021f95f019ce8b"></a>

## Complete configurations — Examples / 5e521336799c / 3

- [Resource](resources--policer--examples--group-001.md#canonical-565925a27623d2d546e173f5125cd5091974da551f35f734923587b43b215f92): valid configuration.

<a id="canonical-f9a52f640b88eab18b9a8185669f392402bd21f3e1c0973309bd799084c95605"></a>

## Next pages — Examples / 5e521336799c / 4

- [Resource](resources--policer--examples--group-001.md#canonical-565925a27623d2d546e173f5125cd5091974da551f35f734923587b43b215f92)
- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)

<a id="canonical-565925a27623d2d546e173f5125cd5091974da551f35f734923587b43b215f92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94b7cf8e8e92918fe11df6314ccd5fa32ccd42fcbc8fb2f100a3372073dcd8b6"></a>

## Resource — Resource / fac3e229c111 / 2

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)
- [Examples](resources--policer--examples--group-001.md#canonical-75881073197b2939e6ca9b8bbdd4c65f83484474a98cf95df09244fd909361b9)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policer/resource.tf`; digest `sha256:771424522ef5cd3615902f335201bca762afca86a171354ca7b9f6787dcaf443`.

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

<a id="canonical-9802ed9ccc37c60e01c511eb1fa65a71093e0ab9cfc45b0a7ef121eeb5cc2b4e"></a>

## Next pages — Resource / fac3e229c111 / 3

- [Examples](resources--policer--examples--group-001.md#canonical-75881073197b2939e6ca9b8bbdd4c65f83484474a98cf95df09244fd909361b9)
- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)
