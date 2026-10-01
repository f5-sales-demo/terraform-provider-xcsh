---
page_title: "xcsh_container_registry landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry landing."
---

# xcsh_container_registry landing

<a id="canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-610abdbd7ba5b544723e4741c36f322897cf2eefd50ccc819234d69fa72c3968"></a>

## xcsh_container_registry — xcsh_container_registry / 3599c75f8a2f / 2

Breadcrumbs:

- xcsh_container_registry

Manages a Container Registry resource in F5 Distributed Cloud for container image registry
configuration.

<a id="canonical-4b6e867be66dc87d02ed4b1a2c2fa147656cd67d851cac80496cbd3ec9f1c4bd"></a>

## Prerequisites — xcsh_container_registry / 3599c75f8a2f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-dac9e9c99a6cddfa8fe2b936a1410e401dcf976c963cfe82e9d38a51dd0c5b56"></a>

## Minimal configuration — xcsh_container_registry / 3599c75f8a2f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ContainerRegistry Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ContainerRegistry by name
data "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"
}

output "container_registry_id" {
  value = data.xcsh_container_registry.example.id
}
```

<a id="canonical-faa2ad6a34b6b020f710b2d529c9b0dd0b0488dc283b6645815f734c66a9e8ca"></a>

## Root configuration — xcsh_container_registry / 3599c75f8a2f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-33a9972a96ed6626852a124a0938b56f03556708ed3590665fccd371c3a166fd"></a>

## Next pages — xcsh_container_registry / 3599c75f8a2f / 6

- [Property reference](../guides/data-sources--container_registry--reference--group-001.md#canonical-e58992eaf4724a8cf0c7f0fa641dc1f519eedb53d24397f6ae8cf4129ef62c19)
- [Examples](../guides/data-sources--container_registry--examples--group-001.md#canonical-b02af979eb8bc9023d2ac42c427813e22c4f214bd2b8b65871fb12f69dd19488)
