---
page_title: "xcsh_protected_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain examples."
---

# xcsh_protected_domain examples

<a id="canonical-4da6f421b79056c24836875bfa0f481d455793d2790413cc6063d3ce821bf0ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36c01412070a3bce0232977f85a50ec8d943cf24dae3c3624cde7eab927f2d81"></a>

## Examples — Examples / 6bfcaf168a57 / 2

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)
- Examples

<a id="canonical-5e492780dd624121cbe78be6df8b4df32dda6d56d97b0bdc9d4e3b43668e8025"></a>

## Complete configurations — Examples / 6bfcaf168a57 / 3

- [Resource](resources--protected_domain--examples--group-001.md#canonical-afa322b7ffa34f17e1571bce367ae96c868d2099aa988231e82298341f8ca71f): valid configuration.

<a id="canonical-2e01acc850f84b8016a604150ff3b3fde183a6c43b8d19f46debd6c6b1f6e6e2"></a>

## Next pages — Examples / 6bfcaf168a57 / 4

- [Resource](resources--protected_domain--examples--group-001.md#canonical-afa322b7ffa34f17e1571bce367ae96c868d2099aa988231e82298341f8ca71f)
- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)

<a id="canonical-afa322b7ffa34f17e1571bce367ae96c868d2099aa988231e82298341f8ca71f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-422a094ca576b40895248db61aacd2bdd58ce1ddc5f5cdc6a2d61b3a98ccaffd"></a>

## Resource — Resource / 34c0d25de4f6 / 2

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)
- [Examples](resources--protected_domain--examples--group-001.md#canonical-4da6f421b79056c24836875bfa0f481d455793d2790413cc6063d3ce821bf0ed)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protected_domain/resource.tf`; digest `sha256:a41646b64f77c1c10c9bfb255c8aeb3c51bbd0f1f5060b667473b425512ade92`.

```terraform
# ProtectedDomain Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedDomain configuration
resource "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"

  protected_domain = "example.com"
}
```

<a id="canonical-9f73b3ce8e3fcc1d6cae156c724313fe3a476d726ab44db0b39005e37430d973"></a>

## Next pages — Resource / 34c0d25de4f6 / 3

- [Examples](resources--protected_domain--examples--group-001.md#canonical-4da6f421b79056c24836875bfa0f481d455793d2790413cc6063d3ce821bf0ed)
- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)
