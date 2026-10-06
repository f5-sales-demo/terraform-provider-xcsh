---
page_title: "xcsh_bot_detection_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_detection_rule examples."
---

# xcsh_bot_detection_rule examples

<a id="canonical-2232113313332232-3000000211312211-0303233303022011-1332301232323032-0121030001232203-3300030203332032-0221121030302230-0100313323330002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md#canonical-2221001030010210-1213310213300133-3302112123210020-1030311021023312-0232211203112101-3232200001000312-1300200232311310-2121130011001102)
- Examples

<a id="canonical-0231301033312002-0122032212223200-3221032202220001-2101331203300011-2132022020011112-2110232230023000-3211000303332010-0100312113021331"></a>

### Complete configurations for `xcsh_bot_detection_rule`

- [Data source](data-sources--bot_detection_rule--examples--group-001.md#canonical-1111333301020220-2003002210000223-2332322201122313-1231303101202201-1202013020013331-1120313321311220-0120012013330230-3230022331003020): valid configuration.

<a id="canonical-1111333301020220-2003002210000223-2332322201122313-1231303101202201-1202013020013331-1120313321311220-0120012013330230-3230022331003020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md#canonical-2221001030010210-1213310213300133-3302112123210020-1030311021023312-0232211203112101-3232200001000312-1300200232311310-2121130011001102)
- [Examples](data-sources--bot_detection_rule--examples--group-001.md#canonical-2232113313332232-3000000211312211-0303233303022011-1332301232323032-0121030001232203-3300030203332032-0221121030302230-0100313323330002)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_detection_rule/data-source.tf`; digest `sha256:fb0cc1608165dd7e57a5b8e5f57c447871327b816ef3b596ef84d5ce798774d2`.

```terraform
# BotDetectionRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDetectionRule by name
data "xcsh_bot_detection_rule" "example" {
  name      = "example-bot-detection-rule"
  namespace = "staging"
}

output "bot_detection_rule_id" {
  value = data.xcsh_bot_detection_rule.example.id
}
```
