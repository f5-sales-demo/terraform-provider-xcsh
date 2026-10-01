---
page_title: "xcsh_device_intelligence_unsubscribe landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_unsubscribe landing."
---

# xcsh_device_intelligence_unsubscribe landing

<a id="canonical-3110012200313320-0120221200221312-0230333311202000-0331023323320332-1101121132213210-3031113003001131-3201203222213110-2210102303020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331101300330231-0122020321031012-3302202112102101-3130030103303010-0333223212031120-2122333002020332-1200323032020013-1010003313131301"></a>

## xcsh_device_intelligence_unsubscribe — xcsh_device_intelligence_unsubscribe / 320120221002 / 2

Breadcrumbs:

- xcsh_device_intelligence_unsubscribe

Resource creation operation.

<a id="canonical-3032330330020132-3100332102203012-1202320233311311-1113331302312101-3212203232231011-1122332230002331-2122310102132313-2010000310220112"></a>

## Prerequisites — xcsh_device_intelligence_unsubscribe / 320120221002 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1321223120023013-3302033312120033-1321310121203301-0131122211221001-0322101233102302-2010232002303003-0231330220130011-3111303013001001"></a>

## Minimal configuration — xcsh_device_intelligence_unsubscribe / 320120221002 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceUnsubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_unsubscribe" "example" {
  config {
  }
}
```

<a id="canonical-1222102313003221-3030211101031122-1022033301131013-1320012100121132-0331333011133211-3300313123303220-0132300232311202-0113332221120322"></a>

## Root configuration — xcsh_device_intelligence_unsubscribe / 320120221002 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-3132311022323300-3220023131221321-2213302102321132-1003120301022323-1002233123231122-3021012332221013-3233232311213023-3221130212023311"></a>

## Next pages — xcsh_device_intelligence_unsubscribe / 320120221002 / 6

- [Property reference](../guides/actions--device_intelligence_unsubscribe--reference--group-001.md#canonical-3010223032203133-1132012221011020-0133203221102333-3300020301000200-3000333120111101-1121230320222222-3031001322003213-2200112221221121)
- [Examples](../guides/actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-3330230022100101-3210012001111211-1111132012321230-0332303032022331-0112322033322021-2100112032212020-0103010331223003-3102323133020132)
- [Lifecycle](../guides/actions--device_intelligence_unsubscribe--lifecycle--group-001.md#canonical-3033213233002220-0133011323120233-2031032000033131-0012000021211031-0313113131231322-0110322103213213-2023032001211113-2313302102122233)
