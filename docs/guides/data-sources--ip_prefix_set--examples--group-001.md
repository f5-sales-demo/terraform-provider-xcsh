---
page_title: "xcsh_ip_prefix_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set examples."
---

# xcsh_ip_prefix_set examples

<a id="canonical-5100a4a30085b44b789ae0073db4bba5a4ec81d101479bf4216798464fee5841"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3beaeba66ad87b65c72f75c883ad616067c309ab2bac785057d9c303c93dafa"></a>

## Examples — Examples / 88d278f00df6 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)
- Examples

<a id="canonical-580f22ecb725c2db6eb2fb7bc92962a1f39ce9fdb1abd75b53deaafb45c5583a"></a>

## Complete configurations — Examples / 88d278f00df6 / 3

- [Data source](data-sources--ip_prefix_set--examples--group-001.md#canonical-1a3b433b9c851819fa75f50b9d36c158eccc5e329484ca8f7702e482a0349200): valid configuration.

<a id="canonical-567e2a9cf20ce17822d067a9c7559fe482c96924e1f68dad95303155963f5130"></a>

## Next pages — Examples / 88d278f00df6 / 4

- [Data source](data-sources--ip_prefix_set--examples--group-001.md#canonical-1a3b433b9c851819fa75f50b9d36c158eccc5e329484ca8f7702e482a0349200)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)

<a id="canonical-1a3b433b9c851819fa75f50b9d36c158eccc5e329484ca8f7702e482a0349200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b42634e5e9ea4993d9e9e8870db5b37f1857111104d99411638c7fb35928fa9"></a>

## Data source — Data source / 6786c7c089f2 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)
- [Examples](data-sources--ip_prefix_set--examples--group-001.md#canonical-5100a4a30085b44b789ae0073db4bba5a4ec81d101479bf4216798464fee5841)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ip_prefix_set/data-source.tf`; digest `sha256:e21d61ff8d692544ea192645c0aec9c60f1857a80969caf8df2d970c69dd061b`.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```

<a id="canonical-40ad7a72d01a24e3ec8609b5917561b742f3120fad2696d3e858ce236c7ce141"></a>

## Next pages — Data source / 6786c7c089f2 / 3

- [Examples](data-sources--ip_prefix_set--examples--group-001.md#canonical-5100a4a30085b44b789ae0073db4bba5a4ec81d101479bf4216798464fee5841)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc)
