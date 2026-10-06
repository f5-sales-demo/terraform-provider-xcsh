---
page_title: "xcsh_proxy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy."
---

# xcsh_proxy

<a id="canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_proxy

Reads Proxy information from F5 Distributed Cloud.

<a id="canonical-0203303013231111-1213232113133101-3112223320011001-2012310022232233-1030303301110131-3201123033003233-0120222000002021-0203332331333220"></a>

### Prerequisites for `xcsh_proxy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3220011332203033-3301330303111303-3030132102310100-0000110131331031-3002002220031212-1203230213330001-3020103003330210-2100333312233230"></a>

### Minimal configuration for `xcsh_proxy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Proxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Proxy by name
data "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}

output "proxy_id" {
  value = data.xcsh_proxy.example.id
}
```

<a id="canonical-3002101112001201-2012010022332213-2202030233332211-2310210212321102-0102210212323102-0313311330221332-0333303231123201-2333313021311133"></a>

### Root configuration for `xcsh_proxy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1300120111201230-1011013210230022-3032003333201311-3012023203231213-1210020233301023-3010322322032210-3011113203232021-2102111123302130"></a>

### Explore this collection for `xcsh_proxy`

- [Property reference](../guides/data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [Examples](../guides/data-sources--proxy--examples--group-001.md#canonical-3101102230333321-3121202112111310-2102101121320321-1110131122110221-3333232113210233-3320003332133013-3212300101003213-0012020323223133)
