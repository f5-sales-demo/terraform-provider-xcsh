---
page_title: "xcsh_managed_client_customer_support_comments examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_managed_client_customer_support_comments examples."
---

# xcsh_managed_client_customer_support_comments examples

<a id="canonical-cc0ea23565c192a388b604de1b0ea61a22dc512875758e6628c231e975bf2815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9875db3419de28440df3eb19ddecb90807c846d147b159f08e05ba79d97219f1"></a>

## Examples — Examples / 37b168f2f3e8 / 2

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)
- Examples

<a id="canonical-cc77de4749db47e482915d76c5d2ea12d289671a97620a5d2500e60d97dc5259"></a>

## Complete configurations — Examples / 37b168f2f3e8 / 3

- [Data source](data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-d01f08d87247474d16b24793c7e806322d58723086f74abb7f8549c15312d5aa): valid configuration.

<a id="canonical-7d29f540afad00f47df38a283d8393e1e69a90506fd22d74fc5325aa54eb3b27"></a>

## Next pages — Examples / 37b168f2f3e8 / 4

- [Data source](data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-d01f08d87247474d16b24793c7e806322d58723086f74abb7f8549c15312d5aa)
- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)

<a id="canonical-d01f08d87247474d16b24793c7e806322d58723086f74abb7f8549c15312d5aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-714fb07f0ca3222211b1658f0f845f34adc6f5fcef4c78295819491928f27239"></a>

## Data source — Data source / 12ea6c76f49f / 2

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)
- [Examples](data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-cc0ea23565c192a388b604de1b0ea61a22dc512875758e6628c231e975bf2815)
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

<a id="canonical-a2b0f5dc20cd7e72639ed1659f62ea0e83eaeccba2d500a7b9db0f90b2677e6b"></a>

## Next pages — Data source / 12ea6c76f49f / 3

- [Examples](data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-cc0ea23565c192a388b604de1b0ea61a22dc512875758e6628c231e975bf2815)
- [xcsh_managed_client_customer_support_comments](../data-sources/managed_client_customer_support_comments.md#canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab)
