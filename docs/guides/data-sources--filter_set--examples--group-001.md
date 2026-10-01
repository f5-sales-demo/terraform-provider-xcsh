---
page_title: "xcsh_filter_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set examples."
---

# xcsh_filter_set examples

<a id="canonical-69f2721702bd33ec6260a574548f5a1651fdf4db913b98cc7442399524e18a89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6c9cf22c8a9b9506a873acc2b7b5c0770deecd256648797caed7e3bdd91526c"></a>

## Examples — Examples / b85960d28880 / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- Examples

<a id="canonical-dd551913ca1752017a2298ce298e8d30f6cc05396b73adb2e609c6ef339c502d"></a>

## Complete configurations — Examples / b85960d28880 / 3

- [Data source](data-sources--filter_set--examples--group-001.md#canonical-98d1bb9d44561dc2aa58a0337688727c998ab214f72a6f1d56b2477585136dfa): valid configuration.

<a id="canonical-e0c042d8418f3f92f1accade9fec5ef0c02d596b8f4bbcbd76d8aa17c50e82ae"></a>

## Next pages — Examples / b85960d28880 / 4

- [Data source](data-sources--filter_set--examples--group-001.md#canonical-98d1bb9d44561dc2aa58a0337688727c998ab214f72a6f1d56b2477585136dfa)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)

<a id="canonical-98d1bb9d44561dc2aa58a0337688727c998ab214f72a6f1d56b2477585136dfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0195f8be09e5dbbc6672f403255cb935847fe78ff05002020e83aae5cdbe0fa9"></a>

## Data source — Data source / 30a4ecdcbf25 / 2

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
- [Examples](data-sources--filter_set--examples--group-001.md#canonical-69f2721702bd33ec6260a574548f5a1651fdf4db913b98cc7442399524e18a89)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_filter_set/data-source.tf`; digest `sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d`.

```terraform
# FilterSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FilterSet by name
data "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"
}

output "filter_set_id" {
  value = data.xcsh_filter_set.example.id
}
```

<a id="canonical-a8a6755bb911f9069f7e29ec660a1421c289f42d487bdf5de55ebd375af6df04"></a>

## Next pages — Data source / 30a4ecdcbf25 / 3

- [Examples](data-sources--filter_set--examples--group-001.md#canonical-69f2721702bd33ec6260a574548f5a1651fdf4db913b98cc7442399524e18a89)
- [xcsh_filter_set](../data-sources/filter_set.md#canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01)
