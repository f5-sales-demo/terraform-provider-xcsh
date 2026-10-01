---
page_title: "xcsh_virtual_host examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host examples."
---

# xcsh_virtual_host examples

<a id="canonical-6478e4ef8c92ee0e5eeebc8e11c97ab12ee08c898ee8603da8ef3f11720a52eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e871b2b0a6f61229754a752e26909dec19cdbff350b01bef05dc48104bcd1da"></a>

## Examples — Examples / 30e0143455b5 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- Examples

<a id="canonical-41fbf9f3efabe737c23eca4d80b118e27593f67cc7b0e2f8a7fb76d25fac2b47"></a>

## Complete configurations — Examples / 30e0143455b5 / 3

- [Resource](resources--virtual_host--examples--group-001.md#canonical-0991096c2829adbf758fad4a791e72f44a6f3e48ff65ddead3139c2ffb0a01c5): valid configuration.

<a id="canonical-79b09252759a0b8319e518916b586951a6af1d6040c667505bb74ea286ced527"></a>

## Next pages — Examples / 30e0143455b5 / 4

- [Resource](resources--virtual_host--examples--group-001.md#canonical-0991096c2829adbf758fad4a791e72f44a6f3e48ff65ddead3139c2ffb0a01c5)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)

<a id="canonical-0991096c2829adbf758fad4a791e72f44a6f3e48ff65ddead3139c2ffb0a01c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31452bcfd95962441afc132869e4e84e8ce6eeaf2257ccb88c43e81f9cb080f0"></a>

## Resource — Resource / 7d9fa8d91644 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
- [Examples](resources--virtual_host--examples--group-001.md#canonical-6478e4ef8c92ee0e5eeebc8e11c97ab12ee08c898ee8603da8ef3f11720a52eb)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_host/resource.tf`; digest `sha256:7132f3f3fb88713f102679821dabd8f6cf4e11c9765e5f1be76eb1b01ecae4a0`.

```terraform
# VirtualHost Resource Example
# Manages virtual host in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualHost configuration
resource "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}
```

<a id="canonical-dde61aa11aba0fe91353e0cdc60c5ea2a8dca2a1b2e86b2aa583eeb23b3328bd"></a>

## Next pages — Resource / 7d9fa8d91644 / 3

- [Examples](resources--virtual_host--examples--group-001.md#canonical-6478e4ef8c92ee0e5eeebc8e11c97ab12ee08c898ee8603da8ef3f11720a52eb)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-eb75e57432f67b77c490109880d93dfba9aae6a7a35dd3e7936370ecd6ba3e2a)
