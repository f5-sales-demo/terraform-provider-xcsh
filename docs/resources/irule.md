---
page_title: "xcsh_irule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule landing."
---

# xcsh_irule landing

<a id="canonical-85fd6f4d086f69670bdf77b16e6524c8e8a0dcea34c8d6b1e850fef55d090d23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8100320f02924942da17deb0f9fe3227f9b609f52d9d209725720dca5a10565"></a>

## xcsh_irule — xcsh_irule / 9b5a535267cd / 2

Breadcrumbs:

- xcsh_irule

Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed
Cloud.

<a id="canonical-bec295b99ae3f223b5f03378937f79fb8b27c44682cf06fcbf3271e773084008"></a>

## Prerequisites — xcsh_irule / 9b5a535267cd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b8b78d2742ac8fed52ffa25c1765b2e1c44d93ff46526af0a4822337c11fc46d"></a>

## Minimal configuration — xcsh_irule / 9b5a535267cd / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```

<a id="canonical-17b7b2cae06523e89858d3bb1674cd94fb3a30545d1c8ad2b7f6e627e32a5808"></a>

## Root configuration — xcsh_irule / 9b5a535267cd / 5

Required root properties: `description_spec`, `irule`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8b3efaf2a817bbda2637b0a9d78d92035ee3163d10a775cc9df30b42327852bc"></a>

## Next pages — xcsh_irule / 9b5a535267cd / 6

- [Property reference](../guides/resources--irule--reference--group-001.md#canonical-4a843a98d007e27057ed2207e6455121fd9d1ff3073d3ba7dce67a1185a2a905)
- [Examples](../guides/resources--irule--examples--group-001.md#canonical-127fa3038056d5fabadf3cc4ce5a2f01422cfdae2d9180f183fce32902d60d6d)
- [Import](../guides/resources--irule--lifecycle--group-001.md#canonical-9e55675d1c73fe8d48b8b4c0dbe14c9f6ca5eade8d4c419d659321e224c9c641)
- [Timeouts](../guides/resources--irule--lifecycle--group-001.md#canonical-8fd54bc91c72ba6ab724dd2eb3d94f61e851723040d3bef208d7273201d1b556)
