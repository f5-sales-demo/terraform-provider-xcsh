---
page_title: "xcsh_virtual_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site examples."
---

# xcsh_virtual_site examples

<a id="canonical-8696865a0851712803d39356ca629443120574cfe724f6c3ff2c9b8d9fb89421"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a048df556b426c6b8e283338355bb7fbf3b966bc624d171e6aaa98772b78d20"></a>

## Examples — Examples / f21e4f98ee5b / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)
- Examples

<a id="canonical-ec4ede6f12a8df7d8068e4342d105167291db6dd1b9c9b2b9be718175ea901fc"></a>

## Complete configurations — Examples / f21e4f98ee5b / 3

- [Resource](resources--virtual_site--examples--group-001.md#canonical-26fe3dbcf8c529b0ce4fdd45f90bc20b83a4f551c0c4143a689b7df45133fec0): valid configuration.

<a id="canonical-b2d1580cce2e04b0e83d7c0d0f4d119e12d58a2f99d79d1cbcd1049391698e2e"></a>

## Next pages — Examples / f21e4f98ee5b / 4

- [Resource](resources--virtual_site--examples--group-001.md#canonical-26fe3dbcf8c529b0ce4fdd45f90bc20b83a4f551c0c4143a689b7df45133fec0)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)

<a id="canonical-26fe3dbcf8c529b0ce4fdd45f90bc20b83a4f551c0c4143a689b7df45133fec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b4a925ae1efc4bcdd7626168aba73cac42b7d6f17780b8702764c7c6fe01ec7"></a>

## Resource — Resource / 3ae0cc83d03a / 2

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)
- [Examples](resources--virtual_site--examples--group-001.md#canonical-8696865a0851712803d39356ca629443120574cfe724f6c3ff2c9b8d9fb89421)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_site/resource.tf`; digest `sha256:d4f2ae5a53db544456cc0750c3ea7b5a8e1e23a8aad03279e065c9ce6ac7f872`.

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

<a id="canonical-5876e62e9a709c60990985a9059a518edd277bd4dde883bb85d311baed6ce483"></a>

## Next pages — Resource / 3ae0cc83d03a / 3

- [Examples](resources--virtual_site--examples--group-001.md#canonical-8696865a0851712803d39356ca629443120574cfe724f6c3ff2c9b8d9fb89421)
- [xcsh_virtual_site](../resources/virtual_site.md#canonical-09806f2f14447d14dc58b9c131307be132d72f5eda3348f98adff903bfddcc45)
