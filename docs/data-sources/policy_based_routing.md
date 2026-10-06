---
page_title: "xcsh_policy_based_routing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing."
---

# xcsh_policy_based_routing

<a id="canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_policy_based_routing

Reads Policy Based Routing information from F5 Distributed Cloud.

<a id="canonical-1230230003321123-0200113022220213-3210002130001211-1013030322133213-2011033031310310-0220210123032021-0111332232313320-0011020201232001"></a>

### Prerequisites for `xcsh_policy_based_routing`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0110311302102200-0131030212010312-0100300031310020-3313221112203121-0233221032333122-2222130013320232-3023133020221101-1220031220120301"></a>

### Minimal configuration for `xcsh_policy_based_routing`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2112200310003022-1200003310101212-1301203103302222-2213123103212120-0031230321320230-2100133001032111-2032202011110130-1331002003203003"></a>

### Root configuration for `xcsh_policy_based_routing`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2011320111011133-0103102111233202-0323300311302103-0133132231020202-2013320212003321-2302310300031233-3130123100022313-0220120111102121"></a>

### Explore this collection for `xcsh_policy_based_routing`

- [Property reference](../guides/data-sources--policy_based_routing--reference--group-001.md#canonical-3020220020102013-1020212212113210-2110032221121310-1233011300323132-3123011100001132-3112102103121120-3020133332302001-0200300102322131)
- [Examples](../guides/data-sources--policy_based_routing--examples--group-001.md#canonical-2001112311223030-0112011300111031-0130103012320210-3101322331101210-0123310022312320-3301110023231332-1011123230100220-2123230230100031)
