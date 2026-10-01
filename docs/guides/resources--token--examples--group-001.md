---
page_title: "xcsh_token examples"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token examples."
---

# xcsh_token examples

<a id="canonical-6d07d12677cc1288658a0cbde21f944c8cff33dfa0d3d0f5cd2ee61b235e0a92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7abf80886971e85238ad8ca562d88272aca79f85f9d48da2a6c2567b6bf539c"></a>

## Examples — Examples / 398353a155e2 / 2

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)
- Examples

<a id="canonical-bc415dcc44c30fc5efa62378d0d9c72bb85535f2f329cd161918505414d59e68"></a>

## Complete configurations — Examples / 398353a155e2 / 3

- [Resource](resources--token--examples--group-001.md#canonical-91f34f623e662bfff6638a53aa5fef3c234409591619b96bb9cadff6d8495d3e): valid configuration.

<a id="canonical-786ddd96909f16e57f158d0836e92da311612e56cad754fabe3ce834429e5fba"></a>

## Next pages — Examples / 398353a155e2 / 4

- [Resource](resources--token--examples--group-001.md#canonical-91f34f623e662bfff6638a53aa5fef3c234409591619b96bb9cadff6d8495d3e)
- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)

<a id="canonical-91f34f623e662bfff6638a53aa5fef3c234409591619b96bb9cadff6d8495d3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-352a4b6df9bddd9e40b95364db18132e2f6e41311f0fd8df0cf8797806778ac5"></a>

## Resource — Resource / 308ec6c6fbac / 2

Breadcrumbs:

- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)
- [Examples](resources--token--examples--group-001.md#canonical-6d07d12677cc1288658a0cbde21f944c8cff33dfa0d3d0f5cd2ee61b235e0a92)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_token/resource.tf`; digest `sha256:e34df2fe171a3a579b8dd3181a5ec00f7f49f2449740ae1edcf8cdaec5c57bb9`.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```

<a id="canonical-7eed1a3bb324d60ea3f2b5a2c217574d175bd421d49b8bdbde0f10fc97534e20"></a>

## Next pages — Resource / 308ec6c6fbac / 3

- [Examples](resources--token--examples--group-001.md#canonical-6d07d12677cc1288658a0cbde21f944c8cff33dfa0d3d0f5cd2ee61b235e0a92)
- [xcsh_token](../resources/token.md#canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8)
