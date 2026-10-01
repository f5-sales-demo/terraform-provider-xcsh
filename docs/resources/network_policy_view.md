---
page_title: "xcsh_network_policy_view landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view landing."
---

# xcsh_network_policy_view landing

<a id="canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200121332213221-1101301003121020-0033033203201022-0110230132312021-1201220333120023-0201130132332120-0112310232200200-3111003320311131"></a>

## xcsh_network_policy_view — xcsh_network_policy_view / 211002212131 / 2

Breadcrumbs:

- xcsh_network_policy_view

Manages a Network Policy View resource in F5 Distributed Cloud for network policy view
specification. configuration.

<a id="canonical-0331030121223030-0221332022023330-3022331230103213-3012330311330031-2031333000322030-0222011232121303-1113133130332032-0302022021133113"></a>

## Prerequisites — xcsh_network_policy_view / 211002212131 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3121112211020122-2130111111012013-1001100233131333-1313011002102120-2102021113010101-2201231211030220-2332132123103100-1033321223303300"></a>

## Minimal configuration — xcsh_network_policy_view / 211002212131 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyView Resource Example
# Manages a Network Policy View resource in F5 Distributed Cloud for network policy view specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyView configuration
resource "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}
```

<a id="canonical-1101202213120103-0020210101203211-3021120132320213-1121301210100123-1000022203032010-3202133233201111-1133200032021002-0001023033020331"></a>

## Root configuration — xcsh_network_policy_view / 211002212131 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1223133122311032-1102022221112001-0133012333011333-3203222231212223-2023111001032133-0320322310123022-3131313112230102-3323101210120300"></a>

## Next pages — xcsh_network_policy_view / 211002212131 / 6

- [Property reference](../guides/resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [Examples](../guides/resources--network_policy_view--examples--group-001.md#canonical-0111310100333132-3203131111203113-0112030133332200-2221321020230210-1230101212322333-3100103111122002-2100330131112123-3230201322132133)
- [Import](../guides/resources--network_policy_view--lifecycle--group-001.md#canonical-0333213131303310-0010230232122223-1210302201222121-1232100200101312-2020321002323232-1002022133033103-2213122131013312-1322321010320023)
- [Timeouts](../guides/resources--network_policy_view--lifecycle--group-001.md#canonical-1110133222000112-1103032211103133-1303112222310121-1002020022012333-3012301211311022-0300022220311233-2310031200312201-3010200201110330)
