---
page_title: "xcsh_fleet"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet."
---

# xcsh_fleet

<a id="canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_fleet

Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

<a id="canonical-2310011312202013-1311023212113311-0310332000321232-0020322100003211-3301221013210031-0331202023211311-1311131321022220-0221230301111013"></a>

### Prerequisites for `xcsh_fleet`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2203102133222032-1120022312130030-1022212020233131-1023331131020110-0300103321331233-1303212002210000-3112211302033323-3323233032122323"></a>

### Minimal configuration for `xcsh_fleet`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Fleet Resource Example
# Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fleet configuration
resource "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"

  fleet_label = "example-value"
}
```

<a id="canonical-2101111123032202-0233231303122000-1120223023100312-2330332010133132-1211303003220100-2103332010333122-0221223033202211-1112002112110032"></a>

### Root configuration for `xcsh_fleet`

Required root properties: `fleet_label`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2032201221021113-3311010321320320-3131011123022311-0322133212120213-1323122131232200-0300331212102031-0301011210202321-3323032232330123"></a>

### Explore this collection for `xcsh_fleet`

- [Property reference](../guides/resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [Examples](../guides/resources--fleet--examples--group-001.md#canonical-3300202033332320-3213223220133221-1232320113200122-3233113313222131-1230323102323110-0100312311110323-1122133332121301-1012301011332211)
- [Import](../guides/resources--fleet--lifecycle--group-001.md#canonical-3221301112320030-3323110332312301-0001311302000302-1300222133101221-2120323300022212-1210133323011120-1330033031131301-3222311203330132)
- [Timeouts](../guides/resources--fleet--lifecycle--group-001.md#canonical-1101022322322211-1011233300033232-0021133201323030-2100313023111233-0110110310121113-3313220221000030-3322020032333301-3110031211322011)
