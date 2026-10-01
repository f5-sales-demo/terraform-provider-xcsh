---
page_title: "xcsh_cloud_credentials landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials landing."
---

# xcsh_cloud_credentials landing

<a id="canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4622bee4ffa63c0e2b9a22e991cc12fd762a4f9d86d30d414c6039e26377bad"></a>

## xcsh_cloud_credentials — xcsh_cloud_credentials / b4030bb86004 / 2

Breadcrumbs:

- xcsh_cloud_credentials

Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud\_credentials
object. configuration.

<a id="canonical-ad4c464447985ee85b6896670d5b07609fc2ed66dfc86c8ef4c8458619c5512f"></a>

## Prerequisites — xcsh_cloud_credentials / b4030bb86004 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-5d4eebd3c2fa5a6c82813cc9746c56bbe92cc0d618beab5161af8514c3f057c5"></a>

## Minimal configuration — xcsh_cloud_credentials / b4030bb86004 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudCredentials Resource Example
# Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudCredentials configuration
resource "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}
```

<a id="canonical-4cb29cd0819ece2cfa1025815b683f80724d7a5ffed103a74683c260e607f77c"></a>

## Root configuration — xcsh_cloud_credentials / b4030bb86004 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-19ac13d7d0e9ea5a9fa0bcb47799f2bda406f8d2913267da2c64cd2e1fb003a3"></a>

## Next pages — xcsh_cloud_credentials / b4030bb86004 / 6

- [Property reference](../guides/resources--cloud_credentials--reference--group-001.md#canonical-e19ca078f045646b07b7307d630d1eebf178fe6edffad4c6ad87aa0a772b6a20)
- [Examples](../guides/resources--cloud_credentials--examples--group-001.md#canonical-d82053d4cca79c9fc1875cc320abd0f9f5cfa38701c92ce1d01723b73fef52c9)
- [Import](../guides/resources--cloud_credentials--lifecycle--group-001.md#canonical-5059158635425194e905f4659a04928b2ebe3549fb964d8e38cd8a883a995dfc)
- [Timeouts](../guides/resources--cloud_credentials--lifecycle--group-001.md#canonical-4526d993c8c8dbd15d278331185bcc46f36ad1fb2ac5b0f2b93ca35d0bdc7336)
