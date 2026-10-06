---
page_title: "xcsh_bgp_routing_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy examples."
---

# xcsh_bgp_routing_policy examples

<a id="canonical-3331233210211201-3320122310001331-0023303210033213-0203331202113223-3301130030300021-1123131122130332-2312033221301123-0012312032021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- Examples

<a id="canonical-3003113132203300-2130033110101030-3320130121130223-2103321230322322-3011021122330123-2000101031121300-1123222100030332-0102123112113221"></a>

### Complete configurations for `xcsh_bgp_routing_policy`

- [Resource](resources--bgp_routing_policy--examples--group-001.md#canonical-3331212200303011-3320120231232031-1121221232110113-2110313120103300-3000112321010011-0322122111210001-2322011202313211-0103301030332021): valid configuration.

<a id="canonical-3331212200303011-3320120231232031-1121221232110113-2110313120103300-3000112321010011-0322122111210001-2322011202313211-0103301030332021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md#canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111)
- [Examples](resources--bgp_routing_policy--examples--group-001.md#canonical-3331233210211201-3320122310001331-0023303210033213-0203331202113223-3301130030300021-1123131122130332-2312033221301123-0012312032021332)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp_routing_policy/resource.tf`; digest `sha256:5114b3e2f71f278003899287dd9fce44709b024db2b4ac16ae2e23054bd84151`.

```terraform
# BGPRoutingPolicy Resource Example
# Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of rules containing match criteria and action to be applied.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPRoutingPolicy configuration
resource "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}
```
