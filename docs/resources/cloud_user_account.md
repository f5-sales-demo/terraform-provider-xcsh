---
page_title: "xcsh_cloud_user_account landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account landing."
---

# xcsh_cloud_user_account landing

<a id="canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8ccf51ec88da9fc7f38acdb8da39b17dfde67592de66b3b6344985501d09c96"></a>

## xcsh_cloud_user_account — xcsh_cloud_user_account / eff10125a25e / 2

Breadcrumbs:

- xcsh_cloud_user_account

Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create
specifications. configuration.

<a id="canonical-037933227a397f7d354805060133e6d46104e39e7c528a0036af9eaadd449c7b"></a>

## Prerequisites — xcsh_cloud_user_account / eff10125a25e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6cce5e6d8a8e0db1f8ca06560aef69443f70672bb2bf5932e48e680f6a3e7f6c"></a>

## Minimal configuration — xcsh_cloud_user_account / eff10125a25e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudUserAccount Resource Example
# Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create specifications.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudUserAccount configuration
resource "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}
```

<a id="canonical-b6e493c4d066e7a74b594b1281892e59f6c6fbce03a5846cb258314f765b40c7"></a>

## Root configuration — xcsh_cloud_user_account / eff10125a25e / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-97086ab1d1783ac39d4249f2c1c3275e2e34592565fb3d71de843ad36b7ab564"></a>

## Next pages — xcsh_cloud_user_account / eff10125a25e / 6

- [Property reference](../guides/resources--cloud_user_account--reference--group-001.md#canonical-d89424604e5e126f68e3bddd26ea136e2f5b9887736c4d4b1504aa04f8961333)
- [Examples](../guides/resources--cloud_user_account--examples--group-001.md#canonical-611103a1d9bd66c091f2e2d67d84ad627dbbe360dec09fb8c6b48462156b5a71)
- [Import](../guides/resources--cloud_user_account--lifecycle--group-001.md#canonical-e16382fa1fb4eb786abdc789a5301f52d12fe0438a33ac41a001e2845fb2476b)
- [Timeouts](../guides/resources--cloud_user_account--lifecycle--group-001.md#canonical-c9d5b2496e14bf695f588b08386b385052726260f00df290318e6834b8ac0d63)
