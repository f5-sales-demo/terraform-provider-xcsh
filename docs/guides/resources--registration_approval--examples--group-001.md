---
page_title: "xcsh_registration_approval examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration_approval examples."
---

# xcsh_registration_approval examples

<a id="canonical-18983f6c2d60c35fbc6941ad7957980b8a60e933dc081010b07873198b27c179"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b7218e20adeb440b8ea609276db462a6d5e301bf40e7ae8f17c77b330f2b96e"></a>

## Examples — Examples / 1b53ab0597b1 / 2

Breadcrumbs:

- [xcsh_registration_approval](../resources/registration_approval.md#canonical-7d4fb674027240378875a2202f657eddfb167f6c0f01b29cddfd793bc70315ba)
- Examples

<a id="canonical-c190f12e646e5d9161af2a1130b46678b3f7fee031f1ab4d582c5d21af30c54c"></a>

## Complete configurations — Examples / 1b53ab0597b1 / 3

- [Resource](resources--registration_approval--examples--group-001.md#canonical-4d52da7f23c249952bf4f42350e0dfff6a34d3c3cf706bfa2bae70cb1e7fce93): valid configuration.

<a id="canonical-d341b07ac3c351f9289535bec9610e2937a32f44f1bf4a28a6873bec6034f5ba"></a>

## Next pages — Examples / 1b53ab0597b1 / 4

- [Resource](resources--registration_approval--examples--group-001.md#canonical-4d52da7f23c249952bf4f42350e0dfff6a34d3c3cf706bfa2bae70cb1e7fce93)
- [xcsh_registration_approval](../resources/registration_approval.md#canonical-7d4fb674027240378875a2202f657eddfb167f6c0f01b29cddfd793bc70315ba)

<a id="canonical-4d52da7f23c249952bf4f42350e0dfff6a34d3c3cf706bfa2bae70cb1e7fce93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-542c7823396ca3b804aba15cf484fb6393f92fbbd1e13032ee4155e24d9dd5be"></a>

## Resource — Resource / 826d3d05356d / 2

Breadcrumbs:

- [xcsh_registration_approval](../resources/registration_approval.md#canonical-7d4fb674027240378875a2202f657eddfb167f6c0f01b29cddfd793bc70315ba)
- [Examples](resources--registration_approval--examples--group-001.md#canonical-18983f6c2d60c35fbc6941ad7957980b8a60e933dc081010b07873198b27c179)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_registration_approval/resource.tf`; digest `sha256:5a69b599688eff851094350bc6a6e7fffd81e09b55087a523abb0c8c970908cd`.

```terraform
# RegistrationApproval Resource Example
# Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RegistrationApproval configuration
resource "xcsh_registration_approval" "example" {
  name      = "example-registration-approval"
  namespace = "staging"

  cluster_size = 1
}
```

<a id="canonical-dae0f964a3586dd5c15eec97a9571e96373de5d30234ae3c41a8ab124dfa1939"></a>

## Next pages — Resource / 826d3d05356d / 3

- [Examples](resources--registration_approval--examples--group-001.md#canonical-18983f6c2d60c35fbc6941ad7957980b8a60e933dc081010b07873198b27c179)
- [xcsh_registration_approval](../resources/registration_approval.md#canonical-7d4fb674027240378875a2202f657eddfb167f6c0f01b29cddfd793bc70315ba)
