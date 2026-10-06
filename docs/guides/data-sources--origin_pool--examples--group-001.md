---
page_title: "xcsh_origin_pool examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool examples."
---

# xcsh_origin_pool examples

<a id="canonical-0212331131013132-1231233202011013-2201121321202120-2001300222023000-2201012103301302-3100013202031002-3021022030203221-0030001203321132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- Examples

<a id="canonical-0310330223110023-1130333023122100-1300022013320010-3123101131130231-1002213301220030-3000211130020223-1010223213100303-1331102031123220"></a>

### Complete configurations for `xcsh_origin_pool`

- [Data source](data-sources--origin_pool--examples--group-001.md#canonical-3122032320030203-3303233222333212-2330222003120311-3332022103233112-3111101123220313-3313010302313031-1303233000223200-0202121233002013): valid configuration.

<a id="canonical-3122032320030203-3303233222333212-2330222003120311-3332022103233112-3111101123220313-3313010302313031-1303233000223200-0202121233002013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Examples](data-sources--origin_pool--examples--group-001.md#canonical-0212331131013132-1231233202011013-2201121321202120-2001300222023000-2201012103301302-3100013202031002-3021022030203221-0030001203321132)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_origin_pool/data-source.tf`; digest `sha256:cae409ea2c1cde6f7ffac297039f556a77e9684a94017d350d26cb02cd85782a`.

```terraform
# OriginPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing OriginPool by name
data "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}

output "origin_pool_id" {
  value = data.xcsh_origin_pool.example.id
}
```
