---
page_title: "xcsh_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery examples."
---

# xcsh_discovery examples

<a id="canonical-af5f4883c598ee98395114fc54a3127f97b306ab0e56ee5dcd42b30eefd63773"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8215d650573441b8a97fa484d13bef4622c63d0c96478a55a711cd3785b1b227"></a>

## Examples — Examples / 73432a35dc7f / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- Examples

<a id="canonical-1a21541ff41b5c1b91b6875089acedb6e9af2c061a1a31c0eb9b176e7072f8fc"></a>

## Complete configurations — Examples / 73432a35dc7f / 3

- [Data source](data-sources--discovery--examples--group-001.md#canonical-593add88bbdd8696cf47e5129abfcba5bc4a5fbac83ef6e7992e5cf61830016d): valid configuration.

<a id="canonical-0073c9e21730aeb56c815174abd7a600babf9a5cd603ffa8ea767fec708a93a6"></a>

## Next pages — Examples / 73432a35dc7f / 4

- [Data source](data-sources--discovery--examples--group-001.md#canonical-593add88bbdd8696cf47e5129abfcba5bc4a5fbac83ef6e7992e5cf61830016d)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)

<a id="canonical-593add88bbdd8696cf47e5129abfcba5bc4a5fbac83ef6e7992e5cf61830016d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-868091c26a88ff3bc8a7d49ddef3f1139e5a148299b96e6e97091efab37b18ac"></a>

## Data source — Data source / c793ced1e07e / 2

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
- [Examples](data-sources--discovery--examples--group-001.md#canonical-af5f4883c598ee98395114fc54a3127f97b306ab0e56ee5dcd42b30eefd63773)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_discovery/data-source.tf`; digest `sha256:de2200d3e23389756392f1e95f9b7544c16ad111ecd7a837e1de4f6caed18861`.

```terraform
# Discovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Discovery by name
data "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}

output "discovery_id" {
  value = data.xcsh_discovery.example.id
}
```

<a id="canonical-5d48df9598251662945fb174009f0a9dba461c9ccf4ba792d412e06730632357"></a>

## Next pages — Data source / c793ced1e07e / 3

- [Examples](data-sources--discovery--examples--group-001.md#canonical-af5f4883c598ee98395114fc54a3127f97b306ab0e56ee5dcd42b30eefd63773)
- [xcsh_discovery](../data-sources/discovery.md#canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea)
