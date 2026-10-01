---
page_title: "xcsh_dns_zone_add_cryptokey landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_add_cryptokey landing."
---

# xcsh_dns_zone_add_cryptokey landing

<a id="canonical-85ff7b5ce0689a3d0497afdd728eedfead55ce6ef5288db5ba75c83a0e169984"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daa5d8727a3ef84ecccbbcd2923a3d6b0fe2141ea80e652ae885e800e7c13ce0"></a>

## xcsh_dns_zone_add_cryptokey — xcsh_dns_zone_add_cryptokey / 48bf3bdbb5b6 / 2

Breadcrumbs:

- xcsh_dns_zone_add_cryptokey

Resource creation operation.

<a id="canonical-26d26732f94a682830cfb3eb4a1166184f952771e2083c39d55d03f06d82958e"></a>

## Prerequisites — xcsh_dns_zone_add_cryptokey / 48bf3bdbb5b6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4d6d873a751627f651cf1f8159183d8d1ecf8fbda870368eb0dae6912ca8a770"></a>

## Minimal configuration — xcsh_dns_zone_add_cryptokey / 48bf3bdbb5b6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneAddCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_add_cryptokey" "example" {
  config {
  }
}
```

<a id="canonical-c4a3c67f7e18a8a775009b41aff65a15a02da80eb2b496c73030e8a61f7f1c09"></a>

## Root configuration — xcsh_dns_zone_add_cryptokey / 48bf3bdbb5b6 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-ce4d7f26fb5a383935ff61f1dabdfc7c2884644cce1125383c3fea8f31da05ef"></a>

## Next pages — xcsh_dns_zone_add_cryptokey / 48bf3bdbb5b6 / 6

- [Property reference](../guides/actions--dns_zone_add_cryptokey--reference--group-001.md#canonical-a2054223eb59702f779639edabed8972fc3dc1170f0410220f43b85cfc6e4f06)
- [Examples](../guides/actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-ffb0ff51487ad67e75b86e76da97595e09cc5e24c235163372580e78b9891690)
- [Lifecycle](../guides/actions--dns_zone_add_cryptokey--lifecycle--group-001.md#canonical-f5c8b9810b0f539d2e5ce59f6b29afae629e46b298c8be3ea29bf2b76ae3b2a3)
