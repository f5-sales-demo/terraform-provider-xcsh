---
page_title: "xcsh_api_testing landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing landing."
---

# xcsh_api_testing landing

<a id="canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-416b00bca0043bee2cfe7c288622fc600dd3c8acb03a1c4f4f81ed54da81ca3b"></a>

## xcsh_api_testing — xcsh_api_testing / a469a93e7058 / 2

Breadcrumbs:

- xcsh_api_testing

Manages a API Testing resource in F5 Distributed Cloud.

<a id="canonical-a42d1ae33e96182b411049e05fd87b3dff46fe7913806c594d994fc4a31a15f7"></a>

## Prerequisites — xcsh_api_testing / a469a93e7058 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-13f85c4a0792f15c4483ecb124d2d8991c7d399bbff9c9520b7c7bd377f2a99a"></a>

## Minimal configuration — xcsh_api_testing / a469a93e7058 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```

<a id="canonical-6c974e6182a48f56856af23aa31f1818f0b42357956e5da2c0d50e797421d4e9"></a>

## Root configuration — xcsh_api_testing / a469a93e7058 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a118320b15e246861b8f26082c52e581abc90f6c5d40ad9f3310fec1ea3c5d59"></a>

## Next pages — xcsh_api_testing / a469a93e7058 / 6

- [Property reference](../guides/resources--api_testing--reference--group-001.md#canonical-7d169202115c4221c0ac36bfa940a8d8036faa89552670aeec6845e2d4f807bf)
- [Examples](../guides/resources--api_testing--examples--group-001.md#canonical-dabec0e7850cf899bd593e2fb91b65674444d08aa39d98a84e5660afac26f91a)
- [Import](../guides/resources--api_testing--lifecycle--group-001.md#canonical-7e2343b64b572e3da43b4817138a69302ab27188baabd023fc3fb102942d01c4)
- [Timeouts](../guides/resources--api_testing--lifecycle--group-001.md#canonical-a4c11db5c0662665788fada7373f5682c8861485f0143b83598be2883513f8f1)
