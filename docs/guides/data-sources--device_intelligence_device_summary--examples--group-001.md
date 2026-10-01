---
page_title: "xcsh_device_intelligence_device_summary examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_summary examples."
---

# xcsh_device_intelligence_device_summary examples

<a id="canonical-f15d9ca0379e4fc7bac628f18fdc6e1e7930d4b49dbab055002e7b17b00a53d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7271ea98faf158fb2bd9ea079e22fdf03c57b7da83209d3f872e7c396280c370"></a>

## Examples — Examples / 4a8d599c7743 / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)
- Examples

<a id="canonical-1332e308964f00c41025366eac027b50309212ab6e64d3027b743d402438ea27"></a>

## Complete configurations — Examples / 4a8d599c7743 / 3

- [Data source](data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-6f3e90309b3d49c8e8f3cfe7fa898bc5afe0dc3594ab814a6bddab1f28acf3bb): valid configuration.

<a id="canonical-58d44d1f3ef5fa36cd5d3ee1baa531786aad0b0196edd6c70771b915fa62c23b"></a>

## Next pages — Examples / 4a8d599c7743 / 4

- [Data source](data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-6f3e90309b3d49c8e8f3cfe7fa898bc5afe0dc3594ab814a6bddab1f28acf3bb)
- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)

<a id="canonical-6f3e90309b3d49c8e8f3cfe7fa898bc5afe0dc3594ab814a6bddab1f28acf3bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8c39ff479f4c736e6a8812c4ab644e03bc11e757f9732709f54eef426144741"></a>

## Data source — Data source / 0e3dc12409a4 / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)
- [Examples](data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-f15d9ca0379e4fc7bac628f18fdc6e1e7930d4b49dbab055002e7b17b00a53d7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_device_summary/data-source.tf`; digest `sha256:04010440b0365c5df9d1e3244c5a3f2f792550e1e47b47ef7b1ff17bae9b7077`.

```terraform
# DeviceIntelligenceDeviceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_summary" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_summary_result" {
  value = data.xcsh_device_intelligence_device_summary.example
}
```

<a id="canonical-2fa4ff9731bdd3077821789da58b3342593deef3b0bc578799ea74749dac4943"></a>

## Next pages — Data source / 0e3dc12409a4 / 3

- [Examples](data-sources--device_intelligence_device_summary--examples--group-001.md#canonical-f15d9ca0379e4fc7bac628f18fdc6e1e7930d4b49dbab055002e7b17b00a53d7)
- [xcsh_device_intelligence_device_summary](../data-sources/device_intelligence_device_summary.md#canonical-f8e0600c589deef40fd9b13736eefcb6123534aaac445791790673b2e06e3989)
