---
page_title: "xcsh_network_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule examples."
---

# xcsh_network_policy_rule examples

<a id="canonical-3e39234cdad1f99cf119f41c27837bd5d429c3fdbdd2062c6c3194ccc84966ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0741e9a98c4f775d495a69b91cd7ff53c58f73be3b6939c7e37eb92a84d595a1"></a>

## Examples — Examples / 5d4e0bce2a26 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- Examples

<a id="canonical-59231856aad3441d40878a9631c7ef4f14b98595a1338f04e4eb55ed37fb9b41"></a>

## Complete configurations — Examples / 5d4e0bce2a26 / 3

- [Resource](resources--network_policy_rule--examples--group-001.md#canonical-2c6fa61de67a5d45fe6a30b27a8619941447d87d42c684e2e91a14063b232eb8): valid configuration.

<a id="canonical-cd0604bc6f726ae848ccbe106094b20834a04a4a93d64744ace27f49fa972cac"></a>

## Next pages — Examples / 5d4e0bce2a26 / 4

- [Resource](resources--network_policy_rule--examples--group-001.md#canonical-2c6fa61de67a5d45fe6a30b27a8619941447d87d42c684e2e91a14063b232eb8)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)

<a id="canonical-2c6fa61de67a5d45fe6a30b27a8619941447d87d42c684e2e91a14063b232eb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c56ea40193712611f641a23b51d3a38a6802ae3612b349fba797d9093f63f42f"></a>

## Resource — Resource / 5567cbf87cd3 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
- [Examples](resources--network_policy_rule--examples--group-001.md#canonical-3e39234cdad1f99cf119f41c27837bd5d429c3fdbdd2062c6c3194ccc84966ad)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy_rule/resource.tf`; digest `sha256:0feaf66dfd5662ae652b573d32c61f50aa8eb63e0f1a54a1e0d029c6ad96ace9`.

```terraform
# NetworkPolicyRule Resource Example
# Manages network policy rule with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyRule configuration
resource "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}
```

<a id="canonical-76ff87eaf3ef3670663e97d33251cdc0f95fd8b0e766842a376290fe0922a56b"></a>

## Next pages — Resource / 5567cbf87cd3 / 3

- [Examples](resources--network_policy_rule--examples--group-001.md#canonical-3e39234cdad1f99cf119f41c27837bd5d429c3fdbdd2062c6c3194ccc84966ad)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-96221053a137de812b14a783903ec828d3abc27ef90a2ca59d6a0d2268ce52e9)
