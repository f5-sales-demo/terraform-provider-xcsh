---
page_title: "xcsh_bot_defense_app_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure examples."
---

# xcsh_bot_defense_app_infrastructure examples

<a id="canonical-2203320102322222-3033130103223000-2221233312330220-3132103002232130-3311323002122313-1102223020232110-3002130020332031-3121022001133130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301131103102103-0232122011321311-2020111210033302-0212020201103303-1113220303312030-1031333103210023-1101323032021133-1232313131203333"></a>

## Examples — Examples / 331313130010 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- Examples

<a id="canonical-1131210332310213-1111003313300030-1230112222321310-0020222031221132-1321202233022302-3220212113101021-0322302111331223-2230211300202123"></a>

## Complete configurations — Examples / 331313130010 / 3

- [Data source](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-3310222001301223-3222221111211002-1103321222223201-0132033202120020-3012311220222302-0210012033102201-0233023132131000-2102300310203212): valid configuration.

<a id="canonical-2330203222230331-0321332132010301-2200022123231113-1012212232330131-1302310323122303-1003311221021321-2223310222212030-3112312023112013"></a>

## Next pages — Examples / 331313130010 / 4

- [Data source](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-3310222001301223-3222221111211002-1103321222223201-0132033202120020-3012311220222302-0210012033102201-0233023132131000-2102300310203212)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)

<a id="canonical-3310222001301223-3222221111211002-1103321222223201-0132033202120020-3012311220222302-0210012033102201-0233023132131000-2102300310203212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110300210202020-1302012223011121-0332113003111310-1220200213120320-0232023013211301-0332203221113203-3111033330003322-2311021231010113"></a>

## Data source — Data source / 132103321110 / 2

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
- [Examples](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-2203320102322222-3033130103223000-2221233312330220-3132103002232130-3311323002122313-1102223020232110-3002130020332031-3121022001133130)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_defense_app_infrastructure/data-source.tf`; digest `sha256:9d1dcbef65f7cbdf8cc89c39b723d01c2f54dd36e7a2d02f47f90c88068c73e0`.

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

<a id="canonical-2202313330131012-0012223113122120-3312002210222231-1320332313110032-3203133200111030-1222102310211222-0230310220113212-3000303133100333"></a>

## Next pages — Data source / 132103321110 / 3

- [Examples](data-sources--bot_defense_app_infrastructure--examples--group-001.md#canonical-2203320102322222-3033130103223000-2221233312330220-3132103002232130-3311323002122313-1102223020232110-3002130020332031-3121022001133130)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md#canonical-3303122223102311-0202101033323122-0020322211123301-2100013330111310-1031210220113020-3201200231330321-3231001213311211-2220120313100211)
