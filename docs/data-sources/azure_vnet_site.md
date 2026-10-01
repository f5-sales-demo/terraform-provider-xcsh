---
page_title: "xcsh_azure_vnet_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site landing."
---

# xcsh_azure_vnet_site landing

<a id="canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232330000300310-0132220132121101-3321230013201220-1021013212230303-3313031333303232-3213013310220113-3013210232130302-1212311323210010"></a>

## xcsh_azure_vnet_site — xcsh_azure_vnet_site / 203013101232 / 2

Breadcrumbs:

- xcsh_azure_vnet_site

Manages a Azure VNET Site resource in F5 Distributed Cloud for deploying F5 sites within Azure
Virtual Network environments.

<a id="canonical-1231313113102001-1022333223123010-1310122313221312-0231020220200333-3010102301022320-2300313110012221-3312311321233001-1303023002331303"></a>

## Prerequisites — xcsh_azure_vnet_site / 203013101232 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: Azure authentication for deployment

<a id="canonical-2222303223332103-3131300303313303-0212000210123030-3120301322320110-3221201220123331-2123002323203203-1132323123213030-1102322111030202"></a>

## Minimal configuration — xcsh_azure_vnet_site / 203013101232 / 4

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

<a id="canonical-2230332320202102-2130322003301120-0302120020103311-0103210032332023-0022013233012133-3333232101301110-0332302231133031-2121202121020333"></a>

## Root configuration — xcsh_azure_vnet_site / 203013101232 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3211021020001133-0033333321321230-0202311020231011-2132003013210311-1220333001122213-3312010130113023-1131030323212231-1112122330130310"></a>

## Next pages — xcsh_azure_vnet_site / 203013101232 / 6

- [Property reference](../guides/data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [Examples](../guides/data-sources--azure_vnet_site--examples--group-001.md#canonical-3310212211210200-1000201300110221-1201100323202021-2101131332302213-3013231100220332-2303033311232213-0211013320200111-1222112023030002)
