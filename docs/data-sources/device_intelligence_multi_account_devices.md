---
page_title: "xcsh_device_intelligence_multi_account_devices"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_multi_account_devices."
---

# xcsh_device_intelligence_multi_account_devices

<a id="canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_device_intelligence_multi_account_devices

Reads Device Intelligence Multi Account Devices information from F5 Distributed Cloud.

<a id="canonical-2113230333320303-3010100223320021-3211112203022000-3313030111223311-2312033001132100-2323010302020020-1321211023321320-1131013200211120"></a>

### Prerequisites for `xcsh_device_intelligence_multi_account_devices`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3222001122213232-0201301301203001-3013032102120100-2112000122010012-0010302101310201-0030301230112121-2131021303103331-1313023321121320"></a>

### Minimal configuration for `xcsh_device_intelligence_multi_account_devices`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceMultiAccountDevices DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_multi_account_devices" "example" {
  namespace = "example-value"
}

output "device_intelligence_multi_account_devices_result" {
  value = data.xcsh_device_intelligence_multi_account_devices.example
}
```

<a id="canonical-0200112102020233-1231030023020022-2122301021212123-3223301111233132-1120110311111221-3002033320003133-2302132223322032-2021211221131302"></a>

### Root configuration for `xcsh_device_intelligence_multi_account_devices`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2123321332232232-3100102012111120-3330102122322203-1320113312220133-0022220311210103-1110111013333102-1311232033010130-2012132112203310"></a>

### Explore this collection for `xcsh_device_intelligence_multi_account_devices`

- [Property reference](../guides/data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0002211303210131-3301122203222131-3121332000231220-3031103123001111-3332222231232120-3000111203301111-1311330311113231-3322102221001210)
- [Examples](../guides/data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-2103201103310120-0321301032000010-0223230302330212-3120120201333300-0132210303031113-1201231000011212-2213131003022303-3203003132131320)
