---
page_title: "xcsh_bot_detection_rule"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_detection_rule."
---

# xcsh_bot_detection_rule

<a id="canonical-2221001030010210-1213310213300133-3302112123210020-1030311021023312-0232211203112101-3232200001000312-1300200232311310-2121130011001102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bot_detection_rule

Reads Bot Detection Rule information from F5 Distributed Cloud.

<a id="canonical-1022132333001300-3030200322223030-2301210023010310-2302003221231002-1103001111013001-2020031033003323-1111100312322013-1000100331012322"></a>

### Prerequisites for `xcsh_bot_detection_rule`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2202332131012312-0200201222211000-1202323010023220-0121102223330000-2001323322001000-3213230020113201-2133130201012122-3010220212120331"></a>

### Minimal configuration for `xcsh_bot_detection_rule`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotDetectionRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDetectionRule by name
data "xcsh_bot_detection_rule" "example" {
  name      = "example-bot-detection-rule"
  namespace = "staging"
}

output "bot_detection_rule_id" {
  value = data.xcsh_bot_detection_rule.example.id
}
```

<a id="canonical-2320211121020211-1213031113021003-3221333303312232-3122212211023032-1103011231011001-3032202210103130-1300031113123231-0013122123021323"></a>

### Root configuration for `xcsh_bot_detection_rule`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1221310003303322-1002200130331300-0021020221001030-2323223331332122-3221020002003031-2213232211330333-3022110103102303-0101011332310031"></a>

### Explore this collection for `xcsh_bot_detection_rule`

- [Property reference](../guides/data-sources--bot_detection_rule--reference--group-001.md#canonical-1300000332213130-3231303130200131-1302120313303120-2230111331211100-2123003010122200-3002313113113213-0321013203213000-2322013230021313)
- [Examples](../guides/data-sources--bot_detection_rule--examples--group-001.md#canonical-2232113313332232-3000000211312211-0303233303022011-1332301232323032-0121030001232203-3300030203332032-0221121030302230-0100313323330002)
