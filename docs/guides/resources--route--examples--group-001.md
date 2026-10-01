---
page_title: "xcsh_route examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route examples."
---

# xcsh_route examples

<a id="canonical-6aded232f88506504f0ff82eecce407706fca028a2976781dc43e66282e9d3c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e8fb6ee8571d08df67f42bbaceaae07809462156b3d348efee47c658f16cd13"></a>

## Examples — Examples / 13978b24aa51 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- Examples

<a id="canonical-2fb590242900484b7a93cfb23fd8cff9f5735decad1cfd7536ae0495fb35ffbf"></a>

## Complete configurations — Examples / 13978b24aa51 / 3

- [Resource](resources--route--examples--group-001.md#canonical-5fbfd5e7ddf017e27a0cf867c3f752dacd44aa5249946ff4cc3d5323a8de61ec): valid configuration.

<a id="canonical-c991f34a7dae5a90d8cdc74350a9b30737d519634443807dfdd0f39c1fd849d9"></a>

## Next pages — Examples / 13978b24aa51 / 4

- [Resource](resources--route--examples--group-001.md#canonical-5fbfd5e7ddf017e27a0cf867c3f752dacd44aa5249946ff4cc3d5323a8de61ec)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)

<a id="canonical-5fbfd5e7ddf017e27a0cf867c3f752dacd44aa5249946ff4cc3d5323a8de61ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a026debbe96f75d3d7a43868a0d3649da5b2e9c32c948638878b08a4cbf2cbaf"></a>

## Resource — Resource / 6b0c170a8c40 / 2

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
- [Examples](resources--route--examples--group-001.md#canonical-6aded232f88506504f0ff82eecce407706fca028a2976781dc43e66282e9d3c3)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_route/resource.tf`; digest `sha256:02c199de787def8c45388d6ef15d06881bc08c78cf9b74e0ee43c7ac9db177f3`.

```terraform
# Route Resource Example
# Manages route object in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Route configuration
resource "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}
```

<a id="canonical-2ce616bc3584b86f999cbe474190a64c385dc9240188f5a25dc60d192aec9653"></a>

## Next pages — Resource / 6b0c170a8c40 / 3

- [Examples](resources--route--examples--group-001.md#canonical-6aded232f88506504f0ff82eecce407706fca028a2976781dc43e66282e9d3c3)
- [xcsh_route](../resources/route.md#canonical-8e20ffe65a1c536995f3a9b20e795999cb3ba46e6df8a20b897c10159c22e936)
