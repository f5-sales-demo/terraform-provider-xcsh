---
page_title: "xcsh_dns_lb_pool examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool examples."
---

# xcsh_dns_lb_pool examples

<a id="canonical-c9d179e71810ad3d046c54203f75de777a65f82b4e362f3fcb56fc10b8a6fcc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-340c335ae230a9bf07beb1f6caea7a7c7f5761a519426b6d97e009551a57415b"></a>

## Examples — Examples / 18a17ad7c754 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- Examples

<a id="canonical-6b8d88a6f9163175f67f9d27c647b7e88fbee36b3be5fd16472f1c4466728a4b"></a>

## Complete configurations — Examples / 18a17ad7c754 / 3

- [Resource](resources--dns_lb_pool--examples--group-001.md#canonical-1d36ec19a6b8b9629a2b18233c7e6200f45866c3aebe1eb16855fb40193d843e): valid configuration.

<a id="canonical-04238722768eb5601796ca582f8695a12abf309f30f6ce8629a5e448b905a90e"></a>

## Next pages — Examples / 18a17ad7c754 / 4

- [Resource](resources--dns_lb_pool--examples--group-001.md#canonical-1d36ec19a6b8b9629a2b18233c7e6200f45866c3aebe1eb16855fb40193d843e)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-1d36ec19a6b8b9629a2b18233c7e6200f45866c3aebe1eb16855fb40193d843e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37995dbab87db90ab16d91d12fd6513de745d94f82c2b04a13aaeeb7c1b7326a"></a>

## Resource — Resource / c42b0873a0a1 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Examples](resources--dns_lb_pool--examples--group-001.md#canonical-c9d179e71810ad3d046c54203f75de777a65f82b4e362f3fcb56fc10b8a6fcc0)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_lb_pool/resource.tf`; digest `sha256:a0f8bbd985a30a9aa4d49bcfaf3107e9c25c24df2c7857f8b4c1d06b0812d763`.

```terraform
# DNSLBPool Resource Example
# Manages DNS Load Balancer Pool in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBPool configuration
resource "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}
```

<a id="canonical-b0848576c8ce8786feb74a07edd51ec0784fbbdb8d959b6832c566176500a36b"></a>

## Next pages — Resource / c42b0873a0a1 / 3

- [Examples](resources--dns_lb_pool--examples--group-001.md#canonical-c9d179e71810ad3d046c54203f75de777a65f82b4e362f3fcb56fc10b8a6fcc0)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
