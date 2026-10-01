---
page_title: "xcsh_mitigated_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain landing."
---

# xcsh_mitigated_domain landing

<a id="canonical-2031233030330203-1311310133233010-1010320023311012-0113331233031011-2233031102231101-3202000023330120-0122302012233333-3031333032012102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200312333022021-0003313300312332-1113010232312011-1021302213202232-3000100321211210-3102100112302221-1132302030232230-3223010131301302"></a>

## xcsh_mitigated_domain — xcsh_mitigated_domain / 121000033122 / 2

Breadcrumbs:

- xcsh_mitigated_domain

Manages Mitigated Domain in F5 Distributed Cloud.

<a id="canonical-2323301032331120-2002321001202031-1213222233031220-3211230021000010-1333032002201100-1123330130223212-3331110230221303-0033303322023210"></a>

## Prerequisites — xcsh_mitigated_domain / 121000033122 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2201301102203323-2132320113223002-0013202333300121-0022220230013103-0230210022111313-3303333032200023-0111203233322213-1331312002023102"></a>

## Minimal configuration — xcsh_mitigated_domain / 121000033122 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MitigatedDomain Resource Example
# Manages Mitigated Domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MitigatedDomain configuration
resource "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"

  mitigated_domain = "example-value"
}
```

<a id="canonical-3122021023133002-3111321103331121-3300202031113112-2210203232212232-3303313123003323-1020210002112212-3121103211333223-1022111220331133"></a>

## Root configuration — xcsh_mitigated_domain / 121000033122 / 5

Required root properties: `mitigated_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3103130013103301-3210322313230203-2103122213103000-0102033113120321-0320313001103030-2001312212302022-1020230310322013-2002231203130133"></a>

## Next pages — xcsh_mitigated_domain / 121000033122 / 6

- [Property reference](../guides/resources--mitigated_domain--reference--group-001.md#canonical-3120221331031233-2101223210101003-1020302123022020-1212033230313332-1110320212031320-1030311212231013-0201130022100233-0012321303020210)
- [Examples](../guides/resources--mitigated_domain--examples--group-001.md#canonical-3320303010302212-2103011113311030-1312202101013113-0131033321031320-2223123322221332-0333230210311320-3010301221032220-3032311011310311)
- [Import](../guides/resources--mitigated_domain--lifecycle--group-001.md#canonical-1221210302132300-2133221330033310-3123130113321223-1131112221133323-0302112133021222-1333330110222111-1102331233322233-0312331032233002)
- [Timeouts](../guides/resources--mitigated_domain--lifecycle--group-001.md#canonical-0320202330230332-2013333132221312-3323122133122121-1233122011023113-3222230321311222-1023230331121231-3021001112003013-1232012302123220)
