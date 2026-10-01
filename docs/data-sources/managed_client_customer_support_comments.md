---
page_title: "xcsh_managed_client_customer_support_comments landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_managed_client_customer_support_comments landing."
---

# xcsh_managed_client_customer_support_comments landing

<a id="canonical-0121323102230033-0320023001323221-1211130012203100-2333331223231001-1203112221121320-3132121331212101-3013023323130212-2212123012212223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332313230032113-3323333120132001-2030230132300333-0103030211221200-0111322022330101-1120133112232232-2300120201213233-2333202210113003"></a>

## xcsh_managed_client_customer_support_comments — xcsh_managed_client_customer_support_comments / 032231303112 / 2

Breadcrumbs:

- xcsh_managed_client_customer_support_comments

Resource retrieval operation.

<a id="canonical-2232132320100212-0013322322201201-2131120030023013-0120333103120232-2222131010000312-2221302000120320-0230101011231230-1233301232011020"></a>

## Prerequisites — xcsh_managed_client_customer_support_comments / 032231303112 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0332002330211112-2331303302330020-0003033001220000-0033321221103311-3202203202223312-2211301231332222-1030112303032230-0103102020231100"></a>

## Minimal configuration — xcsh_managed_client_customer_support_comments / 032231303112 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ManagedClientCustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_managed_client_customer_support_comments" "example" {
  tp_id = "example-value"
}

output "managed_client_customer_support_comments_result" {
  value = data.xcsh_managed_client_customer_support_comments.example
}
```

<a id="canonical-3133202133223323-3100032222231010-2311332000300012-2333000112133303-1111231121023022-0120130101301332-3332020020101200-0122000322211110"></a>

## Root configuration — xcsh_managed_client_customer_support_comments / 032231303112 / 5

Required root properties: `tp_id`. Full root flags and choices appear in the property reference.

<a id="canonical-3030100020030130-0203010113132001-0123130221033332-2201132102010022-2113003032132021-2013323322221201-3202031131003233-0001331203203300"></a>

## Next pages — xcsh_managed_client_customer_support_comments / 032231303112 / 6

- [Property reference](../guides/data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-0000022210111311-2312013223102231-1230223030030013-3002331301021100-0113020132311001-2203322121231303-1313033333102303-3220002111232101)
- [Examples](../guides/data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-3030003222020311-1211300121022203-2020231200103132-0123003222120122-0202313011010220-1311131120321212-0220300203013221-1311233302200111)
