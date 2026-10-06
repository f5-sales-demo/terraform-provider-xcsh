---
page_title: "xcsh_device_intelligence_summary"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_summary."
---

# xcsh_device_intelligence_summary

<a id="canonical-1301033300112123-2331103013013311-3321132033101320-1310222002112032-1233113332311133-1310100202102222-2023210120211201-1120100233333321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_device_intelligence_summary

Reads Device Intelligence Summary information from F5 Distributed Cloud.

<a id="canonical-2213230230110312-1132130212120110-0233303200331301-0313231111230100-2011010020200031-2213111322311211-2112232111133121-1111303113021303"></a>

### Prerequisites for `xcsh_device_intelligence_summary`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2112312301032323-2321031221103131-1110313021300223-0210103110023012-0211223202310121-2100132131033031-0311313222023303-2033200212133111"></a>

### Minimal configuration for `xcsh_device_intelligence_summary`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_summary" "example" {
  namespace = "example-value"
}

output "device_intelligence_summary_result" {
  value = data.xcsh_device_intelligence_summary.example
}
```

<a id="canonical-1032300231033103-0021323133300112-1331120112332122-1321320332110230-2302032212001103-2313100132330132-0010010222312101-3112032333110031"></a>

### Root configuration for `xcsh_device_intelligence_summary`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2003333103111130-3020130333000121-3312132210333230-3211000303223333-0110132103222301-3300222032311210-0333222131202102-3113100112233320"></a>

### Explore this collection for `xcsh_device_intelligence_summary`

- [Property reference](../guides/data-sources--device_intelligence_summary--reference--group-001.md#canonical-0130130001220323-3031111221132001-3222102300212300-0203021003111000-2213312212330210-2310230230130330-0213320113111122-2313223002130011)
- [Examples](../guides/data-sources--device_intelligence_summary--examples--group-001.md#canonical-3321132220332130-1310310320203213-2303202233130231-1232203011031011-2321031101113113-1210300031010010-3013101010002313-1310231302022030)
