---
page_title: "xcsh_bot_peer_top_good_bots examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_good_bots examples."
---

# xcsh_bot_peer_top_good_bots examples

<a id="canonical-3123030123000202-0322300120230232-0200003123301123-1221112202031331-0211032313011311-1310102100131213-3313302112000322-3130220133313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-3030001313332130-0200112302220223-2231210320202203-3301330023332321-1010113230110223-1130212301220213-3012200130111102-0020213232203302)
- Examples

<a id="canonical-1311313303231103-2023200323330223-3121200013113133-1331220113021300-2302110320131111-3302330020232020-1302303001222000-0021203302202200"></a>

### Complete configurations for `xcsh_bot_peer_top_good_bots`

- [Data source](data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-1012323120023221-1201133133303300-1201113000133113-2120332011222220-3132023213221212-2202301130031111-2300232320200221-0013111233133031): valid configuration.

<a id="canonical-1012323120023221-1201133133303300-1201113000133113-2120332011222220-3132023213221212-2202301130031111-2300232320200221-0013111233133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-3030001313332130-0200112302220223-2231210320202203-3301330023332321-1010113230110223-1130212301220213-3012200130111102-0020213232203302)
- [Examples](data-sources--bot_peer_top_good_bots--examples--group-001.md#canonical-3123030123000202-0322300120230232-0200003123301123-1221112202031331-0211032313011311-1310102100131213-3313302112000322-3130220133313121)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_top_good_bots/data-source.tf`; digest `sha256:cbd5508dce9276d9af60eef5a236b141ce290117e051aad973c941688a149920`.

```terraform
# BotPeerTopGoodBots DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_good_bots" "example" {
  namespace = "example-value"
}

output "bot_peer_top_good_bots_result" {
  value = data.xcsh_bot_peer_top_good_bots.example
}
```
