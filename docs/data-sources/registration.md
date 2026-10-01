---
page_title: "xcsh_registration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration landing."
---

# xcsh_registration landing

<a id="canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bef301e08a9664e0930eef5d62840ba40fb31d94a46fa39f2e63603e00071c1"></a>

## xcsh_registration — xcsh_registration / ec73786d331a / 2

Breadcrumbs:

- xcsh_registration

Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this
message, never used by users. configuration.

<a id="canonical-53901e7ac1f8a9ea94da9a15cbae1c7e7c23d2eeaca7a9b0b958dbd9b4e24ede"></a>

## Prerequisites — xcsh_registration / ec73786d331a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-caefeda44a5d3c6228c82a43b82be2436b1847a49df5969064c4b9539bdb1c4c"></a>

## Minimal configuration — xcsh_registration / ec73786d331a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Registration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Registration by name
data "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"
}

output "registration_id" {
  value = data.xcsh_registration.example.id
}
```

<a id="canonical-c159c17d3f0042561f6fc9a91468ad358d693cfd2d5291eef93ea4ab1ba443c4"></a>

## Root configuration — xcsh_registration / ec73786d331a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3df3cb2de0727668a0d7fa825c6abd816c71fd00789e3bc367cd562406e5dfe7"></a>

## Next pages — xcsh_registration / ec73786d331a / 6

- [Property reference](../guides/data-sources--registration--reference--group-001.md#canonical-515aab2ff4416a2f1aa445644e411a2400e1d40ad9d313746445e68b449bf4c6)
- [Examples](../guides/data-sources--registration--examples--group-001.md#canonical-308dd2706c143d052513a67a5ad0066da2434e408835e3cde8ea716d8d0bf9f7)
