---
page_title: "xcsh_rate_limiter"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter."
---

# xcsh_rate_limiter

<a id="canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_rate_limiter

Manages rate\_limiter creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-0221001201113131-1233300031030001-1301103111310301-2310130221203320-1301002000031201-1003013132321102-2000123013210223-2212000222121302"></a>

### Prerequisites for `xcsh_rate_limiter`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `rate_limiter_policy`.

- rate_limiter_policy: Detailed rate limiting rules

<a id="canonical-0100320201313333-1300321020200020-3313002031312321-0322121022030001-3011000131310113-3012310130002120-2212320332001202-0230212121103013"></a>

### Minimal configuration for `xcsh_rate_limiter`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RateLimiter Resource Example
# Manages rate_limiter creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiter configuration
resource "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}
```

<a id="canonical-3213103223120231-3213210331023303-2302311003320302-3221000030210222-3322232311311122-0300311103203212-3130220133213121-3321012321031120"></a>

### Root configuration for `xcsh_rate_limiter`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0101333100330331-1111320101110233-3223311101101020-2000002303120120-1030000132000303-1202332222313100-3030220013201020-3223310110333111"></a>

### Explore this collection for `xcsh_rate_limiter`

- [Property reference](../guides/resources--rate_limiter--reference--group-001.md#canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022)
- [Examples](../guides/resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
- [Import](../guides/resources--rate_limiter--lifecycle--group-001.md#canonical-3213130203103230-3120313131201013-2223130230321301-3222333033311110-3130233311223020-0102202013222001-1320113201303112-2101210023221223)
- [Timeouts](../guides/resources--rate_limiter--lifecycle--group-001.md#canonical-3303320213232022-2201033002022230-3222231123120113-2202131313003122-1311302303211311-1223323333020122-3310003033020022-2230000001222033)
