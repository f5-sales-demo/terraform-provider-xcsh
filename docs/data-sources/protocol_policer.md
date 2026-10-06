---
page_title: "xcsh_protocol_policer"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer."
---

# xcsh_protocol_policer

<a id="canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_protocol_policer

Reads L4 protocol match conditions and associated traffic rate limits.

<a id="canonical-3212010200020120-2121310331312023-1221220131300030-0010131100123223-1322221021303100-2301010013111213-0122311220301331-2230110112103100"></a>

### Prerequisites for `xcsh_protocol_policer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3010021311230111-2210131130100011-3121003100232212-0010022232100203-0030231202033113-2101030221102023-1220330113322102-1131303300231121"></a>

### Minimal configuration for `xcsh_protocol_policer`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
```

<a id="canonical-1103111032023121-1322320000302020-2211303210003101-0113203302033232-2011302113101210-0023310100200011-3001232232313001-3102022120202033"></a>

### Root configuration for `xcsh_protocol_policer`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1012031230120320-0333001220101221-2131121122332323-2300131031120112-3113211011002221-0033311200231311-2033300302333301-3023120323021313"></a>

### Explore this collection for `xcsh_protocol_policer`

- [Property reference](../guides/data-sources--protocol_policer--reference--group-001.md#canonical-0130003321321013-2123100333101302-3322212301112022-3021222300202002-2202101011132112-1302030002103213-1112121312103233-3330032222102101)
- [Examples](../guides/data-sources--protocol_policer--examples--group-001.md#canonical-1213220033330202-0030031031311012-3330210230023311-0231313332003100-3010300113001220-3320131123012020-0310113102113111-0012013123232223)
