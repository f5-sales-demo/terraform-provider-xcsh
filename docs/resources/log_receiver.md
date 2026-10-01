---
page_title: "xcsh_log_receiver landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver landing."
---

# xcsh_log_receiver landing

<a id="canonical-0313113102130312-2233123212112332-3122330131100323-1311101313033113-0013312121123022-2311322331111110-1300111301210312-0111133031122003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303033210203230-2011302203031313-1212312232033331-0221321133233222-0130030222020211-3301120300302130-1233121311013302-0023222311300011"></a>

## xcsh_log_receiver — xcsh_log_receiver / 001200110112 / 2

Breadcrumbs:

- xcsh_log_receiver

Manages new Log Receiver object in F5 Distributed Cloud.

<a id="canonical-0332201320312012-3120311102021121-3303230210302011-0323131203100232-2030133300110101-2301211322313021-3102033023002333-3132120010202123"></a>

## Prerequisites — xcsh_log_receiver / 001200110112 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1201112130123033-0033010222023102-2102133321112121-1222320223112230-3201212221202331-0131233321313311-2203011211232123-2301213101210230"></a>

## Minimal configuration — xcsh_log_receiver / 001200110112 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# LogReceiver Resource Example
# Manages new Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic LogReceiver configuration
resource "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}
```

<a id="canonical-2022130212000200-3010223302203030-0120311202213220-0300023232023013-1212121232310212-2223223110310211-0030132030013103-3010331023323010"></a>

## Root configuration — xcsh_log_receiver / 001200110112 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1330031333312012-1231230232120001-0202122311221212-1320103311111232-3132203133330220-3121110002322211-3111122022010330-1332230230210213"></a>

## Next pages — xcsh_log_receiver / 001200110112 / 6

- [Property reference](../guides/resources--log_receiver--reference--group-001.md#canonical-3233000213200330-1021301100302120-2322002222223311-3211232003101211-3002210023303303-3302330021002111-0231132130121221-0233101103123201)
- [Examples](../guides/resources--log_receiver--examples--group-001.md#canonical-0101332232122303-0002030101233212-3100332320223031-3313332101002120-3223320021112202-0120333120203012-2101332223200021-3100030120332330)
- [Import](../guides/resources--log_receiver--lifecycle--group-001.md#canonical-3123223113211002-2122000132323322-0202030313031023-2331001321000232-0120013032201311-1230233312003002-0122233323122202-3221300302102201)
- [Timeouts](../guides/resources--log_receiver--lifecycle--group-001.md#canonical-3230233203023020-2100232101131133-1332102211333320-1020000113103212-0002323321022100-1011211103213030-1132302033010300-2231100010130121)
