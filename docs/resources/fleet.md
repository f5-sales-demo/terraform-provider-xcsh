---
page_title: "xcsh_fleet landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet landing."
---

# xcsh_fleet landing

<a id="canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4176887752e65f534f80e6e08e900e5f1a4790d3d88b975757792a829b31547"></a>

## xcsh_fleet — xcsh_fleet / 1e2cf1b9651a / 2

Breadcrumbs:

- xcsh_fleet

Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

<a id="canonical-a349fa8e582b670c4a988bdd4bf5d214304f9f6f73982900d69723fbfbbce6bb"></a>

## Prerequisites — xcsh_fleet / 1e2cf1b9651a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-9155b3a22fb7368058acb436bcf847de65cc3a1093f84fda29acf8a55609650e"></a>

## Minimal configuration — xcsh_fleet / 1e2cf1b9651a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Fleet Resource Example
# Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fleet configuration
resource "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"

  fleet_label = "example-value"
}
```

<a id="canonical-8e869257f5139e38dd15b2b53a7e66277b69dba030f6648d311648b9fb3aef1b"></a>

## Root configuration — xcsh_fleet / 1e2cf1b9651a / 5

Required root properties: `fleet_label`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-7177a77db3f522e3db49541642fd5b1c048626b9b72b69b34478c1fd464a0405"></a>

## Next pages — xcsh_fleet / 1e2cf1b9651a / 6

- [Property reference](../guides/resources--fleet--reference--group-001.md#canonical-f6846a0e8eea9a63b350fc210b88d4323acde598409b5a08a6e982a650bfd8f0)
- [Examples](../guides/resources--fleet--examples--group-001.md#canonical-f088ffb8e7ae87e96ee1781aef5f7a9d6ced2ed410db553b5a7fe67146c45fa5)
- [Import](../guides/resources--fleet--lifecycle--group-001.md#canonical-e9c56e0cfb53edb101d7203270a9f46998ef02a6647fb1587c3cd771ead63f1e)
- [Timeouts](../guides/resources--fleet--lifecycle--group-001.md#canonical-512baea545bf03ee097e1ecc90dcb56f14534657f7a2900cfa20eff1d4365e85)
