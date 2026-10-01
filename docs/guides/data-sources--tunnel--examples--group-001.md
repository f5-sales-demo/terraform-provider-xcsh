---
page_title: "xcsh_tunnel examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel examples."
---

# xcsh_tunnel examples

<a id="canonical-f0f2d675aa3ad4df19b03048afbed9c24424795f3a5748460812f25e5fbb78f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5d69e1e5833fc2a94cd47c611ec0f113f512b335fbb89ca9b3347ef7654e8ad"></a>

## Examples — Examples / 86d60933e098 / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- Examples

<a id="canonical-99ae16200725fd9b503a04ae535a41bd2d35575a77f92ffdc317211b9e5fc08b"></a>

## Complete configurations — Examples / 86d60933e098 / 3

- [Data source](data-sources--tunnel--examples--group-001.md#canonical-b4165823a9fc12b159e957fc00de763108b1ad46c5b28626fb7283cf78737d6a): valid configuration.

<a id="canonical-fbb04d8e6c5c11fa05d8db948200acf2a334479eb0522139770ced2f2b57042c"></a>

## Next pages — Examples / 86d60933e098 / 4

- [Data source](data-sources--tunnel--examples--group-001.md#canonical-b4165823a9fc12b159e957fc00de763108b1ad46c5b28626fb7283cf78737d6a)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)

<a id="canonical-b4165823a9fc12b159e957fc00de763108b1ad46c5b28626fb7283cf78737d6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d745afe047902f52bb5c2ed064d62009b927851be06a94feefeb65e976816b57"></a>

## Data source — Data source / f605dcfdb96c / 2

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
- [Examples](data-sources--tunnel--examples--group-001.md#canonical-f0f2d675aa3ad4df19b03048afbed9c24424795f3a5748460812f25e5fbb78f4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tunnel/data-source.tf`; digest `sha256:6bb09ba6599023728c287794941b0e304024407337d302d5ca4e83c7bdd1f81d`.

```terraform
# Tunnel Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Tunnel by name
data "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}

output "tunnel_id" {
  value = data.xcsh_tunnel.example.id
}
```

<a id="canonical-634d5023f8688c5ddad8dcadaa9162bd64968e7c3bc76a42210b53505e9514ec"></a>

## Next pages — Data source / f605dcfdb96c / 3

- [Examples](data-sources--tunnel--examples--group-001.md#canonical-f0f2d675aa3ad4df19b03048afbed9c24424795f3a5748460812f25e5fbb78f4)
- [xcsh_tunnel](../data-sources/tunnel.md#canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081)
