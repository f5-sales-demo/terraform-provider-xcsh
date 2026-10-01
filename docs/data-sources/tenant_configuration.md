---
page_title: "xcsh_tenant_configuration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration landing."
---

# xcsh_tenant_configuration landing

<a id="canonical-6c65d7df25dc79a7517bd40ba3201b123141e0bec2b95c1ab2d50ec783eb0521"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67a6b637ad324ae70ca0900eab5fd63c535434133e0eb4488f0670b24854c390"></a>

## xcsh_tenant_configuration — xcsh_tenant_configuration / 5d860fc13fa2 / 2

Breadcrumbs:

- xcsh_tenant_configuration

Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration
specification. configuration.

<a id="canonical-e0e991bc55a1f8502c872526d373f99da1d9c24bde46ab123d295e0742e47b23"></a>

## Prerequisites — xcsh_tenant_configuration / 5d860fc13fa2 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-727c0d3b14319a45f072d251a927529f659afa2549f3cef1897cad7bec26d9b3"></a>

## Minimal configuration — xcsh_tenant_configuration / 5d860fc13fa2 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TenantConfiguration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TenantConfiguration by name
data "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}

output "tenant_configuration_id" {
  value = data.xcsh_tenant_configuration.example.id
}
```

<a id="canonical-9929b48c6d280c4c5da2a1fb909df9561ff65d2c9d9bdc533d467bf255d0378d"></a>

## Root configuration — xcsh_tenant_configuration / 5d860fc13fa2 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-63b9ddc81b0e89df645e57c4f14e272583c5a2abc85f61e27b4cf03ed0462553"></a>

## Next pages — xcsh_tenant_configuration / 5d860fc13fa2 / 6

- [Property reference](../guides/data-sources--tenant_configuration--reference--group-001.md#canonical-da7b7063eb3f3441709e10134a4ab23d3c579ebcbfb00a4c9f032837ac1cb62e)
- [Examples](../guides/data-sources--tenant_configuration--examples--group-001.md#canonical-ef992db0544684039f98ad0b10918386ea3ddd00f9017e4dffa394b0ee6a5569)
