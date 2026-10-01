---
page_title: "xcsh_cloud_connect landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect landing."
---

# xcsh_cloud_connect landing

<a id="canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212310131113331-0110213313130222-1130322331001313-1321033122202200-1202101112021011-1333213131132002-2121021023320201-2222030200121320"></a>

## xcsh_cloud_connect — xcsh_cloud_connect / 122300223331 / 2

Breadcrumbs:

- xcsh_cloud_connect

Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud
provider networks.

<a id="canonical-3321332030112102-2130100203000030-3321133223130320-2123313132210102-2102013232300032-2333030123132200-2103013311332310-0101002020113111"></a>

## Prerequisites — xcsh_cloud_connect / 122300223331 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0323133032221303-3210320313300032-0220213120202332-3202323300210232-3333330103211322-3232023220010321-1031322210220003-0021312203012200"></a>

## Minimal configuration — xcsh_cloud_connect / 122300223331 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudConnect Resource Example
# Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud provider networks.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudConnect configuration
resource "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}
```

<a id="canonical-1110220102002020-2022323230310033-3221333023300232-2232203322201203-3320201322221131-1130323302120220-3213223312020102-3123220212223200"></a>

## Root configuration — xcsh_cloud_connect / 122300223331 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3201031113311233-1131002010333131-1033013012113002-0322012022023123-3121032012303212-2131132123213202-2300020320131331-2321003302022023"></a>

## Next pages — xcsh_cloud_connect / 122300223331 / 6

- [Property reference](../guides/resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [Examples](../guides/resources--cloud_connect--examples--group-001.md#canonical-0123101101112120-2022133020013020-0302023021001322-2211230202313232-2112121210021023-3202132333203330-1131200300131112-3003102213332121)
- [Import](../guides/resources--cloud_connect--lifecycle--group-001.md#canonical-0303110103131322-0331201202312301-1132313131310220-3312200300323311-1111232310113210-1002300211231013-0030211203302300-2021022023220222)
- [Timeouts](../guides/resources--cloud_connect--lifecycle--group-001.md#canonical-0301022101330330-3113121221220023-3033233123203021-0001230300222012-0320220311021223-2101222202020220-1011312012213132-3022320000100202)
