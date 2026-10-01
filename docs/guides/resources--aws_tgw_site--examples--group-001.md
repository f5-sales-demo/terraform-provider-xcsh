---
page_title: "xcsh_aws_tgw_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site examples."
---

# xcsh_aws_tgw_site examples

<a id="canonical-42c06f13d48332ed5021e5a105e3d74e9187d47b541c18dbff2f2277fff7eadb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c7e1c6409aef5ac2ca3709cbac528dc1356b2d0e75db0fc3aa0e25bb2511750"></a>

## Examples — Examples / 886392e42c86 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- Examples

<a id="canonical-3c68d75d8c308957ed0530776f1108027ec1f9acc8ea68187759da015e88191a"></a>

## Complete configurations — Examples / 886392e42c86 / 3

- [Resource](resources--aws_tgw_site--examples--group-001.md#canonical-28444995e46103eb5232843c73dac2def25addb9aaada2c9d7cf3cb2520bb171): valid configuration.

<a id="canonical-33e8587a4c0161be44958c6ddeb271248ff86ccfc86d2359c89c0ac03f9c03c0"></a>

## Next pages — Examples / 886392e42c86 / 4

- [Resource](resources--aws_tgw_site--examples--group-001.md#canonical-28444995e46103eb5232843c73dac2def25addb9aaada2c9d7cf3cb2520bb171)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-28444995e46103eb5232843c73dac2def25addb9aaada2c9d7cf3cb2520bb171"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f24fe939d310d90dce4b418e0957bcc39b2363b83c1a54909399626b7b2accbd"></a>

## Resource — Resource / d18ec6eb4e3d / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Examples](resources--aws_tgw_site--examples--group-001.md#canonical-42c06f13d48332ed5021e5a105e3d74e9187d47b541c18dbff2f2277fff7eadb)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_aws_tgw_site/resource.tf`; digest `sha256:3c51271321a366128fb5a33731b3df495e257a63396c08d4e66b33465ad7f22a`.

```terraform
# AWSTGWSite Resource Example
# Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS Transit Gateway.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AWSTGWSite configuration
resource "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}
```

<a id="canonical-37dfcd8791c8d32021aba9af0d9f1a5e91a816ae8b633f8e21c04e13c67613b4"></a>

## Next pages — Resource / d18ec6eb4e3d / 3

- [Examples](resources--aws_tgw_site--examples--group-001.md#canonical-42c06f13d48332ed5021e5a105e3d74e9187d47b541c18dbff2f2277fff7eadb)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
