---
page_title: "xcsh_dns_zone_edit_cryptokey"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_edit_cryptokey."
---

# xcsh_dns_zone_edit_cryptokey

<a id="canonical-1132220310333100-2000123203101123-2131221001021322-3330222102211322-3331333302020203-3030331000102011-2023032220122321-2312001020231030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_dns_zone_edit_cryptokey

Edits a cryptographic key in a DNS zone.

<a id="canonical-0012001013023203-0000312100310212-0320020223011201-1210310120221230-0212022331322211-0222300231031232-3302102220221103-2333301121302031"></a>

### Prerequisites for `xcsh_dns_zone_edit_cryptokey`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1313320100312233-2000030010211211-3303320303201320-0003120302300103-1112222101022332-2011233033110212-2330203023231001-1101033223103001"></a>

### Minimal configuration for `xcsh_dns_zone_edit_cryptokey`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneEditCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_edit_cryptokey" "example" {
  config {
  }
}
```

<a id="canonical-2121111221233102-1230011033312313-1332233012320100-2230231010032012-3032321200111232-0033123312311222-2231002300322202-0133200020303023"></a>

### Root configuration for `xcsh_dns_zone_edit_cryptokey`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-0110011220201300-0322131012100213-2002203313211322-1211113022021310-3221021012213022-1002232011313132-1112221003031202-1302103011022121"></a>

### Explore this collection for `xcsh_dns_zone_edit_cryptokey`

- [Property reference](../guides/actions--dns_zone_edit_cryptokey--reference--group-001.md#canonical-2030131320303020-2330130021202001-0203001123031021-3120120210311232-1132000311113113-2200003332022233-3230212003331112-2202112202232310)
- [Examples](../guides/actions--dns_zone_edit_cryptokey--examples--group-001.md#canonical-1013213223203320-0132011103113312-1210300032133231-0200130202001120-1130123320223331-0333013021300110-0212011033121132-3230323323103232)
- [Lifecycle](../guides/actions--dns_zone_edit_cryptokey--lifecycle--group-001.md#canonical-0133311222232302-0130232320301020-2231231223333321-1130330111311102-3003323000122100-1030013331332303-0221201000313232-2222212131312203)
