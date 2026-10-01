---
page_title: "xcsh_code_base_integration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration examples."
---

# xcsh_code_base_integration examples

<a id="canonical-ebd765c530052f7dfefad865c74b821bdddd2ce9c13805dcfc9c9bdea1b87240"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c18966854a9d7e9c337e9d5bb377353272f57e773183a4db9a9bba7f938ddf0"></a>

## Examples — Examples / 7393e081e926 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- Examples

<a id="canonical-276a0615d60d7e8aab4efca8644497587c3c24386edaea5a27020a54bcd1ba18"></a>

## Complete configurations — Examples / 7393e081e926 / 3

- [Resource](resources--code_base_integration--examples--group-001.md#canonical-3274333ea4bb0d2c926a97bacd7cde5cfef826713adccd77d145e217b0d33c94): valid configuration.

<a id="canonical-c2f5e7196dbd9ff78f1a29f8d77516bd005b2703fd6ed969ea07739302e5959f"></a>

## Next pages — Examples / 7393e081e926 / 4

- [Resource](resources--code_base_integration--examples--group-001.md#canonical-3274333ea4bb0d2c926a97bacd7cde5cfef826713adccd77d145e217b0d33c94)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)

<a id="canonical-3274333ea4bb0d2c926a97bacd7cde5cfef826713adccd77d145e217b0d33c94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5230034bdd844f4e186aa941efd0434997d9f7e38d88e43a3ca82adca9a87a76"></a>

## Resource — Resource / e20f71355eac / 2

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
- [Examples](resources--code_base_integration--examples--group-001.md#canonical-ebd765c530052f7dfefad865c74b821bdddd2ce9c13805dcfc9c9bdea1b87240)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_code_base_integration/resource.tf`; digest `sha256:96e41779f04da03e1e24b87e34c40ee5955713fd86df36344fb3cbec779ab18f`.

```terraform
# CodeBaseIntegration Resource Example
# Manages integration details in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CodeBaseIntegration configuration
resource "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}
```

<a id="canonical-c86123661d7cf470e04f025305dd7b20ab84d793be7bf1f82ea571388005a9e2"></a>

## Next pages — Resource / e20f71355eac / 3

- [Examples](resources--code_base_integration--examples--group-001.md#canonical-ebd765c530052f7dfefad865c74b821bdddd2ce9c13805dcfc9c9bdea1b87240)
- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-e465430d1245d444ce5a873302696b28704064b47a369e63e1c6a0047c8a5cc2)
