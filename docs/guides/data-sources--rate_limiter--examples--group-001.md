---
page_title: "xcsh_rate_limiter examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter examples."
---

# xcsh_rate_limiter examples

<a id="canonical-3102222211011003-3321132003022122-0230231131030120-1213200131331313-3100321221330010-0330022033211222-3201000211000203-1131232131111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- Examples

<a id="canonical-1212222313011310-1333210232212322-0022331003222011-1332230232113123-3023032220023210-1121030022323331-2311113100222010-1213112212121202"></a>

### Complete configurations for `xcsh_rate_limiter`

- [Data source](data-sources--rate_limiter--examples--group-001.md#canonical-0022322200220222-0122130010010231-1112213222011300-0122322033120111-0001201000200003-2003010003202010-2223332233120323-2000120213332103): valid configuration.

<a id="canonical-0022322200220222-0122130010010231-1112213222011300-0122322033120111-0001201000200003-2003010003202010-2223332233120323-2000120213332103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130)
- [Examples](data-sources--rate_limiter--examples--group-001.md#canonical-3102222211011003-3321132003022122-0230231131030120-1213200131331313-3100321221330010-0330022033211222-3201000211000203-1131232131111232)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_rate_limiter/data-source.tf`; digest `sha256:6445461d09bf17fe28d100ae80d4a2ae205fe7e6d895306172e50a8556fc3fc6`.

```terraform
# RateLimiter Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiter by name
data "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}

output "rate_limiter_id" {
  value = data.xcsh_rate_limiter.example.id
}
```
