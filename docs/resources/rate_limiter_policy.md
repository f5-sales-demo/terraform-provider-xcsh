---
page_title: "xcsh_rate_limiter_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy landing."
---

# xcsh_rate_limiter_policy landing

<a id="canonical-0001113203213012-3230223203213013-3101130323201300-3231310112330322-1330013011322232-3231222000231002-0131233030332120-1030230203122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012123321330010-2223100133313321-1022131301210203-3003112300213300-1103300012212123-1131313202302231-0300220121020313-2131112123222331"></a>

## xcsh_rate_limiter_policy — xcsh_rate_limiter_policy / 110011321123 / 2

Breadcrumbs:

- xcsh_rate_limiter_policy

Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create
specification. configuration.

<a id="canonical-2001223001221103-1223200202333011-3130132203211311-2223001110112223-0103312320023110-3003001330121100-1333200221112133-1331330112013022"></a>

## Prerequisites — xcsh_rate_limiter_policy / 110011321123 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-2313312230130231-3331130231020002-3033221122220113-0112101302301201-1331113113003121-0231022033130102-3212200230011113-1032003303121313"></a>

## Minimal configuration — xcsh_rate_limiter_policy / 110011321123 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2321231111301020-1103231321123020-1303333220030203-2013130110010110-2223110231131110-1303202101013111-0213303223220020-3301013300120020"></a>

## Root configuration — xcsh_rate_limiter_policy / 110011321123 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2032332333222121-3101102032010211-0120020013110331-2213023212213103-0012301312232012-0313202211022031-2111201211031312-2112301020103313"></a>

## Next pages — xcsh_rate_limiter_policy / 110011321123 / 6

- [Property reference](../guides/resources--rate_limiter_policy--reference--group-001.md#canonical-0300002123113003-0232011230013030-1022121023121332-1310323330300210-0013133131103100-1131102133320333-3103000313033112-0112302033121023)
- [Examples](../guides/resources--rate_limiter_policy--examples--group-001.md#canonical-0213100123320111-2001330133213113-0320130230122202-0312123101222021-3023213322000203-2311002230312032-2133010011033113-1330122112130323)
- [Import](../guides/resources--rate_limiter_policy--lifecycle--group-001.md#canonical-3022030322002201-2133333030111232-3212220102002312-1010010013032310-0113231221123101-3111212122311022-2313312112033301-0223231231100212)
- [Timeouts](../guides/resources--rate_limiter_policy--lifecycle--group-001.md#canonical-3010223313023231-1320313212123203-0211313221212323-3000001311201322-2033012030330202-3300232122213133-1123113213321033-3022020312323302)
