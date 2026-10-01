---
page_title: "xcsh_partner_customer_support_comments landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_partner_customer_support_comments landing."
---

# xcsh_partner_customer_support_comments landing

<a id="canonical-52b42c42fce1130da3cf709e8f2675524c57b5e1f9ab3841fffd8054af69f0a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7750b48ab2eddce8be895f04252f17e5630b419a62c7185f4064da6a5e7b60c"></a>

## xcsh_partner_customer_support_comments — xcsh_partner_customer_support_comments / 94aa15b7640f / 2

Breadcrumbs:

- xcsh_partner_customer_support_comments

Resource retrieval operation.

<a id="canonical-f01bcceb4a625a27324e5720636dc8051a25db2581406ee10592c71b4199e6df"></a>

## Prerequisites — xcsh_partner_customer_support_comments / 94aa15b7640f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-90a00bbcb0c1b5022a3ca04b73d8d06a5edfdde9cab779fd0fc19c66a0235fe7"></a>

## Minimal configuration — xcsh_partner_customer_support_comments / 94aa15b7640f / 4

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

<a id="canonical-d020527cfa594332ddd7a36b840c7c5696122db44c1bbe232f9dc12885c5f27d"></a>

## Root configuration — xcsh_partner_customer_support_comments / 94aa15b7640f / 5

Required root properties: `tp_id`. Full root flags and choices appear in the property reference.

<a id="canonical-5eb225231f02b52ea60dedb46507e318dc723fb517726e59984c57fb167b2236"></a>

## Next pages — xcsh_partner_customer_support_comments / 94aa15b7640f / 6

- [Property reference](../guides/data-sources--partner_customer_support_comments--reference--group-001.md#canonical-2fa65522231cce83a94be0e590bceafb5ed95373feffa05567d61dea8ccd6ef6)
- [Examples](../guides/data-sources--partner_customer_support_comments--examples--group-001.md#canonical-0c592772aba34f4eca08e8e3c45c89ef32f6b5aed46eeddfc3751a68e5555237)
