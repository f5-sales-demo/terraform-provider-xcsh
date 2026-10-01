---
page_title: "xcsh_usb_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy examples."
---

# xcsh_usb_policy examples

<a id="canonical-1adaa97e282905281884b4084b173c708d144e90757a4955cc9d306a91eb2732"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a8c03c9dc0848f70e3511694aa873bed24fcf48c5d7354be06b15a26825a226"></a>

## Examples — Examples / 8f5928393611 / 2

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)
- Examples

<a id="canonical-28ea7992fb7482c7dc35f5cfbe1cf6d0c4787ea1ca0dab9fc6e636a0be17842d"></a>

## Complete configurations — Examples / 8f5928393611 / 3

- [Resource](resources--usb_policy--examples--group-001.md#canonical-49d6ce996ba77fc0f242a1f23a479fb6954bfebdc8903434601aa94552c53136): valid configuration.

<a id="canonical-7215ed6c62cf3ccd8937b93482ee2383448abcd485cba0322a20971121239dab"></a>

## Next pages — Examples / 8f5928393611 / 4

- [Resource](resources--usb_policy--examples--group-001.md#canonical-49d6ce996ba77fc0f242a1f23a479fb6954bfebdc8903434601aa94552c53136)
- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)

<a id="canonical-49d6ce996ba77fc0f242a1f23a479fb6954bfebdc8903434601aa94552c53136"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90c75703afccbb0ce54636ac1f09b8f1739e2131c6704c54268cc3cb3afabfca"></a>

## Resource — Resource / 4c3aa69d5a05 / 2

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)
- [Examples](resources--usb_policy--examples--group-001.md#canonical-1adaa97e282905281884b4084b173c708d144e90757a4955cc9d306a91eb2732)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_usb_policy/resource.tf`; digest `sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e`.

```terraform
# UsbPolicy Resource Example
# Manages new USB policy object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UsbPolicy configuration
resource "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}
```

<a id="canonical-bb0952e9a7b16eb1fb0ce0ecce19820b679b594beed617b1f309acdb917ee921"></a>

## Next pages — Resource / 4c3aa69d5a05 / 3

- [Examples](resources--usb_policy--examples--group-001.md#canonical-1adaa97e282905281884b4084b173c708d144e90757a4955cc9d306a91eb2732)
- [xcsh_usb_policy](../resources/usb_policy.md#canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae)
