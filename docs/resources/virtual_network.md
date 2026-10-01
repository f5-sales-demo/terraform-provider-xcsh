---
page_title: "xcsh_virtual_network landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network landing."
---

# xcsh_virtual_network landing

<a id="canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322232200021313-2003003021010023-2311200130120023-3303331002300001-3323223211023300-2303231103320222-2301112020110232-3111112113113023"></a>

## xcsh_virtual_network — xcsh_virtual_network / 233212230310 / 2

Breadcrumbs:

- xcsh_virtual_network

Manages virtual network in given namespace in F5 Distributed Cloud.

<a id="canonical-2311231030132013-2230212000100122-2012113130302023-0101330311333123-3011120300133002-2323332332121012-0323311230333122-2112122320122003"></a>

## Prerequisites — xcsh_virtual_network / 233212230310 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `network_connector`.

- network_connector: Connect to external networks

<a id="canonical-1003203303223101-1123032022102032-3311232021000110-1202010003322211-3223333230132120-1000120033212012-2013210100213323-3222102011231132"></a>

## Minimal configuration — xcsh_virtual_network / 233212230310 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualNetwork Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}
```

<a id="canonical-0022301133333310-1003121220230322-0203000302333303-1111223211033122-0103202203210310-2201300003211210-2103112321020101-0333132321023123"></a>

## Root configuration — xcsh_virtual_network / 233212230310 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1120130221200103-1312012011032010-0133132333130212-0202301132131101-0132221013231330-1331311013011231-3210102121113330-0000102301333323"></a>

## Next pages — xcsh_virtual_network / 233212230310 / 6

- [Property reference](../guides/resources--virtual_network--reference--group-001.md#canonical-0332221310222333-2002312321101003-3120032212130131-1133112320022023-0013003231103132-3300101032110022-0011221221131201-3330112200303033)
- [Examples](../guides/resources--virtual_network--examples--group-001.md#canonical-0101031222132031-3220122033303210-1331311212220132-3302211031002300-0121302012331202-2313202133033133-2132210103310120-3230232030302201)
- [Import](../guides/resources--virtual_network--lifecycle--group-001.md#canonical-2313133021302110-3012232223033200-1222231230132232-2031032101122323-2131331201003110-1001022232300313-1230230122212021-0020333131003110)
- [Timeouts](../guides/resources--virtual_network--lifecycle--group-001.md#canonical-0111310133201213-3020213011120300-3320102131202030-0002001100332223-1201110320200033-1112123233120322-1223222323203113-3011123321001030)
