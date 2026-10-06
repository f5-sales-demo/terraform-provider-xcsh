---
page_title: "xcsh_rate_limiter_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy examples."
---

# xcsh_rate_limiter_policy examples

<a id="canonical-2122331121300203-0101013233100031-3100200332130301-0010100133112022-0230133200320122-3111332320321112-0213002131311032-2213330231232133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- Examples

<a id="canonical-1322011113201332-2111302100010121-0210011000032101-3132232003211030-2033202001133102-0220003121012311-1022112122322300-1022320222033331"></a>

### Complete configurations for `xcsh_rate_limiter_policy`

- [Data source](data-sources--rate_limiter_policy--examples--group-001.md#canonical-3232132001302223-0202101231322303-2121003002023330-2032000102311231-3221010313312332-1103122100323211-1232301303201101-1303112221111113): valid configuration.

<a id="canonical-3232132001302223-0202101231322303-2121003002023330-2032000102311231-3221010313312332-1103122100323211-1232301303201101-1303112221111113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Examples](data-sources--rate_limiter_policy--examples--group-001.md#canonical-2122331121300203-0101013233100031-3100200332130301-0010100133112022-0230133200320122-3111332320321112-0213002131311032-2213330231232133)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_rate_limiter_policy/data-source.tf`; digest `sha256:631ebdd1844abfdc3c403666a18c881b536cd51ce00cfb0a5d316458e5a24bba`.

```terraform
# RateLimiterPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiterPolicy by name
data "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}

output "rate_limiter_policy_id" {
  value = data.xcsh_rate_limiter_policy.example.id
}
```
