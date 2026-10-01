---
page_title: "xcsh_device_intelligence_high_risk_transactions examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_high_risk_transactions examples."
---

# xcsh_device_intelligence_high_risk_transactions examples

<a id="canonical-d966d8a36af7cefa045e7fbb075fac0386a0f18bc0a3d6461456996da2f07cba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05e48d57340ef74c09e38d229a8c25bd588329cebca3816c7e13638e4f42b945"></a>

## Examples — Examples / 88d2ee994ce0 / 2

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
- Examples

<a id="canonical-0ad5a2d9cc67871b4117f0e6568bb4107607f55a00945fba4126d09dddacb704"></a>

## Complete configurations — Examples / 88d2ee994ce0 / 3

- [Data source](data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-90497bf0c99d7a5b121c15675fcb8c611377025cb460aa9b9536391aab89d33e): valid configuration.

<a id="canonical-f039fc50969d3c6c23110a2f2c9a1c1d2c5fe93743319fd648a8c940eb5b31cc"></a>

## Next pages — Examples / 88d2ee994ce0 / 4

- [Data source](data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-90497bf0c99d7a5b121c15675fcb8c611377025cb460aa9b9536391aab89d33e)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)

<a id="canonical-90497bf0c99d7a5b121c15675fcb8c611377025cb460aa9b9536391aab89d33e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-535762d03fe6d57739a2bfd9fffe47ba2811b78d768e5b4f2f368be6ce7e19c8"></a>

## Data source — Data source / a948430ced39 / 2

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
- [Examples](data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-d966d8a36af7cefa045e7fbb075fac0386a0f18bc0a3d6461456996da2f07cba)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_high_risk_transactions/data-source.tf`; digest `sha256:5f7ef7da7ff2c9994ec1cbab645c6c6f2e697ac5efc56e49ae4d09a5e717abc6`.

```terraform
# DeviceIntelligenceHighRiskTransactions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_high_risk_transactions" "example" {
  namespace = "example-value"
}

output "device_intelligence_high_risk_transactions_result" {
  value = data.xcsh_device_intelligence_high_risk_transactions.example
}
```

<a id="canonical-09598a5f9ca016c9fb00048432169928d9bdcf383eee7890f0ea78e798babaa4"></a>

## Next pages — Data source / a948430ced39 / 3

- [Examples](data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-d966d8a36af7cefa045e7fbb075fac0386a0f18bc0a3d6461456996da2f07cba)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
