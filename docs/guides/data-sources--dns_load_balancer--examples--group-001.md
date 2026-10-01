---
page_title: "xcsh_dns_load_balancer examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer examples."
---

# xcsh_dns_load_balancer examples

<a id="canonical-53377b934401d143eba0bcda241ab1eabf8ad0e63f31cec3dd8a9ba854484274"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-292144784886c1fe4410893b63fcad42c00dd04d9344c7efa6f25655b394be1d"></a>

## Examples — Examples / 33afec71b276 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- Examples

<a id="canonical-698874a0d52d0c54b7b407d815e1b7b61288d83aec521b796e7458a38562df49"></a>

## Complete configurations — Examples / 33afec71b276 / 3

- [Data source](data-sources--dns_load_balancer--examples--group-001.md#canonical-93ca0d79def6781101ba965c63c46b81f42ad08fd6966c39ca612890f20cff53): valid configuration.

<a id="canonical-c5e9debe152d1ab3510663435d6981473421ce3c80c6b46c6875d75ac1fbeecb"></a>

## Next pages — Examples / 33afec71b276 / 4

- [Data source](data-sources--dns_load_balancer--examples--group-001.md#canonical-93ca0d79def6781101ba965c63c46b81f42ad08fd6966c39ca612890f20cff53)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)

<a id="canonical-93ca0d79def6781101ba965c63c46b81f42ad08fd6966c39ca612890f20cff53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33bf65cad319bef42afd2904a749cc4ad333b8b0df83ef48a6ad55d8c125cec3"></a>

## Data source — Data source / 520180c0ad8a / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
- [Examples](data-sources--dns_load_balancer--examples--group-001.md#canonical-53377b934401d143eba0bcda241ab1eabf8ad0e63f31cec3dd8a9ba854484274)
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

<a id="canonical-e4b15401cda7225ac99928939947a9e4d2a90e18d4e700e153cb058ced0d57a6"></a>

## Next pages — Data source / 520180c0ad8a / 3

- [Examples](data-sources--dns_load_balancer--examples--group-001.md#canonical-53377b934401d143eba0bcda241ab1eabf8ad0e63f31cec3dd8a9ba854484274)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-639f0fcc6176dfa735154dc2a6993c8c3f13473206fe7b68e03e66d22c4c0a93)
