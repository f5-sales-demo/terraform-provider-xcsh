---
page_title: "xcsh_dns_zone_delete_cryptokey"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_delete_cryptokey."
---

# xcsh_dns_zone_delete_cryptokey

<a id="canonical-3100330101232133-3311113231220332-0022020131232233-2130012232123321-2121220121030020-0323223001111032-0003233110133020-2000133100010203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_dns_zone_delete_cryptokey

Deletes a cryptographic key from a DNS zone.

<a id="canonical-1121103021133020-0301103032232301-0212130030013231-1013113003231333-3001223120332222-3221322233221301-1302311021123300-3232233032002223"></a>

### Prerequisites for `xcsh_dns_zone_delete_cryptokey`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3030121112010213-3121022020002101-2010021033210312-0100300211101022-0033322110012303-1212311010231120-0201213033232130-3333222301130223"></a>

### Minimal configuration for `xcsh_dns_zone_delete_cryptokey`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneDeleteCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_delete_cryptokey" "example" {
  config {
  }
}
```

<a id="canonical-1130120110031233-2010033301012001-1133310010101200-3131200213222100-1130131201132232-0132213032033101-0123202133302103-3022220021302221"></a>

### Root configuration for `xcsh_dns_zone_delete_cryptokey`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-3110010332301333-0111033222310101-1030222120211120-0312311321032102-2120322302022120-1101013031333111-2012003213113100-3212121312032230"></a>

### Explore this collection for `xcsh_dns_zone_delete_cryptokey`

- [Property reference](../guides/actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-1103032100322021-2001303320000110-2331221323002003-2030210003223133-1201102310333011-1330111231102102-1301022103211223-2321303012132222)
- [Examples](../guides/actions--dns_zone_delete_cryptokey--examples--group-001.md#canonical-3110332130013002-2332000002102303-3002003201132330-2103301202100003-2323332201302020-2330312130012103-2023310212331120-3230010323133203)
- [Lifecycle](../guides/actions--dns_zone_delete_cryptokey--lifecycle--group-001.md#canonical-3031131202220231-3030222302312300-0121122221233012-0031030013010002-1322023112202302-3020013301320100-1221200301013101-3301322221022223)
