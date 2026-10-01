---
page_title: "xcsh_tenant_configuration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration landing."
---

# xcsh_tenant_configuration landing

<a id="canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5896d990233be3d4a335e97b5b4ae946cc5648070cf6ca4a37aef80601a3eb70"></a>

## xcsh_tenant_configuration — xcsh_tenant_configuration / 79f79e865a82 / 2

Breadcrumbs:

- xcsh_tenant_configuration

Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration
specification. configuration.

<a id="canonical-6cd5c338cb7081ee94234e9f1a9d05aab72a081634ca4376d271bef171bd9b16"></a>

## Prerequisites — xcsh_tenant_configuration / 79f79e865a82 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-631f58bc374b3c66b8496d31a9441100523aebf1e4d1369b76bd2fbfb00fc93b"></a>

## Minimal configuration — xcsh_tenant_configuration / 79f79e865a82 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TenantConfiguration Resource Example
# Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TenantConfiguration configuration
resource "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}
```

<a id="canonical-b2e12a7deb56a8cfe7cd2f07bdd6d855f9bd3765a974f4a1b15c9e20d4610c01"></a>

## Root configuration — xcsh_tenant_configuration / 79f79e865a82 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-508aac79a9a5d0dcaaec310d4116eac538062c3c7b13e1e7af69f3cc02da57ae"></a>

## Next pages — xcsh_tenant_configuration / 79f79e865a82 / 6

- [Property reference](../guides/resources--tenant_configuration--reference--group-001.md#canonical-16136964bc77142e101e918b10e24c07bf5c7f99c8c5df90b4efd69072c8b3a2)
- [Examples](../guides/resources--tenant_configuration--examples--group-001.md#canonical-f14acc1257a6c1413786a908104e7dc1d113d5ed4b44ad0345eaf33abc2cf8e7)
- [Import](../guides/resources--tenant_configuration--lifecycle--group-001.md#canonical-95ab19143a5fabf1282cb51237fdf4fc9b81960a07822c4f81d2ed7eba9a9d9e)
- [Timeouts](../guides/resources--tenant_configuration--lifecycle--group-001.md#canonical-c08d92a93ad94f5ce9296e1462258cc516c1bab298feecb540ac6ae0f1880bcb)
