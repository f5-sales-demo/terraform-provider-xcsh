---
page_title: "xcsh_alert_template landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template landing."
---

# xcsh_alert_template landing

<a id="canonical-1003131020323200-2102322313132332-1330100113320011-3213311001122330-0122322231213130-1212212131231012-1332020002133101-0232021330332020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321122002322322-3103320112200330-3313030010310202-3321323322121031-0221113021232023-2103013012010121-2320000132321313-1230122113132201"></a>

## xcsh_alert_template — xcsh_alert_template / 332321222231 / 2

Breadcrumbs:

- xcsh_alert_template

Manages Domain to protect in F5 Distributed Cloud.

<a id="canonical-1212231303112211-1122032113322311-0013331210031033-2112222120121121-0103313010123213-0132023210110110-2100021231103213-1133132232121101"></a>

## Prerequisites — xcsh_alert_template / 332321222231 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3113203100310310-2002123021312011-2123213100213020-0320013033333021-2321120203113111-0013003212303223-0112321121213101-2133311112121330"></a>

## Minimal configuration — xcsh_alert_template / 332321222231 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertTemplate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertTemplate by name
data "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"
}

output "alert_template_id" {
  value = data.xcsh_alert_template.example.id
}
```

<a id="canonical-3111121222302213-2232012011232320-2313300213222030-2113021013103001-0202120203211110-0211202100130322-1212222133112110-1012220203110123"></a>

## Root configuration — xcsh_alert_template / 332321222231 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2331021313303120-1002300203310313-2232001313323023-1320230200200032-0122000302133002-3322203031123110-1313000230122331-2331131213303232"></a>

## Next pages — xcsh_alert_template / 332321222231 / 6

- [Property reference](../guides/data-sources--alert_template--reference--group-001.md#canonical-3212223031032323-0020230301002003-0212033002000323-0113313121231200-0223110303221321-0303212201023311-3001013320003100-1303010003111200)
- [Examples](../guides/data-sources--alert_template--examples--group-001.md#canonical-3213301303320332-3211020211111123-0131003133300303-0130333120222113-3302133011010030-2202002310102011-0320232133100203-2233310233010231)
