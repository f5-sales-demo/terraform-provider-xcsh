---
page_title: "xcsh_api_discovery"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery."
---

# xcsh_api_discovery

<a id="canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_api_discovery

Reads API Discovery information from F5 Distributed Cloud.

<a id="canonical-2211200001301210-1002112010313312-3201001321221101-0032221010002131-3020310000230313-1130320213213000-1103203312022001-1013102220302032"></a>

### Prerequisites for `xcsh_api_discovery`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2123103320323303-3200312303301021-2031321330212301-2111301131002310-0102220302332020-0310333332000013-0301323130301021-3331111333213002"></a>

### Minimal configuration for `xcsh_api_discovery`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDiscovery by name
data "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}

output "api_discovery_id" {
  value = data.xcsh_api_discovery.example.id
}
```

<a id="canonical-2211223213311223-0033233121131120-3012332022333300-1220310102222311-0031002320012233-0002220010130013-3020323311331221-3201321103221120"></a>

### Root configuration for `xcsh_api_discovery`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1233131230213200-3031011012111313-0001101130103320-1213320310332321-2012010312102221-2000123230103310-3130012120320011-2201022023013222"></a>

### Explore this collection for `xcsh_api_discovery`

- [Property reference](../guides/data-sources--api_discovery--reference--group-001.md#canonical-0133003333000300-0333300330222003-3301102231120300-2100101330332332-0003023001213320-1011031212230301-3133222100223120-0311100220023310)
- [Examples](../guides/data-sources--api_discovery--examples--group-001.md#canonical-3301213201311031-3110013301031111-0113212001331001-1333230011010331-3000023322111303-2320221233022331-1103132033021121-3232301013323113)
