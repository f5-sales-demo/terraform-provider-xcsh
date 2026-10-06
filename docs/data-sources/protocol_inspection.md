---
page_title: "xcsh_protocol_inspection"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection."
---

# xcsh_protocol_inspection

<a id="canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_protocol_inspection

Reads an existing Protocol Inspection configuration in its namespace.

<a id="canonical-0203313102131132-3212320332220110-0201223213000012-1120132222303012-3003122100103031-1212300112332021-2233130231322213-3021323301010231"></a>

### Prerequisites for `xcsh_protocol_inspection`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1221022232222022-1002003131222313-0031033332113300-2332003231032332-1012332201233131-2230012003223332-3233320321033003-0010222331310130"></a>

### Minimal configuration for `xcsh_protocol_inspection`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolInspection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolInspection by name
data "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}

output "protocol_inspection_id" {
  value = data.xcsh_protocol_inspection.example.id
}
```

<a id="canonical-3302212001230033-1300321333313303-2120200322020331-3320201030012203-1301210331032122-0130022300103021-2220320201112001-3330231120002320"></a>

### Root configuration for `xcsh_protocol_inspection`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0102212111323120-3101331020020123-0221331131332201-3203013303111103-2333232233131211-1210203022111001-2103313022010310-3322321113012223"></a>

### Explore this collection for `xcsh_protocol_inspection`

- [Property reference](../guides/data-sources--protocol_inspection--reference--group-001.md#canonical-2031133223011203-2331031213330110-1111011320132233-1130032201323122-3030112030303121-2231033112022021-0213331300320023-3130013121123023)
- [Examples](../guides/data-sources--protocol_inspection--examples--group-001.md#canonical-2222200023121001-2120130320002212-3002302310233020-3021231113212123-3120033122221302-2013033122021310-1313023233113113-1023111323332123)
