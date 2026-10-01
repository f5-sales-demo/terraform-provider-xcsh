---
page_title: "xcsh_dns_zone landing"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone landing."
---

# xcsh_dns_zone landing

<a id="canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001330112201023-1200221322322130-1302120222001231-0001013123322321-0221312021132232-0120100122032113-3113221132331213-0210120333230202"></a>

## xcsh_dns_zone — xcsh_dns_zone / 010313113231 / 2

Breadcrumbs:

- xcsh_dns_zone

Manages DNS Zone in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

<a id="canonical-0320333010013111-2030221222001321-0020301303001222-3011211330133320-1100320213211110-2210330120003011-0113000022223033-1232333122111010"></a>

## Prerequisites — xcsh_dns_zone / 010313113231 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `dns_load_balancer`.

- dns_load_balancer: Geographic or weighted DNS routing

<a id="canonical-2122023222003313-1311320312202002-2310223212230101-0012101101122201-3130230201022001-3112330132032222-0023122122223201-3330031101210011"></a>

## Minimal configuration — xcsh_dns_zone / 010313113231 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZone Resource Example
# Manages DNS Zone in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSZone configuration
resource "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"

  primary {
    allow_http_lb_managed_records = true
  }
}
```

<a id="canonical-3300033131030210-0313222212023311-1002122022321212-3323132221001310-0301232333113122-1130102332021210-3232203303031201-2001212131001032"></a>

## Root configuration — xcsh_dns_zone / 010313113231 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3131232032100121-2031210223313222-0011223103312323-3203332130331013-3000212122001323-2222203310102023-2110211211130213-1312003001301313"></a>

## Next pages — xcsh_dns_zone / 010313113231 / 6

- [Property reference](../guides/resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [Examples](../guides/resources--dns_zone--examples--group-001.md#canonical-3202030322001231-1303203113122030-3132231122312030-2110213333302302-0210313030331003-2330021132233031-2332122210201220-1113202010312131)
- [Import](../guides/resources--dns_zone--lifecycle--group-001.md#canonical-0030002023313301-2322000302203010-0323211321011000-2312301301301013-2012312311321222-3323013232021001-0322012302312020-1231013222302331)
- [Timeouts](../guides/resources--dns_zone--lifecycle--group-001.md#canonical-2320303203111311-0222012002210213-1010121321022100-1303131130020100-2203001223110331-0333320310131200-1232133020133013-3121331113122221)
