---
page_title: "xcsh_app_setting examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting examples."
---

# xcsh_app_setting examples

<a id="canonical-f8697b00e9618b953de504386b5babb6260cb03fc79886704ad06f4a7bf39245"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b27705ef1f7cf03b125c69f884e2ec316768e2b68b8550de77ff7440733b3f12"></a>

## Examples — Examples / ba1418bfbdde / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- Examples

<a id="canonical-60b4d9bb92f304df21ba6f0c5b501f2b4d3dbe412f22a17d515774afa01f2537"></a>

## Complete configurations — Examples / ba1418bfbdde / 3

- [Resource](resources--app_setting--examples--group-001.md#canonical-303135ba8286390c47365043319c1a43e2f00d70b2e920feba657d73ca87f623): valid configuration.

<a id="canonical-6b8d321cb11f559e68343bda1b7de39eec52c5081def676a189467d8c5a53c1d"></a>

## Next pages — Examples / ba1418bfbdde / 4

- [Resource](resources--app_setting--examples--group-001.md#canonical-303135ba8286390c47365043319c1a43e2f00d70b2e920feba657d73ca87f623)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)

<a id="canonical-303135ba8286390c47365043319c1a43e2f00d70b2e920feba657d73ca87f623"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3ad0fee87d4fa4f46a82736bf1f8881aad80e1c71b4f4646b7ccaa5b26fe6b7"></a>

## Resource — Resource / aabc005c0b85 / 2

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
- [Examples](resources--app_setting--examples--group-001.md#canonical-f8697b00e9618b953de504386b5babb6260cb03fc79886704ad06f4a7bf39245)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_setting/resource.tf`; digest `sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a`.

```terraform
# AppSetting Resource Example
# Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppSetting configuration
resource "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}
```

<a id="canonical-228cc70a33d868cea8be5038d8481bc17ae4cb76bc631be73a23076bdf89718b"></a>

## Next pages — Resource / aabc005c0b85 / 3

- [Examples](resources--app_setting--examples--group-001.md#canonical-f8697b00e9618b953de504386b5babb6260cb03fc79886704ad06f4a7bf39245)
- [xcsh_app_setting](../resources/app_setting.md#canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94)
