---
page_title: "xcsh_code_base_integration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration landing."
---

# xcsh_code_base_integration landing

<a id="canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011022122303032-1100122203130333-2112331021322222-3101233022223303-0312130323211131-2302220320233311-0302011311312322-3131031303200200"></a>

## xcsh_code_base_integration — xcsh_code_base_integration / 022332231011 / 2

Breadcrumbs:

- xcsh_code_base_integration

Manages integration details in F5 Distributed Cloud.

<a id="canonical-1022323001130110-2132211033012201-1132011320333320-3221313012101020-0023023021223323-3022112203131111-3021221131213311-2301111120313020"></a>

## Prerequisites — xcsh_code_base_integration / 022332231011 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0033301131222223-0110020202223210-1123131211022321-1321133231112201-2320022031030112-0013030230113110-0130030012101032-2100332321310022"></a>

## Minimal configuration — xcsh_code_base_integration / 022332231011 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CodeBaseIntegration Resource Example
# Manages integration details in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CodeBaseIntegration configuration
resource "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}
```

<a id="canonical-0000113221002123-3103132202231221-0222102233200102-3001111120300123-0302112122213120-1303030000032000-3001113310311200-3031211133202003"></a>

## Root configuration — xcsh_code_base_integration / 022332231011 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0121332000131322-1301031021203133-0132320032322211-0121022112213202-0220001302021220-3123021212121102-1312120021113321-1112212330302112"></a>

## Next pages — xcsh_code_base_integration / 022332231011 / 6

- [Property reference](../guides/resources--code_base_integration--reference--group-001.md#canonical-3201211231020022-1303102222210030-3112011321321210-1331302320202012-1010230332303213-3001110101231222-2201322210333200-0201300000023201)
- [Examples](../guides/resources--code_base_integration--examples--group-001.md#canonical-3223311312113011-0300001102331331-3332332231201211-3013102320020123-3131313102303221-3001032000113130-3330213021233132-2201232013021000)
- [Import](../guides/resources--code_base_integration--lifecycle--group-001.md#canonical-3221201113223001-3231030320310302-0302020203002021-2300323111020211-1313111010003203-1011012213310200-0203213003300121-0032101100122213)
- [Timeouts](../guides/resources--code_base_integration--lifecycle--group-001.md#canonical-0222111310312010-2131330212103322-1120331221020202-1131312111120123-3021321201202020-2232121111320023-3023033102222302-1220123101101220)
