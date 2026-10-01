---
page_title: "xcsh_azure_vnet_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site landing."
---

# xcsh_azure_vnet_site landing

<a id="canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301331003003032-3002321323003221-3133101112003012-2010033103233032-3322231313122221-1112101012302103-0222310033121121-1131023212202313"></a>

## xcsh_azure_vnet_site — xcsh_azure_vnet_site / 322300030222 / 2

Breadcrumbs:

- xcsh_azure_vnet_site

Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure
Virtual Network environments.

<a id="canonical-3321223313020303-2011102233201330-2303103300110131-2121122303023210-3331323122231002-2301122130103020-2130230023220122-0213123001222202"></a>

## Prerequisites — xcsh_azure_vnet_site / 322300030222 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: Azure authentication for deployment

<a id="canonical-2302002321222110-0020012332210011-1133010122202230-3000322321201212-0330331133013223-2012112230232311-2211301203022010-0230103113230332"></a>

## Minimal configuration — xcsh_azure_vnet_site / 322300030222 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AzureVNETSite Resource Example
# Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure Virtual Network environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AzureVNETSite configuration
resource "xcsh_azure_vnet_site" "example" {
  name      = "example-azure-vnet-site"
  namespace = "system"

  machine_type   = "example-value"
  resource_group = "example-value"
  ssh_key        = "example-value"
}
```

<a id="canonical-3101302332113210-2010303323310011-3112110230230202-1303221222101233-3011133010010012-3213110032023130-1200223130103113-1323101020230330"></a>

## Root configuration — xcsh_azure_vnet_site / 322300030222 / 5

Required root properties: `machine_type`, `name`, `resource_group`, `ssh_key`. Full root flags and choices appear in the property reference.

<a id="canonical-1333223001120011-0230232131113201-3020031121113022-2223201032110103-3230332233031213-3300120010002200-2323213133303223-2301320203311022"></a>

## Next pages — xcsh_azure_vnet_site / 322300030222 / 6

- [Property reference](../guides/resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [Examples](../guides/resources--azure_vnet_site--examples--group-001.md#canonical-2121031322130301-3000313002230303-0033200031101332-0221120002231122-3220200111201311-0223201002101333-0230022122310231-3320021220101332)
- [Import](../guides/resources--azure_vnet_site--lifecycle--group-001.md#canonical-3312102030220330-2203101332331333-0213222311213313-2123023222023330-2032100010230322-2200030200011010-3000121101322320-1302212112300131)
- [Timeouts](../guides/resources--azure_vnet_site--lifecycle--group-001.md#canonical-1211130211002013-0033110131023330-0012113031031222-1211021101333321-3112023332322113-3132101323311211-0300113223220100-3232011332121031)
