---
page_title: "xcsh_partner_customer_support_comments examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_partner_customer_support_comments examples."
---

# xcsh_partner_customer_support_comments examples

<a id="canonical-0030112102131302-2223220310331032-3022002032203203-3010113020213233-0302331223112232-3110123232313133-3003131101221220-3211111111020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md#canonical-1102231002301002-3330320101030031-2203303313002132-2033021213111102-1030111323113201-3321222303201001-3333333120001110-2233122133002213)
- Examples

<a id="canonical-2200333120101003-3112223121031312-0113022012301100-0120130010102231-1233312012013111-0120010322113220-0303113201131303-1110332202200332"></a>

### Complete configurations for `xcsh_partner_customer_support_comments`

- [Data source](data-sources--partner_customer_support_comments--examples--group-001.md#canonical-2112232030301011-0321211001202223-0133021320113300-1010110020032213-2000033130302233-2230303122323122-0110031303322312-1333010212210112): valid configuration.

<a id="canonical-2112232030301011-0321211001202223-0133021320113300-1010110020032213-2000033130302233-2230303122323122-0110031303322312-1333010212210112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md#canonical-1102231002301002-3330320101030031-2203303313002132-2033021213111102-1030111323113201-3321222303201001-3333333120001110-2233122133002213)
- [Examples](data-sources--partner_customer_support_comments--examples--group-001.md#canonical-0030112102131302-2223220310331032-3022002032203203-3010113020213233-0302331223112232-3110123232313133-3003131101221220-3211111111020313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_partner_customer_support_comments/data-source.tf`; digest `sha256:1a19f6d9101375158d6964da30975931b7c863facf609615fe0cc1c3633c733c`.

```terraform
# PartnerCustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_partner_customer_support_comments" "example" {
  tp_id = "example-value"
}

output "partner_customer_support_comments_result" {
  value = data.xcsh_partner_customer_support_comments.example
}
```
