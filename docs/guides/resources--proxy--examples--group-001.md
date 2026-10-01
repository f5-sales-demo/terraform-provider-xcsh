---
page_title: "xcsh_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy examples."
---

# xcsh_proxy examples

<a id="canonical-76036aac58de73b8f1c98ea53f3c7312dd1d105d4206b6c25fe2f04925ced618"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-954ce15bf2182b49e03cab9d57fd7867b0bf3d8c30213c13b53a5dae6a361fe3"></a>

## Examples — Examples / 366a0f592627 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- Examples

<a id="canonical-a9886e46f7806d313682fc8f8b8585197bcb3ad41c57ede1486e263d61584f35"></a>

## Complete configurations — Examples / 366a0f592627 / 3

- [Resource](resources--proxy--examples--group-001.md#canonical-5d65f73f4d488524036f2ebf0579df8377957c77494d4e02ff2b2012b1e8530b): valid configuration.

<a id="canonical-a527eeeb064942f88c796ea62d537e486622b6d60c2c9449c83dc828fe27ed8c"></a>

## Next pages — Examples / 366a0f592627 / 4

- [Resource](resources--proxy--examples--group-001.md#canonical-5d65f73f4d488524036f2ebf0579df8377957c77494d4e02ff2b2012b1e8530b)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-5d65f73f4d488524036f2ebf0579df8377957c77494d4e02ff2b2012b1e8530b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5aad8fbe94d453207d45aac3530ec6afe527098017b62972ad850641e9d7491"></a>

## Resource — Resource / 88ad30332520 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Examples](resources--proxy--examples--group-001.md#canonical-76036aac58de73b8f1c98ea53f3c7312dd1d105d4206b6c25fe2f04925ced618)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_proxy/resource.tf`; digest `sha256:8d3ef9470714fab754f0f7631e40f5f93a8f5acfb3c6def25dff74cba83c2ef7`.

```terraform
# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Proxy configuration
resource "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}
```

<a id="canonical-391033d3199d59bf2e85a287c48d05b5768db5a86ec4f5a38433ffba1f2690d4"></a>

## Next pages — Resource / 88ad30332520 / 3

- [Examples](resources--proxy--examples--group-001.md#canonical-76036aac58de73b8f1c98ea53f3c7312dd1d105d4206b6c25fe2f04925ced618)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
