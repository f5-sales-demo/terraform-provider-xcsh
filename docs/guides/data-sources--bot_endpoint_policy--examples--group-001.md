---
page_title: "xcsh_bot_endpoint_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_endpoint_policy examples."
---

# xcsh_bot_endpoint_policy examples

<a id="canonical-3303333021232001-0201323230023222-2012023013012233-3023322031312202-3310031131032123-0323231120222213-0220003001000000-0232003211310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- Examples

<a id="canonical-1132222011132301-3113120012300230-2003000003323123-0233220013021100-1011331223333003-3111023110110200-3121130102000031-3203132011102010"></a>

### Complete configurations for `xcsh_bot_endpoint_policy`

- [Data source](data-sources--bot_endpoint_policy--examples--group-001.md#canonical-2013332022003012-1123320013322303-2011313130110113-3110223331110213-0223300030010201-0100322130311131-1321211030302021-0232303101301132): valid configuration.

<a id="canonical-2013332022003012-1123320013322303-2011313130110113-3110223331110213-0223300030010201-0100322130311131-1321211030302021-0232303101301132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Examples](data-sources--bot_endpoint_policy--examples--group-001.md#canonical-3303333021232001-0201323230023222-2012023013012233-3023322031312202-3310031131032123-0323231120222213-0220003001000000-0232003211310310)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_endpoint_policy/data-source.tf`; digest `sha256:742dea18fab3d10fddee899cf0137a64f4bc6d260ffbdc975152403793a768eb`.

```terraform
# BotEndpointPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotEndpointPolicy by name
data "xcsh_bot_endpoint_policy" "example" {
  name      = "example-bot-endpoint-policy"
  namespace = "staging"
}

output "bot_endpoint_policy_id" {
  value = data.xcsh_bot_endpoint_policy.example.id
}
```
