---
page_title: "xcsh_protocol_policer"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer."
---

# xcsh_protocol_policer

<a id="canonical-3330333012213023-3331002310331112-1210012223120003-2212223210101200-3103330300201100-0230232202320033-0111223122201102-2133311232122222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_protocol_policer

Manages protocol\_policer object, protocol\_policer object contains list of L4 protocol match
condition and corresponding traffic rate limits in F5 Distributed Cloud.

<a id="canonical-3011002312232231-3111211001010102-0200300213221201-1133203123013211-0212302303102322-3123203020301013-2010112123031132-0111213002200110"></a>

### Prerequisites for `xcsh_protocol_policer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2301103302120000-2130011211313133-1103323232200323-1202021211313020-2030010231223212-0133113120020212-2033003311310010-0320221132030210"></a>

### Minimal configuration for `xcsh_protocol_policer`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolPolicer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolPolicer configuration
resource "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}
```

<a id="canonical-3223032203312320-1032303220022021-2033330332201213-1002010323330102-0031332123121100-1100022232111003-2000301301122220-1222330232221123"></a>

### Root configuration for `xcsh_protocol_policer`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0111312201211230-1221303133212213-0123230001201311-0113123231113020-0210110333120113-3202331113002320-0211322120132013-3213111101002301"></a>

### Explore this collection for `xcsh_protocol_policer`

- [Property reference](../guides/resources--protocol_policer--reference--group-001.md#canonical-2131200202001010-0000130302120213-0011120213023122-1123322102232000-1023213310323311-0232321203023223-1310032211102313-0221101102322221)
- [Examples](../guides/resources--protocol_policer--examples--group-001.md#canonical-3120323230231032-0233303201201322-3001301212102012-1331313222310222-2110131022000020-0021030121331322-0212120110332220-1100023333311320)
- [Import](../guides/resources--protocol_policer--lifecycle--group-001.md#canonical-3221123312303112-3311203013001203-0033320131200123-3200112203231013-3000220022021011-2310130301103122-1002320123130221-3331020133303130)
- [Timeouts](../guides/resources--protocol_policer--lifecycle--group-001.md#canonical-0031311000210302-3210212120013022-1120330010223111-0311320231312013-3101110213011101-2232320221232323-2100103033032213-3222033322212201)
