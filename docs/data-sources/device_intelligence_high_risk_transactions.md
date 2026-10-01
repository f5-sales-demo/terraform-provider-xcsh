---
page_title: "xcsh_device_intelligence_high_risk_transactions landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_high_risk_transactions landing."
---

# xcsh_device_intelligence_high_risk_transactions landing

<a id="canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321100212111120-3020031203323223-3231003331023222-2121312011200120-3111010332211233-3011320210322221-3333221122330232-1313111120001332"></a>

## xcsh_device_intelligence_high_risk_transactions — xcsh_device_intelligence_high_risk_transactions / 121322003313 / 2

Breadcrumbs:

- xcsh_device_intelligence_high_risk_transactions

Resource creation operation.

<a id="canonical-2202222103011021-1322210322302200-2310122000032133-0221030122032000-2112323000230122-2212302120331012-3312321031130120-2130301333301112"></a>

## Prerequisites — xcsh_device_intelligence_high_risk_transactions / 121322003313 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3321202220000000-2300231113021030-1311332120022302-3232122130111130-1032130021011020-1223320033211013-2322220210022233-1122023223200301"></a>

## Minimal configuration — xcsh_device_intelligence_high_risk_transactions / 121322003313 / 4

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

<a id="canonical-3213201223311322-1002332020023012-2321031221002331-1020312322300221-3122222333333130-1223112032221232-1103322312122221-0231112103023320"></a>

## Root configuration — xcsh_device_intelligence_high_risk_transactions / 121322003313 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2030232202032303-2113211012013001-2023100332303233-0000101233321311-1023323333311332-2021331100120002-0131013112302222-1033320003100001"></a>

## Next pages — xcsh_device_intelligence_high_risk_transactions / 121322003313 / 6

- [Property reference](../guides/data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2030120313301233-0221233133100221-0201300032331331-1120303000311221-2203033322111232-1001320313023121-2003030130133011-0112310120123222)
- [Examples](../guides/data-sources--device_intelligence_high_risk_transactions--examples--group-001.md#canonical-3121121231202203-1222331330323322-0010113213332323-0013113322300003-2012220033012023-3000220331121012-0110111221211231-2202330013302322)
