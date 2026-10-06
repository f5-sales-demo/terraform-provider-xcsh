---
page_title: "xcsh_policy_based_routing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing examples."
---

# xcsh_policy_based_routing examples

<a id="canonical-2001112311223030-0112011300111031-0130103012320210-3101322331101210-0123310022312320-3301110023231332-1011123230100220-2123230230100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- Examples

<a id="canonical-0200111322021112-3302022301201203-2311222032010220-3110130022232302-0101000132301233-0213202000303013-1220233213101130-3030032023113211"></a>

### Complete configurations for `xcsh_policy_based_routing`

- [Data source](data-sources--policy_based_routing--examples--group-001.md#canonical-0322032233221030-1213230211302320-2132120322330020-2331203100000203-0332212112033233-2003312020200311-3321022120333323-2311012110221103): valid configuration.

<a id="canonical-0322032233221030-1213230211302320-2132120322330020-2331203100000203-0332212112033233-2003312020200311-3321022120333323-2311012110221103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md#canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202)
- [Examples](data-sources--policy_based_routing--examples--group-001.md#canonical-2001112311223030-0112011300111031-0130103012320210-3101322331101210-0123310022312320-3301110023231332-1011123230100220-2123230230100031)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_policy_based_routing/data-source.tf`; digest `sha256:d902f95ac0d5d170c71a1f393d0718d60b4c9a9b4856200aec4f7a2fbf5ad0fc`.

```terraform
# PolicyBasedRouting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PolicyBasedRouting by name
data "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}

output "policy_based_routing_id" {
  value = data.xcsh_policy_based_routing.example.id
}
```
