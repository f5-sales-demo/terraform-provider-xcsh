---
page_title: "xcsh_site_mesh_group landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group landing."
---

# xcsh_site_mesh_group landing

<a id="canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201221231320100-1202322103222203-0130201213000100-1111300212110333-1332222331320233-3332323130011230-0030032020201213-1103100202210213"></a>

## xcsh_site_mesh_group — xcsh_site_mesh_group / 021300131131 / 2

Breadcrumbs:

- xcsh_site_mesh_group

Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

<a id="canonical-3331331101012201-3200301233210113-2101302010311311-2300012231030313-1322121333133301-3111210023003102-2301111222011312-2021231233130031"></a>

## Prerequisites — xcsh_site_mesh_group / 021300131131 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `site`.

- site: Sites to include in mesh connectivity

<a id="canonical-1023233012103323-1303222212012331-3011303123322102-2222132313221130-1102213311221013-2012301021032101-1210122231032213-1121001112122120"></a>

## Minimal configuration — xcsh_site_mesh_group / 021300131131 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteMeshGroup Resource Example
# Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SiteMeshGroup configuration
resource "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}
```

<a id="canonical-1031333130032033-3330110212000120-0122320122322300-2321113120333021-3132102202303310-1101312111311003-2200200333111001-3012222000120301"></a>

## Root configuration — xcsh_site_mesh_group / 021300131131 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2000000301131221-1102300102020133-0010020311020122-1023001313211200-1013213032220310-0001203212301113-0230301302221212-2332103021121211"></a>

## Next pages — xcsh_site_mesh_group / 021300131131 / 6

- [Property reference](../guides/resources--site_mesh_group--reference--group-001.md#canonical-0112122103020123-0133230103003320-1203223110023113-0012313123001011-1300131032201323-1011113023033321-0320102011333000-1130221122213113)
- [Examples](../guides/resources--site_mesh_group--examples--group-001.md#canonical-3323110021103110-2201303210222220-2213121010112022-3210303202230223-0003102131301022-3301121021333331-3323112133001111-3023122023131322)
- [Import](../guides/resources--site_mesh_group--lifecycle--group-001.md#canonical-1312210011120113-1303221000111003-1021230032103113-3203130011102120-1211003202133203-1120330320010200-2120311120331102-3303123223322310)
- [Timeouts](../guides/resources--site_mesh_group--lifecycle--group-001.md#canonical-3130213111222112-2223030312130010-0331032031123230-0311020210220111-2322312210100331-2311311312210001-0223213023230313-1311131012331121)
