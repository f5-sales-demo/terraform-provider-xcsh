---
page_title: "xcsh_dc_cluster_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group landing."
---

# xcsh_dc_cluster_group landing

<a id="canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131201332303010-0033021201201023-0001210223310311-2111330112101121-3013010010003230-1010331321121010-0323201133211000-3232201322121102"></a>

## xcsh_dc_cluster_group — xcsh_dc_cluster_group / 130122300302 / 2

Breadcrumbs:

- xcsh_dc_cluster_group

Manages DC Cluster group in given namespace in F5 Distributed Cloud.

<a id="canonical-0231232233310331-2130123232223321-1231332333333103-3312203021022030-0113311313221303-2112130112231312-2330210310223311-0211111011100213"></a>

## Prerequisites — xcsh_dc_cluster_group / 130122300302 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1023203301030203-2123111102221102-0003302220332100-1020201003103031-1031321033013001-0310001231222321-1010023203233312-3302201200121023"></a>

## Minimal configuration — xcsh_dc_cluster_group / 130122300302 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DcClusterGroup Resource Example
# Manages DC Cluster group in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DcClusterGroup configuration
resource "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}
```

<a id="canonical-3303123113303200-2312012230323031-0022133312103123-2203010202113333-0221222021223223-0031330013011311-1203220212113011-2202233330222121"></a>

## Root configuration — xcsh_dc_cluster_group / 130122300302 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1311203131021311-0222231310030312-1113322131131223-2122031321202301-3311332010301303-1230322121100130-2321001231021332-0011222021023320"></a>

## Next pages — xcsh_dc_cluster_group / 130122300302 / 6

- [Property reference](../guides/resources--dc_cluster_group--reference--group-001.md#canonical-3020323102031131-2120211103310111-3012300022003323-3023301231000121-3002020033132213-0201300331100120-3122330131102121-1132110021113321)
- [Examples](../guides/resources--dc_cluster_group--examples--group-001.md#canonical-0233100212100021-2032000111023323-3302131321221223-2322320010130122-3012102222300303-3221330203030323-0100131121203311-0313030123131310)
- [Import](../guides/resources--dc_cluster_group--lifecycle--group-001.md#canonical-2131330032300020-3133131032203122-3320100211332102-1330332333210120-1000313301123331-2101311322302003-0000213012303320-1001311221120210)
- [Timeouts](../guides/resources--dc_cluster_group--lifecycle--group-001.md#canonical-3033320120302333-3022130021202001-1012120310300320-3210221111120203-2223113330311101-1313220202221020-0333222322322121-2202032130203131)
