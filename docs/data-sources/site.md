---
page_title: "xcsh_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site landing."
---

# xcsh_site landing

<a id="canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d96a51b99879e8b656ba87f4d6182243bb503488911696fc5f2d257d85980af"></a>

## xcsh_site — xcsh_site / e3babb2bd3fc / 2

Breadcrumbs:

- xcsh_site

Manages a Site resource in F5 Distributed Cloud for get of site. configuration. (read-only data
source)

<a id="canonical-adc0c5ed77cc38fc2d0ca854e76a4fe5852005aba9ef4862d989061e89941668"></a>

## Prerequisites — xcsh_site / e3babb2bd3fc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `virtual_site`.

- virtual_site: Logical grouping of physical sites

<a id="canonical-84e7e9c8b87a56bc6fc1c9669efb7d4fc13d3fc7a5718a0f37d92cbf419f56a8"></a>

## Minimal configuration — xcsh_site / e3babb2bd3fc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Site Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Site by name
data "xcsh_site" "example" {
  name      = "example-site"
  namespace = "staging"
}

output "site_id" {
  value = data.xcsh_site.example.id
}
```

<a id="canonical-75077b3e28272c881acb1012ad7d314a9c9efb0adfe78176edf9bd138a718b8a"></a>

## Root configuration — xcsh_site / e3babb2bd3fc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-66709d5f6a3d4084599ab9827630827f9a4892b47be226b98c179f1114b99db1"></a>

## Next pages — xcsh_site / e3babb2bd3fc / 6

- [Property reference](../guides/data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [Examples](../guides/data-sources--site--examples--group-001.md#canonical-5dcb0b1f7736d456c2ca9b82b4d4af76c43b08bfe5a9a9e4f3f086ed99c8371a)
