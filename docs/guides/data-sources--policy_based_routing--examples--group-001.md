---
page_title: "xcsh_policy_based_routing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing examples."
---

# xcsh_policy_based_routing examples

<a id="canonical-815b5acc1617054d1c4c6e24d1ebd4641bd0adb8f150bb7e456ec4289bb2c40d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2057a256f22b1863b5a8e128d470abb21101ec6f27880cc768be745ccc38b5e5"></a>

## Examples — Examples / 8a6a5c1d62b9 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- Examples

<a id="canonical-71eb3b032ee6604c497d79ab83c839d2ff45ad0bf1021bb277e50d2e3dd0cd80"></a>

## Complete configurations — Examples / 8a6a5c1d62b9 / 3

- [Data source](data-sources--policy_based_routing--examples--group-001.md#canonical-3a3afa4c67b25cb89e63af08bd8d00233e9963ef83d88835f9298ffbb5194a53): valid configuration.

<a id="canonical-0d9b00e2bcd33e15b23de0ac769e76413c3bd7e002a4e4353e1f6219dd60510d"></a>

## Next pages — Examples / 8a6a5c1d62b9 / 4

- [Data source](data-sources--policy_based_routing--examples--group-001.md#canonical-3a3afa4c67b25cb89e63af08bd8d00233e9963ef83d88835f9298ffbb5194a53)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)

<a id="canonical-3a3afa4c67b25cb89e63af08bd8d00233e9963ef83d88835f9298ffbb5194a53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4439b4c2e1ae941fec893b149cbb1136bb1fc5f9d322f5c1492a90cea1817a5"></a>

## Data source — Data source / 5821ed5490c1 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
- [Examples](data-sources--policy_based_routing--examples--group-001.md#canonical-815b5acc1617054d1c4c6e24d1ebd4641bd0adb8f150bb7e456ec4289bb2c40d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_policy_based_routing/data-source.tf`; digest `sha256:d902f95ac0d5d170c71a1f393d0718d60b4c9a9b4856200aec4f7a2fbf5ad0fc`.

```terraform
# PolicyBasedRouting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PolicyBasedRouting by name
data "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}

output "policy_based_routing_id" {
  value = data.xcsh_policy_based_routing.example.id
}
```

<a id="canonical-d71aa025198db3883162ed8fb7f479bde012b40619f804e4296c7b7e255a6203"></a>

## Next pages — Data source / 5821ed5490c1 / 3

- [Examples](data-sources--policy_based_routing--examples--group-001.md#canonical-815b5acc1617054d1c4c6e24d1ebd4641bd0adb8f150bb7e456ec4289bb2c40d)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-9f484c755ac3f52617430a5b2e5df14772b208bb98b0f23491c13317e615ab22)
