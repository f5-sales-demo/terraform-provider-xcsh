---
page_title: "xcsh_infraprotect_mitigation_ips"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_mitigation_ips."
---

# xcsh_infraprotect_mitigation_ips

<a id="canonical-1223233033223000-1032001012013022-1301320302130033-2311101100202302-3130231002213230-3320300113130330-0113131111131313-3222111203302102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_infraprotect_mitigation_ips

Reads Infraprotect Mitigation Ips information from F5 Distributed Cloud.

<a id="canonical-0333300311102013-1321012212201103-3122023122331332-3010102120322020-0113330303223212-0221010330121020-3313331021321122-0223332030332320"></a>

### Prerequisites for `xcsh_infraprotect_mitigation_ips`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0221303010111213-0112020211231331-1100202232311223-1313331113110320-1102002112111230-3233201322301223-2221000312020000-0001001110001331"></a>

### Minimal configuration for `xcsh_infraprotect_mitigation_ips`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```

<a id="canonical-3132200300031133-1121202223302301-3222001001101233-3131121001310001-3003233301200122-2110021103323001-2223333011220103-0200100202120031"></a>

### Root configuration for `xcsh_infraprotect_mitigation_ips`

Required root properties: `mitigation_id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3133112011130320-1120112112122301-1333213210210300-0103303111101211-3001230131030133-2031002233220132-3311221111113033-1113232311230133"></a>

### Explore this collection for `xcsh_infraprotect_mitigation_ips`

- [Property reference](../guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-3211003101332320-2011321310020021-1301100203211023-1330101230021031-1102102332023332-0031020312103202-2022313110031312-3331302220003310)
- [Examples](../guides/data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-2230332003132330-3003232213123221-1023033101300311-0101332232032132-1020000311020311-2102222232300101-2130223103131322-2223023103332200)
