---
page_title: "xcsh_device_intelligence_device_summary"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_summary."
---

# xcsh_device_intelligence_device_summary

<a id="canonical-3320320012000030-1120213132323310-0033312123010313-0312323233302312-0102031103102222-2230101011132101-1321001213032302-3200123203212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_device_intelligence_device_summary

Reads Device Intelligence Device Summary information from F5 Distributed Cloud.

<a id="canonical-2320301323011213-2223200310111331-2210203111132021-2300320232102010-0211202113113220-1120212120123203-0201021301011123-0111300313330132"></a>

### Prerequisites for `xcsh_device_intelligence_device_summary`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0333021100212012-0330223231222223-3030100230230213-0030101201230011-3212103131122301-0222133113321002-2213223212123310-3221311002113303"></a>

### Minimal configuration for `xcsh_device_intelligence_device_summary`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDeviceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_summary" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_summary_result" {
  value = data.xcsh_device_intelligence_device_summary.example
}
```

<a id="canonical-1102011221311331-2202123120002232-2223010321020230-3031201232311000-0312122103323112-0132132331133013-3030313032320001-1122232132101301"></a>

### Root configuration for `xcsh_device_intelligence_device_summary`

Required root properties: `device_id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3002221032330033-1102321110212010-0112100310020012-1201223313220031-1203132322030330-3123002010230030-3130123132010120-0013103212323002"></a>

### Explore this collection for `xcsh_device_intelligence_device_summary`

- [Property reference](../guides/data-sources--device_intelligence_device_summary--reference--group-001.md#canonical-2103330111212330-1312310023110131-3130101102300031-1201100213003313-0111322033233122-3102333203020332-2002111132021031-0211302211312122)
- [Examples](../guides/data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-3301113121302200-0313213210333013-2322301202203301-2033313012320132-1321030031102310-2131232223001111-0000023213230113-2300002211033113)
