---
page_title: "xcsh_dc_cluster_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group examples."
---

# xcsh_dc_cluster_group examples

<a id="canonical-1022213222022232-3222020310023000-0303221023121000-3320012211022120-1222303133213203-1303220103100323-0121022332201202-1103310100003021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313)
- Examples

<a id="canonical-2210321233031120-1120302030203333-3210111132330130-1130113021013022-1222331332202020-1201333220131012-3103330011312222-3130030103322133"></a>

### Complete configurations for `xcsh_dc_cluster_group`

- [Data source](data-sources--dc_cluster_group--examples--group-001.md#canonical-2233313011020010-1321013321103112-0202103311312302-2101102021231212-3111020232003032-3012321120300301-2201311000111311-3302202102010002): valid configuration.

<a id="canonical-2233313011020010-1321013321103112-0202103311312302-2101102021231212-3111020232003032-3012321120300301-2201311000111311-3302202102010002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md#canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313)
- [Examples](data-sources--dc_cluster_group--examples--group-001.md#canonical-1022213222022232-3222020310023000-0303221023121000-3320012211022120-1222303133213203-1303220103100323-0121022332201202-1103310100003021)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dc_cluster_group/data-source.tf`; digest `sha256:82b040ec9e5742587a1d4d2feb953ac78e9dca3a538c42aff66a3020a25fe4c5`.

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
