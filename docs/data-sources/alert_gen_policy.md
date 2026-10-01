---
page_title: "xcsh_alert_gen_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy landing."
---

# xcsh_alert_gen_policy landing

<a id="canonical-1212223311231031-2303130101102331-2121323302132120-3001003212230233-2233103321211211-1300022131031130-1101122301003332-3330023301222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303323232001230-2302101211232022-1301032201020132-0132031031211103-2023221221130233-1011230211303022-1312122121233232-2232103000011202"></a>

## xcsh_alert_gen_policy — xcsh_alert_gen_policy / 202123101032 / 2

Breadcrumbs:

- xcsh_alert_gen_policy

Manages Alert Generation Policy in F5 Distributed Cloud.

<a id="canonical-2020111012232213-3322100320322320-0202201132333211-0023103203211331-2323120130220033-0113103101313213-2001003212322200-2111310223022113"></a>

## Prerequisites — xcsh_alert_gen_policy / 202123101032 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1321320212130112-1330121323323311-2123030032232313-3211020010311002-0011120203302102-3022131212021101-1310101132232222-3321010202132021"></a>

## Minimal configuration — xcsh_alert_gen_policy / 202123101032 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertGenPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertGenPolicy by name
data "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}

output "alert_gen_policy_id" {
  value = data.xcsh_alert_gen_policy.example.id
}
```

<a id="canonical-1303112103210103-2231310331001210-1110221202131310-3322002321330211-3111231103331000-1020233122302201-1100100112023220-1122023033112000"></a>

## Root configuration — xcsh_alert_gen_policy / 202123101032 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0010013322211313-2132020332313320-2122100213023111-1313323333300002-0233102123103133-3131230231111211-2001133013310313-2112000020222322"></a>

## Next pages — xcsh_alert_gen_policy / 202123101032 / 6

- [Property reference](../guides/data-sources--alert_gen_policy--reference--group-001.md#canonical-0232203323003000-0030100311001232-1132202033111113-0322101230133330-0330113102210010-3303033232211100-3200031120003321-1033030031110112)
- [Examples](../guides/data-sources--alert_gen_policy--examples--group-001.md#canonical-3101133110101332-1212313230200233-0212133030031121-3000033232011331-3002320201122210-1121300201310231-3102220301010033-0300333030203210)
