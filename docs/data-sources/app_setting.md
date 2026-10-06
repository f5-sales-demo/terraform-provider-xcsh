---
page_title: "xcsh_app_setting"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting."
---

# xcsh_app_setting

<a id="canonical-3100101210131302-0032023222032213-0211302130230231-2121201111122113-3321222320011111-1020123133110021-0000133011122133-0323230333123021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_app_setting

Reads App Setting information from F5 Distributed Cloud.

<a id="canonical-0020302222220032-0313000323300113-2003020020231222-1121012133330302-3123002103033313-2031213022012332-0311113101210210-0133030332020103"></a>

### Prerequisites for `xcsh_app_setting`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2323331103110030-2212220303001102-3221312013310100-3321310322010301-0310312033300332-2201202310332011-3102323033120223-0223112023020020"></a>

### Minimal configuration for `xcsh_app_setting`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSetting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppSetting by name
data "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}

output "app_setting_id" {
  value = data.xcsh_app_setting.example.id
}
```

<a id="canonical-1331232021003031-1330310101301221-0001133222022111-1012122313101101-1201130132312012-3313132332220323-1003212203032110-1112113203011310"></a>

### Root configuration for `xcsh_app_setting`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2033133120131300-1331010223303013-0222101011302012-0123320213222322-0123012103230000-0020330013321222-1311200110121331-0133232133121300"></a>

### Explore this collection for `xcsh_app_setting`

- [Property reference](../guides/data-sources--app_setting--reference--group-001.md#canonical-3030321030220121-1133302033120230-3303020323031023-3130002021022220-1312302212132220-1002031103221223-2303312102023000-0130231011221303)
- [Examples](../guides/data-sources--app_setting--examples--group-001.md#canonical-2202211201021322-2102221330322113-3202223232211020-1000321030231200-0020132132111010-1232020333012302-0023302221031020-1233033303312331)
