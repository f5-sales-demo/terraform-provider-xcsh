---
page_title: "xcsh_srv6_network_slice examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice examples."
---

# xcsh_srv6_network_slice examples

<a id="canonical-0202203011323022-1213210303303201-2130020311033333-2223212322220201-3001213012200323-2112301210033031-0033020330120230-0321223030213201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332)
- Examples

<a id="canonical-3000102110223133-2311131323112323-1232200321132101-3123233202212011-0013321200010122-2321223213312002-3133203233302121-1320020232133211"></a>

### Complete configurations for `xcsh_srv6_network_slice`

- [Resource](resources--srv6_network_slice--examples--group-001.md#canonical-1331111120302201-1301121320111201-3320103121320102-0111112211101230-0010013013200103-1200330021230010-3021100201320131-2110302311023232): valid configuration.

<a id="canonical-1331111120302201-1301121320111201-3320103121320102-0111112211101230-0010013013200103-1200330021230010-3021100201320131-2110302311023232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-1003011003113030-2133000032112000-2110300131312112-0323202222213032-3010021101322220-3130233303030103-2300101230220121-0201100002310332)
- [Examples](resources--srv6_network_slice--examples--group-001.md#canonical-0202203011323022-1213210303303201-2130020311033333-2223212322220201-3001213012200323-2112301210033031-0033020330120230-0321223030213201)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_srv6_network_slice/resource.tf`; digest `sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0`.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```
