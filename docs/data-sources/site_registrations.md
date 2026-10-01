---
page_title: "xcsh_site_registrations landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations landing."
---

# xcsh_site_registrations landing

<a id="canonical-2320011012012203-1211121021012323-3333302232010031-0232010323320310-2122323133012221-2310311330012133-3232132103233312-1122133001030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202022330011121-1201133101130100-0323113003332230-3002233000203100-1110301121110110-2132323212031321-1303320302201120-3201201320311113"></a>

## xcsh_site_registrations — xcsh_site_registrations / 303130033322 / 2

Breadcrumbs:

- xcsh_site_registrations

List Customer Edge registrations.

<a id="canonical-1332113021010310-1110212222332321-2322221112020333-3132012101131200-0332000020230320-0131133202013200-0122103101021301-0021320131230310"></a>

## Prerequisites — xcsh_site_registrations / 303130033322 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2322122120031231-0023200100303123-1223100221301331-1030310323333001-0201120302021221-3320200203333301-1303210133102013-1122220310203100"></a>

## Minimal configuration — xcsh_site_registrations / 303130033322 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
```

<a id="canonical-2313310023101221-3021232120111120-0220113021310100-2222320131333022-3200303111210000-2100003332212031-1101122330222113-3200030020221321"></a>

## Root configuration — xcsh_site_registrations / 303130033322 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0202000213010131-3113010033203112-2200012001022113-2100231031233320-3313000302003320-1320232131310100-1213011310003021-3122120012022011"></a>

## Next pages — xcsh_site_registrations / 303130033322 / 6

- [Property reference](../guides/data-sources--site_registrations--reference--group-001.md#canonical-3221303221330313-3221311133231332-1110021131331113-3112110232131000-0111222003200113-2200322113211230-1112020222311121-0102203311103030)
- [Examples](../guides/data-sources--site_registrations--examples--group-001.md#canonical-3032020103011321-0332312120220310-2110113322301312-1231201123320220-2303332131331323-2132310013310121-1131233223031310-1313331210103313)
