---
page_title: "xcsh_udp_loadbalancer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer landing."
---

# xcsh_udp_loadbalancer landing

<a id="canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233200111221120-2221102103132322-2333130000212202-2121331023313032-2103021010123300-0333232020002120-3112002221131311-0231313111000212"></a>

## xcsh_udp_loadbalancer — xcsh_udp_loadbalancer / 013300210333 / 2

Breadcrumbs:

- xcsh_udp_loadbalancer

Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across
origin pools.

<a id="canonical-1210333331201203-1230022322333000-2230021022302231-2302220331322231-3232013121302120-3032311100113132-0223130123133303-1300320030333232"></a>

## Prerequisites — xcsh_udp_loadbalancer / 013300210333 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1022211113110003-3031103121202020-0032323113320020-2112013230130323-2100311232213033-1303003111310232-3112223321031011-1323330231011320"></a>

## Minimal configuration — xcsh_udp_loadbalancer / 013300210333 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UDPLoadBalancer Resource Example
# Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UDPLoadBalancer configuration
resource "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}
```

<a id="canonical-3322131101321110-3120013221110203-1030321002022203-2301030000201232-2020232130230211-0323310110312202-0001220003232221-2131030100220133"></a>

## Root configuration — xcsh_udp_loadbalancer / 013300210333 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0122101332211133-0323320103212100-3301301321101132-0133220013120112-3332010331230232-0310202030120203-1312113321321122-0132020232210220"></a>

## Next pages — xcsh_udp_loadbalancer / 013300210333 / 6

- [Property reference](../guides/resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [Examples](../guides/resources--udp_loadbalancer--examples--group-001.md#canonical-1102221033122331-2330033102010213-3323032213330232-0023202131310322-0301012112303023-1300202102011232-3320002211022203-3211001223100112)
- [Import](../guides/resources--udp_loadbalancer--lifecycle--group-001.md#canonical-1322212200222002-2033010322310030-1221223112313333-3211100323323113-3111001023313010-2313133321101020-0202110122232113-3102232322202330)
- [Timeouts](../guides/resources--udp_loadbalancer--lifecycle--group-001.md#canonical-2121221112222130-0000333112003221-2102021123122200-2110000310310022-3332200311000020-2002312210130120-1033330203333210-1222301321013123)
