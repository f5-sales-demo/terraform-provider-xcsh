---
page_title: "xcsh_virtual_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site landing."
---

# xcsh_virtual_site landing

<a id="canonical-4d9c6e801f11ee0a4dcb7c662b2cdc5dffca2ee94e25d78ffb25ff05789e2980"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f33124c82bcf5ff574774fb7ca2a50b763e8520237f19e472740daff6068fd90"></a>

## xcsh_virtual_site — xcsh_virtual_site / 139d156566bc / 2

Breadcrumbs:

- xcsh_virtual_site

Manages virtual site object in given namespace in F5 Distributed Cloud.

<a id="canonical-bf593f73d4b184fbc12b7f8c77161453949aa06bcf18739ccba6d6ae0cb06f31"></a>

## Prerequisites — xcsh_virtual_site / 139d156566bc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-a289b65db06864fa1225d7b6b98196fdbb9e3af609069a195d6216374100dbd2"></a>

## Minimal configuration — xcsh_virtual_site / 139d156566bc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualSite by name
data "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}

output "virtual_site_id" {
  value = data.xcsh_virtual_site.example.id
}
```

<a id="canonical-f407502d2bdf886775e9b24dd366822465fe16a89fcb2d43bf1925446443c316"></a>

## Root configuration — xcsh_virtual_site / 139d156566bc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3876bd0c09ad02917d22d76e4d4b2ae3cc4437f8efd7d1d0dfc9bca874709cec"></a>

## Next pages — xcsh_virtual_site / 139d156566bc / 6

- [Property reference](../guides/data-sources--virtual_site--reference--group-001.md#canonical-011be626f96641414cc8e65665477f0a7768f907c5427772ed92f6dd2c667a74)
- [Examples](../guides/data-sources--virtual_site--examples--group-001.md#canonical-5bd147df1752d218cbb1af93aa5ac43f096f47268571b17c6a4a9b3e92a11bd6)
