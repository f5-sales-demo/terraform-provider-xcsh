---
page_title: "xcsh_network_policy_view examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view examples."
---

# xcsh_network_policy_view examples

<a id="canonical-15d10fdee37558d71631ffa0a9e48b246c466ebfd04d568290f1d59bec87a79f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-347c3e73a32f1b27c4f826ae813d851fe33c21bc2f9720d37a56a6e1402c60f6"></a>

## Examples — Examples / 946d9681794d / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- Examples

<a id="canonical-50e5a95697555af40f074804e6c4733c5697917b83b9374877e75ef587321b1d"></a>

## Complete configurations — Examples / 946d9681794d / 3

- [Resource](resources--network_policy_view--examples--group-001.md#canonical-e913f2f1144dfda68ff2ce1eaf9fb3f5064c3cdbe68b6235d67233ced23d353e): valid configuration.

<a id="canonical-2ee1b564226238e025f53fed5634bbb06b599308a85ab26171b2febbbfc5ac5e"></a>

## Next pages — Examples / 946d9681794d / 4

- [Resource](resources--network_policy_view--examples--group-001.md#canonical-e913f2f1144dfda68ff2ce1eaf9fb3f5064c3cdbe68b6235d67233ced23d353e)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)

<a id="canonical-e913f2f1144dfda68ff2ce1eaf9fb3f5064c3cdbe68b6235d67233ced23d353e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62e1fd585ae7d5f0668c7f99e43354356f6e65df694567bbaef20a53b77c30ab"></a>

## Resource — Resource / d2246fedb4d6 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
- [Examples](resources--network_policy_view--examples--group-001.md#canonical-15d10fdee37558d71631ffa0a9e48b246c466ebfd04d568290f1d59bec87a79f)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy_view/resource.tf`; digest `sha256:75bf659169fe571161ebe1df55c559339c5fc7697dcbe1c2ae2a1986ba0ba320`.

```terraform
# NetworkPolicyView Resource Example
# Manages a Network Policy View resource in F5 Distributed Cloud for network policy view specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyView configuration
resource "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}
```

<a id="canonical-0acbd54817194292d021f9efb8d7d344acd924dbfac610eeb1438e8fa20f9c3d"></a>

## Next pages — Resource / d2246fedb4d6 / 3

- [Examples](resources--network_policy_view--examples--group-001.md#canonical-15d10fdee37558d71631ffa0a9e48b246c466ebfd04d568290f1d59bec87a79f)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13)
