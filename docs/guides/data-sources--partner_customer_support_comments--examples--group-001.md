---
page_title: "xcsh_partner_customer_support_comments examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_partner_customer_support_comments examples."
---

# xcsh_partner_customer_support_comments examples

<a id="canonical-0c592772aba34f4eca08e8e3c45c89ef32f6b5aed46eeddfc3751a68e5555237"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0fd8443d6ad937617286c50187044ad6fd861d51813a5e8335e177354fa283e"></a>

## Examples — Examples / 5838465bd084 / 2

Breadcrumbs:

- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md#canonical-52b42c42fce1130da3cf709e8f2675524c57b5e1f9ab3841fffd8054af69f0a7)
- Examples

<a id="canonical-805eb4504b99570e00232cc60ff0c1ed19be77de1d98405d860de2ba915e1b42"></a>

## Complete configurations — Examples / 5838465bd084 / 3

- [Data source](data-sources--partner_customer_support_comments--examples--group-001.md#canonical-96b8cc45399418ab1f2785f0445083a7803dccafaccdaeda14373eb67f126916): valid configuration.

<a id="canonical-0009f5c772e206953def6992f2f710fe963a12c4a84050bccfa357854afe9740"></a>

## Next pages — Examples / 5838465bd084 / 4

- [Data source](data-sources--partner_customer_support_comments--examples--group-001.md#canonical-96b8cc45399418ab1f2785f0445083a7803dccafaccdaeda14373eb67f126916)
- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md#canonical-52b42c42fce1130da3cf709e8f2675524c57b5e1f9ab3841fffd8054af69f0a7)

<a id="canonical-96b8cc45399418ab1f2785f0445083a7803dccafaccdaeda14373eb67f126916"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f49ace9b5b1e859481a5ce203e34291791f1612cf5dd7d9191fd9c148e679673"></a>

## Data source — Data source / 0a9192279089 / 2

Breadcrumbs:

- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md#canonical-52b42c42fce1130da3cf709e8f2675524c57b5e1f9ab3841fffd8054af69f0a7)
- [Examples](data-sources--partner_customer_support_comments--examples--group-001.md#canonical-0c592772aba34f4eca08e8e3c45c89ef32f6b5aed46eeddfc3751a68e5555237)
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

<a id="canonical-4682964892ec6110cb0f7a2668ad5a8f3e3d3254093a9eefdf7a15ead3591a64"></a>

## Next pages — Data source / 0a9192279089 / 3

- [Examples](data-sources--partner_customer_support_comments--examples--group-001.md#canonical-0c592772aba34f4eca08e8e3c45c89ef32f6b5aed46eeddfc3751a68e5555237)
- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md#canonical-52b42c42fce1130da3cf709e8f2675524c57b5e1f9ab3841fffd8054af69f0a7)
