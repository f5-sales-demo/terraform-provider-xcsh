---
page_title: "xcsh_network_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule examples."
---

# xcsh_network_policy_rule examples

<a id="canonical-1310323332320123-3101333012201321-2011122013121110-0013233222122212-3331130302323301-1133001222311310-0233221203211132-0022002133333000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- Examples

<a id="canonical-0133133302210123-3102220313322032-2333230113222103-1301322021331233-1321310030110223-2231312333030012-0023303303111232-3133032330130232"></a>

### Complete configurations for `xcsh_network_policy_rule`

- [Data source](data-sources--network_policy_rule--examples--group-001.md#canonical-1331321020003333-3322022322032012-3220312212110232-3221333312323020-0321220020211031-0311033031013131-0022220031323013-2020330210222203): valid configuration.

<a id="canonical-1331321020003333-3322022322032012-3220312212110232-3221333312323020-0321220020211031-0311033031013131-0022220031323013-2020330210222203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103)
- [Examples](data-sources--network_policy_rule--examples--group-001.md#canonical-1310323332320123-3101333012201321-2011122013121110-0013233222122212-3331130302323301-1133001222311310-0233221203211132-0022002133333000)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_rule/data-source.tf`; digest `sha256:e28844bbbc2540f4dac23720e7a2c315eba810ea0e139291452a348547d903cb`.

```terraform
# NetworkPolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyRule by name
data "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}

output "network_policy_rule_id" {
  value = data.xcsh_network_policy_rule.example.id
}
```
