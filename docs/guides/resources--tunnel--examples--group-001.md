---
page_title: "xcsh_tunnel examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel examples."
---

# xcsh_tunnel examples

<a id="canonical-b70f6360292460f47d62d7c611fcc6bc888c2d4745820fdb7eea37be934eda5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06bd3af7b7176d99e2d80f984ff8917e45062b7b2ca2e726d9169008bd5be146"></a>

## Examples — Examples / 06e7aebb30fa / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- Examples

<a id="canonical-50e79c4746db5408c9ddc4930148da15d79994882fb6161d02901475ae85d454"></a>

## Complete configurations — Examples / 06e7aebb30fa / 3

- [Resource](resources--tunnel--examples--group-001.md#canonical-d8e9858a0b76a227f7bfdcdd491b759120c1da7c2a7c1880f4b71516fb9f3f8b): valid configuration.

<a id="canonical-2e325d8732da936a5c9fb797a54cb3abb3afb1a6cdeabdafaa30506f236f2c8e"></a>

## Next pages — Examples / 06e7aebb30fa / 4

- [Resource](resources--tunnel--examples--group-001.md#canonical-d8e9858a0b76a227f7bfdcdd491b759120c1da7c2a7c1880f4b71516fb9f3f8b)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)

<a id="canonical-d8e9858a0b76a227f7bfdcdd491b759120c1da7c2a7c1880f4b71516fb9f3f8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-033ccfe540f2c56c282b37e88facdab7443d15732f13e95bddfd57122cce968b"></a>

## Resource — Resource / 1efc058b0779 / 2

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
- [Examples](resources--tunnel--examples--group-001.md#canonical-b70f6360292460f47d62d7c611fcc6bc888c2d4745820fdb7eea37be934eda5d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tunnel/resource.tf`; digest `sha256:d8686dd99ea2784f4ce424a8e06432fe8502eabcb7549c1f57d3a841fff62174`.

```terraform
# Tunnel Resource Example
# Manages tunnel in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tunnel configuration
resource "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}
```

<a id="canonical-53314753861d5a197f4ce52300271958f8d31ca39350d910417375ffb6050777"></a>

## Next pages — Resource / 1efc058b0779 / 3

- [Examples](resources--tunnel--examples--group-001.md#canonical-b70f6360292460f47d62d7c611fcc6bc888c2d4745820fdb7eea37be934eda5d)
- [xcsh_tunnel](../resources/tunnel.md#canonical-2325b7e0e580c98f73556b64de5a592b581abb0dd80f7c3e6a9d7542d53e801a)
