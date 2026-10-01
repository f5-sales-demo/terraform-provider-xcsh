---
page_title: "xcsh_managed_client_customer_support_comments landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_managed_client_customer_support_comments landing."
---

# xcsh_managed_client_customer_support_comments landing

<a id="canonical-19ed2b0f382c1ee9657068d0bff6bb41635a9678de67d991c72fb726a66c69ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fedec397fbfd87818cb1ec3f13325a6015e8af11587d6baeb06219efbf8a45c3"></a>

## xcsh_managed_client_customer_support_comments — xcsh_managed_client_customer_support_comments / a202023adcd6 / 2

Breadcrumbs:

- xcsh_managed_client_customer_support_comments

Resource retrieval operation.

<a id="canonical-ae7b842607eba8619d60c2c718fd362eaa744036a9c806382c445b6c6fc6e148"></a>

## Prerequisites — xcsh_managed_client_customer_support_comments / a202023adcd6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3e0bc956bdcf2f08033c1a000fe694f5e28e2af6a5c6dfaa4c5b33ac13488b50"></a>

## Minimal configuration — xcsh_managed_client_customer_support_comments / a202023adcd6 / 4

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

<a id="canonical-df89fafbd03aab44b5f80c06bf0167f355b592ca18711c7efe2084601a03a954"></a>

## Root configuration — xcsh_managed_client_customer_support_comments / a202023adcd6 / 5

Required root properties: `tp_id`. Full root flags and choices appear in the property reference.

<a id="canonical-cc40831c231177811b7293fea179210a970ce78987efaa61e235d0ef01f638f0"></a>

## Next pages — xcsh_managed_client_customer_support_comments / a202023adcd6 / 6

- [Property reference](../guides/data-sources--managed_client_customer_support_comments--reference--group-001.md#canonical-002a4575b61eb4ad6cacc307c2f712501721ed41a3e99b73773ff4b3e8095b91)
- [Examples](../guides/data-sources--managed_client_customer_support_comments--examples--group-001.md#canonical-cc0ea23565c192a388b604de1b0ea61a22dc512875758e6628c231e975bf2815)
