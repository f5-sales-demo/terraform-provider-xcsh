---
page_title: "xcsh_cdn_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer landing."
---

# xcsh_cdn_loadbalancer landing

<a id="canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223333203022111-0223221202311001-2322200211113232-0310303302103133-0200102221313111-0321212313303320-0223200033201230-3021322223013221"></a>

## xcsh_cdn_loadbalancer — xcsh_cdn_loadbalancer / 001212320200 / 2

Breadcrumbs:

- xcsh_cdn_loadbalancer

Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching
with load balancing.

<a id="canonical-0332213023210023-3132200223101010-1311100120202010-2022023020101112-0332311000112310-1132113120313301-3023202231333231-2321310132321302"></a>

## Prerequisites — xcsh_cdn_loadbalancer / 001212320200 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cdn_origin_pool`.

- cdn_origin_pool: Origin servers for CDN content

<a id="canonical-3132123113020330-3121202301000023-0323133132333113-3022202231331032-2103022121132023-1233201212122231-1020301113132030-2212211303100300"></a>

## Minimal configuration — xcsh_cdn_loadbalancer / 001212320200 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNLoadBalancer Resource Example
# Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNLoadBalancer configuration
resource "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

<a id="canonical-1122022021331312-0310020322310012-1003230020221121-2100002320333011-1332021121013321-1322020233033210-2003231320212132-0201203133333311"></a>

## Root configuration — xcsh_cdn_loadbalancer / 001212320200 / 5

Required root properties: `domains`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3001331100322231-3302102001312223-1132133031131110-1003301212223313-1230231113323311-0321122032010312-1000321121221133-0131121223223200"></a>

## Next pages — xcsh_cdn_loadbalancer / 001212320200 / 6

- [Property reference](../guides/resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [Examples](../guides/resources--cdn_loadbalancer--examples--group-001.md#canonical-2223333103031111-2331013121000300-2231102333110312-1231112211103303-1003012200322202-1023032201331100-3322113303302113-1213200102031112)
- [Import](../guides/resources--cdn_loadbalancer--lifecycle--group-001.md#canonical-0113203320131311-2001120113302103-1202220232210031-1210210002111130-1222002311032311-2200223111013300-1112121120212322-0211212113231331)
- [Timeouts](../guides/resources--cdn_loadbalancer--lifecycle--group-001.md#canonical-0102122123213023-0022323222303012-2133031130101132-2201120221023002-1213330321103111-3301322222203110-3223300130332103-3302303232132102)
