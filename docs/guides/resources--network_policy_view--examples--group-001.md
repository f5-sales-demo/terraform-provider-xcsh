---
page_title: "xcsh_network_policy_view examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view examples."
---

# xcsh_network_policy_view examples

<a id="canonical-0111310100333132-3203131111203113-0112030133332200-2221321020230210-1230101212322333-3100103111122002-2100330131112123-3230201322132133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- Examples

<a id="canonical-0310133003321303-2203023301230213-3010332002122232-2001033120110133-3203033002012330-0233211302003103-1322111222123201-1000023012003312"></a>

### Complete configurations for `xcsh_network_policy_view`

- [Resource](resources--network_policy_view--examples--group-001.md#canonical-3221010333023301-0110103133312212-2033330230320132-2233213323033311-0012103003303123-3212202312020311-3112130203033032-3102033103110332): valid configuration.

<a id="canonical-3221010333023301-0110103133312212-2033330230320132-2233213323033311-0012103003303123-3212202312020311-3112130203033032-3102033103110332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Examples](resources--network_policy_view--examples--group-001.md#canonical-0111310100333132-3203131111203113-0112030133332200-2221321020230210-1230101212322333-3100103111122002-2100330131112123-3230201322132133)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy_view/resource.tf`; digest `sha256:75bf659169fe571161ebe1df55c559339c5fc7697dcbe1c2ae2a1986ba0ba320`.

```terraform
# NetworkPolicyView Resource Example
# Manages a Network Policy View resource in F5 Distributed Cloud for network policy view specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyView configuration
resource "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}
```
