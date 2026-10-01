---
page_title: "xcsh_allowed_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain examples."
---

# xcsh_allowed_domain examples

<a id="canonical-104c527e8be9f381bf482f9625410f3ecb948b959e49bbecf63ce5facc703ae7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fa8756b75d96d9b9a03a86e7f6455012ec6c44450efdfd28df056d4f7e8fdef"></a>

## Examples — Examples / 75aaf79af822 / 2

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)
- Examples

<a id="canonical-b96284ba4872217096660248066bc829615baf5838669445621b305070413962"></a>

## Complete configurations — Examples / 75aaf79af822 / 3

- [Resource](resources--allowed_domain--examples--group-001.md#canonical-dd985e6f95e30e20bb49b3f880af384e84f3ad92f194a384ae367554b85a1b25): valid configuration.

<a id="canonical-cc7673678f77f3352762414568787372d1a41eef6419b970de005984e3bfa0ff"></a>

## Next pages — Examples / 75aaf79af822 / 4

- [Resource](resources--allowed_domain--examples--group-001.md#canonical-dd985e6f95e30e20bb49b3f880af384e84f3ad92f194a384ae367554b85a1b25)
- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)

<a id="canonical-dd985e6f95e30e20bb49b3f880af384e84f3ad92f194a384ae367554b85a1b25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-655f0431e97b65a8ab117b3d546daeff0423035deed1a9ec0214c1f85e80ad2d"></a>

## Resource — Resource / 57ce15b70fbd / 2

Breadcrumbs:

- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)
- [Examples](resources--allowed_domain--examples--group-001.md#canonical-104c527e8be9f381bf482f9625410f3ecb948b959e49bbecf63ce5facc703ae7)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_allowed_domain/resource.tf`; digest `sha256:311173eaa23e9e9afb5856fa5a592866eb29afd044fe81b255b02bc483f4948e`.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```

<a id="canonical-26d949d73e1186388b018750c7818278694832963275f53d7d6e64fa1be469cb"></a>

## Next pages — Resource / 57ce15b70fbd / 3

- [Examples](resources--allowed_domain--examples--group-001.md#canonical-104c527e8be9f381bf482f9625410f3ecb948b959e49bbecf63ce5facc703ae7)
- [xcsh_allowed_domain](../resources/allowed_domain.md#canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0)
