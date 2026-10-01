---
page_title: "xcsh_dc_cluster_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group landing."
---

# xcsh_dc_cluster_group landing

<a id="canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113331020123111-3200221010133112-1110331200023212-3001103002110221-0301001102323010-3201311231111101-1303001120103233-3013130202031212"></a>

## xcsh_dc_cluster_group — xcsh_dc_cluster_group / 323221022023 / 2

Breadcrumbs:

- xcsh_dc_cluster_group

Manages DC Cluster group in given namespace in F5 Distributed Cloud.

<a id="canonical-3213130221213130-3112231013112002-3121202213111231-2300111031132013-3333001030210033-1001010022202233-3001010030011123-1033132300222322"></a>

## Prerequisites — xcsh_dc_cluster_group / 323221022023 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2203303020110231-3231101331110201-1232032100231312-2112222223200013-1110232333133010-0300202333021221-1030122031120321-2002303333212233"></a>

## Minimal configuration — xcsh_dc_cluster_group / 323221022023 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DcClusterGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DcClusterGroup by name
data "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}

output "dc_cluster_group_id" {
  value = data.xcsh_dc_cluster_group.example.id
}
```

<a id="canonical-3223201102210111-3203202312330311-0010323323033311-1100102331103033-2230322223001201-2030032020233130-2002130231220310-0031032203132111"></a>

## Root configuration — xcsh_dc_cluster_group / 323221022023 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1120012223301022-3200021220013212-0213201223331030-2312110011332103-0322303023303203-0201203023203013-2232320202232032-3332301131100201"></a>

## Next pages — xcsh_dc_cluster_group / 323221022023 / 6

- [Property reference](../guides/data-sources--dc_cluster_group--reference--group-001.md#canonical-1010333130230311-1033232300311001-0022220103320030-1332313110202221-3021223120002030-2000202203133231-3233130200102213-2123030302321023)
- [Examples](../guides/data-sources--dc_cluster_group--examples--group-001.md#canonical-1022213222022232-3222020310023000-0303221023121000-3320012211022120-1222303133213203-1303220103100323-0121022332201202-1103310100003021)
