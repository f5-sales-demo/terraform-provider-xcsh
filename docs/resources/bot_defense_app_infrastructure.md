---
page_title: "xcsh_bot_defense_app_infrastructure landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure landing."
---

# xcsh_bot_defense_app_infrastructure landing

<a id="canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230333013220030-0310133131300120-1101123222031100-0321111000331211-3210001331330120-2100200100100103-2010313012121031-2202032221110312"></a>

## xcsh_bot_defense_app_infrastructure — xcsh_bot_defense_app_infrastructure / 311303233023 / 2

Breadcrumbs:

- xcsh_bot_defense_app_infrastructure

Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

<a id="canonical-2212320132333111-0011300023321113-1120130012221003-2233032002201022-0111030032323332-0200212300303321-1112013233221121-2211321303131321"></a>

## Prerequisites — xcsh_bot_defense_app_infrastructure / 311303233023 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3203331302110123-2111110001221113-0003301012312203-1220333031002113-1003301213113233-3120233001031310-0002220012002202-3321330022132200"></a>

## Minimal configuration — xcsh_bot_defense_app_infrastructure / 311303233023 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotDefenseAppInfrastructure Resource Example
# Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotDefenseAppInfrastructure configuration
resource "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}
```

<a id="canonical-2233011103102331-1122331330110323-0003303213211211-3013013210322132-1111120131332303-3102301112003312-1133021233001231-2003321210320310"></a>

## Root configuration — xcsh_bot_defense_app_infrastructure / 311303233023 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0122031202020323-0213023032011322-1121022333200213-0303123032022002-0200013032322100-3221120133221131-1212323110013111-2101011321331301"></a>

## Next pages — xcsh_bot_defense_app_infrastructure / 311303233023 / 6

- [Property reference](../guides/resources--bot_defense_app_infrastructure--reference--group-001.md#canonical-2223101122200310-0021021103103022-3232011233203112-1202002312211021-1012103011330123-1022210000013032-1301123330002320-3003231322323010)
- [Examples](../guides/resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-2133001221300330-1201323101022110-0302232333033321-0230133312113321-1010131012132012-0233102022313333-2302101232213122-3113330201323310)
- [Import](../guides/resources--bot_defense_app_infrastructure--lifecycle--group-001.md#canonical-3031330333120133-1113200210122200-3223322030233223-1012123120002222-0030320131301013-3131303102032102-0113221130033130-3003212021122000)
- [Timeouts](../guides/resources--bot_defense_app_infrastructure--lifecycle--group-001.md#canonical-2231120213101321-3202323002330103-2233132111213210-0210332230231111-0210000203112210-2301123221323313-1331101110132201-3021033133130100)
