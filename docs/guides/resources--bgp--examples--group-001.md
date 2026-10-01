---
page_title: "xcsh_bgp examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp examples."
---

# xcsh_bgp examples

<a id="canonical-0b5fe495434f4ae93528ca40a2c382856380185f9580f1f16c5b072a7d6a2bca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-925007a7550c38609329df72f35ac04f67f94535476d321482fe51f5ad593bc3"></a>

## Examples — Examples / 4c05f89844ab / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- Examples

<a id="canonical-578fab2010dd2e443eb10f0679a6f02d496f02317de5754ae402b4ddf8995b6a"></a>

## Complete configurations — Examples / 4c05f89844ab / 3

- [Resource](resources--bgp--examples--group-001.md#canonical-b4905ed167122f81305f27b28737a7d3fc9b4c01dae9987e6381dfd77780d06b): valid configuration.

<a id="canonical-332db49ff6d0b3949026c46db4f7f081942f95bc6ea9b3b1907e6e7871ace491"></a>

## Next pages — Examples / 4c05f89844ab / 4

- [Resource](resources--bgp--examples--group-001.md#canonical-b4905ed167122f81305f27b28737a7d3fc9b4c01dae9987e6381dfd77780d06b)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)

<a id="canonical-b4905ed167122f81305f27b28737a7d3fc9b4c01dae9987e6381dfd77780d06b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f5b1d81168d73b8e9e5708a3a561f50da3bfc307e666209872607bfc4e25956"></a>

## Resource — Resource / 28654b376a7c / 2

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
- [Examples](resources--bgp--examples--group-001.md#canonical-0b5fe495434f4ae93528ca40a2c382856380185f9580f1f16c5b072a7d6a2bca)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp/resource.tf`; digest `sha256:7d85805409033bf09b03057cbf852697d9f9c443321c552ee56fb1009994fdd5`.

```terraform
# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGP configuration
resource "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}
```

<a id="canonical-c1b7fb8044347b2b5b0c0b1eb6d5f76856415d1ac59e3f853e2d39f6b5656b95"></a>

## Next pages — Resource / 28654b376a7c / 3

- [Examples](resources--bgp--examples--group-001.md#canonical-0b5fe495434f4ae93528ca40a2c382856380185f9580f1f16c5b072a7d6a2bca)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
