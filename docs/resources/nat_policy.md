---
page_title: "xcsh_nat_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy landing."
---

# xcsh_nat_policy landing

<a id="canonical-7d6d5905e42940d480ae8451e9daddf05997d327a362e559648b22b274603e34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a7f850de11400dc0bf7a852c7ea89f9ebfc90a51c5926627e1c75a1535e11e6"></a>

## xcsh_nat_policy — xcsh_nat_policy / a69d9f495666 / 2

Breadcrumbs:

- xcsh_nat_policy

Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures
nat policy with multiple rules,. configuration.

<a id="canonical-166290be8292c7d16f9671da16bf80a4df808cd5648363585df6010479687735"></a>

## Prerequisites — xcsh_nat_policy / a69d9f495666 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a92b6532e4c58ee0fbb684437929364a1b074c1a28519ea8a5b7033a771cf029"></a>

## Minimal configuration — xcsh_nat_policy / a69d9f495666 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-5f743460c03e853922c613480cbdf5979ec7ff2668029402bc86a0dd0b3c2890"></a>

## Root configuration — xcsh_nat_policy / a69d9f495666 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d30fe186a49ff78d43fb864c0e7bd560836fa96b8316be618029ad87f20362ba"></a>

## Next pages — xcsh_nat_policy / a69d9f495666 / 6

- [Property reference](../guides/resources--nat_policy--reference--group-001.md#canonical-c8a9f7e1caceb1b458b43db01137a3474846e968f5560439c565f578127703e1)
- [Examples](../guides/resources--nat_policy--examples--group-001.md#canonical-9f086a686b98074c7bec7ba77d334f6dcd3a75314c8044dc673ce2c3dc59b167)
- [Import](../guides/resources--nat_policy--lifecycle--group-001.md#canonical-5b8a7ddd04ffb177cc2a41397624379e360e119fbafd892d3c0569a066c8fa57)
- [Timeouts](../guides/resources--nat_policy--lifecycle--group-001.md#canonical-a9629fd10d391559163654ade4b6b0fa1fc9c9c6c110d2f09189c720ff62872b)
