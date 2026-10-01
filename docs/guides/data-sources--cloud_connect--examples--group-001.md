---
page_title: "xcsh_cloud_connect examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect examples."
---

# xcsh_cloud_connect examples

<a id="canonical-0e447ba11f2ec3291f89ffa097e0f893c0f5cb51189a0746945199aa4d3c314f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-541b0c01af93e932dfd42b1a40b45abd3aae32ed2b0bd3244a78b58add0dbaf0"></a>

## Examples — Examples / 6ba3fdda9569 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- Examples

<a id="canonical-3fcc7df7bf4ac368a1e748cb2f3c299d75493da3a7ec4b2eaeb1a8a943b8a8a8"></a>

## Complete configurations — Examples / 6ba3fdda9569 / 3

- [Data source](data-sources--cloud_connect--examples--group-001.md#canonical-caf1cebf162707580ed0ca16d04dbfc95f7f6c9be5fd93e8c0b2582d495535fe): valid configuration.

<a id="canonical-668565c9e304fa524801a302833d51abfdf0663e65c4f269f56eb63c6c732d76"></a>

## Next pages — Examples / 6ba3fdda9569 / 4

- [Data source](data-sources--cloud_connect--examples--group-001.md#canonical-caf1cebf162707580ed0ca16d04dbfc95f7f6c9be5fd93e8c0b2582d495535fe)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)

<a id="canonical-caf1cebf162707580ed0ca16d04dbfc95f7f6c9be5fd93e8c0b2582d495535fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-600b0914db544b7b88f0d80c9cd5ff591f79941e2003249f2b2c6782ef2f95fe"></a>

## Data source — Data source / cb1fc996c787 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
- [Examples](data-sources--cloud_connect--examples--group-001.md#canonical-0e447ba11f2ec3291f89ffa097e0f893c0f5cb51189a0746945199aa4d3c314f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_connect/data-source.tf`; digest `sha256:eed109fcd2f0afc5835242d94d0d30137f2b13d899b31f9e138f00b05427dcd6`.

```terraform
# CloudConnect Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudConnect by name
data "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}

output "cloud_connect_id" {
  value = data.xcsh_cloud_connect.example.id
}
```

<a id="canonical-c95c0c386f541b7fe1401463b54edd9d9251d8db97eaab0ee5ad0442c9846f6d"></a>

## Next pages — Data source / cb1fc996c787 / 3

- [Examples](data-sources--cloud_connect--examples--group-001.md#canonical-0e447ba11f2ec3291f89ffa097e0f893c0f5cb51189a0746945199aa4d3c314f)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-853b4d0ff5eda08da7c860d5cfbd4187898ed900e988089501f5f43cf191ed80)
