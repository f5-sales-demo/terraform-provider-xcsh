---
page_title: "xcsh_ike2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 landing."
---

# xcsh_ike2 landing

<a id="canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-854ce5d6807fbd4da4b358cf69a761eb821e7b769670bc66583108d6a00cc200"></a>

## xcsh_ike2 — xcsh_ike2 / 98426f556a15 / 2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

<a id="canonical-a9c42a8f73847c9361eda5c3de6b4057c8436e870e174fa4018c92a38d3c7fd8"></a>

## Prerequisites — xcsh_ike2 / 98426f556a15 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e5f02f73c581676bfee1b42b33b0dc31f096916ed3f5596c21fc4c9de2c71fa8"></a>

## Minimal configuration — xcsh_ike2 / 98426f556a15 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike2 configuration
resource "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}
```

<a id="canonical-773046225f5e03fb1b2b3bdb0ea06d042f356e757970a241056175c82a74a04f"></a>

## Root configuration — xcsh_ike2 / 98426f556a15 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c0f1fa9a4f973ef5cf748fd72d79f82bbeb3a715a07f9766363a749e29f131cb"></a>

## Next pages — xcsh_ike2 / 98426f556a15 / 6

- [Property reference](../guides/resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- [Examples](../guides/resources--ike2--examples--group-001.md#canonical-1586c8fb99e7237f1fab2a8ef86fd04ecbd9dcd11dbc823cef6bfe24275e7313)
- [Import](../guides/resources--ike2--lifecycle--group-001.md#canonical-d3ff8a38e4f12a9b771edd878bb8d23fc4aec8e14b8c67290efa7396189c15aa)
- [Timeouts](../guides/resources--ike2--lifecycle--group-001.md#canonical-88608e83e8f711609683d69f30d6bc9ba37b2030a7fd7829cc6268a275e5357c)
