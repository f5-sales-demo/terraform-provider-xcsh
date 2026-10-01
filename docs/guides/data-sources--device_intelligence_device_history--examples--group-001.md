---
page_title: "xcsh_device_intelligence_device_history examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_history examples."
---

# xcsh_device_intelligence_device_history examples

<a id="canonical-1951b16ac543cc0787ec0e6e895a4bba92897dd508464ba4197dac078a4fb569"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59d74cd910ca8b9b6c363b27a82475eda1ab178ce07238af3cd57ab7a7be89a1"></a>

## Examples — Examples / ab9ad64759c6 / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- Examples

<a id="canonical-668510ada7368cb2db23680ba685ea0c4404dce15be9921e9dd08f5d424c8b18"></a>

## Complete configurations — Examples / ab9ad64759c6 / 3

- [Data source](data-sources--device_intelligence_device_history--examples--group-001.md#canonical-5cbabc1c8257726aa1e7436916ddbd10e909cef6e30f0520033c3eeaf788a59a): valid configuration.

<a id="canonical-97a5c9bd1798f24891cd1b4e027e86ee6c6af67d055ace00babef7d7c1a259a8"></a>

## Next pages — Examples / ab9ad64759c6 / 4

- [Data source](data-sources--device_intelligence_device_history--examples--group-001.md#canonical-5cbabc1c8257726aa1e7436916ddbd10e909cef6e30f0520033c3eeaf788a59a)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)

<a id="canonical-5cbabc1c8257726aa1e7436916ddbd10e909cef6e30f0520033c3eeaf788a59a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3aaa9a419d43140cb38bf757c9938f1030ca80384cde194b60b7b4e128207f0"></a>

## Data source — Data source / b9e3ede3323c / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- [Examples](data-sources--device_intelligence_device_history--examples--group-001.md#canonical-1951b16ac543cc0787ec0e6e895a4bba92897dd508464ba4197dac078a4fb569)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_device_history/data-source.tf`; digest `sha256:f95f8328fecc18ee316fb7732f60f6fec73cc07388f5c6067f1bad91b8ac0d75`.

```terraform
# DeviceIntelligenceDeviceHistory DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_history" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_history_result" {
  value = data.xcsh_device_intelligence_device_history.example
}
```

<a id="canonical-7fba2f5f268c89c74a4bf6213099b707e5ad48b0e8b8eec69b2992ab97eb1705"></a>

## Next pages — Data source / b9e3ede3323c / 3

- [Examples](data-sources--device_intelligence_device_history--examples--group-001.md#canonical-1951b16ac543cc0787ec0e6e895a4bba92897dd508464ba4197dac078a4fb569)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
