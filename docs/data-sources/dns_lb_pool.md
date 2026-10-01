---
page_title: "xcsh_dns_lb_pool landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool landing."
---

# xcsh_dns_lb_pool landing

<a id="canonical-fd4e55a1af935bf15f770be42dd051ab69016ab1078d5599a235a350bbecd504"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a63fa2e5f338823e1b37e74a508e09293fdc7f196b00599c4a93b9048a1953f"></a>

## xcsh_dns_lb_pool — xcsh_dns_lb_pool / 737348781cea / 2

Breadcrumbs:

- xcsh_dns_lb_pool

Manages DNS Load Balancer Pool in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

<a id="canonical-7a63334769584faab179d4bb4ae08ba1844728e2f30fc07ee5654acdff2bbd5d"></a>

## Prerequisites — xcsh_dns_lb_pool / 737348781cea / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-307b19dbb396e34cfa3ec32fe267fc31e54aef6d2a6fd3e4d8e5acc4ff7fa8bf"></a>

## Minimal configuration — xcsh_dns_lb_pool / 737348781cea / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-010787bd95dcffa2a6f5ac670aed877694410fa2ec35be7ab2299b96003fa1c3"></a>

## Root configuration — xcsh_dns_lb_pool / 737348781cea / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-4a6f9c4770610a800abd5e2eddbc7ce7a2885f34399cf78522224abc5e481fea"></a>

## Next pages — xcsh_dns_lb_pool / 737348781cea / 6

- [Property reference](../guides/data-sources--dns_lb_pool--reference--group-001.md#canonical-a10ef0d228407eacc80437b5921205ffa6f9165ac0a64e1222cf512967e4630f)
- [Examples](../guides/data-sources--dns_lb_pool--examples--group-001.md#canonical-d25e99761fcf8f6618d1dbd05004130cdb72a5ebac2fd0915895e497d5f03a82)
