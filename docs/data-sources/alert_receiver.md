---
page_title: "xcsh_alert_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver landing."
---

# xcsh_alert_receiver landing

<a id="canonical-3012133032112200-3010101132010313-3312231133323113-0303010111120112-3233032202012130-0203311232212023-1011202020122232-2231030013203102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100231103113020-3000200230003230-3131321300200210-0213011120111003-1121212220110133-1303101200030200-1132023300033130-0231221113220330"></a>

## xcsh_alert_receiver — xcsh_alert_receiver / 102303032011 / 2

Breadcrumbs:

- xcsh_alert_receiver

Manages new Alert Receiver object in F5 Distributed Cloud.

<a id="canonical-2002223230321012-1203323101112002-3301110110203303-1101121320222101-0213220201223231-1100022202323321-2002222230203123-3012112222311231"></a>

## Prerequisites — xcsh_alert_receiver / 102303032011 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3310003221333201-3101130232100210-0111312311100221-0200102112133302-0232110001100232-3013121210311223-1212001202302130-1033312221223303"></a>

## Minimal configuration — xcsh_alert_receiver / 102303032011 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertReceiver by name
data "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}

output "alert_receiver_id" {
  value = data.xcsh_alert_receiver.example.id
}
```

<a id="canonical-3010332111230310-3200032333033111-3121113110112202-2103012210100013-0001113011201303-1123233300010010-3110221310003111-0001102120313111"></a>

## Root configuration — xcsh_alert_receiver / 102303032011 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2112120311200022-3031213211120220-1302231230030310-3321100030201222-1212023001311113-0023212331101230-0212033130100300-1201203222201202"></a>

## Next pages — xcsh_alert_receiver / 102303032011 / 6

- [Property reference](../guides/data-sources--alert_receiver--reference--group-001.md#canonical-2132131323221021-2020310203213232-1133303303131310-0202011031123003-3121133011213003-0120331333330031-0012311032301222-0021221020101101)
- [Examples](../guides/data-sources--alert_receiver--examples--group-001.md#canonical-2130120101331130-2120000310312101-0313131112010322-3013100020120003-0221123111021020-1300012221211231-3110312120331222-0223310311022310)
