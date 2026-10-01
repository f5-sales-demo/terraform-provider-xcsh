---
page_title: "xcsh_network_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy landing."
---

# xcsh_network_policy landing

<a id="canonical-3497617cda5e21ac68488f7bf521f577868b56ddc0631029379984fbc02afd8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f86decb5959e9b544966fc6e3cabeabd86e86b5bab7081c3a063e2a58f9a3264"></a>

## xcsh_network_policy — xcsh_network_policy / fdceecd36315 / 2

Breadcrumbs:

- xcsh_network_policy

Manages new network policy with configured parameters in specified namespace in F5 Distributed
Cloud.

<a id="canonical-e6c5b7858eccf86ac713d100371d6fdbaa12d31db158fe474cb476d43e37e01a"></a>

## Prerequisites — xcsh_network_policy / fdceecd36315 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-c5b381f646280cea713ce02729e8bd83ffd8e66ac346a320fad992e2b7b41145"></a>

## Minimal configuration — xcsh_network_policy / fdceecd36315 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-e0aee27502488092cd963855729636a6736c7ad9e5866abf60c4109d5f363fb8"></a>

## Root configuration — xcsh_network_policy / fdceecd36315 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-4757112fb170774c1e6effbf50adffc167ce428e66a4716dbe7ddb0697c5bd17"></a>

## Next pages — xcsh_network_policy / fdceecd36315 / 6

- [Property reference](../guides/data-sources--network_policy--reference--group-001.md#canonical-754e2a526c5d4eaa3c63be9b2b9acdd5575af1fc91af7e96c519bc1b95c7b5ff)
- [Examples](../guides/data-sources--network_policy--examples--group-001.md#canonical-d1d15b038df0af9dd85a494b15b270488299067a984757ebd5ff426db647ed35)
