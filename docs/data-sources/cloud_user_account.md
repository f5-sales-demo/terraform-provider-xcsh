---
page_title: "xcsh_cloud_user_account"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account."
---

# xcsh_cloud_user_account

<a id="canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_cloud_user_account

Reads Cloud User Account information from F5 Distributed Cloud.

<a id="canonical-2313320133003030-2120011021303333-3103131000032311-0010012133012321-3022002212111220-1122113332232031-2321231011013322-1232030322232203"></a>

### Prerequisites for `xcsh_cloud_user_account`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0222100221122300-2130303311103323-0311132220001022-2023123301120003-1332200111101133-2323202122003331-0211001302130131-2120212223323131"></a>

### Minimal configuration for `xcsh_cloud_user_account`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudUserAccount Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudUserAccount by name
data "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}

output "cloud_user_account_id" {
  value = data.xcsh_cloud_user_account.example.id
}
```

<a id="canonical-1332311133231312-3100213021000221-2230302011321300-0320122113203032-3232302232232210-3012113121333012-3102221301130013-3110103001122302"></a>

### Root configuration for `xcsh_cloud_user_account`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2211000322030122-1210232111021003-3322213322001230-3321021031213001-3233022213130130-3133031303230112-2211130332132101-3133333001212013"></a>

### Explore this collection for `xcsh_cloud_user_account`

- [Property reference](../guides/data-sources--cloud_user_account--reference--group-001.md#canonical-2100003020103021-2211222133012112-2230221333311312-3001210132102132-0320331210120102-2221113213322121-0212112321020330-1011211032013131)
- [Examples](../guides/data-sources--cloud_user_account--examples--group-001.md#canonical-1123022031022212-3210322222220211-3221011130302320-0003311301312322-2122213221222103-0021320102020113-1330120011222112-0023310231123313)
