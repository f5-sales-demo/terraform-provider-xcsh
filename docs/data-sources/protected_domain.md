---
page_title: "xcsh_protected_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain landing."
---

# xcsh_protected_domain landing

<a id="canonical-3e543fadedff8ca807776d33c9261ab932e9a9940a7f7de944cf323d9229140f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c2d8ff489a88fd564a3cb09865386466c6d2c62fafc5677a4d45a9857d2f4c3"></a>

## xcsh_protected_domain — xcsh_protected_domain / de3d71f633ed / 2

Breadcrumbs:

- xcsh_protected_domain

Manages Domain to protect in F5 Distributed Cloud.

<a id="canonical-1e05065aaeb33dfa7fea65c52f27af08462873e576677e190c6a312e84dcce3b"></a>

## Prerequisites — xcsh_protected_domain / de3d71f633ed / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c696a392e16c0902292f940418722205813f0decf9d449feeb245ee23c79bef5"></a>

## Minimal configuration — xcsh_protected_domain / de3d71f633ed / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedDomain by name
data "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"
}

output "protected_domain_id" {
  value = data.xcsh_protected_domain.example.id
}
```

<a id="canonical-7de40e4a7c372748f4334392030436892c9f722fc65b23a7fa1ae56f72e1fa06"></a>

## Root configuration — xcsh_protected_domain / de3d71f633ed / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-435f0135e999e7b611de783a188d7f71d983f1a6db7ce815b1cade70144007e0"></a>

## Next pages — xcsh_protected_domain / de3d71f633ed / 6

- [Property reference](../guides/data-sources--protected_domain--reference--group-001.md#canonical-109a5c5b1d87fd443e19544e636c3a201e0ebb9189d96c51395f44228e3d7f1b)
- [Examples](../guides/data-sources--protected_domain--examples--group-001.md#canonical-ccaed106fa00aabd26132c08c45e92c5d9ddd15523136c4c1aa8a1d0612d5b06)
