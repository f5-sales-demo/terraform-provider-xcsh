---
page_title: "xcsh_device_intelligence_device_summary examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_summary examples."
---

# xcsh_device_intelligence_device_summary examples

<a id="canonical-3301113121302200-0313213210333013-2322301202203301-2033313012320132-1321030031102310-2131232223001111-0000023213230113-2300002211033113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-3320320012000030-1120213132323310-0033312123010313-0312323233302312-0102031103102222-2230101011132101-1321001213032302-3200123203212021)
- Examples

<a id="canonical-1302130132222120-3322330111203323-0223312132220013-2132020233313300-0330111323133122-2003020021310333-2013023213300321-1202200030031300"></a>

### Complete configurations for `xcsh_device_intelligence_device_summary`

- [Data source](data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-1233033221000300-2123033110213020-3220330330333213-3322202120233011-2233320031300311-2110222320011022-1223313122230133-0220223033032323): valid configuration.

<a id="canonical-1233033221000300-2123033110213020-3220330330333213-3322202120233011-2233320031300311-2110222320011022-1223313122230133-0220223033032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-3320320012000030-1120213132323310-0033312123010313-0312323233302312-0102031103102222-2230101011132101-1321001213032302-3200123203212021)
- [Examples](data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-3301113121302200-0313213210333013-2322301202203301-2033313012320132-1321030031102310-2131232223001111-0000023213230113-2300002211033113)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_device_summary/data-source.tf`; digest `sha256:04010440b0365c5df9d1e3244c5a3f2f792550e1e47b47ef7b1ff17bae9b7077`.

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
