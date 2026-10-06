---
page_title: "xcsh_log_receiver"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver."
---

# xcsh_log_receiver

<a id="canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_log_receiver

Reads Log Receiver information from F5 Distributed Cloud.

<a id="canonical-2303221233002022-3323000231230130-0310220300323332-1321303121333013-1000121123111201-0022312120310031-0300311312110331-1023110200211313"></a>

### Prerequisites for `xcsh_log_receiver`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3210021130101000-2021031022023123-0322110010230310-3322202011200223-2132322022133013-3100311000130302-3101132222313100-3120302220120210"></a>

### Minimal configuration for `xcsh_log_receiver`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# LogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LogReceiver by name
data "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}

output "log_receiver_id" {
  value = data.xcsh_log_receiver.example.id
}
```

<a id="canonical-3023320011200132-3013033200313222-1312003131200110-2001121231230301-0303112210103132-2011333333132221-3122111130323221-3112330220303303"></a>

### Root configuration for `xcsh_log_receiver`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2121033211000333-2021023102313201-3102032112020013-1113330202203003-1000211112123000-0313210322330013-1112221111232001-0133201303021113"></a>

### Explore this collection for `xcsh_log_receiver`

- [Property reference](../guides/data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [Examples](../guides/data-sources--log_receiver--examples--group-001.md#canonical-2223011232000332-3200032103323313-3320133300222033-2022032003330100-2001123102211220-2202121202312110-3031022223312223-3300202332132331)
