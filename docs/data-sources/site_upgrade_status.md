---
page_title: "xcsh_site_upgrade_status landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_status landing."
---

# xcsh_site_upgrade_status landing

<a id="canonical-0311103030013201-2023013222321300-1123321212232123-1232102210103331-2001110222022222-0103301221312320-3101213121122310-1101333312101011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301212332201332-3303320113032030-1103012130102120-3031313330202132-0021231132032230-0212200103203202-1110013232300000-3213010020133133"></a>

## xcsh_site_upgrade_status — xcsh_site_upgrade_status / 101112323322 / 2

Breadcrumbs:

- xcsh_site_upgrade_status

Observes SMSv2 site upgrade eligibility and waits for explicitly supplied software and
operating-system targets to converge.

<a id="canonical-1201321121310313-1213211300102100-2322222232232311-0220102331111322-3312332233112231-3313011003111012-3111102101223223-1223231301203232"></a>

## Prerequisites — xcsh_site_upgrade_status / 101112323322 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1202223013303121-3320231311201021-1230302013331133-3131131033003323-0103120112033201-3100221033002022-0321300103330323-3003222210332321"></a>

## Minimal configuration — xcsh_site_upgrade_status / 101112323322 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Observe upgrade eligibility or wait for supplied software and OS targets to
# be installed with the site back ONLINE.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 7.3.0"
    }
  }
}

data "xcsh_site_upgrade_status" "site" {
  site = "example-smsv2-site"

  expected_software_version = "crt-20260201-0179"
  expected_os_version       = "9.2026.17"
  wait                      = true
  timeout_seconds           = 7200
  poll_interval_seconds     = 30
}

output "upgrade_converged" {
  value = data.xcsh_site_upgrade_status.site.target_converged
}
```

<a id="canonical-0131332023301203-3320101112313133-0222012310331031-3131131123213132-3313312211310111-1200110313000332-2211010333000330-2011321222002120"></a>

## Root configuration — xcsh_site_upgrade_status / 101112323322 / 5

Required root properties: `site`. Full root flags and choices appear in the property reference.

<a id="canonical-2221302031021313-3310301223203022-3223212211233000-0100121101111311-2201330203030233-0020210030123112-2022011020132131-1022202311020332"></a>

## Next pages — xcsh_site_upgrade_status / 101112323322 / 6

- [Property reference](../guides/data-sources--site_upgrade_status--reference--group-001.md#canonical-0002123201223230-3102222113003200-2112313021101312-3221113322221231-2011313330023232-1101331323313032-0200231210030103-1020333122010212)
- [Examples](../guides/data-sources--site_upgrade_status--examples--group-001.md#canonical-3223111323110230-0123231002213330-0202201232313001-1031031203203312-0001221101110320-3222113013223302-2012232221110223-2002010112100100)
