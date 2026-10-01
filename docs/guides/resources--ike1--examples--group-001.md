---
page_title: "xcsh_ike1 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 examples."
---

# xcsh_ike1 examples

<a id="canonical-a90ace16fce8dfc643a25b1907e73a1d38df1be791702aa2eaa0f7006aee33a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d922b773062f1be7dbe02ea7ef343f356639b30809f45747223dd848408ff90c"></a>

## Examples — Examples / 4b448a36ac1b / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- Examples

<a id="canonical-28dc79aa8438ec848b4fca15072dae10fc66868b6d54ae26e2154a0998c8dd94"></a>

## Complete configurations — Examples / 4b448a36ac1b / 3

- [Resource](resources--ike1--examples--group-001.md#canonical-2688b648fb203d5864f5e411d2994b24cc9c8f1656e264b5828984300604b5a9): valid configuration.

<a id="canonical-97b5f2a957f4ae0478fbbbf1c8d95b4f155ba092289102d98732e9d40b1233bc"></a>

## Next pages — Examples / 4b448a36ac1b / 4

- [Resource](resources--ike1--examples--group-001.md#canonical-2688b648fb203d5864f5e411d2994b24cc9c8f1656e264b5828984300604b5a9)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)

<a id="canonical-2688b648fb203d5864f5e411d2994b24cc9c8f1656e264b5828984300604b5a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5564e583bcc83e3386b03f2a775b32c7a79f516c9b1e98d02c1ca14890c09a23"></a>

## Resource — Resource / 226d670f7829 / 2

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
- [Examples](resources--ike1--examples--group-001.md#canonical-a90ace16fce8dfc643a25b1907e73a1d38df1be791702aa2eaa0f7006aee33a1)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike1/resource.tf`; digest `sha256:3f9c75e223b14e98fb8be4f0bccd9b3e08736eec2e962a472124a5092f03950e`.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```

<a id="canonical-ef53eb73ad51d26dc68422a16cd5e7a914942b56c73ccfc72308e2164e36c9d1"></a>

## Next pages — Resource / 226d670f7829 / 3

- [Examples](resources--ike1--examples--group-001.md#canonical-a90ace16fce8dfc643a25b1907e73a1d38df1be791702aa2eaa0f7006aee33a1)
- [xcsh_ike1](../resources/ike1.md#canonical-81570b716a17db42e7a83bf26108dfb7a5809432f2a942ba3e8a9db577b0a45c)
