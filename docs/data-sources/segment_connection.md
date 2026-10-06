---
page_title: "xcsh_segment_connection"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment_connection."
---

# xcsh_segment_connection

<a id="canonical-0030011210101331-2323210032230002-1133223312021002-2000100321331332-0221323131333131-3323103302112300-2012332113200223-2222011321202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_segment_connection

Reads Segment Connection information from F5 Distributed Cloud.

<a id="canonical-2212312200332232-1100323323310233-1113223110313010-1013201101333201-1101201322021123-0113132212123121-0030100333210320-2133231301020103"></a>

### Prerequisites for `xcsh_segment_connection`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0112030331220313-0010031203103001-1211122203211322-0032130003113033-3221122310301123-2223130023023032-0022302030033302-2032130103320322"></a>

### Minimal configuration for `xcsh_segment_connection`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SegmentConnection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SegmentConnection by name
data "xcsh_segment_connection" "example" {
  name      = "example-segment-connection"
  namespace = "staging"
}

output "segment_connection_id" {
  value = data.xcsh_segment_connection.example.id
}
```

<a id="canonical-2032312022221232-2212301100100113-0103012110312000-0012020022111023-1010122301112321-2300120203331301-2223100201133113-0330130303012013"></a>

### Root configuration for `xcsh_segment_connection`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0133030323023131-2131310310102323-3033113212313131-0203330212200203-0231101310313112-1033112332022322-3221333211022201-2020213330112233"></a>

### Explore this collection for `xcsh_segment_connection`

- [Property reference](../guides/data-sources--segment_connection--reference--group-001.md#canonical-0003210303213121-2323131312120331-1130122231112013-3023230312113302-3230113000232230-2013030312321311-3301212320311012-0202202131321213)
- [Examples](../guides/data-sources--segment_connection--examples--group-001.md#canonical-0210323013221130-1201212220333012-3110212033212013-0023101300000031-2133232101000320-2013333303022112-0013110032021210-1030002313233133)
