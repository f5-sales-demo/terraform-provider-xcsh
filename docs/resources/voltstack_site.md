---
page_title: "xcsh_voltstack_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site landing."
---

# xcsh_voltstack_site landing

<a id="canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb07aeb08e9a95a21efbfadc53aa70ee8d179e50c9ce7c12338dc4d92c50a44a"></a>

## xcsh_voltstack_site — xcsh_voltstack_site / 78574b351342 / 2

Breadcrumbs:

- xcsh_voltstack_site

Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing
sites.

<a id="canonical-5b1aff9a8ddb302d87fb6ef17b63c322bffeac26a91c16904bab8d6d1e4e7fa3"></a>

## Prerequisites — xcsh_voltstack_site / 78574b351342 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d600c301e97a762538f157d95329a47dc6cbd5baeca1b3d43bb947859494d073"></a>

## Minimal configuration — xcsh_voltstack_site / 78574b351342 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VoltstackSite Resource Example
# Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VoltstackSite configuration
resource "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

<a id="canonical-45f96e52d4d5430fdf5d868bd1b52ad3f5ae069132fae3256cf38badde456351"></a>

## Root configuration — xcsh_voltstack_site / 78574b351342 / 5

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

<a id="canonical-4473f81c3e41f3c759f6f3f9d14586b6ea526adc72f4fd6153fd15d765f42a7d"></a>

## Next pages — xcsh_voltstack_site / 78574b351342 / 6

- [Property reference](../guides/resources--voltstack_site--reference--group-001.md#canonical-b3ddccc079feed8c94a9f862590f39f62918f9d739fc6059fb4d79a2465b6a15)
- [Examples](../guides/resources--voltstack_site--examples--group-001.md#canonical-80622daaa5ef0ae14ed426d2603f5d807d6b2b77ce118a081f8210d154ec653a)
- [Import](../guides/resources--voltstack_site--lifecycle--group-001.md#canonical-01c6470fb7a188a21ae90227fad45dda7ded398b1c46fdc05c947b798e2370fb)
- [Timeouts](../guides/resources--voltstack_site--lifecycle--group-001.md#canonical-c1b6ad83b6142b659efd3a65007bbdf9fe03083560168549636e34a84f9ede15)
