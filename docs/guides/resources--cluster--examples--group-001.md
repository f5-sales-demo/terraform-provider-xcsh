---
page_title: "xcsh_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster examples."
---

# xcsh_cluster examples

<a id="canonical-2a64e3e2f1d23626e4a5d1d790bae01ba53b4660306296670e9eb3b6281ad329"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3328037592590aaaf7e47ebb6a66df6f2bcfcedaf1f0de433f864fb123e25417"></a>

## Examples — Examples / 71daf56c72df / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- Examples

<a id="canonical-ac808541d08e4fbb4e9990a49396a388272be1e016fce107d3ef02e1cc802cea"></a>

## Complete configurations — Examples / 71daf56c72df / 3

- [Resource](resources--cluster--examples--group-001.md#canonical-3df8ab9b7a936ebbc522a18fd4d025a26e12185a326921fe3123c194360b2e8a): valid configuration.

<a id="canonical-3fcd54030e10a3e8e32f1743655984adfdf98cb0c14c2a661b0bcf734832a31b"></a>

## Next pages — Examples / 71daf56c72df / 4

- [Resource](resources--cluster--examples--group-001.md#canonical-3df8ab9b7a936ebbc522a18fd4d025a26e12185a326921fe3123c194360b2e8a)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)

<a id="canonical-3df8ab9b7a936ebbc522a18fd4d025a26e12185a326921fe3123c194360b2e8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a5dffc591c39a93d3282f6c3935eb7ae57e65d41c4bc37bfeb19c794d491f20"></a>

## Resource — Resource / 3fb1b9170bc9 / 2

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
- [Examples](resources--cluster--examples--group-001.md#canonical-2a64e3e2f1d23626e4a5d1d790bae01ba53b4660306296670e9eb3b6281ad329)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cluster/resource.tf`; digest `sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e`.

```terraform
# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cluster configuration
resource "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}
```

<a id="canonical-7989e97d3cdf3ec3f4662d3c6de2c640817371b127802fb8e1f78dc984d7894e"></a>

## Next pages — Resource / 3fb1b9170bc9 / 3

- [Examples](resources--cluster--examples--group-001.md#canonical-2a64e3e2f1d23626e4a5d1d790bae01ba53b4660306296670e9eb3b6281ad329)
- [xcsh_cluster](../resources/cluster.md#canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03)
