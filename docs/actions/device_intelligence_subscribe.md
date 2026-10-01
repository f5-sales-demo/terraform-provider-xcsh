---
page_title: "xcsh_device_intelligence_subscribe landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_subscribe landing."
---

# xcsh_device_intelligence_subscribe landing

<a id="canonical-2110103333002230-1333113312300003-3031011200033223-1012111303330101-3123021221010332-0102310002201001-2232032323131011-1213333131130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232031212132300-2100011321122300-3232002033103330-0220002103001102-1332210020211301-3312101030032312-1210021123131321-0102230323122303"></a>

## xcsh_device_intelligence_subscribe — xcsh_device_intelligence_subscribe / 312230100122 / 2

Breadcrumbs:

- xcsh_device_intelligence_subscribe

Resource creation operation.

<a id="canonical-2013112111312232-0332120110212322-0331313131022002-0111230303321113-2311332220001030-1010030033231223-0201221223211000-3313131100032022"></a>

## Prerequisites — xcsh_device_intelligence_subscribe / 312230100122 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3113023330321112-1333103320110213-3300002310110023-3233002331323122-1013202010102020-1231322020032210-2311031202033230-0131122221113022"></a>

## Minimal configuration — xcsh_device_intelligence_subscribe / 312230100122 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceSubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_subscribe" "example" {
  config {
  }
}
```

<a id="canonical-0231022321033311-0303321012202003-3330311232132300-3120021331121003-1020313030010320-2002331200022332-3003022033021231-2222322220102330"></a>

## Root configuration — xcsh_device_intelligence_subscribe / 312230100122 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2100233213220121-3123021023220323-3301221332300203-3210131120321030-0223202302000011-3220223313101123-2011110001322213-2222112320320101"></a>

## Next pages — xcsh_device_intelligence_subscribe / 312230100122 / 6

- [Property reference](../guides/actions--device_intelligence_subscribe--reference--group-001.md#canonical-0320320323021310-3230032032311001-0203030313030101-1313303123010120-1101213323320030-2222312112021322-0330231033330113-3123031132132130)
- [Examples](../guides/actions--device_intelligence_subscribe--examples--group-001.md#canonical-0202210202220322-0020220331121231-2302002000323023-2013002300000311-3121211102111231-1323110303301310-2322132301002230-2100310210332112)
- [Lifecycle](../guides/actions--device_intelligence_subscribe--lifecycle--group-001.md#canonical-3310123332001021-2012232003331012-0330203330033101-3131300122000023-3210032102323112-2301332110321212-1232022122223013-1133320301033031)
