---
page_title: "xcsh_bgp examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp examples."
---

# xcsh_bgp examples

<a id="canonical-05fe067b621ea11533193eb95e9bf3cd816189ec172eb71a345b4eb582563043"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c91a9fcce9f0ef316f7a500601cfea5ddc2a57b699d1aab1fb39ae0fac620377"></a>

## Examples — Examples / 86d1a9b9a296 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- Examples

<a id="canonical-1d6d32ff531142980a36c04958896f56e6fc7857dc3140efc03da26ef187abb0"></a>

## Complete configurations — Examples / 86d1a9b9a296 / 3

- [Data source](data-sources--bgp--examples--group-001.md#canonical-39aebbd936f20bcd5c44c28cbb5668192359c1d5d50462e8badce21321483532): valid configuration.

<a id="canonical-d86f19cf08a0ea25e0a7061b39716f8641b7bd870a0da6738c2584262d91077d"></a>

## Next pages — Examples / 86d1a9b9a296 / 4

- [Data source](data-sources--bgp--examples--group-001.md#canonical-39aebbd936f20bcd5c44c28cbb5668192359c1d5d50462e8badce21321483532)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)

<a id="canonical-39aebbd936f20bcd5c44c28cbb5668192359c1d5d50462e8badce21321483532"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aab1e8895f41cb1c0b8e89ce851c3d985b2af0bdfd7a8f0e8dd1915926eda6c"></a>

## Data source — Data source / 61dfdc2a6396 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
- [Examples](data-sources--bgp--examples--group-001.md#canonical-05fe067b621ea11533193eb95e9bf3cd816189ec172eb71a345b4eb582563043)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp/data-source.tf`; digest `sha256:0e624b6da5c866c4f991e89616b34b5f99775ffa87bb8832e89714aa1bc2809b`.

```terraform
# BGP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGP by name
data "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}

output "bgp_id" {
  value = data.xcsh_bgp.example.id
}
```

<a id="canonical-e0b089ba56565ae968df7cb713e3d05a867f059fe7419a92f3e6c5378db6c133"></a>

## Next pages — Data source / 61dfdc2a6396 / 3

- [Examples](data-sources--bgp--examples--group-001.md#canonical-05fe067b621ea11533193eb95e9bf3cd816189ec172eb71a345b4eb582563043)
- [xcsh_bgp](../data-sources/bgp.md#canonical-3cb34ff3923c8b5bb36d1ac6e7fd3b4e63e06f78cdd4fc62c8d2c79105acc5ad)
