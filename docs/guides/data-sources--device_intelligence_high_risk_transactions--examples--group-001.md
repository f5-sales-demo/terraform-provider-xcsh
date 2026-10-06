---
page_title: "xcsh_device_intelligence_high_risk_transactions examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_high_risk_transactions examples."
---

# xcsh_device_intelligence_high_risk_transactions examples

<a id="canonical-3121121231202203-1222331330323322-0010113213332323-0013113322300003-2012220033012023-3000220331121012-0110111221211231-2202330013302322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103)
- Examples

<a id="canonical-0011321020311113-0310003233131030-0021320320310202-2122203002112331-1120200302213032-2330220320011230-1332010312032032-1033100223211011"></a>

### Complete configurations for `xcsh_device_intelligence_high_risk_transactions`

- [Data source](data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-2100102113233300-3021213113221123-0102013001111213-1133302320301201-0103131300021130-2310120022222123-2111031203210122-2223202131030332): valid configuration.

<a id="canonical-2100102113233300-3021213113221123-0102013001111213-1133302320301201-0103131300021130-2310120022222123-2111031203210122-2223202131030332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103)
- [Examples](data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-3121121231202203-1222331330323322-0010113213332323-0013113322300003-2012220033012023-3000220331121012-0110111221211231-2202330013302322)
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
