---
page_title: "xcsh_site_mesh_group examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group examples."
---

# xcsh_site_mesh_group examples

<a id="canonical-3110301000121100-3230332012301120-0031032013133312-1201232022320210-3002110131103000-1211000110203000-2301303001213332-2030032310001312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- Examples

<a id="canonical-1300013230002133-0323220332210220-2302021331130313-1032031022212313-1210120132122301-0002002231102332-0322301112111320-2033103120223031"></a>

### Complete configurations for `xcsh_site_mesh_group`

- [Data source](data-sources--site_mesh_group--examples--group-001.md#canonical-2030332011320001-0321202110110120-2212212021302132-1021131111021000-1202322131123031-2222333313312121-1023310021330003-3230031010103301): valid configuration.

<a id="canonical-2030332011320001-0321202110110120-2212212021302132-1021131111021000-1202322131123031-2222333313312121-1023310021330003-3230031010103301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202)
- [Examples](data-sources--site_mesh_group--examples--group-001.md#canonical-3110301000121100-3230332012301120-0031032013133312-1201232022320210-3002110131103000-1211000110203000-2301303001213332-2030032310001312)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_mesh_group/data-source.tf`; digest `sha256:57e46715034cefba0a395123af3987e2b164522a8e7ac450c2dc1b9924e5f440`.

```terraform
# SiteMeshGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SiteMeshGroup by name
data "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}

output "site_mesh_group_id" {
  value = data.xcsh_site_mesh_group.example.id
}
```
