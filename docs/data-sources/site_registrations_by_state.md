---
page_title: "xcsh_site_registrations_by_state"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_state."
---

# xcsh_site_registrations_by_state

<a id="canonical-1011310232230021-3303123312003111-3120023223111201-0320322111010320-3032200132321200-3310101021222131-0120032131300331-0201133223313303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_site_registrations_by_state

List Customer Edge registrations by state.

<a id="canonical-0313121301302100-1201002013321030-3030021010131303-3133100130133033-2312121022031303-0102122301123020-3201223011303201-1201301300032230"></a>

### Prerequisites for `xcsh_site_registrations_by_state`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0303111202110122-0011012010020133-0033310202020220-2313323103010312-1120333120112130-3033121333132022-2001230302101110-3303102110132302"></a>

### Minimal configuration for `xcsh_site_registrations_by_state`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrationsByState DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_state" "example" {
  state = "NOTSET"
}

output "site_registrations_by_state_result" {
  value = data.xcsh_site_registrations_by_state.example
}
```

<a id="canonical-2131031330131222-1230012122230023-0001213200310232-0320312123233322-0012211222130100-2223300011111102-0313200310222213-0020100323031001"></a>

### Root configuration for `xcsh_site_registrations_by_state`

Required root properties: `state`. Full root flags and choices appear in the property reference.

<a id="canonical-0232230112130110-2210033202322032-1321131331302323-2130313212201133-3232230131000321-1330322031100303-0322020302303322-2021320011003002"></a>

### Explore this collection for `xcsh_site_registrations_by_state`

- [Property reference](../guides/data-sources--site_registrations_by_state--reference--group-001.md#canonical-0320011213231301-1102033233211220-0302131120233130-2122121213220203-2033211132132211-0333231313221113-3210300123023021-2021122221230130)
- [Examples](../guides/data-sources--site_registrations_by_state--examples--group-001.md#canonical-3303201010310230-2003223220001303-1101321302202101-2030220322223133-1333332012121232-2102231003300003-1210020331110123-0320113132311303)
