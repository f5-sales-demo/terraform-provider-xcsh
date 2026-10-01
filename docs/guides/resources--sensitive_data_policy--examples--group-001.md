---
page_title: "xcsh_sensitive_data_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy examples."
---

# xcsh_sensitive_data_policy examples

<a id="canonical-9a55abc4607becc0302d726239e05592c1088a771f57f7445193af0f804f2f8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-667d3036a6efe0f2bef6f82028ffb977e3d699e5cc82994edf969e13e9bc09ed"></a>

## Examples — Examples / 08a33d7b5b96 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
- Examples

<a id="canonical-e3ecffc6709fab67128143989bbf50690b5847ac82e5ebc6533182e14a9a2fe7"></a>

## Complete configurations — Examples / 08a33d7b5b96 / 3

- [Resource](resources--sensitive_data_policy--examples--group-001.md#canonical-57b01811925f24d273d11864e82dc1f8eac3bee10a36f7b823e2bf80b9adf179): valid configuration.

<a id="canonical-27c785e771f57b262c02c4813f6681a4ed8f139466012f92e6e1b10090153b47"></a>

## Next pages — Examples / 08a33d7b5b96 / 4

- [Resource](resources--sensitive_data_policy--examples--group-001.md#canonical-57b01811925f24d273d11864e82dc1f8eac3bee10a36f7b823e2bf80b9adf179)
- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)

<a id="canonical-57b01811925f24d273d11864e82dc1f8eac3bee10a36f7b823e2bf80b9adf179"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8d534fdda32e1bd458e47ccc2222d479941984e6dd8a4cab21757cbd4854cfe"></a>

## Resource — Resource / 30ccfdf3fb39 / 2

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
- [Examples](resources--sensitive_data_policy--examples--group-001.md#canonical-9a55abc4607becc0302d726239e05592c1088a771f57f7445193af0f804f2f8b)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_sensitive_data_policy/resource.tf`; digest `sha256:1472df2a4b2a381ee1201907e79b19112e2352a74e50b30a503511f47b7c7e1a`.

```terraform
# SensitiveDataPolicy Resource Example
# Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SensitiveDataPolicy configuration
resource "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}
```

<a id="canonical-87eb1a86a7474cca42051a5d3106b2492f0aa39e052cfa6ff913988175281e22"></a>

## Next pages — Resource / 30ccfdf3fb39 / 3

- [Examples](resources--sensitive_data_policy--examples--group-001.md#canonical-9a55abc4607becc0302d726239e05592c1088a771f57f7445193af0f804f2f8b)
- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44)
