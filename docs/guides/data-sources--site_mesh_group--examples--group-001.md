---
page_title: "xcsh_site_mesh_group examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group examples."
---

# xcsh_site_mesh_group examples

<a id="canonical-d4c40650ecf86c580d3877f661b8ae24c251d4c0650148c0b1cc19fe8c3b4076"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-701ec09f3ba3e928b227d7374e34a9b76461e6b1020ad4be3ac565788f4d8acd"></a>

## Examples — Examples / 7afafa615eec / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- Examples

<a id="canonical-092632077c1be67c3155a828be84ec354eb3170641b43f7ac3b70b85883fe235"></a>

## Complete configurations — Examples / 7afafa615eec / 3

- [Data source](data-sources--site_mesh_group--examples--group-001.md#canonical-8cf85e0139894518a6989c9e4975524062e9d6cdaaff7d994bd09f03ec3444f1): valid configuration.

<a id="canonical-a15f493cfc65c553c782873379d4607247f84d0387827ac3e388d9d0b9e3aef6"></a>

## Next pages — Examples / 7afafa615eec / 4

- [Data source](data-sources--site_mesh_group--examples--group-001.md#canonical-8cf85e0139894518a6989c9e4975524062e9d6cdaaff7d994bd09f03ec3444f1)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)

<a id="canonical-8cf85e0139894518a6989c9e4975524062e9d6cdaaff7d994bd09f03ec3444f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74812cac83b13a74590b6ed9b63d41c6d33b01a4705eeb51755d885c1901bec7"></a>

## Data source — Data source / f073e48ec3c3 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
- [Examples](data-sources--site_mesh_group--examples--group-001.md#canonical-d4c40650ecf86c580d3877f661b8ae24c251d4c0650148c0b1cc19fe8c3b4076)
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

<a id="canonical-74f0027fab24966895db75e609a00481f39f6f65f14f3f464e9cec88db6df460"></a>

## Next pages — Data source / f073e48ec3c3 / 3

- [Examples](data-sources--site_mesh_group--examples--group-001.md#canonical-d4c40650ecf86c580d3877f661b8ae24c251d4c0650148c0b1cc19fe8c3b4076)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md#canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22)
