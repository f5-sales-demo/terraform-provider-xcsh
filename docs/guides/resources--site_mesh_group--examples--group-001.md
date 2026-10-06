---
page_title: "xcsh_site_mesh_group examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group examples."
---

# xcsh_site_mesh_group examples

<a id="canonical-3323110021103110-2201303210222220-2213121010112022-3210303202230223-0003102131301022-3301121021333331-3323112133001111-3023122023131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- Examples

<a id="canonical-0313233332000200-2010121332110323-0333321231303010-1101300123001322-0100123113012100-0133333033110200-2132222031331301-1112221120021222"></a>

### Complete configurations for `xcsh_site_mesh_group`

- [Resource](resources--site_mesh_group--examples--group-001.md#canonical-2013331331320111-3100102220130032-1121301310333001-1322300330132123-2323313112222030-0123032021312123-0110130030112030-3322110313311303): valid configuration.

<a id="canonical-2013331331320111-3100102220130032-1121301310333001-1322300330132123-2323313112222030-0123032021312123-0110130030112030-3322110313311303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313)
- [Examples](resources--site_mesh_group--examples--group-001.md#canonical-3323110021103110-2201303210222220-2213121010112022-3210303202230223-0003102131301022-3301121021333331-3323112133001111-3023122023131322)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_site_mesh_group/resource.tf`; digest `sha256:9fd9c5f06a1b22b1d6bf2d57189b47e2ae5e2ce09701525907fff4345875b45e`.

```terraform
# SiteMeshGroup Resource Example
# Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SiteMeshGroup configuration
resource "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}
```
