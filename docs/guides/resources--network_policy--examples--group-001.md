---
page_title: "xcsh_network_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy examples."
---

# xcsh_network_policy examples

<a id="canonical-ed842b56a95a4e0145e96ee3d036d74fdbe1d9e6ffda6cefc8267078b0b528a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d881248f676fe8004567285179a01057bfcdd3b9125cfceb67f0305bb68cea6"></a>

## Examples — Examples / 6112c9db13f3 / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- Examples

<a id="canonical-b5d95a0a023a90bed9339ad26bcb3aa145d6469068bc8048977cf197a7c06ffc"></a>

## Complete configurations — Examples / 6112c9db13f3 / 3

- [Resource](resources--network_policy--examples--group-001.md#canonical-2717ba36fd43aadb64bfe0ef32e9dc55649a1c10cbc1309ab5a264beb71fe0d9): valid configuration.

<a id="canonical-3af2e77b67126950ff48b1743ff3bdaf53c94daf1346cfcdbf803b2540ee5072"></a>

## Next pages — Examples / 6112c9db13f3 / 4

- [Resource](resources--network_policy--examples--group-001.md#canonical-2717ba36fd43aadb64bfe0ef32e9dc55649a1c10cbc1309ab5a264beb71fe0d9)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)

<a id="canonical-2717ba36fd43aadb64bfe0ef32e9dc55649a1c10cbc1309ab5a264beb71fe0d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3934069ae47188c5d952fc0abc40011c7d0b0cb5b14c974182fdcc17ffd20fa"></a>

## Resource — Resource / fcb6531ebe5d / 2

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
- [Examples](resources--network_policy--examples--group-001.md#canonical-ed842b56a95a4e0145e96ee3d036d74fdbe1d9e6ffda6cefc8267078b0b528a9)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy/resource.tf`; digest `sha256:4d106a33f6bcd90b712f1a25c8242900f48156f66650342f586b89690227f887`.

```terraform
# NetworkPolicy Resource Example
# Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicy configuration
resource "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}
```

<a id="canonical-ba8f34d8c2139c007d5af99e1e13cd79cbcd5d1a660370b8ced0c8ebc25d84ac"></a>

## Next pages — Resource / fcb6531ebe5d / 3

- [Examples](resources--network_policy--examples--group-001.md#canonical-ed842b56a95a4e0145e96ee3d036d74fdbe1d9e6ffda6cefc8267078b0b528a9)
- [xcsh_network_policy](../resources/network_policy.md#canonical-f38a470b0a5f76b756805cb3c7f8aafe56d7b84dc1961a59aee5b03d000a7016)
