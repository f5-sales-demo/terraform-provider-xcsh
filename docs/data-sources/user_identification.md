---
page_title: "xcsh_user_identification landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification landing."
---

# xcsh_user_identification landing

<a id="canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45552679873e5b763c2cea43a73f19988378468305215a583a717d19c6ccd6c9"></a>

## xcsh_user_identification — xcsh_user_identification / cc04ca9ed2cd / 2

Breadcrumbs:

- xcsh_user_identification

Manages user\_identification creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-d6e17f6934384d5ecf8d7f7f73267cec0cba054f3a23a9bfe917d1812647edb1"></a>

## Prerequisites — xcsh_user_identification / cc04ca9ed2cd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a5f50baf50d167113e0ea92153faf70e607dd7b0308b7f78364952523573727d"></a>

## Minimal configuration — xcsh_user_identification / cc04ca9ed2cd / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UserIdentification Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UserIdentification by name
data "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}

output "user_identification_id" {
  value = data.xcsh_user_identification.example.id
}
```

<a id="canonical-5f27e27baf91390c9a7f4db0bd4655b594d25f349e516810d44479d3d59e8dd4"></a>

## Root configuration — xcsh_user_identification / cc04ca9ed2cd / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d4cc3f7ee93b6915e01cea4af78ee0fdc9e1a306359c4410fdc1ecf38e7b00f4"></a>

## Next pages — xcsh_user_identification / cc04ca9ed2cd / 6

- [Property reference](../guides/data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [Examples](../guides/data-sources--user_identification--examples--group-001.md#canonical-be2d29176d9518bab3b4e19d382e25c3efaaacb748fe2e85001c9c4043ccc81b)
