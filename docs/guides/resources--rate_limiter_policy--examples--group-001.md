---
page_title: "xcsh_rate_limiter_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy examples."
---

# xcsh_rate_limiter_policy examples

<a id="canonical-0213100123320111-2001330133213113-0320130230122202-0312123101222021-3023213322000203-2311002230312032-2133010011033113-1330122112130323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- Examples

<a id="canonical-0323311101220120-0112311202011200-2202332201111010-1121222013031021-1301231210031311-3111003031332212-1322011322133213-2233220313101301"></a>

### Complete configurations for `xcsh_rate_limiter_policy`

- [Resource](resources--rate_limiter_policy--examples--group-001.md#canonical-2110232211133031-3030230313331213-0331300220013210-3221003011333100-0100303233222330-0013201322210221-0320321012330323-2302323021310220): valid configuration.

<a id="canonical-2110232211133031-3030230313331213-0331300220013210-3221003011333100-0100303233222330-0013201322210221-0320321012330323-2302323021310220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md#canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100)
- [Examples](resources--rate_limiter_policy--examples--group-001.md#canonical-0213100123320111-2001330133213113-0320130230122202-0312123101222021-3023213322000203-2311002230312032-2133010011033113-1330122112130323)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter_policy/resource.tf`; digest `sha256:d8a1a676828e8f463f6383201f395371ac9e0d89d99902337c84f2ee8a150c35`.

```terraform
# RateLimiterPolicy Resource Example
# Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiterPolicy configuration
resource "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}
```
