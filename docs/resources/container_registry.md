---
page_title: "xcsh_container_registry landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry landing."
---

# xcsh_container_registry landing

<a id="canonical-d9354f0149bcec7d60dd61bf8c2faf0289b61a944b3928d453a60e87522be1fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9639f9e7f51de94f94f0ec0f66f70b25fceba00026c613a0c1e7ab9fe08469f0"></a>

## xcsh_container_registry — xcsh_container_registry / d3f97d96b0f2 / 2

Breadcrumbs:

- xcsh_container_registry

Manages a Container Registry resource in F5 Distributed Cloud for container image registry
configuration.

<a id="canonical-6fc0a0d8a103814bd240e66109438b4435a8bb66780d56df5bfff83af7238c73"></a>

## Prerequisites — xcsh_container_registry / d3f97d96b0f2 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-e7c94816c408fc9dad469753f1c55d7dd900ad9425f9e4744791cd89095c47f5"></a>

## Minimal configuration — xcsh_container_registry / d3f97d96b0f2 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-49eb0f54b07df09a7b113de37bd08b3f86074c919f7fd4e3df49aa587029a7c5"></a>

## Root configuration — xcsh_container_registry / d3f97d96b0f2 / 5

Required root properties: `name`, `namespace`, `registry`, `user_name`. Full root flags and choices appear in the property reference.

<a id="canonical-d5b8c00e34e22bacfa55202566000b5b23475075267d900f9f5a21932545e9eb"></a>

## Next pages — xcsh_container_registry / d3f97d96b0f2 / 6

- [Property reference](../guides/resources--container_registry--reference--group-001.md#canonical-4e616e01aa1598e2fc141ba33f7c752da2e7a66e59191d4c102a8adb6557faa3)
- [Examples](../guides/resources--container_registry--examples--group-001.md#canonical-1e63b61b2b9f3be1491b1d0413dd9d88c1be39585fc6f879c3342853a1ae68f4)
- [Import](../guides/resources--container_registry--lifecycle--group-001.md#canonical-6edf4058074763fd900a5411e3721e51635c0bbc1e94e2a670af0074386a1d1a)
- [Timeouts](../guides/resources--container_registry--lifecycle--group-001.md#canonical-b2934a713eac4941a9af207121dadad563b28673bbb4dffd3b32379eb25d1e2c)
