---
page_title: "xcsh_policy_based_routing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing."
---

# xcsh_policy_based_routing

<a id="canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_policy_based_routing

Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing
create specification. configuration.

<a id="canonical-0223201012021011-0133002000123332-0022323211233110-1100010123303112-1330030003300031-1211000213101322-2110120131012201-0122231230121321"></a>

### Prerequisites for `xcsh_policy_based_routing`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3023013130032231-1130102001213213-3220110133302321-1103122203022120-2322202010313323-3110230010032332-0133331230123110-3130000200203210"></a>

### Minimal configuration for `xcsh_policy_based_routing`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2223013333202203-3000312121212212-0313200121110331-3310122331310012-3303103211011231-3323131101031303-0023233103323330-0100322133123221"></a>

### Root configuration for `xcsh_policy_based_routing`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0003132120133032-3313321120303311-2211122201211302-1223102331301233-1100221022320220-1302103223000302-1333102010220012-2311110002331133"></a>

### Explore this collection for `xcsh_policy_based_routing`

- [Property reference](../guides/resources--policy_based_routing--reference--group-001.md#canonical-1322200022123301-3232123312022112-3332200201312123-0021222100000331-1223330323113011-3131111022003132-1223003323103130-3221120203311032)
- [Examples](../guides/resources--policy_based_routing--examples--group-001.md#canonical-0223031333233231-2123100020320123-0312223203333222-3222320213031322-2300012222033000-0200110010212323-3110032213333201-3113233101001131)
- [Import](../guides/resources--policy_based_routing--lifecycle--group-001.md#canonical-1112213111010101-1222223121310031-3300022232322021-0201310110301032-2313331020112132-0132022230312022-2310213321030221-3232001300230132)
- [Timeouts](../guides/resources--policy_based_routing--lifecycle--group-001.md#canonical-1012132132330133-3220103012020021-0230113311303010-2313330311311002-1133022022132020-3023303210233231-2313133123133222-2000111001202030)
