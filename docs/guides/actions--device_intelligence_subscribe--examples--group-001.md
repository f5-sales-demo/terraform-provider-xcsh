---
page_title: "xcsh_device_intelligence_subscribe examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_subscribe examples."
---

# xcsh_device_intelligence_subscribe examples

<a id="canonical-0202210202220322-0020220331121231-2302002000323023-2013002300000311-3121211102111231-1323110303301310-2322132301002230-2100310210332112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_subscribe](../actions/device_intelligence_subscribe.md#canonical-2110103333002230-1333113312300003-3031011200033223-1012111303330101-3123021221010332-0102310002201001-2232032323131011-1213333131130320)
- Examples

<a id="canonical-2201012311013002-2013322303103132-3333012110020323-1323311231113121-2022222311012133-2030221323111011-1200320213331330-2302000230000311"></a>

### Complete configurations for `xcsh_device_intelligence_subscribe`

- [Action](actions--device_intelligence_subscribe--examples--group-001.md#canonical-1121302300110102-3021032313203020-3332332311011113-1011031011003200-2221123022001303-1323323030121002-3221012303212200-3013111111113110): valid configuration.

<a id="canonical-1121302300110102-3021032313203020-3332332311011113-1011031011003200-2221123022001303-1323323030121002-3221012303212200-3013111111113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_device_intelligence_subscribe](../actions/device_intelligence_subscribe.md#canonical-2110103333002230-1333113312300003-3031011200033223-1012111303330101-3123021221010332-0102310002201001-2232032323131011-1213333131130320)
- [Examples](actions--device_intelligence_subscribe--examples--group-001.md#canonical-0202210202220322-0020220331121231-2302002000323023-2013002300000311-3121211102111231-1323110303301310-2322132301002230-2100310210332112)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_device_intelligence_subscribe/action.tf`; digest `sha256:f3174d51768147466c8e2454428347852090bb677186e716ca6cee429d892fbc`.

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
