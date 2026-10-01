---
page_title: "xcsh_global_log_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver landing."
---

# xcsh_global_log_receiver landing

<a id="canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223033111322233-1011231301010332-0232222101231223-1112103122011021-3222001022133302-3212222023100233-1102213033210100-2123033323121123"></a>

## xcsh_global_log_receiver — xcsh_global_log_receiver / 110112223000 / 2

Breadcrumbs:

- xcsh_global_log_receiver

Manages new Global Log Receiver object in F5 Distributed Cloud.

<a id="canonical-2311233201302020-1303213011230100-0223030223323232-2321221321003111-3013300203023130-3223200321313032-0131112330031101-1131200333122331"></a>

## Prerequisites — xcsh_global_log_receiver / 110112223000 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2231220203112032-1202330212333212-0220332313321300-0300021301003311-2323333133013132-0133122203020312-1222110123231333-2122013031123002"></a>

## Minimal configuration — xcsh_global_log_receiver / 110112223000 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GlobalLogReceiver Resource Example
# Manages new Global Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GlobalLogReceiver configuration
resource "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}
```

<a id="canonical-1131102000221202-3132211211322213-3130030131021201-2132333120123321-0120223023131312-2233020321200202-1130103331033300-1012112003030021"></a>

## Root configuration — xcsh_global_log_receiver / 110112223000 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3003322013222231-0100000003333330-1213133210301022-1032220032201113-1322110212203133-1103302123200332-3321210221021003-0212210223122210"></a>

## Next pages — xcsh_global_log_receiver / 110112223000 / 6

- [Property reference](../guides/resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [Examples](../guides/resources--global_log_receiver--examples--group-001.md#canonical-1130021233300031-0131022000130203-1222320002223311-0213001321212132-1132112111222022-2232101012201323-0211121203132212-3221211101222233)
- [Import](../guides/resources--global_log_receiver--lifecycle--group-001.md#canonical-3331120312302021-3201113003013002-1002000000030220-1230313233333002-1203302131113111-0130133021013001-2011203310210230-2022003232312222)
- [Timeouts](../guides/resources--global_log_receiver--lifecycle--group-001.md#canonical-0212220111221322-2101233301200233-3022120110233011-2010133032312133-3031301111023121-3210011022031300-2230321100132130-3202112101133120)
