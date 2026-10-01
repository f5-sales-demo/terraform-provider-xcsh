---
page_title: "xcsh_fast_acl examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl examples."
---

# xcsh_fast_acl examples

<a id="canonical-c6c405700be873f5bc2b525c10a682d3e4871c3edee55da231f7c317c410bf0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28ca08c517d7521174848d1dae917a6ef8b7881f742013759d411f39b73c2f9f"></a>

## Examples — Examples / 7481be855169 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- Examples

<a id="canonical-a7d601749834262a13141b3e8b1ef55ede0f29b1dad22fdb03c607c04ff54be0"></a>

## Complete configurations — Examples / 7481be855169 / 3

- [Resource](resources--fast_acl--examples--group-001.md#canonical-262a77e3da7765076fde8798cae48d564cf48e752c1f3b6fdad9ad8c86cbb7e5): valid configuration.

<a id="canonical-88eafff2d99e2922b8b0cb32fc0a7c8cc26da8b63f577a231744123b8d5bcb0f"></a>

## Next pages — Examples / 7481be855169 / 4

- [Resource](resources--fast_acl--examples--group-001.md#canonical-262a77e3da7765076fde8798cae48d564cf48e752c1f3b6fdad9ad8c86cbb7e5)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)

<a id="canonical-262a77e3da7765076fde8798cae48d564cf48e752c1f3b6fdad9ad8c86cbb7e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-379488c5cf6d3733bca7690fb5c55baad12390183bbc05d2216f48fcc13339ca"></a>

## Resource — Resource / 55fb26826040 / 2

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
- [Examples](resources--fast_acl--examples--group-001.md#canonical-c6c405700be873f5bc2b525c10a682d3e4871c3edee55da231f7c317c410bf0d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_fast_acl/resource.tf`; digest `sha256:5b0dbb55293dd11d1f64bb7c693af3c05672ada5840fb242470c138bb98e2c6d`.

```terraform
# FastACL Resource Example
# Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACL configuration
resource "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}
```

<a id="canonical-e0913a2b8c2c9e712b783a729750a8f4622168b83996cc5dd70f52a94254ecf3"></a>

## Next pages — Resource / 55fb26826040 / 3

- [Examples](resources--fast_acl--examples--group-001.md#canonical-c6c405700be873f5bc2b525c10a682d3e4871c3edee55da231f7c317c410bf0d)
- [xcsh_fast_acl](../resources/fast_acl.md#canonical-ba6c26b78da2e4c8bf56de0adba953c112a8ab6a32e54eb628d06a3077d55892)
