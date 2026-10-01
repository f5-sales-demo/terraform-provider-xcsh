---
page_title: "xcsh_application_profiles examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles examples."
---

# xcsh_application_profiles examples

<a id="canonical-1997211f5f6bcf28b790e55429796b98903d009d88f6006d82ba263143c7649a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49927f316ea72ae6cd9f72b57706d34ed0ccd496fa3d04cd7280c4b9fef4863c"></a>

## Examples — Examples / 16a374a4f38b / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- Examples

<a id="canonical-0936a09d31d61adadf5a95839e4e0f59759c457f2934ac1099c97ef5e187fa1e"></a>

## Complete configurations — Examples / 16a374a4f38b / 3

- [Resource](resources--application_profiles--examples--group-001.md#canonical-38e55f683d2d065fc099d812b2ba3fcde96ecf6b921400045058904e8d79cc62): valid configuration.

<a id="canonical-f7f3d17f3efe767405faa2e6805e4893385568348dfa97e7ec4073ec705021d2"></a>

## Next pages — Examples / 16a374a4f38b / 4

- [Resource](resources--application_profiles--examples--group-001.md#canonical-38e55f683d2d065fc099d812b2ba3fcde96ecf6b921400045058904e8d79cc62)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)

<a id="canonical-38e55f683d2d065fc099d812b2ba3fcde96ecf6b921400045058904e8d79cc62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-305fce299822cf94b65412d0dcab0aad5cada880bcc46aeb0296620329a2633f"></a>

## Resource — Resource / 1b64ded414a9 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
- [Examples](resources--application_profiles--examples--group-001.md#canonical-1997211f5f6bcf28b790e55429796b98903d009d88f6006d82ba263143c7649a)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_application_profiles/resource.tf`; digest `sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f`.

```terraform
# ApplicationProfiles Resource Example
# Manages Application Profiles in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ApplicationProfiles configuration
resource "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}
```

<a id="canonical-dfaed4967cc825746be2acfacb79625f31c95b1924a36232ea870de0fa3d3307"></a>

## Next pages — Resource / 1b64ded414a9 / 3

- [Examples](resources--application_profiles--examples--group-001.md#canonical-1997211f5f6bcf28b790e55429796b98903d009d88f6006d82ba263143c7649a)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-00efa13c12be7b49cefa61c4d967fc1485e14838bc33c17bdcc66248df0d209c)
