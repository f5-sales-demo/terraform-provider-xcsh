---
page_title: "xcsh_bot_endpoint_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_endpoint_policy."
---

# xcsh_bot_endpoint_policy

<a id="canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bot_endpoint_policy

Reads Bot Endpoint Policy information from F5 Distributed Cloud.

<a id="canonical-2200200103020330-0123111323131100-3101132031103310-3030023211222331-2021030121320233-0112021311300100-2300130032231121-0010203313300130"></a>

### Prerequisites for `xcsh_bot_endpoint_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0000330000102122-2301311113031333-3210210233012221-1032011311330112-2022332103321012-3331223202223133-0001320202102232-3013033121223122"></a>

### Minimal configuration for `xcsh_bot_endpoint_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotEndpointPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotEndpointPolicy by name
data "xcsh_bot_endpoint_policy" "example" {
  name      = "example-bot-endpoint-policy"
  namespace = "staging"
}

output "bot_endpoint_policy_id" {
  value = data.xcsh_bot_endpoint_policy.example.id
}
```

<a id="canonical-2332110003013331-2103013011033033-1020230200330303-1121003223002221-2312322111231102-2202232302022123-2033003003120222-0312022120020100"></a>

### Root configuration for `xcsh_bot_endpoint_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3222033320303230-2321112020131223-2133132322103201-2022320033203322-2123112012012121-0301123031000111-0331332222320110-2303331100101130"></a>

### Explore this collection for `xcsh_bot_endpoint_policy`

- [Property reference](../guides/data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [Examples](../guides/data-sources--bot_endpoint_policy--examples--group-001.md#canonical-3303333021232001-0201323230023222-2012023013012233-3023322031312202-3310031131032123-0323231120222213-0220003001000000-0232003211310310)
