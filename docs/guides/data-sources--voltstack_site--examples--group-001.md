---
page_title: "xcsh_voltstack_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site examples."
---

# xcsh_voltstack_site examples

<a id="canonical-731327ff03d150c7c8d69f459d6f85905909d41ecbd50e308ed1feb174cffb5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-437e1af72230ab18c601b011e5a4bef5682cf70067012fb690ed783b007e72f2"></a>

## Examples — Examples / 110708125393 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- Examples

<a id="canonical-25cc1b263b9cfce0f37c5a95b339d4dc9d26904cdbccdb84bda28d79a69f6fb8"></a>

## Complete configurations — Examples / 110708125393 / 3

- [Data source](data-sources--voltstack_site--examples--group-001.md#canonical-12277f946b3f564a91267d844b3898ed3a528d801ef89b94c864646e599af86f): valid configuration.

<a id="canonical-6386a0a8cc34a448fad166746bd6bbacb0cf3e078b313bd985001582f22bb74a"></a>

## Next pages — Examples / 110708125393 / 4

- [Data source](data-sources--voltstack_site--examples--group-001.md#canonical-12277f946b3f564a91267d844b3898ed3a528d801ef89b94c864646e599af86f)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-12277f946b3f564a91267d844b3898ed3a528d801ef89b94c864646e599af86f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8fe43358e57eb3ee8c3d5eeb2dfe5d4980da5cf0a9b2fd595b5b4c39cd6ca98"></a>

## Data source — Data source / 1920e25f4ee6 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Examples](data-sources--voltstack_site--examples--group-001.md#canonical-731327ff03d150c7c8d69f459d6f85905909d41ecbd50e308ed1feb174cffb5e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_voltstack_site/data-source.tf`; digest `sha256:9d6c1f9af1750a76c0f8df78e6df342564c3af32fa86a487ff12441eece04681`.

```terraform
# VoltstackSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VoltstackSite by name
data "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"
}

output "voltstack_site_id" {
  value = data.xcsh_voltstack_site.example.id
}
```

<a id="canonical-0a605236187255f29bc92ff270c2cfc1ea1119436b4f27a1c68cb46c1769f09f"></a>

## Next pages — Data source / 1920e25f4ee6 / 3

- [Examples](data-sources--voltstack_site--examples--group-001.md#canonical-731327ff03d150c7c8d69f459d6f85905909d41ecbd50e308ed1feb174cffb5e)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
