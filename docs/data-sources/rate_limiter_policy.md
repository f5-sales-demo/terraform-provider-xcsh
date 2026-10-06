---
page_title: "xcsh_rate_limiter_policy"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy."
---

# xcsh_rate_limiter_policy

<a id="canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_rate_limiter_policy

Reads Rate Limiter Policy information from F5 Distributed Cloud.

<a id="canonical-0300202203103223-0313020231113310-3110203100003101-2111033331121023-1332201200330230-1333103322323203-3130030002210011-1020233032030131"></a>

### Prerequisites for `xcsh_rate_limiter_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3132013011031321-0310010011111302-3232202101211011-1203133000111222-0320122111020130-3223303213001222-3101300033311001-0132231301321000"></a>

### Minimal configuration for `xcsh_rate_limiter_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1221331223121100-2210010302122303-0322301133322302-2330031232100321-1010023120220233-3330023033330311-0312300033133310-0012210002130320"></a>

### Root configuration for `xcsh_rate_limiter_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2203121021220113-0312102333100123-3120013010030003-3122010002212020-1000222221012030-2110110210013031-0012322303023100-3000300002231201"></a>

### Explore this collection for `xcsh_rate_limiter_policy`

- [Property reference](../guides/data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [Examples](../guides/data-sources--rate_limiter_policy--examples--group-001.md#canonical-2122331121300203-0101013233100031-3100200332130301-0010100133112022-0230133200320122-3111332320321112-0213002131311032-2213330231232133)
