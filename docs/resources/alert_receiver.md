---
page_title: "xcsh_alert_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver landing."
---

# xcsh_alert_receiver landing

<a id="canonical-0033232020121120-2032300130130001-0103310100210332-0222200222113132-2102202301202302-0022022203130310-0311110212100013-2210100123333311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011101031221223-3220301101311020-0003233321221230-1210111022323120-0121203222103100-0002300210001231-2230310030312101-3032332002230111"></a>

## xcsh_alert_receiver — xcsh_alert_receiver / 300200001230 / 2

Breadcrumbs:

- xcsh_alert_receiver

Manages new Alert Receiver object in F5 Distributed Cloud.

<a id="canonical-3201032231002311-2133201300111133-0200022223321332-1303330210301112-2113332122333003-0113002131003203-0203023332311031-2200031201123033"></a>

## Prerequisites — xcsh_alert_receiver / 300200001230 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1201211313313231-1101111023202322-1332112103001303-0100313302122133-0333201033203112-2322100233111012-3230013310012131-3213300310000011"></a>

## Minimal configuration — xcsh_alert_receiver / 300200001230 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertReceiver Resource Example
# Manages new Alert Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertReceiver configuration
resource "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}
```

<a id="canonical-2101323203011003-0220001022121313-2113033310130000-2101001113230131-3033131101130220-2322231330121132-2332113030233032-0320002332213103"></a>

## Root configuration — xcsh_alert_receiver / 300200001230 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1020221230021123-0301003123031333-0333210233020013-1121002333213232-2313000122001311-1202202031000023-3312020201101200-1231111333102230"></a>

## Next pages — xcsh_alert_receiver / 300200001230 / 6

- [Property reference](../guides/resources--alert_receiver--reference--group-001.md#canonical-0323313120200103-0203032321132030-1120210132201312-3223000113332221-3323222122210213-0100202202103113-3212102330201013-1220330230101010)
- [Examples](../guides/resources--alert_receiver--examples--group-001.md#canonical-2322202131233023-2131223013300313-1311001032100001-2201231332310233-0302203023203330-0120223112233121-2021102111032332-0233300211000222)
- [Import](../guides/resources--alert_receiver--lifecycle--group-001.md#canonical-2120230021133013-3220001122131200-3300023323230321-3033322220311321-0323001020101010-2103100010200323-3133011113122323-0010212023102221)
- [Timeouts](../guides/resources--alert_receiver--lifecycle--group-001.md#canonical-3123000220030201-3310112000111320-2031322131111122-3332320010103220-2210121121102130-3030301231110212-1130231301000311-1001022002300333)
