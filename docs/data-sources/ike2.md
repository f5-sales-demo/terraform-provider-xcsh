---
page_title: "xcsh_ike2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 landing."
---

# xcsh_ike2 landing

<a id="canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4385113fd9f2ea83744a43003df4b4797a223cf64c928b05208b07eef0cc5162"></a>

## xcsh_ike2 — xcsh_ike2 / a042d6e726ef / 2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

<a id="canonical-521f593bb22a44d97840ad76a7cbc144d1aff4f1283e7cac373375a237e6d01e"></a>

## Prerequisites — xcsh_ike2 / a042d6e726ef / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e20e89c44fcc19a31bb809b09122e95b200a2745ff2819cbe88687666201e2c6"></a>

## Minimal configuration — xcsh_ike2 / a042d6e726ef / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike2 by name
data "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}

output "ike2_id" {
  value = data.xcsh_ike2.example.id
}
```

<a id="canonical-8aa91200b587e5e1575276d745c75c1ee7cff9c9ea8178cb74bc0aeddfe0b788"></a>

## Root configuration — xcsh_ike2 / a042d6e726ef / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-4d6486dd80e9594d14a0e9914759ca03e0aa8362d7bc30dc83ad7db548506332"></a>

## Next pages — xcsh_ike2 / a042d6e726ef / 6

- [Property reference](../guides/data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- [Examples](../guides/data-sources--ike2--examples--group-001.md#canonical-5c44f3eaa6bf5c1d6489c418cbeccb9f274ea0a5737c1a60a95ec91dd69d099b)
