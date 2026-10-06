---
page_title: "xcsh_device_intelligence_multi_account_devices examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_multi_account_devices examples."
---

# xcsh_device_intelligence_multi_account_devices examples

<a id="canonical-2103201103310120-0321301032000010-0223230302330212-3120120201333300-0132210303031113-1201231000011212-2213131003022303-3203003132131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000)
- Examples

<a id="canonical-0201123003101011-1113232311020301-3221123222223223-2322022203232022-1133110131311323-2312010320010322-2321012300100131-3332110112120031"></a>

### Complete configurations for `xcsh_device_intelligence_multi_account_devices`

- [Data source](data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-1300111002212133-0100000022222012-3220100313023223-2121110011022211-1202021322301122-2203213212102321-1120311111330033-0122320021123312): valid configuration.

<a id="canonical-1300111002212133-0100000022222012-3220100313023223-2121110011022211-1202021322301122-2203213212102321-1120311111330033-0122320021123312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000)
- [Examples](data-sources--device_intelligence_multi_account_devices--examples--group-001.md#canonical-2103201103310120-0321301032000010-0223230302330212-3120120201333300-0132210303031113-1201231000011212-2213131003022303-3203003132131320)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_multi_account_devices/data-source.tf`; digest `sha256:83db70480833f20bc19383610323bcf199f0707dd687a03af490d0098255cfcb`.

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
