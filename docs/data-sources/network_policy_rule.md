---
page_title: "xcsh_network_policy_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule landing."
---

# xcsh_network_policy_rule landing

<a id="canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210211300233300-2302221102233100-2302132000201013-3310303102200011-3213012233222103-2111130102203133-0002101132010200-0101322300102203"></a>

## xcsh_network_policy_rule — xcsh_network_policy_rule / 200232323130 / 2

Breadcrumbs:

- xcsh_network_policy_rule

Manages network policy rule with configured parameters in specified namespace in F5 Distributed
Cloud.

<a id="canonical-1230122322330103-2021302303330023-3321302121102001-3103101023220321-0010112303201023-3320030233012102-2011130213121232-2120312203010313"></a>

## Prerequisites — xcsh_network_policy_rule / 200232323130 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0301323011102330-1202331011130013-3123323011221232-3003110330320012-3021201023020233-3320023311003030-1320020002001100-0330321122132003"></a>

## Minimal configuration — xcsh_network_policy_rule / 200232323130 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1320031031313200-3003133203303132-3101303102100220-3130211110132303-2211202002110222-3002330103012122-1100230033022112-3201330102101302"></a>

## Root configuration — xcsh_network_policy_rule / 200232323130 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3130210130233202-2330122330332232-0100311110300003-0211110003323110-0323233113222033-3031303100203132-3013331103220311-2311120111221313"></a>

## Next pages — xcsh_network_policy_rule / 200232323130 / 6

- [Property reference](../guides/data-sources--network_policy_rule--reference--group-001.md#canonical-0111202221110032-0130300212200201-2331112010002030-3130213113210211-3232312202131201-0020002303223310-0331212101231323-1120301311233300)
- [Examples](../guides/data-sources--network_policy_rule--examples--group-001.md#canonical-1310323332320123-3101333012201321-2011122013121110-0013233222122212-3331130302323301-1133001222311310-0233221203211132-0022002133333000)
