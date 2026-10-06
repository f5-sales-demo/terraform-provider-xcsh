---
page_title: "xcsh_dns_load_balancer examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer examples."
---

# xcsh_dns_load_balancer examples

<a id="canonical-1103031313232103-1010000131011003-3223220023303122-0210012223013222-2333202231003212-0333030130323003-3131202221232220-1110102010021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- Examples

<a id="canonical-0221020110101320-1020201230013332-1010010020210323-1203333022311002-3000003131001031-2103101030133233-2212330211121111-2303211023320131"></a>

### Complete configurations for `xcsh_dns_load_balancer`

- [Data source](data-sources--dns_load_balancer--examples--group-001.md#canonical-2103302200311321-3132331213200101-0001232221121130-1203301012232001-3310022231002033-3112211212300321-3022120102202100-3302003033331103): valid configuration.

<a id="canonical-2103302200311321-3132331213200101-0001232221121130-1203301012232001-3310022231002033-3112211212300321-3022120102202100-3302003033331103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Examples](data-sources--dns_load_balancer--examples--group-001.md#canonical-1103031313232103-1010000131011003-3223220023303122-0210012223013222-2333202231003212-0333030130323003-3131202221232220-1110102010021310)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_load_balancer/data-source.tf`; digest `sha256:94e6c40d8f676ae588608f07a90134d4cad49dd9a1bfd873d996f0534dfb0e30`.

```terraform
# DNSLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLoadBalancer by name
data "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}

output "dns_load_balancer_id" {
  value = data.xcsh_dns_load_balancer.example.id
}
```
