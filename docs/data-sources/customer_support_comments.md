---
page_title: "xcsh_customer_support_comments landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_customer_support_comments landing."
---

# xcsh_customer_support_comments landing

<a id="canonical-eff44c2302047abc21725ce2ae21bd0da5e0bf5c0fbcb6b76347af7cf962e3b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9b4886599c21cb9f65348319d3508453cb4e80efff721f20e2b9b966614d9d2"></a>

## xcsh_customer_support_comments — xcsh_customer_support_comments / 8a6e7f6dcfc5 / 2

Breadcrumbs:

- xcsh_customer_support_comments

Resource retrieval operation.

<a id="canonical-a06a701c36bfe76d19c15d591f4abfb942a1a17e8da8c9beadf45b2b0be3107b"></a>

## Prerequisites — xcsh_customer_support_comments / 8a6e7f6dcfc5 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-73cf6594abcba08031688de8a3f55d281cb61137857803dfd489aa18838b68dd"></a>

## Minimal configuration — xcsh_customer_support_comments / 8a6e7f6dcfc5 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_customer_support_comments" "example" {
  name = "example-value"
}

output "customer_support_comments_result" {
  value = data.xcsh_customer_support_comments.example
}
```

<a id="canonical-f2db09583475c9ec3cce71f6e634d352b0f191ab9770aeb8b86819b94f7f9086"></a>

## Root configuration — xcsh_customer_support_comments / 8a6e7f6dcfc5 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-867b6ca781e8e301da004506d3c13f08834cb5d2d41523afd10aec859df2bca2"></a>

## Next pages — xcsh_customer_support_comments / 8a6e7f6dcfc5 / 6

- [Property reference](../guides/data-sources--customer_support_comments--reference--group-001.md#canonical-12e9887903e3773a47890b40be28cd522622da1fd04215c1250c156eea0f4f47)
- [Examples](../guides/data-sources--customer_support_comments--examples--group-001.md#canonical-1c092da3a2fc15c8a044fa74e96c07cccf0a472cacfd51c5a35a965d897c73bd)
