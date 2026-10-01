---
page_title: "xcsh_dns_load_balancer landing"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer landing."
---

# xcsh_dns_load_balancer landing

<a id="canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322031313020131-2332122323331301-1030301011113021-3130313021333130-0213213310323003-0321300211230133-0333122310332330-3011202232202331"></a>

## xcsh_dns_load_balancer — xcsh_dns_load_balancer / 110333031210 / 2

Breadcrumbs:

- xcsh_dns_load_balancer

Manages DNS Load Balancer in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-3213121031121123-0311103221111023-3033221303331022-0030303001003231-3022031111123313-1321312322322202-2011302033230101-1211032013011211"></a>

## Prerequisites — xcsh_dns_load_balancer / 110333031210 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `dns_zone`.

- dns_zone: Parent zone for DNS records

<a id="canonical-3212111101021303-3211211213311310-2112231010033012-1113100123302213-1010333212133221-1122131031230311-3210131132110020-3333011301002103"></a>

## Minimal configuration — xcsh_dns_load_balancer / 110333031210 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLoadBalancer by name
data "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}

output "dns_load_balancer_id" {
  value = data.xcsh_dns_load_balancer.example.id
}
```

<a id="canonical-3032130032003220-2112233000110321-0220233220311122-0310200212032032-1222002310213321-1033320112322330-1103310001122112-1201200210222221"></a>

## Root configuration — xcsh_dns_load_balancer / 110333031210 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0311133333213213-3110231311232201-3333230120320100-0231123332300233-3331313033112330-3213021120221330-0122012012223310-2112203120023110"></a>

## Next pages — xcsh_dns_load_balancer / 110333031210 / 6

- [Property reference](../guides/data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [Examples](../guides/data-sources--dns_load_balancer--examples--group-001.md#canonical-1103031313232103-1010000131011003-3223220023303122-0210012223013222-2333202231003212-0333030130323003-3131202221232220-1110102010021310)
