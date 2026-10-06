---
page_title: "xcsh_data_type"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type."
---

# xcsh_data_type

<a id="canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_data_type

Manages data\_type creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-3020022022230201-0121101110131202-3111213322220011-0022110200323223-1003231033032210-2233230311303223-1111212020221232-2203120223212130"></a>

### Prerequisites for `xcsh_data_type`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2312122222203010-3311101210211213-3023103312103332-3123333020300010-2321323302002203-3110210103131023-3120202212021032-3012302010103211"></a>

### Minimal configuration for `xcsh_data_type`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```

<a id="canonical-0212330312233331-2312112120032220-0122311130303010-2123300100000033-1002003002232313-1210131203303122-0010031333320002-2023123123332032"></a>

### Root configuration for `xcsh_data_type`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2201302020130211-0233220023333131-2012032312202221-0320332221121201-1312203132003221-1012111212333311-2101032330331101-1000111202200101"></a>

### Explore this collection for `xcsh_data_type`

- [Property reference](../guides/resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [Examples](../guides/resources--data_type--examples--group-001.md#canonical-3312333023112312-2102203021221000-0300101230031301-2301320200100131-0023130210313113-0230300203013330-3023002011013221-3311031012001100)
- [Import](../guides/resources--data_type--lifecycle--group-001.md#canonical-2122031130112123-0233020202310313-0303303003331133-3231010111010321-2010311233333131-3303322010330203-2033303222120322-0213113310013103)
- [Timeouts](../guides/resources--data_type--lifecycle--group-001.md#canonical-0330232323203021-0223201200131021-3133100003122232-0200112320230233-1100321300112230-3131111003330131-3111021323210201-0120011202210313)
