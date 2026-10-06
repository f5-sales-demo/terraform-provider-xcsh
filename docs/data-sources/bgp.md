---
page_title: "xcsh_bgp"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp."
---

# xcsh_bgp

<a id="canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bgp

Reads BGP peering configuration with external BGP servers in the system namespace.

<a id="canonical-0120312330231102-2001033031020113-0122111231120101-0322331321120322-1313111012113023-1223003222333221-0223332211233322-2022311111031223"></a>

### Prerequisites for `xcsh_bgp`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0020301301230123-2223132010210332-3100000302120222-2311013322201123-0131213331230333-1321220333100133-0213301210231323-2223323332012210"></a>

### Minimal configuration for `xcsh_bgp`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGP by name
data "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}

output "bgp_id" {
  value = data.xcsh_bgp.example.id
}
```

<a id="canonical-1231310231002323-1023130301002231-0310101220011310-3102323130200311-3133101303213202-3133203123211202-1200310100110030-3033300110122032"></a>

### Root configuration for `xcsh_bgp`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1121010212331120-2122111030013001-3013221120212033-3320110020011010-3030000201210033-1232012232120120-0232123102311311-1113303222113033"></a>

### Explore this collection for `xcsh_bgp`

- [Property reference](../guides/data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [Examples](../guides/data-sources--bgp--examples--group-001.md#canonical-0011333200121323-1202013222010111-0303012103322321-1132212333033031-2001120120213230-0113023223130122-0310112310322311-2002111203001003)
