---
page_title: "xcsh_bot_suggest_values landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_suggest_values landing."
---

# xcsh_bot_suggest_values landing

<a id="canonical-3223113331303120-3211131233310020-2103101120000000-3302331313210231-3330320100131032-1303233132310303-1120022100002302-3203202022101303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211230323210210-2020200013111103-3112032213221003-3131222020220301-1100333220021313-1330103311110211-2313130120221320-3030221113322130"></a>

## xcsh_bot_suggest_values — xcsh_bot_suggest_values / 201212222312 / 2

Breadcrumbs:

- xcsh_bot_suggest_values

Resource creation operation.

<a id="canonical-0130021020013001-1312302220011202-2221313133302310-1001032302021200-1133320002332311-2302212031111033-2320000202312320-1212303130333000"></a>

## Prerequisites — xcsh_bot_suggest_values / 201212222312 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2312322120322200-3221322203111011-0301101023031120-1030133111210203-2310103310201032-0002233301031122-1123303113203322-2012010113103203"></a>

## Minimal configuration — xcsh_bot_suggest_values / 201212222312 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotSuggestValues DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_suggest_values" "example" {
  namespace = "example-value"
}

output "bot_suggest_values_result" {
  value = data.xcsh_bot_suggest_values.example
}
```

<a id="canonical-1002003312120030-1110031201312120-0331313313332301-0001021330222221-0032311110020322-1013011201133303-3300203000111211-1330303330313210"></a>

## Root configuration — xcsh_bot_suggest_values / 201212222312 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0003110110113103-2220201112102330-3123301022210133-2103333210133010-0000323322011331-1010300300120212-1203331001032111-3322112300220313"></a>

## Next pages — xcsh_bot_suggest_values / 201212222312 / 6

- [Property reference](../guides/data-sources--bot_suggest_values--reference--group-001.md#canonical-0133311012330202-3003220301112011-3121011301000332-0123101221021012-1323031211013031-1311312133320211-3120303313001021-1311000323111302)
- [Examples](../guides/data-sources--bot_suggest_values--examples--group-001.md#canonical-3023011213133012-0111020023130020-1102232223001103-1310002010213032-0302100103210013-3130303103311003-0123122011212000-3022210211021131)
