---
page_title: "xcsh_user_identification examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification examples."
---

# xcsh_user_identification examples

<a id="canonical-be2d29176d9518bab3b4e19d382e25c3efaaacb748fe2e85001c9c4043ccc81b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b7ee3f5046f3175010c79e4d397a8dc42e290a5b038fb462a09cabfd6f526b1"></a>

## Examples — Examples / 78149fd01c1a / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- Examples

<a id="canonical-53b32e4e77c5a2c39355aa8a6b81212cb719c57fef99b7ab2430140cd1880774"></a>

## Complete configurations — Examples / 78149fd01c1a / 3

- [Data source](data-sources--user_identification--examples--group-001.md#canonical-57188526e3513cf6843eb537c7b6565ee4c30bfaf025194b58de62e83f64c67d): valid configuration.

<a id="canonical-50723c6971803f10a2430e0af23caf7a1bf362039624c860b31a5d6285c47274"></a>

## Next pages — Examples / 78149fd01c1a / 4

- [Data source](data-sources--user_identification--examples--group-001.md#canonical-57188526e3513cf6843eb537c7b6565ee4c30bfaf025194b58de62e83f64c67d)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-57188526e3513cf6843eb537c7b6565ee4c30bfaf025194b58de62e83f64c67d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aa4d1f538e602868468e54796682409d656267ac6614ad7c0cfff97797b41e3"></a>

## Data source — Data source / 000e5c9b9edd / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Examples](data-sources--user_identification--examples--group-001.md#canonical-be2d29176d9518bab3b4e19d382e25c3efaaacb748fe2e85001c9c4043ccc81b)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_user_identification/data-source.tf`; digest `sha256:c30020dfd6aecd49ffb0273704b9b9e1c85f494c0e9e28cd883c479b6b6234a3`.

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

<a id="canonical-8ff6bf79ffb140d5dfecbfbd00b84f523c2252d2dce457f3aaf7d350eb6073a6"></a>

## Next pages — Data source / 000e5c9b9edd / 3

- [Examples](data-sources--user_identification--examples--group-001.md#canonical-be2d29176d9518bab3b4e19d382e25c3efaaacb748fe2e85001c9c4043ccc81b)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
