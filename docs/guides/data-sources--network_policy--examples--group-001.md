---
page_title: "xcsh_network_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy examples."
---

# xcsh_network_policy examples

<a id="canonical-d1d15b038df0af9dd85a494b15b270488299067a984757ebd5ff426db647ed35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9864a347b9c66f8a67ca920d418f691e64c97023538d5688a5ad5f38a462ca5"></a>

## Examples — Examples / 963bdc3e5a0d / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- Examples

<a id="canonical-3b6c140651fd46b6ef20e825f4439024f3a42523985b9cf3311c9473bf425c15"></a>

## Complete configurations — Examples / 963bdc3e5a0d / 3

- [Data source](data-sources--network_policy--examples--group-001.md#canonical-9efd6179d43446014690779070594b88261732497d18f214e882f1c75995893c): valid configuration.

<a id="canonical-b78cd7d2fe9f42ddec4eb7c28b3fa5b88a229b9872f68a03f1023d51caac8a1f"></a>

## Next pages — Examples / 963bdc3e5a0d / 4

- [Data source](data-sources--network_policy--examples--group-001.md#canonical-9efd6179d43446014690779070594b88261732497d18f214e882f1c75995893c)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)

<a id="canonical-9efd6179d43446014690779070594b88261732497d18f214e882f1c75995893c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23621d17412b75f6506c5874c382a487c3ed71dd4f031cede4c2b6d0ec94d0e0"></a>

## Data source — Data source / 469952942be7 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
- [Examples](data-sources--network_policy--examples--group-001.md#canonical-d1d15b038df0af9dd85a494b15b270488299067a984757ebd5ff426db647ed35)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy/data-source.tf`; digest `sha256:6c10eea9b1a5304339370438b5c012d3317ee2ecb95c539f946182ee126b8c01`.

```terraform
# NetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicy by name
data "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}

output "network_policy_id" {
  value = data.xcsh_network_policy.example.id
}
```

<a id="canonical-c5c48b4c4f2000ea496ab807e44ab95bc9ff2f5addf6330826efd48a44ed0212"></a>

## Next pages — Data source / 469952942be7 / 3

- [Examples](data-sources--network_policy--examples--group-001.md#canonical-d1d15b038df0af9dd85a494b15b270488299067a984757ebd5ff426db647ed35)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f)
