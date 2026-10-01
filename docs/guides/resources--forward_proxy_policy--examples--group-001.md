---
page_title: "xcsh_forward_proxy_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy examples."
---

# xcsh_forward_proxy_policy examples

<a id="canonical-8c48d5959f98bb2710e87756ddab35325fdfa094b84af6db832080dd50ddbb18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d37968f09bd64c06f8287b6a64dacea2c19ef0946dda2af0180eff40ffd891fd"></a>

## Examples — Examples / 6968abbb98b4 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- Examples

<a id="canonical-85a2e12cd025691c2f1c7e7e85da91fdc0d9e808fa442be59d97060f1f52510a"></a>

## Complete configurations — Examples / 6968abbb98b4 / 3

- [Resource](resources--forward_proxy_policy--examples--group-001.md#canonical-7f081d99243c5434bd0dab5d830832b6d73d98045860f80609d47111b3a96f51): valid configuration.

<a id="canonical-63eb433e8ec94569ce4b6abb1c8add75d158cf431a83994f229fd127f4934678"></a>

## Next pages — Examples / 6968abbb98b4 / 4

- [Resource](resources--forward_proxy_policy--examples--group-001.md#canonical-7f081d99243c5434bd0dab5d830832b6d73d98045860f80609d47111b3a96f51)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-7f081d99243c5434bd0dab5d830832b6d73d98045860f80609d47111b3a96f51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42376932a4b8f7823374c3597a25a2b5aa772175f144b493d7a88ee47486d0d2"></a>

## Resource — Resource / 25bb23a9c6a4 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Examples](resources--forward_proxy_policy--examples--group-001.md#canonical-8c48d5959f98bb2710e87756ddab35325fdfa094b84af6db832080dd50ddbb18)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forward_proxy_policy/resource.tf`; digest `sha256:bcc594a880ce1beadc720b2bd1c76066016eb6d5489e15cf1b1295e1e8d40584`.

```terraform
# ForwardProxyPolicy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardProxyPolicy configuration
resource "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}
```

<a id="canonical-77301c5e00ae34d0a33fbf9580015c1c403ca75a647efa25c786905d4aff966c"></a>

## Next pages — Resource / 25bb23a9c6a4 / 3

- [Examples](resources--forward_proxy_policy--examples--group-001.md#canonical-8c48d5959f98bb2710e87756ddab35325fdfa094b84af6db832080dd50ddbb18)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
