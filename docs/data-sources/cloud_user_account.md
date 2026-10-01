---
page_title: "xcsh_cloud_user_account landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account landing."
---

# xcsh_cloud_user_account landing

<a id="canonical-19961304f17d7e113e05ba91728d23383e280fe37853600cb3d5aced6a71c489"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7e1f0cc98149cffd37403b50419f1b9ca0a65685a5feb8db9b451fa6e33aba3"></a>

## xcsh_cloud_user_account — xcsh_cloud_user_account / 858ab116bffc / 2

Breadcrumbs:

- xcsh_cloud_user_account

Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create
specifications. configuration.

<a id="canonical-2a4296b09ccf54fb357a804a8b6f16037e81545fbb89a0fd2507271d989abedd"></a>

## Prerequisites — xcsh_cloud_user_account / 858ab116bffc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-7ed5fb76d09c9029acc85e70386978ceeecaeba4c65d9fc6d2a71707d44c16b2"></a>

## Minimal configuration — xcsh_cloud_user_account / 858ab116bffc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudUserAccount Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudUserAccount by name
data "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}

output "cloud_user_account_id" {
  value = data.xcsh_cloud_user_account.example.id
}
```

<a id="canonical-a503a31a64b95243fa9fa06cf924d9c1ef2a771cdf373b16a573e791dffc1987"></a>

## Root configuration — xcsh_cloud_user_account / 858ab116bffc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c8ef8a9a783e12f2b33f34e8f692db427f0f91e38170a2cef7f4c3fac0821ebc"></a>

## Next pages — xcsh_cloud_user_account / 858ab116bffc / 6

- [Property reference](../guides/data-sources--cloud_user_account--reference--group-001.md#canonical-900c84c9a5a9f196aca7fd76c191e49e38f64612a95e7e99265b923c4594e1dd)
- [Examples](../guides/data-sources--cloud_user_account--examples--group-001.md#canonical-5b28d2a6e4eaaa25e915ccb803d71dba9a9e9a9309e122177c605a960bd2d6f7)
