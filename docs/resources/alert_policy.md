---
page_title: "xcsh_alert_policy landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy landing."
---

# xcsh_alert_policy landing

<a id="canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331030320002110-2321322110220023-0222232112031232-0113301332230101-0110030223311313-0232001313210210-0311030122313101-1210110230113332"></a>

## xcsh_alert_policy — xcsh_alert_policy / 203313002022 / 2

Breadcrumbs:

- xcsh_alert_policy

Manages new Alert Policy Object in F5 Distributed Cloud.

<a id="canonical-0320221301213122-1220130200011031-2000131020211130-2100031230313231-2211100232233120-1122231101201230-0020033222321110-1321010313302210"></a>

## Prerequisites — xcsh_alert_policy / 203313002022 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3232212313331200-1200322101202321-2030130232101203-0301121233202231-2010223212021221-3210132332303323-1012301303203011-2003121232002131"></a>

## Minimal configuration — xcsh_alert_policy / 203313002022 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertPolicy Resource Example
# Manages new Alert Policy Object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertPolicy configuration
resource "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}
```

<a id="canonical-3020003311110112-2013032103220120-3120011113223302-3032010230320101-1223231203001211-0220111213310220-3310032201020010-1220300130130333"></a>

## Root configuration — xcsh_alert_policy / 203313002022 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0122203333331230-1222213112102301-3303031213122313-2021230322120320-0201021100321133-1232110321013003-1121012212131332-3132211011023001"></a>

## Next pages — xcsh_alert_policy / 203313002022 / 6

- [Property reference](../guides/resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [Examples](../guides/resources--alert_policy--examples--group-001.md#canonical-0203231323102311-3112102032120023-3022320002010022-1132233023231020-3202031330311210-1313302001121023-0323003222313312-3133221332223323)
- [Import](../guides/resources--alert_policy--lifecycle--group-001.md#canonical-0021033122131031-2200011312112322-0213213033232022-1222031301113112-3221011320012321-1120121003133020-0331123133032132-0011203012301301)
- [Timeouts](../guides/resources--alert_policy--lifecycle--group-001.md#canonical-2231223311131012-1112302323232200-0122211100210311-0321303123030232-1332121023331221-2122323011331113-1230322202330230-3301321132001122)
