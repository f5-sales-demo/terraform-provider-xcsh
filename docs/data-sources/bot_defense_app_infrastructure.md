---
page_title: "xcsh_bot_defense_app_infrastructure"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure."
---

# xcsh_bot_defense_app_infrastructure

<a id="canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bot_defense_app_infrastructure

Reads Bot Defense App Infrastructure information from F5 Distributed Cloud.

<a id="canonical-3001113203212113-3020303112220301-2013333123111122-3111030212212202-1122230303133230-2313333220030333-3012303133232113-3310101313022313"></a>

### Prerequisites for `xcsh_bot_defense_app_infrastructure`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1312230222023203-1033301021301020-0233030021221133-3032011112310320-0031200203010100-0012112023101302-3230310231032333-1120230120300322"></a>

### Minimal configuration for `xcsh_bot_defense_app_infrastructure`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotDefenseAppInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDefenseAppInfrastructure by name
data "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}

output "bot_defense_app_infrastructure_id" {
  value = data.xcsh_bot_defense_app_infrastructure.example.id
}
```

<a id="canonical-0301301033211032-0112212302113300-1233123012313330-0031303211223112-0312103020023031-0210311313131111-0000203202323210-3013021303032331"></a>

### Root configuration for `xcsh_bot_defense_app_infrastructure`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2132102103133300-3200221223101232-0100200232322233-2132131200031331-1103022321320123-2233012102030123-0020302330112110-3230332110322130"></a>

### Explore this collection for `xcsh_bot_defense_app_infrastructure`

- [Property reference](../guides/data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [Examples](../guides/data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-2203320102322222-3033130103223000-2221233312330220-3132103002232130-3311323002122313-1102223020232110-3002130020332031-3121022001133130)
