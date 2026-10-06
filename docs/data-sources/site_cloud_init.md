---
page_title: "xcsh_site_cloud_init"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_cloud_init."
---

# xcsh_site_cloud_init

<a id="canonical-2323200302110331-3032331300132323-0030000003201111-2021230223222112-1331220322311201-1311121001102222-2100011132101132-0331131132130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_site_cloud_init

Retrieve Customer Edge cloud-init template.

<a id="canonical-3333130331310232-0223210111130203-3320031330012022-3103030130323002-1012130310210302-0333011000112213-2000012032232213-2231302213030000"></a>

### Prerequisites for `xcsh_site_cloud_init`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3131311113033323-0332333310001233-3031332132122130-3223002012310123-0013022020331122-2021112211020100-0013022001033020-3103331012202021"></a>

### Minimal configuration for `xcsh_site_cloud_init`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteCloudInit DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_cloud_init" "example" {
  provider_ref = "example-value"
  site_name    = "example-value"
}

output "site_cloud_init_result" {
  value     = data.xcsh_site_cloud_init.example
  sensitive = true
}
```

<a id="canonical-0123322002033031-1132110003010031-2310322201202230-3030221313300032-1302212222113310-0310312121012120-0021100120203333-1021321300220013"></a>

### Root configuration for `xcsh_site_cloud_init`

Required root properties: `provider_ref`, `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-0010003102300020-3220311321202012-2210110331121200-2330322233232032-3231013022313003-3131003231331031-1233200113233320-3020021001201222"></a>

### Explore this collection for `xcsh_site_cloud_init`

- [Property reference](../guides/data-sources--site_cloud_init--reference--group-001.md#canonical-1010210110123221-2301023030310012-3131213301230313-0212223132222233-1001200020011312-1122011223101313-1011020230110010-0030103221003122)
- [Examples](../guides/data-sources--site_cloud_init--examples--group-001.md#canonical-0222330133321102-3332331012120322-1322100023222212-3112032330330131-1123001012030022-0120223001321210-3030113310202003-3331003312302210)
