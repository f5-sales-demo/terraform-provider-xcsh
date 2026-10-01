---
page_title: "xcsh_voltstack_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site examples."
---

# xcsh_voltstack_site examples

<a id="canonical-80622daaa5ef0ae14ed426d2603f5d807d6b2b77ce118a081f8210d154ec653a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-331f90b58f4de256bc86f415bfa2e25bfa816f5cd7448bcdd7788437ea4f7099"></a>

## Examples — Examples / 595d65d44446 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- Examples

<a id="canonical-e7d429b1effffc9b8932fb753df0f93fa62db5a2c54cd4e1cc1cb97643edee95"></a>

## Complete configurations — Examples / 595d65d44446 / 3

- [Resource](resources--voltstack_site--examples--group-001.md#canonical-4801c6f6d271cb3b39202591ebb52df0289d30e98fbd0c2baf9019f4e235c672): valid configuration.

<a id="canonical-5bc0c7a3eaabc0b5ac9144b9c0d22d74ba4f2c1270fc11cc63a423b13ddae934"></a>

## Next pages — Examples / 595d65d44446 / 4

- [Resource](resources--voltstack_site--examples--group-001.md#canonical-4801c6f6d271cb3b39202591ebb52df0289d30e98fbd0c2baf9019f4e235c672)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)

<a id="canonical-4801c6f6d271cb3b39202591ebb52df0289d30e98fbd0c2baf9019f4e235c672"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7fc0ad2eb2435b13e546cfce8c864471edd895700c7c629c19806615fce6992"></a>

## Resource — Resource / 5d7c7a7c9cec / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
- [Examples](resources--voltstack_site--examples--group-001.md#canonical-80622daaa5ef0ae14ed426d2603f5d807d6b2b77ce118a081f8210d154ec653a)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_voltstack_site/resource.tf`; digest `sha256:45ebeb1276612423e72f2b472cb13dd3d370a975c3a55fde67381031e1b5970b`.

```terraform
# VoltstackSite Resource Example
# Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VoltstackSite configuration
resource "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

<a id="canonical-e348b94151e7db0ac477a8674668c24a2f8862c936c3d139b36a202154365651"></a>

## Next pages — Resource / 5d7c7a7c9cec / 3

- [Examples](resources--voltstack_site--examples--group-001.md#canonical-80622daaa5ef0ae14ed426d2603f5d807d6b2b77ce118a081f8210d154ec653a)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-b9688217f046b26b7c3c5660128e4d4fb1c2166fd262517b575d5580b7142b17)
