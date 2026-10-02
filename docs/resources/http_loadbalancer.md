---
page_title: "xcsh_http_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer landing."
---

# xcsh_http_loadbalancer landing

<a id="canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223332301230323-2210121033000312-0222313110012322-3012133300210011-3103332112022202-2110203333011002-2333032233023133-3110222113323131"></a>

## xcsh_http_loadbalancer — xcsh_http_loadbalancer / 322302311200 / 2

Breadcrumbs:

- xcsh_http_loadbalancer

Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic
with routing and security controls.

<a id="canonical-0010223003213013-0223113213012002-1032120113312232-2131132213233001-1313023123220123-2223000132023100-1310322100213230-2002301310321300"></a>

## Prerequisites — xcsh_http_loadbalancer / 322302311200 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`, `app_firewall`, `certificate`, `rate_limiter`, `service_policy`, `bot_defense_policy`.

- origin_pool: Backend servers for traffic distribution

- app_firewall: WAF protection (requires WAAP subscription)

- healthcheck: Monitor backend availability

- certificate: TLS termination for HTTPS

- rate_limiter: Protect against traffic spikes

<a id="canonical-3320013231113211-2132001313133021-2302232102030233-1312302320322021-3131231111222221-2011102212102223-3230113012320110-3331202120212013"></a>

## Minimal configuration — xcsh_http_loadbalancer / 322302311200 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# HTTPLoadBalancer Resource Example
# Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with routing and security controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic HTTPLoadBalancer configuration
resource "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

<a id="canonical-3013301201222230-1223312302030331-1301203230310211-0003203123212311-1103002013021012-3302310010022131-1101000032033020-3303113221131123"></a>

## Root configuration — xcsh_http_loadbalancer / 322302311200 / 5

Required root properties: `domains`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2130131011321001-3102030200221112-3013031220111200-1132022000023122-2103002222332131-2303232311222233-3223322233113113-3112221010203110"></a>

## Next pages — xcsh_http_loadbalancer / 322302311200 / 6

- [Property reference](../guides/resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [Examples](../guides/resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- [Import](../guides/resources--http_loadbalancer--lifecycle--group-001.md#canonical-1000101212023023-0103011103132001-3301102330230310-1230010102120212-3223112021023333-2200311133300111-1031101210100332-2211331001010311)
- [Timeouts](../guides/resources--http_loadbalancer--lifecycle--group-001.md#canonical-2232121223123231-2022121130300121-0103113012231013-1023001100002301-2121003313103030-1100033200202130-2111230110133210-3020020103231013)
- [Lifecycle](../guides/resources--http_loadbalancer--lifecycle--group-001.md#canonical-2201010321200222-2110320331131013-1013211030201130-0130323333132231-1030203101130301-0030312322132221-0233000232333000-3101312320200212)
