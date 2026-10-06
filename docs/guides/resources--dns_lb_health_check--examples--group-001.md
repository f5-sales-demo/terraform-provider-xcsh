---
page_title: "xcsh_dns_lb_health_check examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check examples."
---

# xcsh_dns_lb_health_check examples

<a id="canonical-0011113330130022-3300001313003320-3323320131301003-0200111123010021-0002311021212303-1202023122132130-2313132003300203-3211022032122203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- Examples

<a id="canonical-1013231221111200-3212330133002123-0012230033330110-2312003123212203-2233113122111032-1313022201323103-3213110133213111-3333302213302331"></a>

### Complete configurations for `xcsh_dns_lb_health_check`

- [Resource](resources--dns_lb_health_check--examples--group-001.md#canonical-1212210331222111-1213312020222310-0202113023013312-0002130200321233-1212003200221131-3131131133211100-3013031102332123-2121000232231201): valid configuration.

<a id="canonical-1212210331222111-1213312020222310-0202113023013312-0002130200321233-1212003200221131-3131131133211100-3013031102332123-2121000232231201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md#canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333)
- [Examples](resources--dns_lb_health_check--examples--group-001.md#canonical-0011113330130022-3300001313003320-3323320131301003-0200111123010021-0002311021212303-1202023122132130-2313132003300203-3211022032122203)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_lb_health_check/resource.tf`; digest `sha256:219de26c8a2c69aaf219ac20d2e05053708abed48480ac00f40c1cb16da5893e`.

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
