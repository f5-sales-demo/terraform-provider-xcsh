---
page_title: "xcsh_policy_based_routing landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing landing."
---

# xcsh_policy_based_routing landing

<a id="canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b8462451f0806fe0aee5bd45011bcd67c303c0d6502747a9461d1a11ab6c679"></a>

## xcsh_policy_based_routing — xcsh_policy_based_routing / 011d30b1cd1a / 2

Breadcrumbs:

- xcsh_policy_based_routing

Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing
create specification. configuration.

<a id="canonical-cb1dc3ad5c4819e7e851fcb9536a3298ba884dfbd4b043be1ff6c6d4dc0208e4"></a>

## Prerequisites — xcsh_policy_based_routing / 011d30b1cd1a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ab1ff8a3c0d999a63781953df46bdd06f34e516dfb7513730bbd3efc10e9f6e9"></a>

## Minimal configuration — xcsh_policy_based_routing / 011d30b1cd1a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PolicyBasedRouting Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PolicyBasedRouting configuration
resource "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}
```

<a id="canonical-037987cef7e58cf5a56a19726b4bdc6f50a4ae28724eb0327f484a06b5502f5f"></a>

## Root configuration — xcsh_policy_based_routing / 011d30b1cd1a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-5be56c029a018670a3ebf75d8a2b070c185988629da3754c90e6534dec17e861"></a>

## Next pages — xcsh_policy_based_routing / 011d30b1cd1a / 6

- [Property reference](../guides/resources--policy_based_routing--reference--group-001.md#canonical-7a80a6f1ee6f6296fe821d9b09a9003d6bf3b5c5dd54a0de6b0fb4dce9623d4e)
- [Examples](../guides/resources--policy_based_routing--examples--group-001.md#canonical-2b37fbed9b408e1b36ae3feaeae2737ab01aa3c0205049bbd43a7fe1d7bd105d)
- [Import](../guides/resources--policy_based_routing--lifecycle--group-001.md#canonical-569d51116aad9d0df02aee8921d14c4eb7f4859e1e2acd8ab49f9329ee070b1e)
- [Timeouts](../guides/resources--policy_based_routing--lifecycle--group-001.md#canonical-4679ef1fe84c62092c5f5cc4b7f35d425f28a788cbce4bedb77db7ea8054188c)
