---
page_title: "xcsh_protected_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain examples."
---

# xcsh_protected_domain examples

<a id="canonical-ccaed106fa00aabd26132c08c45e92c5d9ddd15523136c4c1aa8a1d0612d5b06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-051a28e410dcd1c0eb79fd3158a3ef37e1d4021105e19f8e8395a145a9b64096"></a>

## Examples — Examples / b97b011bae34 / 2

Breadcrumbs:

- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-3e543fadedff8ca807776d33c9261ab932e9a9940a7f7de944cf323d9229140f)
- Examples

<a id="canonical-8185289c34d5b2a85b0badf53b37dd966ada11cfc582b3c5eb1d569937212b00"></a>

## Complete configurations — Examples / b97b011bae34 / 3

- [Data source](data-sources--protected_domain--examples--group-001.md#canonical-3dfdb8a4156d37f4e694eb763e409b26361114996402611b3e6fb8e08a569952): valid configuration.

<a id="canonical-f9f6b48d8e870eac26dfba46a6dc69adc543ca4378daccf5261980f0bee14cec"></a>

## Next pages — Examples / b97b011bae34 / 4

- [Data source](data-sources--protected_domain--examples--group-001.md#canonical-3dfdb8a4156d37f4e694eb763e409b26361114996402611b3e6fb8e08a569952)
- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-3e543fadedff8ca807776d33c9261ab932e9a9940a7f7de944cf323d9229140f)

<a id="canonical-3dfdb8a4156d37f4e694eb763e409b26361114996402611b3e6fb8e08a569952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1efebabf4fdcdb897438a046f4fa1c8af309648545f74e96575e29e62d002703"></a>

## Data source — Data source / c1128bc192b1 / 2

Breadcrumbs:

- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-3e543fadedff8ca807776d33c9261ab932e9a9940a7f7de944cf323d9229140f)
- [Examples](data-sources--protected_domain--examples--group-001.md#canonical-ccaed106fa00aabd26132c08c45e92c5d9ddd15523136c4c1aa8a1d0612d5b06)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protected_domain/data-source.tf`; digest `sha256:54de5a860ec2aa40dfd3370e44fe450b1b713886394cacf655ec2b67b6b0edd5`.

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

<a id="canonical-b19ee8fe544dd653d7384037157afbe16703e5c5f10c97b673cdabd1d6996c1a"></a>

## Next pages — Data source / c1128bc192b1 / 3

- [Examples](data-sources--protected_domain--examples--group-001.md#canonical-ccaed106fa00aabd26132c08c45e92c5d9ddd15523136c4c1aa8a1d0612d5b06)
- [xcsh_protected_domain](../data-sources/protected_domain.md#canonical-3e543fadedff8ca807776d33c9261ab932e9a9940a7f7de944cf323d9229140f)
