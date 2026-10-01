---
page_title: "xcsh_dns_zone_add_cryptokey examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_add_cryptokey examples."
---

# xcsh_dns_zone_add_cryptokey examples

<a id="canonical-ffb0ff51487ad67e75b86e76da97595e09cc5e24c235163372580e78b9891690"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94af553aad0867655f625b3de1a6839342c4df44463882743c9ad15a09bbd6f8"></a>

## Examples — Examples / f6dd3d77045d / 2

Breadcrumbs:

- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-85ff7b5ce0689a3d0497afdd728eedfead55ce6ef5288db5ba75c83a0e169984)
- Examples

<a id="canonical-1a173615a5c31b14b7965238035bb6c9e8b134735471eb3eb77a8efc0a8139fe"></a>

## Complete configurations — Examples / f6dd3d77045d / 3

- [Action](actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-79536bf7a218f03458663597a6a312db09e24ce61b667eba8fb73937b0d9ab0a): valid configuration.

<a id="canonical-e19a2600be1fc1818920d6a8579519a5d2327eba7e574ba53c6bedac7d6dbb5d"></a>

## Next pages — Examples / f6dd3d77045d / 4

- [Action](actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-79536bf7a218f03458663597a6a312db09e24ce61b667eba8fb73937b0d9ab0a)
- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-85ff7b5ce0689a3d0497afdd728eedfead55ce6ef5288db5ba75c83a0e169984)

<a id="canonical-79536bf7a218f03458663597a6a312db09e24ce61b667eba8fb73937b0d9ab0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3b49467193a2e85309ba16074ed63162410de9ca7b30e59fe802272a6e946c8"></a>

## Action — Action / 2450aaa73313 / 2

Breadcrumbs:

- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-85ff7b5ce0689a3d0497afdd728eedfead55ce6ef5288db5ba75c83a0e169984)
- [Examples](actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-ffb0ff51487ad67e75b86e76da97595e09cc5e24c235163372580e78b9891690)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_dns_zone_add_cryptokey/action.tf`; digest `sha256:41e2e4052bac9b33dbdf74fdf0a795bf2458f8e723d6c49a082ace67a79ec803`.

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

<a id="canonical-89a078823d86c8853004326bdee28a7cab4b416d8966bda6551c5e62002a66e2"></a>

## Next pages — Action / 2450aaa73313 / 3

- [Examples](actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-ffb0ff51487ad67e75b86e76da97595e09cc5e24c235163372580e78b9891690)
- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-85ff7b5ce0689a3d0497afdd728eedfead55ce6ef5288db5ba75c83a0e169984)
