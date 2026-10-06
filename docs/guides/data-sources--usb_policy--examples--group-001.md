---
page_title: "xcsh_usb_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy examples."
---

# xcsh_usb_policy examples

<a id="canonical-0102202312010132-3223213230021232-3100232232131311-1301000132200210-0022013310320310-1011103100213320-3113320310021033-0000003311301213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-3133321203222121-0333232330020202-2311221201233121-1211132233302312-0010122132101103-2031031312101012-3000023221320113-0202220320111000)
- Examples

<a id="canonical-2320010110222022-3102231301002332-2120020020103223-2300113130223021-3122220133022031-2031101031010113-2112311303012203-2010110221122302"></a>

### Complete configurations for `xcsh_usb_policy`

- [Data source](data-sources--usb_policy--examples--group-001.md#canonical-0233311010100122-3103202203302102-2223112121120010-2331221331101103-3002112021222331-3023103233320103-3032001201110233-1233133230012303): valid configuration.

<a id="canonical-0233311010100122-3103202203302102-2223112121120010-2331221331101103-3002112021222331-3023103233320103-3032001201110233-1233133230012303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-3133321203222121-0333232330020202-2311221201233121-1211132233302312-0010122132101103-2031031312101012-3000023221320113-0202220320111000)
- [Examples](data-sources--usb_policy--examples--group-001.md#canonical-0102202312010132-3223213230021232-3100232232131311-1301000132200210-0022013310320310-1011103100213320-3113320310021033-0000003311301213)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_usb_policy/data-source.tf`; digest `sha256:40aac30ea686e00cce92cecb2aee01cb69e14aef947ae725b686bf0b6e9af7b0`.

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
