---
page_title: "xcsh_device_intelligence_device_history examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_history examples."
---

# xcsh_device_intelligence_device_history examples

<a id="canonical-0121110123011222-3011100330300013-2013323000321232-2021112210232322-2102202113313111-0020101210232210-0121133122300013-2022103323111221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-1033022013313010-2230031200002233-1320022023130030-0022221001122220-0201311013330313-2011320120303230-1323223312110203-0123022120001321)
- Examples

<a id="canonical-1121311310303121-0100302220232123-1230031203230213-2220021013113231-2201222301132030-3200130203202233-0330311113222313-2213233220212201"></a>

### Complete configurations for `xcsh_device_intelligence_device_history`

- [Data source](data-sources--device_intelligence_device_history--examples--group-001.md#canonical-1130232223300130-2002111313021222-2201321310031221-0112313123310100-3221002130323312-3203003300110200-0003033003323222-3313202022112122): valid configuration.

<a id="canonical-1130232223300130-2002111313021222-2201321310031221-0112313123310100-3221002130323312-3203003300110200-0003033003323222-3313202022112122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-1033022013313010-2230031200002233-1320022023130030-0022221001122220-0201311013330313-2011320120303230-1323223312110203-0123022120001321)
- [Examples](data-sources--device_intelligence_device_history--examples--group-001.md#canonical-0121110123011222-3011100330300013-2013323000321232-2021112210232322-2102202113313111-0020101210232210-0121133122300013-2022103323111221)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_device_history/data-source.tf`; digest `sha256:f95f8328fecc18ee316fb7732f60f6fec73cc07388f5c6067f1bad91b8ac0d75`.

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
