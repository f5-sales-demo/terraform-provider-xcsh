---
page_title: "xcsh_app_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type examples."
---

# xcsh_app_type examples

<a id="canonical-f7b9298394e83b8e4a094f5b90535372a175dda62f0fa6b17e5bb40ccf6b6b6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-890b19d7f8d9bc5a47302338f4f0310f086df2a5aae2efdf06a083f4f38a712e"></a>

## Examples — Examples / 51a23d035d43 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- Examples

<a id="canonical-890c4e033bfc950437623812b76ff19344d950c1997811f436b5941f74f462d1"></a>

## Complete configurations — Examples / 51a23d035d43 / 3

- [Resource](resources--app_type--examples--group-001.md#canonical-6bea900ea9f2f5ac4959f94ac5c863b3f788e510df870a3b3514e1f53d3bc561): valid configuration.

<a id="canonical-97c49523117973286ddd6b63ebd034656fef3bd57a215949ee22c816a50f2cec"></a>

## Next pages — Examples / 51a23d035d43 / 4

- [Resource](resources--app_type--examples--group-001.md#canonical-6bea900ea9f2f5ac4959f94ac5c863b3f788e510df870a3b3514e1f53d3bc561)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)

<a id="canonical-6bea900ea9f2f5ac4959f94ac5c863b3f788e510df870a3b3514e1f53d3bc561"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c785210940bb4c4c212cbf7163219a662783f60fedde1c7c48ca9e93ac0bd41"></a>

## Resource — Resource / 4ae38d1ab341 / 2

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
- [Examples](resources--app_type--examples--group-001.md#canonical-f7b9298394e83b8e4a094f5b90535372a175dda62f0fa6b17e5bb40ccf6b6b6d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_type/resource.tf`; digest `sha256:9988ab36f101ce764d3f26103a86f6c75f3dcea74d272aeebe958208da89b71d`.

```terraform
# AppType Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppType configuration
resource "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}
```

<a id="canonical-e110c5d802c47a336247e11a74c13e13c9346edd58373928e24f11596f5837ef"></a>

## Next pages — Resource / 4ae38d1ab341 / 3

- [Examples](resources--app_type--examples--group-001.md#canonical-f7b9298394e83b8e4a094f5b90535372a175dda62f0fa6b17e5bb40ccf6b6b6d)
- [xcsh_app_type](../resources/app_type.md#canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692)
