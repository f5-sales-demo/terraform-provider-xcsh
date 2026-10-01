---
page_title: "xcsh_site_mesh_group landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group landing."
---

# xcsh_site_mesh_group landing

<a id="canonical-0012330130320221-3012331110301122-0000031001123103-0200331031322133-2222010022030031-3331120010100133-2001130100013320-3123220212320202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130311101233122-1300120030313230-2222002011031133-0012322012221011-1013321202313311-3302302212200212-3131201331133313-0112022230210232"></a>

## xcsh_site_mesh_group — xcsh_site_mesh_group / 132302103222 / 2

Breadcrumbs:

- xcsh_site_mesh_group

Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

<a id="canonical-0230300303310333-0121203222201310-0211002213133123-0023311302023123-2103021021120010-2210101123202211-1121223012202313-2320123000022333"></a>

## Prerequisites — xcsh_site_mesh_group / 132302103222 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `site`.

- site: Sites to include in mesh connectivity

<a id="canonical-0113100332320101-1233101333110221-2033100102111320-0232130322032001-0201101230012033-3333321001201202-2013003332323001-3302321013001223"></a>

## Minimal configuration — xcsh_site_mesh_group / 132302103222 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2302202011131222-2321030313103232-0310031123212230-0231020101333103-0201220330023120-2103130112132322-2032231133202323-3212101303231013"></a>

## Root configuration — xcsh_site_mesh_group / 132302103222 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0003123302203323-3101022232301103-0031133000022110-2300023113231201-1300003320123121-3313220022331322-1210102032233100-2222032320300231"></a>

## Next pages — xcsh_site_mesh_group / 132302103222 / 6

- [Property reference](../guides/data-sources--site_mesh_group--reference--group-001.md#canonical-3021203322121012-3311333112312003-0213321330020231-1012033300112102-1313110113221131-3020203122221121-3231101122031301-0221001330310011)
- [Examples](../guides/data-sources--site_mesh_group--examples--group-001.md#canonical-3110301000121100-3230332012301120-0031032013133312-1201232022320210-3002110131103000-1211000110203000-2301303001213332-2030032310001312)
