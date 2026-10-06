---
page_title: "xcsh_bgp_routing_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy."
---

# xcsh_bgp_routing_policy

<a id="canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bgp_routing_policy

Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of
rules containing match criteria and action to be applied. these rules help control routes which are
imported or exported to bgp peers. configuration.

<a id="canonical-0010233123020113-2112013102000233-0311110323210031-3332301232131233-3300231032321121-2323313303133332-1321310222331323-0233332303211312"></a>

### Prerequisites for `xcsh_bgp_routing_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1133122312133012-2012123313233300-1320122031030331-3203022112013330-1032320132221320-1320123120232310-3120003010102232-0220311132131303"></a>

### Minimal configuration for `xcsh_bgp_routing_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3300223022220133-0001010233301303-1301231022220212-1132122023030212-3231201201033323-0120123312033220-1002331112120021-2321031113013133"></a>

### Root configuration for `xcsh_bgp_routing_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2333210203300000-0302013112003210-3300213321312031-2210200103321220-3112023101033023-2311120200212021-0111102302231022-0223120300331223"></a>

### Explore this collection for `xcsh_bgp_routing_policy`

- [Property reference](../guides/resources--bgp_routing_policy--reference--group-001.md#canonical-3013023032130133-2230102223310123-1322330103333212-1200321223212002-1322231010210000-3322131101112111-3313131023323000-0210322222303302)
- [Examples](../guides/resources--bgp_routing_policy--examples--group-001.md#canonical-3331233210211201-3320122310001331-0023303210033213-0203331202113223-3301130030300021-1123131122130332-2312033221301123-0012312032021332)
- [Import](../guides/resources--bgp_routing_policy--lifecycle--group-001.md#canonical-2302332132300333-3331212230301322-0010133130121333-3300022001310330-0013021003132133-1012303231133310-2013223122322102-2201023220333122)
- [Timeouts](../guides/resources--bgp_routing_policy--lifecycle--group-001.md#canonical-3011210220032132-0311203232030113-2232331101211323-3021302230332101-3023122011320231-0301202132130313-2013212223101003-1212210112231311)
