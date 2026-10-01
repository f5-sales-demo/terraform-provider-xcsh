---
page_title: "xcsh_tcp_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer landing."
---

# xcsh_tcp_loadbalancer landing

<a id="canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321213323120103-0202332002021330-1311020330030123-3101213132330023-3222131210133103-1020020210130212-0031020333101212-0103223332312132"></a>

## xcsh_tcp_loadbalancer — xcsh_tcp_loadbalancer / 320231130103 / 2

Breadcrumbs:

- xcsh_tcp_loadbalancer

Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across
origin pools.

<a id="canonical-1013013231010030-3100221032112320-2033300031012313-3112203123013123-0110331000230011-2131003121333320-1310222021310110-3021201020221321"></a>

## Prerequisites — xcsh_tcp_loadbalancer / 320231130103 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`.

- origin_pool: Backend servers for TCP/UDP traffic

- healthcheck: Monitor origin server health

<a id="canonical-1113303103311202-1033133110012010-3100302002211122-2302223123321322-2301111220221121-1003122101021112-1133011301033222-0203323322100312"></a>

## Minimal configuration — xcsh_tcp_loadbalancer / 320231130103 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TCPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TCPLoadBalancer by name
data "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}

output "tcp_loadbalancer_id" {
  value = data.xcsh_tcp_loadbalancer.example.id
}
```

<a id="canonical-2133110200222211-0210330132033322-2300111320312101-0101231023013132-1313301101320212-3231110001033102-0232330122032002-0022112103223113"></a>

## Root configuration — xcsh_tcp_loadbalancer / 320231130103 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1321301021203213-0133222131112022-3220020131212130-1032211131010200-0300110210012333-0031130310003001-3320111313131222-0302111320311012"></a>

## Next pages — xcsh_tcp_loadbalancer / 320231130103 / 6

- [Property reference](../guides/data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [Examples](../guides/data-sources--tcp_loadbalancer--examples--group-001.md#canonical-1211231212111202-3031312222230212-0231230230223221-1311201002223122-3011202231020202-1113313023220220-3002202113210302-3322311231033013)
