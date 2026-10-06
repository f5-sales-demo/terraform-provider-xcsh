---
page_title: "xcsh_dns_zone_add_cryptokey"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_add_cryptokey."
---

# xcsh_dns_zone_add_cryptokey

<a id="canonical-2011333313231130-3200122021220331-0010211322333131-1302203232313332-2231111130321232-3311022020312311-2322131130200322-0032011221212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_dns_zone_add_cryptokey

Adds a cryptographic key to a DNS zone.

<a id="canonical-3122221131201302-1322033233201032-3030302323303102-2102032203311223-0033320201100132-2220003212110222-3220201132200000-3213300103303200"></a>

### Prerequisites for `xcsh_dns_zone_add_cryptokey`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0212310212130302-3321102212200220-0300303323033223-1022010112120120-1033211102131301-3202002003300321-3111113100033300-1231200221112032"></a>

### Minimal configuration for `xcsh_dns_zone_add_cryptokey`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneAddCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_add_cryptokey" "example" {
  config {
  }
}
```

<a id="canonical-1031123120130322-1311011202133312-1101303301332001-1121012003312031-0132303320332331-2220130003122032-2300312232122101-0230222022131300"></a>

### Root configuration for `xcsh_dns_zone_add_cryptokey`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-3010220330121333-1332012022202213-1311000021231001-2233331211220111-2200023122200032-2302231021123013-0300030032202212-0133133301300021"></a>

### Explore this collection for `xcsh_dns_zone_add_cryptokey`

- [Property reference](../guides/actions--dns_zone_add_cryptokey--reference--group-001.md#canonical-2202001110020203-3223112113000233-1313211203213231-2223323120211302-3330033130010113-0033001001000202-0033100323201130-3330123210330012)
- [Examples](../guides/actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-3333230033331101-1020132231121332-1311232012321312-3122211311211132-0021303011320210-3002031101120303-1302112000321320-2321202101122100)
- [Lifecycle](../guides/actions--dns_zone_add_cryptokey--lifecycle--group-001.md#canonical-3311302023212001-0023003311032131-0232113032112133-1223022122332232-1202213210122302-2120302023320332-2202212333022313-1222320323022203)
