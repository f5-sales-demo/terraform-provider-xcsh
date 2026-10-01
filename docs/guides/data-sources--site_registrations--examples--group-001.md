---
page_title: "xcsh_site_registrations examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations examples."
---

# xcsh_site_registrations examples

<a id="canonical-ce2131793ed98a34945fac766d85be28b3f9df7b9ed07d195dbeb37477f644f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3d9fcd5bf5819ead74721d5e7dc349f053119b1a581dad22750c3c8e1b0ca71"></a>

## Examples — Examples / 9be6f8c468ff / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- Examples

<a id="canonical-ec4e45a1a2db29b43008d2ce05c2cb7b87a7c45065d1fe8819390e09cf58360a"></a>

## Complete configurations — Examples / 9be6f8c468ff / 3

- [Data source](data-sources--site_registrations--examples--group-001.md#canonical-4e7b8868f91f4b76297c26de204c1f027c3b2b00370f11c6f6148d98493afb33): valid configuration.

<a id="canonical-5b6d9df9f06f3c32449f2fd562cef153d3e25f19bb704f0bc8b06a5842646ff9"></a>

## Next pages — Examples / 9be6f8c468ff / 4

- [Data source](data-sources--site_registrations--examples--group-001.md#canonical-4e7b8868f91f4b76297c26de204c1f027c3b2b00370f11c6f6148d98493afb33)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-4e7b8868f91f4b76297c26de204c1f027c3b2b00370f11c6f6148d98493afb33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be9f232f9389b67aaae536511655922eb18ae3207bcdbc9a027543a3dfb59223"></a>

## Data source — Data source / 9bc9a9a1ac94 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Examples](data-sources--site_registrations--examples--group-001.md#canonical-ce2131793ed98a34945fac766d85be28b3f9df7b9ed07d195dbeb37477f644f7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations/data-source.tf`; digest `sha256:382726257bf8ed644e4ba85a525a0636702e8ea87e1b309384adf95a431060c9`.

```terraform
# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
```

<a id="canonical-9534c4ddaa2012a7c79489e2826bf13b8e79952e4e9c5e518c4a6c8f4bbd7fe2"></a>

## Next pages — Data source / 9bc9a9a1ac94 / 3

- [Examples](data-sources--site_registrations--examples--group-001.md#canonical-ce2131793ed98a34945fac766d85be28b3f9df7b9ed07d195dbeb37477f644f7)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
