---
page_title: "xcsh_bot_network_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_network_policy examples."
---

# xcsh_bot_network_policy examples

<a id="canonical-0002022322030223-0312030020301110-0013302131302112-1230333220320331-3031320232323312-3121201120303031-1030210030322301-1330332213001121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md#canonical-0131231022333300-2321113231232321-2013223130013312-1311321020233021-2120000102000123-3300101133122113-2333131010121013-3020322100131202)
- Examples

<a id="canonical-2323103321012122-2101122011213000-2110210331033010-2002213332221232-2323111323302320-1310210210222011-0002323103211332-2132132211110011"></a>

### Complete configurations for `xcsh_bot_network_policy`

- [Data source](data-sources--bot_network_policy--examples--group-001.md#canonical-1031103302223011-0321323231101300-3331232122030330-1023321010213131-1020033232020221-0312003313133200-0012100321310211-1220333212133131): valid configuration.

<a id="canonical-1031103302223011-0321323231101300-3331232122030330-1023321010213131-1020033232020221-0312003313133200-0012100321310211-1220333212133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_network_policy](../data-sources/bot_network_policy.md#canonical-0131231022333300-2321113231232321-2013223130013312-1311321020233021-2120000102000123-3300101133122113-2333131010121013-3020322100131202)
- [Examples](data-sources--bot_network_policy--examples--group-001.md#canonical-0002022322030223-0312030020301110-0013302131302112-1230333220320331-3031320232323312-3121201120303031-1030210030322301-1330332213001121)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_network_policy/data-source.tf`; digest `sha256:b2aec065924f60efb4e75a54937e4cdc96ee0d4e38bbfe5651736f1aa7a5ffe9`.

```terraform
# BotNetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotNetworkPolicy by name
data "xcsh_bot_network_policy" "example" {
  name      = "example-bot-network-policy"
  namespace = "staging"
}

output "bot_network_policy_id" {
  value = data.xcsh_bot_network_policy.example.id
}
```
