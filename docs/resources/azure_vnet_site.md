---
page_title: "xcsh_azure_vnet_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site landing."
---

# xcsh_azure_vnet_site landing

<a id="canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1f430cec2e7b0e9df4560c6843d3bcefab776a956446c932ad0f6595d2e68b7"></a>

## xcsh_azure_vnet_site — xcsh_azure_vnet_site / 9577a2eb032a / 2

Breadcrumbs:

- xcsh_azure_vnet_site

Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure
Virtual Network environments.

<a id="canonical-f9af7233854af87cb34f051d996b32e4fdedab42b169c4c89cb0ba1a276c1aa2"></a>

## Prerequisites — xcsh_azure_vnet_site / 9577a2eb032a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: Azure authentication for deployment

<a id="canonical-b20b9a94081be9055f11a8acc0eb98663cf5f1eb865acbb5a5c632842c4d7b3e"></a>

## Minimal configuration — xcsh_azure_vnet_site / 9577a2eb032a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-d1cbe5e484cfbd05d652cb2273a6a46fc57c4106e750e2dc60adc4d77b448b3c"></a>

## Root configuration — xcsh_azure_vnet_site / 9577a2eb032a / 5

Required root properties: `machine_type`, `name`, `resource_group`, `ssh_key`. Full root flags and choices appear in the property reference.

<a id="canonical-7fac16052cb9d5e1c83595caab84e513ecfaf367f06040a0bb9dfcebb1e23d4a"></a>

## Next pages — xcsh_azure_vnet_site / 9577a2eb032a / 6

- [Property reference](../guides/resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [Examples](../guides/resources--azure_vnet_site--examples--group-001.md#canonical-9937a731c0dc2b330f80d47e29602b5ae88158752b84247f2c29ad2df826847e)
- [Import](../guides/resources--azure_vnet_site--lifecycle--group-001.md#canonical-f648ca3ca347ef7f27ab59f79b2ea2fc8e404b3aa0320144c0651eb872996c1d)
- [Timeouts](../guides/resources--azure_vnet_site--lifecycle--group-001.md#canonical-657250870f51d2fc065cd36a65251ff9d62fee97de47bd65305eba10ee17e64d)
