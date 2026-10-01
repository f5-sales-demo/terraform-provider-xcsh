---
page_title: "xcsh_api_testing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing examples."
---

# xcsh_api_testing examples

<a id="canonical-dabec0e7850cf899bd593e2fb91b65674444d08aa39d98a84e5660afac26f91a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0f55a1a0d0f999bf532d32b7bf1b2f3ae6901a02f7f9e923d0bf5627a2b20bc"></a>

## Examples — Examples / 680fb2ee0b86 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- Examples

<a id="canonical-2752884eae38f4f6a7d32fa70725ff7ae94ce1992d4a7cd49d056060f71538df"></a>

## Complete configurations — Examples / 680fb2ee0b86 / 3

- [Resource](resources--api_testing--examples--group-001.md#canonical-f9a506b05d97786c5581105bf03ab3e29db26b7e8df39ad7e00933919efa8eef): valid configuration.

<a id="canonical-e88d9116ee4cb3063091cdbddea81eb0105e2b71d1ee32e5a28a8d2dc7622f94"></a>

## Next pages — Examples / 680fb2ee0b86 / 4

- [Resource](resources--api_testing--examples--group-001.md#canonical-f9a506b05d97786c5581105bf03ab3e29db26b7e8df39ad7e00933919efa8eef)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)

<a id="canonical-f9a506b05d97786c5581105bf03ab3e29db26b7e8df39ad7e00933919efa8eef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efeadc22f144ae652459cc371284881f816e61918144732b6783074f73575a5b"></a>

## Resource — Resource / 64afab4a9504 / 2

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
- [Examples](resources--api_testing--examples--group-001.md#canonical-dabec0e7850cf899bd593e2fb91b65674444d08aa39d98a84e5660afac26f91a)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_testing/resource.tf`; digest `sha256:c1227be913c1052c8bc552ce9ebe6740f44bc74215b29a565808e69bd8cea2fb`.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```

<a id="canonical-38f0cde11d80b758eab3565fdbeb714a8b7bf6345c7f92819662962c34abb5f6"></a>

## Next pages — Resource / 64afab4a9504 / 3

- [Examples](resources--api_testing--examples--group-001.md#canonical-dabec0e7850cf899bd593e2fb91b65674444d08aa39d98a84e5660afac26f91a)
- [xcsh_api_testing](../resources/api_testing.md#canonical-78dd6bf18113b2fe02ce024626001569b307a60d3f4c5ce1e4e7e36df50148e2)
