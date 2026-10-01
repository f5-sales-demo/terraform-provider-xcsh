---
page_title: "xcsh_cloud_region landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_region landing."
---

# xcsh_cloud_region landing

<a id="canonical-767677708e7c109e001971c5f50e0bf0785dfe911b57e8c6a2be2079978b8ef3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8314affd9d07d752c3f5362fe36d295b17633d7dcaf119c650b25b7278e5f47"></a>

## xcsh_cloud_region — xcsh_cloud_region / fe99690006cd / 2

Breadcrumbs:

- xcsh_cloud_region

Manages a Cloud Region resource in F5 Distributed Cloud for cloud re specification. configuration.
(read-only data source)

<a id="canonical-3187a6cec766d95a011db73374e3989fc8016b01ec8b8af93dd234f38863e52e"></a>

## Prerequisites — xcsh_cloud_region / fe99690006cd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-fab7bff4a89a963726c9860054423b72dabae1ef01d1a3768aed6324bb4dd414"></a>

## Minimal configuration — xcsh_cloud_region / fe99690006cd / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudRegion by name
data "xcsh_cloud_region" "example" {
  name      = "example-cloud-region"
  namespace = "staging"
}

output "cloud_region_id" {
  value = data.xcsh_cloud_region.example.id
}
```

<a id="canonical-53896097bc4b04ea60c159786d3e60ece25e3d387edc798dd285c9203c16c1f8"></a>

## Root configuration — xcsh_cloud_region / fe99690006cd / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-37551606af15a81c5cf42bf6c734b04b583c42874a5a18649bd9c8f12833bfc7"></a>

## Next pages — xcsh_cloud_region / fe99690006cd / 6

- [Property reference](../guides/data-sources--cloud_region--reference--group-001.md#canonical-5468418e9d8d11db3961c7d28bcb6294a1e394818903d378d0f5a42d91135dd8)
- [Examples](../guides/data-sources--cloud_region--examples--group-001.md#canonical-572733779c4c37f3d2f9db8e1bcb5b1a692dd7dc25e9fb0b77bee03110e2a828)
