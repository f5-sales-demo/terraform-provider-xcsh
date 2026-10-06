---
page_title: "xcsh_device_intelligence_unsubscribe examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_unsubscribe examples."
---

# xcsh_device_intelligence_unsubscribe examples

<a id="canonical-3330230022100101-3210012001111211-1111132012321230-0332303032022331-0112322033322021-2100112032212020-0103010331223003-3102323133020132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_unsubscribe](../actions/device_intelligence_unsubscribe.md#canonical-3110012200313320-0120221200221312-0230333311202000-0331023323320332-1101121132213210-3031113003001131-3201203222213110-2210102303020312)
- Examples

<a id="canonical-2000020113112021-0203320200311121-3030210313121110-3130210300113133-3103320230312222-2122202320130033-1200110210332320-2130111010320311"></a>

### Complete configurations for `xcsh_device_intelligence_unsubscribe`

- [Action](actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-3311111003211210-3123212122310002-3232321300010031-3223330211231201-1201330003122032-3133031110321323-1300222223323000-3201321333320212): valid configuration.

<a id="canonical-3311111003211210-3123212122310002-3232321300010031-3223330211231201-1201330003122032-3133031110321323-1300222223323000-3201321333320212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_device_intelligence_unsubscribe](../actions/device_intelligence_unsubscribe.md#canonical-3110012200313320-0120221200221312-0230333311202000-0331023323320332-1101121132213210-3031113003001131-3201203222213110-2210102303020312)
- [Examples](actions--device_intelligence_unsubscribe--examples--group-001.md#canonical-3330230022100101-3210012001111211-1111132012321230-0332303032022331-0112322033322021-2100112032212020-0103010331223003-3102323133020132)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_device_intelligence_unsubscribe/action.tf`; digest `sha256:391d15fe0fb026a460e8e78fbc03f16af6138839116ddca4fd7a4862c6cad9eb`.

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
