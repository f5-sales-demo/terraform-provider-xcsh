---
page_title: "xcsh_access_active_sessions"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions."
---

# xcsh_access_active_sessions

<a id="canonical-3120002103123020-1021131301111100-3330103332001301-2030202210211031-3012203032110100-2013230030123100-0331113020331110-1012323022312133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_access_active_sessions

Reads active access sessions information from F5 Distributed Cloud.

<a id="canonical-2211323111311113-0110000320303331-0111002312302022-1301330131322211-3111212313030321-1103301211113213-3213212022222223-2131222230301231"></a>

### Prerequisites for `xcsh_access_active_sessions`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0130203102022102-3321103010213032-3110131333130013-0101022101013001-2001203312310201-3210033200122102-3130133333020102-3113200323023103"></a>

### Minimal configuration for `xcsh_access_active_sessions`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_sessions" "example" {
  namespace = "example-value"
}

output "access_active_sessions_result" {
  value = data.xcsh_access_active_sessions.example
}
```

<a id="canonical-1002000322220300-3331310022310123-0003310122302312-2032210331002032-3323112323103033-3220301333132031-2103030033222023-3122110321301202"></a>

### Root configuration for `xcsh_access_active_sessions`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0011102300300112-2333123301230021-3313211333032303-0013221122100203-0300112220132102-3022201201213210-1132110103033102-3111203003110020"></a>

### Explore this collection for `xcsh_access_active_sessions`

- [Property reference](../guides/data-sources--access_active_sessions--reference--group-001.md#canonical-3222033320220301-0321022220333230-1101302312230022-0302301020210223-1300232233122230-0011131023032013-1333300203311223-1320201211302123)
- [Examples](../guides/data-sources--access_active_sessions--examples--group-001.md#canonical-2130122221103112-3022213203212002-0223130112330320-2112021120223032-1012230202021021-3012000133320332-2301032201010003-1302301020010310)
