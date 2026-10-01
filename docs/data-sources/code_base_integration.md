---
page_title: "xcsh_code_base_integration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration landing."
---

# xcsh_code_base_integration landing

<a id="canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201322313302333-3230211121332320-3312321120131211-3121310002313010-0132331323001000-1300122321131333-0201003300320300-3121002203003110"></a>

## xcsh_code_base_integration — xcsh_code_base_integration / 112131322103 / 2

Breadcrumbs:

- xcsh_code_base_integration

Manages integration details in F5 Distributed Cloud.

<a id="canonical-2313001130122112-2102213233312002-2103310311233221-0021220232121010-1013020232010322-1211203321102023-0001313030000003-3223123020233311"></a>

## Prerequisites — xcsh_code_base_integration / 112131322103 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2313030133203303-0133202313330333-3002122002213303-0311000332203001-2022330310302320-1112120001320000-2332020113103221-0202131010300210"></a>

## Minimal configuration — xcsh_code_base_integration / 112131322103 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CodeBaseIntegration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CodeBaseIntegration by name
data "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}

output "code_base_integration_id" {
  value = data.xcsh_code_base_integration.example.id
}
```

<a id="canonical-0232233132301033-3112131200321030-1203233330320322-2213311121210330-3322121313232031-1113221310221022-0230213211112030-0102110211122100"></a>

## Root configuration — xcsh_code_base_integration / 112131322103 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2202322203031323-2221210110132223-3001131131222011-0023120311131113-0310002133010201-1031032122301122-3301012211230310-0033233303001200"></a>

## Next pages — xcsh_code_base_integration / 112131322103 / 6

- [Property reference](../guides/data-sources--code_base_integration--reference--group-001.md#canonical-2333130312100012-0222230000033312-2211223330130202-3230133011333300-3133332212221130-2100321103012223-0120330102301133-3131023123000022)
- [Examples](../guides/data-sources--code_base_integration--examples--group-001.md#canonical-0130211200321301-3120330003030130-0333331123321300-1033122113102103-3121121133011111-1000213313233233-3202013201321223-2331132333232303)
