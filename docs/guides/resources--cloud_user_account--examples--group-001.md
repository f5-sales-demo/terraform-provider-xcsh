---
page_title: "xcsh_cloud_user_account examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account examples."
---

# xcsh_cloud_user_account examples

<a id="canonical-611103a1d9bd66c091f2e2d67d84ad627dbbe360dec09fb8c6b48462156b5a71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7399660cce5a1c65124fbfd3d28f9a275e6d7293d2e664f71546ba32c0ab6a4"></a>

## Examples — Examples / 86db874e8df5 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- Examples

<a id="canonical-31c0d0acee94523a8a06011c4a913f9a4bf36a28a4327e741977c312e2be61ab"></a>

## Complete configurations — Examples / 86db874e8df5 / 3

- [Resource](resources--cloud_user_account--examples--group-001.md#canonical-f84a11f33bbb1fc7ca7e2263810cdbd250cb0d4c875dbfcf89cc29c95b11d304): valid configuration.

<a id="canonical-a7ca51157a3c1cb19c54440d9acbdfcdc94a411489e5b4a1f4d43625e0370b22"></a>

## Next pages — Examples / 86db874e8df5 / 4

- [Resource](resources--cloud_user_account--examples--group-001.md#canonical-f84a11f33bbb1fc7ca7e2263810cdbd250cb0d4c875dbfcf89cc29c95b11d304)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)

<a id="canonical-f84a11f33bbb1fc7ca7e2263810cdbd250cb0d4c875dbfcf89cc29c95b11d304"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-437e1c584acae1ceb785ce8281c847b5b601a73f98ecead3bd0409b6705fe97c"></a>

## Resource — Resource / 97d74dce25a8 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
- [Examples](resources--cloud_user_account--examples--group-001.md#canonical-611103a1d9bd66c091f2e2d67d84ad627dbbe360dec09fb8c6b48462156b5a71)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_user_account/resource.tf`; digest `sha256:10ccdaeaddaf605a3535a274963418653193279aea9eade41b140aeacaaa4d59`.

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

<a id="canonical-1606269cc21df3063af3963da41fb064b3d5afb54c5668d04ce8db40ffafb7a6"></a>

## Next pages — Resource / 97d74dce25a8 / 3

- [Examples](resources--cloud_user_account--examples--group-001.md#canonical-611103a1d9bd66c091f2e2d67d84ad627dbbe360dec09fb8c6b48462156b5a71)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-4d789962c073c5fba355a15dfa877ec397e9a0668d61350ddaf4e09c2085a479)
