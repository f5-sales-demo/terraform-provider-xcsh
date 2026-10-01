---
page_title: "xcsh_dns_lb_pool landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool landing."
---

# xcsh_dns_lb_pool landing

<a id="canonical-3023312023322121-1321331121122122-1001102332013330-2032232201120022-3030001133222220-0010101132132022-1001320122320233-3231010023122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010100023123311-0221323113032212-0221131111003101-1320330101013123-3203011011122201-1321123321323303-2313020312210102-1322113003130320"></a>

## xcsh_dns_lb_pool — xcsh_dns_lb_pool / 001213322111 / 2

Breadcrumbs:

- xcsh_dns_lb_pool

Manages DNS Load Balancer Pool in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-0001123133031312-3010023112302302-0103211213022301-2231030213021000-2023101201231330-3030212201030320-2110220101232203-3030210120303110"></a>

## Prerequisites — xcsh_dns_lb_pool / 001213322111 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0220332123121233-3132122123303103-1213210312213023-2312033233103023-1002112333011331-1030103001010212-1103203121132101-1120321321123213"></a>

## Minimal configuration — xcsh_dns_lb_pool / 001213322111 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLBPool Resource Example
# Manages DNS Load Balancer Pool in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBPool configuration
resource "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}
```

<a id="canonical-3210122331232333-2000022300103320-0003202233020220-3020122303112032-1333030223232330-3130333100020312-2123321002233302-2220110230130022"></a>

## Root configuration — xcsh_dns_lb_pool / 001213322111 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3200101212113312-3332213212220030-3202021231100202-0123103200112222-0032132030331201-0131110121013002-0011103010233120-0121223101133121"></a>

## Next pages — xcsh_dns_lb_pool / 001213322111 / 6

- [Property reference](../guides/resources--dns_lb_pool--reference--group-001.md#canonical-3200110131203200-2231001023331323-3322203230200001-0123131301121313-2113221003233203-3022101200312011-2121220212322312-1003102310203031)
- [Examples](../guides/resources--dns_lb_pool--examples--group-001.md#canonical-3021310113213213-0120010022310331-0010123011100200-0333131131321313-1322121133200223-1032031202330333-3023111233300100-2320221233303000)
- [Import](../guides/resources--dns_lb_pool--lifecycle--group-001.md#canonical-1002022001013301-1011001213213310-1330111101332012-2031213323300312-1300202313100221-1203332221113112-2201022010133213-1003022230011103)
- [Timeouts](../guides/resources--dns_lb_pool--lifecycle--group-001.md#canonical-2212312212223321-1330320020112320-3202212323210200-3203113200113331-1221131301303003-2112223023221301-3323212221022100-2232131210203222)
