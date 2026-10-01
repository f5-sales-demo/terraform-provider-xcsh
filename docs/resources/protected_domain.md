---
page_title: "xcsh_protected_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain landing."
---

# xcsh_protected_domain landing

<a id="canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7b995b38b93e84084834dc4e30ab98e0cb2a66420446d10daa54916abb59030"></a>

## xcsh_protected_domain — xcsh_protected_domain / 68ae6a288a8c / 2

Breadcrumbs:

- xcsh_protected_domain

Manages Domain to protect in F5 Distributed Cloud.

<a id="canonical-d9cd32491889ea3f1245dba97f9a6f94bebff183e78bcf6b8db5f458b9817367"></a>

## Prerequisites — xcsh_protected_domain / 68ae6a288a8c / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-abd9e43122f4e102ffdad1e558906775a9b5e5a4acb2792d9f41ce09d079d3de"></a>

## Minimal configuration — xcsh_protected_domain / 68ae6a288a8c / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedDomain Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedDomain configuration
resource "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"

  protected_domain = "example.com"
}
```

<a id="canonical-2034290c42d0c850dbc6d6f9358bb900fc211699f2e8c7ba225f458855233bed"></a>

## Root configuration — xcsh_protected_domain / 68ae6a288a8c / 5

Required root properties: `name`, `namespace`, `protected_domain`. Full root flags and choices appear in the property reference.

<a id="canonical-6c021675553c0de37eaa1519bac1a02b4bfb7390cd37df349f3ae3dc48b13f22"></a>

## Next pages — xcsh_protected_domain / 68ae6a288a8c / 6

- [Property reference](../guides/resources--protected_domain--reference--group-001.md#canonical-0d8b0bcd910ff49e6d0333184800f1f4d0ef950cb7ec9b09d4025255720e776c)
- [Examples](../guides/resources--protected_domain--examples--group-001.md#canonical-4da6f421b79056c24836875bfa0f481d455793d2790413cc6063d3ce821bf0ed)
- [Import](../guides/resources--protected_domain--lifecycle--group-001.md#canonical-a5b9c6d3a481a8b89710321525ab169b52efd6cff136cb48e6934e7fcf2e55e5)
- [Timeouts](../guides/resources--protected_domain--lifecycle--group-001.md#canonical-7012095e39225da01bf24c1222b711ed69cf59f21c9fa8c1d5bf9b1f369b9726)
