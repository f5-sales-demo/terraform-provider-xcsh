---
page_title: "xcsh_network_regional_edges examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_regional_edges examples."
---

# xcsh_network_regional_edges examples

<a id="canonical-69cbc256fdb601aea7a5046bc5c1d7e5f6bd51f16032486f663d435c93baf01e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95dd6a992e6985b79fe91daa570bea5e716df1b04d5198a84c9e4f26af8a600a"></a>

## Examples — Examples / 4d3fc2fc0b22 / 2

Breadcrumbs:

- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-4b07492253f8f10e61287f15adccadbfc866bfc8be3823bb27bb4e02b5d71089)
- Examples

<a id="canonical-c89d81054a1c22568a5ef11d5e0c43a137e660111e5154cf6cd09c4259288ee8"></a>

## Complete configurations — Examples / 4d3fc2fc0b22 / 3

- [Data source](data-sources--network_regional_edges--examples--group-001.md#canonical-c676f97f17765578c5ebd964a0dffc11ee114baf891e06f6756886269409b94b): valid configuration.

<a id="canonical-1f29b5642b4dee876abd8aafd22769cc58ddd9bb612b7c49bb524cbc31297fa3"></a>

## Next pages — Examples / 4d3fc2fc0b22 / 4

- [Data source](data-sources--network_regional_edges--examples--group-001.md#canonical-c676f97f17765578c5ebd964a0dffc11ee114baf891e06f6756886269409b94b)
- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-4b07492253f8f10e61287f15adccadbfc866bfc8be3823bb27bb4e02b5d71089)

<a id="canonical-c676f97f17765578c5ebd964a0dffc11ee114baf891e06f6756886269409b94b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b0b5479689f4b12c706e026ab8fda17ec7c3927f7261bc1dae0a9c45f68d593"></a>

## Data source — Data source / d5978eb17207 / 2

Breadcrumbs:

- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-4b07492253f8f10e61287f15adccadbfc866bfc8be3823bb27bb4e02b5d71089)
- [Examples](data-sources--network_regional_edges--examples--group-001.md#canonical-69cbc256fdb601aea7a5046bc5c1d7e5f6bd51f16032486f663d435c93baf01e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_regional_edges/data-source.tf`; digest `sha256:1744edb321716a35476a9a25cade01cfd9671a46b90a9dfb2f90ccff555099d6`.

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

# Select the Regional Edge source networks that may initiate HTTPS connections
# to an origin. Omit regions to return all published regions.
data "xcsh_network_regional_edges" "origin_ingress" {
  regions = ["americas", "europe"]
}

output "https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_regional_edges.origin_ingress.cidr_blocks
  }
}
```

<a id="canonical-fac70c63c964ed1aff8442b2df9bf1330dbbd732a8a56a3075fa5cbd10e4eb7d"></a>

## Next pages — Data source / d5978eb17207 / 3

- [Examples](data-sources--network_regional_edges--examples--group-001.md#canonical-69cbc256fdb601aea7a5046bc5c1d7e5f6bd51f16032486f663d435c93baf01e)
- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-4b07492253f8f10e61287f15adccadbfc866bfc8be3823bb27bb4e02b5d71089)
