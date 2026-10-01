---
page_title: "xcsh_app_setting examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting examples."
---

# xcsh_app_setting examples

<a id="canonical-a296127a92a7ce97e2aee94840e4cb600879e5446e23f1b20bca93486f3f3dbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c6ae8f596038db9a0a1753747b8abf0266d59eb6b5c3ba7122ec2b0ceffb995"></a>

## Examples — Examples / ad1f76c6f931 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- Examples

<a id="canonical-a60161a6d884efb59484bf16504a7316926512df8bce66810b92158e07e9469c"></a>

## Complete configurations — Examples / ad1f76c6f931 / 3

- [Data source](data-sources--app_setting--examples--group-001.md#canonical-e738e64750d05c8c4207ac559759fd1a2c5ee22b1c412c9bca8d232d4322741c): valid configuration.

<a id="canonical-6eee2b817d9ae63f42aaa6c50016a10044cb4c759fea4c8d48e36e03a44b246f"></a>

## Next pages — Examples / ad1f76c6f931 / 4

- [Data source](data-sources--app_setting--examples--group-001.md#canonical-e738e64750d05c8c4207ac559759fd1a2c5ee22b1c412c9bca8d232d4322741c)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)

<a id="canonical-e738e64750d05c8c4207ac559759fd1a2c5ee22b1c412c9bca8d232d4322741c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-816e21b413c2f5022c92b48da1d8be48d908ffb35138437c714b23eacb102fb3"></a>

## Data source — Data source / aee38966b6a5 / 2

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
- [Examples](data-sources--app_setting--examples--group-001.md#canonical-a296127a92a7ce97e2aee94840e4cb600879e5446e23f1b20bca93486f3f3dbd)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_setting/data-source.tf`; digest `sha256:4e31d15f083391674e4e309fb8e0ab1858f7a5c2638cd424cc62bee3775dfde7`.

```terraform
# AppSetting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppSetting by name
data "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}

output "app_setting_id" {
  value = data.xcsh_app_setting.example.id
}
```

<a id="canonical-89b831305468af1a5b333c2ae1214ef935a1a0fb57d8bc3ee056bf21820ff59b"></a>

## Next pages — Data source / aee38966b6a5 / 3

- [Examples](data-sources--app_setting--examples--group-001.md#canonical-a296127a92a7ce97e2aee94840e4cb600879e5446e23f1b20bca93486f3f3dbd)
- [xcsh_app_setting](../data-sources/app_setting.md#canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9)
