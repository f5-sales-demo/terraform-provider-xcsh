---
page_title: "xcsh_token"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token."
---

# xcsh_token

<a id="canonical-2000130020133110-1113012332320202-2223121002312130-2001210311231111-3323332201023323-2311120322030210-1131312201010133-0322033200102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_token

Reads a site admission token used during registration and provisioning.

<a id="canonical-2210031320213220-2221033201122022-0231022103102103-2200332230112113-3013201322023232-1032331013232200-3222322323300200-2001222130323033"></a>

### Prerequisites for `xcsh_token`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-0011132222101320-1103133131213002-1313311013002301-3101330013302313-3003303010203000-0330022121032303-2220231223020023-1221111012223132"></a>

### Minimal configuration for `xcsh_token`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Token Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Token by name
data "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
}

output "token_id" {
  value = data.xcsh_token.example.id
}
```

<a id="canonical-3321022300220213-0310231111113313-2331010311010010-0120031133101011-1022102213220130-1312320333321022-1122312013130002-3333212023122201"></a>

### Root configuration for `xcsh_token`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1100332231111133-2023203330203111-1113212203011003-1333121111133033-2030102311313331-3231012103102213-0232013200022001-3021323333111201"></a>

### Explore this collection for `xcsh_token`

- [Property reference](../guides/data-sources--token--reference--group-001.md#canonical-1302021003003321-1033031123003102-3221300113210300-3012332223200031-1300210333011012-3120203201330301-0001020001012101-1203222133001122)
- [Examples](../guides/data-sources--token--examples--group-001.md#canonical-1020321111222301-2033123121331132-3331211233301231-1202000032310201-3201211232323022-2130002310311310-3103122312103310-0031001213131320)
