---
page_title: "xcsh_dns_zone landing"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone landing."
---

# xcsh_dns_zone landing

<a id="canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab819fba5262c16be6f2ca7295b06b6ed0c0adf9f65a2efe81e319767c12f5b7"></a>

## xcsh_dns_zone — xcsh_dns_zone / 53e427031169 / 2

Breadcrumbs:

- xcsh_dns_zone

Manages DNS Zone in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

<a id="canonical-ae8ca94b5a0f42738a1cf8783a8c65fcee1259c7d5bc3f1076dec53af37d9682"></a>

## Prerequisites — xcsh_dns_zone / 53e427031169 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `dns_load_balancer`.

- dns_load_balancer: Geographic or weighted DNS routing

<a id="canonical-59297d8fd69b1daff8994065860ecfae713a1393f40d4442616ed6dc65db53d5"></a>

## Minimal configuration — xcsh_dns_zone / 53e427031169 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZone Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSZone by name
data "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"
}

# Fail closed when this stack depends on an externally owned zone.
resource "terraform_data" "require_managed_records" {
  lifecycle {
    precondition {
      condition = try(
        data.xcsh_dns_zone.example.primary.allow_http_lb_managed_records,
        false
      )
      error_message = "The selected DNS zone must enable HTTP LB managed records."
    }
  }
}

output "dns_zone_id" {
  value = data.xcsh_dns_zone.example.id
}
```

<a id="canonical-75d1ed77fc756d063cb58cdee5d83cdcc9872938da8a2bf3aa65a422d8927d64"></a>

## Root configuration — xcsh_dns_zone / 53e427031169 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-ff5298a3152eb452f992b1008bd07c580d0beaff17a3705afdc81e684f5af9fd"></a>

## Next pages — xcsh_dns_zone / 53e427031169 / 6

- [Property reference](../guides/data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [Examples](../guides/data-sources--dns_zone--examples--group-001.md#canonical-17f8d062238fb35a54de166b256b564642f2aadc0f23fe04b8a8b988b597ea70)
