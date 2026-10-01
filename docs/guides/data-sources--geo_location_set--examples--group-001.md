---
page_title: "xcsh_geo_location_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set examples."
---

# xcsh_geo_location_set examples

<a id="canonical-8d74d833acd81355d521ac9d99256c89b15aa5b670f9266900846362a8dcc8d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e71626f4b0801c1dbe63829fe7b54bd7da41ed7daca6c5a45d2ac06dc0160e6d"></a>

## Examples — Examples / 21b76138f438 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)
- Examples

<a id="canonical-b181b1d0deb73e0e1b8458e9e3c63ad2aed7f4c024f06131d8ed5018a9073d62"></a>

## Complete configurations — Examples / 21b76138f438 / 3

- [Data source](data-sources--geo_location_set--examples--group-001.md#canonical-2c6e3523449abec82a2231e6357b0f9b459e6bb29c52ac8504ace955596d5d39): valid configuration.

<a id="canonical-3f9bf00725727c48a83e5d4fd5ec0c783994efee7ce569e527422a0b69f38b53"></a>

## Next pages — Examples / 21b76138f438 / 4

- [Data source](data-sources--geo_location_set--examples--group-001.md#canonical-2c6e3523449abec82a2231e6357b0f9b459e6bb29c52ac8504ace955596d5d39)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)

<a id="canonical-2c6e3523449abec82a2231e6357b0f9b459e6bb29c52ac8504ace955596d5d39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab8c5e4a8a561de4c8d4c62ff379fa9e3ba2f702dbb40ef9ceeafeaa8458077f"></a>

## Data source — Data source / 0191948d9986 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)
- [Examples](data-sources--geo_location_set--examples--group-001.md#canonical-8d74d833acd81355d521ac9d99256c89b15aa5b670f9266900846362a8dcc8d5)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_geo_location_set/data-source.tf`; digest `sha256:fe5ac2dcce0034284edb5717f6d3da18db76cd12233f57cc1a3e9be26885c3ea`.

```terraform
# GeoLocationSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GeoLocationSet by name
data "xcsh_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "system"
}

output "geo_location_set_id" {
  value = data.xcsh_geo_location_set.example.id
}
```

<a id="canonical-ed8c6664aa33af1e51b57c39e1a296e819929dd9c9c82f210cce20d7583018a6"></a>

## Next pages — Data source / 0191948d9986 / 3

- [Examples](data-sources--geo_location_set--examples--group-001.md#canonical-8d74d833acd81355d521ac9d99256c89b15aa5b670f9266900846362a8dcc8d5)
- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a)
