---
page_title: "xcsh_authorization_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server examples."
---

# xcsh_authorization_server examples

<a id="canonical-1bb99cc587193065495177dec67c564e855be1be071a164a9bfb77b8a6432f7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2f53948d0c019d518f883d12b4d204be743c0ab36a49db25dfd32a80ce9691e"></a>

## Examples — Examples / f7926d857bd2 / 2

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)
- Examples

<a id="canonical-b0771c4a04ea7dcd40975bbf522bb0e347b39474ee363b1924d8fcee27d04d21"></a>

## Complete configurations — Examples / f7926d857bd2 / 3

- [Resource](resources--authorization_server--examples--group-001.md#canonical-546c6ec28076918ca23ba5c4319509f51e11c06be3a292e7ddaea757717e8e47): valid configuration.

<a id="canonical-7265f33a43fe06d197c31a3f23e01c61dea1f83995c1a78ff5b823578a934bb9"></a>

## Next pages — Examples / f7926d857bd2 / 4

- [Resource](resources--authorization_server--examples--group-001.md#canonical-546c6ec28076918ca23ba5c4319509f51e11c06be3a292e7ddaea757717e8e47)
- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)

<a id="canonical-546c6ec28076918ca23ba5c4319509f51e11c06be3a292e7ddaea757717e8e47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebe6e59c5d8f3995b51677044731d351929385b2a36ea436ac72a48e738637d0"></a>

## Resource — Resource / 1d38f1e689be / 2

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)
- [Examples](resources--authorization_server--examples--group-001.md#canonical-1bb99cc587193065495177dec67c564e855be1be071a164a9bfb77b8a6432f7c)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authorization_server/resource.tf`; digest `sha256:bbb05ccd40cd71dbeef4a9b6ed31c1cec2684fa158b0357f99c63ef6b31e6af5`.

```terraform
# AuthorizationServer Resource Example
# Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AuthorizationServer configuration
resource "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"

  jwks_uri = "example-value"
}
```

<a id="canonical-9f20a00debf4a269d6a8eb9efaf3f18f54d2744834869d5c8890ae3be53a4321"></a>

## Next pages — Resource / 1d38f1e689be / 3

- [Examples](resources--authorization_server--examples--group-001.md#canonical-1bb99cc587193065495177dec67c564e855be1be071a164a9bfb77b8a6432f7c)
- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)
