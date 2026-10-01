---
page_title: "xcsh_device_intelligence_device_history landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_history landing."
---

# xcsh_device_intelligence_device_history landing

<a id="canonical-1033022013313010-2230031200002233-1320022023130030-0022221001122220-0201311013330313-2011320120303230-1323223312110203-0123022120001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032023312312120-1021221230213132-0011112120232320-0122322333310300-1031011020220201-2202311331333003-3213003300332010-0001311022200030"></a>

## xcsh_device_intelligence_device_history — xcsh_device_intelligence_device_history / 030223001210 / 2

Breadcrumbs:

- xcsh_device_intelligence_device_history

Resource creation operation.

<a id="canonical-2121233110121023-2231001121303121-2230231220031033-1313101312022213-0033333302313023-1223303102132032-3102130311033000-3320331313312311"></a>

## Prerequisites — xcsh_device_intelligence_device_history / 030223001210 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2200122022230300-0111333320313122-3030231101203003-2122020113210012-0200223202112232-0130112330103223-1101101103100212-1313001010301312"></a>

## Minimal configuration — xcsh_device_intelligence_device_history / 030223001210 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDeviceHistory DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_history" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_history_result" {
  value = data.xcsh_device_intelligence_device_history.example
}
```

<a id="canonical-2331320322000320-2220220020013311-0321133030223001-1000022003003311-1013111212320102-2321323013231113-1120002103301032-2211023220121202"></a>

## Root configuration — xcsh_device_intelligence_device_history / 030223001210 / 5

Required root properties: `device_id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0213332002231320-0310033203201203-1010223033120303-0030210222133121-1311101323002123-3333230112030100-2200322303100022-1011223321111313"></a>

## Next pages — xcsh_device_intelligence_device_history / 030223001210 / 6

- [Property reference](../guides/data-sources--device_intelligence_device_history--reference--group-001.md#canonical-1113221232103010-3013003312221121-0011022102212302-3131012031331210-0020303211131110-3123123112212302-2300102012010332-2102313110312030)
- [Examples](../guides/data-sources--device_intelligence_device_history--examples--group-001.md#canonical-0121110123011222-3011100330300013-2013323000321232-2021112210232322-2102202113313111-0020101210232210-0121133122300013-2022103323111221)
