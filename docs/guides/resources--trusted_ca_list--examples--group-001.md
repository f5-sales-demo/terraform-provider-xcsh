---
page_title: "xcsh_trusted_ca_list examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list examples."
---

# xcsh_trusted_ca_list examples

<a id="canonical-771aef14c4ccf7228b0927677580f5f1cc177af204221a52d4d906931f570b28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ebfd07a579f314f85c099f19570ed96222707cbae59cf091f6b40701bcb1573"></a>

## Examples — Examples / cd2e91860f50 / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)
- Examples

<a id="canonical-8f486d385bcfc270bb5c830d810ea14c49526b2824897346f61dadd8e7ee3ad2"></a>

## Complete configurations — Examples / cd2e91860f50 / 3

- [Resource](resources--trusted_ca_list--examples--group-001.md#canonical-9326753b6e6c94111054b8cd98c9c6b7473415a7169cff951a3748a1019fbcb1): valid configuration.

<a id="canonical-f0028d3176859d98ffc6b7e04a1e37ce7fcdb5a923de3ad5449223b70004a285"></a>

## Next pages — Examples / cd2e91860f50 / 4

- [Resource](resources--trusted_ca_list--examples--group-001.md#canonical-9326753b6e6c94111054b8cd98c9c6b7473415a7169cff951a3748a1019fbcb1)
- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)

<a id="canonical-9326753b6e6c94111054b8cd98c9c6b7473415a7169cff951a3748a1019fbcb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d87e834307c7ae33d200ca12087185a23fdac18d2245feb33199cb01bb459df"></a>

## Resource — Resource / cd8a56df84e5 / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)
- [Examples](resources--trusted_ca_list--examples--group-001.md#canonical-771aef14c4ccf7228b0927677580f5f1cc177af204221a52d4d906931f570b28)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_trusted_ca_list/resource.tf`; digest `sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328`.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```

<a id="canonical-559c82658abc8aaa7deb6cbf053f3c773587bb27a1110bbcb481a803f750697a"></a>

## Next pages — Resource / cd8a56df84e5 / 3

- [Examples](resources--trusted_ca_list--examples--group-001.md#canonical-771aef14c4ccf7228b0927677580f5f1cc177af204221a52d4d906931f570b28)
- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)
