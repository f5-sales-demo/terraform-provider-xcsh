---
page_title: "xcsh_dns_zone_delete_cryptokey landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_delete_cryptokey landing."
---

# xcsh_dns_zone_delete_cryptokey landing

<a id="canonical-d0f11b9ff55eda3e0a21dbaf9c1ae6f999a193083bac154e03bd47c8807d0123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-594c97c8314cebb12670c1ed475c3b7fc1ad8faae9eafa7172d496f0eebce0ab"></a>

## xcsh_dns_zone_delete_cryptokey — xcsh_dns_zone_delete_cryptokey / ee64ce89ef79 / 2

Breadcrumbs:

- xcsh_dns_zone_delete_cryptokey

Resource creation operation.

<a id="canonical-cc656127d92880918424f93610c2544a0fe941b366d44b58219cfb9cffab172b"></a>

## Prerequisites — xcsh_dns_zone_delete_cryptokey / ee64ce89ef79 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-5c61436f843f11815fd04460dd827a905c7617ae1e9ce3d11b89fc93caa09ca9"></a>

## Minimal configuration — xcsh_dns_zone_delete_cryptokey / ee64ce89ef79 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneDeleteCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_delete_cryptokey" "example" {
  config {
  }
}
```

<a id="canonical-d413ec7f153ead114ca9895836d7939298eb2298511cdfd5860e75d0e66763ac"></a>

## Root configuration — xcsh_dns_zone_delete_cryptokey / ee64ce89ef79 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-81f8b742a69625a26ea119f74e4dcbdeacc6c5646a7c10d788a77185f9eaa2d6"></a>

## Next pages — xcsh_dns_zone_delete_cryptokey / ee64ce89ef79 / 6

- [Property reference](../guides/actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-53390e8981cf8014bda7b0838c903adf614b4fc57c56d4927129396bb9cc67aa)
- [Examples](../guides/actions--dns_zone_delete_cryptokey--examples--group-001.md#canonical-d4f9c1c2be0024b3c20e17bc93c62403bbfa1c88bcd9c1938bd26f58ec13b7e3)
- [Lifecycle](../guides/actions--dns_zone_delete_cryptokey--lifecycle--group-001.md#canonical-cd762a2dccab2db0196a9bc60d3071027a2d68b2c81f1e10698311d1f1ea92ab)
