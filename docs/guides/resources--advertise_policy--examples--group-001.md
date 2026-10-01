---
page_title: "xcsh_advertise_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy examples."
---

# xcsh_advertise_policy examples

<a id="canonical-6a3b267374c676ae36bae3d8f4ad9b293170f2b9387c9e12c907fa0a47d6f5eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b76c137ae9bcb61fc03c28b5c918ebbdbb83a4645b6141de236ee76c97e5cb17"></a>

## Examples — Examples / 2893c2dac8c2 / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- Examples

<a id="canonical-eee539f5feac160b17ea59bfedfd04ff685677793edca443c67f33c0832e458a"></a>

## Complete configurations — Examples / 2893c2dac8c2 / 3

- [Resource](resources--advertise_policy--examples--group-001.md#canonical-5ccb686da7f80adba1c7295968f1fdf0d268876d84c951d197e203782acfcab9): valid configuration.

<a id="canonical-77653d4a75877be486d0401de825692808e930a09a38290e2f12a692c402300c"></a>

## Next pages — Examples / 2893c2dac8c2 / 4

- [Resource](resources--advertise_policy--examples--group-001.md#canonical-5ccb686da7f80adba1c7295968f1fdf0d268876d84c951d197e203782acfcab9)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)

<a id="canonical-5ccb686da7f80adba1c7295968f1fdf0d268876d84c951d197e203782acfcab9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-685b65bac6a496e13ab5c1376d8b7cf90c8f5bc9c8cfd21aaf60530533d9e194"></a>

## Resource — Resource / ad9ecf0b507e / 2

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
- [Examples](resources--advertise_policy--examples--group-001.md#canonical-6a3b267374c676ae36bae3d8f4ad9b293170f2b9387c9e12c907fa0a47d6f5eb)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_advertise_policy/resource.tf`; digest `sha256:9e9301fb78d45973929d2d48eacb74419a673af009dabc983da6f509200329a9`.

```terraform
# AdvertisePolicy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AdvertisePolicy configuration
resource "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}
```

<a id="canonical-65f0060b840b4cef82ee5806fc3f06afdcc482db89a6cbff364d2c36f312628f"></a>

## Next pages — Resource / ad9ecf0b507e / 3

- [Examples](resources--advertise_policy--examples--group-001.md#canonical-6a3b267374c676ae36bae3d8f4ad9b293170f2b9387c9e12c907fa0a47d6f5eb)
- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0102af0bcc407b9bfc63412c9cb22411a5c4b8663a8d3c09ad4e18ded1fe9e00)
