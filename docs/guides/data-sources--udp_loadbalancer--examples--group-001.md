---
page_title: "xcsh_udp_loadbalancer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer examples."
---

# xcsh_udp_loadbalancer examples

<a id="canonical-3320030331211200-1233322332302002-3111203123123123-1000011112201332-1323220022102231-1010130310033302-2113232220113311-3310022022322132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- Examples

<a id="canonical-3213212023213322-0031122222312112-0332030012320132-1121121233122312-2032033020313313-0001321130213122-3003122212120200-0312333311311211"></a>

### Complete configurations for `xcsh_udp_loadbalancer`

- [Data source](data-sources--udp_loadbalancer--examples--group-001.md#canonical-0111330112313303-3132033201120320-3023232301303230-0011332301323002-0010010222331002-0222112320131300-2100121111310203-2021302010230113): valid configuration.

<a id="canonical-0111330112313303-3132033201120320-3023232301303230-0011332301323002-0010010222331002-0222112320131300-2100121111310203-2021302010230113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Examples](data-sources--udp_loadbalancer--examples--group-001.md#canonical-3320030331211200-1233322332302002-3111203123123123-1000011112201332-1323220022102231-1010130310033302-2113232220113311-3310022022322132)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_udp_loadbalancer/data-source.tf`; digest `sha256:90436c5e0e42b785b65d7370ba89219e7c36c5236adccb1abeed194590bc9460`.

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
