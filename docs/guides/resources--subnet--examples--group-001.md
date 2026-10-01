---
page_title: "xcsh_subnet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet examples."
---

# xcsh_subnet examples

<a id="canonical-673f0e65e6cb760bcb080dd331c2cf6d8f2132a3a7647cef482d76d0c71001a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-026b468580883be88315bea5239cf75a50f7a54987378bc033516f3959deb12d"></a>

## Examples — Examples / 4163efaa52e3 / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- Examples

<a id="canonical-54b9c112142d2666f23a6152526ccb34b8ca6d7549197f29e4ec8b1f388a115c"></a>

## Complete configurations — Examples / 4163efaa52e3 / 3

- [Resource](resources--subnet--examples--group-001.md#canonical-23144de1af45dbf583364bc736aa6b7be353b841ad02d88a949bad03dd8a308a): valid configuration.

<a id="canonical-a0e54eecb47de03349b0ca2b9596763a611aac72e90d1d71d850513eb8429aac"></a>

## Next pages — Examples / 4163efaa52e3 / 4

- [Resource](resources--subnet--examples--group-001.md#canonical-23144de1af45dbf583364bc736aa6b7be353b841ad02d88a949bad03dd8a308a)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)

<a id="canonical-23144de1af45dbf583364bc736aa6b7be353b841ad02d88a949bad03dd8a308a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be2c645708f404142a7f747f4bdf5da091736fc31a3d367c4d023ddc54ee1382"></a>

## Resource — Resource / 8b4ae55a504f / 2

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
- [Examples](resources--subnet--examples--group-001.md#canonical-673f0e65e6cb760bcb080dd331c2cf6d8f2132a3a7647cef482d76d0c71001a7)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_subnet/resource.tf`; digest `sha256:4b17c21430cf56ad4295d650386f0dba8c0f28367f60769889a0c3628d334ae0`.

```terraform
# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}
```

<a id="canonical-7fb43bc2a164f4c63a33d24aba88f236a1d778b79ea2461ed59edf1efd2f308c"></a>

## Next pages — Resource / 8b4ae55a504f / 3

- [Examples](resources--subnet--examples--group-001.md#canonical-673f0e65e6cb760bcb080dd331c2cf6d8f2132a3a7647cef482d76d0c71001a7)
- [xcsh_subnet](../resources/subnet.md#canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41)
