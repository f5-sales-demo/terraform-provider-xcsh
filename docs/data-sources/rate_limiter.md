---
page_title: "xcsh_rate_limiter landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter landing."
---

# xcsh_rate_limiter landing

<a id="canonical-0000133313221131-0223011310120301-2223003312312002-1323011333202102-3322002222213120-2323100011202011-1122313333010000-3320032012132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333011110230213-0211131222233001-0212222000213232-0311020030213021-0203233121031222-2323212120020303-0232233020113330-2000022013320023"></a>

## xcsh_rate_limiter — xcsh_rate_limiter / 112100123123 / 2

Breadcrumbs:

- xcsh_rate_limiter

Manages rate\_limiter creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-3321202023022330-0233302300112201-2001330023200120-1121011031032310-2011330133233202-0100232131110021-0101012313231130-0321302100330300"></a>

## Prerequisites — xcsh_rate_limiter / 112100123123 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `rate_limiter_policy`.

- rate_limiter_policy: Detailed rate limiting rules

<a id="canonical-2321123220330031-2321233023012323-3233212033031321-0213301331321100-0111021112000033-1301230220233130-0202302312211031-3313122110332000"></a>

## Minimal configuration — xcsh_rate_limiter / 112100123123 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3320132113202123-0213311321310213-2110010332303221-3032001310212123-3112021002311323-2122113011101100-2030210022230133-1123233013032101"></a>

## Root configuration — xcsh_rate_limiter / 112100123123 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3131011101022001-0222131200120321-2113221222200120-3010102210032130-1310103003113223-1032323201231221-1301012111130231-3102032210013303"></a>

## Next pages — xcsh_rate_limiter / 112100123123 / 6

- [Property reference](../guides/data-sources--rate_limiter--reference--group-001.md#canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223)
- [Examples](../guides/data-sources--rate_limiter--examples--group-001.md#canonical-3102222211011003-3321132003022122-0230231131030120-1213200131331313-3100321221330010-0330022033211222-3201000211000203-1131232131111232)
