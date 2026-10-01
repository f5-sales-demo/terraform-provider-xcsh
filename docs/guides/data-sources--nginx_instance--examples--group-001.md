---
page_title: "xcsh_nginx_instance examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_instance examples."
---

# xcsh_nginx_instance examples

<a id="canonical-f1cfd45c9eb52b9043c8dddc78a39f8e6d4c57431fafe737175b7666f359758d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6103b605987ef013576e9056882dec3a8a36acb7c80817781c32247f0c9f71f"></a>

## Examples — Examples / ca406d14b26e / 2

Breadcrumbs:

- [xcsh_nginx_instance](../data-sources/nginx_instance.md#canonical-b9ef2669c0bf7b4c379f2af4596962e7273bf8b58edf51272ccb89570543335a)
- Examples

<a id="canonical-7a6ad47aa0dd8876d57a1812959537719aef90e692edda650330b8057fba61d9"></a>

## Complete configurations — Examples / ca406d14b26e / 3

- [Data source](data-sources--nginx_instance--examples--group-001.md#canonical-ee5adf47193967874ec777e39dd3a2ae1e42374fe76ee735c7d02c22195387df): valid configuration.

<a id="canonical-bb4a5318b3a06a46d238904d6c7a57541630c384a9febbb08d2b55c0d115d501"></a>

## Next pages — Examples / ca406d14b26e / 4

- [Data source](data-sources--nginx_instance--examples--group-001.md#canonical-ee5adf47193967874ec777e39dd3a2ae1e42374fe76ee735c7d02c22195387df)
- [xcsh_nginx_instance](../data-sources/nginx_instance.md#canonical-b9ef2669c0bf7b4c379f2af4596962e7273bf8b58edf51272ccb89570543335a)

<a id="canonical-ee5adf47193967874ec777e39dd3a2ae1e42374fe76ee735c7d02c22195387df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16e1b5a61cc179c06daf263919afa5e076923ceb31b501c42fd6ad8eb43a27f3"></a>

## Data source — Data source / 7b7a7f913d0c / 2

Breadcrumbs:

- [xcsh_nginx_instance](../data-sources/nginx_instance.md#canonical-b9ef2669c0bf7b4c379f2af4596962e7273bf8b58edf51272ccb89570543335a)
- [Examples](data-sources--nginx_instance--examples--group-001.md#canonical-f1cfd45c9eb52b9043c8dddc78a39f8e6d4c57431fafe737175b7666f359758d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_instance/data-source.tf`; digest `sha256:772405375dce9fbbb8eabe6970ee7d9b3a944178189b317c587b8f2983eaf8fd`.

```terraform
# NginxInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxInstance by name
data "xcsh_nginx_instance" "example" {
  name      = "example-nginx-instance"
  namespace = "staging"
}

output "nginx_instance_id" {
  value = data.xcsh_nginx_instance.example.id
}
```

<a id="canonical-770cf13cce4c7cfe7d8f854ab5f519f7056e07211d01d777628bea5df7246966"></a>

## Next pages — Data source / 7b7a7f913d0c / 3

- [Examples](data-sources--nginx_instance--examples--group-001.md#canonical-f1cfd45c9eb52b9043c8dddc78a39f8e6d4c57431fafe737175b7666f359758d)
- [xcsh_nginx_instance](../data-sources/nginx_instance.md#canonical-b9ef2669c0bf7b4c379f2af4596962e7273bf8b58edf51272ccb89570543335a)
