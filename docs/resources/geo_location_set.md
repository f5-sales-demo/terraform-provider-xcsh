---
page_title: "xcsh_geo_location_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set landing."
---

# xcsh_geo_location_set landing

<a id="canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e22c01d40f29e9956fbfa8f7a66e877f7f660487a4299a72bc2e889988acb371"></a>

## xcsh_geo_location_set — xcsh_geo_location_set / 095021391dfc / 2

Breadcrumbs:

- xcsh_geo_location_set

Manages Geolocation Set in F5 Distributed Cloud.

<a id="canonical-f9ce2b583515c20667e0402d15ca9b99a3ad82a364b20c255e65e7f6a50ca6f9"></a>

## Prerequisites — xcsh_geo_location_set / 095021391dfc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-aaa4b291b1694d445cd0456513fe36154c898510ed1be744c138d5d265e24087"></a>

## Minimal configuration — xcsh_geo_location_set / 095021391dfc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GeoLocationSet Resource Example
# Manages Geolocation Set in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GeoLocationSet configuration
resource "xcsh_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "system"
}
```

<a id="canonical-0ab723bf302d20845f49af873e0c0d7e747b218d2f12f12917fdb347d9877951"></a>

## Root configuration — xcsh_geo_location_set / 095021391dfc / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-a397cd7d6f89a701f9c14fb66c23408adb55a47580cdeea7c67762a92e800eb0"></a>

## Next pages — xcsh_geo_location_set / 095021391dfc / 6

- [Property reference](../guides/resources--geo_location_set--reference--group-001.md#canonical-5b59e1bdb82766035879753e3dec625b78e3fc0781f7694c35dffba49ee9e586)
- [Examples](../guides/resources--geo_location_set--examples--group-001.md#canonical-2bf875c901f9ce083510b2a48b22347952a1d45c6497a43a24374cffa0e1aec7)
- [Import](../guides/resources--geo_location_set--lifecycle--group-001.md#canonical-f139e01c4e7e5a4b8d5196be04abebb21c25f334db4454b4e5690699f15eeb02)
- [Timeouts](../guides/resources--geo_location_set--lifecycle--group-001.md#canonical-dbb2a2c4e1d0b862cda1cd04fe2157f06479a4f973bd37ba0d8892e2894af68e)
