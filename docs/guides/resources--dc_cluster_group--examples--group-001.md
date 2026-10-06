---
page_title: "xcsh_dc_cluster_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group examples."
---

# xcsh_dc_cluster_group examples

<a id="canonical-0233100212100021-2032000111023323-3302131321221223-2322320010130122-3012102222300303-3221330203030323-0100131121203311-0313030123131310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321)
- Examples

<a id="canonical-1310320013102201-2312100300020131-0100211120222323-0120210201021120-1210120132121132-1213222211300111-2231300321120202-1211220103001221"></a>

### Complete configurations for `xcsh_dc_cluster_group`

- [Resource](resources--dc_cluster_group--examples--group-001.md#canonical-0223321230310323-2220020220301011-0123200033210203-2212212210021211-1120230132131231-0101003332213021-3112103031101223-1031333233113323): valid configuration.

<a id="canonical-0223321230310323-2220020220301011-0123200033210203-2212212210021211-1120230132131231-0101003332213021-3112103031101223-1031333233113323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321)
- [Examples](resources--dc_cluster_group--examples--group-001.md#canonical-0233100212100021-2032000111023323-3302131321221223-2322320010130122-3012102222300303-3221330203030323-0100131121203311-0313030123131310)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dc_cluster_group/resource.tf`; digest `sha256:8a0876385980d3e39be1eb4ce88a5f57c0c3917ba68858a952fc5d4f99aaef53`.

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
