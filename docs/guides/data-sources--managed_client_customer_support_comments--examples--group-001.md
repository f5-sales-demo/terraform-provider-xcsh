---
page_title: "xcsh_managed_client_customer_support_comments examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_managed_client_customer_support_comments examples."
---

# xcsh_managed_client_customer_support_comments examples

<a id="canonical-3030003222020311-1211300121022203-2020231200103132-0123003222120122-0202313011010220-1311131120321212-0220300203013221-1311233302200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-0121323102230033-0320023001323221-1211130012203100-2333331223231001-1203112221121320-3132121331212101-3013023323130212-2212123012212223)
- Examples

<a id="canonical-2120131131230310-0121313202201010-0031330332230121-3131323023210020-0013302010123101-1013230111213300-2032001123221321-3121130201213301"></a>

### Complete configurations for `xcsh_managed_client_customer_support_comments`

- [Data source](data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-3100013300203120-1302101310131031-0112230210132103-3013322000120302-0231112013020300-2012331310222323-1333201110213001-1103010231112222): valid configuration.

<a id="canonical-3100013300203120-1302101310131031-0112230210132103-3013322000120302-0231112013020300-2012331310222323-1333201110213001-1103010231112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-0121323102230033-0320023001323221-1211130012203100-2333331223231001-1203112221121320-3132121331212101-3013023323130212-2212123012212223)
- [Examples](data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-3030003222020311-1211300121022203-2020231200103132-0123003222120122-0202313011010220-1311131120321212-0220300203013221-1311233302200111)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_managed_client_customer_support_comments/data-source.tf`; digest `sha256:12767ce24f59f5ccb7ea0d90d9e1667a91d868bfd5e02b49dd1c582b7f041c1d`.

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
