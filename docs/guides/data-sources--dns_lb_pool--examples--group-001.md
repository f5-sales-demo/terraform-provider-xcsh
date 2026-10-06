---
page_title: "xcsh_dns_lb_pool examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool examples."
---

# xcsh_dns_lb_pool examples

<a id="canonical-3102113221211312-0133303320331212-0120310131233100-1100001001030030-3123130222113223-2230023331002101-1120211132102113-3111330003222002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- Examples

<a id="canonical-3310110102132211-1233221333102030-1311122030113223-0001000110330123-2130130231011322-0003132022301121-1311201321302011-3020033003101200"></a>

### Complete configurations for `xcsh_dns_lb_pool`

- [Data source](data-sources--dns_lb_pool--examples--group-001.md#canonical-3023200330133301-3130231230213011-0313212203321200-0110222110132101-1212200320230102-3323131102332331-1103313303213021-1033012000322312): valid configuration.

<a id="canonical-3023200330133301-3130231230213011-0313212203321200-0110222110132101-1212200320230102-3323131102332331-1103313303213021-1033012000322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Examples](data-sources--dns_lb_pool--examples--group-001.md#canonical-3102113221211312-0133303320331212-0120310131233100-1100001001030030-3123130222113223-2230023331002101-1120211132102113-3111330003222002)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_lb_pool/data-source.tf`; digest `sha256:9dddbd7b299906535a7d4a7f3de8f5eb6fae61134e8af1bbe57c031c456e2613`.

```terraform
# DNSLBPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLBPool by name
data "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}

output "dns_lb_pool_id" {
  value = data.xcsh_dns_lb_pool.example.id
}
```
