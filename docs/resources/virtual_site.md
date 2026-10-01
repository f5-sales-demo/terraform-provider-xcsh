---
page_title: "xcsh_virtual_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site landing."
---

# xcsh_virtual_site landing

<a id="canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4d4e39dafc68fa268f26dbaf4f77da823ed7e29739cf25ba9798eb929314b31"></a>

## xcsh_virtual_site — xcsh_virtual_site / dbc9df3809b5 / 2

Breadcrumbs:

- xcsh_virtual_site

Manages virtual site object in given namespace in F5 Distributed Cloud.

<a id="canonical-7fb263a0e2140858d197b9c0ba8acb3c47709944bcc560e317b5e4dc6f2ee524"></a>

## Prerequisites — xcsh_virtual_site / dbc9df3809b5 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-ed1613788d419b50b86402a5d22ebc6034dcdb4816b1370f23801d42be47cc38"></a>

## Minimal configuration — xcsh_virtual_site / dbc9df3809b5 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualSite Resource Example
# Manages virtual site object in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualSite configuration
resource "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}
```

<a id="canonical-051a94ce47cd0cba9a21d708af2bf3b38c57f4a73b80f9cee4876d3a860b8611"></a>

## Root configuration — xcsh_virtual_site / dbc9df3809b5 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c0b74d58522c7a7b37f40e351179018bc4f0490ef6649177a2cc36fc2c37465b"></a>

## Next pages — xcsh_virtual_site / dbc9df3809b5 / 6

- [Property reference](../guides/resources--virtual_site--reference--group-001.md#canonical-c6e1c58e8ed5d2e1ac06a6cd0b224843ecf7d102b47216aca5fb83c491698b76)
- [Examples](../guides/resources--virtual_site--examples--group-001.md#canonical-8696865a0851712803d39356ca629443120574cfe724f6c3ff2c9b8d9fb89421)
- [Import](../guides/resources--virtual_site--lifecycle--group-001.md#canonical-a34e00f807eb98bf6a164322104192bd256a48bd0210ddadb3991a7f09112061)
- [Timeouts](../guides/resources--virtual_site--lifecycle--group-001.md#canonical-765adc4a44cab036b871f4206f41e1f04ca1ad0cebe54ce919ca994963a6eaad)
