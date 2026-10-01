---
page_title: "xcsh_fast_acl landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl landing."
---

# xcsh_fast_acl landing

<a id="canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08ab6eb5b8649383f7a0414ef50f5b8055c2bb98648bf25b1cf052b5dd935b84"></a>

## xcsh_fast_acl — xcsh_fast_acl / e6e8031fb93f / 2

Breadcrumbs:

- xcsh_fast_acl

Manages object, object contains rules to protect site from denial of service It has
destination\{destination IP, destination port) and references to in F5 Distributed Cloud.

<a id="canonical-7aeac6f035803be9b4f0ebf9e1512a04d8e1ca5761e8d373cabe7b423ed2d3fa"></a>

## Prerequisites — xcsh_fast_acl / e6e8031fb93f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1f1b4b9e51f4d87380f1f5df6a56ddbe76150ffc563dd1b41e5eff8dfe9e5a23"></a>

## Minimal configuration — xcsh_fast_acl / e6e8031fb93f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```

<a id="canonical-3916342f3345adae271f93ed2f4713837c4fbb3e7f154649f292e9c143f018c5"></a>

## Root configuration — xcsh_fast_acl / e6e8031fb93f / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-953606decc33e49f3a521c3ddb0402efb8a2eefd57a7176915560464a8556bee"></a>

## Next pages — xcsh_fast_acl / e6e8031fb93f / 6

- [Property reference](../guides/data-sources--fast_acl--reference--group-001.md#canonical-d091f7959c7f948a310039fc3d26b683c81c8576f27196686c6e4e997cece7f5)
- [Examples](../guides/data-sources--fast_acl--examples--group-001.md#canonical-6f43747349ebe04c37600937c82d7383a504fe9eb260f386303fa0fec99920fb)
