---
page_title: "xcsh_device_intelligence_devices examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_devices examples."
---

# xcsh_device_intelligence_devices examples

<a id="canonical-3202213332101103-3233313320200230-0033320332100123-0103231020300312-3210010033102102-2213010331130131-3001113011213133-2113013001302102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-1321010130030011-2203222211110110-0002220011211330-1110301000223122-2213011310020223-2303303332130221-2212020001131033-0330011210030333)
- Examples

<a id="canonical-0222330232313100-3102313231112132-1021303030020203-0302313110220233-3133300130301321-1210013320202311-3331110133210003-0001010211003230"></a>

### Complete configurations for `xcsh_device_intelligence_devices`

- [Data source](data-sources--device_intelligence_devices--examples--group-001.md#canonical-2323220013020300-2331302230113033-2133211132113032-0103013320230332-0220103031101102-0201112211333231-1012111230223230-1210320000330233): valid configuration.

<a id="canonical-2323220013020300-2331302230113033-2133211132113032-0103013320230332-0220103031101102-0201112211333231-1012111230223230-1210320000330233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-1321010130030011-2203222211110110-0002220011211330-1110301000223122-2213011310020223-2303303332130221-2212020001131033-0330011210030333)
- [Examples](data-sources--device_intelligence_devices--examples--group-001.md#canonical-3202213332101103-3233313320200230-0033320332100123-0103231020300312-3210010033102102-2213010331130131-3001113011213133-2113013001302102)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_devices/data-source.tf`; digest `sha256:59c1a13d7be54a7427ce8e310798e9737dbdf20c08bef5ed916ec4310ce79e92`.

```terraform
# DeviceIntelligenceDevices DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_devices" "example" {
  namespace = "example-value"
}

output "device_intelligence_devices_result" {
  value = data.xcsh_device_intelligence_devices.example
}
```
