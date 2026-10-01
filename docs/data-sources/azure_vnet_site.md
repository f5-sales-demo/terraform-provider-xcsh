---
page_title: "xcsh_azure_vnet_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site landing."
---

# xcsh_azure_vnet_site landing

<a id="canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aef00c341ea1e651f9b07868491e6b33f737fceee71f4a17c792e73266d7b904"></a>

## xcsh_azure_vnet_site — xcsh_azure_vnet_site / 0bc1d68c746e / 2

Breadcrumbs:

- xcsh_azure_vnet_site

Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure
Virtual Network environments.

<a id="canonical-6ddd74814afeb6c4746b7a762d22883fc44b12b8b0dd41a9f6d79bc1732c2f73"></a>

## Prerequisites — xcsh_azure_vnet_site / 0bc1d68c746e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: Azure authentication for deployment

<a id="canonical-aacebf93ddc33df3260246ccd8c7ae14e98686fd9b0bb8e35eedb9cc52e95322"></a>

## Minimal configuration — xcsh_azure_vnet_site / 0bc1d68c746e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AzureVNETSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AzureVNETSite by name
data "xcsh_azure_vnet_site" "example" {
  name      = "example-azure-vnet-site"
  namespace = "system"
}

output "azure_vnet_site_id" {
  value = data.xcsh_azure_vnet_site.example.id
}
```

<a id="canonical-acfb88929ce83c58326084f51390ef8b0a1ef19fffb91c543ecad7cd9989923f"></a>

## Root configuration — xcsh_azure_vnet_site / 0bc1d68c746e / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-e524805f0fff9e6c22d48b459e0c793568fc16a7f611c5cb5d33b9ad566bc734"></a>

## Next pages — xcsh_azure_vnet_site / 0bc1d68c746e / 6

- [Property reference](../guides/data-sources--azure_vnet_site--reference--group-001.md#canonical-a4424f966fa597f1188eace4a92ab52837b368faae830053b1a037d77fbacc17)
- [Examples](../guides/data-sources--azure_vnet_site--examples--group-001.md#canonical-f49a5920408705296143b8899177eca7c7b50a3eb33f5ba7251f88156a58b302)
