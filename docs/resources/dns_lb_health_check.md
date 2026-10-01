---
page_title: "xcsh_dns_lb_health_check landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check landing."
---

# xcsh_dns_lb_health_check landing

<a id="canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010031131321123-2331122320013210-3233111101031331-3311121012300230-0030303201130201-3011023003020302-2332332113121103-0302002033200003"></a>

## xcsh_dns_lb_health_check — xcsh_dns_lb_health_check / 112301022031 / 2

Breadcrumbs:

- xcsh_dns_lb_health_check

Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a
error in F5 Distributed Cloud.

<a id="canonical-3001301011310002-2212000111322131-1233310222201031-3202301031130210-3123001020332232-2120321113220323-3223021003132232-2121123220211012"></a>

## Prerequisites — xcsh_dns_lb_health_check / 112301022031 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3010120133033220-1212120033323223-3003223133312303-1100210011022012-0103321332102020-2222232100203322-2301212132012002-1020331130122301"></a>

## Minimal configuration — xcsh_dns_lb_health_check / 112301022031 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLBHealthCheck Resource Example
# Manages DNS Load Balancer Health Check in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBHealthCheck configuration
resource "xcsh_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}
```

<a id="canonical-1232130033003032-1333210022020012-3220311101000033-1312101223323321-2020021022023122-0023032301230121-1233002133332303-0113223233123032"></a>

## Root configuration — xcsh_dns_lb_health_check / 112301022031 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2002331031232010-2032132110311103-2101101131013132-0211101300320031-3230000311003103-1010301202311323-2121301202133220-3320121231101010"></a>

## Next pages — xcsh_dns_lb_health_check / 112301022031 / 6

- [Property reference](../guides/resources--dns_lb_health_check--reference--group-001.md#canonical-1123001131012002-2102110230010013-2322302212231031-1022032210330113-3031302203112210-3133030011231302-3030013110113220-2123320331203100)
- [Examples](../guides/resources--dns_lb_health_check--examples--group-001.md#canonical-0011113330130022-3300001313003320-3323320131301003-0200111123010021-0002311021212303-1202023122132130-2313132003300203-3211022032122203)
- [Import](../guides/resources--dns_lb_health_check--lifecycle--group-001.md#canonical-3333031110123202-2201201202302210-3130133010220102-2202330101010300-0311012112221312-3323322102322330-2012213322300031-2120211320330223)
- [Timeouts](../guides/resources--dns_lb_health_check--lifecycle--group-001.md#canonical-0203031303320223-1303201222123000-3210123031301122-1221013302300213-2232113330100230-0231221211301313-3323303330103111-1310110331112010)
