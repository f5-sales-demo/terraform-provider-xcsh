---
page_title: "xcsh_tcp_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer examples."
---

# xcsh_tcp_loadbalancer examples

<a id="canonical-1211231212111202-3031312222230212-0231230230223221-1311201002223122-3011202231020202-1113313023220220-3002202113210302-3322311231033013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- Examples

<a id="canonical-1111000020221101-3300332130321011-2210022313200232-2023201111300002-0020231211123331-1311211223012020-2131332120322103-0331210311331021"></a>

### Complete configurations for `xcsh_tcp_loadbalancer`

- [Data source](data-sources--tcp_loadbalancer--examples--group-001.md#canonical-0233201203120113-2030322203213203-0121222200220001-1100113313030002-1132311101133310-0203002030003103-3221013332201111-2231100001010010): valid configuration.

<a id="canonical-0233201203120113-2030322203213203-0121222200220001-1100113313030002-1132311101133310-0203002030003103-3221013332201111-2231100001010010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Examples](data-sources--tcp_loadbalancer--examples--group-001.md#canonical-1211231212111202-3031312222230212-0231230230223221-1311201002223122-3011202231020202-1113313023220220-3002202113210302-3322311231033013)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tcp_loadbalancer/data-source.tf`; digest `sha256:8aadc1f29b777e99791395cd3d6a0a6d69101be8986adffcec19d0c93ccfd1f9`.

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
