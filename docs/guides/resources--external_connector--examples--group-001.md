---
page_title: "xcsh_external_connector examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector examples."
---

# xcsh_external_connector examples

<a id="canonical-4b2095e73af6f07a0640ac9c476f2bec2e3a1067e307b42dfc92f1a141556abf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e6af5c84ea55914d938d6196b480f4855d6fcccf9f1ff7ee10fed7cadbf5242"></a>

## Examples — Examples / 1cd2a3997f71 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- Examples

<a id="canonical-8809a0cc17fbf6d03494b3e51d13bbf37ee8eb5177febbda242002d161b6d3f6"></a>

## Complete configurations — Examples / 1cd2a3997f71 / 3

- [Resource](resources--external_connector--examples--group-001.md#canonical-7593b9c0886866e419ba64f9906d513188b6c7bf002e3b3dcf74240a07f6eb3d): valid configuration.

<a id="canonical-39e250033de22443f99c98572d926b5b630ff48d5c3dce1dbfb07cd3a02a7342"></a>

## Next pages — Examples / 1cd2a3997f71 / 4

- [Resource](resources--external_connector--examples--group-001.md#canonical-7593b9c0886866e419ba64f9906d513188b6c7bf002e3b3dcf74240a07f6eb3d)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)

<a id="canonical-7593b9c0886866e419ba64f9906d513188b6c7bf002e3b3dcf74240a07f6eb3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91f1986063bf3135dfd2e26b0b687b95cfd32615c4177c6178b3e8a9b18ed015"></a>

## Resource — Resource / d0adb18537f4 / 2

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
- [Examples](resources--external_connector--examples--group-001.md#canonical-4b2095e73af6f07a0640ac9c476f2bec2e3a1067e307b42dfc92f1a141556abf)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_external_connector/resource.tf`; digest `sha256:e6d206bd3e4355ffab92fe542ac16d79135373b63491d27d87ac9997db3c291a`.

```terraform
# ExternalConnector Resource Example
# Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ExternalConnector configuration
resource "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}
```

<a id="canonical-86e9fd0f4fd949138411f100dec0f8b395180f07985679569b847c8e62cba8c7"></a>

## Next pages — Resource / d0adb18537f4 / 3

- [Examples](resources--external_connector--examples--group-001.md#canonical-4b2095e73af6f07a0640ac9c476f2bec2e3a1067e307b42dfc92f1a141556abf)
- [xcsh_external_connector](../resources/external_connector.md#canonical-dad89a1cf0c7335cb27686cad0a7d14c290e321ef65c59324d0c5f5417b1fbea)
