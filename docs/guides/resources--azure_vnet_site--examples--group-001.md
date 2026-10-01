---
page_title: "xcsh_azure_vnet_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site examples."
---

# xcsh_azure_vnet_site examples

<a id="canonical-9937a731c0dc2b330f80d47e29602b5ae88158752b84247f2c29ad2df826847e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0840455b81c8008e2db5449c5d9bba1254431b75b64eafe93ea26823d5f27a8a"></a>

## Examples — Examples / 98e2e63629a4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- Examples

<a id="canonical-99bb9153f683937d2ed20047314bd813d8352545acde40e178f5d022e2fefdf0"></a>

## Complete configurations — Examples / 98e2e63629a4 / 3

- [Resource](resources--azure_vnet_site--examples--group-001.md#canonical-3888fe076c94a48b604fe4b18f2ea07c0393905f4a6feacb15e8a12a16c325bf): valid configuration.

<a id="canonical-09cf17ba4557b82326251f8726e73bc0052caae4f7eeb3d5734d927a84ed0c27"></a>

## Next pages — Examples / 98e2e63629a4 / 4

- [Resource](resources--azure_vnet_site--examples--group-001.md#canonical-3888fe076c94a48b604fe4b18f2ea07c0393905f4a6feacb15e8a12a16c325bf)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3888fe076c94a48b604fe4b18f2ea07c0393905f4a6feacb15e8a12a16c325bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d9a6b4ae4a3c125b1d39f38f5cc4a8a7349cd220d0ded9cbbad25111486dc2c"></a>

## Resource — Resource / a143ca62b93e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Examples](resources--azure_vnet_site--examples--group-001.md#canonical-9937a731c0dc2b330f80d47e29602b5ae88158752b84247f2c29ad2df826847e)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_azure_vnet_site/resource.tf`; digest `sha256:ba22e659e1847f8f68a606a677c06fc95a71eb38fbebadbdb91aec35b0057b71`.

```terraform
# AzureVNETSite Resource Example
# Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure Virtual Network environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AzureVNETSite configuration
resource "xcsh_azure_vnet_site" "example" {
  name      = "example-azure-vnet-site"
  namespace = "system"

  machine_type   = "example-value"
  resource_group = "example-value"
  ssh_key        = "example-value"
}
```

<a id="canonical-a159eddc8592a671f3f3461b1fdba264105d5bd16d8202d4f03f7ed51b710cad"></a>

## Next pages — Resource / a143ca62b93e / 3

- [Examples](resources--azure_vnet_site--examples--group-001.md#canonical-9937a731c0dc2b330f80d47e29602b5ae88158752b84247f2c29ad2df826847e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
