---
page_title: "xcsh_bot_peer_threat_types examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_threat_types examples."
---

# xcsh_bot_peer_threat_types examples

<a id="canonical-1103132310203222-1120313020330113-0011321100311123-1333222112230011-3120013033102012-3030123330103203-1122100100321000-2210030210131031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-2013010020311013-0321123032202111-3000312123121121-3301311011300221-1213110110313232-0030311231210211-3033301232310333-0010002333121001)
- Examples

<a id="canonical-2110011323123203-3102200001032300-3120032003331000-2130331320303320-2000302020003022-2130311011312030-2032333002221310-0003211221303133"></a>

### Complete configurations for `xcsh_bot_peer_threat_types`

- [Data source](data-sources--bot_peer_threat_types--examples--group-001.md#canonical-3303302102011220-1331032122300301-3302330113313021-3120312120332211-0110023003331220-2003113313201210-2011131202320023-2303030320200122): valid configuration.

<a id="canonical-3303302102011220-1331032122300301-3302330113313021-3120312120332211-0110023003331220-2003113313201210-2011131202320023-2303030320200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-2013010020311013-0321123032202111-3000312123121121-3301311011300221-1213110110313232-0030311231210211-3033301232310333-0010002333121001)
- [Examples](data-sources--bot_peer_threat_types--examples--group-001.md#canonical-1103132310203222-1120313020330113-0011321100311123-1333222112230011-3120013033102012-3030123330103203-1122100100321000-2210030210131031)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_threat_types/data-source.tf`; digest `sha256:340fefb3e7c9dd8e5383ef09986b56ee68d03ed73a4bcf9a8c9ec927ce63eafa`.

```terraform
# BotPeerThreatTypes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_threat_types" "example" {
  namespace = "example-value"
}

output "bot_peer_threat_types_result" {
  value = data.xcsh_bot_peer_threat_types.example
}
```
