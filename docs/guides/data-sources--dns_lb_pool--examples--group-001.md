---
page_title: "xcsh_dns_lb_pool examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool examples."
---

# xcsh_dns_lb_pool examples

<a id="canonical-d25e99761fcf8f6618d1dbd05004130cdb72a5ebac2fd0915895e497d5f03a82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f45127a56fa7f48c7568c5eb01014f1b9c72d17a0378ac5975879c85c83c3460"></a>

## Examples — Examples / 9a8472388262 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- Examples

<a id="canonical-52541755777ae503c90a7fa069921c831e24bc728603d45d4e9187d5d8938ff1"></a>

## Complete configurations — Examples / 9a8472388262 / 3

- [Data source](data-sources--dns_lb_pool--examples--group-001.md#canonical-cb83c7f1dcb6c9c5379a3e6014a9479166838b12fb752fbd53df39c94f180eb6): valid configuration.

<a id="canonical-1c3dc43d5ea6c6ee48158eb0e17dffdb2195f64e6a8018f1c843779ecf7ec454"></a>

## Next pages — Examples / 9a8472388262 / 4

- [Data source](data-sources--dns_lb_pool--examples--group-001.md#canonical-cb83c7f1dcb6c9c5379a3e6014a9479166838b12fb752fbd53df39c94f180eb6)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)

<a id="canonical-cb83c7f1dcb6c9c5379a3e6014a9479166838b12fb752fbd53df39c94f180eb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8005b1b54c5f40973d6e96b949fdb420381ae43995587eac31d6782cf236861e"></a>

## Data source — Data source / 243e7a6a2488 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
- [Examples](data-sources--dns_lb_pool--examples--group-001.md#canonical-d25e99761fcf8f6618d1dbd05004130cdb72a5ebac2fd0915895e497d5f03a82)
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

<a id="canonical-3ea91378d1eb1ccbe04e3892b7de05a8276838126994cac66fa72728a536c1ef"></a>

## Next pages — Data source / 243e7a6a2488 / 3

- [Examples](data-sources--dns_lb_pool--examples--group-001.md#canonical-d25e99761fcf8f6618d1dbd05004130cdb72a5ebac2fd0915895e497d5f03a82)
- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504)
