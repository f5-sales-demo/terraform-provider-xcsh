---
page_title: "xcsh_network_policy_view examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view examples."
---

# xcsh_network_policy_view examples

<a id="canonical-3c9c66dfd1fa47b761e332e76bd2fd4801b1dd120cf96d2162cd28460590654d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e27950c01e8dd20e83fcc8aa14d760879b6d11ad7f0c6572ad7f2659d55c1d14"></a>

## Examples — Examples / 64316716d9d4 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- Examples

<a id="canonical-a3bd8b9f7f213b990ff16f246c4e889ad12ceef17720d96d869d537f716f2a40"></a>

## Complete configurations — Examples / 64316716d9d4 / 3

- [Data source](data-sources--network_policy_view--examples--group-001.md#canonical-59aff966f44ce63105edb3557c0165b01c9482ded14a6c77c1eb29b7f0896ecc): valid configuration.

<a id="canonical-2dcd1bd6e18f31449fd0129abcf23f5491ec6030716e0d010ab29dbcca584ace"></a>

## Next pages — Examples / 64316716d9d4 / 4

- [Data source](data-sources--network_policy_view--examples--group-001.md#canonical-59aff966f44ce63105edb3557c0165b01c9482ded14a6c77c1eb29b7f0896ecc)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)

<a id="canonical-59aff966f44ce63105edb3557c0165b01c9482ded14a6c77c1eb29b7f0896ecc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd3062850542577be17d415415d7e3e2d75be838bb220a4bb154ff2980b4239b"></a>

## Data source — Data source / 1fc2ffd28b57 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
- [Examples](data-sources--network_policy_view--examples--group-001.md#canonical-3c9c66dfd1fa47b761e332e76bd2fd4801b1dd120cf96d2162cd28460590654d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_view/data-source.tf`; digest `sha256:961fb42feeb0126081bc0ae8691b145c93619816626daf020f96266ec095e50a`.

```terraform
# NetworkPolicyView Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyView by name
data "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.xcsh_network_policy_view.example.id
}
```

<a id="canonical-82bc37d3ac92ba06c6b0952392d54a04028b2de88299d20e518f5959039415aa"></a>

## Next pages — Data source / 1fc2ffd28b57 / 3

- [Examples](data-sources--network_policy_view--examples--group-001.md#canonical-3c9c66dfd1fa47b761e332e76bd2fd4801b1dd120cf96d2162cd28460590654d)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18)
