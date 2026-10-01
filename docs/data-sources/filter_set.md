---
page_title: "xcsh_filter_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set landing."
---

# xcsh_filter_set landing

<a id="canonical-e8e33e1d1dd12455d2f08a08ea045d5f0ebbbb80a4ffd1555414d839959cfb01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8a41edbe67788eb173157ea11dd766ef53a30e64eb6fe27cd2c95d0d3199156"></a>

## xcsh_filter_set — xcsh_filter_set / 93b0b478e250 / 2

Breadcrumbs:

- xcsh_filter_set

Manages specification in F5 Distributed Cloud.

<a id="canonical-abb32287a07f829551654cd57ed9c48d4efcc9a34af255364fe3fcc969f77c0f"></a>

## Prerequisites — xcsh_filter_set / 93b0b478e250 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-dd568be3c62cd4488905fe71e61f55e0a9eba9dbb7cd3aaa7e79c1e06210b208"></a>

## Minimal configuration — xcsh_filter_set / 93b0b478e250 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0eee6348811cf2a233bde962bbc24865fed62b0051d056f1615b0211124fda03"></a>

## Root configuration — xcsh_filter_set / 93b0b478e250 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0aa60949ecc355059eaaee7d1cc15bac648b7d9d6815d38eed05f316342e0e3b"></a>

## Next pages — xcsh_filter_set / 93b0b478e250 / 6

- [Property reference](../guides/data-sources--filter_set--reference--group-001.md#canonical-d3a2053da6cf7c53a8d6afe3a753166143941836e72ebaa16ebb086119667599)
- [Examples](../guides/data-sources--filter_set--examples--group-001.md#canonical-69f2721702bd33ec6260a574548f5a1651fdf4db913b98cc7442399524e18a89)
