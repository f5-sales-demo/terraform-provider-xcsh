---
page_title: "xcsh_workload landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload landing."
---

# xcsh_workload landing

<a id="canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14d835a00c21949e901333e80f73d52d073557ecce3652d0ec572d8ccca10c2b"></a>

## xcsh_workload — xcsh_workload / 9cdf9d17a30c / 2

Breadcrumbs:

- xcsh_workload

Manages a Workload resource in F5 Distributed Cloud for workload. configuration.

<a id="canonical-003c2fab0857e39d73beaad873824fa7600d8bb7f9d892e8dfe564aed44263bd"></a>

## Prerequisites — xcsh_workload / 9cdf9d17a30c / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_k8s`.

- virtual_k8s: Namespace for workload deployment

<a id="canonical-9e41de33f47195981bb801f87e7628fc2e843889ab7d9f0884e88a6f11f2ab37"></a>

## Minimal configuration — xcsh_workload / 9cdf9d17a30c / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Workload Resource Example
# Manages a Workload resource in F5 Distributed Cloud for workload.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Workload configuration
resource "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}
```

<a id="canonical-7c6af6459ca603a2562e3d7c10c48e927baff8ab3020a37a91f58021439defe4"></a>

## Root configuration — xcsh_workload / 9cdf9d17a30c / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-7d559c8d3845039aa9d3da21ac7439a86ad0821984f42c8d15bf5e25a593618d"></a>

## Next pages — xcsh_workload / 9cdf9d17a30c / 6

- [Property reference](../guides/resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [Examples](../guides/resources--workload--examples--group-001.md#canonical-e109bf1c1cc0b1f5fac1e13f4e816b72d88664d03443751d6c8ad0dc16247a38)
- [Import](../guides/resources--workload--lifecycle--group-001.md#canonical-e5ab166e489361432d4b9f5c5481c604ddc241f1bdd57148e55013d8894626ec)
- [Timeouts](../guides/resources--workload--lifecycle--group-001.md#canonical-80d5ca92cc9c7af92dcc212b52205c6bb1e045840dc750d97c021cccb1425a95)
