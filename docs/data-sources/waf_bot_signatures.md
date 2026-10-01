---
page_title: "xcsh_waf_bot_signatures landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_bot_signatures landing."
---

# xcsh_waf_bot_signatures landing

<a id="canonical-1232013323203331-2030333000300300-3031302123320310-3302000013332010-1200120322302012-2021223130131323-2112213021213111-3103303123120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222220223311211-3321132201013310-1001233023332131-3011111202303213-0230010323131011-3032103030230321-0322321002232211-3011113023131113"></a>

## xcsh_waf_bot_signatures — xcsh_waf_bot_signatures / 210323200022 / 2

Breadcrumbs:

- xcsh_waf_bot_signatures

Bot detection and defense configuration.

<a id="canonical-1200103230011002-0212231011013103-0102100012211220-2301113121323001-3001320211330333-3200003322323322-0110211223121223-0332333330002221"></a>

## Prerequisites — xcsh_waf_bot_signatures / 210323200022 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3331032120132003-2102313130131021-0310123221112012-1312303031301013-1110302002301213-1130211313230302-1213302323003001-0001121200232323"></a>

## Minimal configuration — xcsh_waf_bot_signatures / 210323200022 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFBotSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_bot_signatures" "example" {
}

output "waf_bot_signatures_result" {
  value = data.xcsh_waf_bot_signatures.example
}
```

<a id="canonical-3120201233302112-0121021013111032-2210211332300332-2031113231322301-2221231121021031-0332033303231233-3032012120013122-1313201113032302"></a>

## Root configuration — xcsh_waf_bot_signatures / 210323200022 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2023320302300123-2003200210321201-2301303210033121-1113112022232121-1210212102123001-2221310011101202-1001011032201230-1110030213222020"></a>

## Next pages — xcsh_waf_bot_signatures / 210323200022 / 6

- [Property reference](../guides/data-sources--waf_bot_signatures--reference--group-001.md#canonical-0202302311100230-1213232002201200-1112333313213300-1120210203201311-3310221133022103-2132301023122012-0230333001330220-0113223112221303)
- [Examples](../guides/data-sources--waf_bot_signatures--examples--group-001.md#canonical-0221322122210100-2300022313321102-2033010210311101-2201221100231320-3020012201200310-1222222312200201-3033013013221222-3211032233120313)
