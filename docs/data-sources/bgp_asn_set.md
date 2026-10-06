---
page_title: "xcsh_bgp_asn_set"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set."
---

# xcsh_bgp_asn_set

<a id="canonical-0101132211310203-2233232230121033-0201030033302303-2112033202323033-1200310311200222-2223131321000021-3211120032230011-1102013013320333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bgp_asn_set

Reads BGP Asn Set information from F5 Distributed Cloud.

<a id="canonical-3002322300001103-0032202000101331-3002220322111020-2233112033310310-2012230210030311-2333003331133111-1002222200202101-1200303122202033"></a>

### Prerequisites for `xcsh_bgp_asn_set`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0312320130310111-3303231110132201-2021333212033310-2031311300220103-0112021303111003-0032200210323011-2302123322110121-3022022110010231"></a>

### Minimal configuration for `xcsh_bgp_asn_set`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPAsnSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPAsnSet by name
data "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"
}

output "bgp_asn_set_id" {
  value = data.xcsh_bgp_asn_set.example.id
}
```

<a id="canonical-1300232020231322-2330120201210233-3330110203001110-1103212213000123-2300222202002213-0003333102221031-3111311320232031-0313012121220122"></a>

### Root configuration for `xcsh_bgp_asn_set`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1320213132013122-0021103302000020-1311333231213031-1312302012210020-3001311212313011-0301202230211123-3022031310313330-1300311320301020"></a>

### Explore this collection for `xcsh_bgp_asn_set`

- [Property reference](../guides/data-sources--bgp_asn_set--reference--group-001.md#canonical-3313133110321332-0200122330231112-0330010031231222-0131121232003313-0202132231312021-2212233212012031-1311232213212101-1310201022103223)
- [Examples](../guides/data-sources--bgp_asn_set--examples--group-001.md#canonical-0121101003311303-3302213232300110-0000312323132103-3123212202002220-2302122013101120-3233013032232201-3111302230331100-1001302210011031)
