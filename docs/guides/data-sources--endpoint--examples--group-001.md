---
page_title: "xcsh_endpoint examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint examples."
---

# xcsh_endpoint examples

<a id="canonical-b10c566b27db1e8b6a31ee834ea7dbdeee24513c193a9ead9e8c97bbdddcba7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64bf1f107658b15fa91c2d6671d6ae34ee05a86d991af64a3be5561c00c77baa"></a>

## Examples — Examples / 055d52ed4a51 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- Examples

<a id="canonical-b095cd3cdb29dd80c41c784d9fdfa494175cb446b228f58bc1ffde100a77bae2"></a>

## Complete configurations — Examples / 055d52ed4a51 / 3

- [Data source](data-sources--endpoint--examples--group-001.md#canonical-212be56da55839318036d8033029d7ef18759021e605b43c97a0636ec9e8df8a): valid configuration.

<a id="canonical-e227857faf013b1babe0c2e989bc0ac12450af45fbd3a05d96c1c7bdd6100074"></a>

## Next pages — Examples / 055d52ed4a51 / 4

- [Data source](data-sources--endpoint--examples--group-001.md#canonical-212be56da55839318036d8033029d7ef18759021e605b43c97a0636ec9e8df8a)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)

<a id="canonical-212be56da55839318036d8033029d7ef18759021e605b43c97a0636ec9e8df8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c8f9bef7d7167cf3d963f493cfc4189743903ee7d80643c9b60fc4581490327"></a>

## Data source — Data source / eac92aeda616 / 2

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
- [Examples](data-sources--endpoint--examples--group-001.md#canonical-b10c566b27db1e8b6a31ee834ea7dbdeee24513c193a9ead9e8c97bbdddcba7f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_endpoint/data-source.tf`; digest `sha256:995586c12b63c50200cc6d03f85d5a9e36c9682dfdd230f380379061443ed02b`.

```terraform
# Endpoint Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Endpoint by name
data "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}

output "endpoint_id" {
  value = data.xcsh_endpoint.example.id
}
```

<a id="canonical-a67e57d6d0be0404140208aec54a7ffb4478dc4660f78576317fa13fdcc108b9"></a>

## Next pages — Data source / eac92aeda616 / 3

- [Examples](data-sources--endpoint--examples--group-001.md#canonical-b10c566b27db1e8b6a31ee834ea7dbdeee24513c193a9ead9e8c97bbdddcba7f)
- [xcsh_endpoint](../data-sources/endpoint.md#canonical-6d5a0ef1abda847b165d88a4c5ee11a8b4e789245312baf595e9fd561e4893ec)
