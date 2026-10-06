---
page_title: "xcsh_policy_based_routing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing examples."
---

# xcsh_policy_based_routing examples

<a id="canonical-0223031333233231-2123100020320123-0312223203333222-3222320213031322-2300012222033000-0200110010212323-3110032213333201-3113233101001131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- Examples

<a id="canonical-0221333202113101-0012211210010101-0333000330013233-0213132221100220-3122300012210103-2313232130030113-3223022220122121-2100210012303231"></a>

### Complete configurations for `xcsh_policy_based_routing`

- [Resource](resources--policy_based_routing--examples--group-001.md#canonical-3310100020110201-2313301030311313-2123101220300301-3311301012232032-0302121032101013-3003300102331202-1220111310000022-2200130313332003): valid configuration.

<a id="canonical-3310100020110201-2313301030311313-2123101220300301-3311301012232032-0302121032101013-3003300102331202-1220111310000022-2200130313332003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102)
- [Examples](resources--policy_based_routing--examples--group-001.md#canonical-0223031333233231-2123100020320123-0312223203333222-3222320213031322-2300012222033000-0200110010212323-3110032213333201-3113233101001131)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policy_based_routing/resource.tf`; digest `sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968`.

```terraform
# PolicyBasedRouting Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PolicyBasedRouting configuration
resource "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}
```
