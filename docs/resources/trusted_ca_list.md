---
page_title: "xcsh_trusted_ca_list landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list landing."
---

# xcsh_trusted_ca_list landing

<a id="canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a89009d5893b0fe46e714cf6bc2f35aa386b15efee448683ea0a7981a9164210"></a>

## xcsh_trusted_ca_list — xcsh_trusted_ca_list / b1872cfe6e8d / 2

Breadcrumbs:

- xcsh_trusted_ca_list

Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list
management.

<a id="canonical-97109eb1ca4322695f60e476f80426d6c314c8e6157f906c8b4d109293857b74"></a>

## Prerequisites — xcsh_trusted_ca_list / b1872cfe6e8d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-749fceabfb55e6d5750454465860f4d2f4f27e6ecde0bf51d04a80af5cf6c219"></a>

## Minimal configuration — xcsh_trusted_ca_list / b1872cfe6e8d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-b28cdf85409235f9322f750a67cda8fb573011c9a0b1ac089aba24297ce07ae6"></a>

## Root configuration — xcsh_trusted_ca_list / b1872cfe6e8d / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-269b3d81e406ac33594d5292531e3f9062eea04eb8cb5b30202cb8860a3fb942"></a>

## Next pages — xcsh_trusted_ca_list / b1872cfe6e8d / 6

- [Property reference](../guides/resources--trusted_ca_list--reference--group-001.md#canonical-85409a678c309b8757816905fb4db6f52278501f089928c98decb8ccad0ed016)
- [Examples](../guides/resources--trusted_ca_list--examples--group-001.md#canonical-771aef14c4ccf7228b0927677580f5f1cc177af204221a52d4d906931f570b28)
- [Import](../guides/resources--trusted_ca_list--lifecycle--group-001.md#canonical-90349cffb6492bab5ca64e2f0f0c9b5d553d96785a772040d9f97aa2ea9a5357)
- [Timeouts](../guides/resources--trusted_ca_list--lifecycle--group-001.md#canonical-9f9488b0edc2f220cbd73e586d1d0ad1ba4ed01d3c29da97d88ca47ee42a5770)
