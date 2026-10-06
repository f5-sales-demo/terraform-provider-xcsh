---
page_title: "xcsh_network_policy_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_set examples."
---

# xcsh_network_policy_set examples

<a id="canonical-1232113121123000-3300302111221013-1210121120012010-3032211003322111-3112000330100221-2332033101132211-0232032222131202-2010201203203112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_policy_set](../data-sources/network_policy_set.md#canonical-3022111020001102-2220213113331210-1332233320322003-0003223303012213-1120233210001232-0202220312332012-2312220103333300-0203313111321331)
- Examples

<a id="canonical-3201002110321112-3111113130211120-3111102123030021-3201321332011202-3200320210203220-1130220203033103-3121002112303321-3021102011220212"></a>

### Complete configurations for `xcsh_network_policy_set`

- [Data source](data-sources--network_policy_set--examples--group-001.md#canonical-1110122112330003-0202010302130112-1032221102200032-0032231131301032-2011230313113113-1132200310203222-1113103100301323-3200111212221102): valid configuration.

<a id="canonical-1110122112330003-0202010302130112-1032221102200032-0032231131301032-2011230313113113-1132200310203222-1113103100301323-3200111212221102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_policy_set](../data-sources/network_policy_set.md#canonical-3022111020001102-2220213113331210-1332233320322003-0003223303012213-1120233210001232-0202220312332012-2312220103333300-0203313111321331)
- [Examples](data-sources--network_policy_set--examples--group-001.md#canonical-1232113121123000-3300302111221013-1210121120012010-3032211003322111-3112000330100221-2332033101132211-0232032222131202-2010201203203112)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_set/data-source.tf`; digest `sha256:292e820cf31aaf1ca4db02cc9e7027c764e59f12b6fedea7ebfc71f60117c57e`.

```terraform
# NetworkPolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicySet by name
data "xcsh_network_policy_set" "example" {
  name      = "example-network-policy-set"
  namespace = "staging"
}

output "network_policy_set_id" {
  value = data.xcsh_network_policy_set.example.id
}
```
