---
page_title: "xcsh_udp_loadbalancer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer landing."
---

# xcsh_udp_loadbalancer landing

<a id="canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301303021013312-3101303022030300-2113112132123312-1202121101223212-0220301320101003-0013023032303120-3013211310330213-2312032023110220"></a>

## xcsh_udp_loadbalancer — xcsh_udp_loadbalancer / 023312302133 / 2

Breadcrumbs:

- xcsh_udp_loadbalancer

Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across
origin pools.

<a id="canonical-1332321222132201-0203202231131311-3022013030313130-3100020110033300-0232003020331012-3330132120122333-3203000311020002-3301220103010220"></a>

## Prerequisites — xcsh_udp_loadbalancer / 023312302133 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1023221211030332-1030202023210122-0123001323303132-3302112233330312-3010332001010332-3103232021330212-2233023022002020-2303210002011020"></a>

## Minimal configuration — xcsh_udp_loadbalancer / 023312302133 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UDPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UDPLoadBalancer by name
data "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}

output "udp_loadbalancer_id" {
  value = data.xcsh_udp_loadbalancer.example.id
}
```

<a id="canonical-2201112130212001-3222003122223223-3023112231302202-3031233131121223-3120003231311321-2321322023303333-3022031032122202-2102131211231200"></a>

## Root configuration — xcsh_udp_loadbalancer / 023312302133 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3220312330301111-3312020122301121-2032021001310323-3320130001002023-3312003012322201-1221002110010223-1133231033212301-0111100330311232"></a>

## Next pages — xcsh_udp_loadbalancer / 023312302133 / 6

- [Property reference](../guides/data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [Examples](../guides/data-sources--udp_loadbalancer--examples--group-001.md#canonical-3320030331211200-1233322332302002-3111203123123123-1000011112201332-1323220022102231-1010130310033302-2113232220113311-3310022022322132)
