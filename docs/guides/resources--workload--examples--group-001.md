---
page_title: "xcsh_workload examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload examples."
---

# xcsh_workload examples

<a id="canonical-e109bf1c1cc0b1f5fac1e13f4e816b72d88664d03443751d6c8ad0dc16247a38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3019e740004c82b5ac065eb30582433981876ef91d6ac953e9b9a284446d29b6"></a>

## Examples — Examples / d1ec9090362c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- Examples

<a id="canonical-40a79ec09924fd227938d464cbc03d437b8f6644ec2d27d446b00a6ae4d4cdcd"></a>

## Complete configurations — Examples / d1ec9090362c / 3

- [Resource](resources--workload--examples--group-001.md#canonical-e2c483db40d48fcba0f05dc4e5f3dd9aed48afaf0bdf82fbb9f81a5e01785ef8): valid configuration.

<a id="canonical-01db094f9358269f004ea6555cd45740d7404311da7e4bb7a281e5f6678d3217"></a>

## Next pages — Examples / d1ec9090362c / 4

- [Resource](resources--workload--examples--group-001.md#canonical-e2c483db40d48fcba0f05dc4e5f3dd9aed48afaf0bdf82fbb9f81a5e01785ef8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e2c483db40d48fcba0f05dc4e5f3dd9aed48afaf0bdf82fbb9f81a5e01785ef8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c95366ffe77847a295ea327a5dfb4170ff3f29b4b8975893c8c7c9e396ac137"></a>

## Resource — Resource / ee1a05ea0f4c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Examples](resources--workload--examples--group-001.md#canonical-e109bf1c1cc0b1f5fac1e13f4e816b72d88664d03443751d6c8ad0dc16247a38)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_workload/resource.tf`; digest `sha256:95bc586cf100a6ef22147aba592992dc22645f4b290f9226bf45165f3414e58c`.

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

<a id="canonical-22f6187fae5a357917a278ce0ea27cad8cbaa8b42e620a334504c682b47e217b"></a>

## Next pages — Resource / ee1a05ea0f4c / 3

- [Examples](resources--workload--examples--group-001.md#canonical-e109bf1c1cc0b1f5fac1e13f4e816b72d88664d03443751d6c8ad0dc16247a38)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
