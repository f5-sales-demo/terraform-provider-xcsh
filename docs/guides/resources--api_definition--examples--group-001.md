---
page_title: "xcsh_api_definition examples"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition examples."
---

# xcsh_api_definition examples

<a id="canonical-49827350fd565d874def3c4d3aa62c89a8251685c5428c58bd74db2481e14124"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b81f2d85b1865bad77fc0eb5ff89afcbd8ab0387f3acd29f0041e0b53b88804"></a>

## Examples — Examples / 5c32456f687f / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- Examples

<a id="canonical-9ee2a926c51c3db58c59eb4dffddc3c09391e7b721e4647310127e5f798c72a1"></a>

## Complete configurations — Examples / 5c32456f687f / 3

- [Resource](resources--api_definition--examples--group-001.md#canonical-670567b3ff1594d0e4beafcb3dc24e336c0f673a90390165a13e9e9185689130): valid configuration.

<a id="canonical-88530e4064de6bf8fe596735d2e92f4fe8f71f6f2be4f0cde545754ae13adc63"></a>

## Next pages — Examples / 5c32456f687f / 4

- [Resource](resources--api_definition--examples--group-001.md#canonical-670567b3ff1594d0e4beafcb3dc24e336c0f673a90390165a13e9e9185689130)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)

<a id="canonical-670567b3ff1594d0e4beafcb3dc24e336c0f673a90390165a13e9e9185689130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d74baabda443ab60c1c7cd202011da3086c185e1e857dcea8da02093520b27a1"></a>

## Resource — Resource / 70e81a0d16f5 / 2

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
- [Examples](resources--api_definition--examples--group-001.md#canonical-49827350fd565d874def3c4d3aa62c89a8251685c5428c58bd74db2481e14124)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_definition/resource.tf`; digest `sha256:8eaaa98845c6177dba9c28e17da78fe9cbe8a7d10b85b7c4e55e7857d3d8ba64`.

```terraform
# APIDefinition Resource Example
# Manages API Definition in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDefinition configuration
resource "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}
```

<a id="canonical-d7a07012f211017c410768f285a668126c28c3d92b6482cd3a2aa902f5e6b92b"></a>

## Next pages — Resource / 70e81a0d16f5 / 3

- [Examples](resources--api_definition--examples--group-001.md#canonical-49827350fd565d874def3c4d3aa62c89a8251685c5428c58bd74db2481e14124)
- [xcsh_api_definition](../resources/api_definition.md#canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0)
