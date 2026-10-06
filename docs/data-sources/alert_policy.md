---
page_title: "xcsh_alert_policy"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy."
---

# xcsh_alert_policy

<a id="canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_alert_policy

Reads Alert Policy information from F5 Distributed Cloud.

<a id="canonical-0032310100001320-0003100320332330-0321131313303332-0322000120120102-0231121202220123-0333203011222231-0320112102313132-3131322002221233"></a>

### Prerequisites for `xcsh_alert_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-2023130201312201-0121213302231332-1001331310130212-1011331130222110-1333011000210200-1012020030300313-2301302110200300-2113201310230323"></a>

### Minimal configuration for `xcsh_alert_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertPolicy by name
data "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}

output "alert_policy_id" {
  value = data.xcsh_alert_policy.example.id
}
```

<a id="canonical-0201201200120122-3010133112023031-0313001000123130-1203011312320321-3320113312223032-2121210330202013-3122001102311233-3132011211103223"></a>

### Root configuration for `xcsh_alert_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1033221110330022-3313202003101233-1120032202210130-1030013023213033-1201021211201312-1213100312230320-1321030130011202-1312032021311131"></a>

### Explore this collection for `xcsh_alert_policy`

- [Property reference](../guides/data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [Examples](../guides/data-sources--alert_policy--examples--group-001.md#canonical-0331200133120203-2223103130311131-1313003123311020-2320233233111301-0201023120200322-1100113321132123-0333222333022022-0211023101101221)
