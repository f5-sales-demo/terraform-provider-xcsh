---
page_title: "xcsh_network_bot_defense examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_bot_defense examples."
---

# xcsh_network_bot_defense examples

<a id="canonical-0002130000312332-1110233230003121-2103233333033223-2312121023022123-2020010323101203-1010300303231210-0003103302003303-2112021022311230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md#canonical-0033010203002120-1333301320122310-3023133010230310-1013120133033132-3201111033200121-2113002210321222-0020023233301212-2222022300321031)
- Examples

<a id="canonical-1320333022012213-2021200013322020-1000001323332023-3201311123322300-3312100322330122-1012011001022031-3002320100310131-3001010132321102"></a>

### Complete configurations for `xcsh_network_bot_defense`

- [Data source](data-sources--network_bot_defense--examples--group-001.md#canonical-2110231032332131-1321211201201120-0310201033330000-2032302212220211-3330213230110222-0011120001300010-1101323001121133-0200120122011032): valid configuration.

<a id="canonical-2110231032332131-1321211201201120-0310201033330000-2032302212220211-3330213230110222-0011120001300010-1101323001121133-0200120122011032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md#canonical-0033010203002120-1333301320122310-3023133010230310-1013120133033132-3201111033200121-2113002210321222-0020023233301212-2222022300321031)
- [Examples](data-sources--network_bot_defense--examples--group-001.md#canonical-0002130000312332-1110233230003121-2103233333033223-2312121023022123-2020010323101203-1010300303231210-0003103302003303-2112021022311230)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_bot_defense/data-source.tf`; digest `sha256:a428d9d820fed630e17a7faf15c103fb2b68704b5517319077da6fc62dccc081`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_bot_defense" "proxy" {}

# Configure these exact domains in a suffix-aware proxy or FQDN firewall.
output "bot_defense_https_proxy_rule" {
  value = {
    direction = "egress"
    protocol  = "tcp"
    port      = 443
    domains   = data.xcsh_network_bot_defense.proxy.domains
  }
}
```
