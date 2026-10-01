---
page_title: "xcsh_dns_zone landing"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone landing."
---

# xcsh_dns_zone landing

<a id="canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1f1684b60a7ae9c7262a06d011dbeb929d897ae1841a397d7a5ef672463fb22"></a>

## xcsh_dns_zone — xcsh_dns_zone / fc40251375ed / 2

Breadcrumbs:

- xcsh_dns_zone

Manages DNS Zone in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

<a id="canonical-38fc41d58ca6a07908c7306ac597c7f850e27954a4f180c51700aacf6efda544"></a>

## Prerequisites — xcsh_dns_zone / fc40251375ed / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `dns_load_balancer`.

- dns_load_balancer: Geographic or weighted DNS routing

<a id="canonical-9a2ea0f775e36882b4ae6b11064516a1dcb21281d6f1e3aa0b69aae1fc351905"></a>

## Minimal configuration — xcsh_dns_zone / fc40251375ed / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZone Resource Example
# Manages DNS Zone in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSZone configuration
resource "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"

  primary {
    allow_http_lb_managed_records = true
  }
}
```

<a id="canonical-f03dd32437aa62f54268ae66fb7a907431bbf5da5c4be264ee8f33618199d04e"></a>

## Root configuration — xcsh_dns_zone / fc40251375ed / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-ddb8e4198d92bdea05ad3dbbe3f9cf47c099a07baa8f448b94965727760c1c77"></a>

## Next pages — xcsh_dns_zone / fc40251375ed / 6

- [Property reference](../guides/resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [Examples](../guides/resources--dns_zone--examples--group-001.md#canonical-e233a06d738d768cdeb5ad8c949ffcb224dccf43bc25ebcdbe6a486857884d9d)
- [Import](../guides/resources--dns_zone--lifecycle--group-001.md#canonical-0c08bdf1ba0328c43b979140b6c71c4786db5e6afb1ee2413a1b2d886d1eacbd)
- [Timeouts](../guides/resources--dns_zone--lifecycle--group-001.md#canonical-b8ce35752a182927446792907375c210a306b53d3fe347606e7c87c7d9f576a9)
