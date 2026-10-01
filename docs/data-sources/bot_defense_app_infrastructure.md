---
page_title: "xcsh_bot_defense_app_infrastructure landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure landing."
---

# xcsh_bot_defense_app_infrastructure landing

<a id="canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001113203212113-3020303112220301-2013333123111122-3111030212212202-1122230303133230-2313333220030333-3012303133232113-3310101313022313"></a>

## xcsh_bot_defense_app_infrastructure — xcsh_bot_defense_app_infrastructure / 320333302110 / 2

Breadcrumbs:

- xcsh_bot_defense_app_infrastructure

Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

<a id="canonical-1312230222023203-1033301021301020-0233030021221133-3032011112310320-0031200203010100-0012112023101302-3230310231032333-1120230120300322"></a>

## Prerequisites — xcsh_bot_defense_app_infrastructure / 320333302110 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0301301033211032-0112212302113300-1233123012313330-0031303211223112-0312103020023031-0210311313131111-0000203202323210-3013021303032331"></a>

## Minimal configuration — xcsh_bot_defense_app_infrastructure / 320333302110 / 4

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

<a id="canonical-2132102103133300-3200221223101232-0100200232322233-2132131200031331-1103022321320123-2233012102030123-0020302330112110-3230332110322130"></a>

## Root configuration — xcsh_bot_defense_app_infrastructure / 320333302110 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0323203012220310-2311000231000201-1130231230211002-1003210312120311-3211320323322221-1311133001002130-1223312310120123-2311022033133122"></a>

## Next pages — xcsh_bot_defense_app_infrastructure / 320333302110 / 6

- [Property reference](../guides/data-sources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2132310132221331-3023323033302023-2020313301002322-3022221322122223-3020302132133032-3320001331003303-2321302323320112-2100121000233122)
- [Examples](../guides/data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-2203320102322222-3033130103223000-2221233312330220-3132103002232130-3311323002122313-1102223020232110-3002130020332031-3121022001133130)
