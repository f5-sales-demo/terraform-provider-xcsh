---
page_title: "xcsh_rate_limiter_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy examples."
---

# xcsh_rate_limiter_policy examples

<a id="canonical-2741be1581f1f9d73872c6a2366d1a89cb9fa023b50acd8e9f1053d77c69673b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bd51a1816d62160a2fa154459a8734971b64375d50cdfa67a17a7e7afa37471"></a>

## Examples — Examples / 87953f9c598a / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- Examples

<a id="canonical-6d6ce62af0be6e42739ee498fe4ff661320d2a7751566760984f49a3761aff3f"></a>

## Complete configurations — Examples / 87953f9c598a / 3

- [Resource](resources--rate_limiter_policy--examples--group-001.md#canonical-94ba57cdccb37f673dc281e4e90c5fd010cefabc0787a92938e46f3bb2ec9d28): valid configuration.

<a id="canonical-bf7d544189683f799f14e627db7ac85b9a4a365b89483e66f6dd8f650f035916"></a>

## Next pages — Examples / 87953f9c598a / 4

- [Resource](resources--rate_limiter_policy--examples--group-001.md#canonical-94ba57cdccb37f673dc281e4e90c5fd010cefabc0787a92938e46f3bb2ec9d28)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)

<a id="canonical-94ba57cdccb37f673dc281e4e90c5fd010cefabc0787a92938e46f3bb2ec9d28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49e90bc1b048df7688127633b4d20b713f95b37b489f23d6a52a88bef3751730"></a>

## Resource — Resource / 89dd4920cd23 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
- [Examples](resources--rate_limiter_policy--examples--group-001.md#canonical-2741be1581f1f9d73872c6a2366d1a89cb9fa023b50acd8e9f1053d77c69673b)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter_policy/resource.tf`; digest `sha256:d8a1a676828e8f463f6383201f395371ac9e0d89d99902337c84f2ee8a150c35`.

```terraform
# RateLimiterPolicy Resource Example
# Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiterPolicy configuration
resource "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}
```

<a id="canonical-2390a5cf0ee9fa8e15488befba4cd1f99d110423ba44a40943c4aad33582ba66"></a>

## Next pages — Resource / 89dd4920cd23 / 3

- [Examples](resources--rate_limiter_policy--examples--group-001.md#canonical-2741be1581f1f9d73872c6a2366d1a89cb9fa023b50acd8e9f1053d77c69673b)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-015e39c6ecae39c7d173b870edd16f3a7c1c5eaeeda80b421dbccf984cb23690)
