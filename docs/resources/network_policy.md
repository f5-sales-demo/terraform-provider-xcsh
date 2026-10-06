---
page_title: "xcsh_network_policy"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy."
---

# xcsh_network_policy

<a id="canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_policy

Manages new network policy with configured parameters in specified namespace in F5 Distributed
Cloud.

<a id="canonical-0303012231113232-2300321011121121-1031301302200030-2130103033200001-2230233300122000-1200113022101102-0022202331010103-3002113031002100"></a>

### Prerequisites for `xcsh_network_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1010011231211132-1220333331033211-2221010310201133-0121213133320011-3003122122330331-0333301332200103-0011013013213302-1310320302021132"></a>

### Minimal configuration for `xcsh_network_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicy Resource Example
# Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicy configuration
resource "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}
```

<a id="canonical-3210320301013333-1013133331211012-3232102332120023-1122321033000331-2023313231232020-1332123332013100-2030301230232332-2121021113112031"></a>

### Root configuration for `xcsh_network_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2330023130210003-2202333212331320-3001331000103023-0223323033001222-0100112010001323-0311201213300220-0321320031130323-3110101120323231"></a>

### Explore this collection for `xcsh_network_policy`

- [Property reference](../guides/resources--network_policy--reference--group-001.md#canonical-1122330320000323-1310100310030230-0103300322011221-0212210110212101-2310031033131221-2122113320133222-3230212213211330-2210120331000103)
- [Examples](../guides/resources--network_policy--examples--group-001.md#canonical-3231201002231112-2221112210320001-1011322112323203-3100031231131033-3123320131213212-3333312212303233-3020021213001320-2300231102202221)
- [Import](../guides/resources--network_policy--lifecycle--group-001.md#canonical-0203203102310102-3230133322230311-0211213233220222-0120312020122001-0231021211322131-3011220331203322-2332111033312010-3103313001010222)
- [Timeouts](../guides/resources--network_policy--lifecycle--group-001.md#canonical-2313312223012201-0101023111003103-2130313010211100-2122230022030130-1031202332102210-2213211310312131-3012302230323300-1230323031000110)
