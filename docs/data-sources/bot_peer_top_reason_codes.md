---
page_title: "xcsh_bot_peer_top_reason_codes landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_reason_codes landing."
---

# xcsh_bot_peer_top_reason_codes landing

<a id="canonical-0013220013020012-2103133222320300-2320112130123102-0221023121031312-3010203123123211-1231322223322030-3232013030300102-1302220311210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320222111320310-0100210022211323-3101020000212132-1010022001002310-3332132210113313-3200211323012230-3202313200233201-0123202231013310"></a>

## xcsh_bot_peer_top_reason_codes — xcsh_bot_peer_top_reason_codes / 322221322222 / 2

Breadcrumbs:

- xcsh_bot_peer_top_reason_codes

Resource creation operation.

<a id="canonical-2202132201313220-0103232023321300-2031200002132000-2231331230011000-3130321311220030-3222223233322012-1133203310333023-0301323200203332"></a>

## Prerequisites — xcsh_bot_peer_top_reason_codes / 322221322222 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3012110300111322-2312133322212232-2133123310021302-3010300201003321-2003310312032023-3131122220300021-0023332132211220-2112032003112220"></a>

## Minimal configuration — xcsh_bot_peer_top_reason_codes / 322221322222 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerTopReasonCodes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_reason_codes" "example" {
  namespace = "example-value"
}

output "bot_peer_top_reason_codes_result" {
  value = data.xcsh_bot_peer_top_reason_codes.example
}
```

<a id="canonical-1332020120111032-2031210200232310-0033333320310120-1232301300003330-2322310313233133-3333230300202112-2003213131311123-2311030012312213"></a>

## Root configuration — xcsh_bot_peer_top_reason_codes / 322221322222 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3113121121103030-2132013002020302-2222331001133033-3321032323122233-2233121103221301-0221021221021300-2221313331120233-0113231233200002"></a>

## Next pages — xcsh_bot_peer_top_reason_codes / 322221322222 / 6

- [Property reference](../guides/data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-0110112023203000-0031022023020133-2033031221131113-2012122332320101-0201000032302131-3222020302002200-2121313003302130-3022233222000230)
- [Examples](../guides/data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-0320323030313213-2101030110312101-0330110202201203-0031013210210312-0333111222310230-0321233033312200-1220010221201302-0102132110333220)
