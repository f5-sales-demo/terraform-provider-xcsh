---
page_title: "xcsh_network_policy_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_set examples."
---

# xcsh_network_policy_set examples

<a id="canonical-6e5d96c0f0c95a4764658184ce943e95d603c429be3d17a52e3aa762848638d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1094e56d55dc958d549b309e1e7e162e0e248e85ca233d3d9096cf9c9485a26"></a>

## Examples — Examples / c1d5a757c3ee / 2

Breadcrumbs:

- [xcsh_network_policy_set](../data-sources/network_policy_set.md#canonical-ca548052a89d7f647ebf8e8303af31a758be406e22a36f86b6a13ff023dd5e7d)
- Examples

<a id="canonical-6d7c73c4877dba6ea9b4ecfd76ad72e44ad34545b6120836c3c0f5350fe20323"></a>

## Complete configurations — Examples / c1d5a757c3ee / 3

- [Data source](data-sources--network_policy_set--examples--group-001.md#canonical-54696f03221327164ea5280e0eb5dc4e85b375d75e8348ea574d0c7be0566a52): valid configuration.

<a id="canonical-4a649bf4b6192e82342af92f986578a152d3e7b6b0ce99b283c5e182b0d35362"></a>

## Next pages — Examples / c1d5a757c3ee / 4

- [Data source](data-sources--network_policy_set--examples--group-001.md#canonical-54696f03221327164ea5280e0eb5dc4e85b375d75e8348ea574d0c7be0566a52)
- [xcsh_network_policy_set](../data-sources/network_policy_set.md#canonical-ca548052a89d7f647ebf8e8303af31a758be406e22a36f86b6a13ff023dd5e7d)

<a id="canonical-54696f03221327164ea5280e0eb5dc4e85b375d75e8348ea574d0c7be0566a52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72e1530b8c30c27f4ab06dc93aefb04a9bb139d190bb9724981be48159a4499a"></a>

## Data source — Data source / 3f843f7bd71f / 2

Breadcrumbs:

- [xcsh_network_policy_set](../data-sources/network_policy_set.md#canonical-ca548052a89d7f647ebf8e8303af31a758be406e22a36f86b6a13ff023dd5e7d)
- [Examples](data-sources--network_policy_set--examples--group-001.md#canonical-6e5d96c0f0c95a4764658184ce943e95d603c429be3d17a52e3aa762848638d6)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_set/data-source.tf`; digest `sha256:292e820cf31aaf1ca4db02cc9e7027c764e59f12b6fedea7ebfc71f60117c57e`.

```terraform
# NetworkPolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicySet by name
data "xcsh_network_policy_set" "example" {
  name      = "example-network-policy-set"
  namespace = "staging"
}

output "network_policy_set_id" {
  value = data.xcsh_network_policy_set.example.id
}
```

<a id="canonical-4dca8832179674e7843f61ce598d43e9597b83e46a2d72f6e6c23d4d8dcd1891"></a>

## Next pages — Data source / 3f843f7bd71f / 3

- [Examples](data-sources--network_policy_set--examples--group-001.md#canonical-6e5d96c0f0c95a4764658184ce943e95d603c429be3d17a52e3aa762848638d6)
- [xcsh_network_policy_set](../data-sources/network_policy_set.md#canonical-ca548052a89d7f647ebf8e8303af31a758be406e22a36f86b6a13ff023dd5e7d)
