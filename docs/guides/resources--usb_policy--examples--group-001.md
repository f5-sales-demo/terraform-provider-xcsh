---
page_title: "xcsh_usb_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy examples."
---

# xcsh_usb_policy examples

<a id="canonical-0122312222211332-0220022100110220-0120201023100020-1023011303301300-2031011010322100-1311132210211111-3030213103001222-2101322302130302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-3313033212003300-1002332031103020-1323001313000200-3003230033322232-3111031201110211-0323302100321032-3212002010220013-3201110212312232)
- Examples

<a id="canonical-0322203000033021-3130002010203313-0032031101011221-1022222013032332-3102103330331020-3011311303111023-3200122301112202-1220021122020212"></a>

### Complete configurations for `xcsh_usb_policy`

- [Resource](resources--usb_policy--examples--group-001.md#canonical-1021311230322121-1223221313333000-3302100222013302-0322101321332312-2111102333322331-3020210003100310-1200012222211011-1102301103010312): valid configuration.

<a id="canonical-1021311230322121-1223221313333000-3302100222013302-0322101321332312-2111102333322331-3020210003100310-1200012222211011-1102301103010312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-3313033212003300-1002332031103020-1323001313000200-3003230033322232-3111031201110211-0323302100321032-3212002010220013-3201110212312232)
- [Examples](resources--usb_policy--examples--group-001.md#canonical-0122312222211332-0220022100110220-0120201023100020-1023011303301300-2031011010322100-1311132210211111-3030213103001222-2101322302130302)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_usb_policy/resource.tf`; digest `sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e`.

```terraform
# UsbPolicy Resource Example
# Manages new USB policy object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UsbPolicy configuration
resource "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}
```
