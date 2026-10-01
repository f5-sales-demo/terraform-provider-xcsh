---
page_title: "xcsh_network_data_intelligence examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_data_intelligence examples."
---

# xcsh_network_data_intelligence examples

<a id="canonical-3b7b4a98c89d8243a85c0a9412ae43f816de2f4deca4d3e685794d6ea3227e8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edc9563b6202a3242f53ecfd993dd4255709779b0aace473dd0b1672dab43569"></a>

## Examples — Examples / 316208141070 / 2

Breadcrumbs:

- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-db1c6e7c4254880692807159269c6d76d21575ca6575b838319cf6c66703a647)
- Examples

<a id="canonical-88a1a266f758ce1c9ca3700369e29e1da1bff598fe209956438443952090520b"></a>

## Complete configurations — Examples / 316208141070 / 3

- [Data source](data-sources--network_data_intelligence--examples--group-001.md#canonical-b47dfbca9e5aa0b51de84fc5df49a0096b559b88eb718e13b6377ea07a7a532c): valid configuration.

<a id="canonical-fd9297e8cfce04707505e8c7547ce0408e578aa47feb4f59de8e0e220038cd16"></a>

## Next pages — Examples / 316208141070 / 4

- [Data source](data-sources--network_data_intelligence--examples--group-001.md#canonical-b47dfbca9e5aa0b51de84fc5df49a0096b559b88eb718e13b6377ea07a7a532c)
- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-db1c6e7c4254880692807159269c6d76d21575ca6575b838319cf6c66703a647)

<a id="canonical-b47dfbca9e5aa0b51de84fc5df49a0096b559b88eb718e13b6377ea07a7a532c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78e0fad380259f917505cb62064dd4b1c344e7da01719fdf9fe70be41383b6e8"></a>

## Data source — Data source / 1402d9db87b8 / 2

Breadcrumbs:

- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-db1c6e7c4254880692807159269c6d76d21575ca6575b838319cf6c66703a647)
- [Examples](data-sources--network_data_intelligence--examples--group-001.md#canonical-3b7b4a98c89d8243a85c0a9412ae43f816de2f4deca4d3e685794d6ea3227e8b)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_data_intelligence/data-source.tf`; digest `sha256:529d85fb608129abcf053afa1a4c7516767c8567633cbc06efe8d845d135d233`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_data_intelligence" "us" {
  regions = ["us"]
}

output "data_intelligence_https_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_data_intelligence.us.cidr_blocks
  }
}
```

<a id="canonical-3a6d72e007e8d129f2f4860611ac5049b96c1c6dac010953c2ff488bfd7ff774"></a>

## Next pages — Data source / 1402d9db87b8 / 3

- [Examples](data-sources--network_data_intelligence--examples--group-001.md#canonical-3b7b4a98c89d8243a85c0a9412ae43f816de2f4deca4d3e685794d6ea3227e8b)
- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-db1c6e7c4254880692807159269c6d76d21575ca6575b838319cf6c66703a647)
