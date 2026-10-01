---
page_title: "xcsh_app_api_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group examples."
---

# xcsh_app_api_group examples

<a id="canonical-3241f436996013f20af4a7b6f436914d3e43c74b62cf9d26ba502193fd04cc4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1170b1a9a0f64b9dbeab63e3165741223774743a42dfdfb61766f82c35b55f61"></a>

## Examples — Examples / c54e2b0c6b27 / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- Examples

<a id="canonical-f4559ebb708586cebf344b081458c76b4bd8f8c6cd01d35acfeeb4ce3ef17e9e"></a>

## Complete configurations — Examples / c54e2b0c6b27 / 3

- [Data source](data-sources--app_api_group--examples--group-001.md#canonical-f72ffb492a20f07ed3458e2dc93587ac556ba554f12d3b882c008da80d82a4f7): valid configuration.

<a id="canonical-969b0f2fe429036bcb2696a4029fba89529c5c4d5cd5a7889d92cbe2d6f168f9"></a>

## Next pages — Examples / c54e2b0c6b27 / 4

- [Data source](data-sources--app_api_group--examples--group-001.md#canonical-f72ffb492a20f07ed3458e2dc93587ac556ba554f12d3b882c008da80d82a4f7)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)

<a id="canonical-f72ffb492a20f07ed3458e2dc93587ac556ba554f12d3b882c008da80d82a4f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3433a6ea22d9c6fb78246a6e88e484cc774d35c923dc6d8050ccd035833d6744"></a>

## Data source — Data source / dc34ad54a3ca / 2

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
- [Examples](data-sources--app_api_group--examples--group-001.md#canonical-3241f436996013f20af4a7b6f436914d3e43c74b62cf9d26ba502193fd04cc4e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_api_group/data-source.tf`; digest `sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312`.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```

<a id="canonical-a7c0c4775b1aa4bda234593c35907477f12f1365d83a409732a4e19a3242dbfd"></a>

## Next pages — Data source / dc34ad54a3ca / 3

- [Examples](data-sources--app_api_group--examples--group-001.md#canonical-3241f436996013f20af4a7b6f436914d3e43c74b62cf9d26ba502193fd04cc4e)
- [xcsh_app_api_group](../data-sources/app_api_group.md#canonical-c5d129b096928f377cebf14f55dc9234d951dfe9603ea8c7996ce9f18b787644)
