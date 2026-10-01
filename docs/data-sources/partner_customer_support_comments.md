---
page_title: "xcsh_partner_customer_support_comments landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_partner_customer_support_comments landing."
---

# xcsh_partner_customer_support_comments landing

<a id="canonical-1102231002301002-3330320101030031-2203303313002132-2033021213111102-1030111323113201-3321222303201001-3333333120001110-2233122133002213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113131100231020-2223023231313032-2023322021113300-1002110233011332-1112030023100121-2212023013012011-3310001210312212-2211321323120030"></a>

## xcsh_partner_customer_support_comments — xcsh_partner_customer_support_comments / 231312100033 / 2

Breadcrumbs:

- xcsh_partner_customer_support_comments

Resource retrieval operation.

<a id="canonical-3300012330303223-1022120211220213-0302103211130200-1203123130200011-0122021131230211-2001100012323201-0011210230130123-1001212132123133"></a>

## Prerequisites — xcsh_partner_customer_support_comments / 231312100033 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2100220000232330-2300300123110002-0222033022001023-1303312031001222-1132313331313221-3022231313213331-0033300121301212-2200020311333213"></a>

## Minimal configuration — xcsh_partner_customer_support_comments / 231312100033 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3100020011021330-3322112110030302-3131311322031223-2010003013301112-2112010202312310-1030012323320203-0233213130010220-2011301133021331"></a>

## Root configuration — xcsh_partner_customer_support_comments / 231312100033 / 5

Required root properties: `tp_id`. Full root flags and choices appear in the property reference.

<a id="canonical-1132230202110203-0133000223110232-2212003132312310-1211001332030120-3130130203332311-0113130212321121-2120103011133323-0112132302020312"></a>

## Next pages — xcsh_partner_customer_support_comments / 231312100033 / 6

- [Property reference](../guides/data-sources--partner_customer_support_comments--reference--group-001.md#canonical-0233221211110202-0203013030322003-2221102332003211-2100233032223323-1132312111031303-3332333322001111-1213311201313222-2030303112323312)
- [Examples](../guides/data-sources--partner_customer_support_comments--examples--group-001.md#canonical-0030112102131302-2223220310331032-3022002032203203-3010113020213233-0302331223112232-3110123232313133-3003131101221220-3211111111020313)
