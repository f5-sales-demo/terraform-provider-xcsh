---
page_title: "xcsh_container_registry examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry examples."
---

# xcsh_container_registry examples

<a id="canonical-1e63b61b2b9f3be1491b1d0413dd9d88c1be39585fc6f879c3342853a1ae68f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8e431b3c79ff3c59fc8da346e45fc96aa55eab7a70ff3ccb44315ae6dc15516"></a>

## Examples — Examples / 64eee1404188 / 2

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
- Examples

<a id="canonical-8f9fce83f045657865f2ee2a29989ed4e7c3f62d47f1735e2f1fbf519560ac47"></a>

## Complete configurations — Examples / 64eee1404188 / 3

- [Resource](resources--container_registry--examples--group-001.md#canonical-9cc6db3871b4eb6f0ed48ddac71ba2261c991a6e067bda502e2bbed48f14546e): valid configuration.

<a id="canonical-9b958eeb1580ed1960e2aaf79fb40cbc78955957427ec545a417cb07c1e2d318"></a>

## Next pages — Examples / 64eee1404188 / 4

- [Resource](resources--container_registry--examples--group-001.md#canonical-9cc6db3871b4eb6f0ed48ddac71ba2261c991a6e067bda502e2bbed48f14546e)
- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)

<a id="canonical-9cc6db3871b4eb6f0ed48ddac71ba2261c991a6e067bda502e2bbed48f14546e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14362f6e199f713f7efd3f5bf4be9b05b44389b7677b7404c2b9ba571ce6749c"></a>

## Resource — Resource / 25248237ec77 / 2

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
- [Examples](resources--container_registry--examples--group-001.md#canonical-1e63b61b2b9f3be1491b1d0413dd9d88c1be39585fc6f879c3342853a1ae68f4)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_container_registry/resource.tf`; digest `sha256:333d8100537c966abe2c10d8b88b6ea4b034dc0115a736110297064d1894694d`.

```terraform
# ContainerRegistry Resource Example
# Manages a Container Registry resource in F5 Distributed Cloud for container image registry configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ContainerRegistry configuration
resource "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"

  registry  = "example-value"
  user_name = "example-value"
}
```

<a id="canonical-a0710fe4031f7a20db30b6726752db7aa2480c2ac3f22c23080e97f8b5737593"></a>

## Next pages — Resource / 25248237ec77 / 3

- [Examples](resources--container_registry--examples--group-001.md#canonical-1e63b61b2b9f3be1491b1d0413dd9d88c1be39585fc6f879c3342853a1ae68f4)
- [xcsh_container_registry](../resources/container_registry.md#canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe)
