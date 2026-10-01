---
page_title: "xcsh_site_registrations_by_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_site examples."
---

# xcsh_site_registrations_by_site examples

<a id="canonical-9dfe40fa0db7c48dfa71d784e8447fe909be11dbb249be1956696f9f35cef12f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-963606aede31a5d1122675b7d9c7be3cc3ab58b275759f8bedf3df3594080056"></a>

## Examples — Examples / 1ea01003edab / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- Examples

<a id="canonical-0f7746016bcbf735ebf878310ada7abc6f95858b653bd35227efa0599c9cdfd4"></a>

## Complete configurations — Examples / 1ea01003edab / 3

- [Data source](data-sources--site_registrations_by_site--examples--group-001.md#canonical-fdf84bd88785e416e322b96555ffc9930eb6461c0f2ffa1be7ce6b1a1f43d5b5): valid configuration.

<a id="canonical-17f16a01ff4ce83ee00cdeb209f9aba08d54cf90b57f790cf7786cd21c1f6015"></a>

## Next pages — Examples / 1ea01003edab / 4

- [Data source](data-sources--site_registrations_by_site--examples--group-001.md#canonical-fdf84bd88785e416e322b96555ffc9930eb6461c0f2ffa1be7ce6b1a1f43d5b5)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-fdf84bd88785e416e322b96555ffc9930eb6461c0f2ffa1be7ce6b1a1f43d5b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cdec3aba6bdaa7ae712d5e5353e1cb971847cf475481406a672c863240cad31"></a>

## Data source — Data source / 806645140358 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Examples](data-sources--site_registrations_by_site--examples--group-001.md#canonical-9dfe40fa0db7c48dfa71d784e8447fe909be11dbb249be1956696f9f35cef12f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations_by_site/data-source.tf`; digest `sha256:c5f91ec913e86f87149f588895f436b17703e7ae2ae9655f9226a6f3b281b896`.

```terraform
# SiteRegistrationsBySite DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_site" "example" {
  site_name = "example-value"
}

output "site_registrations_by_site_result" {
  value = data.xcsh_site_registrations_by_site.example
}
```

<a id="canonical-187a408fe680ec4ce3c7ba0ca892e39b42a62717bdc361917dd4b1d494690597"></a>

## Next pages — Data source / 806645140358 / 3

- [Examples](data-sources--site_registrations_by_site--examples--group-001.md#canonical-9dfe40fa0db7c48dfa71d784e8447fe909be11dbb249be1956696f9f35cef12f)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
