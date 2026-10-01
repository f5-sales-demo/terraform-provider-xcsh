---
page_title: "xcsh_cloud_credentials examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials examples."
---

# xcsh_cloud_credentials examples

<a id="canonical-d82053d4cca79c9fc1875cc320abd0f9f5cfa38701c92ce1d01723b73fef52c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63c1712f5549f60ccb4d4bff48adabf0af44709aa8a68debaac853dfba1619fb"></a>

## Examples — Examples / 9ecd430cfe03 / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- Examples

<a id="canonical-9c856032ea3fc1717364ad647f81f4a82a09b2cddc34bf9a0c9c19e05e3f5b40"></a>

## Complete configurations — Examples / 9ecd430cfe03 / 3

- [Resource](resources--cloud_credentials--examples--group-001.md#canonical-75c9ba697831fe60b6467afd6c905dde54b4dd211f30c0d2a684609020bd03b0): valid configuration.

<a id="canonical-4d86798b2770a0b961fe302edf7871d73ff08ba8230dc49519eead2c96ec2221"></a>

## Next pages — Examples / 9ecd430cfe03 / 4

- [Resource](resources--cloud_credentials--examples--group-001.md#canonical-75c9ba697831fe60b6467afd6c905dde54b4dd211f30c0d2a684609020bd03b0)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)

<a id="canonical-75c9ba697831fe60b6467afd6c905dde54b4dd211f30c0d2a684609020bd03b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29c5b60908927b7c28b053a6c0427581e23a42bc1aa0276f3e030e7f98b66ac9"></a>

## Resource — Resource / 926bee69a23a / 2

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
- [Examples](resources--cloud_credentials--examples--group-001.md#canonical-d82053d4cca79c9fc1875cc320abd0f9f5cfa38701c92ce1d01723b73fef52c9)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_credentials/resource.tf`; digest `sha256:59ca9345886180c4c5420989ec1dac0bd8bd93a36d37acaaca6a81918a8c91f1`.

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

<a id="canonical-f7a9ba16e33b858a996f7e3c5f8b6f95830e149cd03645a48f9f431f26284888"></a>

## Next pages — Resource / 926bee69a23a / 3

- [Examples](resources--cloud_credentials--examples--group-001.md#canonical-d82053d4cca79c9fc1875cc320abd0f9f5cfa38701c92ce1d01723b73fef52c9)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-1cb4bf6ac71d4f09833d90c0cb47c40c59691818c5c6c69f8111a26572dc4e18)
