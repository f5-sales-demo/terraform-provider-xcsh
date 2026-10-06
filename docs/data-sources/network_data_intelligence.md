---
page_title: "xcsh_network_data_intelligence"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_data_intelligence."
---

# xcsh_network_data_intelligence

<a id="canonical-3123013012321330-1002111020200012-2102200013011121-0212213012311312-3102011113113022-1211131123200320-0301213033123012-1213000322121013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_data_intelligence

Regional Data Intelligence IPv4 destinations. Values are bundled from the pinned OpenAPI release;
this data source performs no network request. Ports and traffic direction are not encoded in the
manifest.

<a id="canonical-1121320102122322-2123010311322220-0030233122003000-1102213230323112-2010201103200232-1202231030011103-0012232111030331-0221030331222011"></a>

### Prerequisites for `xcsh_network_data_intelligence`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3211010122033101-1311100033322030-0133032011120002-1122313321011202-2300131103322313-2333120332023303-3103220223033331-1113232210221103"></a>

### Minimal configuration for `xcsh_network_data_intelligence`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_data_intelligence" "us" {
  regions = ["us"]
}

output "data_intelligence_https_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_data_intelligence.us.cidr_blocks
  }
}
```

<a id="canonical-3303330012330202-0213132211312020-0232021112222230-3031320300301003-1223020220202030-0201100031232030-3021013130020021-0132031000230331"></a>

### Root configuration for `xcsh_network_data_intelligence`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2112210030011023-1013203100111222-3220010000213332-2110131112133310-3113201330102130-3320333113023013-0011121300013201-0132333100130112"></a>

### Explore this collection for `xcsh_network_data_intelligence`

- [Property reference](../guides/data-sources--network_data_intelligence--reference--group-001.md#canonical-3012112022123323-3332203123233211-3331023122230032-2300033323333020-2001232011111103-0002122020120221-0122010012220211-0130201010131103)
- [Examples](../guides/data-sources--network_data_intelligence--examples--group-001.md#canonical-0323132310222120-3020213120021003-2220113000222110-0102223210033320-0112313202331031-3230221031033212-2011132110311232-2203020213322023)
