---
page_title: "xcsh_fleet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet examples."
---

# xcsh_fleet examples

<a id="canonical-6b4f913b2d40c1f26fc8e814243719a8906bcedb349bf38a8429fa1910f02b67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d1fe6d01d32d5acabf7a0453c9da7ee919b42c2e58ba60f332b400440fb0550"></a>

## Examples — Examples / c2db385392ba / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- Examples

<a id="canonical-4d337453006fb128c880c71225893fb785edd392564217d3870d2ab47d38b561"></a>

## Complete configurations — Examples / c2db385392ba / 3

- [Data source](data-sources--fleet--examples--group-001.md#canonical-7b1071edc3465befc4618c402c65e6ff7d17f08d654a7ee2b21bf20f8f1f866f): valid configuration.

<a id="canonical-4fb59f4a7ba5fb5c19aa5c0b48efcf84d6c4203f317645b6928e97c41b7f21f2"></a>

## Next pages — Examples / c2db385392ba / 4

- [Data source](data-sources--fleet--examples--group-001.md#canonical-7b1071edc3465befc4618c402c65e6ff7d17f08d654a7ee2b21bf20f8f1f866f)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)

<a id="canonical-7b1071edc3465befc4618c402c65e6ff7d17f08d654a7ee2b21bf20f8f1f866f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29a15567057aa433bb881f0e91c0e7aa683047d875d4fb4acf469dc4efc47eb7"></a>

## Data source — Data source / 0c80eee8bc56 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
- [Examples](data-sources--fleet--examples--group-001.md#canonical-6b4f913b2d40c1f26fc8e814243719a8906bcedb349bf38a8429fa1910f02b67)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fleet/data-source.tf`; digest `sha256:6d30c4d1a598a1f7f950e3ba739fc0d0cef6e7cc4d4b114700efeb50ec416c17`.

```terraform
# Fleet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Fleet by name
data "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"
}

output "fleet_id" {
  value = data.xcsh_fleet.example.id
}
```

<a id="canonical-404f62b9845a99f2c44f31e6c67b5c7a687decbde5eae5312194d110ef1a824d"></a>

## Next pages — Data source / 0c80eee8bc56 / 3

- [Examples](data-sources--fleet--examples--group-001.md#canonical-6b4f913b2d40c1f26fc8e814243719a8906bcedb349bf38a8429fa1910f02b67)
- [xcsh_fleet](../data-sources/fleet.md#canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a)
