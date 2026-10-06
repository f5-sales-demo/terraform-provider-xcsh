---
page_title: "xcsh_cloud_user_account examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account examples."
---

# xcsh_cloud_user_account examples

<a id="canonical-1123022031022212-3210322222220211-3221011130302320-0003311301312322-2122213221222103-0021320102020113-1330120011222112-0023310231123313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- Examples

<a id="canonical-3022323120031221-0202211121322030-2332110132310101-0113023002330300-0323133123211100-3203101123000203-2230331312312223-1231331313120132"></a>

### Complete configurations for `xcsh_cloud_user_account`

- [Data source](data-sources--cloud_user_account--examples--group-001.md#canonical-0320012310310023-2232332310013003-0010130032332200-0303302101222303-2002231232020011-1311312300232133-2201330011113332-3321030110103222): valid configuration.

<a id="canonical-0320012310310023-2232332310013003-0010130032332200-0303302101222303-2002231232020011-1311312300232133-2201330011113332-3321030110103222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md#canonical-0121211201030010-3301133113320101-0332001123222101-1302203102030320-0332022000333203-1320110312000030-2303311122303231-1222130130102021)
- [Examples](data-sources--cloud_user_account--examples--group-001.md#canonical-1123022031022212-3210322222220211-3221011130302320-0003311301312322-2122213221222103-0021320102020113-1330120011222112-0023310231123313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_user_account/data-source.tf`; digest `sha256:4bd66f34c8654fcf2e99224154cac4393ee6bb933699a10ecdb50c1d8d402e99`.

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
