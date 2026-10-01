---
page_title: "xcsh_allowed_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain landing."
---

# xcsh_allowed_domain landing

<a id="canonical-cd1b8e8b9763ec41046ec5b23378de3e4eb43cba5cd5576863af43284e5075e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0db7f7fbefd8faebff81b09cc4bdfd0b9cf8a1586e6f8a5afe052d8f03797a6"></a>

## xcsh_allowed_domain — xcsh_allowed_domain / 09e690ac3e38 / 2

Breadcrumbs:

- xcsh_allowed_domain

Manages allowed domain in F5 Distributed Cloud.

<a id="canonical-31dfafd5be70444ae1e57d0910646de1df3ccac0a266ca56a70dda46e82f58a2"></a>

## Prerequisites — xcsh_allowed_domain / 09e690ac3e38 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e051185a2df2b9b3efeb973ccc5e2de013ca9fdaf8963f8a4ebda478594b976e"></a>

## Minimal configuration — xcsh_allowed_domain / 09e690ac3e38 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
```

<a id="canonical-18612dc9105d925976d6c6ba0ff571753a256552c575fee8022da025581ae07a"></a>

## Root configuration — xcsh_allowed_domain / 09e690ac3e38 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-097b8455258a8ae2fbc868e862567dea772e040f1c9d4c9ed25d1bac0b5287da"></a>

## Next pages — xcsh_allowed_domain / 09e690ac3e38 / 6

- [Property reference](../guides/data-sources--allowed_domain--reference--group-001.md#canonical-814fabc4d196c20da75111723783348fc11b0934b815355a67be6505acee85a4)
- [Examples](../guides/data-sources--allowed_domain--examples--group-001.md#canonical-72efddb7d276e7706910b3e4eef629cc5c74854545163ae33946456ab3c63799)
