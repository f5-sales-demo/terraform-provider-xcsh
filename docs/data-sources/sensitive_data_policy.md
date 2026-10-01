---
page_title: "xcsh_sensitive_data_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy landing."
---

# xcsh_sensitive_data_policy landing

<a id="canonical-b99cef409e9dd36f90bead99171d0e4c451e4223833ed625f792b1c61c5daa20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-709906cbe4a1fe6268b19fb7a33171016461cdb25077ef134f5a115806428e2b"></a>

## xcsh_sensitive_data_policy — xcsh_sensitive_data_policy / b523e818067f / 2

Breadcrumbs:

- xcsh_sensitive_data_policy

Manages sensitive\_data\_policy creates a new object in the storage backend for metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-16055801940fd64bc7ccf15e3d3cab43537213d782997b61134b6bde819fd263"></a>

## Prerequisites — xcsh_sensitive_data_policy / b523e818067f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-e2c53a6e6643e6586dc79ef1bf02e0d6fa52fc07a7589b4f5851157b494abbd0"></a>

## Minimal configuration — xcsh_sensitive_data_policy / b523e818067f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-72f2a5f9af268c1a7e8a44d96814a7bc13ea02f258a1146a445b0f93bb33cb98"></a>

## Root configuration — xcsh_sensitive_data_policy / b523e818067f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-edfbb7e8041ac080e72286e82fc51ac2b3897d8861ecc4330df2c3b6b100837f"></a>

## Next pages — xcsh_sensitive_data_policy / b523e818067f / 6

- [Property reference](../guides/data-sources--sensitive_data_policy--reference--group-001.md#canonical-86a29b3194d50ff67495d3075172b43382c2563aef377018db9add9b0f5caf8f)
- [Examples](../guides/data-sources--sensitive_data_policy--examples--group-001.md#canonical-dc5a4d74e8b698d6a613694d8df01710e3f8a93a09db35642dad8f6088734943)
