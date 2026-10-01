---
page_title: "xcsh_dns_load_balancer landing"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer landing."
---

# xcsh_dns_load_balancer landing

<a id="canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220110323322002-1021302021331331-1301220312300231-2130000123223022-1011223003001111-3303233131331111-1032031332213131-0000332103202033"></a>

## xcsh_dns_load_balancer — xcsh_dns_load_balancer / 323312201001 / 2

Breadcrumbs:

- xcsh_dns_load_balancer

Manages DNS Load Balancer in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-1300021131312321-0111101101230231-1231001132011320-3031320121013000-0313211220013332-1111222331001201-2130011023023202-1110331200121002"></a>

## Prerequisites — xcsh_dns_load_balancer / 323312201001 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `dns_zone`.

- dns_zone: Parent zone for DNS records

<a id="canonical-0112312302221320-3020203022110122-1131212111311020-0311111313202120-2210332212223213-2223210230302300-1310102101121002-0113112000230021"></a>

## Minimal configuration — xcsh_dns_load_balancer / 323312201001 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLoadBalancer Resource Example
# Manages DNS Load Balancer in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLoadBalancer configuration
resource "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}
```

<a id="canonical-0130013233220033-1311022233220311-2110021000010131-0020013300200032-1200120013120121-3110231223013322-1010313233303131-2120322313210223"></a>

## Root configuration — xcsh_dns_load_balancer / 323312201001 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2012012200002200-3023132123010000-2101323021133331-2212103021311130-1220232232013102-3322021111313033-1333321322002030-3302232333101032"></a>

## Next pages — xcsh_dns_load_balancer / 323312201001 / 6

- [Property reference](../guides/resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [Examples](../guides/resources--dns_load_balancer--examples--group-001.md#canonical-2003320123011322-3333121122213202-3231212110113333-2111133101120023-0032220332120233-1033311212000332-3323130012313222-2011322223300201)
- [Import](../guides/resources--dns_load_balancer--lifecycle--group-001.md#canonical-0130210010103222-1210113203203231-3123131212201010-0212231011303020-0311103213211033-1102302312321331-2031301213323123-1323130320030110)
- [Timeouts](../guides/resources--dns_load_balancer--lifecycle--group-001.md#canonical-2023332301100233-3231132221301023-3311210031320012-3232232133322230-3332031130021023-0103222301012031-3031302033131120-3030001222321203)
