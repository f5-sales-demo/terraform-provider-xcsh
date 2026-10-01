---
page_title: "xcsh_cdn_cache_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule examples."
---

# xcsh_cdn_cache_rule examples

<a id="canonical-5bb95f551d7b0fd1f82574637552008edccd3cf016caf1cf657dbfc5e5b9b9ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c926b25d5bd757c075dbf7c2a6262ca24e77b8889f64e04432212a4cc7177bc"></a>

## Examples — Examples / f8630e7a5f82 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- Examples

<a id="canonical-938777dffe03c7a3744dde9dd586bcc4f8f3f4f63972e2f97aa76c899fe71695"></a>

## Complete configurations — Examples / f8630e7a5f82 / 3

- [Data source](data-sources--cdn_cache_rule--examples--group-001.md#canonical-dd78befcad06da9b6c29e4cdeb35b111cf0b17ca26ef8ae2dad64ab24754df9b): valid configuration.

<a id="canonical-9e850c28a516940f7a2642e1f66b0d7719278ccd46085329560ddd245eaec802"></a>

## Next pages — Examples / f8630e7a5f82 / 4

- [Data source](data-sources--cdn_cache_rule--examples--group-001.md#canonical-dd78befcad06da9b6c29e4cdeb35b111cf0b17ca26ef8ae2dad64ab24754df9b)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)

<a id="canonical-dd78befcad06da9b6c29e4cdeb35b111cf0b17ca26ef8ae2dad64ab24754df9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e858bbd7a150d030f280787110b6fa3c975eb9583f7965b5cb2c3cdef1fa6e3"></a>

## Data source — Data source / da7ce2dc835d / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
- [Examples](data-sources--cdn_cache_rule--examples--group-001.md#canonical-5bb95f551d7b0fd1f82574637552008edccd3cf016caf1cf657dbfc5e5b9b9ef)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_cache_rule/data-source.tf`; digest `sha256:fc686ceb8d4fdca2d4b78094b5b3751ef65d5c692bd05f9089c3561d05468ad9`.

```terraform
# CDNCacheRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNCacheRule by name
data "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}

output "cdn_cache_rule_id" {
  value = data.xcsh_cdn_cache_rule.example.id
}
```

<a id="canonical-a00f78332e900bf3b81fa2db869c15e57a34abe7a6ae88f828bbfd1c34b2ad69"></a>

## Next pages — Data source / da7ce2dc835d / 3

- [Examples](data-sources--cdn_cache_rule--examples--group-001.md#canonical-5bb95f551d7b0fd1f82574637552008edccd3cf016caf1cf657dbfc5e5b9b9ef)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-727b5aab8761041dba2c3d26acd033e191cb47a8d8343b9c5c9ea22286e70589)
