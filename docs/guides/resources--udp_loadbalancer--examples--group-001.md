---
page_title: "xcsh_udp_loadbalancer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer examples."
---

# xcsh_udp_loadbalancer examples

<a id="canonical-1102221033122331-2330033102010213-3323032213330232-0023202131310322-0301012112303023-1300202102011232-3320002211022203-3211001223100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- Examples

<a id="canonical-0120111113221312-2102020012300320-3300013301103321-0123012000131110-2123003002132220-2021303001120001-0030100010231220-2232321213312221"></a>

### Complete configurations for `xcsh_udp_loadbalancer`

- [Resource](resources--udp_loadbalancer--examples--group-001.md#canonical-1301313203231213-3220311132100110-0330103211030102-0333331101201213-3012011110303223-3023132303012223-3010301333232332-0310230313101110): valid configuration.

<a id="canonical-1301313203231213-3220311132100110-0330103211030102-0333331101201213-3012011110303223-3023132303012223-3010301333232332-0310230313101110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Examples](resources--udp_loadbalancer--examples--group-001.md#canonical-1102221033122331-2330033102010213-3323032213330232-0023202131310322-0301012112303023-1300202102011232-3320002211022203-3211001223100112)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_udp_loadbalancer/resource.tf`; digest `sha256:f54b33fab2f01b75ab0f74a179601bfe368063a44f1d292cd4b49a5f0936b681`.

```terraform
# UDPLoadBalancer Resource Example
# Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UDPLoadBalancer configuration
resource "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}
```
