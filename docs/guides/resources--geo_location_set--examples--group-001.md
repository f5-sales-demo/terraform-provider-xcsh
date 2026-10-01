---
page_title: "xcsh_geo_location_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set examples."
---

# xcsh_geo_location_set examples

<a id="canonical-2bf875c901f9ce083510b2a48b22347952a1d45c6497a43a24374cffa0e1aec7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77446f4f5d68dff4a3aabcb2b056be0498cd23e6bf8643b47192c90c6b1de480"></a>

## Examples — Examples / 94fed3f3e179 / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
- Examples

<a id="canonical-9c9e71cf9b84997d27af908297727d8e9517138cd6b37938048982fd31a44033"></a>

## Complete configurations — Examples / 94fed3f3e179 / 3

- [Resource](resources--geo_location_set--examples--group-001.md#canonical-3c9d9bb56d812ed4a2db987bd9e82771bfa18d362733c7a8ec81f8bf7f09bb82): valid configuration.

<a id="canonical-cc84ca00f456cd436f9587b4bedc4235edcd9932594b474aaa29769b761c6fe2"></a>

## Next pages — Examples / 94fed3f3e179 / 4

- [Resource](resources--geo_location_set--examples--group-001.md#canonical-3c9d9bb56d812ed4a2db987bd9e82771bfa18d362733c7a8ec81f8bf7f09bb82)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)

<a id="canonical-3c9d9bb56d812ed4a2db987bd9e82771bfa18d362733c7a8ec81f8bf7f09bb82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8295367c84f54bd34603bc9511bc38f393bf1c013bb24a7a42077e3a8612e332"></a>

## Resource — Resource / 95809893b16f / 2

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
- [Examples](resources--geo_location_set--examples--group-001.md#canonical-2bf875c901f9ce083510b2a48b22347952a1d45c6497a43a24374cffa0e1aec7)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_geo_location_set/resource.tf`; digest `sha256:0e7cc4dbb2f138697448312f9d3840d2e253e947097376d3ac1a000ac5c54714`.

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

<a id="canonical-f1aa8b565259d2913d0db516d5e8bf2d641c23c162a74bab8bedd2492af4b97e"></a>

## Next pages — Resource / 95809893b16f / 3

- [Examples](resources--geo_location_set--examples--group-001.md#canonical-2bf875c901f9ce083510b2a48b22347952a1d45c6497a43a24374cffa0e1aec7)
- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-d316af30af83a2e4202cd927edfa772e4725c3cf87f11306f130cc11f9ddcd50)
