---
page_title: "xcsh_device_intelligence_high_risk_transactions landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_high_risk_transactions landing."
---

# xcsh_device_intelligence_high_risk_transactions landing

<a id="canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9426558c8363eebed0fd2ea99d85818d513e96fc5e24ea9ffa5af2e7755807e"></a>

## xcsh_device_intelligence_high_risk_transactions — xcsh_device_intelligence_high_risk_transactions / ef0f0367a0f7 / 2

Breadcrumbs:

- xcsh_device_intelligence_high_risk_transactions

Resource creation operation.

<a id="canonical-a2a931497a93aca0b468039f2931a38096ec0b1aa6c98f46f6e4d7189cc7fc56"></a>

## Prerequisites — xcsh_device_intelligence_high_risk_transactions / ef0f0367a0f7 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f98a8000b0b5724c75f982b2ee69c55c4e7091486be0f947baa242af5a2eb831"></a>

## Minimal configuration — xcsh_device_intelligence_high_risk_transactions / ef0f0367a0f7 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-e786bd7a42f882c6b93690bd48dbac29daabffdc6b58ea6e53eb66a92d5932f8"></a>

## Root configuration — xcsh_device_intelligence_high_risk_transactions / ef0f0367a0f7 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8cba23b3979461c18b43ecef0046fe754beffd7e89f506021d1d6caa4fe03401"></a>

## Next pages — xcsh_device_intelligence_high_risk_transactions / ef0f0367a0f7 / 6

- [Property reference](../guides/data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea)
- [Examples](../guides/data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-d966d8a36af7cefa045e7fbb075fac0386a0f18bc0a3d6461456996da2f07cba)
