---
page_title: "xcsh_bot_peer_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_status examples."
---

# xcsh_bot_peer_status examples

<a id="canonical-1013303221323110-1301112223102022-3233003023303032-3303121320330131-1103031223202233-0201233102030332-0223000332130021-1011032031033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md#canonical-1213323001223033-2102310131020031-2130002023320333-3233002313233120-0102113233101212-0211233322133022-3213211320322133-3330331011313210)
- Examples

<a id="canonical-1332213122330100-3232201101002220-3213122031121111-2003111001102033-3312321003000333-0032121213202220-2003210120110212-2030211013102023"></a>

### Complete configurations for `xcsh_bot_peer_status`

- [Data source](data-sources--bot_peer_status--examples--group-001.md#canonical-0221003332320320-1313110133232103-2211003222031301-0331101231032103-3030101102000201-0300120000321230-2311001010331233-1221013233132003): valid configuration.

<a id="canonical-0221003332320320-1313110133232103-2211003222031301-0331101231032103-3030101102000201-0300120000321230-2311001010331233-1221013233132003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_peer_status](../data-sources/bot_peer_status.md#canonical-1213323001223033-2102310131020031-2130002023320333-3233002313233120-0102113233101212-0211233322133022-3213211320322133-3330331011313210)
- [Examples](data-sources--bot_peer_status--examples--group-001.md#canonical-1013303221323110-1301112223102022-3233003023303032-3303121320330131-1103031223202233-0201233102030332-0223000332130021-1011032031033213)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_status/data-source.tf`; digest `sha256:97b1bb9bbcf7e35e2e9068fc82fcedf34184ffea0125e8cf3c1304beb4255fcc`.

```terraform
# BotPeerStatus DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_status" "example" {
  namespace = "example-value"
}

output "bot_peer_status_result" {
  value = data.xcsh_bot_peer_status.example
}
```
