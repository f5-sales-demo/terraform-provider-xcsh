---
page_title: "xcsh_dns_load_balancer examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer examples."
---

# xcsh_dns_load_balancer examples

<a id="canonical-2003320123011322-3333121122213202-3231212110113333-2111133101120023-0032220332120233-1033311212000332-3323130012313222-2011322223300201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- Examples

<a id="canonical-2002301032213202-1023333133211320-3303012013323202-1002330302322022-1232000320300332-1223120011320301-3132003131120222-2330223202222020"></a>

### Complete configurations for `xcsh_dns_load_balancer`

- [Resource](resources--dns_load_balancer--examples--group-001.md#canonical-3122312210120211-2301233313110112-1223001302033003-3321333322000211-0233313321130303-3302100022220202-2101330001313100-0300131302133001): valid configuration.

<a id="canonical-3122312210120211-2301233313110112-1223001302033003-3321333322000211-0233313321130303-3302100022220202-2101330001313100-0300131302133001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Examples](resources--dns_load_balancer--examples--group-001.md#canonical-2003320123011322-3333121122213202-3231212110113333-2111133101120023-0032220332120233-1033311212000332-3323130012313222-2011322223300201)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_load_balancer/resource.tf`; digest `sha256:103fabb6495b3c7befdb807ff9ebc6d11ed6c79f3f4a4083e4f2815709c3a01c`.

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
