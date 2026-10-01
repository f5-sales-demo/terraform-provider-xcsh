---
page_title: "xcsh_irule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule landing."
---

# xcsh_irule landing

<a id="canonical-2d46463305002f54ff3a1ae4bb70fcb0aa618cbcb1816cb90bb6c2a6ed549962"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d1ea6c730d0a71c0695423fd3d2bbeafbb59ebda3cbfa5ee08a3030ffc179bd"></a>

## xcsh_irule — xcsh_irule / f6f98d09b867 / 2

Breadcrumbs:

- xcsh_irule

Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed
Cloud.

<a id="canonical-3ece3b776c61b9d1544cb419c600cc179585c83a9a790e1454950fbe3fba21b9"></a>

## Prerequisites — xcsh_irule / f6f98d09b867 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8557da19ad31f83dc0572b4c22b705e3de477b0dc1fc31e60b571b996eb5eff4"></a>

## Minimal configuration — xcsh_irule / f6f98d09b867 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Irule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Irule by name
data "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"
}

output "irule_id" {
  value = data.xcsh_irule.example.id
}
```

<a id="canonical-de5ef5f1f76868e1784031810019a507a2b7fb9e2d9057c9c786121fad7e0eab"></a>

## Root configuration — xcsh_irule / f6f98d09b867 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-92ba14ca5e62f0a29daa9b2799c6af560b7b0fed5e0d0b21f91bc30720933c0b"></a>

## Next pages — xcsh_irule / f6f98d09b867 / 6

- [Property reference](../guides/data-sources--irule--reference--group-001.md#canonical-86a3e148141db31b69bea2546604e814213c814fc5ca09e8be422e905b31ed84)
- [Examples](../guides/data-sources--irule--examples--group-001.md#canonical-2de7ae51f53e24a111116c100bfd19a60484b4b51f3ad177bb6724e12299aa5d)
