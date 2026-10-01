---
page_title: "xcsh_api_testing landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing landing."
---

# xcsh_api_testing landing

<a id="canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f4b50ad267767ab36bd03366d1bf06887c511b38fed96331282b993c00e1b8b"></a>

## xcsh_api_testing — xcsh_api_testing / 992b39eaacad / 2

Breadcrumbs:

- xcsh_api_testing

Manages a API Testing resource in F5 Distributed Cloud.

<a id="canonical-01f5cc44a635fb3eb959b1cc246d7479924d5acae1efda5bec03c8ef97126c77"></a>

## Prerequisites — xcsh_api_testing / 992b39eaacad / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ec189eb75517e0dd1ffb27cd2b4d2910e17d42e21ecc0e2c3f3fe061aa6e266d"></a>

## Minimal configuration — xcsh_api_testing / 992b39eaacad / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APITesting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APITesting by name
data "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}

output "api_testing_id" {
  value = data.xcsh_api_testing.example.id
}
```

<a id="canonical-25fd128079eb3e40535213be4f296203da351edf0be17e762da60a11eccfcebe"></a>

## Root configuration — xcsh_api_testing / 992b39eaacad / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-5c560038b08d546355163309f07b76cb637e42add5f3fc56992c8e55fe494e53"></a>

## Next pages — xcsh_api_testing / 992b39eaacad / 6

- [Property reference](../guides/data-sources--api_testing--reference--group-001.md#canonical-c29fb04ff0e7f10a5a7e03af6728d874d4b9396f13a83d733b7df848936cd493)
- [Examples](../guides/data-sources--api_testing--examples--group-001.md#canonical-beb504e58ed4f5cdc2da1cb373b5d71c3498aaee616e4b912c35bcb206ab55b3)
