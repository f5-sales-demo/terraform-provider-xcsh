---
page_title: "xcsh_nat_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy examples."
---

# xcsh_nat_policy examples

<a id="canonical-9f086a686b98074c7bec7ba77d334f6dcd3a75314c8044dc673ce2c3dc59b167"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49ef80084a7e412d352f7a08844baaa6a7a2a2faaf26eb849e4c965f41b73ffe"></a>

## Examples — Examples / 3372050254ee / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- Examples

<a id="canonical-b94470c0e283202cd67c74d01a3cc357a7687bc9ce91ede1dac0795bacc6be6b"></a>

## Complete configurations — Examples / 3372050254ee / 3

- [Resource](resources--nat_policy--examples--group-001.md#canonical-b2e10afb533d18af85f92686b9dda20f739773d356e9eafd795d8af302d40c58): valid configuration.

<a id="canonical-901d0e9f64be6dd6d634f8d34f44e9ecc42dad000ac004613510594032412bce"></a>

## Next pages — Examples / 3372050254ee / 4

- [Resource](resources--nat_policy--examples--group-001.md#canonical-b2e10afb533d18af85f92686b9dda20f739773d356e9eafd795d8af302d40c58)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)

<a id="canonical-b2e10afb533d18af85f92686b9dda20f739773d356e9eafd795d8af302d40c58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2e6c954ccf9e86489a10312ceabd6f133a4b6f2d872e1ad89bcb067c6460d54"></a>

## Resource — Resource / c60e0d80e8ed / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
- [Examples](resources--nat_policy--examples--group-001.md#canonical-9f086a686b98074c7bec7ba77d334f6dcd3a75314c8044dc673ce2c3dc59b167)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nat_policy/resource.tf`; digest `sha256:9976ae5d67503e6c32870094ecdf09b471ea80cfbc484860002722af35c9de08`.

```terraform
# NATPolicy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NATPolicy configuration
resource "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}
```

<a id="canonical-3f781681f4d835d1e1f9fa4f3a2f8440457ebf302827ca70c9174f7594fbeafb"></a>

## Next pages — Resource / c60e0d80e8ed / 3

- [Examples](resources--nat_policy--examples--group-001.md#canonical-9f086a686b98074c7bec7ba77d334f6dcd3a75314c8044dc673ce2c3dc59b167)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34)
