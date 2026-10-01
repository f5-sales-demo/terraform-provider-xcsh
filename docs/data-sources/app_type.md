---
page_title: "xcsh_app_type landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type landing."
---

# xcsh_app_type landing

<a id="canonical-a05f527d2022e16c05ae64b94126a5cf0528b771352bcca895e3afb7446c725c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d26ce86fabab993eb59a3bbfd714fb51f8d5ca58b6c5325e156356df9318dbb6"></a>

## xcsh_app_type — xcsh_app_type / 95ec58c37eea / 2

Breadcrumbs:

- xcsh_app_type

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-f7b4ead8030991de0e8ab5658f13a0e7c56bdd105d3be8b342dca8b236920837"></a>

## Prerequisites — xcsh_app_type / 95ec58c37eea / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6c642a5b2366b2844032e0e11a9167233a5713a057ccb0af2be182acd6cf48e7"></a>

## Minimal configuration — xcsh_app_type / 95ec58c37eea / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppType by name
data "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}

output "app_type_id" {
  value = data.xcsh_app_type.example.id
}
```

<a id="canonical-4a3bffa0b3d4aec164522f2c780ed7c968ad3112dd216abdd9d06ef5120b7b0b"></a>

## Root configuration — xcsh_app_type / 95ec58c37eea / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c90cd5ae10357f6e7a1f82dd8ba1a1198573cc7ba304ed3465641b208ebfc7ec"></a>

## Next pages — xcsh_app_type / 95ec58c37eea / 6

- [Property reference](../guides/data-sources--app_type--reference--group-001.md#canonical-e94d90df6cbb06549a51c7d3e2057461f524b7c1d655589f838727813a6e5329)
- [Examples](../guides/data-sources--app_type--examples--group-001.md#canonical-92963a326d3f64840ee203fb8602b7ecd63f4db1fd0ee3492c5862d8f217a946)
