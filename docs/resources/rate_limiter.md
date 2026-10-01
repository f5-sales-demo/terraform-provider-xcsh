---
page_title: "xcsh_rate_limiter landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter landing."
---

# xcsh_rate_limiter landing

<a id="canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-290615dd6fc0d301714d5d31b47298f871080361431dee52806c792ba602a672"></a>

## xcsh_rate_limiter — xcsh_rate_limiter / b5133d7dc470 / 2

Breadcrumbs:

- xcsh_rate_limiter

Manages rate\_limiter creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-10e21dff70e48808f708ddb93a64a301c501dd17c6d1c098a6e3e0622c9994c7"></a>

## Prerequisites — xcsh_rate_limiter / b5133d7dc470 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `rate_limiter_policy`.

- rate_limiter_policy: Detailed rate limiting rules

<a id="canonical-e74eb62de793d2f3b2d43e32e900c92afabb5d5a30d538e6dca1f9d9f91b9358"></a>

## Minimal configuration — xcsh_rate_limiter / b5133d7dc470 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RateLimiter Resource Example
# Manages rate_limiter creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiter configuration
resource "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}
```

<a id="canonical-11fd0f3d55e1152febd51448800b36184c01e03362faadd0cca07848ebd14fd5"></a>

## Root configuration — xcsh_rate_limiter / b5133d7dc470 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a4399abbf52d4b2db634abd05bcaa2cbfb73df1d3fb18e79b096f51fb4cc916a"></a>

## Next pages — xcsh_rate_limiter / b5133d7dc470 / 6

- [Property reference](../guides/resources--rate_limiter--reference--group-001.md#canonical-9be48129c49b834e84b2c46957ecf1750c1fe42efcca8c4b1f1a0ed0929a4fca)
- [Examples](../guides/resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [Import](../guides/resources--rate_limiter--lifecycle--group-001.md#canonical-e77234ecd8ddd847ab72ce71eafcfd54dcbf5ac812887a81785e1cd69190ba6b)
- [Timeouts](../guides/resources--rate_limiter--lifecycle--group-001.md#canonical-f3e27b8aa13c22aceab5b617a27770da75cb39756beff21af40cf20aac001a8f)
