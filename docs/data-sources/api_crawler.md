---
page_title: "xcsh_api_crawler"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler."
---

# xcsh_api_crawler

<a id="canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_api_crawler

Reads API Crawler information from F5 Distributed Cloud.

<a id="canonical-2221310030212111-0212133102202202-0001133031203122-0211201222313303-3102203332031020-3032102000121003-0202012120320223-2333233203123213"></a>

### Prerequisites for `xcsh_api_crawler`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0302000202132012-2031033212013202-2231033130323102-3303232022213010-2111323012303021-1013131100010332-0231110323003113-3012122123102131"></a>

### Minimal configuration for `xcsh_api_crawler`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APICrawler Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APICrawler by name
data "xcsh_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"
}

output "api_crawler_id" {
  value = data.xcsh_api_crawler.example.id
}
```

<a id="canonical-1221231102010331-2202202023132331-3022332031332100-1201301002331133-0223302212213213-2130222200202331-1222113302213233-3211002300021221"></a>

### Root configuration for `xcsh_api_crawler`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3031010021202303-1300333002033200-0002233232210300-2333313003032201-0103101002233023-3221201230300322-0323232113032212-3011133133231323"></a>

### Explore this collection for `xcsh_api_crawler`

- [Property reference](../guides/data-sources--api_crawler--reference--group-001.md#canonical-0220220322223001-1320313010233203-3300000030200300-0201002302301013-0210212312130310-0300333002303302-1311323021220123-2132303221000023)
- [Examples](../guides/data-sources--api_crawler--examples--group-001.md#canonical-3033332101013230-0112123023331122-1221102333031200-0302111301111031-1120302110301131-1023232102311101-1120221301330013-3333200122200111)
