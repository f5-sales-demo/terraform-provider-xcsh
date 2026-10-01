---
page_title: "xcsh_securemesh_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site landing."
---

# xcsh_securemesh_site landing

<a id="canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c8951b0aaecdc1134372ddf8991726f482b811c82fc5a687f37ce316f7c3031"></a>

## xcsh_securemesh_site — xcsh_securemesh_site / c58701380656 / 2

Breadcrumbs:

- xcsh_securemesh_site

Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with
distributed security.

<a id="canonical-8e4c64bbb6e7bda6a9b0025209bcf54ecca1fc3f477177de6b560f9fc26d6874"></a>

## Prerequisites — xcsh_securemesh_site / c58701380656 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3f5aa4d6cd42207948d4494e005bc8e21f087018c5d888dd870af69779510edd"></a>

## Minimal configuration — xcsh_securemesh_site / c58701380656 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSite by name
data "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"
}

output "securemesh_site_id" {
  value = data.xcsh_securemesh_site.example.id
}
```

<a id="canonical-adff79a8438f121bf2ad336be6d85431290a510c5a6eb5c6b853f3d2d98f5f2a"></a>

## Root configuration — xcsh_securemesh_site / c58701380656 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-338bd004b50a1a184c4fdf0ce9549eb3b73cfd283cc39bffb017e2189bfb879d"></a>

## Next pages — xcsh_securemesh_site / c58701380656 / 6

- [Property reference](../guides/data-sources--securemesh_site--reference--group-001.md#canonical-70f2b876f8c6ed3295459dbe10e8713248e8af1fefb7e4604690cfecd9889552)
- [Examples](../guides/data-sources--securemesh_site--examples--group-001.md#canonical-e9fdcf3f99896d5d2d8fc3260c18902073b98d1b6b84881b2fcdb33ae6a33ade)
