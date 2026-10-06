---
page_title: "xcsh_cdn_purge_command"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command."
---

# xcsh_cdn_purge_command

<a id="canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_cdn_purge_command

Reads CDN Purge Command information from F5 Distributed Cloud.

<a id="canonical-0202201201202120-1331012100102100-0301211313132323-3221000221202101-3112012220311313-1311211010110200-1122122331311030-0330102133113311"></a>

### Prerequisites for `xcsh_cdn_purge_command`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3120123212001313-2320021011023030-0133300010303021-1233101322211001-3232331003102010-1100033023203003-2001230231232303-1211111230312233"></a>

### Minimal configuration for `xcsh_cdn_purge_command`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNPurgeCommand Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNPurgeCommand by name
data "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}

output "cdn_purge_command_id" {
  value = data.xcsh_cdn_purge_command.example.id
}
```

<a id="canonical-0222001122231202-1002020112101202-3313110303323021-0111012311230203-3012011333023331-2230202203303211-3200021332202131-2312210210131031"></a>

### Root configuration for `xcsh_cdn_purge_command`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2223113102323303-2030323111123210-1113011300332301-0021033232221213-0021303103321022-1311000322131101-0223103232111223-3233220331113100"></a>

### Explore this collection for `xcsh_cdn_purge_command`

- [Property reference](../guides/data-sources--cdn_purge_command--reference--group-001.md#canonical-2333230322013111-3233211221333000-1330202200011321-1003000031212213-1200303112020123-3333032203113133-3010313323203233-1003330132300032)
- [Examples](../guides/data-sources--cdn_purge_command--examples--group-001.md#canonical-2122121333200302-2331113010221332-2131222102331332-1212022030030013-0300301001103103-1111302312103002-1313003100101122-0210003320133232)
