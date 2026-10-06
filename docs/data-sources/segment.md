---
page_title: "xcsh_segment"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment."
---

# xcsh_segment

<a id="canonical-1322202313211133-3232200032023212-3030131113213211-1232130001132103-2232301233230330-2022002332021123-3031321121233232-3101311003121100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_segment

Reads Segment information from F5 Distributed Cloud.

<a id="canonical-1100330210032022-2031031120110020-2301213303032022-0030120201112331-2130031330231012-0321011103333113-1111310023210120-3013320311000020"></a>

### Prerequisites for `xcsh_segment`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0222131030303110-0223320112211210-2200133021001313-1220011033101322-1331300332313311-0100000230303122-2223100002021022-1133030020323311"></a>

### Minimal configuration for `xcsh_segment`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Segment Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Segment by name
data "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.xcsh_segment.example.id
}
```

<a id="canonical-1023120033220022-1001003123230122-1222023031300133-0122132131321010-2200003310010023-2210001002202201-1123010222222210-2220232230103233"></a>

### Root configuration for `xcsh_segment`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2230133003100332-0232020320331033-0130211011232212-2223133201221331-1232030300133110-3121220012013011-1030300312200020-2310221121032012"></a>

### Explore this collection for `xcsh_segment`

- [Property reference](../guides/data-sources--segment--reference--group-001.md#canonical-1110023031331322-0210111201230333-1322101312131120-3013310220132303-3223212001102311-1012321231332121-2223100100323021-0213000200231233)
- [Examples](../guides/data-sources--segment--examples--group-001.md#canonical-0210203003102130-2132223030332301-3223100302300300-3131312102123022-1311222120010123-0220311130321330-1300000003331333-1121130121113311)
