---
page_title: "xcsh_sensitive_data_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy examples."
---

# xcsh_sensitive_data_policy examples

<a id="canonical-dc5a4d74e8b698d6a613694d8df01710e3f8a93a09db35642dad8f6088734943"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55e8004f83d19b4cfd06d61ab4ffa1b53a27b173eadb3b8b5434ad2812d05a8e"></a>

## Examples — Examples / 5a2907960a52 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)
- Examples

<a id="canonical-6043be62cb7b4f6d67b0adbaeae34857a2310077ad31d7da006b4c279ade72de"></a>

## Complete configurations — Examples / 5a2907960a52 / 3

- [Data source](data-sources--sensitive_data_policy--examples--group-001.md#canonical-c05c0124d466768b3c0e9176a7d1c23db1ed21dc0d24ea23b2d0e306cb8b31d1): valid configuration.

<a id="canonical-ebbbc8e45f8cfbb459092c921976a4a1fb1b8bc03e4c3f67d0ed37867fcf8366"></a>

## Next pages — Examples / 5a2907960a52 / 4

- [Data source](data-sources--sensitive_data_policy--examples--group-001.md#canonical-c05c0124d466768b3c0e9176a7d1c23db1ed21dc0d24ea23b2d0e306cb8b31d1)
- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)

<a id="canonical-c05c0124d466768b3c0e9176a7d1c23db1ed21dc0d24ea23b2d0e306cb8b31d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bf245b8d208e3220a204d1f269f39bf2c7b6dd82171fbcc13ea80e1e7276bb1"></a>

## Data source — Data source / 8f371a836151 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)
- [Examples](data-sources--sensitive_data_policy--examples--group-001.md#canonical-dc5a4d74e8b698d6a613694d8df01710e3f8a93a09db35642dad8f6088734943)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_sensitive_data_policy/data-source.tf`; digest `sha256:cfa788d1ad2d015969f47a23be03161a048878a194ca7cc1ede99b43ed21acb8`.

```terraform
# SensitiveDataPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SensitiveDataPolicy by name
data "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}

output "sensitive_data_policy_id" {
  value = data.xcsh_sensitive_data_policy.example.id
}
```

<a id="canonical-a2f5f48841ff63d47ccb85b1e20fe79a1709fd14603d399dba24d707dd280de3"></a>

## Next pages — Data source / 8f371a836151 / 3

- [Examples](data-sources--sensitive_data_policy--examples--group-001.md#canonical-dc5a4d74e8b698d6a613694d8df01710e3f8a93a09db35642dad8f6088734943)
- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20)
