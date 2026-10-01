---
page_title: "xcsh_azure_vnet_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site examples."
---

# xcsh_azure_vnet_site examples

<a id="canonical-f49a5920408705296143b8899177eca7c7b50a3eb33f5ba7251f88156a58b302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e2a5beb2eb15cda81800c542ce2a67983bcc49cf8ae12248c70f00075cf602b"></a>

## Examples — Examples / 00a76384fca7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- Examples

<a id="canonical-063330e0f00ad7c10389bf44365e2c5caa5df8b748f4a70e9448ca9a603f855e"></a>

## Complete configurations — Examples / 00a76384fca7 / 3

- [Data source](data-sources--azure_vnet_site--examples--group-001.md#canonical-770f33799f54422e53e122f620ce9d3aafb5d07eb82624a5f5fbc743b17b303e): valid configuration.

<a id="canonical-2e3da36af7af6a42beffce437ebebe7422c2c4f323f8b53f024f4632948a371c"></a>

## Next pages — Examples / 00a76384fca7 / 4

- [Data source](data-sources--azure_vnet_site--examples--group-001.md#canonical-770f33799f54422e53e122f620ce9d3aafb5d07eb82624a5f5fbc743b17b303e)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)

<a id="canonical-770f33799f54422e53e122f620ce9d3aafb5d07eb82624a5f5fbc743b17b303e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d0d736aa3d040a78a1043d92d55ba090ae2a93c51a18b53a5838ef00a7d5b13"></a>

## Data source — Data source / fb257cea03b0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
- [Examples](data-sources--azure_vnet_site--examples--group-001.md#canonical-f49a5920408705296143b8899177eca7c7b50a3eb33f5ba7251f88156a58b302)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_azure_vnet_site/data-source.tf`; digest `sha256:c010494f98bc90e6b2715fc8293de32922c6d80e43d822b5b3e3b2a72ebc3f82`.

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

<a id="canonical-55e32ce4023e2a2a8931bec3d07ef7e77bcf7107c27f2d8d1001656ef1a33a85"></a>

## Next pages — Data source / fb257cea03b0 / 3

- [Examples](data-sources--azure_vnet_site--examples--group-001.md#canonical-f49a5920408705296143b8899177eca7c7b50a3eb33f5ba7251f88156a58b302)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-7cb92499b57a54d5442d3ee07b93b09fafe4dea77fa6f1c73ebef54dbe104934)
