---
page_title: "xcsh_geo_location_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set landing."
---

# xcsh_geo_location_set landing

<a id="canonical-414009b5ef1ee0d7aeaab0bd5faaea0251b9f1a55924a8d8f2f05c764bd2b79a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dedcbf056f6b684361d339d35baa26adf518d911e529671fe7deafd7df18707f"></a>

## xcsh_geo_location_set — xcsh_geo_location_set / c85d91516540 / 2

Breadcrumbs:

- xcsh_geo_location_set

Manages Geolocation Set in F5 Distributed Cloud.

<a id="canonical-5b8697b28efe53f7638de7c73255281bce2fa1b666e4cb1727accfca5c492c60"></a>

## Prerequisites — xcsh_geo_location_set / c85d91516540 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a08eb8136a7c86695b67f51297f7f24757d59667e3ec0820e8a29f2b500513a0"></a>

## Minimal configuration — xcsh_geo_location_set / c85d91516540 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-7976021785d60bae34133357af2e9a8d3f0db771fba341093590401f9cad9099"></a>

## Root configuration — xcsh_geo_location_set / c85d91516540 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-f00a05efb77dd15ff87729919d0854632f88a194c3f00451f8bcd70686a86d0c"></a>

## Next pages — xcsh_geo_location_set / c85d91516540 / 6

- [Property reference](../guides/data-sources--geo_location_set--reference--group-001.md#canonical-5bf5901a30d1ac937b093d57aa613f518f7da14cd2317e1a7f4c81cd568647fc)
- [Examples](../guides/data-sources--geo_location_set--examples--group-001.md#canonical-8d74d833acd81355d521ac9d99256c89b15aa5b670f9266900846362a8dcc8d5)
