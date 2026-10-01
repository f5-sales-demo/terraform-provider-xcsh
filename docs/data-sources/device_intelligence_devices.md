---
page_title: "xcsh_device_intelligence_devices landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_devices landing."
---

# xcsh_device_intelligence_devices landing

<a id="canonical-1321010130030011-2203222211110110-0002220011211330-1110301000223122-2213011310020223-2303303332130221-2212020001131033-0330011210030333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012301031211012-3322301003232000-0200133220002233-2132332123331013-2213112011122030-1312022232011203-1102220213303102-2312002303311201"></a>

## xcsh_device_intelligence_devices — xcsh_device_intelligence_devices / 032032212101 / 2

Breadcrumbs:

- xcsh_device_intelligence_devices

Resource creation operation.

<a id="canonical-3021303130222120-2112112313023212-2111230102031232-3203201330002232-3100311131302303-3313312000003002-1033303022112202-2010013102300310"></a>

## Prerequisites — xcsh_device_intelligence_devices / 032032212101 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1212003130320201-0213201231331030-0000010310333300-3132030322321303-3110020300230020-0102330213121331-0333023113002001-1212323321022201"></a>

## Minimal configuration — xcsh_device_intelligence_devices / 032032212101 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3030320022101233-1323221033322111-1002013123233302-0220022200121020-1002223231102222-2211310130302032-3021331020211303-0020300103323300"></a>

## Root configuration — xcsh_device_intelligence_devices / 032032212101 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1220123023100212-3130032000112222-2203130222232331-2123000210233010-3311010003132202-1011230011232032-1300230033000003-3332210010233011"></a>

## Next pages — xcsh_device_intelligence_devices / 032032212101 / 6

- [Property reference](../guides/data-sources--device_intelligence_devices--reference--group-001.md#canonical-1010221202313233-3310120333232013-0103200231213012-2012130133200323-1120003102201202-2222210221033201-2030231010113300-1231012332201303)
- [Examples](../guides/data-sources--device_intelligence_devices--examples--group-001.md#canonical-3202213332101103-3233313320200230-0033320332100123-0103231020300312-3210010033102102-2213010331130131-3001113011213133-2113013001302102)
