---
page_title: "xcsh_namespace examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace examples."
---

# xcsh_namespace examples

<a id="canonical-13126340d43a7fef70e3c42c1e8d7b7fb39680618a33f6e79757825b1db49b7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-131a2ab98f57f5a5d477d9656550c48d1e66be02d12b371efef54f4dd1b34936"></a>

## Examples — Examples / b3e83447176c / 2

Breadcrumbs:

- [xcsh_namespace](../data-sources/namespace.md#canonical-87303facfb5a86c45f03fd2edff3e82a1b37ef90f0d3be2e23ec03de74d32096)
- Examples

<a id="canonical-d98382438c552bfb87d54869f94adb28c59cac3888fa7b3f0fe66c91a730e6c9"></a>

## Complete configurations — Examples / b3e83447176c / 3

- [Data source](data-sources--namespace--examples--group-001.md#canonical-15781e0a0bc91c84f808fb90a443783ca159ab17c66976627f17521b86319630): valid configuration.

<a id="canonical-e464c0a3ed7b6f4a6f22ed1c26af7e4c5e4920df9d116e423e346f4fff3bfc22"></a>

## Next pages — Examples / b3e83447176c / 4

- [Data source](data-sources--namespace--examples--group-001.md#canonical-15781e0a0bc91c84f808fb90a443783ca159ab17c66976627f17521b86319630)
- [xcsh_namespace](../data-sources/namespace.md#canonical-87303facfb5a86c45f03fd2edff3e82a1b37ef90f0d3be2e23ec03de74d32096)

<a id="canonical-15781e0a0bc91c84f808fb90a443783ca159ab17c66976627f17521b86319630"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-758183fa011e333abc2c3e2fd037a60563ca8990d5ac9a0f12a1cf0aa76530f8"></a>

## Data source — Data source / 38df7dabe7bb / 2

Breadcrumbs:

- [xcsh_namespace](../data-sources/namespace.md#canonical-87303facfb5a86c45f03fd2edff3e82a1b37ef90f0d3be2e23ec03de74d32096)
- [Examples](data-sources--namespace--examples--group-001.md#canonical-13126340d43a7fef70e3c42c1e8d7b7fb39680618a33f6e79757825b1db49b7e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_namespace/data-source.tf`; digest `sha256:a2418f75dbce2847a8b4c2579a5eb124cd8b027cba5de6019e63eeb934c35492`.

```terraform
# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name = "example-namespace"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```

<a id="canonical-0d86ed32b571764284d0ddd983e4923b1ef8366d323627424a1f58f2df47d375"></a>

## Next pages — Data source / 38df7dabe7bb / 3

- [Examples](data-sources--namespace--examples--group-001.md#canonical-13126340d43a7fef70e3c42c1e8d7b7fb39680618a33f6e79757825b1db49b7e)
- [xcsh_namespace](../data-sources/namespace.md#canonical-87303facfb5a86c45f03fd2edff3e82a1b37ef90f0d3be2e23ec03de74d32096)
