---
page_title: "xcsh_proxy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy landing."
---

# xcsh_proxy landing

<a id="canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c13297201ccdaaa8c803ea53b7b1fc4a12a0495471a086bafabd550a8d07108a"></a>

## xcsh_proxy — xcsh_proxy / 4ce9ea69d2cc / 2

Breadcrumbs:

- xcsh_proxy

Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.
configuration.

<a id="canonical-3fc389c6516454b7e4bd54937bfba54dfc25ccb6c40f28150f0599fa47c1a770"></a>

## Prerequisites — xcsh_proxy / 4ce9ea69d2cc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-356b1387b8d215722e1a4b1008185b1ac67196402b6680f85d4a715928e84ccb"></a>

## Minimal configuration — xcsh_proxy / 4ce9ea69d2cc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-92fa053e87f5efae27c110b1bb8a3d918951057306650d9685db6592792a3d5a"></a>

## Root configuration — xcsh_proxy / 4ce9ea69d2cc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b615d5ba062447afaf3a875a3d1da1a76d28a3f4de47c3d7cdfb7fe2f1390cb6"></a>

## Next pages — xcsh_proxy / 4ce9ea69d2cc / 6

- [Property reference](../guides/resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [Examples](../guides/resources--proxy--examples--group-001.md#canonical-76036aac58de73b8f1c98ea53f3c7312dd1d105d4206b6c25fe2f04925ced618)
- [Import](../guides/resources--proxy--lifecycle--group-001.md#canonical-af23da54e8e5dc8e151b0f51fb1785f7274f730be5e09e5a47e21e97a81c8d9b)
- [Timeouts](../guides/resources--proxy--lifecycle--group-001.md#canonical-9777ca2b81a3b69b19b04622d413e25e03fd001b4511cba34b0acb7cafe0065e)
