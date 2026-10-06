---
page_title: "xcsh_bot_peer_top_reason_codes examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_reason_codes examples."
---

# xcsh_bot_peer_top_reason_codes examples

<a id="canonical-0320323030313213-2101030110312101-0330110202201203-0031013210210312-0333111222310230-0321233033312200-1220010221201302-0102132110333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-0013220013020012-2103133222320300-2320112130123102-0221023121031312-3010203123123211-1231322223322030-3232013030300102-1302220311210333)
- Examples

<a id="canonical-1121123223003201-2131110011013313-1200323222220332-1233000013113233-1101112110203211-1331332231322111-0231112030330110-2201312120000132"></a>

### Complete configurations for `xcsh_bot_peer_top_reason_codes`

- [Data source](data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-2230133111230322-1202023033331030-3200213012023022-3302133031002223-3110033012111233-0101223211133032-1310121002113011-2023333223200301): valid configuration.

<a id="canonical-2230133111230322-1202023033331030-3200213012023022-3302133031002223-3110033012111233-0101223211133032-1310121002113011-2023333223200301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-0013220013020012-2103133222320300-2320112130123102-0221023121031312-3010203123123211-1231322223322030-3232013030300102-1302220311210333)
- [Examples](data-sources--bot_peer_top_reason_codes--examples--group-001.md#canonical-0320323030313213-2101030110312101-0330110202201203-0031013210210312-0333111222310230-0321233033312200-1220010221201302-0102132110333220)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_top_reason_codes/data-source.tf`; digest `sha256:088c7908202451a1670b0688576af84d4f3a2cb203efd005dc0803bb8dc3d516`.

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
