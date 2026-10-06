---
page_title: "xcsh_protocol_inspection"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection."
---

# xcsh_protocol_inspection

<a id="canonical-3030212022213332-3331201033021331-2130011232020201-3310322223233130-3212232100333301-3201232230312101-1003032323133112-1133013223230102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_protocol_inspection

Manages Protocol Inspection Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

<a id="canonical-2232311013102122-3210101113030221-3201131003203030-0000331222030031-1313103112313103-1111322020313103-0011323121020002-0011330313313133"></a>

### Prerequisites for `xcsh_protocol_inspection`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1022031231302010-2210200133133320-3212122230303211-1020311333031333-0113020130013021-0131211003333000-0133113002232203-0223232121312113"></a>

### Minimal configuration for `xcsh_protocol_inspection`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolInspection Resource Example
# Manages Protocol Inspection Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolInspection configuration
resource "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}
```

<a id="canonical-3011023302330132-3113022220200103-2220210203322323-3033123101121220-2013312001210221-0130100013202302-3111021310310222-3002220101101311"></a>

### Root configuration for `xcsh_protocol_inspection`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1202213003223200-0033320010111022-1013000022103310-1222012020302101-1312132233020201-1122100032133032-1330312223311002-2133202001021300"></a>

### Explore this collection for `xcsh_protocol_inspection`

- [Property reference](../guides/resources--protocol_inspection--reference--group-001.md#canonical-3210031210322130-3133210121332000-1300320030213333-1220303203113101-3030330132202312-1103302221030102-1202010201212331-3030222101313220)
- [Examples](../guides/resources--protocol_inspection--examples--group-001.md#canonical-3122220013221013-1023121113312331-1202030321022111-3032013321232312-3210321300200303-3121131203021231-1020101200200202-3102211232331033)
- [Import](../guides/resources--protocol_inspection--lifecycle--group-001.md#canonical-1021321310223030-0232333021020310-1333200330213323-3322220200023113-2130212001333221-3302001022012101-2332331022300321-3010003212102123)
- [Timeouts](../guides/resources--protocol_inspection--lifecycle--group-001.md#canonical-0330301323110300-1031203122033132-0220031123123302-2000033332203031-0131030201233012-0032132331020213-2132132100003332-3212112131311220)
