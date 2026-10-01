---
page_title: "xcsh_dns_lb_pool landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool landing."
---

# xcsh_dns_lb_pool landing

<a id="canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022120333220232-1133030320200203-3201230313321310-2211002032002102-2103333130133301-2112230000112121-3010222103232100-1020220121110333"></a>

## xcsh_dns_lb_pool — xcsh_dns_lb_pool / 132001303222 / 2

Breadcrumbs:

- xcsh_dns_lb_pool

Manages DNS Load Balancer Pool in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-1322120303031013-1221112010332222-2301132131102323-1022320020232201-2010101302203202-3303003330001332-3211121110223031-3333022323311131"></a>

## Prerequisites — xcsh_dns_lb_pool / 132001303222 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0300132301213123-2303211232031030-3322033230030233-3202121333300301-3211102232331231-0222123331033210-3120321122303010-3333133322202333"></a>

## Minimal configuration — xcsh_dns_lb_pool / 132001303222 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLBPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLBPool by name
data "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}

output "dns_lb_pool_id" {
  value = data.xcsh_dns_lb_pool.example.id
}
```

<a id="canonical-0001001320132331-2111313033332202-2212331122301213-0022323120131312-2110100100332202-3230031123321322-2302022121232112-0000033322013003"></a>

## Root configuration — xcsh_dns_lb_pool / 132001303222 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1022123321301013-1300120100222000-0022233111320232-3131233013303213-2202202011330310-0321213033132011-0202020210222330-1132102001333222"></a>

## Next pages — xcsh_dns_lb_pool / 132001303222 / 6

- [Property reference](../guides/data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [Examples](../guides/data-sources--dns_lb_pool--examples--group-001.md#canonical-3102113221211312-0133303320331212-0120310131233100-1100001001030030-3123130222113223-2230023331002101-1120211132102113-3111330003222002)
