---
page_title: "xcsh_certificate examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate examples."
---

# xcsh_certificate examples

<a id="canonical-fadf3f9d1d214e0a699b1607ad361706f2c97596ce1ac016dffdb3f418d41ffe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f7ae5eacaca80ca0333f499d31610daa1027eef22da63d315200d977ce8f4a1"></a>

## Examples — Examples / 4a4225361e8d / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- Examples

<a id="canonical-f1cd3c5515ee4b568efbdfbb4be0d446f0eedad542c53fd20fb4b18e6d028ed1"></a>

## Complete configurations — Examples / 4a4225361e8d / 3

- [Resource](resources--certificate--examples--group-001.md#canonical-56a7294d8d4da130412016cae380c0ecc9b3cce0aeb31acab5db18aa43fafc49): valid configuration.

<a id="canonical-dbc055a67f4a240686951ec6b138c3c8045cfb489c801573ae47e8747f53ed9b"></a>

## Next pages — Examples / 4a4225361e8d / 4

- [Resource](resources--certificate--examples--group-001.md#canonical-56a7294d8d4da130412016cae380c0ecc9b3cce0aeb31acab5db18aa43fafc49)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)

<a id="canonical-56a7294d8d4da130412016cae380c0ecc9b3cce0aeb31acab5db18aa43fafc49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53d5a9691b274f6dab71e7cb499167d6da86b3a06959a809ab44023b8973843a"></a>

## Resource — Resource / fbe1cdf5c433 / 2

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
- [Examples](resources--certificate--examples--group-001.md#canonical-fadf3f9d1d214e0a699b1607ad361706f2c97596ce1ac016dffdb3f418d41ffe)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate/resource.tf`; digest `sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092`.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```

<a id="canonical-dd486bc779e7902b4c0e442617e0d2a2c567d76bd913e27be1d07916fc7f8cb9"></a>

## Next pages — Resource / fbe1cdf5c433 / 3

- [Examples](resources--certificate--examples--group-001.md#canonical-fadf3f9d1d214e0a699b1607ad361706f2c97596ce1ac016dffdb3f418d41ffe)
- [xcsh_certificate](../resources/certificate.md#canonical-843c7ef2d56e85f1ba6b346483224fce2e27fa0dc0213cd0e259331947e5281c)
