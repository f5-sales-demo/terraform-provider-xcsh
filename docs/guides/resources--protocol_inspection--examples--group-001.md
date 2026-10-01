---
page_title: "xcsh_protocol_inspection examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection examples."
---

# xcsh_protocol_inspection examples

<a id="canonical-daa07a474b657dbd62339295ce1f9bb6e4e70833d976326d48460822d296ef4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33484113c1d3def91b9b34607c235ac3df4e5ccdbe8a1d4ba96b7c392200f12e"></a>

## Examples — Examples / 9e2d5c65d920 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- Examples

<a id="canonical-f7cef283d530996602b3700ef9016b70ac419e2aaa7a4be1d24fdb1c9888582e"></a>

## Complete configurations — Examples / 9e2d5c65d920 / 3

- [Resource](resources--protocol_inspection--examples--group-001.md#canonical-1d853d91b28ca24d1611f77a2a59dfb29acca4a9df2159e9c0412b71dd985d10): valid configuration.

<a id="canonical-909ff3aebd2cffe328f822ee3517393edf4617cbc9db314ac2d599545947a2f0"></a>

## Next pages — Examples / 9e2d5c65d920 / 4

- [Resource](resources--protocol_inspection--examples--group-001.md#canonical-1d853d91b28ca24d1611f77a2a59dfb29acca4a9df2159e9c0412b71dd985d10)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)

<a id="canonical-1d853d91b28ca24d1611f77a2a59dfb29acca4a9df2159e9c0412b71dd985d10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c98a429d87e9b772248b072e0f665780be0c3c9de0a067757725145c0461db5"></a>

## Resource — Resource / f9fc21f3c32d / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
- [Examples](resources--protocol_inspection--examples--group-001.md#canonical-daa07a474b657dbd62339295ce1f9bb6e4e70833d976326d48460822d296ef4f)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protocol_inspection/resource.tf`; digest `sha256:0edd9213402a7bd99b02d8b9c2f3a4ee63eb0305d083b63f69f415300d8acfa9`.

```terraform
# ProtocolInspection Resource Example
# Manages Protocol Inspection Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolInspection configuration
resource "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}
```

<a id="canonical-629b6578c7dcf4e6a0a160b06bb6e8e5d416018f88991bd31206ab41768330eb"></a>

## Next pages — Resource / f9fc21f3c32d / 3

- [Examples](resources--protocol_inspection--examples--group-001.md#canonical-daa07a474b657dbd62339295ce1f9bb6e4e70833d976326d48460822d296ef4f)
- [xcsh_protocol_inspection](../resources/protocol_inspection.md#canonical-cc98a9fefd84f27d9c16e221f4eabbdce6b90ff1e1bacd91433bb7d65f1ebb12)
