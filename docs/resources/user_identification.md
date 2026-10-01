---
page_title: "xcsh_user_identification landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification landing."
---

# xcsh_user_identification landing

<a id="canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0c6d0c28d0511c211a552842644d674aae83c406124ca70d2b58994251691ea"></a>

## xcsh_user_identification — xcsh_user_identification / 2e8e3f0008ec / 2

Breadcrumbs:

- xcsh_user_identification

Manages user\_identification creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-0bead206a50d5f32612afb438b3bf98ecfa94f53cbd7da1efadd80bf66dc58be"></a>

## Prerequisites — xcsh_user_identification / 2e8e3f0008ec / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-70199b8e22ad7d3b4fc056b25310f194d3e0c661a1323f9768d3f2f7efe45d32"></a>

## Minimal configuration — xcsh_user_identification / 2e8e3f0008ec / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UserIdentification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UserIdentification configuration
resource "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}
```

<a id="canonical-46b282b98f70cb6530ddc08a02b575cffbb7ec03bf2a7bfb20ff3f50fce7897a"></a>

## Root configuration — xcsh_user_identification / 2e8e3f0008ec / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ccaae688f68cee73ef070eaa3a12546131b774b503711d17c65450670dcbc0e8"></a>

## Next pages — xcsh_user_identification / 2e8e3f0008ec / 6

- [Property reference](../guides/resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [Examples](../guides/resources--user_identification--examples--group-001.md#canonical-5e1a761927919765e78808a5740beea6304ea7c033514dc585440689af13116e)
- [Import](../guides/resources--user_identification--lifecycle--group-001.md#canonical-10449741f8c1f3a2446923944269989374112bdfc082e03fd431af06b0c2bf6c)
- [Timeouts](../guides/resources--user_identification--lifecycle--group-001.md#canonical-204d4a6732481533225628a2d1d1f2009a6c85b61c5fed835b5199bbbef1b8b4)
