---
page_title: "xcsh_data_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group landing."
---

# xcsh_data_group landing

<a id="canonical-8fa44c0b348749f6556f32f242093dd90beea888fbbeb6165a302dd0d3a951cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7319206f38d572d79a209967be20cf8af5738ffdd28269c2ac27fc30e2709981"></a>

## xcsh_data_group — xcsh_data_group / c56901fe0bec / 2

Breadcrumbs:

- xcsh_data_group

Manages data group in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-871a0b2d842e38eba7f9b7476025f4fd7db7b7db5592ad0f241922dc4612818a"></a>

## Prerequisites — xcsh_data_group / c56901fe0bec / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ca3c01bd464e21d95a5abf6100c38af5d50b6ba7abbe5f0538c064be36dac0a3"></a>

## Minimal configuration — xcsh_data_group / c56901fe0bec / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```

<a id="canonical-f5338865819073604d358ccd8b36633e46ab241604ebfa7bbcceb5c82fc99298"></a>

## Root configuration — xcsh_data_group / c56901fe0bec / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-dbdbc44f486715b47e7f4458aedca74086a14b7729ef601843913c59e3e13094"></a>

## Next pages — xcsh_data_group / c56901fe0bec / 6

- [Property reference](../guides/resources--data_group--reference--group-001.md#canonical-ed344f915adc27bbb56a72ed0c2fdcae2a9f2a3364048364e25bb8dc6b6f309e)
- [Examples](../guides/resources--data_group--examples--group-001.md#canonical-1d293247dc1bbafe31dca58399cf1cf445f54893e37cb987fa1aae1f27900bb7)
- [Import](../guides/resources--data_group--lifecycle--group-001.md#canonical-956f582ae35bd6ba68bf106ad37ec0701f0e92158d3347507dbc126a914ca166)
- [Timeouts](../guides/resources--data_group--lifecycle--group-001.md#canonical-a7e57f32081f18308109b8b393dd7ebd58630cad04970e6107487adce1cbea53)
