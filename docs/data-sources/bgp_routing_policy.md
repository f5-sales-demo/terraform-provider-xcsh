---
page_title: "xcsh_bgp_routing_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy."
---

# xcsh_bgp_routing_policy

<a id="canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bgp_routing_policy

Reads BGP routing rules with match criteria and actions that control imported or exported routes.

<a id="canonical-3101331032301321-1133133323103302-2233300222232202-1112033022223330-3332013131331023-0102132132303302-3211310023213111-2120300102130011"></a>

### Prerequisites for `xcsh_bgp_routing_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0211132013111321-2210332103302212-3032123130302003-1321302000031321-0331320133222030-2333212203203000-3323122202031222-3210030333112330"></a>

### Minimal configuration for `xcsh_bgp_routing_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPRoutingPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPRoutingPolicy by name
data "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}

output "bgp_routing_policy_id" {
  value = data.xcsh_bgp_routing_policy.example.id
}
```

<a id="canonical-0233322022300200-0123123103213321-2200002000233210-3032233331011132-0301213300021322-0012010121223133-3230203320032120-3310232001201113"></a>

### Root configuration for `xcsh_bgp_routing_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2331301130200201-0013122102301133-2122233322301330-0320021303200000-2011012230021021-2222003321202213-1010131321033332-0332200211231332"></a>

### Explore this collection for `xcsh_bgp_routing_policy`

- [Property reference](../guides/data-sources--bgp_routing_policy--reference--group-001.md#canonical-3230302232032223-2102321002200133-1031202010123210-2300211332133213-2202101301330130-1113310322000130-1210320002212323-0301110213023323)
- [Examples](../guides/data-sources--bgp_routing_policy--examples--group-001.md#canonical-0020013202220313-1110310122121231-1203301102111331-0112321221000103-2101220000303302-1202101110220103-3131301330003222-3113021201320212)
