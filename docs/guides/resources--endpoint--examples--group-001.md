---
page_title: "xcsh_endpoint examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint examples."
---

# xcsh_endpoint examples

<a id="canonical-59458105bf2576af045aeb7a0d428b8b8147593e44f513c8b62c4f392b72181a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75878e66c366e3eb7a656b573a3bb15900e3750dfe62c8bcecf52315b739b3bf"></a>

## Examples — Examples / e5fe98f00e58 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- Examples

<a id="canonical-0524fa53e6ee3d7ddbefed0fd0c809008e80fc033205453facd09ddda8da68e7"></a>

## Complete configurations — Examples / e5fe98f00e58 / 3

- [Resource](resources--endpoint--examples--group-001.md#canonical-a68846d49d0aef443989a8d589bce64565b50a0c8d567bd31eb4eb89a0944b6e): valid configuration.

<a id="canonical-4e4c8076fd7c84bce8046984905d7a983ee63d11cff695bb10210b2a6519e37d"></a>

## Next pages — Examples / e5fe98f00e58 / 4

- [Resource](resources--endpoint--examples--group-001.md#canonical-a68846d49d0aef443989a8d589bce64565b50a0c8d567bd31eb4eb89a0944b6e)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)

<a id="canonical-a68846d49d0aef443989a8d589bce64565b50a0c8d567bd31eb4eb89a0944b6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9a247d245173c5c6cbfa2809df239f392b386c51e7e20cede72fe63a5734a42"></a>

## Resource — Resource / 9cc70f52ea92 / 2

Breadcrumbs:

- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
- [Examples](resources--endpoint--examples--group-001.md#canonical-59458105bf2576af045aeb7a0d428b8b8147593e44f513c8b62c4f392b72181a)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_endpoint/resource.tf`; digest `sha256:cf40cc690b6bb6c6e69c1acdd091e1041b483bbc473b7b87f06961cff98e132e`.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```

<a id="canonical-249028ed5209748b37dbbbd306ea8390f3c697dd74b9fb3ea5c3d8e6ecd270d6"></a>

## Next pages — Resource / 9cc70f52ea92 / 3

- [Examples](resources--endpoint--examples--group-001.md#canonical-59458105bf2576af045aeb7a0d428b8b8147593e44f513c8b62c4f392b72181a)
- [xcsh_endpoint](../resources/endpoint.md#canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506)
