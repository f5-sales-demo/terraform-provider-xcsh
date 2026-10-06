---
page_title: "xcsh_usb_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy."
---

# xcsh_usb_policy

<a id="canonical-3313033212003300-1002332031103020-1323001313000200-3003230033322232-3111031201110211-0323302100321032-3212002010220013-3201110212312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_usb_policy

Manages new USB policy object in F5 Distributed Cloud.

<a id="canonical-3102130032010333-2111001202012220-3223120320000321-2312312322033002-0201122133012132-1033222312310230-3113310022223103-3101033320213123"></a>

### Prerequisites for `xcsh_usb_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2003201133321101-3113221303202021-2101122322222131-3101300012032303-1310301301033313-3221122032222203-1112132323102032-1222303003332223"></a>

### Minimal configuration for `xcsh_usb_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2332010100302012-2313032213213231-2012131023221103-3023022200122122-3132121120020011-2003132200003302-0111120211113213-1332121202123120"></a>

### Root configuration for `xcsh_usb_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0220022130131111-2120210113210230-0321033103132101-3201000101221013-1211121330010001-1211222122220323-3122003002312021-3212111320131101"></a>

### Explore this collection for `xcsh_usb_policy`

- [Property reference](../guides/resources--usb_policy--reference--group-001.md#canonical-0203202331120213-3031223102221310-3303300333312020-2233220001321002-3220003200013332-1202131202021320-0033221202012231-3201003230011131)
- [Examples](../guides/resources--usb_policy--examples--group-001.md#canonical-0122312222211332-0220022100110220-0120201023100020-1023011303301300-2031011010322100-1311132210211111-3030213103001222-2101322302130302)
- [Import](../guides/resources--usb_policy--lifecycle--group-001.md#canonical-1111023303231103-3021112330033232-2131012323020111-3311313110103332-2032122300013112-3111103113310103-0003300013211111-0033330013131102)
- [Timeouts](../guides/resources--usb_policy--lifecycle--group-001.md#canonical-0011001131103310-1310111000023110-3131311212123012-3020022323100012-3110201320013200-2303131021101101-3012120322230023-2332322131022122)
