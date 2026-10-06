---
page_title: "xcsh_usb_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy."
---

# xcsh_usb_policy

<a id="canonical-3133321203222121-0333232330020202-2311221201233121-1211132233302312-0010122132101103-2031031312101012-3000023221320113-0202220320111000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_usb_policy

Reads USB Policy information from F5 Distributed Cloud.

<a id="canonical-2122200012113103-3133321121322230-0210011200213313-1010122001123302-2102211220303222-3233211210123201-2211113323210230-0210213120032200"></a>

### Prerequisites for `xcsh_usb_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0122211320213132-1313220031221310-2222113203112011-2231220001230312-1023221230333100-0113003020201020-3123030130213300-3203232200012022"></a>

### Minimal configuration for `xcsh_usb_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UsbPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UsbPolicy by name
data "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}

output "usb_policy_id" {
  value = data.xcsh_usb_policy.example.id
}
```

<a id="canonical-2131133132233230-3333000312032113-1002112102232112-0133312303222311-2212023213010200-0111002020113231-2100301302131322-3201220100100123"></a>

### Root configuration for `xcsh_usb_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3331313012322112-1133333013032021-0101312102202222-3303113021011201-0231230031030023-1303032330200122-0100131010320230-3222121003012012"></a>

### Explore this collection for `xcsh_usb_policy`

- [Property reference](../guides/data-sources--usb_policy--reference--group-001.md#canonical-1000132301302331-1210321310212233-2332001022201121-3133211123011103-2123333120333133-1000202023301002-2333222023302121-1001000003033000)
- [Examples](../guides/data-sources--usb_policy--examples--group-001.md#canonical-0102202312010132-3223213230021232-3100232232131311-1301000132200210-0022013310320310-1011103100213320-3113320310021033-0000003311301213)
